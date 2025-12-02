// Copyright © 2025 Scott Wiederhold <s.e.wiederhold@gmail.com>
// All rights reserved.
//
// Licensed under the Mozilla Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	https://mozilla.org/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: MPL-2.0
package client

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/go-resty/resty/v2"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
)

// DatapowerClientTaretConfig describes the target model.
type DatapowerClientTargetConfig struct {
	Hostname types.String `tfsdk:"hostname"`
	Port     types.Int32  `tfsdk:"port"`
	Username types.String `tfsdk:"username"`
	Password types.String `tfsdk:"password"`
	Insecure types.Bool   `tfsdk:"insecure"`
}

// DatapowerClientConfig describes the provider data model.
type DatapowerClientConfig struct {
	Hostname types.String                            `tfsdk:"hostname"`
	Port     types.Int32                             `tfsdk:"port"`
	Username types.String                            `tfsdk:"username"`
	Password types.String                            `tfsdk:"password"`
	Insecure types.Bool                              `tfsdk:"insecure"`
	Targets  map[string]*DatapowerClientTargetConfig `tfsdk:"targets"`
}

// DatapowerClient describes the data maintained by the provider.
type DatapowerClient struct {
	config *DatapowerClientConfig
	client map[string]*resty.Client
}

func NewClient(config *DatapowerClientConfig) (*DatapowerClient, error) {
	var baseUsername types.String
	if (config.Username.IsUnknown() || config.Username.ValueString() == "") && os.Getenv("DP_USERNAME") == "" {
		return nil, fmt.Errorf("client: Cannot use unknown/empty value as username")
	}
	if !config.Username.IsNull() && config.Username.ValueString() != "" {
		baseUsername = config.Username
	} else {
		baseUsername = types.StringValue(os.Getenv("DP_USERNAME"))
	}

	var basePassword types.String
	if (config.Password.IsUnknown() || config.Password.ValueString() == "") && os.Getenv("DP_PASSWORD") == "" {
		return nil, fmt.Errorf("client: Cannot use unknown/empty value as password")
	}
	if !config.Password.IsNull() && config.Password.ValueString() != "" {
		basePassword = config.Password
	} else {
		basePassword = types.StringValue(os.Getenv("DP_PASSWORD"))
	}

	// Set Base URL
	var baseHostname types.String
	if !config.Hostname.IsUnknown() && config.Hostname.ValueString() != "" {
		baseHostname = config.Hostname
	} else if env := os.Getenv("DP_HOSTNAME"); env != "" {
		baseHostname = types.StringValue(env)
	} else {
		baseHostname = types.StringValue("localhost")
	}

	var basePort types.Int32
	if !config.Port.IsUnknown() && config.Port.ValueInt32() > 0 {
		basePort = config.Port
	} else if env := os.Getenv("DP_PORT"); env != "" {
		p, err := strconv.Atoi(env)
		if err != nil {
			return nil, fmt.Errorf("client: Invalid port value %s", env)
		}
		basePort = types.Int32Value(int32(p))
	} else {
		basePort = types.Int32Value(5554)
	}

	if basePort.ValueInt32() < 1 || basePort.ValueInt32() > 65535 {
		return nil, fmt.Errorf("invalid port: %d (must be 1-65535)", basePort.ValueInt32())
	}

	baseInsecure := config.Insecure
	if env := os.Getenv("DP_INSECURE"); env != "" {
		b, err := strconv.ParseBool(env)
		if err != nil {
			return nil, fmt.Errorf("invalid DP_INSECURE value: %s", env)
		}
		baseInsecure = types.BoolValue(baseInsecure.ValueBool() || b)
	}

	var restyClients = make(map[string]*resty.Client)

	restyClients["_base_"] = newClient(
		DatapowerClientTargetConfig{
			Hostname: baseHostname,
			Port:     basePort,
			Username: baseUsername,
			Password: basePassword,
			Insecure: baseInsecure,
		},
	)

	for targetName, target := range config.Targets {
		restyClients[targetName] = newClient(
			DatapowerClientTargetConfig{
				Hostname: stringOr(target.Hostname, baseHostname),
				Port:     int32Or(target.Port, basePort),
				Username: stringOr(target.Username, baseUsername),
				Password: stringOr(target.Password, basePassword),
				Insecure: boolOr(target.Insecure, baseInsecure),
			},
		)
	}

	return &DatapowerClient{
		config: config,
		client: restyClients,
	}, nil
}

// Get performs a GET request to the specified path.
func (c *DatapowerClient) Get(path string, target types.String) (gjson.Result, error) {
	resp, err := c.client[stringOr(target, types.StringValue("_base_")).ValueString()].R().
		SetHeader("Accept", "application/json").
		Get(normalizePath(path))
	if err != nil {
		return gjson.Result{}, fmt.Errorf("GET request failed for path %s: %w", path, err)
	}
	if resp.IsError() {
		return gjson.Result{}, fmt.Errorf("GET request to %s returned status %d: %s", path, resp.StatusCode(), resp.String())
	}
	result := gjson.ParseBytes(resp.Body())
	if !result.Exists() {
		return gjson.Result{}, fmt.Errorf("invalid JSON response from %s: %s", path, resp.String())
	}
	return result, nil
}

// Post performs a POST request to the specified path with the provided body.
func (c *DatapowerClient) Post(path string, body interface{}, target types.String) (*resty.Response, error) {
	resp, err := c.client[stringOr(target, types.StringValue("_base_")).ValueString()].R().
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		Post(normalizePath(path))
	if err != nil {
		return nil, fmt.Errorf("POST request failed for path %s: %w", path, err)
	}
	if resp.IsError() {
		return resp, fmt.Errorf("POST request to %s returned status %d: %s", path, resp.StatusCode(), resp.String())
	}
	return resp, nil
}

// Put performs a PUT request to the specified path with the provided body.
func (c *DatapowerClient) Put(path string, body interface{}, target types.String) (*resty.Response, error) {
	resp, err := c.client[stringOr(target, types.StringValue("_base_")).ValueString()].R().
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		Put(normalizePath(path))
	if err != nil {
		return nil, fmt.Errorf("PUT request failed for path %s: %w", path, err)
	}
	if resp.IsError() {
		return resp, fmt.Errorf("PUT request to %s returned status %d: %s\n\nJSON: %v\nBody: %v", path, resp.StatusCode(), resp.String(), body, resp.Request.Body)
	}
	return resp, nil
}

// Delete performs a DELETE request to the specified path.
func (c *DatapowerClient) Delete(path string, target types.String) (*resty.Response, error) {
	resp, err := c.client[stringOr(target, types.StringValue("_base_")).ValueString()].R().
		Delete(normalizePath(path))
	if err != nil {
		return nil, fmt.Errorf("DELETE request failed for path %s: %w", path, err)
	}
	if resp.IsError() {
		return resp, fmt.Errorf("DELETE request to %s returned status %d: %s", path, resp.StatusCode(), resp.String())
	}
	return resp, nil
}

// Checks if target is valid
func (c *DatapowerClient) ValidTarget(target string) bool {
	if _, ok := c.client[target]; ok {
		return true
	}
	return false
}

// Helper to get list of valid target names for better error messages
func (c *DatapowerClient) GetTargetNames() []string {
	names := make([]string, 0, len(c.client))
	for name := range c.client {
		if name != "_base_" {
			names = append(names, name)
		}
	}
	return names
}

// normalizePath ensures the path is properly formatted (removes leading/trailing slashes).
func normalizePath(path string) string {
	// Remove single / from the start, if present.
	r, size := utf8.DecodeRuneInString(path)
	if size > 0 && strings.ContainsRune("/", r) {
		path = path[size:]
	}

	// Remove single / from the send, if present.
	if len(path) == 0 {
		return path
	}
	r, size = utf8.DecodeLastRuneInString(path)
	if size > 0 && strings.ContainsRune("/", r) {
		path = path[:len(path)-size]
	}
	return "/" + path
}

func newClient(config DatapowerClientTargetConfig) *resty.Client {
	// Initialize resty client
	restyClient := resty.New()

	restyClient.SetTransport(&http.Transport{
		// Necessary to eliminate error log entries on DataPower due resty closing the connection without sending a proper close_notify alert
		DisableKeepAlives: true,
	})

	baseURL := fmt.Sprintf("https://%s:%d", config.Hostname.ValueString(), config.Port.ValueInt32())
	restyClient.SetBaseURL(baseURL)

	if config.Insecure.ValueBool() {
		restyClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	}

	// Set Basic Auth
	restyClient.SetBasicAuth(config.Username.ValueString(), config.Password.ValueString())

	restyClient.
		AddRetryCondition(func(r *resty.Response, err error) bool {
			return err != nil || r.StatusCode() >= 500
		}).
		SetRetryCount(3).
		SetRetryWaitTime(2 * time.Second).
		SetRetryMaxWaitTime(10 * time.Second)

	restyClient.SetTimeout(30 * time.Second)

	return restyClient
}

func stringOr(val, fallback types.String) types.String {
	if val.IsNull() || val.IsUnknown() {
		return fallback
	}
	return val
}

func int32Or(val, fallback types.Int32) types.Int32 {
	if val.IsNull() || val.IsUnknown() {
		return fallback
	}
	return val
}

func boolOr(val, fallback types.Bool) types.Bool {
	if val.IsNull() || val.IsUnknown() {
		return fallback
	}
	return val
}

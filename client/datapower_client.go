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

// DatapowerClientConfig describes the provider data model.
type DatapowerClientConfig struct {
	Hostname types.String `tfsdk:"hostname"`
	Port     types.Int32  `tfsdk:"port"`
	Username types.String `tfsdk:"username"`
	Password types.String `tfsdk:"password"`
	Insecure types.Bool   `tfsdk:"insecure"`
}

// DatapowerClient describes the data maintained by the provider.
type DatapowerClient struct {
	config *DatapowerClientConfig
	client *resty.Client
}

func NewClient(config *DatapowerClientConfig) (*DatapowerClient, error) {
	// Initialize resty client
	restyClient := resty.New()

	restyClient.SetTransport(&http.Transport{
		// Necessary to eliminate error log entries on DataPower due resty closing the connection without sending a proper close_notify alert
		DisableKeepAlives: true,
	})

	var username string
	if (config.Username.IsUnknown() || config.Username.ValueString() == "") && os.Getenv("DP_USERNAME") == "" {
		return nil, fmt.Errorf("client: Cannot use unknown/empty value as username")
	}
	if !config.Username.IsNull() && config.Username.ValueString() != "" {
		username = config.Username.ValueString()
	} else {
		username = os.Getenv("DP_USERNAME")
	}

	var password string
	if (config.Password.IsUnknown() || config.Password.ValueString() == "") && os.Getenv("DP_PASSWORD") == "" {
		return nil, fmt.Errorf("client: Cannot use unknown/empty value as password")
	}
	if !config.Password.IsNull() && config.Password.ValueString() != "" {
		password = config.Password.ValueString()
	} else {
		password = os.Getenv("DP_PASSWORD")
	}

	// Set Base URL
	var hostname string
	if !config.Hostname.IsUnknown() && config.Hostname.ValueString() != "" {
		hostname = config.Hostname.ValueString()
	} else if env := os.Getenv("DP_HOSTNAME"); env != "" {
		hostname = env
	} else {
		hostname = "localhost"
	}

	var port int
	if !config.Port.IsUnknown() && config.Port.ValueInt32() > 0 {
		port = int(config.Port.ValueInt32())
	} else if env := os.Getenv("DP_PORT"); env != "" {
		p, err := strconv.Atoi(env)
		if err != nil {
			return nil, fmt.Errorf("client: Invalid port value %s", env)
		}
		port = p
	} else {
		port = 5554
	}

	if port < 1 || port > 65535 {
		return nil, fmt.Errorf("invalid port: %d (must be 1-65535)", port)
	}

	baseURL := fmt.Sprintf("https://%s:%d", hostname, port)
	restyClient.SetBaseURL(baseURL)

	insecure := config.Insecure.ValueBool()
	if env := os.Getenv("DP_INSECURE"); env != "" {
		b, err := strconv.ParseBool(env)
		if err != nil {
			return nil, fmt.Errorf("invalid DP_INSECURE value: %s", env)
		}
		insecure = insecure || b
	}
	if insecure {
		restyClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	}

	// Set Basic Auth
	restyClient.SetBasicAuth(username, password)

	restyClient.
		AddRetryCondition(func(r *resty.Response, err error) bool {
			return err != nil || r.StatusCode() >= 500
		}).
		SetRetryCount(3).
		SetRetryWaitTime(2 * time.Second).
		SetRetryMaxWaitTime(10 * time.Second)

	restyClient.SetTimeout(30 * time.Second)

	return &DatapowerClient{
		config: config,
		client: restyClient,
	}, nil
}

// Get performs a GET request to the specified path.
func (c *DatapowerClient) Get(path string) (gjson.Result, error) {
	resp, err := c.client.R().
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
func (c *DatapowerClient) Post(path string, body interface{}) (*resty.Response, error) {
	resp, err := c.client.R().
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
func (c *DatapowerClient) Put(path string, body interface{}) (*resty.Response, error) {
	resp, err := c.client.R().
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
func (c *DatapowerClient) Delete(path string) (*resty.Response, error) {
	resp, err := c.client.R().
		Delete(normalizePath(path))
	if err != nil {
		return nil, fmt.Errorf("DELETE request failed for path %s: %w", path, err)
	}
	if resp.IsError() {
		return resp, fmt.Errorf("DELETE request to %s returned status %d: %s", path, resp.StatusCode(), resp.String())
	}
	return resp, nil
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

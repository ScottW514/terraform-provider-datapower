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

	"github.com/go-resty/resty/v2"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/pkg/errors"
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
		DisableKeepAlives: true,
	})

	var username string
	if (config.Username.IsUnknown() || (config.Username.ValueString() == "")) && (os.Getenv("DP_USERNAME") == "") {
		return nil, errors.Errorf("Client: Cannot use unknown/empty value as username")
	}

	if config.Username.IsNull() {
		username = os.Getenv("DP_USERNAME")
	} else {
		username = config.Username.ValueString()
	}

	var password string
	if (config.Password.IsUnknown() || (config.Password.ValueString() == "")) && (os.Getenv("DP_PASSWORD") == "") {
		return nil, errors.Errorf("Client: Cannot use unknown/empty value as password")
	}

	if config.Password.IsNull() {
		password = os.Getenv("DP_PASSWORD")
	} else {
		password = config.Password.ValueString()
	}

	// Set Base URL
	var hostname string
	if os.Getenv("DP_HOSTNAME") != "" {
		hostname = os.Getenv("DP_HOSTNAME")
	} else if !config.Hostname.IsUnknown() && (config.Hostname.ValueString() != "") {
		hostname = config.Hostname.ValueString()
	} else {
		hostname = "localhost"
	}

	var port int
	if os.Getenv("DP_PORT") != "" {
		p, err := strconv.Atoi(os.Getenv("DP_PORT"))
		if err != nil {
			return nil, errors.Errorf("Client: Invalid port value %s", os.Getenv("DP_PORT"))
		} else {
			port = p
		}
	} else if !config.Port.IsUnknown() && config.Port.ValueInt32() > 0 {
		port = int(config.Port.ValueInt32())
	} else {
		port = 5554
	}

	baseURL := fmt.Sprintf("https://%s:%d", hostname, port)
	restyClient.SetBaseURL(baseURL)
	if config.Insecure.ValueBool() || (os.Getenv("DP_INSECURE") != "") {
		restyClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	}

	// Set Basic Auth
	restyClient.SetBasicAuth(username, password)

	restyClient.
		AddRetryCondition(func(r *resty.Response, err error) bool {
			return err != nil || r.StatusCode() >= 500
		})

	return &DatapowerClient{
		config: config,
		client: restyClient,
	}, nil
}

// Get performs a GET request to the specified path.
func (c *DatapowerClient) Get(path string) (gjson.Result, error) {
	pxy := os.Getenv("no_proxy")
	if pxy == "" {

	}
	resp, err := c.client.R().
		SetHeader("Accept", "application/json").
		Get(path)
	if err != nil {
		return gjson.Result{}, errors.Wrapf(err, "GET request failed for path %s", path)
	}
	if resp.IsError() {
		return gjson.Result{}, errors.Errorf("GET request to %s returned status %d: %s", path, resp.StatusCode(), resp.String())
	}
	return gjson.ParseBytes(resp.Body()), nil
}

// Post performs a POST request to the specified path with the provided body.
func (c *DatapowerClient) Post(path string, body interface{}) (*resty.Response, error) {
	resp, err := c.client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		Post(normalizePath(path))
	if err != nil {
		return nil, errors.Wrapf(err, "POST request failed for path %s", path)
	}
	if resp.IsError() {
		return resp, errors.Errorf("POST request to %s returned status %d: %s", path, resp.StatusCode(), resp.String())
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
		return nil, errors.Wrapf(err, "PUT request failed for path %s", path)
	}
	if resp.IsError() {
		return resp, errors.Errorf("PUT request to %s returned status %d: %s\n\nJSON:%vBody:%v", path, resp.StatusCode(), resp.String(), body, resp.Request.Body)
	}
	return resp, nil
}

// Delete performs a DELETE request to the specified path.
func (c *DatapowerClient) Delete(path string) (*resty.Response, error) {
	resp, err := c.client.R().
		Delete(normalizePath(path))
	if err != nil {
		return nil, errors.Wrapf(err, "DELETE request failed for path %s", path)
	}
	if resp.IsError() {
		return resp, errors.Errorf("DELETE request to %s returned status %d: %s", path, resp.StatusCode(), resp.String())
	}
	return resp, nil
}

// normalizePath ensures the path is properly formatted (removes leading/trailing slashes).
func normalizePath(path string) string {
	return "/" + strings.Trim(path, "/")
}

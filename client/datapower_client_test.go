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
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stretchr/testify/assert"
	"github.com/tidwall/gjson"
)

func TestNewClient(t *testing.T) {
	tests := []struct {
		name        string
		config      *DatapowerClientConfig
		envVars     map[string]string
		expectedErr string
		validate    func(t *testing.T, client *DatapowerClient)
	}{
		{
			name: "Valid config with all fields",
			config: &DatapowerClientConfig{
				Hostname: types.StringValue("example.com"),
				Port:     types.Int32Value(8443),
				Username: types.StringValue("user"),
				Password: types.StringValue("pass"),
				Insecure: types.BoolValue(true),
			},
			expectedErr: "",
			validate: func(t *testing.T, client *DatapowerClient) {
				assert.NotNil(t, client)
				assert.Equal(t, "https://datapower-dev:8443", client.client.BaseURL)
				// Verify TLS config with a test request
				server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusOK)
				}))
				defer server.Close()
				originalBaseURL := client.client.BaseURL
				client.client.SetBaseURL(server.URL)
				defer client.client.SetBaseURL(originalBaseURL)
				_, err := client.client.R().Get("/")
				assert.NoError(t, err, "Request should succeed with InsecureSkipVerify")
			},
		},
		{
			name: "Config with environment variable overrides",
			config: &DatapowerClientConfig{
				Username: types.StringNull(),
				Password: types.StringNull(),
			},
			envVars: map[string]string{
				"DP_USERNAME": "envuser",
				"DP_PASSWORD": "envpass",
				"DP_HOSTNAME": "envhost.com",
				"DP_PORT":     "8080",
			},
			expectedErr: "",
			validate: func(t *testing.T, client *DatapowerClient) {
				assert.NotNil(t, client)
				assert.Equal(t, "https://envhost.com:8080", client.client.BaseURL)
			},
		},
		{
			name: "Missing username",
			config: &DatapowerClientConfig{
				Password: types.StringValue("pass"),
			},
			envVars: map[string]string{
				"DP_USERNAME": "",
			},
			expectedErr: "Client: Cannot use unknown/empty value as username",
		},
		{
			name: "Missing password",
			config: &DatapowerClientConfig{
				Username: types.StringValue("user"),
			},
			expectedErr: "Client: Cannot use unknown/empty value as password",
		},
		{
			name: "Default values",
			config: &DatapowerClientConfig{
				Username: types.StringValue("user"),
				Password: types.StringValue("pass"),
			},
			expectedErr: "",
			validate: func(t *testing.T, client *DatapowerClient) {
				assert.NotNil(t, client)
				assert.Equal(t, "https://localhost:5554", client.client.BaseURL)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set environment variables
			for k, v := range tt.envVars {
				os.Setenv(k, v)
			}
			defer func() {
				for k := range tt.envVars {
					os.Unsetenv(k)
				}
			}()

			client, err := NewClient(tt.config)
			if tt.expectedErr != "" {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedErr)
				assert.Nil(t, client)
			} else {
				assert.NoError(t, err)
				if tt.validate != nil {
					tt.validate(t, client)
				}
			}
		})
	}
}

func TestHTTPMethods(t *testing.T) {
	config := &DatapowerClientConfig{
		Username: types.StringValue("user"),
		Password: types.StringValue("pass"),
		Insecure: types.BoolValue(true),
	}
	client, err := NewClient(config)
	assert.NoError(t, err)

	tests := []struct {
		name           string
		setupServer    func(t *testing.T, mux *http.ServeMux)
		method         string
		path           string
		body           interface{}
		expectedStatus int
		expectedResult interface{}
		expectedErr    string
	}{
		{
			name: "Successful GET",
			setupServer: func(t *testing.T, mux *http.ServeMux) {
				mux.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
					assert.Equal(t, "GET", r.Method)
					w.Header().Set("Content-Type", "application/json")
					json.NewEncoder(w).Encode(map[string]string{"key": "value"})
				})
			},
			method:         "Get",
			path:           "/test",
			expectedStatus: 200,
			expectedResult: gjson.Result{Type: gjson.JSON, Raw: "{\"key\":\"value\"}\n"},
		},
		{
			name: "GET with error status",
			setupServer: func(t *testing.T, mux *http.ServeMux) {
				mux.HandleFunc("/error", func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusInternalServerError)
					w.Write([]byte("server error"))
				})
			},
			method:      "Get",
			path:        "/error",
			expectedErr: "GET request to /error returned status 500: server error",
		},
		{
			name: "Successful POST",
			setupServer: func(t *testing.T, mux *http.ServeMux) {
				mux.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
					assert.Equal(t, "POST", r.Method)
					var body map[string]string
					json.NewDecoder(r.Body).Decode(&body)
					assert.Equal(t, "value", body["key"])
					w.WriteHeader(http.StatusCreated)
				})
			},
			method:         "Post",
			path:           "/test",
			body:           map[string]string{"key": "value"},
			expectedStatus: 201,
		},
		{
			name: "Successful PUT",
			setupServer: func(t *testing.T, mux *http.ServeMux) {
				mux.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
					assert.Equal(t, "PUT", r.Method)
					var body map[string]string
					json.NewDecoder(r.Body).Decode(&body)
					assert.Equal(t, "value", body["key"])
					w.WriteHeader(http.StatusOK)
				})
			},
			method:         "Put",
			path:           "/test",
			body:           map[string]string{"key": "value"},
			expectedStatus: 200,
		},
		{
			name: "Successful DELETE",
			setupServer: func(t *testing.T, mux *http.ServeMux) {
				mux.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
					assert.Equal(t, "DELETE", r.Method)
					w.WriteHeader(http.StatusNoContent)
				})
			},
			method:         "Delete",
			path:           "/test",
			expectedStatus: 204,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mux := http.NewServeMux()
			server := httptest.NewTLSServer(mux)
			defer server.Close()

			originalBaseURL := client.client.BaseURL
			client.client.SetBaseURL(server.URL)
			defer client.client.SetBaseURL(originalBaseURL)

			if tt.setupServer != nil {
				tt.setupServer(t, mux)
			}

			var result interface{}
			var err error

			switch tt.method {
			case "Get":
				var gjsonResult gjson.Result
				gjsonResult, err = client.Get(tt.path)
				result = gjsonResult
			case "Post":
				var resp *resty.Response
				resp, err = client.Post(tt.path, tt.body)
				if resp != nil {
					result = resp.StatusCode()
				}
			case "Put":
				var resp *resty.Response
				resp, err = client.Put(tt.path, tt.body)
				if resp != nil {
					result = resp.StatusCode()
				}
			case "Delete":
				var resp *resty.Response
				resp, err = client.Delete(tt.path)
				if resp != nil {
					result = resp.StatusCode()
				}
			}

			if tt.expectedErr != "" {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedErr)
			} else {
				assert.NoError(t, err)
				if tt.method == "Get" {
					assert.Equal(t, tt.expectedResult.(gjson.Result).Raw, result.(gjson.Result).Raw)
				} else {
					assert.Equal(t, tt.expectedStatus, result)
				}
			}
		})
	}
}

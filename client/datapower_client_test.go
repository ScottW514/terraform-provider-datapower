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
	"encoding/base64"
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

// relevantEnvVars lists all environment variables used by NewClient.
var relevantEnvVars = []string{
	"DP_USERNAME",
	"DP_PASSWORD",
	"DP_HOSTNAME",
	"DP_PORT",
	"DP_INSECURE",
}

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
				assert.Equal(t, "https://example.com:8443", client.client.BaseURL)
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
			name: "Secure TLS fails on self-signed",
			config: &DatapowerClientConfig{
				Hostname: types.StringValue("localhost"),
				Port:     types.Int32Value(5554),
				Username: types.StringValue("user"),
				Password: types.StringValue("pass"),
				Insecure: types.BoolValue(false),
			},
			expectedErr: "",
			validate: func(t *testing.T, client *DatapowerClient) {
				server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
				defer server.Close()
				originalBaseURL := client.client.BaseURL
				client.client.SetBaseURL(server.URL)
				defer client.client.SetBaseURL(originalBaseURL)
				_, err := client.client.R().Get("/")
				assert.Error(t, err)
				if assert.Error(t, err) {
					assert.Contains(t, err.Error(), "certificate signed by unknown authority")
				}
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
			name: "Invalid port from env",
			config: &DatapowerClientConfig{
				Username: types.StringValue("user"),
				Password: types.StringValue("pass"),
			},
			envVars: map[string]string{
				"DP_PORT": "abc",
			},
			expectedErr: "client: Invalid port value abc",
		},
		{
			name: "Invalid DP_INSECURE value",
			config: &DatapowerClientConfig{
				Username: types.StringValue("user"),
				Password: types.StringValue("pass"),
			},
			envVars: map[string]string{
				"DP_INSECURE": "notabool",
			},
			expectedErr: "invalid DP_INSECURE value: notabool",
		},
		{
			name: "DP_INSECURE true overrides config false",
			config: &DatapowerClientConfig{
				Hostname: types.StringValue("localhost"),
				Port:     types.Int32Value(5554),
				Username: types.StringValue("user"),
				Password: types.StringValue("pass"),
				Insecure: types.BoolValue(false),
			},
			envVars: map[string]string{
				"DP_INSECURE": "true",
			},
			expectedErr: "",
			validate: func(t *testing.T, client *DatapowerClient) {
				server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusOK)
				}))
				defer server.Close()
				originalBaseURL := client.client.BaseURL
				client.client.SetBaseURL(server.URL)
				defer client.client.SetBaseURL(originalBaseURL)
				_, err := client.client.R().Get("/")
				assert.NoError(t, err, "Request should succeed due to DP_INSECURE=true")
			},
		},
		{
			name: "DP_INSECURE false does not override config true",
			config: &DatapowerClientConfig{
				Hostname: types.StringValue("localhost"),
				Port:     types.Int32Value(5554),
				Username: types.StringValue("user"),
				Password: types.StringValue("pass"),
				Insecure: types.BoolValue(true),
			},
			envVars: map[string]string{
				"DP_INSECURE": "false",
			},
			expectedErr: "",
			validate: func(t *testing.T, client *DatapowerClient) {
				server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusOK)
				}))
				defer server.Close()
				originalBaseURL := client.client.BaseURL
				client.client.SetBaseURL(server.URL)
				defer client.client.SetBaseURL(originalBaseURL)
				_, err := client.client.R().Get("/")
				assert.NoError(t, err, "Request should succeed due to config Insecure=true")
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
			expectedErr: "client: Cannot use unknown/empty value as username",
		},
		{
			name: "Missing password",
			config: &DatapowerClientConfig{
				Username: types.StringValue("user"),
			},
			expectedErr: "client: Cannot use unknown/empty value as password",
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
			t.Parallel()
			// Unset relevant environment variables to isolate tests
			for _, envVar := range relevantEnvVars {
				os.Unsetenv(envVar)
			}
			// Set test-specific environment variables
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
				if assert.Error(t, err) {
					assert.Contains(t, err.Error(), tt.expectedErr)
				}
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
	// Unset relevant environment variables before creating client
	for _, envVar := range relevantEnvVars {
		os.Unsetenv(envVar)
	}
	config := &DatapowerClientConfig{
		Username: types.StringValue("user"),
		Password: types.StringValue("pass"),
		Insecure: types.BoolValue(true),
	}
	client, err := NewClient(config)
	assert.NoError(t, err)

	expectedAuth := "Basic " + base64.StdEncoding.EncodeToString([]byte("user:pass"))

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
					assert.Equal(t, expectedAuth, r.Header.Get("Authorization"))
					assert.Equal(t, "application/json", r.Header.Get("Accept"))
					w.Header().Set("Content-Type", "application/json")
					json.NewEncoder(w).Encode(map[string]string{"key": "value"})
				})
			},
			method:         "Get",
			path:           "/test",
			expectedResult: gjson.Result{Type: gjson.JSON, Raw: "{\"key\":\"value\"}\n"},
		},
		{
			name: "GET with error status",
			setupServer: func(t *testing.T, mux *http.ServeMux) {
				mux.HandleFunc("/error", func(w http.ResponseWriter, r *http.Request) {
					assert.Equal(t, expectedAuth, r.Header.Get("Authorization"))
					w.WriteHeader(http.StatusInternalServerError)
					w.Write([]byte("server error"))
				})
			},
			method:      "Get",
			path:        "/error",
			expectedErr: "GET request to /error returned status 500: server error",
		},
		{
			name: "GET with retry on 500",
			setupServer: func(t *testing.T, mux *http.ServeMux) {
				count := 0
				mux.HandleFunc("/retry", func(w http.ResponseWriter, r *http.Request) {
					assert.Equal(t, expectedAuth, r.Header.Get("Authorization"))
					count++
					if count < 3 {
						w.WriteHeader(500)
						return
					}
					w.Header().Set("Content-Type", "application/json")
					json.NewEncoder(w).Encode(map[string]string{"key": "retried"})
				})
			},
			method:         "Get",
			path:           "/retry",
			expectedResult: gjson.Result{Type: gjson.JSON, Raw: "{\"key\":\"retried\"}\n"},
		},
		{
			name: "Successful GET with trailing slash",
			setupServer: func(t *testing.T, mux *http.ServeMux) {
				mux.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
					assert.Equal(t, expectedAuth, r.Header.Get("Authorization"))
					w.Header().Set("Content-Type", "application/json")
					json.NewEncoder(w).Encode(map[string]string{"key": "value"})
				})
			},
			method:         "Get",
			path:           "/test/",
			expectedResult: gjson.Result{Type: gjson.JSON, Raw: "{\"key\":\"value\"}\n"},
		},
		{
			name: "Successful GET without leading slash",
			setupServer: func(t *testing.T, mux *http.ServeMux) {
				mux.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
					assert.Equal(t, expectedAuth, r.Header.Get("Authorization"))
					w.Header().Set("Content-Type", "application/json")
					json.NewEncoder(w).Encode(map[string]string{"key": "value"})
				})
			},
			method:         "Get",
			path:           "test",
			expectedResult: gjson.Result{Type: gjson.JSON, Raw: "{\"key\":\"value\"}\n"},
		},
		{
			name: "Successful POST",
			setupServer: func(t *testing.T, mux *http.ServeMux) {
				mux.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
					assert.Equal(t, "POST", r.Method)
					assert.Equal(t, expectedAuth, r.Header.Get("Authorization"))
					assert.Equal(t, "application/json", r.Header.Get("Content-Type"))
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
			name: "POST with error status",
			setupServer: func(t *testing.T, mux *http.ServeMux) {
				mux.HandleFunc("/error", func(w http.ResponseWriter, r *http.Request) {
					assert.Equal(t, expectedAuth, r.Header.Get("Authorization"))
					w.WriteHeader(http.StatusBadRequest)
					w.Write([]byte("bad request"))
				})
			},
			method:      "Post",
			path:        "/error",
			body:        map[string]string{"key": "value"},
			expectedErr: "POST request to /error returned status 400: bad request",
		},
		{
			name: "Successful PUT",
			setupServer: func(t *testing.T, mux *http.ServeMux) {
				mux.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
					assert.Equal(t, "PUT", r.Method)
					assert.Equal(t, expectedAuth, r.Header.Get("Authorization"))
					assert.Equal(t, "application/json", r.Header.Get("Content-Type"))
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
			name: "PUT with error status",
			setupServer: func(t *testing.T, mux *http.ServeMux) {
				mux.HandleFunc("/error", func(w http.ResponseWriter, r *http.Request) {
					assert.Equal(t, expectedAuth, r.Header.Get("Authorization"))
					w.WriteHeader(http.StatusBadRequest)
					w.Write([]byte("bad request"))
				})
			},
			method:      "Put",
			path:        "/error",
			body:        map[string]string{"key": "value"},
			expectedErr: "PUT request to /error returned status 400: bad request\n\nJSON: map[key:value]\nBody: map[key:value]",
		},
		{
			name: "Successful DELETE",
			setupServer: func(t *testing.T, mux *http.ServeMux) {
				mux.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
					assert.Equal(t, "DELETE", r.Method)
					assert.Equal(t, expectedAuth, r.Header.Get("Authorization"))
					w.WriteHeader(http.StatusNoContent)
				})
			},
			method:         "Delete",
			path:           "/test",
			expectedStatus: 204,
		},
		{
			name: "DELETE with error status",
			setupServer: func(t *testing.T, mux *http.ServeMux) {
				mux.HandleFunc("/error", func(w http.ResponseWriter, r *http.Request) {
					assert.Equal(t, expectedAuth, r.Header.Get("Authorization"))
					w.WriteHeader(http.StatusBadRequest)
					w.Write([]byte("bad request"))
				})
			},
			method:      "Delete",
			path:        "/error",
			expectedErr: "DELETE request to /error returned status 400: bad request",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Removed t.Parallel() to avoid race conditions on shared client
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
				if assert.Error(t, err) {
					assert.Contains(t, err.Error(), tt.expectedErr)
				}
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

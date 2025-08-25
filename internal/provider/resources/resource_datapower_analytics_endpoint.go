// Copyright © 2025 Scott Wiederhold <s.e.wiederhold@gmail.com>
// All rights reserved.
//
// Licensed under the Mozilla Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://mozilla.org/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: MPL-2.0

// This file is generated "gen/generator.go"
// !!CHANGES TO THIS FILE WILL BE OVERWRITTEN!!

package resources

import (
	"context"
	"fmt"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/modifiers"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/validators"
)

var _ resource.Resource = &AnalyticsEndpointResource{}

func NewAnalyticsEndpointResource() resource.Resource {
	return &AnalyticsEndpointResource{}
}

type AnalyticsEndpointResource struct {
	pData *tfutils.ProviderData
}

func (r *AnalyticsEndpointResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_analytics_endpoint"
}

func (r *AnalyticsEndpointResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("An analytics endpoint buffers API event data and offloads the collected data as a bulk activity log to a remote server. When offloaded, you can use this data for display and analysis.", "analytics-endpoint", "").String,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Name of the object. Must be unique among object types in application domain.", "", "").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 128),
					stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9_-]+$"), "Must match :"+"^[a-zA-Z0-9_-]+$"),
				},
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"app_domain": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("The name of the application domain the object belongs to", "", "").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 128),
					stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9_-]+$"), "Must match :"+"^[a-zA-Z0-9_-]+$"),
				},
				PlanModifiers: []planmodifier.String{
					modifiers.ImmutableAfterSet(),
				},
			},
			"user_summary": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Comments", "summary", "").String,
				Optional:            true,
			},
			"analytics_server_url": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the URL to offload the collected API event data. The URL can start with <tt>http</tt> or <tt>https</tt> for an Elasticsearch server or start with <tt>dpkafka</tt> for a Kafka server. <ul><li>For an Elasticsearch server, specify the full URL to the endpoint starting with the <tt>http</tt> or <tt>https</tt> protocol identifier. With HTTPS, you must assign a TLS client profile.</li><li>For a Kafka server, specify only the name of the existing Kafka cluster configuration after the <tt>dpkafka</tt> protocol identifier. To complete the URL, you must specify which request topic to offload analytics data.</li></ul>", "analytics-server-url", "").String,
				Required:            true,
			},
			"ssl_client": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("TLS client profile", "ssl-client", "ssl_client_profile").AddRequiredWhen(models.AnalyticsEndpointSSLClientCondVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.AnalyticsEndpointSSLClientCondVal, validators.Evaluation{}, false),
				},
			},
			"request_topic": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Request topic", "request-topic", "").AddRequiredWhen(models.AnalyticsEndpointRequestTopicCondVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.AnalyticsEndpointRequestTopicCondVal, validators.Evaluation{}, false),
				},
			},
			"max_records": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the maximum number of records that can be buffered for each API gateway. The collected analytics data for an API gateway is offloaded when 80% of this value or the defined interval is reached. The value must be a power of 2. Enter a value in the range 256 - 65536. The default value is 1024.", "max-records", "").AddIntegerRange(256, 65536).AddDefaultValue("1024").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(256, 65536),
				},
				Default: int64default.StaticInt64(1024),
			},
			"max_records_memory_kb": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the maximum size for each record in KB. Enter a value in the range 4 - 1024. The default value is 512.", "max-record-size", "").AddIntegerRange(4, 1024).AddDefaultValue("512").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(4, 1024),
				},
				Default: int64default.StaticInt64(512),
			},
			"max_delivery_memory_mb": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the maximum size for each delivery in MB. Enter a value in the range 1 - 1024. The default value is 512.", "max-delivery-size", "").AddIntegerRange(1, 1024).AddDefaultValue("512").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 1024),
				},
				Default: int64default.StaticInt64(512),
			},
			"interval": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the interval in seconds between offloads. Data is offloaded at this interval or when an API gateway reaches 80% of the value set for maximum records. Enter a value in the range 1 - 3600. The default value is 600", "interval", "").AddIntegerRange(1, 3600).AddDefaultValue("600").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 3600),
				},
				Default: int64default.StaticInt64(600),
			},
			"delivery_connections": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the number of connections to establish per delivery to the remote server to offload analytics data. Each connection can carry a bulk activity log. Enter a value in the range 1 - 100. The default value is 1.", "delivery-connections", "").AddIntegerRange(1, 100).AddDefaultValue("1").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 100),
				},
				Default: int64default.StaticInt64(1),
			},
			"enable_jwt": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enable JWT feature sending logs to analytics server.", "enable-jwt", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"management_url": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the URL of management platform endpoint to retrieve a JWT. The URL must use the <tt>http</tt> or <tt>https</tt> protocol.", "management-url", "").AddRequiredWhen(models.AnalyticsEndpointManagementURLCondVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.AnalyticsEndpointManagementURLCondVal, validators.Evaluation{}, false),
				},
			},
			"management_url_ssl_client": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Management platform TLS client profile", "management-ssl-client", "ssl_client_profile").AddRequiredWhen(models.AnalyticsEndpointManagementURL_SSLClientCondVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.AnalyticsEndpointManagementURL_SSLClientCondVal, validators.Evaluation{}, false),
				},
			},
			"client_id": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Client ID", "clientid", "").AddRequiredWhen(models.AnalyticsEndpointClientIDCondVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.AnalyticsEndpointClientIDCondVal, validators.Evaluation{}, false),
				},
			},
			"client_secret_alias": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Client secret", "client-secret-alias", "password_alias").AddRequiredWhen(models.AnalyticsEndpointClientSecretAliasCondVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(0, 127),
					validators.ConditionalRequiredString(models.AnalyticsEndpointClientSecretAliasCondVal, validators.Evaluation{}, false),
				},
			},
			"grant_type": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the grant type for requesting JWT tokens. Only the client credentials grant type is supported.", "grant-type", "").AddStringEnum("implicit", "password", "application", "accessCode").AddRequiredWhen(models.AnalyticsEndpointGrantTypeCondVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("implicit", "password", "application", "accessCode"),
					validators.ConditionalRequiredString(models.AnalyticsEndpointGrantTypeCondVal, validators.Evaluation{}, false),
				},
			},
			"scope": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the scope for requesting JWT tokens. The value is in the <tt>openid analytics_subsystem_ID/name</tt> format.", "scope", "").AddRequiredWhen(models.AnalyticsEndpointScopeCondVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.AnalyticsEndpointScopeCondVal, validators.Evaluation{}, false),
				},
			},
			"persistent_connection": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to negotiate persistent connections. By default, persistent connections are enabled. The HTTP/2 protocol controls persistent connections and reuse. Therefore, these settings are ignored.", "persistent-connection", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"timeout": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the intra-transaction timeout for connections, which is the maximum idle time to allow in a transaction. This timer monitors idle time in the data transfer process. When the idle time is exceeded, the connection is torn down. Enter a value in the range 1 - 86400. The default value is 90.", "timeout", "").AddIntegerRange(1, 86400).AddDefaultValue("90").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 86400),
				},
				Default: int64default.StaticInt64(90),
			},
			"persistent_timeout": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the inter-transaction timeout for connections, which is the maximum idle time to allow between the completion of a TCP transaction and the initiation of a new TCP transaction. When the idle time is exceeded, the connection is torn down. Enter a value in the range 1 - 86400. The default value is 60.", "persistent-timeout", "").AddIntegerRange(1, 86400).AddDefaultValue("60").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 86400),
				},
				Default: int64default.StaticInt64(60),
			},
			"dependency_actions": actions.ActionsSchema,
		},
	}
}

func (r *AnalyticsEndpointResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (r *AnalyticsEndpointResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.AnalyticsEndpoint
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Create, false)
	if resp.Diagnostics.HasError() {
		return
	}

	body := data.ToBody(ctx, `AnalyticsEndpoint`)
	_, err := r.pData.Client.Post(data.GetPath(), body)

	if err != nil && !strings.Contains(err.Error(), "status 409") {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to create object (%s), got error: %s", "POST", err))
		return
	}
	actions.PostProcess(ctx, &resp.Diagnostics, data.DependencyActions, actions.Create)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AnalyticsEndpointResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.AnalyticsEndpoint
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	res, err := r.pData.Client.Get(data.GetPath() + "/" + data.Id.ValueString())
	if err != nil && (strings.Contains(err.Error(), "status 404") || strings.Contains(err.Error(), "status 406") || strings.Contains(err.Error(), "status 500") || strings.Contains(err.Error(), "status 400")) {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s", err))
		return
	}

	if data.IsNull() {
		// Import
		data.FromBody(ctx, `AnalyticsEndpoint`, res)
	} else {
		// Update
		data.UpdateFromBody(ctx, `AnalyticsEndpoint`, res)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AnalyticsEndpointResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.AnalyticsEndpoint
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Update, false)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.pData.Client.Put(data.GetPath()+"/"+data.Id.ValueString(), data.ToBody(ctx, `AnalyticsEndpoint`))
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to update object (PUT), got error: %s", err))
		return
	}

	actions.PostProcess(ctx, &resp.Diagnostics, data.DependencyActions, actions.Update)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AnalyticsEndpointResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.AnalyticsEndpoint
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Delete, false)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.pData.Client.Delete(data.GetPath() + "/" + data.Id.ValueString())
	if err != nil && !strings.Contains(err.Error(), "status 404") && !strings.Contains(err.Error(), "status 409") {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object (DELETE), got error: %s", err))
		return
	}

	actions.PostProcess(ctx, &resp.Diagnostics, data.DependencyActions, actions.Delete)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.State.RemoveResource(ctx)
}

func (r *AnalyticsEndpointResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data models.AnalyticsEndpoint

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.ValidateConfig(ctx, &resp.Diagnostics, data.DependencyActions)
}

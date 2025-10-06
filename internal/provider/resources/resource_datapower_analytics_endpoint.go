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
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/modifiers"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/validators"
)

var _ resource.Resource = &AnalyticsEndpointResource{}
var _ resource.ResourceWithImportState = &AnalyticsEndpointResource{}

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
				MarkdownDescription: tfutils.NewAttributeDescription("TLS client profile", "ssl-client", "ssl_client_profile").AddRequiredWhen(models.AnalyticsEndpointSSLClientCondVal.String()).AddNotValidWhen(models.AnalyticsEndpointSSLClientIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.AnalyticsEndpointSSLClientCondVal, models.AnalyticsEndpointSSLClientIgnoreVal, false),
				},
			},
			"request_topic": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Request topic", "request-topic", "").AddRequiredWhen(models.AnalyticsEndpointRequestTopicCondVal.String()).AddNotValidWhen(models.AnalyticsEndpointRequestTopicIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.AnalyticsEndpointRequestTopicCondVal, models.AnalyticsEndpointRequestTopicIgnoreVal, false),
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
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the number of connections to establish per delivery to the remote server to offload analytics data. Each connection can carry a bulk activity log. Enter a value in the range 1 - 100. The default value is 1.", "delivery-connections", "").AddIntegerRange(1, 100).AddDefaultValue("1").AddNotValidWhen(models.AnalyticsEndpointDeliveryConnectionsIgnoreVal.String()).String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 100),
					validators.ConditionalRequiredInt64(validators.Evaluation{}, models.AnalyticsEndpointDeliveryConnectionsIgnoreVal, true),
				},
				Default: int64default.StaticInt64(1),
			},
			"enable_jwt": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enable JWT feature sending logs to analytics server.", "enable-jwt", "").AddDefaultValue("false").AddNotValidWhen(models.AnalyticsEndpointEnableJWTIgnoreVal.String()).String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"management_url": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the URL of management platform endpoint to retrieve a JWT. The URL must use the <tt>http</tt> or <tt>https</tt> protocol.", "management-url", "").AddRequiredWhen(models.AnalyticsEndpointManagementURLCondVal.String()).AddNotValidWhen(models.AnalyticsEndpointManagementURLIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.AnalyticsEndpointManagementURLCondVal, models.AnalyticsEndpointManagementURLIgnoreVal, false),
				},
			},
			"management_url_ssl_client": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Management platform TLS client profile", "management-ssl-client", "ssl_client_profile").AddRequiredWhen(models.AnalyticsEndpointManagementURL_SSLClientCondVal.String()).AddNotValidWhen(models.AnalyticsEndpointManagementURL_SSLClientIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.AnalyticsEndpointManagementURL_SSLClientCondVal, models.AnalyticsEndpointManagementURL_SSLClientIgnoreVal, false),
				},
			},
			"client_id": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Client ID", "clientid", "").AddRequiredWhen(models.AnalyticsEndpointClientIDCondVal.String()).AddNotValidWhen(models.AnalyticsEndpointClientIDIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.AnalyticsEndpointClientIDCondVal, models.AnalyticsEndpointClientIDIgnoreVal, false),
				},
			},
			"client_secret_alias": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Client secret", "client-secret-alias", "password_alias").AddRequiredWhen(models.AnalyticsEndpointClientSecretAliasCondVal.String()).AddNotValidWhen(models.AnalyticsEndpointClientSecretAliasIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(0, 127),
					validators.ConditionalRequiredString(models.AnalyticsEndpointClientSecretAliasCondVal, models.AnalyticsEndpointClientSecretAliasIgnoreVal, false),
				},
			},
			"grant_type": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the grant type for requesting JWT tokens. Only the client credentials grant type is supported.", "grant-type", "").AddStringEnum("implicit", "password", "application", "accessCode").AddRequiredWhen(models.AnalyticsEndpointGrantTypeCondVal.String()).AddNotValidWhen(models.AnalyticsEndpointGrantTypeIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("implicit", "password", "application", "accessCode"),
					validators.ConditionalRequiredString(models.AnalyticsEndpointGrantTypeCondVal, models.AnalyticsEndpointGrantTypeIgnoreVal, false),
				},
			},
			"scope": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the scope for requesting JWT tokens. The value is in the <tt>openid analytics_subsystem_ID/name</tt> format.", "scope", "").AddRequiredWhen(models.AnalyticsEndpointScopeCondVal.String()).AddNotValidWhen(models.AnalyticsEndpointScopeIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.AnalyticsEndpointScopeCondVal, models.AnalyticsEndpointScopeIgnoreVal, false),
				},
			},
			"persistent_connection": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to negotiate persistent connections. By default, persistent connections are enabled. The HTTP/2 protocol controls persistent connections and reuse. Therefore, these settings are ignored.", "persistent-connection", "").AddDefaultValue("true").AddNotValidWhen(models.AnalyticsEndpointPersistentConnectionIgnoreVal.String()).String,
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
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the inter-transaction timeout for connections, which is the maximum idle time to allow between the completion of a TCP transaction and the initiation of a new TCP transaction. When the idle time is exceeded, the connection is torn down. Enter a value in the range 1 - 86400. The default value is 60.", "persistent-timeout", "").AddIntegerRange(1, 86400).AddDefaultValue("60").AddNotValidWhen(models.AnalyticsEndpointPersistentTimeoutIgnoreVal.String()).String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 86400),
					validators.ConditionalRequiredInt64(validators.Evaluation{}, models.AnalyticsEndpointPersistentTimeoutIgnoreVal, true),
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
	if err != nil {
		if strings.Contains(err.Error(), "status 409") {
			_, err := r.pData.Client.Put(data.GetPath()+"/"+data.Id.ValueString(), body)
			if err != nil {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Resource already exists. Failed to update resource, got error: %s", err))
				return
			}
			resp.Diagnostics.AddWarning("Warning", "Resource already exists. Existing resource was updated.")
		} else {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to create resource, got error: %s", err))
			return
		}
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

	data.UpdateFromBody(ctx, `AnalyticsEndpoint`, res)

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

func (r *AnalyticsEndpointResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()
	parts := strings.Split(req.ID, "/")
	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		resp.Diagnostics.AddError("Invalid Import ID", "Expected format: <app_domain>/<id>. Got: "+req.ID)
		return
	}

	appDomain := parts[0]
	id := parts[1]

	if !regexp.MustCompile("^[a-zA-Z0-9_-]+$").MatchString(id) || len(id) < 1 || len(id) > 128 {
		resp.Diagnostics.AddError("Invalid ID", "ID must be 1-128 characters and match pattern ^[a-zA-Z0-9_-]+$")
		return
	}
	if !regexp.MustCompile("^[a-zA-Z0-9_-]+$").MatchString(appDomain) || len(appDomain) < 1 || len(appDomain) > 128 {
		resp.Diagnostics.AddError("Invalid Application Domain", "Application domain must be 1-128 characters and match pattern ^[a-zA-Z0-9_-]+$")
		return
	}

	var data models.AnalyticsEndpoint
	data.AppDomain = types.StringValue(appDomain)
	data.Id = types.StringValue(id)
	res, err := r.pData.Client.Get(data.GetPath() + "/" + data.Id.ValueString())
	if err != nil {
		if strings.Contains(err.Error(), "status 404") {
			resp.Diagnostics.AddError("Resource Not Found", fmt.Sprintf("Resource was not found, got error: %s", err))
		} else {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		}
		return
	}

	data.FromBody(ctx, `AnalyticsEndpoint`, res)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AnalyticsEndpointResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data models.AnalyticsEndpoint

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.ValidateConfig(ctx, &resp.Diagnostics, data.DependencyActions)
}

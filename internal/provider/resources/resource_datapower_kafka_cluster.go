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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/modifiers"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/validators"
)

var _ resource.Resource = &KafkaClusterResource{}
var _ resource.ResourceWithImportState = &KafkaClusterResource{}

func NewKafkaClusterResource() resource.Resource {
	return &KafkaClusterResource{}
}

type KafkaClusterResource struct {
	pData *tfutils.ProviderData
}

func (r *KafkaClusterResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_kafka_cluster"
}

func (r *KafkaClusterResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("Define the Kafka cluster that is responsible for the messaging services. The Kafka cluster periodically monitors and polls topics. The Kafka cluster ensures that sent messages are directed to the correct response topic or are routed to another server.", "kafka-cluster", "").String,
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
			"protocol": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the transport protocol for the Kafka bootstrap connection. The selected protocol is used for the exchange of information between the Kafka server and the bootstrap server. By default, uses a non-encrypted transport.", "protocol", "").AddStringEnum("plaintext", "ssl", "sasl_plaintext", "sasl_ssl").AddDefaultValue("plaintext").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("plaintext", "ssl", "sasl_plaintext", "sasl_ssl"),
				},
				Default: stringdefault.StaticString("plaintext"),
			},
			"endpoint": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the endpoints for the bootstrap process. A bootstrap server uses a host name or IP address and a port to define an endpoint address. You can add multiple nondefault bootstrap servers. For failover capability, the endpoints must be members of the same cluster.", "endpoint", "").String,
				NestedObject:        models.GetDmKafkaEndpointResourceSchema(),
				Required:            true,
			},
			"sasl_mechanism": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the Simple Authentication and Security Layer (SASL) mechanism to communicate with the Kafka cluster. By default, uses a clear text password.", "sasl-mechanism", "").AddStringEnum("plain", "scram-sha-256", "scram-sha-512").AddDefaultValue("plain").AddRequiredWhen(models.KafkaClusterSASLMechanismCondVal.String()).AddNotValidWhen(models.KafkaClusterSASLMechanismIgnoreVal.String()).String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("plain", "scram-sha-256", "scram-sha-512"),
					validators.ConditionalRequiredString(models.KafkaClusterSASLMechanismCondVal, models.KafkaClusterSASLMechanismIgnoreVal, true),
				},
				Default: stringdefault.StaticString("plain"),
			},
			"user_name": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Username", "username", "").AddRequiredWhen(models.KafkaClusterUserNameCondVal.String()).AddNotValidWhen(models.KafkaClusterUserNameIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.RegexMatches(regexp.MustCompile("^[^ ]+$"), "Must match :"+"^[^ ]+$"),
					validators.ConditionalRequiredString(models.KafkaClusterUserNameCondVal, models.KafkaClusterUserNameIgnoreVal, false),
				},
			},
			"password_alias": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Password alias", "password-alias", "password_alias").AddRequiredWhen(models.KafkaClusterPasswordAliasCondVal.String()).AddNotValidWhen(models.KafkaClusterPasswordAliasIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.KafkaClusterPasswordAliasCondVal, models.KafkaClusterPasswordAliasIgnoreVal, false),
				},
			},
			"autocommit": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to commit offsets at the defined interval or at process-completion. <ul><li>When enabled, commits offsets at the defined interval. The default interval is 5 seconds. To change the interval, set the <tt>auto.commit.interval.ms</tt> property.</li><li>When disabled, commits offsets at process-completion. You can use the batch size setting for the Kafka handle to define the number of messages to attempt to receive from the consumer.</li></ul>", "autocommit", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"ssl_client": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("TLS client profile", "ssl-client", "ssl_client_profile").AddRequiredWhen(models.KafkaClusterSSLClientCondVal.String()).AddNotValidWhen(models.KafkaClusterSSLClientIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.KafkaClusterSSLClientCondVal, models.KafkaClusterSSLClientIgnoreVal, false),
				},
			},
			"memory_threshold": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the maximum memory to allocate in bytes. Enter a value in the range 10485760 - 1073741824. The default value is 268435456.", "memory-threshold", "").AddIntegerRange(10485760, 1073741824).AddDefaultValue("268435456").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(10485760, 1073741824),
				},
				Default: int64default.StaticInt64(268435456),
			},
			"maximum_message_size": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the maximum message size in bytes. Enter a value in the range 0 - 1073741824. The default value is 1048576. A value of 0 disables the enforcement of a maximum message size.", "maximum-message-size", "").AddIntegerRange(0, 1073741824).AddDefaultValue("1048576").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 1073741824),
				},
				Default: int64default.StaticInt64(1048576),
			},
			"auto_retry": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Automatic retry", "auto-retry", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"retry_interval": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the interval between attempts to reestablish a connection in seconds. Enter a value in the range 1 - 65535. The default value is 10.", "retry-interval", "").AddIntegerRange(1, 65535).AddDefaultValue("10").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 65535),
				},
				Default: int64default.StaticInt64(10),
			},
			"property": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify extra property to configure the connection to the Kafka server. Use this property for each extra property that is required. Some properties are unsupported and will cause a configuration failure.", "property", "").String,
				NestedObject:        models.GetDmKafkaPropertyResourceSchema(),
				Optional:            true,
			},
			"dependency_actions": actions.ActionsSchema,
		},
	}
}

func (r *KafkaClusterResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (r *KafkaClusterResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.KafkaCluster
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Create, false)
	if resp.Diagnostics.HasError() {
		return
	}

	body := data.ToBody(ctx, `KafkaCluster`)
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

func (r *KafkaClusterResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.KafkaCluster
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

	data.UpdateFromBody(ctx, `KafkaCluster`, res)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *KafkaClusterResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.KafkaCluster
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Update, false)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.pData.Client.Put(data.GetPath()+"/"+data.Id.ValueString(), data.ToBody(ctx, `KafkaCluster`))
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

func (r *KafkaClusterResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.KafkaCluster
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

func (r *KafkaClusterResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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

	var data models.KafkaCluster
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

	data.FromBody(ctx, `KafkaCluster`, res)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *KafkaClusterResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data models.KafkaCluster

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.ValidateConfig(ctx, &resp.Diagnostics, data.DependencyActions)
}

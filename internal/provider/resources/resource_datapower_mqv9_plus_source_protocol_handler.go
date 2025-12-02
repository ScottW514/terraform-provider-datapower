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

var (
	_ resource.Resource                   = &MQv9PlusSourceProtocolHandlerResource{}
	_ resource.ResourceWithValidateConfig = &MQv9PlusSourceProtocolHandlerResource{}
	_ resource.ResourceWithImportState    = &MQv9PlusSourceProtocolHandlerResource{}
)

func NewMQv9PlusSourceProtocolHandlerResource() resource.Resource {
	return &MQv9PlusSourceProtocolHandlerResource{}
}

type MQv9PlusSourceProtocolHandlerResource struct {
	pData *tfutils.ProviderData
}

func (r *MQv9PlusSourceProtocolHandlerResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_mqv9_plus_source_protocol_handler"
}

func (r *MQv9PlusSourceProtocolHandlerResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("Configure the IBM MQ v9+ handler to manage IBM MQ protocol communications.", "source-idg-mq", "").AddActions("quiesce").String,
		Attributes: map[string]schema.Attribute{
			"provider_target": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Target host for this resource. If not set, provider will use the top level settings.", "", "").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 128),
					stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9_-]+$"), "Must match :"+"^[a-zA-Z0-9_-]+$"),
				},
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
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
			"queue_manager": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the name of the queue manager that provides messaging services for communicating applications by periodically monitoring or polling queues and by ensuring that messages are directed to the correct receive queue or routed to another queue manager. The local queue manager corresponds to a queue manager running on another host on the network.", "queue-manager", "").String,
				Required:            true,
			},
			"get_queue": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the name of the get queue associated with the queue manager. The handler gets messages from this queue.", "get-queue", "").AddRequiredWhen(models.MQv9PlusSourceProtocolHandlerGetQueueCondVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.MQv9PlusSourceProtocolHandlerGetQueueCondVal, validators.Evaluation{}, false),
				},
			},
			"subscribe_topic_string": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the the topic string associated with the queue manager. The handler subscribes to this topic string and gets messages from it.", "subscribe-topic-string", "").AddRequiredWhen(models.MQv9PlusSourceProtocolHandlerSubscribeTopicStringCondVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.MQv9PlusSourceProtocolHandlerSubscribeTopicStringCondVal, validators.Evaluation{}, false),
				},
			},
			"subscription_name": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the subscription name of a durable subscription associated with the queue manager. This name identifies the subscription after reestablishing a lost connection.", "sub-name", "").String,
				Optional:            true,
			},
			"put_queue": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the the name of the put queue associated with the queue manager. The handler puts server-originated reply messages to this queue.", "put-queue", "").String,
				Optional:            true,
			},
			"publish_topic_string": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the publish topic string associated with the queue manager. The topic string describes the subject of the information that is published in a publish or subscribe message. The handler publishes messages to this topic string. If the put queue is specified, this property is ignored.", "publish-topic-string", "").String,
				Optional:            true,
			},
			"code_page": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Specify the coded character set identifier to which the remote IBM MQ queue manager converts output data. This property is meaningful only when the queue manager has the convert property set to on. The CCSI will be put in the CodeCharSetId field of MQMD.</p><p>The default CCSI is for ISO-8859-1 (latin-1).</p><p>For MQCCSI_EMBEDDED enter 4294967295; for MQCCSI_INHERIT enter 4294967294. For the other CCSIDs, refer to the IBM Code Pages.</p>", "ccsi", "").AddIntegerRange(0, 4294967295).AddDefaultValue("0").String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 4294967295),
				},
				Default: int64default.StaticInt64(0),
			},
			"get_message_options": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the cumulative value of the MQGET options that are applicable to an IBM MQ message in decimal or hex format. The value is passed directly to the IBM MQ API. The default value is 1, which is the decimal value for the MQGMO_WAIT option.", "get-message-options", "").AddIntegerRange(0, 4294967295).AddDefaultValue("1").String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 4294967295),
				},
				Default: int64default.StaticInt64(1),
			},
			"message_selector": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the selector that filters the messages from a queue or a subscription by their properties. The selector uses the SQL92 conditional expression syntax. The selector enables the handler to filter messages from a queue or a subscription.", "selector", "").String,
				Optional:            true,
			},
			"parse_properties": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to parse the properties of the incoming messages from a queue or a subscription. By default, properties are not parsed.", "parse-properties", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"async_put": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to asynchronously put a message to a queue without waiting for a response from the queue manager.", "async-put", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"exclude_headers": models.GetDmMQHeadersResourceSchema("Specify the headers after MQMD to strip from the message. By default only the MQMD header is parsed.", "exclude-headers", "", false),
			"concurrent_connections": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the number of concurrent IBM MQ conversations to allocate. The default value is 1 but can be increased to improve performance.", "concurrent-connections", "").AddIntegerRange(1, 65535).AddDefaultValue("1").AddNotValidWhen(models.MQv9PlusSourceProtocolHandlerConcurrentConnectionsIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 65535),
					validators.ConditionalRequiredInt64(validators.Evaluation{}, models.MQv9PlusSourceProtocolHandlerConcurrentConnectionsIgnoreVal, true),
				},
				Default: int64default.StaticInt64(1),
			},
			"polling_interval": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the duration in seconds to wait after processing all messages before attempting to retrieve messages from the get queue.", "polling-interval", "").AddIntegerRange(0, 65535).AddDefaultValue("30").String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 65535),
				},
				Default: int64default.StaticInt64(30),
			},
			"batch_size": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the number of messages to process as a batch. The handler gathers the specified number of messages and processes them as a batch.", "batch-size", "").AddIntegerRange(0, 65535).AddDefaultValue("0").String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 65535),
				},
				Default: int64default.StaticInt64(0),
			},
			"content_type_header": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Header to extract Content-Type", "content-type-header", "").AddStringEnum("None", "MQRFH", "MQRFH2").AddDefaultValue("None").String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("None", "MQRFH", "MQRFH2"),
				},
				Default: stringdefault.StaticString("None"),
			},
			"content_type_xpath": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("XPath expression to extract Content-Type from IBM MQ header", "content-type-xpath", "").AddRequiredWhen(models.MQv9PlusSourceProtocolHandlerContentTypeXPathCondVal.String()).AddNotValidWhen(models.MQv9PlusSourceProtocolHandlerContentTypeXPathIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.MQv9PlusSourceProtocolHandlerContentTypeXPathCondVal, models.MQv9PlusSourceProtocolHandlerContentTypeXPathIgnoreVal, false),
				},
			},
			"retrieve_backout_settings": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to retrieve backout setting from the IBM MQ server. <p>When enabled, retrieves the <b>Backout threshold</b> and <b>Backout requeue queue name</b> settings from the IBM MQ server and checks these values. On a reattempt, the handler uses the higher priority backout settings from the server. If the server does not contain backout settings, The handler uses any existing backout values, either empty or populated, from the local IBM MQ queue manager. If there are no backout settings, the backout function is disabled.</p><p>When an alias queue is used, its attributes are retrieved, not those of the base queue.</p>", "retrieve-backout-settings", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"use_qm_name_in_url": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether the var://service/URL-in variable returns the name of the local queue manager or queue manager group when this configuration defines a queue manager group as the queue manager. <ul><li>When enabled, the variable returns the name of the queue manager.</li><li>When not enabled, the variable returns the name of the queue manager group. This setting is the default value.</li></ul>", "use-qm-in-url", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"dependency_actions": actions.ActionsSchema,
		},
	}
}

func (r *MQv9PlusSourceProtocolHandlerResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (r *MQv9PlusSourceProtocolHandlerResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.MQv9PlusSourceProtocolHandler
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if data.ProviderTarget.ValueString() != "" && !r.pData.Client.ValidTarget(data.ProviderTarget.ValueString()) {
		resp.Diagnostics.AddError("Invalid provider_target", fmt.Sprintf(`Target %q is not defined in the provider's 'targets' block. Available targets: %v`, data.ProviderTarget.ValueString(), r.pData.Client.GetTargetNames()))
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Create, false, data.ProviderTarget)
	if resp.Diagnostics.HasError() {
		return
	}

	body := data.ToBody(ctx, `MQv9PlusSourceProtocolHandler`)
	_, err := r.pData.Client.Post(data.GetPath(), body, data.ProviderTarget)
	if err != nil {
		if strings.Contains(err.Error(), "status 409") {
			_, err := r.pData.Client.Put(data.GetPath()+"/"+data.Id.ValueString(), body, data.ProviderTarget)
			if err != nil {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Resource already exists. Failed to update resource, got error: %s", err))
				return
			}
			resp.Diagnostics.AddWarning("Warning", "Resource already exists. Existing resource was updated.")
		} else if strings.Contains(err.Error(), "status 401") {
			_ = tfutils.DomainCredentialTest(r.pData.Client, &resp.Diagnostics, data.AppDomain.ValueString(), data.ProviderTarget)
			if !resp.Diagnostics.HasError() {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Application Domain '%s' does not exist", data.AppDomain.ValueString()))
			}
			return
		} else {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to create resource, got error: %s", err))
			return
		}
	}
	actions.PostProcess(ctx, &resp.Diagnostics, data.DependencyActions, actions.Create, data.ProviderTarget)
	if resp.Diagnostics.HasError() {
		return
	}

	tfutils.SaveDomain(ctx, &resp.Diagnostics, r.pData.Client, data.AppDomain, data.ProviderTarget)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *MQv9PlusSourceProtocolHandlerResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.MQv9PlusSourceProtocolHandler
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if data.ProviderTarget.ValueString() != "" && !r.pData.Client.ValidTarget(data.ProviderTarget.ValueString()) {
		resp.Diagnostics.AddError("Invalid provider_target", fmt.Sprintf(`Target %q is not defined in the provider's 'targets' block. Available targets: %v`, data.ProviderTarget.ValueString(), r.pData.Client.GetTargetNames()))
		return
	}
	res, err := r.pData.Client.Get(data.GetPath()+"/"+data.Id.ValueString(), data.ProviderTarget)
	if err != nil {
		if strings.Contains(err.Error(), "status 404") || strings.Contains(err.Error(), "status 406") || strings.Contains(err.Error(), "status 500") || strings.Contains(err.Error(), "status 400") {
			resp.State.RemoveResource(ctx)
			return
		} else if strings.Contains(err.Error(), "status 401") {
			_ = tfutils.DomainCredentialTest(r.pData.Client, &resp.Diagnostics, data.AppDomain.ValueString(), data.ProviderTarget)
			if !resp.Diagnostics.HasError() {
				resp.State.RemoveResource(ctx)
			}
			return
		} else {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s", err))
			return
		}
	}

	data.UpdateFromBody(ctx, `MQv9PlusSourceProtocolHandler`, res)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *MQv9PlusSourceProtocolHandlerResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.MQv9PlusSourceProtocolHandler
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if data.ProviderTarget.ValueString() != "" && !r.pData.Client.ValidTarget(data.ProviderTarget.ValueString()) {
		resp.Diagnostics.AddError("Invalid provider_target", fmt.Sprintf(`Target %q is not defined in the provider's 'targets' block. Available targets: %v`, data.ProviderTarget.ValueString(), r.pData.Client.GetTargetNames()))
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Update, false, data.ProviderTarget)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.pData.Client.Put(data.GetPath()+"/"+data.Id.ValueString(), data.ToBody(ctx, `MQv9PlusSourceProtocolHandler`), data.ProviderTarget)
	if err != nil {
		if strings.Contains(err.Error(), "status 401") {
			_ = tfutils.DomainCredentialTest(r.pData.Client, &resp.Diagnostics, data.AppDomain.ValueString(), data.ProviderTarget)
			if !resp.Diagnostics.HasError() {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Application Domain '%s' does not exist", data.AppDomain.ValueString()))
			}
			return
		} else {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to update object (PUT), got error: %s", err))
			return
		}
	}

	actions.PostProcess(ctx, &resp.Diagnostics, data.DependencyActions, actions.Update, data.ProviderTarget)
	if resp.Diagnostics.HasError() {
		return
	}

	tfutils.SaveDomain(ctx, &resp.Diagnostics, r.pData.Client, data.AppDomain, data.ProviderTarget)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *MQv9PlusSourceProtocolHandlerResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.MQv9PlusSourceProtocolHandler
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if data.ProviderTarget.ValueString() != "" && !r.pData.Client.ValidTarget(data.ProviderTarget.ValueString()) {
		resp.Diagnostics.AddError("Invalid provider_target", fmt.Sprintf(`Target %q is not defined in the provider's 'targets' block. Available targets: %v`, data.ProviderTarget.ValueString(), r.pData.Client.GetTargetNames()))
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Delete, false, data.ProviderTarget)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.pData.Client.Delete(data.GetPath()+"/"+data.Id.ValueString(), data.ProviderTarget)
	if err != nil {
		if strings.Contains(err.Error(), "status 409") {
			resp.Diagnostics.AddWarning("Resource Conflict", fmt.Sprintf("Resource is no longer tracked by Terraform, but may need to be manually deleted on DataPower host. Got error: %s", err))
		} else if strings.Contains(err.Error(), "status 401") {
			_ = tfutils.DomainCredentialTest(r.pData.Client, &resp.Diagnostics, data.AppDomain.ValueString(), data.ProviderTarget)
			if !resp.Diagnostics.HasError() {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Application Domain '%s' does not exist", data.AppDomain.ValueString()))
			}
			return
		} else if !strings.Contains(err.Error(), "status 404") {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete resource, got error: %s", err))
			return
		}
	}

	actions.PostProcess(ctx, &resp.Diagnostics, data.DependencyActions, actions.Delete, data.ProviderTarget)
	if resp.Diagnostics.HasError() {
		return
	}

	tfutils.SaveDomain(ctx, &resp.Diagnostics, r.pData.Client, data.AppDomain, data.ProviderTarget)

	resp.State.RemoveResource(ctx)
}

func (r *MQv9PlusSourceProtocolHandlerResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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

	var data models.MQv9PlusSourceProtocolHandler
	data.AppDomain = types.StringValue(appDomain)
	data.Id = types.StringValue(id)
	res, err := r.pData.Client.Get(data.GetPath()+"/"+data.Id.ValueString(), data.ProviderTarget)
	if err != nil {
		if strings.Contains(err.Error(), "status 404") {
			resp.Diagnostics.AddError("Resource Not Found", fmt.Sprintf("Resource was not found, got error: %s", err))
		} else if strings.Contains(err.Error(), "status 401") {
			_ = tfutils.DomainCredentialTest(r.pData.Client, &resp.Diagnostics, data.AppDomain.ValueString(), data.ProviderTarget)
			if !resp.Diagnostics.HasError() {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Application Domain '%s' does not exist", data.AppDomain.ValueString()))
			}
		} else {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		}
		return
	}

	data.FromBody(ctx, `MQv9PlusSourceProtocolHandler`, res)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
func (r *MQv9PlusSourceProtocolHandlerResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data models.MQv9PlusSourceProtocolHandler

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.ValidateConfig(ctx, &resp.Diagnostics, data.DependencyActions)
}

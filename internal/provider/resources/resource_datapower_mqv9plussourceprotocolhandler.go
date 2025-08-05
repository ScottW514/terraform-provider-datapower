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
	"github.com/scottw514/terraform-provider-datapower/client"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/modifiers"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
)

var _ resource.Resource = &MQv9PlusSourceProtocolHandlerResource{}

func NewMQv9PlusSourceProtocolHandlerResource() resource.Resource {
	return &MQv9PlusSourceProtocolHandlerResource{}
}

type MQv9PlusSourceProtocolHandlerResource struct {
	client *client.DatapowerClient
}

func (r *MQv9PlusSourceProtocolHandlerResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_mqv9plussourceprotocolhandler"
}

func (r *MQv9PlusSourceProtocolHandlerResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("IBM MQ v9+ handler", "source-idg-mq", "").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Name of the object. Must be unique among object types in application domain.", "", "").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 128),
					stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z0-9_-]+$`), ""),
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
					stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z0-9_-]+$`), ""),
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
				MarkdownDescription: tfutils.NewAttributeDescription("Queue manager (reference to MQManger or MQManagerGroup)", "queue-manager", "").String,
				Required:            true,
			},
			"get_queue": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Get queue", "get-queue", "").String,
				Optional:            true,
			},
			"subscribe_topic_string": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Subscribe topic string", "subscribe-topic-string", "").String,
				Optional:            true,
			},
			"subscription_name": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Subscription name", "sub-name", "").String,
				Optional:            true,
			},
			"put_queue": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Put queue", "put-queue", "").String,
				Optional:            true,
			},
			"publish_topic_string": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Publish topic string", "publish-topic-string", "").String,
				Optional:            true,
			},
			"code_page": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("CCSI", "ccsi", "").AddIntegerRange(0, 4294967295).AddDefaultValue("0").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(0, 4294967295),
				},
				Default: int64default.StaticInt64(0),
			},
			"get_message_options": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Get message options", "get-message-options", "").AddIntegerRange(0, 4294967295).AddDefaultValue("1").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(0, 4294967295),
				},
				Default: int64default.StaticInt64(1),
			},
			"message_selector": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Selector", "selector", "").String,
				Optional:            true,
			},
			"parse_properties": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Parse properties", "parse-properties", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"async_put": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Async put", "async-put", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"exclude_headers": models.GetDmMQHeadersResourceSchema("Exclude message headers", "exclude-headers", "", false),
			"concurrent_connections": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Concurrent conversations", "concurrent-connections", "").AddIntegerRange(1, 65535).AddDefaultValue("1").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(1, 65535),
				},
				Default: int64default.StaticInt64(1),
			},
			"polling_interval": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Polling interval", "polling-interval", "").AddIntegerRange(0, 65535).AddDefaultValue("30").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(0, 65535),
				},
				Default: int64default.StaticInt64(30),
			},
			"batch_size": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Batch size", "batch-size", "").AddIntegerRange(0, 65535).AddDefaultValue("0").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(0, 65535),
				},
				Default: int64default.StaticInt64(0),
			},
			"content_type_header": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Header to extract Content-Type", "content-type-header", "").AddStringEnum("None", "MQRFH", "MQRFH2").AddDefaultValue("None").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("None", "MQRFH", "MQRFH2"),
				},
				Default: stringdefault.StaticString("None"),
			},
			"content_type_x_path": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("XPath expression to extract Content-Type from IBM MQ header", "content-type-xpath", "").String,
				Optional:            true,
			},
			"retrieve_backout_settings": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Retrieve backout settings", "retrieve-backout-settings", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"use_qm_name_in_url": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Use queue manager in URL", "use-qm-in-url", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
		},
	}
}

func (r *MQv9PlusSourceProtocolHandlerResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = *req.ProviderData.(**client.DatapowerClient)
}

func (r *MQv9PlusSourceProtocolHandlerResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data models.MQv9PlusSourceProtocolHandler

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	body := data.ToBody(ctx, `MQv9PlusSourceProtocolHandler`)
	_, err := r.client.Post(data.GetPath(), body)

	if err != nil && !strings.Contains(err.Error(), "status 409") {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to create object (%s), got error: %s", "POST", err))
		return
	}

	_, err = r.client.Post("/mgmt/actionqueue/"+data.AppDomain.ValueString(), "{\"SaveConfig\": 0}")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to save object (%s), got error: %s", "POST", err))
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *MQv9PlusSourceProtocolHandlerResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data models.MQv9PlusSourceProtocolHandler

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	res, err := r.client.Get(data.GetPath() + "/" + data.Id.ValueString())
	if err != nil && (strings.Contains(err.Error(), "status 404") || strings.Contains(err.Error(), "status 406") || strings.Contains(err.Error(), "status 500") || strings.Contains(err.Error(), "status 400")) {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s", err))
		return
	}

	if data.IsNull() {
		// Import
		data.FromBody(ctx, `MQv9PlusSourceProtocolHandler`, res)
	} else {
		// Update
		data.UpdateFromBody(ctx, `MQv9PlusSourceProtocolHandler`, res)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *MQv9PlusSourceProtocolHandlerResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data models.MQv9PlusSourceProtocolHandler

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.client.Put(data.GetPath()+"/"+data.Id.ValueString(), data.ToBody(ctx, `MQv9PlusSourceProtocolHandler`))
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to update object (PUT), got error: %s", err))
		return
	}
	_, err = r.client.Post("/mgmt/actionqueue/"+data.AppDomain.ValueString(), "{\"SaveConfig\": 0}")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to save object (%s), got error: %s", "POST", err))
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *MQv9PlusSourceProtocolHandlerResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data models.MQv9PlusSourceProtocolHandler

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.client.Delete(data.GetPath() + "/" + data.Id.ValueString())
	if err != nil && !strings.Contains(err.Error(), "status 404") && !strings.Contains(err.Error(), "status 409") {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object (DELETE), got error: %s", err))
		return
	}

	resp.State.RemoveResource(ctx)
}

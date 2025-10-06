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

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/modifiers"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
)

var _ resource.Resource = &WebSphereJMSSourceProtocolHandlerResource{}
var _ resource.ResourceWithValidateConfig = &WebSphereJMSSourceProtocolHandlerResource{}
var _ resource.ResourceWithImportState = &WebSphereJMSSourceProtocolHandlerResource{}

func NewWebSphereJMSSourceProtocolHandlerResource() resource.Resource {
	return &WebSphereJMSSourceProtocolHandlerResource{}
}

type WebSphereJMSSourceProtocolHandlerResource struct {
	pData *tfutils.ProviderData
}

func (r *WebSphereJMSSourceProtocolHandlerResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_web_sphere_jms_source_protocol_handler"
}

func (r *WebSphereJMSSourceProtocolHandlerResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("Configure the WebSphere JMS handler to manage WebSphere JMS protocol communications.", "source-wasjms", "").AddActions("quiesce").String,
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
			"server": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Select the WebSphere JMS server object supported by this handler.", "server", "web_sphere_jms_server").String,
				Required:            true,
			},
			"request_topic_space": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Use this property to disambiguate a topic if the request destination is a topic whose name appears in multiple topic spaces.</p><p>A topic space is a hierarchy of topics used for publish/subscribe messaging. Topics with the same name can exist in multiple topic spaces, but there can be only one topic space with a given name in a service integration bus.</p><p>For example, consider a topic hierarchy split into the following topic spaces:</p><ul><li>library - topics for document management</li><li>sales - topics for marketing and sales tracking</li><li>engineering - topics for engineering and technology</li></ul><p>The topic <em>volumes</em> can appear in all three topic spaces, and have a different meaning in each.</p><p>Enter the name of the target topic space if necessary.</p>", "request-topic-space", "").String,
				Optional:            true,
			},
			"reply_topic_space": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Use this property to disambiguate a topic if the response destination is a topic whose name appears in multiple topic spaces.</p><p>A topic space is a hierarchy of topics used for publish/subscribe messaging. Topics with the same name can exist in multiple topic spaces, but there can be only one topic space with a given name in a service integration bus.</p><p>For example, consider a topic hierarchy split into the following topic spaces:</p><ul><li>library - topics for document management</li><li>sales - topics for marketing and sales tracking</li><li>engineering - topics for engineering and technology</li></ul><p>The topic <em>volumes</em> can appear in all three topic spaces, and have a different meaning in each.</p><p>Enter the name of the target topic space if necessary.</p>", "reply-topic-space", "").String,
				Optional:            true,
			},
			"strict_message_order": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Enable to work smoothly with WebSphere server when the \"Strict Message Ordering\" option in the corresponding destination is checked.</p>", "strict-message-order", "").AddDefaultValue("false").AddNotValidWhen(models.WebSphereJMSSourceProtocolHandlerStrictMessageOrderIgnoreVal.String()).String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"user_summary": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Comments", "summary", "").String,
				Optional:            true,
			},
			"get_queue": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Enter the name of the get queue associated with the WebSphere JMS or TIBCO EMS object supported by this handler.</p><p>The handler monitors the get queue for incoming client requests. Upon message receipt, The handler forwards the extracted message to the DataPower object that will gateway the message to a remote message provider.</p>", "get-queue", "").String,
				Required:            true,
			},
			"put_queue": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Enter the name of the put queue associated with the WebSphere JMS or TIBCO EMS object supported by this handler.</p><p>The put queue contains server-originated WAS JMS or TIBCO EMS reply messages.</p><p>Such messages are originated by a remote WAS JMS or TIBCO EMS message provider and put into this queue by a local WebSphere JMS or TIBCO EMS object.</p><p>Configuration of a put queue is optional.</p><p>A put queue should be configured if server replies are expected; if reply messages are not expected, a put queue need not be configured.</p><p>In the absence of a put queue, any received replies are dropped.</p>", "put-queue", "").String,
				Optional:            true,
			},
			"selector": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Provide an SQL-like expression to filter messages from the GET queue.</p><p>For example, <tt>DeliveryMode LIKE PERSISTENT</tt></p><p>This expression specifies that only client requests that have a DeliveryMode of PERSISTENT are forwarded to the WebSphere JMS or TIBCO EMS object for processing; all other messages are dropped from the get queue.</p><p>The message selector is a conditional expression based on a subset of SQL92 conditional expression syntax. The conditional expression enables the handler to identify <em>messages of interest</em> .</p><p>The conditional expression does not operate on the body of the message, rather it examines message headers and properties (proprietary user-created headers that might appear between the required headers and the message body).</p><p>The required headers are as follows:</p><ul><li><tt>Destination</tt> - contains the destination (queue) to which the message is being sent</li><li><tt>DeliveryMode</tt> - contains the delivery mode (PERSISTENT or NON_PERSISTENT)</li><li><tt>Expiration</tt> - contains a message TTL or a value of 0 indicating an unlimited TTL</li><li><tt>Priority</tt> - contains the message priority expressed as a digit from 0 (lowest priority) to 9 (highest priority)</li><li><tt>MessageID</tt> - contains a unique message identifier starting with the prefix ID:, or a null value, effectively disabling message ID</li><li><tt>Timestamp</tt> - contains the time the message was handed off for transmission, not the time it was actually sent</li><li><tt>CorrelationID</tt> - contains a means of associating one message (for example, a response) with another message (for example, the original request)</li><li><tt>ReplyTo</tt> - contains the destination (queue) to which a reply to this message should be sent</li><li><tt>Type</tt> - contains a message identifier provided by the application</li><li><tt>Redelivered</tt> - contains a boolean indicating that the message has been delivered in the past, but not yet acknowledged</li></ul><p>Configuration of a message selector is optional. If a message selector is not specified, all incoming client request messages are transferred by The handler to the DataPower object for processing.</p>", "selector", "").String,
				Optional:            true,
			},
			"async_message_processing": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>If enabled messages taken from the get queue will be processed not necessarily in the same order as they were queued.</p><p>This property may be set to improve performance only if associated Multi-Protocol Gateway or WS-Proxy isn't configured to process messages in order.</p>", "async-message-processing", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"dependency_actions": actions.ActionsSchema,
		},
	}
}

func (r *WebSphereJMSSourceProtocolHandlerResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (r *WebSphereJMSSourceProtocolHandlerResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.WebSphereJMSSourceProtocolHandler
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Create, false)
	if resp.Diagnostics.HasError() {
		return
	}

	body := data.ToBody(ctx, `WebSphereJMSSourceProtocolHandler`)
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

func (r *WebSphereJMSSourceProtocolHandlerResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.WebSphereJMSSourceProtocolHandler
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

	data.UpdateFromBody(ctx, `WebSphereJMSSourceProtocolHandler`, res)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *WebSphereJMSSourceProtocolHandlerResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.WebSphereJMSSourceProtocolHandler
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Update, false)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.pData.Client.Put(data.GetPath()+"/"+data.Id.ValueString(), data.ToBody(ctx, `WebSphereJMSSourceProtocolHandler`))
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

func (r *WebSphereJMSSourceProtocolHandlerResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.WebSphereJMSSourceProtocolHandler
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

func (r *WebSphereJMSSourceProtocolHandlerResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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

	var data models.WebSphereJMSSourceProtocolHandler
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

	data.FromBody(ctx, `WebSphereJMSSourceProtocolHandler`, res)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *WebSphereJMSSourceProtocolHandlerResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data models.WebSphereJMSSourceProtocolHandler

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.ValidateConfig(ctx, &resp.Diagnostics, data.DependencyActions)
}

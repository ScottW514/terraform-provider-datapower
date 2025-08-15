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
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/modifiers"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
)

var _ resource.Resource = &AMQPBrokerResource{}

func NewAMQPBrokerResource() resource.Resource {
	return &AMQPBrokerResource{}
}

type AMQPBrokerResource struct {
	pData *tfutils.ProviderData
}

func (r *AMQPBrokerResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_amqp_broker"
}

func (r *AMQPBrokerResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("In AMQP, distributed source and target termini are managed by a broker. The broker provides messaging services for communicating applications by periodically monitoring and polling termini. The broker ensures that sent messages are directed to the correct target terminus or are routed to another server. The AMQP broker configuration corresponds to an AMQP broker that is running on another host in the network. The configured properties enable communication between the DataPower Gateway and the remote AMQP broker.", "amqp-broker", "").String,
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
			"host_name": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Host", "hostname", "").String,
				Required:            true,
			},
			"port": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Port", "port", "").AddIntegerRange(1, 65535).AddDefaultValue("5672").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(1, 65535),
				},
				Default: int64default.StaticInt64(5672),
			},
			"xml_manager": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the XML manager to control access to the remote AMQP server. The XML manager obtains and manages documents.", "xml-manager", "xml_manager").AddDefaultValue("default").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("default"),
			},
			"container_id": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("AMQP container ID", "container-id", "").String,
				Optional:            true,
			},
			"authorization": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the SASL layer that the AMQP broker uses to authenticate with the AMQP server. The default setting in no authentication.", "authorization", "").AddStringEnum("none", "anonymous", "plain").AddDefaultValue("none").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("none", "anonymous", "plain"),
				},
				Default: stringdefault.StaticString("none"),
			},
			"user_name": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Username", "user", "").String,
				Optional:            true,
			},
			"password_alias": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Password alias", "password-alias", "password_alias").String,
				Optional:            true,
			},
			"maximum_frame_size": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the maximum frame size in bytes to allow. Frames Frames that are larger are rejected. When rejected, the connection is closed. Enter a value in the range 512 - 104857600. The default value is 104857600.", "maximum-frame-size", "").AddIntegerRange(512, 104857600).AddDefaultValue("104857600").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(512, 104857600),
				},
				Default: int64default.StaticInt64(104857600),
			},
			"auto_retry": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to enable the automatic retry procedure after an AMQP connection failure. By default, the automatic retry behavior is enabled. This setting does not affect attempts over an established connection.", "auto-retry", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"retry_interval": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the interval in seconds to wait before attempting to reestablish a failed connection. After the number of attempts is reached, attempts to reestablish a failed connection use the interval that is defined by the long retry interval. Enter a value in the range 1 - 65535. The default value is 10. <p>This setting does not affect attempts over an established connection.</p>", "retry-interval", "").AddIntegerRange(1, 86400).AddDefaultValue("10").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(1, 86400),
				},
				Default: int64default.StaticInt64(10),
			},
			"retry_attempts": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the number of attempts for a failed connection to the remote AMQP server. After the number of attempts is reached, the long retry interval is used. Enter a value in the range 0 - 65535. The default value is 6. The special value of 0 disables the long interval, where the retry interval is used forever.", "retry-attempts", "").AddDefaultValue("6").String,
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(6),
			},
			"long_retry_interval": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the interval in seconds to use after the number of attempts is reached to attempt to reestablish a failed connection. Enter a value in the range 1 - 65535. The default value is 600. <p>This setting does not affect attempts over an established connection.</p>", "long-retry-interval", "").AddIntegerRange(1, 86400).AddDefaultValue("600").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(1, 86400),
				},
				Default: int64default.StaticInt64(600),
			},
			"reporting_interval": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the interval in seconds between the writing of identical log message. Enter a value in the range 1 - 65535. The default value is 10.", "reporting-interval", "").AddIntegerRange(1, 86400).AddDefaultValue("10").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(1, 86400),
				},
				Default: int64default.StaticInt64(10),
			},
			"ssl_client": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("TLS client profile", "ssl-client", "ssl_client_profile").String,
				Optional:            true,
			},
			"dependency_actions": actions.ActionsSchema,
		},
	}
}

func (r *AMQPBrokerResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (r *AMQPBrokerResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.AMQPBroker
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Create, false)
	if resp.Diagnostics.HasError() {
		return
	}

	body := data.ToBody(ctx, `AMQPBroker`)
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

func (r *AMQPBrokerResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.AMQPBroker
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
		data.FromBody(ctx, `AMQPBroker`, res)
	} else {
		// Update
		data.UpdateFromBody(ctx, `AMQPBroker`, res)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AMQPBrokerResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.AMQPBroker
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Update, false)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.pData.Client.Put(data.GetPath()+"/"+data.Id.ValueString(), data.ToBody(ctx, `AMQPBroker`))
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

func (r *AMQPBrokerResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.AMQPBroker
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

func (r *AMQPBrokerResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data models.AMQPBroker

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.ValidateConfig(ctx, &resp.Diagnostics, data.DependencyActions)
}

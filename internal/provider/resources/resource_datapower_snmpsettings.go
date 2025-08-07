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
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/client"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
)

var _ resource.Resource = &SNMPSettingsResource{}

func NewSNMPSettingsResource() resource.Resource {
	return &SNMPSettingsResource{}
}

type SNMPSettingsResource struct {
	client *client.DatapowerClient
}

func (r *SNMPSettingsResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_snmpsettings"
}

func (r *SNMPSettingsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("SNMP Settings (`default` domain only)", "snmp", "").String,
		Attributes: map[string]schema.Attribute{
			"enabled": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Administrative state", "admin-state", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"user_summary": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Comments", "summary", "").String,
				Optional:            true,
			},
			"local_address": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Local IP Address", "ip-address", "").AddDefaultValue("0.0.0.0").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("0.0.0.0"),
			},
			"local_port": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Local Port", "port", "").AddDefaultValue("161").String,
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(161),
			},
			"policies": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("SNMPv1/v2c Communities", "community", "").String,
				NestedObject:        models.DmSnmpPolicyResourceSchema,
				Optional:            true,
			},
			"policies_mq": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("SNMPv1/v2c Communities", "snmp-community", "").String,
				NestedObject:        models.DmSnmpPolicyMQResourceSchema,
				Optional:            true,
			},
			"targets": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Trap and Notification Targets", "trap-target", "").String,
				NestedObject:        models.DmSnmpTargetResourceSchema,
				Optional:            true,
			},
			"users": schema.ListAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("SNMPv3 Users", "user", "user").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"contexts": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("SNMPv3 Contexts", "context", "").String,
				NestedObject:        models.DmSnmpContextResourceSchema,
				Optional:            true,
			},
			"security_level": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("SNMPv3 Security Level", "security-level", "").AddStringEnum("noAuthNoPriv", "authNoPriv", "authPriv").AddDefaultValue("authPriv").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("noAuthNoPriv", "authNoPriv", "authPriv"),
				},
				Default: stringdefault.StaticString("authPriv"),
			},
			"access_level": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("SNMPv3 Access Level", "access-level", "").AddStringEnum("none", "read-only", "read-write").AddDefaultValue("read-only").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("none", "read-only", "read-write"),
				},
				Default: stringdefault.StaticString("read-only"),
			},
			"enable_default_trap_subscriptions": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enable Default Event Subscriptions", "trap-default-subscriptions", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"trap_priority": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Minimum Priority", "trap-priority", "").AddStringEnum("emerg", "alert", "critic", "error", "warn", "notice", "info", "debug").AddDefaultValue("error").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("emerg", "alert", "critic", "error", "warn", "notice", "info", "debug"),
				},
				Default: stringdefault.StaticString("error"),
			},
			"trap_event_code": schema.ListAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Event Subscriptions", "trap-code", "").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"config_mib": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Configuration", "config-mib", "").AddDefaultValue("/drConfigMIB.txt").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("/drConfigMIB.txt"),
			},
			"config_mib_mq": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Configuration", "config-mib", "").AddDefaultValue("/mqConfigMIB.txt").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("/mqConfigMIB.txt"),
			},
			"status_mib": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Status", "status-mib", "").AddDefaultValue("/drStatusMIB.txt").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("/drStatusMIB.txt"),
			},
			"status_mib_mq": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Status", "status-mib", "").AddDefaultValue("/mqStatusMIB.txt").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("/mqStatusMIB.txt"),
			},
			"notif_mib": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Notifications", "notification-mib", "").AddDefaultValue("/drNotificationMIB.txt").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("/drNotificationMIB.txt"),
			},
			"notif_mib_mq": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Notifications", "notification-mib", "").AddDefaultValue("/mqNotificationMIB.txt").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("/mqNotificationMIB.txt"),
			},
			"dependency_actions": actions.ActionsSchema,
		},
	}
}

func (r *SNMPSettingsResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = *req.ProviderData.(**client.DatapowerClient)
}

func (r *SNMPSettingsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data models.SNMPSettings

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, r.client, "default", data.DependencyActions, actions.Create)

	body := data.ToBody(ctx, `SNMPSettings`)
	_, err := r.client.Put(data.GetPath(), body)

	if err != nil && !strings.Contains(err.Error(), "status 409") {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to create object (%s), got error: %s", "PUT", err))
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SNMPSettingsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data models.SNMPSettings

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	res, err := r.client.Get(data.GetPath())
	if err != nil && (strings.Contains(err.Error(), "status 404") || strings.Contains(err.Error(), "status 406") || strings.Contains(err.Error(), "status 500") || strings.Contains(err.Error(), "status 400")) {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s", err))
		return
	}

	if data.IsNull() {
		// Import
		data.FromBody(ctx, `SNMPSettings`, res)
	} else {
		// Update
		data.UpdateFromBody(ctx, `SNMPSettings`, res)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SNMPSettingsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data models.SNMPSettings

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, r.client, "default", data.DependencyActions, actions.Update)
	_, err := r.client.Put(data.GetPath(), data.ToBody(ctx, `SNMPSettings`))
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to update object (PUT), got error: %s", err))
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SNMPSettingsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

	resp.State.RemoveResource(ctx)
}

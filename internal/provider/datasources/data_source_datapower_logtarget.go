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

package datasources

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
)

type LogTargetList struct {
	AppDomain types.String `tfsdk:"app_domain"`
	Result    types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &LogTargetDataSource{}
	_ datasource.DataSourceWithConfigure = &LogTargetDataSource{}
)

func NewLogTargetDataSource() datasource.DataSource {
	return &LogTargetDataSource{}
}

type LogTargetDataSource struct {
	pData *tfutils.ProviderData
}

func (d *LogTargetDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_logtarget"
}

func (d *LogTargetDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Log Target",
		Attributes: map[string]schema.Attribute{
			"app_domain": schema.StringAttribute{
				MarkdownDescription: "The name of the application domain the object belongs to",
				Required:            true,
			},
			"result": schema.ListNestedAttribute{
				MarkdownDescription: "List of objects",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: "Name of the object. Must be unique among object types in application domain.",
							Computed:            true,
						},
						"app_domain": schema.StringAttribute{
							MarkdownDescription: "The name of the application domain the object belongs to",
							Computed:            true,
						},
						"backup": schema.StringAttribute{
							MarkdownDescription: "Backup log",
							Computed:            true,
						},
						"log_events": schema.ListNestedAttribute{
							MarkdownDescription: "Event subscriptions",
							NestedObject:        models.DmLogEventDataSourceSchema,
							Computed:            true,
						},
						"user_summary": schema.StringAttribute{
							MarkdownDescription: "Comments",
							Computed:            true,
						},
						"type": schema.StringAttribute{
							MarkdownDescription: "Target type",
							Computed:            true,
						},
						"priority": schema.StringAttribute{
							MarkdownDescription: "Logging priority",
							Computed:            true,
						},
						"soap_version": schema.StringAttribute{
							MarkdownDescription: "SOAP version",
							Computed:            true,
						},
						"format": schema.StringAttribute{
							MarkdownDescription: "Log format",
							Computed:            true,
						},
						"timestamp_format": schema.StringAttribute{
							MarkdownDescription: "Timestamp format",
							Computed:            true,
						},
						"fixed_format": schema.BoolAttribute{
							MarkdownDescription: "Fixed format",
							Computed:            true,
						},
						"local_identifier": schema.StringAttribute{
							MarkdownDescription: "Local identifier",
							Computed:            true,
						},
						"email_address": schema.StringAttribute{
							MarkdownDescription: "Recipient email address",
							Computed:            true,
						},
						"sender_address": schema.StringAttribute{
							MarkdownDescription: "Sender email address",
							Computed:            true,
						},
						"smtp_domain": schema.StringAttribute{
							MarkdownDescription: "SMTP client domain",
							Computed:            true,
						},
						"size": schema.Int64Attribute{
							MarkdownDescription: "Log size",
							Computed:            true,
						},
						"url": schema.StringAttribute{
							MarkdownDescription: "URL",
							Computed:            true,
						},
						"nfs_mount": schema.StringAttribute{
							MarkdownDescription: "NFS static mount",
							Computed:            true,
						},
						"local_file": schema.StringAttribute{
							MarkdownDescription: "File name",
							Computed:            true,
						},
						"nfs_file": schema.StringAttribute{
							MarkdownDescription: "NFS file path",
							Computed:            true,
						},
						"archive_mode": schema.StringAttribute{
							MarkdownDescription: "Archive mode",
							Computed:            true,
						},
						"upload_method": schema.StringAttribute{
							MarkdownDescription: "Upload protocol",
							Computed:            true,
						},
						"rotate": schema.Int64Attribute{
							MarkdownDescription: "Number of rotations",
							Computed:            true,
						},
						"use_ansi_color": schema.BoolAttribute{
							MarkdownDescription: "Use ANSI color scheme",
							Computed:            true,
						},
						"remote_address": schema.StringAttribute{
							MarkdownDescription: "Remote address",
							Computed:            true,
						},
						"remote_port": schema.Int64Attribute{
							MarkdownDescription: "Remote port",
							Computed:            true,
						},
						"remote_login": schema.StringAttribute{
							MarkdownDescription: "Remote login",
							Computed:            true,
						},
						"remote_password": schema.StringAttribute{
							MarkdownDescription: "Remote password",
							Computed:            true,
						},
						"remote_directory": schema.StringAttribute{
							MarkdownDescription: "Remote directory",
							Computed:            true,
						},
						"local_address": schema.StringAttribute{
							MarkdownDescription: "Local address",
							Computed:            true,
						},
						"syslog_facility": schema.StringAttribute{
							MarkdownDescription: "syslog facility",
							Computed:            true,
						},
						"rate_limit": schema.Int64Attribute{
							MarkdownDescription: "Rate limit",
							Computed:            true,
						},
						"max_connections": schema.Int64Attribute{
							MarkdownDescription: "Maximum connections",
							Computed:            true,
						},
						"connect_timeout": schema.Int64Attribute{
							MarkdownDescription: "Connect timeout",
							Computed:            true,
						},
						"idle_timeout": schema.Int64Attribute{
							MarkdownDescription: "Idle timeout",
							Computed:            true,
						},
						"active_timeout": schema.Int64Attribute{
							MarkdownDescription: "Active timeout",
							Computed:            true,
						},
						"feedback_detection": schema.BoolAttribute{
							MarkdownDescription: "Feedback detection",
							Computed:            true,
						},
						"log_event_code": schema.ListAttribute{
							MarkdownDescription: "Event subscription filter",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"log_event_filter": schema.ListAttribute{
							MarkdownDescription: "Event suppression filter",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"log_objects": schema.ListNestedAttribute{
							MarkdownDescription: "Object filters",
							NestedObject:        models.DmLogObjectDataSourceSchema,
							Computed:            true,
						},
						"log_ip_filter": schema.ListNestedAttribute{
							MarkdownDescription: "IP address filters",
							NestedObject:        models.DmLogIPFilterDataSourceSchema,
							Computed:            true,
						},
						"log_triggers": schema.ListNestedAttribute{
							MarkdownDescription: "Event triggers",
							NestedObject:        models.DmLogTriggerDataSourceSchema,
							Computed:            true,
						},
						"ssl_client_profile": schema.StringAttribute{
							MarkdownDescription: "TLS client profile",
							Computed:            true,
						},
						"ssl_client_config_type": schema.StringAttribute{
							MarkdownDescription: "TLS client type",
							Computed:            true,
						},
						"retry_interval": schema.Int64Attribute{
							MarkdownDescription: "Retry interval",
							Computed:            true,
						},
						"retry_attempts": schema.Int64Attribute{
							MarkdownDescription: "Retry attempts",
							Computed:            true,
						},
						"long_retry_interval": schema.Int64Attribute{
							MarkdownDescription: "Long retry interval",
							Computed:            true,
						},
						"log_precision": schema.StringAttribute{
							MarkdownDescription: "Log timestamp precision",
							Computed:            true,
						},
						"event_buffer_size": schema.StringAttribute{
							MarkdownDescription: "Buffer size",
							Computed:            true,
						},
						"dependency_actions": actions.ActionsSchema,
					},
				},
			},
		},
	}
}

func (d *LogTargetDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *LogTargetDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data LogTargetList
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.LogTarget{
		AppDomain: data.AppDomain,
	}

	res, err := d.pData.Client.Get(o.GetPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}
	l := []models.LogTarget{}
	if value := res.Get(`LogTarget`); value.Exists() {
		for _, v := range value.Array() {
			item := models.LogTarget{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.LogTargetObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.LogTargetObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

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
	"github.com/scottw514/terraform-provider-datapower/client"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/modifiers"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
)

var _ resource.Resource = &LogTargetResource{}

func NewLogTargetResource() resource.Resource {
	return &LogTargetResource{}
}

type LogTargetResource struct {
	client *client.DatapowerClient
}

func (r *LogTargetResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_logtarget"
}

func (r *LogTargetResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("Log Target", "logging target", "").String,

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
			"backup": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Backup log", "backup", "logtarget").String,
				Optional:            true,
			},
			"log_events": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Event subscriptions", "event", "").String,
				NestedObject:        models.DmLogEventResourceSchema,
				Optional:            true,
			},
			"user_summary": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Comments", "summary", "").String,
				Optional:            true,
			},
			"type": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Target type", "type", "").AddStringEnum("console", "cache", "syslog", "syslog-tcp", "smtp", "file", "soap", "snmp", "nfs").AddDefaultValue("file").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("console", "cache", "syslog", "syslog-tcp", "smtp", "file", "soap", "snmp", "nfs"),
				},
				Default: stringdefault.StaticString("file"),
			},
			"priority": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Logging priority", "priority", "").AddStringEnum("unknown", "high-min", "high", "high-max", "normal-min", "normal", "normal-max", "low-min", "low", "low-max").AddDefaultValue("normal").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("unknown", "high-min", "high", "high-max", "normal-min", "normal", "normal-max", "low-min", "low", "low-max"),
				},
				Default: stringdefault.StaticString("normal"),
			},
			"soap_version": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("SOAP version", "soap-version", "").AddStringEnum("soap11", "soap12").AddDefaultValue("soap11").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("soap11", "soap12"),
				},
				Default: stringdefault.StaticString("soap11"),
			},
			"format": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Log format", "format", "").AddStringEnum("text", "raw", "xml", "json-icp", "cbe", "csv", "audit", "diag").AddDefaultValue("xml").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("text", "raw", "xml", "json-icp", "cbe", "csv", "audit", "diag"),
				},
				Default: stringdefault.StaticString("xml"),
			},
			"timestamp_format": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Timestamp format", "timestamp", "").AddStringEnum("syslog", "numeric", "zulu").AddDefaultValue("zulu").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("syslog", "numeric", "zulu"),
				},
				Default: stringdefault.StaticString("zulu"),
			},
			"fixed_format": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Fixed format", "fixed-format", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"local_identifier": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Local identifier", "local-ident", "").String,
				Optional:            true,
			},
			"email_address": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Recipient email address", "email-address", "").String,
				Optional:            true,
			},
			"sender_address": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Sender email address", "sender-address", "").String,
				Optional:            true,
			},
			"smtp_domain": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("SMTP client domain", "smtp-domain", "").String,
				Optional:            true,
			},
			"size": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Log size", "size", "").AddIntegerRange(100, 50000).AddDefaultValue("500").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(100, 50000),
				},
				Default: int64default.StaticInt64(500),
			},
			"url": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("URL", "url", "").String,
				Optional:            true,
			},
			"nfs_mount": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("NFS static mount", "nfs-static-mount", "nfsstaticmount").String,
				Optional:            true,
			},
			"local_file": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("File name", "local-file", "").String,
				Optional:            true,
			},
			"nfs_file": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("NFS file path", "nfs-file", "").String,
				Optional:            true,
			},
			"archive_mode": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Archive mode", "archive-mode", "").AddStringEnum("rotate", "upload").AddDefaultValue("rotate").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("rotate", "upload"),
				},
				Default: stringdefault.StaticString("rotate"),
			},
			"upload_method": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Upload protocol", "upload-method", "").AddStringEnum("ftp", "scp", "sftp", "smtp").AddDefaultValue("ftp").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("ftp", "scp", "sftp", "smtp"),
				},
				Default: stringdefault.StaticString("ftp"),
			},
			"rotate": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Number of rotations", "rotate", "").AddIntegerRange(1, 100).AddDefaultValue("3").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(1, 100),
				},
				Default: int64default.StaticInt64(3),
			},
			"use_ansi_color": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Use ANSI color scheme", "ansi-color", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"remote_address": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Remote address", "remote-address", "").String,
				Optional:            true,
			},
			"remote_port": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Remote port", "remote-port", "").AddIntegerRange(1, 65535).String,
				Optional:            true,
				Validators: []validator.Int64{

					int64validator.Between(1, 65535),
				},
			},
			"remote_login": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Remote login", "remote-login", "").String,
				Optional:            true,
			},
			"remote_password": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Remote password", "", "").String,
				Optional:            true,
			},
			"remote_directory": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Remote directory", "remote-directory", "").String,
				Optional:            true,
			},
			"local_address": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Local address", "local-address", "").String,
				Optional:            true,
			},
			"syslog_facility": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("syslog facility", "facility", "").AddStringEnum("user", "security", "authpriv", "local0", "local1", "local2", "local3", "local4", "local5", "local6", "local7").AddDefaultValue("user").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("user", "security", "authpriv", "local0", "local1", "local2", "local3", "local4", "local5", "local6", "local7"),
				},
				Default: stringdefault.StaticString("user"),
			},
			"rate_limit": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Rate limit", "rate-limit", "").AddIntegerRange(0, 1000).AddDefaultValue("100").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(0, 1000),
				},
				Default: int64default.StaticInt64(100),
			},
			"max_connections": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Maximum connections", "maximum-connections", "").AddIntegerRange(1, 100).AddDefaultValue("1").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(1, 100),
				},
				Default: int64default.StaticInt64(1),
			},
			"connect_timeout": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Connect timeout", "connect-timeout", "").AddIntegerRange(1, 90).AddDefaultValue("60").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(1, 90),
				},
				Default: int64default.StaticInt64(60),
			},
			"idle_timeout": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Idle timeout", "idle-timeout", "").AddIntegerRange(1, 600).AddDefaultValue("15").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(1, 600),
				},
				Default: int64default.StaticInt64(15),
			},
			"active_timeout": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Active timeout", "active-timeout", "").AddIntegerRange(0, 60).AddDefaultValue("0").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(0, 60),
				},
				Default: int64default.StaticInt64(0),
			},
			"feedback_detection": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Feedback detection", "feedback-detection", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"log_event_code": schema.ListAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Event subscription filter", "event-code", "").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"log_event_filter": schema.ListAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Event suppression filter", "event-filter", "").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"log_objects": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Object filters", "object", "").String,
				NestedObject:        models.DmLogObjectResourceSchema,
				Optional:            true,
			},
			"log_ip_filter": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("IP address filters", "ip-filter", "").String,
				NestedObject:        models.DmLogIPFilterResourceSchema,
				Optional:            true,
			},
			"log_triggers": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Event triggers", "trigger", "").String,
				NestedObject:        models.DmLogTriggerResourceSchema,
				Optional:            true,
			},
			"ssl_client_profile": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("TLS client profile", "ssl-client", "sslclientprofile").String,
				Optional:            true,
			},
			"ssl_client_config_type": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("TLS client type", "ssl-client-type", "").AddStringEnum("client").AddDefaultValue("client").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("client"),
				},
				Default: stringdefault.StaticString("client"),
			},
			"retry_interval": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Retry interval", "retry-interval", "").AddIntegerRange(1, 600).AddDefaultValue("1").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(1, 600),
				},
				Default: int64default.StaticInt64(1),
			},
			"retry_attempts": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Retry attempts", "retry-attempts", "").AddIntegerRange(1, 100).AddDefaultValue("1").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(1, 100),
				},
				Default: int64default.StaticInt64(1),
			},
			"long_retry_interval": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Long retry interval", "long-retry-interval", "").AddIntegerRange(1, 600).AddDefaultValue("20").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(1, 600),
				},
				Default: int64default.StaticInt64(20),
			},
			"log_precision": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Log timestamp precision", "precision", "").AddStringEnum("second", "microsecond").AddDefaultValue("second").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("second", "microsecond"),
				},
				Default: stringdefault.StaticString("second"),
			},
			"event_buffer_size": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Buffer size", "buffer-size", "").AddStringEnum("2048", "16384", "65536", "131072", "262144", "524288", "1048576", "2097152", "4194304", "8388608").AddDefaultValue("2048").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("2048", "16384", "65536", "131072", "262144", "524288", "1048576", "2097152", "4194304", "8388608"),
				},
				Default: stringdefault.StaticString("2048"),
			},
			"object_actions": actions.ActionsSchema,
		},
	}
}

func (r *LogTargetResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = *req.ProviderData.(**client.DatapowerClient)
}

func (r *LogTargetResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data models.LogTarget

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, r.client, data.ObjectActions, actions.Create)

	body := data.ToBody(ctx, `LogTarget`)
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

func (r *LogTargetResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data models.LogTarget

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
		data.FromBody(ctx, `LogTarget`, res)
	} else {
		// Update
		data.UpdateFromBody(ctx, `LogTarget`, res)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LogTargetResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data models.LogTarget

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, r.client, data.ObjectActions, actions.Update)
	_, err := r.client.Put(data.GetPath()+"/"+data.Id.ValueString(), data.ToBody(ctx, `LogTarget`))
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

func (r *LogTargetResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data models.LogTarget

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, r.client, data.ObjectActions, actions.Delete)
	_, err := r.client.Delete(data.GetPath() + "/" + data.Id.ValueString())
	if err != nil && !strings.Contains(err.Error(), "status 404") && !strings.Contains(err.Error(), "status 409") {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object (DELETE), got error: %s", err))
		return
	}

	resp.State.RemoveResource(ctx)
}

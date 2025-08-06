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
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/modifiers"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
)

var _ resource.Resource = &AS3SourceProtocolHandlerResource{}

func NewAS3SourceProtocolHandlerResource() resource.Resource {
	return &AS3SourceProtocolHandlerResource{}
}

type AS3SourceProtocolHandlerResource struct {
	client *client.DatapowerClient
}

func (r *AS3SourceProtocolHandlerResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_as3sourceprotocolhandler"
}

func (r *AS3SourceProtocolHandlerResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("AS3 handler", "source-as3", "").String,

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
			"local_address": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Local IP address", "address", "").AddDefaultValue("0.0.0.0").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("0.0.0.0"),
			},
			"local_port": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Local port", "port", "").AddIntegerRange(1, 65535).AddDefaultValue("21").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(1, 65535),
				},
				Default: int64default.StaticInt64(21),
			},
			"filesystem_type": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("File system type", "filesystem", "").AddStringEnum("virtual-ephemeral", "virtual-persistent", "transparent").AddDefaultValue("virtual-ephemeral").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("virtual-ephemeral", "virtual-persistent", "transparent"),
				},
				Default: stringdefault.StaticString("virtual-ephemeral"),
			},
			"persistent_filesystem_timeout": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Persistent timeout", "persistent-filesystem-timeout", "").AddIntegerRange(1, 43200).AddDefaultValue("600").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(1, 43200),
				},
				Default: int64default.StaticInt64(600),
			},
			"virtual_directories": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Virtual directories", "virtual-directory", "").String,
				NestedObject:        models.DmFTPServerVirtualDirectoryResourceSchema,
				Optional:            true,
			},
			"default_directory": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Default directory", "default-directory", "").AddDefaultValue("/").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("/"),
			},
			"max_filename_length": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Max file name length", "max-filename-len", "").AddIntegerRange(1, 4000).AddDefaultValue("256").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(1, 4000),
				},
				Default: int64default.StaticInt64(256),
			},
			"acl": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Access control list", "acl", "accesscontrollist").String,
				Optional:            true,
			},
			"require_tls": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Require TLS", "require-tls", "").AddStringEnum("off", "explicit", "implicit").AddDefaultValue("off").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("off", "explicit", "implicit"),
				},
				Default: stringdefault.StaticString("off"),
			},
			"password_aaa_policy": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Username-password AAA policy", "password-aaa-policy", "aaapolicy").String,
				Optional:            true,
			},
			"certificate_aaa_policy": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Password-required AAA policy", "certificate-aaa-policy", "aaapolicy").String,
				Optional:            true,
			},
			"allow_ccc": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Allow CCC command", "allow-ccc", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"passive": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Passive (PASV) command", "passive", "").AddStringEnum("allow", "disallow", "require").AddDefaultValue("allow").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("allow", "disallow", "require"),
				},
				Default: stringdefault.StaticString("allow"),
			},
			"use_pasv_port_range": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Limit port range for passive connections", "passive-port-range", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"pasv_min_port": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Min passive port", "passive-port-min", "").AddIntegerRange(1024, 65534).AddDefaultValue("1024").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(1024, 65534),
				},
				Default: int64default.StaticInt64(1024),
			},
			"pasv_max_port": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Max passive port", "passive-port-max", "").AddIntegerRange(1024, 65534).AddDefaultValue("1050").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(1024, 65534),
				},
				Default: int64default.StaticInt64(1050),
			},
			"pasv_idle_time_out": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Passive data connection idle timeout", "passive-idle-timeout", "").AddIntegerRange(5, 300).AddDefaultValue("60").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(5, 300),
				},
				Default: int64default.StaticInt64(60),
			},
			"disable_pasvip_check": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Disable passive data connection (PASV) IP security check", "passive-promiscuous", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"disable_portip_check": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Disable active data connection (PORT) IP security check", "port-promiscuous", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"use_alternate_pasv_addr": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Use alternate PASV IP address", "allow-passive-addr", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"alternate_pasv_addr": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Alternate PASV IP address", "passive-addr", "").String,
				Optional:            true,
			},
			"allow_list_cmd": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enable LIST command support", "list-cmd", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"allow_dele_cmd": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enable DELE command support", "dele-cmd", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"data_encryption": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("File transfer data encryption", "data-encryption", "").AddStringEnum("disallow", "allow", "require").AddDefaultValue("allow").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("disallow", "allow", "require"),
				},
				Default: stringdefault.StaticString("allow"),
			},
			"allow_compression": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Allow compression", "allow-compression", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"allow_stou": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Allow unique file name (STOU)", "allow-unique-filename", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"unique_filename_prefix": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Unique file name prefix", "unique-filename-prefix", "").String,
				Optional:            true,
			},
			"allow_rest": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Allow restart (REST)", "allow-restart", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"restart_timeout": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Restart timeout", "restart-timeout", "").AddDefaultValue("240").String,
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(240),
			},
			"idle_timeout": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Idle timeout", "idle-timeout", "").AddIntegerRange(0, 65535).String,
				Optional:            true,
				Validators: []validator.Int64{

					int64validator.Between(0, 65535),
				},
			},
			"response_type": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Response type", "response-type", "").AddStringEnum("none", "virtual-filesystem").AddDefaultValue("none").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("none", "virtual-filesystem"),
				},
				Default: stringdefault.StaticString("none"),
			},
			"response_storage": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Response storage", "response-storage", "").AddStringEnum("temporary", "nfs").AddDefaultValue("temporary").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("temporary", "nfs"),
				},
				Default: stringdefault.StaticString("temporary"),
			},
			"temporary_storage_size": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Temporary storage size", "filesystem-size", "").AddIntegerRange(1, 2048).AddDefaultValue("32").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(1, 2048),
				},
				Default: int64default.StaticInt64(32),
			},
			"response_nfs_mount": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Response NFS mount", "response-nfs-mount", "nfsstaticmount").String,
				Optional:            true,
			},
			"response_suffix": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Response suffix", "response-suffix", "").String,
				Optional:            true,
			},
			"response_url": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Response URL", "response-url", "").String,
				Optional:            true,
			},
			"ssl_server_config_type": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("TLS server type", "ssl-config-type", "").AddStringEnum("server", "sni").AddDefaultValue("server").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("server", "sni"),
				},
				Default: stringdefault.StaticString("server"),
			},
			"ssl_server": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("TLS server profile", "ssl-server", "sslserverprofile").String,
				Optional:            true,
			},
			"sslsni_server": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("TLS SNI server profile", "ssl-sni-server", "sslsniserverprofile").String,
				Optional:            true,
			},
			"object_actions": actions.ActionsSchema,
		},
	}
}

func (r *AS3SourceProtocolHandlerResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = *req.ProviderData.(**client.DatapowerClient)
}

func (r *AS3SourceProtocolHandlerResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data models.AS3SourceProtocolHandler

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, r.client, data.ObjectActions, actions.Create)

	body := data.ToBody(ctx, `AS3SourceProtocolHandler`)
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

func (r *AS3SourceProtocolHandlerResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data models.AS3SourceProtocolHandler

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
		data.FromBody(ctx, `AS3SourceProtocolHandler`, res)
	} else {
		// Update
		data.UpdateFromBody(ctx, `AS3SourceProtocolHandler`, res)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AS3SourceProtocolHandlerResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data models.AS3SourceProtocolHandler

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, r.client, data.ObjectActions, actions.Update)
	_, err := r.client.Put(data.GetPath()+"/"+data.Id.ValueString(), data.ToBody(ctx, `AS3SourceProtocolHandler`))
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

func (r *AS3SourceProtocolHandlerResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data models.AS3SourceProtocolHandler

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

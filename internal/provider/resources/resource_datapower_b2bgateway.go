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
)

var _ resource.Resource = &B2BGatewayResource{}
var _ resource.ResourceWithValidateConfig = &B2BGatewayResource{}

func NewB2BGatewayResource() resource.Resource {
	return &B2BGatewayResource{}
}

type B2BGatewayResource struct {
	pData *tfutils.ProviderData
}

func (r *B2BGatewayResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_b2bgateway"
}

func (r *B2BGatewayResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("B2B gateway", "b2bgw", "").AddActions("quiesce").String,
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
			"priority": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Service priority", "priority", "").AddStringEnum("unknown", "high-min", "high", "high-max", "normal-min", "normal", "normal-max", "low-min", "low", "low-max").AddDefaultValue("normal").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("unknown", "high-min", "high", "high-max", "normal-min", "normal", "normal-max", "low-min", "low", "low-max"),
				},
				Default: stringdefault.StaticString("normal"),
			},
			"doc_store_location": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Document storage location", "doc-location", "").AddDefaultValue("(default)").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("(default)"),
			},
			"as_front_protocol": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Protocol handlers", "as-fsph", "").String,
				NestedObject:        models.DmASFrontProtocolResourceSchema,
				Optional:            true,
			},
			"as1mdn_email": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Default AS1 MDN return email", "as1-mdn-email", "").String,
				Optional:            true,
			},
			"as1mdnsmtp_server_connection": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("AS1 MDN SMTP server connection", "as1-mdn-smtp-server-connection", "smtpserverconnection").String,
				Optional:            true,
			},
			"as2mdnurl": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Default AS2 MDN Return URL", "as2-mdn-url", "").String,
				Optional:            true,
			},
			"as3mdnurl": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Default AS3 MDN return URL", "as3-mdn-url", "").String,
				Optional:            true,
			},
			"b2b_profiles": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Active partner profiles", "b2b-profile", "").String,
				NestedObject:        models.DmB2BActiveProfileResourceSchema,
				Optional:            true,
			},
			"b2b_groups": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Active profile groups", "b2b-group", "").String,
				NestedObject:        models.DmB2BActiveGroupResourceSchema,
				Optional:            true,
			},
			"document_routing_preprocessor_type": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Processor type", "document-routing-preprocessor-type", "").AddStringEnum("stylesheet", "gatewayscript").AddDefaultValue("stylesheet").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("stylesheet", "gatewayscript"),
				},
				Default: stringdefault.StaticString("stylesheet"),
			},
			"document_routing_preprocessor": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("File location", "document-routing-preprocessor", "").AddDefaultValue("store:///b2b-routing.xsl").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("store:///b2b-routing.xsl"),
			},
			"document_routing_preprocessor_debug": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enable GatewayScript debugger", "document-routing-preprocessor-debug", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"archive_mode": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Purge mode", "arch-mode", "").AddStringEnum("archpurge", "purgeonly").AddDefaultValue("archpurge").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("archpurge", "purgeonly"),
				},
				Default: stringdefault.StaticString("archpurge"),
			},
			"archive_location": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Archive location", "arch-dir", "").String,
				Optional:            true,
			},
			"archive_file_name": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Archive file base name", "arch-file", "").String,
				Optional:            true,
			},
			"archive_minimum_size": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Min size", "arch-minimum-size", "").AddDefaultValue("1024").String,
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(1024),
			},
			"archive_document_age": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Document age", "arch-document-age", "").AddIntegerRange(1, 3650).AddDefaultValue("90").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(1, 3650),
				},
				Default: int64default.StaticInt64(90),
			},
			"archive_minimum_documents": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Min documents", "arch-minimum-documents", "").AddIntegerRange(1, 65535).AddDefaultValue("100").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(1, 65535),
				},
				Default: int64default.StaticInt64(100),
			},
			"disk_use_check_interval": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Check interval", "diskuse-check-interval", "").AddIntegerRange(1, 1440).AddDefaultValue("60").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(1, 1440),
				},
				Default: int64default.StaticInt64(60),
			},
			"max_document_disk_use": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Max document storage", "max-diskuse", "").AddIntegerRange(1, 4294967295).AddDefaultValue("25165824").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(1, 4294967295),
				},
				Default: int64default.StaticInt64(25165824),
			},
			"archive_monitor": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Monitor during archival", "arch-monitor", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"shaping_threshold": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Monitor threshold", "arch-shaping-threshold", "").AddIntegerRange(10, 10000).AddDefaultValue("200").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(10, 10000),
				},
				Default: int64default.StaticInt64(200),
			},
			"archive_backup_documents": models.GetDmB2BBackupMsgTypeResourceSchema("Document types to archive", "arch-backup-documents", "", false),
			"x_path_routing_policies": schema.ListAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("XPath routing policies", "xpath-routing", "b2bxpathroutingpolicy").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"xml_manager": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("XML manager", "xml-manager", "xmlmanager").AddDefaultValue("default").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("default"),
			},
			"debug_mode": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Probe setting", "debug-mode", "").AddStringEnum("on", "off", "unbounded").AddDefaultValue("off").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("on", "off", "unbounded"),
				},
				Default: stringdefault.StaticString("off"),
			},
			"debug_history": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Transaction history", "debug-history", "").AddIntegerRange(10, 250).AddDefaultValue("25").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(10, 250),
				},
				Default: int64default.StaticInt64(25),
			},
			"cpa_entries": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("CPA", "cpa-entry", "").String,
				NestedObject:        models.DmB2BCPAEntryResourceSchema,
				Optional:            true,
			},
			"sql_data_source": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("SQL data source", "sql-source", "sqldatasource").String,
				Optional:            true,
			},
			"front_side_timeout": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Front timeout", "front-side-timeout", "").AddIntegerRange(1, 86400).AddDefaultValue("120").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(1, 86400),
				},
				Default: int64default.StaticInt64(120),
			},
			"dependency_actions": actions.ActionsSchema,
		},
	}
}

func (r *B2BGatewayResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (r *B2BGatewayResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.B2BGateway
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Create, false)
	if resp.Diagnostics.HasError() {
		return
	}

	body := data.ToBody(ctx, `B2BGateway`)
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

func (r *B2BGatewayResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.B2BGateway
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
		data.FromBody(ctx, `B2BGateway`, res)
	} else {
		// Update
		data.UpdateFromBody(ctx, `B2BGateway`, res)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *B2BGatewayResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.B2BGateway
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Update, false)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.pData.Client.Put(data.GetPath()+"/"+data.Id.ValueString(), data.ToBody(ctx, `B2BGateway`))
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

func (r *B2BGatewayResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.B2BGateway
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

func (r *B2BGatewayResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data models.B2BGateway

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.ValidateConfig(ctx, &resp.Diagnostics, data.DependencyActions)
}

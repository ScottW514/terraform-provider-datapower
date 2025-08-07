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

var _ resource.Resource = &TAMResource{}

func NewTAMResource() resource.Resource {
	return &TAMResource{}
}

type TAMResource struct {
	client *client.DatapowerClient
}

func (r *TAMResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_tam"
}

func (r *TAMResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("Access Manager Client", "tam", "").String,
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
			"ad_use_ad": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Use Active Directory", "ad-use-ad", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"tam_version": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Access Manager Client Version", "tam-version", "").AddStringEnum("default", "v70", "v801", "v901", "v903", "v1005").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("default", "v70", "v801", "v901", "v903", "v1005"),
				},
			},
			"configuration_file": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Configuration File for Access Manager", "file", "").String,
				Required:            true,
			},
			"ad_configuration_file": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Configuration File for Directories", "reg_file", "").String,
				Optional:            true,
			},
			"ssl_key_file": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("TLS Key File", "ssl-key", "").String,
				Required:            true,
			},
			"ssl_key_stash_file": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("TLS Key Stash File", "ssl-key-stash", "").String,
				Required:            true,
			},
			"use_local_mode": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Use Local Policy Database", "use-local-mode", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"poll_interval": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Local Database Refresh Interval", "cache-refresh-interval", "").AddDefaultValue("default").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("default"),
			},
			"listen_mode": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Accept Update Notifications", "listen-mode", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"listen_port": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Update Notification Port", "listen-port", "").String,
				Optional:            true,
			},
			"returning_user_attributes": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Return User Attributes", "return-attributes", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"ldap_use_ssl": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Use TLS with Registry Server", "use-ldap-ssl", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"ldapssl_port": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("LDAP TLS Port", "ldap-ssl-port", "").AddDefaultValue("636").String,
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(636),
			},
			"ldapssl_key_file": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Registry Server TLS Key File", "ldap-ssl-key-file", "").String,
				Optional:            true,
			},
			"ldapssl_key_file_password_alias": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Registry Server TLS Key File Password Alias", "ldap-ssl-key-file-password-alias", "passwordalias").String,
				Optional:            true,
			},
			"ldapssl_key_file_label": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Registry Server TLS Key File Label", "ldap-ssl-key-file-dn", "").String,
				Optional:            true,
			},
			"tam_use_fips": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Run in FIPS Mode", "use-fips", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"tam_choose_nist": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Select a NIST Compliance Mode", "choose-nist", "").AddStringEnum("fips", "sp800_131_transition", "sp800_131_strict", "suite_b_128", "suite_b_192").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("fips", "sp800_131_transition", "sp800_131_strict", "suite_b_128", "suite_b_192"),
				},
			},
			"tam_use_basic_user": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enable basic user mode", "use-basic-user", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"user_principal_attribute": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("User principal attribute", "user-principal-attribute", "").AddDefaultValue("uid").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("uid"),
			},
			"user_no_duplicates": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Disallow duplicate principals", "user-no-duplicates", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"user_search_suffixes": schema.ListAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Principal search suffixes", "user-search-suffixes", "").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"user_suffix_optimiser": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enable suffix optimization", "user-suffix-optimiser", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"tam_fed_dirs": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Federated directories", "federated-directory", "").String,
				NestedObject:        models.DmTAMFedDirResourceSchema,
				Optional:            true,
			},
			"tamaz_replicas": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Authorization Server Replicas", "replica", "").String,
				NestedObject:        models.DmTAMAZReplicaResourceSchema,
				Optional:            true,
			},
			"tamras_trace": models.GetDmTAMRASTraceResourceSchema("Trace Logging", "tam-ras-trace", "", false),
			"auto_retry": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enable Automatic Restart Attempts", "auto-retry", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"retry_interval": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Attempt Interval", "retry-interval", "").AddIntegerRange(1, 3600).AddDefaultValue("180").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(1, 3600),
				},
				Default: int64default.StaticInt64(180),
			},
			"retry_attempts": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Number of Attempts", "retry-attempts", "").AddDefaultValue("3").String,
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(3),
			},
			"long_retry_interval": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Long Attempt Interval", "long-retry-interval", "").AddIntegerRange(1, 3600).AddDefaultValue("900").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(1, 3600),
				},
				Default: int64default.StaticInt64(900),
			},
			"dependency_actions": actions.ActionsSchema,
		},
	}
}

func (r *TAMResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = *req.ProviderData.(**client.DatapowerClient)
}

func (r *TAMResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data models.TAM

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, r.client, data.AppDomain.ValueString(), data.DependencyActions, actions.Create)

	body := data.ToBody(ctx, `TAM`)
	_, err := r.client.Post(data.GetPath(), body)

	if err != nil && !strings.Contains(err.Error(), "status 409") {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to create object (%s), got error: %s", "POST", err))
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *TAMResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data models.TAM

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
		data.FromBody(ctx, `TAM`, res)
	} else {
		// Update
		data.UpdateFromBody(ctx, `TAM`, res)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *TAMResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data models.TAM

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, r.client, data.AppDomain.ValueString(), data.DependencyActions, actions.Update)
	_, err := r.client.Put(data.GetPath()+"/"+data.Id.ValueString(), data.ToBody(ctx, `TAM`))
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to update object (PUT), got error: %s", err))
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *TAMResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data models.TAM

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, r.client, data.AppDomain.ValueString(), data.DependencyActions, actions.Delete)
	_, err := r.client.Delete(data.GetPath() + "/" + data.Id.ValueString())
	if err != nil && !strings.Contains(err.Error(), "status 404") && !strings.Contains(err.Error(), "status 409") {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object (DELETE), got error: %s", err))
		return
	}

	resp.State.RemoveResource(ctx)
}

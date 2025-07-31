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

package provider

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
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
)

var _ resource.Resource = &APILDAPRegistryResource{}

func NewAPILDAPRegistryResource() resource.Resource {
	return &APILDAPRegistryResource{}
}

type APILDAPRegistryResource struct {
	client *client.DatapowerClient
}

func (r *APILDAPRegistryResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_apildapregistry"
}

func (r *APILDAPRegistryResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("API LDAP registry", "api-ldap-reg", "").String,

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
					ImmutableAfterSet(),
				},
			},
			"user_summary": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Comments", "summary", "").String,
				Optional:            true,
			},
			"ldap_host": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Host", "ldap-host", "").String,
				Required:            true,
			},
			"ldap_port": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Port", "ldap-port", "").AddDefaultValue("636").String,
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(636),
			},
			"ssl_client_profile": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("TLS client profile", "ssl-client", "sslclientprofile").String,
				Optional:            true,
			},
			"ldap_version": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("LDAP version", "ldap-version", "").AddStringEnum("v2", "v3").AddDefaultValue("v3").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("v2", "v3"),
				},
				Default: stringdefault.StaticString("v3"),
			},
			"ldap_auth_method": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("LDAP authentication method", "ldap-auth-method", "").AddStringEnum("composeDN", "composeUPN", "searchDN").AddDefaultValue("searchDN").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("composeDN", "composeUPN", "searchDN"),
				},
				Default: stringdefault.StaticString("searchDN"),
			},
			"ldap_bind_dn": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("LDAP bind DN", "ldap-bind-dn", "").String,
				Optional:            true,
			},
			"ldap_bind_password_alias": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("LDAP bind password alias", "ldap-bind-password-alias", "passwordalias").String,
				Optional:            true,
			},
			"ldap_search_parameters": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("LDAP search parameters", "ldap-search-param", "ldapsearchparameters").String,
				Required:            true,
			},
			"ldap_read_timeout": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("LDAP read timeout", "ldap-readtimeout", "").AddIntegerRange(0, 86400).AddDefaultValue("60").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(0, 86400),
				},
				Default: int64default.StaticInt64(60),
			},
			"ldap_group_auth_enabled": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enable LDAP group authentication", "ldap-group-auth-enabled", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"ldap_group_auth_type": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("LDAP group authentication type", "ldap-group-auth-type", "").AddStringEnum("dynamic", "static").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("dynamic", "static"),
				},
			},
			"ldap_group_scope": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("LDAP group scope", "ldap-group-scope", "").AddStringEnum("subtree", "one-level", "base").AddDefaultValue("subtree").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("subtree", "one-level", "base"),
				},
				Default: stringdefault.StaticString("subtree"),
			},
			"ldap_group_base_dn": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("LDAP static group base DN", "ldap-group-base-dn", "").String,
				Optional:            true,
			},
			"ldap_group_filter_prefix": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("LDAP static group filter prefix", "ldap-group-filter-prefix", "").String,
				Optional:            true,
			},
			"ldap_group_filter_suffix": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("LDAP static group filter suffix", "ldap-group-filter-suffix", "").String,
				Optional:            true,
			},
			"ldap_group_dynamic_filter": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("LDAP dynamic filter", "ldap-group-dynamic-filter", "").String,
				Optional:            true,
			},
		},
	}
}

func (r *APILDAPRegistryResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = *req.ProviderData.(**client.DatapowerClient)
}

func (r *APILDAPRegistryResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data models.APILDAPRegistry

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	body := data.ToBody(ctx, `APILDAPRegistry`)
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

func (r *APILDAPRegistryResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data models.APILDAPRegistry

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
		data.FromBody(ctx, `APILDAPRegistry`, res)
	} else {
		// Update
		data.UpdateFromBody(ctx, `APILDAPRegistry`, res)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *APILDAPRegistryResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data models.APILDAPRegistry

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.client.Put(data.GetPath()+"/"+data.Id.ValueString(), data.ToBody(ctx, `APILDAPRegistry`))
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

func (r *APILDAPRegistryResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data models.APILDAPRegistry

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

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
	"github.com/scottw514/terraform-provider-datapower/internal/provider/validators"
)

var _ resource.Resource = &APILDAPRegistryResource{}
var _ resource.ResourceWithImportState = &APILDAPRegistryResource{}

func NewAPILDAPRegistryResource() resource.Resource {
	return &APILDAPRegistryResource{}
}

type APILDAPRegistryResource struct {
	pData *tfutils.ProviderData
}

func (r *APILDAPRegistryResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_api_ldap_registry"
}

func (r *APILDAPRegistryResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("Configure and manage the API LDAP registry.", "api-ldap-reg", "").String,
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
			"user_summary": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Comments", "summary", "").String,
				Optional:            true,
			},
			"ldap_host": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Host", "ldap-host", "").String,
				Required:            true,
			},
			"ldap_port": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the listening port on the LDAP server. The default value is 636.", "ldap-port", "").AddDefaultValue("636").String,
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(636),
			},
			"ssl_client_profile": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("TLS client profile", "ssl-client", "ssl_client_profile").String,
				Optional:            true,
			},
			"ldap_version": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the LDAP protocol version for bind operation. The default value is v3.", "ldap-version", "").AddStringEnum("v2", "v3").AddDefaultValue("v3").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("v2", "v3"),
				},
				Default: stringdefault.StaticString("v3"),
			},
			"ldap_auth_method": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the method to create the user for authentication. <ul><li>When compose DN, the DN can be composed from the username. For example, <tt>uid=john,ou=People,dc=company,dc=com</tt> is a DN format that can be composed from the username.</li><li>When compose UPN, the UPN can be composed from the username. For example, <tt>john@example.com</tt> is a UPN format that can be composed from the username.</li><li>When search DN, the DN cannot be composed from the username. You must use an LDAP search to retrieve information that matches the username.</li></ul><p>By default, queries the LDAP server to retrieve user information. Before deciding on the method, contact your LDAP administrator.</p>", "ldap-auth-method", "").AddStringEnum("composeDN", "composeUPN", "searchDN").AddDefaultValue("searchDN").String,
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
				MarkdownDescription: tfutils.NewAttributeDescription("LDAP bind password alias", "ldap-bind-password-alias", "password_alias").String,
				Optional:            true,
			},
			"ldap_search_parameters": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("LDAP search parameters", "ldap-search-param", "ldap_search_parameters").String,
				Required:            true,
			},
			"ldap_read_timeout": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the time to wait for a response from the LDAP server before the connection is closed. Enter a value in the range 0 - 86400. The default value is 60. A value of 0 indicates that the connection never times out.", "ldap-readtimeout", "").AddIntegerRange(0, 86400).AddDefaultValue("60").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 86400),
				},
				Default: int64default.StaticInt64(60),
			},
			"ldap_group_auth_enabled": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to enable LDAP group authentication to use to check group membership for a user. The default value is off.", "ldap-group-auth-enabled", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"ldap_group_auth_type": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the type of group authentication configuration to use. The default value is static.", "ldap-group-auth-type", "").AddStringEnum("dynamic", "static").AddRequiredWhen(models.APILDAPRegistryLDAPGroupAuthTypeCondVal.String()).AddNotValidWhen(models.APILDAPRegistryLDAPGroupAuthTypeIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("dynamic", "static"),
					validators.ConditionalRequiredString(models.APILDAPRegistryLDAPGroupAuthTypeCondVal, models.APILDAPRegistryLDAPGroupAuthTypeIgnoreVal, false),
				},
			},
			"ldap_group_scope": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the depth of the LDAP group search. The default value is subtree.", "ldap-group-scope", "").AddStringEnum("subtree", "one-level", "base").AddDefaultValue("subtree").AddNotValidWhen(models.APILDAPRegistryLDAPGroupScopeIgnoreVal.String()).String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("subtree", "one-level", "base"),
					validators.ConditionalRequiredString(validators.Evaluation{}, models.APILDAPRegistryLDAPGroupScopeIgnoreVal, true),
				},
				Default: stringdefault.StaticString("subtree"),
			},
			"ldap_group_base_dn": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the base DN name to begin the group authentication search. This value identifies the entry level of the tree used by the LDAP group scope.", "ldap-group-base-dn", "").AddRequiredWhen(models.APILDAPRegistryLDAPGroupBaseDNCondVal.String()).AddNotValidWhen(models.APILDAPRegistryLDAPGroupBaseDNIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.APILDAPRegistryLDAPGroupBaseDNCondVal, models.APILDAPRegistryLDAPGroupBaseDNIgnoreVal, false),
				},
			},
			"ldap_group_filter_prefix": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the prefix of the LDAP group filter expression. An LDAP group filter expression is composed by <tt>prefix + user DN + suffix</tt> . <p>When the prefix is <tt>(&amp;(objectclass=group)(member=</tt> and the user DN is <tt>CN=bob,DN=ibm,DN=com</tt> , the LDAP search filter is <tt>(&amp;(objectclass=group)(member=CN=bob,DN=ibm,DN=com))</tt> .</p>", "ldap-group-filter-prefix", "").AddRequiredWhen(models.APILDAPRegistryLDAPGroupFilterPrefixCondVal.String()).AddNotValidWhen(models.APILDAPRegistryLDAPGroupFilterPrefixIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.APILDAPRegistryLDAPGroupFilterPrefixCondVal, models.APILDAPRegistryLDAPGroupFilterPrefixIgnoreVal, false),
				},
			},
			"ldap_group_filter_suffix": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the suffix of the LDAP group filter expression. <p>When the prefix is <tt>&amp;(objectclass=group)(member=</tt> , the user DN is <tt>CN=bob,DN=ibm,DN=com</tt> , and the suffix is <tt>)(CN=ibm-group))</tt> , the LDAP search filter is <tt>(&amp;(objectclass=group)(member=CN=bob,DN=ibm,DN=com)(CN=ibm-group))</tt> .</p>", "ldap-group-filter-suffix", "").AddRequiredWhen(models.APILDAPRegistryLDAPGroupFilterSuffixCondVal.String()).AddNotValidWhen(models.APILDAPRegistryLDAPGroupFilterSuffixIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.APILDAPRegistryLDAPGroupFilterSuffixCondVal, models.APILDAPRegistryLDAPGroupFilterSuffixIgnoreVal, false),
				},
			},
			"ldap_group_dynamic_filter": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the filter expression of the LDAP dynamic group configuration. Only for dynamic. <p>When the filter is <tt>(memberOf=CN=ibm-group,DC=ibm,DC=com)</tt> , the value is used verbatim for LDAP group dynamic search.</p>", "ldap-group-dynamic-filter", "").AddRequiredWhen(models.APILDAPRegistryLDAPGroupDynamicFilterCondVal.String()).AddNotValidWhen(models.APILDAPRegistryLDAPGroupDynamicFilterIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.APILDAPRegistryLDAPGroupDynamicFilterCondVal, models.APILDAPRegistryLDAPGroupDynamicFilterIgnoreVal, false),
				},
			},
			"dependency_actions": actions.ActionsSchema,
		},
	}
}

func (r *APILDAPRegistryResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (r *APILDAPRegistryResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.APILDAPRegistry
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Create, false)
	if resp.Diagnostics.HasError() {
		return
	}

	body := data.ToBody(ctx, `APILDAPRegistry`)
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

func (r *APILDAPRegistryResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.APILDAPRegistry
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

	data.UpdateFromBody(ctx, `APILDAPRegistry`, res)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *APILDAPRegistryResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.APILDAPRegistry
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Update, false)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.pData.Client.Put(data.GetPath()+"/"+data.Id.ValueString(), data.ToBody(ctx, `APILDAPRegistry`))
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

func (r *APILDAPRegistryResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.APILDAPRegistry
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

func (r *APILDAPRegistryResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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

	var data models.APILDAPRegistry
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

	data.FromBody(ctx, `APILDAPRegistry`, res)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *APILDAPRegistryResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data models.APILDAPRegistry

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.ValidateConfig(ctx, &resp.Diagnostics, data.DependencyActions)
}

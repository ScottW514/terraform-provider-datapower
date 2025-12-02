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
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
)

type OAuthSupportedClientGroupList struct {
	ProviderTarget types.String `tfsdk:"provider_target"`
	Id             types.String `tfsdk:"id"`
	AppDomain      types.String `tfsdk:"app_domain"`
	Result         types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &OAuthSupportedClientGroupDataSource{}
	_ datasource.DataSourceWithConfigure = &OAuthSupportedClientGroupDataSource{}
)

func NewOAuthSupportedClientGroupDataSource() datasource.DataSource {
	return &OAuthSupportedClientGroupDataSource{}
}

type OAuthSupportedClientGroupDataSource struct {
	pData *tfutils.ProviderData
}

func (d *OAuthSupportedClientGroupDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_oauth_supported_client_group"
}

func (d *OAuthSupportedClientGroupDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "<p>To support the OAuth 2.0 protocol, an AAA policy requires the configuration of an OAuth client group. An OAuth client group contains the configured OAuth clients that the DataPower Gateway accepts requests from.</p><p>When creating an OAuth client group for an AAA policy, the OAuth client group consists of one or more OAuth clients with the same OAuth roles.</p>",
		Attributes: map[string]schema.Attribute{
			"provider_target": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Target host to retrieve this data from. If not set, provider will use the top level settings.", "", "").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 128),
					stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9_-]+$"), "Must match :"+"^[a-zA-Z0-9_-]+$"),
				},
			},
			"id": schema.StringAttribute{
				MarkdownDescription: "The name of the object to retrieve.",
				Optional:            true,
			},
			"app_domain": schema.StringAttribute{
				MarkdownDescription: "The name of the application domain the object belongs to.",
				Required:            true,
			},
			"result": schema.ListNestedAttribute{
				MarkdownDescription: "List of objects. If `id` was provided and it exists, it will be the only item in the list.",
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
						"user_summary": schema.StringAttribute{
							MarkdownDescription: "Specifies a brief comment that describes the configuration.",
							Computed:            true,
						},
						"customized": schema.BoolAttribute{
							MarkdownDescription: "Indicates whether the configuration is for a customized OAuth client group.",
							Computed:            true,
						},
						"customized_type": schema.StringAttribute{
							MarkdownDescription: "Sets the method to customize an OAuth client.",
							Computed:            true,
						},
						"oauth_role": models.GetDmOAuthRoleDataSourceSchema("Identifies the roles of clients in the group. This property is mutually exclusive to the <b>Customized OAuth</b> property.", "oauth-role", ""),
						"client": schema.ListAttribute{
							MarkdownDescription: "Manages the group of OAuth clients. Use the controls to add or remove clients from the group.",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"template_process_url": schema.StringAttribute{
							MarkdownDescription: "<p>Specifies the location of the stylesheet or GatewayScript file that defines the OAuth client. You can define parts of the configuration parameters in the stylesheet or GatewayScript file and then specify an OAuth client template to derive the remaining information. Note that the stylesheet or GatewayScript file must at least provide the \"client-id\" node. The stylesheet or GatewayScript file must be in the local: or store: directory.</p><p>The stylesheet or GatewayScript file must follow the guidelines when it returns the information: <ul><li>The stylesheet or GatewayScript file must return the &lt;client-id> element.</li><li>If the &lt;customized> element value is set to \"on\", the &lt;customized-process-url> element value must be provided by either the stylesheet or GatewayScript file or the template.</li><li>If the &lt;customized> element value is set to \"on\" in the client template, the &lt;customized-process-url> element cannot unset this value.</li><li>If the &lt;use-validation-url> element value is set to \"on\", the &lt;validation-url> element value must be provided by either the stylesheet or GatewayScript file or the template.</li><li>If the &lt;custom-scope-check> element value is set to \"on\", the &lt;scope-url> element value must be provided by either the stylesheet or GatewayScript file or the template.</li><li>If the &lt;custom-resource-owner> element value is set to \"on\", the &lt;resource-owner-url> element value must be provided by either the stylesheet or GatewayScript file or the template.</li><li>If the &lt;role> element value is set, the value must be the same or a subset of what is defined in the template.</li><li>If the &lt;client-type> element value is set, the value must be the same or a subset of what is defined in the template.</li><li>If the &lt;az-grant> element value is set, the value must be the same or a subset of what is defined in the template.</li><li>If the &lt;az-grant> element value is set to \"+code+\" or \"+token+\", the &lt;local-az-page-url> element value must be provided by either the stylesheet or GatewayScript file or the template.</li><li>If the &lt;caching> element value is set to \"custom\", the &lt;additional-oauth-processing-url> element value must be provided by either the stylesheet or GatewayScript file or the template.</li><li>If the &lt;refresh-token-allowed> is set to a non-zero value, the &lt;refresh-token-lifetime> element value must be provided by either the stylesheet or GatewayScript file or the template.</li><li>If the &lt;check-client-credential> element value is set to \"on\", the &lt;client-authen-method> element value must be provided by either the stylesheet or GatewayScript file or the template.</li><li>If the &lt;client-authen-method> element value is set to \"secret\", the &lt;client-secret> element value must be provided by either the stylesheet or GatewayScript file or the template.</li><li>If the &lt;client-authen-method> element value is set to \"ssl\", the &lt;client-valcred> element value must be provided by either the stylesheet or GatewayScript file or the template.</li></ul></p>",
							Computed:            true,
						},
						"client_template": schema.StringAttribute{
							MarkdownDescription: "<p>Specifies the OAuth client template that is used to derive the configuration parameters that are not specified in the stylesheet or GatewayScript file of the OAuth client.</p>",
							Computed:            true,
						},
						"dependency_actions": actions.ActionsSchema,
						"provider_target": schema.StringAttribute{
							MarkdownDescription: tfutils.NewAttributeDescription("Target host to retrieve this data from. If not set, provider will use the top level settings.", "", "").String,
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *OAuthSupportedClientGroupDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *OAuthSupportedClientGroupDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data OAuthSupportedClientGroupList
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if data.ProviderTarget.ValueString() != "" && !d.pData.Client.ValidTarget(data.ProviderTarget.ValueString()) {
		resp.Diagnostics.AddError("Invalid provider_target", fmt.Sprintf(`Target %q is not defined in the provider's 'targets' block. Available targets: %v`, data.ProviderTarget.ValueString(), d.pData.Client.GetTargetNames()))
		return
	}
	o := models.OAuthSupportedClientGroup{
		AppDomain: data.AppDomain,
	}

	path := o.GetPath()
	if !data.Id.IsNull() {
		path = path + "/" + data.Id.ValueString()
	}

	res, err := d.pData.Client.Get(path, data.ProviderTarget)
	resFound := true
	if err != nil {
		if strings.Contains(err.Error(), "status 401") {
			_ = tfutils.DomainCredentialTest(d.pData.Client, &resp.Diagnostics, data.AppDomain.ValueString(), data.ProviderTarget)
			if !resp.Diagnostics.HasError() {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Application Domain '%s' does not exist", data.AppDomain.ValueString()))
			}
			return
		} else if !strings.Contains(err.Error(), "status 404") {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
			return
		} else {
			resFound = false
		}
	}
	l := []models.OAuthSupportedClientGroup{}
	if resFound {
		if value := res.Get(`OAuthSupportedClientGroup`); value.Exists() {
			for _, v := range value.Array() {
				item := models.OAuthSupportedClientGroup{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.OAuthSupportedClientGroupObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.OAuthSupportedClientGroupObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

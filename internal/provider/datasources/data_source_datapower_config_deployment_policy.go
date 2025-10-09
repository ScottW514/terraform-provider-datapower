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
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
)

type ConfigDeploymentPolicyList struct {
	Id        types.String `tfsdk:"id"`
	AppDomain types.String `tfsdk:"app_domain"`
	Result    types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &ConfigDeploymentPolicyDataSource{}
	_ datasource.DataSourceWithConfigure = &ConfigDeploymentPolicyDataSource{}
)

func NewConfigDeploymentPolicyDataSource() datasource.DataSource {
	return &ConfigDeploymentPolicyDataSource{}
}

type ConfigDeploymentPolicyDataSource struct {
	pData *tfutils.ProviderData
}

func (d *ConfigDeploymentPolicyDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_config_deployment_policy"
}

func (d *ConfigDeploymentPolicyDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "A deployment policy contains a set of rules that are applied during a configuration import. A deployment policy can be used to accept, filter, or modify configuration during import.",
		Attributes: map[string]schema.Attribute{
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
							MarkdownDescription: "Comments",
							Computed:            true,
						},
						"accepted_config": schema.ListAttribute{
							MarkdownDescription: "Matching configuration is accepted during import. To create a match statement, type a correctly formatted resource match in the horizontal text box or select Build. Selecting Build displays the Configuration Match Builder in a popup window. <p>A match statement takes the following general form: <br/><i>Addr</i> / <i>Domain</i> / <i>Resource</i> [? <i>Name=resource-name</i> &amp; <i>Property=property-name</i> &amp; <i>Value=property-value</i> ]</p><table><tr><td valign=\"top\">Addr</td><td>Device Address. Specifies IP address or host alias. The value (*) matches all IP addresses.</td></tr><tr><td valign=\"top\">Domain</td><td>Application Domain. The name of the application domain. The value (*) matches all domains.</td></tr><tr><td valign=\"top\">Resource</td><td>Resource Type. The name of the resource type. The value (*) matches all resource types.</td></tr><tr><td valign=\"top\">Name=resource-name</td><td>An additional specification field, such as \"Name\". Limits the match statement to resources of the specified name. Use a PCRE to select groups of resource instances. For example, \"Name=foo.*\" would match all resources with names that start with \"foo\".</td></tr><tr><td valign=\"top\">Property=property-name</td><td>Property Name. The name of the configuration property. Limits the match statement to resources of the specified property. If change specified, set property-name to null string.</td></tr><tr><td valign=\"top\">Value=property-value</td><td>Property Value. Specifies the value for the configuration property. This property limits the match statement to resources with the specified property value.</td></tr></table>",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"filtered_config": schema.ListAttribute{
							MarkdownDescription: "Matching configuration is filtered during import. Match statements do not control whether to import files that are part of the imported configuration package. Files are imported regardless of whether the filter disallows the resources with which the files are associated. <p>To create a match statement, type a correctly formatted resource match in the horizontal text box or select Build. Selecting Build displays the Configuration Match Builder in a popup window.</p><p>A match statement takes the following general form: <br/><i>Addr</i> / <i>Domain</i> / <i>Resource</i> [? <i>Name=resource-name</i> &amp; <i>Property=property-name</i> &amp; <i>Value=property-value</i> ]</p><table><tr><td valign=\"top\">Addr</td><td>Device Address. Specifies IP address or host alias. The value (*) matches all IP addresses.</td></tr><tr><td valign=\"top\">Domain</td><td>Application Domain. The name of the application domain. The value (*) matches all domains.</td></tr><tr><td valign=\"top\">Resource</td><td>Resource Type. The name of the resource type. The value (*) matches all resource types.</td></tr><tr><td valign=\"top\">Name=resource-name</td><td>An additional specification field, such as \"Name\". Limits the match statement to resources of the specified name. Use a PCRE to select groups of resource instances. For example, \"Name=foo.*\" would match all resources with names that start with \"foo\".</td></tr><tr><td valign=\"top\">Property=property-name</td><td>Property Name. The name of the configuration property. Limits the match statement to resources of the specified property. If change specified, set property-name to null string.</td></tr><tr><td valign=\"top\">Value=property-value</td><td>Property Value. Specifies the value for the configuration property. This property limits the match statement to resources with the specified property value.</td></tr></table>",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"modified_config": schema.ListNestedAttribute{
							MarkdownDescription: "Matching configuration is modified during import. The matching configuration may be changed, added, or deleted. To create a match statement, type a correctly formatted resource match in the horizontal text box or select Build. Selecting Build displays the Configuration Match Builder in a popup window. <p>A match statement takes the following general form: <br/><i>Addr</i> / <i>Domain</i> / <i>Resource</i> [? <i>Name=resource-name</i> &amp; <i>Property=property-name</i> &amp; <i>Value=property-value</i> ]</p><table><tr><td valign=\"top\">Addr</td><td>Device Address. Specifies IP address or host alias. The value (*) matches all IP addresses.</td></tr><tr><td valign=\"top\">Domain</td><td>Application Domain. The name of the application domain. The value (*) matches all domains.</td></tr><tr><td valign=\"top\">Resource</td><td>Resource Type. The name of the resource type. The value (*) matches all resource types.</td></tr><tr><td valign=\"top\">Name=resource-name</td><td>An additional specification field, such as \"Name\". Limits the match statement to resources of the specified name. Use a PCRE to select groups of resource instances. For example, \"Name=foo.*\" would match all resources with names that start with \"foo\".</td></tr><tr><td valign=\"top\">Property=property-name</td><td>Property Name. The name of the configuration property. Limits the match statement to resources of the specified property. If change specified, set property-name to null string.</td></tr><tr><td valign=\"top\">Value=property-value</td><td>Property Value. Specifies the value for the configuration property. This property limits the match statement to resources with the specified property value.</td></tr></table>",
							NestedObject:        models.GetDmConfigModifyTypeDataSourceSchema(),
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *ConfigDeploymentPolicyDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *ConfigDeploymentPolicyDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data ConfigDeploymentPolicyList
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.ConfigDeploymentPolicy{
		AppDomain: data.AppDomain,
	}

	path := o.GetPath()
	if !data.Id.IsNull() {
		path = path + "/" + data.Id.ValueString()
	}

	res, err := d.pData.Client.Get(path)
	resFound := true
	if err != nil {
		if !strings.Contains(err.Error(), "status 404") {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
			return
		} else {
			resFound = false
		}
	}
	l := []models.ConfigDeploymentPolicy{}
	if resFound {
		if value := res.Get(`ConfigDeploymentPolicy`); value.Exists() {
			for _, v := range value.Array() {
				item := models.ConfigDeploymentPolicy{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.ConfigDeploymentPolicyObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.ConfigDeploymentPolicyObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

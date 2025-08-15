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

type NameValueProfileList struct {
	AppDomain types.String `tfsdk:"app_domain"`
	Result    types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &NameValueProfileDataSource{}
	_ datasource.DataSourceWithConfigure = &NameValueProfileDataSource{}
)

func NewNameValueProfileDataSource() datasource.DataSource {
	return &NameValueProfileDataSource{}
}

type NameValueProfileDataSource struct {
	pData *tfutils.ProviderData
}

func (d *NameValueProfileDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_name_value_profile"
}

func (d *NameValueProfileDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Many HTTP things are expressed as name value pairs. These include HTTP headers, cookie values, url-encoded query strings, and url-encoded request messages. This profile provides a mechanism for what kinds of names are expected and for each kind of name what properties should be enforced on the corresponding values. When a name-value pair is not validated successfully that may generate an error, the pair might be stripped from the transaction, or the value may be mapped to another default value.",
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
						"user_summary": schema.StringAttribute{
							MarkdownDescription: "Comments",
							Computed:            true,
						},
						"max_attributes": schema.Int64Attribute{
							MarkdownDescription: "The maximum number of name value pairs allowed in a single entity (header, cookie set, body, and so forth).",
							Computed:            true,
						},
						"max_aggregate_size": schema.Int64Attribute{
							MarkdownDescription: "The lengths of all the names and values in a single entity (header, cookie set, body, query string, and so forth) in aggregate must not exceed this property.",
							Computed:            true,
						},
						"max_name_size": schema.Int64Attribute{
							MarkdownDescription: "The maximum size of a name attribute used in this profile.",
							Computed:            true,
						},
						"max_value_size": schema.Int64Attribute{
							MarkdownDescription: "The maximum size of a value attribute used in this profile.",
							Computed:            true,
						},
						"validation_list": schema.ListNestedAttribute{
							MarkdownDescription: "Each pair submitted to this profile consults this validation list, looking for the first regular expression match of the name against the name expression in the list. When that is found, the corresponding value constraint is matched against the value portion of the name-value pair. If that does not match, the policy applies the 'fixup' attribute to the submitted value. That may result in no change, the pair being removed, an error being generated, or the value being mapped to a known constant.",
							NestedObject:        models.DmValidationTypeDataSourceSchema,
							Computed:            true,
						},
						"default_fixup": schema.StringAttribute{
							MarkdownDescription: "Select the action to taken when no matching entry in the validation list is found. The default is Strip.",
							Computed:            true,
						},
						"default_map_value": schema.StringAttribute{
							MarkdownDescription: "An value that does not have a matching entry in the validation list is changed to this value if the no match policy is 'set'.",
							Computed:            true,
						},
						"default_xss": schema.BoolAttribute{
							MarkdownDescription: "This property allows the value to be checked for Cross Site Scripting (XSS) signatures. These signatures are malicious attempts to input client-side script as the input to a web application. If this client-side script is later displayed in a browser, the script executes and can perform malicious activities. Enable this feature to filter input for malicious content that might get stored and displayed again later, such as the contents of a comment form. The check looks for invalid characters and various forms of the term &lt;script that is often used to engage JavaScript on a browser without the user knowing.",
							Computed:            true,
						},
						"no_match_xss_patterns_file": schema.StringAttribute{
							MarkdownDescription: "Specifies the patterns file that will be used by the XSS filter when No Match XSS is selected. The default file, store:///XSS-Patterns.xml, checks for invalid characters and various forms of the term &lt;script. Specify a custom XML patterns file with PCRE patterns to be used by the XSS filter.",
							Computed:            true,
						},
						"dependency_actions": actions.ActionsSchema,
					},
				},
			},
		},
	}
}

func (d *NameValueProfileDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *NameValueProfileDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data NameValueProfileList
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.NameValueProfile{
		AppDomain: data.AppDomain,
	}

	res, err := d.pData.Client.Get(o.GetPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}
	l := []models.NameValueProfile{}
	if value := res.Get(`NameValueProfile`); value.Exists() {
		for _, v := range value.Array() {
			item := models.NameValueProfile{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.NameValueProfileObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.NameValueProfileObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

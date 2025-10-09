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

type CookieAttributePolicyList struct {
	Id        types.String `tfsdk:"id"`
	AppDomain types.String `tfsdk:"app_domain"`
	Result    types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &CookieAttributePolicyDataSource{}
	_ datasource.DataSourceWithConfigure = &CookieAttributePolicyDataSource{}
)

func NewCookieAttributePolicyDataSource() datasource.DataSource {
	return &CookieAttributePolicyDataSource{}
}

type CookieAttributePolicyDataSource struct {
	pData *tfutils.ProviderData
}

func (d *CookieAttributePolicyDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cookie_attribute_policy"
}

func (d *CookieAttributePolicyDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Cookie attribute policy manages pre-defined and custom attributes of a cookie.",
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
							MarkdownDescription: "A descriptive summary for the configuration.",
							Computed:            true,
						},
						"cookie_attribute": models.GetDmCookieAttributeDataSourceSchema("Specifies the attributes to include in the cookie. Acceptable values include domain, path, secure, httponly, max-age, expires, and custom. When an attribute is enabled, you can enter its value in the corresponding property and the enabled attribute is included in the cookie.", "cookie-attribute", ""),
						"domain": schema.StringAttribute{
							MarkdownDescription: "Identifies domain to which a cookie belongs. A browser accepts cookies only when the current domain matches the value you enter here. The maximum length of the domain is 256 characters.",
							Computed:            true,
						},
						"path": schema.StringAttribute{
							MarkdownDescription: "Identifies path attribute of a cookie. A browser accepts cookies only when the current path matches the value you enter there. If this policy object is attached to HTML Forms Login Policy, this property overrides Form POST Action URL property. The maximum length of the path is 256 characters.",
							Computed:            true,
						},
						"interval": schema.Int64Attribute{
							MarkdownDescription: "<p>Sets the cookie's maximum age and/or the cookie's expiration date as an interval of seconds, relative to the time the transaction occurred on the object. For example, if this value is set to 3600 and the transaction on this object occurred on Feb 10, 2014 12:00:00 GMT, then the maximum age of the cookie is 3600 seconds and the expiration date is Feb 10, 2014 13:00:00 GMT, depending on whether the Max-Age and the Expires attribute are included.</p><p>When the maximum age or the expiration date is reached, the cookie is deleted. Enter a value in the range 1 - 2678400. The default value is 3600. Note that the Max-Age attribute in this policy overrides Inactivity Timeout and Session Lifetime attributes in HTML Forms Login policy.</p>",
							Computed:            true,
						},
						"custom_attribute": schema.StringAttribute{
							MarkdownDescription: "The additional attributes to include in the cookie. Enter each attribute in name-value pair. When you enter multiple pairs, use a semicolon (;) to separate them. A name-value pair with an empty value (name-only portion) can also be specified here. You can use variables instead of name-value pair(s). Enter a context variable as var://variablename",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *CookieAttributePolicyDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *CookieAttributePolicyDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data CookieAttributePolicyList
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.CookieAttributePolicy{
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
	l := []models.CookieAttributePolicy{}
	if resFound {
		if value := res.Get(`CookieAttributePolicy`); value.Exists() {
			for _, v := range value.Array() {
				item := models.CookieAttributePolicy{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.CookieAttributePolicyObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.CookieAttributePolicyObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

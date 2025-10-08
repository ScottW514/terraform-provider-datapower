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
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
)

type APISecurityHTTPSchemeList struct {
	Id        types.String `tfsdk:"id"`
	AppDomain types.String `tfsdk:"app_domain"`
	Result    types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &APISecurityHTTPSchemeDataSource{}
	_ datasource.DataSourceWithConfigure = &APISecurityHTTPSchemeDataSource{}
)

func NewAPISecurityHTTPSchemeDataSource() datasource.DataSource {
	return &APISecurityHTTPSchemeDataSource{}
}

type APISecurityHTTPSchemeDataSource struct {
	pData *tfutils.ProviderData
}

func (d *APISecurityHTTPSchemeDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_api_security_http_scheme"
}

func (d *APISecurityHTTPSchemeDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Configure and manage an HTTP bearer authentication security definition with the Bearer authentication security scheme in an API as defined in OpenAPI 3.0. Bearer authentication is an HTTP Authentication scheme and was originally part of RFC 6750: The OAuth 2.0 Authorization Framework: Bearer Token Usage.",
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
						"scheme": schema.StringAttribute{
							MarkdownDescription: "Specify the HTTP authentication security scheme. The only supported value is <tt>bearer</tt> .",
							Computed:            true,
						},
						"bearer_format": schema.StringAttribute{
							MarkdownDescription: "Specify how the bearer token is formatted. The format is primarily a hint.",
							Computed:            true,
						},
						"bearer_validation_method": schema.StringAttribute{
							MarkdownDescription: "Bearer validation method",
							Computed:            true,
						},
						"bearer_validation_endpoint": schema.StringAttribute{
							MarkdownDescription: "Specify the URL of the validation endpoint. When this connection uses the HTTPS protocol, specify the TLS client profile to secure the connection.",
							Computed:            true,
						},
						"bearer_validation_tls_profile": schema.StringAttribute{
							MarkdownDescription: "Specify the TLS client profile to secure the connection to the validation endpoint. You must specify a client profile when you connect to the validation endpoint with the HTTPS protocol.",
							Computed:            true,
						},
						"dependency_actions": actions.ActionsSchema,
					},
				},
			},
		},
	}
}

func (d *APISecurityHTTPSchemeDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *APISecurityHTTPSchemeDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data APISecurityHTTPSchemeList
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.APISecurityHTTPScheme{
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
	l := []models.APISecurityHTTPScheme{}
	if resFound {
		if value := res.Get(`APISecurityHTTPScheme`); value.Exists() {
			for _, v := range value.Array() {
				item := models.APISecurityHTTPScheme{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.APISecurityHTTPSchemeObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.APISecurityHTTPSchemeObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

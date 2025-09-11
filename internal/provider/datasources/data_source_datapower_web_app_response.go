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

type WebAppResponseList struct {
	Id        types.String `tfsdk:"id"`
	AppDomain types.String `tfsdk:"app_domain"`
	Result    types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &WebAppResponseDataSource{}
	_ datasource.DataSourceWithConfigure = &WebAppResponseDataSource{}
)

func NewWebAppResponseDataSource() datasource.DataSource {
	return &WebAppResponseDataSource{}
}

type WebAppResponseDataSource struct {
	pData *tfutils.ProviderData
}

func (d *WebAppResponseDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_web_app_response"
}

func (d *WebAppResponseDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "The web response profile specifies various properties about the response side of the transaction that must be satisfied.",
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
						"policy_type": schema.StringAttribute{
							MarkdownDescription: "Select the satisfaction policy for the profile. The default is Admission.",
							Computed:            true,
						},
						"ok_codes":    models.GetDmHTTPResponseCodesDataSourceSchema("Specify which response codes from the backend server are acceptable.", "response-codes", ""),
						"ok_versions": models.GetDmHTTPVersionMaskDataSourceSchema("Specify which HTTP versions are acceptable from the backend server.", "response-versions", ""),
						"min_body_size": schema.Int64Attribute{
							MarkdownDescription: "Specify the minimum size of the response body.",
							Computed:            true,
						},
						"max_body_size": schema.Int64Attribute{
							MarkdownDescription: "Specify the maximum size of the response body.",
							Computed:            true,
						},
						"header_gnvc": schema.StringAttribute{
							MarkdownDescription: "The validation profile allows you to specify what headers are expected, what headers should be stripped, and what headers should be mapped to known values. If no profile is specified, any header is allowed.",
							Computed:            true,
						},
						"content_types": schema.ListAttribute{
							MarkdownDescription: "A list of PCRE regular expressions that indicate acceptable content-type MIME headers on the response. If this list is empty, any content-type is acceptable. If the response does not have a content type that will be represented as an empty string for matching purposes. An empty list will match all content types.",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"xml_policy": schema.StringAttribute{
							MarkdownDescription: "Specify how the device handles responses that contain an XML MIME type.",
							Computed:            true,
						},
						"xml_rule": schema.StringAttribute{
							MarkdownDescription: "This is the transformation rule that is run when the response contains an XML MIME type and the XML processing policy is set to XML or SOAP.",
							Computed:            true,
						},
						"non_xml_policy": schema.StringAttribute{
							MarkdownDescription: "Specify how the device handles responses that do not contain an XML MIME type.",
							Computed:            true,
						},
						"non_xml_rule": schema.StringAttribute{
							MarkdownDescription: "This is the transformation rule that is run when the response does not contain an XML MIME type and the Non-XML processing policy is set to binary or side-effect.",
							Computed:            true,
						},
						"error_policy": schema.StringAttribute{
							MarkdownDescription: "If this response policy is violated the firewall error policy will be invoked unless this more specific error policy is provided, in which case this policy takes precedence.",
							Computed:            true,
						},
						"dependency_actions": actions.ActionsSchema,
					},
				},
			},
		},
	}
}

func (d *WebAppResponseDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *WebAppResponseDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data WebAppResponseList
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.WebAppResponse{
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
	l := []models.WebAppResponse{}
	if resFound {
		if value := res.Get(`WebAppResponse`); value.Exists() {
			for _, v := range value.Array() {
				item := models.WebAppResponse{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.WebAppResponseObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.WebAppResponseObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

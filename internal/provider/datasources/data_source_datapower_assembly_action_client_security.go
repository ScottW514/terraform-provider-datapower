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

type AssemblyActionClientSecurityList struct {
	Id        types.String `tfsdk:"id"`
	AppDomain types.String `tfsdk:"app_domain"`
	Result    types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &AssemblyActionClientSecurityDataSource{}
	_ datasource.DataSourceWithConfigure = &AssemblyActionClientSecurityDataSource{}
)

func NewAssemblyActionClientSecurityDataSource() datasource.DataSource {
	return &AssemblyActionClientSecurityDataSource{}
}

type AssemblyActionClientSecurityDataSource struct {
	pData *tfutils.ProviderData
}

func (d *AssemblyActionClientSecurityDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_assembly_action_client_security"
}

func (d *AssemblyActionClientSecurityDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Client security assembly action",
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
						"stop_on_error": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to stop processing if client security fails. If failed, stops the assembly and return an error.",
							Computed:            true,
						},
						"secret_required": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to require the client secret. When required, the secret is compared to the registered secret on the application that is identified by the client ID.",
							Computed:            true,
						},
						"extract_credential_method": schema.StringAttribute{
							MarkdownDescription: "<p>Specify the method to extract client credentials from the request.</p><ul><li>For all methods except HTTP, use the ID name and the secret name to specify the locations that contains the ID and the location that contain the secret. <ul><li>When cookie, specify which cookie.</li><li>When context variable, specify which runtime context variable.</li><li>When form data, specify the form data.</li><li>When header, specify which header.</li><li>When query parameter, specify which query parameter.</li></ul></li><li>For the HTTP method, use the HTTP type to specify the format of the <tt>Authorization</tt> header, which expects the basic form in the <tt>Basic <i>base64_id:secret</i></tt> format.</li></ul>",
							Computed:            true,
						},
						"id_name": schema.StringAttribute{
							MarkdownDescription: "<p>Specify the location where to find the client ID to extract.</p><ul><li>When cookie, specify which cookie.</li><li>When context variable, specify which runtime context variable.</li><li>When form data, specify the form data.</li><li>When header, specify which header.</li><li>When query parameter, specify which query parameter.</li></ul>",
							Computed:            true,
						},
						"secret_name": schema.StringAttribute{
							MarkdownDescription: "<p>Specify the location where to find the secret to extract.</p><ul><li>When cookie, specify which cookie.</li><li>When context variable, specify which runtime context variable.</li><li>When form data, specify the form data.</li><li>When header, specify which header.</li><li>When query parameter, specify which query parameter.</li></ul>",
							Computed:            true,
						},
						"http_type": schema.StringAttribute{
							MarkdownDescription: "HTTP type",
							Computed:            true,
						},
						"authenticate_client_method": schema.StringAttribute{
							MarkdownDescription: "Specify the method to authenticate the extracted client credentials. When third-party, specify the user-registry to authenticate the extracted client credentials.",
							Computed:            true,
						},
						"user_registry": schema.StringAttribute{
							MarkdownDescription: "Specify the API registry to authenticate the extracted client credentials. The supported registries are API authentication URL and API LDAP.",
							Computed:            true,
						},
						"user_summary": schema.StringAttribute{
							MarkdownDescription: "Comments",
							Computed:            true,
						},
						"title": schema.StringAttribute{
							MarkdownDescription: "Title",
							Computed:            true,
						},
						"correlation_path": schema.StringAttribute{
							MarkdownDescription: "Specify the path that correlates the API action to a specific part of the API specification. The correlation path specifies the part of the API definition that correlates with the API action. This path is exposed in the debug data by the API gateway for use by debugging tools. For example, for an API configuration that is retrieved from API Connect and specified in an OpenAPI document with IBM extensions, this path is the JSON path to the assembly policy in the IBM extensions section of the document. The path can be expressed in any form that the debugging tool can correlate to the API definition.",
							Computed:            true,
						},
						"action_debug": schema.BoolAttribute{
							MarkdownDescription: "<p>Specify whether to enable the GatewayScript debugger to troubleshoot the following GatewayScript files or script.</p><ul><li>Troubleshoot a GatewayScript file that is called from the GatewayScript assembly action.</li><li>Troubleshoot a GatewayScript file that is called from an XSLT assembly action that uses the <tt>gatewayscript()</tt> extension function.</li><li>Troubleshoot a GatewayScript script that is called through the <tt>value</tt> or <tt>default</tt> property in the JSON file from the map assembly action.</li></ul><p>To debug a file or script, the following conditions must be met.</p><ul><li>The file contains one or more <tt>debugger;</tt> statements at the points in your script where you want to start debugging.</li><li>The GatewayScript debugger is enabled.</li></ul><p>You run the <tt>debug-action</tt> command.</p>",
							Computed:            true,
						},
						"dependency_actions": actions.ActionsSchema,
					},
				},
			},
		},
	}
}

func (d *AssemblyActionClientSecurityDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *AssemblyActionClientSecurityDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data AssemblyActionClientSecurityList
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.AssemblyActionClientSecurity{
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
	l := []models.AssemblyActionClientSecurity{}
	if resFound {
		if value := res.Get(`AssemblyActionClientSecurity`); value.Exists() {
			for _, v := range value.Array() {
				item := models.AssemblyActionClientSecurity{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.AssemblyActionClientSecurityObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.AssemblyActionClientSecurityObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

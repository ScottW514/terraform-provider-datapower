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

type ImportPackageList struct {
	Id        types.String `tfsdk:"id"`
	AppDomain types.String `tfsdk:"app_domain"`
	Result    types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &ImportPackageDataSource{}
	_ datasource.DataSourceWithConfigure = &ImportPackageDataSource{}
)

func NewImportPackageDataSource() datasource.DataSource {
	return &ImportPackageDataSource{}
}

type ImportPackageDataSource struct {
	pData *tfutils.ProviderData
}

func (d *ImportPackageDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_import_package"
}

func (d *ImportPackageDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "An import package allows the system to import a configuration package from an external server, such as a centralized configuration server, when the configuration is reloaded (such as during a restart). This package can optionally overwrite existing files or objects. <p>An import package specifies a source, content type, and import parameters for a single bundle.</p>",
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
						"url": schema.StringAttribute{
							MarkdownDescription: "Specify the URL of the import package. The tool does not support SCP and SFTP protocols. All other URL protocols are available; for example, HTTP, HTTPS, or FTP.",
							Computed:            true,
						},
						"import_format": schema.StringAttribute{
							MarkdownDescription: "Specify the format of the import package. The default value is ZIP.",
							Computed:            true,
						},
						"overwrite_files": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to overwrite existing files. When enabled, files in the package overwrite any file of the same path and name that already exist. The default behavior is to overwrite files.",
							Computed:            true,
						},
						"overwrite_objects": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to overwrite existing objects. When enabled, objects in the package overwrite any object of the same class and name that already exist. The default behavior is to overwrite objects.",
							Computed:            true,
						},
						"destination_domain": schema.StringAttribute{
							MarkdownDescription: "Destination domain",
							Computed:            true,
						},
						"deployment_policy": schema.StringAttribute{
							MarkdownDescription: "Deployment policy",
							Computed:            true,
						},
						"deployment_policy_parameters": schema.StringAttribute{
							MarkdownDescription: "Deployment policy variables",
							Computed:            true,
						},
						"local_ip_rewrite": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to rewrite local IP addresses on import. When enabled, local IP addresses in the import package are rewritten to match the local configuration on import. In other words, a service bound to eth1 in the import package are rewritten to bind to eth1 on the local system on import. The default behavior is to rewrite IP addresses.",
							Computed:            true,
						},
						"on_startup": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to import the import package on startup. The default behavior is to import on startup. <ul><li>When enabled, the import package is imported at startup. The configuration is marked external and cannot be saved locally. This setting is equivalent to 'import-always'.</li><li>When disabled, the import must be started manually. The configuration is not marked external and can be saved locally. This setting is equivalent to 'import-once'.</li></ul>",
							Computed:            true,
						},
						"dependency_actions": actions.ActionsSchema,
					},
				},
			},
		},
	}
}

func (d *ImportPackageDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *ImportPackageDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data ImportPackageList
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.ImportPackage{
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
	l := []models.ImportPackage{}
	if resFound {
		if value := res.Get(`ImportPackage`); value.Exists() {
			for _, v := range value.Array() {
				item := models.ImportPackage{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.ImportPackageObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.ImportPackageObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

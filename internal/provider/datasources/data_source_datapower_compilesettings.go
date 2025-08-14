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

type CompileSettingsList struct {
	AppDomain types.String `tfsdk:"app_domain"`
	Result    types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &CompileSettingsDataSource{}
	_ datasource.DataSourceWithConfigure = &CompileSettingsDataSource{}
)

func NewCompileSettingsDataSource() datasource.DataSource {
	return &CompileSettingsDataSource{}
}

type CompileSettingsDataSource struct {
	pData *tfutils.ProviderData
}

func (d *CompileSettingsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_compilesettings"
}

func (d *CompileSettingsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Compile Settings",
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
						"xslt_version": schema.StringAttribute{
							MarkdownDescription: "XSLT version",
							Computed:            true,
						},
						"strict": schema.BoolAttribute{
							MarkdownDescription: "Strict",
							Computed:            true,
						},
						"profile": schema.BoolAttribute{
							MarkdownDescription: "Profile rule",
							Computed:            true,
						},
						"debug": schema.BoolAttribute{
							MarkdownDescription: "Debug rule",
							Computed:            true,
						},
						"stream": schema.BoolAttribute{
							MarkdownDescription: "Streaming rule",
							Computed:            true,
						},
						"try_stream": schema.BoolAttribute{
							MarkdownDescription: "Attempt streaming rule",
							Computed:            true,
						},
						"minimum_escaping": schema.BoolAttribute{
							MarkdownDescription: "Minimum output escaping rule",
							Computed:            true,
						},
						"stack_size": schema.Int64Attribute{
							MarkdownDescription: "Maximum stack size",
							Computed:            true,
						},
						"wsi_validation": schema.StringAttribute{
							MarkdownDescription: "WS-I Basic Profile validation",
							Computed:            true,
						},
						"wsdl_validate_body": schema.StringAttribute{
							MarkdownDescription: "Validate message body",
							Computed:            true,
						},
						"wsdl_validate_headers": schema.StringAttribute{
							MarkdownDescription: "Validate message headers",
							Computed:            true,
						},
						"wsdl_validate_faults": schema.StringAttribute{
							MarkdownDescription: "Validate message fault details",
							Computed:            true,
						},
						"wsdl_wrapped_faults": schema.BoolAttribute{
							MarkdownDescription: "Require wrappers on fault details specified by type",
							Computed:            true,
						},
						"allow_soap_enc_array": schema.BoolAttribute{
							MarkdownDescription: "Specifically allow xsi:type='SOAP-ENC:Array' rule",
							Computed:            true,
						},
						"validate_soap_enc_array": schema.BoolAttribute{
							MarkdownDescription: "Validate SOAP 1.1 encoding rule",
							Computed:            true,
						},
						"wildcards_ignore_xsi_type": schema.BoolAttribute{
							MarkdownDescription: "Wildcards ignore xsi:type rule",
							Computed:            true,
						},
						"wsdl_strict_soap_version": schema.BoolAttribute{
							MarkdownDescription: "Strict SOAP envelope version",
							Computed:            true,
						},
						"xacml_debug": schema.BoolAttribute{
							MarkdownDescription: "Debug XACML policy",
							Computed:            true,
						},
						"allow_xop_include": schema.BoolAttribute{
							MarkdownDescription: "Accept MTOM/XOP optimized messages",
							Computed:            true,
						},
						"dependency_actions": actions.ActionsSchema,
					},
				},
			},
		},
	}
}

func (d *CompileSettingsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *CompileSettingsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data CompileSettingsList
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.CompileSettings{
		AppDomain: data.AppDomain,
	}

	res, err := d.pData.Client.Get(o.GetPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}
	l := []models.CompileSettings{}
	if value := res.Get(`CompileSettings`); value.Exists() {
		for _, v := range value.Array() {
			item := models.CompileSettings{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.CompileSettingsObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.CompileSettingsObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

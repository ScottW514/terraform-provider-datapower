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

type CompileOptionsPolicyList struct {
	AppDomain types.String `tfsdk:"app_domain"`
	Result    types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &CompileOptionsPolicyDataSource{}
	_ datasource.DataSourceWithConfigure = &CompileOptionsPolicyDataSource{}
)

func NewCompileOptionsPolicyDataSource() datasource.DataSource {
	return &CompileOptionsPolicyDataSource{}
}

type CompileOptionsPolicyDataSource struct {
	pData *tfutils.ProviderData
}

func (d *CompileOptionsPolicyDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_compileoptionspolicy"
}

func (d *CompileOptionsPolicyDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Compile Options Policy",
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
							MarkdownDescription: "XSLT Version",
							Computed:            true,
						},
						"strict": schema.BoolAttribute{
							MarkdownDescription: "Strict",
							Computed:            true,
						},
						"profile": schema.StringAttribute{
							MarkdownDescription: "Profile Rule",
							Computed:            true,
						},
						"debug": schema.StringAttribute{
							MarkdownDescription: "Debug Rule",
							Computed:            true,
						},
						"stream": schema.StringAttribute{
							MarkdownDescription: "Streaming Rule",
							Computed:            true,
						},
						"try_stream": schema.StringAttribute{
							MarkdownDescription: "Attempt Streaming Rule",
							Computed:            true,
						},
						"minimum_escaping": schema.StringAttribute{
							MarkdownDescription: "Minimum Output Escaping Rule",
							Computed:            true,
						},
						"stack_size": schema.Int64Attribute{
							MarkdownDescription: "Maximum Stack Size",
							Computed:            true,
						},
						"prefer_xg4": schema.StringAttribute{
							MarkdownDescription: "XML Hardware Acceleration Preferred Rule",
							Computed:            true,
						},
						"disallow_xg4": schema.StringAttribute{
							MarkdownDescription: "XML Hardware Acceleration Disallowed Rule",
							Computed:            true,
						},
						"wsi_validation": schema.StringAttribute{
							MarkdownDescription: "WS-I Basic Profile Validation",
							Computed:            true,
						},
						"wsdl_validate_body": schema.StringAttribute{
							MarkdownDescription: "Validate Message Body",
							Computed:            true,
						},
						"wsdl_validate_headers": schema.StringAttribute{
							MarkdownDescription: "Validate Message Headers",
							Computed:            true,
						},
						"wsdl_validate_faults": schema.StringAttribute{
							MarkdownDescription: "Validate Message Fault details",
							Computed:            true,
						},
						"wsdl_wrapped_faults": schema.BoolAttribute{
							MarkdownDescription: "Require wrappers on fault-details specified by type",
							Computed:            true,
						},
						"allow_soap_enc_array": schema.StringAttribute{
							MarkdownDescription: "Specifically Allow xsi:type='SOAP-ENC:Array' Rule",
							Computed:            true,
						},
						"validate_soap_enc_array": schema.StringAttribute{
							MarkdownDescription: "Validate SOAP 1.1 Encoding Rule",
							Computed:            true,
						},
						"wildcards_ignore_xsi_type": schema.StringAttribute{
							MarkdownDescription: "Wildcards Ignore xsi:type Rule",
							Computed:            true,
						},
						"wsdl_strict_soap_version": schema.BoolAttribute{
							MarkdownDescription: "Strict SOAP Envelope Version",
							Computed:            true,
						},
						"xacml_debug": schema.BoolAttribute{
							MarkdownDescription: "Debug XACML Policy",
							Computed:            true,
						},
						"allow_xop_include": schema.StringAttribute{
							MarkdownDescription: "Accept MTOM/XOP Optimized Messages",
							Computed:            true,
						},
						"dependency_actions": actions.ActionsSchema,
					},
				},
			},
		},
	}
}

func (d *CompileOptionsPolicyDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *CompileOptionsPolicyDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data CompileOptionsPolicyList
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.CompileOptionsPolicy{
		AppDomain: data.AppDomain,
	}

	res, err := d.pData.Client.Get(o.GetPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}
	l := []models.CompileOptionsPolicy{}
	if value := res.Get(`CompileOptionsPolicy`); value.Exists() {
		for _, v := range value.Array() {
			item := models.CompileOptionsPolicy{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.CompileOptionsPolicyObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.CompileOptionsPolicyObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

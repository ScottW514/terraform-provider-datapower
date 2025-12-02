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

type ParseSettingsList struct {
	ProviderTarget types.String `tfsdk:"provider_target"`
	Id             types.String `tfsdk:"id"`
	AppDomain      types.String `tfsdk:"app_domain"`
	Result         types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &ParseSettingsDataSource{}
	_ datasource.DataSourceWithConfigure = &ParseSettingsDataSource{}
)

func NewParseSettingsDataSource() datasource.DataSource {
	return &ParseSettingsDataSource{}
}

type ParseSettingsDataSource struct {
	pData *tfutils.ProviderData
}

func (d *ParseSettingsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_parse_settings"
}

func (d *ParseSettingsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Parse settings define the constraints to parsie documents. These constraints overwrite the parser limits in the XML manager. These settings provide threat protection.",
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
							MarkdownDescription: "Comments",
							Computed:            true,
						},
						"document_type": schema.StringAttribute{
							MarkdownDescription: "Specify the type of document to parse. By default, the document type is automatically detected.",
							Computed:            true,
						},
						"document_size": schema.Int64Attribute{
							MarkdownDescription: "Specify the maximum document size in bytes. A document is rejected when its size exceeds the maximum size. Enter a value in the range 0 - 5368709121. The default value is 4194304. A value of 0 indicates unlimited. When 0, the action does not return the document size. This setting is applicable to binary, JSON, XML, and GraphQL documents.",
							Computed:            true,
						},
						"nesting_depth": schema.Int64Attribute{
							MarkdownDescription: "Specify the maximum level of nested element depth. A document is rejected when its nesting depth exceeds the maximum depth. Enter a value in the range 0 - 4096. The default value is 512. A value of 0 indicates unlimited. When 0, the action does not return the nesting depth. This setting is applicable to XML, JSON, and GraphQL documents. <ul><li>For XML, applies to the maximum level of element depth.</li><li>For JSON, applies to the maximum number of nested label-value pairs, the maximum number of nested arrays, or the maximum number of combination of label-value pairs and arrays.</li><li>For GraphQL, applies to the maximum level of nested selection sets.</li></ul>",
							Computed:            true,
						},
						"width": schema.Int64Attribute{
							MarkdownDescription: "Specify the maximum width of the payload. A document is rejected when its width exceeds the maximum width. Enter a value in the range 0 - 65535. The default value is 4096. A value of 0 indicates unlimited. When 0, the action does not return the width. This setting is application to XML, JSON, and GraphQL documents. <ul><li>For XML applies to the maximum number of attributes on an element and the maximum number of child elements for an element.</li><li>For JSON, applies to the maximum number of properties on a JSON object and the maximum number of JSON items in a JSON array</li><li>For GraphQL, applies to the maximum number of selections in a selection set.</li></ul>",
							Computed:            true,
						},
						"name_length": schema.Int64Attribute{
							MarkdownDescription: "Specify the maximum name length in bytes. A document is rejected when its name length exceeds the maximum length. Enter a value in the range 0 - 8192. The default value is 256. A value of 0 indicates unlimited. When 0, the action does not return the name length. This setting is applicable to XML, JSON, and GraphQL documents. <ul><li>For XML, applies to the length of the name portion of a tag.</li><li>For JSON, applies to the length of the label portion of the JSON label-value pair.</li><li>For GraphQL, applies to the length of the identifiers, which includes field and directive names.</li></ul><p>The length includes any white space between tags in XML or quotation marks in JSON.</p>",
							Computed:            true,
						},
						"value_length": schema.Int64Attribute{
							MarkdownDescription: "Specify the maximum value length in bytes. A document is rejected when its value length exceeds the maximum length. Enter a value in the range 0 - 5368709121. The default value is 8192. A value of 0 indicates unlimited. When 0, the action does not return the value length. This setting is applicable to XML, JSON, and GraphQL documents. <ul><li>For XML, applies to the length of an attribute or text value.</li><li>For JSON, applies to the length of a string value.</li><li>For GraphQL, applies to the number of bytes in any string from which GraphQL is composed.</li></ul><p>The length includes any white space between tags in XML or quotation marks in JSON.</p>",
							Computed:            true,
						},
						"unique_prefixes": schema.Int64Attribute{
							MarkdownDescription: "Specify the maximum number of unique XML namespace prefixes. This count includes multiple prefixes for the same namespace, but not multiple namespaces in different parts of the document under a single prefix. A document is rejected when its number of unique prefixes exceeds the maximum number. Enter a value in the range 0 - 262143. The default value is 1024. A value of 0 indicates unlimited. When 0, the action does not return the number of unique prefixes.",
							Computed:            true,
						},
						"unique_namespaces": schema.Int64Attribute{
							MarkdownDescription: "Specify the maximum number of unique XML namespace URIs. This count includes all XML namespaces. A document is rejected when its number of unique namespaces exceeds the maximum number. Enter a value in the range 0 - 65535. The default value is 1024. A value of 0 indicates an unlimited. When 0, the action does not return the number of unique namespaces.",
							Computed:            true,
						},
						"unique_names": schema.Int64Attribute{
							MarkdownDescription: "Specify the maximum number of unique names. A document is rejected when its number of unique names exceeds the maximum number. Enter a value in the range 0 - 1048575. The default value is 1024. A value of 0 indicates an unlimited. When 0, the action does not return the number of unique names. This setting applies to XML and JSON documents. <ul><li>For XML, applies to the number of unique XML local names.</li><li>For JSON, applies to the number of unique JSON labels.</li></ul>",
							Computed:            true,
						},
						"number_length": schema.Int64Attribute{
							MarkdownDescription: "Specify the maximum number length in bytes for the value portion of a label-value pair when the value is a number. The number must be a contiguous string of bytes that contain no white space. The number can include a minus sign and a positive or negative exponent. A document is rejected when the number length in the document exceeds the maximum length. Enter a value in the range 0 - 256. The default value is 128. A value of 0 indicates unlimited. When 0, the action does not return the number length.",
							Computed:            true,
						},
						"strict_utf8encoding": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to enforce strict UTF-8 encoding throughout the entire JSON document. When enabled, the entire document is checked for valid UTF-8 encoding. When disabled, only the first few bytes are checked for proper encoding and the rest of the document is assumed to be in the same encoding.",
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

func (d *ParseSettingsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *ParseSettingsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data ParseSettingsList
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if data.ProviderTarget.ValueString() != "" && !d.pData.Client.ValidTarget(data.ProviderTarget.ValueString()) {
		resp.Diagnostics.AddError("Invalid provider_target", fmt.Sprintf(`Target %q is not defined in the provider's 'targets' block. Available targets: %v`, data.ProviderTarget.ValueString(), d.pData.Client.GetTargetNames()))
		return
	}
	o := models.ParseSettings{
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
	l := []models.ParseSettings{}
	if resFound {
		if value := res.Get(`ParseSettings`); value.Exists() {
			for _, v := range value.Array() {
				item := models.ParseSettings{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.ParseSettingsObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.ParseSettingsObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

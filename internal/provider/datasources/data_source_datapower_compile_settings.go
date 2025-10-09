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

type CompileSettingsList struct {
	Id        types.String `tfsdk:"id"`
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
	resp.TypeName = req.ProviderTypeName + "_compile_settings"
}

func (d *CompileSettingsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Configure customized compile settings.",
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
						"xslt_version": schema.StringAttribute{
							MarkdownDescription: "Specifies the XSLT processor version. The default value is XSLT10.",
							Computed:            true,
						},
						"strict": schema.BoolAttribute{
							MarkdownDescription: "Specifies whether to enable strict XSLT error checking. Non-strict operations attempt to recover from certain errors, such as use of undeclared variables, calling undeclared templates, and so forth. By default, strict XSLT error checking is enabled.",
							Computed:            true,
						},
						"profile": schema.BoolAttribute{
							MarkdownDescription: "Specifies whether to enable stylesheet profiling. This option should not be used in production environments. By default, stylesheet profiling is disabled.",
							Computed:            true,
						},
						"debug": schema.BoolAttribute{
							MarkdownDescription: "Specifies whether to run the stylesheet, XQuery script, and JSONiq script in debug mode. When a stylesheet, XQuery script, or JSONiq script is run in debug mode, it generates a custom web page instead of displaying its normal output. The web page details exactly what occurred during execution, including the values of variables and where particular pieces of the output came from. This option should not be used in production environments. By default, debug mode is disabled.",
							Computed:            true,
						},
						"stream": schema.BoolAttribute{
							MarkdownDescription: "Specifies whether the stylesheet must be run in streaming mode. Transformation of the document begins before the input is fully parsed. Not all stylesheets can be streamed. If the stylesheet cannot be streamed, an error is generated and the input is not processed. By default, streaming mode is disabled.",
							Computed:            true,
						},
						"try_stream": schema.BoolAttribute{
							MarkdownDescription: "Specifies whether to attempt to run the stylesheet in streaming mode. Transformation of the document begins before the input is fully parsed. Not all stylesheets can be streamed. If the stylesheet cannot be streamed, a warning is generated during compilation and the stylesheet is read in the entire input as normal at execution time. By default, attempting to run the stylesheet in streaming mode is disabled.",
							Computed:            true,
						},
						"minimum_escaping": schema.BoolAttribute{
							MarkdownDescription: "Specifies whether to escape output produced from the stylesheet during processing. Minimal escaping is particularly useful when handling non-English character sets. By default, minimum escaping is disabled.",
							Computed:            true,
						},
						"stack_size": schema.Int64Attribute{
							MarkdownDescription: "Indicates the maximum number of bytes that the stack is allowed to use while executing a stylesheet or other compiled content. This setting is used to block infinite recursion. The minimum value is 10 kilobytes, or 10,240 bytes. The maximum value is 100 megabytes, or 104,857,600 bytes. The default value is 1 megabyte, or 1,048,576 bytes.",
							Computed:            true,
						},
						"ws_i_validation": schema.StringAttribute{
							MarkdownDescription: "Specifies the validation behavior to apply to WSDL files that are checked for conformance to section 5 of WS-I Basic Profile (version 1.0, April 2004). The default setting is Warn.",
							Computed:            true,
						},
						"wsdl_validate_body": schema.StringAttribute{
							MarkdownDescription: "Specifies the validation behavior for the soap:Body. The default setting is Strict.",
							Computed:            true,
						},
						"wsdl_validate_headers": schema.StringAttribute{
							MarkdownDescription: "Specifies the validation behavior for the soap:Header. The default setting is Lax.",
							Computed:            true,
						},
						"wsdl_validate_faults": schema.StringAttribute{
							MarkdownDescription: "Specifies the validation behavior for the fault detail. The default setting is Strict.",
							Computed:            true,
						},
						"wsdl_wrapped_faults": schema.BoolAttribute{
							MarkdownDescription: "Specifies whether to require compatibility with RPC-style wrappers. By default, RPC-style wrappers are not required.",
							Computed:            true,
						},
						"allow_soap_enc_array": schema.BoolAttribute{
							MarkdownDescription: "Specifies whether to allow the schema to accept most uses of elements with xsi:type='SOAP-ENC:Array' consistent with SOAP 1.1 Section 5, even when these attributes violate the XML Schema specification. Normally the xsi:type attribute must name a type equal to or derived from the actual type of the element. For schemas compiled with this option, xsi:type is accepted specifically for the SOAP 1.1 Encoding 'Array' complex type if the element type is derived from SOAP-ENC:Array. The opposite is the normal allowable case. By default, elements with xsi:type='SOAP-ENC:Array' are not accepted.",
							Computed:            true,
						},
						"validate_soap_enc_array": schema.BoolAttribute{
							MarkdownDescription: "Specifies whether to perform extra schema validation following the encoding rules in SOAP 1.1 Section 5. When enabled, members of SOAP arrays are validated, attributes such as @id and @href are allowed even if they are not allowed by the schema, and @href values are checked to ensure that they have a corresponding @id element. By default, the extra validation is not performed.",
							Computed:            true,
						},
						"wildcards_ignore_xsi_type": schema.BoolAttribute{
							MarkdownDescription: "Specifies whether xs:any elements in the schema validate only child elements by name. The XML Schema specification requires that, if a wildcard matches an element but that element does not have an element declaration, the element is instead validated according to an xsi:type attribute on it. This option ignores those xsi:type attributes. It should be used for cases such as SOAP envelope validation where a further validation step will validate the contents matching the wildcard, possibly using the SOAP 1.1 encoding rules. By default, xsi:type attributes are not ignored.",
							Computed:            true,
						},
						"wsdl_strict_soap_version": schema.BoolAttribute{
							MarkdownDescription: "Specifies whether to strictly follow the SOAP binding in the WSDL. When enabled, only messages bound to SOAP 1.2 appear in SOAP 1.2 envelopes and only messages bound to SOAP 1.1 appear in SOAP 1.1 envelopes. By default, strict SOAP binding is disabled.",
							Computed:            true,
						},
						"xacml_debug": schema.BoolAttribute{
							MarkdownDescription: "Specifies whether to compile XACML policies with debug information. Note that the XACML debugging messages are also controlled by the log event in the XACML category. Use the debug log level to view the full XACML debugging messages. By default, XACML policies are not compiled with debug information.",
							Computed:            true,
						},
						"allow_xop_include": schema.BoolAttribute{
							MarkdownDescription: "<p>Specifies whether the schema or WSDL document accepts messages where base64-encoded binary content was optimized according to the MTOM/XOP specifications. XOP binary-optimization replaces base64-encoded binary data with an xop:Include reference element that references the unencoded binary data located in an attachment. By default, MTOM/XOP optimized messages are disabled.</p><ul><li>When disabled, such optimized messages are rejected by validation of the optimized form. Rejection occurs because the schema specifies a simple type that accepts base64-encoded data, such as xs:base64Binary or xs:string, but the message contains an xop:Include element instead.</li><li>When enabled, an xop:Include element can optionally appear in place of content for any XML Schema simple type that validates base64-encoded binary data. The xop:Include element itself will be validated according to the built-in schema in store:///schemas/xop.xsd.</li></ul>",
							Computed:            true,
						},
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
	l := []models.CompileSettings{}
	if resFound {
		if value := res.Get(`CompileSettings`); value.Exists() {
			for _, v := range value.Array() {
				item := models.CompileSettings{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
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

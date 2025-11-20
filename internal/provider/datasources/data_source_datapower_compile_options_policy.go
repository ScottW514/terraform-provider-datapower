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

type CompileOptionsPolicyList struct {
	Id        types.String `tfsdk:"id"`
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
	resp.TypeName = req.ProviderTypeName + "_compile_options_policy"
}

func (d *CompileOptionsPolicyDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Configure/Edit Compile Option Policies",
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
							MarkdownDescription: "Select the XSLT processor version.",
							Computed:            true,
						},
						"strict": schema.BoolAttribute{
							MarkdownDescription: "Enable strict XSLT error-checking; non-strict operation attempts to recover from certain errors such as use of undeclared variables, calling undeclared templates, and so forth.",
							Computed:            true,
						},
						"profile": schema.StringAttribute{
							MarkdownDescription: "Designates a set of stylesheets to be profiled based on their URL. This should not be used in production environments.",
							Computed:            true,
						},
						"debug": schema.StringAttribute{
							MarkdownDescription: "Designates a set of stylesheets, XQuery scripts, and JSONiq scripts to be run in debug mode. When a stylesheet, XQuery script, or JSONiq script is run in debug mode, it generates a custom web page instead of displaying its normal output. The web page details exactly what occurred during execution, including the values of variables and where particular pieces of the output came from. This should not be used in production environments.",
							Computed:            true,
						},
						"stream": schema.StringAttribute{
							MarkdownDescription: "Designates a set of stylesheets which must be run in streaming mode. Transformation of the document begins before the input is fully parsed. Not all stylesheets can be streamed; if a stylesheet cannot be streamed, an error will be issued and the input will not be processed. See the DataPower manual for suggestions on producing streamable stylesheets.",
							Computed:            true,
						},
						"try_stream": schema.StringAttribute{
							MarkdownDescription: "Designates a set of stylesheets to attempt to run in streaming mode. Transformation of the document begins before the input is fully parsed. Not all stylesheets can be streamed; if a stylesheet cannot be streamed, a warning will be issued during compilation and the stylesheet will read in the entire input as normal at execution time. See the DataPower manual for suggestions on producing streamable stylesheets.",
							Computed:            true,
						},
						"minimum_escaping": schema.StringAttribute{
							MarkdownDescription: "Select a URL Map from the list. Output produced from stylesheets that meet the URL map criteria are not escaped during processing. Escaping is enabled by default. Minimal escaping is particularly useful when handling non-English character sets.",
							Computed:            true,
						},
						"stack_size": schema.Int64Attribute{
							MarkdownDescription: "Indicates the maximum number of bytes the stack is allowed to use while executing a stylesheet or other compiled content. This blocks infinite recursion. The minimum value is 10 kilobytes, or 10,240 bytes. The default is half a megabyte, or 524,288 bytes.",
							Computed:            true,
						},
						"prefer_xg4": schema.StringAttribute{
							MarkdownDescription: "<p>Designates a set of stylesheets or schemas that must use XML hardware acceleration when possible. XML hardware acceleration will be used only if the first action in a processing policy is a schema validation against a fixed schema. Any schemas specified here will use XML hardware acceleration if hardware resources are available. Schemas not specified here may also use XML hardware acceleration, but those schemas will be displaced in favor of schemas specified here.</p><p>When XML hardware acceleration is disabled in the Systems Settings, the XML Hardware Acceleration Preferred Rule has no effect.</p>",
							Computed:            true,
						},
						"disallow_xg4": schema.StringAttribute{
							MarkdownDescription: "<p>Designates a set of stylesheets or schemas that will never use XML hardware acceleration. XML hardware acceleration will be used only if the first action in a processing policy is a schema validation against a fixed schema. Any schemas specified here will not use XML hardware acceleration even if hardware resources are available.</p><p>When XML hardware acceleration is disabled in the Systems Settings, the XML Hardware Acceleration Disallowed Rule has no effect.</p>",
							Computed:            true,
						},
						"ws_i_validation": schema.StringAttribute{
							MarkdownDescription: "Select the validation behavior to apply to WSDL files that are checked for conformance to section 5 of WS-I Basic Profile (version 1.0, April 2004). The default is Ignore.",
							Computed:            true,
						},
						"wsdl_validate_body": schema.StringAttribute{
							MarkdownDescription: "Select the validation behavior for the soap:Body. The default is Strict.",
							Computed:            true,
						},
						"wsdl_validate_headers": schema.StringAttribute{
							MarkdownDescription: "Select the validation behavior for the soap:Header. The default is Lax.",
							Computed:            true,
						},
						"wsdl_validate_faults": schema.StringAttribute{
							MarkdownDescription: "Select the validation behavior for the fault detail. The default is Strict.",
							Computed:            true,
						},
						"wsdl_wrapped_faults": schema.BoolAttribute{
							MarkdownDescription: "For compatibility, require a rpc-style wrapper around fault details that are specified by type.",
							Computed:            true,
						},
						"allow_soap_enc_array": schema.StringAttribute{
							MarkdownDescription: "Designates a set of schemas that will accept most uses of elements with xsi:type='SOAP-ENC:Array' consistent with SOAP 1.1 Section 5, even when these attributes violate the XML Schema specification. Normally the xsi:type attribute must name a type equal to or derived from the actual type of the element. For schemas compiled with this option, xsi:type is accepted specifically for the SOAP 1.1 Encoding 'Array' complex type if the element's type is derived from SOAP-ENC:Array; this is the opposite of the normal allowable case.",
							Computed:            true,
						},
						"validate_soap_enc_array": schema.StringAttribute{
							MarkdownDescription: "Designates a set of schemas that will do extra validation following the encoding rules in SOAP 1.1 Section 5. This validates members of SOAP arrays, allows attributes such as @id and @href even if not allowed by the schema, and checks that @href values have a corresponding @id element.",
							Computed:            true,
						},
						"wildcards_ignore_xsi_type": schema.StringAttribute{
							MarkdownDescription: "Designates a set of schemas where wildcards (xs:any elements) only validate children by element name. The XML Schema specification requires that, if a wildcard matches an element but that element does not have an element declaration, the element is instead validated according to an xsi:type attribute on it. This option ignores those xsi:type attributes. It should be used for cases such as SOAP envelope validation where a further validation step will validate the contents matching the wildcard, possibly using the SOAP 1.1 encoding rules.",
							Computed:            true,
						},
						"wsdl_strict_soap_version": schema.BoolAttribute{
							MarkdownDescription: "When on, follow the version of the SOAP Binding in the WSDL, allowing only messages bound to SOAP 1.2 to appear in SOAP 1.2 envelopes and messages bound to SOAP 1.1 to appear in SOAP 1.1 envelopes. The default is off.",
							Computed:            true,
						},
						"xacml_debug": schema.BoolAttribute{
							MarkdownDescription: "Setting to \"on\" to make the XACML compiler to put more debugging information when evaluate a policy. Note that the XACML debugging messages are also controlled by the log event in the 'XACML' category. Use the \"debug\" log level to view the full XACML debugging messages.",
							Computed:            true,
						},
						"allow_xop_include": schema.StringAttribute{
							MarkdownDescription: "<p>Designates a set of schema or WSDL documents that accept messages where base64-encoded binary content was optimized according to the MTOM/XOP specifications. XOP binary-optimization replaces base64-encoded binary data with an xop:Include reference element that references the unencoded binary data located in an attachment.</p><ul><li>When disabled, such optimized messages are rejected by validation of the optimized form. Rejection occurs because the schema specifies a simple type that accepts base64-encoded data, such as xs:base64Binary or xs:string, but the message contains an xop:Include element instead.</li><li>When enabled, an xop:Include element can optionally appear in place of content for any XML Schema simple type that validates base64-encoded binary data. The xop:Include element itself will be validated according to the built-in schema in store:///schemas/xop.xsd.</li></ul>",
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

	path := o.GetPath()
	if !data.Id.IsNull() {
		path = path + "/" + data.Id.ValueString()
	}

	res, err := d.pData.Client.Get(path)
	resFound := true
	if err != nil {
		if strings.Contains(err.Error(), "status 401") {
			_ = tfutils.DomainCredentialTest(d.pData.Client, &resp.Diagnostics, data.AppDomain.ValueString())
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
	l := []models.CompileOptionsPolicy{}
	if resFound {
		if value := res.Get(`CompileOptionsPolicy`); value.Exists() {
			for _, v := range value.Array() {
				item := models.CompileOptionsPolicy{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
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

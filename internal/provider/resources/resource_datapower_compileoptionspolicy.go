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

package resources

import (
	"context"
	"fmt"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/modifiers"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
)

var _ resource.Resource = &CompileOptionsPolicyResource{}

func NewCompileOptionsPolicyResource() resource.Resource {
	return &CompileOptionsPolicyResource{}
}

type CompileOptionsPolicyResource struct {
	pData *tfutils.ProviderData
}

func (r *CompileOptionsPolicyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_compileoptionspolicy"
}

func (r *CompileOptionsPolicyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("Configure/Edit Compile Option Policies", "compile-options", "").String,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Name of the object. Must be unique among object types in application domain.", "", "").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 128),
					stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z0-9_-]+$`), ""),
				},
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"app_domain": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("The name of the application domain the object belongs to", "", "").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 128),
					stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z0-9_-]+$`), ""),
				},
				PlanModifiers: []planmodifier.String{
					modifiers.ImmutableAfterSet(),
				},
			},
			"user_summary": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Comments", "summary", "").String,
				Optional:            true,
			},
			"xslt_version": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Select the XSLT processor version.", "xslt-version", "").AddStringEnum("XSLT10", "XSLT10_IT23272", "XSLT20", "StylesheetSpecified").AddDefaultValue("XSLT10").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("XSLT10", "XSLT10_IT23272", "XSLT20", "StylesheetSpecified"),
				},
				Default: stringdefault.StaticString("XSLT10"),
			},
			"strict": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enable strict XSLT error-checking; non-strict operation attempts to recover from certain errors such as use of undeclared variables, calling undeclared templates, and so forth.", "strict", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"profile": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Designates a set of stylesheets to be profiled based on their URL. This should not be used in production environments.", "profile", "urlmap").String,
				Optional:            true,
			},
			"debug": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Designates a set of stylesheets, XQuery scripts, and JSONiq scripts to be run in debug mode. When a stylesheet, XQuery script, or JSONiq script is run in debug mode, it generates a custom web page instead of displaying its normal output. The web page details exactly what occurred during execution, including the values of variables and where particular pieces of the output came from. This should not be used in production environments.", "debug", "urlmap").String,
				Optional:            true,
			},
			"stream": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Designates a set of stylesheets which must be run in streaming mode. Transformation of the document begins before the input is fully parsed. Not all stylesheets can be streamed; if a stylesheet cannot be streamed, an error will be issued and the input will not be processed. See the DataPower manual for suggestions on producing streamable stylesheets.", "stream", "urlmap").String,
				Optional:            true,
			},
			"try_stream": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Designates a set of stylesheets to attempt to run in streaming mode. Transformation of the document begins before the input is fully parsed. Not all stylesheets can be streamed; if a stylesheet cannot be streamed, a warning will be issued during compilation and the stylesheet will read in the entire input as normal at execution time. See the DataPower manual for suggestions on producing streamable stylesheets.", "try-stream", "urlmap").String,
				Optional:            true,
			},
			"minimum_escaping": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Select a URL Map from the list. Output produced from stylesheets that meet the URL map criteria are not escaped during processing. Escaping is enabled by default. Minimal escaping is particularly useful when handling non-English character sets.", "minesc", "urlmap").String,
				Optional:            true,
			},
			"stack_size": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Indicates the maximum number of bytes the stack is allowed to use while executing a stylesheet or other compiled content. This blocks infinite recursion. The minimum value is 10 kilobytes, or 10,240 bytes. The default is half a megabyte, or 524,288 bytes.", "stack-size", "").AddIntegerRange(10240, 104857600).AddDefaultValue("524288").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(10240, 104857600),
				},
				Default: int64default.StaticInt64(524288),
			},
			"prefer_xg4": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Designates a set of stylesheets or schemas that must use XML hardware acceleration when possible. XML hardware acceleration will be used only if the first action in a processing policy is a schema validation against a fixed schema. Any schemas specified here will use XML hardware acceleration if hardware resources are available. Schemas not specified here may also use XML hardware acceleration, but those schemas will be displaced in favor of schemas specified here.</p><p>When XML hardware acceleration is disabled in the Systems Settings, the XML Hardware Acceleration Preferred Rule has no effect.</p>", "prefer-xg4", "urlmap").String,
				Optional:            true,
			},
			"disallow_xg4": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Designates a set of stylesheets or schemas that will never use XML hardware acceleration. XML hardware acceleration will be used only if the first action in a processing policy is a schema validation against a fixed schema. Any schemas specified here will not use XML hardware acceleration even if hardware resources are available.</p><p>When XML hardware acceleration is disabled in the Systems Settings, the XML Hardware Acceleration Disallowed Rule has no effect.</p>", "disallow-xg4", "urlmap").String,
				Optional:            true,
			},
			"wsi_validation": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Select the validation behavior to apply to WSDL files that are checked for conformance to section 5 of WS-I Basic Profile (version 1.0, April 2004). The default is Ignore.", "wsi-validate", "").AddStringEnum("ignore", "warn", "fail").AddDefaultValue("ignore").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("ignore", "warn", "fail"),
				},
				Default: stringdefault.StaticString("ignore"),
			},
			"wsdl_validate_body": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Select the validation behavior for the soap:Body. The default is Strict.", "wsdl-validate-body", "").AddStringEnum("strict", "lax", "skip").AddDefaultValue("strict").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("strict", "lax", "skip"),
				},
				Default: stringdefault.StaticString("strict"),
			},
			"wsdl_validate_headers": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Select the validation behavior for the soap:Header. The default is Lax.", "wsdl-validate-headers", "").AddStringEnum("strict", "lax", "skip").AddDefaultValue("lax").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("strict", "lax", "skip"),
				},
				Default: stringdefault.StaticString("lax"),
			},
			"wsdl_validate_faults": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Select the validation behavior for the fault detail. The default is Strict.", "wsdl-validate-faults", "").AddStringEnum("strict", "lax", "skip").AddDefaultValue("strict").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("strict", "lax", "skip"),
				},
				Default: stringdefault.StaticString("strict"),
			},
			"wsdl_wrapped_faults": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("For compatibility, require a rpc-style wrapper around fault details that are specified by type.", "wsdl-wrapped-faults", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"allow_soap_enc_array": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Designates a set of schemas that will accept most uses of elements with xsi:type='SOAP-ENC:Array' consistent with SOAP 1.1 Section 5, even when these attributes violate the XML Schema specification. Normally the xsi:type attribute must name a type equal to or derived from the actual type of the element. For schemas compiled with this option, xsi:type is accepted specifically for the SOAP 1.1 Encoding 'Array' complex type if the element's type is derived from SOAP-ENC:Array; this is the opposite of the normal allowable case.", "allow-soap-enc-array", "urlmap").String,
				Optional:            true,
			},
			"validate_soap_enc_array": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Designates a set of schemas that will do extra validation following the encoding rules in SOAP 1.1 Section 5. This validates members of SOAP arrays, allows attributes such as @id and @href even if not allowed by the schema, and checks that @href values have a corresponding @id element.", "validate-soap-enc-array", "urlmap").String,
				Optional:            true,
			},
			"wildcards_ignore_xsi_type": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Designates a set of schemas where wildcards (xs:any elements) only validate children by element name. The XML Schema specification requires that, if a wildcard matches an element but that element does not have an element declaration, the element is instead validated according to an xsi:type attribute on it. This option ignores those xsi:type attributes. It should be used for cases such as SOAP envelope validation where a further validation step will validate the contents matching the wildcard, possibly using the SOAP 1.1 encoding rules.", "wildcards-ignore-xsi-type", "urlmap").String,
				Optional:            true,
			},
			"wsdl_strict_soap_version": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("When on, follow the version of the SOAP Binding in the WSDL, allowing only messages bound to SOAP 1.2 to appear in SOAP 1.2 envelopes and messages bound to SOAP 1.1 to appear in SOAP 1.1 envelopes. The default is off.", "wsdl-strict-soap-version", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"xacml_debug": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Setting to \"on\" to make the XACML compiler to put more debugging information when evaluate a policy. Note that the XACML debugging messages are also controlled by the log event in the 'XACML' category. Use the \"debug\" log level to view the full XACML debugging messages.", "xacml-debug", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"allow_xop_include": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Designates a set of schema or WSDL documents that accept messages where base64-encoded binary content was optimized according to the MTOM/XOP specifications. XOP binary-optimization replaces base64-encoded binary data with an xop:Include reference element that references the unencoded binary data located in an attachment.</p><ul><li>When disabled, such optimized messages are rejected by validation of the optimized form. Rejection occurs because the schema specifies a simple type that accepts base64-encoded data, such as xs:base64Binary or xs:string, but the message contains an xop:Include element instead.</li><li>When enabled, an xop:Include element can optionally appear in place of content for any XML Schema simple type that validates base64-encoded binary data. The xop:Include element itself will be validated according to the built-in schema in store:///schemas/xop.xsd.</li></ul>", "allow-xop-include", "urlmap").String,
				Optional:            true,
			},
			"dependency_actions": actions.ActionsSchema,
		},
	}
}

func (r *CompileOptionsPolicyResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (r *CompileOptionsPolicyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.CompileOptionsPolicy
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Create, false)
	if resp.Diagnostics.HasError() {
		return
	}

	body := data.ToBody(ctx, `CompileOptionsPolicy`)
	_, err := r.pData.Client.Post(data.GetPath(), body)

	if err != nil && !strings.Contains(err.Error(), "status 409") {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to create object (%s), got error: %s", "POST", err))
		return
	}
	actions.PostProcess(ctx, &resp.Diagnostics, data.DependencyActions, actions.Create)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CompileOptionsPolicyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.CompileOptionsPolicy
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	res, err := r.pData.Client.Get(data.GetPath() + "/" + data.Id.ValueString())
	if err != nil && (strings.Contains(err.Error(), "status 404") || strings.Contains(err.Error(), "status 406") || strings.Contains(err.Error(), "status 500") || strings.Contains(err.Error(), "status 400")) {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s", err))
		return
	}

	if data.IsNull() {
		// Import
		data.FromBody(ctx, `CompileOptionsPolicy`, res)
	} else {
		// Update
		data.UpdateFromBody(ctx, `CompileOptionsPolicy`, res)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CompileOptionsPolicyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.CompileOptionsPolicy
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Update, false)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.pData.Client.Put(data.GetPath()+"/"+data.Id.ValueString(), data.ToBody(ctx, `CompileOptionsPolicy`))
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to update object (PUT), got error: %s", err))
		return
	}

	actions.PostProcess(ctx, &resp.Diagnostics, data.DependencyActions, actions.Update)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CompileOptionsPolicyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.CompileOptionsPolicy
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Delete, false)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.pData.Client.Delete(data.GetPath() + "/" + data.Id.ValueString())
	if err != nil && !strings.Contains(err.Error(), "status 404") && !strings.Contains(err.Error(), "status 409") {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object (DELETE), got error: %s", err))
		return
	}

	actions.PostProcess(ctx, &resp.Diagnostics, data.DependencyActions, actions.Delete)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.State.RemoveResource(ctx)
}

func (r *CompileOptionsPolicyResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data models.CompileOptionsPolicy

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.ValidateConfig(ctx, &resp.Diagnostics, data.DependencyActions)
}

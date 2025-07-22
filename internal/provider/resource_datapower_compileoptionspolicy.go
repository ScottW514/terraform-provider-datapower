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

package provider

import (
	"context"
	"fmt"
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
	"github.com/scottw514/terraform-provider-datapower/client"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
)

var _ resource.Resource = &CompileOptionsPolicyResource{}

func NewCompileOptionsPolicyResource() resource.Resource {
	return &CompileOptionsPolicyResource{}
}

type CompileOptionsPolicyResource struct {
	client *client.DatapowerClient
}

func (r *CompileOptionsPolicyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_compileoptionspolicy"
}

func (r *CompileOptionsPolicyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("Compile Options Policy", "compile-options", "").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Name of the object. Must be unique among object types in application domain.", "", "").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"app_domain": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("The name of the application domain the object belongs to", "", "").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					ImmutableAfterSet(),
				},
			},
			"user_summary": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Comments", "summary", "").String,
				Optional:            true,
			},
			"xslt_version": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("XSLT Version", "xslt-version", "").AddStringEnum("XSLT10", "XSLT10_IT23272", "XSLT20", "StylesheetSpecified").AddDefaultValue("XSLT10").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("XSLT10", "XSLT10_IT23272", "XSLT20", "StylesheetSpecified"),
				},
				Default: stringdefault.StaticString("XSLT10"),
			},
			"strict": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Strict", "strict", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"profile": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Profile Rule", "profile", "urlmap").String,
				Optional:            true,
			},
			"debug": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Debug Rule", "debug", "urlmap").String,
				Optional:            true,
			},
			"stream": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Streaming Rule", "stream", "urlmap").String,
				Optional:            true,
			},
			"try_stream": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Attempt Streaming Rule", "try-stream", "urlmap").String,
				Optional:            true,
			},
			"minimum_escaping": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Minimum Output Escaping Rule", "minesc", "urlmap").String,
				Optional:            true,
			},
			"stack_size": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Maximum Stack Size", "stack-size", "").AddIntegerRange(10240, 104857600).AddDefaultValue("524288").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(10240, 104857600),
				},
				Default: int64default.StaticInt64(524288),
			},
			"prefer_xg4": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("XML Hardware Acceleration Preferred Rule", "prefer-xg4", "urlmap").String,
				Optional:            true,
			},
			"disallow_xg4": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("XML Hardware Acceleration Disallowed Rule", "disallow-xg4", "urlmap").String,
				Optional:            true,
			},
			"wsi_validation": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("WS-I Basic Profile Validation", "wsi-validate", "").AddStringEnum("ignore", "warn", "fail").AddDefaultValue("ignore").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("ignore", "warn", "fail"),
				},
				Default: stringdefault.StaticString("ignore"),
			},
			"wsdl_validate_body": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Validate Message Body", "wsdl-validate-body", "").AddStringEnum("strict", "lax", "skip").AddDefaultValue("strict").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("strict", "lax", "skip"),
				},
				Default: stringdefault.StaticString("strict"),
			},
			"wsdl_validate_headers": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Validate Message Headers", "wsdl-validate-headers", "").AddStringEnum("strict", "lax", "skip").AddDefaultValue("lax").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("strict", "lax", "skip"),
				},
				Default: stringdefault.StaticString("lax"),
			},
			"wsdl_validate_faults": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Validate Message Fault details", "wsdl-validate-faults", "").AddStringEnum("strict", "lax", "skip").AddDefaultValue("strict").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("strict", "lax", "skip"),
				},
				Default: stringdefault.StaticString("strict"),
			},
			"wsdl_wrapped_faults": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Require wrappers on fault-details specified by type", "wsdl-wrapped-faults", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"allow_soap_enc_array": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specifically Allow xsi:type='SOAP-ENC:Array' Rule", "allow-soap-enc-array", "urlmap").String,
				Optional:            true,
			},
			"validate_soap_enc_array": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Validate SOAP 1.1 Encoding Rule", "validate-soap-enc-array", "urlmap").String,
				Optional:            true,
			},
			"wildcards_ignore_xsi_type": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Wildcards Ignore xsi:type Rule", "wildcards-ignore-xsi-type", "urlmap").String,
				Optional:            true,
			},
			"wsdl_strict_soap_version": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Strict SOAP Envelope Version", "wsdl-strict-soap-version", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"xacml_debug": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Debug XACML Policy", "xacml-debug", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"allow_xop_include": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Accept MTOM/XOP Optimized Messages", "allow-xop-include", "urlmap").String,
				Optional:            true,
			},
		},
	}
}

func (r *CompileOptionsPolicyResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = *req.ProviderData.(**client.DatapowerClient)
}

func (r *CompileOptionsPolicyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data models.CompileOptionsPolicy

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	body := data.ToBody(ctx, `CompileOptionsPolicy`)
	res, err := r.client.Post(data.GetPath(), body)

	if err != nil && !strings.Contains(err.Error(), "status 409") {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to create object (%s), got error: %s, %s", "POST", err, res.String()))
		return
	}

	res, err = r.client.Post("/mgmt/actionqueue/"+data.AppDomain.ValueString(), "{\"SaveConfig\": 0}")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to save object (%s), got error: %s, %s", "POST", err, res.String()))
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CompileOptionsPolicyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data models.CompileOptionsPolicy

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	res, err := r.client.Get(data.GetPath() + "/" + data.Id.ValueString())
	if err != nil && (strings.Contains(err.Error(), "status 404") || strings.Contains(err.Error(), "status 406") || strings.Contains(err.Error(), "status 500") || strings.Contains(err.Error(), "status 400")) {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
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
	var data models.CompileOptionsPolicy

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	res, err := r.client.Put(data.GetPath()+"/"+data.Id.ValueString(), data.ToBody(ctx, `CompileOptionsPolicy`))
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to update object (PUT), got error: %s, %s", err, res.String()))
		return
	}
	res, err = r.client.Post("/mgmt/actionqueue/"+data.AppDomain.ValueString(), "{\"SaveConfig\": 0}")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to save object (%s), got error: %s, %s", "POST", err, res.String()))
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CompileOptionsPolicyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data models.CompileOptionsPolicy

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	res, err := r.client.Delete(data.GetPath() + "/" + data.Id.ValueString())
	if err != nil && !strings.Contains(err.Error(), "status 404") && !strings.Contains(err.Error(), "status 409") {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object (DELETE), got error: %s, %s", err, res.String()))
		return
	}

	resp.State.RemoveResource(ctx)
}

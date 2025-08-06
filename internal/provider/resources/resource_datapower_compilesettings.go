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
	"github.com/scottw514/terraform-provider-datapower/client"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/modifiers"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
)

var _ resource.Resource = &CompileSettingsResource{}

func NewCompileSettingsResource() resource.Resource {
	return &CompileSettingsResource{}
}

type CompileSettingsResource struct {
	client *client.DatapowerClient
}

func (r *CompileSettingsResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_compilesettings"
}

func (r *CompileSettingsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("Compile Settings", "compile-settings", "").String,

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
				MarkdownDescription: tfutils.NewAttributeDescription("XSLT version", "xslt-version", "").AddStringEnum("XSLT10", "XSLT10_IT23272", "XSLT20", "StylesheetSpecified").AddDefaultValue("XSLT10").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("XSLT10", "XSLT10_IT23272", "XSLT20", "StylesheetSpecified"),
				},
				Default: stringdefault.StaticString("XSLT10"),
			},
			"strict": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Strict", "strict", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"profile": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Profile rule", "profile", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"debug": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Debug rule", "debug", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"stream": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Streaming rule", "stream", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"try_stream": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Attempt streaming rule", "try-stream", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"minimum_escaping": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Minimum output escaping rule", "minesc", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"stack_size": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Maximum stack size", "stack-size", "").AddIntegerRange(10240, 104857600).AddDefaultValue("1048576").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(10240, 104857600),
				},
				Default: int64default.StaticInt64(1048576),
			},
			"wsi_validation": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("WS-I Basic Profile validation", "wsi-validate", "").AddStringEnum("ignore", "warn", "fail").AddDefaultValue("warn").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("ignore", "warn", "fail"),
				},
				Default: stringdefault.StaticString("warn"),
			},
			"wsdl_validate_body": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Validate message body", "wsdl-validate-body", "").AddStringEnum("strict", "lax", "skip").AddDefaultValue("strict").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("strict", "lax", "skip"),
				},
				Default: stringdefault.StaticString("strict"),
			},
			"wsdl_validate_headers": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Validate message headers", "wsdl-validate-headers", "").AddStringEnum("strict", "lax", "skip").AddDefaultValue("lax").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("strict", "lax", "skip"),
				},
				Default: stringdefault.StaticString("lax"),
			},
			"wsdl_validate_faults": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Validate message fault details", "wsdl-validate-faults", "").AddStringEnum("strict", "lax", "skip").AddDefaultValue("strict").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("strict", "lax", "skip"),
				},
				Default: stringdefault.StaticString("strict"),
			},
			"wsdl_wrapped_faults": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Require wrappers on fault details specified by type", "wsdl-wrapped-faults", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"allow_soap_enc_array": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specifically allow xsi:type='SOAP-ENC:Array' rule", "allow-soap-enc-array", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"validate_soap_enc_array": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Validate SOAP 1.1 encoding rule", "validate-soap-enc-array", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"wildcards_ignore_xsi_type": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Wildcards ignore xsi:type rule", "wildcards-ignore-xsi-type", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"wsdl_strict_soap_version": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Strict SOAP envelope version", "wsdl-strict-soap-version", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"xacml_debug": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Debug XACML policy", "xacml-debug", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"allow_xop_include": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Accept MTOM/XOP optimized messages", "allow-xop-include", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"object_actions": actions.ActionsSchema,
		},
	}
}

func (r *CompileSettingsResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = *req.ProviderData.(**client.DatapowerClient)
}

func (r *CompileSettingsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data models.CompileSettings

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, r.client, data.ObjectActions, actions.Create)

	body := data.ToBody(ctx, `CompileSettings`)
	_, err := r.client.Post(data.GetPath(), body)

	if err != nil && !strings.Contains(err.Error(), "status 409") {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to create object (%s), got error: %s", "POST", err))
		return
	}

	_, err = r.client.Post("/mgmt/actionqueue/"+data.AppDomain.ValueString(), "{\"SaveConfig\": 0}")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to save object (%s), got error: %s", "POST", err))
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CompileSettingsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data models.CompileSettings

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	res, err := r.client.Get(data.GetPath() + "/" + data.Id.ValueString())
	if err != nil && (strings.Contains(err.Error(), "status 404") || strings.Contains(err.Error(), "status 406") || strings.Contains(err.Error(), "status 500") || strings.Contains(err.Error(), "status 400")) {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s", err))
		return
	}

	if data.IsNull() {
		// Import
		data.FromBody(ctx, `CompileSettings`, res)
	} else {
		// Update
		data.UpdateFromBody(ctx, `CompileSettings`, res)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CompileSettingsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data models.CompileSettings

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, r.client, data.ObjectActions, actions.Update)
	_, err := r.client.Put(data.GetPath()+"/"+data.Id.ValueString(), data.ToBody(ctx, `CompileSettings`))
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to update object (PUT), got error: %s", err))
		return
	}
	_, err = r.client.Post("/mgmt/actionqueue/"+data.AppDomain.ValueString(), "{\"SaveConfig\": 0}")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to save object (%s), got error: %s", "POST", err))
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CompileSettingsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data models.CompileSettings

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, r.client, data.ObjectActions, actions.Delete)
	_, err := r.client.Delete(data.GetPath() + "/" + data.Id.ValueString())
	if err != nil && !strings.Contains(err.Error(), "status 404") && !strings.Contains(err.Error(), "status 409") {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object (DELETE), got error: %s", err))
		return
	}

	resp.State.RemoveResource(ctx)
}

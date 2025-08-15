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

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/modifiers"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
)

var _ resource.Resource = &AssemblyActionClientSecurityResource{}

func NewAssemblyActionClientSecurityResource() resource.Resource {
	return &AssemblyActionClientSecurityResource{}
}

type AssemblyActionClientSecurityResource struct {
	pData *tfutils.ProviderData
}

func (r *AssemblyActionClientSecurityResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_assembly_action_client_security"
}

func (r *AssemblyActionClientSecurityResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("Client security assembly action", "assembly-client-security", "").String,
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
			"stop_on_error": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to stop processing if client security fails. If failed, stops the assembly and return an error.", "stop-on-error", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"secret_required": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to require the client secret. When required, the secret is compared to the registered secret on the application that is identified by the client ID.", "secret-required", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"extract_credential_method": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Specify the method to extract client credentials from the request.</p><ul><li>For all methods except HTTP, use the ID name and the secret name to specify the locations that contains the ID and the location that contain the secret. <ul><li>When cookie, specify which cookie.</li><li>When context variable, specify which runtime context variable.</li><li>When form data, specify the form data.</li><li>When header, specify which header.</li><li>When query parameter, specify which query parameter.</li></ul></li><li>For the HTTP method, use the HTTP type to specify the format of the <tt>Authorization</tt> header, which expects the basic form in the <tt>Basic <i>base64_id:secret</i></tt> format.</li></ul>", "extract-credential-method", "").AddStringEnum("header", "query", "form", "cookie", "http", "context-var").AddDefaultValue("header").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("header", "query", "form", "cookie", "http", "context-var"),
				},
				Default: stringdefault.StaticString("header"),
			},
			"id_name": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Specify the location where to find the client ID to extract.</p><ul><li>When cookie, specify which cookie.</li><li>When context variable, specify which runtime context variable.</li><li>When form data, specify the form data.</li><li>When header, specify which header.</li><li>When query parameter, specify which query parameter.</li></ul>", "id-name", "").String,
				Optional:            true,
			},
			"secret_name": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Specify the location where to find the secret to extract.</p><ul><li>When cookie, specify which cookie.</li><li>When context variable, specify which runtime context variable.</li><li>When form data, specify the form data.</li><li>When header, specify which header.</li><li>When query parameter, specify which query parameter.</li></ul>", "secret-name", "").String,
				Optional:            true,
			},
			"http_type": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("HTTP type", "http-type", "").AddStringEnum("basic").AddDefaultValue("basic").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("basic"),
				},
				Default: stringdefault.StaticString("basic"),
			},
			"authenticate_client_method": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the method to authenticate the extracted client credentials. When third-party, specify the user-registry to authenticate the extracted client credentials.", "client-auth-method", "").AddStringEnum("native", "third-party").AddDefaultValue("native").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("native", "third-party"),
				},
				Default: stringdefault.StaticString("native"),
			},
			"user_registry": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the API registry to authenticate the extracted client credentials. The supported registries are API authentication URL and API LDAP.", "user-registry", "").String,
				Optional:            true,
			},
			"user_summary": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Comments", "summary", "").String,
				Optional:            true,
			},
			"title": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Title", "title", "").String,
				Optional:            true,
			},
			"correlation_path": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the path that correlates the API action to a specific part of the API specification. The correlation path specifies the part of the API definition that correlates with the API action. This path is exposed in the debug data by the API gateway for use by debugging tools. For example, for an API configuration that is retrieved from API Connect and specified in an OpenAPI document with IBM extensions, this path is the JSON path to the assembly policy in the IBM extensions section of the document. The path can be expressed in any form that the debugging tool can correlate to the API definition.", "correlation-path", "").String,
				Optional:            true,
			},
			"action_debug": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Specify whether to enable the GatewayScript debugger to troubleshoot the following GatewayScript files or script.</p><ul><li>Troubleshoot a GatewayScript file that is called from the GatewayScript assembly action.</li><li>Troubleshoot a GatewayScript file that is called from an XSLT assembly action that uses the <tt>gatewayscript()</tt> extension function.</li><li>Troubleshoot a GatewayScript script that is called through the <tt>value</tt> or <tt>default</tt> property in the JSON file from the map assembly action.</li></ul><p>To debug a file or script, the following conditions must be met.</p><ul><li>The file contains one or more <tt>debugger;</tt> statements at the points in your script where you want to start debugging.</li><li>The GatewayScript debugger is enabled.</li></ul><p>You run the <tt>debug-action</tt> command.</p>", "debug", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"dependency_actions": actions.ActionsSchema,
		},
	}
}

func (r *AssemblyActionClientSecurityResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (r *AssemblyActionClientSecurityResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.AssemblyActionClientSecurity
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Create, false)
	if resp.Diagnostics.HasError() {
		return
	}

	body := data.ToBody(ctx, `AssemblyActionClientSecurity`)
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

func (r *AssemblyActionClientSecurityResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.AssemblyActionClientSecurity
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
		data.FromBody(ctx, `AssemblyActionClientSecurity`, res)
	} else {
		// Update
		data.UpdateFromBody(ctx, `AssemblyActionClientSecurity`, res)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AssemblyActionClientSecurityResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.AssemblyActionClientSecurity
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Update, false)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.pData.Client.Put(data.GetPath()+"/"+data.Id.ValueString(), data.ToBody(ctx, `AssemblyActionClientSecurity`))
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

func (r *AssemblyActionClientSecurityResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.AssemblyActionClientSecurity
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

func (r *AssemblyActionClientSecurityResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data models.AssemblyActionClientSecurity

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.ValidateConfig(ctx, &resp.Diagnostics, data.DependencyActions)
}

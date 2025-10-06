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
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/modifiers"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/validators"
)

var _ resource.Resource = &NameValueProfileResource{}
var _ resource.ResourceWithImportState = &NameValueProfileResource{}

func NewNameValueProfileResource() resource.Resource {
	return &NameValueProfileResource{}
}

type NameValueProfileResource struct {
	pData *tfutils.ProviderData
}

func (r *NameValueProfileResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_name_value_profile"
}

func (r *NameValueProfileResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("Many HTTP things are expressed as name value pairs. These include HTTP headers, cookie values, url-encoded query strings, and url-encoded request messages. This profile provides a mechanism for what kinds of names are expected and for each kind of name what properties should be enforced on the corresponding values. When a name-value pair is not validated successfully that may generate an error, the pair might be stripped from the transaction, or the value may be mapped to another default value.", "webapp-gnvc", "").String,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Name of the object. Must be unique among object types in application domain.", "", "").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 128),
					stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9_-]+$"), "Must match :"+"^[a-zA-Z0-9_-]+$"),
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
					stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9_-]+$"), "Must match :"+"^[a-zA-Z0-9_-]+$"),
				},
				PlanModifiers: []planmodifier.String{
					modifiers.ImmutableAfterSet(),
				},
			},
			"user_summary": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Comments", "summary", "").String,
				Optional:            true,
			},
			"max_attributes": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("The maximum number of name value pairs allowed in a single entity (header, cookie set, body, and so forth).", "max-attributes", "").AddIntegerRange(1, 4294967295).AddDefaultValue("256").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 4294967295),
				},
				Default: int64default.StaticInt64(256),
			},
			"max_aggregate_size": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("The lengths of all the names and values in a single entity (header, cookie set, body, query string, and so forth) in aggregate must not exceed this property.", "max-aggregate-size", "").AddIntegerRange(1, 4294967295).AddDefaultValue("128000").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 4294967295),
				},
				Default: int64default.StaticInt64(128000),
			},
			"max_name_size": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("The maximum size of a name attribute used in this profile.", "max-name-size", "").AddIntegerRange(1, 4294967295).AddDefaultValue("512").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 4294967295),
				},
				Default: int64default.StaticInt64(512),
			},
			"max_value_size": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("The maximum size of a value attribute used in this profile.", "max-value-size", "").AddIntegerRange(1, 4294967295).AddDefaultValue("1024").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 4294967295),
				},
				Default: int64default.StaticInt64(1024),
			},
			"validation_list": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Each pair submitted to this profile consults this validation list, looking for the first regular expression match of the name against the name expression in the list. When that is found, the corresponding value constraint is matched against the value portion of the name-value pair. If that does not match, the policy applies the 'fixup' attribute to the submitted value. That may result in no change, the pair being removed, an error being generated, or the value being mapped to a known constant.", "validation", "").String,
				NestedObject:        models.GetDmValidationTypeResourceSchema(),
				Optional:            true,
			},
			"default_fixup": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Select the action to taken when no matching entry in the validation list is found. The default is Strip.", "unvalidated-fixup-policy", "").AddStringEnum("passthrough", "strip", "error", "set").AddDefaultValue("strip").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("passthrough", "strip", "error", "set"),
				},
				Default: stringdefault.StaticString("strip"),
			},
			"default_map_value": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("An value that does not have a matching entry in the validation list is changed to this value if the no match policy is 'set'.", "unvalidated-fixup-map", "").AddNotValidWhen(models.NameValueProfileDefaultMapValueIgnoreVal.String()).String,
				Optional:            true,
			},
			"default_xss": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("This property allows the value to be checked for Cross Site Scripting (XSS) signatures. These signatures are malicious attempts to input client-side script as the input to a web application. If this client-side script is later displayed in a browser, the script executes and can perform malicious activities. Enable this feature to filter input for malicious content that might get stored and displayed again later, such as the contents of a comment form. The check looks for invalid characters and various forms of the term &lt;script that is often used to engage JavaScript on a browser without the user knowing.", "unvalidated-xss-check", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"no_match_xss_patterns_file": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specifies the patterns file that will be used by the XSS filter when No Match XSS is selected. The default file, store:///XSS-Patterns.xml, checks for invalid characters and various forms of the term &lt;script. Specify a custom XML patterns file with PCRE patterns to be used by the XSS filter.", "unvalidated-xss-patternsfile", "").AddDefaultValue("store:///XSS-Patterns.xml").AddRequiredWhen(models.NameValueProfileNoMatchXSSPatternsFileCondVal.String()).AddNotValidWhen(models.NameValueProfileNoMatchXSSPatternsFileIgnoreVal.String()).String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.NameValueProfileNoMatchXSSPatternsFileCondVal, models.NameValueProfileNoMatchXSSPatternsFileIgnoreVal, true),
				},
				Default: stringdefault.StaticString("store:///XSS-Patterns.xml"),
			},
			"dependency_actions": actions.ActionsSchema,
		},
	}
}

func (r *NameValueProfileResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (r *NameValueProfileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.NameValueProfile
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Create, false)
	if resp.Diagnostics.HasError() {
		return
	}

	body := data.ToBody(ctx, `NameValueProfile`)
	_, err := r.pData.Client.Post(data.GetPath(), body)
	if err != nil {
		if strings.Contains(err.Error(), "status 409") {
			_, err := r.pData.Client.Put(data.GetPath()+"/"+data.Id.ValueString(), body)
			if err != nil {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Resource already exists. Failed to update resource, got error: %s", err))
				return
			}
			resp.Diagnostics.AddWarning("Warning", "Resource already exists. Existing resource was updated.")
		} else {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to create resource, got error: %s", err))
			return
		}
	}
	actions.PostProcess(ctx, &resp.Diagnostics, data.DependencyActions, actions.Create)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NameValueProfileResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.NameValueProfile
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

	data.UpdateFromBody(ctx, `NameValueProfile`, res)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NameValueProfileResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.NameValueProfile
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Update, false)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.pData.Client.Put(data.GetPath()+"/"+data.Id.ValueString(), data.ToBody(ctx, `NameValueProfile`))
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

func (r *NameValueProfileResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.NameValueProfile
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Delete, false)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.pData.Client.Delete(data.GetPath() + "/" + data.Id.ValueString())
	if err != nil {
		if strings.Contains(err.Error(), "status 409") {
			resp.Diagnostics.AddWarning("Resource Conflict", fmt.Sprintf("Resource is no longer tracked by Terraform, but may need to be manually deleted on DataPower host. Got error: %s", err))
		} else if !strings.Contains(err.Error(), "status 404") {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete resource, got error: %s", err))
			return
		}
	}

	actions.PostProcess(ctx, &resp.Diagnostics, data.DependencyActions, actions.Delete)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.State.RemoveResource(ctx)
}

func (r *NameValueProfileResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()
	parts := strings.Split(req.ID, "/")
	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		resp.Diagnostics.AddError("Invalid Import ID", "Expected format: <app_domain>/<id>. Got: "+req.ID)
		return
	}

	appDomain := parts[0]
	id := parts[1]

	if !regexp.MustCompile("^[a-zA-Z0-9_-]+$").MatchString(id) || len(id) < 1 || len(id) > 128 {
		resp.Diagnostics.AddError("Invalid ID", "ID must be 1-128 characters and match pattern ^[a-zA-Z0-9_-]+$")
		return
	}
	if !regexp.MustCompile("^[a-zA-Z0-9_-]+$").MatchString(appDomain) || len(appDomain) < 1 || len(appDomain) > 128 {
		resp.Diagnostics.AddError("Invalid Application Domain", "Application domain must be 1-128 characters and match pattern ^[a-zA-Z0-9_-]+$")
		return
	}

	var data models.NameValueProfile
	data.AppDomain = types.StringValue(appDomain)
	data.Id = types.StringValue(id)
	res, err := r.pData.Client.Get(data.GetPath() + "/" + data.Id.ValueString())
	if err != nil {
		if strings.Contains(err.Error(), "status 404") {
			resp.Diagnostics.AddError("Resource Not Found", fmt.Sprintf("Resource was not found, got error: %s", err))
		} else {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		}
		return
	}

	data.FromBody(ctx, `NameValueProfile`, res)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NameValueProfileResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data models.NameValueProfile

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.ValidateConfig(ctx, &resp.Diagnostics, data.DependencyActions)
}

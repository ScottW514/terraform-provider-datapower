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

var _ resource.Resource = &ConfigSequenceResource{}

func NewConfigSequenceResource() resource.Resource {
	return &ConfigSequenceResource{}
}

type ConfigSequenceResource struct {
	pData *tfutils.ProviderData
}

func (r *ConfigSequenceResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_config_sequence"
}

func (r *ConfigSequenceResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("A configuration sequence defines a series of configuration files to load after the startup configuration. By default, changes in configuration files are detected and reloaded.", "config-sequence", "").String,
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
			"locations": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the locations to watch for changes and the permissions for each location. Each entry specifies a directory where the configuration files to match are stored. The DataPower Gateway watches the location and reloads the configuration when a change is detected that match the PCRE match pattern. The entries are processed in the listed order. The assess profile indicates the permissions for processing.", "location", "").String,
				NestedObject:        models.DmConfigSequenceLocationResourceSchema,
				Optional:            true,
			},
			"match_pattern": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the PCRE pattern to determine whether a file is considered part of the location match. For example, when the configuration files to match are <tt>NNNNNN.input</tt> , the PCRE pattern is <tt>\"([0-9]{6})\\.input$\"</tt> .", "match-pattern", "").AddDefaultValue("(.*)\\\\.cfg$").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("(.*)\\\\.cfg$"),
			},
			"result_name_pattern": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the PCRE pattern to name the result file. This pattern normally has a back-reference to the base input file name. For example, when input files are <tt>NNNNNN.input</tt> and the wanted result file name is <tt>NNNNNN.result</tt> , the pattern is <tt>\"$1.result\"</tt> .", "result-name-pattern", "").AddDefaultValue("$1.log").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("$1.log"),
			},
			"status_name_pattern": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the PCRE pattern to name the status file. This pattern normally has a back-reference to the base input file name. For example, when the input files are <tt>NNNNNN.input</tt> and the wanted status file name is <tt>NNNNNN.json</tt> , the pattern is <tt>\"$1.json\"</tt> .", "status-name-pattern", "").AddDefaultValue("$1.status").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("$1.status"),
			},
			"watch": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to watch the specified directory for configuration file changes and automatically reload the configuration when a change is detected. By default, the specified directory is watched.", "watch", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"use_output_location": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to place output log and status files in the configured output location after processing. By default, output files are stored in the input file location.", "use-output-location", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"output_location": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the directory to store the output files that processing generates. When not specified, the input file location is used.", "output-location", "").AddDefaultValue("logtemp:///").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("logtemp:///"),
			},
			"delete_unused": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to clean up objects that are no longer needed. When enabled, the configuration sequence detects and attempts to delete objects that are no longer modified by any configuration file. By default, the configuration sequence does not delete unneeded objects.", "delete-unused", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"run_sequence_interval": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the interval in milliseconds between the processing of changes. This delay enables multiple file events to be aggregated and handled within the same sequence run. Enter a value in the range 100 - 60000. The default value is 100.", "run-sequence-interval", "").AddIntegerRange(100, 60000).AddDefaultValue("100").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(100, 60000),
				},
				Default: int64default.StaticInt64(100),
			},
			"capabilities":       models.GetDmConfigSequenceCapabilitiesResourceSchema("Capabilities", "", "", false),
			"dependency_actions": actions.ActionsSchema,
		},
	}
}

func (r *ConfigSequenceResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (r *ConfigSequenceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.ConfigSequence
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Create, false)
	if resp.Diagnostics.HasError() {
		return
	}

	body := data.ToBody(ctx, `ConfigSequence`)
	_, err := r.pData.Client.Post(data.GetPath(), body)

	if err != nil && !strings.Contains(err.Error(), "status 409") {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to create object (%s), got error: %s", "POST", err))
		return
	}
	getRes, getErr := r.pData.Client.Get(data.GetPath() + "/" + data.Id.ValueString())
	if getErr != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object after creation (GET), got error: %s", getErr))
		return
	}
	data.UpdateUnknownFromBody(ctx, `ConfigSequence`, getRes)
	actions.PostProcess(ctx, &resp.Diagnostics, data.DependencyActions, actions.Create)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ConfigSequenceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.ConfigSequence
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
		data.FromBody(ctx, `ConfigSequence`, res)
	} else {
		// Update
		data.UpdateFromBody(ctx, `ConfigSequence`, res)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ConfigSequenceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.ConfigSequence
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Update, false)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.pData.Client.Put(data.GetPath()+"/"+data.Id.ValueString(), data.ToBody(ctx, `ConfigSequence`))
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

func (r *ConfigSequenceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.ConfigSequence
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

func (r *ConfigSequenceResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data models.ConfigSequence

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.ValidateConfig(ctx, &resp.Diagnostics, data.DependencyActions)
}

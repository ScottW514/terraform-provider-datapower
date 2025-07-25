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

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/client"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
)

var _ resource.Resource = &SLMRsrcClassResource{}

func NewSLMRsrcClassResource() resource.Resource {
	return &SLMRsrcClassResource{}
}

type SLMRsrcClassResource struct {
	client *client.DatapowerClient
}

func (r *SLMRsrcClassResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_slmrsrcclass"
}

func (r *SLMRsrcClassResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("SLM resource class", "slm-rsrc", "").String,

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
			"rsrc_type": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Resource type", "type", "").AddStringEnum("aaa-mapped-resource", "front-url", "destination-url", "xpath-filter", "request-message", "response-message", "soap-fault", "errorcode", "custom-stylesheet", "concurrent-connections", "concurrent-transactions", "wsdl", "wsdl-service", "wsdl-port", "wsdl-operation", "request-mq-qname", "reply-mq-qname", "uddi-subscription", "wsrr-subscription", "wsrr-saved-search-subscription").AddDefaultValue("aaa-mapped-resource").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("aaa-mapped-resource", "front-url", "destination-url", "xpath-filter", "request-message", "response-message", "soap-fault", "errorcode", "custom-stylesheet", "concurrent-connections", "concurrent-transactions", "wsdl", "wsdl-service", "wsdl-port", "wsdl-operation", "request-mq-qname", "reply-mq-qname", "uddi-subscription", "wsrr-subscription", "wsrr-saved-search-subscription"),
				},
				Default: stringdefault.StaticString("aaa-mapped-resource"),
			},
			"rsrc_match_type": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Match type", "match-type", "").AddStringEnum("per-extracted-value", "exact-match", "regexp-match").AddDefaultValue("per-extracted-value").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("per-extracted-value", "exact-match", "regexp-match"),
				},
				Default: stringdefault.StaticString("per-extracted-value"),
			},
			"rsrc_value": schema.ListAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Resource value", "value", "").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"stylesheet": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Custom stylesheet", "stylesheet", "").String,
				Optional:            true,
			},
			"x_path_filter": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("XPath filter", "xpath-filter", "").String,
				Optional:            true,
			},
			"subscription": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("UDDI subscription (deprecated)", "subscription", "").String,
				Optional:            true,
			},
			"wsrr_subscription": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("WSRR subscription", "wsrr-subscription", "wsrrsubscription").String,
				Optional:            true,
			},
			"wsrr_saved_search_subscription": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("WSRR saved search subscription", "wsrr-saved-search-subscription", "wsrrsavedsearchsubscription").String,
				Optional:            true,
			},
		},
	}
}

func (r *SLMRsrcClassResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = *req.ProviderData.(**client.DatapowerClient)
}

func (r *SLMRsrcClassResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data models.SLMRsrcClass

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	body := data.ToBody(ctx, `SLMRsrcClass`)
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

func (r *SLMRsrcClassResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data models.SLMRsrcClass

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
		data.FromBody(ctx, `SLMRsrcClass`, res)
	} else {
		// Update
		data.UpdateFromBody(ctx, `SLMRsrcClass`, res)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SLMRsrcClassResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data models.SLMRsrcClass

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.client.Put(data.GetPath()+"/"+data.Id.ValueString(), data.ToBody(ctx, `SLMRsrcClass`))
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

func (r *SLMRsrcClassResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data models.SLMRsrcClass

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.client.Delete(data.GetPath() + "/" + data.Id.ValueString())
	if err != nil && !strings.Contains(err.Error(), "status 404") && !strings.Contains(err.Error(), "status 409") {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object (DELETE), got error: %s", err))
		return
	}

	resp.State.RemoveResource(ctx)
}

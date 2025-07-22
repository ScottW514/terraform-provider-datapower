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
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/client"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
)

var _ resource.Resource = &XMLManagerResource{}

func NewXMLManagerResource() resource.Resource {
	return &XMLManagerResource{}
}

type XMLManagerResource struct {
	client *client.DatapowerClient
}

func (r *XMLManagerResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_xmlmanager"
}

func (r *XMLManagerResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("XML Manager", "xmlmgr", "").String,

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
			"url_refresh_policy": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("URL Refresh Policy", "xslrefresh", "urlrefreshpolicy").String,
				Optional:            true,
			},
			"compile_options_policy": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Compile Options Policy", "xslconfig", "compileoptionspolicy").String,
				Optional:            true,
			},
			"cache_memory_size": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Stylesheet cache size", "xsl cache memorysize", "").AddDefaultValue("2147483647").String,
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(2147483647),
			},
			"cache_size": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Stylesheet cache count", "xsl cache size", "").AddIntegerRange(5, 250000).AddDefaultValue("256").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(5, 250000),
				},
				Default: int64default.StaticInt64(256),
			},
			"sha1_caching": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("SHA1 Caching", "xsl checksummed cache", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"static_document_calls": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Static Document Call", "static-document-calls", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"search_results": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("XSLT Expression Optimization", "search results", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"virtual_servers": schema.ListAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Load Balancer Groups", "loadbalancer-group", "loadbalancergroup").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"parser_limits_bytes_scanned": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("XML Bytes Scanned", "bytes-scanned", "").AddDefaultValue("4194304").String,
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(4194304),
			},
			"parser_limits_element_depth": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("XML Element Depth", "element-depth", "").AddDefaultValue("512").String,
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(512),
			},
			"parser_limits_attribute_count": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("XML Attribute Count", "attribute-count", "").AddDefaultValue("128").String,
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(128),
			},
			"parser_limits_max_node_size": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("XML Maximum Node Size", "max-node-size", "").AddIntegerRange(1024, 4294967295).AddDefaultValue("33554432").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(1024, 4294967295),
				},
				Default: int64default.StaticInt64(33554432),
			},
			"parser_limits_forbid_external_references": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("XML Forbid External References", "forbid-external-references", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"parser_limits_external_references": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("XML External Reference Handling", "external-references", "").AddStringEnum("forbid", "ignore", "allow").AddDefaultValue("forbid").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("forbid", "ignore", "allow"),
				},
				Default: stringdefault.StaticString("forbid"),
			},
			"parser_limits_max_prefixes": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("XML Maximum Distinct Prefixes", "max-prefixes", "").AddIntegerRange(0, 262143).AddDefaultValue("1024").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(0, 262143),
				},
				Default: int64default.StaticInt64(1024),
			},
			"parser_limits_max_namespaces": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("XML Maximum Distinct Namespaces", "max-namespaces", "").AddIntegerRange(0, 65535).AddDefaultValue("1024").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(0, 65535),
				},
				Default: int64default.StaticInt64(1024),
			},
			"parser_limits_max_local_names": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("XML Maximum Distinct Local Names", "max-local-names", "").AddIntegerRange(0, 1048575).AddDefaultValue("60000").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(0, 1048575),
				},
				Default: int64default.StaticInt64(60000),
			},
			"doc_cache_max_docs": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Document Cache Count", "maxdocs", "").AddIntegerRange(1, 250000).AddDefaultValue("5000").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(1, 250000),
				},
				Default: int64default.StaticInt64(5000),
			},
			"doc_cache_size": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Document Cache Size", "size", "").String,
				Optional:            true,
			},
			"doc_max_writes": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Maximum concurrent writes", "max-writes", "").AddIntegerRange(1, 32768).AddDefaultValue("32768").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(1, 32768),
				},
				Default: int64default.StaticInt64(32768),
			},
			"doc_cache_policy": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Document Cache Policy", "policy", "").String,
				NestedObject:        models.DmDocCachePolicyResourceSchema,
				Optional:            true,
			},
			"schema_validation": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Schema Validation Rules", "validate", "").String,
				NestedObject:        models.DmSchemaValidationResourceSchema,
				Optional:            true,
			},
			"scheduled_rule": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Scheduled Processing Policy Rule", "schedule-rule", "").String,
				NestedObject:        models.DmScheduledRuleResourceSchema,
				Optional:            true,
			},
			"user_agent": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("User Agent Configuration", "user-agent", "httpuseragent").AddDefaultValue("default").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("default"),
			},
			"json_parser_settings": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("JSON Settings", "json-parser-settings", "jsonsettings").String,
				Optional:            true,
			},
			"ldap_conn_pool": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("LDAP Connection Pool", "ldap-pool", "ldapconnectionpool").String,
				Optional:            true,
			},
		},
	}
}

func (r *XMLManagerResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = *req.ProviderData.(**client.DatapowerClient)
}

func (r *XMLManagerResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data models.XMLManager

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	body := data.ToBody(ctx, `XMLManager`)
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

func (r *XMLManagerResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data models.XMLManager

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
		data.FromBody(ctx, `XMLManager`, res)
	} else {
		// Update
		data.UpdateFromBody(ctx, `XMLManager`, res)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *XMLManagerResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data models.XMLManager

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.client.Put(data.GetPath()+"/"+data.Id.ValueString(), data.ToBody(ctx, `XMLManager`))
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

func (r *XMLManagerResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data models.XMLManager

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

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

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/client"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
)

type XMLManagerList struct {
	AppDomain types.String `tfsdk:"app_domain"`
	Result    types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &XMLManagerDataSource{}
	_ datasource.DataSourceWithConfigure = &XMLManagerDataSource{}
)

func NewXMLManagerDataSource() datasource.DataSource {
	return &XMLManagerDataSource{}
}

type XMLManagerDataSource struct {
	client *client.DatapowerClient
}

func (d *XMLManagerDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_xmlmanager"
}

func (d *XMLManagerDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "XML Manager",
		Attributes: map[string]schema.Attribute{
			"app_domain": schema.StringAttribute{
				MarkdownDescription: "The name of the application domain the object belongs to",
				Required:            true,
			},
			"result": schema.ListNestedAttribute{
				MarkdownDescription: "List of objects",
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
						"url_refresh_policy": schema.StringAttribute{
							MarkdownDescription: "URL Refresh Policy",
							Computed:            true,
						},
						"compile_options_policy": schema.StringAttribute{
							MarkdownDescription: "Compile Options Policy",
							Computed:            true,
						},
						"cache_memory_size": schema.Int64Attribute{
							MarkdownDescription: "Stylesheet cache size",
							Computed:            true,
						},
						"cache_size": schema.Int64Attribute{
							MarkdownDescription: "Stylesheet cache count",
							Computed:            true,
						},
						"sha1_caching": schema.BoolAttribute{
							MarkdownDescription: "SHA1 Caching",
							Computed:            true,
						},
						"static_document_calls": schema.BoolAttribute{
							MarkdownDescription: "Static Document Call",
							Computed:            true,
						},
						"search_results": schema.BoolAttribute{
							MarkdownDescription: "XSLT Expression Optimization",
							Computed:            true,
						},
						"support_tx_warn": schema.BoolAttribute{
							MarkdownDescription: "Support ITX Warnings",
							Computed:            true,
						},
						"virtual_servers": schema.ListAttribute{
							MarkdownDescription: "Load Balancer Groups",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"parser_limits_bytes_scanned": schema.Int64Attribute{
							MarkdownDescription: "XML Bytes Scanned",
							Computed:            true,
						},
						"parser_limits_element_depth": schema.Int64Attribute{
							MarkdownDescription: "XML Element Depth",
							Computed:            true,
						},
						"parser_limits_attribute_count": schema.Int64Attribute{
							MarkdownDescription: "XML Attribute Count",
							Computed:            true,
						},
						"parser_limits_max_node_size": schema.Int64Attribute{
							MarkdownDescription: "XML Maximum Node Size",
							Computed:            true,
						},
						"parser_limits_forbid_external_references": schema.BoolAttribute{
							MarkdownDescription: "XML Forbid External References",
							Computed:            true,
						},
						"parser_limits_external_references": schema.StringAttribute{
							MarkdownDescription: "XML External Reference Handling",
							Computed:            true,
						},
						"parser_limits_max_prefixes": schema.Int64Attribute{
							MarkdownDescription: "XML Maximum Distinct Prefixes",
							Computed:            true,
						},
						"parser_limits_max_namespaces": schema.Int64Attribute{
							MarkdownDescription: "XML Maximum Distinct Namespaces",
							Computed:            true,
						},
						"parser_limits_max_local_names": schema.Int64Attribute{
							MarkdownDescription: "XML Maximum Distinct Local Names",
							Computed:            true,
						},
						"doc_cache_max_docs": schema.Int64Attribute{
							MarkdownDescription: "Document Cache Count",
							Computed:            true,
						},
						"doc_cache_size": schema.Int64Attribute{
							MarkdownDescription: "Document Cache Size",
							Computed:            true,
						},
						"doc_max_writes": schema.Int64Attribute{
							MarkdownDescription: "Maximum concurrent writes",
							Computed:            true,
						},
						"doc_cache_policy": schema.ListNestedAttribute{
							MarkdownDescription: "Document Cache Policy",
							NestedObject:        models.DmDocCachePolicyDataSourceSchema,
							Computed:            true,
						},
						"schema_validation": schema.ListNestedAttribute{
							MarkdownDescription: "Schema Validation Rules",
							NestedObject:        models.DmSchemaValidationDataSourceSchema,
							Computed:            true,
						},
						"scheduled_rule": schema.ListNestedAttribute{
							MarkdownDescription: "Scheduled Processing Policy Rule",
							NestedObject:        models.DmScheduledRuleDataSourceSchema,
							Computed:            true,
						},
						"user_agent": schema.StringAttribute{
							MarkdownDescription: "User Agent Configuration",
							Computed:            true,
						},
						"json_parser_settings": schema.StringAttribute{
							MarkdownDescription: "JSON Settings",
							Computed:            true,
						},
						"ldap_conn_pool": schema.StringAttribute{
							MarkdownDescription: "LDAP Connection Pool",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *XMLManagerDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = *req.ProviderData.(**client.DatapowerClient)
}

func (d *XMLManagerDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data XMLManagerList

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.XMLManager{
		AppDomain: data.AppDomain,
	}

	res, err := d.client.Get(o.GetPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}
	l := []models.XMLManager{}
	if value := res.Get(`XMLManager`); value.Exists() {
		for _, v := range value.Array() {
			item := models.XMLManager{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.XMLManagerObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.XMLManagerObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

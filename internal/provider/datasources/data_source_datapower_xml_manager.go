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
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
)

type XMLManagerList struct {
	Id        types.String `tfsdk:"id"`
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
	pData *tfutils.ProviderData
}

func (d *XMLManagerDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_xml_manager"
}

func (d *XMLManagerDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "An XML Manager manages the compilation and caching of stylesheets, the caching of documents, and provides configuration constraints on the size and parsing depth of documents. You can enable streaming operation by configuring an XML Manager to use a Streaming Compile Option Policy. More than one firewall can use the same XML Manager.",
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
						"url_refresh_policy": schema.StringAttribute{
							MarkdownDescription: "Optionally, assigns a Stylesheet Refresh Policy to the XML manager. Stylesheets cached by this manager are refreshed in accordance with policy rules.",
							Computed:            true,
						},
						"compile_options_policy": schema.StringAttribute{
							MarkdownDescription: "<p>Optionally, assigns a Compilation Options Policy to the XML Manager. Stylesheets cached by this manager are compiled in accordance with policy rules. Compile Options allow the administrator to manage profiling and debug modes, which are helpful during development and troubleshooting sessions. Profiling and debug modes should not be left on for production purposes.</p><p>Compile Options Policy also enables streaming mode, which might be a valid choice for production environments that regularly encounter very large documents.</p>",
							Computed:            true,
						},
						"cache_memory_size": schema.Int64Attribute{
							MarkdownDescription: "Enter the maximum size of the stylesheet cache. The default value is 2147483647. A value of 0 disables caching. Stylesheets are purged when either the cache size or the cache count is reached.",
							Computed:            true,
						},
						"cache_size": schema.Int64Attribute{
							MarkdownDescription: "Enter the maximum number of stylesheets to cache. Enter a value in the range 5 - 250000. The default value is 256. Stylesheets are purged when either the cache size or the cache count is reached.",
							Computed:            true,
						},
						"sha1caching": schema.BoolAttribute{
							MarkdownDescription: "Enables/disables SHA1-assisted stylesheet caching. With SHA1 caching enabled, stylesheets are cached by both URL and SHA1 message digest value. With SHA1 caching disabled, stylesheets are cached only by URL.",
							Computed:            true,
						},
						"static_document_calls": schema.BoolAttribute{
							MarkdownDescription: "The latest XSLT specifications require that multiple document calls in the same transformation return the same result. Disable to allow all document calls to operate independently.",
							Computed:            true,
						},
						"search_results": schema.BoolAttribute{
							MarkdownDescription: "Configures optimization of '//' in XPath expressions",
							Computed:            true,
						},
						"virtual_servers": schema.ListAttribute{
							MarkdownDescription: "Select a Load Balancer Group and assign that group to this XML manager. A Load Balancer Group, or server pool, provides redundancy among backside resources.",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"parser_limits_bytes_scanned": schema.Int64Attribute{
							MarkdownDescription: "Enter the maximum number of bytes scanned in one message. This property applies to any XML document that is parsed. If the value is 0, no limit is enforced. The default value is 4194304.",
							Computed:            true,
						},
						"parser_limits_element_depth": schema.Int64Attribute{
							MarkdownDescription: "Enter the maximum depth of element nesting in XML parser. The default value is 512.",
							Computed:            true,
						},
						"parser_limits_attribute_count": schema.Int64Attribute{
							MarkdownDescription: "Enter the maximum number of attributes of a given element. The default value is 128.",
							Computed:            true,
						},
						"parser_limits_max_node_size": schema.Int64Attribute{
							MarkdownDescription: "Enter the maximum size that any one node can consume. Enter a value in the range 1024 - 4294967295. The default value is 33554432. Sizes which are powers of two result in the best performance. Although you define a value, the DataPower Gateway uses a value that is the rounded-down to the largest power of two that is smaller than the defined value.",
							Computed:            true,
						},
						"parser_limits_external_references": schema.StringAttribute{
							MarkdownDescription: "Select the special handling for input documents that contain external references, such as an external entity or external DTD definition.",
							Computed:            true,
						},
						"parser_limits_max_prefixes": schema.Int64Attribute{
							MarkdownDescription: "Enter the maximum number of distinct XML namespace prefixes in a document. This limit counts multiple prefixes defined for the same namespace, but does not count multiple namespaces defined in different parts of the input document under a single prefix. Enter a value in the range 0 - 262143. The default value is 1024. A value of 0 indicates that the limit is 1024.",
							Computed:            true,
						},
						"parser_limits_max_namespaces": schema.Int64Attribute{
							MarkdownDescription: "Enter the maximum number of distinct XML namespace URIs in a document. This limit counts all XML namespaces, regardless of how many prefixes are used to declare them. Enter a value in the range 0 - 65535. The default value is 1024. A value of 0 indicates that the limit is 1024.",
							Computed:            true,
						},
						"parser_limits_max_local_names": schema.Int64Attribute{
							MarkdownDescription: "Enter the maximum number of distinct XML local names in a document. This limit counts all local names, independent of the namespaces they are associated with. Enter a value in the range 0 - 1048575. The default value is 60000. A value of 0 indicates that the limit is 60000.",
							Computed:            true,
						},
						"doc_cache_max_docs": schema.Int64Attribute{
							MarkdownDescription: "Enter the maximum number of documents to cache. Enter a value in the range 1 - 250000. The default value is 5000.",
							Computed:            true,
						},
						"doc_cache_size": schema.Int64Attribute{
							MarkdownDescription: "Enter the maximum size of the document cache. Regardless of the specified size, no document that is greater than 1073741824 bytes is cached. This restriction applies even if the cache has available space.",
							Computed:            true,
						},
						"doc_max_writes": schema.Int64Attribute{
							MarkdownDescription: "Enter the maximum number of concurrent write requests to create documents or refresh expired documents in the document cache. Enter a value in the range 1 - 32768. The default value is 32768. After the maximum number is reached, requests are forwarded to the target server and the response is not written to the cache.",
							Computed:            true,
						},
						"doc_cache_policy": schema.ListNestedAttribute{
							MarkdownDescription: "A document cache policy allows the administrator to determine how documents are cached by the XML manager. The policy offers time-to-live, priority, and type configuration values. This document cache is distinct from the stylesheet cache for the XML manager. Documents and stylesheets that the XML manager retrieves can be cached.",
							NestedObject:        models.GetDmDocCachePolicyDataSourceSchema(),
							Computed:            true,
						},
						"schema_validation": schema.ListNestedAttribute{
							MarkdownDescription: "",
							NestedObject:        models.GetDmSchemaValidationDataSourceSchema(),
							Computed:            true,
						},
						"scheduled_rule": schema.ListNestedAttribute{
							MarkdownDescription: "Certain applications might require the running of a scheduled processing rule. Integration with a CA Unicenter Manager is facilitated by a regularly scheduled processing rule that obtains relationship data from the Unicenter Manager.",
							NestedObject:        models.GetDmScheduledRuleDataSourceSchema(),
							Computed:            true,
						},
						"user_agent": schema.StringAttribute{
							MarkdownDescription: "Identifies the User Agent configuration the XML manager uses to connect to external servers to retrieve resources.",
							Computed:            true,
						},
						"json_parser_settings": schema.StringAttribute{
							MarkdownDescription: "Identifies the JSON Settings used by the XML manager when processing a JSON body. If no JSON Settings are specified, the default limits are enforced.",
							Computed:            true,
						},
						"ldap_conn_pool": schema.StringAttribute{
							MarkdownDescription: "Identifies the LDAP connection pool to be used by the XML Manager when contacting LDAP",
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

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *XMLManagerDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data XMLManagerList
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.XMLManager{
		AppDomain: data.AppDomain,
	}

	path := o.GetPath()
	if !data.Id.IsNull() {
		path = path + "/" + data.Id.ValueString()
	}

	res, err := d.pData.Client.Get(path)
	resFound := true
	if err != nil {
		if !strings.Contains(err.Error(), "status 404") {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
			return
		} else {
			resFound = false
		}
	}
	l := []models.XMLManager{}
	if resFound {
		if value := res.Get(`XMLManager`); value.Exists() {
			for _, v := range value.Array() {
				item := models.XMLManager{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
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

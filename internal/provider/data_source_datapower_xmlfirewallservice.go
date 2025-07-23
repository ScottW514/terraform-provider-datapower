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

type XMLFirewallServiceList struct {
	AppDomain types.String `tfsdk:"app_domain"`
	Result    types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &XMLFirewallServiceDataSource{}
	_ datasource.DataSourceWithConfigure = &XMLFirewallServiceDataSource{}
)

func NewXMLFirewallServiceDataSource() datasource.DataSource {
	return &XMLFirewallServiceDataSource{}
}

type XMLFirewallServiceDataSource struct {
	client *client.DatapowerClient
}

func (d *XMLFirewallServiceDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_xmlfirewallservice"
}

func (d *XMLFirewallServiceDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "XML Firewall",
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
						"type": schema.StringAttribute{
							MarkdownDescription: "Type",
							Computed:            true,
						},
						"xml_manager": schema.StringAttribute{
							MarkdownDescription: "XML Manager",
							Computed:            true,
						},
						"url_rewrite_policy": schema.StringAttribute{
							MarkdownDescription: "URL Rewrite Policy",
							Computed:            true,
						},
						"ssl_proxy": schema.StringAttribute{
							MarkdownDescription: "TLS proxy profile (deprecated)",
							Computed:            true,
						},
						"style_policy": schema.StringAttribute{
							MarkdownDescription: "Processing Policy",
							Computed:            true,
						},
						"max_message_size": schema.Int64Attribute{
							MarkdownDescription: "Maximum Message Size",
							Computed:            true,
						},
						"request_type": schema.StringAttribute{
							MarkdownDescription: "Request Type",
							Computed:            true,
						},
						"response_type": schema.StringAttribute{
							MarkdownDescription: "Response Type",
							Computed:            true,
						},
						"fw_cred": schema.StringAttribute{
							MarkdownDescription: "Firewall Credentials",
							Computed:            true,
						},
						"service_monitors": schema.ListAttribute{
							MarkdownDescription: "Service Monitors",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"request_attachments": schema.StringAttribute{
							MarkdownDescription: "Request attachment processing mode",
							Computed:            true,
						},
						"response_attachments": schema.StringAttribute{
							MarkdownDescription: "Response attachment processing mode",
							Computed:            true,
						},
						"root_part_not_first_action": schema.StringAttribute{
							MarkdownDescription: "Action when required root part is not first",
							Computed:            true,
						},
						"front_attachment_format": schema.StringAttribute{
							MarkdownDescription: "Front attachment processing format",
							Computed:            true,
						},
						"back_attachment_format": schema.StringAttribute{
							MarkdownDescription: "Back attachment processing format",
							Computed:            true,
						},
						"mime_headers": schema.BoolAttribute{
							MarkdownDescription: "MIME Header Processing",
							Computed:            true,
						},
						"rewrite_errors": schema.BoolAttribute{
							MarkdownDescription: "Rewrite Error Messages",
							Computed:            true,
						},
						"delay_errors": schema.BoolAttribute{
							MarkdownDescription: "Delay Error Messages",
							Computed:            true,
						},
						"delay_errors_duration": schema.Int64Attribute{
							MarkdownDescription: "Duration to Delay Error Messages",
							Computed:            true,
						},
						"soap_schema_url": schema.StringAttribute{
							MarkdownDescription: "SOAP Schema URL",
							Computed:            true,
						},
						"wsdl_response_policy": schema.StringAttribute{
							MarkdownDescription: "WSDL Response Policy",
							Computed:            true,
						},
						"wsdl_file_location": schema.StringAttribute{
							MarkdownDescription: "WSDL file location",
							Computed:            true,
						},
						"firewall_parser_limits": schema.BoolAttribute{
							MarkdownDescription: "Firewall parser limits",
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
						"parser_limits_attachment_byte_count": schema.Int64Attribute{
							MarkdownDescription: "Attachment Byte Count Limit",
							Computed:            true,
						},
						"parser_limits_attachment_package_byte_count": schema.Int64Attribute{
							MarkdownDescription: "Attachment Package Byte Count Limit",
							Computed:            true,
						},
						"parser_limits_external_references": schema.StringAttribute{
							MarkdownDescription: "XML External Reference Handling",
							Computed:            true,
						},
						"credential_charset": schema.StringAttribute{
							MarkdownDescription: "Credential Character Set",
							Computed:            true,
						},
						"ssl_config_type": schema.StringAttribute{
							MarkdownDescription: "TLS type",
							Computed:            true,
						},
						"ssl_server": schema.StringAttribute{
							MarkdownDescription: "TLS server profile",
							Computed:            true,
						},
						"sslsni_server": schema.StringAttribute{
							MarkdownDescription: "TLS SNI server profile",
							Computed:            true,
						},
						"ssl_client": schema.StringAttribute{
							MarkdownDescription: "TLS client profile",
							Computed:            true,
						},
						"user_summary": schema.StringAttribute{
							MarkdownDescription: "Comments",
							Computed:            true,
						},
						"priority": schema.StringAttribute{
							MarkdownDescription: "Service Priority",
							Computed:            true,
						},
						"local_port": schema.Int64Attribute{
							MarkdownDescription: "Port Number",
							Computed:            true,
						},
						"remote_address": schema.StringAttribute{
							MarkdownDescription: "Remote Host",
							Computed:            true,
						},
						"remote_port": schema.Int64Attribute{
							MarkdownDescription: "Remote Port",
							Computed:            true,
						},
						"acl": schema.StringAttribute{
							MarkdownDescription: "Access Control List",
							Computed:            true,
						},
						"http_timeout": schema.Int64Attribute{
							MarkdownDescription: "HTTP Timeout",
							Computed:            true,
						},
						"http_persist_timeout": schema.Int64Attribute{
							MarkdownDescription: "HTTP Persistent Timeout",
							Computed:            true,
						},
						"do_host_rewrite": schema.BoolAttribute{
							MarkdownDescription: "Host Rewrite",
							Computed:            true,
						},
						"suppress_http_warnings": schema.BoolAttribute{
							MarkdownDescription: "HTTP Warning Suppression",
							Computed:            true,
						},
						"http_compression": schema.BoolAttribute{
							MarkdownDescription: "HTTP Compression",
							Computed:            true,
						},
						"http_include_response_type_encoding": schema.BoolAttribute{
							MarkdownDescription: "HTTP Include charset in response-type",
							Computed:            true,
						},
						"always_show_errors": schema.BoolAttribute{
							MarkdownDescription: "Always provide full errors",
							Computed:            true,
						},
						"disallow_get": schema.BoolAttribute{
							MarkdownDescription: "Disallow GET (and HEAD)",
							Computed:            true,
						},
						"disallow_empty_response": schema.BoolAttribute{
							MarkdownDescription: "Don't allow empty response bodies",
							Computed:            true,
						},
						"http_persistent_connections": schema.BoolAttribute{
							MarkdownDescription: "HTTP Persistent Connections",
							Computed:            true,
						},
						"http_client_ip_label": schema.StringAttribute{
							MarkdownDescription: "HTTP Client IP Label",
							Computed:            true,
						},
						"http_log_cor_id_label": schema.StringAttribute{
							MarkdownDescription: "HTTP Global Transaction ID Label",
							Computed:            true,
						},
						"http_proxy_host": schema.StringAttribute{
							MarkdownDescription: "Proxy Host",
							Computed:            true,
						},
						"http_proxy_port": schema.Int64Attribute{
							MarkdownDescription: "Proxy Port",
							Computed:            true,
						},
						"http_version": models.GetDmHTTPClientServerVersionDataSourceSchema("HTTP Version", "version", ""),
						"do_chunked_upload": schema.BoolAttribute{
							MarkdownDescription: "Allow Chunked Uploads",
							Computed:            true,
						},
						"header_injection": schema.ListNestedAttribute{
							MarkdownDescription: "HTTP Header Injection",
							NestedObject:        models.DmHeaderInjectionDataSourceSchema,
							Computed:            true,
						},
						"header_suppression": schema.ListNestedAttribute{
							MarkdownDescription: "HTTP Header Suppression",
							NestedObject:        models.DmHeaderSuppressionDataSourceSchema,
							Computed:            true,
						},
						"stylesheet_parameters": schema.ListNestedAttribute{
							MarkdownDescription: "Stylesheet Parameter",
							NestedObject:        models.DmStylesheetParameterDataSourceSchema,
							Computed:            true,
						},
						"default_param_namespace": schema.StringAttribute{
							MarkdownDescription: "Default parameter namespace",
							Computed:            true,
						},
						"query_param_namespace": schema.StringAttribute{
							MarkdownDescription: "Query parameter namespace",
							Computed:            true,
						},
						"force_policy_exec": schema.BoolAttribute{
							MarkdownDescription: "Process Messages Whose Body Is Empty",
							Computed:            true,
						},
						"count_monitors": schema.ListAttribute{
							MarkdownDescription: "Count Monitors",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"duration_monitors": schema.ListAttribute{
							MarkdownDescription: "Duration Monitors",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"monitor_processing_policy": schema.StringAttribute{
							MarkdownDescription: "Monitors evaluation method",
							Computed:            true,
						},
						"debug_mode": schema.StringAttribute{
							MarkdownDescription: "Probe setting",
							Computed:            true,
						},
						"debug_history": schema.Int64Attribute{
							MarkdownDescription: "Transaction History",
							Computed:            true,
						},
						"debug_trigger": schema.ListNestedAttribute{
							MarkdownDescription: "Probe Triggers",
							NestedObject:        models.DmMSDebugTriggerTypeDataSourceSchema,
							Computed:            true,
						},
						"local_address": schema.StringAttribute{
							MarkdownDescription: "Local address",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *XMLFirewallServiceDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = *req.ProviderData.(**client.DatapowerClient)
}

func (d *XMLFirewallServiceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data XMLFirewallServiceList

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.XMLFirewallService{
		AppDomain: data.AppDomain,
	}

	res, err := d.client.Get(o.GetPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}
	l := []models.XMLFirewallService{}
	if value := res.Get(`XMLFirewallService`); value.Exists() {
		for _, v := range value.Array() {
			item := models.XMLFirewallService{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.XMLFirewallServiceObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.XMLFirewallServiceObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

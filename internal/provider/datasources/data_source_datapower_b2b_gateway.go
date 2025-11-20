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
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
)

type B2BGatewayList struct {
	Id        types.String `tfsdk:"id"`
	AppDomain types.String `tfsdk:"app_domain"`
	Result    types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &B2BGatewayDataSource{}
	_ datasource.DataSourceWithConfigure = &B2BGatewayDataSource{}
)

func NewB2BGatewayDataSource() datasource.DataSource {
	return &B2BGatewayDataSource{}
}

type B2BGatewayDataSource struct {
	pData *tfutils.ProviderData
}

func (d *B2BGatewayDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_b2b_gateway"
}

func (d *B2BGatewayDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "A B2B gateway can handle B2B messages transmitted over a variety of AS and non-AS protocols for inbound and outbound flows, as well as MDN messages.",
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
						"priority": schema.StringAttribute{
							MarkdownDescription: "Specify the priority of service-scheduling. When system resources are in high demand, high priority services are favored over lower priority services.",
							Computed:            true,
						},
						"doc_store_location": schema.StringAttribute{
							MarkdownDescription: "Specify the location for document storage, which saves copies of inbound, outbound, and intermediate documents that might be needed for a retransmit operation. If unspecified, documents are stored in the encrypted area on the RAID volume. For storage on the RAID volume, set the maximum disk usage to 30 GB.",
							Computed:            true,
						},
						"as_front_protocol": schema.ListNestedAttribute{
							MarkdownDescription: "Protocol handlers",
							NestedObject:        models.GetDmASFrontProtocolDataSourceSchema(),
							Computed:            true,
						},
						"as1_mdn_email": schema.StringAttribute{
							MarkdownDescription: "For AS1 asynchronous MDN scenarios, specify the default email address for the AS1 asynchronous MDN. When sending outbound AS1 email messages that request an MDN, this email address can be the default email address for the response MDN. An email address in the destination overrides this value.",
							Computed:            true,
						},
						"as1_mdn_smtp_server_connection": schema.StringAttribute{
							MarkdownDescription: "When an incoming AS1 message requests an MDN as an email request, specify the SMTP server connection for asynchronous MDN responses.",
							Computed:            true,
						},
						"as2_mdn_url": schema.StringAttribute{
							MarkdownDescription: "For AS2 asynchronous MDN scenarios, specify the default URL for the AS2 asynchronous MDN. This URL can point to that gateway itself or a firewall or proxy that routes the message to the gateway. A URL in the destination overrides this value.",
							Computed:            true,
						},
						"as3_mdn_url": schema.StringAttribute{
							MarkdownDescription: "For AS3 asynchronous MDN scenarios, specify the default URL for the AS3 asynchronous MDN. This URL can point to that gateway itself or a firewall or proxy that routes the message to the gateway. A URL in the destination overrides this value.",
							Computed:            true,
						},
						"b2b_profiles": schema.ListNestedAttribute{
							MarkdownDescription: "Active partner profiles",
							NestedObject:        models.GetDmB2BActiveProfileDataSourceSchema(),
							Computed:            true,
						},
						"b2b_groups": schema.ListNestedAttribute{
							MarkdownDescription: "Active profile groups",
							NestedObject:        models.GetDmB2BActiveGroupDataSourceSchema(),
							Computed:            true,
						},
						"document_routing_preprocessor_type": schema.StringAttribute{
							MarkdownDescription: "Specify the file type of the document-routing preprocessor file to run against messages that are not received through AS or ebMS protocols. The default value is stylesheet.",
							Computed:            true,
						},
						"document_routing_preprocessor": schema.StringAttribute{
							MarkdownDescription: "Specify the location of the document-routing preprocessor file, which is a stylesheet or a GatewayScript file to run against transactions that cannot be handled by B2B-specific handlers. <ul><li>A stylesheet examines information from transport headers and other non-content sources to select relevant trading partners. The default stylesheet is <tt>store:///b2b-routing.xsl</tt> .</li><li>A GatewayScript examines information from transport headers and payloads to select relevant trading partners. A GatewayScript can parse messages of different data types, such as JSON, XML, and non-XML. A sample GatewayScript file is <tt>store:///gateayscript/example-b2b-routing.js</tt> .</li></ul>",
							Computed:            true,
						},
						"document_routing_preprocessor_debug": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to enable the GatewayScript debugger to detect and diagnose errors in the document routing preprocessor file. To debug a file, ensure that the file contains one or more <tt>debugger;</tt> statements where you want to start debugging.",
							Computed:            true,
						},
						"archive_mode": schema.StringAttribute{
							MarkdownDescription: "Purge mode",
							Computed:            true,
						},
						"archive_location": schema.StringAttribute{
							MarkdownDescription: "Specify the location for archive file. Enter the fully qualified name of the directory. To copy the archive file to an FTP server, ensure that the FTP policies in the XML manager enable image (binary) data transfer.",
							Computed:            true,
						},
						"archive_file_name": schema.StringAttribute{
							MarkdownDescription: "Specify the base file name for archive file. When archiving, the operation appends the current timestamp.",
							Computed:            true,
						},
						"archive_minimum_size": schema.Int64Attribute{
							MarkdownDescription: "Specify the minimum remaining size in KB of document storage that triggers archival. The default value is 1024.",
							Computed:            true,
						},
						"archive_document_age": schema.Int64Attribute{
							MarkdownDescription: "Specify the maximum duration in days to retain processed documents. Enter a value in the range 1 - 3650. The default value is 90.",
							Computed:            true,
						},
						"archive_minimum_documents": schema.Int64Attribute{
							MarkdownDescription: "Specify the minimum number of documents to retain in document storage after archival. The minimum value is 1. The default value is 100.",
							Computed:            true,
						},
						"disk_use_check_interval": schema.Int64Attribute{
							MarkdownDescription: "Specify the interval in minutes between checks for documents that exceed the maximum age. During the check, documents that exceed the maximum age are purged. Enter a value in the range 1 - 1440. The default value is 60.",
							Computed:            true,
						},
						"max_document_disk_use": schema.Int64Attribute{
							MarkdownDescription: "Specify the maximum size in KB for document storage. When storage exceeds this value, documents are purged. The default value is 25165824.",
							Computed:            true,
						},
						"archive_monitor": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to use a monitor during archival. The monitor limits the message injection rate to prevent problems in a critical situation like performance testing or on a heavily loaded system.",
							Computed:            true,
						},
						"shaping_threshold": schema.Int64Attribute{
							MarkdownDescription: "Specify the maximum TPS to allow during archival. When the threshold is reached, the service queues transactions. When the queue is full, the service rejects transactions and generates a log message. Enter a value in the range 10 - 10000. The default value is 200.",
							Computed:            true,
						},
						"archive_backup_documents": models.GetDmB2BBackupMsgTypeDataSourceSchema("Specify the types of documents to archive. This property does not indicate the inbound or outbound transaction to archive.", "arch-backup-documents", ""),
						"xpath_routing_policies": schema.ListAttribute{
							MarkdownDescription: "XPath routing policies",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"xml_manager": schema.StringAttribute{
							MarkdownDescription: "XML manager",
							Computed:            true,
						},
						"debug_mode": schema.StringAttribute{
							MarkdownDescription: "Specify whether to enable diagnostics. Diagnostics are not intended for use in production environments. Diagnostics consume significant resources that can slow down processing.",
							Computed:            true,
						},
						"debug_history": schema.Int64Attribute{
							MarkdownDescription: "Specify the number of transactions to capture for diagnostics. Enter a value in the range 10 - 250. The default value is 25.",
							Computed:            true,
						},
						"cpa_entries": schema.ListNestedAttribute{
							MarkdownDescription: "Specify CPA entries. Each CPA entry binds an ebXML messaging service (ebMS) to provide partnership interactions between the internal and partner.",
							NestedObject:        models.GetDmB2BCPAEntryDataSourceSchema(),
							Computed:            true,
						},
						"sql_data_source": schema.StringAttribute{
							MarkdownDescription: "SQL data source",
							Computed:            true,
						},
						"front_side_timeout": schema.Int64Attribute{
							MarkdownDescription: "Specify the duration that a client connection can be idle before the connection is closed. For outbound transactions, this connection is between an internal application and the B2B gateway. For inbound transaction, this connection is between an external partner and a B2B gateway.",
							Computed:            true,
						},
						"dependency_actions": actions.ActionsSchema,
					},
				},
			},
		},
	}
}

func (d *B2BGatewayDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *B2BGatewayDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data B2BGatewayList
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.B2BGateway{
		AppDomain: data.AppDomain,
	}

	path := o.GetPath()
	if !data.Id.IsNull() {
		path = path + "/" + data.Id.ValueString()
	}

	res, err := d.pData.Client.Get(path)
	resFound := true
	if err != nil {
		if strings.Contains(err.Error(), "status 401") {
			_ = tfutils.DomainCredentialTest(d.pData.Client, &resp.Diagnostics, data.AppDomain.ValueString())
			if !resp.Diagnostics.HasError() {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Application Domain '%s' does not exist", data.AppDomain.ValueString()))
			}
			return
		} else if !strings.Contains(err.Error(), "status 404") {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
			return
		} else {
			resFound = false
		}
	}
	l := []models.B2BGateway{}
	if resFound {
		if value := res.Get(`B2BGateway`); value.Exists() {
			for _, v := range value.Array() {
				item := models.B2BGateway{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.B2BGatewayObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.B2BGatewayObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

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

type FTPFilePollerSourceProtocolHandlerList struct {
	Id        types.String `tfsdk:"id"`
	AppDomain types.String `tfsdk:"app_domain"`
	Result    types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &FTPFilePollerSourceProtocolHandlerDataSource{}
	_ datasource.DataSourceWithConfigure = &FTPFilePollerSourceProtocolHandlerDataSource{}
)

func NewFTPFilePollerSourceProtocolHandlerDataSource() datasource.DataSource {
	return &FTPFilePollerSourceProtocolHandlerDataSource{}
}

type FTPFilePollerSourceProtocolHandlerDataSource struct {
	pData *tfutils.ProviderData
}

func (d *FTPFilePollerSourceProtocolHandlerDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ftp_file_poller_source_protocol_handler"
}

func (d *FTPFilePollerSourceProtocolHandlerDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "An FTP poller handler has file-based input from a configured FTP directory. The directory must exist and have the appropriate permission to perform the designated operations. For these configured directories, ensure that the read, write, and delete permissions are available, as required.",
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
						"target_directory": schema.StringAttribute{
							MarkdownDescription: "<p>Specify the directory to poll. The path must end in a slash. The path denotes a directory.</p><ul><li>FTP examples: <ul><li>Absolute to the root directory: <tt>\"ftp://user:password@host:port/%2Fpath/\"</tt><p>If the username or password contains the characters colon (:), at symbol (@), or slash (/), use their URL-encoded values in accordance with the specification.</p></li><li>Relative to the home directory of the user: <tt>\"ftp://user:password@host:port/path/\"</tt></li></ul><p>Include a password in the URL with caution. The use of user:/password@host results in a server connection. However, with this configuration, the connection could be unable to send multiple commands to the FTP server. For a stable connection, define a basic authentication policy in the user agent. The user agent is in the XML manager associated with the DataPower service.</p></li><li>NFS example: <ul><li><tt>\"dpnfs://static-mount-name/path/\"</tt></li></ul></li><li>SFTP examples: <ul><li>Absolute to the root directory: <tt>\"sftp://host:port/path/\"</tt></li><li>Relative to the home directory of the user: <tt>\"sftp://host:port/~/path/\"</tt></li></ul></li></ul><p>Do not configure one poller to point at a host name that is a virtual name of a load balancer group. This configuration is not the correct way to poll multiple hosts. To poll multiple hosts, use the same DataPower service and configure one poller object for each real host.</p>",
							Computed:            true,
						},
						"delay_between_polls": schema.Int64Attribute{
							MarkdownDescription: "Specify the interval in milliseconds to wait after the completion of one poll sequence before the next one is started. Enter a value in the range 25 - 100000000. The default value is 60000. <p>A <em>polling sequence</em> is the actual polling action plus the time to complete all transactions that were started by the poll action. The next polling action will start the specified number of milliseconds after the last transaction completes.</p>",
							Computed:            true,
						},
						"input_file_match_pattern": schema.StringAttribute{
							MarkdownDescription: "<p>Enter the PCRE to use to match the contents of the directory being polled. If there is file-renaming or there is a response, this PCRE must create PCRE back references using () pairs.</p><p>For example, if the input files are <tt>NNNNNN.input</tt> , the match pattern would be <tt>\"([0-9]{6})\\.input$\"</tt> .</p>",
							Computed:            true,
						},
						"processing_rename_pattern": schema.StringAttribute{
							MarkdownDescription: "<p>Enter the PCRE to use to rename a file that is being processed. This functionality allows multiple pollers to poll the same directory with the same match pattern. There is no lack of atomicity if the rename operation on the server is atomic. The poller that succeeds in renaming the input file will proceed to process the file. Any other poller that tries to rename the file at the same time will fail to rename the file and will proceed to try the next file that matches the specified match pattern.</p><p>To ensure uniqueness, the resulting file name will be in the following format:</p><p><em>filename</em> . <em>hostname</em> . <em>serial</em> . <em>domain</em> . <em>poller</em> . <em>timestamp</em></p><p><dl><dt><em>filename</em></dt><dd>The file name for the renamed input file.</dd><dt><em>hostname</em></dt><dd>The hostname (system identifier) of the configured DataPower device.</dd><dt><em>serial</em></dt><dd>The serial number of the configured DataPower device.</dd><dt><em>domain</em></dt><dd>The domain of the polling object.</dd><dt><em>poller</em></dt><dd>The name of the polling object.</dd><dt><em>timestamp</em></dt><dd>The timestamp</dd></dl></p><p><b>Note:</b> File renaming cannot be used with an FTP server that supports only 8.3 file names.</p><p>For example if the input files are <tt>NNNNNN.input</tt> and you want to rename them to <tt>NNNNNN.processing</tt> , then the match pattern would be <tt>\"([0-9]{6})\\.input$\"</tt> and the processing pattern would be <tt>\"$1.processing\"</tt> . The resultant file name on the server would be:</p><p><tt>NNNNNN.processing. <em>hostname</em> . <em>serial</em> . <em>domain</em> . <em>poller</em> . <em>timestamp</em></tt></p><p><b>Note:</b> If no processing rename pattern is configured, the file will still be renamed. The only difference is that the <em>filename</em> portion of the resulting file is the name of the original input file, not the renamed input file.</p>",
							Computed:            true,
						},
						"delete_on_success": schema.BoolAttribute{
							MarkdownDescription: "<p>Select whether to delete the input file after successful processing.</p><ul><li>When enabled, deletes the input file.</li><li>When not enabled, renames the input file using the renaming pattern specified by Success File Renaming Pattern.</li></ul>",
							Computed:            true,
						},
						"success_rename_pattern": schema.StringAttribute{
							MarkdownDescription: "<p>When Delete File on Success is off, enter the PCRE to use to rename the input file on success. This PCRE will normally have a back reference for the base input file name. For instance, if input files are <tt>NNNNNN.input</tt> and you want to rename them to <tt>NNNNNN.processed</tt> , the match pattern would be <tt>\"([0-9]{6})\\.input$\"</tt> and the rename pattern would be <tt>\"$1.processed\"</tt> .</p><p>Some servers might allow this pattern to indicate a path that puts the file in a different directory, if it allows cross-directory renames. For instance, the match pattern would be <tt>\"(.*)\"</tt> and the rename pattern would be <tt>\"../processed/$1\"</tt> .</p>",
							Computed:            true,
						},
						"delete_on_error": schema.BoolAttribute{
							MarkdownDescription: "<p>Select whether to delete the input or processing rename file when it could not be processed.</p><ul><li>When enabled, deletes the file.</li><li>When not enabled, renames the input or processing rename file using the renaming pattern specified by Error File Renaming Pattern.</li></ul>",
							Computed:            true,
						},
						"error_rename_pattern": schema.StringAttribute{
							MarkdownDescription: "When Delete File on Processing Error is not enabled, enter the PCRE to use to rename a file when it could not be processed.",
							Computed:            true,
						},
						"generate_result_file": schema.BoolAttribute{
							MarkdownDescription: "<p>Select whether to create a result file after processing an input file.</p><ul><li>When enabled, creates the result file using the naming pattern specified by Result File Name Pattern.</li><li>When not enabled, does not create the result file.</li></ul>",
							Computed:            true,
						},
						"result_name_pattern": schema.StringAttribute{
							MarkdownDescription: "<p>When Generate Result File is on, enter the PCRE to use as the match pattern to build the name of the result file. This PCRE will normally have a back reference for the base input file name. For instance, if input files are <tt>NNNNNN.input</tt> and the desired result file name is <tt>NNNNNN.result</tt> , then the match pattern would be <tt>\"([0-9]{6})\\.input$\"</tt> and the result pattern would be <tt>\"$1.result\"</tt> .</p><p>Some servers might allow this pattern to indicate a path that puts the file in a different directory, if it allows cross-directory renames. For instance, the match pattern would be <tt>\"(.*)\"</tt> and the result pattern would be <tt>\"../result/$1\"</tt> .</p>",
							Computed:            true,
						},
						"processing_seize_timeout": schema.Int64Attribute{
							MarkdownDescription: "<p>Specify the duration in seconds to wait to process a file that is in the processing state. Enter a value in the range 0 - 1000. The default value is 0, which means disabled.</p><p>Processing seizure allows failure handling of a poller when more than one poller polls the same target. If another poller renames a file and does not process and rename or delete it in the specified time, another poller can take over processing. A poller attempts to take over processing when the following conditions are met when compared to the processing seize pattern.</p><ol><li>The seize pattern includes the portion of the file name with the configured processing suffix to match.</li><li>The time stamp is further in the past than the wait time specified by the timeout.</li></ol><p>When these conditions are met, another poller renames the file with a fresh time stamp and processes the file. The processing assumes that the rename operation succeeded.</p>",
							Computed:            true,
						},
						"processing_seize_pattern": schema.StringAttribute{
							MarkdownDescription: "<p>Enter the PCRE to find files that were renamed to indicate that they are in the \"being processed\" state but the processing was never completed.</p><p>The seize pattern contains two phrases. The first phrase is the portion of the file name with the configured processing suffix. The second phrase is the time stamp.</p><p>For example: <tt>(.*.processing).*[.*]([0-9]*)</tt> . This assumes that <tt>$1.processing</tt> was supplied as the renaming pattern.</p>",
							Computed:            true,
						},
						"xml_manager": schema.StringAttribute{
							MarkdownDescription: "An XML Manager manages the compilation and caching of stylesheets and documents. The XML Manager can also control the size and depth of messages processed by this host. Specify an existing XML Manager. More than one service may use the same XML Manager.",
							Computed:            true,
						},
						"max_transfers_per_poll": schema.Int64Attribute{
							MarkdownDescription: "<p>The number of allowed concurrent client connections in a polling sequence.</p><p>Enter a value in the range 0 - 100. The value must be less than the number of simultaneous connections that the polled server accepts. The default value is 0 which means unlimited number of connections based on available system resources. To avoid the consumption of all the systems resources, enter a value other than 0.</p>",
							Computed:            true,
						},
						"dependency_actions": actions.ActionsSchema,
					},
				},
			},
		},
	}
}

func (d *FTPFilePollerSourceProtocolHandlerDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *FTPFilePollerSourceProtocolHandlerDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data FTPFilePollerSourceProtocolHandlerList
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.FTPFilePollerSourceProtocolHandler{
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
	l := []models.FTPFilePollerSourceProtocolHandler{}
	if resFound {
		if value := res.Get(`FTPFilePollerSourceProtocolHandler`); value.Exists() {
			for _, v := range value.Array() {
				item := models.FTPFilePollerSourceProtocolHandler{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.FTPFilePollerSourceProtocolHandlerObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.FTPFilePollerSourceProtocolHandlerObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

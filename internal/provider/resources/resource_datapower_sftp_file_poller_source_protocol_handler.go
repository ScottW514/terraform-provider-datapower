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
	"github.com/scottw514/terraform-provider-datapower/internal/provider/validators"
)

var _ resource.Resource = &SFTPFilePollerSourceProtocolHandlerResource{}
var _ resource.ResourceWithValidateConfig = &SFTPFilePollerSourceProtocolHandlerResource{}

func NewSFTPFilePollerSourceProtocolHandlerResource() resource.Resource {
	return &SFTPFilePollerSourceProtocolHandlerResource{}
}

type SFTPFilePollerSourceProtocolHandlerResource struct {
	pData *tfutils.ProviderData
}

func (r *SFTPFilePollerSourceProtocolHandlerResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sftp_file_poller_source_protocol_handler"
}

func (r *SFTPFilePollerSourceProtocolHandlerResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("Use an SFTP poller handler to have file based input from a configured SFTP directory. The directory must exist and have the appropriate permission to perform the designated operations. For these configured directories, ensure that the read, write, and delete permissions are available, as required.", "source-sftp-poller", "").AddActions("quiesce").String,
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
			"ssh_client_connection": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specifies the name of the SSH client profile which describes how SSH connections are made from the SSH client to an SSH server.", "ssh-client-connection", "ssh_client_profile").String,
				Required:            true,
			},
			"user_summary": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Comments", "summary", "").String,
				Optional:            true,
			},
			"target_directory": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Specify the directory to poll. The path must end in a slash. The path denotes a directory.</p><ul><li>FTP examples: <ul><li>Absolute to the root directory: <tt>\"ftp://user:password@host:port/%2Fpath/\"</tt><p>If the username or password contains the characters colon (:), at symbol (@), or slash (/), use their URL-encoded values in accordance with the specification.</p></li><li>Relative to the home directory of the user: <tt>\"ftp://user:password@host:port/path/\"</tt></li></ul><p>Include a password in the URL with caution. The use of user:/password@host results in a server connection. However, with this configuration, the connection could be unable to send multiple commands to the FTP server. For a stable connection, define a basic authentication policy in the user agent. The user agent is in the XML manager associated with the DataPower service.</p></li><li>NFS example: <ul><li><tt>\"dpnfs://static-mount-name/path/\"</tt></li></ul></li><li>SFTP examples: <ul><li>Absolute to the root directory: <tt>\"sftp://host:port/path/\"</tt></li><li>Relative to the home directory of the user: <tt>\"sftp://host:port/~/path/\"</tt></li></ul></li></ul><p>Do not configure one poller to point at a host name that is a virtual name of a load balancer group. This configuration is not the correct way to poll multiple hosts. To poll multiple hosts, use the same DataPower service and configure one poller object for each real host.</p>", "target-dir", "").String,
				Required:            true,
			},
			"delay_between_polls": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the interval in milliseconds to wait after the completion of one poll sequence before the next one is started. Enter a value in the range 25 - 100000000. The default value is 60000. <p>A <em>polling sequence</em> is the actual polling action plus the time to complete all transactions that were started by the poll action. The next polling action will start the specified number of milliseconds after the last transaction completes.</p>", "delay-time", "").AddIntegerRange(25, 100000000).AddDefaultValue("60000").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(25, 100000000),
				},
				Default: int64default.StaticInt64(60000),
			},
			"input_file_match_pattern": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Enter the PCRE to use to match the contents of the directory being polled. If there is file-renaming or there is a response, this PCRE must create PCRE back references using () pairs.</p><p>For example, if the input files are <tt>NNNNNN.input</tt> , the match pattern would be <tt>\"([0-9]{6})\\.input$\"</tt> .</p>", "match-pattern", "").String,
				Required:            true,
			},
			"processing_rename_pattern": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Enter the PCRE to use to rename a file that is being processed. This functionality allows multiple pollers to poll the same directory with the same match pattern. There is no lack of atomicity if the rename operation on the server is atomic. The poller that succeeds in renaming the input file will proceed to process the file. Any other poller that tries to rename the file at the same time will fail to rename the file and will proceed to try the next file that matches the specified match pattern.</p><p>To ensure uniqueness, the resulting file name will be in the following format:</p><p><em>filename</em> . <em>hostname</em> . <em>serial</em> . <em>domain</em> . <em>poller</em> . <em>timestamp</em></p><p><dl><dt><em>filename</em></dt><dd>The file name for the renamed input file.</dd><dt><em>hostname</em></dt><dd>The hostname (system identifier) of the configured DataPower device.</dd><dt><em>serial</em></dt><dd>The serial number of the configured DataPower device.</dd><dt><em>domain</em></dt><dd>The domain of the polling object.</dd><dt><em>poller</em></dt><dd>The name of the polling object.</dd><dt><em>timestamp</em></dt><dd>The timestamp</dd></dl></p><p><b>Note:</b> File renaming cannot be used with an FTP server that supports only 8.3 file names.</p><p>For example if the input files are <tt>NNNNNN.input</tt> and you want to rename them to <tt>NNNNNN.processing</tt> , then the match pattern would be <tt>\"([0-9]{6})\\.input$\"</tt> and the processing pattern would be <tt>\"$1.processing\"</tt> . The resultant file name on the server would be:</p><p><tt>NNNNNN.processing. <em>hostname</em> . <em>serial</em> . <em>domain</em> . <em>poller</em> . <em>timestamp</em></tt></p><p><b>Note:</b> If no processing rename pattern is configured, the file will still be renamed. The only difference is that the <em>filename</em> portion of the resulting file is the name of the original input file, not the renamed input file.</p>", "processing-rename-pattern", "").String,
				Optional:            true,
			},
			"delete_on_success": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Select whether to delete the input file after successful processing.</p><ul><li>When enabled, deletes the input file.</li><li>When not enabled, renames the input file using the renaming pattern specified by Success File Renaming Pattern.</li></ul>", "success-delete", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"success_rename_pattern": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>When Delete File on Success is off, enter the PCRE to use to rename the input file on success. This PCRE will normally have a back reference for the base input file name. For instance, if input files are <tt>NNNNNN.input</tt> and you want to rename them to <tt>NNNNNN.processed</tt> , the match pattern would be <tt>\"([0-9]{6})\\.input$\"</tt> and the rename pattern would be <tt>\"$1.processed\"</tt> .</p><p>Some servers might allow this pattern to indicate a path that puts the file in a different directory, if it allows cross-directory renames. For instance, the match pattern would be <tt>\"(.*)\"</tt> and the rename pattern would be <tt>\"../processed/$1\"</tt> .</p>", "success-rename-pattern", "").AddDefaultValue("$1.processed.ok").AddNotValidWhen(models.SFTPFilePollerSourceProtocolHandlerSuccessRenamePatternIgnoreVal.String()).String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("$1.processed.ok"),
			},
			"delete_on_error": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Select whether to delete the input or processing rename file when it could not be processed.</p><ul><li>When enabled, deletes the file.</li><li>When not enabled, renames the input or processing rename file using the renaming pattern specified by Error File Renaming Pattern.</li></ul>", "error-delete", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"error_rename_pattern": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("When Delete File on Processing Error is not enabled, enter the PCRE to use to rename a file when it could not be processed.", "error-rename-pattern", "").AddDefaultValue("$0.processed.error").AddNotValidWhen(models.SFTPFilePollerSourceProtocolHandlerErrorRenamePatternIgnoreVal.String()).String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("$0.processed.error"),
			},
			"generate_result_file": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Select whether to create a result file after processing an input file.</p><ul><li>When enabled, creates the result file using the naming pattern specified by Result File Name Pattern.</li><li>When not enabled, does not create the result file.</li></ul>", "result", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"result_name_pattern": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>When Generate Result File is on, enter the PCRE to use as the match pattern to build the name of the result file. This PCRE will normally have a back reference for the base input file name. For instance, if input files are <tt>NNNNNN.input</tt> and the desired result file name is <tt>NNNNNN.result</tt> , then the match pattern would be <tt>\"([0-9]{6})\\.input$\"</tt> and the result pattern would be <tt>\"$1.result\"</tt> .</p><p>Some servers might allow this pattern to indicate a path that puts the file in a different directory, if it allows cross-directory renames. For instance, the match pattern would be <tt>\"(.*)\"</tt> and the result pattern would be <tt>\"../result/$1\"</tt> .</p>", "result-name-pattern", "").AddRequiredWhen(models.SFTPFilePollerSourceProtocolHandlerResultNamePatternCondVal.String()).AddNotValidWhen(models.SFTPFilePollerSourceProtocolHandlerResultNamePatternIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.SFTPFilePollerSourceProtocolHandlerResultNamePatternCondVal, models.SFTPFilePollerSourceProtocolHandlerResultNamePatternIgnoreVal, false),
				},
			},
			"processing_seize_timeout": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Specify the duration in seconds to wait to process a file that is in the processing state. Enter a value in the range 0 - 1000. The default value is 0, which means disabled.</p><p>Processing seizure allows failure handling of a poller when more than one poller polls the same target. If another poller renames a file and does not process and rename or delete it in the specified time, another poller can take over processing. A poller attempts to take over processing when the following conditions are met when compared to the processing seize pattern.</p><ol><li>The seize pattern includes the portion of the file name with the configured processing suffix to match.</li><li>The time stamp is further in the past than the wait time specified by the timeout.</li></ol><p>When these conditions are met, another poller renames the file with a fresh time stamp and processes the file. The processing assumes that the rename operation succeeded.</p>", "processing-seize-timeout", "").AddIntegerRange(0, 1000).AddDefaultValue("0").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 1000),
				},
				Default: int64default.StaticInt64(0),
			},
			"processing_seize_pattern": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Enter the PCRE to find files that were renamed to indicate that they are in the \"being processed\" state but the processing was never completed.</p><p>The seize pattern contains two phrases. The first phrase is the portion of the file name with the configured processing suffix. The second phrase is the time stamp.</p><p>For example: <tt>(.*.processing).*[.*]([0-9]*)</tt> . This assumes that <tt>$1.processing</tt> was supplied as the renaming pattern.</p>", "processing-seize-pattern", "").AddRequiredWhen(models.SFTPFilePollerSourceProtocolHandlerProcessingSeizePatternCondVal.String()).AddNotValidWhen(models.SFTPFilePollerSourceProtocolHandlerProcessingSeizePatternIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.SFTPFilePollerSourceProtocolHandlerProcessingSeizePatternCondVal, models.SFTPFilePollerSourceProtocolHandlerProcessingSeizePatternIgnoreVal, false),
				},
			},
			"xml_manager": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("An XML Manager manages the compilation and caching of stylesheets and documents. The XML Manager can also control the size and depth of messages processed by this host. Specify an existing XML Manager. More than one service may use the same XML Manager.", "xml-manager", "xml_manager").AddDefaultValue("default").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("default"),
			},
			"max_transfers_per_poll": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>The number of allowed concurrent client connections in a polling sequence.</p><p>Enter a value in the range 0 - 100. The value must be less than the number of simultaneous connections that the polled server accepts. The default value is 0 which means unlimited number of connections based on available system resources. To avoid the consumption of all the systems resources, enter a value other than 0.</p>", "max-transfers-per-poll", "").AddIntegerRange(0, 100).AddDefaultValue("0").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 100),
				},
				Default: int64default.StaticInt64(0),
			},
			"dependency_actions": actions.ActionsSchema,
		},
	}
}

func (r *SFTPFilePollerSourceProtocolHandlerResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (r *SFTPFilePollerSourceProtocolHandlerResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.SFTPFilePollerSourceProtocolHandler
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Create, false)
	if resp.Diagnostics.HasError() {
		return
	}

	body := data.ToBody(ctx, `SFTPFilePollerSourceProtocolHandler`)
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

func (r *SFTPFilePollerSourceProtocolHandlerResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.SFTPFilePollerSourceProtocolHandler
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
		data.FromBody(ctx, `SFTPFilePollerSourceProtocolHandler`, res)
	} else {
		// Update
		data.UpdateFromBody(ctx, `SFTPFilePollerSourceProtocolHandler`, res)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SFTPFilePollerSourceProtocolHandlerResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.SFTPFilePollerSourceProtocolHandler
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Update, false)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.pData.Client.Put(data.GetPath()+"/"+data.Id.ValueString(), data.ToBody(ctx, `SFTPFilePollerSourceProtocolHandler`))
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

func (r *SFTPFilePollerSourceProtocolHandlerResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.SFTPFilePollerSourceProtocolHandler
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

func (r *SFTPFilePollerSourceProtocolHandlerResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data models.SFTPFilePollerSourceProtocolHandler

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.ValidateConfig(ctx, &resp.Diagnostics, data.DependencyActions)
}

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

var _ resource.Resource = &B2BGatewayResource{}
var _ resource.ResourceWithValidateConfig = &B2BGatewayResource{}

func NewB2BGatewayResource() resource.Resource {
	return &B2BGatewayResource{}
}

type B2BGatewayResource struct {
	pData *tfutils.ProviderData
}

func (r *B2BGatewayResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_b2b_gateway"
}

func (r *B2BGatewayResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("A B2B gateway can handle B2B messages transmitted over a variety of AS and non-AS protocols for inbound and outbound flows, as well as MDN messages.", "b2bgw", "").AddActions("quiesce").String,
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
			"priority": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the priority of service-scheduling. When system resources are in high demand, high priority services are favored over lower priority services.", "priority", "").AddStringEnum("unknown", "high-min", "high", "high-max", "normal-min", "normal", "normal-max", "low-min", "low", "low-max").AddDefaultValue("normal").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("unknown", "high-min", "high", "high-max", "normal-min", "normal", "normal-max", "low-min", "low", "low-max"),
				},
				Default: stringdefault.StaticString("normal"),
			},
			"doc_store_location": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the location for document storage, which saves copies of inbound, outbound, and intermediate documents that might be needed for a retransmit operation. If unspecified, documents are stored in the encrypted area on the RAID volume. For storage on the RAID volume, set the maximum disk usage to 30 GB.", "doc-location", "").AddDefaultValue("(default)").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("(default)"),
			},
			"as_front_protocol": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Protocol handlers", "as-fsph", "").String,
				NestedObject:        models.GetDmASFrontProtocolResourceSchema(),
				Optional:            true,
			},
			"as1mdn_email": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("For AS1 asynchronous MDN scenarios, specify the default email address for the AS1 asynchronous MDN. When sending outbound AS1 email messages that request an MDN, this email address can be the default email address for the response MDN. An email address in the destination overrides this value.", "as1-mdn-email", "").String,
				Optional:            true,
			},
			"as1mdnsmtp_server_connection": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("When an incoming AS1 message requests an MDN as an email request, specify the SMTP server connection for asynchronous MDN responses.", "as1-mdn-smtp-server-connection", "smtp_server_connection").String,
				Optional:            true,
			},
			"as2mdnurl": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("For AS2 asynchronous MDN scenarios, specify the default URL for the AS2 asynchronous MDN. This URL can point to that gateway itself or a firewall or proxy that routes the message to the gateway. A URL in the destination overrides this value.", "as2-mdn-url", "").String,
				Optional:            true,
			},
			"as3mdnurl": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("For AS3 asynchronous MDN scenarios, specify the default URL for the AS3 asynchronous MDN. This URL can point to that gateway itself or a firewall or proxy that routes the message to the gateway. A URL in the destination overrides this value.", "as3-mdn-url", "").String,
				Optional:            true,
			},
			"b2b_profiles": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Active partner profiles", "b2b-profile", "").AddRequiredWhen(models.B2BGatewayB2BProfilesCondVal.String()).String,
				NestedObject:        models.GetDmB2BActiveProfileResourceSchema(),
				Optional:            true,
				Validators: []validator.List{
					validators.ConditionalRequiredList(models.B2BGatewayB2BProfilesCondVal, validators.Evaluation{}, false),
				},
			},
			"b2b_groups": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Active profile groups", "b2b-group", "").String,
				NestedObject:        models.GetDmB2BActiveGroupResourceSchema(),
				Optional:            true,
			},
			"document_routing_preprocessor_type": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the file type of the document-routing preprocessor file to run against messages that are not received through AS or ebMS protocols. The default value is stylesheet.", "document-routing-preprocessor-type", "").AddStringEnum("stylesheet", "gatewayscript").AddDefaultValue("stylesheet").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("stylesheet", "gatewayscript"),
				},
				Default: stringdefault.StaticString("stylesheet"),
			},
			"document_routing_preprocessor": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the location of the document-routing preprocessor file, which is a stylesheet or a GatewayScript file to run against transactions that cannot be handled by B2B-specific handlers. <ul><li>A stylesheet examines information from transport headers and other non-content sources to select relevant trading partners. The default stylesheet is <tt>store:///b2b-routing.xsl</tt> .</li><li>A GatewayScript examines information from transport headers and payloads to select relevant trading partners. A GatewayScript can parse messages of different data types, such as JSON, XML, and non-XML. A sample GatewayScript file is <tt>store:///gateayscript/example-b2b-routing.js</tt> .</li></ul>", "document-routing-preprocessor", "").AddDefaultValue("store:///b2b-routing.xsl").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("store:///b2b-routing.xsl"),
			},
			"document_routing_preprocessor_debug": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to enable the GatewayScript debugger to detect and diagnose errors in the document routing preprocessor file. To debug a file, ensure that the file contains one or more <tt>debugger;</tt> statements where you want to start debugging.", "document-routing-preprocessor-debug", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"archive_mode": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Purge mode", "arch-mode", "").AddStringEnum("archpurge", "purgeonly").AddDefaultValue("archpurge").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("archpurge", "purgeonly"),
				},
				Default: stringdefault.StaticString("archpurge"),
			},
			"archive_location": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the location for archive file. Enter the fully qualified name of the directory. To copy the archive file to an FTP server, ensure that the FTP policies in the XML manager enable image (binary) data transfer.", "arch-dir", "").AddRequiredWhen(models.B2BGatewayArchiveLocationCondVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.B2BGatewayArchiveLocationCondVal, validators.Evaluation{}, false),
				},
			},
			"archive_file_name": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the base file name for archive file. When archiving, the operation appends the current timestamp.", "arch-file", "").AddRequiredWhen(models.B2BGatewayArchiveFileNameCondVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(0, 128),
					stringvalidator.RegexMatches(regexp.MustCompile("^[_a-zA-Z0-9.-]+$"), "Must match :"+"^[_a-zA-Z0-9.-]+$"),
					validators.ConditionalRequiredString(models.B2BGatewayArchiveFileNameCondVal, validators.Evaluation{}, false),
				},
			},
			"archive_minimum_size": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the minimum remaining size in KB of document storage that triggers archival. The default value is 1024.", "arch-minimum-size", "").AddDefaultValue("1024").String,
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(1024),
			},
			"archive_document_age": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the maximum duration in days to retain processed documents. Enter a value in the range 1 - 3650. The default value is 90.", "arch-document-age", "").AddIntegerRange(1, 3650).AddDefaultValue("90").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 3650),
				},
				Default: int64default.StaticInt64(90),
			},
			"archive_minimum_documents": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the minimum number of documents to retain in document storage after archival. The minimum value is 1. The default value is 100.", "arch-minimum-documents", "").AddIntegerRange(1, 65535).AddDefaultValue("100").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 65535),
				},
				Default: int64default.StaticInt64(100),
			},
			"disk_use_check_interval": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the interval in minutes between checks for documents that exceed the maximum age. During the check, documents that exceed the maximum age are purged. Enter a value in the range 1 - 1440. The default value is 60.", "diskuse-check-interval", "").AddIntegerRange(1, 1440).AddDefaultValue("60").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 1440),
				},
				Default: int64default.StaticInt64(60),
			},
			"max_document_disk_use": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the maximum size in KB for document storage. When storage exceeds this value, documents are purged. The default value is 25165824.", "max-diskuse", "").AddIntegerRange(1, 4294967295).AddDefaultValue("25165824").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 4294967295),
				},
				Default: int64default.StaticInt64(25165824),
			},
			"archive_monitor": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to use a monitor during archival. The monitor limits the message injection rate to prevent problems in a critical situation like performance testing or on a heavily loaded system.", "arch-monitor", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"shaping_threshold": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the maximum TPS to allow during archival. When the threshold is reached, the service queues transactions. When the queue is full, the service rejects transactions and generates a log message. Enter a value in the range 10 - 10000. The default value is 200.", "arch-shaping-threshold", "").AddIntegerRange(10, 10000).AddDefaultValue("200").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(10, 10000),
				},
				Default: int64default.StaticInt64(200),
			},
			"archive_backup_documents": models.GetDmB2BBackupMsgTypeResourceSchema("Specify the types of documents to archive. This property does not indicate the inbound or outbound transaction to archive.", "arch-backup-documents", "", false),
			"x_path_routing_policies": schema.ListAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("XPath routing policies", "xpath-routing", "b2b_xpath_routing_policy").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"xml_manager": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("XML manager", "xml-manager", "xml_manager").AddDefaultValue("default").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("default"),
			},
			"debug_mode": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to enable diagnostics. Diagnostics are not intended for use in production environments. Diagnostics consume significant resources that can slow down processing.", "debug-mode", "").AddStringEnum("on", "off", "unbounded").AddDefaultValue("off").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("on", "off", "unbounded"),
				},
				Default: stringdefault.StaticString("off"),
			},
			"debug_history": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the number of transactions to capture for diagnostics. Enter a value in the range 10 - 250. The default value is 25.", "debug-history", "").AddIntegerRange(10, 250).AddDefaultValue("25").AddRequiredWhen(models.B2BGatewayDebugHistoryCondVal.String()).String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(10, 250),
					validators.ConditionalRequiredInt64(models.B2BGatewayDebugHistoryCondVal, validators.Evaluation{}, true),
				},
				Default: int64default.StaticInt64(25),
			},
			"cpa_entries": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify CPA entries. Each CPA entry binds an ebXML messaging service (ebMS) to provide partnership interactions between the internal and partner.", "cpa-entry", "").String,
				NestedObject:        models.GetDmB2BCPAEntryResourceSchema(),
				Optional:            true,
			},
			"sql_data_source": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("SQL data source", "sql-source", "sql_data_source").String,
				Optional:            true,
			},
			"front_side_timeout": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the duration that a client connection can be idle before the connection is closed. For outbound transactions, this connection is between an internal application and the B2B gateway. For inbound transaction, this connection is between an external partner and a B2B gateway.", "front-side-timeout", "").AddIntegerRange(1, 86400).AddDefaultValue("120").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 86400),
				},
				Default: int64default.StaticInt64(120),
			},
			"dependency_actions": actions.ActionsSchema,
		},
	}
}

func (r *B2BGatewayResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (r *B2BGatewayResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.B2BGateway
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Create, false)
	if resp.Diagnostics.HasError() {
		return
	}

	body := data.ToBody(ctx, `B2BGateway`)
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

func (r *B2BGatewayResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.B2BGateway
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
		data.FromBody(ctx, `B2BGateway`, res)
	} else {
		// Update
		data.UpdateFromBody(ctx, `B2BGateway`, res)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *B2BGatewayResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.B2BGateway
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Update, false)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.pData.Client.Put(data.GetPath()+"/"+data.Id.ValueString(), data.ToBody(ctx, `B2BGateway`))
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

func (r *B2BGatewayResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.B2BGateway
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

func (r *B2BGatewayResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data models.B2BGateway

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.ValidateConfig(ctx, &resp.Diagnostics, data.DependencyActions)
}

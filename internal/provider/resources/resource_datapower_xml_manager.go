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
)

var _ resource.Resource = &XMLManagerResource{}
var _ resource.ResourceWithValidateConfig = &XMLManagerResource{}

func NewXMLManagerResource() resource.Resource {
	return &XMLManagerResource{}
}

type XMLManagerResource struct {
	pData *tfutils.ProviderData
}

func (r *XMLManagerResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_xml_manager"
}

func (r *XMLManagerResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("An XML Manager manages the compilation and caching of stylesheets, the caching of documents, and provides configuration constraints on the size and parsing depth of documents. You can enable streaming operation by configuring an XML Manager to use a Streaming Compile Option Policy. More than one firewall can use the same XML Manager.", "xmlmgr", "").AddActions("flush_stylesheet_cache", "flush_document_cache", "flush_ldap_pool_cache").String,
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
			"url_refresh_policy": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Optionally, assigns a Stylesheet Refresh Policy to the XML manager. Stylesheets cached by this manager are refreshed in accordance with policy rules.", "xslrefresh", "url_refresh_policy").String,
				Optional:            true,
			},
			"compile_options_policy": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Optionally, assigns a Compilation Options Policy to the XML Manager. Stylesheets cached by this manager are compiled in accordance with policy rules. Compile Options allow the administrator to manage profiling and debug modes, which are helpful during development and troubleshooting sessions. Profiling and debug modes should not be left on for production purposes.</p><p>Compile Options Policy also enables streaming mode, which might be a valid choice for production environments that regularly encounter very large documents.</p>", "xslconfig", "compile_options_policy").String,
				Optional:            true,
			},
			"cache_memory_size": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enter the maximum size of the stylesheet cache. The default value is 2147483647. A value of 0 disables caching. Stylesheets are purged when either the cache size or the cache count is reached.", "xsl cache memorysize", "").AddDefaultValue("2147483647").String,
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(2147483647),
			},
			"cache_size": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enter the maximum number of stylesheets to cache. Enter a value in the range 5 - 250000. The default value is 256. Stylesheets are purged when either the cache size or the cache count is reached.", "xsl cache size", "").AddIntegerRange(5, 250000).AddDefaultValue("256").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(5, 250000),
				},
				Default: int64default.StaticInt64(256),
			},
			"sha1caching": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enables/disables SHA1-assisted stylesheet caching. With SHA1 caching enabled, stylesheets are cached by both URL and SHA1 message digest value. With SHA1 caching disabled, stylesheets are cached only by URL.", "xsl checksummed cache", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"static_document_calls": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("The latest XSLT specifications require that multiple document calls in the same transformation return the same result. Disable to allow all document calls to operate independently.", "static-document-calls", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"search_results": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Configures optimization of '//' in XPath expressions", "search results", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"virtual_servers": schema.ListAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Select a Load Balancer Group and assign that group to this XML manager. A Load Balancer Group, or server pool, provides redundancy among backside resources.", "loadbalancer-group", "load_balancer_group").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"parser_limits_bytes_scanned": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enter the maximum number of bytes scanned in one message. This property applies to any XML document that is parsed. If the value is 0, no limit is enforced. The default value is 4194304.", "bytes-scanned", "").AddDefaultValue("4194304").String,
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(4194304),
			},
			"parser_limits_element_depth": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enter the maximum depth of element nesting in XML parser. The default value is 512.", "element-depth", "").AddDefaultValue("512").String,
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(512),
			},
			"parser_limits_attribute_count": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enter the maximum number of attributes of a given element. The default value is 128.", "attribute-count", "").AddDefaultValue("128").String,
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(128),
			},
			"parser_limits_max_node_size": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enter the maximum size that any one node can consume. Enter a value in the range 1024 - 4294967295. The default value is 33554432. Sizes which are powers of two result in the best performance. Although you define a value, the DataPower Gateway uses a value that is the rounded-down to the largest power of two that is smaller than the defined value.", "max-node-size", "").AddIntegerRange(1024, 4294967295).AddDefaultValue("33554432").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(1024, 4294967295),
				},
				Default: int64default.StaticInt64(33554432),
			},
			"parser_limits_external_references": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Select the special handling for input documents that contain external references, such as an external entity or external DTD definition.", "external-references", "").AddStringEnum("forbid", "ignore", "allow").AddDefaultValue("forbid").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("forbid", "ignore", "allow"),
				},
				Default: stringdefault.StaticString("forbid"),
			},
			"parser_limits_max_prefixes": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enter the maximum number of distinct XML namespace prefixes in a document. This limit counts multiple prefixes defined for the same namespace, but does not count multiple namespaces defined in different parts of the input document under a single prefix. Enter a value in the range 0 - 262143. The default value is 1024. A value of 0 indicates that the limit is 1024.", "max-prefixes", "").AddIntegerRange(0, 262143).AddDefaultValue("1024").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 262143),
				},
				Default: int64default.StaticInt64(1024),
			},
			"parser_limits_max_namespaces": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enter the maximum number of distinct XML namespace URIs in a document. This limit counts all XML namespaces, regardless of how many prefixes are used to declare them. Enter a value in the range 0 - 65535. The default value is 1024. A value of 0 indicates that the limit is 1024.", "max-namespaces", "").AddIntegerRange(0, 65535).AddDefaultValue("1024").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 65535),
				},
				Default: int64default.StaticInt64(1024),
			},
			"parser_limits_max_local_names": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enter the maximum number of distinct XML local names in a document. This limit counts all local names, independent of the namespaces they are associated with. Enter a value in the range 0 - 1048575. The default value is 60000. A value of 0 indicates that the limit is 60000.", "max-local-names", "").AddIntegerRange(0, 1048575).AddDefaultValue("60000").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 1048575),
				},
				Default: int64default.StaticInt64(60000),
			},
			"doc_cache_max_docs": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enter the maximum number of documents to cache. Enter a value in the range 1 - 250000. The default value is 5000.", "maxdocs", "").AddIntegerRange(1, 250000).AddDefaultValue("5000").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 250000),
				},
				Default: int64default.StaticInt64(5000),
			},
			"doc_cache_size": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enter the maximum size of the document cache. Regardless of the specified size, no document that is greater than 1073741824 bytes is cached. This restriction applies even if the cache has available space.", "size", "").String,
				Optional:            true,
			},
			"doc_max_writes": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enter the maximum number of concurrent write requests to create documents or refresh expired documents in the document cache. Enter a value in the range 1 - 32768. The default value is 32768. After the maximum number is reached, requests are forwarded to the target server and the response is not written to the cache.", "max-writes", "").AddIntegerRange(1, 32768).AddDefaultValue("32768").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 32768),
				},
				Default: int64default.StaticInt64(32768),
			},
			"doc_cache_policy": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("A document cache policy allows the administrator to determine how documents are cached by the XML manager. The policy offers time-to-live, priority, and type configuration values. This document cache is distinct from the stylesheet cache for the XML manager. Documents and stylesheets that the XML manager retrieves can be cached.", "policy", "").String,
				NestedObject:        models.GetDmDocCachePolicyResourceSchema(),
				Optional:            true,
			},
			"schema_validation": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("", "validate", "").String,
				NestedObject:        models.GetDmSchemaValidationResourceSchema(),
				Optional:            true,
			},
			"scheduled_rule": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Certain applications might require the running of a scheduled processing rule. Integration with a CA Unicenter Manager is facilitated by a regularly scheduled processing rule that obtains relationship data from the Unicenter Manager.", "schedule-rule", "").String,
				NestedObject:        models.GetDmScheduledRuleResourceSchema(),
				Optional:            true,
			},
			"user_agent": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Identifies the User Agent configuration the XML manager uses to connect to external servers to retrieve resources.", "user-agent", "http_user_agent").AddDefaultValue("default").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("default"),
			},
			"json_parser_settings": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Identifies the JSON Settings used by the XML manager when processing a JSON body. If no JSON Settings are specified, the default limits are enforced.", "json-parser-settings", "json_settings").String,
				Optional:            true,
			},
			"ldap_conn_pool": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Identifies the LDAP connection pool to be used by the XML Manager when contacting LDAP", "ldap-pool", "ldap_connection_pool").String,
				Optional:            true,
			},
			"dependency_actions": actions.ActionsSchema,
		},
	}
}

func (r *XMLManagerResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (r *XMLManagerResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.XMLManager
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Create, false)
	if resp.Diagnostics.HasError() {
		return
	}

	body := data.ToBody(ctx, `XMLManager`)
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

func (r *XMLManagerResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.XMLManager
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
		data.FromBody(ctx, `XMLManager`, res)
	} else {
		// Update
		data.UpdateFromBody(ctx, `XMLManager`, res)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *XMLManagerResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.XMLManager
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Update, false)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.pData.Client.Put(data.GetPath()+"/"+data.Id.ValueString(), data.ToBody(ctx, `XMLManager`))
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

func (r *XMLManagerResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.XMLManager
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

func (r *XMLManagerResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data models.XMLManager

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.ValidateConfig(ctx, &resp.Diagnostics, data.DependencyActions)
}

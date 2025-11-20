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

type APICollectionList struct {
	Id        types.String `tfsdk:"id"`
	AppDomain types.String `tfsdk:"app_domain"`
	Result    types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &APICollectionDataSource{}
	_ datasource.DataSourceWithConfigure = &APICollectionDataSource{}
)

func NewAPICollectionDataSource() datasource.DataSource {
	return &APICollectionDataSource{}
}

type APICollectionDataSource struct {
	pData *tfutils.ProviderData
}

func (d *APICollectionDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_api_collection"
}

func (d *APICollectionDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "An API collection is a logical partition of an API gateway that packages the plans and subscribers to make APIs available to a specific group of clients. An API collection corresponds to a catalog in the API manager.",
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
						"sandbox": schema.BoolAttribute{
							MarkdownDescription: "Specify whether the APIs in this catalog are for test purposes. By default, a catalog is not for test purposes.",
							Computed:            true,
						},
						"org_id": schema.StringAttribute{
							MarkdownDescription: "Organization ID",
							Computed:            true,
						},
						"org_name": schema.StringAttribute{
							MarkdownDescription: "Organization name",
							Computed:            true,
						},
						"catalog_id": schema.StringAttribute{
							MarkdownDescription: "Catalog ID",
							Computed:            true,
						},
						"catalog_name": schema.StringAttribute{
							MarkdownDescription: "Catalog name",
							Computed:            true,
						},
						"dev_portal_endpoint": schema.StringAttribute{
							MarkdownDescription: "Specify the URL of the Developer Portal endpoint. This endpoint can be used to provide security credentials for access to an API.",
							Computed:            true,
						},
						"cache_capacity": schema.Int64Attribute{
							MarkdownDescription: "Specify the maximum number of subscriber entries to cache. Enter a value in the range 8 - 51200. The default value is 128. When the limit is exceeded, the least recently used (LRU) entry is removed.",
							Computed:            true,
						},
						"routing_prefix": schema.ListNestedAttribute{
							MarkdownDescription: "Specify the routing prefix to determine which API collection to route the request. You can use routing prefixes to organize your APIs and plans into collections and subcollections. For example, if you have a collection of APIs serving for a certain purpose, and the APIs are to be used by two segments of your organization, you might create two API collections with the organization name, purpose name, and segment name in the routing prefix. If the organization name is <tt>myorg</tt> , the APIs serve for purpose <tt>purpose1</tt> , and the two segments under the organization is <tt>section1</tt> and <tt>section2</tt> , the resulting URL routing prefixes are <tt>/myorg/purpose1/section1</tt> and <tt>/myorg/purpose1/section2</tt> . The resulting hostname routing prefixes are <tt>section1.purpose1.myorg</tt> and <tt>section2.purpose1.myorg</tt> . <p>The API gateway uses the routing prefix to form the complete URI <tt>routing_prefix/base_path/operation_path</tt> and accepts only the incoming requests with this URI. In the complete URI, <tt>base_path</tt> is the base path on which the API is served, and <tt>operation_path</tt> is the relative path to the base path where the operations are available.</p><p>The default routing prefix is slash (/) when the type is URI and blank when the type is hostname. An API collection becomes the default API collection in the API Gateway when the API collection has a default routing prefix. The API gateway routes a request to the default API collection when other API collections do not match. An API gateway can have only one default API collection. Therefore, regardless of the prefix type, only one API collection can be configured with the default routing prefix.</p>",
							NestedObject:        models.GetDmRoutingPrefixDataSourceSchema(),
							Computed:            true,
						},
						"use_rate_limit_group": schema.BoolAttribute{
							MarkdownDescription: "Use rate limit group",
							Computed:            true,
						},
						"default_rate_limit": schema.ListNestedAttribute{
							MarkdownDescription: "Specify the default rate limit scheme for API requests without API keys for client identification. When not defined, requests without API keys are rejected.",
							NestedObject:        models.GetDmAPIRateLimitDataSourceSchema(),
							Computed:            true,
						},
						"rate_limit_group": schema.StringAttribute{
							MarkdownDescription: "Specify the default rate limit group for API requests without API keys for client identification. When not defined, requests without API keys are rejected.",
							Computed:            true,
						},
						"assembly_burst_limit": schema.ListNestedAttribute{
							MarkdownDescription: "Assembly burst limit",
							NestedObject:        models.GetDmAPIBurstLimitDataSourceSchema(),
							Computed:            true,
						},
						"assembly_rate_limit": schema.ListNestedAttribute{
							MarkdownDescription: "Assembly rate limit",
							NestedObject:        models.GetDmAPIRateLimitDataSourceSchema(),
							Computed:            true,
						},
						"assembly_count_limit": schema.ListNestedAttribute{
							MarkdownDescription: "Assembly count limit",
							NestedObject:        models.GetDmAPICountLimitDataSourceSchema(),
							Computed:            true,
						},
						"enforce_pre_assembly_rate_limits": schema.BoolAttribute{
							MarkdownDescription: "Specify whether to enforce the API rate and burst limits from the plan. When disabled, only the limits specified in a rate limit assembly action are applied to this API.",
							Computed:            true,
						},
						"api_processing_rule": schema.StringAttribute{
							MarkdownDescription: "Specify the processing rule to process API requests. When your collection requires custom processing, use API Connect global policies to define the custom rules.",
							Computed:            true,
						},
						"api_error_rule": schema.StringAttribute{
							MarkdownDescription: "Specify the processing rule to handle errors during API processing. When your collection requires custom processing, use API Connect global policies to define the custom rules.",
							Computed:            true,
						},
						"assembly_preflow": schema.StringAttribute{
							MarkdownDescription: "Specify the processing rule to run before the assembly rule. When your collection requires custom processing, use API Connect global policies to configure the assembly.",
							Computed:            true,
						},
						"assembly_postflow": schema.StringAttribute{
							MarkdownDescription: "Specify the processing rule to run after the assembly rule. When your collection requires custom processing, use API Connect global policies to configure the assembly.",
							Computed:            true,
						},
						"plan": schema.ListAttribute{
							MarkdownDescription: "Specify the API plans for the collection. Each plan contains a list of APIs and defines the rate limit for the API operations.",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"analytics_endpoint": schema.StringAttribute{
							MarkdownDescription: "Analytic endpoint",
							Computed:            true,
						},
						"application_type": schema.ListAttribute{
							MarkdownDescription: "Application types",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"parse_settings_reference": models.GetDmDynamicParseSettingsReferenceDataSourceSchema("Specify the constraints to parse documents. Precedence rules apply when the constraint for the same aspect of an input document is configured with more than one method. <ul><li>You can specify a URL reference from which to retrieve the constraints definition.</li><li>You can specify a literal configuration string in XML management interface or REST management interface format that contains the constraints definition.</li><li>You can specify a parse settings configuration object to retrieve the constraints definition.</li></ul>", "parse-settings-reference", ""),
						"dependency_actions":       actions.ActionsSchema,
					},
				},
			},
		},
	}
}

func (d *APICollectionDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *APICollectionDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data APICollectionList
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.APICollection{
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
	l := []models.APICollection{}
	if resFound {
		if value := res.Get(`APICollection`); value.Exists() {
			for _, v := range value.Array() {
				item := models.APICollection{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.APICollectionObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.APICollectionObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

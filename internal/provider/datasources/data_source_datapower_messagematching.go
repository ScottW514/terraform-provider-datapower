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

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
)

type MessageMatchingList struct {
	AppDomain types.String `tfsdk:"app_domain"`
	Result    types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &MessageMatchingDataSource{}
	_ datasource.DataSourceWithConfigure = &MessageMatchingDataSource{}
)

func NewMessageMatchingDataSource() datasource.DataSource {
	return &MessageMatchingDataSource{}
}

type MessageMatchingDataSource struct {
	pData *tfutils.ProviderData
}

func (d *MessageMatchingDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_messagematching"
}

func (d *MessageMatchingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "<p>Message Matching determines what messages will be monitored by any monitor that uses the Message Matching object. It is a definition of the traffic flow that is monitored.</p><p>To capture all messages, set the HTTP Method field to 'any' and leave all other fields blank.</p>",
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
						"ip_address": schema.StringAttribute{
							MarkdownDescription: "Use an IP network address and prefix length to define a contiguous range of IP addresses included in this traffic definition (for example, 10.10.100.0/28 specifies IP addresses 10.10.100.0 through 10.10.100.15, and 10.10.100.19/32 specifies a single host address).",
							Computed:            true,
						},
						"ip_exclude": schema.StringAttribute{
							MarkdownDescription: "Use an IP network address and prefix length to define a contiguous range of IP addresses excluded from this traffic definition (for example, 10.10.100.0/28 specifies IP addresses 10.10.100.0 through 10.10.100.15, and 10.10.100.19/32 specifies a single host address).",
							Computed:            true,
						},
						"http_method": schema.StringAttribute{
							MarkdownDescription: "Select the HTTP method included in this traffic definition. The default value is any, which indicates that HTTP method is not a match criteria.",
							Computed:            true,
						},
						"http_header": schema.ListNestedAttribute{
							MarkdownDescription: "Click this tab to add match criteria for inclusive HTTP header fields to the traffic definition. The field with its corresponding value must appear in the HTTP header of the message to qualify for inclusion.",
							NestedObject:        models.DmHTTPHeaderDataSourceSchema,
							Computed:            true,
						},
						"http_header_exclude": schema.ListNestedAttribute{
							MarkdownDescription: "Click this tab to add match criteria for exclusive HTTP header field to the traffic definition. The field with its corresponding value cannot appear in the HTTP header of the message to qualify for exclusion.",
							NestedObject:        models.DmHTTPHeaderDataSourceSchema,
							Computed:            true,
						},
						"request_url": schema.StringAttribute{
							MarkdownDescription: "Provide a literal or wildcard expression to define a URL set included in this traffic definition. The following wildcard characters are available when defining the URL set. <table><tr><td valign=\"top\">asterisk (*)</td><td valign=\"top\">Matches 0 or more occurrences of any character</td></tr><tr><td valign=\"top\">question mark (?)</td><td valign=\"top\">Matches one occurrence of any single character</td></tr><tr><td valign=\"top\">brackets ( [ ] )</td><td valign=\"top\">Defines a character or numeric range. For example, [1-5] matches 1, 2, 3, 4, or 5, while xs[dl] matches xsd or xsl.</td></tr></table>",
							Computed:            true,
						},
						"dependency_actions": actions.ActionsSchema,
					},
				},
			},
		},
	}
}

func (d *MessageMatchingDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *MessageMatchingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data MessageMatchingList
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.MessageMatching{
		AppDomain: data.AppDomain,
	}

	res, err := d.pData.Client.Get(o.GetPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}
	l := []models.MessageMatching{}
	if value := res.Get(`MessageMatching`); value.Exists() {
		for _, v := range value.Array() {
			item := models.MessageMatching{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.MessageMatchingObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.MessageMatchingObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

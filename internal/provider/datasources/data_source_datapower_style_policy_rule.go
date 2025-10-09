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

type StylePolicyRuleList struct {
	Id        types.String `tfsdk:"id"`
	AppDomain types.String `tfsdk:"app_domain"`
	Result    types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &StylePolicyRuleDataSource{}
	_ datasource.DataSourceWithConfigure = &StylePolicyRuleDataSource{}
)

func NewStylePolicyRuleDataSource() datasource.DataSource {
	return &StylePolicyRuleDataSource{}
}

type StylePolicyRuleDataSource struct {
	pData *tfutils.ProviderData
}

func (d *StylePolicyRuleDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_style_policy_rule"
}

func (d *StylePolicyRuleDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Create, edit, or delete processing policy rules. Rules can also be created and edited using a drag and drop interface by opening the Services menu area on the left-hand navigation bar. Click the appropriate Policy link which will then present an opportunity to create both Rules and Actions.",
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
						"actions": schema.ListAttribute{
							MarkdownDescription: "Define XSL filters and/or transformations. Select the desired existing action from the list and click Add to add the action to the rule. Actions are executed in the order (top to bottom) in which they are listed in the box. Click the + button to create a new Action; select an action and click the ... button to edit an existing action. To delete an action, select it from the list and click Delete.",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"direction": schema.StringAttribute{
							MarkdownDescription: "Select a rule type. The default is Rule.",
							Computed:            true,
						},
						"input_format": schema.StringAttribute{
							MarkdownDescription: "Select a decompression algorithm to apply to all messages before any other processing occurs (a preprocess step). All messages are decompressed using the selected algorithm. If the message is not compressed with this algorithm, an error is raised. This setting is independent of transport-level decompression. The default is None.",
							Computed:            true,
						},
						"output_format": schema.StringAttribute{
							MarkdownDescription: "Select a compression algorithm to apply to all messages after any other processing occurs (a post-process step). All messages are compressed using the selected algorithm. The resulting archive contains only one file. This setting is independent of transport-level compression. The default is None.",
							Computed:            true,
						},
						"non_xml_processing": schema.BoolAttribute{
							MarkdownDescription: "Normally, processing is only performed on XML content. This option allows processing actions to be performed on Non-XML content as well.",
							Computed:            true,
						},
						"unprocessed": schema.BoolAttribute{
							MarkdownDescription: "Permit rule to pass-through data unprocessed",
							Computed:            true,
						},
						"user_summary": schema.StringAttribute{
							MarkdownDescription: "Comments",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *StylePolicyRuleDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *StylePolicyRuleDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data StylePolicyRuleList
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.StylePolicyRule{
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
	l := []models.StylePolicyRule{}
	if resFound {
		if value := res.Get(`StylePolicyRule`); value.Exists() {
			for _, v := range value.Array() {
				item := models.StylePolicyRule{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.StylePolicyRuleObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.StylePolicyRuleObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

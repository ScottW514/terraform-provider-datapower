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

type PolicyAttachmentsList struct {
	AppDomain types.String `tfsdk:"app_domain"`
	Result    types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &PolicyAttachmentsDataSource{}
	_ datasource.DataSourceWithConfigure = &PolicyAttachmentsDataSource{}
)

func NewPolicyAttachmentsDataSource() datasource.DataSource {
	return &PolicyAttachmentsDataSource{}
}

type PolicyAttachmentsDataSource struct {
	pData *tfutils.ProviderData
}

func (d *PolicyAttachmentsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_policy_attachments"
}

func (d *PolicyAttachmentsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Create and configure ws-policy attachments for WSDLs",
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
						"enforcement_mode": schema.StringAttribute{
							MarkdownDescription: "Enforcement Mode defines how the service uses WS-Policy to ensure that messages meet security requirements. The default behavior is enforce.",
							Computed:            true,
						},
						"policy_references": schema.BoolAttribute{
							MarkdownDescription: "Enable policies attached to WSDL using PolicyURI attributes and PolicyReference elements. These attachments are sometimes called XML element attachments. If 'off', all PolicyURI attributes and PolicyReference elements are ignored and only external policies are enforced.",
							Computed:            true,
						},
						"ignored_policy_attachment_points": schema.ListNestedAttribute{
							MarkdownDescription: "Disable all policies attached by policy references at a configured attachment point leaving other policy references intact.",
							NestedObject:        models.GetDmPolicyAttachmentPointDataSourceSchema(),
							Computed:            true,
						},
						"external_policy": schema.ListNestedAttribute{
							MarkdownDescription: "Associate an external policy with a service.",
							NestedObject:        models.GetDmExternalAttachedPolicyDataSourceSchema(),
							Computed:            true,
						},
						"sla_enforcement_mode": schema.StringAttribute{
							MarkdownDescription: "SLA Enforcement Mode controls the application of SLA Policies to transactions. Transactions are either allowed or rejected based on whether an SLA rule is applied to the transaction.",
							Computed:            true,
						},
						"dependency_actions": actions.ActionsSchema,
					},
				},
			},
		},
	}
}

func (d *PolicyAttachmentsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *PolicyAttachmentsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data PolicyAttachmentsList
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.PolicyAttachments{
		AppDomain: data.AppDomain,
	}

	res, err := d.pData.Client.Get(o.GetPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}
	l := []models.PolicyAttachments{}
	if value := res.Get(`PolicyAttachments`); value.Exists() {
		for _, v := range value.Array() {
			item := models.PolicyAttachments{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.PolicyAttachmentsObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.PolicyAttachmentsObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

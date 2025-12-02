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
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
)

type B2BCPACollaborationList struct {
	ProviderTarget types.String `tfsdk:"provider_target"`
	Id             types.String `tfsdk:"id"`
	AppDomain      types.String `tfsdk:"app_domain"`
	Result         types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &B2BCPACollaborationDataSource{}
	_ datasource.DataSourceWithConfigure = &B2BCPACollaborationDataSource{}
)

func NewB2BCPACollaborationDataSource() datasource.DataSource {
	return &B2BCPACollaborationDataSource{}
}

type B2BCPACollaborationDataSource struct {
	pData *tfutils.ProviderData
}

func (d *B2BCPACollaborationDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_b2b_cpa_collaboration"
}

func (d *B2BCPACollaborationDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "B2B CPA collaboration",
		Attributes: map[string]schema.Attribute{
			"provider_target": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Target host to retrieve this data from. If not set, provider will use the top level settings.", "", "").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 128),
					stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9_-]+$"), "Must match :"+"^[a-zA-Z0-9_-]+$"),
				},
			},
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
						"internal_role": schema.StringAttribute{
							MarkdownDescription: "Specify the name of the authorized role of the internal partner in a business collaboration service. Each role is authorized for specific actions. For example, a <tt>Buyer</tt> role has the authority for purchasing actions.",
							Computed:            true,
						},
						"external_role": schema.StringAttribute{
							MarkdownDescription: "Specify the name of the authorized role of the external partner in a business collaboration service. Each role is authorized for specific actions. For example, a <tt>Supplier</tt> role has the authority for selling actions.",
							Computed:            true,
						},
						"process_specification": schema.StringAttribute{
							MarkdownDescription: "Specify the location of the process specification document that defines the interactions between the internal and external partners. For example, <tt>http://www.rosettanet.org/processes/3A4</tt> .",
							Computed:            true,
						},
						"service": schema.StringAttribute{
							MarkdownDescription: "Specify the value of the service that acts on the message. The value is used to specify and identify the value of the Service element in the outbound and inbound ebMS message header. The service is one of the following types. <ul><li>A business collaboration service for exchanging business messages.</li><li>An MSH signal service for exchanging MSH signals.</li></ul><p>The value of <tt>urn:oasis:names:tc:ebxml-msg:service;</tt> is an MSH signal service. Any other value represents a business collaboration service.</p>",
							Computed:            true,
						},
						"service_type": schema.StringAttribute{
							MarkdownDescription: "Specify the value of the service type. If you specify the type, the value is present in the type attribute of the <tt>Service</tt> element within the message to be sent. If the type is empty, the value of the <tt>Service</tt> element must be a URI.",
							Computed:            true,
						},
						"sender_msh_setting": schema.StringAttribute{
							MarkdownDescription: "Specify the name of the default MSH signal sender to send ebMS MSH signals to send MSH level signals that include <tt>Acknowledgment</tt> , <tt>Error</tt> , <tt>StatusRequest</tt> , <tt>StatusResponse</tt> , <tt>Ping</tt> , and <tt>Pong</tt> .",
							Computed:            true,
						},
						"receiver_msh_setting": schema.StringAttribute{
							MarkdownDescription: "Specify the name of the default MSH signal receiver to receive ebMS MSH signals to receive MSH level signals that include <tt>Acknowledgment</tt> , <tt>Error</tt> , <tt>StatusRequest</tt> , <tt>StatusResponse</tt> , <tt>Ping</tt> , and <tt>Pong</tt> .",
							Computed:            true,
						},
						"actions": schema.ListNestedAttribute{
							MarkdownDescription: "Specify CPA actions to bind. For a business collaboration, each action entry identifies a business message that a party can send or receive. For a collaboration of MSH level signal, the action overrides the sending or receiving behaviors of the default sender setting or default receiver setting.",
							NestedObject:        models.GetDmCPACollaborationActionDataSourceSchema(),
							Computed:            true,
						},
						"dependency_actions": actions.ActionsSchema,
						"provider_target": schema.StringAttribute{
							MarkdownDescription: tfutils.NewAttributeDescription("Target host to retrieve this data from. If not set, provider will use the top level settings.", "", "").String,
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *B2BCPACollaborationDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *B2BCPACollaborationDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data B2BCPACollaborationList
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if data.ProviderTarget.ValueString() != "" && !d.pData.Client.ValidTarget(data.ProviderTarget.ValueString()) {
		resp.Diagnostics.AddError("Invalid provider_target", fmt.Sprintf(`Target %q is not defined in the provider's 'targets' block. Available targets: %v`, data.ProviderTarget.ValueString(), d.pData.Client.GetTargetNames()))
		return
	}
	o := models.B2BCPACollaboration{
		AppDomain: data.AppDomain,
	}

	path := o.GetPath()
	if !data.Id.IsNull() {
		path = path + "/" + data.Id.ValueString()
	}

	res, err := d.pData.Client.Get(path, data.ProviderTarget)
	resFound := true
	if err != nil {
		if strings.Contains(err.Error(), "status 401") {
			_ = tfutils.DomainCredentialTest(d.pData.Client, &resp.Diagnostics, data.AppDomain.ValueString(), data.ProviderTarget)
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
	l := []models.B2BCPACollaboration{}
	if resFound {
		if value := res.Get(`B2BCPACollaboration`); value.Exists() {
			for _, v := range value.Array() {
				item := models.B2BCPACollaboration{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.B2BCPACollaborationObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.B2BCPACollaborationObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

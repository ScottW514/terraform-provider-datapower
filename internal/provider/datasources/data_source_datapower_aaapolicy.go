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

type AAAPolicyList struct {
	AppDomain types.String `tfsdk:"app_domain"`
	Result    types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &AAAPolicyDataSource{}
	_ datasource.DataSourceWithConfigure = &AAAPolicyDataSource{}
)

func NewAAAPolicyDataSource() datasource.DataSource {
	return &AAAPolicyDataSource{}
}

type AAAPolicyDataSource struct {
	pData *tfutils.ProviderData
}

func (d *AAAPolicyDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_aaapolicy"
}

func (d *AAAPolicyDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "AAA policy",
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
						"authorized_counter": schema.StringAttribute{
							MarkdownDescription: "Authorized counter",
							Computed:            true,
						},
						"rejected_counter": schema.StringAttribute{
							MarkdownDescription: "Rejected counter",
							Computed:            true,
						},
						"namespace_mapping": schema.ListNestedAttribute{
							MarkdownDescription: "Namespace mapping",
							NestedObject:        models.DmNamespaceMappingDataSourceSchema,
							Computed:            true,
						},
						"extract_identity": models.GetDmAAAPExtractIdentityDataSourceSchema("Identity extraction", "extract-identity", ""),
						"authenticate":     models.GetDmAAAPAuthenticateDataSourceSchema("Authentication", "authenticate", ""),
						"map_credentials":  models.GetDmAAAPMapCredentialsDataSourceSchema("Credential mapping", "map-credentials", ""),
						"extract_resource": models.GetDmAAAPExtractResourceDataSourceSchema("Resource extraction", "extract-resource", ""),
						"map_resource":     models.GetDmAAAPMapResourceDataSourceSchema("Resource mapping", "map-resource", ""),
						"authorize":        models.GetDmAAAPAuthorizeDataSourceSchema("Authorization", "authorize", ""),
						"post_process":     models.GetDmAAAPPostProcessDataSourceSchema("Postprocessing", "post-process", ""),
						"saml_attribute": schema.ListNestedAttribute{
							MarkdownDescription: "SAML attributes",
							NestedObject:        models.DmSAMLAttributeNameAndValueDataSourceSchema,
							Computed:            true,
						},
						"ltpa_attributes": schema.ListNestedAttribute{
							MarkdownDescription: "LTPA user attributes",
							NestedObject:        models.DmLTPAUserAttributeNameAndValueDataSourceSchema,
							Computed:            true,
						},
						"transaction_priority": schema.ListNestedAttribute{
							MarkdownDescription: "Transaction priority",
							NestedObject:        models.DmAAATransactionPriorityDataSourceSchema,
							Computed:            true,
						},
						"saml_valcred": schema.StringAttribute{
							MarkdownDescription: "SAML signature validation credentials",
							Computed:            true,
						},
						"saml_signing_key": schema.StringAttribute{
							MarkdownDescription: "SAML message signing key",
							Computed:            true,
						},
						"saml_signing_cert": schema.StringAttribute{
							MarkdownDescription: "SAML message signing certificate",
							Computed:            true,
						},
						"saml_signing_hash_alg": schema.StringAttribute{
							MarkdownDescription: "SAML signing message digest algorithm",
							Computed:            true,
						},
						"saml_signing_alg": schema.StringAttribute{
							MarkdownDescription: "SAML message signing algorithm",
							Computed:            true,
						},
						"lda_psuffix": schema.StringAttribute{
							MarkdownDescription: "LDAP suffix",
							Computed:            true,
						},
						"log_allowed": schema.BoolAttribute{
							MarkdownDescription: "Log allowed",
							Computed:            true,
						},
						"log_allowed_level": schema.StringAttribute{
							MarkdownDescription: "Log allowed level",
							Computed:            true,
						},
						"log_rejected": schema.BoolAttribute{
							MarkdownDescription: "Log rejected",
							Computed:            true,
						},
						"log_rejected_level": schema.StringAttribute{
							MarkdownDescription: "Log rejected level",
							Computed:            true,
						},
						"ws_secure_conversation_crypto_key": schema.StringAttribute{
							MarkdownDescription: "WS-Trust encryption recipient certificate",
							Computed:            true,
						},
						"saml_source_id_mapping_file": schema.StringAttribute{
							MarkdownDescription: "SAML Artifact mapping file",
							Computed:            true,
						},
						"ping_identity_compatibility": schema.BoolAttribute{
							MarkdownDescription: "PingFederate compatibility",
							Computed:            true,
						},
						"saml2_metadata_file": schema.StringAttribute{
							MarkdownDescription: "SAML 2.0 metadata file",
							Computed:            true,
						},
						"do_s_valve": schema.Int64Attribute{
							MarkdownDescription: "DoS flooding attack valve",
							Computed:            true,
						},
						"ldap_version": schema.StringAttribute{
							MarkdownDescription: "LDAP version",
							Computed:            true,
						},
						"enforce_soap_actor": schema.BoolAttribute{
							MarkdownDescription: "Enforce actor or role for WS-Security message",
							Computed:            true,
						},
						"ws_sec_actor_role_id": schema.StringAttribute{
							MarkdownDescription: "WS-Security actor or role identifier",
							Computed:            true,
						},
						"ausmhttp_header": schema.ListAttribute{
							MarkdownDescription: "HTTP headers",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"azsmhttp_header": schema.ListAttribute{
							MarkdownDescription: "HTTP headers",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"dyn_config": schema.StringAttribute{
							MarkdownDescription: "Dynamic configuration type",
							Computed:            true,
						},
						"external_aaa_template": schema.StringAttribute{
							MarkdownDescription: "External AAA policy template",
							Computed:            true,
						},
						"dyn_config_custom_url": schema.StringAttribute{
							MarkdownDescription: "Dynamic configuration custom URL",
							Computed:            true,
						},
						"dependency_actions": actions.ActionsSchema,
					},
				},
			},
		},
	}
}

func (d *AAAPolicyDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *AAAPolicyDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data AAAPolicyList
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.AAAPolicy{
		AppDomain: data.AppDomain,
	}

	res, err := d.pData.Client.Get(o.GetPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}
	l := []models.AAAPolicy{}
	if value := res.Get(`AAAPolicy`); value.Exists() {
		for _, v := range value.Array() {
			item := models.AAAPolicy{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.AAAPolicyObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.AAAPolicyObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

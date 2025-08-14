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

type StylePolicyActionList struct {
	AppDomain types.String `tfsdk:"app_domain"`
	Result    types.List   `tfsdk:"result"`
}

var (
	_ datasource.DataSource              = &StylePolicyActionDataSource{}
	_ datasource.DataSourceWithConfigure = &StylePolicyActionDataSource{}
)

func NewStylePolicyActionDataSource() datasource.DataSource {
	return &StylePolicyActionDataSource{}
}

type StylePolicyActionDataSource struct {
	pData *tfutils.ProviderData
}

func (d *StylePolicyActionDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_stylepolicyaction"
}

func (d *StylePolicyActionDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Processing action",
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
						"type": schema.StringAttribute{
							MarkdownDescription: "Action type",
							Computed:            true,
						},
						"input": schema.StringAttribute{
							MarkdownDescription: "Input",
							Computed:            true,
						},
						"transform": schema.StringAttribute{
							MarkdownDescription: "Transform file",
							Computed:            true,
						},
						"parse_settings_reference": models.GetDmDynamicParseSettingsReferenceDataSourceSchema("Parse settings", "parse-settings-reference", ""),
						"parse_metrics_result_type": schema.StringAttribute{
							MarkdownDescription: "Parse metrics result type",
							Computed:            true,
						},
						"parse_metrics_result_location": schema.StringAttribute{
							MarkdownDescription: "Parse metrics result location",
							Computed:            true,
						},
						"input_language": schema.StringAttribute{
							MarkdownDescription: "Input language",
							Computed:            true,
						},
						"dfdl_input_root_name": schema.StringAttribute{
							MarkdownDescription: "DFDL input root name",
							Computed:            true,
						},
						"input_descriptor": schema.StringAttribute{
							MarkdownDescription: "Input descriptor",
							Computed:            true,
						},
						"output_descriptor": schema.StringAttribute{
							MarkdownDescription: "Output descriptor",
							Computed:            true,
						},
						"transform_language": schema.StringAttribute{
							MarkdownDescription: "Transform language",
							Computed:            true,
						},
						"output_language": schema.StringAttribute{
							MarkdownDescription: "Output language",
							Computed:            true,
						},
						"tx_map": schema.StringAttribute{
							MarkdownDescription: "ITX map file",
							Computed:            true,
						},
						"gateway_script_location": schema.StringAttribute{
							MarkdownDescription: "GatewayScript file",
							Computed:            true,
						},
						"action_debug": schema.BoolAttribute{
							MarkdownDescription: "Enable GatewayScript debug",
							Computed:            true,
						},
						"tx_top_level_map": schema.StringAttribute{
							MarkdownDescription: "ITX top-level map name",
							Computed:            true,
						},
						"tx_mode": schema.StringAttribute{
							MarkdownDescription: "ITX map mode",
							Computed:            true,
						},
						"tx_audit_log": schema.StringAttribute{
							MarkdownDescription: "ITX audit log",
							Computed:            true,
						},
						"output": schema.StringAttribute{
							MarkdownDescription: "Output",
							Computed:            true,
						},
						"no_transcode_utf8": schema.BoolAttribute{
							MarkdownDescription: "Retain input encoding",
							Computed:            true,
						},
						"named_in_out_location_type": schema.StringAttribute{
							MarkdownDescription: "Locate named inputs and outputs",
							Computed:            true,
						},
						"named_inputs": schema.ListNestedAttribute{
							MarkdownDescription: "Named inputs",
							NestedObject:        models.DmNamedInOutDataSourceSchema,
							Computed:            true,
						},
						"named_outputs": schema.ListNestedAttribute{
							MarkdownDescription: "Named outputs",
							NestedObject:        models.DmNamedInOutDataSourceSchema,
							Computed:            true,
						},
						"destination": schema.StringAttribute{
							MarkdownDescription: "External URL",
							Computed:            true,
						},
						"schema_url": schema.StringAttribute{
							MarkdownDescription: "Schema URL",
							Computed:            true,
						},
						"json_schema_url": schema.StringAttribute{
							MarkdownDescription: "JSON schema URL",
							Computed:            true,
						},
						"wsdl_url": schema.StringAttribute{
							MarkdownDescription: "WSDL URL",
							Computed:            true,
						},
						"policy": schema.StringAttribute{
							MarkdownDescription: "URL rewrite policy",
							Computed:            true,
						},
						"aaa": schema.StringAttribute{
							MarkdownDescription: "AAA policy",
							Computed:            true,
						},
						"dynamic_schema": schema.StringAttribute{
							MarkdownDescription: "Dynamic schema",
							Computed:            true,
						},
						"dynamic_stylesheet": schema.StringAttribute{
							MarkdownDescription: "Dynamic stylesheet",
							Computed:            true,
						},
						"input_conversion": schema.StringAttribute{
							MarkdownDescription: "Input conversion",
							Computed:            true,
						},
						"x_path": schema.StringAttribute{
							MarkdownDescription: "XPath",
							Computed:            true,
						},
						"variable": schema.StringAttribute{
							MarkdownDescription: "Variable name",
							Computed:            true,
						},
						"value": schema.StringAttribute{
							MarkdownDescription: "Variable value",
							Computed:            true,
						},
						"ssl_client_config_type": schema.StringAttribute{
							MarkdownDescription: "TLS client type",
							Computed:            true,
						},
						"ssl_client_cred": schema.StringAttribute{
							MarkdownDescription: "TLS client profile",
							Computed:            true,
						},
						"attachment_uri": schema.StringAttribute{
							MarkdownDescription: "Attachment URI",
							Computed:            true,
						},
						"stylesheet_parameters": schema.ListNestedAttribute{
							MarkdownDescription: "Stylesheet parameters",
							NestedObject:        models.DmStylesheetParameterDataSourceSchema,
							Computed:            true,
						},
						"error_mode": schema.StringAttribute{
							MarkdownDescription: "Error mode",
							Computed:            true,
						},
						"error_input": schema.StringAttribute{
							MarkdownDescription: "Error input",
							Computed:            true,
						},
						"error_output": schema.StringAttribute{
							MarkdownDescription: "Error output",
							Computed:            true,
						},
						"rule": schema.StringAttribute{
							MarkdownDescription: "Processing rule",
							Computed:            true,
						},
						"output_type": schema.StringAttribute{
							MarkdownDescription: "Output type",
							Computed:            true,
						},
						"log_level": schema.StringAttribute{
							MarkdownDescription: "Log level",
							Computed:            true,
						},
						"log_type": schema.StringAttribute{
							MarkdownDescription: "Log type",
							Computed:            true,
						},
						"transactional": schema.BoolAttribute{
							MarkdownDescription: "Transactional",
							Computed:            true,
						},
						"checkpoint_event": schema.StringAttribute{
							MarkdownDescription: "Event",
							Computed:            true,
						},
						"slm_policy": schema.StringAttribute{
							MarkdownDescription: "SLM policy",
							Computed:            true,
						},
						"sql_data_source": schema.StringAttribute{
							MarkdownDescription: "SQL Data Source",
							Computed:            true,
						},
						"sql_text": schema.StringAttribute{
							MarkdownDescription: "SQL text",
							Computed:            true,
						},
						"soap_validation": schema.StringAttribute{
							MarkdownDescription: "SOAP validation",
							Computed:            true,
						},
						"sql_source_type": schema.StringAttribute{
							MarkdownDescription: "SQL input method",
							Computed:            true,
						},
						"jose_serialization_type": schema.StringAttribute{
							MarkdownDescription: "Serialization",
							Computed:            true,
						},
						"jwe_enc_algorithm": schema.StringAttribute{
							MarkdownDescription: "Algorithm",
							Computed:            true,
						},
						"jws_signature_object": schema.StringAttribute{
							MarkdownDescription: "Signature",
							Computed:            true,
						},
						"jwe_header_object": schema.StringAttribute{
							MarkdownDescription: "JWE Header",
							Computed:            true,
						},
						"jose_verify_type": schema.StringAttribute{
							MarkdownDescription: "Identifier type",
							Computed:            true,
						},
						"jose_decrypt_type": schema.StringAttribute{
							MarkdownDescription: "Identifier Type",
							Computed:            true,
						},
						"signature_identifier": schema.ListAttribute{
							MarkdownDescription: "Signature Identifiers",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"recipient_identifier": schema.ListAttribute{
							MarkdownDescription: "Recipient Identifiers",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"single_certificate": schema.StringAttribute{
							MarkdownDescription: "Certificate",
							Computed:            true,
						},
						"single_key": schema.StringAttribute{
							MarkdownDescription: "Private Key",
							Computed:            true,
						},
						"single_ss_key": schema.StringAttribute{
							MarkdownDescription: "Shared Secret Key",
							Computed:            true,
						},
						"jwe_direct_key_object": schema.StringAttribute{
							MarkdownDescription: "Direct Key",
							Computed:            true,
						},
						"jws_verify_strip_signature": schema.BoolAttribute{
							MarkdownDescription: "Strip Signature",
							Computed:            true,
						},
						"asynchronous": schema.BoolAttribute{
							MarkdownDescription: "Asynchronous",
							Computed:            true,
						},
						"condition": schema.ListNestedAttribute{
							MarkdownDescription: "Condition",
							NestedObject:        models.DmConditionDataSourceSchema,
							Computed:            true,
						},
						"results_mode": schema.StringAttribute{
							MarkdownDescription: "Multi-way results mode",
							Computed:            true,
						},
						"retry_count": schema.Int64Attribute{
							MarkdownDescription: "Number of retries",
							Computed:            true,
						},
						"retry_interval": schema.Int64Attribute{
							MarkdownDescription: "Retry interval",
							Computed:            true,
						},
						"multiple_outputs": schema.BoolAttribute{
							MarkdownDescription: "Use multiple outputs",
							Computed:            true,
						},
						"iterator_type": schema.StringAttribute{
							MarkdownDescription: "Iterator type",
							Computed:            true,
						},
						"iterator_expression": schema.StringAttribute{
							MarkdownDescription: "XPath expression",
							Computed:            true,
						},
						"iterator_count": schema.Int64Attribute{
							MarkdownDescription: "Iterator count",
							Computed:            true,
						},
						"loop_action": schema.StringAttribute{
							MarkdownDescription: "Loop action",
							Computed:            true,
						},
						"async_action": schema.ListAttribute{
							MarkdownDescription: "Asynchronous actions",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"timeout": schema.Int64Attribute{
							MarkdownDescription: "Timeout",
							Computed:            true,
						},
						"wsdl_port_q_name": schema.StringAttribute{
							MarkdownDescription: "WSDL Port QName",
							Computed:            true,
						},
						"wsdl_operation_name": schema.StringAttribute{
							MarkdownDescription: "WSDL Operation Name",
							Computed:            true,
						},
						"wsdl_message_direction_or_name": schema.StringAttribute{
							MarkdownDescription: "WSDL Message Direction or Name",
							Computed:            true,
						},
						"wsdl_attachment_part": schema.StringAttribute{
							MarkdownDescription: "WSDL Attachment Part Name",
							Computed:            true,
						},
						"method_rewrite_type": schema.StringAttribute{
							MarkdownDescription: "Method",
							Computed:            true,
						},
						"method_type": schema.StringAttribute{
							MarkdownDescription: "Method",
							Computed:            true,
						},
						"method_type2": schema.StringAttribute{
							MarkdownDescription: "Method",
							Computed:            true,
						},
						"policy_key": schema.StringAttribute{
							MarkdownDescription: "Policy identifier (for internal use)",
							Computed:            true,
						},
						"dependency_actions": actions.ActionsSchema,
					},
				},
			},
		},
	}
}

func (d *StylePolicyActionDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (d *StylePolicyActionDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	d.pData.Mu.Lock()
	defer d.pData.Mu.Unlock()

	var data StylePolicyActionList
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	o := models.StylePolicyAction{
		AppDomain: data.AppDomain,
	}

	res, err := d.pData.Client.Get(o.GetPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}
	l := []models.StylePolicyAction{}
	if value := res.Get(`StylePolicyAction`); value.Exists() {
		for _, v := range value.Array() {
			item := models.StylePolicyAction{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
	}
	if len(l) > 0 {
		var diag diag.Diagnostics
		data.Result, diag = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: models.StylePolicyActionObjectType}, l)
		resp.Diagnostics.Append(diag...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else {
		data.Result = types.ListNull(types.ObjectType{AttrTypes: models.StylePolicyActionObjectType})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

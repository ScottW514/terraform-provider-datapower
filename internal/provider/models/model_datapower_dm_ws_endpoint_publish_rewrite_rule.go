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

package models

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	DataSourceSchema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	ResourceSchema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type DmWSEndpointPublishRewriteRule struct {
	ServicePortMatchRegexp    types.String `tfsdk:"service_port_match_regexp"`
	PublishedEndpointProtocol types.String `tfsdk:"published_endpoint_protocol"`
	PublishedEndpointHostname types.String `tfsdk:"published_endpoint_hostname"`
	PublishedEndpointPort     types.Int64  `tfsdk:"published_endpoint_port"`
	PublishedEndpointUri      types.String `tfsdk:"published_endpoint_uri"`
}

var DmWSEndpointPublishRewriteRuleObjectType = map[string]attr.Type{
	"service_port_match_regexp":   types.StringType,
	"published_endpoint_protocol": types.StringType,
	"published_endpoint_hostname": types.StringType,
	"published_endpoint_port":     types.Int64Type,
	"published_endpoint_uri":      types.StringType,
}
var DmWSEndpointPublishRewriteRuleObjectDefault = map[string]attr.Value{
	"service_port_match_regexp":   types.StringValue(".*"),
	"published_endpoint_protocol": types.StringValue("default"),
	"published_endpoint_hostname": types.StringNull(),
	"published_endpoint_port":     types.Int64Null(),
	"published_endpoint_uri":      types.StringNull(),
}

func GetDmWSEndpointPublishRewriteRuleDataSourceSchema() DataSourceSchema.NestedAttributeObject {
	var DmWSEndpointPublishRewriteRuleDataSourceSchema = DataSourceSchema.NestedAttributeObject{
		Attributes: map[string]DataSourceSchema.Attribute{
			"service_port_match_regexp": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify a PCRE as the match criteria to identify the web services port. When the PCRE matches a Service Port in WSDL, the rewrite rules is applied.", "service-port-match", "").AddDefaultValue(".*").String,
				Computed:            true,
			},
			"published_endpoint_protocol": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Select the protocol portion of the rewritten web service binding used by the publish endpoint. The protocol can be different from the one in the WSDL.", "published-endpoint-protocol", "").AddStringEnum("default", "http", "https").AddDefaultValue("default").String,
				Computed:            true,
			},
			"published_endpoint_hostname": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the URL portion of the rewritten web service binding that specifies the host name or IP address.", "published-endpoint-hostname", "").String,
				Computed:            true,
			},
			"published_endpoint_port": DataSourceSchema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the URL portion of the rewritten web service binding that specifies the port. If 0, uses the value from the WSDL.", "published-endpoint-port", "").String,
				Computed:            true,
			},
			"published_endpoint_uri": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the URL portion of the rewritten web service binding that specifies the local path. If not specified, uses the value from the WSDL.", "published-endpoint-path", "").String,
				Computed:            true,
			},
		},
	}
	return DmWSEndpointPublishRewriteRuleDataSourceSchema
}
func GetDmWSEndpointPublishRewriteRuleResourceSchema() ResourceSchema.NestedAttributeObject {
	var DmWSEndpointPublishRewriteRuleResourceSchema = ResourceSchema.NestedAttributeObject{
		Attributes: map[string]ResourceSchema.Attribute{
			"service_port_match_regexp": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify a PCRE as the match criteria to identify the web services port. When the PCRE matches a Service Port in WSDL, the rewrite rules is applied.", "service-port-match", "").AddDefaultValue(".*").String,
				Computed:            true,
				Optional:            true,
				Default:             stringdefault.StaticString(".*"),
			},
			"published_endpoint_protocol": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Select the protocol portion of the rewritten web service binding used by the publish endpoint. The protocol can be different from the one in the WSDL.", "published-endpoint-protocol", "").AddStringEnum("default", "http", "https").AddDefaultValue("default").String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("default", "http", "https"),
				},
				Default: stringdefault.StaticString("default"),
			},
			"published_endpoint_hostname": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the URL portion of the rewritten web service binding that specifies the host name or IP address.", "published-endpoint-hostname", "").String,
				Optional:            true,
			},
			"published_endpoint_port": ResourceSchema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the URL portion of the rewritten web service binding that specifies the port. If 0, uses the value from the WSDL.", "published-endpoint-port", "").String,
				Optional:            true,
			},
			"published_endpoint_uri": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the URL portion of the rewritten web service binding that specifies the local path. If not specified, uses the value from the WSDL.", "published-endpoint-path", "").String,
				Optional:            true,
			},
		},
	}
	return DmWSEndpointPublishRewriteRuleResourceSchema
}

func (data DmWSEndpointPublishRewriteRule) IsNull() bool {
	if !data.ServicePortMatchRegexp.IsNull() {
		return false
	}
	if !data.PublishedEndpointProtocol.IsNull() {
		return false
	}
	if !data.PublishedEndpointHostname.IsNull() {
		return false
	}
	if !data.PublishedEndpointPort.IsNull() {
		return false
	}
	if !data.PublishedEndpointUri.IsNull() {
		return false
	}
	return true
}

func (data DmWSEndpointPublishRewriteRule) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""

	if !data.ServicePortMatchRegexp.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ServicePortMatchRegexp`, data.ServicePortMatchRegexp.ValueString())
	}
	if !data.PublishedEndpointProtocol.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PublishedEndpointProtocol`, data.PublishedEndpointProtocol.ValueString())
	}
	if !data.PublishedEndpointHostname.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PublishedEndpointHostname`, data.PublishedEndpointHostname.ValueString())
	}
	if !data.PublishedEndpointPort.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PublishedEndpointPort`, data.PublishedEndpointPort.ValueInt64())
	}
	if !data.PublishedEndpointUri.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PublishedEndpointURI`, data.PublishedEndpointUri.ValueString())
	}
	return body
}

func (data *DmWSEndpointPublishRewriteRule) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `ServicePortMatchRegexp`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ServicePortMatchRegexp = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ServicePortMatchRegexp = types.StringValue(".*")
	}
	if value := res.Get(pathRoot + `PublishedEndpointProtocol`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PublishedEndpointProtocol = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PublishedEndpointProtocol = types.StringValue("default")
	}
	if value := res.Get(pathRoot + `PublishedEndpointHostname`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PublishedEndpointHostname = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PublishedEndpointHostname = types.StringNull()
	}
	if value := res.Get(pathRoot + `PublishedEndpointPort`); value.Exists() {
		data.PublishedEndpointPort = types.Int64Value(value.Int())
	} else {
		data.PublishedEndpointPort = types.Int64Null()
	}
	if value := res.Get(pathRoot + `PublishedEndpointURI`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PublishedEndpointUri = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PublishedEndpointUri = types.StringNull()
	}
}

func (data *DmWSEndpointPublishRewriteRule) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `ServicePortMatchRegexp`); value.Exists() && !data.ServicePortMatchRegexp.IsNull() {
		data.ServicePortMatchRegexp = tfutils.ParseStringFromGJSON(value)
	} else if data.ServicePortMatchRegexp.ValueString() != ".*" {
		data.ServicePortMatchRegexp = types.StringNull()
	}
	if value := res.Get(pathRoot + `PublishedEndpointProtocol`); value.Exists() && !data.PublishedEndpointProtocol.IsNull() {
		data.PublishedEndpointProtocol = tfutils.ParseStringFromGJSON(value)
	} else if data.PublishedEndpointProtocol.ValueString() != "default" {
		data.PublishedEndpointProtocol = types.StringNull()
	}
	if value := res.Get(pathRoot + `PublishedEndpointHostname`); value.Exists() && !data.PublishedEndpointHostname.IsNull() {
		data.PublishedEndpointHostname = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PublishedEndpointHostname = types.StringNull()
	}
	if value := res.Get(pathRoot + `PublishedEndpointPort`); value.Exists() && !data.PublishedEndpointPort.IsNull() {
		data.PublishedEndpointPort = types.Int64Value(value.Int())
	} else {
		data.PublishedEndpointPort = types.Int64Null()
	}
	if value := res.Get(pathRoot + `PublishedEndpointURI`); value.Exists() && !data.PublishedEndpointUri.IsNull() {
		data.PublishedEndpointUri = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PublishedEndpointUri = types.StringNull()
	}
}

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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/validators"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type DmWSEndpointSubscriptionLocalRule struct {
	Subscription          types.String `tfsdk:"subscription"`
	LocalEndpointProtocol types.String `tfsdk:"local_endpoint_protocol"`
	LocalEndpointHostname types.String `tfsdk:"local_endpoint_hostname"`
	LocalEndpointPort     types.Int64  `tfsdk:"local_endpoint_port"`
	LocalEndpointUri      types.String `tfsdk:"local_endpoint_uri"`
	FrontProtocol         types.String `tfsdk:"front_protocol"`
	UseFrontProtocol      types.Bool   `tfsdk:"use_front_protocol"`
	WsdlBindingProtocol   types.String `tfsdk:"wsdl_binding_protocol"`
	FrontsidePortSuffix   types.String `tfsdk:"frontside_port_suffix"`
}

var DmWSEndpointSubscriptionLocalRuleFrontProtocolCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "use_front_protocol",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"true"},
}
var DmWSEndpointSubscriptionLocalRuleLocalEndpointProtocolIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "use_front_protocol",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"false"},
}
var DmWSEndpointSubscriptionLocalRuleLocalEndpointHostnameIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "use_front_protocol",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"false"},
}
var DmWSEndpointSubscriptionLocalRuleLocalEndpointPortIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "use_front_protocol",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"false"},
}
var DmWSEndpointSubscriptionLocalRuleFrontProtocolIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var DmWSEndpointSubscriptionLocalRuleObjectType = map[string]attr.Type{
	"subscription":            types.StringType,
	"local_endpoint_protocol": types.StringType,
	"local_endpoint_hostname": types.StringType,
	"local_endpoint_port":     types.Int64Type,
	"local_endpoint_uri":      types.StringType,
	"front_protocol":          types.StringType,
	"use_front_protocol":      types.BoolType,
	"wsdl_binding_protocol":   types.StringType,
	"frontside_port_suffix":   types.StringType,
}
var DmWSEndpointSubscriptionLocalRuleObjectDefault = map[string]attr.Value{
	"subscription":            types.StringNull(),
	"local_endpoint_protocol": types.StringValue("default"),
	"local_endpoint_hostname": types.StringValue("0.0.0.0"),
	"local_endpoint_port":     types.Int64Null(),
	"local_endpoint_uri":      types.StringNull(),
	"front_protocol":          types.StringNull(),
	"use_front_protocol":      types.BoolValue(false),
	"wsdl_binding_protocol":   types.StringValue("default"),
	"frontside_port_suffix":   types.StringNull(),
}

func GetDmWSEndpointSubscriptionLocalRuleDataSourceSchema() DataSourceSchema.NestedAttributeObject {
	var DmWSEndpointSubscriptionLocalRuleDataSourceSchema = DataSourceSchema.NestedAttributeObject{
		Attributes: map[string]DataSourceSchema.Attribute{
			"subscription": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Select the subscription and configure its endpoint.", "subscription", "").String,
				Computed:            true,
			},
			"local_endpoint_protocol": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Select the protocol portion of the rewritten web service binding used by the local endpoint. The protocol can be different from the one in the WSDL.", "local-endpoint-protocol", "").AddStringEnum("default", "http", "https").AddDefaultValue("default").AddNotValidWhen(DmWSEndpointSubscriptionLocalRuleLocalEndpointProtocolIgnoreVal.String()).String,
				Computed:            true,
			},
			"local_endpoint_hostname": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the URL protion of the rewritten web service binding that specifies the host name or IP address. A value of 0.0.0.0 indicates that the Web Service Proxy listens on all of the interfaces. Alternatively, you can specify a Host Alias.", "local-endpoint-hostname", "").AddDefaultValue("0.0.0.0").AddNotValidWhen(DmWSEndpointSubscriptionLocalRuleLocalEndpointHostnameIgnoreVal.String()).String,
				Computed:            true,
			},
			"local_endpoint_port": DataSourceSchema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the URL portion of the rewritten web service binding that specifies the port. If 0, uses the value from the WSDL.", "local-endpoint-port", "").AddNotValidWhen(DmWSEndpointSubscriptionLocalRuleLocalEndpointPortIgnoreVal.String()).String,
				Computed:            true,
			},
			"local_endpoint_uri": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the URL portion of the rewritten web service binding that specifies the local path. If not specified, uses the value that is specified in the WSDL.", "local-endpoint-uri", "").String,
				Computed:            true,
			},
			"front_protocol": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Select the Front Side Handler to use to determine the IP address, port, and protocol.", "", "").AddRequiredWhen(DmWSEndpointSubscriptionLocalRuleFrontProtocolCondVal.String()).AddNotValidWhen(DmWSEndpointSubscriptionLocalRuleFrontProtocolIgnoreVal.String()).String,
				Computed:            true,
			},
			"use_front_protocol": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Use a Front Side Handler to determine the IP address, port, and protocol for matching WSDL service port. Selecting this mode overrides these values in this rewrite rule.", "", "").AddDefaultValue("false").String,
				Computed:            true,
			},
			"wsdl_binding_protocol": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the WSDL binding protocol to use in the rewritten Web service.", "", "").AddStringEnum("default", "soap-11", "soap-12", "http-get", "http-post").AddDefaultValue("default").String,
				Computed:            true,
			},
			"frontside_port_suffix": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify a suffix to add to the name of the WSDL port that will be used to represent this service endpoint in the rewritten Web service. If empty, rewrite the original port. The original port can only be rewritten once.", "", "").String,
				Computed:            true,
			},
		},
	}
	return DmWSEndpointSubscriptionLocalRuleDataSourceSchema
}
func GetDmWSEndpointSubscriptionLocalRuleResourceSchema() ResourceSchema.NestedAttributeObject {
	var DmWSEndpointSubscriptionLocalRuleResourceSchema = ResourceSchema.NestedAttributeObject{
		Attributes: map[string]ResourceSchema.Attribute{
			"subscription": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Select the subscription and configure its endpoint.", "subscription", "").String,
				Optional:            true,
			},
			"local_endpoint_protocol": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Select the protocol portion of the rewritten web service binding used by the local endpoint. The protocol can be different from the one in the WSDL.", "local-endpoint-protocol", "").AddStringEnum("default", "http", "https").AddDefaultValue("default").AddNotValidWhen(DmWSEndpointSubscriptionLocalRuleLocalEndpointProtocolIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("default", "http", "https"),
					validators.ConditionalRequiredString(validators.Evaluation{}, DmWSEndpointSubscriptionLocalRuleLocalEndpointProtocolIgnoreVal, true),
				},
				Default: stringdefault.StaticString("default"),
			},
			"local_endpoint_hostname": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the URL protion of the rewritten web service binding that specifies the host name or IP address. A value of 0.0.0.0 indicates that the Web Service Proxy listens on all of the interfaces. Alternatively, you can specify a Host Alias.", "local-endpoint-hostname", "").AddDefaultValue("0.0.0.0").AddNotValidWhen(DmWSEndpointSubscriptionLocalRuleLocalEndpointHostnameIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Default:             stringdefault.StaticString("0.0.0.0"),
			},
			"local_endpoint_port": ResourceSchema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the URL portion of the rewritten web service binding that specifies the port. If 0, uses the value from the WSDL.", "local-endpoint-port", "").AddNotValidWhen(DmWSEndpointSubscriptionLocalRuleLocalEndpointPortIgnoreVal.String()).String,
				Optional:            true,
			},
			"local_endpoint_uri": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the URL portion of the rewritten web service binding that specifies the local path. If not specified, uses the value that is specified in the WSDL.", "local-endpoint-uri", "").String,
				Optional:            true,
			},
			"front_protocol": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Select the Front Side Handler to use to determine the IP address, port, and protocol.", "", "").AddRequiredWhen(DmWSEndpointSubscriptionLocalRuleFrontProtocolCondVal.String()).AddNotValidWhen(DmWSEndpointSubscriptionLocalRuleFrontProtocolIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(DmWSEndpointSubscriptionLocalRuleFrontProtocolCondVal, DmWSEndpointSubscriptionLocalRuleFrontProtocolIgnoreVal, false),
				},
			},
			"use_front_protocol": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Use a Front Side Handler to determine the IP address, port, and protocol for matching WSDL service port. Selecting this mode overrides these values in this rewrite rule.", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"wsdl_binding_protocol": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the WSDL binding protocol to use in the rewritten Web service.", "", "").AddStringEnum("default", "soap-11", "soap-12", "http-get", "http-post").AddDefaultValue("default").String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("default", "soap-11", "soap-12", "http-get", "http-post"),
				},
				Default: stringdefault.StaticString("default"),
			},
			"frontside_port_suffix": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify a suffix to add to the name of the WSDL port that will be used to represent this service endpoint in the rewritten Web service. If empty, rewrite the original port. The original port can only be rewritten once.", "", "").String,
				Optional:            true,
			},
		},
	}
	return DmWSEndpointSubscriptionLocalRuleResourceSchema
}

func (data DmWSEndpointSubscriptionLocalRule) IsNull() bool {
	if !data.Subscription.IsNull() {
		return false
	}
	if !data.LocalEndpointProtocol.IsNull() {
		return false
	}
	if !data.LocalEndpointHostname.IsNull() {
		return false
	}
	if !data.LocalEndpointPort.IsNull() {
		return false
	}
	if !data.LocalEndpointUri.IsNull() {
		return false
	}
	if !data.FrontProtocol.IsNull() {
		return false
	}
	if !data.UseFrontProtocol.IsNull() {
		return false
	}
	if !data.WsdlBindingProtocol.IsNull() {
		return false
	}
	if !data.FrontsidePortSuffix.IsNull() {
		return false
	}
	return true
}

func (data DmWSEndpointSubscriptionLocalRule) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""

	if !data.Subscription.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Subscription`, data.Subscription.ValueString())
	}
	if !data.LocalEndpointProtocol.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LocalEndpointProtocol`, data.LocalEndpointProtocol.ValueString())
	}
	if !data.LocalEndpointHostname.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LocalEndpointHostname`, data.LocalEndpointHostname.ValueString())
	}
	if !data.LocalEndpointPort.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LocalEndpointPort`, data.LocalEndpointPort.ValueInt64())
	}
	if !data.LocalEndpointUri.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LocalEndpointURI`, data.LocalEndpointUri.ValueString())
	}
	if !data.FrontProtocol.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`FrontProtocol`, data.FrontProtocol.ValueString())
	}
	if !data.UseFrontProtocol.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UseFrontProtocol`, tfutils.StringFromBool(data.UseFrontProtocol, ""))
	}
	if !data.WsdlBindingProtocol.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`WSDLBindingProtocol`, data.WsdlBindingProtocol.ValueString())
	}
	if !data.FrontsidePortSuffix.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`FrontsidePortSuffix`, data.FrontsidePortSuffix.ValueString())
	}
	return body
}

func (data *DmWSEndpointSubscriptionLocalRule) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Subscription`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Subscription = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Subscription = types.StringNull()
	}
	if value := res.Get(pathRoot + `LocalEndpointProtocol`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.LocalEndpointProtocol = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LocalEndpointProtocol = types.StringValue("default")
	}
	if value := res.Get(pathRoot + `LocalEndpointHostname`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.LocalEndpointHostname = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LocalEndpointHostname = types.StringValue("0.0.0.0")
	}
	if value := res.Get(pathRoot + `LocalEndpointPort`); value.Exists() {
		data.LocalEndpointPort = types.Int64Value(value.Int())
	} else {
		data.LocalEndpointPort = types.Int64Null()
	}
	if value := res.Get(pathRoot + `LocalEndpointURI`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.LocalEndpointUri = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LocalEndpointUri = types.StringNull()
	}
	if value := res.Get(pathRoot + `FrontProtocol`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.FrontProtocol = tfutils.ParseStringFromGJSON(value)
	} else {
		data.FrontProtocol = types.StringNull()
	}
	if value := res.Get(pathRoot + `UseFrontProtocol`); value.Exists() {
		data.UseFrontProtocol = tfutils.BoolFromString(value.String())
	} else {
		data.UseFrontProtocol = types.BoolNull()
	}
	if value := res.Get(pathRoot + `WSDLBindingProtocol`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.WsdlBindingProtocol = tfutils.ParseStringFromGJSON(value)
	} else {
		data.WsdlBindingProtocol = types.StringValue("default")
	}
	if value := res.Get(pathRoot + `FrontsidePortSuffix`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.FrontsidePortSuffix = tfutils.ParseStringFromGJSON(value)
	} else {
		data.FrontsidePortSuffix = types.StringNull()
	}
}

func (data *DmWSEndpointSubscriptionLocalRule) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Subscription`); value.Exists() && !data.Subscription.IsNull() {
		data.Subscription = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Subscription = types.StringNull()
	}
	if value := res.Get(pathRoot + `LocalEndpointProtocol`); value.Exists() && !data.LocalEndpointProtocol.IsNull() {
		data.LocalEndpointProtocol = tfutils.ParseStringFromGJSON(value)
	} else if data.LocalEndpointProtocol.ValueString() != "default" {
		data.LocalEndpointProtocol = types.StringNull()
	}
	if value := res.Get(pathRoot + `LocalEndpointHostname`); value.Exists() && !data.LocalEndpointHostname.IsNull() {
		data.LocalEndpointHostname = tfutils.ParseStringFromGJSON(value)
	} else if data.LocalEndpointHostname.ValueString() != "0.0.0.0" {
		data.LocalEndpointHostname = types.StringNull()
	}
	if value := res.Get(pathRoot + `LocalEndpointPort`); value.Exists() && !data.LocalEndpointPort.IsNull() {
		data.LocalEndpointPort = types.Int64Value(value.Int())
	} else {
		data.LocalEndpointPort = types.Int64Null()
	}
	if value := res.Get(pathRoot + `LocalEndpointURI`); value.Exists() && !data.LocalEndpointUri.IsNull() {
		data.LocalEndpointUri = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LocalEndpointUri = types.StringNull()
	}
	if value := res.Get(pathRoot + `FrontProtocol`); value.Exists() && !data.FrontProtocol.IsNull() {
		data.FrontProtocol = tfutils.ParseStringFromGJSON(value)
	} else {
		data.FrontProtocol = types.StringNull()
	}
	if value := res.Get(pathRoot + `UseFrontProtocol`); value.Exists() && !data.UseFrontProtocol.IsNull() {
		data.UseFrontProtocol = tfutils.BoolFromString(value.String())
	} else if data.UseFrontProtocol.ValueBool() {
		data.UseFrontProtocol = types.BoolNull()
	}
	if value := res.Get(pathRoot + `WSDLBindingProtocol`); value.Exists() && !data.WsdlBindingProtocol.IsNull() {
		data.WsdlBindingProtocol = tfutils.ParseStringFromGJSON(value)
	} else if data.WsdlBindingProtocol.ValueString() != "default" {
		data.WsdlBindingProtocol = types.StringNull()
	}
	if value := res.Get(pathRoot + `FrontsidePortSuffix`); value.Exists() && !data.FrontsidePortSuffix.IsNull() {
		data.FrontsidePortSuffix = tfutils.ParseStringFromGJSON(value)
	} else {
		data.FrontsidePortSuffix = types.StringNull()
	}
}

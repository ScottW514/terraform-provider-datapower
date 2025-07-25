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

type DmWSEndpointSubscriptionRemoteRule struct {
	Subscription           types.String `tfsdk:"subscription"`
	RemoteEndpointProtocol types.String `tfsdk:"remote_endpoint_protocol"`
	RemoteEndpointHostname types.String `tfsdk:"remote_endpoint_hostname"`
	RemoteEndpointPort     types.Int64  `tfsdk:"remote_endpoint_port"`
	RemoteEndpointUri      types.String `tfsdk:"remote_endpoint_uri"`
	RemoteMqqm             types.String `tfsdk:"remote_mqqm"`
	RemoteMqManager        types.String `tfsdk:"remote_mq_manager"`
	RemoteTibcoEms         types.String `tfsdk:"remote_tibco_ems"`
	RemoteWebSphereJms     types.String `tfsdk:"remote_web_sphere_jms"`
}

var DmWSEndpointSubscriptionRemoteRuleObjectType = map[string]attr.Type{
	"subscription":             types.StringType,
	"remote_endpoint_protocol": types.StringType,
	"remote_endpoint_hostname": types.StringType,
	"remote_endpoint_port":     types.Int64Type,
	"remote_endpoint_uri":      types.StringType,
	"remote_mqqm":              types.StringType,
	"remote_mq_manager":        types.StringType,
	"remote_tibco_ems":         types.StringType,
	"remote_web_sphere_jms":    types.StringType,
}
var DmWSEndpointSubscriptionRemoteRuleObjectDefault = map[string]attr.Value{
	"subscription":             types.StringNull(),
	"remote_endpoint_protocol": types.StringValue("default"),
	"remote_endpoint_hostname": types.StringNull(),
	"remote_endpoint_port":     types.Int64Null(),
	"remote_endpoint_uri":      types.StringNull(),
	"remote_mqqm":              types.StringNull(),
	"remote_mq_manager":        types.StringNull(),
	"remote_tibco_ems":         types.StringNull(),
	"remote_web_sphere_jms":    types.StringNull(),
}
var DmWSEndpointSubscriptionRemoteRuleDataSourceSchema = DataSourceSchema.NestedAttributeObject{
	Attributes: map[string]DataSourceSchema.Attribute{
		"subscription": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Subscription", "subscription", "").String,
			Computed:            true,
		},
		"remote_endpoint_protocol": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Remote Endpoint Protocol", "remote-endpoint-protocol", "").AddStringEnum("default", "http", "https", "dpmq", "mq", "idgmq", "dptibems", "tibems", "dpwasjms").AddDefaultValue("default").String,
			Computed:            true,
		},
		"remote_endpoint_hostname": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Remote Endpoint Host", "remote-endpoint-hostname", "").String,
			Computed:            true,
		},
		"remote_endpoint_port": DataSourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Remote Endpoint Port", "remote-endpoint-port", "").String,
			Computed:            true,
		},
		"remote_endpoint_uri": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Remote Endpoint URI", "remote-endpoint-uri", "").String,
			Computed:            true,
		},
		"remote_mqqm": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Remote IBM MQ Queue Manager", "remote-mq-qm", "").String,
			Computed:            true,
		},
		"remote_mq_manager": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Remote IBM MQ v9+ Queue Manager", "remote-idg-mq-qm", "").String,
			Computed:            true,
		},
		"remote_tibco_ems": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Remote TIBCO EMS", "remote-tibems-server", "tibcoemsserver").String,
			Computed:            true,
		},
		"remote_web_sphere_jms": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Remote WebSphere JMS", "remote-wasjms-server", "webspherejmsserver").String,
			Computed:            true,
		},
	},
}
var DmWSEndpointSubscriptionRemoteRuleResourceSchema = ResourceSchema.NestedAttributeObject{
	Attributes: map[string]ResourceSchema.Attribute{
		"subscription": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Subscription", "subscription", "").String,
			Optional:            true,
		},
		"remote_endpoint_protocol": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Remote Endpoint Protocol", "remote-endpoint-protocol", "").AddStringEnum("default", "http", "https", "dpmq", "mq", "idgmq", "dptibems", "tibems", "dpwasjms").AddDefaultValue("default").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("default", "http", "https", "dpmq", "mq", "idgmq", "dptibems", "tibems", "dpwasjms"),
			},
			Default: stringdefault.StaticString("default"),
		},
		"remote_endpoint_hostname": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Remote Endpoint Host", "remote-endpoint-hostname", "").String,
			Optional:            true,
		},
		"remote_endpoint_port": ResourceSchema.Int64Attribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Remote Endpoint Port", "remote-endpoint-port", "").String,
			Optional:            true,
		},
		"remote_endpoint_uri": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Remote Endpoint URI", "remote-endpoint-uri", "").String,
			Optional:            true,
		},
		"remote_mqqm": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Remote IBM MQ Queue Manager", "remote-mq-qm", "").String,
			Optional:            true,
		},
		"remote_mq_manager": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Remote IBM MQ v9+ Queue Manager", "remote-idg-mq-qm", "").String,
			Optional:            true,
		},
		"remote_tibco_ems": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Remote TIBCO EMS", "remote-tibems-server", "tibcoemsserver").String,
			Optional:            true,
		},
		"remote_web_sphere_jms": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Remote WebSphere JMS", "remote-wasjms-server", "webspherejmsserver").String,
			Optional:            true,
		},
	},
}

func (data DmWSEndpointSubscriptionRemoteRule) IsNull() bool {
	if !data.Subscription.IsNull() {
		return false
	}
	if !data.RemoteEndpointProtocol.IsNull() {
		return false
	}
	if !data.RemoteEndpointHostname.IsNull() {
		return false
	}
	if !data.RemoteEndpointPort.IsNull() {
		return false
	}
	if !data.RemoteEndpointUri.IsNull() {
		return false
	}
	if !data.RemoteMqqm.IsNull() {
		return false
	}
	if !data.RemoteMqManager.IsNull() {
		return false
	}
	if !data.RemoteTibcoEms.IsNull() {
		return false
	}
	if !data.RemoteWebSphereJms.IsNull() {
		return false
	}
	return true
}

func (data DmWSEndpointSubscriptionRemoteRule) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	if !data.Subscription.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Subscription`, data.Subscription.ValueString())
	}
	if !data.RemoteEndpointProtocol.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`RemoteEndpointProtocol`, data.RemoteEndpointProtocol.ValueString())
	}
	if !data.RemoteEndpointHostname.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`RemoteEndpointHostname`, data.RemoteEndpointHostname.ValueString())
	}
	if !data.RemoteEndpointPort.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`RemoteEndpointPort`, data.RemoteEndpointPort.ValueInt64())
	}
	if !data.RemoteEndpointUri.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`RemoteEndpointURI`, data.RemoteEndpointUri.ValueString())
	}
	if !data.RemoteMqqm.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`RemoteMQQM`, data.RemoteMqqm.ValueString())
	}
	if !data.RemoteMqManager.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`RemoteMQManager`, data.RemoteMqManager.ValueString())
	}
	if !data.RemoteTibcoEms.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`RemoteTibcoEMS`, data.RemoteTibcoEms.ValueString())
	}
	if !data.RemoteWebSphereJms.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`RemoteWebSphereJMS`, data.RemoteWebSphereJms.ValueString())
	}
	return body
}

func (data *DmWSEndpointSubscriptionRemoteRule) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Subscription`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Subscription = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Subscription = types.StringNull()
	}
	if value := res.Get(pathRoot + `RemoteEndpointProtocol`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.RemoteEndpointProtocol = tfutils.ParseStringFromGJSON(value)
	} else {
		data.RemoteEndpointProtocol = types.StringValue("default")
	}
	if value := res.Get(pathRoot + `RemoteEndpointHostname`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.RemoteEndpointHostname = tfutils.ParseStringFromGJSON(value)
	} else {
		data.RemoteEndpointHostname = types.StringNull()
	}
	if value := res.Get(pathRoot + `RemoteEndpointPort`); value.Exists() {
		data.RemoteEndpointPort = types.Int64Value(value.Int())
	} else {
		data.RemoteEndpointPort = types.Int64Null()
	}
	if value := res.Get(pathRoot + `RemoteEndpointURI`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.RemoteEndpointUri = tfutils.ParseStringFromGJSON(value)
	} else {
		data.RemoteEndpointUri = types.StringNull()
	}
	if value := res.Get(pathRoot + `RemoteMQQM`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.RemoteMqqm = tfutils.ParseStringFromGJSON(value)
	} else {
		data.RemoteMqqm = types.StringNull()
	}
	if value := res.Get(pathRoot + `RemoteMQManager`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.RemoteMqManager = tfutils.ParseStringFromGJSON(value)
	} else {
		data.RemoteMqManager = types.StringNull()
	}
	if value := res.Get(pathRoot + `RemoteTibcoEMS`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.RemoteTibcoEms = tfutils.ParseStringFromGJSON(value)
	} else {
		data.RemoteTibcoEms = types.StringNull()
	}
	if value := res.Get(pathRoot + `RemoteWebSphereJMS`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.RemoteWebSphereJms = tfutils.ParseStringFromGJSON(value)
	} else {
		data.RemoteWebSphereJms = types.StringNull()
	}
}

func (data *DmWSEndpointSubscriptionRemoteRule) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Subscription`); value.Exists() && !data.Subscription.IsNull() {
		data.Subscription = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Subscription = types.StringNull()
	}
	if value := res.Get(pathRoot + `RemoteEndpointProtocol`); value.Exists() && !data.RemoteEndpointProtocol.IsNull() {
		data.RemoteEndpointProtocol = tfutils.ParseStringFromGJSON(value)
	} else if data.RemoteEndpointProtocol.ValueString() != "default" {
		data.RemoteEndpointProtocol = types.StringNull()
	}
	if value := res.Get(pathRoot + `RemoteEndpointHostname`); value.Exists() && !data.RemoteEndpointHostname.IsNull() {
		data.RemoteEndpointHostname = tfutils.ParseStringFromGJSON(value)
	} else {
		data.RemoteEndpointHostname = types.StringNull()
	}
	if value := res.Get(pathRoot + `RemoteEndpointPort`); value.Exists() && !data.RemoteEndpointPort.IsNull() {
		data.RemoteEndpointPort = types.Int64Value(value.Int())
	} else {
		data.RemoteEndpointPort = types.Int64Null()
	}
	if value := res.Get(pathRoot + `RemoteEndpointURI`); value.Exists() && !data.RemoteEndpointUri.IsNull() {
		data.RemoteEndpointUri = tfutils.ParseStringFromGJSON(value)
	} else {
		data.RemoteEndpointUri = types.StringNull()
	}
	if value := res.Get(pathRoot + `RemoteMQQM`); value.Exists() && !data.RemoteMqqm.IsNull() {
		data.RemoteMqqm = tfutils.ParseStringFromGJSON(value)
	} else {
		data.RemoteMqqm = types.StringNull()
	}
	if value := res.Get(pathRoot + `RemoteMQManager`); value.Exists() && !data.RemoteMqManager.IsNull() {
		data.RemoteMqManager = tfutils.ParseStringFromGJSON(value)
	} else {
		data.RemoteMqManager = types.StringNull()
	}
	if value := res.Get(pathRoot + `RemoteTibcoEMS`); value.Exists() && !data.RemoteTibcoEms.IsNull() {
		data.RemoteTibcoEms = tfutils.ParseStringFromGJSON(value)
	} else {
		data.RemoteTibcoEms = types.StringNull()
	}
	if value := res.Get(pathRoot + `RemoteWebSphereJMS`); value.Exists() && !data.RemoteWebSphereJms.IsNull() {
		data.RemoteWebSphereJms = tfutils.ParseStringFromGJSON(value)
	} else {
		data.RemoteWebSphereJms = types.StringNull()
	}
}

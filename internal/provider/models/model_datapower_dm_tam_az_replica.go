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

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	DataSourceSchema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	ResourceSchema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type DmTAMAZReplica struct {
	TamAzReplica       types.String `tfsdk:"tam_az_replica"`
	TamAzReplicaPort   types.Int64  `tfsdk:"tam_az_replica_port"`
	TamAzReplicaWeight types.Int64  `tfsdk:"tam_az_replica_weight"`
}

var DmTAMAZReplicaObjectType = map[string]attr.Type{
	"tam_az_replica":        types.StringType,
	"tam_az_replica_port":   types.Int64Type,
	"tam_az_replica_weight": types.Int64Type,
}
var DmTAMAZReplicaObjectDefault = map[string]attr.Value{
	"tam_az_replica":        types.StringNull(),
	"tam_az_replica_port":   types.Int64Value(7136),
	"tam_az_replica_weight": types.Int64Value(10),
}

func GetDmTAMAZReplicaDataSourceSchema() DataSourceSchema.NestedAttributeObject {
	var DmTAMAZReplicaDataSourceSchema = DataSourceSchema.NestedAttributeObject{
		Attributes: map[string]DataSourceSchema.Attribute{
			"tam_az_replica": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the TCP host name of the authorization server replica.", "host", "").String,
				Computed:            true,
			},
			"tam_az_replica_port": DataSourceSchema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the listening port on the authorization server replica. The default value is 7136.", "port", "").AddDefaultValue("7136").String,
				Computed:            true,
			},
			"tam_az_replica_weight": DataSourceSchema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the weight of the authorization server replica. The greater the weight, the higher the preference. Enter a value in the range 1 - 10. The default value is 10.", "weight", "").AddIntegerRange(1, 10).AddDefaultValue("10").String,
				Computed:            true,
			},
		},
	}
	return DmTAMAZReplicaDataSourceSchema
}
func GetDmTAMAZReplicaResourceSchema() ResourceSchema.NestedAttributeObject {
	var DmTAMAZReplicaResourceSchema = ResourceSchema.NestedAttributeObject{
		Attributes: map[string]ResourceSchema.Attribute{
			"tam_az_replica": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the TCP host name of the authorization server replica.", "host", "").String,
				Required:            true,
			},
			"tam_az_replica_port": ResourceSchema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the listening port on the authorization server replica. The default value is 7136.", "port", "").AddDefaultValue("7136").String,
				Computed:            true,
				Optional:            true,
				Default:             int64default.StaticInt64(7136),
			},
			"tam_az_replica_weight": ResourceSchema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the weight of the authorization server replica. The greater the weight, the higher the preference. Enter a value in the range 1 - 10. The default value is 10.", "weight", "").AddIntegerRange(1, 10).AddDefaultValue("10").String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 10),
				},
				Default: int64default.StaticInt64(10),
			},
		},
	}
	return DmTAMAZReplicaResourceSchema
}

func (data DmTAMAZReplica) IsNull() bool {
	if !data.TamAzReplica.IsNull() {
		return false
	}
	if !data.TamAzReplicaPort.IsNull() {
		return false
	}
	if !data.TamAzReplicaWeight.IsNull() {
		return false
	}
	return true
}

func (data DmTAMAZReplica) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""

	if !data.TamAzReplica.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`TAMAZReplica`, data.TamAzReplica.ValueString())
	}
	if !data.TamAzReplicaPort.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`TAMAZReplicaPort`, data.TamAzReplicaPort.ValueInt64())
	}
	if !data.TamAzReplicaWeight.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`TAMAZReplicaWeight`, data.TamAzReplicaWeight.ValueInt64())
	}
	return body
}

func (data *DmTAMAZReplica) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `TAMAZReplica`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.TamAzReplica = tfutils.ParseStringFromGJSON(value)
	} else {
		data.TamAzReplica = types.StringNull()
	}
	if value := res.Get(pathRoot + `TAMAZReplicaPort`); value.Exists() {
		data.TamAzReplicaPort = types.Int64Value(value.Int())
	} else {
		data.TamAzReplicaPort = types.Int64Value(7136)
	}
	if value := res.Get(pathRoot + `TAMAZReplicaWeight`); value.Exists() {
		data.TamAzReplicaWeight = types.Int64Value(value.Int())
	} else {
		data.TamAzReplicaWeight = types.Int64Value(10)
	}
}

func (data *DmTAMAZReplica) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `TAMAZReplica`); value.Exists() && !data.TamAzReplica.IsNull() {
		data.TamAzReplica = tfutils.ParseStringFromGJSON(value)
	} else {
		data.TamAzReplica = types.StringNull()
	}
	if value := res.Get(pathRoot + `TAMAZReplicaPort`); value.Exists() && !data.TamAzReplicaPort.IsNull() {
		data.TamAzReplicaPort = types.Int64Value(value.Int())
	} else if data.TamAzReplicaPort.ValueInt64() != 7136 {
		data.TamAzReplicaPort = types.Int64Null()
	}
	if value := res.Get(pathRoot + `TAMAZReplicaWeight`); value.Exists() && !data.TamAzReplicaWeight.IsNull() {
		data.TamAzReplicaWeight = types.Int64Value(value.Int())
	} else if data.TamAzReplicaWeight.ValueInt64() != 10 {
		data.TamAzReplicaWeight = types.Int64Null()
	}
}

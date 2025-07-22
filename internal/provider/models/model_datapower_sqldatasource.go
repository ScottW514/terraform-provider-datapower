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
	"net/url"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type SQLDataSource struct {
	Id                         types.String `tfsdk:"id"`
	AppDomain                  types.String `tfsdk:"app_domain"`
	UserSummary                types.String `tfsdk:"user_summary"`
	Database                   types.String `tfsdk:"database"`
	Username                   types.String `tfsdk:"username"`
	PasswordAlias              types.String `tfsdk:"password_alias"`
	DataSourceId               types.String `tfsdk:"data_source_id"`
	DataSourceHost             types.String `tfsdk:"data_source_host"`
	DataSourcePort             types.Int64  `tfsdk:"data_source_port"`
	LimitReturnedData          types.Bool   `tfsdk:"limit_returned_data"`
	LimitReturnedDataSize      types.Int64  `tfsdk:"limit_returned_data_size"`
	SqlDataSourceConfigNvPairs types.List   `tfsdk:"sql_data_source_config_nv_pairs"`
	MaxConnection              types.Int64  `tfsdk:"max_connection"`
	OracleDataSourceType       types.String `tfsdk:"oracle_data_source_type"`
	OracleObjects              types.Bool   `tfsdk:"oracle_objects"`
	ConnectTimeout             types.Int64  `tfsdk:"connect_timeout"`
	QueryTimeout               types.Int64  `tfsdk:"query_timeout"`
	IdleTimeout                types.Int64  `tfsdk:"idle_timeout"`
	LoadBalancing              types.Bool   `tfsdk:"load_balancing"`
	EncryptionMethodMssql      types.String `tfsdk:"encryption_method_mssql"`
	EncryptionMethodOracle     types.String `tfsdk:"encryption_method_oracle"`
	EncryptionMethodDb2        types.String `tfsdk:"encryption_method_db2"`
	TruststoreRef              types.String `tfsdk:"truststore_ref"`
	ValidateServerCertificate  types.String `tfsdk:"validate_server_certificate"`
	HostNameInCertificate      types.String `tfsdk:"host_name_in_certificate"`
	ValidateHostName           types.Bool   `tfsdk:"validate_host_name"`
	KeystoreRef                types.String `tfsdk:"keystore_ref"`
}

var SQLDataSourceObjectType = map[string]attr.Type{
	"id":                              types.StringType,
	"app_domain":                      types.StringType,
	"user_summary":                    types.StringType,
	"database":                        types.StringType,
	"username":                        types.StringType,
	"password_alias":                  types.StringType,
	"data_source_id":                  types.StringType,
	"data_source_host":                types.StringType,
	"data_source_port":                types.Int64Type,
	"limit_returned_data":             types.BoolType,
	"limit_returned_data_size":        types.Int64Type,
	"sql_data_source_config_nv_pairs": types.ListType{ElemType: types.ObjectType{AttrTypes: DmSQLDataSourceConfigNVPairObjectType}},
	"max_connection":                  types.Int64Type,
	"oracle_data_source_type":         types.StringType,
	"oracle_objects":                  types.BoolType,
	"connect_timeout":                 types.Int64Type,
	"query_timeout":                   types.Int64Type,
	"idle_timeout":                    types.Int64Type,
	"load_balancing":                  types.BoolType,
	"encryption_method_mssql":         types.StringType,
	"encryption_method_oracle":        types.StringType,
	"encryption_method_db2":           types.StringType,
	"truststore_ref":                  types.StringType,
	"validate_server_certificate":     types.StringType,
	"host_name_in_certificate":        types.StringType,
	"validate_host_name":              types.BoolType,
	"keystore_ref":                    types.StringType,
}

func (data SQLDataSource) GetPath() string {
	rest_path := "/mgmt/config/{domain}/SQLDataSource"
	rest_path = strings.ReplaceAll(rest_path, "{name}", url.QueryEscape(data.Id.ValueString()))
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data SQLDataSource) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.Database.IsNull() {
		return false
	}
	if !data.Username.IsNull() {
		return false
	}
	if !data.PasswordAlias.IsNull() {
		return false
	}
	if !data.DataSourceId.IsNull() {
		return false
	}
	if !data.DataSourceHost.IsNull() {
		return false
	}
	if !data.DataSourcePort.IsNull() {
		return false
	}
	if !data.LimitReturnedData.IsNull() {
		return false
	}
	if !data.LimitReturnedDataSize.IsNull() {
		return false
	}
	if !data.SqlDataSourceConfigNvPairs.IsNull() {
		return false
	}
	if !data.MaxConnection.IsNull() {
		return false
	}
	if !data.OracleDataSourceType.IsNull() {
		return false
	}
	if !data.OracleObjects.IsNull() {
		return false
	}
	if !data.ConnectTimeout.IsNull() {
		return false
	}
	if !data.QueryTimeout.IsNull() {
		return false
	}
	if !data.IdleTimeout.IsNull() {
		return false
	}
	if !data.LoadBalancing.IsNull() {
		return false
	}
	if !data.EncryptionMethodMssql.IsNull() {
		return false
	}
	if !data.EncryptionMethodOracle.IsNull() {
		return false
	}
	if !data.EncryptionMethodDb2.IsNull() {
		return false
	}
	if !data.TruststoreRef.IsNull() {
		return false
	}
	if !data.ValidateServerCertificate.IsNull() {
		return false
	}
	if !data.HostNameInCertificate.IsNull() {
		return false
	}
	if !data.ValidateHostName.IsNull() {
		return false
	}
	if !data.KeystoreRef.IsNull() {
		return false
	}
	return true
}

func (data SQLDataSource) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	if !data.Id.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`name`, data.Id.ValueString())
	}
	if !data.UserSummary.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UserSummary`, data.UserSummary.ValueString())
	}
	if !data.Database.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Database`, data.Database.ValueString())
	}
	if !data.Username.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Username`, data.Username.ValueString())
	}
	if !data.PasswordAlias.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PasswordAlias`, data.PasswordAlias.ValueString())
	}
	if !data.DataSourceId.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DataSourceID`, data.DataSourceId.ValueString())
	}
	if !data.DataSourceHost.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DataSourceHost`, data.DataSourceHost.ValueString())
	}
	if !data.DataSourcePort.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DataSourcePort`, data.DataSourcePort.ValueInt64())
	}
	if !data.LimitReturnedData.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LimitReturnedData`, tfutils.StringFromBool(data.LimitReturnedData, false))
	}
	if !data.LimitReturnedDataSize.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LimitReturnedDataSize`, data.LimitReturnedDataSize.ValueInt64())
	}
	if !data.SqlDataSourceConfigNvPairs.IsNull() {
		var values []DmSQLDataSourceConfigNVPair
		data.SqlDataSourceConfigNvPairs.ElementsAs(ctx, &values, false)
		for _, val := range values {
			body, _ = sjson.SetRaw(body, pathRoot+`SQLDataSourceConfigNVPairs`+".-1", val.ToBody(ctx, ""))
		}
	}
	if !data.MaxConnection.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MaxConnection`, data.MaxConnection.ValueInt64())
	}
	if !data.OracleDataSourceType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`OracleDataSourceType`, data.OracleDataSourceType.ValueString())
	}
	if !data.OracleObjects.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`OracleObjects`, tfutils.StringFromBool(data.OracleObjects, false))
	}
	if !data.ConnectTimeout.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ConnectTimeout`, data.ConnectTimeout.ValueInt64())
	}
	if !data.QueryTimeout.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`QueryTimeout`, data.QueryTimeout.ValueInt64())
	}
	if !data.IdleTimeout.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`IdleTimeout`, data.IdleTimeout.ValueInt64())
	}
	if !data.LoadBalancing.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LoadBalancing`, tfutils.StringFromBool(data.LoadBalancing, false))
	}
	if !data.EncryptionMethodMssql.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EncryptionMethodMSSQL`, data.EncryptionMethodMssql.ValueString())
	}
	if !data.EncryptionMethodOracle.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EncryptionMethodOracle`, data.EncryptionMethodOracle.ValueString())
	}
	if !data.EncryptionMethodDb2.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EncryptionMethodDB2`, data.EncryptionMethodDb2.ValueString())
	}
	if !data.TruststoreRef.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`TruststoreRef`, data.TruststoreRef.ValueString())
	}
	if !data.ValidateServerCertificate.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ValidateServerCertificate`, data.ValidateServerCertificate.ValueString())
	}
	if !data.HostNameInCertificate.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`HostNameInCertificate`, data.HostNameInCertificate.ValueString())
	}
	if !data.ValidateHostName.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ValidateHostName`, tfutils.StringFromBool(data.ValidateHostName, false))
	}
	if !data.KeystoreRef.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`KeystoreRef`, data.KeystoreRef.ValueString())
	}
	return body
}

func (data *SQLDataSource) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `name`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Id = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Id = types.StringNull()
	}
	if value := res.Get(pathRoot + `UserSummary`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.UserSummary = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UserSummary = types.StringNull()
	}
	if value := res.Get(pathRoot + `Database`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Database = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Database = types.StringNull()
	}
	if value := res.Get(pathRoot + `Username`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Username = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Username = types.StringNull()
	}
	if value := res.Get(pathRoot + `PasswordAlias`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PasswordAlias = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PasswordAlias = types.StringNull()
	}
	if value := res.Get(pathRoot + `DataSourceID`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.DataSourceId = tfutils.ParseStringFromGJSON(value)
	} else {
		data.DataSourceId = types.StringNull()
	}
	if value := res.Get(pathRoot + `DataSourceHost`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.DataSourceHost = tfutils.ParseStringFromGJSON(value)
	} else {
		data.DataSourceHost = types.StringNull()
	}
	if value := res.Get(pathRoot + `DataSourcePort`); value.Exists() {
		data.DataSourcePort = types.Int64Value(value.Int())
	} else {
		data.DataSourcePort = types.Int64Null()
	}
	if value := res.Get(pathRoot + `LimitReturnedData`); value.Exists() {
		data.LimitReturnedData = tfutils.BoolFromString(value.String())
	} else {
		data.LimitReturnedData = types.BoolNull()
	}
	if value := res.Get(pathRoot + `LimitReturnedDataSize`); value.Exists() {
		data.LimitReturnedDataSize = types.Int64Value(value.Int())
	} else {
		data.LimitReturnedDataSize = types.Int64Value(128)
	}
	if value := res.Get(pathRoot + `SQLDataSourceConfigNVPairs`); value.Exists() {
		l := []DmSQLDataSourceConfigNVPair{}
		if value := res.Get(`SQLDataSourceConfigNVPairs`); value.Exists() {
			for _, v := range value.Array() {
				item := DmSQLDataSourceConfigNVPair{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.SqlDataSourceConfigNvPairs, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmSQLDataSourceConfigNVPairObjectType}, l)
		} else {
			data.SqlDataSourceConfigNvPairs = types.ListNull(types.ObjectType{AttrTypes: DmSQLDataSourceConfigNVPairObjectType})
		}
	} else {
		data.SqlDataSourceConfigNvPairs = types.ListNull(types.ObjectType{AttrTypes: DmSQLDataSourceConfigNVPairObjectType})
	}
	if value := res.Get(pathRoot + `MaxConnection`); value.Exists() {
		data.MaxConnection = types.Int64Value(value.Int())
	} else {
		data.MaxConnection = types.Int64Value(10)
	}
	if value := res.Get(pathRoot + `OracleDataSourceType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.OracleDataSourceType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.OracleDataSourceType = types.StringValue("SID")
	}
	if value := res.Get(pathRoot + `OracleObjects`); value.Exists() {
		data.OracleObjects = tfutils.BoolFromString(value.String())
	} else {
		data.OracleObjects = types.BoolNull()
	}
	if value := res.Get(pathRoot + `ConnectTimeout`); value.Exists() {
		data.ConnectTimeout = types.Int64Value(value.Int())
	} else {
		data.ConnectTimeout = types.Int64Value(15)
	}
	if value := res.Get(pathRoot + `QueryTimeout`); value.Exists() {
		data.QueryTimeout = types.Int64Value(value.Int())
	} else {
		data.QueryTimeout = types.Int64Null()
	}
	if value := res.Get(pathRoot + `IdleTimeout`); value.Exists() {
		data.IdleTimeout = types.Int64Value(value.Int())
	} else {
		data.IdleTimeout = types.Int64Value(180)
	}
	if value := res.Get(pathRoot + `LoadBalancing`); value.Exists() {
		data.LoadBalancing = tfutils.BoolFromString(value.String())
	} else {
		data.LoadBalancing = types.BoolNull()
	}
	if value := res.Get(pathRoot + `EncryptionMethodMSSQL`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.EncryptionMethodMssql = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EncryptionMethodMssql = types.StringValue("NoEncryption")
	}
	if value := res.Get(pathRoot + `EncryptionMethodOracle`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.EncryptionMethodOracle = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EncryptionMethodOracle = types.StringValue("NoEncryption")
	}
	if value := res.Get(pathRoot + `EncryptionMethodDB2`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.EncryptionMethodDb2 = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EncryptionMethodDb2 = types.StringValue("NoEncryption")
	}
	if value := res.Get(pathRoot + `TruststoreRef`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.TruststoreRef = tfutils.ParseStringFromGJSON(value)
	} else {
		data.TruststoreRef = types.StringNull()
	}
	if value := res.Get(pathRoot + `ValidateServerCertificate`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ValidateServerCertificate = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ValidateServerCertificate = types.StringValue("Enabled")
	}
	if value := res.Get(pathRoot + `HostNameInCertificate`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.HostNameInCertificate = tfutils.ParseStringFromGJSON(value)
	} else {
		data.HostNameInCertificate = types.StringNull()
	}
	if value := res.Get(pathRoot + `ValidateHostName`); value.Exists() {
		data.ValidateHostName = tfutils.BoolFromString(value.String())
	} else {
		data.ValidateHostName = types.BoolNull()
	}
	if value := res.Get(pathRoot + `KeystoreRef`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.KeystoreRef = tfutils.ParseStringFromGJSON(value)
	} else {
		data.KeystoreRef = types.StringNull()
	}
}

func (data *SQLDataSource) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `name`); value.Exists() && !data.Id.IsNull() {
		data.Id = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Id = types.StringNull()
	}
	if value := res.Get(pathRoot + `UserSummary`); value.Exists() && !data.UserSummary.IsNull() {
		data.UserSummary = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UserSummary = types.StringNull()
	}
	if value := res.Get(pathRoot + `Database`); value.Exists() && !data.Database.IsNull() {
		data.Database = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Database = types.StringNull()
	}
	if value := res.Get(pathRoot + `Username`); value.Exists() && !data.Username.IsNull() {
		data.Username = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Username = types.StringNull()
	}
	if value := res.Get(pathRoot + `PasswordAlias`); value.Exists() && !data.PasswordAlias.IsNull() {
		data.PasswordAlias = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PasswordAlias = types.StringNull()
	}
	if value := res.Get(pathRoot + `DataSourceID`); value.Exists() && !data.DataSourceId.IsNull() {
		data.DataSourceId = tfutils.ParseStringFromGJSON(value)
	} else {
		data.DataSourceId = types.StringNull()
	}
	if value := res.Get(pathRoot + `DataSourceHost`); value.Exists() && !data.DataSourceHost.IsNull() {
		data.DataSourceHost = tfutils.ParseStringFromGJSON(value)
	} else {
		data.DataSourceHost = types.StringNull()
	}
	if value := res.Get(pathRoot + `DataSourcePort`); value.Exists() && !data.DataSourcePort.IsNull() {
		data.DataSourcePort = types.Int64Value(value.Int())
	} else {
		data.DataSourcePort = types.Int64Null()
	}
	if value := res.Get(pathRoot + `LimitReturnedData`); value.Exists() && !data.LimitReturnedData.IsNull() {
		data.LimitReturnedData = tfutils.BoolFromString(value.String())
	} else if data.LimitReturnedData.ValueBool() {
		data.LimitReturnedData = types.BoolNull()
	}
	if value := res.Get(pathRoot + `LimitReturnedDataSize`); value.Exists() && !data.LimitReturnedDataSize.IsNull() {
		data.LimitReturnedDataSize = types.Int64Value(value.Int())
	} else if data.LimitReturnedDataSize.ValueInt64() != 128 {
		data.LimitReturnedDataSize = types.Int64Null()
	}
	if value := res.Get(pathRoot + `SQLDataSourceConfigNVPairs`); value.Exists() && !data.SqlDataSourceConfigNvPairs.IsNull() {
		l := []DmSQLDataSourceConfigNVPair{}
		for _, v := range value.Array() {
			item := DmSQLDataSourceConfigNVPair{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
		if len(l) > 0 {
			data.SqlDataSourceConfigNvPairs, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmSQLDataSourceConfigNVPairObjectType}, l)
		} else {
			data.SqlDataSourceConfigNvPairs = types.ListNull(types.ObjectType{AttrTypes: DmSQLDataSourceConfigNVPairObjectType})
		}
	} else {
		data.SqlDataSourceConfigNvPairs = types.ListNull(types.ObjectType{AttrTypes: DmSQLDataSourceConfigNVPairObjectType})
	}
	if value := res.Get(pathRoot + `MaxConnection`); value.Exists() && !data.MaxConnection.IsNull() {
		data.MaxConnection = types.Int64Value(value.Int())
	} else if data.MaxConnection.ValueInt64() != 10 {
		data.MaxConnection = types.Int64Null()
	}
	if value := res.Get(pathRoot + `OracleDataSourceType`); value.Exists() && !data.OracleDataSourceType.IsNull() {
		data.OracleDataSourceType = tfutils.ParseStringFromGJSON(value)
	} else if data.OracleDataSourceType.ValueString() != "SID" {
		data.OracleDataSourceType = types.StringNull()
	}
	if value := res.Get(pathRoot + `OracleObjects`); value.Exists() && !data.OracleObjects.IsNull() {
		data.OracleObjects = tfutils.BoolFromString(value.String())
	} else if data.OracleObjects.ValueBool() {
		data.OracleObjects = types.BoolNull()
	}
	if value := res.Get(pathRoot + `ConnectTimeout`); value.Exists() && !data.ConnectTimeout.IsNull() {
		data.ConnectTimeout = types.Int64Value(value.Int())
	} else if data.ConnectTimeout.ValueInt64() != 15 {
		data.ConnectTimeout = types.Int64Null()
	}
	if value := res.Get(pathRoot + `QueryTimeout`); value.Exists() && !data.QueryTimeout.IsNull() {
		data.QueryTimeout = types.Int64Value(value.Int())
	} else {
		data.QueryTimeout = types.Int64Null()
	}
	if value := res.Get(pathRoot + `IdleTimeout`); value.Exists() && !data.IdleTimeout.IsNull() {
		data.IdleTimeout = types.Int64Value(value.Int())
	} else if data.IdleTimeout.ValueInt64() != 180 {
		data.IdleTimeout = types.Int64Null()
	}
	if value := res.Get(pathRoot + `LoadBalancing`); value.Exists() && !data.LoadBalancing.IsNull() {
		data.LoadBalancing = tfutils.BoolFromString(value.String())
	} else if data.LoadBalancing.ValueBool() {
		data.LoadBalancing = types.BoolNull()
	}
	if value := res.Get(pathRoot + `EncryptionMethodMSSQL`); value.Exists() && !data.EncryptionMethodMssql.IsNull() {
		data.EncryptionMethodMssql = tfutils.ParseStringFromGJSON(value)
	} else if data.EncryptionMethodMssql.ValueString() != "NoEncryption" {
		data.EncryptionMethodMssql = types.StringNull()
	}
	if value := res.Get(pathRoot + `EncryptionMethodOracle`); value.Exists() && !data.EncryptionMethodOracle.IsNull() {
		data.EncryptionMethodOracle = tfutils.ParseStringFromGJSON(value)
	} else if data.EncryptionMethodOracle.ValueString() != "NoEncryption" {
		data.EncryptionMethodOracle = types.StringNull()
	}
	if value := res.Get(pathRoot + `EncryptionMethodDB2`); value.Exists() && !data.EncryptionMethodDb2.IsNull() {
		data.EncryptionMethodDb2 = tfutils.ParseStringFromGJSON(value)
	} else if data.EncryptionMethodDb2.ValueString() != "NoEncryption" {
		data.EncryptionMethodDb2 = types.StringNull()
	}
	if value := res.Get(pathRoot + `TruststoreRef`); value.Exists() && !data.TruststoreRef.IsNull() {
		data.TruststoreRef = tfutils.ParseStringFromGJSON(value)
	} else {
		data.TruststoreRef = types.StringNull()
	}
	if value := res.Get(pathRoot + `ValidateServerCertificate`); value.Exists() && !data.ValidateServerCertificate.IsNull() {
		data.ValidateServerCertificate = tfutils.ParseStringFromGJSON(value)
	} else if data.ValidateServerCertificate.ValueString() != "Enabled" {
		data.ValidateServerCertificate = types.StringNull()
	}
	if value := res.Get(pathRoot + `HostNameInCertificate`); value.Exists() && !data.HostNameInCertificate.IsNull() {
		data.HostNameInCertificate = tfutils.ParseStringFromGJSON(value)
	} else {
		data.HostNameInCertificate = types.StringNull()
	}
	if value := res.Get(pathRoot + `ValidateHostName`); value.Exists() && !data.ValidateHostName.IsNull() {
		data.ValidateHostName = tfutils.BoolFromString(value.String())
	} else if !data.ValidateHostName.ValueBool() {
		data.ValidateHostName = types.BoolNull()
	}
	if value := res.Get(pathRoot + `KeystoreRef`); value.Exists() && !data.KeystoreRef.IsNull() {
		data.KeystoreRef = tfutils.ParseStringFromGJSON(value)
	} else {
		data.KeystoreRef = types.StringNull()
	}
}

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
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type ODRConnectorGroup struct {
	Id                     types.String                `tfsdk:"id"`
	UserSummary            types.String                `tfsdk:"user_summary"`
	OdrGroupConnectors     types.List                  `tfsdk:"odr_group_connectors"`
	MaxRetryInterval       types.Int64                 `tfsdk:"max_retry_interval"`
	XmlManager             types.String                `tfsdk:"xml_manager"`
	OdrConnGroupProperties types.List                  `tfsdk:"odr_conn_group_properties"`
	SslClientConfigType    types.String                `tfsdk:"ssl_client_config_type"`
	SslClient              types.String                `tfsdk:"ssl_client"`
	DependencyActions      []*actions.DependencyAction `tfsdk:"dependency_actions"`
}

var ODRConnectorGroupObjectType = map[string]attr.Type{
	"id":                        types.StringType,
	"user_summary":              types.StringType,
	"odr_group_connectors":      types.ListType{ElemType: types.ObjectType{AttrTypes: DmODRConnectorObjectType}},
	"max_retry_interval":        types.Int64Type,
	"xml_manager":               types.StringType,
	"odr_conn_group_properties": types.ListType{ElemType: types.ObjectType{AttrTypes: DmODRConnPropertyObjectType}},
	"ssl_client_config_type":    types.StringType,
	"ssl_client":                types.StringType,
	"dependency_actions":        actions.ActionsListType,
}

func (data ODRConnectorGroup) GetPath() string {
	rest_path := "/mgmt/config/default/ODRConnectorGroup"
	rest_path = strings.ReplaceAll(rest_path, "{name}", url.QueryEscape(data.Id.ValueString()))
	return rest_path
}

func (data ODRConnectorGroup) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.OdrGroupConnectors.IsNull() {
		return false
	}
	if !data.MaxRetryInterval.IsNull() {
		return false
	}
	if !data.XmlManager.IsNull() {
		return false
	}
	if !data.OdrConnGroupProperties.IsNull() {
		return false
	}
	if !data.SslClientConfigType.IsNull() {
		return false
	}
	if !data.SslClient.IsNull() {
		return false
	}
	return true
}

func (data ODRConnectorGroup) ToBody(ctx context.Context, pathRoot string) string {
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
	if !data.OdrGroupConnectors.IsNull() {
		var values []DmODRConnector
		data.OdrGroupConnectors.ElementsAs(ctx, &values, false)
		for _, val := range values {
			body, _ = sjson.SetRaw(body, pathRoot+`ODRGroupConnectors`+".-1", val.ToBody(ctx, ""))
		}
	}
	if !data.MaxRetryInterval.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MaxRetryInterval`, data.MaxRetryInterval.ValueInt64())
	}
	if !data.XmlManager.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`XMLManager`, data.XmlManager.ValueString())
	}
	if !data.OdrConnGroupProperties.IsNull() {
		var values []DmODRConnProperty
		data.OdrConnGroupProperties.ElementsAs(ctx, &values, false)
		for _, val := range values {
			body, _ = sjson.SetRaw(body, pathRoot+`ODRConnGroupProperties`+".-1", val.ToBody(ctx, ""))
		}
	}
	if !data.SslClientConfigType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SSLClientConfigType`, data.SslClientConfigType.ValueString())
	}
	if !data.SslClient.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SSLClient`, data.SslClient.ValueString())
	}
	return body
}

func (data *ODRConnectorGroup) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `ODRGroupConnectors`); value.Exists() {
		l := []DmODRConnector{}
		if value := res.Get(`ODRGroupConnectors`); value.Exists() {
			for _, v := range value.Array() {
				item := DmODRConnector{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.OdrGroupConnectors, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmODRConnectorObjectType}, l)
		} else {
			data.OdrGroupConnectors = types.ListNull(types.ObjectType{AttrTypes: DmODRConnectorObjectType})
		}
	} else {
		data.OdrGroupConnectors = types.ListNull(types.ObjectType{AttrTypes: DmODRConnectorObjectType})
	}
	if value := res.Get(pathRoot + `MaxRetryInterval`); value.Exists() {
		data.MaxRetryInterval = types.Int64Value(value.Int())
	} else {
		data.MaxRetryInterval = types.Int64Value(60)
	}
	if value := res.Get(pathRoot + `XMLManager`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.XmlManager = tfutils.ParseStringFromGJSON(value)
	} else {
		data.XmlManager = types.StringValue("default")
	}
	if value := res.Get(pathRoot + `ODRConnGroupProperties`); value.Exists() {
		l := []DmODRConnProperty{}
		if value := res.Get(`ODRConnGroupProperties`); value.Exists() {
			for _, v := range value.Array() {
				item := DmODRConnProperty{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.OdrConnGroupProperties, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmODRConnPropertyObjectType}, l)
		} else {
			data.OdrConnGroupProperties = types.ListNull(types.ObjectType{AttrTypes: DmODRConnPropertyObjectType})
		}
	} else {
		data.OdrConnGroupProperties = types.ListNull(types.ObjectType{AttrTypes: DmODRConnPropertyObjectType})
	}
	if value := res.Get(pathRoot + `SSLClientConfigType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SslClientConfigType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SslClientConfigType = types.StringValue("client")
	}
	if value := res.Get(pathRoot + `SSLClient`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SslClient = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SslClient = types.StringNull()
	}
}

func (data *ODRConnectorGroup) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `ODRGroupConnectors`); value.Exists() && !data.OdrGroupConnectors.IsNull() {
		l := []DmODRConnector{}
		for _, v := range value.Array() {
			item := DmODRConnector{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
		if len(l) > 0 {
			data.OdrGroupConnectors, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmODRConnectorObjectType}, l)
		} else {
			data.OdrGroupConnectors = types.ListNull(types.ObjectType{AttrTypes: DmODRConnectorObjectType})
		}
	} else {
		data.OdrGroupConnectors = types.ListNull(types.ObjectType{AttrTypes: DmODRConnectorObjectType})
	}
	if value := res.Get(pathRoot + `MaxRetryInterval`); value.Exists() && !data.MaxRetryInterval.IsNull() {
		data.MaxRetryInterval = types.Int64Value(value.Int())
	} else if data.MaxRetryInterval.ValueInt64() != 60 {
		data.MaxRetryInterval = types.Int64Null()
	}
	if value := res.Get(pathRoot + `XMLManager`); value.Exists() && !data.XmlManager.IsNull() {
		data.XmlManager = tfutils.ParseStringFromGJSON(value)
	} else if data.XmlManager.ValueString() != "default" {
		data.XmlManager = types.StringNull()
	}
	if value := res.Get(pathRoot + `ODRConnGroupProperties`); value.Exists() && !data.OdrConnGroupProperties.IsNull() {
		l := []DmODRConnProperty{}
		for _, v := range value.Array() {
			item := DmODRConnProperty{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
		if len(l) > 0 {
			data.OdrConnGroupProperties, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmODRConnPropertyObjectType}, l)
		} else {
			data.OdrConnGroupProperties = types.ListNull(types.ObjectType{AttrTypes: DmODRConnPropertyObjectType})
		}
	} else {
		data.OdrConnGroupProperties = types.ListNull(types.ObjectType{AttrTypes: DmODRConnPropertyObjectType})
	}
	if value := res.Get(pathRoot + `SSLClientConfigType`); value.Exists() && !data.SslClientConfigType.IsNull() {
		data.SslClientConfigType = tfutils.ParseStringFromGJSON(value)
	} else if data.SslClientConfigType.ValueString() != "client" {
		data.SslClientConfigType = types.StringNull()
	}
	if value := res.Get(pathRoot + `SSLClient`); value.Exists() && !data.SslClient.IsNull() {
		data.SslClient = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SslClient = types.StringNull()
	}
}

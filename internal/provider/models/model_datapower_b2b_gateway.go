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
	"github.com/scottw514/terraform-provider-datapower/internal/provider/validators"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type B2BGateway struct {
	Id                               types.String                `tfsdk:"id"`
	AppDomain                        types.String                `tfsdk:"app_domain"`
	UserSummary                      types.String                `tfsdk:"user_summary"`
	Priority                         types.String                `tfsdk:"priority"`
	DocStoreLocation                 types.String                `tfsdk:"doc_store_location"`
	AsFrontProtocol                  types.List                  `tfsdk:"as_front_protocol"`
	As1MdnEmail                      types.String                `tfsdk:"as1_mdn_email"`
	As1MdnSmtpServerConnection       types.String                `tfsdk:"as1_mdn_smtp_server_connection"`
	As2MdnUrl                        types.String                `tfsdk:"as2_mdn_url"`
	As3MdnUrl                        types.String                `tfsdk:"as3_mdn_url"`
	B2bProfiles                      types.List                  `tfsdk:"b2b_profiles"`
	B2bGroups                        types.List                  `tfsdk:"b2b_groups"`
	DocumentRoutingPreprocessorType  types.String                `tfsdk:"document_routing_preprocessor_type"`
	DocumentRoutingPreprocessor      types.String                `tfsdk:"document_routing_preprocessor"`
	DocumentRoutingPreprocessorDebug types.Bool                  `tfsdk:"document_routing_preprocessor_debug"`
	ArchiveMode                      types.String                `tfsdk:"archive_mode"`
	ArchiveLocation                  types.String                `tfsdk:"archive_location"`
	ArchiveFileName                  types.String                `tfsdk:"archive_file_name"`
	ArchiveMinimumSize               types.Int64                 `tfsdk:"archive_minimum_size"`
	ArchiveDocumentAge               types.Int64                 `tfsdk:"archive_document_age"`
	ArchiveMinimumDocuments          types.Int64                 `tfsdk:"archive_minimum_documents"`
	DiskUseCheckInterval             types.Int64                 `tfsdk:"disk_use_check_interval"`
	MaxDocumentDiskUse               types.Int64                 `tfsdk:"max_document_disk_use"`
	ArchiveMonitor                   types.Bool                  `tfsdk:"archive_monitor"`
	ShapingThreshold                 types.Int64                 `tfsdk:"shaping_threshold"`
	ArchiveBackupDocuments           *DmB2BBackupMsgType         `tfsdk:"archive_backup_documents"`
	XpathRoutingPolicies             types.List                  `tfsdk:"xpath_routing_policies"`
	XmlManager                       types.String                `tfsdk:"xml_manager"`
	DebugMode                        types.String                `tfsdk:"debug_mode"`
	DebugHistory                     types.Int64                 `tfsdk:"debug_history"`
	CpaEntries                       types.List                  `tfsdk:"cpa_entries"`
	SqlDataSource                    types.String                `tfsdk:"sql_data_source"`
	FrontSideTimeout                 types.Int64                 `tfsdk:"front_side_timeout"`
	DependencyActions                []*actions.DependencyAction `tfsdk:"dependency_actions"`
}

var B2BGatewayB2BProfilesCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "b2b_groups",
	AttrType:    "List",
	AttrDefault: "",
	Value:       []string{""},
}
var B2BGatewayArchiveLocationCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "archive_mode",
	AttrType:    "String",
	AttrDefault: "archpurge",
	Value:       []string{"archpurge"},
}
var B2BGatewayArchiveFileNameCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "archive_mode",
	AttrType:    "String",
	AttrDefault: "archpurge",
	Value:       []string{"archpurge"},
}
var B2BGatewayDebugHistoryCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "debug_mode",
	AttrType:    "String",
	AttrDefault: "off",
	Value:       []string{"true"},
}

var B2BGatewayObjectType = map[string]attr.Type{
	"id":                                  types.StringType,
	"app_domain":                          types.StringType,
	"user_summary":                        types.StringType,
	"priority":                            types.StringType,
	"doc_store_location":                  types.StringType,
	"as_front_protocol":                   types.ListType{ElemType: types.ObjectType{AttrTypes: DmASFrontProtocolObjectType}},
	"as1_mdn_email":                       types.StringType,
	"as1_mdn_smtp_server_connection":      types.StringType,
	"as2_mdn_url":                         types.StringType,
	"as3_mdn_url":                         types.StringType,
	"b2b_profiles":                        types.ListType{ElemType: types.ObjectType{AttrTypes: DmB2BActiveProfileObjectType}},
	"b2b_groups":                          types.ListType{ElemType: types.ObjectType{AttrTypes: DmB2BActiveGroupObjectType}},
	"document_routing_preprocessor_type":  types.StringType,
	"document_routing_preprocessor":       types.StringType,
	"document_routing_preprocessor_debug": types.BoolType,
	"archive_mode":                        types.StringType,
	"archive_location":                    types.StringType,
	"archive_file_name":                   types.StringType,
	"archive_minimum_size":                types.Int64Type,
	"archive_document_age":                types.Int64Type,
	"archive_minimum_documents":           types.Int64Type,
	"disk_use_check_interval":             types.Int64Type,
	"max_document_disk_use":               types.Int64Type,
	"archive_monitor":                     types.BoolType,
	"shaping_threshold":                   types.Int64Type,
	"archive_backup_documents":            types.ObjectType{AttrTypes: DmB2BBackupMsgTypeObjectType},
	"xpath_routing_policies":              types.ListType{ElemType: types.StringType},
	"xml_manager":                         types.StringType,
	"debug_mode":                          types.StringType,
	"debug_history":                       types.Int64Type,
	"cpa_entries":                         types.ListType{ElemType: types.ObjectType{AttrTypes: DmB2BCPAEntryObjectType}},
	"sql_data_source":                     types.StringType,
	"front_side_timeout":                  types.Int64Type,
	"dependency_actions":                  actions.ActionsListType,
}

func (data B2BGateway) GetPath() string {
	rest_path := "/mgmt/config/{domain}/B2BGateway"
	rest_path = strings.ReplaceAll(rest_path, "{name}", url.QueryEscape(data.Id.ValueString()))
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data B2BGateway) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.Priority.IsNull() {
		return false
	}
	if !data.DocStoreLocation.IsNull() {
		return false
	}
	if !data.AsFrontProtocol.IsNull() {
		return false
	}
	if !data.As1MdnEmail.IsNull() {
		return false
	}
	if !data.As1MdnSmtpServerConnection.IsNull() {
		return false
	}
	if !data.As2MdnUrl.IsNull() {
		return false
	}
	if !data.As3MdnUrl.IsNull() {
		return false
	}
	if !data.B2bProfiles.IsNull() {
		return false
	}
	if !data.B2bGroups.IsNull() {
		return false
	}
	if !data.DocumentRoutingPreprocessorType.IsNull() {
		return false
	}
	if !data.DocumentRoutingPreprocessor.IsNull() {
		return false
	}
	if !data.DocumentRoutingPreprocessorDebug.IsNull() {
		return false
	}
	if !data.ArchiveMode.IsNull() {
		return false
	}
	if !data.ArchiveLocation.IsNull() {
		return false
	}
	if !data.ArchiveFileName.IsNull() {
		return false
	}
	if !data.ArchiveMinimumSize.IsNull() {
		return false
	}
	if !data.ArchiveDocumentAge.IsNull() {
		return false
	}
	if !data.ArchiveMinimumDocuments.IsNull() {
		return false
	}
	if !data.DiskUseCheckInterval.IsNull() {
		return false
	}
	if !data.MaxDocumentDiskUse.IsNull() {
		return false
	}
	if !data.ArchiveMonitor.IsNull() {
		return false
	}
	if !data.ShapingThreshold.IsNull() {
		return false
	}
	if data.ArchiveBackupDocuments != nil {
		if !data.ArchiveBackupDocuments.IsNull() {
			return false
		}
	}
	if !data.XpathRoutingPolicies.IsNull() {
		return false
	}
	if !data.XmlManager.IsNull() {
		return false
	}
	if !data.DebugMode.IsNull() {
		return false
	}
	if !data.DebugHistory.IsNull() {
		return false
	}
	if !data.CpaEntries.IsNull() {
		return false
	}
	if !data.SqlDataSource.IsNull() {
		return false
	}
	if !data.FrontSideTimeout.IsNull() {
		return false
	}
	return true
}

func (data B2BGateway) ToBody(ctx context.Context, pathRoot string) string {
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
	if !data.Priority.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Priority`, data.Priority.ValueString())
	}
	if !data.DocStoreLocation.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DocStoreLocation`, data.DocStoreLocation.ValueString())
	}
	if !data.AsFrontProtocol.IsNull() {
		var dataValues []DmASFrontProtocol
		data.AsFrontProtocol.ElementsAs(ctx, &dataValues, false)
		for _, val := range dataValues {
			body, _ = sjson.SetRaw(body, pathRoot+`ASFrontProtocol`+".-1", val.ToBody(ctx, ""))
		}
	}
	if !data.As1MdnEmail.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AS1MDNEmail`, data.As1MdnEmail.ValueString())
	}
	if !data.As1MdnSmtpServerConnection.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AS1MDNSMTPServerConnection`, data.As1MdnSmtpServerConnection.ValueString())
	}
	if !data.As2MdnUrl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AS2MDNURL`, data.As2MdnUrl.ValueString())
	}
	if !data.As3MdnUrl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AS3MDNURL`, data.As3MdnUrl.ValueString())
	}
	if !data.B2bProfiles.IsNull() {
		var dataValues []DmB2BActiveProfile
		data.B2bProfiles.ElementsAs(ctx, &dataValues, false)
		for _, val := range dataValues {
			body, _ = sjson.SetRaw(body, pathRoot+`B2BProfiles`+".-1", val.ToBody(ctx, ""))
		}
	}
	if !data.B2bGroups.IsNull() {
		var dataValues []DmB2BActiveGroup
		data.B2bGroups.ElementsAs(ctx, &dataValues, false)
		for _, val := range dataValues {
			body, _ = sjson.SetRaw(body, pathRoot+`B2BGroups`+".-1", val.ToBody(ctx, ""))
		}
	}
	if !data.DocumentRoutingPreprocessorType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DocumentRoutingPreprocessorType`, data.DocumentRoutingPreprocessorType.ValueString())
	}
	if !data.DocumentRoutingPreprocessor.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DocumentRoutingPreprocessor`, data.DocumentRoutingPreprocessor.ValueString())
	}
	if !data.DocumentRoutingPreprocessorDebug.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DocumentRoutingPreprocessorDebug`, tfutils.StringFromBool(data.DocumentRoutingPreprocessorDebug, ""))
	}
	if !data.ArchiveMode.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ArchiveMode`, data.ArchiveMode.ValueString())
	}
	if !data.ArchiveLocation.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ArchiveLocation`, data.ArchiveLocation.ValueString())
	}
	if !data.ArchiveFileName.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ArchiveFileName`, data.ArchiveFileName.ValueString())
	}
	if !data.ArchiveMinimumSize.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ArchiveMinimumSize`, data.ArchiveMinimumSize.ValueInt64())
	}
	if !data.ArchiveDocumentAge.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ArchiveDocumentAge`, data.ArchiveDocumentAge.ValueInt64())
	}
	if !data.ArchiveMinimumDocuments.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ArchiveMinimumDocuments`, data.ArchiveMinimumDocuments.ValueInt64())
	}
	if !data.DiskUseCheckInterval.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DiskUseCheckInterval`, data.DiskUseCheckInterval.ValueInt64())
	}
	if !data.MaxDocumentDiskUse.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MaxDocumentDiskUse`, data.MaxDocumentDiskUse.ValueInt64())
	}
	if !data.ArchiveMonitor.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ArchiveMonitor`, tfutils.StringFromBool(data.ArchiveMonitor, ""))
	}
	if !data.ShapingThreshold.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ShapingThreshold`, data.ShapingThreshold.ValueInt64())
	}
	if data.ArchiveBackupDocuments != nil {
		if !data.ArchiveBackupDocuments.IsNull() {
			body, _ = sjson.SetRaw(body, pathRoot+`ArchiveBackupDocuments`, data.ArchiveBackupDocuments.ToBody(ctx, ""))
		}
	}
	if !data.XpathRoutingPolicies.IsNull() {
		var dataValues []string
		data.XpathRoutingPolicies.ElementsAs(ctx, &dataValues, false)
		for _, val := range dataValues {
			body, _ = sjson.Set(body, pathRoot+`XPathRoutingPolicies`+".-1", map[string]string{"value": val})
		}
	}
	if !data.XmlManager.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`XMLManager`, data.XmlManager.ValueString())
	}
	if !data.DebugMode.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DebugMode`, data.DebugMode.ValueString())
	}
	if !data.DebugHistory.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DebugHistory`, data.DebugHistory.ValueInt64())
	}
	if !data.CpaEntries.IsNull() {
		var dataValues []DmB2BCPAEntry
		data.CpaEntries.ElementsAs(ctx, &dataValues, false)
		for _, val := range dataValues {
			body, _ = sjson.SetRaw(body, pathRoot+`CPAEntries`+".-1", val.ToBody(ctx, ""))
		}
	}
	if !data.SqlDataSource.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SQLDataSource`, data.SqlDataSource.ValueString())
	}
	if !data.FrontSideTimeout.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`FrontSideTimeout`, data.FrontSideTimeout.ValueInt64())
	}
	return body
}

func (data *B2BGateway) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `Priority`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Priority = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Priority = types.StringValue("normal")
	}
	if value := res.Get(pathRoot + `DocStoreLocation`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.DocStoreLocation = tfutils.ParseStringFromGJSON(value)
	} else {
		data.DocStoreLocation = types.StringValue("(default)")
	}
	if value := res.Get(pathRoot + `ASFrontProtocol`); value.Exists() {
		l := []DmASFrontProtocol{}
		if value := res.Get(`ASFrontProtocol`); value.Exists() {
			for _, v := range value.Array() {
				item := DmASFrontProtocol{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.AsFrontProtocol, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmASFrontProtocolObjectType}, l)
		} else {
			data.AsFrontProtocol = types.ListNull(types.ObjectType{AttrTypes: DmASFrontProtocolObjectType})
		}
	} else {
		data.AsFrontProtocol = types.ListNull(types.ObjectType{AttrTypes: DmASFrontProtocolObjectType})
	}
	if value := res.Get(pathRoot + `AS1MDNEmail`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.As1MdnEmail = tfutils.ParseStringFromGJSON(value)
	} else {
		data.As1MdnEmail = types.StringNull()
	}
	if value := res.Get(pathRoot + `AS1MDNSMTPServerConnection`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.As1MdnSmtpServerConnection = tfutils.ParseStringFromGJSON(value)
	} else {
		data.As1MdnSmtpServerConnection = types.StringNull()
	}
	if value := res.Get(pathRoot + `AS2MDNURL`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.As2MdnUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.As2MdnUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `AS3MDNURL`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.As3MdnUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.As3MdnUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `B2BProfiles`); value.Exists() {
		l := []DmB2BActiveProfile{}
		if value := res.Get(`B2BProfiles`); value.Exists() {
			for _, v := range value.Array() {
				item := DmB2BActiveProfile{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.B2bProfiles, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmB2BActiveProfileObjectType}, l)
		} else {
			data.B2bProfiles = types.ListNull(types.ObjectType{AttrTypes: DmB2BActiveProfileObjectType})
		}
	} else {
		data.B2bProfiles = types.ListNull(types.ObjectType{AttrTypes: DmB2BActiveProfileObjectType})
	}
	if value := res.Get(pathRoot + `B2BGroups`); value.Exists() {
		l := []DmB2BActiveGroup{}
		if value := res.Get(`B2BGroups`); value.Exists() {
			for _, v := range value.Array() {
				item := DmB2BActiveGroup{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.B2bGroups, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmB2BActiveGroupObjectType}, l)
		} else {
			data.B2bGroups = types.ListNull(types.ObjectType{AttrTypes: DmB2BActiveGroupObjectType})
		}
	} else {
		data.B2bGroups = types.ListNull(types.ObjectType{AttrTypes: DmB2BActiveGroupObjectType})
	}
	if value := res.Get(pathRoot + `DocumentRoutingPreprocessorType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.DocumentRoutingPreprocessorType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.DocumentRoutingPreprocessorType = types.StringValue("stylesheet")
	}
	if value := res.Get(pathRoot + `DocumentRoutingPreprocessor`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.DocumentRoutingPreprocessor = tfutils.ParseStringFromGJSON(value)
	} else {
		data.DocumentRoutingPreprocessor = types.StringValue("store:///b2b-routing.xsl")
	}
	if value := res.Get(pathRoot + `DocumentRoutingPreprocessorDebug`); value.Exists() {
		data.DocumentRoutingPreprocessorDebug = tfutils.BoolFromString(value.String())
	} else {
		data.DocumentRoutingPreprocessorDebug = types.BoolNull()
	}
	if value := res.Get(pathRoot + `ArchiveMode`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ArchiveMode = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ArchiveMode = types.StringValue("archpurge")
	}
	if value := res.Get(pathRoot + `ArchiveLocation`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ArchiveLocation = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ArchiveLocation = types.StringNull()
	}
	if value := res.Get(pathRoot + `ArchiveFileName`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ArchiveFileName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ArchiveFileName = types.StringNull()
	}
	if value := res.Get(pathRoot + `ArchiveMinimumSize`); value.Exists() {
		data.ArchiveMinimumSize = types.Int64Value(value.Int())
	} else {
		data.ArchiveMinimumSize = types.Int64Value(1024)
	}
	if value := res.Get(pathRoot + `ArchiveDocumentAge`); value.Exists() {
		data.ArchiveDocumentAge = types.Int64Value(value.Int())
	} else {
		data.ArchiveDocumentAge = types.Int64Value(90)
	}
	if value := res.Get(pathRoot + `ArchiveMinimumDocuments`); value.Exists() {
		data.ArchiveMinimumDocuments = types.Int64Value(value.Int())
	} else {
		data.ArchiveMinimumDocuments = types.Int64Value(100)
	}
	if value := res.Get(pathRoot + `DiskUseCheckInterval`); value.Exists() {
		data.DiskUseCheckInterval = types.Int64Value(value.Int())
	} else {
		data.DiskUseCheckInterval = types.Int64Value(60)
	}
	if value := res.Get(pathRoot + `MaxDocumentDiskUse`); value.Exists() {
		data.MaxDocumentDiskUse = types.Int64Value(value.Int())
	} else {
		data.MaxDocumentDiskUse = types.Int64Value(25165824)
	}
	if value := res.Get(pathRoot + `ArchiveMonitor`); value.Exists() {
		data.ArchiveMonitor = tfutils.BoolFromString(value.String())
	} else {
		data.ArchiveMonitor = types.BoolNull()
	}
	if value := res.Get(pathRoot + `ShapingThreshold`); value.Exists() {
		data.ShapingThreshold = types.Int64Value(value.Int())
	} else {
		data.ShapingThreshold = types.Int64Value(200)
	}
	if value := res.Get(pathRoot + `ArchiveBackupDocuments`); value.Exists() {
		data.ArchiveBackupDocuments = &DmB2BBackupMsgType{}
		data.ArchiveBackupDocuments.FromBody(ctx, "", value)
	} else {
		data.ArchiveBackupDocuments = nil
	}
	if value := res.Get(pathRoot + `XPathRoutingPolicies`); value.Exists() {
		data.XpathRoutingPolicies = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.XpathRoutingPolicies = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `XMLManager`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.XmlManager = tfutils.ParseStringFromGJSON(value)
	} else {
		data.XmlManager = types.StringValue("default")
	}
	if value := res.Get(pathRoot + `DebugMode`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.DebugMode = tfutils.ParseStringFromGJSON(value)
	} else {
		data.DebugMode = types.StringValue("off")
	}
	if value := res.Get(pathRoot + `DebugHistory`); value.Exists() {
		data.DebugHistory = types.Int64Value(value.Int())
	} else {
		data.DebugHistory = types.Int64Value(25)
	}
	if value := res.Get(pathRoot + `CPAEntries`); value.Exists() {
		l := []DmB2BCPAEntry{}
		if value := res.Get(`CPAEntries`); value.Exists() {
			for _, v := range value.Array() {
				item := DmB2BCPAEntry{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.CpaEntries, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmB2BCPAEntryObjectType}, l)
		} else {
			data.CpaEntries = types.ListNull(types.ObjectType{AttrTypes: DmB2BCPAEntryObjectType})
		}
	} else {
		data.CpaEntries = types.ListNull(types.ObjectType{AttrTypes: DmB2BCPAEntryObjectType})
	}
	if value := res.Get(pathRoot + `SQLDataSource`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SqlDataSource = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SqlDataSource = types.StringNull()
	}
	if value := res.Get(pathRoot + `FrontSideTimeout`); value.Exists() {
		data.FrontSideTimeout = types.Int64Value(value.Int())
	} else {
		data.FrontSideTimeout = types.Int64Value(120)
	}
}

func (data *B2BGateway) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `Priority`); value.Exists() && !data.Priority.IsNull() {
		data.Priority = tfutils.ParseStringFromGJSON(value)
	} else if data.Priority.ValueString() != "normal" {
		data.Priority = types.StringNull()
	}
	if value := res.Get(pathRoot + `DocStoreLocation`); value.Exists() && !data.DocStoreLocation.IsNull() {
		data.DocStoreLocation = tfutils.ParseStringFromGJSON(value)
	} else if data.DocStoreLocation.ValueString() != "(default)" {
		data.DocStoreLocation = types.StringNull()
	}
	if value := res.Get(pathRoot + `ASFrontProtocol`); value.Exists() && !data.AsFrontProtocol.IsNull() {
		l := []DmASFrontProtocol{}
		for _, v := range value.Array() {
			item := DmASFrontProtocol{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
		if len(l) > 0 {
			data.AsFrontProtocol, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmASFrontProtocolObjectType}, l)
		} else {
			data.AsFrontProtocol = types.ListNull(types.ObjectType{AttrTypes: DmASFrontProtocolObjectType})
		}
	} else {
		data.AsFrontProtocol = types.ListNull(types.ObjectType{AttrTypes: DmASFrontProtocolObjectType})
	}
	if value := res.Get(pathRoot + `AS1MDNEmail`); value.Exists() && !data.As1MdnEmail.IsNull() {
		data.As1MdnEmail = tfutils.ParseStringFromGJSON(value)
	} else {
		data.As1MdnEmail = types.StringNull()
	}
	if value := res.Get(pathRoot + `AS1MDNSMTPServerConnection`); value.Exists() && !data.As1MdnSmtpServerConnection.IsNull() {
		data.As1MdnSmtpServerConnection = tfutils.ParseStringFromGJSON(value)
	} else {
		data.As1MdnSmtpServerConnection = types.StringNull()
	}
	if value := res.Get(pathRoot + `AS2MDNURL`); value.Exists() && !data.As2MdnUrl.IsNull() {
		data.As2MdnUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.As2MdnUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `AS3MDNURL`); value.Exists() && !data.As3MdnUrl.IsNull() {
		data.As3MdnUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.As3MdnUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `B2BProfiles`); value.Exists() && !data.B2bProfiles.IsNull() {
		l := []DmB2BActiveProfile{}
		for _, v := range value.Array() {
			item := DmB2BActiveProfile{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
		if len(l) > 0 {
			data.B2bProfiles, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmB2BActiveProfileObjectType}, l)
		} else {
			data.B2bProfiles = types.ListNull(types.ObjectType{AttrTypes: DmB2BActiveProfileObjectType})
		}
	} else {
		data.B2bProfiles = types.ListNull(types.ObjectType{AttrTypes: DmB2BActiveProfileObjectType})
	}
	if value := res.Get(pathRoot + `B2BGroups`); value.Exists() && !data.B2bGroups.IsNull() {
		l := []DmB2BActiveGroup{}
		for _, v := range value.Array() {
			item := DmB2BActiveGroup{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
		if len(l) > 0 {
			data.B2bGroups, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmB2BActiveGroupObjectType}, l)
		} else {
			data.B2bGroups = types.ListNull(types.ObjectType{AttrTypes: DmB2BActiveGroupObjectType})
		}
	} else {
		data.B2bGroups = types.ListNull(types.ObjectType{AttrTypes: DmB2BActiveGroupObjectType})
	}
	if value := res.Get(pathRoot + `DocumentRoutingPreprocessorType`); value.Exists() && !data.DocumentRoutingPreprocessorType.IsNull() {
		data.DocumentRoutingPreprocessorType = tfutils.ParseStringFromGJSON(value)
	} else if data.DocumentRoutingPreprocessorType.ValueString() != "stylesheet" {
		data.DocumentRoutingPreprocessorType = types.StringNull()
	}
	if value := res.Get(pathRoot + `DocumentRoutingPreprocessor`); value.Exists() && !data.DocumentRoutingPreprocessor.IsNull() {
		data.DocumentRoutingPreprocessor = tfutils.ParseStringFromGJSON(value)
	} else if data.DocumentRoutingPreprocessor.ValueString() != "store:///b2b-routing.xsl" {
		data.DocumentRoutingPreprocessor = types.StringNull()
	}
	if value := res.Get(pathRoot + `DocumentRoutingPreprocessorDebug`); value.Exists() && !data.DocumentRoutingPreprocessorDebug.IsNull() {
		data.DocumentRoutingPreprocessorDebug = tfutils.BoolFromString(value.String())
	} else if data.DocumentRoutingPreprocessorDebug.ValueBool() {
		data.DocumentRoutingPreprocessorDebug = types.BoolNull()
	}
	if value := res.Get(pathRoot + `ArchiveMode`); value.Exists() && !data.ArchiveMode.IsNull() {
		data.ArchiveMode = tfutils.ParseStringFromGJSON(value)
	} else if data.ArchiveMode.ValueString() != "archpurge" {
		data.ArchiveMode = types.StringNull()
	}
	if value := res.Get(pathRoot + `ArchiveLocation`); value.Exists() && !data.ArchiveLocation.IsNull() {
		data.ArchiveLocation = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ArchiveLocation = types.StringNull()
	}
	if value := res.Get(pathRoot + `ArchiveFileName`); value.Exists() && !data.ArchiveFileName.IsNull() {
		data.ArchiveFileName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ArchiveFileName = types.StringNull()
	}
	if value := res.Get(pathRoot + `ArchiveMinimumSize`); value.Exists() && !data.ArchiveMinimumSize.IsNull() {
		data.ArchiveMinimumSize = types.Int64Value(value.Int())
	} else if data.ArchiveMinimumSize.ValueInt64() != 1024 {
		data.ArchiveMinimumSize = types.Int64Null()
	}
	if value := res.Get(pathRoot + `ArchiveDocumentAge`); value.Exists() && !data.ArchiveDocumentAge.IsNull() {
		data.ArchiveDocumentAge = types.Int64Value(value.Int())
	} else if data.ArchiveDocumentAge.ValueInt64() != 90 {
		data.ArchiveDocumentAge = types.Int64Null()
	}
	if value := res.Get(pathRoot + `ArchiveMinimumDocuments`); value.Exists() && !data.ArchiveMinimumDocuments.IsNull() {
		data.ArchiveMinimumDocuments = types.Int64Value(value.Int())
	} else if data.ArchiveMinimumDocuments.ValueInt64() != 100 {
		data.ArchiveMinimumDocuments = types.Int64Null()
	}
	if value := res.Get(pathRoot + `DiskUseCheckInterval`); value.Exists() && !data.DiskUseCheckInterval.IsNull() {
		data.DiskUseCheckInterval = types.Int64Value(value.Int())
	} else if data.DiskUseCheckInterval.ValueInt64() != 60 {
		data.DiskUseCheckInterval = types.Int64Null()
	}
	if value := res.Get(pathRoot + `MaxDocumentDiskUse`); value.Exists() && !data.MaxDocumentDiskUse.IsNull() {
		data.MaxDocumentDiskUse = types.Int64Value(value.Int())
	} else if data.MaxDocumentDiskUse.ValueInt64() != 25165824 {
		data.MaxDocumentDiskUse = types.Int64Null()
	}
	if value := res.Get(pathRoot + `ArchiveMonitor`); value.Exists() && !data.ArchiveMonitor.IsNull() {
		data.ArchiveMonitor = tfutils.BoolFromString(value.String())
	} else if !data.ArchiveMonitor.ValueBool() {
		data.ArchiveMonitor = types.BoolNull()
	}
	if value := res.Get(pathRoot + `ShapingThreshold`); value.Exists() && !data.ShapingThreshold.IsNull() {
		data.ShapingThreshold = types.Int64Value(value.Int())
	} else if data.ShapingThreshold.ValueInt64() != 200 {
		data.ShapingThreshold = types.Int64Null()
	}
	if value := res.Get(pathRoot + `ArchiveBackupDocuments`); value.Exists() {
		data.ArchiveBackupDocuments.UpdateFromBody(ctx, "", value)
	} else {
		data.ArchiveBackupDocuments = nil
	}
	if value := res.Get(pathRoot + `XPathRoutingPolicies`); value.Exists() && !data.XpathRoutingPolicies.IsNull() {
		data.XpathRoutingPolicies = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.XpathRoutingPolicies = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `XMLManager`); value.Exists() && !data.XmlManager.IsNull() {
		data.XmlManager = tfutils.ParseStringFromGJSON(value)
	} else if data.XmlManager.ValueString() != "default" {
		data.XmlManager = types.StringNull()
	}
	if value := res.Get(pathRoot + `DebugMode`); value.Exists() && !data.DebugMode.IsNull() {
		data.DebugMode = tfutils.ParseStringFromGJSON(value)
	} else if data.DebugMode.ValueString() != "off" {
		data.DebugMode = types.StringNull()
	}
	if value := res.Get(pathRoot + `DebugHistory`); value.Exists() && !data.DebugHistory.IsNull() {
		data.DebugHistory = types.Int64Value(value.Int())
	} else if data.DebugHistory.ValueInt64() != 25 {
		data.DebugHistory = types.Int64Null()
	}
	if value := res.Get(pathRoot + `CPAEntries`); value.Exists() && !data.CpaEntries.IsNull() {
		l := []DmB2BCPAEntry{}
		for _, v := range value.Array() {
			item := DmB2BCPAEntry{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
		if len(l) > 0 {
			data.CpaEntries, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmB2BCPAEntryObjectType}, l)
		} else {
			data.CpaEntries = types.ListNull(types.ObjectType{AttrTypes: DmB2BCPAEntryObjectType})
		}
	} else {
		data.CpaEntries = types.ListNull(types.ObjectType{AttrTypes: DmB2BCPAEntryObjectType})
	}
	if value := res.Get(pathRoot + `SQLDataSource`); value.Exists() && !data.SqlDataSource.IsNull() {
		data.SqlDataSource = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SqlDataSource = types.StringNull()
	}
	if value := res.Get(pathRoot + `FrontSideTimeout`); value.Exists() && !data.FrontSideTimeout.IsNull() {
		data.FrontSideTimeout = types.Int64Value(value.Int())
	} else if data.FrontSideTimeout.ValueInt64() != 120 {
		data.FrontSideTimeout = types.Int64Null()
	}
}

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

type AAAPolicy struct {
	Id                            types.String                `tfsdk:"id"`
	AppDomain                     types.String                `tfsdk:"app_domain"`
	UserSummary                   types.String                `tfsdk:"user_summary"`
	AuthorizedCounter             types.String                `tfsdk:"authorized_counter"`
	RejectedCounter               types.String                `tfsdk:"rejected_counter"`
	NamespaceMapping              types.List                  `tfsdk:"namespace_mapping"`
	ExtractIdentity               *DmAAAPExtractIdentity      `tfsdk:"extract_identity"`
	Authenticate                  *DmAAAPAuthenticate         `tfsdk:"authenticate"`
	MapCredentials                *DmAAAPMapCredentials       `tfsdk:"map_credentials"`
	ExtractResource               *DmAAAPExtractResource      `tfsdk:"extract_resource"`
	MapResource                   *DmAAAPMapResource          `tfsdk:"map_resource"`
	Authorize                     *DmAAAPAuthorize            `tfsdk:"authorize"`
	PostProcess                   *DmAAAPPostProcess          `tfsdk:"post_process"`
	SamlAttribute                 types.List                  `tfsdk:"saml_attribute"`
	LtpaAttributes                types.List                  `tfsdk:"ltpa_attributes"`
	TransactionPriority           types.List                  `tfsdk:"transaction_priority"`
	SamlValcred                   types.String                `tfsdk:"saml_valcred"`
	SamlSigningKey                types.String                `tfsdk:"saml_signing_key"`
	SamlSigningCert               types.String                `tfsdk:"saml_signing_cert"`
	SamlSigningHashAlg            types.String                `tfsdk:"saml_signing_hash_alg"`
	SamlSigningAlg                types.String                `tfsdk:"saml_signing_alg"`
	LdapSuffix                    types.String                `tfsdk:"ldap_suffix"`
	LogAllowed                    types.Bool                  `tfsdk:"log_allowed"`
	LogAllowedLevel               types.String                `tfsdk:"log_allowed_level"`
	LogRejected                   types.Bool                  `tfsdk:"log_rejected"`
	LogRejectedLevel              types.String                `tfsdk:"log_rejected_level"`
	WsSecureConversationCryptoKey types.String                `tfsdk:"ws_secure_conversation_crypto_key"`
	SamlSourceIdMappingFile       types.String                `tfsdk:"saml_source_id_mapping_file"`
	PingIdentityCompatibility     types.Bool                  `tfsdk:"ping_identity_compatibility"`
	Saml2MetadataFile             types.String                `tfsdk:"saml2_metadata_file"`
	DoSValve                      types.Int64                 `tfsdk:"do_s_valve"`
	LdapVersion                   types.String                `tfsdk:"ldap_version"`
	EnforceSoapActor              types.Bool                  `tfsdk:"enforce_soap_actor"`
	WsSecActorRoleId              types.String                `tfsdk:"ws_sec_actor_role_id"`
	AuSmHttpHeader                types.List                  `tfsdk:"au_sm_http_header"`
	AzSmHttpHeader                types.List                  `tfsdk:"az_sm_http_header"`
	DynConfig                     types.String                `tfsdk:"dyn_config"`
	ExternalAaaTemplate           types.String                `tfsdk:"external_aaa_template"`
	DynConfigCustomUrl            types.String                `tfsdk:"dyn_config_custom_url"`
	DependencyActions             []*actions.DependencyAction `tfsdk:"dependency_actions"`
}

var AAAPolicyLogAllowedLevelCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "log_allowed",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"true"},
}
var AAAPolicyLogRejectedLevelCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "log_rejected",
	AttrType:    "Bool",
	AttrDefault: "true",
	Value:       []string{"true"},
}
var AAAPolicyExternalAAATemplateCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "dyn_config",
	AttrType:    "String",
	AttrDefault: "none",
	Value:       []string{"external-aaa"},
}
var AAAPolicyDynConfigCustomURLCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "dyn_config",
	AttrType:    "String",
	AttrDefault: "none",
	Value:       []string{"current-aaa", "external-aaa"},
}

var AAAPolicyObjectType = map[string]attr.Type{
	"id":                                types.StringType,
	"app_domain":                        types.StringType,
	"user_summary":                      types.StringType,
	"authorized_counter":                types.StringType,
	"rejected_counter":                  types.StringType,
	"namespace_mapping":                 types.ListType{ElemType: types.ObjectType{AttrTypes: DmNamespaceMappingObjectType}},
	"extract_identity":                  types.ObjectType{AttrTypes: DmAAAPExtractIdentityObjectType},
	"authenticate":                      types.ObjectType{AttrTypes: DmAAAPAuthenticateObjectType},
	"map_credentials":                   types.ObjectType{AttrTypes: DmAAAPMapCredentialsObjectType},
	"extract_resource":                  types.ObjectType{AttrTypes: DmAAAPExtractResourceObjectType},
	"map_resource":                      types.ObjectType{AttrTypes: DmAAAPMapResourceObjectType},
	"authorize":                         types.ObjectType{AttrTypes: DmAAAPAuthorizeObjectType},
	"post_process":                      types.ObjectType{AttrTypes: DmAAAPPostProcessObjectType},
	"saml_attribute":                    types.ListType{ElemType: types.ObjectType{AttrTypes: DmSAMLAttributeNameAndValueObjectType}},
	"ltpa_attributes":                   types.ListType{ElemType: types.ObjectType{AttrTypes: DmLTPAUserAttributeNameAndValueObjectType}},
	"transaction_priority":              types.ListType{ElemType: types.ObjectType{AttrTypes: DmAAATransactionPriorityObjectType}},
	"saml_valcred":                      types.StringType,
	"saml_signing_key":                  types.StringType,
	"saml_signing_cert":                 types.StringType,
	"saml_signing_hash_alg":             types.StringType,
	"saml_signing_alg":                  types.StringType,
	"ldap_suffix":                       types.StringType,
	"log_allowed":                       types.BoolType,
	"log_allowed_level":                 types.StringType,
	"log_rejected":                      types.BoolType,
	"log_rejected_level":                types.StringType,
	"ws_secure_conversation_crypto_key": types.StringType,
	"saml_source_id_mapping_file":       types.StringType,
	"ping_identity_compatibility":       types.BoolType,
	"saml2_metadata_file":               types.StringType,
	"do_s_valve":                        types.Int64Type,
	"ldap_version":                      types.StringType,
	"enforce_soap_actor":                types.BoolType,
	"ws_sec_actor_role_id":              types.StringType,
	"au_sm_http_header":                 types.ListType{ElemType: types.StringType},
	"az_sm_http_header":                 types.ListType{ElemType: types.StringType},
	"dyn_config":                        types.StringType,
	"external_aaa_template":             types.StringType,
	"dyn_config_custom_url":             types.StringType,
	"dependency_actions":                actions.ActionsListType,
}

func (data AAAPolicy) GetPath() string {
	rest_path := "/mgmt/config/{domain}/AAAPolicy"
	rest_path = strings.ReplaceAll(rest_path, "{name}", url.QueryEscape(data.Id.ValueString()))
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data AAAPolicy) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.AuthorizedCounter.IsNull() {
		return false
	}
	if !data.RejectedCounter.IsNull() {
		return false
	}
	if !data.NamespaceMapping.IsNull() {
		return false
	}
	if data.ExtractIdentity != nil {
		if !data.ExtractIdentity.IsNull() {
			return false
		}
	}
	if data.Authenticate != nil {
		if !data.Authenticate.IsNull() {
			return false
		}
	}
	if data.MapCredentials != nil {
		if !data.MapCredentials.IsNull() {
			return false
		}
	}
	if data.ExtractResource != nil {
		if !data.ExtractResource.IsNull() {
			return false
		}
	}
	if data.MapResource != nil {
		if !data.MapResource.IsNull() {
			return false
		}
	}
	if data.Authorize != nil {
		if !data.Authorize.IsNull() {
			return false
		}
	}
	if data.PostProcess != nil {
		if !data.PostProcess.IsNull() {
			return false
		}
	}
	if !data.SamlAttribute.IsNull() {
		return false
	}
	if !data.LtpaAttributes.IsNull() {
		return false
	}
	if !data.TransactionPriority.IsNull() {
		return false
	}
	if !data.SamlValcred.IsNull() {
		return false
	}
	if !data.SamlSigningKey.IsNull() {
		return false
	}
	if !data.SamlSigningCert.IsNull() {
		return false
	}
	if !data.SamlSigningHashAlg.IsNull() {
		return false
	}
	if !data.SamlSigningAlg.IsNull() {
		return false
	}
	if !data.LdapSuffix.IsNull() {
		return false
	}
	if !data.LogAllowed.IsNull() {
		return false
	}
	if !data.LogAllowedLevel.IsNull() {
		return false
	}
	if !data.LogRejected.IsNull() {
		return false
	}
	if !data.LogRejectedLevel.IsNull() {
		return false
	}
	if !data.WsSecureConversationCryptoKey.IsNull() {
		return false
	}
	if !data.SamlSourceIdMappingFile.IsNull() {
		return false
	}
	if !data.PingIdentityCompatibility.IsNull() {
		return false
	}
	if !data.Saml2MetadataFile.IsNull() {
		return false
	}
	if !data.DoSValve.IsNull() {
		return false
	}
	if !data.LdapVersion.IsNull() {
		return false
	}
	if !data.EnforceSoapActor.IsNull() {
		return false
	}
	if !data.WsSecActorRoleId.IsNull() {
		return false
	}
	if !data.AuSmHttpHeader.IsNull() {
		return false
	}
	if !data.AzSmHttpHeader.IsNull() {
		return false
	}
	if !data.DynConfig.IsNull() {
		return false
	}
	if !data.ExternalAaaTemplate.IsNull() {
		return false
	}
	if !data.DynConfigCustomUrl.IsNull() {
		return false
	}
	return true
}

func (data AAAPolicy) ToBody(ctx context.Context, pathRoot string) string {
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
	if !data.AuthorizedCounter.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AuthorizedCounter`, data.AuthorizedCounter.ValueString())
	}
	if !data.RejectedCounter.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`RejectedCounter`, data.RejectedCounter.ValueString())
	}
	if !data.NamespaceMapping.IsNull() {
		var dataValues []DmNamespaceMapping
		data.NamespaceMapping.ElementsAs(ctx, &dataValues, false)
		for _, val := range dataValues {
			body, _ = sjson.SetRaw(body, pathRoot+`NamespaceMapping`+".-1", val.ToBody(ctx, ""))
		}
	}
	if data.ExtractIdentity != nil {
		if !data.ExtractIdentity.IsNull() {
			body, _ = sjson.SetRaw(body, pathRoot+`ExtractIdentity`, data.ExtractIdentity.ToBody(ctx, ""))
		}
	}
	if data.Authenticate != nil {
		if !data.Authenticate.IsNull() {
			body, _ = sjson.SetRaw(body, pathRoot+`Authenticate`, data.Authenticate.ToBody(ctx, ""))
		}
	}
	if data.MapCredentials != nil {
		if !data.MapCredentials.IsNull() {
			body, _ = sjson.SetRaw(body, pathRoot+`MapCredentials`, data.MapCredentials.ToBody(ctx, ""))
		}
	}
	if data.ExtractResource != nil {
		if !data.ExtractResource.IsNull() {
			body, _ = sjson.SetRaw(body, pathRoot+`ExtractResource`, data.ExtractResource.ToBody(ctx, ""))
		}
	}
	if data.MapResource != nil {
		if !data.MapResource.IsNull() {
			body, _ = sjson.SetRaw(body, pathRoot+`MapResource`, data.MapResource.ToBody(ctx, ""))
		}
	}
	if data.Authorize != nil {
		if !data.Authorize.IsNull() {
			body, _ = sjson.SetRaw(body, pathRoot+`Authorize`, data.Authorize.ToBody(ctx, ""))
		}
	}
	if data.PostProcess != nil {
		if !data.PostProcess.IsNull() {
			body, _ = sjson.SetRaw(body, pathRoot+`PostProcess`, data.PostProcess.ToBody(ctx, ""))
		}
	}
	if !data.SamlAttribute.IsNull() {
		var dataValues []DmSAMLAttributeNameAndValue
		data.SamlAttribute.ElementsAs(ctx, &dataValues, false)
		for _, val := range dataValues {
			body, _ = sjson.SetRaw(body, pathRoot+`SAMLAttribute`+".-1", val.ToBody(ctx, ""))
		}
	}
	if !data.LtpaAttributes.IsNull() {
		var dataValues []DmLTPAUserAttributeNameAndValue
		data.LtpaAttributes.ElementsAs(ctx, &dataValues, false)
		for _, val := range dataValues {
			body, _ = sjson.SetRaw(body, pathRoot+`LTPAAttributes`+".-1", val.ToBody(ctx, ""))
		}
	}
	if !data.TransactionPriority.IsNull() {
		var dataValues []DmAAATransactionPriority
		data.TransactionPriority.ElementsAs(ctx, &dataValues, false)
		for _, val := range dataValues {
			body, _ = sjson.SetRaw(body, pathRoot+`TransactionPriority`+".-1", val.ToBody(ctx, ""))
		}
	}
	if !data.SamlValcred.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SAMLValcred`, data.SamlValcred.ValueString())
	}
	if !data.SamlSigningKey.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SAMLSigningKey`, data.SamlSigningKey.ValueString())
	}
	if !data.SamlSigningCert.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SAMLSigningCert`, data.SamlSigningCert.ValueString())
	}
	if !data.SamlSigningHashAlg.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SAMLSigningHashAlg`, data.SamlSigningHashAlg.ValueString())
	}
	if !data.SamlSigningAlg.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SAMLSigningAlg`, data.SamlSigningAlg.ValueString())
	}
	if !data.LdapSuffix.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LDAPsuffix`, data.LdapSuffix.ValueString())
	}
	if !data.LogAllowed.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LogAllowed`, tfutils.StringFromBool(data.LogAllowed, ""))
	}
	if !data.LogAllowedLevel.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LogAllowedLevel`, data.LogAllowedLevel.ValueString())
	}
	if !data.LogRejected.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LogRejected`, tfutils.StringFromBool(data.LogRejected, ""))
	}
	if !data.LogRejectedLevel.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LogRejectedLevel`, data.LogRejectedLevel.ValueString())
	}
	if !data.WsSecureConversationCryptoKey.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`WSSecureConversationCryptoKey`, data.WsSecureConversationCryptoKey.ValueString())
	}
	if !data.SamlSourceIdMappingFile.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SAMLSourceIDMappingFile`, data.SamlSourceIdMappingFile.ValueString())
	}
	if !data.PingIdentityCompatibility.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PingIdentityCompatibility`, tfutils.StringFromBool(data.PingIdentityCompatibility, ""))
	}
	if !data.Saml2MetadataFile.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SAML2MetadataFile`, data.Saml2MetadataFile.ValueString())
	}
	if !data.DoSValve.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DoSValve`, data.DoSValve.ValueInt64())
	}
	if !data.LdapVersion.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LDAPVersion`, data.LdapVersion.ValueString())
	}
	if !data.EnforceSoapActor.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EnforceSOAPActor`, tfutils.StringFromBool(data.EnforceSoapActor, ""))
	}
	if !data.WsSecActorRoleId.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`WSSecActorRoleID`, data.WsSecActorRoleId.ValueString())
	}
	if !data.AuSmHttpHeader.IsNull() {
		var dataValues []string
		data.AuSmHttpHeader.ElementsAs(ctx, &dataValues, false)
		for _, val := range dataValues {
			body, _ = sjson.Set(body, pathRoot+`AUSMHTTPHeader`+".-1", map[string]string{"value": val})
		}
	}
	if !data.AzSmHttpHeader.IsNull() {
		var dataValues []string
		data.AzSmHttpHeader.ElementsAs(ctx, &dataValues, false)
		for _, val := range dataValues {
			body, _ = sjson.Set(body, pathRoot+`AZSMHTTPHeader`+".-1", map[string]string{"value": val})
		}
	}
	if !data.DynConfig.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DynConfig`, data.DynConfig.ValueString())
	}
	if !data.ExternalAaaTemplate.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ExternalAAATemplate`, data.ExternalAaaTemplate.ValueString())
	}
	if !data.DynConfigCustomUrl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DynConfigCustomURL`, data.DynConfigCustomUrl.ValueString())
	}
	return body
}

func (data *AAAPolicy) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `AuthorizedCounter`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AuthorizedCounter = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuthorizedCounter = types.StringNull()
	}
	if value := res.Get(pathRoot + `RejectedCounter`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.RejectedCounter = tfutils.ParseStringFromGJSON(value)
	} else {
		data.RejectedCounter = types.StringNull()
	}
	if value := res.Get(pathRoot + `NamespaceMapping`); value.Exists() {
		l := []DmNamespaceMapping{}
		if value := res.Get(`NamespaceMapping`); value.Exists() {
			for _, v := range value.Array() {
				item := DmNamespaceMapping{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.NamespaceMapping, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmNamespaceMappingObjectType}, l)
		} else {
			data.NamespaceMapping = types.ListNull(types.ObjectType{AttrTypes: DmNamespaceMappingObjectType})
		}
	} else {
		data.NamespaceMapping = types.ListNull(types.ObjectType{AttrTypes: DmNamespaceMappingObjectType})
	}
	if value := res.Get(pathRoot + `ExtractIdentity`); value.Exists() {
		data.ExtractIdentity = &DmAAAPExtractIdentity{}
		data.ExtractIdentity.FromBody(ctx, "", value)
	} else {
		data.ExtractIdentity = nil
	}
	if value := res.Get(pathRoot + `Authenticate`); value.Exists() {
		data.Authenticate = &DmAAAPAuthenticate{}
		data.Authenticate.FromBody(ctx, "", value)
	} else {
		data.Authenticate = nil
	}
	if value := res.Get(pathRoot + `MapCredentials`); value.Exists() {
		data.MapCredentials = &DmAAAPMapCredentials{}
		data.MapCredentials.FromBody(ctx, "", value)
	} else {
		data.MapCredentials = nil
	}
	if value := res.Get(pathRoot + `ExtractResource`); value.Exists() {
		data.ExtractResource = &DmAAAPExtractResource{}
		data.ExtractResource.FromBody(ctx, "", value)
	} else {
		data.ExtractResource = nil
	}
	if value := res.Get(pathRoot + `MapResource`); value.Exists() {
		data.MapResource = &DmAAAPMapResource{}
		data.MapResource.FromBody(ctx, "", value)
	} else {
		data.MapResource = nil
	}
	if value := res.Get(pathRoot + `Authorize`); value.Exists() {
		data.Authorize = &DmAAAPAuthorize{}
		data.Authorize.FromBody(ctx, "", value)
	} else {
		data.Authorize = nil
	}
	if value := res.Get(pathRoot + `PostProcess`); value.Exists() {
		data.PostProcess = &DmAAAPPostProcess{}
		data.PostProcess.FromBody(ctx, "", value)
	} else {
		data.PostProcess = nil
	}
	if value := res.Get(pathRoot + `SAMLAttribute`); value.Exists() {
		l := []DmSAMLAttributeNameAndValue{}
		if value := res.Get(`SAMLAttribute`); value.Exists() {
			for _, v := range value.Array() {
				item := DmSAMLAttributeNameAndValue{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.SamlAttribute, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmSAMLAttributeNameAndValueObjectType}, l)
		} else {
			data.SamlAttribute = types.ListNull(types.ObjectType{AttrTypes: DmSAMLAttributeNameAndValueObjectType})
		}
	} else {
		data.SamlAttribute = types.ListNull(types.ObjectType{AttrTypes: DmSAMLAttributeNameAndValueObjectType})
	}
	if value := res.Get(pathRoot + `LTPAAttributes`); value.Exists() {
		l := []DmLTPAUserAttributeNameAndValue{}
		if value := res.Get(`LTPAAttributes`); value.Exists() {
			for _, v := range value.Array() {
				item := DmLTPAUserAttributeNameAndValue{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.LtpaAttributes, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmLTPAUserAttributeNameAndValueObjectType}, l)
		} else {
			data.LtpaAttributes = types.ListNull(types.ObjectType{AttrTypes: DmLTPAUserAttributeNameAndValueObjectType})
		}
	} else {
		data.LtpaAttributes = types.ListNull(types.ObjectType{AttrTypes: DmLTPAUserAttributeNameAndValueObjectType})
	}
	if value := res.Get(pathRoot + `TransactionPriority`); value.Exists() {
		l := []DmAAATransactionPriority{}
		if value := res.Get(`TransactionPriority`); value.Exists() {
			for _, v := range value.Array() {
				item := DmAAATransactionPriority{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.TransactionPriority, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmAAATransactionPriorityObjectType}, l)
		} else {
			data.TransactionPriority = types.ListNull(types.ObjectType{AttrTypes: DmAAATransactionPriorityObjectType})
		}
	} else {
		data.TransactionPriority = types.ListNull(types.ObjectType{AttrTypes: DmAAATransactionPriorityObjectType})
	}
	if value := res.Get(pathRoot + `SAMLValcred`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SamlValcred = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SamlValcred = types.StringNull()
	}
	if value := res.Get(pathRoot + `SAMLSigningKey`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SamlSigningKey = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SamlSigningKey = types.StringNull()
	}
	if value := res.Get(pathRoot + `SAMLSigningCert`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SamlSigningCert = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SamlSigningCert = types.StringNull()
	}
	if value := res.Get(pathRoot + `SAMLSigningHashAlg`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SamlSigningHashAlg = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SamlSigningHashAlg = types.StringNull()
	}
	if value := res.Get(pathRoot + `SAMLSigningAlg`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SamlSigningAlg = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SamlSigningAlg = types.StringNull()
	}
	if value := res.Get(pathRoot + `LDAPsuffix`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.LdapSuffix = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LdapSuffix = types.StringNull()
	}
	if value := res.Get(pathRoot + `LogAllowed`); value.Exists() {
		data.LogAllowed = tfutils.BoolFromString(value.String())
	} else {
		data.LogAllowed = types.BoolNull()
	}
	if value := res.Get(pathRoot + `LogAllowedLevel`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.LogAllowedLevel = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LogAllowedLevel = types.StringValue("info")
	}
	if value := res.Get(pathRoot + `LogRejected`); value.Exists() {
		data.LogRejected = tfutils.BoolFromString(value.String())
	} else {
		data.LogRejected = types.BoolNull()
	}
	if value := res.Get(pathRoot + `LogRejectedLevel`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.LogRejectedLevel = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LogRejectedLevel = types.StringValue("warn")
	}
	if value := res.Get(pathRoot + `WSSecureConversationCryptoKey`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.WsSecureConversationCryptoKey = tfutils.ParseStringFromGJSON(value)
	} else {
		data.WsSecureConversationCryptoKey = types.StringNull()
	}
	if value := res.Get(pathRoot + `SAMLSourceIDMappingFile`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SamlSourceIdMappingFile = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SamlSourceIdMappingFile = types.StringNull()
	}
	if value := res.Get(pathRoot + `PingIdentityCompatibility`); value.Exists() {
		data.PingIdentityCompatibility = tfutils.BoolFromString(value.String())
	} else {
		data.PingIdentityCompatibility = types.BoolNull()
	}
	if value := res.Get(pathRoot + `SAML2MetadataFile`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Saml2MetadataFile = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Saml2MetadataFile = types.StringNull()
	}
	if value := res.Get(pathRoot + `DoSValve`); value.Exists() {
		data.DoSValve = types.Int64Value(value.Int())
	} else {
		data.DoSValve = types.Int64Value(3)
	}
	if value := res.Get(pathRoot + `LDAPVersion`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.LdapVersion = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LdapVersion = types.StringValue("v2")
	}
	if value := res.Get(pathRoot + `EnforceSOAPActor`); value.Exists() {
		data.EnforceSoapActor = tfutils.BoolFromString(value.String())
	} else {
		data.EnforceSoapActor = types.BoolNull()
	}
	if value := res.Get(pathRoot + `WSSecActorRoleID`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.WsSecActorRoleId = tfutils.ParseStringFromGJSON(value)
	} else {
		data.WsSecActorRoleId = types.StringNull()
	}
	if value := res.Get(pathRoot + `AUSMHTTPHeader`); value.Exists() {
		data.AuSmHttpHeader = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.AuSmHttpHeader = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `AZSMHTTPHeader`); value.Exists() {
		data.AzSmHttpHeader = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.AzSmHttpHeader = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `DynConfig`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.DynConfig = tfutils.ParseStringFromGJSON(value)
	} else {
		data.DynConfig = types.StringValue("none")
	}
	if value := res.Get(pathRoot + `ExternalAAATemplate`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ExternalAaaTemplate = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ExternalAaaTemplate = types.StringNull()
	}
	if value := res.Get(pathRoot + `DynConfigCustomURL`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.DynConfigCustomUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.DynConfigCustomUrl = types.StringNull()
	}
}

func (data *AAAPolicy) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `AuthorizedCounter`); value.Exists() && !data.AuthorizedCounter.IsNull() {
		data.AuthorizedCounter = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AuthorizedCounter = types.StringNull()
	}
	if value := res.Get(pathRoot + `RejectedCounter`); value.Exists() && !data.RejectedCounter.IsNull() {
		data.RejectedCounter = tfutils.ParseStringFromGJSON(value)
	} else {
		data.RejectedCounter = types.StringNull()
	}
	if value := res.Get(pathRoot + `NamespaceMapping`); value.Exists() && !data.NamespaceMapping.IsNull() {
		l := []DmNamespaceMapping{}
		for _, v := range value.Array() {
			item := DmNamespaceMapping{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
		if len(l) > 0 {
			data.NamespaceMapping, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmNamespaceMappingObjectType}, l)
		} else {
			data.NamespaceMapping = types.ListNull(types.ObjectType{AttrTypes: DmNamespaceMappingObjectType})
		}
	} else {
		data.NamespaceMapping = types.ListNull(types.ObjectType{AttrTypes: DmNamespaceMappingObjectType})
	}
	if value := res.Get(pathRoot + `ExtractIdentity`); value.Exists() {
		data.ExtractIdentity.UpdateFromBody(ctx, "", value)
	} else {
		data.ExtractIdentity = nil
	}
	if value := res.Get(pathRoot + `Authenticate`); value.Exists() {
		data.Authenticate.UpdateFromBody(ctx, "", value)
	} else {
		data.Authenticate = nil
	}
	if value := res.Get(pathRoot + `MapCredentials`); value.Exists() {
		data.MapCredentials.UpdateFromBody(ctx, "", value)
	} else {
		data.MapCredentials = nil
	}
	if value := res.Get(pathRoot + `ExtractResource`); value.Exists() {
		data.ExtractResource.UpdateFromBody(ctx, "", value)
	} else {
		data.ExtractResource = nil
	}
	if value := res.Get(pathRoot + `MapResource`); value.Exists() {
		data.MapResource.UpdateFromBody(ctx, "", value)
	} else {
		data.MapResource = nil
	}
	if value := res.Get(pathRoot + `Authorize`); value.Exists() {
		data.Authorize.UpdateFromBody(ctx, "", value)
	} else {
		data.Authorize = nil
	}
	if value := res.Get(pathRoot + `PostProcess`); value.Exists() {
		data.PostProcess.UpdateFromBody(ctx, "", value)
	} else {
		data.PostProcess = nil
	}
	if value := res.Get(pathRoot + `SAMLAttribute`); value.Exists() && !data.SamlAttribute.IsNull() {
		l := []DmSAMLAttributeNameAndValue{}
		for _, v := range value.Array() {
			item := DmSAMLAttributeNameAndValue{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
		if len(l) > 0 {
			data.SamlAttribute, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmSAMLAttributeNameAndValueObjectType}, l)
		} else {
			data.SamlAttribute = types.ListNull(types.ObjectType{AttrTypes: DmSAMLAttributeNameAndValueObjectType})
		}
	} else {
		data.SamlAttribute = types.ListNull(types.ObjectType{AttrTypes: DmSAMLAttributeNameAndValueObjectType})
	}
	if value := res.Get(pathRoot + `LTPAAttributes`); value.Exists() && !data.LtpaAttributes.IsNull() {
		l := []DmLTPAUserAttributeNameAndValue{}
		for _, v := range value.Array() {
			item := DmLTPAUserAttributeNameAndValue{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
		if len(l) > 0 {
			data.LtpaAttributes, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmLTPAUserAttributeNameAndValueObjectType}, l)
		} else {
			data.LtpaAttributes = types.ListNull(types.ObjectType{AttrTypes: DmLTPAUserAttributeNameAndValueObjectType})
		}
	} else {
		data.LtpaAttributes = types.ListNull(types.ObjectType{AttrTypes: DmLTPAUserAttributeNameAndValueObjectType})
	}
	if value := res.Get(pathRoot + `TransactionPriority`); value.Exists() && !data.TransactionPriority.IsNull() {
		l := []DmAAATransactionPriority{}
		for _, v := range value.Array() {
			item := DmAAATransactionPriority{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
		if len(l) > 0 {
			data.TransactionPriority, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmAAATransactionPriorityObjectType}, l)
		} else {
			data.TransactionPriority = types.ListNull(types.ObjectType{AttrTypes: DmAAATransactionPriorityObjectType})
		}
	} else {
		data.TransactionPriority = types.ListNull(types.ObjectType{AttrTypes: DmAAATransactionPriorityObjectType})
	}
	if value := res.Get(pathRoot + `SAMLValcred`); value.Exists() && !data.SamlValcred.IsNull() {
		data.SamlValcred = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SamlValcred = types.StringNull()
	}
	if value := res.Get(pathRoot + `SAMLSigningKey`); value.Exists() && !data.SamlSigningKey.IsNull() {
		data.SamlSigningKey = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SamlSigningKey = types.StringNull()
	}
	if value := res.Get(pathRoot + `SAMLSigningCert`); value.Exists() && !data.SamlSigningCert.IsNull() {
		data.SamlSigningCert = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SamlSigningCert = types.StringNull()
	}
	if value := res.Get(pathRoot + `SAMLSigningHashAlg`); value.Exists() && !data.SamlSigningHashAlg.IsNull() {
		data.SamlSigningHashAlg = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SamlSigningHashAlg = types.StringNull()
	}
	if value := res.Get(pathRoot + `SAMLSigningAlg`); value.Exists() && !data.SamlSigningAlg.IsNull() {
		data.SamlSigningAlg = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SamlSigningAlg = types.StringNull()
	}
	if value := res.Get(pathRoot + `LDAPsuffix`); value.Exists() && !data.LdapSuffix.IsNull() {
		data.LdapSuffix = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LdapSuffix = types.StringNull()
	}
	if value := res.Get(pathRoot + `LogAllowed`); value.Exists() && !data.LogAllowed.IsNull() {
		data.LogAllowed = tfutils.BoolFromString(value.String())
	} else if data.LogAllowed.ValueBool() {
		data.LogAllowed = types.BoolNull()
	}
	if value := res.Get(pathRoot + `LogAllowedLevel`); value.Exists() && !data.LogAllowedLevel.IsNull() {
		data.LogAllowedLevel = tfutils.ParseStringFromGJSON(value)
	} else if data.LogAllowedLevel.ValueString() != "info" {
		data.LogAllowedLevel = types.StringNull()
	}
	if value := res.Get(pathRoot + `LogRejected`); value.Exists() && !data.LogRejected.IsNull() {
		data.LogRejected = tfutils.BoolFromString(value.String())
	} else if !data.LogRejected.ValueBool() {
		data.LogRejected = types.BoolNull()
	}
	if value := res.Get(pathRoot + `LogRejectedLevel`); value.Exists() && !data.LogRejectedLevel.IsNull() {
		data.LogRejectedLevel = tfutils.ParseStringFromGJSON(value)
	} else if data.LogRejectedLevel.ValueString() != "warn" {
		data.LogRejectedLevel = types.StringNull()
	}
	if value := res.Get(pathRoot + `WSSecureConversationCryptoKey`); value.Exists() && !data.WsSecureConversationCryptoKey.IsNull() {
		data.WsSecureConversationCryptoKey = tfutils.ParseStringFromGJSON(value)
	} else {
		data.WsSecureConversationCryptoKey = types.StringNull()
	}
	if value := res.Get(pathRoot + `SAMLSourceIDMappingFile`); value.Exists() && !data.SamlSourceIdMappingFile.IsNull() {
		data.SamlSourceIdMappingFile = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SamlSourceIdMappingFile = types.StringNull()
	}
	if value := res.Get(pathRoot + `PingIdentityCompatibility`); value.Exists() && !data.PingIdentityCompatibility.IsNull() {
		data.PingIdentityCompatibility = tfutils.BoolFromString(value.String())
	} else if data.PingIdentityCompatibility.ValueBool() {
		data.PingIdentityCompatibility = types.BoolNull()
	}
	if value := res.Get(pathRoot + `SAML2MetadataFile`); value.Exists() && !data.Saml2MetadataFile.IsNull() {
		data.Saml2MetadataFile = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Saml2MetadataFile = types.StringNull()
	}
	if value := res.Get(pathRoot + `DoSValve`); value.Exists() && !data.DoSValve.IsNull() {
		data.DoSValve = types.Int64Value(value.Int())
	} else if data.DoSValve.ValueInt64() != 3 {
		data.DoSValve = types.Int64Null()
	}
	if value := res.Get(pathRoot + `LDAPVersion`); value.Exists() && !data.LdapVersion.IsNull() {
		data.LdapVersion = tfutils.ParseStringFromGJSON(value)
	} else if data.LdapVersion.ValueString() != "v2" {
		data.LdapVersion = types.StringNull()
	}
	if value := res.Get(pathRoot + `EnforceSOAPActor`); value.Exists() && !data.EnforceSoapActor.IsNull() {
		data.EnforceSoapActor = tfutils.BoolFromString(value.String())
	} else if !data.EnforceSoapActor.ValueBool() {
		data.EnforceSoapActor = types.BoolNull()
	}
	if value := res.Get(pathRoot + `WSSecActorRoleID`); value.Exists() && !data.WsSecActorRoleId.IsNull() {
		data.WsSecActorRoleId = tfutils.ParseStringFromGJSON(value)
	} else {
		data.WsSecActorRoleId = types.StringNull()
	}
	if value := res.Get(pathRoot + `AUSMHTTPHeader`); value.Exists() && !data.AuSmHttpHeader.IsNull() {
		data.AuSmHttpHeader = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.AuSmHttpHeader = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `AZSMHTTPHeader`); value.Exists() && !data.AzSmHttpHeader.IsNull() {
		data.AzSmHttpHeader = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.AzSmHttpHeader = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `DynConfig`); value.Exists() && !data.DynConfig.IsNull() {
		data.DynConfig = tfutils.ParseStringFromGJSON(value)
	} else if data.DynConfig.ValueString() != "none" {
		data.DynConfig = types.StringNull()
	}
	if value := res.Get(pathRoot + `ExternalAAATemplate`); value.Exists() && !data.ExternalAaaTemplate.IsNull() {
		data.ExternalAaaTemplate = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ExternalAaaTemplate = types.StringNull()
	}
	if value := res.Get(pathRoot + `DynConfigCustomURL`); value.Exists() && !data.DynConfigCustomUrl.IsNull() {
		data.DynConfigCustomUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.DynConfigCustomUrl = types.StringNull()
	}
}

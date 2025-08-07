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

type APIDefinition struct {
	Id                               types.String      `tfsdk:"id"`
	AppDomain                        types.String      `tfsdk:"app_domain"`
	UserSummary                      types.String      `tfsdk:"user_summary"`
	ApiId                            types.String      `tfsdk:"api_id"`
	Name                             types.String      `tfsdk:"name"`
	Version                          types.String      `tfsdk:"version"`
	BasePath                         types.String      `tfsdk:"base_path"`
	HtmlPage                         types.String      `tfsdk:"html_page"`
	Type                             types.String      `tfsdk:"type"`
	Assembly                         types.String      `tfsdk:"assembly"`
	Path                             types.List        `tfsdk:"path"`
	Consume                          types.List        `tfsdk:"consume"`
	Produce                          types.List        `tfsdk:"produce"`
	SwaggerLocation                  types.String      `tfsdk:"swagger_location"`
	GraphQlSchema                    types.String      `tfsdk:"graph_ql_schema"`
	WsdlAdvertisedSchemaLocation     types.String      `tfsdk:"wsdl_advertised_schema_location"`
	WsdlValidationSchema             types.String      `tfsdk:"wsdl_validation_schema"`
	SecurityRequirement              types.List        `tfsdk:"security_requirement"`
	RequireApiMutualTls              types.Bool        `tfsdk:"require_api_mutual_tls"`
	ApiMutualTlsSource               types.List        `tfsdk:"api_mutual_tls_source"`
	ApiMutualTlsHeaderName           types.String      `tfsdk:"api_mutual_tls_header_name"`
	Properties                       types.List        `tfsdk:"properties"`
	Schemas                          types.List        `tfsdk:"schemas"`
	CorsToggle                       types.Bool        `tfsdk:"cors_toggle"`
	CorsPolicy                       types.String      `tfsdk:"cors_policy"`
	ActivityLogToggle                types.Bool        `tfsdk:"activity_log_toggle"`
	Content                          types.String      `tfsdk:"content"`
	ErrorContent                     types.String      `tfsdk:"error_content"`
	PreservedRequestHeader           types.List        `tfsdk:"preserved_request_header"`
	PreservedResponseHeader          types.List        `tfsdk:"preserved_response_header"`
	MessageBuffering                 types.Bool        `tfsdk:"message_buffering"`
	DeploymentState                  types.String      `tfsdk:"deployment_state"`
	ShareRateLimitCount              types.String      `tfsdk:"share_rate_limit_count"`
	ReturnV5Responses                types.Bool        `tfsdk:"return_v5_responses"`
	CopyIdHeadersToMessage           types.Bool        `tfsdk:"copy_id_headers_to_message"`
	EnforceRequiredParams            types.Bool        `tfsdk:"enforce_required_params"`
	AllowChunkedUploads              types.Bool        `tfsdk:"allow_chunked_uploads"`
	SetV5RequestHeaders              types.Bool        `tfsdk:"set_v5_request_headers"`
	GetRawBodyValue                  types.Bool        `tfsdk:"get_raw_body_value"`
	AllowedApiProtocols              *DmAPIProtocols   `tfsdk:"allowed_api_protocols"`
	AllowTrailingSlash               types.Bool        `tfsdk:"allow_trailing_slash"`
	EnforceAllHeadersCaseInsensitive types.Bool        `tfsdk:"enforce_all_headers_case_insensitive"`
	EnforceFormDataParameter         types.Bool        `tfsdk:"enforce_form_data_parameter"`
	ForceHttp500ForSoap11            types.Bool        `tfsdk:"force_http500_for_soap11"`
	DependencyActions                []*actions.Action `tfsdk:"dependency_actions"`
}

var APIDefinitionObjectType = map[string]attr.Type{
	"id":                                   types.StringType,
	"app_domain":                           types.StringType,
	"user_summary":                         types.StringType,
	"api_id":                               types.StringType,
	"name":                                 types.StringType,
	"version":                              types.StringType,
	"base_path":                            types.StringType,
	"html_page":                            types.StringType,
	"type":                                 types.StringType,
	"assembly":                             types.StringType,
	"path":                                 types.ListType{ElemType: types.StringType},
	"consume":                              types.ListType{ElemType: types.StringType},
	"produce":                              types.ListType{ElemType: types.StringType},
	"swagger_location":                     types.StringType,
	"graph_ql_schema":                      types.StringType,
	"wsdl_advertised_schema_location":      types.StringType,
	"wsdl_validation_schema":               types.StringType,
	"security_requirement":                 types.ListType{ElemType: types.StringType},
	"require_api_mutual_tls":               types.BoolType,
	"api_mutual_tls_source":                types.ListType{ElemType: types.StringType},
	"api_mutual_tls_header_name":           types.StringType,
	"properties":                           types.ListType{ElemType: types.ObjectType{AttrTypes: DmAPIPropertyObjectType}},
	"schemas":                              types.ListType{ElemType: types.ObjectType{AttrTypes: DmAPIDataTypeDefinitionObjectType}},
	"cors_toggle":                          types.BoolType,
	"cors_policy":                          types.StringType,
	"activity_log_toggle":                  types.BoolType,
	"content":                              types.StringType,
	"error_content":                        types.StringType,
	"preserved_request_header":             types.ListType{ElemType: types.StringType},
	"preserved_response_header":            types.ListType{ElemType: types.StringType},
	"message_buffering":                    types.BoolType,
	"deployment_state":                     types.StringType,
	"share_rate_limit_count":               types.StringType,
	"return_v5_responses":                  types.BoolType,
	"copy_id_headers_to_message":           types.BoolType,
	"enforce_required_params":              types.BoolType,
	"allow_chunked_uploads":                types.BoolType,
	"set_v5_request_headers":               types.BoolType,
	"get_raw_body_value":                   types.BoolType,
	"allowed_api_protocols":                types.ObjectType{AttrTypes: DmAPIProtocolsObjectType},
	"allow_trailing_slash":                 types.BoolType,
	"enforce_all_headers_case_insensitive": types.BoolType,
	"enforce_form_data_parameter":          types.BoolType,
	"force_http500_for_soap11":             types.BoolType,
	"dependency_actions":                   actions.ActionsListType,
}

func (data APIDefinition) GetPath() string {
	rest_path := "/mgmt/config/{domain}/APIDefinition"
	rest_path = strings.ReplaceAll(rest_path, "{name}", url.QueryEscape(data.Id.ValueString()))
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data APIDefinition) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.ApiId.IsNull() {
		return false
	}
	if !data.Name.IsNull() {
		return false
	}
	if !data.Version.IsNull() {
		return false
	}
	if !data.BasePath.IsNull() {
		return false
	}
	if !data.HtmlPage.IsNull() {
		return false
	}
	if !data.Type.IsNull() {
		return false
	}
	if !data.Assembly.IsNull() {
		return false
	}
	if !data.Path.IsNull() {
		return false
	}
	if !data.Consume.IsNull() {
		return false
	}
	if !data.Produce.IsNull() {
		return false
	}
	if !data.SwaggerLocation.IsNull() {
		return false
	}
	if !data.GraphQlSchema.IsNull() {
		return false
	}
	if !data.WsdlAdvertisedSchemaLocation.IsNull() {
		return false
	}
	if !data.WsdlValidationSchema.IsNull() {
		return false
	}
	if !data.SecurityRequirement.IsNull() {
		return false
	}
	if !data.RequireApiMutualTls.IsNull() {
		return false
	}
	if !data.ApiMutualTlsSource.IsNull() {
		return false
	}
	if !data.ApiMutualTlsHeaderName.IsNull() {
		return false
	}
	if !data.Properties.IsNull() {
		return false
	}
	if !data.Schemas.IsNull() {
		return false
	}
	if !data.CorsToggle.IsNull() {
		return false
	}
	if !data.CorsPolicy.IsNull() {
		return false
	}
	if !data.ActivityLogToggle.IsNull() {
		return false
	}
	if !data.Content.IsNull() {
		return false
	}
	if !data.ErrorContent.IsNull() {
		return false
	}
	if !data.PreservedRequestHeader.IsNull() {
		return false
	}
	if !data.PreservedResponseHeader.IsNull() {
		return false
	}
	if !data.MessageBuffering.IsNull() {
		return false
	}
	if !data.DeploymentState.IsNull() {
		return false
	}
	if !data.ShareRateLimitCount.IsNull() {
		return false
	}
	if !data.ReturnV5Responses.IsNull() {
		return false
	}
	if !data.CopyIdHeadersToMessage.IsNull() {
		return false
	}
	if !data.EnforceRequiredParams.IsNull() {
		return false
	}
	if !data.AllowChunkedUploads.IsNull() {
		return false
	}
	if !data.SetV5RequestHeaders.IsNull() {
		return false
	}
	if !data.GetRawBodyValue.IsNull() {
		return false
	}
	if data.AllowedApiProtocols != nil {
		if !data.AllowedApiProtocols.IsNull() {
			return false
		}
	}
	if !data.AllowTrailingSlash.IsNull() {
		return false
	}
	if !data.EnforceAllHeadersCaseInsensitive.IsNull() {
		return false
	}
	if !data.EnforceFormDataParameter.IsNull() {
		return false
	}
	if !data.ForceHttp500ForSoap11.IsNull() {
		return false
	}
	return true
}

func (data APIDefinition) ToBody(ctx context.Context, pathRoot string) string {
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
	if !data.ApiId.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ID`, data.ApiId.ValueString())
	}
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Name`, data.Name.ValueString())
	}
	if !data.Version.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Version`, data.Version.ValueString())
	}
	if !data.BasePath.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`BasePath`, data.BasePath.ValueString())
	}
	if !data.HtmlPage.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`HtmlPage`, data.HtmlPage.ValueString())
	}
	if !data.Type.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Type`, data.Type.ValueString())
	}
	if !data.Assembly.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Assembly`, data.Assembly.ValueString())
	}
	if !data.Path.IsNull() {
		var values []string
		data.Path.ElementsAs(ctx, &values, false)
		for _, val := range values {
			body, _ = sjson.Set(body, pathRoot+`Path`+".-1", map[string]string{"value": val})
		}
	}
	if !data.Consume.IsNull() {
		var values []string
		data.Consume.ElementsAs(ctx, &values, false)
		for _, val := range values {
			body, _ = sjson.Set(body, pathRoot+`Consume`+".-1", map[string]string{"value": val})
		}
	}
	if !data.Produce.IsNull() {
		var values []string
		data.Produce.ElementsAs(ctx, &values, false)
		for _, val := range values {
			body, _ = sjson.Set(body, pathRoot+`Produce`+".-1", map[string]string{"value": val})
		}
	}
	if !data.SwaggerLocation.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SwaggerLocation`, data.SwaggerLocation.ValueString())
	}
	if !data.GraphQlSchema.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`GraphQLSchema`, data.GraphQlSchema.ValueString())
	}
	if !data.WsdlAdvertisedSchemaLocation.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`WsdlAdvertisedSchemaLocation`, data.WsdlAdvertisedSchemaLocation.ValueString())
	}
	if !data.WsdlValidationSchema.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`WsdlValidationSchema`, data.WsdlValidationSchema.ValueString())
	}
	if !data.SecurityRequirement.IsNull() {
		var values []string
		data.SecurityRequirement.ElementsAs(ctx, &values, false)
		for _, val := range values {
			body, _ = sjson.Set(body, pathRoot+`SecurityRequirement`+".-1", map[string]string{"value": val})
		}
	}
	if !data.RequireApiMutualTls.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`RequireAPIMutualTLS`, tfutils.StringFromBool(data.RequireApiMutualTls, ""))
	}
	if !data.ApiMutualTlsSource.IsNull() {
		var values []string
		data.ApiMutualTlsSource.ElementsAs(ctx, &values, false)
		for _, val := range values {
			body, _ = sjson.Set(body, pathRoot+`APIMutualTLSSource`+".-1", map[string]string{"value": val})
		}
	}
	if !data.ApiMutualTlsHeaderName.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`APIMutualTLSHeaderName`, data.ApiMutualTlsHeaderName.ValueString())
	}
	if !data.Properties.IsNull() {
		var values []DmAPIProperty
		data.Properties.ElementsAs(ctx, &values, false)
		for _, val := range values {
			body, _ = sjson.SetRaw(body, pathRoot+`Properties`+".-1", val.ToBody(ctx, ""))
		}
	}
	if !data.Schemas.IsNull() {
		var values []DmAPIDataTypeDefinition
		data.Schemas.ElementsAs(ctx, &values, false)
		for _, val := range values {
			body, _ = sjson.SetRaw(body, pathRoot+`Schemas`+".-1", val.ToBody(ctx, ""))
		}
	}
	if !data.CorsToggle.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`CORSToggle`, tfutils.StringFromBool(data.CorsToggle, ""))
	}
	if !data.CorsPolicy.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`CORSPolicy`, data.CorsPolicy.ValueString())
	}
	if !data.ActivityLogToggle.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ActivityLogToggle`, tfutils.StringFromBool(data.ActivityLogToggle, ""))
	}
	if !data.Content.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Content`, data.Content.ValueString())
	}
	if !data.ErrorContent.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ErrorContent`, data.ErrorContent.ValueString())
	}
	if !data.PreservedRequestHeader.IsNull() {
		var values []string
		data.PreservedRequestHeader.ElementsAs(ctx, &values, false)
		for _, val := range values {
			body, _ = sjson.Set(body, pathRoot+`PreservedRequestHeader`+".-1", map[string]string{"value": val})
		}
	}
	if !data.PreservedResponseHeader.IsNull() {
		var values []string
		data.PreservedResponseHeader.ElementsAs(ctx, &values, false)
		for _, val := range values {
			body, _ = sjson.Set(body, pathRoot+`PreservedResponseHeader`+".-1", map[string]string{"value": val})
		}
	}
	if !data.MessageBuffering.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MessageBuffering`, tfutils.StringFromBool(data.MessageBuffering, ""))
	}
	if !data.DeploymentState.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DeploymentState`, data.DeploymentState.ValueString())
	}
	if !data.ShareRateLimitCount.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ShareRateLimitCount`, data.ShareRateLimitCount.ValueString())
	}
	if !data.ReturnV5Responses.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ReturnV5Responses`, tfutils.StringFromBool(data.ReturnV5Responses, ""))
	}
	if !data.CopyIdHeadersToMessage.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`CopyIdHeadersToMessage`, tfutils.StringFromBool(data.CopyIdHeadersToMessage, ""))
	}
	if !data.EnforceRequiredParams.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EnforceRequiredParams`, tfutils.StringFromBool(data.EnforceRequiredParams, ""))
	}
	if !data.AllowChunkedUploads.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AllowChunkedUploads`, tfutils.StringFromBool(data.AllowChunkedUploads, ""))
	}
	if !data.SetV5RequestHeaders.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SetV5RequestHeaders`, tfutils.StringFromBool(data.SetV5RequestHeaders, ""))
	}
	if !data.GetRawBodyValue.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`GetRawBodyValue`, tfutils.StringFromBool(data.GetRawBodyValue, ""))
	}
	if data.AllowedApiProtocols != nil {
		if !data.AllowedApiProtocols.IsNull() {
			body, _ = sjson.SetRaw(body, pathRoot+`AllowedAPIProtocols`, data.AllowedApiProtocols.ToBody(ctx, ""))
		}
	}
	if !data.AllowTrailingSlash.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AllowTrailingSlash`, tfutils.StringFromBool(data.AllowTrailingSlash, ""))
	}
	if !data.EnforceAllHeadersCaseInsensitive.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EnforceAllHeadersCaseInsensitive`, tfutils.StringFromBool(data.EnforceAllHeadersCaseInsensitive, ""))
	}
	if !data.EnforceFormDataParameter.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EnforceFormDataParameter`, tfutils.StringFromBool(data.EnforceFormDataParameter, ""))
	}
	if !data.ForceHttp500ForSoap11.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ForceHttp500ForSoap11`, tfutils.StringFromBool(data.ForceHttp500ForSoap11, ""))
	}
	return body
}

func (data *APIDefinition) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `ID`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ApiId = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ApiId = types.StringNull()
	}
	if value := res.Get(pathRoot + `Name`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Name = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get(pathRoot + `Version`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Version = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Version = types.StringValue("1.0.0")
	}
	if value := res.Get(pathRoot + `BasePath`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.BasePath = tfutils.ParseStringFromGJSON(value)
	} else {
		data.BasePath = types.StringValue("/")
	}
	if value := res.Get(pathRoot + `HtmlPage`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.HtmlPage = tfutils.ParseStringFromGJSON(value)
	} else {
		data.HtmlPage = types.StringNull()
	}
	if value := res.Get(pathRoot + `Type`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Type = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Type = types.StringValue("standard")
	}
	if value := res.Get(pathRoot + `Assembly`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Assembly = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Assembly = types.StringNull()
	}
	if value := res.Get(pathRoot + `Path`); value.Exists() {
		data.Path = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.Path = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `Consume`); value.Exists() {
		data.Consume = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.Consume = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `Produce`); value.Exists() {
		data.Produce = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.Produce = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `SwaggerLocation`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SwaggerLocation = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SwaggerLocation = types.StringNull()
	}
	if value := res.Get(pathRoot + `GraphQLSchema`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.GraphQlSchema = tfutils.ParseStringFromGJSON(value)
	} else {
		data.GraphQlSchema = types.StringNull()
	}
	if value := res.Get(pathRoot + `WsdlAdvertisedSchemaLocation`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.WsdlAdvertisedSchemaLocation = tfutils.ParseStringFromGJSON(value)
	} else {
		data.WsdlAdvertisedSchemaLocation = types.StringNull()
	}
	if value := res.Get(pathRoot + `WsdlValidationSchema`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.WsdlValidationSchema = tfutils.ParseStringFromGJSON(value)
	} else {
		data.WsdlValidationSchema = types.StringNull()
	}
	if value := res.Get(pathRoot + `SecurityRequirement`); value.Exists() {
		data.SecurityRequirement = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.SecurityRequirement = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `RequireAPIMutualTLS`); value.Exists() {
		data.RequireApiMutualTls = tfutils.BoolFromString(value.String())
	} else {
		data.RequireApiMutualTls = types.BoolNull()
	}
	if value := res.Get(pathRoot + `APIMutualTLSSource`); value.Exists() {
		data.ApiMutualTlsSource = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.ApiMutualTlsSource = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `APIMutualTLSHeaderName`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ApiMutualTlsHeaderName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ApiMutualTlsHeaderName = types.StringValue("X-Client-Certificate")
	}
	if value := res.Get(pathRoot + `Properties`); value.Exists() {
		l := []DmAPIProperty{}
		if value := res.Get(`Properties`); value.Exists() {
			for _, v := range value.Array() {
				item := DmAPIProperty{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.Properties, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmAPIPropertyObjectType}, l)
		} else {
			data.Properties = types.ListNull(types.ObjectType{AttrTypes: DmAPIPropertyObjectType})
		}
	} else {
		data.Properties = types.ListNull(types.ObjectType{AttrTypes: DmAPIPropertyObjectType})
	}
	if value := res.Get(pathRoot + `Schemas`); value.Exists() {
		l := []DmAPIDataTypeDefinition{}
		if value := res.Get(`Schemas`); value.Exists() {
			for _, v := range value.Array() {
				item := DmAPIDataTypeDefinition{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.Schemas, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmAPIDataTypeDefinitionObjectType}, l)
		} else {
			data.Schemas = types.ListNull(types.ObjectType{AttrTypes: DmAPIDataTypeDefinitionObjectType})
		}
	} else {
		data.Schemas = types.ListNull(types.ObjectType{AttrTypes: DmAPIDataTypeDefinitionObjectType})
	}
	if value := res.Get(pathRoot + `CORSToggle`); value.Exists() {
		data.CorsToggle = tfutils.BoolFromString(value.String())
	} else {
		data.CorsToggle = types.BoolNull()
	}
	if value := res.Get(pathRoot + `CORSPolicy`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.CorsPolicy = tfutils.ParseStringFromGJSON(value)
	} else {
		data.CorsPolicy = types.StringNull()
	}
	if value := res.Get(pathRoot + `ActivityLogToggle`); value.Exists() {
		data.ActivityLogToggle = tfutils.BoolFromString(value.String())
	} else {
		data.ActivityLogToggle = types.BoolNull()
	}
	if value := res.Get(pathRoot + `Content`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Content = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Content = types.StringValue("activity")
	}
	if value := res.Get(pathRoot + `ErrorContent`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ErrorContent = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ErrorContent = types.StringValue("payload")
	}
	if value := res.Get(pathRoot + `PreservedRequestHeader`); value.Exists() {
		data.PreservedRequestHeader = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.PreservedRequestHeader = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `PreservedResponseHeader`); value.Exists() {
		data.PreservedResponseHeader = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.PreservedResponseHeader = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `MessageBuffering`); value.Exists() {
		data.MessageBuffering = tfutils.BoolFromString(value.String())
	} else {
		data.MessageBuffering = types.BoolNull()
	}
	if value := res.Get(pathRoot + `DeploymentState`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.DeploymentState = tfutils.ParseStringFromGJSON(value)
	} else {
		data.DeploymentState = types.StringValue("running")
	}
	if value := res.Get(pathRoot + `ShareRateLimitCount`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ShareRateLimitCount = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ShareRateLimitCount = types.StringValue("unset")
	}
	if value := res.Get(pathRoot + `ReturnV5Responses`); value.Exists() {
		data.ReturnV5Responses = tfutils.BoolFromString(value.String())
	} else {
		data.ReturnV5Responses = types.BoolNull()
	}
	if value := res.Get(pathRoot + `CopyIdHeadersToMessage`); value.Exists() {
		data.CopyIdHeadersToMessage = tfutils.BoolFromString(value.String())
	} else {
		data.CopyIdHeadersToMessage = types.BoolNull()
	}
	if value := res.Get(pathRoot + `EnforceRequiredParams`); value.Exists() {
		data.EnforceRequiredParams = tfutils.BoolFromString(value.String())
	} else {
		data.EnforceRequiredParams = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AllowChunkedUploads`); value.Exists() {
		data.AllowChunkedUploads = tfutils.BoolFromString(value.String())
	} else {
		data.AllowChunkedUploads = types.BoolNull()
	}
	if value := res.Get(pathRoot + `SetV5RequestHeaders`); value.Exists() {
		data.SetV5RequestHeaders = tfutils.BoolFromString(value.String())
	} else {
		data.SetV5RequestHeaders = types.BoolNull()
	}
	if value := res.Get(pathRoot + `GetRawBodyValue`); value.Exists() {
		data.GetRawBodyValue = tfutils.BoolFromString(value.String())
	} else {
		data.GetRawBodyValue = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AllowedAPIProtocols`); value.Exists() {
		data.AllowedApiProtocols = &DmAPIProtocols{}
		data.AllowedApiProtocols.FromBody(ctx, "", value)
	} else {
		data.AllowedApiProtocols = nil
	}
	if value := res.Get(pathRoot + `AllowTrailingSlash`); value.Exists() {
		data.AllowTrailingSlash = tfutils.BoolFromString(value.String())
	} else {
		data.AllowTrailingSlash = types.BoolNull()
	}
	if value := res.Get(pathRoot + `EnforceAllHeadersCaseInsensitive`); value.Exists() {
		data.EnforceAllHeadersCaseInsensitive = tfutils.BoolFromString(value.String())
	} else {
		data.EnforceAllHeadersCaseInsensitive = types.BoolNull()
	}
	if value := res.Get(pathRoot + `EnforceFormDataParameter`); value.Exists() {
		data.EnforceFormDataParameter = tfutils.BoolFromString(value.String())
	} else {
		data.EnforceFormDataParameter = types.BoolNull()
	}
	if value := res.Get(pathRoot + `ForceHttp500ForSoap11`); value.Exists() {
		data.ForceHttp500ForSoap11 = tfutils.BoolFromString(value.String())
	} else {
		data.ForceHttp500ForSoap11 = types.BoolNull()
	}
}

func (data *APIDefinition) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `ID`); value.Exists() && !data.ApiId.IsNull() {
		data.ApiId = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ApiId = types.StringNull()
	}
	if value := res.Get(pathRoot + `Name`); value.Exists() && !data.Name.IsNull() {
		data.Name = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get(pathRoot + `Version`); value.Exists() && !data.Version.IsNull() {
		data.Version = tfutils.ParseStringFromGJSON(value)
	} else if data.Version.ValueString() != "1.0.0" {
		data.Version = types.StringNull()
	}
	if value := res.Get(pathRoot + `BasePath`); value.Exists() && !data.BasePath.IsNull() {
		data.BasePath = tfutils.ParseStringFromGJSON(value)
	} else if data.BasePath.ValueString() != "/" {
		data.BasePath = types.StringNull()
	}
	if value := res.Get(pathRoot + `HtmlPage`); value.Exists() && !data.HtmlPage.IsNull() {
		data.HtmlPage = tfutils.ParseStringFromGJSON(value)
	} else {
		data.HtmlPage = types.StringNull()
	}
	if value := res.Get(pathRoot + `Type`); value.Exists() && !data.Type.IsNull() {
		data.Type = tfutils.ParseStringFromGJSON(value)
	} else if data.Type.ValueString() != "standard" {
		data.Type = types.StringNull()
	}
	if value := res.Get(pathRoot + `Assembly`); value.Exists() && !data.Assembly.IsNull() {
		data.Assembly = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Assembly = types.StringNull()
	}
	if value := res.Get(pathRoot + `Path`); value.Exists() && !data.Path.IsNull() {
		data.Path = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.Path = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `Consume`); value.Exists() && !data.Consume.IsNull() {
		data.Consume = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.Consume = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `Produce`); value.Exists() && !data.Produce.IsNull() {
		data.Produce = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.Produce = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `SwaggerLocation`); value.Exists() && !data.SwaggerLocation.IsNull() {
		data.SwaggerLocation = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SwaggerLocation = types.StringNull()
	}
	if value := res.Get(pathRoot + `GraphQLSchema`); value.Exists() && !data.GraphQlSchema.IsNull() {
		data.GraphQlSchema = tfutils.ParseStringFromGJSON(value)
	} else {
		data.GraphQlSchema = types.StringNull()
	}
	if value := res.Get(pathRoot + `WsdlAdvertisedSchemaLocation`); value.Exists() && !data.WsdlAdvertisedSchemaLocation.IsNull() {
		data.WsdlAdvertisedSchemaLocation = tfutils.ParseStringFromGJSON(value)
	} else {
		data.WsdlAdvertisedSchemaLocation = types.StringNull()
	}
	if value := res.Get(pathRoot + `WsdlValidationSchema`); value.Exists() && !data.WsdlValidationSchema.IsNull() {
		data.WsdlValidationSchema = tfutils.ParseStringFromGJSON(value)
	} else {
		data.WsdlValidationSchema = types.StringNull()
	}
	if value := res.Get(pathRoot + `SecurityRequirement`); value.Exists() && !data.SecurityRequirement.IsNull() {
		data.SecurityRequirement = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.SecurityRequirement = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `RequireAPIMutualTLS`); value.Exists() && !data.RequireApiMutualTls.IsNull() {
		data.RequireApiMutualTls = tfutils.BoolFromString(value.String())
	} else if data.RequireApiMutualTls.ValueBool() {
		data.RequireApiMutualTls = types.BoolNull()
	}
	if value := res.Get(pathRoot + `APIMutualTLSSource`); value.Exists() && !data.ApiMutualTlsSource.IsNull() {
		data.ApiMutualTlsSource = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.ApiMutualTlsSource = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `APIMutualTLSHeaderName`); value.Exists() && !data.ApiMutualTlsHeaderName.IsNull() {
		data.ApiMutualTlsHeaderName = tfutils.ParseStringFromGJSON(value)
	} else if data.ApiMutualTlsHeaderName.ValueString() != "X-Client-Certificate" {
		data.ApiMutualTlsHeaderName = types.StringNull()
	}
	if value := res.Get(pathRoot + `Properties`); value.Exists() && !data.Properties.IsNull() {
		l := []DmAPIProperty{}
		for _, v := range value.Array() {
			item := DmAPIProperty{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
		if len(l) > 0 {
			data.Properties, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmAPIPropertyObjectType}, l)
		} else {
			data.Properties = types.ListNull(types.ObjectType{AttrTypes: DmAPIPropertyObjectType})
		}
	} else {
		data.Properties = types.ListNull(types.ObjectType{AttrTypes: DmAPIPropertyObjectType})
	}
	if value := res.Get(pathRoot + `Schemas`); value.Exists() && !data.Schemas.IsNull() {
		l := []DmAPIDataTypeDefinition{}
		for _, v := range value.Array() {
			item := DmAPIDataTypeDefinition{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
		if len(l) > 0 {
			data.Schemas, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmAPIDataTypeDefinitionObjectType}, l)
		} else {
			data.Schemas = types.ListNull(types.ObjectType{AttrTypes: DmAPIDataTypeDefinitionObjectType})
		}
	} else {
		data.Schemas = types.ListNull(types.ObjectType{AttrTypes: DmAPIDataTypeDefinitionObjectType})
	}
	if value := res.Get(pathRoot + `CORSToggle`); value.Exists() && !data.CorsToggle.IsNull() {
		data.CorsToggle = tfutils.BoolFromString(value.String())
	} else if data.CorsToggle.ValueBool() {
		data.CorsToggle = types.BoolNull()
	}
	if value := res.Get(pathRoot + `CORSPolicy`); value.Exists() && !data.CorsPolicy.IsNull() {
		data.CorsPolicy = tfutils.ParseStringFromGJSON(value)
	} else {
		data.CorsPolicy = types.StringNull()
	}
	if value := res.Get(pathRoot + `ActivityLogToggle`); value.Exists() && !data.ActivityLogToggle.IsNull() {
		data.ActivityLogToggle = tfutils.BoolFromString(value.String())
	} else if data.ActivityLogToggle.ValueBool() {
		data.ActivityLogToggle = types.BoolNull()
	}
	if value := res.Get(pathRoot + `Content`); value.Exists() && !data.Content.IsNull() {
		data.Content = tfutils.ParseStringFromGJSON(value)
	} else if data.Content.ValueString() != "activity" {
		data.Content = types.StringNull()
	}
	if value := res.Get(pathRoot + `ErrorContent`); value.Exists() && !data.ErrorContent.IsNull() {
		data.ErrorContent = tfutils.ParseStringFromGJSON(value)
	} else if data.ErrorContent.ValueString() != "payload" {
		data.ErrorContent = types.StringNull()
	}
	if value := res.Get(pathRoot + `PreservedRequestHeader`); value.Exists() && !data.PreservedRequestHeader.IsNull() {
		data.PreservedRequestHeader = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.PreservedRequestHeader = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `PreservedResponseHeader`); value.Exists() && !data.PreservedResponseHeader.IsNull() {
		data.PreservedResponseHeader = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.PreservedResponseHeader = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `MessageBuffering`); value.Exists() && !data.MessageBuffering.IsNull() {
		data.MessageBuffering = tfutils.BoolFromString(value.String())
	} else if data.MessageBuffering.ValueBool() {
		data.MessageBuffering = types.BoolNull()
	}
	if value := res.Get(pathRoot + `DeploymentState`); value.Exists() && !data.DeploymentState.IsNull() {
		data.DeploymentState = tfutils.ParseStringFromGJSON(value)
	} else if data.DeploymentState.ValueString() != "running" {
		data.DeploymentState = types.StringNull()
	}
	if value := res.Get(pathRoot + `ShareRateLimitCount`); value.Exists() && !data.ShareRateLimitCount.IsNull() {
		data.ShareRateLimitCount = tfutils.ParseStringFromGJSON(value)
	} else if data.ShareRateLimitCount.ValueString() != "unset" {
		data.ShareRateLimitCount = types.StringNull()
	}
	if value := res.Get(pathRoot + `ReturnV5Responses`); value.Exists() && !data.ReturnV5Responses.IsNull() {
		data.ReturnV5Responses = tfutils.BoolFromString(value.String())
	} else if data.ReturnV5Responses.ValueBool() {
		data.ReturnV5Responses = types.BoolNull()
	}
	if value := res.Get(pathRoot + `CopyIdHeadersToMessage`); value.Exists() && !data.CopyIdHeadersToMessage.IsNull() {
		data.CopyIdHeadersToMessage = tfutils.BoolFromString(value.String())
	} else if data.CopyIdHeadersToMessage.ValueBool() {
		data.CopyIdHeadersToMessage = types.BoolNull()
	}
	if value := res.Get(pathRoot + `EnforceRequiredParams`); value.Exists() && !data.EnforceRequiredParams.IsNull() {
		data.EnforceRequiredParams = tfutils.BoolFromString(value.String())
	} else if !data.EnforceRequiredParams.ValueBool() {
		data.EnforceRequiredParams = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AllowChunkedUploads`); value.Exists() && !data.AllowChunkedUploads.IsNull() {
		data.AllowChunkedUploads = tfutils.BoolFromString(value.String())
	} else if !data.AllowChunkedUploads.ValueBool() {
		data.AllowChunkedUploads = types.BoolNull()
	}
	if value := res.Get(pathRoot + `SetV5RequestHeaders`); value.Exists() && !data.SetV5RequestHeaders.IsNull() {
		data.SetV5RequestHeaders = tfutils.BoolFromString(value.String())
	} else if data.SetV5RequestHeaders.ValueBool() {
		data.SetV5RequestHeaders = types.BoolNull()
	}
	if value := res.Get(pathRoot + `GetRawBodyValue`); value.Exists() && !data.GetRawBodyValue.IsNull() {
		data.GetRawBodyValue = tfutils.BoolFromString(value.String())
	} else if data.GetRawBodyValue.ValueBool() {
		data.GetRawBodyValue = types.BoolNull()
	}
	if value := res.Get(pathRoot + `AllowedAPIProtocols`); value.Exists() {
		data.AllowedApiProtocols.UpdateFromBody(ctx, "", value)
	} else {
		data.AllowedApiProtocols = nil
	}
	if value := res.Get(pathRoot + `AllowTrailingSlash`); value.Exists() && !data.AllowTrailingSlash.IsNull() {
		data.AllowTrailingSlash = tfutils.BoolFromString(value.String())
	} else if data.AllowTrailingSlash.ValueBool() {
		data.AllowTrailingSlash = types.BoolNull()
	}
	if value := res.Get(pathRoot + `EnforceAllHeadersCaseInsensitive`); value.Exists() && !data.EnforceAllHeadersCaseInsensitive.IsNull() {
		data.EnforceAllHeadersCaseInsensitive = tfutils.BoolFromString(value.String())
	} else if data.EnforceAllHeadersCaseInsensitive.ValueBool() {
		data.EnforceAllHeadersCaseInsensitive = types.BoolNull()
	}
	if value := res.Get(pathRoot + `EnforceFormDataParameter`); value.Exists() && !data.EnforceFormDataParameter.IsNull() {
		data.EnforceFormDataParameter = tfutils.BoolFromString(value.String())
	} else if !data.EnforceFormDataParameter.ValueBool() {
		data.EnforceFormDataParameter = types.BoolNull()
	}
	if value := res.Get(pathRoot + `ForceHttp500ForSoap11`); value.Exists() && !data.ForceHttp500ForSoap11.IsNull() {
		data.ForceHttp500ForSoap11 = tfutils.BoolFromString(value.String())
	} else if data.ForceHttp500ForSoap11.ValueBool() {
		data.ForceHttp500ForSoap11 = types.BoolNull()
	}
}

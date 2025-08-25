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

	"github.com/hashicorp/terraform-plugin-framework/attr"
	DataSourceSchema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	ResourceSchema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type DmOAuthComponents struct {
	None                     types.Bool `tfsdk:"none"`
	OAuthValidateRequest     types.Bool `tfsdk:"o_auth_validate_request"`
	OAuthGenerateAzCode      types.Bool `tfsdk:"o_auth_generate_az_code"`
	OAuthVerifyAzCode        types.Bool `tfsdk:"o_auth_verify_az_code"`
	OAuthVerifyRefreshToken  types.Bool `tfsdk:"o_auth_verify_refresh_token"`
	OAuthCollectMetadata     types.Bool `tfsdk:"o_auth_collect_metadata"`
	OAuthGenerateAccessToken types.Bool `tfsdk:"o_auth_generate_access_token"`
	OAuthIntrospectToken     types.Bool `tfsdk:"o_auth_introspect_token"`
	OAuthRevokeToken         types.Bool `tfsdk:"o_auth_revoke_token"`
}

var DmOAuthComponentsObjectType = map[string]attr.Type{
	"none":                         types.BoolType,
	"o_auth_validate_request":      types.BoolType,
	"o_auth_generate_az_code":      types.BoolType,
	"o_auth_verify_az_code":        types.BoolType,
	"o_auth_verify_refresh_token":  types.BoolType,
	"o_auth_collect_metadata":      types.BoolType,
	"o_auth_generate_access_token": types.BoolType,
	"o_auth_introspect_token":      types.BoolType,
	"o_auth_revoke_token":          types.BoolType,
}
var DmOAuthComponentsObjectDefault = map[string]attr.Value{
	"none":                         types.BoolValue(false),
	"o_auth_validate_request":      types.BoolValue(false),
	"o_auth_generate_az_code":      types.BoolValue(false),
	"o_auth_verify_az_code":        types.BoolValue(false),
	"o_auth_verify_refresh_token":  types.BoolValue(false),
	"o_auth_collect_metadata":      types.BoolValue(false),
	"o_auth_generate_access_token": types.BoolValue(false),
	"o_auth_introspect_token":      types.BoolValue(false),
	"o_auth_revoke_token":          types.BoolValue(false),
}

func GetDmOAuthComponentsDataSourceSchema(description string, cliAlias string, referenceTo string) DataSourceSchema.SingleNestedAttribute {
	var DmOAuthComponentsDataSourceSchema = DataSourceSchema.SingleNestedAttribute{
		Computed: true,
		Attributes: map[string]DataSourceSchema.Attribute{
			"none": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("", "", "").AddDefaultValue("false").String,
				Computed:            true,
			},
			"o_auth_validate_request": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Validate request", "", "").AddDefaultValue("false").String,
				Computed:            true,
			},
			"o_auth_generate_az_code": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Generate authorization code", "", "").AddDefaultValue("false").String,
				Computed:            true,
			},
			"o_auth_verify_az_code": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Verify authorization code", "", "").AddDefaultValue("false").String,
				Computed:            true,
			},
			"o_auth_verify_refresh_token": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Verify refresh token", "", "").AddDefaultValue("false").String,
				Computed:            true,
			},
			"o_auth_collect_metadata": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Collect Metadata", "", "").AddDefaultValue("false").String,
				Computed:            true,
			},
			"o_auth_generate_access_token": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Generate access token", "", "").AddDefaultValue("false").String,
				Computed:            true,
			},
			"o_auth_introspect_token": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Introspect token", "", "").AddDefaultValue("false").String,
				Computed:            true,
			},
			"o_auth_revoke_token": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Revoke token", "", "").AddDefaultValue("false").String,
				Computed:            true,
			},
		},
	}
	DmOAuthComponentsDataSourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, referenceTo).String
	return DmOAuthComponentsDataSourceSchema
}
func GetDmOAuthComponentsResourceSchema(description string, cliAlias string, referenceTo string, required bool) ResourceSchema.SingleNestedAttribute {
	var DmOAuthComponentsResourceSchema = ResourceSchema.SingleNestedAttribute{
		Default: objectdefault.StaticValue(
			types.ObjectValueMust(
				DmOAuthComponentsObjectType,
				DmOAuthComponentsObjectDefault,
			)),
		Attributes: map[string]ResourceSchema.Attribute{
			"none": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"o_auth_validate_request": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Validate request", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"o_auth_generate_az_code": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Generate authorization code", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"o_auth_verify_az_code": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Verify authorization code", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"o_auth_verify_refresh_token": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Verify refresh token", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"o_auth_collect_metadata": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Collect Metadata", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"o_auth_generate_access_token": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Generate access token", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"o_auth_introspect_token": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Introspect token", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"o_auth_revoke_token": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Revoke token", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
		},
	}
	DmOAuthComponentsResourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, referenceTo).String
	if required {
		DmOAuthComponentsResourceSchema.Required = true
	} else {
		DmOAuthComponentsResourceSchema.Optional = true
		DmOAuthComponentsResourceSchema.Computed = true
	}
	return DmOAuthComponentsResourceSchema
}

func (data DmOAuthComponents) IsNull() bool {
	if !data.None.IsNull() {
		return false
	}
	if !data.OAuthValidateRequest.IsNull() {
		return false
	}
	if !data.OAuthGenerateAzCode.IsNull() {
		return false
	}
	if !data.OAuthVerifyAzCode.IsNull() {
		return false
	}
	if !data.OAuthVerifyRefreshToken.IsNull() {
		return false
	}
	if !data.OAuthCollectMetadata.IsNull() {
		return false
	}
	if !data.OAuthGenerateAccessToken.IsNull() {
		return false
	}
	if !data.OAuthIntrospectToken.IsNull() {
		return false
	}
	if !data.OAuthRevokeToken.IsNull() {
		return false
	}
	return true
}

func (data DmOAuthComponents) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""

	if !data.None.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`none`, tfutils.StringFromBool(data.None, ""))
	}
	if !data.OAuthValidateRequest.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`OAuthValidateRequest`, tfutils.StringFromBool(data.OAuthValidateRequest, ""))
	}
	if !data.OAuthGenerateAzCode.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`OAuthGenerateAZCode`, tfutils.StringFromBool(data.OAuthGenerateAzCode, ""))
	}
	if !data.OAuthVerifyAzCode.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`OAuthVerifyAZCode`, tfutils.StringFromBool(data.OAuthVerifyAzCode, ""))
	}
	if !data.OAuthVerifyRefreshToken.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`OAuthVerifyRefreshToken`, tfutils.StringFromBool(data.OAuthVerifyRefreshToken, ""))
	}
	if !data.OAuthCollectMetadata.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`OAuthCollectMetadata`, tfutils.StringFromBool(data.OAuthCollectMetadata, ""))
	}
	if !data.OAuthGenerateAccessToken.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`OAuthGenerateAccessToken`, tfutils.StringFromBool(data.OAuthGenerateAccessToken, ""))
	}
	if !data.OAuthIntrospectToken.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`OAuthIntrospectToken`, tfutils.StringFromBool(data.OAuthIntrospectToken, ""))
	}
	if !data.OAuthRevokeToken.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`OAuthRevokeToken`, tfutils.StringFromBool(data.OAuthRevokeToken, ""))
	}
	return body
}

func (data *DmOAuthComponents) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `none`); value.Exists() {
		data.None = tfutils.BoolFromString(value.String())
	} else {
		data.None = types.BoolNull()
	}
	if value := res.Get(pathRoot + `OAuthValidateRequest`); value.Exists() {
		data.OAuthValidateRequest = tfutils.BoolFromString(value.String())
	} else {
		data.OAuthValidateRequest = types.BoolNull()
	}
	if value := res.Get(pathRoot + `OAuthGenerateAZCode`); value.Exists() {
		data.OAuthGenerateAzCode = tfutils.BoolFromString(value.String())
	} else {
		data.OAuthGenerateAzCode = types.BoolNull()
	}
	if value := res.Get(pathRoot + `OAuthVerifyAZCode`); value.Exists() {
		data.OAuthVerifyAzCode = tfutils.BoolFromString(value.String())
	} else {
		data.OAuthVerifyAzCode = types.BoolNull()
	}
	if value := res.Get(pathRoot + `OAuthVerifyRefreshToken`); value.Exists() {
		data.OAuthVerifyRefreshToken = tfutils.BoolFromString(value.String())
	} else {
		data.OAuthVerifyRefreshToken = types.BoolNull()
	}
	if value := res.Get(pathRoot + `OAuthCollectMetadata`); value.Exists() {
		data.OAuthCollectMetadata = tfutils.BoolFromString(value.String())
	} else {
		data.OAuthCollectMetadata = types.BoolNull()
	}
	if value := res.Get(pathRoot + `OAuthGenerateAccessToken`); value.Exists() {
		data.OAuthGenerateAccessToken = tfutils.BoolFromString(value.String())
	} else {
		data.OAuthGenerateAccessToken = types.BoolNull()
	}
	if value := res.Get(pathRoot + `OAuthIntrospectToken`); value.Exists() {
		data.OAuthIntrospectToken = tfutils.BoolFromString(value.String())
	} else {
		data.OAuthIntrospectToken = types.BoolNull()
	}
	if value := res.Get(pathRoot + `OAuthRevokeToken`); value.Exists() {
		data.OAuthRevokeToken = tfutils.BoolFromString(value.String())
	} else {
		data.OAuthRevokeToken = types.BoolNull()
	}
}

func (data *DmOAuthComponents) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `none`); value.Exists() && !data.None.IsNull() {
		data.None = tfutils.BoolFromString(value.String())
	} else if data.None.ValueBool() {
		data.None = types.BoolNull()
	}
	if value := res.Get(pathRoot + `OAuthValidateRequest`); value.Exists() && !data.OAuthValidateRequest.IsNull() {
		data.OAuthValidateRequest = tfutils.BoolFromString(value.String())
	} else if data.OAuthValidateRequest.ValueBool() {
		data.OAuthValidateRequest = types.BoolNull()
	}
	if value := res.Get(pathRoot + `OAuthGenerateAZCode`); value.Exists() && !data.OAuthGenerateAzCode.IsNull() {
		data.OAuthGenerateAzCode = tfutils.BoolFromString(value.String())
	} else if data.OAuthGenerateAzCode.ValueBool() {
		data.OAuthGenerateAzCode = types.BoolNull()
	}
	if value := res.Get(pathRoot + `OAuthVerifyAZCode`); value.Exists() && !data.OAuthVerifyAzCode.IsNull() {
		data.OAuthVerifyAzCode = tfutils.BoolFromString(value.String())
	} else if data.OAuthVerifyAzCode.ValueBool() {
		data.OAuthVerifyAzCode = types.BoolNull()
	}
	if value := res.Get(pathRoot + `OAuthVerifyRefreshToken`); value.Exists() && !data.OAuthVerifyRefreshToken.IsNull() {
		data.OAuthVerifyRefreshToken = tfutils.BoolFromString(value.String())
	} else if data.OAuthVerifyRefreshToken.ValueBool() {
		data.OAuthVerifyRefreshToken = types.BoolNull()
	}
	if value := res.Get(pathRoot + `OAuthCollectMetadata`); value.Exists() && !data.OAuthCollectMetadata.IsNull() {
		data.OAuthCollectMetadata = tfutils.BoolFromString(value.String())
	} else if data.OAuthCollectMetadata.ValueBool() {
		data.OAuthCollectMetadata = types.BoolNull()
	}
	if value := res.Get(pathRoot + `OAuthGenerateAccessToken`); value.Exists() && !data.OAuthGenerateAccessToken.IsNull() {
		data.OAuthGenerateAccessToken = tfutils.BoolFromString(value.String())
	} else if data.OAuthGenerateAccessToken.ValueBool() {
		data.OAuthGenerateAccessToken = types.BoolNull()
	}
	if value := res.Get(pathRoot + `OAuthIntrospectToken`); value.Exists() && !data.OAuthIntrospectToken.IsNull() {
		data.OAuthIntrospectToken = tfutils.BoolFromString(value.String())
	} else if data.OAuthIntrospectToken.ValueBool() {
		data.OAuthIntrospectToken = types.BoolNull()
	}
	if value := res.Get(pathRoot + `OAuthRevokeToken`); value.Exists() && !data.OAuthRevokeToken.IsNull() {
		data.OAuthRevokeToken = tfutils.BoolFromString(value.String())
	} else if data.OAuthRevokeToken.ValueBool() {
		data.OAuthRevokeToken = types.BoolNull()
	}
}

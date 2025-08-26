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
	OauthValidateRequest     types.Bool `tfsdk:"oauth_validate_request"`
	OauthGenerateAzCode      types.Bool `tfsdk:"oauth_generate_az_code"`
	OauthVerifyAzCode        types.Bool `tfsdk:"oauth_verify_az_code"`
	OauthVerifyRefreshToken  types.Bool `tfsdk:"oauth_verify_refresh_token"`
	OauthCollectMetadata     types.Bool `tfsdk:"oauth_collect_metadata"`
	OauthGenerateAccessToken types.Bool `tfsdk:"oauth_generate_access_token"`
	OauthIntrospectToken     types.Bool `tfsdk:"oauth_introspect_token"`
	OauthRevokeToken         types.Bool `tfsdk:"oauth_revoke_token"`
}

var DmOAuthComponentsObjectType = map[string]attr.Type{
	"none":                        types.BoolType,
	"oauth_validate_request":      types.BoolType,
	"oauth_generate_az_code":      types.BoolType,
	"oauth_verify_az_code":        types.BoolType,
	"oauth_verify_refresh_token":  types.BoolType,
	"oauth_collect_metadata":      types.BoolType,
	"oauth_generate_access_token": types.BoolType,
	"oauth_introspect_token":      types.BoolType,
	"oauth_revoke_token":          types.BoolType,
}
var DmOAuthComponentsObjectDefault = map[string]attr.Value{
	"none":                        types.BoolValue(false),
	"oauth_validate_request":      types.BoolValue(false),
	"oauth_generate_az_code":      types.BoolValue(false),
	"oauth_verify_az_code":        types.BoolValue(false),
	"oauth_verify_refresh_token":  types.BoolValue(false),
	"oauth_collect_metadata":      types.BoolValue(false),
	"oauth_generate_access_token": types.BoolValue(false),
	"oauth_introspect_token":      types.BoolValue(false),
	"oauth_revoke_token":          types.BoolValue(false),
}

func GetDmOAuthComponentsDataSourceSchema(description string, cliAlias string, referenceTo string) DataSourceSchema.SingleNestedAttribute {
	var DmOAuthComponentsDataSourceSchema = DataSourceSchema.SingleNestedAttribute{
		Computed: true,
		Attributes: map[string]DataSourceSchema.Attribute{
			"none": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("", "", "").AddDefaultValue("false").String,
				Computed:            true,
			},
			"oauth_validate_request": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Validate request", "", "").AddDefaultValue("false").String,
				Computed:            true,
			},
			"oauth_generate_az_code": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Generate authorization code", "", "").AddDefaultValue("false").String,
				Computed:            true,
			},
			"oauth_verify_az_code": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Verify authorization code", "", "").AddDefaultValue("false").String,
				Computed:            true,
			},
			"oauth_verify_refresh_token": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Verify refresh token", "", "").AddDefaultValue("false").String,
				Computed:            true,
			},
			"oauth_collect_metadata": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Collect Metadata", "", "").AddDefaultValue("false").String,
				Computed:            true,
			},
			"oauth_generate_access_token": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Generate access token", "", "").AddDefaultValue("false").String,
				Computed:            true,
			},
			"oauth_introspect_token": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Introspect token", "", "").AddDefaultValue("false").String,
				Computed:            true,
			},
			"oauth_revoke_token": DataSourceSchema.BoolAttribute{
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
			"oauth_validate_request": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Validate request", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"oauth_generate_az_code": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Generate authorization code", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"oauth_verify_az_code": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Verify authorization code", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"oauth_verify_refresh_token": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Verify refresh token", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"oauth_collect_metadata": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Collect Metadata", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"oauth_generate_access_token": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Generate access token", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"oauth_introspect_token": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Introspect token", "", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"oauth_revoke_token": ResourceSchema.BoolAttribute{
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
	if !data.OauthValidateRequest.IsNull() {
		return false
	}
	if !data.OauthGenerateAzCode.IsNull() {
		return false
	}
	if !data.OauthVerifyAzCode.IsNull() {
		return false
	}
	if !data.OauthVerifyRefreshToken.IsNull() {
		return false
	}
	if !data.OauthCollectMetadata.IsNull() {
		return false
	}
	if !data.OauthGenerateAccessToken.IsNull() {
		return false
	}
	if !data.OauthIntrospectToken.IsNull() {
		return false
	}
	if !data.OauthRevokeToken.IsNull() {
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
	if !data.OauthValidateRequest.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`OAuthValidateRequest`, tfutils.StringFromBool(data.OauthValidateRequest, ""))
	}
	if !data.OauthGenerateAzCode.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`OAuthGenerateAZCode`, tfutils.StringFromBool(data.OauthGenerateAzCode, ""))
	}
	if !data.OauthVerifyAzCode.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`OAuthVerifyAZCode`, tfutils.StringFromBool(data.OauthVerifyAzCode, ""))
	}
	if !data.OauthVerifyRefreshToken.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`OAuthVerifyRefreshToken`, tfutils.StringFromBool(data.OauthVerifyRefreshToken, ""))
	}
	if !data.OauthCollectMetadata.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`OAuthCollectMetadata`, tfutils.StringFromBool(data.OauthCollectMetadata, ""))
	}
	if !data.OauthGenerateAccessToken.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`OAuthGenerateAccessToken`, tfutils.StringFromBool(data.OauthGenerateAccessToken, ""))
	}
	if !data.OauthIntrospectToken.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`OAuthIntrospectToken`, tfutils.StringFromBool(data.OauthIntrospectToken, ""))
	}
	if !data.OauthRevokeToken.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`OAuthRevokeToken`, tfutils.StringFromBool(data.OauthRevokeToken, ""))
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
		data.OauthValidateRequest = tfutils.BoolFromString(value.String())
	} else {
		data.OauthValidateRequest = types.BoolNull()
	}
	if value := res.Get(pathRoot + `OAuthGenerateAZCode`); value.Exists() {
		data.OauthGenerateAzCode = tfutils.BoolFromString(value.String())
	} else {
		data.OauthGenerateAzCode = types.BoolNull()
	}
	if value := res.Get(pathRoot + `OAuthVerifyAZCode`); value.Exists() {
		data.OauthVerifyAzCode = tfutils.BoolFromString(value.String())
	} else {
		data.OauthVerifyAzCode = types.BoolNull()
	}
	if value := res.Get(pathRoot + `OAuthVerifyRefreshToken`); value.Exists() {
		data.OauthVerifyRefreshToken = tfutils.BoolFromString(value.String())
	} else {
		data.OauthVerifyRefreshToken = types.BoolNull()
	}
	if value := res.Get(pathRoot + `OAuthCollectMetadata`); value.Exists() {
		data.OauthCollectMetadata = tfutils.BoolFromString(value.String())
	} else {
		data.OauthCollectMetadata = types.BoolNull()
	}
	if value := res.Get(pathRoot + `OAuthGenerateAccessToken`); value.Exists() {
		data.OauthGenerateAccessToken = tfutils.BoolFromString(value.String())
	} else {
		data.OauthGenerateAccessToken = types.BoolNull()
	}
	if value := res.Get(pathRoot + `OAuthIntrospectToken`); value.Exists() {
		data.OauthIntrospectToken = tfutils.BoolFromString(value.String())
	} else {
		data.OauthIntrospectToken = types.BoolNull()
	}
	if value := res.Get(pathRoot + `OAuthRevokeToken`); value.Exists() {
		data.OauthRevokeToken = tfutils.BoolFromString(value.String())
	} else {
		data.OauthRevokeToken = types.BoolNull()
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
	if value := res.Get(pathRoot + `OAuthValidateRequest`); value.Exists() && !data.OauthValidateRequest.IsNull() {
		data.OauthValidateRequest = tfutils.BoolFromString(value.String())
	} else if data.OauthValidateRequest.ValueBool() {
		data.OauthValidateRequest = types.BoolNull()
	}
	if value := res.Get(pathRoot + `OAuthGenerateAZCode`); value.Exists() && !data.OauthGenerateAzCode.IsNull() {
		data.OauthGenerateAzCode = tfutils.BoolFromString(value.String())
	} else if data.OauthGenerateAzCode.ValueBool() {
		data.OauthGenerateAzCode = types.BoolNull()
	}
	if value := res.Get(pathRoot + `OAuthVerifyAZCode`); value.Exists() && !data.OauthVerifyAzCode.IsNull() {
		data.OauthVerifyAzCode = tfutils.BoolFromString(value.String())
	} else if data.OauthVerifyAzCode.ValueBool() {
		data.OauthVerifyAzCode = types.BoolNull()
	}
	if value := res.Get(pathRoot + `OAuthVerifyRefreshToken`); value.Exists() && !data.OauthVerifyRefreshToken.IsNull() {
		data.OauthVerifyRefreshToken = tfutils.BoolFromString(value.String())
	} else if data.OauthVerifyRefreshToken.ValueBool() {
		data.OauthVerifyRefreshToken = types.BoolNull()
	}
	if value := res.Get(pathRoot + `OAuthCollectMetadata`); value.Exists() && !data.OauthCollectMetadata.IsNull() {
		data.OauthCollectMetadata = tfutils.BoolFromString(value.String())
	} else if data.OauthCollectMetadata.ValueBool() {
		data.OauthCollectMetadata = types.BoolNull()
	}
	if value := res.Get(pathRoot + `OAuthGenerateAccessToken`); value.Exists() && !data.OauthGenerateAccessToken.IsNull() {
		data.OauthGenerateAccessToken = tfutils.BoolFromString(value.String())
	} else if data.OauthGenerateAccessToken.ValueBool() {
		data.OauthGenerateAccessToken = types.BoolNull()
	}
	if value := res.Get(pathRoot + `OAuthIntrospectToken`); value.Exists() && !data.OauthIntrospectToken.IsNull() {
		data.OauthIntrospectToken = tfutils.BoolFromString(value.String())
	} else if data.OauthIntrospectToken.ValueBool() {
		data.OauthIntrospectToken = types.BoolNull()
	}
	if value := res.Get(pathRoot + `OAuthRevokeToken`); value.Exists() && !data.OauthRevokeToken.IsNull() {
		data.OauthRevokeToken = tfutils.BoolFromString(value.String())
	} else if data.OauthRevokeToken.ValueBool() {
		data.OauthRevokeToken = types.BoolNull()
	}
}

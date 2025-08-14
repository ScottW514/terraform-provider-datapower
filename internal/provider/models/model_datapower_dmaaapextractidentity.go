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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type DmAAAPExtractIdentity struct {
	EiBitmap                      *DmAAAPEIBitmap `tfsdk:"ei_bitmap"`
	EiCustomUrl                   types.String    `tfsdk:"ei_custom_url"`
	EixPath                       types.String    `tfsdk:"eix_path"`
	EiSignerDnValcred             types.String    `tfsdk:"ei_signer_dn_valcred"`
	EiCookieName                  types.String    `tfsdk:"ei_cookie_name"`
	EiBasicAuthRealm              types.String    `tfsdk:"ei_basic_auth_realm"`
	EiUseWsSec                    types.Bool      `tfsdk:"ei_use_ws_sec"`
	EiMetadata                    types.String    `tfsdk:"ei_metadata"`
	EiAllowRemoteTokenReference   types.Bool      `tfsdk:"ei_allow_remote_token_reference"`
	EiRemoteTokenProcessService   types.String    `tfsdk:"ei_remote_token_process_service"`
	EiPasswordRetrievalMechanism  types.String    `tfsdk:"ei_password_retrieval_mechanism"`
	EiPasswordRetrievalCustomUrl  types.String    `tfsdk:"ei_password_retrieval_custom_url"`
	EiPasswordRetrievalAaaInfoUrl types.String    `tfsdk:"ei_password_retrieval_aaa_info_url"`
	EiFormsLoginPolicy            types.String    `tfsdk:"ei_forms_login_policy"`
	EioAuthClientGroup            types.String    `tfsdk:"eio_auth_client_group"`
	EisslClientConfigType         types.String    `tfsdk:"eissl_client_config_type"`
	EisslClientProfile            types.String    `tfsdk:"eissl_client_profile"`
	EijwtValidator                types.String    `tfsdk:"eijwt_validator"`
	EiSocialLoginPolicy           types.String    `tfsdk:"ei_social_login_policy"`
	EisamlResponseValCred         types.String    `tfsdk:"eisaml_response_val_cred"`
}

var DmAAAPExtractIdentityObjectType = map[string]attr.Type{
	"ei_bitmap":                          types.ObjectType{AttrTypes: DmAAAPEIBitmapObjectType},
	"ei_custom_url":                      types.StringType,
	"eix_path":                           types.StringType,
	"ei_signer_dn_valcred":               types.StringType,
	"ei_cookie_name":                     types.StringType,
	"ei_basic_auth_realm":                types.StringType,
	"ei_use_ws_sec":                      types.BoolType,
	"ei_metadata":                        types.StringType,
	"ei_allow_remote_token_reference":    types.BoolType,
	"ei_remote_token_process_service":    types.StringType,
	"ei_password_retrieval_mechanism":    types.StringType,
	"ei_password_retrieval_custom_url":   types.StringType,
	"ei_password_retrieval_aaa_info_url": types.StringType,
	"ei_forms_login_policy":              types.StringType,
	"eio_auth_client_group":              types.StringType,
	"eissl_client_config_type":           types.StringType,
	"eissl_client_profile":               types.StringType,
	"eijwt_validator":                    types.StringType,
	"ei_social_login_policy":             types.StringType,
	"eisaml_response_val_cred":           types.StringType,
}
var DmAAAPExtractIdentityObjectDefault = map[string]attr.Value{
	"ei_bitmap":                          types.ObjectValueMust(DmAAAPEIBitmapObjectType, DmAAAPEIBitmapObjectDefault),
	"ei_custom_url":                      types.StringNull(),
	"eix_path":                           types.StringNull(),
	"ei_signer_dn_valcred":               types.StringNull(),
	"ei_cookie_name":                     types.StringNull(),
	"ei_basic_auth_realm":                types.StringValue("login"),
	"ei_use_ws_sec":                      types.BoolValue(false),
	"ei_metadata":                        types.StringNull(),
	"ei_allow_remote_token_reference":    types.BoolValue(false),
	"ei_remote_token_process_service":    types.StringNull(),
	"ei_password_retrieval_mechanism":    types.StringValue("xmlfile"),
	"ei_password_retrieval_custom_url":   types.StringNull(),
	"ei_password_retrieval_aaa_info_url": types.StringNull(),
	"ei_forms_login_policy":              types.StringNull(),
	"eio_auth_client_group":              types.StringNull(),
	"eissl_client_config_type":           types.StringValue("proxy"),
	"eissl_client_profile":               types.StringNull(),
	"eijwt_validator":                    types.StringNull(),
	"ei_social_login_policy":             types.StringNull(),
	"eisaml_response_val_cred":           types.StringNull(),
}
var DmAAAPExtractIdentityDataSourceSchema = DataSourceSchema.SingleNestedAttribute{
	Computed: true,
	Attributes: map[string]DataSourceSchema.Attribute{
		"ei_bitmap": GetDmAAAPEIBitmapDataSourceSchema("Specify the methods to extract identifies.", "method", ""),
		"ei_custom_url": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the location of the processing file. This file is the stylesheet or GatewayScript that extracts identity information from the candidate XML document.", "custom-url", "").String,
			Computed:            true,
		},
		"eix_path": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the XPath expression to apply to the message. The extracted identity uses the string value for this expression.", "xpath", "").String,
			Computed:            true,
		},
		"ei_signer_dn_valcred": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the validation credentials to verify the signer certificate.", "valcred", "cryptovalcred").String,
			Computed:            true,
		},
		"ei_cookie_name": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the cookie in the <tt>Cookie</tt> header to extract and use as the identity token.", "cookie-name", "").String,
			Computed:            true,
		},
		"ei_basic_auth_realm": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the name of the HTTP basic authentication realm as described by RFC 2617. A web browser might display this name to help the user determine which credentials must be supplied. The default value is <tt>login</tt> .", "basic-auth-realm", "").AddDefaultValue("login").String,
			Computed:            true,
		},
		"ei_use_ws_sec": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("<p>Specify whether to use the WS-Security token first. By default, this feature is not enabled.</p><ul><li>When enabled, use the token from the appropriate WS-Security <tt>Security</tt> header. For example, use when an LTPA token is extracted from a <tt>Cookie</tt> header and you want to use the one that is available from the <tt>Security</tt> header.</li><li>When not enabled, use the token that is extracted somewhere other than the WS-Security <tt>Security</tt> header.</li></ul>", "use-wssec-token", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"ei_metadata": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("<p>Specify the processing metadata, which defines items such as variables and protocol headers.</p><ul><li>When specified, processing returns only the defined metadata items.</li><li>When not specified, processing returns all metadata items for the current processing rule.</li></ul>", "metadata", "processingmetadata").String,
			Computed:            true,
		},
		"ei_allow_remote_token_reference": DataSourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to contact a remote location for the final WS-Security <tt>Security</tt> token. Processing might need the final token when the message indicates that the token is at a remote location. By default, processing cannot contact a remote location.", "remote-token-allowed", "").AddDefaultValue("false").String,
			Computed:            true,
		},
		"ei_remote_token_process_service": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("<p>Specify the URL of the service that can process the remote security token as a SOAP call and, if successful, respond with the final security token.</p><p>The remote WS-Security token can be signed, encrypted, or encoded. A service with different actions can process the remote token. For example, this token can decrypt pieces of a remote SAML assertion, run an XSLT transform, or the AAA policy can assert the token.</p>", "remote-token-url", "").String,
			Computed:            true,
		},
		"ei_password_retrieval_mechanism": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the method to obtain the user password. The password is needed to calculate the derived symmetric key.", "password-retrieval-method", "").AddStringEnum("xmlfile", "custom").AddDefaultValue("xmlfile").String,
			Computed:            true,
		},
		"ei_password_retrieval_custom_url": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the location of custom file to obtain the password.", "password-retrieval-custom-url", "").String,
			Computed:            true,
		},
		"ei_password_retrieval_aaa_info_url": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the location of AAA information file to obtain the password.", "password-retrieval-xmlfile-url", "").String,
			Computed:            true,
		},
		"ei_forms_login_policy": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the HTML forms-based login policy. This policy defines the forms that collect user and password information.", "forms-login-policy", "formsloginpolicy").String,
			Computed:            true,
		},
		"eio_auth_client_group": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the name of the OAuth client group.", "oauth-client-group", "oauthsupportedclientgroup").String,
			Computed:            true,
		},
		"eissl_client_config_type": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the TLS profile type to secure connections.", "ssl-client-type", "").AddStringEnum("proxy", "client").AddDefaultValue("proxy").String,
			Computed:            true,
		},
		"eissl_client_profile": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the TLS client profile to secure connections to targets", "ssl-client", "sslclientprofile").String,
			Computed:            true,
		},
		"eijwt_validator": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the JWT validator to validate the JWT.", "validate-jwt", "aaajwtvalidator").String,
			Computed:            true,
		},
		"ei_social_login_policy": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the social login policy. To choose a policy at run time, set the value with the <tt>var:///context/AAA/social-login-policy-name</tt> variable. When the value is set with a variable before the AAA action, the variable takes precedence over this value.", "social-login-policy", "socialloginpolicy").String,
			Computed:            true,
		},
		"eisaml_response_val_cred": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the validation credentials to verify the signature of the SAML response.", "saml-response-valcred", "cryptovalcred").String,
			Computed:            true,
		},
	},
}
var DmAAAPExtractIdentityResourceSchema = ResourceSchema.SingleNestedAttribute{
	Default: objectdefault.StaticValue(
		types.ObjectValueMust(
			DmAAAPExtractIdentityObjectType,
			DmAAAPExtractIdentityObjectDefault,
		)),
	Attributes: map[string]ResourceSchema.Attribute{
		"ei_bitmap": GetDmAAAPEIBitmapResourceSchema("Specify the methods to extract identifies.", "method", "", false),
		"ei_custom_url": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the location of the processing file. This file is the stylesheet or GatewayScript that extracts identity information from the candidate XML document.", "custom-url", "").String,
			Optional:            true,
		},
		"eix_path": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the XPath expression to apply to the message. The extracted identity uses the string value for this expression.", "xpath", "").String,
			Optional:            true,
		},
		"ei_signer_dn_valcred": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the validation credentials to verify the signer certificate.", "valcred", "cryptovalcred").String,
			Optional:            true,
		},
		"ei_cookie_name": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the cookie in the <tt>Cookie</tt> header to extract and use as the identity token.", "cookie-name", "").String,
			Optional:            true,
		},
		"ei_basic_auth_realm": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the name of the HTTP basic authentication realm as described by RFC 2617. A web browser might display this name to help the user determine which credentials must be supplied. The default value is <tt>login</tt> .", "basic-auth-realm", "").AddDefaultValue("login").String,
			Computed:            true,
			Optional:            true,
			Default:             stringdefault.StaticString("login"),
		},
		"ei_use_ws_sec": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("<p>Specify whether to use the WS-Security token first. By default, this feature is not enabled.</p><ul><li>When enabled, use the token from the appropriate WS-Security <tt>Security</tt> header. For example, use when an LTPA token is extracted from a <tt>Cookie</tt> header and you want to use the one that is available from the <tt>Security</tt> header.</li><li>When not enabled, use the token that is extracted somewhere other than the WS-Security <tt>Security</tt> header.</li></ul>", "use-wssec-token", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"ei_metadata": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("<p>Specify the processing metadata, which defines items such as variables and protocol headers.</p><ul><li>When specified, processing returns only the defined metadata items.</li><li>When not specified, processing returns all metadata items for the current processing rule.</li></ul>", "metadata", "processingmetadata").String,
			Optional:            true,
		},
		"ei_allow_remote_token_reference": ResourceSchema.BoolAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to contact a remote location for the final WS-Security <tt>Security</tt> token. Processing might need the final token when the message indicates that the token is at a remote location. By default, processing cannot contact a remote location.", "remote-token-allowed", "").AddDefaultValue("false").String,
			Computed:            true,
			Optional:            true,
			Default:             booldefault.StaticBool(false),
		},
		"ei_remote_token_process_service": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("<p>Specify the URL of the service that can process the remote security token as a SOAP call and, if successful, respond with the final security token.</p><p>The remote WS-Security token can be signed, encrypted, or encoded. A service with different actions can process the remote token. For example, this token can decrypt pieces of a remote SAML assertion, run an XSLT transform, or the AAA policy can assert the token.</p>", "remote-token-url", "").String,
			Optional:            true,
		},
		"ei_password_retrieval_mechanism": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the method to obtain the user password. The password is needed to calculate the derived symmetric key.", "password-retrieval-method", "").AddStringEnum("xmlfile", "custom").AddDefaultValue("xmlfile").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("xmlfile", "custom"),
			},
			Default: stringdefault.StaticString("xmlfile"),
		},
		"ei_password_retrieval_custom_url": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the location of custom file to obtain the password.", "password-retrieval-custom-url", "").String,
			Optional:            true,
		},
		"ei_password_retrieval_aaa_info_url": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the location of AAA information file to obtain the password.", "password-retrieval-xmlfile-url", "").String,
			Optional:            true,
		},
		"ei_forms_login_policy": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the HTML forms-based login policy. This policy defines the forms that collect user and password information.", "forms-login-policy", "formsloginpolicy").String,
			Optional:            true,
		},
		"eio_auth_client_group": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the name of the OAuth client group.", "oauth-client-group", "oauthsupportedclientgroup").String,
			Optional:            true,
		},
		"eissl_client_config_type": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the TLS profile type to secure connections.", "ssl-client-type", "").AddStringEnum("proxy", "client").AddDefaultValue("proxy").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("proxy", "client"),
			},
			Default: stringdefault.StaticString("proxy"),
		},
		"eissl_client_profile": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the TLS client profile to secure connections to targets", "ssl-client", "sslclientprofile").String,
			Optional:            true,
		},
		"eijwt_validator": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the JWT validator to validate the JWT.", "validate-jwt", "aaajwtvalidator").String,
			Optional:            true,
		},
		"ei_social_login_policy": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the social login policy. To choose a policy at run time, set the value with the <tt>var:///context/AAA/social-login-policy-name</tt> variable. When the value is set with a variable before the AAA action, the variable takes precedence over this value.", "social-login-policy", "socialloginpolicy").String,
			Optional:            true,
		},
		"eisaml_response_val_cred": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the validation credentials to verify the signature of the SAML response.", "saml-response-valcred", "cryptovalcred").String,
			Optional:            true,
		},
	},
}

func (data DmAAAPExtractIdentity) IsNull() bool {
	if data.EiBitmap != nil {
		if !data.EiBitmap.IsNull() {
			return false
		}
	}
	if !data.EiCustomUrl.IsNull() {
		return false
	}
	if !data.EixPath.IsNull() {
		return false
	}
	if !data.EiSignerDnValcred.IsNull() {
		return false
	}
	if !data.EiCookieName.IsNull() {
		return false
	}
	if !data.EiBasicAuthRealm.IsNull() {
		return false
	}
	if !data.EiUseWsSec.IsNull() {
		return false
	}
	if !data.EiMetadata.IsNull() {
		return false
	}
	if !data.EiAllowRemoteTokenReference.IsNull() {
		return false
	}
	if !data.EiRemoteTokenProcessService.IsNull() {
		return false
	}
	if !data.EiPasswordRetrievalMechanism.IsNull() {
		return false
	}
	if !data.EiPasswordRetrievalCustomUrl.IsNull() {
		return false
	}
	if !data.EiPasswordRetrievalAaaInfoUrl.IsNull() {
		return false
	}
	if !data.EiFormsLoginPolicy.IsNull() {
		return false
	}
	if !data.EioAuthClientGroup.IsNull() {
		return false
	}
	if !data.EisslClientConfigType.IsNull() {
		return false
	}
	if !data.EisslClientProfile.IsNull() {
		return false
	}
	if !data.EijwtValidator.IsNull() {
		return false
	}
	if !data.EiSocialLoginPolicy.IsNull() {
		return false
	}
	if !data.EisamlResponseValCred.IsNull() {
		return false
	}
	return true
}
func GetDmAAAPExtractIdentityDataSourceSchema(description string, cliAlias string, referenceTo string) DataSourceSchema.NestedAttribute {
	DmAAAPExtractIdentityDataSourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, referenceTo).String
	return DmAAAPExtractIdentityDataSourceSchema
}

func GetDmAAAPExtractIdentityResourceSchema(description string, cliAlias string, referenceTo string, required bool) ResourceSchema.NestedAttribute {
	if required {
		DmAAAPExtractIdentityResourceSchema.Required = true
	} else {
		DmAAAPExtractIdentityResourceSchema.Optional = true
		DmAAAPExtractIdentityResourceSchema.Computed = true
	}
	DmAAAPExtractIdentityResourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, "").String
	return DmAAAPExtractIdentityResourceSchema
}

func (data DmAAAPExtractIdentity) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	if data.EiBitmap != nil {
		if !data.EiBitmap.IsNull() {
			body, _ = sjson.SetRaw(body, pathRoot+`EIBitmap`, data.EiBitmap.ToBody(ctx, ""))
		}
	}
	if !data.EiCustomUrl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EICustomURL`, data.EiCustomUrl.ValueString())
	}
	if !data.EixPath.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EIXPath`, data.EixPath.ValueString())
	}
	if !data.EiSignerDnValcred.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EISignerDNValcred`, data.EiSignerDnValcred.ValueString())
	}
	if !data.EiCookieName.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EICookieName`, data.EiCookieName.ValueString())
	}
	if !data.EiBasicAuthRealm.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EIBasicAuthRealm`, data.EiBasicAuthRealm.ValueString())
	}
	if !data.EiUseWsSec.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EIUseWSSec`, tfutils.StringFromBool(data.EiUseWsSec, ""))
	}
	if !data.EiMetadata.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EIMetadata`, data.EiMetadata.ValueString())
	}
	if !data.EiAllowRemoteTokenReference.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EIAllowRemoteTokenReference`, tfutils.StringFromBool(data.EiAllowRemoteTokenReference, ""))
	}
	if !data.EiRemoteTokenProcessService.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EIRemoteTokenProcessService`, data.EiRemoteTokenProcessService.ValueString())
	}
	if !data.EiPasswordRetrievalMechanism.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EIPasswordRetrievalMechanism`, data.EiPasswordRetrievalMechanism.ValueString())
	}
	if !data.EiPasswordRetrievalCustomUrl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EIPasswordRetrievalCustomURL`, data.EiPasswordRetrievalCustomUrl.ValueString())
	}
	if !data.EiPasswordRetrievalAaaInfoUrl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EIPasswordRetrievalAAAInfoURL`, data.EiPasswordRetrievalAaaInfoUrl.ValueString())
	}
	if !data.EiFormsLoginPolicy.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EIFormsLoginPolicy`, data.EiFormsLoginPolicy.ValueString())
	}
	if !data.EioAuthClientGroup.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EIOAuthClientGroup`, data.EioAuthClientGroup.ValueString())
	}
	if !data.EisslClientConfigType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EISSLClientConfigType`, data.EisslClientConfigType.ValueString())
	}
	if !data.EisslClientProfile.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EISSLClientProfile`, data.EisslClientProfile.ValueString())
	}
	if !data.EijwtValidator.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EIJWTValidator`, data.EijwtValidator.ValueString())
	}
	if !data.EiSocialLoginPolicy.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EISocialLoginPolicy`, data.EiSocialLoginPolicy.ValueString())
	}
	if !data.EisamlResponseValCred.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EISAMLResponseValCred`, data.EisamlResponseValCred.ValueString())
	}
	return body
}

func (data *DmAAAPExtractIdentity) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `EIBitmap`); value.Exists() {
		data.EiBitmap = &DmAAAPEIBitmap{}
		data.EiBitmap.FromBody(ctx, "", value)
	} else {
		data.EiBitmap = nil
	}
	if value := res.Get(pathRoot + `EICustomURL`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.EiCustomUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EiCustomUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `EIXPath`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.EixPath = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EixPath = types.StringNull()
	}
	if value := res.Get(pathRoot + `EISignerDNValcred`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.EiSignerDnValcred = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EiSignerDnValcred = types.StringNull()
	}
	if value := res.Get(pathRoot + `EICookieName`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.EiCookieName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EiCookieName = types.StringNull()
	}
	if value := res.Get(pathRoot + `EIBasicAuthRealm`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.EiBasicAuthRealm = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EiBasicAuthRealm = types.StringValue("login")
	}
	if value := res.Get(pathRoot + `EIUseWSSec`); value.Exists() {
		data.EiUseWsSec = tfutils.BoolFromString(value.String())
	} else {
		data.EiUseWsSec = types.BoolNull()
	}
	if value := res.Get(pathRoot + `EIMetadata`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.EiMetadata = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EiMetadata = types.StringNull()
	}
	if value := res.Get(pathRoot + `EIAllowRemoteTokenReference`); value.Exists() {
		data.EiAllowRemoteTokenReference = tfutils.BoolFromString(value.String())
	} else {
		data.EiAllowRemoteTokenReference = types.BoolNull()
	}
	if value := res.Get(pathRoot + `EIRemoteTokenProcessService`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.EiRemoteTokenProcessService = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EiRemoteTokenProcessService = types.StringNull()
	}
	if value := res.Get(pathRoot + `EIPasswordRetrievalMechanism`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.EiPasswordRetrievalMechanism = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EiPasswordRetrievalMechanism = types.StringValue("xmlfile")
	}
	if value := res.Get(pathRoot + `EIPasswordRetrievalCustomURL`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.EiPasswordRetrievalCustomUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EiPasswordRetrievalCustomUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `EIPasswordRetrievalAAAInfoURL`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.EiPasswordRetrievalAaaInfoUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EiPasswordRetrievalAaaInfoUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `EIFormsLoginPolicy`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.EiFormsLoginPolicy = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EiFormsLoginPolicy = types.StringNull()
	}
	if value := res.Get(pathRoot + `EIOAuthClientGroup`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.EioAuthClientGroup = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EioAuthClientGroup = types.StringNull()
	}
	if value := res.Get(pathRoot + `EISSLClientConfigType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.EisslClientConfigType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EisslClientConfigType = types.StringValue("proxy")
	}
	if value := res.Get(pathRoot + `EISSLClientProfile`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.EisslClientProfile = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EisslClientProfile = types.StringNull()
	}
	if value := res.Get(pathRoot + `EIJWTValidator`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.EijwtValidator = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EijwtValidator = types.StringNull()
	}
	if value := res.Get(pathRoot + `EISocialLoginPolicy`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.EiSocialLoginPolicy = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EiSocialLoginPolicy = types.StringNull()
	}
	if value := res.Get(pathRoot + `EISAMLResponseValCred`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.EisamlResponseValCred = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EisamlResponseValCred = types.StringNull()
	}
}

func (data *DmAAAPExtractIdentity) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `EIBitmap`); value.Exists() {
		data.EiBitmap.UpdateFromBody(ctx, "", value)
	} else {
		data.EiBitmap = nil
	}
	if value := res.Get(pathRoot + `EICustomURL`); value.Exists() && !data.EiCustomUrl.IsNull() {
		data.EiCustomUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EiCustomUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `EIXPath`); value.Exists() && !data.EixPath.IsNull() {
		data.EixPath = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EixPath = types.StringNull()
	}
	if value := res.Get(pathRoot + `EISignerDNValcred`); value.Exists() && !data.EiSignerDnValcred.IsNull() {
		data.EiSignerDnValcred = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EiSignerDnValcred = types.StringNull()
	}
	if value := res.Get(pathRoot + `EICookieName`); value.Exists() && !data.EiCookieName.IsNull() {
		data.EiCookieName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EiCookieName = types.StringNull()
	}
	if value := res.Get(pathRoot + `EIBasicAuthRealm`); value.Exists() && !data.EiBasicAuthRealm.IsNull() {
		data.EiBasicAuthRealm = tfutils.ParseStringFromGJSON(value)
	} else if data.EiBasicAuthRealm.ValueString() != "login" {
		data.EiBasicAuthRealm = types.StringNull()
	}
	if value := res.Get(pathRoot + `EIUseWSSec`); value.Exists() && !data.EiUseWsSec.IsNull() {
		data.EiUseWsSec = tfutils.BoolFromString(value.String())
	} else if data.EiUseWsSec.ValueBool() {
		data.EiUseWsSec = types.BoolNull()
	}
	if value := res.Get(pathRoot + `EIMetadata`); value.Exists() && !data.EiMetadata.IsNull() {
		data.EiMetadata = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EiMetadata = types.StringNull()
	}
	if value := res.Get(pathRoot + `EIAllowRemoteTokenReference`); value.Exists() && !data.EiAllowRemoteTokenReference.IsNull() {
		data.EiAllowRemoteTokenReference = tfutils.BoolFromString(value.String())
	} else if data.EiAllowRemoteTokenReference.ValueBool() {
		data.EiAllowRemoteTokenReference = types.BoolNull()
	}
	if value := res.Get(pathRoot + `EIRemoteTokenProcessService`); value.Exists() && !data.EiRemoteTokenProcessService.IsNull() {
		data.EiRemoteTokenProcessService = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EiRemoteTokenProcessService = types.StringNull()
	}
	if value := res.Get(pathRoot + `EIPasswordRetrievalMechanism`); value.Exists() && !data.EiPasswordRetrievalMechanism.IsNull() {
		data.EiPasswordRetrievalMechanism = tfutils.ParseStringFromGJSON(value)
	} else if data.EiPasswordRetrievalMechanism.ValueString() != "xmlfile" {
		data.EiPasswordRetrievalMechanism = types.StringNull()
	}
	if value := res.Get(pathRoot + `EIPasswordRetrievalCustomURL`); value.Exists() && !data.EiPasswordRetrievalCustomUrl.IsNull() {
		data.EiPasswordRetrievalCustomUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EiPasswordRetrievalCustomUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `EIPasswordRetrievalAAAInfoURL`); value.Exists() && !data.EiPasswordRetrievalAaaInfoUrl.IsNull() {
		data.EiPasswordRetrievalAaaInfoUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EiPasswordRetrievalAaaInfoUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `EIFormsLoginPolicy`); value.Exists() && !data.EiFormsLoginPolicy.IsNull() {
		data.EiFormsLoginPolicy = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EiFormsLoginPolicy = types.StringNull()
	}
	if value := res.Get(pathRoot + `EIOAuthClientGroup`); value.Exists() && !data.EioAuthClientGroup.IsNull() {
		data.EioAuthClientGroup = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EioAuthClientGroup = types.StringNull()
	}
	if value := res.Get(pathRoot + `EISSLClientConfigType`); value.Exists() && !data.EisslClientConfigType.IsNull() {
		data.EisslClientConfigType = tfutils.ParseStringFromGJSON(value)
	} else if data.EisslClientConfigType.ValueString() != "proxy" {
		data.EisslClientConfigType = types.StringNull()
	}
	if value := res.Get(pathRoot + `EISSLClientProfile`); value.Exists() && !data.EisslClientProfile.IsNull() {
		data.EisslClientProfile = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EisslClientProfile = types.StringNull()
	}
	if value := res.Get(pathRoot + `EIJWTValidator`); value.Exists() && !data.EijwtValidator.IsNull() {
		data.EijwtValidator = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EijwtValidator = types.StringNull()
	}
	if value := res.Get(pathRoot + `EISocialLoginPolicy`); value.Exists() && !data.EiSocialLoginPolicy.IsNull() {
		data.EiSocialLoginPolicy = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EiSocialLoginPolicy = types.StringNull()
	}
	if value := res.Get(pathRoot + `EISAMLResponseValCred`); value.Exists() && !data.EisamlResponseValCred.IsNull() {
		data.EisamlResponseValCred = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EisamlResponseValCred = types.StringNull()
	}
}

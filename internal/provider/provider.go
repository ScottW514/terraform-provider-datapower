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

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/int32validator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/client"
)

type DatapowerProvider struct {
	version string
}

type DatapowerProviderModel struct {
	Hostname types.String `tfsdk:"hostname"`
	Port     types.Int32  `tfsdk:"port"`
	Username types.String `tfsdk:"username"`
	Password types.String `tfsdk:"password"`
	Insecure types.Bool   `tfsdk:"insecure"`
}

func (p *DatapowerProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "datapower"
	resp.Version = p.version
}

func (p *DatapowerProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"hostname": schema.StringAttribute{
				MarkdownDescription: "Hostname of the Datapower instance. Can be set via DP_HOSTNAME evnvironment variable. Defaults to `localhost`",
				Optional:            true,
			},
			"port": schema.Int32Attribute{
				MarkdownDescription: "Port for REST API calls. Can be set via DP_PORT evnvironment variable. Defaults to `5554`.",
				Optional:            true,
				Validators: []validator.Int32{
					int32validator.Between(1024, 65535),
				},
			},
			"username": schema.StringAttribute{
				MarkdownDescription: "Username for the Datapower instance. Can be set via DP_USERNAME evnvironment variable.",
				Optional:            true,
			},
			"password": schema.StringAttribute{
				MarkdownDescription: "Password for the Datapower instance. Can be set via DP_PASSWORD evnvironment variable.",
				Optional:            true,
				Sensitive:           true,
			},
			"insecure": schema.BoolAttribute{
				MarkdownDescription: "Allow insecure HTTPS client. Can be set via DP_INSECURE evnvironment variable. Defaults to `false`.",
				Optional:            true,
			},
		},
	}
}

func (p *DatapowerProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var config client.DatapowerClientConfig
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Create a new Datapower client and set it to the provider client
	c, err := client.NewClient(&config)
	if err != nil {
		resp.Diagnostics.AddError(
			"Client Error",
			"When creating datapower client: "+err.Error(),
		)
		return
	}

	resp.DataSourceData = &c
	resp.ResourceData = &c
}

func (p *DatapowerProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewFileResource,
		NewAAAJWTGeneratorResource,
		NewAAAJWTValidatorResource,
		NewAAAPolicyResource,
		NewAccessControlListResource,
		NewAccessProfileResource,
		NewAuditLogResource,
		NewCompileOptionsPolicyResource,
		NewConfigDeploymentPolicyResource,
		NewCookieAttributePolicyResource,
		NewCountMonitorResource,
		NewCryptoCertificateResource,
		NewCryptoFWCredResource,
		NewCryptoIdentCredResource,
		NewCryptoKerberosKeytabResource,
		NewCryptoKeyResource,
		NewCryptoSSKeyResource,
		NewCryptoValCredResource,
		NewDeploymentPolicyParametersBindingResource,
		NewDomainResource,
		NewDomainAvailabilityResource,
		NewDurationMonitorResource,
		NewFTPQuoteCommandsResource,
		NewFilterActionResource,
		NewFormsLoginPolicyResource,
		NewHTTPInputConversionMapResource,
		NewHTTPSSourceProtocolHandlerResource,
		NewHTTPSourceProtocolHandlerResource,
		NewHTTPUserAgentResource,
		NewJOSERecipientIdentifierResource,
		NewJOSESignatureIdentifierResource,
		NewJSONSettingsResource,
		NewJWEHeaderResource,
		NewJWERecipientResource,
		NewJWSSignatureResource,
		NewLDAPConnectionPoolResource,
		NewLDAPSearchParametersResource,
		NewLoadBalancerGroupResource,
		NewLogLabelResource,
		NewMPGWErrorActionResource,
		NewMPGWErrorHandlingPolicyResource,
		NewMatchingResource,
		NewMessageContentFiltersResource,
		NewMessageMatchingResource,
		NewMessageTypeResource,
		NewMultiProtocolGatewayResource,
		NewOAuthSupportedClientResource,
		NewOAuthSupportedClientGroupResource,
		NewParseSettingsResource,
		NewPasswordAliasResource,
		NewPolicyAttachmentsResource,
		NewPolicyParametersResource,
		NewProcessingMetadataResource,
		NewSAMLAttributesResource,
		NewSQLDataSourceResource,
		NewSSHClientProfileResource,
		NewSSLClientProfileResource,
		NewSSLSNIMappingResource,
		NewSSLSNIServerProfileResource,
		NewSSLServerProfileResource,
		NewSocialLoginPolicyResource,
		NewStylePolicyResource,
		NewStylePolicyActionResource,
		NewStylePolicyRuleResource,
		NewTAMResource,
		NewTFIMEndpointResource,
		NewURLMapResource,
		NewURLRefreshPolicyResource,
		NewURLRewritePolicyResource,
		NewUserResource,
		NewUserGroupResource,
		NewWCCServiceResource,
		NewWSRRSavedSearchSubscriptionResource,
		NewWSRRServerResource,
		NewWSRRSubscriptionResource,
		NewWebServiceMonitorResource,
		NewXACMLPDPResource,
		NewXMLManagerResource,
		NewZosNSSClientResource,
	}
}

func (p *DatapowerProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewAAAJWTGeneratorDataSource,
		NewAAAJWTValidatorDataSource,
		NewAAAPolicyDataSource,
		NewAccessControlListDataSource,
		NewAccessProfileDataSource,
		NewAuditLogDataSource,
		NewCompileOptionsPolicyDataSource,
		NewConfigDeploymentPolicyDataSource,
		NewCookieAttributePolicyDataSource,
		NewCountMonitorDataSource,
		NewCryptoCertificateDataSource,
		NewCryptoFWCredDataSource,
		NewCryptoIdentCredDataSource,
		NewCryptoKerberosKeytabDataSource,
		NewCryptoKeyDataSource,
		NewCryptoSSKeyDataSource,
		NewCryptoValCredDataSource,
		NewDeploymentPolicyParametersBindingDataSource,
		NewDomainDataSource,
		NewDomainAvailabilityDataSource,
		NewDurationMonitorDataSource,
		NewFTPQuoteCommandsDataSource,
		NewFilterActionDataSource,
		NewFormsLoginPolicyDataSource,
		NewHTTPInputConversionMapDataSource,
		NewHTTPSSourceProtocolHandlerDataSource,
		NewHTTPSourceProtocolHandlerDataSource,
		NewHTTPUserAgentDataSource,
		NewJOSERecipientIdentifierDataSource,
		NewJOSESignatureIdentifierDataSource,
		NewJSONSettingsDataSource,
		NewJWEHeaderDataSource,
		NewJWERecipientDataSource,
		NewJWSSignatureDataSource,
		NewLDAPConnectionPoolDataSource,
		NewLDAPSearchParametersDataSource,
		NewLoadBalancerGroupDataSource,
		NewLogLabelDataSource,
		NewMPGWErrorActionDataSource,
		NewMPGWErrorHandlingPolicyDataSource,
		NewMatchingDataSource,
		NewMessageContentFiltersDataSource,
		NewMessageMatchingDataSource,
		NewMessageTypeDataSource,
		NewMultiProtocolGatewayDataSource,
		NewOAuthSupportedClientDataSource,
		NewOAuthSupportedClientGroupDataSource,
		NewParseSettingsDataSource,
		NewPasswordAliasDataSource,
		NewPolicyAttachmentsDataSource,
		NewPolicyParametersDataSource,
		NewProcessingMetadataDataSource,
		NewSAMLAttributesDataSource,
		NewSQLDataSourceDataSource,
		NewSSHClientProfileDataSource,
		NewSSLClientProfileDataSource,
		NewSSLSNIMappingDataSource,
		NewSSLSNIServerProfileDataSource,
		NewSSLServerProfileDataSource,
		NewSocialLoginPolicyDataSource,
		NewStylePolicyDataSource,
		NewStylePolicyActionDataSource,
		NewStylePolicyRuleDataSource,
		NewTAMDataSource,
		NewTFIMEndpointDataSource,
		NewURLMapDataSource,
		NewURLRefreshPolicyDataSource,
		NewURLRewritePolicyDataSource,
		NewUserDataSource,
		NewUserGroupDataSource,
		NewWCCServiceDataSource,
		NewWSRRSavedSearchSubscriptionDataSource,
		NewWSRRServerDataSource,
		NewWSRRSubscriptionDataSource,
		NewWebServiceMonitorDataSource,
		NewXACMLPDPDataSource,
		NewXMLManagerDataSource,
		NewZosNSSClientDataSource,
	}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &DatapowerProvider{
			version: version,
		}
	}
}

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

type OAuthSupportedClient struct {
	Id                         types.String                `tfsdk:"id"`
	AppDomain                  types.String                `tfsdk:"app_domain"`
	UserSummary                types.String                `tfsdk:"user_summary"`
	Customized                 types.Bool                  `tfsdk:"customized"`
	CustomizedProcessUrl       types.String                `tfsdk:"customized_process_url"`
	OauthRole                  *DmOAuthRole                `tfsdk:"oauth_role"`
	AzGrant                    *DmOAuthAZGrantType         `tfsdk:"az_grant"`
	ClientType                 types.String                `tfsdk:"client_type"`
	CheckClientCredential      types.Bool                  `tfsdk:"check_client_credential"`
	UseValidationUrl           types.Bool                  `tfsdk:"use_validation_url"`
	ClientAuthenMethod         types.String                `tfsdk:"client_authen_method"`
	ClientValCred              types.String                `tfsdk:"client_val_cred"`
	GenerateClientSecret       types.Bool                  `tfsdk:"generate_client_secret"`
	ClientSecretWo             types.String                `tfsdk:"client_secret_wo"`
	ClientSecretWoVersion      types.Int64                 `tfsdk:"client_secret_wo_version"`
	Caching                    types.String                `tfsdk:"caching"`
	ValidationUrl              types.String                `tfsdk:"validation_url"`
	ValidationFeatures         *DmValidationFeatures       `tfsdk:"validation_features"`
	RedirectUri                types.List                  `tfsdk:"redirect_uri"`
	CustomScopeCheck           types.Bool                  `tfsdk:"custom_scope_check"`
	Scope                      types.String                `tfsdk:"scope"`
	ScopeUrl                   types.String                `tfsdk:"scope_url"`
	DefaultScope               types.String                `tfsdk:"default_scope"`
	TokenSecret                types.String                `tfsdk:"token_secret"`
	LocalAzPageUrl             types.String                `tfsdk:"local_az_page_url"`
	DpStateLifeTime            types.Int64                 `tfsdk:"dp_state_life_time"`
	AuCodeLifeTime             types.Int64                 `tfsdk:"au_code_life_time"`
	AccessTokenLifeTime        types.Int64                 `tfsdk:"access_token_life_time"`
	RefreshTokenAllowed        types.Int64                 `tfsdk:"refresh_token_allowed"`
	RefreshTokenLifeTime       types.Int64                 `tfsdk:"refresh_token_life_time"`
	MaxConsentLifeTime         types.Int64                 `tfsdk:"max_consent_life_time"`
	CustomResourceOwner        types.Bool                  `tfsdk:"custom_resource_owner"`
	ResourceOwnerUrl           types.String                `tfsdk:"resource_owner_url"`
	AdditionalOauthProcessUrl  types.String                `tfsdk:"additional_oauth_process_url"`
	RsSetHeader                *DmOAuthRSSetHeader         `tfsdk:"rs_set_header"`
	ValidationUrlSslClientType types.String                `tfsdk:"validation_url_ssl_client_type"`
	ValidationUrlSslClient     types.String                `tfsdk:"validation_url_ssl_client"`
	JwtGrantValidator          types.String                `tfsdk:"jwt_grant_validator"`
	ClientJwtValidator         types.String                `tfsdk:"client_jwt_validator"`
	OidcIdTokenGenerator       types.String                `tfsdk:"oidc_id_token_generator"`
	OauthFeatures              *DmOAuthFeatures            `tfsdk:"oauth_features"`
	DependencyActions          []*actions.DependencyAction `tfsdk:"dependency_actions"`
}
type OAuthSupportedClientWO struct {
	Id                         types.String                `tfsdk:"id"`
	AppDomain                  types.String                `tfsdk:"app_domain"`
	UserSummary                types.String                `tfsdk:"user_summary"`
	Customized                 types.Bool                  `tfsdk:"customized"`
	CustomizedProcessUrl       types.String                `tfsdk:"customized_process_url"`
	OauthRole                  *DmOAuthRole                `tfsdk:"oauth_role"`
	AzGrant                    *DmOAuthAZGrantType         `tfsdk:"az_grant"`
	ClientType                 types.String                `tfsdk:"client_type"`
	CheckClientCredential      types.Bool                  `tfsdk:"check_client_credential"`
	UseValidationUrl           types.Bool                  `tfsdk:"use_validation_url"`
	ClientAuthenMethod         types.String                `tfsdk:"client_authen_method"`
	ClientValCred              types.String                `tfsdk:"client_val_cred"`
	GenerateClientSecret       types.Bool                  `tfsdk:"generate_client_secret"`
	Caching                    types.String                `tfsdk:"caching"`
	ValidationUrl              types.String                `tfsdk:"validation_url"`
	ValidationFeatures         *DmValidationFeatures       `tfsdk:"validation_features"`
	RedirectUri                types.List                  `tfsdk:"redirect_uri"`
	CustomScopeCheck           types.Bool                  `tfsdk:"custom_scope_check"`
	Scope                      types.String                `tfsdk:"scope"`
	ScopeUrl                   types.String                `tfsdk:"scope_url"`
	DefaultScope               types.String                `tfsdk:"default_scope"`
	TokenSecret                types.String                `tfsdk:"token_secret"`
	LocalAzPageUrl             types.String                `tfsdk:"local_az_page_url"`
	DpStateLifeTime            types.Int64                 `tfsdk:"dp_state_life_time"`
	AuCodeLifeTime             types.Int64                 `tfsdk:"au_code_life_time"`
	AccessTokenLifeTime        types.Int64                 `tfsdk:"access_token_life_time"`
	RefreshTokenAllowed        types.Int64                 `tfsdk:"refresh_token_allowed"`
	RefreshTokenLifeTime       types.Int64                 `tfsdk:"refresh_token_life_time"`
	MaxConsentLifeTime         types.Int64                 `tfsdk:"max_consent_life_time"`
	CustomResourceOwner        types.Bool                  `tfsdk:"custom_resource_owner"`
	ResourceOwnerUrl           types.String                `tfsdk:"resource_owner_url"`
	AdditionalOauthProcessUrl  types.String                `tfsdk:"additional_oauth_process_url"`
	RsSetHeader                *DmOAuthRSSetHeader         `tfsdk:"rs_set_header"`
	ValidationUrlSslClientType types.String                `tfsdk:"validation_url_ssl_client_type"`
	ValidationUrlSslClient     types.String                `tfsdk:"validation_url_ssl_client"`
	JwtGrantValidator          types.String                `tfsdk:"jwt_grant_validator"`
	ClientJwtValidator         types.String                `tfsdk:"client_jwt_validator"`
	OidcIdTokenGenerator       types.String                `tfsdk:"oidc_id_token_generator"`
	OauthFeatures              *DmOAuthFeatures            `tfsdk:"oauth_features"`
	DependencyActions          []*actions.DependencyAction `tfsdk:"dependency_actions"`
}

var OAuthSupportedClientCustomizedProcessUrlCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "customized",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"true"},
}

var OAuthSupportedClientCustomizedProcessUrlIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var OAuthSupportedClientOAuthRoleCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "customized",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"false"},
}

var OAuthSupportedClientOAuthRoleIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var OAuthSupportedClientAZGrantCondVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "customized",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"true"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "oauth_role",
			AttrType:    "DmOAuthRole",
			AttrDefault: "",
			Value:       []string{"azsvr"},
		},
	},
}

var OAuthSupportedClientAZGrantIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var OAuthSupportedClientClientTypeCondVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "customized",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"true"},
		},
		{
			Evaluation: "logical-and",
			Conditions: []validators.Evaluation{

				{
					Evaluation:  "property-value-in-list",
					Attribute:   "customized",
					AttrType:    "Bool",
					AttrDefault: "false",
					Value:       []string{"false"},
				},

				{
					Evaluation:  "property-value-in-list",
					Attribute:   "oauth_role",
					AttrType:    "DmOAuthRole",
					AttrDefault: "",
					Value:       []string{"azsvr"},
				},

				{
					Evaluation:  "property-value-in-list",
					Attribute:   "az_grant",
					AttrType:    "DmOAuthAZGrantType",
					AttrDefault: "",
					Value:       []string{"code", "implicit", "password", "jwt"},
				},
			},
		},
		{
			Evaluation: "logical-and",
			Conditions: []validators.Evaluation{

				{
					Evaluation:  "property-value-in-list",
					Attribute:   "customized",
					AttrType:    "Bool",
					AttrDefault: "false",
					Value:       []string{"false"},
				},

				{
					Evaluation:  "property-value-in-list",
					Attribute:   "oauth_role",
					AttrType:    "DmOAuthRole",
					AttrDefault: "",
					Value:       []string{"rssvr"},
				},

				{
					Evaluation: "logical-or",
					Conditions: []validators.Evaluation{

						{
							Evaluation:  "property-value-in-list",
							Attribute:   "use_validation_url",
							AttrType:    "Bool",
							AttrDefault: "false",
							Value:       []string{"true"},
						},

						{
							Evaluation:  "property-value-in-list",
							Attribute:   "check_client_credential",
							AttrType:    "Bool",
							AttrDefault: "false",
							Value:       []string{"true"},
						},
					},
				},
			},
		},
	},
}

var OAuthSupportedClientClientTypeIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var OAuthSupportedClientCheckClientCredentialIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "customized",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"true"},
		},
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "oauth_role",
			AttrType:    "DmOAuthRole",
			AttrDefault: "",
			Value:       []string{"rssvr"},
		},
	},
}

var OAuthSupportedClientUseValidationUrlIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "customized",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"true"},
		},
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "oauth_role",
			AttrType:    "DmOAuthRole",
			AttrDefault: "",
			Value:       []string{"rssvr"},
		},
	},
}

var OAuthSupportedClientClientAuthenMethodCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "customized",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"false"},
		},
		{
			Evaluation: "logical-or",
			Conditions: []validators.Evaluation{

				{
					Evaluation:  "property-value-in-list",
					Attribute:   "client_type",
					AttrType:    "String",
					AttrDefault: "confidential",
					Value:       []string{"confidential"},
				},

				{
					Evaluation: "logical-and",
					Conditions: []validators.Evaluation{

						{
							Evaluation:  "property-value-in-list",
							Attribute:   "oauth_role",
							AttrType:    "DmOAuthRole",
							AttrDefault: "",
							Value:       []string{"azsvr"},
						},

						{
							Evaluation:  "property-value-in-list",
							Attribute:   "az_grant",
							AttrType:    "DmOAuthAZGrantType",
							AttrDefault: "",
							Value:       []string{"client"},
						},
					},
				},
			},
		},
		{
			Evaluation: "logical-or",
			Conditions: []validators.Evaluation{

				{
					Evaluation: "logical-and",
					Conditions: []validators.Evaluation{

						{
							Evaluation:  "property-value-in-list",
							Attribute:   "oauth_role",
							AttrType:    "DmOAuthRole",
							AttrDefault: "",
							Value:       []string{"azsvr"},
						},

						{
							Evaluation:  "property-value-in-list",
							Attribute:   "az_grant",
							AttrType:    "DmOAuthAZGrantType",
							AttrDefault: "",
							Value:       []string{"code", "implicit", "password", "client", "jwt"},
						},
					},
				},

				{
					Evaluation: "logical-and",
					Conditions: []validators.Evaluation{

						{
							Evaluation:  "property-value-in-list",
							Attribute:   "oauth_role",
							AttrType:    "DmOAuthRole",
							AttrDefault: "",
							Value:       []string{"rssvr"},
						},

						{
							Evaluation: "logical-or",
							Conditions: []validators.Evaluation{

								{
									Evaluation:  "property-value-in-list",
									Attribute:   "use_validation_url",
									AttrType:    "Bool",
									AttrDefault: "false",
									Value:       []string{"true"},
								},

								{
									Evaluation:  "property-value-in-list",
									Attribute:   "check_client_credential",
									AttrType:    "Bool",
									AttrDefault: "false",
									Value:       []string{"true"},
								},
							},
						},
					},
				},
			},
		},
	},
}

var OAuthSupportedClientClientAuthenMethodIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var OAuthSupportedClientClientValCredCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "customized",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"false"},
		},
		{
			Evaluation: "logical-or",
			Conditions: []validators.Evaluation{

				{
					Evaluation:  "property-value-in-list",
					Attribute:   "client_type",
					AttrType:    "String",
					AttrDefault: "confidential",
					Value:       []string{"confidential"},
				},

				{
					Evaluation: "logical-and",
					Conditions: []validators.Evaluation{

						{
							Evaluation:  "property-value-in-list",
							Attribute:   "oauth_role",
							AttrType:    "DmOAuthRole",
							AttrDefault: "",
							Value:       []string{"azsvr"},
						},

						{
							Evaluation:  "property-value-in-list",
							Attribute:   "az_grant",
							AttrType:    "DmOAuthAZGrantType",
							AttrDefault: "",
							Value:       []string{"client"},
						},
					},
				},
			},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "client_authen_method",
			AttrType:    "String",
			AttrDefault: "secret",
			Value:       []string{"ssl"},
		},
		{
			Evaluation: "logical-or",
			Conditions: []validators.Evaluation{

				{
					Evaluation: "logical-and",
					Conditions: []validators.Evaluation{

						{
							Evaluation:  "property-value-in-list",
							Attribute:   "oauth_role",
							AttrType:    "DmOAuthRole",
							AttrDefault: "",
							Value:       []string{"azsvr"},
						},

						{
							Evaluation:  "property-value-in-list",
							Attribute:   "az_grant",
							AttrType:    "DmOAuthAZGrantType",
							AttrDefault: "",
							Value:       []string{"code", "implicit", "password", "client", "jwt"},
						},
					},
				},

				{
					Evaluation: "logical-and",
					Conditions: []validators.Evaluation{

						{
							Evaluation:  "property-value-in-list",
							Attribute:   "oauth_role",
							AttrType:    "DmOAuthRole",
							AttrDefault: "",
							Value:       []string{"rssvr"},
						},

						{
							Evaluation:  "property-value-not-in-list",
							Attribute:   "oauth_role",
							AttrType:    "DmOAuthRole",
							AttrDefault: "",
							Value:       []string{"azsvr"},
						},

						{
							Evaluation:  "property-value-in-list",
							Attribute:   "check_client_credential",
							AttrType:    "Bool",
							AttrDefault: "false",
							Value:       []string{"true"},
						},
					},
				},
			},
		},
	},
}

var OAuthSupportedClientClientValCredIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var OAuthSupportedClientGenerateClientSecretCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "customized",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"false"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "oauth_role",
			AttrType:    "DmOAuthRole",
			AttrDefault: "",
			Value:       []string{"azsvr"},
		},
		{
			Evaluation: "logical-or",
			Conditions: []validators.Evaluation{

				{
					Evaluation:  "property-value-in-list",
					Attribute:   "client_type",
					AttrType:    "String",
					AttrDefault: "confidential",
					Value:       []string{"confidential"},
				},

				{
					Evaluation: "logical-and",
					Conditions: []validators.Evaluation{

						{
							Evaluation:  "property-value-in-list",
							Attribute:   "oauth_role",
							AttrType:    "DmOAuthRole",
							AttrDefault: "",
							Value:       []string{"azsvr"},
						},

						{
							Evaluation:  "property-value-in-list",
							Attribute:   "az_grant",
							AttrType:    "DmOAuthAZGrantType",
							AttrDefault: "",
							Value:       []string{"client"},
						},
					},
				},
			},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "client_secret_wo",
			AttrType:    "String",
			AttrDefault: "",
			Value:       []string{""},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "az_grant",
			AttrType:    "DmOAuthAZGrantType",
			AttrDefault: "",
			Value:       []string{"code", "implicit", "password", "client", "jwt"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "client_authen_method",
			AttrType:    "String",
			AttrDefault: "secret",
			Value:       []string{"secret"},
		},
	},
}

var OAuthSupportedClientGenerateClientSecretIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var OAuthSupportedClientClientSecretCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "customized",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"false"},
		},
		{
			Evaluation: "logical-or",
			Conditions: []validators.Evaluation{

				{
					Evaluation:  "property-value-in-list",
					Attribute:   "client_type",
					AttrType:    "String",
					AttrDefault: "confidential",
					Value:       []string{"confidential"},
				},

				{
					Evaluation: "logical-and",
					Conditions: []validators.Evaluation{

						{
							Evaluation:  "property-value-in-list",
							Attribute:   "oauth_role",
							AttrType:    "DmOAuthRole",
							AttrDefault: "",
							Value:       []string{"azsvr"},
						},

						{
							Evaluation:  "property-value-in-list",
							Attribute:   "az_grant",
							AttrType:    "DmOAuthAZGrantType",
							AttrDefault: "",
							Value:       []string{"client"},
						},
					},
				},
			},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "client_authen_method",
			AttrType:    "String",
			AttrDefault: "secret",
			Value:       []string{"secret"},
		},
		{
			Evaluation: "logical-or",
			Conditions: []validators.Evaluation{

				{
					Evaluation: "logical-and",
					Conditions: []validators.Evaluation{

						{
							Evaluation:  "property-value-in-list",
							Attribute:   "oauth_role",
							AttrType:    "DmOAuthRole",
							AttrDefault: "",
							Value:       []string{"azsvr"},
						},

						{
							Evaluation:  "property-value-in-list",
							Attribute:   "generate_client_secret",
							AttrType:    "Bool",
							AttrDefault: "true",
							Value:       []string{"false"},
						},

						{
							Evaluation:  "property-value-in-list",
							Attribute:   "az_grant",
							AttrType:    "DmOAuthAZGrantType",
							AttrDefault: "",
							Value:       []string{"code", "implicit", "password", "client", "jwt", "saml20bearer"},
						},
					},
				},

				{
					Evaluation: "logical-and",
					Conditions: []validators.Evaluation{

						{
							Evaluation:  "property-value-in-list",
							Attribute:   "oauth_role",
							AttrType:    "DmOAuthRole",
							AttrDefault: "",
							Value:       []string{"rssvr"},
						},

						{
							Evaluation:  "property-value-not-in-list",
							Attribute:   "oauth_role",
							AttrType:    "DmOAuthRole",
							AttrDefault: "",
							Value:       []string{"azsvr"},
						},

						{
							Evaluation: "logical-or",
							Conditions: []validators.Evaluation{

								{
									Evaluation:  "property-value-in-list",
									Attribute:   "check_client_credential",
									AttrType:    "Bool",
									AttrDefault: "false",
									Value:       []string{"true"},
								},

								{
									Evaluation:  "property-value-in-list",
									Attribute:   "use_validation_url",
									AttrType:    "Bool",
									AttrDefault: "false",
									Value:       []string{"true"},
								},
							},
						},
					},
				},
			},
		},
	},
}

var OAuthSupportedClientClientSecretIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var OAuthSupportedClientCachingIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "customized",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"true"},
		},
		{
			Evaluation: "logical-and",
			Conditions: []validators.Evaluation{

				{
					Evaluation:  "property-value-not-in-list",
					Attribute:   "oauth_role",
					AttrType:    "DmOAuthRole",
					AttrDefault: "",
					Value:       []string{"azsvr"},
				},

				{
					Evaluation:  "property-value-in-list",
					Attribute:   "oauth_role",
					AttrType:    "DmOAuthRole",
					AttrDefault: "",
					Value:       []string{"rssvr"},
				},

				{
					Evaluation:  "property-value-in-list",
					Attribute:   "use_validation_url",
					AttrType:    "Bool",
					AttrDefault: "false",
					Value:       []string{"true"},
				},
			},
		},
	},
}

var OAuthSupportedClientValidationURLCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "customized",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"false"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "oauth_role",
			AttrType:    "DmOAuthRole",
			AttrDefault: "",
			Value:       []string{"rssvr"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "use_validation_url",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"true"},
		},
	},
}

var OAuthSupportedClientValidationURLIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var OAuthSupportedClientValidationFeaturesIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "customized",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"true"},
		},
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "oauth_role",
			AttrType:    "DmOAuthRole",
			AttrDefault: "",
			Value:       []string{"azsvr"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "az_grant",
			AttrType:    "DmOAuthAZGrantType",
			AttrDefault: "",
			Value:       []string{"novalidate"},
		},
	},
}

var OAuthSupportedClientRedirectURICondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "customized",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"false"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "oauth_role",
			AttrType:    "DmOAuthRole",
			AttrDefault: "",
			Value:       []string{"azsvr"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "az_grant",
			AttrType:    "DmOAuthAZGrantType",
			AttrDefault: "",
			Value:       []string{"implicit", "code"},
		},
	},
}

var OAuthSupportedClientRedirectURIIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "customized",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"true"},
		},
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "oauth_role",
			AttrType:    "DmOAuthRole",
			AttrDefault: "",
			Value:       []string{"azsvr"},
		},
		{
			Evaluation: "logical-and",
			Conditions: []validators.Evaluation{

				{
					Evaluation:  "property-value-in-list",
					Attribute:   "customized",
					AttrType:    "Bool",
					AttrDefault: "false",
					Value:       []string{"false"},
				},

				{
					Evaluation:  "property-value-not-in-list",
					Attribute:   "az_grant",
					AttrType:    "DmOAuthAZGrantType",
					AttrDefault: "",
					Value:       []string{"code", "implicit"},
				},
			},
		},
	},
}

var OAuthSupportedClientCustomScopeCheckIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "customized",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"true"},
		},
		{
			Evaluation: "logical-and",
			Conditions: []validators.Evaluation{

				{
					Evaluation:  "property-value-not-in-list",
					Attribute:   "oauth_role",
					AttrType:    "DmOAuthRole",
					AttrDefault: "",
					Value:       []string{"azsvr"},
				},

				{
					Evaluation:  "property-value-in-list",
					Attribute:   "use_validation_url",
					AttrType:    "Bool",
					AttrDefault: "false",
					Value:       []string{"true"},
				},
			},
		},
	},
}

var OAuthSupportedClientScopeCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "customized",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"false"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "custom_scope_check",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"false"},
		},
		{
			Evaluation: "logical-or",
			Conditions: []validators.Evaluation{

				{
					Evaluation:  "property-value-in-list",
					Attribute:   "oauth_role",
					AttrType:    "DmOAuthRole",
					AttrDefault: "",
					Value:       []string{"azsvr"},
				},

				{
					Evaluation:  "property-value-in-list",
					Attribute:   "use_validation_url",
					AttrType:    "Bool",
					AttrDefault: "false",
					Value:       []string{"false"},
				},
			},
		},
	},
}

var OAuthSupportedClientScopeIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var OAuthSupportedClientScopeUrlCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "customized",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"false"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "custom_scope_check",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"true"},
		},
		{
			Evaluation: "logical-or",
			Conditions: []validators.Evaluation{

				{
					Evaluation:  "property-value-in-list",
					Attribute:   "oauth_role",
					AttrType:    "DmOAuthRole",
					AttrDefault: "",
					Value:       []string{"azsvr"},
				},

				{
					Evaluation:  "property-value-in-list",
					Attribute:   "use_validation_url",
					AttrType:    "Bool",
					AttrDefault: "false",
					Value:       []string{"false"},
				},
			},
		},
	},
}

var OAuthSupportedClientScopeUrlIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var OAuthSupportedClientDefaultScopeIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "customized",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"true"},
		},
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "oauth_role",
			AttrType:    "DmOAuthRole",
			AttrDefault: "",
			Value:       []string{"azsvr"},
		},
	},
}

var OAuthSupportedClientTokenSecretCondVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation: "logical-and",
			Conditions: []validators.Evaluation{

				{
					Evaluation:  "property-value-in-list",
					Attribute:   "customized",
					AttrType:    "Bool",
					AttrDefault: "false",
					Value:       []string{"false"},
				},

				{
					Evaluation: "logical-or",
					Conditions: []validators.Evaluation{

						{
							Evaluation:  "property-value-in-list",
							Attribute:   "oauth_role",
							AttrType:    "DmOAuthRole",
							AttrDefault: "",
							Value:       []string{"azsvr"},
						},

						{
							Evaluation: "logical-and",
							Conditions: []validators.Evaluation{

								{
									Evaluation:  "property-value-not-in-list",
									Attribute:   "oauth_role",
									AttrType:    "DmOAuthRole",
									AttrDefault: "",
									Value:       []string{"azsrv"},
								},

								{
									Evaluation:  "property-value-in-list",
									Attribute:   "use_validation_url",
									AttrType:    "Bool",
									AttrDefault: "false",
									Value:       []string{"false"},
								},
							},
						},
					},
				},
			},
		},
		{
			Evaluation: "logical-and",
			Conditions: []validators.Evaluation{

				{
					Evaluation:  "property-value-in-list",
					Attribute:   "customized",
					AttrType:    "Bool",
					AttrDefault: "false",
					Value:       []string{"true"},
				},

				{
					Evaluation:  "property-value-in-list",
					Attribute:   "az_grant",
					AttrType:    "DmOAuthAZGrantType",
					AttrDefault: "",
					Value:       []string{"code", "implicit"},
				},

				{
					Evaluation: "logical-or",
					Conditions: []validators.Evaluation{

						{
							Evaluation:  "property-value-in-list",
							Attribute:   "oauth_role",
							AttrType:    "DmOAuthRole",
							AttrDefault: "",
							Value:       []string{"azsvr"},
						},

						{
							Evaluation: "logical-and",
							Conditions: []validators.Evaluation{

								{
									Evaluation:  "property-value-not-in-list",
									Attribute:   "oauth_role",
									AttrType:    "DmOAuthRole",
									AttrDefault: "",
									Value:       []string{"azsvr"},
								},

								{
									Evaluation:  "property-value-in-list",
									Attribute:   "use_validation_url",
									AttrType:    "Bool",
									AttrDefault: "false",
									Value:       []string{"false"},
								},
							},
						},
					},
				},
			},
		},
	},
}

var OAuthSupportedClientTokenSecretIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var OAuthSupportedClientLocalAZPageUrlCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation: "logical-or",
			Conditions: []validators.Evaluation{

				{
					Evaluation:  "property-value-in-list",
					Attribute:   "customized",
					AttrType:    "Bool",
					AttrDefault: "false",
					Value:       []string{"true"},
				},

				{
					Evaluation:  "property-value-in-list",
					Attribute:   "oauth_role",
					AttrType:    "DmOAuthRole",
					AttrDefault: "",
					Value:       []string{"azsvr"},
				},
			},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "az_grant",
			AttrType:    "DmOAuthAZGrantType",
			AttrDefault: "",
			Value:       []string{"code", "implicit"},
		},
	},
}

var OAuthSupportedClientLocalAZPageUrlIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var OAuthSupportedClientDPStateLifeTimeCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation: "logical-or",
			Conditions: []validators.Evaluation{

				{
					Evaluation:  "property-value-in-list",
					Attribute:   "customized",
					AttrType:    "Bool",
					AttrDefault: "false",
					Value:       []string{"true"},
				},

				{
					Evaluation:  "property-value-in-list",
					Attribute:   "oauth_role",
					AttrType:    "DmOAuthRole",
					AttrDefault: "",
					Value:       []string{"azsvr"},
				},
			},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "az_grant",
			AttrType:    "DmOAuthAZGrantType",
			AttrDefault: "",
			Value:       []string{"code", "implicit"},
		},
	},
}

var OAuthSupportedClientDPStateLifeTimeIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var OAuthSupportedClientAUCodeLifeTimeCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "customized",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"false"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "oauth_role",
			AttrType:    "DmOAuthRole",
			AttrDefault: "",
			Value:       []string{"azsvr"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "az_grant",
			AttrType:    "DmOAuthAZGrantType",
			AttrDefault: "",
			Value:       []string{"code"},
		},
	},
}

var OAuthSupportedClientAUCodeLifeTimeIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var OAuthSupportedClientAccessTokenLifeTimeCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "customized",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"false"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "oauth_role",
			AttrType:    "DmOAuthRole",
			AttrDefault: "",
			Value:       []string{"azsvr"},
		},
	},
}

var OAuthSupportedClientAccessTokenLifeTimeIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var OAuthSupportedClientRefreshTokenAllowedCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "customized",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"false"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "oauth_role",
			AttrType:    "DmOAuthRole",
			AttrDefault: "",
			Value:       []string{"azsvr"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "az_grant",
			AttrType:    "DmOAuthAZGrantType",
			AttrDefault: "",
			Value:       []string{"code", "password", "jwt"},
		},
	},
}

var OAuthSupportedClientRefreshTokenAllowedIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var OAuthSupportedClientRefreshTokenLifeTimeCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "customized",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"false"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "oauth_role",
			AttrType:    "DmOAuthRole",
			AttrDefault: "",
			Value:       []string{"azsvr"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "az_grant",
			AttrType:    "DmOAuthAZGrantType",
			AttrDefault: "",
			Value:       []string{"code", "password", "jwt"},
		},
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "refresh_token_allowed",
			AttrType:    "Int64",
			AttrDefault: "0",
			Value:       []string{"0"},
		},
	},
}

var OAuthSupportedClientRefreshTokenLifeTimeIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var OAuthSupportedClientMaxConsentLifeTimeCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "customized",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"false"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "oauth_role",
			AttrType:    "DmOAuthRole",
			AttrDefault: "",
			Value:       []string{"azsvr"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "az_grant",
			AttrType:    "DmOAuthAZGrantType",
			AttrDefault: "",
			Value:       []string{"code", "password", "jwt"},
		},
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "refresh_token_allowed",
			AttrType:    "Int64",
			AttrDefault: "0",
			Value:       []string{"0"},
		},
	},
}

var OAuthSupportedClientMaxConsentLifeTimeIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var OAuthSupportedClientCustomResourceOwnerIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "customized",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"true"},
		},
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "oauth_role",
			AttrType:    "DmOAuthRole",
			AttrDefault: "",
			Value:       []string{"azsvr"},
		},
	},
}

var OAuthSupportedClientResourceOwnerUrlCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "customized",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"false"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "oauth_role",
			AttrType:    "DmOAuthRole",
			AttrDefault: "",
			Value:       []string{"azsvr"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "custom_resource_owner",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"true"},
		},
	},
}

var OAuthSupportedClientResourceOwnerUrlIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var OAuthSupportedClientAdditionalOAuthProcessUrlCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "customized",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"false"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "caching",
			AttrType:    "String",
			AttrDefault: "replay",
			Value:       []string{"custom"},
		},
	},
}

var OAuthSupportedClientAdditionalOAuthProcessUrlIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "customized",
	AttrType:    "Bool",
	AttrDefault: "false",
	Value:       []string{"true"},
}

var OAuthSupportedClientRSSetHeaderIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "customized",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"true"},
		},
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "oauth_role",
			AttrType:    "DmOAuthRole",
			AttrDefault: "",
			Value:       []string{"rssvr"},
		},
	},
}

var OAuthSupportedClientValidationURLSSLClientTypeCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "customized",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"false"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "oauth_role",
			AttrType:    "DmOAuthRole",
			AttrDefault: "",
			Value:       []string{"rssvr"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "use_validation_url",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"true"},
		},
	},
}

var OAuthSupportedClientValidationURLSSLClientTypeIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var OAuthSupportedClientValidationURLSSLClientCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "customized",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"false"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "oauth_role",
			AttrType:    "DmOAuthRole",
			AttrDefault: "",
			Value:       []string{"rssvr"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "use_validation_url",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"true"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "validation_url_ssl_client_type",
			AttrType:    "String",
			AttrDefault: "client",
			Value:       []string{"client"},
		},
	},
}

var OAuthSupportedClientValidationURLSSLClientIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var OAuthSupportedClientJWTGrantValidatorCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "customized",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"false"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "oauth_role",
			AttrType:    "DmOAuthRole",
			AttrDefault: "",
			Value:       []string{"azsvr"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "az_grant",
			AttrType:    "DmOAuthAZGrantType",
			AttrDefault: "",
			Value:       []string{"jwt"},
		},
	},
}

var OAuthSupportedClientJWTGrantValidatorIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var OAuthSupportedClientClientJWTValidatorCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "customized",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"false"},
		},
		{
			Evaluation: "logical-or",
			Conditions: []validators.Evaluation{

				{
					Evaluation:  "property-value-in-list",
					Attribute:   "client_type",
					AttrType:    "String",
					AttrDefault: "confidential",
					Value:       []string{"confidential"},
				},

				{
					Evaluation: "logical-and",
					Conditions: []validators.Evaluation{

						{
							Evaluation:  "property-value-in-list",
							Attribute:   "oauth_role",
							AttrType:    "DmOAuthRole",
							AttrDefault: "",
							Value:       []string{"azsvr"},
						},

						{
							Evaluation:  "property-value-in-list",
							Attribute:   "az_grant",
							AttrType:    "DmOAuthAZGrantType",
							AttrDefault: "",
							Value:       []string{"client"},
						},
					},
				},
			},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "client_authen_method",
			AttrType:    "String",
			AttrDefault: "secret",
			Value:       []string{"jwt"},
		},
		{
			Evaluation: "logical-or",
			Conditions: []validators.Evaluation{

				{
					Evaluation: "logical-and",
					Conditions: []validators.Evaluation{

						{
							Evaluation:  "property-value-in-list",
							Attribute:   "oauth_role",
							AttrType:    "DmOAuthRole",
							AttrDefault: "",
							Value:       []string{"azsvr"},
						},

						{
							Evaluation:  "property-value-in-list",
							Attribute:   "az_grant",
							AttrType:    "DmOAuthAZGrantType",
							AttrDefault: "",
							Value:       []string{"code", "implicit", "password", "client", "jwt"},
						},
					},
				},

				{
					Evaluation: "logical-and",
					Conditions: []validators.Evaluation{

						{
							Evaluation:  "property-value-in-list",
							Attribute:   "oauth_role",
							AttrType:    "DmOAuthRole",
							AttrDefault: "",
							Value:       []string{"rssvr"},
						},

						{
							Evaluation:  "property-value-not-in-list",
							Attribute:   "oauth_role",
							AttrType:    "DmOAuthRole",
							AttrDefault: "",
							Value:       []string{"azsvr"},
						},

						{
							Evaluation:  "property-value-in-list",
							Attribute:   "check_client_credential",
							AttrType:    "Bool",
							AttrDefault: "false",
							Value:       []string{"true"},
						},
					},
				},
			},
		},
	},
}

var OAuthSupportedClientClientJWTValidatorIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var OAuthSupportedClientOIDCIDTokenGeneratorCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "customized",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"false"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "oauth_role",
			AttrType:    "DmOAuthRole",
			AttrDefault: "",
			Value:       []string{"azsvr"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "az_grant",
			AttrType:    "DmOAuthAZGrantType",
			AttrDefault: "",
			Value:       []string{"oidc"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "az_grant",
			AttrType:    "DmOAuthAZGrantType",
			AttrDefault: "",
			Value:       []string{"code", "implicit", "password", "client", "jwt"},
		},
	},
}

var OAuthSupportedClientOIDCIDTokenGeneratorIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}

var OAuthSupportedClientObjectType = map[string]attr.Type{
	"id":                             types.StringType,
	"app_domain":                     types.StringType,
	"user_summary":                   types.StringType,
	"customized":                     types.BoolType,
	"customized_process_url":         types.StringType,
	"oauth_role":                     types.ObjectType{AttrTypes: DmOAuthRoleObjectType},
	"az_grant":                       types.ObjectType{AttrTypes: DmOAuthAZGrantTypeObjectType},
	"client_type":                    types.StringType,
	"check_client_credential":        types.BoolType,
	"use_validation_url":             types.BoolType,
	"client_authen_method":           types.StringType,
	"client_val_cred":                types.StringType,
	"generate_client_secret":         types.BoolType,
	"client_secret_wo":               types.StringType,
	"client_secret_wo_version":       types.Int64Type,
	"caching":                        types.StringType,
	"validation_url":                 types.StringType,
	"validation_features":            types.ObjectType{AttrTypes: DmValidationFeaturesObjectType},
	"redirect_uri":                   types.ListType{ElemType: types.StringType},
	"custom_scope_check":             types.BoolType,
	"scope":                          types.StringType,
	"scope_url":                      types.StringType,
	"default_scope":                  types.StringType,
	"token_secret":                   types.StringType,
	"local_az_page_url":              types.StringType,
	"dp_state_life_time":             types.Int64Type,
	"au_code_life_time":              types.Int64Type,
	"access_token_life_time":         types.Int64Type,
	"refresh_token_allowed":          types.Int64Type,
	"refresh_token_life_time":        types.Int64Type,
	"max_consent_life_time":          types.Int64Type,
	"custom_resource_owner":          types.BoolType,
	"resource_owner_url":             types.StringType,
	"additional_oauth_process_url":   types.StringType,
	"rs_set_header":                  types.ObjectType{AttrTypes: DmOAuthRSSetHeaderObjectType},
	"validation_url_ssl_client_type": types.StringType,
	"validation_url_ssl_client":      types.StringType,
	"jwt_grant_validator":            types.StringType,
	"client_jwt_validator":           types.StringType,
	"oidc_id_token_generator":        types.StringType,
	"oauth_features":                 types.ObjectType{AttrTypes: DmOAuthFeaturesObjectType},
	"dependency_actions":             actions.ActionsListType,
}
var OAuthSupportedClientObjectTypeWO = map[string]attr.Type{
	"id":                             types.StringType,
	"app_domain":                     types.StringType,
	"user_summary":                   types.StringType,
	"customized":                     types.BoolType,
	"customized_process_url":         types.StringType,
	"oauth_role":                     types.ObjectType{AttrTypes: DmOAuthRoleObjectType},
	"az_grant":                       types.ObjectType{AttrTypes: DmOAuthAZGrantTypeObjectType},
	"client_type":                    types.StringType,
	"check_client_credential":        types.BoolType,
	"use_validation_url":             types.BoolType,
	"client_authen_method":           types.StringType,
	"client_val_cred":                types.StringType,
	"generate_client_secret":         types.BoolType,
	"caching":                        types.StringType,
	"validation_url":                 types.StringType,
	"validation_features":            types.ObjectType{AttrTypes: DmValidationFeaturesObjectType},
	"redirect_uri":                   types.ListType{ElemType: types.StringType},
	"custom_scope_check":             types.BoolType,
	"scope":                          types.StringType,
	"scope_url":                      types.StringType,
	"default_scope":                  types.StringType,
	"token_secret":                   types.StringType,
	"local_az_page_url":              types.StringType,
	"dp_state_life_time":             types.Int64Type,
	"au_code_life_time":              types.Int64Type,
	"access_token_life_time":         types.Int64Type,
	"refresh_token_allowed":          types.Int64Type,
	"refresh_token_life_time":        types.Int64Type,
	"max_consent_life_time":          types.Int64Type,
	"custom_resource_owner":          types.BoolType,
	"resource_owner_url":             types.StringType,
	"additional_oauth_process_url":   types.StringType,
	"rs_set_header":                  types.ObjectType{AttrTypes: DmOAuthRSSetHeaderObjectType},
	"validation_url_ssl_client_type": types.StringType,
	"validation_url_ssl_client":      types.StringType,
	"jwt_grant_validator":            types.StringType,
	"client_jwt_validator":           types.StringType,
	"oidc_id_token_generator":        types.StringType,
	"oauth_features":                 types.ObjectType{AttrTypes: DmOAuthFeaturesObjectType},
	"dependency_actions":             actions.ActionsListType,
}

func (data OAuthSupportedClient) GetPath() string {
	rest_path := "/mgmt/config/{domain}/OAuthSupportedClient"
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data OAuthSupportedClientWO) GetPath() string {
	rest_path := "/mgmt/config/{domain}/OAuthSupportedClient"
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data OAuthSupportedClient) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.Customized.IsNull() {
		return false
	}
	if !data.CustomizedProcessUrl.IsNull() {
		return false
	}
	if data.OauthRole != nil {
		if !data.OauthRole.IsNull() {
			return false
		}
	}
	if data.AzGrant != nil {
		if !data.AzGrant.IsNull() {
			return false
		}
	}
	if !data.ClientType.IsNull() {
		return false
	}
	if !data.CheckClientCredential.IsNull() {
		return false
	}
	if !data.UseValidationUrl.IsNull() {
		return false
	}
	if !data.ClientAuthenMethod.IsNull() {
		return false
	}
	if !data.ClientValCred.IsNull() {
		return false
	}
	if !data.GenerateClientSecret.IsNull() {
		return false
	}
	if !data.ClientSecretWo.IsNull() {
		return false
	}
	if !data.Caching.IsNull() {
		return false
	}
	if !data.ValidationUrl.IsNull() {
		return false
	}
	if data.ValidationFeatures != nil {
		if !data.ValidationFeatures.IsNull() {
			return false
		}
	}
	if !data.RedirectUri.IsNull() {
		return false
	}
	if !data.CustomScopeCheck.IsNull() {
		return false
	}
	if !data.Scope.IsNull() {
		return false
	}
	if !data.ScopeUrl.IsNull() {
		return false
	}
	if !data.DefaultScope.IsNull() {
		return false
	}
	if !data.TokenSecret.IsNull() {
		return false
	}
	if !data.LocalAzPageUrl.IsNull() {
		return false
	}
	if !data.DpStateLifeTime.IsNull() {
		return false
	}
	if !data.AuCodeLifeTime.IsNull() {
		return false
	}
	if !data.AccessTokenLifeTime.IsNull() {
		return false
	}
	if !data.RefreshTokenAllowed.IsNull() {
		return false
	}
	if !data.RefreshTokenLifeTime.IsNull() {
		return false
	}
	if !data.MaxConsentLifeTime.IsNull() {
		return false
	}
	if !data.CustomResourceOwner.IsNull() {
		return false
	}
	if !data.ResourceOwnerUrl.IsNull() {
		return false
	}
	if !data.AdditionalOauthProcessUrl.IsNull() {
		return false
	}
	if data.RsSetHeader != nil {
		if !data.RsSetHeader.IsNull() {
			return false
		}
	}
	if !data.ValidationUrlSslClientType.IsNull() {
		return false
	}
	if !data.ValidationUrlSslClient.IsNull() {
		return false
	}
	if !data.JwtGrantValidator.IsNull() {
		return false
	}
	if !data.ClientJwtValidator.IsNull() {
		return false
	}
	if !data.OidcIdTokenGenerator.IsNull() {
		return false
	}
	if data.OauthFeatures != nil {
		if !data.OauthFeatures.IsNull() {
			return false
		}
	}
	return true
}

func (data OAuthSupportedClientWO) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.Customized.IsNull() {
		return false
	}
	if !data.CustomizedProcessUrl.IsNull() {
		return false
	}
	if data.OauthRole != nil {
		if !data.OauthRole.IsNull() {
			return false
		}
	}
	if data.AzGrant != nil {
		if !data.AzGrant.IsNull() {
			return false
		}
	}
	if !data.ClientType.IsNull() {
		return false
	}
	if !data.CheckClientCredential.IsNull() {
		return false
	}
	if !data.UseValidationUrl.IsNull() {
		return false
	}
	if !data.ClientAuthenMethod.IsNull() {
		return false
	}
	if !data.ClientValCred.IsNull() {
		return false
	}
	if !data.GenerateClientSecret.IsNull() {
		return false
	}
	if !data.Caching.IsNull() {
		return false
	}
	if !data.ValidationUrl.IsNull() {
		return false
	}
	if data.ValidationFeatures != nil {
		if !data.ValidationFeatures.IsNull() {
			return false
		}
	}
	if !data.RedirectUri.IsNull() {
		return false
	}
	if !data.CustomScopeCheck.IsNull() {
		return false
	}
	if !data.Scope.IsNull() {
		return false
	}
	if !data.ScopeUrl.IsNull() {
		return false
	}
	if !data.DefaultScope.IsNull() {
		return false
	}
	if !data.TokenSecret.IsNull() {
		return false
	}
	if !data.LocalAzPageUrl.IsNull() {
		return false
	}
	if !data.DpStateLifeTime.IsNull() {
		return false
	}
	if !data.AuCodeLifeTime.IsNull() {
		return false
	}
	if !data.AccessTokenLifeTime.IsNull() {
		return false
	}
	if !data.RefreshTokenAllowed.IsNull() {
		return false
	}
	if !data.RefreshTokenLifeTime.IsNull() {
		return false
	}
	if !data.MaxConsentLifeTime.IsNull() {
		return false
	}
	if !data.CustomResourceOwner.IsNull() {
		return false
	}
	if !data.ResourceOwnerUrl.IsNull() {
		return false
	}
	if !data.AdditionalOauthProcessUrl.IsNull() {
		return false
	}
	if data.RsSetHeader != nil {
		if !data.RsSetHeader.IsNull() {
			return false
		}
	}
	if !data.ValidationUrlSslClientType.IsNull() {
		return false
	}
	if !data.ValidationUrlSslClient.IsNull() {
		return false
	}
	if !data.JwtGrantValidator.IsNull() {
		return false
	}
	if !data.ClientJwtValidator.IsNull() {
		return false
	}
	if !data.OidcIdTokenGenerator.IsNull() {
		return false
	}
	if data.OauthFeatures != nil {
		if !data.OauthFeatures.IsNull() {
			return false
		}
	}
	return true
}

func (data OAuthSupportedClient) ToBody(ctx context.Context, pathRoot string, config *OAuthSupportedClient) string {
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
	if !data.Customized.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Customized`, tfutils.StringFromBool(data.Customized, ""))
	}
	if !data.CustomizedProcessUrl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`CustomizedProcessUrl`, data.CustomizedProcessUrl.ValueString())
	}
	if data.OauthRole != nil {
		if !data.OauthRole.IsNull() {
			body, _ = sjson.SetRaw(body, pathRoot+`OAuthRole`, data.OauthRole.ToBody(ctx, ""))
		}
	}
	if data.AzGrant != nil {
		if !data.AzGrant.IsNull() {
			body, _ = sjson.SetRaw(body, pathRoot+`AZGrant`, data.AzGrant.ToBody(ctx, ""))
		}
	}
	if !data.ClientType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ClientType`, data.ClientType.ValueString())
	}
	if !data.CheckClientCredential.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`CheckClientCredential`, tfutils.StringFromBool(data.CheckClientCredential, ""))
	}
	if !data.UseValidationUrl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UseValidationUrl`, tfutils.StringFromBool(data.UseValidationUrl, ""))
	}
	if !data.ClientAuthenMethod.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ClientAuthenMethod`, data.ClientAuthenMethod.ValueString())
	}
	if !data.ClientValCred.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ClientValCred`, data.ClientValCred.ValueString())
	}
	if !data.GenerateClientSecret.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`GenerateClientSecret`, tfutils.StringFromBool(data.GenerateClientSecret, ""))
	}
	if !data.ClientSecretWo.IsNull() || !data.ClientSecretWoVersion.IsNull() {
		if data.ClientSecretWo.IsNull() && config != nil {
			data.ClientSecretWo = config.ClientSecretWo
		}
		body, _ = sjson.Set(body, pathRoot+`ClientSecret`, data.ClientSecretWo.ValueString())
	}
	if !data.Caching.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Caching`, data.Caching.ValueString())
	}
	if !data.ValidationUrl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ValidationURL`, data.ValidationUrl.ValueString())
	}
	if data.ValidationFeatures != nil {
		if !data.ValidationFeatures.IsNull() {
			body, _ = sjson.SetRaw(body, pathRoot+`ValidationFeatures`, data.ValidationFeatures.ToBody(ctx, ""))
		}
	}
	if !data.RedirectUri.IsNull() {
		var dataValues []string
		data.RedirectUri.ElementsAs(ctx, &dataValues, false)
		if len(dataValues) > 0 {
			for _, val := range dataValues {
				body, _ = sjson.Set(body, pathRoot+`RedirectURI`+".-1", map[string]string{"value": val})
			}
		} else {
			body, _ = sjson.SetRaw(body, pathRoot+`RedirectURI`, "[]")
		}
	} else {
		body, _ = sjson.SetRaw(body, pathRoot+`RedirectURI`, "[]")
	}
	if !data.CustomScopeCheck.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`CustomScopeCheck`, tfutils.StringFromBool(data.CustomScopeCheck, ""))
	}
	if !data.Scope.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Scope`, data.Scope.ValueString())
	}
	if !data.ScopeUrl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ScopeUrl`, data.ScopeUrl.ValueString())
	}
	if !data.DefaultScope.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DefaultScope`, data.DefaultScope.ValueString())
	}
	if !data.TokenSecret.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`TokenSecret`, data.TokenSecret.ValueString())
	}
	if !data.LocalAzPageUrl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LocalAZPageUrl`, data.LocalAzPageUrl.ValueString())
	}
	if !data.DpStateLifeTime.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DPStateLifeTime`, data.DpStateLifeTime.ValueInt64())
	}
	if !data.AuCodeLifeTime.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AUCodeLifeTime`, data.AuCodeLifeTime.ValueInt64())
	}
	if !data.AccessTokenLifeTime.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AccessTokenLifeTime`, data.AccessTokenLifeTime.ValueInt64())
	}
	if !data.RefreshTokenAllowed.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`RefreshTokenAllowed`, data.RefreshTokenAllowed.ValueInt64())
	}
	if !data.RefreshTokenLifeTime.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`RefreshTokenLifeTime`, data.RefreshTokenLifeTime.ValueInt64())
	}
	if !data.MaxConsentLifeTime.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MaxConsentLifeTime`, data.MaxConsentLifeTime.ValueInt64())
	}
	if !data.CustomResourceOwner.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`CustomResourceOwner`, tfutils.StringFromBool(data.CustomResourceOwner, ""))
	}
	if !data.ResourceOwnerUrl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ResourceOwnerUrl`, data.ResourceOwnerUrl.ValueString())
	}
	if !data.AdditionalOauthProcessUrl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AdditionalOAuthProcessUrl`, data.AdditionalOauthProcessUrl.ValueString())
	}
	if data.RsSetHeader != nil {
		if !data.RsSetHeader.IsNull() {
			body, _ = sjson.SetRaw(body, pathRoot+`RSSetHeader`, data.RsSetHeader.ToBody(ctx, ""))
		}
	}
	if !data.ValidationUrlSslClientType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ValidationURLSSLClientType`, data.ValidationUrlSslClientType.ValueString())
	}
	if !data.ValidationUrlSslClient.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ValidationURLSSLClient`, data.ValidationUrlSslClient.ValueString())
	}
	if !data.JwtGrantValidator.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`JWTGrantValidator`, data.JwtGrantValidator.ValueString())
	}
	if !data.ClientJwtValidator.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ClientJWTValidator`, data.ClientJwtValidator.ValueString())
	}
	if !data.OidcIdTokenGenerator.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`OIDCIDTokenGenerator`, data.OidcIdTokenGenerator.ValueString())
	}
	if data.OauthFeatures != nil {
		if !data.OauthFeatures.IsNull() {
			body, _ = sjson.SetRaw(body, pathRoot+`OAuthFeatures`, data.OauthFeatures.ToBody(ctx, ""))
		}
	}
	return body
}

func (data *OAuthSupportedClient) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `Customized`); value.Exists() {
		data.Customized = tfutils.BoolFromString(value.String())
	} else {
		data.Customized = types.BoolNull()
	}
	if value := res.Get(pathRoot + `CustomizedProcessUrl`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.CustomizedProcessUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.CustomizedProcessUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `OAuthRole`); value.Exists() {
		data.OauthRole = &DmOAuthRole{}
		data.OauthRole.FromBody(ctx, "", value)
	} else {
		data.OauthRole = nil
	}
	if value := res.Get(pathRoot + `AZGrant`); value.Exists() {
		data.AzGrant = &DmOAuthAZGrantType{}
		data.AzGrant.FromBody(ctx, "", value)
	} else {
		data.AzGrant = nil
	}
	if value := res.Get(pathRoot + `ClientType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ClientType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ClientType = types.StringValue("confidential")
	}
	if value := res.Get(pathRoot + `CheckClientCredential`); value.Exists() {
		data.CheckClientCredential = tfutils.BoolFromString(value.String())
	} else {
		data.CheckClientCredential = types.BoolNull()
	}
	if value := res.Get(pathRoot + `UseValidationUrl`); value.Exists() {
		data.UseValidationUrl = tfutils.BoolFromString(value.String())
	} else {
		data.UseValidationUrl = types.BoolNull()
	}
	if value := res.Get(pathRoot + `ClientAuthenMethod`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ClientAuthenMethod = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ClientAuthenMethod = types.StringValue("secret")
	}
	if value := res.Get(pathRoot + `ClientValCred`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ClientValCred = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ClientValCred = types.StringNull()
	}
	if value := res.Get(pathRoot + `GenerateClientSecret`); value.Exists() {
		data.GenerateClientSecret = tfutils.BoolFromString(value.String())
	} else {
		data.GenerateClientSecret = types.BoolNull()
	}
	if value := res.Get(pathRoot + `ClientSecret`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ClientSecretWo = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ClientSecretWo = types.StringNull()
	}
	if value := res.Get(pathRoot + `Caching`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Caching = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Caching = types.StringValue("replay")
	}
	if value := res.Get(pathRoot + `ValidationURL`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ValidationUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ValidationUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `ValidationFeatures`); value.Exists() {
		data.ValidationFeatures = &DmValidationFeatures{}
		data.ValidationFeatures.FromBody(ctx, "", value)
	} else {
		data.ValidationFeatures = nil
	}
	if value := res.Get(pathRoot + `RedirectURI`); value.Exists() {
		data.RedirectUri = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.RedirectUri = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `CustomScopeCheck`); value.Exists() {
		data.CustomScopeCheck = tfutils.BoolFromString(value.String())
	} else {
		data.CustomScopeCheck = types.BoolNull()
	}
	if value := res.Get(pathRoot + `Scope`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Scope = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Scope = types.StringNull()
	}
	if value := res.Get(pathRoot + `ScopeUrl`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ScopeUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ScopeUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `DefaultScope`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.DefaultScope = tfutils.ParseStringFromGJSON(value)
	} else {
		data.DefaultScope = types.StringNull()
	}
	if value := res.Get(pathRoot + `TokenSecret`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.TokenSecret = tfutils.ParseStringFromGJSON(value)
	} else {
		data.TokenSecret = types.StringNull()
	}
	if value := res.Get(pathRoot + `LocalAZPageUrl`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.LocalAzPageUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LocalAzPageUrl = types.StringValue("store:///OAuth-Generate-HTML.xsl")
	}
	if value := res.Get(pathRoot + `DPStateLifeTime`); value.Exists() {
		data.DpStateLifeTime = types.Int64Value(value.Int())
	} else {
		data.DpStateLifeTime = types.Int64Value(300)
	}
	if value := res.Get(pathRoot + `AUCodeLifeTime`); value.Exists() {
		data.AuCodeLifeTime = types.Int64Value(value.Int())
	} else {
		data.AuCodeLifeTime = types.Int64Value(300)
	}
	if value := res.Get(pathRoot + `AccessTokenLifeTime`); value.Exists() {
		data.AccessTokenLifeTime = types.Int64Value(value.Int())
	} else {
		data.AccessTokenLifeTime = types.Int64Value(3600)
	}
	if value := res.Get(pathRoot + `RefreshTokenAllowed`); value.Exists() {
		data.RefreshTokenAllowed = types.Int64Value(value.Int())
	} else {
		data.RefreshTokenAllowed = types.Int64Value(0)
	}
	if value := res.Get(pathRoot + `RefreshTokenLifeTime`); value.Exists() {
		data.RefreshTokenLifeTime = types.Int64Value(value.Int())
	} else {
		data.RefreshTokenLifeTime = types.Int64Value(5400)
	}
	if value := res.Get(pathRoot + `MaxConsentLifeTime`); value.Exists() {
		data.MaxConsentLifeTime = types.Int64Value(value.Int())
	} else {
		data.MaxConsentLifeTime = types.Int64Null()
	}
	if value := res.Get(pathRoot + `CustomResourceOwner`); value.Exists() {
		data.CustomResourceOwner = tfutils.BoolFromString(value.String())
	} else {
		data.CustomResourceOwner = types.BoolNull()
	}
	if value := res.Get(pathRoot + `ResourceOwnerUrl`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ResourceOwnerUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ResourceOwnerUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `AdditionalOAuthProcessUrl`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AdditionalOauthProcessUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AdditionalOauthProcessUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `RSSetHeader`); value.Exists() {
		data.RsSetHeader = &DmOAuthRSSetHeader{}
		data.RsSetHeader.FromBody(ctx, "", value)
	} else {
		data.RsSetHeader = nil
	}
	if value := res.Get(pathRoot + `ValidationURLSSLClientType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ValidationUrlSslClientType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ValidationUrlSslClientType = types.StringValue("client")
	}
	if value := res.Get(pathRoot + `ValidationURLSSLClient`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ValidationUrlSslClient = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ValidationUrlSslClient = types.StringNull()
	}
	if value := res.Get(pathRoot + `JWTGrantValidator`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.JwtGrantValidator = tfutils.ParseStringFromGJSON(value)
	} else {
		data.JwtGrantValidator = types.StringNull()
	}
	if value := res.Get(pathRoot + `ClientJWTValidator`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ClientJwtValidator = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ClientJwtValidator = types.StringNull()
	}
	if value := res.Get(pathRoot + `OIDCIDTokenGenerator`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.OidcIdTokenGenerator = tfutils.ParseStringFromGJSON(value)
	} else {
		data.OidcIdTokenGenerator = types.StringNull()
	}
	if value := res.Get(pathRoot + `OAuthFeatures`); value.Exists() {
		data.OauthFeatures = &DmOAuthFeatures{}
		data.OauthFeatures.FromBody(ctx, "", value)
	} else {
		data.OauthFeatures = nil
	}
}

func (data *OAuthSupportedClientWO) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `Customized`); value.Exists() {
		data.Customized = tfutils.BoolFromString(value.String())
	} else {
		data.Customized = types.BoolNull()
	}
	if value := res.Get(pathRoot + `CustomizedProcessUrl`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.CustomizedProcessUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.CustomizedProcessUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `OAuthRole`); value.Exists() {
		data.OauthRole = &DmOAuthRole{}
		data.OauthRole.FromBody(ctx, "", value)
	} else {
		data.OauthRole = nil
	}
	if value := res.Get(pathRoot + `AZGrant`); value.Exists() {
		data.AzGrant = &DmOAuthAZGrantType{}
		data.AzGrant.FromBody(ctx, "", value)
	} else {
		data.AzGrant = nil
	}
	if value := res.Get(pathRoot + `ClientType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ClientType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ClientType = types.StringValue("confidential")
	}
	if value := res.Get(pathRoot + `CheckClientCredential`); value.Exists() {
		data.CheckClientCredential = tfutils.BoolFromString(value.String())
	} else {
		data.CheckClientCredential = types.BoolNull()
	}
	if value := res.Get(pathRoot + `UseValidationUrl`); value.Exists() {
		data.UseValidationUrl = tfutils.BoolFromString(value.String())
	} else {
		data.UseValidationUrl = types.BoolNull()
	}
	if value := res.Get(pathRoot + `ClientAuthenMethod`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ClientAuthenMethod = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ClientAuthenMethod = types.StringValue("secret")
	}
	if value := res.Get(pathRoot + `ClientValCred`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ClientValCred = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ClientValCred = types.StringNull()
	}
	if value := res.Get(pathRoot + `GenerateClientSecret`); value.Exists() {
		data.GenerateClientSecret = tfutils.BoolFromString(value.String())
	} else {
		data.GenerateClientSecret = types.BoolNull()
	}
	if value := res.Get(pathRoot + `Caching`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Caching = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Caching = types.StringValue("replay")
	}
	if value := res.Get(pathRoot + `ValidationURL`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ValidationUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ValidationUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `ValidationFeatures`); value.Exists() {
		data.ValidationFeatures = &DmValidationFeatures{}
		data.ValidationFeatures.FromBody(ctx, "", value)
	} else {
		data.ValidationFeatures = nil
	}
	if value := res.Get(pathRoot + `RedirectURI`); value.Exists() {
		data.RedirectUri = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.RedirectUri = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `CustomScopeCheck`); value.Exists() {
		data.CustomScopeCheck = tfutils.BoolFromString(value.String())
	} else {
		data.CustomScopeCheck = types.BoolNull()
	}
	if value := res.Get(pathRoot + `Scope`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Scope = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Scope = types.StringNull()
	}
	if value := res.Get(pathRoot + `ScopeUrl`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ScopeUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ScopeUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `DefaultScope`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.DefaultScope = tfutils.ParseStringFromGJSON(value)
	} else {
		data.DefaultScope = types.StringNull()
	}
	if value := res.Get(pathRoot + `TokenSecret`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.TokenSecret = tfutils.ParseStringFromGJSON(value)
	} else {
		data.TokenSecret = types.StringNull()
	}
	if value := res.Get(pathRoot + `LocalAZPageUrl`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.LocalAzPageUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LocalAzPageUrl = types.StringValue("store:///OAuth-Generate-HTML.xsl")
	}
	if value := res.Get(pathRoot + `DPStateLifeTime`); value.Exists() {
		data.DpStateLifeTime = types.Int64Value(value.Int())
	} else {
		data.DpStateLifeTime = types.Int64Value(300)
	}
	if value := res.Get(pathRoot + `AUCodeLifeTime`); value.Exists() {
		data.AuCodeLifeTime = types.Int64Value(value.Int())
	} else {
		data.AuCodeLifeTime = types.Int64Value(300)
	}
	if value := res.Get(pathRoot + `AccessTokenLifeTime`); value.Exists() {
		data.AccessTokenLifeTime = types.Int64Value(value.Int())
	} else {
		data.AccessTokenLifeTime = types.Int64Value(3600)
	}
	if value := res.Get(pathRoot + `RefreshTokenAllowed`); value.Exists() {
		data.RefreshTokenAllowed = types.Int64Value(value.Int())
	} else {
		data.RefreshTokenAllowed = types.Int64Value(0)
	}
	if value := res.Get(pathRoot + `RefreshTokenLifeTime`); value.Exists() {
		data.RefreshTokenLifeTime = types.Int64Value(value.Int())
	} else {
		data.RefreshTokenLifeTime = types.Int64Value(5400)
	}
	if value := res.Get(pathRoot + `MaxConsentLifeTime`); value.Exists() {
		data.MaxConsentLifeTime = types.Int64Value(value.Int())
	} else {
		data.MaxConsentLifeTime = types.Int64Null()
	}
	if value := res.Get(pathRoot + `CustomResourceOwner`); value.Exists() {
		data.CustomResourceOwner = tfutils.BoolFromString(value.String())
	} else {
		data.CustomResourceOwner = types.BoolNull()
	}
	if value := res.Get(pathRoot + `ResourceOwnerUrl`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ResourceOwnerUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ResourceOwnerUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `AdditionalOAuthProcessUrl`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AdditionalOauthProcessUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AdditionalOauthProcessUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `RSSetHeader`); value.Exists() {
		data.RsSetHeader = &DmOAuthRSSetHeader{}
		data.RsSetHeader.FromBody(ctx, "", value)
	} else {
		data.RsSetHeader = nil
	}
	if value := res.Get(pathRoot + `ValidationURLSSLClientType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ValidationUrlSslClientType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ValidationUrlSslClientType = types.StringValue("client")
	}
	if value := res.Get(pathRoot + `ValidationURLSSLClient`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ValidationUrlSslClient = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ValidationUrlSslClient = types.StringNull()
	}
	if value := res.Get(pathRoot + `JWTGrantValidator`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.JwtGrantValidator = tfutils.ParseStringFromGJSON(value)
	} else {
		data.JwtGrantValidator = types.StringNull()
	}
	if value := res.Get(pathRoot + `ClientJWTValidator`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ClientJwtValidator = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ClientJwtValidator = types.StringNull()
	}
	if value := res.Get(pathRoot + `OIDCIDTokenGenerator`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.OidcIdTokenGenerator = tfutils.ParseStringFromGJSON(value)
	} else {
		data.OidcIdTokenGenerator = types.StringNull()
	}
	if value := res.Get(pathRoot + `OAuthFeatures`); value.Exists() {
		data.OauthFeatures = &DmOAuthFeatures{}
		data.OauthFeatures.FromBody(ctx, "", value)
	} else {
		data.OauthFeatures = nil
	}
}

func (data *OAuthSupportedClient) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `Customized`); value.Exists() && !data.Customized.IsNull() {
		data.Customized = tfutils.BoolFromString(value.String())
	} else if data.Customized.ValueBool() {
		data.Customized = types.BoolNull()
	}
	if value := res.Get(pathRoot + `CustomizedProcessUrl`); value.Exists() && !data.CustomizedProcessUrl.IsNull() {
		data.CustomizedProcessUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.CustomizedProcessUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `OAuthRole`); value.Exists() {
		data.OauthRole.UpdateFromBody(ctx, "", value)
	} else {
		data.OauthRole = nil
	}
	if value := res.Get(pathRoot + `AZGrant`); value.Exists() {
		data.AzGrant.UpdateFromBody(ctx, "", value)
	} else {
		data.AzGrant = nil
	}
	if value := res.Get(pathRoot + `ClientType`); value.Exists() && !data.ClientType.IsNull() {
		data.ClientType = tfutils.ParseStringFromGJSON(value)
	} else if data.ClientType.ValueString() != "confidential" {
		data.ClientType = types.StringNull()
	}
	if value := res.Get(pathRoot + `CheckClientCredential`); value.Exists() && !data.CheckClientCredential.IsNull() {
		data.CheckClientCredential = tfutils.BoolFromString(value.String())
	} else if data.CheckClientCredential.ValueBool() {
		data.CheckClientCredential = types.BoolNull()
	}
	if value := res.Get(pathRoot + `UseValidationUrl`); value.Exists() && !data.UseValidationUrl.IsNull() {
		data.UseValidationUrl = tfutils.BoolFromString(value.String())
	} else if data.UseValidationUrl.ValueBool() {
		data.UseValidationUrl = types.BoolNull()
	}
	if value := res.Get(pathRoot + `ClientAuthenMethod`); value.Exists() && !data.ClientAuthenMethod.IsNull() {
		data.ClientAuthenMethod = tfutils.ParseStringFromGJSON(value)
	} else if data.ClientAuthenMethod.ValueString() != "secret" {
		data.ClientAuthenMethod = types.StringNull()
	}
	if value := res.Get(pathRoot + `ClientValCred`); value.Exists() && !data.ClientValCred.IsNull() {
		data.ClientValCred = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ClientValCred = types.StringNull()
	}
	if value := res.Get(pathRoot + `GenerateClientSecret`); value.Exists() && !data.GenerateClientSecret.IsNull() {
		data.GenerateClientSecret = tfutils.BoolFromString(value.String())
	} else if !data.GenerateClientSecret.ValueBool() {
		data.GenerateClientSecret = types.BoolNull()
	}
	if value := res.Get(pathRoot + `ClientSecret`); value.Exists() && !data.ClientSecretWo.IsNull() {
		data.ClientSecretWo = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ClientSecretWo = types.StringNull()
	}
	if value := res.Get(pathRoot + `Caching`); value.Exists() && !data.Caching.IsNull() {
		data.Caching = tfutils.ParseStringFromGJSON(value)
	} else if data.Caching.ValueString() != "replay" {
		data.Caching = types.StringNull()
	}
	if value := res.Get(pathRoot + `ValidationURL`); value.Exists() && !data.ValidationUrl.IsNull() {
		data.ValidationUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ValidationUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `ValidationFeatures`); value.Exists() {
		data.ValidationFeatures.UpdateFromBody(ctx, "", value)
	} else {
		data.ValidationFeatures = nil
	}
	if value := res.Get(pathRoot + `RedirectURI`); value.Exists() && !data.RedirectUri.IsNull() {
		data.RedirectUri = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.RedirectUri = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `CustomScopeCheck`); value.Exists() && !data.CustomScopeCheck.IsNull() {
		data.CustomScopeCheck = tfutils.BoolFromString(value.String())
	} else if data.CustomScopeCheck.ValueBool() {
		data.CustomScopeCheck = types.BoolNull()
	}
	if value := res.Get(pathRoot + `Scope`); value.Exists() && !data.Scope.IsNull() {
		data.Scope = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Scope = types.StringNull()
	}
	if value := res.Get(pathRoot + `ScopeUrl`); value.Exists() && !data.ScopeUrl.IsNull() {
		data.ScopeUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ScopeUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `DefaultScope`); value.Exists() && !data.DefaultScope.IsNull() {
		data.DefaultScope = tfutils.ParseStringFromGJSON(value)
	} else {
		data.DefaultScope = types.StringNull()
	}
	if value := res.Get(pathRoot + `TokenSecret`); value.Exists() && !data.TokenSecret.IsNull() {
		data.TokenSecret = tfutils.ParseStringFromGJSON(value)
	} else {
		data.TokenSecret = types.StringNull()
	}
	if value := res.Get(pathRoot + `LocalAZPageUrl`); value.Exists() && !data.LocalAzPageUrl.IsNull() {
		data.LocalAzPageUrl = tfutils.ParseStringFromGJSON(value)
	} else if data.LocalAzPageUrl.ValueString() != "store:///OAuth-Generate-HTML.xsl" {
		data.LocalAzPageUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `DPStateLifeTime`); value.Exists() && !data.DpStateLifeTime.IsNull() {
		data.DpStateLifeTime = types.Int64Value(value.Int())
	} else if data.DpStateLifeTime.ValueInt64() != 300 {
		data.DpStateLifeTime = types.Int64Null()
	}
	if value := res.Get(pathRoot + `AUCodeLifeTime`); value.Exists() && !data.AuCodeLifeTime.IsNull() {
		data.AuCodeLifeTime = types.Int64Value(value.Int())
	} else if data.AuCodeLifeTime.ValueInt64() != 300 {
		data.AuCodeLifeTime = types.Int64Null()
	}
	if value := res.Get(pathRoot + `AccessTokenLifeTime`); value.Exists() && !data.AccessTokenLifeTime.IsNull() {
		data.AccessTokenLifeTime = types.Int64Value(value.Int())
	} else if data.AccessTokenLifeTime.ValueInt64() != 3600 {
		data.AccessTokenLifeTime = types.Int64Null()
	}
	if value := res.Get(pathRoot + `RefreshTokenAllowed`); value.Exists() && !data.RefreshTokenAllowed.IsNull() {
		data.RefreshTokenAllowed = types.Int64Value(value.Int())
	} else if data.RefreshTokenAllowed.ValueInt64() != 0 {
		data.RefreshTokenAllowed = types.Int64Null()
	}
	if value := res.Get(pathRoot + `RefreshTokenLifeTime`); value.Exists() && !data.RefreshTokenLifeTime.IsNull() {
		data.RefreshTokenLifeTime = types.Int64Value(value.Int())
	} else if data.RefreshTokenLifeTime.ValueInt64() != 5400 {
		data.RefreshTokenLifeTime = types.Int64Null()
	}
	if value := res.Get(pathRoot + `MaxConsentLifeTime`); value.Exists() && !data.MaxConsentLifeTime.IsNull() {
		data.MaxConsentLifeTime = types.Int64Value(value.Int())
	} else {
		data.MaxConsentLifeTime = types.Int64Null()
	}
	if value := res.Get(pathRoot + `CustomResourceOwner`); value.Exists() && !data.CustomResourceOwner.IsNull() {
		data.CustomResourceOwner = tfutils.BoolFromString(value.String())
	} else if data.CustomResourceOwner.ValueBool() {
		data.CustomResourceOwner = types.BoolNull()
	}
	if value := res.Get(pathRoot + `ResourceOwnerUrl`); value.Exists() && !data.ResourceOwnerUrl.IsNull() {
		data.ResourceOwnerUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ResourceOwnerUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `AdditionalOAuthProcessUrl`); value.Exists() && !data.AdditionalOauthProcessUrl.IsNull() {
		data.AdditionalOauthProcessUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AdditionalOauthProcessUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `RSSetHeader`); value.Exists() {
		data.RsSetHeader.UpdateFromBody(ctx, "", value)
	} else {
		data.RsSetHeader = nil
	}
	if value := res.Get(pathRoot + `ValidationURLSSLClientType`); value.Exists() && !data.ValidationUrlSslClientType.IsNull() {
		data.ValidationUrlSslClientType = tfutils.ParseStringFromGJSON(value)
	} else if data.ValidationUrlSslClientType.ValueString() != "client" {
		data.ValidationUrlSslClientType = types.StringNull()
	}
	if value := res.Get(pathRoot + `ValidationURLSSLClient`); value.Exists() && !data.ValidationUrlSslClient.IsNull() {
		data.ValidationUrlSslClient = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ValidationUrlSslClient = types.StringNull()
	}
	if value := res.Get(pathRoot + `JWTGrantValidator`); value.Exists() && !data.JwtGrantValidator.IsNull() {
		data.JwtGrantValidator = tfutils.ParseStringFromGJSON(value)
	} else {
		data.JwtGrantValidator = types.StringNull()
	}
	if value := res.Get(pathRoot + `ClientJWTValidator`); value.Exists() && !data.ClientJwtValidator.IsNull() {
		data.ClientJwtValidator = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ClientJwtValidator = types.StringNull()
	}
	if value := res.Get(pathRoot + `OIDCIDTokenGenerator`); value.Exists() && !data.OidcIdTokenGenerator.IsNull() {
		data.OidcIdTokenGenerator = tfutils.ParseStringFromGJSON(value)
	} else {
		data.OidcIdTokenGenerator = types.StringNull()
	}
	if value := res.Get(pathRoot + `OAuthFeatures`); value.Exists() {
		data.OauthFeatures.UpdateFromBody(ctx, "", value)
	} else {
		data.OauthFeatures = nil
	}
}
func (data *OAuthSupportedClient) UpdateUnknownFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if data.Id.IsUnknown() {
		if value := res.Get(pathRoot + `name`); value.Exists() && !data.Id.IsNull() {
			data.Id = tfutils.ParseStringFromGJSON(value)
		} else {
			data.Id = types.StringNull()
		}
	}
	if data.UserSummary.IsUnknown() {
		if value := res.Get(pathRoot + `UserSummary`); value.Exists() && !data.UserSummary.IsNull() {
			data.UserSummary = tfutils.ParseStringFromGJSON(value)
		} else {
			data.UserSummary = types.StringNull()
		}
	}
	if data.Customized.IsUnknown() {
		if value := res.Get(pathRoot + `Customized`); value.Exists() && !data.Customized.IsNull() {
			data.Customized = tfutils.BoolFromString(value.String())
		} else {
			data.Customized = types.BoolNull()
		}
	}
	if data.CustomizedProcessUrl.IsUnknown() {
		if value := res.Get(pathRoot + `CustomizedProcessUrl`); value.Exists() && !data.CustomizedProcessUrl.IsNull() {
			data.CustomizedProcessUrl = tfutils.ParseStringFromGJSON(value)
		} else {
			data.CustomizedProcessUrl = types.StringNull()
		}
	}
	if data.OauthRole == nil {
		if value := res.Get(pathRoot + `OAuthRole`); value.Exists() {
			d := DmOAuthRole{}
			d.UpdateFromBody(ctx, "", value)
			if !d.IsNull() {
				data.OauthRole = &d
			}
		}
	}
	if data.AzGrant == nil {
		if value := res.Get(pathRoot + `AZGrant`); value.Exists() {
			d := DmOAuthAZGrantType{}
			d.UpdateFromBody(ctx, "", value)
			if !d.IsNull() {
				data.AzGrant = &d
			}
		}
	}
	if data.ClientType.IsUnknown() {
		if value := res.Get(pathRoot + `ClientType`); value.Exists() && !data.ClientType.IsNull() {
			data.ClientType = tfutils.ParseStringFromGJSON(value)
		} else if data.ClientType.ValueString() != "confidential" {
			data.ClientType = types.StringNull()
		}
	}
	if data.CheckClientCredential.IsUnknown() {
		if value := res.Get(pathRoot + `CheckClientCredential`); value.Exists() && !data.CheckClientCredential.IsNull() {
			data.CheckClientCredential = tfutils.BoolFromString(value.String())
		} else {
			data.CheckClientCredential = types.BoolNull()
		}
	}
	if data.UseValidationUrl.IsUnknown() {
		if value := res.Get(pathRoot + `UseValidationUrl`); value.Exists() && !data.UseValidationUrl.IsNull() {
			data.UseValidationUrl = tfutils.BoolFromString(value.String())
		} else {
			data.UseValidationUrl = types.BoolNull()
		}
	}
	if data.ClientAuthenMethod.IsUnknown() {
		if value := res.Get(pathRoot + `ClientAuthenMethod`); value.Exists() && !data.ClientAuthenMethod.IsNull() {
			data.ClientAuthenMethod = tfutils.ParseStringFromGJSON(value)
		} else if data.ClientAuthenMethod.ValueString() != "secret" {
			data.ClientAuthenMethod = types.StringNull()
		}
	}
	if data.ClientValCred.IsUnknown() {
		if value := res.Get(pathRoot + `ClientValCred`); value.Exists() && !data.ClientValCred.IsNull() {
			data.ClientValCred = tfutils.ParseStringFromGJSON(value)
		} else {
			data.ClientValCred = types.StringNull()
		}
	}
	if data.GenerateClientSecret.IsUnknown() {
		if value := res.Get(pathRoot + `GenerateClientSecret`); value.Exists() && !data.GenerateClientSecret.IsNull() {
			data.GenerateClientSecret = tfutils.BoolFromString(value.String())
		} else {
			data.GenerateClientSecret = types.BoolNull()
		}
	}
	if data.ClientSecretWo.IsUnknown() {
		if value := res.Get(pathRoot + `ClientSecret`); value.Exists() && !data.ClientSecretWo.IsNull() {
			data.ClientSecretWo = tfutils.ParseStringFromGJSON(value)
		} else {
			data.ClientSecretWo = types.StringNull()
		}
	}
	if data.Caching.IsUnknown() {
		if value := res.Get(pathRoot + `Caching`); value.Exists() && !data.Caching.IsNull() {
			data.Caching = tfutils.ParseStringFromGJSON(value)
		} else if data.Caching.ValueString() != "replay" {
			data.Caching = types.StringNull()
		}
	}
	if data.ValidationUrl.IsUnknown() {
		if value := res.Get(pathRoot + `ValidationURL`); value.Exists() && !data.ValidationUrl.IsNull() {
			data.ValidationUrl = tfutils.ParseStringFromGJSON(value)
		} else {
			data.ValidationUrl = types.StringNull()
		}
	}
	if data.ValidationFeatures == nil {
		if value := res.Get(pathRoot + `ValidationFeatures`); value.Exists() {
			d := DmValidationFeatures{}
			d.UpdateFromBody(ctx, "", value)
			if !d.IsNull() {
				data.ValidationFeatures = &d
			}
		}
	}
	if data.RedirectUri.IsUnknown() {
		if value := res.Get(pathRoot + `RedirectURI`); value.Exists() && !data.RedirectUri.IsNull() {
			data.RedirectUri = tfutils.ParseStringListFromGJSON(value)
		} else {
			data.RedirectUri = types.ListNull(types.StringType)
		}
	}
	if data.CustomScopeCheck.IsUnknown() {
		if value := res.Get(pathRoot + `CustomScopeCheck`); value.Exists() && !data.CustomScopeCheck.IsNull() {
			data.CustomScopeCheck = tfutils.BoolFromString(value.String())
		} else {
			data.CustomScopeCheck = types.BoolNull()
		}
	}
	if data.Scope.IsUnknown() {
		if value := res.Get(pathRoot + `Scope`); value.Exists() && !data.Scope.IsNull() {
			data.Scope = tfutils.ParseStringFromGJSON(value)
		} else {
			data.Scope = types.StringNull()
		}
	}
	if data.ScopeUrl.IsUnknown() {
		if value := res.Get(pathRoot + `ScopeUrl`); value.Exists() && !data.ScopeUrl.IsNull() {
			data.ScopeUrl = tfutils.ParseStringFromGJSON(value)
		} else {
			data.ScopeUrl = types.StringNull()
		}
	}
	if data.DefaultScope.IsUnknown() {
		if value := res.Get(pathRoot + `DefaultScope`); value.Exists() && !data.DefaultScope.IsNull() {
			data.DefaultScope = tfutils.ParseStringFromGJSON(value)
		} else {
			data.DefaultScope = types.StringNull()
		}
	}
	if data.TokenSecret.IsUnknown() {
		if value := res.Get(pathRoot + `TokenSecret`); value.Exists() && !data.TokenSecret.IsNull() {
			data.TokenSecret = tfutils.ParseStringFromGJSON(value)
		} else {
			data.TokenSecret = types.StringNull()
		}
	}
	if data.LocalAzPageUrl.IsUnknown() {
		if value := res.Get(pathRoot + `LocalAZPageUrl`); value.Exists() && !data.LocalAzPageUrl.IsNull() {
			data.LocalAzPageUrl = tfutils.ParseStringFromGJSON(value)
		} else if data.LocalAzPageUrl.ValueString() != "store:///OAuth-Generate-HTML.xsl" {
			data.LocalAzPageUrl = types.StringNull()
		}
	}
	if data.DpStateLifeTime.IsUnknown() {
		if value := res.Get(pathRoot + `DPStateLifeTime`); value.Exists() && !data.DpStateLifeTime.IsNull() {
			data.DpStateLifeTime = types.Int64Value(value.Int())
		} else if data.DpStateLifeTime.ValueInt64() != 300 {
			data.DpStateLifeTime = types.Int64Null()
		}
	}
	if data.AuCodeLifeTime.IsUnknown() {
		if value := res.Get(pathRoot + `AUCodeLifeTime`); value.Exists() && !data.AuCodeLifeTime.IsNull() {
			data.AuCodeLifeTime = types.Int64Value(value.Int())
		} else if data.AuCodeLifeTime.ValueInt64() != 300 {
			data.AuCodeLifeTime = types.Int64Null()
		}
	}
	if data.AccessTokenLifeTime.IsUnknown() {
		if value := res.Get(pathRoot + `AccessTokenLifeTime`); value.Exists() && !data.AccessTokenLifeTime.IsNull() {
			data.AccessTokenLifeTime = types.Int64Value(value.Int())
		} else if data.AccessTokenLifeTime.ValueInt64() != 3600 {
			data.AccessTokenLifeTime = types.Int64Null()
		}
	}
	if data.RefreshTokenAllowed.IsUnknown() {
		if value := res.Get(pathRoot + `RefreshTokenAllowed`); value.Exists() && !data.RefreshTokenAllowed.IsNull() {
			data.RefreshTokenAllowed = types.Int64Value(value.Int())
		} else if data.RefreshTokenAllowed.ValueInt64() != 0 {
			data.RefreshTokenAllowed = types.Int64Null()
		}
	}
	if data.RefreshTokenLifeTime.IsUnknown() {
		if value := res.Get(pathRoot + `RefreshTokenLifeTime`); value.Exists() && !data.RefreshTokenLifeTime.IsNull() {
			data.RefreshTokenLifeTime = types.Int64Value(value.Int())
		} else if data.RefreshTokenLifeTime.ValueInt64() != 5400 {
			data.RefreshTokenLifeTime = types.Int64Null()
		}
	}
	if data.MaxConsentLifeTime.IsUnknown() {
		if value := res.Get(pathRoot + `MaxConsentLifeTime`); value.Exists() && !data.MaxConsentLifeTime.IsNull() {
			data.MaxConsentLifeTime = types.Int64Value(value.Int())
		} else {
			data.MaxConsentLifeTime = types.Int64Null()
		}
	}
	if data.CustomResourceOwner.IsUnknown() {
		if value := res.Get(pathRoot + `CustomResourceOwner`); value.Exists() && !data.CustomResourceOwner.IsNull() {
			data.CustomResourceOwner = tfutils.BoolFromString(value.String())
		} else {
			data.CustomResourceOwner = types.BoolNull()
		}
	}
	if data.ResourceOwnerUrl.IsUnknown() {
		if value := res.Get(pathRoot + `ResourceOwnerUrl`); value.Exists() && !data.ResourceOwnerUrl.IsNull() {
			data.ResourceOwnerUrl = tfutils.ParseStringFromGJSON(value)
		} else {
			data.ResourceOwnerUrl = types.StringNull()
		}
	}
	if data.AdditionalOauthProcessUrl.IsUnknown() {
		if value := res.Get(pathRoot + `AdditionalOAuthProcessUrl`); value.Exists() && !data.AdditionalOauthProcessUrl.IsNull() {
			data.AdditionalOauthProcessUrl = tfutils.ParseStringFromGJSON(value)
		} else {
			data.AdditionalOauthProcessUrl = types.StringNull()
		}
	}
	if data.RsSetHeader == nil {
		if value := res.Get(pathRoot + `RSSetHeader`); value.Exists() {
			d := DmOAuthRSSetHeader{}
			d.UpdateFromBody(ctx, "", value)
			if !d.IsNull() {
				data.RsSetHeader = &d
			}
		}
	}
	if data.ValidationUrlSslClientType.IsUnknown() {
		if value := res.Get(pathRoot + `ValidationURLSSLClientType`); value.Exists() && !data.ValidationUrlSslClientType.IsNull() {
			data.ValidationUrlSslClientType = tfutils.ParseStringFromGJSON(value)
		} else if data.ValidationUrlSslClientType.ValueString() != "client" {
			data.ValidationUrlSslClientType = types.StringNull()
		}
	}
	if data.ValidationUrlSslClient.IsUnknown() {
		if value := res.Get(pathRoot + `ValidationURLSSLClient`); value.Exists() && !data.ValidationUrlSslClient.IsNull() {
			data.ValidationUrlSslClient = tfutils.ParseStringFromGJSON(value)
		} else {
			data.ValidationUrlSslClient = types.StringNull()
		}
	}
	if data.JwtGrantValidator.IsUnknown() {
		if value := res.Get(pathRoot + `JWTGrantValidator`); value.Exists() && !data.JwtGrantValidator.IsNull() {
			data.JwtGrantValidator = tfutils.ParseStringFromGJSON(value)
		} else {
			data.JwtGrantValidator = types.StringNull()
		}
	}
	if data.ClientJwtValidator.IsUnknown() {
		if value := res.Get(pathRoot + `ClientJWTValidator`); value.Exists() && !data.ClientJwtValidator.IsNull() {
			data.ClientJwtValidator = tfutils.ParseStringFromGJSON(value)
		} else {
			data.ClientJwtValidator = types.StringNull()
		}
	}
	if data.OidcIdTokenGenerator.IsUnknown() {
		if value := res.Get(pathRoot + `OIDCIDTokenGenerator`); value.Exists() && !data.OidcIdTokenGenerator.IsNull() {
			data.OidcIdTokenGenerator = tfutils.ParseStringFromGJSON(value)
		} else {
			data.OidcIdTokenGenerator = types.StringNull()
		}
	}
	if data.OauthFeatures == nil {
		if value := res.Get(pathRoot + `OAuthFeatures`); value.Exists() {
			d := DmOAuthFeatures{}
			d.UpdateFromBody(ctx, "", value)
			if !d.IsNull() {
				data.OauthFeatures = &d
			}
		}
	}
}

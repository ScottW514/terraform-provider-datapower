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

package resources

import (
	"context"
	"fmt"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/modifiers"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/validators"
)

var _ resource.Resource = &SSLClientProfileResource{}

func NewSSLClientProfileResource() resource.Resource {
	return &SSLClientProfileResource{}
}

type SSLClientProfileResource struct {
	pData *tfutils.ProviderData
}

func (r *SSLClientProfileResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ssl_client_profile"
}

func (r *SSLClientProfileResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("A TLS client profile define how to secure a connection remote endpoint.", "ssl-client", "").String,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Name of the object. Must be unique among object types in application domain.", "", "").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 128),
					stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9_-]+$"), "Must match :"+"^[a-zA-Z0-9_-]+$"),
				},
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"app_domain": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("The name of the application domain the object belongs to", "", "").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 128),
					stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9_-]+$"), "Must match :"+"^[a-zA-Z0-9_-]+$"),
				},
				PlanModifiers: []planmodifier.String{
					modifiers.ImmutableAfterSet(),
				},
			},
			"user_summary": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Comments", "summary", "").String,
				Optional:            true,
			},
			"protocols": models.GetDmSSLProtoVersionsBitmapResourceSchema("Protocols", "protocols", "", false),
			"ciphers": schema.ListAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the list of cipher suites to support in preference order. Ensure that the displayed order of cipher suites are in preference order. Otherwise, change the sequence to meet your preference.", "ciphers", "").AddStringEnum("RSA_WITH_NULL_MD5", "RSA_WITH_NULL_SHA", "RSA_WITH_RC4_128_MD5", "RSA_WITH_RC4_128_SHA", "RSA_WITH_DES_CBC_SHA", "RSA_WITH_3DES_EDE_CBC_SHA", "DHE_DSS_WITH_DES_CBC_SHA", "DHE_DSS_WITH_3DES_EDE_CBC_SHA", "DHE_RSA_WITH_DES_CBC_SHA", "DHE_RSA_WITH_3DES_EDE_CBC_SHA", "RSA_WITH_AES_128_CBC_SHA", "DHE_DSS_WITH_AES_128_CBC_SHA", "DHE_RSA_WITH_AES_128_CBC_SHA", "RSA_WITH_AES_256_CBC_SHA", "DHE_DSS_WITH_AES_256_CBC_SHA", "DHE_RSA_WITH_AES_256_CBC_SHA", "RSA_WITH_NULL_SHA256", "RSA_WITH_AES_128_CBC_SHA256", "RSA_WITH_AES_256_CBC_SHA256", "DHE_DSS_WITH_AES_128_CBC_SHA256", "DHE_RSA_WITH_AES_128_CBC_SHA256", "DHE_DSS_WITH_AES_256_CBC_SHA256", "DHE_RSA_WITH_AES_256_CBC_SHA256", "RSA_WITH_AES_128_GCM_SHA256", "RSA_WITH_AES_256_GCM_SHA384", "DHE_RSA_WITH_AES_128_GCM_SHA256", "DHE_RSA_WITH_AES_256_GCM_SHA384", "DHE_DSS_WITH_AES_128_GCM_SHA256", "DHE_DSS_WITH_AES_256_GCM_SHA384", "AES_128_GCM_SHA256", "AES_256_GCM_SHA384", "CHACHA20_POLY1305_SHA256", "AES_128_CCM_SHA256", "AES_128_CCM_8_SHA256", "ECDHE_RSA_WITH_NULL_SHA", "ECDHE_RSA_WITH_RC4_128_SHA", "ECDHE_RSA_WITH_3DES_EDE_CBC_SHA", "ECDHE_RSA_WITH_AES_128_CBC_SHA", "ECDHE_RSA_WITH_AES_256_CBC_SHA", "ECDHE_RSA_WITH_AES_128_CBC_SHA256", "ECDHE_RSA_WITH_AES_256_CBC_SHA384", "ECDHE_RSA_WITH_AES_128_GCM_SHA256", "ECDHE_RSA_WITH_AES_256_GCM_SHA384", "ECDHE_ECDSA_WITH_NULL_SHA", "ECDHE_ECDSA_WITH_RC4_128_SHA", "ECDHE_ECDSA_WITH_3DES_EDE_CBC_SHA", "ECDHE_ECDSA_WITH_AES_128_CBC_SHA", "ECDHE_ECDSA_WITH_AES_256_CBC_SHA", "ECDHE_ECDSA_WITH_AES_128_CBC_SHA256", "ECDHE_ECDSA_WITH_AES_256_CBC_SHA384", "ECDHE_ECDSA_WITH_AES_128_GCM_SHA256", "ECDHE_ECDSA_WITH_AES_256_GCM_SHA384").String,
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Validators: []validator.List{
					listvalidator.ValueStringsAre(
						stringvalidator.OneOf("RSA_WITH_NULL_MD5", "RSA_WITH_NULL_SHA", "RSA_WITH_RC4_128_MD5", "RSA_WITH_RC4_128_SHA", "RSA_WITH_DES_CBC_SHA", "RSA_WITH_3DES_EDE_CBC_SHA", "DHE_DSS_WITH_DES_CBC_SHA", "DHE_DSS_WITH_3DES_EDE_CBC_SHA", "DHE_RSA_WITH_DES_CBC_SHA", "DHE_RSA_WITH_3DES_EDE_CBC_SHA", "RSA_WITH_AES_128_CBC_SHA", "DHE_DSS_WITH_AES_128_CBC_SHA", "DHE_RSA_WITH_AES_128_CBC_SHA", "RSA_WITH_AES_256_CBC_SHA", "DHE_DSS_WITH_AES_256_CBC_SHA", "DHE_RSA_WITH_AES_256_CBC_SHA", "RSA_WITH_NULL_SHA256", "RSA_WITH_AES_128_CBC_SHA256", "RSA_WITH_AES_256_CBC_SHA256", "DHE_DSS_WITH_AES_128_CBC_SHA256", "DHE_RSA_WITH_AES_128_CBC_SHA256", "DHE_DSS_WITH_AES_256_CBC_SHA256", "DHE_RSA_WITH_AES_256_CBC_SHA256", "RSA_WITH_AES_128_GCM_SHA256", "RSA_WITH_AES_256_GCM_SHA384", "DHE_RSA_WITH_AES_128_GCM_SHA256", "DHE_RSA_WITH_AES_256_GCM_SHA384", "DHE_DSS_WITH_AES_128_GCM_SHA256", "DHE_DSS_WITH_AES_256_GCM_SHA384", "AES_128_GCM_SHA256", "AES_256_GCM_SHA384", "CHACHA20_POLY1305_SHA256", "AES_128_CCM_SHA256", "AES_128_CCM_8_SHA256", "ECDHE_RSA_WITH_NULL_SHA", "ECDHE_RSA_WITH_RC4_128_SHA", "ECDHE_RSA_WITH_3DES_EDE_CBC_SHA", "ECDHE_RSA_WITH_AES_128_CBC_SHA", "ECDHE_RSA_WITH_AES_256_CBC_SHA", "ECDHE_RSA_WITH_AES_128_CBC_SHA256", "ECDHE_RSA_WITH_AES_256_CBC_SHA384", "ECDHE_RSA_WITH_AES_128_GCM_SHA256", "ECDHE_RSA_WITH_AES_256_GCM_SHA384", "ECDHE_ECDSA_WITH_NULL_SHA", "ECDHE_ECDSA_WITH_RC4_128_SHA", "ECDHE_ECDSA_WITH_3DES_EDE_CBC_SHA", "ECDHE_ECDSA_WITH_AES_128_CBC_SHA", "ECDHE_ECDSA_WITH_AES_256_CBC_SHA", "ECDHE_ECDSA_WITH_AES_128_CBC_SHA256", "ECDHE_ECDSA_WITH_AES_256_CBC_SHA384", "ECDHE_ECDSA_WITH_AES_128_GCM_SHA256", "ECDHE_ECDSA_WITH_AES_256_GCM_SHA384"),
					),
				},
				Default: listdefault.StaticValue(types.ListValueMust(types.StringType, []attr.Value{
					types.StringValue("AES_256_GCM_SHA384"),
					types.StringValue("CHACHA20_POLY1305_SHA256"),
					types.StringValue("AES_128_GCM_SHA256"),
					types.StringValue("ECDHE_ECDSA_WITH_AES_256_GCM_SHA384"),
					types.StringValue("ECDHE_ECDSA_WITH_AES_256_CBC_SHA384"),
					types.StringValue("ECDHE_ECDSA_WITH_AES_128_GCM_SHA256"),
					types.StringValue("ECDHE_ECDSA_WITH_AES_128_CBC_SHA256"),
					types.StringValue("ECDHE_ECDSA_WITH_AES_256_CBC_SHA"),
					types.StringValue("ECDHE_ECDSA_WITH_AES_128_CBC_SHA"),
					types.StringValue("ECDHE_RSA_WITH_AES_256_GCM_SHA384"),
					types.StringValue("ECDHE_RSA_WITH_AES_256_CBC_SHA384"),
					types.StringValue("ECDHE_RSA_WITH_AES_128_GCM_SHA256"),
					types.StringValue("ECDHE_RSA_WITH_AES_128_CBC_SHA256"),
					types.StringValue("ECDHE_RSA_WITH_AES_256_CBC_SHA"),
					types.StringValue("ECDHE_RSA_WITH_AES_128_CBC_SHA"),
					types.StringValue("DHE_RSA_WITH_AES_256_GCM_SHA384"),
					types.StringValue("DHE_RSA_WITH_AES_256_CBC_SHA256"),
					types.StringValue("DHE_RSA_WITH_AES_128_GCM_SHA256"),
					types.StringValue("DHE_RSA_WITH_AES_128_CBC_SHA256"),
					types.StringValue("DHE_RSA_WITH_AES_256_CBC_SHA"),
					types.StringValue("DHE_RSA_WITH_AES_128_CBC_SHA"),
				})),
			},
			"idcred": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Identification credentials", "idcred", "crypto_ident_cred").String,
				Optional:            true,
			},
			"validate_server_cert": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Validate server certificate", "validate-server-cert", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"valcred": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Validation credentials", "valcred", "crypto_val_cred").AddRequiredWhen(models.SSLClientProfileValcredCondVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.SSLClientProfileValcredCondVal, validators.Evaluation{}, false),
				},
			},
			"caching": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enable session caching", "caching", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"cache_timeout": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the duration in seconds that TLS sessions remain in the session cache before they are removed. Enter a value in the range 1 - 86400. The default value is 300.", "cache-timeout", "").AddIntegerRange(1, 86400).AddDefaultValue("300").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 86400),
				},
				Default: int64default.StaticInt64(300),
			},
			"cache_size": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the maximum number of entries of the session cache. Enter a value in the range 1 - 500000. The default value is 100.", "cache-size", "").AddIntegerRange(1, 500000).AddDefaultValue("100").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 500000),
				},
				Default: int64default.StaticInt64(100),
			},
			"ssl_client_features": models.GetDmSSLClientFeaturesResourceSchema("Features", "ssl-client-features", "", false),
			"elliptic_curves": schema.ListAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Elliptic curves", "curves", "").AddStringEnum("sect163k1", "sect163r1", "sect163r2", "sect193r1", "sect193r2", "sect233k1", "sect233r1", "sect239k1", "sect283k1", "sect283r1", "sect409k1", "sect409r1", "sect571k1", "sect571r1", "secp160k1", "secp160r1", "secp160r2", "secp192k1", "secp192r1", "secp224k1", "secp224r1", "secp256k1", "secp256r1", "secp384r1", "secp521r1", "brainpoolP256r1", "brainpoolP384r1", "brainpoolP512r1").String,
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Validators: []validator.List{
					listvalidator.ValueStringsAre(
						stringvalidator.OneOf("sect163k1", "sect163r1", "sect163r2", "sect193r1", "sect193r2", "sect233k1", "sect233r1", "sect239k1", "sect283k1", "sect283r1", "sect409k1", "sect409r1", "sect571k1", "sect571r1", "secp160k1", "secp160r1", "secp160r2", "secp192k1", "secp192r1", "secp224k1", "secp224r1", "secp256k1", "secp256r1", "secp384r1", "secp521r1", "brainpoolP256r1", "brainpoolP384r1", "brainpoolP512r1"),
					),
				},
				Default: listdefault.StaticValue(types.ListValueMust(types.StringType, []attr.Value{
					types.StringValue("secp521r1"),
					types.StringValue("secp384r1"),
					types.StringValue("secp256k1"),
					types.StringValue("secp256r1"),
				})),
			},
			"use_custom_sni_hostname": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to use a custom server name in the SNI extension in the TLS client <tt>hello</tt> message. By default, the hostname of the target is used in the SNI extension.", "use-custom-sni-hostname", "").AddDefaultValue("false").AddRequiredWhen(models.SSLClientProfileUseCustomSNIHostnameCondVal.String()).String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"custom_sni_hostname": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Custom SNI hostname", "custom-sni-hostname", "").AddRequiredWhen(models.SSLClientProfileCustomSNIHostnameCondVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.SSLClientProfileCustomSNIHostnameCondVal, validators.Evaluation{}, false),
				},
			},
			"validate_hostname": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Validate server hostname", "validate-hostname", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"hostname_validation_flags": models.GetDmSSLHostnameValidationFlagsResourceSchema("Specify the flags that fine tune the validation methods and settings during the handshake. The default behavior uses the subject DN only when the <tt>Subject Alternative Name</tt> (SAN) extension contains no DNS name.", "hostname-validation-flags", "", false),
			"hostname_validation_fail_on_error": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to terminate the handshake when hostname validation fails or to ignore the failure, log an event, and continue with server certificate validation. The default behavior is to ignore the failure, log an event, and continue with any configured server certificate validation.", "hostname-validation-fail", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"enable_tls13_compat": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to send extra TLS messages to look more like TLSv1.2. Some network middleboxes might not recognize TLSv1.3 and drop the connection. Enable this option to send dummy Change Cipher Spec messages, which makes TLSv1.3 look more like TLSv1.2.", "enable-tls13-compat", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"disable_renegotiation": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Disable renegotiation", "disable-renegotiation", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"sig_algs": schema.ListAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the list of signature algorithms to advertise and support. An empty list implies the use of all of the default algorithms.", "sign-alg", "").AddStringEnum("ecdsa_secp256r1_sha256", "ecdsa_secp384r1_sha384", "ecdsa_secp521r1_sha512", "ed25519", "ed448", "ecdsa_sha224", "ecdsa_sha1", "rsa_pss_rsae_sha256", "rsa_pss_rsae_sha384", "rsa_pss_rsae_sha512", "rsa_pss_pss_sha256", "rsa_pss_pss_sha384", "rsa_pss_pss_sha512", "rsa_pkcs1_sha256", "rsa_pkcs1_sha384", "rsa_pkcs1_sha512", "rsa_pkcs1_sha224", "rsa_pkcs1_sha1", "dsa_sha256", "dsa_sha384", "dsa_sha512", "dsa_sha224", "dsa_sha1").String,
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Validators: []validator.List{
					listvalidator.ValueStringsAre(
						stringvalidator.OneOf("ecdsa_secp256r1_sha256", "ecdsa_secp384r1_sha384", "ecdsa_secp521r1_sha512", "ed25519", "ed448", "ecdsa_sha224", "ecdsa_sha1", "rsa_pss_rsae_sha256", "rsa_pss_rsae_sha384", "rsa_pss_rsae_sha512", "rsa_pss_pss_sha256", "rsa_pss_pss_sha384", "rsa_pss_pss_sha512", "rsa_pkcs1_sha256", "rsa_pkcs1_sha384", "rsa_pkcs1_sha512", "rsa_pkcs1_sha224", "rsa_pkcs1_sha1", "dsa_sha256", "dsa_sha384", "dsa_sha512", "dsa_sha224", "dsa_sha1"),
					),
				},
				Default: listdefault.StaticValue(types.ListValueMust(types.StringType, []attr.Value{
					types.StringValue("ecdsa_secp256r1_sha256"),
					types.StringValue("ecdsa_secp384r1_sha384"),
					types.StringValue("ecdsa_secp521r1_sha512"),
					types.StringValue("ed25519"),
					types.StringValue("ed448"),
					types.StringValue("rsa_pss_rsae_sha256"),
					types.StringValue("rsa_pss_rsae_sha384"),
					types.StringValue("rsa_pss_rsae_sha512"),
					types.StringValue("rsa_pss_pss_sha256"),
					types.StringValue("rsa_pss_pss_sha384"),
					types.StringValue("rsa_pss_pss_sha512"),
					types.StringValue("rsa_pkcs1_sha256"),
					types.StringValue("rsa_pkcs1_sha384"),
					types.StringValue("rsa_pkcs1_sha512"),
					types.StringValue("ecdsa_sha224"),
					types.StringValue("ecdsa_sha1"),
					types.StringValue("rsa_pkcs1_sha224"),
					types.StringValue("rsa_pkcs1_sha1"),
					types.StringValue("dsa_sha256"),
					types.StringValue("dsa_sha384"),
					types.StringValue("dsa_sha512"),
				})),
			},
			"require_closure_notification": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Require closure notification", "require-closure-notification", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"dependency_actions": actions.ActionsSchema,
		},
	}
}

func (r *SSLClientProfileResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (r *SSLClientProfileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.SSLClientProfile
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Create, false)
	if resp.Diagnostics.HasError() {
		return
	}

	body := data.ToBody(ctx, `SSLClientProfile`)
	_, err := r.pData.Client.Post(data.GetPath(), body)

	if err != nil && !strings.Contains(err.Error(), "status 409") {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to create object (%s), got error: %s", "POST", err))
		return
	}
	actions.PostProcess(ctx, &resp.Diagnostics, data.DependencyActions, actions.Create)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SSLClientProfileResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.SSLClientProfile
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	res, err := r.pData.Client.Get(data.GetPath() + "/" + data.Id.ValueString())
	if err != nil && (strings.Contains(err.Error(), "status 404") || strings.Contains(err.Error(), "status 406") || strings.Contains(err.Error(), "status 500") || strings.Contains(err.Error(), "status 400")) {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s", err))
		return
	}

	if data.IsNull() {
		// Import
		data.FromBody(ctx, `SSLClientProfile`, res)
	} else {
		// Update
		data.UpdateFromBody(ctx, `SSLClientProfile`, res)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SSLClientProfileResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.SSLClientProfile
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Update, false)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.pData.Client.Put(data.GetPath()+"/"+data.Id.ValueString(), data.ToBody(ctx, `SSLClientProfile`))
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to update object (PUT), got error: %s", err))
		return
	}

	actions.PostProcess(ctx, &resp.Diagnostics, data.DependencyActions, actions.Update)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SSLClientProfileResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.SSLClientProfile
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Delete, false)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.pData.Client.Delete(data.GetPath() + "/" + data.Id.ValueString())
	if err != nil && !strings.Contains(err.Error(), "status 404") && !strings.Contains(err.Error(), "status 409") {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object (DELETE), got error: %s", err))
		return
	}

	actions.PostProcess(ctx, &resp.Diagnostics, data.DependencyActions, actions.Delete)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.State.RemoveResource(ctx)
}

func (r *SSLClientProfileResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data models.SSLClientProfile

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.ValidateConfig(ctx, &resp.Diagnostics, data.DependencyActions)
}

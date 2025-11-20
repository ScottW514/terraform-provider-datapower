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
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
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

var _ resource.Resource = &SSLServerProfileResource{}
var _ resource.ResourceWithImportState = &SSLServerProfileResource{}

func NewSSLServerProfileResource() resource.Resource {
	return &SSLServerProfileResource{}
}

type SSLServerProfileResource struct {
	pData *tfutils.ProviderData
}

func (r *SSLServerProfileResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ssl_server_profile"
}

func (r *SSLServerProfileResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("The TLS server profile secures connections with clients.", "ssl-server", "").String,
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
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the cipher suites to support in preference order. Ensure that the order of cipher suites are in preference order. Otherwise, change the sequence order to meet your preference.", "ciphers", "").AddStringEnum("RSA_WITH_NULL_MD5", "RSA_WITH_NULL_SHA", "RSA_WITH_RC4_128_MD5", "RSA_WITH_RC4_128_SHA", "RSA_WITH_DES_CBC_SHA", "RSA_WITH_3DES_EDE_CBC_SHA", "DHE_DSS_WITH_DES_CBC_SHA", "DHE_DSS_WITH_3DES_EDE_CBC_SHA", "DHE_RSA_WITH_DES_CBC_SHA", "DHE_RSA_WITH_3DES_EDE_CBC_SHA", "RSA_WITH_AES_128_CBC_SHA", "DHE_DSS_WITH_AES_128_CBC_SHA", "DHE_RSA_WITH_AES_128_CBC_SHA", "RSA_WITH_AES_256_CBC_SHA", "DHE_DSS_WITH_AES_256_CBC_SHA", "DHE_RSA_WITH_AES_256_CBC_SHA", "RSA_WITH_NULL_SHA256", "RSA_WITH_AES_128_CBC_SHA256", "RSA_WITH_AES_256_CBC_SHA256", "DHE_DSS_WITH_AES_128_CBC_SHA256", "DHE_RSA_WITH_AES_128_CBC_SHA256", "DHE_DSS_WITH_AES_256_CBC_SHA256", "DHE_RSA_WITH_AES_256_CBC_SHA256", "RSA_WITH_AES_128_GCM_SHA256", "RSA_WITH_AES_256_GCM_SHA384", "DHE_RSA_WITH_AES_128_GCM_SHA256", "DHE_RSA_WITH_AES_256_GCM_SHA384", "DHE_DSS_WITH_AES_128_GCM_SHA256", "DHE_DSS_WITH_AES_256_GCM_SHA384", "AES_128_GCM_SHA256", "AES_256_GCM_SHA384", "CHACHA20_POLY1305_SHA256", "AES_128_CCM_SHA256", "AES_128_CCM_8_SHA256", "ECDHE_RSA_WITH_NULL_SHA", "ECDHE_RSA_WITH_RC4_128_SHA", "ECDHE_RSA_WITH_3DES_EDE_CBC_SHA", "ECDHE_RSA_WITH_AES_128_CBC_SHA", "ECDHE_RSA_WITH_AES_256_CBC_SHA", "ECDHE_RSA_WITH_AES_128_CBC_SHA256", "ECDHE_RSA_WITH_AES_256_CBC_SHA384", "ECDHE_RSA_WITH_AES_128_GCM_SHA256", "ECDHE_RSA_WITH_AES_256_GCM_SHA384", "ECDHE_ECDSA_WITH_NULL_SHA", "ECDHE_ECDSA_WITH_RC4_128_SHA", "ECDHE_ECDSA_WITH_3DES_EDE_CBC_SHA", "ECDHE_ECDSA_WITH_AES_128_CBC_SHA", "ECDHE_ECDSA_WITH_AES_256_CBC_SHA", "ECDHE_ECDSA_WITH_AES_128_CBC_SHA256", "ECDHE_ECDSA_WITH_AES_256_CBC_SHA384", "ECDHE_ECDSA_WITH_AES_128_GCM_SHA256", "ECDHE_ECDSA_WITH_AES_256_GCM_SHA384").String,
				ElementType:         types.StringType,
				Required:            true,
				Validators: []validator.List{
					listvalidator.ValueStringsAre(
						stringvalidator.OneOf("RSA_WITH_NULL_MD5", "RSA_WITH_NULL_SHA", "RSA_WITH_RC4_128_MD5", "RSA_WITH_RC4_128_SHA", "RSA_WITH_DES_CBC_SHA", "RSA_WITH_3DES_EDE_CBC_SHA", "DHE_DSS_WITH_DES_CBC_SHA", "DHE_DSS_WITH_3DES_EDE_CBC_SHA", "DHE_RSA_WITH_DES_CBC_SHA", "DHE_RSA_WITH_3DES_EDE_CBC_SHA", "RSA_WITH_AES_128_CBC_SHA", "DHE_DSS_WITH_AES_128_CBC_SHA", "DHE_RSA_WITH_AES_128_CBC_SHA", "RSA_WITH_AES_256_CBC_SHA", "DHE_DSS_WITH_AES_256_CBC_SHA", "DHE_RSA_WITH_AES_256_CBC_SHA", "RSA_WITH_NULL_SHA256", "RSA_WITH_AES_128_CBC_SHA256", "RSA_WITH_AES_256_CBC_SHA256", "DHE_DSS_WITH_AES_128_CBC_SHA256", "DHE_RSA_WITH_AES_128_CBC_SHA256", "DHE_DSS_WITH_AES_256_CBC_SHA256", "DHE_RSA_WITH_AES_256_CBC_SHA256", "RSA_WITH_AES_128_GCM_SHA256", "RSA_WITH_AES_256_GCM_SHA384", "DHE_RSA_WITH_AES_128_GCM_SHA256", "DHE_RSA_WITH_AES_256_GCM_SHA384", "DHE_DSS_WITH_AES_128_GCM_SHA256", "DHE_DSS_WITH_AES_256_GCM_SHA384", "AES_128_GCM_SHA256", "AES_256_GCM_SHA384", "CHACHA20_POLY1305_SHA256", "AES_128_CCM_SHA256", "AES_128_CCM_8_SHA256", "ECDHE_RSA_WITH_NULL_SHA", "ECDHE_RSA_WITH_RC4_128_SHA", "ECDHE_RSA_WITH_3DES_EDE_CBC_SHA", "ECDHE_RSA_WITH_AES_128_CBC_SHA", "ECDHE_RSA_WITH_AES_256_CBC_SHA", "ECDHE_RSA_WITH_AES_128_CBC_SHA256", "ECDHE_RSA_WITH_AES_256_CBC_SHA384", "ECDHE_RSA_WITH_AES_128_GCM_SHA256", "ECDHE_RSA_WITH_AES_256_GCM_SHA384", "ECDHE_ECDSA_WITH_NULL_SHA", "ECDHE_ECDSA_WITH_RC4_128_SHA", "ECDHE_ECDSA_WITH_3DES_EDE_CBC_SHA", "ECDHE_ECDSA_WITH_AES_128_CBC_SHA", "ECDHE_ECDSA_WITH_AES_256_CBC_SHA", "ECDHE_ECDSA_WITH_AES_128_CBC_SHA256", "ECDHE_ECDSA_WITH_AES_256_CBC_SHA384", "ECDHE_ECDSA_WITH_AES_128_GCM_SHA256", "ECDHE_ECDSA_WITH_AES_256_GCM_SHA384"),
					),
				},
			},
			"idcred": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Identification credentials", "idcred", "crypto_ident_cred").String,
				Required:            true,
			},
			"request_client_auth": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Request client authentication", "request-client-auth", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"require_client_auth": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to require client authentication during the TLS handshake. When enabled, the handshake is aborted if the client certificate is not provided. Otherwise, the request does not fail when there is no client certificate.", "require-client-auth", "").AddDefaultValue("true").AddNotValidWhen(models.SSLServerProfileRequireClientAuthIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(true),
			},
			"validate_client_cert": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Validate client certificate", "validate-client-cert", "").AddDefaultValue("true").AddNotValidWhen(models.SSLServerProfileValidateClientCertIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(true),
			},
			"send_client_auth_ca_list": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Send client authentication CA list", "send-client-auth-ca-list", "").AddDefaultValue("true").AddNotValidWhen(models.SSLServerProfileSendClientAuthCAListIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(true),
			},
			"valcred": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Validation credentials", "valcred", "crypto_val_cred").AddRequiredWhen(models.SSLServerProfileValcredCondVal.String()).AddNotValidWhen(models.SSLServerProfileValcredIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.SSLServerProfileValcredCondVal, models.SSLServerProfileValcredIgnoreVal, false),
				},
			},
			"caching": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enable session caching", "caching", "").AddDefaultValue("true").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(true),
			},
			"cache_timeout": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the number of seconds that TLS sessions are allowed to remain in the TLS session cache before they are removed. Enter a value in the range 1 - 86400. The default value is 300.", "cache-timeout", "").AddIntegerRange(1, 86400).AddDefaultValue("300").AddNotValidWhen(models.SSLServerProfileCacheTimeoutIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 86400),
					validators.ConditionalRequiredInt64(validators.Evaluation{}, models.SSLServerProfileCacheTimeoutIgnoreVal, true),
				},
				Default: int64default.StaticInt64(300),
			},
			"cache_size": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the maximum number of entries (multiplied by 1024) in the session cache. Enter a value in the range 1 - 500. The default value is 20.", "cache-size", "").AddIntegerRange(1, 500).AddDefaultValue("20").AddNotValidWhen(models.SSLServerProfileCacheSizeIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 500),
					validators.ConditionalRequiredInt64(validators.Evaluation{}, models.SSLServerProfileCacheSizeIgnoreVal, true),
				},
				Default: int64default.StaticInt64(20),
			},
			"ssl_options": models.GetDmSSLOptionsResourceSchema("Specify the options to apply to the TLS connection. These options have negative impact on the performance.", "ssl-options", "", false),
			"max_ssl_duration": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the maximum duration of an established TLS session. The TLS connection terminates when the duration is reached. Enter a value in the range 1 - 11520. The default value is 60.", "max-duration", "").AddIntegerRange(1, 11520).AddDefaultValue("60").AddNotValidWhen(models.SSLServerProfileMaxSSLDurationIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 11520),
					validators.ConditionalRequiredInt64(validators.Evaluation{}, models.SSLServerProfileMaxSSLDurationIgnoreVal, true),
				},
				Default: int64default.StaticInt64(60),
			},
			"disable_renegotiation": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Disable renegotiation", "disable-renegotiation", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"number_of_renegotiation_allowed": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the maximum number of client initiated renegotiations to allow. Enter a value in the range 0 - 512. The default value is 0, which indicates TLS client initiated renegotiation is not allowed.", "max-renegotiation-allowed", "").AddIntegerRange(0, 512).AddDefaultValue("0").AddNotValidWhen(models.SSLServerProfileNumberOfRenegotiationAllowedIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 512),
					validators.ConditionalRequiredInt64(validators.Evaluation{}, models.SSLServerProfileNumberOfRenegotiationAllowedIgnoreVal, true),
				},
				Default: int64default.StaticInt64(0),
			},
			"prohibit_resume_on_reneg": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Prohibit session resumption on renegotiation", "prohibit-resume-on-reneg", "").AddDefaultValue("false").AddNotValidWhen(models.SSLServerProfileProhibitResumeOnRenegIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"compression": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to enable TLS compression. TLS compression in HTTPS is dangerous because the connection becomes vulnerable to the CRIME attack.", "compression", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"allow_legacy_renegotiation": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to allow TLS renegotiation with TLS clients that do not support RFC 5746. By default, this support is disabled because renegotiation is vulnerable to man-in-the-middle attacks as documented in CVE-2009-3555. Renegotiation with TLS clients that support RFC 5746 is permitted regardless of the setting.", "allow-legacy-renegotiation", "").AddDefaultValue("false").AddNotValidWhen(models.SSLServerProfileAllowLegacyRenegotiationIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"prefer_server_ciphers": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Use server cipher suite preferences", "prefer-server-ciphers", "").AddDefaultValue("true").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(true),
			},
			"elliptic_curves": schema.ListAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Elliptic curves", "curves", "").AddStringEnum("sect163k1", "sect163r1", "sect163r2", "sect193r1", "sect193r2", "sect233k1", "sect233r1", "sect239k1", "sect283k1", "sect283r1", "sect409k1", "sect409r1", "sect571k1", "sect571r1", "secp160k1", "secp160r1", "secp160r2", "secp192k1", "secp192r1", "secp224k1", "secp224r1", "secp256k1", "secp256r1", "secp384r1", "secp521r1", "brainpoolP256r1", "brainpoolP384r1", "brainpoolP512r1").String,
				ElementType:         types.StringType,
				Optional:            true,
				Validators: []validator.List{
					listvalidator.ValueStringsAre(
						stringvalidator.OneOf("sect163k1", "sect163r1", "sect163r2", "sect193r1", "sect193r2", "sect233k1", "sect233r1", "sect239k1", "sect283k1", "sect283r1", "sect409k1", "sect409r1", "sect571k1", "sect571r1", "secp160k1", "secp160r1", "secp160r2", "secp192k1", "secp192r1", "secp224k1", "secp224r1", "secp256k1", "secp256r1", "secp384r1", "secp521r1", "brainpoolP256r1", "brainpoolP384r1", "brainpoolP512r1"),
					),
				},
			},
			"prioritize_cha_cha": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to move ChaCha20-Poly1305 cipher to the top of preference list sent to the client when this cipher is at the top of client cipher list When server cipher suite preferences is in effect, enabling this property temporarily moves the ChaCha20-Poly1305 cipher to the top of preference list when clients that present ChaCha20-Poly1305 cipher have this cipher at the top of their preference list. This setting allows the client to negotiate ChaCha20-Poly1305 cipher while other clients can use other ciphers.", "prioritize-chacha", "").AddDefaultValue("false").AddNotValidWhen(models.SSLServerProfilePrioritizeChaChaIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"sig_algs": schema.ListAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the signature algorithms to advertise and support. An empty list uses the default algorithms.", "sign-alg", "").AddStringEnum("ecdsa_secp256r1_sha256", "ecdsa_secp384r1_sha384", "ecdsa_secp521r1_sha512", "ed25519", "ed448", "ecdsa_sha224", "ecdsa_sha1", "rsa_pss_rsae_sha256", "rsa_pss_rsae_sha384", "rsa_pss_rsae_sha512", "rsa_pss_pss_sha256", "rsa_pss_pss_sha384", "rsa_pss_pss_sha512", "rsa_pkcs1_sha256", "rsa_pkcs1_sha384", "rsa_pkcs1_sha512", "rsa_pkcs1_sha224", "rsa_pkcs1_sha1", "dsa_sha256", "dsa_sha384", "dsa_sha512", "dsa_sha224", "dsa_sha1").String,
				ElementType:         types.StringType,
				Optional:            true,
				Validators: []validator.List{
					listvalidator.ValueStringsAre(
						stringvalidator.OneOf("ecdsa_secp256r1_sha256", "ecdsa_secp384r1_sha384", "ecdsa_secp521r1_sha512", "ed25519", "ed448", "ecdsa_sha224", "ecdsa_sha1", "rsa_pss_rsae_sha256", "rsa_pss_rsae_sha384", "rsa_pss_rsae_sha512", "rsa_pss_pss_sha256", "rsa_pss_pss_sha384", "rsa_pss_pss_sha512", "rsa_pkcs1_sha256", "rsa_pkcs1_sha384", "rsa_pkcs1_sha512", "rsa_pkcs1_sha224", "rsa_pkcs1_sha1", "dsa_sha256", "dsa_sha384", "dsa_sha512", "dsa_sha224", "dsa_sha1"),
					),
				},
			},
			"require_closure_notification": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Require closure notification", "require-closure-notification", "").AddDefaultValue("true").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(true),
			},
			"dependency_actions": actions.ActionsSchema,
		},
	}
}

func (r *SSLServerProfileResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (r *SSLServerProfileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.SSLServerProfile
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Create, false)
	if resp.Diagnostics.HasError() {
		return
	}

	body := data.ToBody(ctx, `SSLServerProfile`)
	_, err := r.pData.Client.Post(data.GetPath(), body)
	if err != nil {
		if strings.Contains(err.Error(), "status 409") {
			_, err := r.pData.Client.Put(data.GetPath()+"/"+data.Id.ValueString(), body)
			if err != nil {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Resource already exists. Failed to update resource, got error: %s", err))
				return
			}
			resp.Diagnostics.AddWarning("Warning", "Resource already exists. Existing resource was updated.")
		} else if strings.Contains(err.Error(), "status 401") {
			_ = tfutils.DomainCredentialTest(r.pData.Client, &resp.Diagnostics, data.AppDomain.ValueString())
			if !resp.Diagnostics.HasError() {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Application Domain '%s' does not exist", data.AppDomain.ValueString()))
			}
			return
		} else {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to create resource, got error: %s", err))
			return
		}
	}
	actions.PostProcess(ctx, &resp.Diagnostics, data.DependencyActions, actions.Create)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SSLServerProfileResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.SSLServerProfile
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	res, err := r.pData.Client.Get(data.GetPath() + "/" + data.Id.ValueString())
	if err != nil {
		if strings.Contains(err.Error(), "status 404") || strings.Contains(err.Error(), "status 406") || strings.Contains(err.Error(), "status 500") || strings.Contains(err.Error(), "status 400") {
			resp.State.RemoveResource(ctx)
			return
		} else if strings.Contains(err.Error(), "status 401") {
			_ = tfutils.DomainCredentialTest(r.pData.Client, &resp.Diagnostics, data.AppDomain.ValueString())
			if !resp.Diagnostics.HasError() {
				resp.State.RemoveResource(ctx)
			}
			return
		} else {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s", err))
			return
		}
	}

	data.UpdateFromBody(ctx, `SSLServerProfile`, res)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SSLServerProfileResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.SSLServerProfile
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Update, false)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.pData.Client.Put(data.GetPath()+"/"+data.Id.ValueString(), data.ToBody(ctx, `SSLServerProfile`))
	if err != nil {
		if strings.Contains(err.Error(), "status 401") {
			_ = tfutils.DomainCredentialTest(r.pData.Client, &resp.Diagnostics, data.AppDomain.ValueString())
			if !resp.Diagnostics.HasError() {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Application Domain '%s' does not exist", data.AppDomain.ValueString()))
			}
			return
		} else {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to update object (PUT), got error: %s", err))
			return
		}
	}

	actions.PostProcess(ctx, &resp.Diagnostics, data.DependencyActions, actions.Update)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SSLServerProfileResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.SSLServerProfile
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Delete, false)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.pData.Client.Delete(data.GetPath() + "/" + data.Id.ValueString())
	if err != nil {
		if strings.Contains(err.Error(), "status 409") {
			resp.Diagnostics.AddWarning("Resource Conflict", fmt.Sprintf("Resource is no longer tracked by Terraform, but may need to be manually deleted on DataPower host. Got error: %s", err))
		} else if strings.Contains(err.Error(), "status 401") {
			_ = tfutils.DomainCredentialTest(r.pData.Client, &resp.Diagnostics, data.AppDomain.ValueString())
			if !resp.Diagnostics.HasError() {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Application Domain '%s' does not exist", data.AppDomain.ValueString()))
			}
			return
		} else if !strings.Contains(err.Error(), "status 404") {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete resource, got error: %s", err))
			return
		}
	}

	actions.PostProcess(ctx, &resp.Diagnostics, data.DependencyActions, actions.Delete)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.State.RemoveResource(ctx)
}

func (r *SSLServerProfileResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()
	parts := strings.Split(req.ID, "/")
	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		resp.Diagnostics.AddError("Invalid Import ID", "Expected format: <app_domain>/<id>. Got: "+req.ID)
		return
	}

	appDomain := parts[0]
	id := parts[1]

	if !regexp.MustCompile("^[a-zA-Z0-9_-]+$").MatchString(id) || len(id) < 1 || len(id) > 128 {
		resp.Diagnostics.AddError("Invalid ID", "ID must be 1-128 characters and match pattern ^[a-zA-Z0-9_-]+$")
		return
	}
	if !regexp.MustCompile("^[a-zA-Z0-9_-]+$").MatchString(appDomain) || len(appDomain) < 1 || len(appDomain) > 128 {
		resp.Diagnostics.AddError("Invalid Application Domain", "Application domain must be 1-128 characters and match pattern ^[a-zA-Z0-9_-]+$")
		return
	}

	var data models.SSLServerProfile
	data.AppDomain = types.StringValue(appDomain)
	data.Id = types.StringValue(id)
	res, err := r.pData.Client.Get(data.GetPath() + "/" + data.Id.ValueString())
	if err != nil {
		if strings.Contains(err.Error(), "status 404") {
			resp.Diagnostics.AddError("Resource Not Found", fmt.Sprintf("Resource was not found, got error: %s", err))
		} else if strings.Contains(err.Error(), "status 401") {
			_ = tfutils.DomainCredentialTest(r.pData.Client, &resp.Diagnostics, data.AppDomain.ValueString())
			if !resp.Diagnostics.HasError() {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Application Domain '%s' does not exist", data.AppDomain.ValueString()))
			}
		} else {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		}
		return
	}

	data.FromBody(ctx, `SSLServerProfile`, res)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SSLServerProfileResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data models.SSLServerProfile

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.ValidateConfig(ctx, &resp.Diagnostics, data.DependencyActions)
}

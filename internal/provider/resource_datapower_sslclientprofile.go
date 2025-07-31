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
	"github.com/scottw514/terraform-provider-datapower/client"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
)

var _ resource.Resource = &SSLClientProfileResource{}

func NewSSLClientProfileResource() resource.Resource {
	return &SSLClientProfileResource{}
}

type SSLClientProfileResource struct {
	client *client.DatapowerClient
}

func (r *SSLClientProfileResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sslclientprofile"
}

func (r *SSLClientProfileResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("TLS client profile", "ssl-client", "").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Name of the object. Must be unique among object types in application domain.", "", "").String,
				Required:            true,
				Validators: []validator.String{

					stringvalidator.LengthBetween(1, 128),
					stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z0-9_-]+$`), ""),
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
					stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z0-9_-]+$`), ""),
				},
				PlanModifiers: []planmodifier.String{
					ImmutableAfterSet(),
				},
			},
			"user_summary": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Comments", "summary", "").String,
				Optional:            true,
			},
			"protocols": models.GetDmSSLProtoVersionsBitmapResourceSchema("Protocols", "protocols", "", false),
			"ciphers": schema.ListAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Ciphers", "ciphers", "").AddStringEnum("RSA_WITH_NULL_MD5", "RSA_WITH_NULL_SHA", "RSA_WITH_RC4_128_MD5", "RSA_WITH_RC4_128_SHA", "RSA_WITH_DES_CBC_SHA", "RSA_WITH_3DES_EDE_CBC_SHA", "DHE_DSS_WITH_DES_CBC_SHA", "DHE_DSS_WITH_3DES_EDE_CBC_SHA", "DHE_RSA_WITH_DES_CBC_SHA", "DHE_RSA_WITH_3DES_EDE_CBC_SHA", "RSA_WITH_AES_128_CBC_SHA", "DHE_DSS_WITH_AES_128_CBC_SHA", "DHE_RSA_WITH_AES_128_CBC_SHA", "RSA_WITH_AES_256_CBC_SHA", "DHE_DSS_WITH_AES_256_CBC_SHA", "DHE_RSA_WITH_AES_256_CBC_SHA", "RSA_WITH_NULL_SHA256", "RSA_WITH_AES_128_CBC_SHA256", "RSA_WITH_AES_256_CBC_SHA256", "DHE_DSS_WITH_AES_128_CBC_SHA256", "DHE_RSA_WITH_AES_128_CBC_SHA256", "DHE_DSS_WITH_AES_256_CBC_SHA256", "DHE_RSA_WITH_AES_256_CBC_SHA256", "RSA_WITH_AES_128_GCM_SHA256", "RSA_WITH_AES_256_GCM_SHA384", "DHE_RSA_WITH_AES_128_GCM_SHA256", "DHE_RSA_WITH_AES_256_GCM_SHA384", "DHE_DSS_WITH_AES_128_GCM_SHA256", "DHE_DSS_WITH_AES_256_GCM_SHA384", "AES_128_GCM_SHA256", "AES_256_GCM_SHA384", "CHACHA20_POLY1305_SHA256", "AES_128_CCM_SHA256", "AES_128_CCM_8_SHA256", "ECDHE_RSA_WITH_NULL_SHA", "ECDHE_RSA_WITH_RC4_128_SHA", "ECDHE_RSA_WITH_3DES_EDE_CBC_SHA", "ECDHE_RSA_WITH_AES_128_CBC_SHA", "ECDHE_RSA_WITH_AES_256_CBC_SHA", "ECDHE_RSA_WITH_AES_128_CBC_SHA256", "ECDHE_RSA_WITH_AES_256_CBC_SHA384", "ECDHE_RSA_WITH_AES_128_GCM_SHA256", "ECDHE_RSA_WITH_AES_256_GCM_SHA384", "ECDHE_ECDSA_WITH_NULL_SHA", "ECDHE_ECDSA_WITH_RC4_128_SHA", "ECDHE_ECDSA_WITH_3DES_EDE_CBC_SHA", "ECDHE_ECDSA_WITH_AES_128_CBC_SHA", "ECDHE_ECDSA_WITH_AES_256_CBC_SHA", "ECDHE_ECDSA_WITH_AES_128_CBC_SHA256", "ECDHE_ECDSA_WITH_AES_256_CBC_SHA384", "ECDHE_ECDSA_WITH_AES_128_GCM_SHA256", "ECDHE_ECDSA_WITH_AES_256_GCM_SHA384").String,
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Validators: []validator.List{
					listvalidator.ValueStringsAre(stringvalidator.OneOf("RSA_WITH_NULL_MD5", "RSA_WITH_NULL_SHA", "RSA_WITH_RC4_128_MD5", "RSA_WITH_RC4_128_SHA", "RSA_WITH_DES_CBC_SHA", "RSA_WITH_3DES_EDE_CBC_SHA", "DHE_DSS_WITH_DES_CBC_SHA", "DHE_DSS_WITH_3DES_EDE_CBC_SHA", "DHE_RSA_WITH_DES_CBC_SHA", "DHE_RSA_WITH_3DES_EDE_CBC_SHA", "RSA_WITH_AES_128_CBC_SHA", "DHE_DSS_WITH_AES_128_CBC_SHA", "DHE_RSA_WITH_AES_128_CBC_SHA", "RSA_WITH_AES_256_CBC_SHA", "DHE_DSS_WITH_AES_256_CBC_SHA", "DHE_RSA_WITH_AES_256_CBC_SHA", "RSA_WITH_NULL_SHA256", "RSA_WITH_AES_128_CBC_SHA256", "RSA_WITH_AES_256_CBC_SHA256", "DHE_DSS_WITH_AES_128_CBC_SHA256", "DHE_RSA_WITH_AES_128_CBC_SHA256", "DHE_DSS_WITH_AES_256_CBC_SHA256", "DHE_RSA_WITH_AES_256_CBC_SHA256", "RSA_WITH_AES_128_GCM_SHA256", "RSA_WITH_AES_256_GCM_SHA384", "DHE_RSA_WITH_AES_128_GCM_SHA256", "DHE_RSA_WITH_AES_256_GCM_SHA384", "DHE_DSS_WITH_AES_128_GCM_SHA256", "DHE_DSS_WITH_AES_256_GCM_SHA384", "AES_128_GCM_SHA256", "AES_256_GCM_SHA384", "CHACHA20_POLY1305_SHA256", "AES_128_CCM_SHA256", "AES_128_CCM_8_SHA256", "ECDHE_RSA_WITH_NULL_SHA", "ECDHE_RSA_WITH_RC4_128_SHA", "ECDHE_RSA_WITH_3DES_EDE_CBC_SHA", "ECDHE_RSA_WITH_AES_128_CBC_SHA", "ECDHE_RSA_WITH_AES_256_CBC_SHA", "ECDHE_RSA_WITH_AES_128_CBC_SHA256", "ECDHE_RSA_WITH_AES_256_CBC_SHA384", "ECDHE_RSA_WITH_AES_128_GCM_SHA256", "ECDHE_RSA_WITH_AES_256_GCM_SHA384", "ECDHE_ECDSA_WITH_NULL_SHA", "ECDHE_ECDSA_WITH_RC4_128_SHA", "ECDHE_ECDSA_WITH_3DES_EDE_CBC_SHA", "ECDHE_ECDSA_WITH_AES_128_CBC_SHA", "ECDHE_ECDSA_WITH_AES_256_CBC_SHA", "ECDHE_ECDSA_WITH_AES_128_CBC_SHA256", "ECDHE_ECDSA_WITH_AES_256_CBC_SHA384", "ECDHE_ECDSA_WITH_AES_128_GCM_SHA256", "ECDHE_ECDSA_WITH_AES_256_GCM_SHA384")),
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
				MarkdownDescription: tfutils.NewAttributeDescription("Identification credentials", "idcred", "cryptoidentcred").String,
				Optional:            true,
			},
			"validate_server_cert": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Validate server certificate", "validate-server-cert", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"valcred": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Validation credentials", "valcred", "cryptovalcred").String,
				Optional:            true,
			},
			"caching": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enable session caching", "caching", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"cache_timeout": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Session cache timeout", "cache-timeout", "").AddIntegerRange(1, 86400).AddDefaultValue("300").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(1, 86400),
				},
				Default: int64default.StaticInt64(300),
			},
			"cache_size": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Session cache size", "cache-size", "").AddIntegerRange(1, 500000).AddDefaultValue("100").String,
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
					listvalidator.ValueStringsAre(stringvalidator.OneOf("sect163k1", "sect163r1", "sect163r2", "sect193r1", "sect193r2", "sect233k1", "sect233r1", "sect239k1", "sect283k1", "sect283r1", "sect409k1", "sect409r1", "sect571k1", "sect571r1", "secp160k1", "secp160r1", "secp160r2", "secp192k1", "secp192r1", "secp224k1", "secp224r1", "secp256k1", "secp256r1", "secp384r1", "secp521r1", "brainpoolP256r1", "brainpoolP384r1", "brainpoolP512r1")),
				},
				Default: listdefault.StaticValue(types.ListValueMust(types.StringType, []attr.Value{
					types.StringValue("secp521r1"),
					types.StringValue("secp384r1"),
					types.StringValue("secp256k1"),
					types.StringValue("secp256r1"),
				})),
			},
			"use_custom_sni_hostname": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Use custom SNI host name", "use-custom-sni-hostname", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"custom_sni_hostname": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Custom SNI hostname", "custom-sni-hostname", "").String,
				Optional:            true,
			},
			"validate_hostname": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Validate server hostname", "validate-hostname", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"hostname_validation_flags": models.GetDmSSLHostnameValidationFlagsResourceSchema("Hostname validation flags", "hostname-validation-flags", "", false),
			"hostname_validation_fail_on_error": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Hostname validation fail on error", "hostname-validation-fail", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"enable_tls13_compat": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enable TLSv1.3 compatibility", "enable-tls13-compat", "").AddDefaultValue("true").String,
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
				MarkdownDescription: tfutils.NewAttributeDescription("Signature algorithms", "sign-alg", "").AddStringEnum("ecdsa_secp256r1_sha256", "ecdsa_secp384r1_sha384", "ecdsa_secp521r1_sha512", "ed25519", "ed448", "ecdsa_sha224", "ecdsa_sha1", "rsa_pss_rsae_sha256", "rsa_pss_rsae_sha384", "rsa_pss_rsae_sha512", "rsa_pss_pss_sha256", "rsa_pss_pss_sha384", "rsa_pss_pss_sha512", "rsa_pkcs1_sha256", "rsa_pkcs1_sha384", "rsa_pkcs1_sha512", "rsa_pkcs1_sha224", "rsa_pkcs1_sha1", "dsa_sha256", "dsa_sha384", "dsa_sha512", "dsa_sha224", "dsa_sha1").String,
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Validators: []validator.List{
					listvalidator.ValueStringsAre(stringvalidator.OneOf("ecdsa_secp256r1_sha256", "ecdsa_secp384r1_sha384", "ecdsa_secp521r1_sha512", "ed25519", "ed448", "ecdsa_sha224", "ecdsa_sha1", "rsa_pss_rsae_sha256", "rsa_pss_rsae_sha384", "rsa_pss_rsae_sha512", "rsa_pss_pss_sha256", "rsa_pss_pss_sha384", "rsa_pss_pss_sha512", "rsa_pkcs1_sha256", "rsa_pkcs1_sha384", "rsa_pkcs1_sha512", "rsa_pkcs1_sha224", "rsa_pkcs1_sha1", "dsa_sha256", "dsa_sha384", "dsa_sha512", "dsa_sha224", "dsa_sha1")),
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
		},
	}
}

func (r *SSLClientProfileResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = *req.ProviderData.(**client.DatapowerClient)
}

func (r *SSLClientProfileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data models.SSLClientProfile

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	body := data.ToBody(ctx, `SSLClientProfile`)
	_, err := r.client.Post(data.GetPath(), body)

	if err != nil && !strings.Contains(err.Error(), "status 409") {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to create object (%s), got error: %s", "POST", err))
		return
	}

	_, err = r.client.Post("/mgmt/actionqueue/"+data.AppDomain.ValueString(), "{\"SaveConfig\": 0}")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to save object (%s), got error: %s", "POST", err))
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SSLClientProfileResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data models.SSLClientProfile

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	res, err := r.client.Get(data.GetPath() + "/" + data.Id.ValueString())
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
	var data models.SSLClientProfile

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.client.Put(data.GetPath()+"/"+data.Id.ValueString(), data.ToBody(ctx, `SSLClientProfile`))
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to update object (PUT), got error: %s", err))
		return
	}
	_, err = r.client.Post("/mgmt/actionqueue/"+data.AppDomain.ValueString(), "{\"SaveConfig\": 0}")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to save object (%s), got error: %s", "POST", err))
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SSLClientProfileResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data models.SSLClientProfile

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.client.Delete(data.GetPath() + "/" + data.Id.ValueString())
	if err != nil && !strings.Contains(err.Error(), "status 404") && !strings.Contains(err.Error(), "status 409") {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object (DELETE), got error: %s", err))
		return
	}

	resp.State.RemoveResource(ctx)
}

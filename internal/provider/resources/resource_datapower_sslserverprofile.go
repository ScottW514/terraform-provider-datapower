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
	"github.com/scottw514/terraform-provider-datapower/client"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/modifiers"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
)

var _ resource.Resource = &SSLServerProfileResource{}

func NewSSLServerProfileResource() resource.Resource {
	return &SSLServerProfileResource{}
}

type SSLServerProfileResource struct {
	client *client.DatapowerClient
}

func (r *SSLServerProfileResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sslserverprofile"
}

func (r *SSLServerProfileResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("TLS server profile", "ssl-server", "").String,

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
					modifiers.ImmutableAfterSet(),
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
				Required:            true,
				Validators: []validator.List{
					listvalidator.ValueStringsAre(stringvalidator.OneOf("RSA_WITH_NULL_MD5", "RSA_WITH_NULL_SHA", "RSA_WITH_RC4_128_MD5", "RSA_WITH_RC4_128_SHA", "RSA_WITH_DES_CBC_SHA", "RSA_WITH_3DES_EDE_CBC_SHA", "DHE_DSS_WITH_DES_CBC_SHA", "DHE_DSS_WITH_3DES_EDE_CBC_SHA", "DHE_RSA_WITH_DES_CBC_SHA", "DHE_RSA_WITH_3DES_EDE_CBC_SHA", "RSA_WITH_AES_128_CBC_SHA", "DHE_DSS_WITH_AES_128_CBC_SHA", "DHE_RSA_WITH_AES_128_CBC_SHA", "RSA_WITH_AES_256_CBC_SHA", "DHE_DSS_WITH_AES_256_CBC_SHA", "DHE_RSA_WITH_AES_256_CBC_SHA", "RSA_WITH_NULL_SHA256", "RSA_WITH_AES_128_CBC_SHA256", "RSA_WITH_AES_256_CBC_SHA256", "DHE_DSS_WITH_AES_128_CBC_SHA256", "DHE_RSA_WITH_AES_128_CBC_SHA256", "DHE_DSS_WITH_AES_256_CBC_SHA256", "DHE_RSA_WITH_AES_256_CBC_SHA256", "RSA_WITH_AES_128_GCM_SHA256", "RSA_WITH_AES_256_GCM_SHA384", "DHE_RSA_WITH_AES_128_GCM_SHA256", "DHE_RSA_WITH_AES_256_GCM_SHA384", "DHE_DSS_WITH_AES_128_GCM_SHA256", "DHE_DSS_WITH_AES_256_GCM_SHA384", "AES_128_GCM_SHA256", "AES_256_GCM_SHA384", "CHACHA20_POLY1305_SHA256", "AES_128_CCM_SHA256", "AES_128_CCM_8_SHA256", "ECDHE_RSA_WITH_NULL_SHA", "ECDHE_RSA_WITH_RC4_128_SHA", "ECDHE_RSA_WITH_3DES_EDE_CBC_SHA", "ECDHE_RSA_WITH_AES_128_CBC_SHA", "ECDHE_RSA_WITH_AES_256_CBC_SHA", "ECDHE_RSA_WITH_AES_128_CBC_SHA256", "ECDHE_RSA_WITH_AES_256_CBC_SHA384", "ECDHE_RSA_WITH_AES_128_GCM_SHA256", "ECDHE_RSA_WITH_AES_256_GCM_SHA384", "ECDHE_ECDSA_WITH_NULL_SHA", "ECDHE_ECDSA_WITH_RC4_128_SHA", "ECDHE_ECDSA_WITH_3DES_EDE_CBC_SHA", "ECDHE_ECDSA_WITH_AES_128_CBC_SHA", "ECDHE_ECDSA_WITH_AES_256_CBC_SHA", "ECDHE_ECDSA_WITH_AES_128_CBC_SHA256", "ECDHE_ECDSA_WITH_AES_256_CBC_SHA384", "ECDHE_ECDSA_WITH_AES_128_GCM_SHA256", "ECDHE_ECDSA_WITH_AES_256_GCM_SHA384")),
				},
			},
			"idcred": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Identification credentials", "idcred", "cryptoidentcred").String,
				Required:            true,
			},
			"request_client_auth": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Request client authentication", "request-client-auth", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"require_client_auth": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Require client authentication", "require-client-auth", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"validate_client_cert": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Validate client certificate", "validate-client-cert", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"send_client_auth_ca_list": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Send client authentication CA list", "send-client-auth-ca-list", "").AddDefaultValue("true").String,
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
				MarkdownDescription: tfutils.NewAttributeDescription("Session cache size", "cache-size", "").AddIntegerRange(1, 500).AddDefaultValue("20").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(1, 500),
				},
				Default: int64default.StaticInt64(20),
			},
			"ssl_options": models.GetDmSSLOptionsResourceSchema("Advanced TLS options", "ssl-options", "", false),
			"max_ssl_duration": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Maximum TLS session duration", "max-duration", "").AddIntegerRange(1, 11520).AddDefaultValue("60").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(1, 11520),
				},
				Default: int64default.StaticInt64(60),
			},
			"disable_renegotiation": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Disable renegotiation", "disable-renegotiation", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"number_of_renegotiation_allowed": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Maximum client initiated renegotiations", "max-renegotiation-allowed", "").AddIntegerRange(0, 512).AddDefaultValue("0").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(0, 512),
				},
				Default: int64default.StaticInt64(0),
			},
			"prohibit_resume_on_reneg": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Prohibit session resumption on renegotiation", "prohibit-resume-on-reneg", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"compression": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enable compression", "compression", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"allow_legacy_renegotiation": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Allow legacy renegotiation", "allow-legacy-renegotiation", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"prefer_server_ciphers": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Use server cipher suite preferences", "prefer-server-ciphers", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"elliptic_curves": schema.ListAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Elliptic curves", "curves", "").AddStringEnum("sect163k1", "sect163r1", "sect163r2", "sect193r1", "sect193r2", "sect233k1", "sect233r1", "sect239k1", "sect283k1", "sect283r1", "sect409k1", "sect409r1", "sect571k1", "sect571r1", "secp160k1", "secp160r1", "secp160r2", "secp192k1", "secp192r1", "secp224k1", "secp224r1", "secp256k1", "secp256r1", "secp384r1", "secp521r1", "brainpoolP256r1", "brainpoolP384r1", "brainpoolP512r1").String,
				ElementType:         types.StringType,
				Optional:            true,
				Validators: []validator.List{
					listvalidator.ValueStringsAre(stringvalidator.OneOf("sect163k1", "sect163r1", "sect163r2", "sect193r1", "sect193r2", "sect233k1", "sect233r1", "sect239k1", "sect283k1", "sect283r1", "sect409k1", "sect409r1", "sect571k1", "sect571r1", "secp160k1", "secp160r1", "secp160r2", "secp192k1", "secp192r1", "secp224k1", "secp224r1", "secp256k1", "secp256r1", "secp384r1", "secp521r1", "brainpoolP256r1", "brainpoolP384r1", "brainpoolP512r1")),
				},
			},
			"prioritize_cha_cha": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Prioritize ChaCha20-Poly1305 cipher", "prioritize-chacha", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"sig_algs": schema.ListAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Signature algorithms", "sign-alg", "").AddStringEnum("ecdsa_secp256r1_sha256", "ecdsa_secp384r1_sha384", "ecdsa_secp521r1_sha512", "ed25519", "ed448", "ecdsa_sha224", "ecdsa_sha1", "rsa_pss_rsae_sha256", "rsa_pss_rsae_sha384", "rsa_pss_rsae_sha512", "rsa_pss_pss_sha256", "rsa_pss_pss_sha384", "rsa_pss_pss_sha512", "rsa_pkcs1_sha256", "rsa_pkcs1_sha384", "rsa_pkcs1_sha512", "rsa_pkcs1_sha224", "rsa_pkcs1_sha1", "dsa_sha256", "dsa_sha384", "dsa_sha512", "dsa_sha224", "dsa_sha1").String,
				ElementType:         types.StringType,
				Optional:            true,
				Validators: []validator.List{
					listvalidator.ValueStringsAre(stringvalidator.OneOf("ecdsa_secp256r1_sha256", "ecdsa_secp384r1_sha384", "ecdsa_secp521r1_sha512", "ed25519", "ed448", "ecdsa_sha224", "ecdsa_sha1", "rsa_pss_rsae_sha256", "rsa_pss_rsae_sha384", "rsa_pss_rsae_sha512", "rsa_pss_pss_sha256", "rsa_pss_pss_sha384", "rsa_pss_pss_sha512", "rsa_pkcs1_sha256", "rsa_pkcs1_sha384", "rsa_pkcs1_sha512", "rsa_pkcs1_sha224", "rsa_pkcs1_sha1", "dsa_sha256", "dsa_sha384", "dsa_sha512", "dsa_sha224", "dsa_sha1")),
				},
			},
			"require_closure_notification": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Require closure notification", "require-closure-notification", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"object_actions": actions.ActionsSchema,
		},
	}
}

func (r *SSLServerProfileResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = *req.ProviderData.(**client.DatapowerClient)
}

func (r *SSLServerProfileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data models.SSLServerProfile

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, r.client, data.ObjectActions, actions.Create)

	body := data.ToBody(ctx, `SSLServerProfile`)
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

func (r *SSLServerProfileResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data models.SSLServerProfile

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
		data.FromBody(ctx, `SSLServerProfile`, res)
	} else {
		// Update
		data.UpdateFromBody(ctx, `SSLServerProfile`, res)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SSLServerProfileResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data models.SSLServerProfile

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, r.client, data.ObjectActions, actions.Update)
	_, err := r.client.Put(data.GetPath()+"/"+data.Id.ValueString(), data.ToBody(ctx, `SSLServerProfile`))
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

func (r *SSLServerProfileResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data models.SSLServerProfile

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, r.client, data.ObjectActions, actions.Delete)
	_, err := r.client.Delete(data.GetPath() + "/" + data.Id.ValueString())
	if err != nil && !strings.Contains(err.Error(), "status 404") && !strings.Contains(err.Error(), "status 409") {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object (DELETE), got error: %s", err))
		return
	}

	resp.State.RemoveResource(ctx)
}

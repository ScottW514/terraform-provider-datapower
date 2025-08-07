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
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/scottw514/terraform-provider-datapower/client"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/modifiers"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
)

var _ resource.Resource = &B2BCPASenderSettingResource{}

func NewB2BCPASenderSettingResource() resource.Resource {
	return &B2BCPASenderSettingResource{}
}

type B2BCPASenderSettingResource struct {
	client *client.DatapowerClient
}

func (r *B2BCPASenderSettingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_b2bcpasendersetting"
}

func (r *B2BCPASenderSettingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("CPA sender setting", "cpa-sender-setting", "").String,
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
			"enabled_doc_type": models.GetDmB2BEnabledDocTypeResourceSchema("Enabled document types", "enabled-doc-type", "", false),
			"dest_endpoint_url": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Destination URL", "dest-url", "").String,
				Optional:            true,
			},
			"user_name": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Username", "username", "").String,
				Optional:            true,
			},
			"password_alias": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Password alias", "password-alias", "passwordalias").String,
				Optional:            true,
			},
			"connection_timeout": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Connection timeout", "timeout", "").AddIntegerRange(3, 7200).AddDefaultValue("300").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(3, 7200),
				},
				Default: int64default.StaticInt64(300),
			},
			"sync_reply_mode": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Sync reply mode", "syncreply-mode", "").AddStringEnum("mshSignalsOnly", "none").AddDefaultValue("none").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("mshSignalsOnly", "none"),
				},
				Default: stringdefault.StaticString("none"),
			},
			"duplicate_elimination": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Duplicate elimination", "duplicate-elimination", "").AddStringEnum("never", "always").AddDefaultValue("always").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("never", "always"),
				},
				Default: stringdefault.StaticString("always"),
			},
			"ack_requested": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Request acknowledgment", "ack-requested", "").AddStringEnum("never", "always").AddDefaultValue("never").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("never", "always"),
				},
				Default: stringdefault.StaticString("never"),
			},
			"ack_signature_requested": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Request signed acknowledgment", "ack-signature-requested", "").AddStringEnum("never", "always").AddDefaultValue("never").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("never", "always"),
				},
				Default: stringdefault.StaticString("never"),
			},
			"retry": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Retransmit unacknowledged messages", "retry", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"max_retries": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Retransmit attempts", "max-retries", "").AddIntegerRange(1, 30).AddDefaultValue("3").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(1, 30),
				},
				Default: int64default.StaticInt64(3),
			},
			"retry_interval": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Retransmit interval", "retry-interval", "").AddIntegerRange(1, 86400).AddDefaultValue("1800").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(1, 86400),
				},
				Default: int64default.StaticInt64(1800),
			},
			"persist_duration": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Persistence duration", "persist-duration", "").AddIntegerRange(0, 6000000).String,
				Optional:            true,
				Validators: []validator.Int64{

					int64validator.Between(0, 6000000),
				},
			},
			"include_time_to_live": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Include TimeToLive element", "include-time-to-live", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"encryption_required": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Require encryption", "encrypt-required", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"encrypt_cert": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Encryption certificate", "encrypt-cert", "cryptocertificate").String,
				Optional:            true,
			},
			"encrypt_algorithm": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Encryption algorithm", "encrypt-algorithm", "").AddStringEnum("http://www.w3.org/2001/04/xmlenc#tripledes-cbc", "http://www.w3.org/2001/04/xmlenc#aes128-cbc", "http://www.w3.org/2001/04/xmlenc#aes192-cbc", "http://www.w3.org/2001/04/xmlenc#aes256-cbc", "http://www.w3.org/2009/xmlenc11#aes128-gcm", "http://www.w3.org/2009/xmlenc11#aes192-gcm", "http://www.w3.org/2009/xmlenc11#aes256-gcm").AddDefaultValue("http://www.w3.org/2001/04/xmlenc#tripledes-cbc").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("http://www.w3.org/2001/04/xmlenc#tripledes-cbc", "http://www.w3.org/2001/04/xmlenc#aes128-cbc", "http://www.w3.org/2001/04/xmlenc#aes192-cbc", "http://www.w3.org/2001/04/xmlenc#aes256-cbc", "http://www.w3.org/2009/xmlenc11#aes128-gcm", "http://www.w3.org/2009/xmlenc11#aes192-gcm", "http://www.w3.org/2009/xmlenc11#aes256-gcm"),
				},
				Default: stringdefault.StaticString("http://www.w3.org/2001/04/xmlenc#tripledes-cbc"),
			},
			"signature_required": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Require signature", "sign-required", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"sign_id_cred": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Signature identification credentials", "sign-idcred", "cryptoidentcred").String,
				Optional:            true,
			},
			"signature_algorithm": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Signature algorithm", "sign-algorithm", "").AddStringEnum("dsa-sha1", "rsa-sha1", "rsa-sha256", "rsa-sha384", "rsa-sha512", "rsa-ripemd160", "rsa-ripemd160-2010", "sha256-rsa-MGF1", "rsa-md5", "ecdsa-sha1", "ecdsa-sha224", "ecdsa-sha256", "ecdsa-sha384", "ecdsa-sha512").AddDefaultValue("rsa-sha1").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("dsa-sha1", "rsa-sha1", "rsa-sha256", "rsa-sha384", "rsa-sha512", "rsa-ripemd160", "rsa-ripemd160-2010", "sha256-rsa-MGF1", "rsa-md5", "ecdsa-sha1", "ecdsa-sha224", "ecdsa-sha256", "ecdsa-sha384", "ecdsa-sha512"),
				},
				Default: stringdefault.StaticString("rsa-sha1"),
			},
			"sign_digest_algorithm": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Signing digest algorithm", "sign-digest-algorithm", "").AddStringEnum("sha1", "sha256", "sha512", "ripemd160", "sha224", "sha384", "md5").AddDefaultValue("sha1").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("sha1", "sha256", "sha512", "ripemd160", "sha224", "sha384", "md5"),
				},
				Default: stringdefault.StaticString("sha1"),
			},
			"signature_c14n_algorithm": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Signature canonicalization method", "sign-c14n-algorithm", "").AddStringEnum("c14n", "exc-c14n", "c14n-comments", "exc-c14n-comments", "c14n11", "c14n11-comments").AddDefaultValue("exc-c14n").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("c14n", "exc-c14n", "c14n-comments", "exc-c14n-comments", "c14n11", "c14n11-comments"),
				},
				Default: stringdefault.StaticString("exc-c14n"),
			},
			"ssl_client_config_type": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("TLS client type", "ssl-client-type", "").AddStringEnum("client").AddDefaultValue("client").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("client"),
				},
				Default: stringdefault.StaticString("client"),
			},
			"ssl_client": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("TLS client profile", "ssl-client", "sslclientprofile").String,
				Optional:            true,
			},
			"dependency_actions": actions.ActionsSchema,
		},
	}
}

func (r *B2BCPASenderSettingResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = *req.ProviderData.(**client.DatapowerClient)
}

func (r *B2BCPASenderSettingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data models.B2BCPASenderSetting

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, r.client, data.AppDomain.ValueString(), data.DependencyActions, actions.Create, false)

	body := data.ToBody(ctx, `B2BCPASenderSetting`)
	_, err := r.client.Post(data.GetPath(), body)

	if err != nil && !strings.Contains(err.Error(), "status 409") {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to create object (%s), got error: %s", "POST", err))
		return
	}
	actions.PostProcess(ctx, &resp.Diagnostics, r.client, data.DependencyActions, actions.Create)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *B2BCPASenderSettingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data models.B2BCPASenderSetting

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
		data.FromBody(ctx, `B2BCPASenderSetting`, res)
	} else {
		// Update
		data.UpdateFromBody(ctx, `B2BCPASenderSetting`, res)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *B2BCPASenderSettingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data models.B2BCPASenderSetting

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, r.client, data.AppDomain.ValueString(), data.DependencyActions, actions.Update, false)
	_, err := r.client.Put(data.GetPath()+"/"+data.Id.ValueString(), data.ToBody(ctx, `B2BCPASenderSetting`))
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to update object (PUT), got error: %s", err))
		return
	}

	actions.PostProcess(ctx, &resp.Diagnostics, r.client, data.DependencyActions, actions.Create)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *B2BCPASenderSettingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data models.B2BCPASenderSetting

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, r.client, data.AppDomain.ValueString(), data.DependencyActions, actions.Delete, false)
	_, err := r.client.Delete(data.GetPath() + "/" + data.Id.ValueString())
	if err != nil && !strings.Contains(err.Error(), "status 404") && !strings.Contains(err.Error(), "status 409") {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object (DELETE), got error: %s", err))
		return
	}
	actions.PostProcess(ctx, &resp.Diagnostics, r.client, data.DependencyActions, actions.Create)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.State.RemoveResource(ctx)
}

func (r *B2BCPASenderSettingResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data models.B2BCPASenderSetting

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.ValidateConfig(ctx, &resp.Diagnostics, data.DependencyActions)
}

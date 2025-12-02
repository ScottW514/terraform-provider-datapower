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
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/modifiers"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/validators"
)

var (
	_ resource.Resource                = &B2BCPASenderSettingResource{}
	_ resource.ResourceWithImportState = &B2BCPASenderSettingResource{}
)

func NewB2BCPASenderSettingResource() resource.Resource {
	return &B2BCPASenderSettingResource{}
}

type B2BCPASenderSettingResource struct {
	pData *tfutils.ProviderData
}

func (r *B2BCPASenderSettingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_b2b_cpa_sender_setting"
}

func (r *B2BCPASenderSettingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("CPA sender setting", "cpa-sender-setting", "").String,
		Attributes: map[string]schema.Attribute{
			"provider_target": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Target host for this resource. If not set, provider will use the top level settings.", "", "").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 128),
					stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9_-]+$"), "Must match :"+"^[a-zA-Z0-9_-]+$"),
				},
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
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
			"enabled_doc_type": models.GetDmB2BEnabledDocTypeResourceSchema("Enabled document types", "enabled-doc-type", "", false),
			"dest_endpoint_url": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the destination URL for sending the message to the external party endpoint. For load distribution, use the name of the load-balancing group instead of the address-port pair in the URL.", "dest-url", "").String,
				Optional:            true,
			},
			"user_name": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Username", "username", "").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(0, 128),
					stringvalidator.RegexMatches(regexp.MustCompile("^[_a-zA-Z0-9-.]+$"), "Must match :"+"^[_a-zA-Z0-9-.]+$"),
				},
			},
			"password_alias": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Password alias", "password-alias", "password_alias").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(0, 128),
				},
			},
			"connection_timeout": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the duration in seconds to maintain an idle connection. Enter a value in the range 3 - 7200. The default value is 300.", "timeout", "").AddIntegerRange(3, 7200).AddDefaultValue("300").String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(3, 7200),
				},
				Default: int64default.StaticInt64(300),
			},
			"sync_reply_mode": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Sync reply mode", "syncreply-mode", "").AddStringEnum("mshSignalsOnly", "none").AddDefaultValue("none").String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("mshSignalsOnly", "none"),
				},
				Default: stringdefault.StaticString("none"),
			},
			"duplicate_elimination": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("For an outbound ebMS message, specify whether the internal sending party requests the external receiving party to check duplicate elimination. The request is made by presenting the <tt>DuplicateElimination</tt> element in the <tt>MessageHeader</tt> element in the ebMS SOAP header. <p>When imported from CPA, the <tt>duplicateElimination</tt> attribute on the internal party <tt>DeliveryChannel</tt> element in the <tt>MessagingCharacteristics</tt> element.</p>", "duplicate-elimination", "").AddStringEnum("never", "always").AddDefaultValue("always").String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("never", "always"),
				},
				Default: stringdefault.StaticString("always"),
			},
			"ack_requested": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Request acknowledgment", "ack-requested", "").AddStringEnum("never", "always").AddDefaultValue("never").String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("never", "always"),
				},
				Default: stringdefault.StaticString("never"),
			},
			"ack_signature_requested": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Request signed acknowledgment", "ack-signature-requested", "").AddStringEnum("never", "always").AddDefaultValue("never").String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("never", "always"),
				},
				Default: stringdefault.StaticString("never"),
			},
			"retry": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Retransmit unacknowledged messages", "retry", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"max_retries": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the number of attempts to retransmit an unacknowledged message. Enter a value in the range 1 - 30. The default value is 3.", "max-retries", "").AddIntegerRange(1, 30).AddDefaultValue("3").AddRequiredWhen(models.B2BCPASenderSettingMaxRetriesCondVal.String()).AddNotValidWhen(models.B2BCPASenderSettingMaxRetriesIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 30),
					validators.ConditionalRequiredInt64(models.B2BCPASenderSettingMaxRetriesCondVal, models.B2BCPASenderSettingMaxRetriesIgnoreVal, true),
				},
				Default: int64default.StaticInt64(3),
			},
			"retry_interval": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the interval in seconds between retransmit attempts. Enter a value in the range 1 - 86400. The default value in 1800.", "retry-interval", "").AddIntegerRange(1, 86400).AddDefaultValue("1800").AddRequiredWhen(models.B2BCPASenderSettingRetryIntervalCondVal.String()).AddNotValidWhen(models.B2BCPASenderSettingRetryIntervalIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 86400),
					validators.ConditionalRequiredInt64(models.B2BCPASenderSettingRetryIntervalCondVal, models.B2BCPASenderSettingRetryIntervalIgnoreVal, true),
				},
				Default: int64default.StaticInt64(1800),
			},
			"persist_duration": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the duration in seconds to retain messages in persistent storage. This value is used to compute the <tt>TimeToLive</tt> value. Until the value of the <tt>TimeToLive</tt> element elapses, the message cannot be archived.", "persist-duration", "").AddIntegerRange(0, 6000000).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 6000000),
				},
			},
			"include_time_to_live": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to include the <tt>TimeToLive</tt> element in the outbound messages. This element indicates when the message expires. The receiving partner can accept the message only when it has not expired.", "include-time-to-live", "").AddDefaultValue("true").AddNotValidWhen(models.B2BCPASenderSettingIncludeTimeToLiveIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(true),
			},
			"encryption_required": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to encrypt outbound messages. Encryption does not apply to MSH level signals.", "encrypt-required", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"encrypt_cert": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Encryption certificate", "encrypt-cert", "crypto_certificate").AddRequiredWhen(models.B2BCPASenderSettingEncryptCertCondVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.B2BCPASenderSettingEncryptCertCondVal, validators.Evaluation{}, false),
				},
			},
			"encrypt_algorithm": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Encryption algorithm", "encrypt-algorithm", "").AddStringEnum("http://www.w3.org/2001/04/xmlenc#tripledes-cbc", "http://www.w3.org/2001/04/xmlenc#aes128-cbc", "http://www.w3.org/2001/04/xmlenc#aes192-cbc", "http://www.w3.org/2001/04/xmlenc#aes256-cbc", "http://www.w3.org/2009/xmlenc11#aes128-gcm", "http://www.w3.org/2009/xmlenc11#aes192-gcm", "http://www.w3.org/2009/xmlenc11#aes256-gcm").AddDefaultValue("http://www.w3.org/2001/04/xmlenc#tripledes-cbc").AddRequiredWhen(models.B2BCPASenderSettingEncryptAlgorithmCondVal.String()).String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("http://www.w3.org/2001/04/xmlenc#tripledes-cbc", "http://www.w3.org/2001/04/xmlenc#aes128-cbc", "http://www.w3.org/2001/04/xmlenc#aes192-cbc", "http://www.w3.org/2001/04/xmlenc#aes256-cbc", "http://www.w3.org/2009/xmlenc11#aes128-gcm", "http://www.w3.org/2009/xmlenc11#aes192-gcm", "http://www.w3.org/2009/xmlenc11#aes256-gcm"),
					validators.ConditionalRequiredString(models.B2BCPASenderSettingEncryptAlgorithmCondVal, validators.Evaluation{}, true),
				},
				Default: stringdefault.StaticString("http://www.w3.org/2001/04/xmlenc#tripledes-cbc"),
			},
			"signature_required": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Require signature", "sign-required", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"sign_id_cred": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Signature identification credentials", "sign-idcred", "crypto_ident_cred").AddRequiredWhen(models.B2BCPASenderSettingSignIdCredCondVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.B2BCPASenderSettingSignIdCredCondVal, validators.Evaluation{}, false),
				},
			},
			"signature_algorithm": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Signature algorithm", "sign-algorithm", "").AddStringEnum("dsa-sha1", "rsa-sha1", "rsa-sha256", "rsa-sha384", "rsa-sha512", "rsa-ripemd160", "rsa-ripemd160-2010", "sha256-rsa-MGF1", "rsa-md5", "ecdsa-sha1", "ecdsa-sha224", "ecdsa-sha256", "ecdsa-sha384", "ecdsa-sha512").AddDefaultValue("rsa-sha1").AddRequiredWhen(models.B2BCPASenderSettingSignatureAlgorithmCondVal.String()).String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("dsa-sha1", "rsa-sha1", "rsa-sha256", "rsa-sha384", "rsa-sha512", "rsa-ripemd160", "rsa-ripemd160-2010", "sha256-rsa-MGF1", "rsa-md5", "ecdsa-sha1", "ecdsa-sha224", "ecdsa-sha256", "ecdsa-sha384", "ecdsa-sha512"),
					validators.ConditionalRequiredString(models.B2BCPASenderSettingSignatureAlgorithmCondVal, validators.Evaluation{}, true),
				},
				Default: stringdefault.StaticString("rsa-sha1"),
			},
			"sign_digest_algorithm": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Signing digest algorithm", "sign-digest-algorithm", "").AddStringEnum("sha1", "sha256", "sha512", "ripemd160", "sha224", "sha384", "md5").AddDefaultValue("sha1").AddRequiredWhen(models.B2BCPASenderSettingSignDigestAlgorithmCondVal.String()).String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("sha1", "sha256", "sha512", "ripemd160", "sha224", "sha384", "md5"),
					validators.ConditionalRequiredString(models.B2BCPASenderSettingSignDigestAlgorithmCondVal, validators.Evaluation{}, true),
				},
				Default: stringdefault.StaticString("sha1"),
			},
			"signature_c14n_algorithm": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Signature canonicalization method", "sign-c14n-algorithm", "").AddStringEnum("c14n", "exc-c14n", "c14n-comments", "exc-c14n-comments", "c14n11", "c14n11-comments").AddDefaultValue("exc-c14n").AddRequiredWhen(models.B2BCPASenderSettingSignatureC14NAlgorithmCondVal.String()).String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("c14n", "exc-c14n", "c14n-comments", "exc-c14n-comments", "c14n11", "c14n11-comments"),
					validators.ConditionalRequiredString(models.B2BCPASenderSettingSignatureC14NAlgorithmCondVal, validators.Evaluation{}, true),
				},
				Default: stringdefault.StaticString("exc-c14n"),
			},
			"ssl_client_config_type": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("TLS client type", "ssl-client-type", "").AddStringEnum("client").AddDefaultValue("client").AddNotValidWhen(models.B2BCPASenderSettingSSLClientConfigTypeIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("client"),
					validators.ConditionalRequiredString(validators.Evaluation{}, models.B2BCPASenderSettingSSLClientConfigTypeIgnoreVal, true),
				},
				Default: stringdefault.StaticString("client"),
			},
			"ssl_client": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("TLS client profile", "ssl-client", "ssl_client_profile").AddRequiredWhen(models.B2BCPASenderSettingSSLClientCondVal.String()).AddNotValidWhen(models.B2BCPASenderSettingSSLClientIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.B2BCPASenderSettingSSLClientCondVal, models.B2BCPASenderSettingSSLClientIgnoreVal, false),
				},
			},
			"dependency_actions": actions.ActionsSchema,
		},
	}
}

func (r *B2BCPASenderSettingResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (r *B2BCPASenderSettingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.B2BCPASenderSetting
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if data.ProviderTarget.ValueString() != "" && !r.pData.Client.ValidTarget(data.ProviderTarget.ValueString()) {
		resp.Diagnostics.AddError("Invalid provider_target", fmt.Sprintf(`Target %q is not defined in the provider's 'targets' block. Available targets: %v`, data.ProviderTarget.ValueString(), r.pData.Client.GetTargetNames()))
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Create, false, data.ProviderTarget)
	if resp.Diagnostics.HasError() {
		return
	}

	body := data.ToBody(ctx, `B2BCPASenderSetting`)
	_, err := r.pData.Client.Post(data.GetPath(), body, data.ProviderTarget)
	if err != nil {
		if strings.Contains(err.Error(), "status 409") {
			_, err := r.pData.Client.Put(data.GetPath()+"/"+data.Id.ValueString(), body, data.ProviderTarget)
			if err != nil {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Resource already exists. Failed to update resource, got error: %s", err))
				return
			}
			resp.Diagnostics.AddWarning("Warning", "Resource already exists. Existing resource was updated.")
		} else if strings.Contains(err.Error(), "status 401") {
			_ = tfutils.DomainCredentialTest(r.pData.Client, &resp.Diagnostics, data.AppDomain.ValueString(), data.ProviderTarget)
			if !resp.Diagnostics.HasError() {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Application Domain '%s' does not exist", data.AppDomain.ValueString()))
			}
			return
		} else {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to create resource, got error: %s", err))
			return
		}
	}
	actions.PostProcess(ctx, &resp.Diagnostics, data.DependencyActions, actions.Create, data.ProviderTarget)
	if resp.Diagnostics.HasError() {
		return
	}

	tfutils.SaveDomain(ctx, &resp.Diagnostics, r.pData.Client, data.AppDomain, data.ProviderTarget)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *B2BCPASenderSettingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.B2BCPASenderSetting
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if data.ProviderTarget.ValueString() != "" && !r.pData.Client.ValidTarget(data.ProviderTarget.ValueString()) {
		resp.Diagnostics.AddError("Invalid provider_target", fmt.Sprintf(`Target %q is not defined in the provider's 'targets' block. Available targets: %v`, data.ProviderTarget.ValueString(), r.pData.Client.GetTargetNames()))
		return
	}
	res, err := r.pData.Client.Get(data.GetPath()+"/"+data.Id.ValueString(), data.ProviderTarget)
	if err != nil {
		if strings.Contains(err.Error(), "status 404") || strings.Contains(err.Error(), "status 406") || strings.Contains(err.Error(), "status 500") || strings.Contains(err.Error(), "status 400") {
			resp.State.RemoveResource(ctx)
			return
		} else if strings.Contains(err.Error(), "status 401") {
			_ = tfutils.DomainCredentialTest(r.pData.Client, &resp.Diagnostics, data.AppDomain.ValueString(), data.ProviderTarget)
			if !resp.Diagnostics.HasError() {
				resp.State.RemoveResource(ctx)
			}
			return
		} else {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s", err))
			return
		}
	}

	data.UpdateFromBody(ctx, `B2BCPASenderSetting`, res)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *B2BCPASenderSettingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.B2BCPASenderSetting
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if data.ProviderTarget.ValueString() != "" && !r.pData.Client.ValidTarget(data.ProviderTarget.ValueString()) {
		resp.Diagnostics.AddError("Invalid provider_target", fmt.Sprintf(`Target %q is not defined in the provider's 'targets' block. Available targets: %v`, data.ProviderTarget.ValueString(), r.pData.Client.GetTargetNames()))
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Update, false, data.ProviderTarget)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.pData.Client.Put(data.GetPath()+"/"+data.Id.ValueString(), data.ToBody(ctx, `B2BCPASenderSetting`), data.ProviderTarget)
	if err != nil {
		if strings.Contains(err.Error(), "status 401") {
			_ = tfutils.DomainCredentialTest(r.pData.Client, &resp.Diagnostics, data.AppDomain.ValueString(), data.ProviderTarget)
			if !resp.Diagnostics.HasError() {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Application Domain '%s' does not exist", data.AppDomain.ValueString()))
			}
			return
		} else {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to update object (PUT), got error: %s", err))
			return
		}
	}

	actions.PostProcess(ctx, &resp.Diagnostics, data.DependencyActions, actions.Update, data.ProviderTarget)
	if resp.Diagnostics.HasError() {
		return
	}

	tfutils.SaveDomain(ctx, &resp.Diagnostics, r.pData.Client, data.AppDomain, data.ProviderTarget)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *B2BCPASenderSettingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.B2BCPASenderSetting
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if data.ProviderTarget.ValueString() != "" && !r.pData.Client.ValidTarget(data.ProviderTarget.ValueString()) {
		resp.Diagnostics.AddError("Invalid provider_target", fmt.Sprintf(`Target %q is not defined in the provider's 'targets' block. Available targets: %v`, data.ProviderTarget.ValueString(), r.pData.Client.GetTargetNames()))
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Delete, false, data.ProviderTarget)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.pData.Client.Delete(data.GetPath()+"/"+data.Id.ValueString(), data.ProviderTarget)
	if err != nil {
		if strings.Contains(err.Error(), "status 409") {
			resp.Diagnostics.AddWarning("Resource Conflict", fmt.Sprintf("Resource is no longer tracked by Terraform, but may need to be manually deleted on DataPower host. Got error: %s", err))
		} else if strings.Contains(err.Error(), "status 401") {
			_ = tfutils.DomainCredentialTest(r.pData.Client, &resp.Diagnostics, data.AppDomain.ValueString(), data.ProviderTarget)
			if !resp.Diagnostics.HasError() {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Application Domain '%s' does not exist", data.AppDomain.ValueString()))
			}
			return
		} else if !strings.Contains(err.Error(), "status 404") {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete resource, got error: %s", err))
			return
		}
	}

	actions.PostProcess(ctx, &resp.Diagnostics, data.DependencyActions, actions.Delete, data.ProviderTarget)
	if resp.Diagnostics.HasError() {
		return
	}

	tfutils.SaveDomain(ctx, &resp.Diagnostics, r.pData.Client, data.AppDomain, data.ProviderTarget)

	resp.State.RemoveResource(ctx)
}

func (r *B2BCPASenderSettingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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

	var data models.B2BCPASenderSetting
	data.AppDomain = types.StringValue(appDomain)
	data.Id = types.StringValue(id)
	res, err := r.pData.Client.Get(data.GetPath()+"/"+data.Id.ValueString(), data.ProviderTarget)
	if err != nil {
		if strings.Contains(err.Error(), "status 404") {
			resp.Diagnostics.AddError("Resource Not Found", fmt.Sprintf("Resource was not found, got error: %s", err))
		} else if strings.Contains(err.Error(), "status 401") {
			_ = tfutils.DomainCredentialTest(r.pData.Client, &resp.Diagnostics, data.AppDomain.ValueString(), data.ProviderTarget)
			if !resp.Diagnostics.HasError() {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Application Domain '%s' does not exist", data.AppDomain.ValueString()))
			}
		} else {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		}
		return
	}

	data.FromBody(ctx, `B2BCPASenderSetting`, res)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

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
)

var _ resource.Resource = &AAAPolicyResource{}
var _ resource.ResourceWithValidateConfig = &AAAPolicyResource{}

func NewAAAPolicyResource() resource.Resource {
	return &AAAPolicyResource{}
}

type AAAPolicyResource struct {
	pData *tfutils.ProviderData
}

func (r *AAAPolicyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_aaapolicy"
}

func (r *AAAPolicyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("AAA policy", "aaapolicy", "").AddActions("flush_cache").String,
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
			"authorized_counter": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Authorized counter", "authorized-counter", "countmonitor").String,
				Optional:            true,
			},
			"rejected_counter": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Rejected counter", "rejected-counter", "countmonitor").String,
				Optional:            true,
			},
			"namespace_mapping": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Namespace mapping", "namespace-mapping", "").String,
				NestedObject:        models.DmNamespaceMappingResourceSchema,
				Optional:            true,
			},
			"extract_identity": models.GetDmAAAPExtractIdentityResourceSchema("Identity extraction", "extract-identity", "", false),
			"authenticate":     models.GetDmAAAPAuthenticateResourceSchema("Authentication", "authenticate", "", false),
			"map_credentials":  models.GetDmAAAPMapCredentialsResourceSchema("Credential mapping", "map-credentials", "", false),
			"extract_resource": models.GetDmAAAPExtractResourceResourceSchema("Resource extraction", "extract-resource", "", false),
			"map_resource":     models.GetDmAAAPMapResourceResourceSchema("Resource mapping", "map-resource", "", false),
			"authorize":        models.GetDmAAAPAuthorizeResourceSchema("Authorization", "authorize", "", false),
			"post_process":     models.GetDmAAAPPostProcessResourceSchema("Postprocessing", "post-process", "", false),
			"saml_attribute": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("SAML attributes", "saml-attribute", "").String,
				NestedObject:        models.DmSAMLAttributeNameAndValueResourceSchema,
				Optional:            true,
			},
			"ltpa_attributes": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("LTPA user attributes", "ltpa-attribute", "").String,
				NestedObject:        models.DmLTPAUserAttributeNameAndValueResourceSchema,
				Optional:            true,
			},
			"transaction_priority": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Transaction priority", "transaction-priority", "").String,
				NestedObject:        models.DmAAATransactionPriorityResourceSchema,
				Optional:            true,
			},
			"saml_valcred": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("SAML signature validation credentials", "saml-valcred", "cryptovalcred").String,
				Optional:            true,
			},
			"saml_signing_key": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("SAML message signing key", "saml-sign-key", "cryptokey").String,
				Optional:            true,
			},
			"saml_signing_cert": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("SAML message signing certificate", "saml-sign-cert", "cryptocertificate").String,
				Optional:            true,
			},
			"saml_signing_hash_alg": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("SAML signing message digest algorithm", "saml-sign-hash", "").AddStringEnum("sha1", "sha256", "sha512", "ripemd160", "sha224", "sha384", "md5").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("sha1", "sha256", "sha512", "ripemd160", "sha224", "sha384", "md5"),
				},
			},
			"saml_signing_alg": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("SAML message signing algorithm", "saml-sign-alg", "").AddStringEnum("rsa-sha1", "dsa-sha1", "rsa-sha256", "rsa-sha384", "rsa-sha512", "rsa-ripemd160", "rsa-ripemd160-2010", "sha256-rsa-MGF1", "rsa-md5", "rsa", "dsa", "ecdsa-sha1", "ecdsa-sha224", "ecdsa-sha256", "ecdsa-sha384", "ecdsa-sha512").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("rsa-sha1", "dsa-sha1", "rsa-sha256", "rsa-sha384", "rsa-sha512", "rsa-ripemd160", "rsa-ripemd160-2010", "sha256-rsa-MGF1", "rsa-md5", "rsa", "dsa", "ecdsa-sha1", "ecdsa-sha224", "ecdsa-sha256", "ecdsa-sha384", "ecdsa-sha512"),
				},
			},
			"lda_psuffix": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("LDAP suffix", "ldap-suffix", "").String,
				Optional:            true,
			},
			"log_allowed": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Log allowed", "log-allowed", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"log_allowed_level": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Log allowed level", "log-allowed-level", "").AddStringEnum("emerg", "alert", "critic", "error", "warn", "notice", "info", "debug").AddDefaultValue("info").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("emerg", "alert", "critic", "error", "warn", "notice", "info", "debug"),
				},
				Default: stringdefault.StaticString("info"),
			},
			"log_rejected": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Log rejected", "log-rejected", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"log_rejected_level": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Log rejected level", "log-rejected-level", "").AddStringEnum("emerg", "alert", "critic", "error", "warn", "notice", "info", "debug").AddDefaultValue("warn").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("emerg", "alert", "critic", "error", "warn", "notice", "info", "debug"),
				},
				Default: stringdefault.StaticString("warn"),
			},
			"ws_secure_conversation_crypto_key": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("WS-Trust encryption recipient certificate", "wstrust-encrypt-key", "cryptocertificate").String,
				Optional:            true,
			},
			"saml_source_id_mapping_file": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("SAML Artifact mapping file", "saml-artifact-mapping", "").String,
				Optional:            true,
			},
			"ping_identity_compatibility": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("PingFederate compatibility", "ping-identity-compatibility", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"saml2_metadata_file": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("SAML 2.0 metadata file", "saml2-metadata", "").String,
				Optional:            true,
			},
			"do_s_valve": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("DoS flooding attack valve", "dos-valve", "").AddIntegerRange(1, 1000).AddDefaultValue("3").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(1, 1000),
				},
				Default: int64default.StaticInt64(3),
			},
			"ldap_version": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("LDAP version", "ldap-version", "").AddStringEnum("v2", "v3").AddDefaultValue("v2").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("v2", "v3"),
				},
				Default: stringdefault.StaticString("v2"),
			},
			"enforce_soap_actor": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enforce actor or role for WS-Security message", "enforce-actor-role", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"ws_sec_actor_role_id": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("WS-Security actor or role identifier", "actor-role-id", "").String,
				Optional:            true,
			},
			"ausmhttp_header": schema.ListAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("HTTP headers", "au-sm-http-header", "").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"azsmhttp_header": schema.ListAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("HTTP headers", "az-sm-http-header", "").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"dyn_config": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Dynamic configuration type", "dyn-config", "").AddStringEnum("none", "current-aaa", "external-aaa").AddDefaultValue("none").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("none", "current-aaa", "external-aaa"),
				},
				Default: stringdefault.StaticString("none"),
			},
			"external_aaa_template": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("External AAA policy template", "external-aaa-template", "aaapolicy").String,
				Optional:            true,
			},
			"dyn_config_custom_url": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Dynamic configuration custom URL", "dyn-config-custom-url", "").String,
				Optional:            true,
			},
			"dependency_actions": actions.ActionsSchema,
		},
	}
}

func (r *AAAPolicyResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (r *AAAPolicyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.AAAPolicy
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Create, false)
	if resp.Diagnostics.HasError() {
		return
	}

	body := data.ToBody(ctx, `AAAPolicy`)
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

func (r *AAAPolicyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.AAAPolicy
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
		data.FromBody(ctx, `AAAPolicy`, res)
	} else {
		// Update
		data.UpdateFromBody(ctx, `AAAPolicy`, res)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AAAPolicyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.AAAPolicy
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Update, false)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.pData.Client.Put(data.GetPath()+"/"+data.Id.ValueString(), data.ToBody(ctx, `AAAPolicy`))
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

func (r *AAAPolicyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.AAAPolicy
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

func (r *AAAPolicyResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data models.AAAPolicy

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.ValidateConfig(ctx, &resp.Diagnostics, data.DependencyActions)
}

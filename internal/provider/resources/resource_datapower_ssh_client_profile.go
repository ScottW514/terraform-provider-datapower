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

var _ resource.Resource = &SSHClientProfileResource{}
var _ resource.ResourceWithImportState = &SSHClientProfileResource{}

func NewSSHClientProfileResource() resource.Resource {
	return &SSHClientProfileResource{}
}

type SSHClientProfileResource struct {
	pData *tfutils.ProviderData
}

func (r *SSHClientProfileResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ssh_client_profile"
}

func (r *SSHClientProfileResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("An SSH client profile defines the authentication type, credentials, and cipher suites to use for an SSH client connection.", "sshclientprofile", "").String,
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
			"user_name": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("User", "user-name", "").String,
				Required:            true,
			},
			"profile_usage": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the usage of the profile. Only SFTP is supported.", "profile-usage", "").AddStringEnum("sftp", "scc").AddDefaultValue("sftp").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("sftp", "scc"),
				},
				Default: stringdefault.StaticString("sftp"),
			},
			"ssh_user_authentication": models.GetDmSSHUserAuthenticationMethodsResourceSchema("Specify the available types of SSH user authentication for the SSH client. Authentication can be public key or password or both public key and password. When both methods are defined, public key authentication is attempted first.", "user-auth", "", false),
			"user_private_key": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the private key for public key authentication. User private keys must not be password protected.", "user-private-key", "crypto_key").AddRequiredWhen(models.SSHClientProfileUserPrivateKeyCondVal.String()).AddNotValidWhen(models.SSHClientProfileUserPrivateKeyIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.SSHClientProfileUserPrivateKeyCondVal, models.SSHClientProfileUserPrivateKeyIgnoreVal, false),
				},
			},
			"password_alias": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Password Alias", "user-password-alias", "password_alias").AddRequiredWhen(models.SSHClientProfilePasswordAliasCondVal.String()).AddNotValidWhen(models.SSHClientProfilePasswordAliasIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.SSHClientProfilePasswordAliasCondVal, models.SSHClientProfilePasswordAliasIgnoreVal, false),
				},
			},
			"persistent_connections": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to support persistent connections. By default, persistent connections are enabled. <ul><li>When enabled, new requests reuse the connection of a previous session without reauthentication.</li><li>When not enabled, new request must reauthenticate.</li></ul>", "persistent-connections", "").AddDefaultValue("true").AddRequiredWhen(models.SSHClientProfilePersistentConnectionsCondVal.String()).AddNotValidWhen(models.SSHClientProfilePersistentConnectionsIgnoreVal.String()).String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"persistent_connection_timeout": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the idle duration in seconds for a persistent connection. When the connection remains idle for the specified duration, the connection is closed. Enter any value in the range 1 - 86000. The default value is 120.", "persistent-connection-timeout", "").AddIntegerRange(1, 86400).AddDefaultValue("120").AddRequiredWhen(models.SSHClientProfilePersistentConnectionTimeoutCondVal.String()).AddNotValidWhen(models.SSHClientProfilePersistentConnectionTimeoutIgnoreVal.String()).String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 86400),
					validators.ConditionalRequiredInt64(models.SSHClientProfilePersistentConnectionTimeoutCondVal, models.SSHClientProfilePersistentConnectionTimeoutIgnoreVal, true),
				},
				Default: int64default.StaticInt64(120),
			},
			"strict_host_key_checking": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify how to check host keys during the connection and authentication phases. By default strict host key checking is not enabled. <ul><li>When enabled, checks the host key against the known hosts list. Host keys that are not in the known host list are rejected.</li><li>When not enabled, checks the host key against the known hosts list. Host keys that are not in the known host list are added to the known hosts list and accepted.</li></ul>", "strict-host-key-checking", "").AddDefaultValue("false").AddRequiredWhen(models.SSHClientProfileStrictHostKeyCheckingCondVal.String()).AddNotValidWhen(models.SSHClientProfileStrictHostKeyCheckingIgnoreVal.String()).String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"ciphers": schema.ListAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the SSH cipher suites to support.", "ciphers", "").AddStringEnum("CHACHA20-POLY1305_AT_OPENSSH.COM", "AES128-CTR", "AES192-CTR", "AES256-CTR", "AES128-GCM_AT_OPENSSH.COM", "AES256-GCM_AT_OPENSSH.COM").AddNotValidWhen(models.SSHClientProfileCiphersIgnoreVal.String()).String,
				ElementType:         types.StringType,
				Optional:            true,
				Validators: []validator.List{
					listvalidator.ValueStringsAre(
						stringvalidator.OneOf("CHACHA20-POLY1305_AT_OPENSSH.COM", "AES128-CTR", "AES192-CTR", "AES256-CTR", "AES128-GCM_AT_OPENSSH.COM", "AES256-GCM_AT_OPENSSH.COM"),
					),
					validators.ConditionalRequiredList(validators.Evaluation{}, models.SSHClientProfileCiphersIgnoreVal, false),
				},
			},
			"kex_alg": schema.ListAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the key exchange (KEX) algorithms to support.", "kex-alg", "").AddStringEnum("DIFFIE-HELLMAN-GROUP-EXCHANGE-SHA256", "ECDH-SHA2-NISTP256", "ECDH-SHA2-NISTP384", "ECDH-SHA2-NISTP521", "CURVE25519-SHA256_AT_LIBSSH.ORG").AddNotValidWhen(models.SSHClientProfileKEXAlgIgnoreVal.String()).String,
				ElementType:         types.StringType,
				Optional:            true,
				Validators: []validator.List{
					listvalidator.ValueStringsAre(
						stringvalidator.OneOf("DIFFIE-HELLMAN-GROUP-EXCHANGE-SHA256", "ECDH-SHA2-NISTP256", "ECDH-SHA2-NISTP384", "ECDH-SHA2-NISTP521", "CURVE25519-SHA256_AT_LIBSSH.ORG"),
					),
					validators.ConditionalRequiredList(validators.Evaluation{}, models.SSHClientProfileKEXAlgIgnoreVal, false),
				},
			},
			"mac_alg": schema.ListAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the message authentication codes (MAC) to support.", "mac-alg", "").AddStringEnum("HMAC-SHA1", "HMAC-SHA2-256", "HMAC-SHA2-512", "UMAC-64_AT_OPENSSH.COM", "UMAC-128_AT_OPENSSH.COM", "HMAC-SHA1-ETM_AT_OPENSSH.COM", "HMAC-SHA2-256-ETM_AT_OPENSSH.COM", "HMAC-SHA2-512-ETM_AT_OPENSSH.COM", "UMAC-64-ETM_AT_OPENSSH.COM", "UMAC-128-ETM_AT_OPENSSH.COM").AddNotValidWhen(models.SSHClientProfileMACAlgIgnoreVal.String()).String,
				ElementType:         types.StringType,
				Optional:            true,
				Validators: []validator.List{
					listvalidator.ValueStringsAre(
						stringvalidator.OneOf("HMAC-SHA1", "HMAC-SHA2-256", "HMAC-SHA2-512", "UMAC-64_AT_OPENSSH.COM", "UMAC-128_AT_OPENSSH.COM", "HMAC-SHA1-ETM_AT_OPENSSH.COM", "HMAC-SHA2-256-ETM_AT_OPENSSH.COM", "HMAC-SHA2-512-ETM_AT_OPENSSH.COM", "UMAC-64-ETM_AT_OPENSSH.COM", "UMAC-128-ETM_AT_OPENSSH.COM"),
					),
					validators.ConditionalRequiredList(validators.Evaluation{}, models.SSHClientProfileMACAlgIgnoreVal, false),
				},
			},
			"dependency_actions": actions.ActionsSchema,
		},
	}
}

func (r *SSHClientProfileResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (r *SSHClientProfileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.SSHClientProfile
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Create, false)
	if resp.Diagnostics.HasError() {
		return
	}

	body := data.ToBody(ctx, `SSHClientProfile`)
	_, err := r.pData.Client.Post(data.GetPath(), body)
	if err != nil {
		if strings.Contains(err.Error(), "status 409") {
			_, err := r.pData.Client.Put(data.GetPath()+"/"+data.Id.ValueString(), body)
			if err != nil {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Resource already exists. Failed to update resource, got error: %s", err))
				return
			}
			resp.Diagnostics.AddWarning("Warning", "Resource already exists. Existing resource was updated.")
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

func (r *SSHClientProfileResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.SSHClientProfile
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

	data.UpdateFromBody(ctx, `SSHClientProfile`, res)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SSHClientProfileResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.SSHClientProfile
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Update, false)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.pData.Client.Put(data.GetPath()+"/"+data.Id.ValueString(), data.ToBody(ctx, `SSHClientProfile`))
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

func (r *SSHClientProfileResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.SSHClientProfile
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

func (r *SSHClientProfileResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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

	var data models.SSHClientProfile
	data.AppDomain = types.StringValue(appDomain)
	data.Id = types.StringValue(id)
	res, err := r.pData.Client.Get(data.GetPath() + "/" + data.Id.ValueString())
	if err != nil {
		if strings.Contains(err.Error(), "status 404") {
			resp.Diagnostics.AddError("Resource Not Found", fmt.Sprintf("Resource was not found, got error: %s", err))
		} else {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		}
		return
	}

	data.FromBody(ctx, `SSHClientProfile`, res)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SSHClientProfileResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data models.SSHClientProfile

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.ValidateConfig(ctx, &resp.Diagnostics, data.DependencyActions)
}

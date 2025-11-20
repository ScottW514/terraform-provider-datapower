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

var _ resource.Resource = &TAMResource{}
var _ resource.ResourceWithImportState = &TAMResource{}

func NewTAMResource() resource.Resource {
	return &TAMResource{}
}

type TAMResource struct {
	pData *tfutils.ProviderData
}

func (r *TAMResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_tam"
}

func (r *TAMResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("<p>A DataPower Gateway can support multiple IBM Security Access Manager registry types: LDAP or Microsoft Active Directory. Support for the registry type is established for each Access Manager client that starts according to its configuration.</p>", "tam", "").String,
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
				MarkdownDescription: tfutils.NewAttributeDescription("Enter the descriptive summary for the configuration.", "summary", "").String,
				Optional:            true,
			},
			"ad_use_ad": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Select whether the Access Manager client uses Microsoft Active Directory instead of LDAP as the registry. The default registry for an Access Manager deployment is LDAP. This selection will cause this client to use Microsoft Active Directory. Active Directory type is not supported after ISAM 7.0 .</p><p><b>Note:</b> The type of registry that an Access Manager deployment supports is determined by the configuration of the Access Manager server. The registry that you define in this configuration is for a client and must match the registry of the server.</p>", "ad-use-ad", "").AddDefaultValue("false").AddRequiredWhen(models.TAMADUseADCondVal.String()).AddNotValidWhen(models.TAMADUseADIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"tam_version": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Select the Access Manager client version to use. The default value is Default.</p>", "tam-version", "").AddStringEnum("default", "v70", "v801", "v901", "v903", "v1005").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("default", "v70", "v801", "v901", "v903", "v1005"),
				},
			},
			"configuration_file": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Select the location of the configuration file for the Access Manager client. To be available for selection, files must have .conf or .cfg as their file extension.", "file", "").String,
				Required:            true,
			},
			"ad_configuration_file": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Select the location of the configuration file for user directories. To be available for selection, files must have .conf or .cfg as their file extension.", "reg_file", "").AddRequiredWhen(models.TAMADConfigurationFileCondVal.String()).AddNotValidWhen(models.TAMADConfigurationFileIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.TAMADConfigurationFileCondVal, models.TAMADConfigurationFileIgnoreVal, false),
				},
			},
			"ssl_key_file": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Select the location of the key file for TLS communication. To be available for selection, files must have .kdb as their file extension. Generally, these files are in the cert: directory or the sharedcert: directory.", "ssl-key", "").String,
				Required:            true,
			},
			"ssl_key_stash_file": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Select the location of the key stash file for TLS communication. To be available for selection, files must have .sth as their file extension. Generally, these files are in the cert: directory or the sharedcert: directory.", "ssl-key-stash", "").String,
				Required:            true,
			},
			"use_local_mode": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Select whether to create a local copy of the policy database. Set this property to cache the policy database locally instead of accessing the remote policy server. This property must match the behavior defined in the configuration files for the Access Manager client.", "use-local-mode", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"poll_interval": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enter the interval between requests to update the local policy database from the remote policy server. <ul><li><b>default</b> - Uses the default value, which is 600 seconds.</li><li><b>disable</b> - Disables requests to the policy database for requests.</li><li><i>seconds</i> - Specifies the time interval in seconds.</li></ul>", "cache-refresh-interval", "").AddDefaultValue("default").AddRequiredWhen(models.TAMPollIntervalCondVal.String()).AddNotValidWhen(models.TAMPollIntervalIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.TAMPollIntervalCondVal, models.TAMPollIntervalIgnoreVal, true),
				},
				Default: stringdefault.StaticString("default"),
			},
			"listen_mode": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Select whether to accept notifications to update the local policy database from the policy server. When you set this property, it overrides the behavior defined in configuration files for the Access Manager client.", "listen-mode", "").AddDefaultValue("false").AddRequiredWhen(models.TAMListenModeCondVal.String()).AddNotValidWhen(models.TAMListenModeIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"listen_port": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enter the listening port on the DataPower Gateway to receive update notifications from the remote policy server.", "listen-port", "").AddNotValidWhen(models.TAMListenPortIgnoreVal.String()).String,
				Optional:            true,
			},
			"returning_user_attributes": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Select whether the registry returns users attributes for successful authorization requests. When set, the registry returns user attributes.", "return-attributes", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"ldap_use_ssl": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Select whether to use TLS communication between the Access Manager client and the LDAP or Active Directory server.", "use-ldap-ssl", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"ldap_ssl_port": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enter the listening port that the LDAP server uses for TLS communication. This property does not apply for TLS communication with an Active Directory server.", "ldap-ssl-port", "").AddDefaultValue("636").AddNotValidWhen(models.TAMLDAPSSLPortIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Default:             int64default.StaticInt64(636),
			},
			"ldap_ssl_key_file": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Select the location of the key file that contains the certificates for TLS communication with the registry server.</p><ul><li>For server-only authentication, the key file must contain the signer certificate for the registry server.</li><li>For mutual authentication, the key file must also contain a personal certificate that the registry server can validate. If the personal certificate is not the default personal certificate in the key file, you must enter the label of the personal certificate.</li></ul><p>This file must be in the cert: or the sharedcert: directory.</p>", "ldap-ssl-key-file", "").AddRequiredWhen(models.TAMLDAPSSLKeyFileCondVal.String()).AddNotValidWhen(models.TAMLDAPSSLKeyFileIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.TAMLDAPSSLKeyFileCondVal, models.TAMLDAPSSLKeyFileIgnoreVal, false),
				},
			},
			"ldap_ssl_key_file_password_alias": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enter the password alias of the password for the key file that contains the certificates for TLS communication with the registry server.", "ldap-ssl-key-file-password-alias", "password_alias").AddNotValidWhen(models.TAMLDAPSSLKeyFilePasswordAliasIgnoreVal.String()).String,
				Optional:            true,
			},
			"ldap_ssl_key_file_label": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Enter the label of the personal certificate in the key file for client authentication.</p><ul><li>When using mutual authentication with the registry server and the personal certificate is not the default personal certificate in the key file, enter the label of the personal certificate. The personal certificate allows client authentication.</li><li>For server-only authentication, do not enter a value.</li></ul>", "ldap-ssl-key-file-dn", "").AddNotValidWhen(models.TAMLDAPSSLKeyFileLabelIgnoreVal.String()).String,
				Optional:            true,
			},
			"tam_use_fips": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Select whether the secure communication between the Access Manager client and the authorization server runs in FIPS mode.", "use-fips", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"tam_choose_nist": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("In FIPS mode, there is a mandatory NIST compliance level to select. Select a level that is compatible with the corresponding Access Manager servers and registry servers. Note that the NIST and NSA options are available only in Access Manager versions 7.0 and later.", "choose-nist", "").AddStringEnum("fips", "sp800_131_transition", "sp800_131_strict", "suite_b_128", "suite_b_192").AddNotValidWhen(models.TAMTAMChooseNISTIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("fips", "sp800_131_transition", "sp800_131_strict", "suite_b_128", "suite_b_192"),
					validators.ConditionalRequiredString(validators.Evaluation{}, models.TAMTAMChooseNISTIgnoreVal, false),
				},
			},
			"tam_use_basic_user": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Control whether to support basic users in the client. When enabled, you can use LDAP user entries for authentication or authorization without importing them into the ISAM domain.", "use-basic-user", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"user_principal_attribute": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the attribute that identifies the basic user in the LDAP user entry. The default value is uid.", "user-principal-attribute", "").AddDefaultValue("uid").AddRequiredWhen(models.TAMUserPrincipalAttributeCondVal.String()).AddNotValidWhen(models.TAMUserPrincipalAttributeIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.TAMUserPrincipalAttributeCondVal, models.TAMUserPrincipalAttributeIgnoreVal, true),
				},
				Default: stringdefault.StaticString("uid"),
			},
			"user_no_duplicates": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Control whether to allow duplicate principals.</p><ul><li>When enabled, the search for basic users covers all suffixes to ensure that no users with the same name are found. If duplicate principals are found in this client, the system returns an error.</li><li>When disabled, the search for basic users ignores possible duplicates. By default, duplicate principals are not allowed.</li></ul>", "user-no-duplicates", "").AddDefaultValue("true").AddNotValidWhen(models.TAMUserNoDuplicatesIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(true),
			},
			"user_search_suffixes": schema.ListAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Specify the ordered list of LDAP suffixes to be searched for principals. When specified and suffix optimization is disabled, the suffixes are searched in entry order. If suffix optimization is enabled, this order is overridden by the suffix optimization order.</p><p>If you do not specify any suffixes, the system searches all available suffixes.</p>", "user-search-suffixes", "").AddNotValidWhen(models.TAMUserSearchSuffixesIgnoreVal.String()).String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"user_suffix_optimiser": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Control whether to search the suffixes in an optimized order.</p><ul><li>When enabled and uplicate principals are allowed, the suffixes are searched in an optimized order based on hit count, with the most hit of the suffix at the head of the search suffix list. This can help reduce the number of suffixes searched. If duplicate principals are not allowed, the suffix optimization setting is disregarded and all suffixes are searched to check for duplicates.</li><li>When disabled, the search order is provided by the order that is defined by the search suffixes property.</li></ul>", "user-suffix-optimiser", "").AddDefaultValue("false").AddNotValidWhen(models.TAMUserSuffixOptimiserIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"tam_fed_dirs": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify a list of federated directories. Each entry describes a unique set of LDAP suffixes and LDAP server. Federated directories define all the suffixes that can be searched for user identities.", "federated-directory", "").AddNotValidWhen(models.TAMTAMFedDirsIgnoreVal.String()).String,
				NestedObject:        models.GetDmTAMFedDirResourceSchema(),
				Optional:            true,
			},
			"tam_az_replicas": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Replicas indicate the network location of remote authorization servers. You must configure at least one replica. You can configure additional replicas for failover purposes.</p><p><b>Note:</b> If you uploaded a file that was created previously, it must define at least one replica.</p>", "replica", "").String,
				NestedObject:        models.GetDmTAMAZReplicaResourceSchema(),
				Optional:            true,
			},
			"tam_ras_trace": models.GetDmTAMRASTraceResourceSchema("<p>Trace logging is a useful debugging tool. By default, trace logging is not enabled. Trace logging collects large amounts of data in a short amount of time and might result in a significant performance degradation. Enable trace logging only at the direction of IBM Support.</p><p>When enabled, the DataPower Gateway creates two trace files for each library. The DataPower Gateway writes the files cyclically. Double the size of the files to obtain the total allowable file size.</p>", "tam-ras-trace", "", false),
			"auto_retry": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Select whether to attempt starting the Access Manager client after an initial failure. The DataPower Gateway automatically attempts to start the client after a critical error. This property controls the behavior after the initial attempt to restart.</p><ul><li>When enabled, the DataPower Gateway attempts to start the client with the defined configuration.</li><li>When disabled, the client is marked as <tt>down</tt> .</li></ul><p>The default behavior is to not attempt to start the client after an initial failure.</p>", "auto-retry", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"retry_interval": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Enter the number of seconds to wait between attempts to start the Access Manager client. Enter a value in the range 1 - 65535. The default value is 180.</p>", "retry-interval", "").AddIntegerRange(1, 3600).AddDefaultValue("180").AddNotValidWhen(models.TAMRetryIntervalIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 3600),
					validators.ConditionalRequiredInt64(validators.Evaluation{}, models.TAMRetryIntervalIgnoreVal, true),
				},
				Default: int64default.StaticInt64(180),
			},
			"retry_attempts": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enter the number of attempts to perform for the Access Manager client. After performing these attempts and the client did not start, each additional attempt waits the number of seconds defined by the long interval. Enter a value in the range 0 - 65535. A value of 0 disables the long interval. The default value is 3.", "retry-attempts", "").AddDefaultValue("3").AddNotValidWhen(models.TAMRetryAttemptsIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Default:             int64default.StaticInt64(3),
			},
			"long_retry_interval": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specifies the number of seconds to wait after reaching the number of retry attempts. Enter a value in the range 1 - 65535. The default value is 900.", "long-retry-interval", "").AddIntegerRange(1, 3600).AddDefaultValue("900").AddNotValidWhen(models.TAMLongRetryIntervalIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 3600),
					validators.ConditionalRequiredInt64(validators.Evaluation{}, models.TAMLongRetryIntervalIgnoreVal, true),
				},
				Default: int64default.StaticInt64(900),
			},
			"dependency_actions": actions.ActionsSchema,
		},
	}
}

func (r *TAMResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (r *TAMResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.TAM
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Create, false)
	if resp.Diagnostics.HasError() {
		return
	}

	body := data.ToBody(ctx, `TAM`)
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

func (r *TAMResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.TAM
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

	data.UpdateFromBody(ctx, `TAM`, res)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *TAMResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.TAM
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Update, false)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.pData.Client.Put(data.GetPath()+"/"+data.Id.ValueString(), data.ToBody(ctx, `TAM`))
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

func (r *TAMResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.TAM
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

func (r *TAMResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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

	var data models.TAM
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

	data.FromBody(ctx, `TAM`, res)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *TAMResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data models.TAM

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.ValidateConfig(ctx, &resp.Diagnostics, data.DependencyActions)
}

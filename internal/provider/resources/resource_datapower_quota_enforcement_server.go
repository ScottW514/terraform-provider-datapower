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
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/validators"
)

var _ resource.Resource = &QuotaEnforcementServerResource{}
var _ resource.ResourceWithImportState = &QuotaEnforcementServerResource{}

func NewQuotaEnforcementServerResource() resource.Resource {
	return &QuotaEnforcementServerResource{}
}

type QuotaEnforcementServerResource struct {
	pData *tfutils.ProviderData
}

func (r *QuotaEnforcementServerResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_quota_enforcement_server"
}

func (r *QuotaEnforcementServerResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("<p>On each DataPower Gateway, you configure the quota enforcement server to store the thresholds and associated metadata in memory or persist them on the RAID volume.</p><p>The quota enforcement server can work in standalone mode or peer group mode.</p><p>A peer group is collection of at least three nodes across which quota enforcement is implemented. In each peer group, one DataPower Gateway is the primary node and others are replicas. Failover might occur when the primary node becomes unavailable.</p><p>When you enable the peer group mode, the appropriate configuration properties are displayed. You must configure the connection among peers.</p><p>Based on your requirements for quota enforcement, you can enable or disable strict mode. Strict mode affects data-consistency across the peer group.</p><p>When strict mode is enabled, the following effects are caused: <ul><li>When the primary node is operational and when strict mode of all nodes in a peer group is enabled, threshold synchronization is more frequent to ensure data-consistency across the peer group. However, more network bandwidth is used. Therefore, strict mode is suitable for peers in the same data center.</li><li>When the primary node becomes unavailable, before failover occurs, the replica with enabled strict mode cannot process the request. <p>If service performance and availability are more important than data-consistency, you can disable strict mode for the replica so that this replica can process the request. The replica with disabled strict mode writes the threshold and associated metadata to the local data storage. After failover occurs, the connection is resumed between the replica and the new primary node. The threshold and associated metadata stored by the replica might be overwritten by the new primary node when the new primay node synchronizes the data to all replicas. Data-consistency might be affected across the peer group.</p></li></ul></p><p>By default, data is stored in memory and quota enforcement server works in standalone mode.</p>", "quota-enforcement-server", "").String,
		Attributes: map[string]schema.Attribute{
			"enabled": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>The administrative state of the configuration.</p><ul><li>To make active, set to enabled.</li><li>To make inactive, set to disabled.</li></ul>", "admin-state", "").AddDefaultValue("true").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(true),
			},
			"user_summary": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specifies a brief descriptive summary for the configuration.", "summary", "").String,
				Optional:            true,
			},
			"password_alias": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specifies the password alias to secure the data store. If not specified, a system default is used.", "password-alias", "password_alias").String,
				Optional:            true,
			},
			"raid_volume": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Specifies whether data storage is persistent or in-memory. <ul><li>For persistent storage, select the RAID volume that must be <tt>raid0</tt> .</li><li>For in-memory storage, do not select the RAID volume.</li></ul></p>", "raid-volume", "raid_volume").String,
				Optional:            true,
			},
			"server_port": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Indicates the listening port used by the GatewayScript module to communicate with the quota enforcement server. The default port value is 16379.", "server-port", "").AddDefaultValue("16379").String,
				Computed:            true,
				Optional:            true,
				Default:             int64default.StaticInt64(16379),
			},
			"monitor_port": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Indicates the listening port for operational state monitoring for the quota enforcement server. The default port is 26379.", "monitor-port", "").AddDefaultValue("26379").String,
				Computed:            true,
				Optional:            true,
				Default:             int64default.StaticInt64(26379),
			},
			"enable_peer_group": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Indicates whether the DataPower Gateway is enabled to work in peer group mode. By default, peer group mode is disabled, which indicates that the DataPower Gateway works in standalone mode.", "enable-peer-group", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"enable_ssl": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Indicates whether TLS is used to secure connection among the peers of the peer group. By default, the TLS is enabled.", "enable-ssl", "").AddDefaultValue("true").AddNotValidWhen(models.QuotaEnforcementServerEnableSSLIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(true),
			},
			"ssl_crypto_key": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Indicates the key alias for the DataPower Gateway to authenticate a peer of the DataPower Gateway during the TLS handshake.", "ssl-key", "crypto_key").AddRequiredWhen(models.QuotaEnforcementServerSSLCryptoKeyCondVal.String()).AddNotValidWhen(models.QuotaEnforcementServerSSLCryptoKeyIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.QuotaEnforcementServerSSLCryptoKeyCondVal, models.QuotaEnforcementServerSSLCryptoKeyIgnoreVal, false),
				},
			},
			"ssl_crypto_certificate": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Indicates the certificate alias that is sent to a peer when the DataPower Gateway negotiates a TLS connection with the peer.", "ssl-cert", "crypto_certificate").AddRequiredWhen(models.QuotaEnforcementServerSSLCryptoCertificateCondVal.String()).AddNotValidWhen(models.QuotaEnforcementServerSSLCryptoCertificateIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(models.QuotaEnforcementServerSSLCryptoCertificateCondVal, models.QuotaEnforcementServerSSLCryptoCertificateIgnoreVal, false),
				},
			},
			"ip_address": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Indicates the IP address of the DataPower Gateway for other peers to connect to. The IP address can be the IP address on any interface of the DataPower Gateway and must be accessible by other peers in the peer group. The IP address cannot be 127.0.0.1, 0.0.0.0 or ::. This IP address uniquely identifies the DataPower Gateway.</p><p>You can use a local host alias instead of a static IP address. A host alias resolves a locally configured alias to a static IP address.</p>", "ip-address", "").AddRequiredWhen(models.QuotaEnforcementServerIPAddressCondVal.String()).AddNotValidWhen(models.QuotaEnforcementServerIPAddressIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.NoneOf([]string{"127.0.0.1", "0.0.0.0", "::", "::1"}...),
					validators.ConditionalRequiredString(models.QuotaEnforcementServerIPAddressCondVal, models.QuotaEnforcementServerIPAddressIgnoreVal, false),
				},
			},
			"peers": schema.ListAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Specifies peers of the DataPower Gateway in the peer group. The DataPower Gateway connects to each peer in the order in which peers are added in the list. It is not necessary to specify all peers in the Peers list.</p><ul><li>When the DataPower Gateway connects to no peer in the list, this DataPower Gateway is the first active server and joins the peer group as the primary node.</li><li>When the DataPower Gateway connects to any peer in the list, it joins the peer group as a replica.</li></ul><p>You can use a local host alias instead of a static IP address. A host alias resolves a locally configured alias to a static IP address. Aliasing can help when you move configurations among appliances.</p><p>Note: Do not specify the IP address or hostname of this DataPower Gateway.</p>", "peer", "").AddNotValidWhen(models.QuotaEnforcementServerPeersIgnoreVal.String()).String,
				ElementType:         types.StringType,
				Optional:            true,
				Validators: []validator.List{
					listvalidator.ValueStringsAre(
						stringvalidator.NoneOf([]string{"127.0.0.1", "0.0.0.0", "::", "::1"}...),
					),
					validators.ConditionalRequiredList(validators.Evaluation{}, models.QuotaEnforcementServerPeersIgnoreVal, false),
				},
			},
			"priority": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Indicates the priority that is used to decide which replica is promoted to the primary node when failover occurs.</p><p>Enter a value in range 0 - 255. The default value is 100. The replica with the lowest priority number is promoted. A replica with the value of 0 can never be promoted.</p>", "priority", "").AddIntegerRange(0, 255).AddDefaultValue("100").AddRequiredWhen(models.QuotaEnforcementServerPriorityCondVal.String()).AddNotValidWhen(models.QuotaEnforcementServerPriorityIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 255),
					validators.ConditionalRequiredInt64(models.QuotaEnforcementServerPriorityCondVal, models.QuotaEnforcementServerPriorityIgnoreVal, true),
				},
				Default: int64default.StaticInt64(100),
			},
			"strict_mode": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Based on your requirements for quota enforcement, enable or disable strict mode. By default, the strict mode is enabled.", "strict-mode", "").AddDefaultValue("true").AddNotValidWhen(models.QuotaEnforcementServerStrictModeIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(true),
			},
			"dependency_actions": actions.ActionsSchema,
		},
	}
}

func (r *QuotaEnforcementServerResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (r *QuotaEnforcementServerResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.QuotaEnforcementServer
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, "default", data.DependencyActions, actions.Create, false)
	if resp.Diagnostics.HasError() {
		return
	}

	body := data.ToBody(ctx, `QuotaEnforcementServer`)
	_, err := r.pData.Client.Put(data.GetPath(), body)
	if err != nil {
		if strings.Contains(err.Error(), "status 409") {
			_, err := r.pData.Client.Put(data.GetPath(), body)
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

func (r *QuotaEnforcementServerResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.QuotaEnforcementServer
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	res, err := r.pData.Client.Get(data.GetPath())
	if err != nil {
		if strings.Contains(err.Error(), "status 404") || strings.Contains(err.Error(), "status 406") || strings.Contains(err.Error(), "status 500") || strings.Contains(err.Error(), "status 400") {
			resp.State.RemoveResource(ctx)
			return
		} else {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s", err))
			return
		}
	}

	data.UpdateFromBody(ctx, `QuotaEnforcementServer`, res)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *QuotaEnforcementServerResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.QuotaEnforcementServer
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, "default", data.DependencyActions, actions.Update, false)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.pData.Client.Put(data.GetPath(), data.ToBody(ctx, `QuotaEnforcementServer`))
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

func (r *QuotaEnforcementServerResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.QuotaEnforcementServer
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, "default", data.DependencyActions, actions.Delete, false)
	if resp.Diagnostics.HasError() {
		return
	}
	data.ToDefault()
	_, err := r.pData.Client.Put(data.GetPath(), data.ToBody(ctx, `QuotaEnforcementServer`))
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to restore singleton to default, got error: %s", err))
		return
	}

	actions.PostProcess(ctx, &resp.Diagnostics, data.DependencyActions, actions.Delete)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.State.RemoveResource(ctx)
}

func (r *QuotaEnforcementServerResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()
	appDomain := req.ID
	if appDomain != "default" {
		resp.Diagnostics.AddError("Invalid Application Domain", "This resourece supported on the 'default' domain only.")
		return
	}
	if !regexp.MustCompile("^[a-zA-Z0-9_-]+$").MatchString(appDomain) || len(appDomain) < 1 || len(appDomain) > 128 {
		resp.Diagnostics.AddError("Invalid Application Domain", "Application domain must be 1-128 characters and match pattern ^[a-zA-Z0-9_-]+$")
		return
	}

	var data models.QuotaEnforcementServer
	res, err := r.pData.Client.Get(data.GetPath())
	if err != nil {
		if strings.Contains(err.Error(), "status 404") {
			resp.Diagnostics.AddError("Resource Not Found", fmt.Sprintf("Resource was not found, got error: %s", err))
		} else {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		}
		return
	}

	data.FromBody(ctx, `QuotaEnforcementServer`, res)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *QuotaEnforcementServerResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data models.QuotaEnforcementServer

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.ValidateConfig(ctx, &resp.Diagnostics, data.DependencyActions)
}

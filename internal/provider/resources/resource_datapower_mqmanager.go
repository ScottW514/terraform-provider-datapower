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
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/modifiers"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
)

var _ resource.Resource = &MQManagerResource{}

func NewMQManagerResource() resource.Resource {
	return &MQManagerResource{}
}

type MQManagerResource struct {
	client *client.DatapowerClient
}

func (r *MQManagerResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_mqmanager"
}

func (r *MQManagerResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("IBM MQ v9+ queue manager", "idg-mq-qm", "").String,

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
			"host_name": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Host", "hostname", "").String,
				Required:            true,
			},
			"q_mname": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Queue manager name", "queue-manager", "").String,
				Optional:            true,
			},
			"ccsid": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Coded character set ID", "ccsid", "").AddIntegerRange(0, 65535).AddDefaultValue("819").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(0, 65535),
				},
				Default: int64default.StaticInt64(819),
			},
			"channel_name": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Channel name", "channel-name", "").AddDefaultValue("SYSTEM.DEF.SVRCONN").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("SYSTEM.DEF.SVRCONN"),
			},
			"csp_user_id": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("MQCSP user ID", "mqcsp-userid", "").String,
				Optional:            true,
			},
			"csp_password_alias": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("MQCSP password alias", "mqcsp-password-alias", "passwordalias").String,
				Optional:            true,
			},
			"heartbeat": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Channel heartbeat", "heartbeat", "").AddIntegerRange(0, 999999).AddDefaultValue("300").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(0, 999999),
				},
				Default: int64default.StaticInt64(300),
			},
			"maximum_message_size": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Max message size", "maximum-message-size", "").AddIntegerRange(1024, 104857600).AddDefaultValue("1048576").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(1024, 104857600),
				},
				Default: int64default.StaticInt64(1048576),
			},
			"cache_timeout": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Cache timeout", "cache-timeout", "").AddIntegerRange(0, 65535).AddDefaultValue("60").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(0, 65535),
				},
				Default: int64default.StaticInt64(60),
			},
			"ffst_size": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("FFST file size", "ffst-size", "").AddIntegerRange(100, 50000).AddDefaultValue("500").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(100, 50000),
				},
				Default: int64default.StaticInt64(500),
			},
			"ffst_rotate": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Number of FFST file rotations", "ffst-rotate", "").AddIntegerRange(3, 5).AddDefaultValue("3").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(3, 5),
				},
				Default: int64default.StaticInt64(3),
			},
			"units_of_work": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Units-of-work", "units-of-work", "").AddIntegerRange(0, 1).AddDefaultValue("0").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(0, 1),
				},
				Default: int64default.StaticInt64(0),
			},
			"automatic_backout": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Automatic backout", "automatic-backout", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"backout_threshold": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Backout threshold", "backout-threshold", "").AddIntegerRange(1, 65535).String,
				Optional:            true,
				Validators: []validator.Int64{

					int64validator.Between(1, 65535),
				},
			},
			"backout_queue_name": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Backout queue", "backout-queue", "").String,
				Optional:            true,
			},
			"total_connection_limit": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Total connection limit", "total-connection-limit", "").AddIntegerRange(1, 10000).AddDefaultValue("250").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(1, 10000),
				},
				Default: int64default.StaticInt64(250),
			},
			"initial_connections": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Initial connections", "initial-connections", "").AddIntegerRange(0, 10000).AddDefaultValue("1").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(0, 10000),
				},
				Default: int64default.StaticInt64(1),
			},
			"sharing_conversations": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Sharing conversations", "sharing-conversations", "").AddIntegerRange(1, 5000).AddDefaultValue("1").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(1, 5000),
				},
				Default: int64default.StaticInt64(1),
			},
			"ss_lkey": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("TLS key repository", "ssl-key", "").String,
				Optional:            true,
			},
			"permit_insecure_servers": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Permit insecure connections", "permit-insecure-servers", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"ss_lcipher": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("TLS cipher specification", "ssl-cipher", "").AddStringEnum("none", "NULL_MD5", "NULL_SHA", "RC4_MD5_EXPORT", "RC4_MD5_US", "RC4_SHA_US", "RC2_MD5_EXPORT", "DES_SHA_EXPORT", "RC4_56_SHA_EXPORT1024", "DES_SHA_EXPORT1024", "TRIPLE_DES_SHA_US", "TLS_RSA_WITH_AES_128_CBC_SHA", "TLS_RSA_WITH_AES_256_CBC_SHA", "AES_SHA_US", "TLS_RSA_WITH_NULL_SHA256", "TLS_RSA_WITH_AES_128_CBC_SHA256", "TLS_RSA_WITH_AES_256_CBC_SHA256", "ECDHE_ECDSA_AES_128_CBC_SHA256", "ECDHE_RSA_AES_128_CBC_SHA256", "TLS_RSA_WITH_AES_128_GCM_SHA256", "TLS_RSA_WITH_AES_256_GCM_SHA384", "ECDHE_ECDSA_AES_256_CBC_SHA384", "ECDHE_ECDSA_AES_128_GCM_SHA256", "ECDHE_ECDSA_AES_256_GCM_SHA384", "ECDHE_RSA_AES_256_CBC_SHA384", "ECDHE_RSA_AES_128_GCM_SHA256", "ECDHE_RSA_AES_256_GCM_SHA384", "TLS_AES_128_GCM_SHA256", "TLS_AES_256_GCM_SHA384", "TLS_CHACHA20_POLY1305_SHA256", "TLS_AES_128_CCM_SHA256", "TLS_AES_128_CCM_8_SHA256").AddDefaultValue("none").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("none", "NULL_MD5", "NULL_SHA", "RC4_MD5_EXPORT", "RC4_MD5_US", "RC4_SHA_US", "RC2_MD5_EXPORT", "DES_SHA_EXPORT", "RC4_56_SHA_EXPORT1024", "DES_SHA_EXPORT1024", "TRIPLE_DES_SHA_US", "TLS_RSA_WITH_AES_128_CBC_SHA", "TLS_RSA_WITH_AES_256_CBC_SHA", "AES_SHA_US", "TLS_RSA_WITH_NULL_SHA256", "TLS_RSA_WITH_AES_128_CBC_SHA256", "TLS_RSA_WITH_AES_256_CBC_SHA256", "ECDHE_ECDSA_AES_128_CBC_SHA256", "ECDHE_RSA_AES_128_CBC_SHA256", "TLS_RSA_WITH_AES_128_GCM_SHA256", "TLS_RSA_WITH_AES_256_GCM_SHA384", "ECDHE_ECDSA_AES_256_CBC_SHA384", "ECDHE_ECDSA_AES_128_GCM_SHA256", "ECDHE_ECDSA_AES_256_GCM_SHA384", "ECDHE_RSA_AES_256_CBC_SHA384", "ECDHE_RSA_AES_128_GCM_SHA256", "ECDHE_RSA_AES_256_GCM_SHA384", "TLS_AES_128_GCM_SHA256", "TLS_AES_256_GCM_SHA384", "TLS_CHACHA20_POLY1305_SHA256", "TLS_AES_128_CCM_SHA256", "TLS_AES_128_CCM_8_SHA256"),
				},
				Default: stringdefault.StaticString("none"),
			},
			"ssl_cert_label": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("TLS certificate label", "ssl-cert-label", "").String,
				Optional:            true,
			},
			"convert_input": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Convert input", "convert", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"auto_retry": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Automatic retry", "auto-retry", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"retry_interval": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Retry interval", "retry-interval", "").AddIntegerRange(1, 65535).AddDefaultValue("10").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(1, 65535),
				},
				Default: int64default.StaticInt64(10),
			},
			"retry_attempts": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Retry attempts", "retry-attempts", "").AddIntegerRange(0, 65535).AddDefaultValue("6").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(0, 65535),
				},
				Default: int64default.StaticInt64(6),
			},
			"long_retry_interval": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Long retry interval", "long-retry-interval", "").AddIntegerRange(1, 65535).AddDefaultValue("600").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(1, 65535),
				},
				Default: int64default.StaticInt64(600),
			},
			"reporting_interval": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Reporting interval", "reporting-interval", "").AddIntegerRange(1, 65535).AddDefaultValue("10").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(1, 65535),
				},
				Default: int64default.StaticInt64(10),
			},
			"alternate_user": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Alternate user", "alternate-user", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"local_address": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Local address", "local-address", "").String,
				Optional:            true,
			},
			"xml_manager": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("XML manager", "xml-manager", "xmlmanager").AddDefaultValue("default").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("default"),
			},
			"ssl_client": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("TLS client profile", "ssl-client", "sslclientprofile").String,
				Optional:            true,
			},
			"outbound_sni": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Outbound SNI", "outbound-sni", "").AddStringEnum("Channel", "Hostname").AddDefaultValue("Channel").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("Channel", "Hostname"),
				},
				Default: stringdefault.StaticString("Channel"),
			},
			"ocsp_check_extensions": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Check OCSP extensions", "ocsp-check-extensions", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"ocsp_authentication": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("OCSP authentication", "ocsp-authentication", "").AddStringEnum("REQUIRED", "OPTIONAL", "WARN").AddDefaultValue("REQUIRED").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("REQUIRED", "OPTIONAL", "WARN"),
				},
				Default: stringdefault.StaticString("REQUIRED"),
			},
			"cdp_check_extensions": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Check CDP extensions", "cdp-check-extensions", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"client_revocation_checks": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Client revocation checking", "client-revocation-checks", "").AddStringEnum("REQUIRED", "OPTIONAL", "DISABLED").AddDefaultValue("REQUIRED").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("REQUIRED", "OPTIONAL", "DISABLED"),
				},
				Default: stringdefault.StaticString("REQUIRED"),
			},
		},
	}
}

func (r *MQManagerResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = *req.ProviderData.(**client.DatapowerClient)
}

func (r *MQManagerResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data models.MQManager

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	body := data.ToBody(ctx, `MQManager`)
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

func (r *MQManagerResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data models.MQManager

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
		data.FromBody(ctx, `MQManager`, res)
	} else {
		// Update
		data.UpdateFromBody(ctx, `MQManager`, res)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *MQManagerResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data models.MQManager

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.client.Put(data.GetPath()+"/"+data.Id.ValueString(), data.ToBody(ctx, `MQManager`))
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

func (r *MQManagerResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data models.MQManager

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

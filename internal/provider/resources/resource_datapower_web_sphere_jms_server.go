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
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/models"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/modifiers"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/validators"
)

var _ resource.Resource = &WebSphereJMSServerResource{}

func NewWebSphereJMSServerResource() resource.Resource {
	return &WebSphereJMSServerResource{}
}

type WebSphereJMSServerResource struct {
	pData *tfutils.ProviderData
}

func (r *WebSphereJMSServerResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_web_sphere_jms_server"
}

func (r *WebSphereJMSServerResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("Define the connection details to the remote WebSphere JMS server environment. This configuration is responsible for JMS messaging by periodically monitoring and polling queues to ensure that sent messages are directed to the correct receive queue or are routed to another queue manager.", "wasjms-server", "").String,
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
			"endpoint": schema.ListNestedAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Specify JMS nondefault bootstrap server endpoints by host, port, and protocol to use for the bootstrap process.</p><p>A service integration bus (SIB) supports applications that use message-based and service-oriented architectures. A bus is a group of interconnected servers and clusters that were added as members. Applications connect to a bus at one of the messaging engines that is associated with its bus members.</p><p>A messaging engine is a component that runs inside a server. The messaging engine manages messaging resources for a bus member. Applications connect to a messaging engine when a SIB is accessed.</p><p>Applications, such as the WebSphere JMS message provider, that run outside the WebSphere Application Server environment cannot directly locate a suitable messaging engine to connect to the target bus. In such cases, the remote clients or servers must access the bus through a bootstrap server that is a member of the target bus. A bootstrap server is an application server that runs the SIB process. The bootstrap server does not need to run any message engines. Rather the bootstrap server selects a messaging engine that is running in an application server that supports the bootstrap protocol that was requested.</p><p>The connection sequence to a messaging engine is as follows. <ol><li>The remote application first connects to a bootstrap server.</li><li>The bootstrap server selects a messaging engine.</li><li>The bootstrap server tells the client application to connect to that message engine to gain bus access.</li></ol></p><p>A bootstrap server uses a host name or IP address, with a port number and a bootstrap transport chain. The transport chain identifies the protocol stack that is offered by the bootstrap server to define an endpoint address.</p><p>The protocol stack that is used to access the bootstrap server does not need to be the same protocol stack that is used for actual message transfer over the bus.</p><p>You can add multiple nondefault bootstrap servers. For failover capability, the endpoints must be members of the same WebSphere Application Server cluster.</p>", "endpoint", "").String,
				NestedObject:        models.GetDmWebSphereJMSEndpointResourceSchema(),
				Required:            true,
			},
			"target_transport_chain": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Specify the predefined transport chain to use for WebSphere JMS message exchanges. The transport chain used for message exchange need not match the chain used for bootstrap access.</p><p>This property takes one of the following values. The default value is <tt>InboundBasicMessaging</tt> . <dl><dt><tt>InboundBasicMessaging</tt></dt><dd>Specifies the predefined <tt>InboundBasicMessaging</tt> transport chain (JFAP-TCP/IP).</dd><dt><tt>InboundHTTPMessaging</tt></dt><dd>Specifies the predefined <tt>InboundHTTPMessaging</tt> transport chain (tunnels JFAP using HTTP wrappers).</dd><dt><tt>InboundSecureMessaging</tt></dt><dd>Specifies the predefined <tt>InboundSecureMessaging</tt> transport chain (JFAP-SSL-TCP/IP).</dd></dl></p><p>If you have access to the WebSphere Administrative Console, you can view transport chain information through the <b>Application Servers/serverName/Transport Chain</b> menu.</p>", "target-transport-chain", "").AddDefaultValue("InboundBasicMessaging").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("InboundBasicMessaging"),
			},
			"messaging_bus": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Specify the SIB to access the remote WebSphere JMS service provider.</p><p>A SIB supports applications that use message-based and service-oriented architectures. A bus is a group of interconnected servers and clusters that were added as members of the bus. Applications connect to a bus at one of the messaging engines that is associated with its bus members.</p><p>If you have access to the WebSphere Application Server administrative console, you can view bus information through the <b>Service integration/Buses</b> menu. This information includes bus members and messaging engines, queues, topics, and the bus-specific default topic space.</p>", "messaging-bus", "").String,
				Required:            true,
			},
			"ssl_cipher": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the IBM cipher specification when the TLS profile establishes a secure connection with the WebSphere JMS message provider. When you specify a TLS profile, the cipher suite that is associated with the profile is replaced by an IBM default cipher specification. The default value is <tt>SSL_RSA_WITH_NULL_MD5</tt> .", "ssl-cipher", "").AddStringEnum("SSL_RSA_WITH_NULL_MD5", "SSL_RSA_EXPORT_WITH_RC2_CBC_40_MD5", "SSL_RSA_EXPORT_WITH_RC4_40_MD5", "SSL_RSA_WITH_RC4_128_MD5", "SSL_RSA_WITH_NULL_SHA", "SSL_RSA_EXPORT1024_WITH_RC4_56_SHA", "SSL_RSA_WITH_RC4_128_SHA", "SSL_RSA_WITH_DES_CBC_SHA", "SSL_RSA_EXPORT1024_WITH_DES_CBC_SHA", "SSL_RSA_FIPS_WITH_DES_CBC_SHA", "SSL_RSA_WITH_3DES_EDE_CBC_SHA", "SSL_RSA_FIPS_WITH_3DES_EDE_CBC_SHA", "TLS_RSA_WITH_DES_CBC_SHA", "TLS_RSA_WITH_3DES_EDE_CBC_SHA", "TLS_RSA_WITH_AES_128_CBC_SHA", "TLS_RSA_WITH_AES_256_CBC_SHA", "TLS_RSA_WITH_AES_128_CBC_SHA256", "TLS_RSA_WITH_AES_256_CBC_SHA256", "TLS_RSA_WITH_NULL_SHA256").AddNotValidWhen(models.WebSphereJMSServerSSLCipherIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("SSL_RSA_WITH_NULL_MD5", "SSL_RSA_EXPORT_WITH_RC2_CBC_40_MD5", "SSL_RSA_EXPORT_WITH_RC4_40_MD5", "SSL_RSA_WITH_RC4_128_MD5", "SSL_RSA_WITH_NULL_SHA", "SSL_RSA_EXPORT1024_WITH_RC4_56_SHA", "SSL_RSA_WITH_RC4_128_SHA", "SSL_RSA_WITH_DES_CBC_SHA", "SSL_RSA_EXPORT1024_WITH_DES_CBC_SHA", "SSL_RSA_FIPS_WITH_DES_CBC_SHA", "SSL_RSA_WITH_3DES_EDE_CBC_SHA", "SSL_RSA_FIPS_WITH_3DES_EDE_CBC_SHA", "TLS_RSA_WITH_DES_CBC_SHA", "TLS_RSA_WITH_3DES_EDE_CBC_SHA", "TLS_RSA_WITH_AES_128_CBC_SHA", "TLS_RSA_WITH_AES_256_CBC_SHA", "TLS_RSA_WITH_AES_128_CBC_SHA256", "TLS_RSA_WITH_AES_256_CBC_SHA256", "TLS_RSA_WITH_NULL_SHA256"),
					validators.ConditionalRequiredString(validators.Evaluation{}, models.WebSphereJMSServerSSLCipherIgnoreVal, false),
				},
			},
			"fips": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to require the assignment of a FIPS-compliant cipher for the WebSphere JMS message provider. By default, a FIPS-compliant cipher is not required. When enabled, which prevents the use of non-FIPs compliance ciphers, requires the use of one of the following ciphers. <ul><li><tt>TLS_RSA_WITH_AES_128_CBC_SHA</tt></li><li><tt>TLS_RSA_WITH_AES_256_CBC_SHA</tt></li><li><tt>TLS_RSA_WITH_AES_128_CBC_SHA256</tt></li><li><tt>TLS_RSA_WITH_AES_256_CBC_SHA256</tt></li></ul>", "ssl-fips", "").AddDefaultValue("false").AddNotValidWhen(models.WebSphereJMSServerFIPSIgnoreVal.String()).String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"user_summary": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Comments", "summary", "").String,
				Optional:            true,
			},
			"user_name": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Username", "username", "").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.RegexMatches(regexp.MustCompile("^[^ ]+$"), "Must match :"+"^[^ ]+$"),
				},
			},
			"password_alias": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Password alias", "password-alias", "password_alias").String,
				Optional:            true,
			},
			"transactional": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify whether to acknowledge messages only after the transaction succeeds. When enabled, messages are acknowledged only after the transaction succeeds. By default, transaction-based processing is disabled.", "transactional", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"memory_threshold": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the maximum memory allocation for pending messages in bytes. Enter a value in the range 10485760 - 1073741824. The default value is 268435456.", "memory-threshold", "").AddIntegerRange(10485760, 1073741824).AddDefaultValue("268435456").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(10485760, 1073741824),
				},
				Default: int64default.StaticInt64(268435456),
			},
			"maximum_message_size": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the maximum message size in bytes. Enter a value in the range 0 - 1073741824. The default value is 1048576. A value of 0 disables the enforcement of a maximum message size.", "maximum-message-size", "").AddIntegerRange(0, 1073741824).AddDefaultValue("1048576").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 1073741824),
				},
				Default: int64default.StaticInt64(1048576),
			},
			"default_message_type": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Specify the default JMS message type. This message type is provided only when the message type cannot be determined from the JMS message headers. By default, the message payload is accessed as a Java byte array.</p><p>On the transaction level, the <tt>DP_JMSMessageType</tt> header can change this setting.</p>", "default-message-type", "").AddStringEnum("byte", "text").AddDefaultValue("byte").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("byte", "text"),
				},
				Default: stringdefault.StaticString("byte"),
			},
			"total_connection_limit": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the maximum number of simultaneous open connections. Enter a value in the range 1 - 65535. The default value is 64.", "total-connection-limit", "").AddIntegerRange(1, 65535).AddDefaultValue("64").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 65535),
				},
				Default: int64default.StaticInt64(64),
			},
			"sessions_per_connection": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the maximum number of concurrent multiplexed sessions that a single connection can support. Enter a value in the range 1 - 65535. The default is 100. <p>Session requests in excess of the defined value trigger the establishment of a new connection to the server. A new connection cannot be established unless the number of current connections is less than the value set for total connections.</p><p>For example, with values of 20 sessions per connection, 5 total connections, and 3 active fully subscribed connections, a new session request generates the establishment of a fourth connection.</p>", "sessions-per-connection", "").AddIntegerRange(1, 65535).AddDefaultValue("100").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 65535),
				},
				Default: int64default.StaticInt64(100),
			},
			"auto_retry": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify how to control the automatic error-recovery procedure for the connection. By default, automatic error-recovery is enabled to reestablish a connection that was broken in response to an error condition.", "auto-retry", "").AddDefaultValue("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"retry_interval": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the interval between attempts to reestablish a connection in seconds. Enter a value in the range 1 - 65535. The default value is 1.", "retry-interval", "").AddIntegerRange(1, 65535).AddDefaultValue("1").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 65535),
				},
				Default: int64default.StaticInt64(1),
			},
			"enable_logging": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify how to control JMS-specific expanded logging. By default, expanded logging is disabled.", "enable-logging", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
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
				MarkdownDescription: tfutils.NewAttributeDescription("TLS client profile", "ssl-client", "ssl_client_profile").AddNotValidWhen(models.WebSphereJMSServerSSLClientIgnoreVal.String()).String,
				Optional:            true,
			},
			"dependency_actions": actions.ActionsSchema,
		},
	}
}

func (r *WebSphereJMSServerResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (r *WebSphereJMSServerResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.WebSphereJMSServer
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Create, false)
	if resp.Diagnostics.HasError() {
		return
	}

	body := data.ToBody(ctx, `WebSphereJMSServer`)
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

func (r *WebSphereJMSServerResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.WebSphereJMSServer
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
		data.FromBody(ctx, `WebSphereJMSServer`, res)
	} else {
		// Update
		data.UpdateFromBody(ctx, `WebSphereJMSServer`, res)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *WebSphereJMSServerResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.WebSphereJMSServer
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Update, false)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.pData.Client.Put(data.GetPath()+"/"+data.Id.ValueString(), data.ToBody(ctx, `WebSphereJMSServer`))
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

func (r *WebSphereJMSServerResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.WebSphereJMSServer
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

func (r *WebSphereJMSServerResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data models.WebSphereJMSServer

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.ValidateConfig(ctx, &resp.Diagnostics, data.DependencyActions)
}

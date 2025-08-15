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
)

var _ resource.Resource = &WXSGridResource{}

func NewWXSGridResource() resource.Resource {
	return &WXSGridResource{}
}

type WXSGridResource struct {
	pData *tfutils.ProviderData
}

func (r *WXSGridResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_wxs_grid"
}

func (r *WXSGridResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: tfutils.NewAttributeDescription("The configuration of an eXtreme Scale Grid defines the connection details to an eXtreme Scale grid in an eXtreme Scale collective. To define this configuration, you must define the eXtreme Scale collective, the grid name, and the user and password for the user account who connects to the eXtreme Scale. If you need to secure connections to eXtreme Scale, you must assign a TLS Proxy Profile.", "wxs-grid", "").String,
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
				MarkdownDescription: tfutils.NewAttributeDescription("Specify a brief, but descriptive, summary of the configuration.", "summary", "").String,
				Optional:            true,
			},
			"collective": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the Load Balancer Group that contains members in the collective of eXtreme Scale. You must define at least one member in the collective.", "collective", "load_balancer_group").String,
				Required:            true,
			},
			"grid": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the name of the eXtreme Scale grid. The value cannot contain whitespace or the following characters: <tt>^ . \\ / , # $ @ : ; * ? &lt; > | = + &amp; % [ ] \" '</tt> .", "grid", "").String,
				Required:            true,
			},
			"user_name": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("<p>Specify the user account of the eXtreme Scale user who connects to the eXtreme Scale collective. The value can be up to 64 characters in length and cannot be blank. You can use all alphanumeric characters and most special characters. You cannot use spaces or the following special characters: <tt># &lt;</tt> .</p><p>The user must have sufficient eXtreme Scale permissions to access the grid.</p>", "username", "").String,
				Required:            true,
			},
			"password_alias": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the password alias to use to look up the password of the eXtreme Scale user who connects to the eXtreme Scale collective.", "password-alias", "password_alias").String,
				Required:            true,
			},
			"timeout": schema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the maximum time to wait to establish a connection to an eXtreme Scale. If unable to establish a connection, the operation fails. Enter a value in the range 10 - 86400000. The default value is 1000.", "timeout", "").AddIntegerRange(10, 86400000).AddDefaultValue("1000").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.Int64{

					int64validator.Between(10, 86400000),
				},
				Default: int64default.StaticInt64(1000),
			},
			"ssl_client": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("The TLS client profile to secure connections between the DataPower Gateway and its targets.", "ssl-client", "ssl_client_profile").String,
				Optional:            true,
			},
			"encrypt": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Indicates whether the data in the eXtreme Scale data grid is encrypted. If encrypted, the data is encrypted, when writing to, and decrypted, when reading from, the eXtreme Scale data grid.", "encrypt", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"encrypt_ss_key": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the shared secret for PKCS #7 encryption and decryption. When writing data to the data grid, encrypts the data. When reading data from the eXtreme Scale data grid, decrypts the data.", "encrypt-sskey", "crypto_sskey_").String,
				Optional:            true,
			},
			"encrypt_alg": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the PKCS #7 algorithm for encryption and decryption. When writing data to the data grid, encrypts the data. When reading data from the eXtreme Scale data grid, decrypts the data.", "encrypt-alg", "").AddStringEnum("tripledes-cbc", "aes128-cbc", "aes192-cbc", "aes256-cbc", "rc2-40-cbc", "rc2-64-cbc", "rc2-cbc").AddDefaultValue("tripledes-cbc").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("tripledes-cbc", "aes128-cbc", "aes192-cbc", "aes256-cbc", "rc2-40-cbc", "rc2-64-cbc", "rc2-cbc"),
				},
				Default: stringdefault.StaticString("tripledes-cbc"),
			},
			"key_obfuscation": schema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Indicate whether to apply a hash algorithm to obfuscate keys before reading data from or writing data to the eXtreme Scale data grid.", "key-obfuscation", "").AddDefaultValue("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"key_obfuscation_alg": schema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Specify the hash algorithm to obfuscate keys before reading data from or writing data to the eXtreme Scale data grid.", "key-obfuscation-alg", "").AddStringEnum("sha1", "sha256", "sha512", "ripemd160", "sha224", "sha384", "md5").AddDefaultValue("sha256").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("sha1", "sha256", "sha512", "ripemd160", "sha224", "sha384", "md5"),
				},
				Default: stringdefault.StaticString("sha256"),
			},
			"dependency_actions": actions.ActionsSchema,
		},
	}
}

func (r *WXSGridResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.pData = req.ProviderData.(*tfutils.ProviderData)
}

func (r *WXSGridResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.WXSGrid
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Create, false)
	if resp.Diagnostics.HasError() {
		return
	}

	body := data.ToBody(ctx, `WXSGrid`)
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

func (r *WXSGridResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.WXSGrid
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
		data.FromBody(ctx, `WXSGrid`, res)
	} else {
		// Update
		data.UpdateFromBody(ctx, `WXSGrid`, res)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *WXSGridResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.WXSGrid
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.PreProcess(ctx, &resp.Diagnostics, data.AppDomain.ValueString(), data.DependencyActions, actions.Update, false)
	if resp.Diagnostics.HasError() {
		return
	}
	_, err := r.pData.Client.Put(data.GetPath()+"/"+data.Id.ValueString(), data.ToBody(ctx, `WXSGrid`))
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

func (r *WXSGridResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	r.pData.Mu.Lock()
	defer r.pData.Mu.Unlock()

	var data models.WXSGrid
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

func (r *WXSGridResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data models.WXSGrid

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	actions.ValidateConfig(ctx, &resp.Diagnostics, data.DependencyActions)
}

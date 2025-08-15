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

package models

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	DataSourceSchema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	ResourceSchema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type DmSMTPPolicy struct {
	RegExp    types.String   `tfsdk:"reg_exp"`
	Recipient types.String   `tfsdk:"recipient"`
	Sender    types.String   `tfsdk:"sender"`
	Subject   types.String   `tfsdk:"subject"`
	Options   *DmSMTPOptions `tfsdk:"options"`
	Auth      types.String   `tfsdk:"auth"`
}

var DmSMTPPolicyObjectType = map[string]attr.Type{
	"reg_exp":   types.StringType,
	"recipient": types.StringType,
	"sender":    types.StringType,
	"subject":   types.StringType,
	"options":   types.ObjectType{AttrTypes: DmSMTPOptionsObjectType},
	"auth":      types.StringType,
}
var DmSMTPPolicyObjectDefault = map[string]attr.Value{
	"reg_exp":   types.StringValue("dpsmtp://*"),
	"recipient": types.StringNull(),
	"sender":    types.StringNull(),
	"subject":   types.StringNull(),
	"options":   types.ObjectValueMust(DmSMTPOptionsObjectType, DmSMTPOptionsObjectDefault),
	"auth":      types.StringValue("plain"),
}
var DmSMTPPolicyDataSourceSchema = DataSourceSchema.NestedAttributeObject{
	Attributes: map[string]DataSourceSchema.Attribute{
		"reg_exp": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the shell-style expression to define the URL set.", "", "").AddDefaultValue("dpsmtp://*").String,
			Computed:            true,
		},
		"recipient": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the e-mail address of the recipient (\"To:\")", "", "").String,
			Computed:            true,
		},
		"sender": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the e-mail address of the sender (\"From:\")", "", "").String,
			Computed:            true,
		},
		"subject": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the subject line of the e-mail.", "", "").String,
			Computed:            true,
		},
		"options": GetDmSMTPOptionsDataSourceSchema("Specify the SMTP options to enable for the client.", "", ""),
		"auth": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the method to authenticate the SMTP client.", "", "").AddStringEnum("plain", "login").AddDefaultValue("plain").String,
			Computed:            true,
		},
	},
}
var DmSMTPPolicyResourceSchema = ResourceSchema.NestedAttributeObject{
	Attributes: map[string]ResourceSchema.Attribute{
		"reg_exp": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the shell-style expression to define the URL set.", "", "").AddDefaultValue("dpsmtp://*").String,
			Computed:            true,
			Optional:            true,
			Default:             stringdefault.StaticString("dpsmtp://*"),
		},
		"recipient": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the e-mail address of the recipient (\"To:\")", "", "").String,
			Optional:            true,
		},
		"sender": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the e-mail address of the sender (\"From:\")", "", "").String,
			Optional:            true,
		},
		"subject": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the subject line of the e-mail.", "", "").String,
			Optional:            true,
		},
		"options": GetDmSMTPOptionsResourceSchema("Specify the SMTP options to enable for the client.", "", "", false),
		"auth": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Specify the method to authenticate the SMTP client.", "", "").AddStringEnum("plain", "login").AddDefaultValue("plain").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("plain", "login"),
			},
			Default: stringdefault.StaticString("plain"),
		},
	},
}

func (data DmSMTPPolicy) IsNull() bool {
	if !data.RegExp.IsNull() {
		return false
	}
	if !data.Recipient.IsNull() {
		return false
	}
	if !data.Sender.IsNull() {
		return false
	}
	if !data.Subject.IsNull() {
		return false
	}
	if data.Options != nil {
		if !data.Options.IsNull() {
			return false
		}
	}
	if !data.Auth.IsNull() {
		return false
	}
	return true
}

func (data DmSMTPPolicy) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	if !data.RegExp.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`RegExp`, data.RegExp.ValueString())
	}
	if !data.Recipient.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Recipient`, data.Recipient.ValueString())
	}
	if !data.Sender.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Sender`, data.Sender.ValueString())
	}
	if !data.Subject.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Subject`, data.Subject.ValueString())
	}
	if data.Options != nil {
		if !data.Options.IsNull() {
			body, _ = sjson.SetRaw(body, pathRoot+`Options`, data.Options.ToBody(ctx, ""))
		}
	}
	if !data.Auth.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Auth`, data.Auth.ValueString())
	}
	return body
}

func (data *DmSMTPPolicy) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `RegExp`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.RegExp = tfutils.ParseStringFromGJSON(value)
	} else {
		data.RegExp = types.StringValue("dpsmtp://*")
	}
	if value := res.Get(pathRoot + `Recipient`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Recipient = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Recipient = types.StringNull()
	}
	if value := res.Get(pathRoot + `Sender`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Sender = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Sender = types.StringNull()
	}
	if value := res.Get(pathRoot + `Subject`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Subject = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Subject = types.StringNull()
	}
	if value := res.Get(pathRoot + `Options`); value.Exists() {
		data.Options = &DmSMTPOptions{}
		data.Options.FromBody(ctx, "", value)
	} else {
		data.Options = nil
	}
	if value := res.Get(pathRoot + `Auth`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Auth = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Auth = types.StringValue("plain")
	}
}

func (data *DmSMTPPolicy) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `RegExp`); value.Exists() && !data.RegExp.IsNull() {
		data.RegExp = tfutils.ParseStringFromGJSON(value)
	} else if data.RegExp.ValueString() != "dpsmtp://*" {
		data.RegExp = types.StringNull()
	}
	if value := res.Get(pathRoot + `Recipient`); value.Exists() && !data.Recipient.IsNull() {
		data.Recipient = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Recipient = types.StringNull()
	}
	if value := res.Get(pathRoot + `Sender`); value.Exists() && !data.Sender.IsNull() {
		data.Sender = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Sender = types.StringNull()
	}
	if value := res.Get(pathRoot + `Subject`); value.Exists() && !data.Subject.IsNull() {
		data.Subject = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Subject = types.StringNull()
	}
	if value := res.Get(pathRoot + `Options`); value.Exists() {
		data.Options.UpdateFromBody(ctx, "", value)
	} else {
		data.Options = nil
	}
	if value := res.Get(pathRoot + `Auth`); value.Exists() && !data.Auth.IsNull() {
		data.Auth = tfutils.ParseStringFromGJSON(value)
	} else if data.Auth.ValueString() != "plain" {
		data.Auth = types.StringNull()
	}
}

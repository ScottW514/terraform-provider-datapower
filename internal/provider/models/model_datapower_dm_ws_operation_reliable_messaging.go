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

type DmWSOperationReliableMessaging struct {
	Options                             *DmWSRMPolicyBitmap `tfsdk:"options"`
	DeliveryAssuranceType               types.String        `tfsdk:"delivery_assurance_type"`
	ReliableMessagingWsdlComponentType  types.String        `tfsdk:"reliable_messaging_wsdl_component_type"`
	ReliableMessagingWsdlComponentValue types.String        `tfsdk:"reliable_messaging_wsdl_component_value"`
	ReliableMessagingSubscription       types.String        `tfsdk:"reliable_messaging_subscription"`
	ReliableMessagingFragmentId         types.String        `tfsdk:"reliable_messaging_fragment_id"`
}

var DmWSOperationReliableMessagingObjectType = map[string]attr.Type{
	"options":                                 types.ObjectType{AttrTypes: DmWSRMPolicyBitmapObjectType},
	"delivery_assurance_type":                 types.StringType,
	"reliable_messaging_wsdl_component_type":  types.StringType,
	"reliable_messaging_wsdl_component_value": types.StringType,
	"reliable_messaging_subscription":         types.StringType,
	"reliable_messaging_fragment_id":          types.StringType,
}
var DmWSOperationReliableMessagingObjectDefault = map[string]attr.Value{
	"options":                                 types.ObjectValueMust(DmWSRMPolicyBitmapObjectType, DmWSRMPolicyBitmapObjectDefault),
	"delivery_assurance_type":                 types.StringValue("exactly-once"),
	"reliable_messaging_wsdl_component_type":  types.StringValue("all"),
	"reliable_messaging_wsdl_component_value": types.StringNull(),
	"reliable_messaging_subscription":         types.StringNull(),
	"reliable_messaging_fragment_id":          types.StringNull(),
}
var DmWSOperationReliableMessagingDataSourceSchema = DataSourceSchema.NestedAttributeObject{
	Attributes: map[string]DataSourceSchema.Attribute{
		"options": GetDmWSRMPolicyBitmapDataSourceSchema("Reliable Messaging Options", "", ""),
		"delivery_assurance_type": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Reliable Messaging", "", "").AddStringEnum("exactly-once").AddDefaultValue("exactly-once").String,
			Computed:            true,
		},
		"reliable_messaging_wsdl_component_type": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Select a type of WSDL Component. The default is All.", "", "").AddStringEnum("all", "subscription", "wsdl", "service", "port", "operation", "fragmentid").AddDefaultValue("all").String,
			Computed:            true,
		},
		"reliable_messaging_wsdl_component_value": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Enter the name of a WSDL-defined component of the type selected in the WSDL Component Type field.", "", "").String,
			Computed:            true,
		},
		"reliable_messaging_subscription": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Select a subscription.", "", "").String,
			Computed:            true,
		},
		"reliable_messaging_fragment_id": DataSourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Matches Fragment Identifier", "", "").String,
			Computed:            true,
		},
	},
}
var DmWSOperationReliableMessagingResourceSchema = ResourceSchema.NestedAttributeObject{
	Attributes: map[string]ResourceSchema.Attribute{
		"options": GetDmWSRMPolicyBitmapResourceSchema("Reliable Messaging Options", "", "", false),
		"delivery_assurance_type": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Reliable Messaging", "", "").AddStringEnum("exactly-once").AddDefaultValue("exactly-once").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("exactly-once"),
			},
			Default: stringdefault.StaticString("exactly-once"),
		},
		"reliable_messaging_wsdl_component_type": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Select a type of WSDL Component. The default is All.", "", "").AddStringEnum("all", "subscription", "wsdl", "service", "port", "operation", "fragmentid").AddDefaultValue("all").String,
			Computed:            true,
			Optional:            true,
			Validators: []validator.String{
				stringvalidator.OneOf("all", "subscription", "wsdl", "service", "port", "operation", "fragmentid"),
			},
			Default: stringdefault.StaticString("all"),
		},
		"reliable_messaging_wsdl_component_value": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Enter the name of a WSDL-defined component of the type selected in the WSDL Component Type field.", "", "").String,
			Optional:            true,
		},
		"reliable_messaging_subscription": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Select a subscription.", "", "").String,
			Optional:            true,
		},
		"reliable_messaging_fragment_id": ResourceSchema.StringAttribute{
			MarkdownDescription: tfutils.NewAttributeDescription("Matches Fragment Identifier", "", "").String,
			Optional:            true,
		},
	},
}

func (data DmWSOperationReliableMessaging) IsNull() bool {
	if data.Options != nil {
		if !data.Options.IsNull() {
			return false
		}
	}
	if !data.DeliveryAssuranceType.IsNull() {
		return false
	}
	if !data.ReliableMessagingWsdlComponentType.IsNull() {
		return false
	}
	if !data.ReliableMessagingWsdlComponentValue.IsNull() {
		return false
	}
	if !data.ReliableMessagingSubscription.IsNull() {
		return false
	}
	if !data.ReliableMessagingFragmentId.IsNull() {
		return false
	}
	return true
}

func (data DmWSOperationReliableMessaging) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""
	if data.Options != nil {
		if !data.Options.IsNull() {
			body, _ = sjson.SetRaw(body, pathRoot+`Options`, data.Options.ToBody(ctx, ""))
		}
	}
	if !data.DeliveryAssuranceType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DeliveryAssuranceType`, data.DeliveryAssuranceType.ValueString())
	}
	if !data.ReliableMessagingWsdlComponentType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ReliableMessagingWSDLComponentType`, data.ReliableMessagingWsdlComponentType.ValueString())
	}
	if !data.ReliableMessagingWsdlComponentValue.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ReliableMessagingWSDLComponentValue`, data.ReliableMessagingWsdlComponentValue.ValueString())
	}
	if !data.ReliableMessagingSubscription.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ReliableMessagingSubscription`, data.ReliableMessagingSubscription.ValueString())
	}
	if !data.ReliableMessagingFragmentId.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ReliableMessagingFragmentID`, data.ReliableMessagingFragmentId.ValueString())
	}
	return body
}

func (data *DmWSOperationReliableMessaging) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Options`); value.Exists() {
		data.Options = &DmWSRMPolicyBitmap{}
		data.Options.FromBody(ctx, "", value)
	} else {
		data.Options = nil
	}
	if value := res.Get(pathRoot + `DeliveryAssuranceType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.DeliveryAssuranceType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.DeliveryAssuranceType = types.StringValue("exactly-once")
	}
	if value := res.Get(pathRoot + `ReliableMessagingWSDLComponentType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ReliableMessagingWsdlComponentType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ReliableMessagingWsdlComponentType = types.StringValue("all")
	}
	if value := res.Get(pathRoot + `ReliableMessagingWSDLComponentValue`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ReliableMessagingWsdlComponentValue = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ReliableMessagingWsdlComponentValue = types.StringNull()
	}
	if value := res.Get(pathRoot + `ReliableMessagingSubscription`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ReliableMessagingSubscription = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ReliableMessagingSubscription = types.StringNull()
	}
	if value := res.Get(pathRoot + `ReliableMessagingFragmentID`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ReliableMessagingFragmentId = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ReliableMessagingFragmentId = types.StringNull()
	}
}

func (data *DmWSOperationReliableMessaging) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Options`); value.Exists() {
		data.Options.UpdateFromBody(ctx, "", value)
	} else {
		data.Options = nil
	}
	if value := res.Get(pathRoot + `DeliveryAssuranceType`); value.Exists() && !data.DeliveryAssuranceType.IsNull() {
		data.DeliveryAssuranceType = tfutils.ParseStringFromGJSON(value)
	} else if data.DeliveryAssuranceType.ValueString() != "exactly-once" {
		data.DeliveryAssuranceType = types.StringNull()
	}
	if value := res.Get(pathRoot + `ReliableMessagingWSDLComponentType`); value.Exists() && !data.ReliableMessagingWsdlComponentType.IsNull() {
		data.ReliableMessagingWsdlComponentType = tfutils.ParseStringFromGJSON(value)
	} else if data.ReliableMessagingWsdlComponentType.ValueString() != "all" {
		data.ReliableMessagingWsdlComponentType = types.StringNull()
	}
	if value := res.Get(pathRoot + `ReliableMessagingWSDLComponentValue`); value.Exists() && !data.ReliableMessagingWsdlComponentValue.IsNull() {
		data.ReliableMessagingWsdlComponentValue = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ReliableMessagingWsdlComponentValue = types.StringNull()
	}
	if value := res.Get(pathRoot + `ReliableMessagingSubscription`); value.Exists() && !data.ReliableMessagingSubscription.IsNull() {
		data.ReliableMessagingSubscription = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ReliableMessagingSubscription = types.StringNull()
	}
	if value := res.Get(pathRoot + `ReliableMessagingFragmentID`); value.Exists() && !data.ReliableMessagingFragmentId.IsNull() {
		data.ReliableMessagingFragmentId = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ReliableMessagingFragmentId = types.StringNull()
	}
}

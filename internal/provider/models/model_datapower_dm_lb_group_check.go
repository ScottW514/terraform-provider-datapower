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

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	DataSourceSchema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	ResourceSchema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/validators"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type DmLBGroupCheck struct {
	Active                          types.Bool   `tfsdk:"active"`
	Uri                             types.String `tfsdk:"uri"`
	Port                            types.Int64  `tfsdk:"port"`
	Ssl                             types.String `tfsdk:"ssl"`
	Post                            types.Bool   `tfsdk:"post"`
	Input                           types.String `tfsdk:"input"`
	Timeout                         types.Int64  `tfsdk:"timeout"`
	Frequency                       types.Int64  `tfsdk:"frequency"`
	Xpath                           types.String `tfsdk:"xpath"`
	Filter                          types.String `tfsdk:"filter"`
	EnforceTimeout                  types.Bool   `tfsdk:"enforce_timeout"`
	IndependentChecks               types.Bool   `tfsdk:"independent_checks"`
	GatewayScriptChecks             types.Bool   `tfsdk:"gateway_script_checks"`
	GatewayScriptReqMethod          types.String `tfsdk:"gateway_script_req_method"`
	GatewayScriptCustomReqMethod    types.String `tfsdk:"gateway_script_custom_req_method"`
	GatewayScriptReqDoc             types.String `tfsdk:"gateway_script_req_doc"`
	GatewayScriptReqContentType     types.String `tfsdk:"gateway_script_req_content_type"`
	GatewayScriptRspHandlerMetadata types.String `tfsdk:"gateway_script_rsp_handler_metadata"`
	GatewayScriptRspHandler         types.String `tfsdk:"gateway_script_rsp_handler"`
	TcpConnectionType               types.String `tfsdk:"tcp_connection_type"`
	SslClientConfigType             types.String `tfsdk:"ssl_client_config_type"`
	SslClient                       types.String `tfsdk:"ssl_client"`
}

var DmLBGroupCheckURIIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "ssl",
	AttrType:    "String",
	AttrDefault: "",
	Value:       []string{"LDAP", "IMSConnect", "TCPConnection"},
}

var DmLBGroupCheckPostIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "ssl",
			AttrType:    "String",
			AttrDefault: "",
			Value:       []string{"LDAP", "IMSConnect", "TCPConnection"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "gateway_script_checks",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"true"},
		},
	},
}

var DmLBGroupCheckInputIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "post",
			AttrType:    "Bool",
			AttrDefault: "true",
			Value:       []string{"false"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "gateway_script_checks",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"true"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "ssl",
			AttrType:    "String",
			AttrDefault: "",
			Value:       []string{"LDAP", "IMSConnect", "TCPConnection"},
		},
	},
}

var DmLBGroupCheckXPathIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "ssl",
			AttrType:    "String",
			AttrDefault: "",
			Value:       []string{"LDAP", "IMSConnect", "TCPConnection"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "gateway_script_checks",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"true"},
		},
	},
}

var DmLBGroupCheckFilterIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "ssl",
			AttrType:    "String",
			AttrDefault: "",
			Value:       []string{"LDAP", "IMSConnect", "TCPConnection"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "gateway_script_checks",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"true"},
		},
	},
}

var DmLBGroupCheckEnforceTimeoutIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "ssl",
	AttrType:    "String",
	AttrDefault: "",
	Value:       []string{"LDAP", "IMSConnect", "TCPConnection"},
}

var DmLBGroupCheckIndependentChecksIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "ssl",
	AttrType:    "String",
	AttrDefault: "",
	Value:       []string{"LDAP", "IMSConnect", "TCPConnection"},
}

var DmLBGroupCheckGatewayScriptChecksIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "ssl",
	AttrType:    "String",
	AttrDefault: "",
	Value:       []string{"LDAP", "IMSConnect", "TCPConnection"},
}

var DmLBGroupCheckGatewayScriptReqMethodIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "ssl",
			AttrType:    "String",
			AttrDefault: "",
			Value:       []string{"LDAP", "IMSConnect", "TCPConnection"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "gateway_script_checks",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"false"},
		},
	},
}

var DmLBGroupCheckGatewayScriptCustomReqMethodCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "ssl",
			AttrType:    "String",
			AttrDefault: "",
			Value:       []string{"Standard"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "gateway_script_checks",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"true"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "gateway_script_req_method",
			AttrType:    "String",
			AttrDefault: "",
			Value:       []string{"Custom"},
		},
	},
}

var DmLBGroupCheckGatewayScriptCustomReqMethodIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "ssl",
			AttrType:    "String",
			AttrDefault: "",
			Value:       []string{"LDAP", "IMSConnect", "TCPConnection"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "gateway_script_checks",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"false"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "gateway_script_req_method",
			AttrType:    "String",
			AttrDefault: "",
			Value:       []string{"GET", "POST", "PUT", "HEAD"},
		},
	},
}

var DmLBGroupCheckGatewayScriptReqDocCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "ssl",
			AttrType:    "String",
			AttrDefault: "",
			Value:       []string{"Standard"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "gateway_script_checks",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"true"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "gateway_script_req_method",
			AttrType:    "String",
			AttrDefault: "",
			Value:       []string{"POST", "PUT"},
		},
	},
}

var DmLBGroupCheckGatewayScriptReqDocIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "ssl",
			AttrType:    "String",
			AttrDefault: "",
			Value:       []string{"LDAP", "IMSConnect", "TCPConnection"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "gateway_script_checks",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"false"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "gateway_script_req_method",
			AttrType:    "String",
			AttrDefault: "",
			Value:       []string{"GET", "HEAD"},
		},
	},
}

var DmLBGroupCheckGatewayScriptReqContentTypeIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "ssl",
			AttrType:    "String",
			AttrDefault: "",
			Value:       []string{"LDAP", "IMSConnect", "TCPConnection"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "gateway_script_checks",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"false"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "gateway_script_req_method",
			AttrType:    "String",
			AttrDefault: "",
			Value:       []string{"GET", "HEAD"},
		},
	},
}

var DmLBGroupCheckGatewayScriptRspHandlerMetadataIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "ssl",
			AttrType:    "String",
			AttrDefault: "",
			Value:       []string{"LDAP", "IMSConnect", "TCPConnection"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "gateway_script_checks",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"false"},
		},
	},
}

var DmLBGroupCheckGatewayScriptRspHandlerCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "gateway_script_checks",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"true"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "ssl",
			AttrType:    "String",
			AttrDefault: "",
			Value:       []string{"Standard"},
		},
	},
}

var DmLBGroupCheckGatewayScriptRspHandlerIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "ssl",
			AttrType:    "String",
			AttrDefault: "",
			Value:       []string{"LDAP", "IMSConnect", "TCPConnection"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "gateway_script_checks",
			AttrType:    "Bool",
			AttrDefault: "false",
			Value:       []string{"false"},
		},
	},
}

var DmLBGroupCheckTCPConnectionTypeCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "ssl",
	AttrType:    "String",
	AttrDefault: "",
	Value:       []string{"TCPConnection"},
}

var DmLBGroupCheckTCPConnectionTypeIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "ssl",
	AttrType:    "String",
	AttrDefault: "",
	Value:       []string{"LDAP", "IMSConnect", "Standard"},
}

var DmLBGroupCheckSSLClientConfigTypeIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "ssl",
	AttrType:    "String",
	AttrDefault: "",
	Value:       []string{"Standard"},
}

var DmLBGroupCheckSSLClientIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "ssl_client_config_type",
			AttrType:    "String",
			AttrDefault: "client",
			Value:       []string{"client"},
		},
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "ssl",
			AttrType:    "String",
			AttrDefault: "",
			Value:       []string{"Standard"},
		},
	},
}

var DmLBGroupCheckObjectType = map[string]attr.Type{
	"active":                              types.BoolType,
	"uri":                                 types.StringType,
	"port":                                types.Int64Type,
	"ssl":                                 types.StringType,
	"post":                                types.BoolType,
	"input":                               types.StringType,
	"timeout":                             types.Int64Type,
	"frequency":                           types.Int64Type,
	"xpath":                               types.StringType,
	"filter":                              types.StringType,
	"enforce_timeout":                     types.BoolType,
	"independent_checks":                  types.BoolType,
	"gateway_script_checks":               types.BoolType,
	"gateway_script_req_method":           types.StringType,
	"gateway_script_custom_req_method":    types.StringType,
	"gateway_script_req_doc":              types.StringType,
	"gateway_script_req_content_type":     types.StringType,
	"gateway_script_rsp_handler_metadata": types.StringType,
	"gateway_script_rsp_handler":          types.StringType,
	"tcp_connection_type":                 types.StringType,
	"ssl_client_config_type":              types.StringType,
	"ssl_client":                          types.StringType,
}
var DmLBGroupCheckObjectDefault = map[string]attr.Value{
	"active":                              types.BoolValue(false),
	"uri":                                 types.StringValue("/"),
	"port":                                types.Int64Value(80),
	"ssl":                                 types.StringNull(),
	"post":                                types.BoolValue(true),
	"input":                               types.StringValue("store:///healthcheck.xml"),
	"timeout":                             types.Int64Value(10),
	"frequency":                           types.Int64Value(180),
	"xpath":                               types.StringValue("/"),
	"filter":                              types.StringValue("store:///healthcheck.xsl"),
	"enforce_timeout":                     types.BoolValue(false),
	"independent_checks":                  types.BoolValue(false),
	"gateway_script_checks":               types.BoolValue(false),
	"gateway_script_req_method":           types.StringNull(),
	"gateway_script_custom_req_method":    types.StringNull(),
	"gateway_script_req_doc":              types.StringValue("store:///healthcheck.json"),
	"gateway_script_req_content_type":     types.StringValue("application/json"),
	"gateway_script_rsp_handler_metadata": types.StringNull(),
	"gateway_script_rsp_handler":          types.StringValue("store:///healthcheck.js"),
	"tcp_connection_type":                 types.StringNull(),
	"ssl_client_config_type":              types.StringValue("client"),
	"ssl_client":                          types.StringNull(),
}

func GetDmLBGroupCheckDataSourceSchema(description string, cliAlias string, referenceTo string) DataSourceSchema.SingleNestedAttribute {
	var DmLBGroupCheckDataSourceSchema = DataSourceSchema.SingleNestedAttribute{
		Computed: true,
		Attributes: map[string]DataSourceSchema.Attribute{
			"active": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enables or disables a health check. The check will not run unless enabled", "admin-state", "").AddDefaultValue("false").String,
				Computed:            true,
			},
			"uri": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("The relative URI to test for each member of the group. XXX in http://www.foobar.com/XXX.", "target-uri", "").AddDefaultValue("/").AddNotValidWhen(DmLBGroupCheckURIIgnoreVal.String()).String,
				Computed:            true,
			},
			"port": DataSourceSchema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enter the TCP Port number to test. The remote port and the member server health port interact differently depending on the type of the health check. <ul><li>Standard: The member server health port is used if specified. If not specified, the remote port in the load balancer group health check is used.</li><li>LDAP: The remote port in the load balancer group health check only is used. A specification of a health port for the member server is ignored.</li><li>IMS: The health port of the member only is used. The remote port in the load balancer group health check is ignored.</li></ul>", "target-port", "").AddDefaultValue("80").String,
				Computed:            true,
			},
			"ssl": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Type of health check to perform. In TCP mode (TCP Connection health check type), a health check is performed with a TCP connection request. In HTTP mode (Standard health check type), a health check is performed with an HTTP GET or HTTP POST.", "type", "").AddStringEnum("Standard", "LDAP", "IMSConnect", "TCPConnection", "on", "off").String,
				Computed:            true,
			},
			"post": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("When set to on, the SOAP Request Document will be submitted to the web service via an HTTP POST operation. When set to off (without a request document), an HTTP Get operation is sent. The expected response is always valid XML and it is then analyzed using the Xpath Expression to verify the state of health of the server.", "use-soap", "").AddDefaultValue("true").AddNotValidWhen(DmLBGroupCheckPostIgnoreVal.String()).String,
				Computed:            true,
			},
			"input": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enter the URL of the SOAP Request Document. This may be a file on the local flash filesystem (for example, local:///healthcheck.xml). This document is POSTed to the server if Send SOAP Request is on.", "send-soap", "").AddDefaultValue("store:///healthcheck.xml").AddNotValidWhen(DmLBGroupCheckInputIgnoreVal.String()).String,
				Computed:            true,
			},
			"timeout": DataSourceSchema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enter an integer between 2 and 86400 to indicate the number of seconds to wait for a response to a health check post. The default is 10.", "timeout", "").AddIntegerRange(2, 86400).AddDefaultValue("10").String,
				Computed:            true,
			},
			"frequency": DataSourceSchema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enter an integer between 5 and 86400 to indicate the number of seconds to wait between health check posts. The default is 180.", "frequency", "").AddIntegerRange(5, 86400).AddDefaultValue("180").String,
				Computed:            true,
			},
			"xpath": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enter an XPath Expression that verifies the health of the server. The expression is applied to the server's response to the health check POST. If the expression is true, the server is healthy. Otherwise, the server health state is softdown until it passes a health check.", "xpath", "").AddDefaultValue("/").AddNotValidWhen(DmLBGroupCheckXPathIgnoreVal.String()).String,
				Computed:            true,
			},
			"filter": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enter the URL of the stylesheet that evaluates the server response using the XPath Expression entered above. The default is supplied on this DataPower device.", "filter", "").AddDefaultValue("store:///healthcheck.xsl").AddNotValidWhen(DmLBGroupCheckFilterIgnoreVal.String()).String,
				Computed:            true,
			},
			"enforce_timeout": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Indicates whether the health check timeout value is used to interrupt and end a health transaction. By default, the health check timeout only compares the actual time that the request took. With the default behavior, the health check timeout is not used to interrupt the transaction.", "enforce-timeout", "").AddDefaultValue("false").AddNotValidWhen(DmLBGroupCheckEnforceTimeoutIgnoreVal.String()).String,
				Computed:            true,
			},
			"independent_checks": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Indicates whether the health checks within a Load Balancer Group run independently of one another. By default, health checks within a Load Balancer Group run sequentially. With the default behavior, if a server hangs, all the other health checks for the other members are delayed.", "independent-checks", "").AddDefaultValue("false").AddNotValidWhen(DmLBGroupCheckIndependentChecksIgnoreVal.String()).String,
				Computed:            true,
			},
			"gateway_script_checks": DataSourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enables or disables GatewayScript health check processing. If enabled, content of health check requests and responses can be of any type and the response evaluator will be a GatewayScript file.", "gatewayscript-checks", "").AddDefaultValue("false").AddNotValidWhen(DmLBGroupCheckGatewayScriptChecksIgnoreVal.String()).String,
				Computed:            true,
			},
			"gateway_script_req_method": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("The HTTP method to use for Standard health checks with GatewayScript. If Custom is selected, a custom method name must also be specified.", "request-method", "").AddStringEnum("GET", "HEAD", "POST", "PUT", "Custom").AddNotValidWhen(DmLBGroupCheckGatewayScriptReqMethodIgnoreVal.String()).String,
				Computed:            true,
			},
			"gateway_script_custom_req_method": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("The name of the custom method to use for Standard health check requests when GatewayScript processing is enabled.", "request-custom-method", "").AddRequiredWhen(DmLBGroupCheckGatewayScriptCustomReqMethodCondVal.String()).AddNotValidWhen(DmLBGroupCheckGatewayScriptCustomReqMethodIgnoreVal.String()).String,
				Computed:            true,
			},
			"gateway_script_req_doc": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enter the URL of the Request Document. This might be a file on the local flash filesystem (for example, local:///healthcheck.xml). This document is included in the request to the server if the POST, PUT, or Custom method is selected for Standard health checks with GatewayScript. If the HTTP method is Custom, the request document is optional.", "request-doc", "").AddDefaultValue("store:///healthcheck.json").AddRequiredWhen(DmLBGroupCheckGatewayScriptReqDocCondVal.String()).AddNotValidWhen(DmLBGroupCheckGatewayScriptReqDocIgnoreVal.String()).String,
				Computed:            true,
			},
			"gateway_script_req_content_type": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enter the content type of the Request Document for Standard health checks with GatewayScript.", "request-content-type", "").AddDefaultValue("application/json").AddNotValidWhen(DmLBGroupCheckGatewayScriptReqContentTypeIgnoreVal.String()).String,
				Computed:            true,
			},
			"gateway_script_rsp_handler_metadata": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enter a string to be used by the GatewayScript Response Evaluator to help determine the state of the server's health.", "response-evaluator-metadata", "").AddNotValidWhen(DmLBGroupCheckGatewayScriptRspHandlerMetadataIgnoreVal.String()).String,
				Computed:            true,
			},
			"gateway_script_rsp_handler": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enter the URL of the GatewayScript file that evaluates the server response using the Response Evaluator Metadata. The default GatewayScript file is supplied on the DataPower Gateway.", "response-evaluator", "").AddDefaultValue("store:///healthcheck.js").AddRequiredWhen(DmLBGroupCheckGatewayScriptRspHandlerCondVal.String()).AddNotValidWhen(DmLBGroupCheckGatewayScriptRspHandlerIgnoreVal.String()).String,
				Computed:            true,
			},
			"tcp_connection_type": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("When set to full, a TCP Connection health check will perform a complete three-way handshake to establish a connection, followed by a four-way handshake to close the connection. When set to partial, a TCP Connection health check will perform a half-open connection to check a member's health. The three-way handshake to establish a connection will not be completed for partial connection types.", "tcp-conn-type", "").AddStringEnum("Full", "Partial").AddRequiredWhen(DmLBGroupCheckTCPConnectionTypeCondVal.String()).AddNotValidWhen(DmLBGroupCheckTCPConnectionTypeIgnoreVal.String()).String,
				Computed:            true,
			},
			"ssl_client_config_type": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("The TLS profile type to secure connections between the DataPower Gateway and its targets.", "ssl-client-type", "").AddStringEnum("client").AddDefaultValue("client").AddNotValidWhen(DmLBGroupCheckSSLClientConfigTypeIgnoreVal.String()).String,
				Computed:            true,
			},
			"ssl_client": DataSourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("The TLS client profile to secure connections between the DataPower Gateway and its targets.", "ssl-client", "ssl_client_profile").AddNotValidWhen(DmLBGroupCheckSSLClientIgnoreVal.String()).String,
				Computed:            true,
			},
		},
	}
	DmLBGroupCheckDataSourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, referenceTo).String
	return DmLBGroupCheckDataSourceSchema
}
func GetDmLBGroupCheckResourceSchema(description string, cliAlias string, referenceTo string, required bool) ResourceSchema.SingleNestedAttribute {
	var DmLBGroupCheckResourceSchema = ResourceSchema.SingleNestedAttribute{
		Default: objectdefault.StaticValue(
			types.ObjectValueMust(
				DmLBGroupCheckObjectType,
				DmLBGroupCheckObjectDefault,
			)),
		Attributes: map[string]ResourceSchema.Attribute{
			"active": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enables or disables a health check. The check will not run unless enabled", "admin-state", "").AddDefaultValue("false").String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"uri": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("The relative URI to test for each member of the group. XXX in http://www.foobar.com/XXX.", "target-uri", "").AddDefaultValue("/").AddNotValidWhen(DmLBGroupCheckURIIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Default:             stringdefault.StaticString("/"),
			},
			"port": ResourceSchema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enter the TCP Port number to test. The remote port and the member server health port interact differently depending on the type of the health check. <ul><li>Standard: The member server health port is used if specified. If not specified, the remote port in the load balancer group health check is used.</li><li>LDAP: The remote port in the load balancer group health check only is used. A specification of a health port for the member server is ignored.</li><li>IMS: The health port of the member only is used. The remote port in the load balancer group health check is ignored.</li></ul>", "target-port", "").AddDefaultValue("80").String,
				Computed:            true,
				Optional:            true,
				Default:             int64default.StaticInt64(80),
			},
			"ssl": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Type of health check to perform. In TCP mode (TCP Connection health check type), a health check is performed with a TCP connection request. In HTTP mode (Standard health check type), a health check is performed with an HTTP GET or HTTP POST.", "type", "").AddStringEnum("Standard", "LDAP", "IMSConnect", "TCPConnection", "on", "off").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("Standard", "LDAP", "IMSConnect", "TCPConnection", "on", "off"),
				},
			},
			"post": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("When set to on, the SOAP Request Document will be submitted to the web service via an HTTP POST operation. When set to off (without a request document), an HTTP Get operation is sent. The expected response is always valid XML and it is then analyzed using the Xpath Expression to verify the state of health of the server.", "use-soap", "").AddDefaultValue("true").AddNotValidWhen(DmLBGroupCheckPostIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(true),
			},
			"input": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enter the URL of the SOAP Request Document. This may be a file on the local flash filesystem (for example, local:///healthcheck.xml). This document is POSTed to the server if Send SOAP Request is on.", "send-soap", "").AddDefaultValue("store:///healthcheck.xml").AddNotValidWhen(DmLBGroupCheckInputIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Default:             stringdefault.StaticString("store:///healthcheck.xml"),
			},
			"timeout": ResourceSchema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enter an integer between 2 and 86400 to indicate the number of seconds to wait for a response to a health check post. The default is 10.", "timeout", "").AddIntegerRange(2, 86400).AddDefaultValue("10").String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(2, 86400),
				},
				Default: int64default.StaticInt64(10),
			},
			"frequency": ResourceSchema.Int64Attribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enter an integer between 5 and 86400 to indicate the number of seconds to wait between health check posts. The default is 180.", "frequency", "").AddIntegerRange(5, 86400).AddDefaultValue("180").String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(5, 86400),
				},
				Default: int64default.StaticInt64(180),
			},
			"xpath": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enter an XPath Expression that verifies the health of the server. The expression is applied to the server's response to the health check POST. If the expression is true, the server is healthy. Otherwise, the server health state is softdown until it passes a health check.", "xpath", "").AddDefaultValue("/").AddNotValidWhen(DmLBGroupCheckXPathIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Default:             stringdefault.StaticString("/"),
			},
			"filter": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enter the URL of the stylesheet that evaluates the server response using the XPath Expression entered above. The default is supplied on this DataPower device.", "filter", "").AddDefaultValue("store:///healthcheck.xsl").AddNotValidWhen(DmLBGroupCheckFilterIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Default:             stringdefault.StaticString("store:///healthcheck.xsl"),
			},
			"enforce_timeout": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Indicates whether the health check timeout value is used to interrupt and end a health transaction. By default, the health check timeout only compares the actual time that the request took. With the default behavior, the health check timeout is not used to interrupt the transaction.", "enforce-timeout", "").AddDefaultValue("false").AddNotValidWhen(DmLBGroupCheckEnforceTimeoutIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"independent_checks": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Indicates whether the health checks within a Load Balancer Group run independently of one another. By default, health checks within a Load Balancer Group run sequentially. With the default behavior, if a server hangs, all the other health checks for the other members are delayed.", "independent-checks", "").AddDefaultValue("false").AddNotValidWhen(DmLBGroupCheckIndependentChecksIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"gateway_script_checks": ResourceSchema.BoolAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enables or disables GatewayScript health check processing. If enabled, content of health check requests and responses can be of any type and the response evaluator will be a GatewayScript file.", "gatewayscript-checks", "").AddDefaultValue("false").AddNotValidWhen(DmLBGroupCheckGatewayScriptChecksIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
			},
			"gateway_script_req_method": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("The HTTP method to use for Standard health checks with GatewayScript. If Custom is selected, a custom method name must also be specified.", "request-method", "").AddStringEnum("GET", "HEAD", "POST", "PUT", "Custom").AddNotValidWhen(DmLBGroupCheckGatewayScriptReqMethodIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("GET", "HEAD", "POST", "PUT", "Custom"),
					validators.ConditionalRequiredString(validators.Evaluation{}, DmLBGroupCheckGatewayScriptReqMethodIgnoreVal, false),
				},
			},
			"gateway_script_custom_req_method": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("The name of the custom method to use for Standard health check requests when GatewayScript processing is enabled.", "request-custom-method", "").AddRequiredWhen(DmLBGroupCheckGatewayScriptCustomReqMethodCondVal.String()).AddNotValidWhen(DmLBGroupCheckGatewayScriptCustomReqMethodIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(DmLBGroupCheckGatewayScriptCustomReqMethodCondVal, DmLBGroupCheckGatewayScriptCustomReqMethodIgnoreVal, false),
				},
			},
			"gateway_script_req_doc": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enter the URL of the Request Document. This might be a file on the local flash filesystem (for example, local:///healthcheck.xml). This document is included in the request to the server if the POST, PUT, or Custom method is selected for Standard health checks with GatewayScript. If the HTTP method is Custom, the request document is optional.", "request-doc", "").AddDefaultValue("store:///healthcheck.json").AddRequiredWhen(DmLBGroupCheckGatewayScriptReqDocCondVal.String()).AddNotValidWhen(DmLBGroupCheckGatewayScriptReqDocIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(DmLBGroupCheckGatewayScriptReqDocCondVal, DmLBGroupCheckGatewayScriptReqDocIgnoreVal, true),
				},
				Default: stringdefault.StaticString("store:///healthcheck.json"),
			},
			"gateway_script_req_content_type": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enter the content type of the Request Document for Standard health checks with GatewayScript.", "request-content-type", "").AddDefaultValue("application/json").AddNotValidWhen(DmLBGroupCheckGatewayScriptReqContentTypeIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Default:             stringdefault.StaticString("application/json"),
			},
			"gateway_script_rsp_handler_metadata": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enter a string to be used by the GatewayScript Response Evaluator to help determine the state of the server's health.", "response-evaluator-metadata", "").AddNotValidWhen(DmLBGroupCheckGatewayScriptRspHandlerMetadataIgnoreVal.String()).String,
				Optional:            true,
			},
			"gateway_script_rsp_handler": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("Enter the URL of the GatewayScript file that evaluates the server response using the Response Evaluator Metadata. The default GatewayScript file is supplied on the DataPower Gateway.", "response-evaluator", "").AddDefaultValue("store:///healthcheck.js").AddRequiredWhen(DmLBGroupCheckGatewayScriptRspHandlerCondVal.String()).AddNotValidWhen(DmLBGroupCheckGatewayScriptRspHandlerIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					validators.ConditionalRequiredString(DmLBGroupCheckGatewayScriptRspHandlerCondVal, DmLBGroupCheckGatewayScriptRspHandlerIgnoreVal, true),
				},
				Default: stringdefault.StaticString("store:///healthcheck.js"),
			},
			"tcp_connection_type": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("When set to full, a TCP Connection health check will perform a complete three-way handshake to establish a connection, followed by a four-way handshake to close the connection. When set to partial, a TCP Connection health check will perform a half-open connection to check a member's health. The three-way handshake to establish a connection will not be completed for partial connection types.", "tcp-conn-type", "").AddStringEnum("Full", "Partial").AddRequiredWhen(DmLBGroupCheckTCPConnectionTypeCondVal.String()).AddNotValidWhen(DmLBGroupCheckTCPConnectionTypeIgnoreVal.String()).String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("Full", "Partial"),
					validators.ConditionalRequiredString(DmLBGroupCheckTCPConnectionTypeCondVal, DmLBGroupCheckTCPConnectionTypeIgnoreVal, false),
				},
			},
			"ssl_client_config_type": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("The TLS profile type to secure connections between the DataPower Gateway and its targets.", "ssl-client-type", "").AddStringEnum("client").AddDefaultValue("client").AddNotValidWhen(DmLBGroupCheckSSLClientConfigTypeIgnoreVal.String()).String,
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("client"),
					validators.ConditionalRequiredString(validators.Evaluation{}, DmLBGroupCheckSSLClientConfigTypeIgnoreVal, true),
				},
				Default: stringdefault.StaticString("client"),
			},
			"ssl_client": ResourceSchema.StringAttribute{
				MarkdownDescription: tfutils.NewAttributeDescription("The TLS client profile to secure connections between the DataPower Gateway and its targets.", "ssl-client", "ssl_client_profile").AddNotValidWhen(DmLBGroupCheckSSLClientIgnoreVal.String()).String,
				Optional:            true,
			},
		},
	}
	DmLBGroupCheckResourceSchema.MarkdownDescription = tfutils.NewAttributeDescription(description, cliAlias, referenceTo).String
	if required {
		DmLBGroupCheckResourceSchema.Required = true
	} else {
		DmLBGroupCheckResourceSchema.Optional = true
		DmLBGroupCheckResourceSchema.Computed = true
	}
	return DmLBGroupCheckResourceSchema
}

func (data DmLBGroupCheck) IsNull() bool {
	if !data.Active.IsNull() {
		return false
	}
	if !data.Uri.IsNull() {
		return false
	}
	if !data.Port.IsNull() {
		return false
	}
	if !data.Ssl.IsNull() {
		return false
	}
	if !data.Post.IsNull() {
		return false
	}
	if !data.Input.IsNull() {
		return false
	}
	if !data.Timeout.IsNull() {
		return false
	}
	if !data.Frequency.IsNull() {
		return false
	}
	if !data.Xpath.IsNull() {
		return false
	}
	if !data.Filter.IsNull() {
		return false
	}
	if !data.EnforceTimeout.IsNull() {
		return false
	}
	if !data.IndependentChecks.IsNull() {
		return false
	}
	if !data.GatewayScriptChecks.IsNull() {
		return false
	}
	if !data.GatewayScriptReqMethod.IsNull() {
		return false
	}
	if !data.GatewayScriptCustomReqMethod.IsNull() {
		return false
	}
	if !data.GatewayScriptReqDoc.IsNull() {
		return false
	}
	if !data.GatewayScriptReqContentType.IsNull() {
		return false
	}
	if !data.GatewayScriptRspHandlerMetadata.IsNull() {
		return false
	}
	if !data.GatewayScriptRspHandler.IsNull() {
		return false
	}
	if !data.TcpConnectionType.IsNull() {
		return false
	}
	if !data.SslClientConfigType.IsNull() {
		return false
	}
	if !data.SslClient.IsNull() {
		return false
	}
	return true
}

func (data DmLBGroupCheck) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""

	if !data.Active.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Active`, tfutils.StringFromBool(data.Active, ""))
	}
	if !data.Uri.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`URI`, data.Uri.ValueString())
	}
	if !data.Port.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Port`, data.Port.ValueInt64())
	}
	if !data.Ssl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SSL`, data.Ssl.ValueString())
	}
	if !data.Post.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Post`, tfutils.StringFromBool(data.Post, ""))
	}
	if !data.Input.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Input`, data.Input.ValueString())
	}
	if !data.Timeout.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Timeout`, data.Timeout.ValueInt64())
	}
	if !data.Frequency.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Frequency`, data.Frequency.ValueInt64())
	}
	if !data.Xpath.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`XPath`, data.Xpath.ValueString())
	}
	if !data.Filter.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Filter`, data.Filter.ValueString())
	}
	if !data.EnforceTimeout.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EnforceTimeout`, tfutils.StringFromBool(data.EnforceTimeout, ""))
	}
	if !data.IndependentChecks.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`IndependentChecks`, tfutils.StringFromBool(data.IndependentChecks, ""))
	}
	if !data.GatewayScriptChecks.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`GatewayScriptChecks`, tfutils.StringFromBool(data.GatewayScriptChecks, ""))
	}
	if !data.GatewayScriptReqMethod.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`GatewayScriptReqMethod`, data.GatewayScriptReqMethod.ValueString())
	}
	if !data.GatewayScriptCustomReqMethod.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`GatewayScriptCustomReqMethod`, data.GatewayScriptCustomReqMethod.ValueString())
	}
	if !data.GatewayScriptReqDoc.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`GatewayScriptReqDoc`, data.GatewayScriptReqDoc.ValueString())
	}
	if !data.GatewayScriptReqContentType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`GatewayScriptReqContentType`, data.GatewayScriptReqContentType.ValueString())
	}
	if !data.GatewayScriptRspHandlerMetadata.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`GatewayScriptRspHandlerMetadata`, data.GatewayScriptRspHandlerMetadata.ValueString())
	}
	if !data.GatewayScriptRspHandler.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`GatewayScriptRspHandler`, data.GatewayScriptRspHandler.ValueString())
	}
	if !data.TcpConnectionType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`TCPConnectionType`, data.TcpConnectionType.ValueString())
	}
	if !data.SslClientConfigType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SSLClientConfigType`, data.SslClientConfigType.ValueString())
	}
	if !data.SslClient.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SSLClient`, data.SslClient.ValueString())
	}
	return body
}

func (data *DmLBGroupCheck) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Active`); value.Exists() {
		data.Active = tfutils.BoolFromString(value.String())
	} else {
		data.Active = types.BoolNull()
	}
	if value := res.Get(pathRoot + `URI`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Uri = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Uri = types.StringValue("/")
	}
	if value := res.Get(pathRoot + `Port`); value.Exists() {
		data.Port = types.Int64Value(value.Int())
	} else {
		data.Port = types.Int64Value(80)
	}
	if value := res.Get(pathRoot + `SSL`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Ssl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Ssl = types.StringNull()
	}
	if value := res.Get(pathRoot + `Post`); value.Exists() {
		data.Post = tfutils.BoolFromString(value.String())
	} else {
		data.Post = types.BoolNull()
	}
	if value := res.Get(pathRoot + `Input`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Input = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Input = types.StringValue("store:///healthcheck.xml")
	}
	if value := res.Get(pathRoot + `Timeout`); value.Exists() {
		data.Timeout = types.Int64Value(value.Int())
	} else {
		data.Timeout = types.Int64Value(10)
	}
	if value := res.Get(pathRoot + `Frequency`); value.Exists() {
		data.Frequency = types.Int64Value(value.Int())
	} else {
		data.Frequency = types.Int64Value(180)
	}
	if value := res.Get(pathRoot + `XPath`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Xpath = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Xpath = types.StringValue("/")
	}
	if value := res.Get(pathRoot + `Filter`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Filter = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Filter = types.StringValue("store:///healthcheck.xsl")
	}
	if value := res.Get(pathRoot + `EnforceTimeout`); value.Exists() {
		data.EnforceTimeout = tfutils.BoolFromString(value.String())
	} else {
		data.EnforceTimeout = types.BoolNull()
	}
	if value := res.Get(pathRoot + `IndependentChecks`); value.Exists() {
		data.IndependentChecks = tfutils.BoolFromString(value.String())
	} else {
		data.IndependentChecks = types.BoolNull()
	}
	if value := res.Get(pathRoot + `GatewayScriptChecks`); value.Exists() {
		data.GatewayScriptChecks = tfutils.BoolFromString(value.String())
	} else {
		data.GatewayScriptChecks = types.BoolNull()
	}
	if value := res.Get(pathRoot + `GatewayScriptReqMethod`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.GatewayScriptReqMethod = tfutils.ParseStringFromGJSON(value)
	} else {
		data.GatewayScriptReqMethod = types.StringNull()
	}
	if value := res.Get(pathRoot + `GatewayScriptCustomReqMethod`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.GatewayScriptCustomReqMethod = tfutils.ParseStringFromGJSON(value)
	} else {
		data.GatewayScriptCustomReqMethod = types.StringNull()
	}
	if value := res.Get(pathRoot + `GatewayScriptReqDoc`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.GatewayScriptReqDoc = tfutils.ParseStringFromGJSON(value)
	} else {
		data.GatewayScriptReqDoc = types.StringValue("store:///healthcheck.json")
	}
	if value := res.Get(pathRoot + `GatewayScriptReqContentType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.GatewayScriptReqContentType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.GatewayScriptReqContentType = types.StringValue("application/json")
	}
	if value := res.Get(pathRoot + `GatewayScriptRspHandlerMetadata`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.GatewayScriptRspHandlerMetadata = tfutils.ParseStringFromGJSON(value)
	} else {
		data.GatewayScriptRspHandlerMetadata = types.StringNull()
	}
	if value := res.Get(pathRoot + `GatewayScriptRspHandler`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.GatewayScriptRspHandler = tfutils.ParseStringFromGJSON(value)
	} else {
		data.GatewayScriptRspHandler = types.StringValue("store:///healthcheck.js")
	}
	if value := res.Get(pathRoot + `TCPConnectionType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.TcpConnectionType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.TcpConnectionType = types.StringNull()
	}
	if value := res.Get(pathRoot + `SSLClientConfigType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SslClientConfigType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SslClientConfigType = types.StringValue("client")
	}
	if value := res.Get(pathRoot + `SSLClient`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SslClient = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SslClient = types.StringNull()
	}
}

func (data *DmLBGroupCheck) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `Active`); value.Exists() && !data.Active.IsNull() {
		data.Active = tfutils.BoolFromString(value.String())
	} else if data.Active.ValueBool() {
		data.Active = types.BoolNull()
	}
	if value := res.Get(pathRoot + `URI`); value.Exists() && !data.Uri.IsNull() {
		data.Uri = tfutils.ParseStringFromGJSON(value)
	} else if data.Uri.ValueString() != "/" {
		data.Uri = types.StringNull()
	}
	if value := res.Get(pathRoot + `Port`); value.Exists() && !data.Port.IsNull() {
		data.Port = types.Int64Value(value.Int())
	} else if data.Port.ValueInt64() != 80 {
		data.Port = types.Int64Null()
	}
	if value := res.Get(pathRoot + `SSL`); value.Exists() && !data.Ssl.IsNull() {
		data.Ssl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Ssl = types.StringNull()
	}
	if value := res.Get(pathRoot + `Post`); value.Exists() && !data.Post.IsNull() {
		data.Post = tfutils.BoolFromString(value.String())
	} else if !data.Post.ValueBool() {
		data.Post = types.BoolNull()
	}
	if value := res.Get(pathRoot + `Input`); value.Exists() && !data.Input.IsNull() {
		data.Input = tfutils.ParseStringFromGJSON(value)
	} else if data.Input.ValueString() != "store:///healthcheck.xml" {
		data.Input = types.StringNull()
	}
	if value := res.Get(pathRoot + `Timeout`); value.Exists() && !data.Timeout.IsNull() {
		data.Timeout = types.Int64Value(value.Int())
	} else if data.Timeout.ValueInt64() != 10 {
		data.Timeout = types.Int64Null()
	}
	if value := res.Get(pathRoot + `Frequency`); value.Exists() && !data.Frequency.IsNull() {
		data.Frequency = types.Int64Value(value.Int())
	} else if data.Frequency.ValueInt64() != 180 {
		data.Frequency = types.Int64Null()
	}
	if value := res.Get(pathRoot + `XPath`); value.Exists() && !data.Xpath.IsNull() {
		data.Xpath = tfutils.ParseStringFromGJSON(value)
	} else if data.Xpath.ValueString() != "/" {
		data.Xpath = types.StringNull()
	}
	if value := res.Get(pathRoot + `Filter`); value.Exists() && !data.Filter.IsNull() {
		data.Filter = tfutils.ParseStringFromGJSON(value)
	} else if data.Filter.ValueString() != "store:///healthcheck.xsl" {
		data.Filter = types.StringNull()
	}
	if value := res.Get(pathRoot + `EnforceTimeout`); value.Exists() && !data.EnforceTimeout.IsNull() {
		data.EnforceTimeout = tfutils.BoolFromString(value.String())
	} else if data.EnforceTimeout.ValueBool() {
		data.EnforceTimeout = types.BoolNull()
	}
	if value := res.Get(pathRoot + `IndependentChecks`); value.Exists() && !data.IndependentChecks.IsNull() {
		data.IndependentChecks = tfutils.BoolFromString(value.String())
	} else if data.IndependentChecks.ValueBool() {
		data.IndependentChecks = types.BoolNull()
	}
	if value := res.Get(pathRoot + `GatewayScriptChecks`); value.Exists() && !data.GatewayScriptChecks.IsNull() {
		data.GatewayScriptChecks = tfutils.BoolFromString(value.String())
	} else if data.GatewayScriptChecks.ValueBool() {
		data.GatewayScriptChecks = types.BoolNull()
	}
	if value := res.Get(pathRoot + `GatewayScriptReqMethod`); value.Exists() && !data.GatewayScriptReqMethod.IsNull() {
		data.GatewayScriptReqMethod = tfutils.ParseStringFromGJSON(value)
	} else {
		data.GatewayScriptReqMethod = types.StringNull()
	}
	if value := res.Get(pathRoot + `GatewayScriptCustomReqMethod`); value.Exists() && !data.GatewayScriptCustomReqMethod.IsNull() {
		data.GatewayScriptCustomReqMethod = tfutils.ParseStringFromGJSON(value)
	} else {
		data.GatewayScriptCustomReqMethod = types.StringNull()
	}
	if value := res.Get(pathRoot + `GatewayScriptReqDoc`); value.Exists() && !data.GatewayScriptReqDoc.IsNull() {
		data.GatewayScriptReqDoc = tfutils.ParseStringFromGJSON(value)
	} else if data.GatewayScriptReqDoc.ValueString() != "store:///healthcheck.json" {
		data.GatewayScriptReqDoc = types.StringNull()
	}
	if value := res.Get(pathRoot + `GatewayScriptReqContentType`); value.Exists() && !data.GatewayScriptReqContentType.IsNull() {
		data.GatewayScriptReqContentType = tfutils.ParseStringFromGJSON(value)
	} else if data.GatewayScriptReqContentType.ValueString() != "application/json" {
		data.GatewayScriptReqContentType = types.StringNull()
	}
	if value := res.Get(pathRoot + `GatewayScriptRspHandlerMetadata`); value.Exists() && !data.GatewayScriptRspHandlerMetadata.IsNull() {
		data.GatewayScriptRspHandlerMetadata = tfutils.ParseStringFromGJSON(value)
	} else {
		data.GatewayScriptRspHandlerMetadata = types.StringNull()
	}
	if value := res.Get(pathRoot + `GatewayScriptRspHandler`); value.Exists() && !data.GatewayScriptRspHandler.IsNull() {
		data.GatewayScriptRspHandler = tfutils.ParseStringFromGJSON(value)
	} else if data.GatewayScriptRspHandler.ValueString() != "store:///healthcheck.js" {
		data.GatewayScriptRspHandler = types.StringNull()
	}
	if value := res.Get(pathRoot + `TCPConnectionType`); value.Exists() && !data.TcpConnectionType.IsNull() {
		data.TcpConnectionType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.TcpConnectionType = types.StringNull()
	}
	if value := res.Get(pathRoot + `SSLClientConfigType`); value.Exists() && !data.SslClientConfigType.IsNull() {
		data.SslClientConfigType = tfutils.ParseStringFromGJSON(value)
	} else if data.SslClientConfigType.ValueString() != "client" {
		data.SslClientConfigType = types.StringNull()
	}
	if value := res.Get(pathRoot + `SSLClient`); value.Exists() && !data.SslClient.IsNull() {
		data.SslClient = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SslClient = types.StringNull()
	}
}

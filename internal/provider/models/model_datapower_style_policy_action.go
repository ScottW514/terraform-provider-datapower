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
	"net/url"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/actions"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/tfutils"
	"github.com/scottw514/terraform-provider-datapower/internal/provider/validators"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type StylePolicyAction struct {
	Id                         types.String                     `tfsdk:"id"`
	AppDomain                  types.String                     `tfsdk:"app_domain"`
	UserSummary                types.String                     `tfsdk:"user_summary"`
	Type                       types.String                     `tfsdk:"type"`
	Input                      types.String                     `tfsdk:"input"`
	Transform                  types.String                     `tfsdk:"transform"`
	ParseSettingsReference     *DmDynamicParseSettingsReference `tfsdk:"parse_settings_reference"`
	ParseMetricsResultType     types.String                     `tfsdk:"parse_metrics_result_type"`
	ParseMetricsResultLocation types.String                     `tfsdk:"parse_metrics_result_location"`
	InputLanguage              types.String                     `tfsdk:"input_language"`
	DfdlInputRootName          types.String                     `tfsdk:"dfdl_input_root_name"`
	InputDescriptor            types.String                     `tfsdk:"input_descriptor"`
	OutputDescriptor           types.String                     `tfsdk:"output_descriptor"`
	TransformLanguage          types.String                     `tfsdk:"transform_language"`
	OutputLanguage             types.String                     `tfsdk:"output_language"`
	TxMap                      types.String                     `tfsdk:"tx_map"`
	GatewayScriptLocation      types.String                     `tfsdk:"gateway_script_location"`
	ActionDebug                types.Bool                       `tfsdk:"action_debug"`
	TxTopLevelMap              types.String                     `tfsdk:"tx_top_level_map"`
	TxMode                     types.String                     `tfsdk:"tx_mode"`
	TxAuditLog                 types.String                     `tfsdk:"tx_audit_log"`
	Output                     types.String                     `tfsdk:"output"`
	NoTranscodeUtf8            types.Bool                       `tfsdk:"no_transcode_utf8"`
	NamedInOutLocationType     types.String                     `tfsdk:"named_in_out_location_type"`
	NamedInputs                types.List                       `tfsdk:"named_inputs"`
	NamedOutputs               types.List                       `tfsdk:"named_outputs"`
	Destination                types.String                     `tfsdk:"destination"`
	SchemaUrl                  types.String                     `tfsdk:"schema_url"`
	JsonSchemaUrl              types.String                     `tfsdk:"json_schema_url"`
	WsdlUrl                    types.String                     `tfsdk:"wsdl_url"`
	Policy                     types.String                     `tfsdk:"policy"`
	Aaa                        types.String                     `tfsdk:"aaa"`
	DynamicSchema              types.String                     `tfsdk:"dynamic_schema"`
	DynamicStylesheet          types.String                     `tfsdk:"dynamic_stylesheet"`
	InputConversion            types.String                     `tfsdk:"input_conversion"`
	Xpath                      types.String                     `tfsdk:"xpath"`
	Variable                   types.String                     `tfsdk:"variable"`
	Value                      types.String                     `tfsdk:"value"`
	SslClientConfigType        types.String                     `tfsdk:"ssl_client_config_type"`
	SslClientCred              types.String                     `tfsdk:"ssl_client_cred"`
	AttachmentUri              types.String                     `tfsdk:"attachment_uri"`
	StylesheetParameters       types.List                       `tfsdk:"stylesheet_parameters"`
	ErrorMode                  types.String                     `tfsdk:"error_mode"`
	ErrorInput                 types.String                     `tfsdk:"error_input"`
	ErrorOutput                types.String                     `tfsdk:"error_output"`
	Rule                       types.String                     `tfsdk:"rule"`
	OutputType                 types.String                     `tfsdk:"output_type"`
	LogLevel                   types.String                     `tfsdk:"log_level"`
	LogType                    types.String                     `tfsdk:"log_type"`
	Transactional              types.Bool                       `tfsdk:"transactional"`
	CheckpointEvent            types.String                     `tfsdk:"checkpoint_event"`
	SlmPolicy                  types.String                     `tfsdk:"slm_policy"`
	SqlDataSource              types.String                     `tfsdk:"sql_data_source"`
	SqlText                    types.String                     `tfsdk:"sql_text"`
	SoapValidation             types.String                     `tfsdk:"soap_validation"`
	SqlSourceType              types.String                     `tfsdk:"sql_source_type"`
	JoseSerializationType      types.String                     `tfsdk:"jose_serialization_type"`
	JweEncAlgorithm            types.String                     `tfsdk:"jwe_enc_algorithm"`
	JwsSignatureObject         types.String                     `tfsdk:"jws_signature_object"`
	JweHeaderObject            types.String                     `tfsdk:"jwe_header_object"`
	JoseVerifyType             types.String                     `tfsdk:"jose_verify_type"`
	JoseDecryptType            types.String                     `tfsdk:"jose_decrypt_type"`
	SignatureIdentifier        types.List                       `tfsdk:"signature_identifier"`
	RecipientIdentifier        types.List                       `tfsdk:"recipient_identifier"`
	SingleCertificate          types.String                     `tfsdk:"single_certificate"`
	SingleKey                  types.String                     `tfsdk:"single_key"`
	SingleSskey                types.String                     `tfsdk:"single_sskey"`
	JweDirectKeyObject         types.String                     `tfsdk:"jwe_direct_key_object"`
	JwsVerifyStripSignature    types.Bool                       `tfsdk:"jws_verify_strip_signature"`
	Asynchronous               types.Bool                       `tfsdk:"asynchronous"`
	Condition                  types.List                       `tfsdk:"condition"`
	ResultsMode                types.String                     `tfsdk:"results_mode"`
	RetryCount                 types.Int64                      `tfsdk:"retry_count"`
	RetryInterval              types.Int64                      `tfsdk:"retry_interval"`
	MultipleOutputs            types.Bool                       `tfsdk:"multiple_outputs"`
	IteratorType               types.String                     `tfsdk:"iterator_type"`
	IteratorExpression         types.String                     `tfsdk:"iterator_expression"`
	IteratorCount              types.Int64                      `tfsdk:"iterator_count"`
	LoopAction                 types.String                     `tfsdk:"loop_action"`
	AsyncAction                types.List                       `tfsdk:"async_action"`
	Timeout                    types.Int64                      `tfsdk:"timeout"`
	WsdlPortQName              types.String                     `tfsdk:"wsdl_port_q_name"`
	WsdlOperationName          types.String                     `tfsdk:"wsdl_operation_name"`
	WsdlMessageDirectionOrName types.String                     `tfsdk:"wsdl_message_direction_or_name"`
	WsdlAttachmentPart         types.String                     `tfsdk:"wsdl_attachment_part"`
	MethodRewriteType          types.String                     `tfsdk:"method_rewrite_type"`
	MethodType                 types.String                     `tfsdk:"method_type"`
	MethodType2                types.String                     `tfsdk:"method_type2"`
	PolicyKey                  types.String                     `tfsdk:"policy_key"`
	DependencyActions          []*actions.DependencyAction      `tfsdk:"dependency_actions"`
}

var StylePolicyActionInputCondVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "type",
	AttrType:    "String",
	AttrDefault: "xform",
	Value:       []string{"rewrite", "route-set", "on-error", "method-rewrite"},
}
var StylePolicyActionTransformCondVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "type",
			AttrType:    "String",
			AttrDefault: "xform",
			Value:       []string{"filter", "cryptobin"},
		},
		{
			Evaluation: "logical-and",
			Conditions: []validators.Evaluation{
				{
					Evaluation:  "property-value-in-list",
					Attribute:   "type",
					AttrType:    "String",
					AttrDefault: "xform",
					Value:       []string{"xformng"},
				},
				{
					Evaluation:  "property-value-not-in-list",
					Attribute:   "transform_language",
					AttrType:    "String",
					AttrDefault: "none",
					Value:       []string{"none"},
				},
			},
		},
	},
}
var StylePolicyActionParseMetricsResultLocationCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "type",
			AttrType:    "String",
			AttrDefault: "xform",
			Value:       []string{"parse"},
		},
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "parse_metrics_result_type",
			AttrType:    "String",
			AttrDefault: "none",
			Value:       []string{"none"},
		},
	},
}
var StylePolicyActionInputLanguageCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "type",
	AttrType:    "String",
	AttrDefault: "xform",
	Value:       []string{"xformng"},
}
var StylePolicyActionTransformLanguageCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "type",
	AttrType:    "String",
	AttrDefault: "xform",
	Value:       []string{"xformng"},
}
var StylePolicyActionGatewayScriptLocationCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "type",
	AttrType:    "String",
	AttrDefault: "xform",
	Value:       []string{"gatewayscript", "jose-sign", "jose-verify", "jose-encrypt", "jose-decrypt"},
}
var StylePolicyActionOutputCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "type",
	AttrType:    "String",
	AttrDefault: "xform",
	Value:       []string{"xformpi", "xformbin", "xformng", "cryptobin", "xform", "convert-http", "fetch", "extract", "call", "gatewayscript", "jose-sign", "jose-verify", "jose-encrypt", "jose-decrypt"},
}
var StylePolicyActionDestinationCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "type",
	AttrType:    "String",
	AttrDefault: "xform",
	Value:       []string{"results-async", "fetch", "route-set"},
}
var StylePolicyActionPolicyCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "type",
	AttrType:    "String",
	AttrDefault: "xform",
	Value:       []string{"rewrite"},
}
var StylePolicyActionAAACondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "type",
	AttrType:    "String",
	AttrDefault: "xform",
	Value:       []string{"aaa"},
}
var StylePolicyActionErrorModeCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "type",
	AttrType:    "String",
	AttrDefault: "xform",
	Value:       []string{"on-error"},
}
var StylePolicyActionRuleCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "type",
	AttrType:    "String",
	AttrDefault: "xform",
	Value:       []string{"call"},
}
var StylePolicyActionLogLevelCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "type",
	AttrType:    "String",
	AttrDefault: "xform",
	Value:       []string{"log"},
}
var StylePolicyActionLogTypeCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "type",
	AttrType:    "String",
	AttrDefault: "xform",
	Value:       []string{"log"},
}
var StylePolicyActionCheckpointEventCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "type",
	AttrType:    "String",
	AttrDefault: "xform",
	Value:       []string{"checkpoint"},
}
var StylePolicyActionSLMPolicyCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "type",
	AttrType:    "String",
	AttrDefault: "xform",
	Value:       []string{"slm"},
}
var StylePolicyActionSQLDataSourceCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "type",
	AttrType:    "String",
	AttrDefault: "xform",
	Value:       []string{"sql"},
}
var StylePolicyActionSQLTextCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "sql_source_type",
			AttrType:    "String",
			AttrDefault: "static",
			Value:       []string{"static"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "type",
			AttrType:    "String",
			AttrDefault: "xform",
			Value:       []string{"sql"},
		},
	},
}
var StylePolicyActionSQLSourceTypeCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "type",
	AttrType:    "String",
	AttrDefault: "xform",
	Value:       []string{"sql"},
}
var StylePolicyActionJOSESerializationTypeCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "type",
	AttrType:    "String",
	AttrDefault: "xform",
	Value:       []string{"jose-sign", "jose-encrypt"},
}
var StylePolicyActionJWEEncAlgorithmCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "type",
	AttrType:    "String",
	AttrDefault: "xform",
	Value:       []string{"jose-encrypt"},
}
var StylePolicyActionJWSSignatureObjectCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "type",
	AttrType:    "String",
	AttrDefault: "xform",
	Value:       []string{"jose-sign"},
}
var StylePolicyActionJWEHeaderObjectCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "type",
	AttrType:    "String",
	AttrDefault: "xform",
	Value:       []string{"jose-encrypt"},
}
var StylePolicyActionSignatureIdentifierCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "type",
			AttrType:    "String",
			AttrDefault: "xform",
			Value:       []string{"jose-verify"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "jose_verify_type",
			AttrType:    "String",
			AttrDefault: "",
			Value:       []string{"identifiers"},
		},
	},
}
var StylePolicyActionRecipientIdentifierCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "type",
			AttrType:    "String",
			AttrDefault: "xform",
			Value:       []string{"jose-decrypt"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "jose_decrypt_type",
			AttrType:    "String",
			AttrDefault: "",
			Value:       []string{"identifiers"},
		},
	},
}
var StylePolicyActionSingleCertificateCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "type",
			AttrType:    "String",
			AttrDefault: "xform",
			Value:       []string{"jose-verify"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "jose_verify_type",
			AttrType:    "String",
			AttrDefault: "",
			Value:       []string{"single-cert"},
		},
	},
}
var StylePolicyActionSingleKeyCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "type",
			AttrType:    "String",
			AttrDefault: "xform",
			Value:       []string{"jose-decrypt"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "jose_decrypt_type",
			AttrType:    "String",
			AttrDefault: "",
			Value:       []string{"single-key"},
		},
	},
}
var StylePolicyActionSingleSSKeyCondVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation: "logical-and",
			Conditions: []validators.Evaluation{
				{
					Evaluation:  "property-value-in-list",
					Attribute:   "type",
					AttrType:    "String",
					AttrDefault: "xform",
					Value:       []string{"jose-decrypt"},
				},
				{
					Evaluation:  "property-value-in-list",
					Attribute:   "jose_decrypt_type",
					AttrType:    "String",
					AttrDefault: "",
					Value:       []string{"single-sskey"},
				},
			},
		},
		{
			Evaluation: "logical-and",
			Conditions: []validators.Evaluation{
				{
					Evaluation:  "property-value-in-list",
					Attribute:   "type",
					AttrType:    "String",
					AttrDefault: "xform",
					Value:       []string{"jose-verify"},
				},
				{
					Evaluation:  "property-value-in-list",
					Attribute:   "jose_verify_type",
					AttrType:    "String",
					AttrDefault: "",
					Value:       []string{"single-sskey"},
				},
			},
		},
	},
}
var StylePolicyActionJWEDirectKeyObjectCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "type",
			AttrType:    "String",
			AttrDefault: "xform",
			Value:       []string{"jose-decrypt"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "jose_decrypt_type",
			AttrType:    "String",
			AttrDefault: "",
			Value:       []string{"direct-key"},
		},
	},
}
var StylePolicyActionIteratorTypeCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "type",
	AttrType:    "String",
	AttrDefault: "xform",
	Value:       []string{"for-each"},
}
var StylePolicyActionIteratorExpressionCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "type",
			AttrType:    "String",
			AttrDefault: "xform",
			Value:       []string{"for-each"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "iterator_type",
			AttrType:    "String",
			AttrDefault: "XPATH",
			Value:       []string{"XPATH"},
		},
	},
}
var StylePolicyActionIteratorCountCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "type",
			AttrType:    "String",
			AttrDefault: "xform",
			Value:       []string{"for-each"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "iterator_type",
			AttrType:    "String",
			AttrDefault: "XPATH",
			Value:       []string{"COUNT"},
		},
	},
}
var StylePolicyActionLoopActionCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "type",
	AttrType:    "String",
	AttrDefault: "xform",
	Value:       []string{"for-each"},
}
var StylePolicyActionMethodRewriteTypeCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "type",
	AttrType:    "String",
	AttrDefault: "xform",
	Value:       []string{"method-rewrite", "fetch"},
}
var StylePolicyActionMethodTypeCondVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "type",
			AttrType:    "String",
			AttrDefault: "xform",
			Value:       []string{"results-async"},
		},
		{
			Evaluation: "logical-and",
			Conditions: []validators.Evaluation{
				{
					Evaluation:  "property-value-in-list",
					Attribute:   "type",
					AttrType:    "String",
					AttrDefault: "xform",
					Value:       []string{"results"},
				},
				{
					Evaluation:  "property-value-not-in-list",
					Attribute:   "destination",
					AttrType:    "String",
					AttrDefault: "",
					Value:       []string{""},
				},
			},
		},
	},
}
var StylePolicyActionMethodType2CondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "type",
	AttrType:    "String",
	AttrDefault: "xform",
	Value:       []string{"log"},
}
var StylePolicyActionInputIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}
var StylePolicyActionTransformIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation: "logical-and",
			Conditions: []validators.Evaluation{
				{
					Evaluation:  "property-value-in-list",
					Attribute:   "type",
					AttrType:    "String",
					AttrDefault: "xform",
					Value:       []string{"xformng"},
				},
				{
					Evaluation:  "property-value-in-list",
					Attribute:   "transform_language",
					AttrType:    "String",
					AttrDefault: "none",
					Value:       []string{"none"},
				},
			},
		},
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "type",
			AttrType:    "String",
			AttrDefault: "xform",
			Value:       []string{"route-action", "xformpi", "xformbin", "xformng", "xform", "sql"},
		},
	},
}
var StylePolicyActionParseSettingsReferenceIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "type",
	AttrType:    "String",
	AttrDefault: "xform",
	Value:       []string{"parse"},
}
var StylePolicyActionParseMetricsResultTypeIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "type",
	AttrType:    "String",
	AttrDefault: "xform",
	Value:       []string{"parse"},
}
var StylePolicyActionParseMetricsResultLocationIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "type",
	AttrType:    "String",
	AttrDefault: "xform",
	Value:       []string{"parse"},
}
var StylePolicyActionInputLanguageIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "type",
	AttrType:    "String",
	AttrDefault: "xform",
	Value:       []string{"xformng"},
}
var StylePolicyActionDFDLInputRootNameIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "type",
			AttrType:    "String",
			AttrDefault: "xform",
			Value:       []string{"xformng"},
		},
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "input_language",
			AttrType:    "String",
			AttrDefault: "xml",
			Value:       []string{"dfdl"},
		},
	},
}
var StylePolicyActionInputDescriptorIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "type",
			AttrType:    "String",
			AttrDefault: "xform",
			Value:       []string{"xformng"},
		},
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "input_language",
			AttrType:    "String",
			AttrDefault: "xml",
			Value:       []string{"xsd", "dfdl"},
		},
	},
}
var StylePolicyActionOutputDescriptorIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "type",
			AttrType:    "String",
			AttrDefault: "xform",
			Value:       []string{"xformng"},
		},
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "output_language",
			AttrType:    "String",
			AttrDefault: "default",
			Value:       []string{"dfdl"},
		},
	},
}
var StylePolicyActionTransformLanguageIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "type",
	AttrType:    "String",
	AttrDefault: "xform",
	Value:       []string{"xformng"},
}
var StylePolicyActionOutputLanguageIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "type",
	AttrType:    "String",
	AttrDefault: "xform",
	Value:       []string{"xformng"},
}
var StylePolicyActionTxMapIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "type",
	AttrType:    "String",
	AttrDefault: "xform",
	Value:       []string{"xformbin"},
}
var StylePolicyActionGatewayScriptLocationIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}
var StylePolicyActionActionDebugIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "type",
	AttrType:    "String",
	AttrDefault: "xform",
	Value:       []string{"gatewayscript"},
}
var StylePolicyActionTxTopLevelMapIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "type",
			AttrType:    "String",
			AttrDefault: "xform",
			Value:       []string{"xformbin"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "tx_mode",
			AttrType:    "String",
			AttrDefault: "",
			Value:       []string{"dpa"},
		},
	},
}
var StylePolicyActionTxModeIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "type",
	AttrType:    "String",
	AttrDefault: "xform",
	Value:       []string{"xformbin"},
}
var StylePolicyActionTxAuditLogIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "type",
			AttrType:    "String",
			AttrDefault: "xform",
			Value:       []string{"xformbin"},
		},
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "tx_mode",
			AttrType:    "String",
			AttrDefault: "",
			Value:       []string{"dpa"},
		},
	},
}
var StylePolicyActionOutputIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "type",
	AttrType:    "String",
	AttrDefault: "xform",
	Value:       []string{"results-async", "setvar", "rewrite", "route-set", "strip-attachments", "on-error", "checkpoint", "conditional", "event-sink", "method-rewrite"},
}
var StylePolicyActionNoTranscodeUtf8IgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "type",
	AttrType:    "String",
	AttrDefault: "xform",
	Value:       []string{"convert-http"},
}
var StylePolicyActionNamedInOutLocationTypeIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "type",
	AttrType:    "String",
	AttrDefault: "xform",
	Value:       []string{"xformbin"},
}
var StylePolicyActionNamedInputsIgnoreVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "named_in_out_location_type",
			AttrType:    "String",
			AttrDefault: "default",
			Value:       []string{"explicit"},
		},
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "type",
			AttrType:    "String",
			AttrDefault: "xform",
			Value:       []string{"xformbin"},
		},
	},
}
var StylePolicyActionNamedOutputsIgnoreVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "named_in_out_location_type",
			AttrType:    "String",
			AttrDefault: "default",
			Value:       []string{"explicit"},
		},
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "type",
			AttrType:    "String",
			AttrDefault: "xform",
			Value:       []string{"xformbin"},
		},
	},
}
var StylePolicyActionDestinationIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "type",
	AttrType:    "String",
	AttrDefault: "xform",
	Value:       []string{"results", "log"},
}
var StylePolicyActionSchemaURLIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "type",
	AttrType:    "String",
	AttrDefault: "xform",
	Value:       []string{"validate"},
}
var StylePolicyActionJSONSchemaURLIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "type",
	AttrType:    "String",
	AttrDefault: "xform",
	Value:       []string{"validate"},
}
var StylePolicyActionWsdlURLIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "type",
	AttrType:    "String",
	AttrDefault: "xform",
	Value:       []string{"validate"},
}
var StylePolicyActionPolicyIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "type",
	AttrType:    "String",
	AttrDefault: "xform",
	Value:       []string{"validate", "xformpi", "xformbin", "xformng", "xform"},
}
var StylePolicyActionAAAIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}
var StylePolicyActionDynamicSchemaIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "type",
	AttrType:    "String",
	AttrDefault: "xform",
	Value:       []string{"validate"},
}
var StylePolicyActionDynamicStylesheetIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "type",
	AttrType:    "String",
	AttrDefault: "xform",
	Value:       []string{"filter", "route-action", "xformpi", "xformbin", "xformng", "xform"},
}
var StylePolicyActionInputConversionIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "type",
	AttrType:    "String",
	AttrDefault: "xform",
	Value:       []string{"convert-http"},
}
var StylePolicyActionXPathIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "type",
	AttrType:    "String",
	AttrDefault: "xform",
	Value:       []string{"extract"},
}
var StylePolicyActionVariableIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "type",
			AttrType:    "String",
			AttrDefault: "xform",
			Value:       []string{"extract", "setvar", "sql"},
		},
		{
			Evaluation: "logical-and",
			Conditions: []validators.Evaluation{
				{
					Evaluation:  "property-value-in-list",
					Attribute:   "type",
					AttrType:    "String",
					AttrDefault: "xform",
					Value:       []string{"sql"},
				},
				{
					Evaluation:  "property-value-not-in-list",
					Attribute:   "sql_source_type",
					AttrType:    "String",
					AttrDefault: "static",
					Value:       []string{"variable"},
				},
			},
		},
	},
}
var StylePolicyActionValueIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "type",
	AttrType:    "String",
	AttrDefault: "xform",
	Value:       []string{"setvar"},
}
var StylePolicyActionSSLClientConfigTypeIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "type",
	AttrType:    "String",
	AttrDefault: "xform",
	Value:       []string{"route-set"},
}
var StylePolicyActionSSLClientCredIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "type",
			AttrType:    "String",
			AttrDefault: "xform",
			Value:       []string{"route-set"},
		},
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "ssl_client_config_type",
			AttrType:    "String",
			AttrDefault: "client",
			Value:       []string{"client"},
		},
	},
}
var StylePolicyActionAttachmentURIIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "type",
	AttrType:    "String",
	AttrDefault: "xform",
	Value:       []string{"strip-attachments"},
}
var StylePolicyActionStylesheetParametersIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "type",
	AttrType:    "String",
	AttrDefault: "xform",
	Value:       []string{"filter", "route-action", "xformpi", "xformbin", "xformng", "cryptobin", "xform", "gatewayscript"},
}
var StylePolicyActionErrorModeIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}
var StylePolicyActionErrorInputIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "type",
	AttrType:    "String",
	AttrDefault: "xform",
	Value:       []string{"on-error"},
}
var StylePolicyActionErrorOutputIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "type",
	AttrType:    "String",
	AttrDefault: "xform",
	Value:       []string{"on-error"},
}
var StylePolicyActionRuleIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "type",
	AttrType:    "String",
	AttrDefault: "xform",
	Value:       []string{"on-error"},
}
var StylePolicyActionOutputTypeIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "type",
	AttrType:    "String",
	AttrDefault: "xform",
	Value:       []string{"xformpi", "xformbin", "xformng", "xform", "results", "log", "fetch"},
}
var StylePolicyActionLogLevelIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}
var StylePolicyActionLogTypeIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}
var StylePolicyActionTransactionalIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "type",
	AttrType:    "String",
	AttrDefault: "xform",
	Value:       []string{"results", "results-async"},
}
var StylePolicyActionCheckpointEventIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}
var StylePolicyActionSLMPolicyIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}
var StylePolicyActionSQLDataSourceIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}
var StylePolicyActionSQLTextIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "type",
			AttrType:    "String",
			AttrDefault: "xform",
			Value:       []string{"sql"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "sql_source_type",
			AttrType:    "String",
			AttrDefault: "static",
			Value:       []string{"stylesheet", "variable"},
		},
	},
}
var StylePolicyActionSOAPValidationIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "type",
			AttrType:    "String",
			AttrDefault: "xform",
			Value:       []string{"validate"},
		},
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "wsdl_url",
			AttrType:    "String",
			AttrDefault: "",
			Value:       []string{""},
		},
	},
}
var StylePolicyActionSQLSourceTypeIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "type",
	AttrType:    "String",
	AttrDefault: "xform",
	Value:       []string{"sql"},
}
var StylePolicyActionJOSESerializationTypeIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}
var StylePolicyActionJWEEncAlgorithmIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}
var StylePolicyActionJWSSignatureObjectIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}
var StylePolicyActionJWEHeaderObjectIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}
var StylePolicyActionJOSEVerifyTypeIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "type",
	AttrType:    "String",
	AttrDefault: "xform",
	Value:       []string{"jose-verify"},
}
var StylePolicyActionJOSEDecryptTypeIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "type",
	AttrType:    "String",
	AttrDefault: "xform",
	Value:       []string{"jose-decrypt"},
}
var StylePolicyActionSignatureIdentifierIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}
var StylePolicyActionRecipientIdentifierIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}
var StylePolicyActionSingleCertificateIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}
var StylePolicyActionSingleKeyIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}
var StylePolicyActionSingleSSKeyIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}
var StylePolicyActionJWEDirectKeyObjectIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}
var StylePolicyActionJWSVerifyStripSignatureIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "type",
	AttrType:    "String",
	AttrDefault: "xform",
	Value:       []string{"jose-verify"},
}
var StylePolicyActionAsynchronousIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "type",
	AttrType:    "String",
	AttrDefault: "xform",
	Value:       []string{"results-async", "event-sink", "method-rewrite", "gatewayscript", "jose-sign", "jose-verify", "jose-encrypt", "jose-decrypt"},
}
var StylePolicyActionConditionIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "type",
	AttrType:    "String",
	AttrDefault: "xform",
	Value:       []string{"conditional"},
}
var StylePolicyActionResultsModeIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "type",
	AttrType:    "String",
	AttrDefault: "xform",
	Value:       []string{"results", "results-async"},
}
var StylePolicyActionRetryCountIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "type",
	AttrType:    "String",
	AttrDefault: "xform",
	Value:       []string{"results", "results-async"},
}
var StylePolicyActionRetryIntervalIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "type",
	AttrType:    "String",
	AttrDefault: "xform",
	Value:       []string{"results", "results-async"},
}
var StylePolicyActionMultipleOutputsIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "type",
			AttrType:    "String",
			AttrDefault: "xform",
			Value:       []string{"results", "for-each"},
		},
		{
			Evaluation: "logical-and",
			Conditions: []validators.Evaluation{
				{
					Evaluation:  "property-value-in-list",
					Attribute:   "type",
					AttrType:    "String",
					AttrDefault: "xform",
					Value:       []string{"results"},
				},
				{
					Evaluation:  "property-value-in-list",
					Attribute:   "destination",
					AttrType:    "String",
					AttrDefault: "",
					Value:       []string{""},
				},
			},
		},
	},
}
var StylePolicyActionIteratorTypeIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "type",
	AttrType:    "String",
	AttrDefault: "xform",
	Value:       []string{"for-each"},
}
var StylePolicyActionIteratorExpressionIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "type",
			AttrType:    "String",
			AttrDefault: "xform",
			Value:       []string{"for-each"},
		},
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "iterator_type",
			AttrType:    "String",
			AttrDefault: "XPATH",
			Value:       []string{"XPATH"},
		},
	},
}
var StylePolicyActionIteratorCountIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "type",
			AttrType:    "String",
			AttrDefault: "xform",
			Value:       []string{"for-each"},
		},
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "iterator_type",
			AttrType:    "String",
			AttrDefault: "XPATH",
			Value:       []string{"COUNT"},
		},
	},
}
var StylePolicyActionLoopActionIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}
var StylePolicyActionAsyncActionIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "type",
	AttrType:    "String",
	AttrDefault: "xform",
	Value:       []string{"event-sink"},
}
var StylePolicyActionTimeoutIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "type",
	AttrType:    "String",
	AttrDefault: "xform",
	Value:       []string{"event-sink"},
}
var StylePolicyActionWSDLPortQNameIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "type",
	AttrType:    "String",
	AttrDefault: "xform",
	Value:       []string{"validate"},
}
var StylePolicyActionWSDLOperationNameIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "type",
	AttrType:    "String",
	AttrDefault: "xform",
	Value:       []string{"validate"},
}
var StylePolicyActionWSDLMessageDirectionOrNameIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "type",
	AttrType:    "String",
	AttrDefault: "xform",
	Value:       []string{"validate"},
}
var StylePolicyActionWSDLAttachmentPartIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "type",
	AttrType:    "String",
	AttrDefault: "xform",
	Value:       []string{"validate"},
}
var StylePolicyActionMethodRewriteTypeIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "type",
	AttrType:    "String",
	AttrDefault: "xform",
	Value:       []string{"method-rewrite", "fetch"},
}
var StylePolicyActionMethodTypeIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "type",
			AttrType:    "String",
			AttrDefault: "xform",
			Value:       []string{"results", "results-async"},
		},
		{
			Evaluation: "logical-and",
			Conditions: []validators.Evaluation{
				{
					Evaluation:  "property-value-in-list",
					Attribute:   "type",
					AttrType:    "String",
					AttrDefault: "xform",
					Value:       []string{"results"},
				},
				{
					Evaluation:  "property-value-in-list",
					Attribute:   "destination",
					AttrType:    "String",
					AttrDefault: "",
					Value:       []string{""},
				},
			},
		},
	},
}
var StylePolicyActionMethodType2IgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "type",
	AttrType:    "String",
	AttrDefault: "xform",
	Value:       []string{"log"},
}

var StylePolicyActionObjectType = map[string]attr.Type{
	"id":                             types.StringType,
	"app_domain":                     types.StringType,
	"user_summary":                   types.StringType,
	"type":                           types.StringType,
	"input":                          types.StringType,
	"transform":                      types.StringType,
	"parse_settings_reference":       types.ObjectType{AttrTypes: DmDynamicParseSettingsReferenceObjectType},
	"parse_metrics_result_type":      types.StringType,
	"parse_metrics_result_location":  types.StringType,
	"input_language":                 types.StringType,
	"dfdl_input_root_name":           types.StringType,
	"input_descriptor":               types.StringType,
	"output_descriptor":              types.StringType,
	"transform_language":             types.StringType,
	"output_language":                types.StringType,
	"tx_map":                         types.StringType,
	"gateway_script_location":        types.StringType,
	"action_debug":                   types.BoolType,
	"tx_top_level_map":               types.StringType,
	"tx_mode":                        types.StringType,
	"tx_audit_log":                   types.StringType,
	"output":                         types.StringType,
	"no_transcode_utf8":              types.BoolType,
	"named_in_out_location_type":     types.StringType,
	"named_inputs":                   types.ListType{ElemType: types.ObjectType{AttrTypes: DmNamedInOutObjectType}},
	"named_outputs":                  types.ListType{ElemType: types.ObjectType{AttrTypes: DmNamedInOutObjectType}},
	"destination":                    types.StringType,
	"schema_url":                     types.StringType,
	"json_schema_url":                types.StringType,
	"wsdl_url":                       types.StringType,
	"policy":                         types.StringType,
	"aaa":                            types.StringType,
	"dynamic_schema":                 types.StringType,
	"dynamic_stylesheet":             types.StringType,
	"input_conversion":               types.StringType,
	"xpath":                          types.StringType,
	"variable":                       types.StringType,
	"value":                          types.StringType,
	"ssl_client_config_type":         types.StringType,
	"ssl_client_cred":                types.StringType,
	"attachment_uri":                 types.StringType,
	"stylesheet_parameters":          types.ListType{ElemType: types.ObjectType{AttrTypes: DmStylesheetParameterObjectType}},
	"error_mode":                     types.StringType,
	"error_input":                    types.StringType,
	"error_output":                   types.StringType,
	"rule":                           types.StringType,
	"output_type":                    types.StringType,
	"log_level":                      types.StringType,
	"log_type":                       types.StringType,
	"transactional":                  types.BoolType,
	"checkpoint_event":               types.StringType,
	"slm_policy":                     types.StringType,
	"sql_data_source":                types.StringType,
	"sql_text":                       types.StringType,
	"soap_validation":                types.StringType,
	"sql_source_type":                types.StringType,
	"jose_serialization_type":        types.StringType,
	"jwe_enc_algorithm":              types.StringType,
	"jws_signature_object":           types.StringType,
	"jwe_header_object":              types.StringType,
	"jose_verify_type":               types.StringType,
	"jose_decrypt_type":              types.StringType,
	"signature_identifier":           types.ListType{ElemType: types.StringType},
	"recipient_identifier":           types.ListType{ElemType: types.StringType},
	"single_certificate":             types.StringType,
	"single_key":                     types.StringType,
	"single_sskey":                   types.StringType,
	"jwe_direct_key_object":          types.StringType,
	"jws_verify_strip_signature":     types.BoolType,
	"asynchronous":                   types.BoolType,
	"condition":                      types.ListType{ElemType: types.ObjectType{AttrTypes: DmConditionObjectType}},
	"results_mode":                   types.StringType,
	"retry_count":                    types.Int64Type,
	"retry_interval":                 types.Int64Type,
	"multiple_outputs":               types.BoolType,
	"iterator_type":                  types.StringType,
	"iterator_expression":            types.StringType,
	"iterator_count":                 types.Int64Type,
	"loop_action":                    types.StringType,
	"async_action":                   types.ListType{ElemType: types.StringType},
	"timeout":                        types.Int64Type,
	"wsdl_port_q_name":               types.StringType,
	"wsdl_operation_name":            types.StringType,
	"wsdl_message_direction_or_name": types.StringType,
	"wsdl_attachment_part":           types.StringType,
	"method_rewrite_type":            types.StringType,
	"method_type":                    types.StringType,
	"method_type2":                   types.StringType,
	"policy_key":                     types.StringType,
	"dependency_actions":             actions.ActionsListType,
}

func (data StylePolicyAction) GetPath() string {
	rest_path := "/mgmt/config/{domain}/StylePolicyAction"
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data StylePolicyAction) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.Type.IsNull() {
		return false
	}
	if !data.Input.IsNull() {
		return false
	}
	if !data.Transform.IsNull() {
		return false
	}
	if data.ParseSettingsReference != nil {
		if !data.ParseSettingsReference.IsNull() {
			return false
		}
	}
	if !data.ParseMetricsResultType.IsNull() {
		return false
	}
	if !data.ParseMetricsResultLocation.IsNull() {
		return false
	}
	if !data.InputLanguage.IsNull() {
		return false
	}
	if !data.DfdlInputRootName.IsNull() {
		return false
	}
	if !data.InputDescriptor.IsNull() {
		return false
	}
	if !data.OutputDescriptor.IsNull() {
		return false
	}
	if !data.TransformLanguage.IsNull() {
		return false
	}
	if !data.OutputLanguage.IsNull() {
		return false
	}
	if !data.TxMap.IsNull() {
		return false
	}
	if !data.GatewayScriptLocation.IsNull() {
		return false
	}
	if !data.ActionDebug.IsNull() {
		return false
	}
	if !data.TxTopLevelMap.IsNull() {
		return false
	}
	if !data.TxMode.IsNull() {
		return false
	}
	if !data.TxAuditLog.IsNull() {
		return false
	}
	if !data.Output.IsNull() {
		return false
	}
	if !data.NoTranscodeUtf8.IsNull() {
		return false
	}
	if !data.NamedInOutLocationType.IsNull() {
		return false
	}
	if !data.NamedInputs.IsNull() {
		return false
	}
	if !data.NamedOutputs.IsNull() {
		return false
	}
	if !data.Destination.IsNull() {
		return false
	}
	if !data.SchemaUrl.IsNull() {
		return false
	}
	if !data.JsonSchemaUrl.IsNull() {
		return false
	}
	if !data.WsdlUrl.IsNull() {
		return false
	}
	if !data.Policy.IsNull() {
		return false
	}
	if !data.Aaa.IsNull() {
		return false
	}
	if !data.DynamicSchema.IsNull() {
		return false
	}
	if !data.DynamicStylesheet.IsNull() {
		return false
	}
	if !data.InputConversion.IsNull() {
		return false
	}
	if !data.Xpath.IsNull() {
		return false
	}
	if !data.Variable.IsNull() {
		return false
	}
	if !data.Value.IsNull() {
		return false
	}
	if !data.SslClientConfigType.IsNull() {
		return false
	}
	if !data.SslClientCred.IsNull() {
		return false
	}
	if !data.AttachmentUri.IsNull() {
		return false
	}
	if !data.StylesheetParameters.IsNull() {
		return false
	}
	if !data.ErrorMode.IsNull() {
		return false
	}
	if !data.ErrorInput.IsNull() {
		return false
	}
	if !data.ErrorOutput.IsNull() {
		return false
	}
	if !data.Rule.IsNull() {
		return false
	}
	if !data.OutputType.IsNull() {
		return false
	}
	if !data.LogLevel.IsNull() {
		return false
	}
	if !data.LogType.IsNull() {
		return false
	}
	if !data.Transactional.IsNull() {
		return false
	}
	if !data.CheckpointEvent.IsNull() {
		return false
	}
	if !data.SlmPolicy.IsNull() {
		return false
	}
	if !data.SqlDataSource.IsNull() {
		return false
	}
	if !data.SqlText.IsNull() {
		return false
	}
	if !data.SoapValidation.IsNull() {
		return false
	}
	if !data.SqlSourceType.IsNull() {
		return false
	}
	if !data.JoseSerializationType.IsNull() {
		return false
	}
	if !data.JweEncAlgorithm.IsNull() {
		return false
	}
	if !data.JwsSignatureObject.IsNull() {
		return false
	}
	if !data.JweHeaderObject.IsNull() {
		return false
	}
	if !data.JoseVerifyType.IsNull() {
		return false
	}
	if !data.JoseDecryptType.IsNull() {
		return false
	}
	if !data.SignatureIdentifier.IsNull() {
		return false
	}
	if !data.RecipientIdentifier.IsNull() {
		return false
	}
	if !data.SingleCertificate.IsNull() {
		return false
	}
	if !data.SingleKey.IsNull() {
		return false
	}
	if !data.SingleSskey.IsNull() {
		return false
	}
	if !data.JweDirectKeyObject.IsNull() {
		return false
	}
	if !data.JwsVerifyStripSignature.IsNull() {
		return false
	}
	if !data.Asynchronous.IsNull() {
		return false
	}
	if !data.Condition.IsNull() {
		return false
	}
	if !data.ResultsMode.IsNull() {
		return false
	}
	if !data.RetryCount.IsNull() {
		return false
	}
	if !data.RetryInterval.IsNull() {
		return false
	}
	if !data.MultipleOutputs.IsNull() {
		return false
	}
	if !data.IteratorType.IsNull() {
		return false
	}
	if !data.IteratorExpression.IsNull() {
		return false
	}
	if !data.IteratorCount.IsNull() {
		return false
	}
	if !data.LoopAction.IsNull() {
		return false
	}
	if !data.AsyncAction.IsNull() {
		return false
	}
	if !data.Timeout.IsNull() {
		return false
	}
	if !data.WsdlPortQName.IsNull() {
		return false
	}
	if !data.WsdlOperationName.IsNull() {
		return false
	}
	if !data.WsdlMessageDirectionOrName.IsNull() {
		return false
	}
	if !data.WsdlAttachmentPart.IsNull() {
		return false
	}
	if !data.MethodRewriteType.IsNull() {
		return false
	}
	if !data.MethodType.IsNull() {
		return false
	}
	if !data.MethodType2.IsNull() {
		return false
	}
	if !data.PolicyKey.IsNull() {
		return false
	}
	return true
}

func (data StylePolicyAction) ToBody(ctx context.Context, pathRoot string) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""

	if !data.Id.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`name`, data.Id.ValueString())
	}
	if !data.UserSummary.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UserSummary`, data.UserSummary.ValueString())
	}
	if !data.Type.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Type`, data.Type.ValueString())
	}
	if !data.Input.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Input`, data.Input.ValueString())
	}
	if !data.Transform.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Transform`, data.Transform.ValueString())
	}
	if data.ParseSettingsReference != nil {
		if !data.ParseSettingsReference.IsNull() {
			body, _ = sjson.SetRaw(body, pathRoot+`ParseSettingsReference`, data.ParseSettingsReference.ToBody(ctx, ""))
		}
	}
	if !data.ParseMetricsResultType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ParseMetricsResultType`, data.ParseMetricsResultType.ValueString())
	}
	if !data.ParseMetricsResultLocation.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ParseMetricsResultLocation`, data.ParseMetricsResultLocation.ValueString())
	}
	if !data.InputLanguage.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`InputLanguage`, data.InputLanguage.ValueString())
	}
	if !data.DfdlInputRootName.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DFDLInputRootName`, data.DfdlInputRootName.ValueString())
	}
	if !data.InputDescriptor.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`InputDescriptor`, data.InputDescriptor.ValueString())
	}
	if !data.OutputDescriptor.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`OutputDescriptor`, data.OutputDescriptor.ValueString())
	}
	if !data.TransformLanguage.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`TransformLanguage`, data.TransformLanguage.ValueString())
	}
	if !data.OutputLanguage.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`OutputLanguage`, data.OutputLanguage.ValueString())
	}
	if !data.TxMap.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`TxMap`, data.TxMap.ValueString())
	}
	if !data.GatewayScriptLocation.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`GatewayScriptLocation`, data.GatewayScriptLocation.ValueString())
	}
	if !data.ActionDebug.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ActionDebug`, tfutils.StringFromBool(data.ActionDebug, ""))
	}
	if !data.TxTopLevelMap.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`TxTopLevelMap`, data.TxTopLevelMap.ValueString())
	}
	if !data.TxMode.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`TxMode`, data.TxMode.ValueString())
	}
	if !data.TxAuditLog.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`TxAuditLog`, data.TxAuditLog.ValueString())
	}
	if !data.Output.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Output`, data.Output.ValueString())
	}
	if !data.NoTranscodeUtf8.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`NoTranscodeUtf8`, tfutils.StringFromBool(data.NoTranscodeUtf8, ""))
	}
	if !data.NamedInOutLocationType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`NamedInOutLocationType`, data.NamedInOutLocationType.ValueString())
	}
	if !data.NamedInputs.IsNull() {
		var dataValues []DmNamedInOut
		data.NamedInputs.ElementsAs(ctx, &dataValues, false)
		for _, val := range dataValues {
			body, _ = sjson.SetRaw(body, pathRoot+`NamedInputs`+".-1", val.ToBody(ctx, ""))
		}
	}
	if !data.NamedOutputs.IsNull() {
		var dataValues []DmNamedInOut
		data.NamedOutputs.ElementsAs(ctx, &dataValues, false)
		for _, val := range dataValues {
			body, _ = sjson.SetRaw(body, pathRoot+`NamedOutputs`+".-1", val.ToBody(ctx, ""))
		}
	}
	if !data.Destination.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Destination`, data.Destination.ValueString())
	}
	if !data.SchemaUrl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SchemaURL`, data.SchemaUrl.ValueString())
	}
	if !data.JsonSchemaUrl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`JSONSchemaURL`, data.JsonSchemaUrl.ValueString())
	}
	if !data.WsdlUrl.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`WsdlURL`, data.WsdlUrl.ValueString())
	}
	if !data.Policy.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Policy`, data.Policy.ValueString())
	}
	if !data.Aaa.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AAA`, data.Aaa.ValueString())
	}
	if !data.DynamicSchema.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DynamicSchema`, data.DynamicSchema.ValueString())
	}
	if !data.DynamicStylesheet.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`DynamicStylesheet`, data.DynamicStylesheet.ValueString())
	}
	if !data.InputConversion.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`InputConversion`, data.InputConversion.ValueString())
	}
	if !data.Xpath.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`XPath`, data.Xpath.ValueString())
	}
	if !data.Variable.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Variable`, data.Variable.ValueString())
	}
	if !data.Value.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Value`, data.Value.ValueString())
	}
	if !data.SslClientConfigType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SSLClientConfigType`, data.SslClientConfigType.ValueString())
	}
	if !data.SslClientCred.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SSLClientCred`, data.SslClientCred.ValueString())
	}
	if !data.AttachmentUri.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`AttachmentURI`, data.AttachmentUri.ValueString())
	}
	if !data.StylesheetParameters.IsNull() {
		var dataValues []DmStylesheetParameter
		data.StylesheetParameters.ElementsAs(ctx, &dataValues, false)
		for _, val := range dataValues {
			body, _ = sjson.SetRaw(body, pathRoot+`StylesheetParameters`+".-1", val.ToBody(ctx, ""))
		}
	}
	if !data.ErrorMode.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ErrorMode`, data.ErrorMode.ValueString())
	}
	if !data.ErrorInput.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ErrorInput`, data.ErrorInput.ValueString())
	}
	if !data.ErrorOutput.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ErrorOutput`, data.ErrorOutput.ValueString())
	}
	if !data.Rule.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Rule`, data.Rule.ValueString())
	}
	if !data.OutputType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`OutputType`, data.OutputType.ValueString())
	}
	if !data.LogLevel.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LogLevel`, data.LogLevel.ValueString())
	}
	if !data.LogType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LogType`, data.LogType.ValueString())
	}
	if !data.Transactional.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Transactional`, tfutils.StringFromBool(data.Transactional, ""))
	}
	if !data.CheckpointEvent.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`CheckpointEvent`, data.CheckpointEvent.ValueString())
	}
	if !data.SlmPolicy.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SLMPolicy`, data.SlmPolicy.ValueString())
	}
	if !data.SqlDataSource.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SQLDataSource`, data.SqlDataSource.ValueString())
	}
	if !data.SqlText.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SQLText`, data.SqlText.ValueString())
	}
	if !data.SoapValidation.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SOAPValidation`, data.SoapValidation.ValueString())
	}
	if !data.SqlSourceType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SQLSourceType`, data.SqlSourceType.ValueString())
	}
	if !data.JoseSerializationType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`JOSESerializationType`, data.JoseSerializationType.ValueString())
	}
	if !data.JweEncAlgorithm.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`JWEEncAlgorithm`, data.JweEncAlgorithm.ValueString())
	}
	if !data.JwsSignatureObject.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`JWSSignatureObject`, data.JwsSignatureObject.ValueString())
	}
	if !data.JweHeaderObject.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`JWEHeaderObject`, data.JweHeaderObject.ValueString())
	}
	if !data.JoseVerifyType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`JOSEVerifyType`, data.JoseVerifyType.ValueString())
	}
	if !data.JoseDecryptType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`JOSEDecryptType`, data.JoseDecryptType.ValueString())
	}
	if !data.SignatureIdentifier.IsNull() {
		var dataValues []string
		data.SignatureIdentifier.ElementsAs(ctx, &dataValues, false)
		for _, val := range dataValues {
			body, _ = sjson.Set(body, pathRoot+`SignatureIdentifier`+".-1", map[string]string{"value": val})
		}
	}
	if !data.RecipientIdentifier.IsNull() {
		var dataValues []string
		data.RecipientIdentifier.ElementsAs(ctx, &dataValues, false)
		for _, val := range dataValues {
			body, _ = sjson.Set(body, pathRoot+`RecipientIdentifier`+".-1", map[string]string{"value": val})
		}
	}
	if !data.SingleCertificate.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SingleCertificate`, data.SingleCertificate.ValueString())
	}
	if !data.SingleKey.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SingleKey`, data.SingleKey.ValueString())
	}
	if !data.SingleSskey.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SingleSSKey`, data.SingleSskey.ValueString())
	}
	if !data.JweDirectKeyObject.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`JWEDirectKeyObject`, data.JweDirectKeyObject.ValueString())
	}
	if !data.JwsVerifyStripSignature.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`JWSVerifyStripSignature`, tfutils.StringFromBool(data.JwsVerifyStripSignature, ""))
	}
	if !data.Asynchronous.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Asynchronous`, tfutils.StringFromBool(data.Asynchronous, ""))
	}
	if !data.Condition.IsNull() {
		var dataValues []DmCondition
		data.Condition.ElementsAs(ctx, &dataValues, false)
		for _, val := range dataValues {
			body, _ = sjson.SetRaw(body, pathRoot+`Condition`+".-1", val.ToBody(ctx, ""))
		}
	}
	if !data.ResultsMode.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ResultsMode`, data.ResultsMode.ValueString())
	}
	if !data.RetryCount.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`RetryCount`, data.RetryCount.ValueInt64())
	}
	if !data.RetryInterval.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`RetryInterval`, data.RetryInterval.ValueInt64())
	}
	if !data.MultipleOutputs.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MultipleOutputs`, tfutils.StringFromBool(data.MultipleOutputs, ""))
	}
	if !data.IteratorType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`IteratorType`, data.IteratorType.ValueString())
	}
	if !data.IteratorExpression.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`IteratorExpression`, data.IteratorExpression.ValueString())
	}
	if !data.IteratorCount.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`IteratorCount`, data.IteratorCount.ValueInt64())
	}
	if !data.LoopAction.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LoopAction`, data.LoopAction.ValueString())
	}
	if !data.AsyncAction.IsNull() {
		var dataValues []string
		data.AsyncAction.ElementsAs(ctx, &dataValues, false)
		for _, val := range dataValues {
			body, _ = sjson.Set(body, pathRoot+`AsyncAction`+".-1", map[string]string{"value": val})
		}
	}
	if !data.Timeout.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Timeout`, data.Timeout.ValueInt64())
	}
	if !data.WsdlPortQName.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`WSDLPortQName`, data.WsdlPortQName.ValueString())
	}
	if !data.WsdlOperationName.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`WSDLOperationName`, data.WsdlOperationName.ValueString())
	}
	if !data.WsdlMessageDirectionOrName.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`WSDLMessageDirectionOrName`, data.WsdlMessageDirectionOrName.ValueString())
	}
	if !data.WsdlAttachmentPart.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`WSDLAttachmentPart`, data.WsdlAttachmentPart.ValueString())
	}
	if !data.MethodRewriteType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MethodRewriteType`, data.MethodRewriteType.ValueString())
	}
	if !data.MethodType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MethodType`, data.MethodType.ValueString())
	}
	if !data.MethodType2.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MethodType2`, data.MethodType2.ValueString())
	}
	if !data.PolicyKey.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`PolicyKey`, data.PolicyKey.ValueString())
	}
	return body
}

func (data *StylePolicyAction) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `name`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Id = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Id = types.StringNull()
	}
	if value := res.Get(pathRoot + `UserSummary`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.UserSummary = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UserSummary = types.StringNull()
	}
	if value := res.Get(pathRoot + `Type`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Type = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Type = types.StringValue("xform")
	}
	if value := res.Get(pathRoot + `Input`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Input = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Input = types.StringNull()
	}
	if value := res.Get(pathRoot + `Transform`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Transform = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Transform = types.StringNull()
	}
	if value := res.Get(pathRoot + `ParseSettingsReference`); value.Exists() {
		data.ParseSettingsReference = &DmDynamicParseSettingsReference{}
		data.ParseSettingsReference.FromBody(ctx, "", value)
	} else {
		data.ParseSettingsReference = nil
	}
	if value := res.Get(pathRoot + `ParseMetricsResultType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ParseMetricsResultType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ParseMetricsResultType = types.StringValue("none")
	}
	if value := res.Get(pathRoot + `ParseMetricsResultLocation`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ParseMetricsResultLocation = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ParseMetricsResultLocation = types.StringNull()
	}
	if value := res.Get(pathRoot + `InputLanguage`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.InputLanguage = tfutils.ParseStringFromGJSON(value)
	} else {
		data.InputLanguage = types.StringValue("xml")
	}
	if value := res.Get(pathRoot + `DFDLInputRootName`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.DfdlInputRootName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.DfdlInputRootName = types.StringNull()
	}
	if value := res.Get(pathRoot + `InputDescriptor`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.InputDescriptor = tfutils.ParseStringFromGJSON(value)
	} else {
		data.InputDescriptor = types.StringNull()
	}
	if value := res.Get(pathRoot + `OutputDescriptor`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.OutputDescriptor = tfutils.ParseStringFromGJSON(value)
	} else {
		data.OutputDescriptor = types.StringNull()
	}
	if value := res.Get(pathRoot + `TransformLanguage`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.TransformLanguage = tfutils.ParseStringFromGJSON(value)
	} else {
		data.TransformLanguage = types.StringValue("none")
	}
	if value := res.Get(pathRoot + `OutputLanguage`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.OutputLanguage = tfutils.ParseStringFromGJSON(value)
	} else {
		data.OutputLanguage = types.StringValue("default")
	}
	if value := res.Get(pathRoot + `TxMap`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.TxMap = tfutils.ParseStringFromGJSON(value)
	} else {
		data.TxMap = types.StringNull()
	}
	if value := res.Get(pathRoot + `GatewayScriptLocation`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.GatewayScriptLocation = tfutils.ParseStringFromGJSON(value)
	} else {
		data.GatewayScriptLocation = types.StringNull()
	}
	if value := res.Get(pathRoot + `ActionDebug`); value.Exists() {
		data.ActionDebug = tfutils.BoolFromString(value.String())
	} else {
		data.ActionDebug = types.BoolNull()
	}
	if value := res.Get(pathRoot + `TxTopLevelMap`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.TxTopLevelMap = tfutils.ParseStringFromGJSON(value)
	} else {
		data.TxTopLevelMap = types.StringNull()
	}
	if value := res.Get(pathRoot + `TxMode`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.TxMode = tfutils.ParseStringFromGJSON(value)
	} else {
		data.TxMode = types.StringNull()
	}
	if value := res.Get(pathRoot + `TxAuditLog`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.TxAuditLog = tfutils.ParseStringFromGJSON(value)
	} else {
		data.TxAuditLog = types.StringNull()
	}
	if value := res.Get(pathRoot + `Output`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Output = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Output = types.StringNull()
	}
	if value := res.Get(pathRoot + `NoTranscodeUtf8`); value.Exists() {
		data.NoTranscodeUtf8 = tfutils.BoolFromString(value.String())
	} else {
		data.NoTranscodeUtf8 = types.BoolNull()
	}
	if value := res.Get(pathRoot + `NamedInOutLocationType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.NamedInOutLocationType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.NamedInOutLocationType = types.StringValue("default")
	}
	if value := res.Get(pathRoot + `NamedInputs`); value.Exists() {
		l := []DmNamedInOut{}
		if value := res.Get(`NamedInputs`); value.Exists() {
			for _, v := range value.Array() {
				item := DmNamedInOut{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.NamedInputs, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmNamedInOutObjectType}, l)
		} else {
			data.NamedInputs = types.ListNull(types.ObjectType{AttrTypes: DmNamedInOutObjectType})
		}
	} else {
		data.NamedInputs = types.ListNull(types.ObjectType{AttrTypes: DmNamedInOutObjectType})
	}
	if value := res.Get(pathRoot + `NamedOutputs`); value.Exists() {
		l := []DmNamedInOut{}
		if value := res.Get(`NamedOutputs`); value.Exists() {
			for _, v := range value.Array() {
				item := DmNamedInOut{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.NamedOutputs, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmNamedInOutObjectType}, l)
		} else {
			data.NamedOutputs = types.ListNull(types.ObjectType{AttrTypes: DmNamedInOutObjectType})
		}
	} else {
		data.NamedOutputs = types.ListNull(types.ObjectType{AttrTypes: DmNamedInOutObjectType})
	}
	if value := res.Get(pathRoot + `Destination`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Destination = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Destination = types.StringNull()
	}
	if value := res.Get(pathRoot + `SchemaURL`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SchemaUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SchemaUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `JSONSchemaURL`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.JsonSchemaUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.JsonSchemaUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `WsdlURL`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.WsdlUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.WsdlUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `Policy`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Policy = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Policy = types.StringNull()
	}
	if value := res.Get(pathRoot + `AAA`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Aaa = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Aaa = types.StringNull()
	}
	if value := res.Get(pathRoot + `DynamicSchema`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.DynamicSchema = tfutils.ParseStringFromGJSON(value)
	} else {
		data.DynamicSchema = types.StringNull()
	}
	if value := res.Get(pathRoot + `DynamicStylesheet`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.DynamicStylesheet = tfutils.ParseStringFromGJSON(value)
	} else {
		data.DynamicStylesheet = types.StringNull()
	}
	if value := res.Get(pathRoot + `InputConversion`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.InputConversion = tfutils.ParseStringFromGJSON(value)
	} else {
		data.InputConversion = types.StringNull()
	}
	if value := res.Get(pathRoot + `XPath`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Xpath = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Xpath = types.StringNull()
	}
	if value := res.Get(pathRoot + `Variable`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Variable = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Variable = types.StringNull()
	}
	if value := res.Get(pathRoot + `Value`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Value = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Value = types.StringNull()
	}
	if value := res.Get(pathRoot + `SSLClientConfigType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SslClientConfigType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SslClientConfigType = types.StringValue("client")
	}
	if value := res.Get(pathRoot + `SSLClientCred`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SslClientCred = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SslClientCred = types.StringNull()
	}
	if value := res.Get(pathRoot + `AttachmentURI`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.AttachmentUri = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AttachmentUri = types.StringNull()
	}
	if value := res.Get(pathRoot + `StylesheetParameters`); value.Exists() {
		l := []DmStylesheetParameter{}
		if value := res.Get(`StylesheetParameters`); value.Exists() {
			for _, v := range value.Array() {
				item := DmStylesheetParameter{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.StylesheetParameters, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmStylesheetParameterObjectType}, l)
		} else {
			data.StylesheetParameters = types.ListNull(types.ObjectType{AttrTypes: DmStylesheetParameterObjectType})
		}
	} else {
		data.StylesheetParameters = types.ListNull(types.ObjectType{AttrTypes: DmStylesheetParameterObjectType})
	}
	if value := res.Get(pathRoot + `ErrorMode`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ErrorMode = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ErrorMode = types.StringNull()
	}
	if value := res.Get(pathRoot + `ErrorInput`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ErrorInput = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ErrorInput = types.StringNull()
	}
	if value := res.Get(pathRoot + `ErrorOutput`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ErrorOutput = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ErrorOutput = types.StringNull()
	}
	if value := res.Get(pathRoot + `Rule`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Rule = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Rule = types.StringNull()
	}
	if value := res.Get(pathRoot + `OutputType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.OutputType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.OutputType = types.StringNull()
	}
	if value := res.Get(pathRoot + `LogLevel`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.LogLevel = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LogLevel = types.StringNull()
	}
	if value := res.Get(pathRoot + `LogType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.LogType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LogType = types.StringNull()
	}
	if value := res.Get(pathRoot + `Transactional`); value.Exists() {
		data.Transactional = tfutils.BoolFromString(value.String())
	} else {
		data.Transactional = types.BoolNull()
	}
	if value := res.Get(pathRoot + `CheckpointEvent`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.CheckpointEvent = tfutils.ParseStringFromGJSON(value)
	} else {
		data.CheckpointEvent = types.StringNull()
	}
	if value := res.Get(pathRoot + `SLMPolicy`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SlmPolicy = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SlmPolicy = types.StringNull()
	}
	if value := res.Get(pathRoot + `SQLDataSource`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SqlDataSource = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SqlDataSource = types.StringNull()
	}
	if value := res.Get(pathRoot + `SQLText`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SqlText = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SqlText = types.StringNull()
	}
	if value := res.Get(pathRoot + `SOAPValidation`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SoapValidation = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SoapValidation = types.StringValue("body")
	}
	if value := res.Get(pathRoot + `SQLSourceType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SqlSourceType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SqlSourceType = types.StringValue("static")
	}
	if value := res.Get(pathRoot + `JOSESerializationType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.JoseSerializationType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.JoseSerializationType = types.StringNull()
	}
	if value := res.Get(pathRoot + `JWEEncAlgorithm`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.JweEncAlgorithm = tfutils.ParseStringFromGJSON(value)
	} else {
		data.JweEncAlgorithm = types.StringNull()
	}
	if value := res.Get(pathRoot + `JWSSignatureObject`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.JwsSignatureObject = tfutils.ParseStringFromGJSON(value)
	} else {
		data.JwsSignatureObject = types.StringNull()
	}
	if value := res.Get(pathRoot + `JWEHeaderObject`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.JweHeaderObject = tfutils.ParseStringFromGJSON(value)
	} else {
		data.JweHeaderObject = types.StringNull()
	}
	if value := res.Get(pathRoot + `JOSEVerifyType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.JoseVerifyType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.JoseVerifyType = types.StringNull()
	}
	if value := res.Get(pathRoot + `JOSEDecryptType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.JoseDecryptType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.JoseDecryptType = types.StringNull()
	}
	if value := res.Get(pathRoot + `SignatureIdentifier`); value.Exists() {
		data.SignatureIdentifier = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.SignatureIdentifier = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `RecipientIdentifier`); value.Exists() {
		data.RecipientIdentifier = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.RecipientIdentifier = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `SingleCertificate`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SingleCertificate = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SingleCertificate = types.StringNull()
	}
	if value := res.Get(pathRoot + `SingleKey`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SingleKey = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SingleKey = types.StringNull()
	}
	if value := res.Get(pathRoot + `SingleSSKey`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SingleSskey = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SingleSskey = types.StringNull()
	}
	if value := res.Get(pathRoot + `JWEDirectKeyObject`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.JweDirectKeyObject = tfutils.ParseStringFromGJSON(value)
	} else {
		data.JweDirectKeyObject = types.StringNull()
	}
	if value := res.Get(pathRoot + `JWSVerifyStripSignature`); value.Exists() {
		data.JwsVerifyStripSignature = tfutils.BoolFromString(value.String())
	} else {
		data.JwsVerifyStripSignature = types.BoolNull()
	}
	if value := res.Get(pathRoot + `Asynchronous`); value.Exists() {
		data.Asynchronous = tfutils.BoolFromString(value.String())
	} else {
		data.Asynchronous = types.BoolNull()
	}
	if value := res.Get(pathRoot + `Condition`); value.Exists() {
		l := []DmCondition{}
		if value := res.Get(`Condition`); value.Exists() {
			for _, v := range value.Array() {
				item := DmCondition{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.Condition, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmConditionObjectType}, l)
		} else {
			data.Condition = types.ListNull(types.ObjectType{AttrTypes: DmConditionObjectType})
		}
	} else {
		data.Condition = types.ListNull(types.ObjectType{AttrTypes: DmConditionObjectType})
	}
	if value := res.Get(pathRoot + `ResultsMode`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ResultsMode = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ResultsMode = types.StringValue("first-available")
	}
	if value := res.Get(pathRoot + `RetryCount`); value.Exists() {
		data.RetryCount = types.Int64Value(value.Int())
	} else {
		data.RetryCount = types.Int64Null()
	}
	if value := res.Get(pathRoot + `RetryInterval`); value.Exists() {
		data.RetryInterval = types.Int64Value(value.Int())
	} else {
		data.RetryInterval = types.Int64Value(1000)
	}
	if value := res.Get(pathRoot + `MultipleOutputs`); value.Exists() {
		data.MultipleOutputs = tfutils.BoolFromString(value.String())
	} else {
		data.MultipleOutputs = types.BoolNull()
	}
	if value := res.Get(pathRoot + `IteratorType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.IteratorType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.IteratorType = types.StringValue("XPATH")
	}
	if value := res.Get(pathRoot + `IteratorExpression`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.IteratorExpression = tfutils.ParseStringFromGJSON(value)
	} else {
		data.IteratorExpression = types.StringNull()
	}
	if value := res.Get(pathRoot + `IteratorCount`); value.Exists() {
		data.IteratorCount = types.Int64Value(value.Int())
	} else {
		data.IteratorCount = types.Int64Null()
	}
	if value := res.Get(pathRoot + `LoopAction`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.LoopAction = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LoopAction = types.StringNull()
	}
	if value := res.Get(pathRoot + `AsyncAction`); value.Exists() {
		data.AsyncAction = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.AsyncAction = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `Timeout`); value.Exists() {
		data.Timeout = types.Int64Value(value.Int())
	} else {
		data.Timeout = types.Int64Null()
	}
	if value := res.Get(pathRoot + `WSDLPortQName`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.WsdlPortQName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.WsdlPortQName = types.StringNull()
	}
	if value := res.Get(pathRoot + `WSDLOperationName`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.WsdlOperationName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.WsdlOperationName = types.StringNull()
	}
	if value := res.Get(pathRoot + `WSDLMessageDirectionOrName`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.WsdlMessageDirectionOrName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.WsdlMessageDirectionOrName = types.StringNull()
	}
	if value := res.Get(pathRoot + `WSDLAttachmentPart`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.WsdlAttachmentPart = tfutils.ParseStringFromGJSON(value)
	} else {
		data.WsdlAttachmentPart = types.StringNull()
	}
	if value := res.Get(pathRoot + `MethodRewriteType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.MethodRewriteType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.MethodRewriteType = types.StringValue("GET")
	}
	if value := res.Get(pathRoot + `MethodType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.MethodType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.MethodType = types.StringValue("POST")
	}
	if value := res.Get(pathRoot + `MethodType2`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.MethodType2 = tfutils.ParseStringFromGJSON(value)
	} else {
		data.MethodType2 = types.StringValue("POST")
	}
	if value := res.Get(pathRoot + `PolicyKey`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.PolicyKey = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PolicyKey = types.StringNull()
	}
}

func (data *StylePolicyAction) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `name`); value.Exists() && !data.Id.IsNull() {
		data.Id = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Id = types.StringNull()
	}
	if value := res.Get(pathRoot + `UserSummary`); value.Exists() && !data.UserSummary.IsNull() {
		data.UserSummary = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UserSummary = types.StringNull()
	}
	if value := res.Get(pathRoot + `Type`); value.Exists() && !data.Type.IsNull() {
		data.Type = tfutils.ParseStringFromGJSON(value)
	} else if data.Type.ValueString() != "xform" {
		data.Type = types.StringNull()
	}
	if value := res.Get(pathRoot + `Input`); value.Exists() && !data.Input.IsNull() {
		data.Input = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Input = types.StringNull()
	}
	if value := res.Get(pathRoot + `Transform`); value.Exists() && !data.Transform.IsNull() {
		data.Transform = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Transform = types.StringNull()
	}
	if value := res.Get(pathRoot + `ParseSettingsReference`); value.Exists() {
		data.ParseSettingsReference.UpdateFromBody(ctx, "", value)
	} else {
		data.ParseSettingsReference = nil
	}
	if value := res.Get(pathRoot + `ParseMetricsResultType`); value.Exists() && !data.ParseMetricsResultType.IsNull() {
		data.ParseMetricsResultType = tfutils.ParseStringFromGJSON(value)
	} else if data.ParseMetricsResultType.ValueString() != "none" {
		data.ParseMetricsResultType = types.StringNull()
	}
	if value := res.Get(pathRoot + `ParseMetricsResultLocation`); value.Exists() && !data.ParseMetricsResultLocation.IsNull() {
		data.ParseMetricsResultLocation = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ParseMetricsResultLocation = types.StringNull()
	}
	if value := res.Get(pathRoot + `InputLanguage`); value.Exists() && !data.InputLanguage.IsNull() {
		data.InputLanguage = tfutils.ParseStringFromGJSON(value)
	} else if data.InputLanguage.ValueString() != "xml" {
		data.InputLanguage = types.StringNull()
	}
	if value := res.Get(pathRoot + `DFDLInputRootName`); value.Exists() && !data.DfdlInputRootName.IsNull() {
		data.DfdlInputRootName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.DfdlInputRootName = types.StringNull()
	}
	if value := res.Get(pathRoot + `InputDescriptor`); value.Exists() && !data.InputDescriptor.IsNull() {
		data.InputDescriptor = tfutils.ParseStringFromGJSON(value)
	} else {
		data.InputDescriptor = types.StringNull()
	}
	if value := res.Get(pathRoot + `OutputDescriptor`); value.Exists() && !data.OutputDescriptor.IsNull() {
		data.OutputDescriptor = tfutils.ParseStringFromGJSON(value)
	} else {
		data.OutputDescriptor = types.StringNull()
	}
	if value := res.Get(pathRoot + `TransformLanguage`); value.Exists() && !data.TransformLanguage.IsNull() {
		data.TransformLanguage = tfutils.ParseStringFromGJSON(value)
	} else if data.TransformLanguage.ValueString() != "none" {
		data.TransformLanguage = types.StringNull()
	}
	if value := res.Get(pathRoot + `OutputLanguage`); value.Exists() && !data.OutputLanguage.IsNull() {
		data.OutputLanguage = tfutils.ParseStringFromGJSON(value)
	} else if data.OutputLanguage.ValueString() != "default" {
		data.OutputLanguage = types.StringNull()
	}
	if value := res.Get(pathRoot + `TxMap`); value.Exists() && !data.TxMap.IsNull() {
		data.TxMap = tfutils.ParseStringFromGJSON(value)
	} else {
		data.TxMap = types.StringNull()
	}
	if value := res.Get(pathRoot + `GatewayScriptLocation`); value.Exists() && !data.GatewayScriptLocation.IsNull() {
		data.GatewayScriptLocation = tfutils.ParseStringFromGJSON(value)
	} else {
		data.GatewayScriptLocation = types.StringNull()
	}
	if value := res.Get(pathRoot + `ActionDebug`); value.Exists() && !data.ActionDebug.IsNull() {
		data.ActionDebug = tfutils.BoolFromString(value.String())
	} else if data.ActionDebug.ValueBool() {
		data.ActionDebug = types.BoolNull()
	}
	if value := res.Get(pathRoot + `TxTopLevelMap`); value.Exists() && !data.TxTopLevelMap.IsNull() {
		data.TxTopLevelMap = tfutils.ParseStringFromGJSON(value)
	} else {
		data.TxTopLevelMap = types.StringNull()
	}
	if value := res.Get(pathRoot + `TxMode`); value.Exists() && !data.TxMode.IsNull() {
		data.TxMode = tfutils.ParseStringFromGJSON(value)
	} else {
		data.TxMode = types.StringNull()
	}
	if value := res.Get(pathRoot + `TxAuditLog`); value.Exists() && !data.TxAuditLog.IsNull() {
		data.TxAuditLog = tfutils.ParseStringFromGJSON(value)
	} else {
		data.TxAuditLog = types.StringNull()
	}
	if value := res.Get(pathRoot + `Output`); value.Exists() && !data.Output.IsNull() {
		data.Output = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Output = types.StringNull()
	}
	if value := res.Get(pathRoot + `NoTranscodeUtf8`); value.Exists() && !data.NoTranscodeUtf8.IsNull() {
		data.NoTranscodeUtf8 = tfutils.BoolFromString(value.String())
	} else if data.NoTranscodeUtf8.ValueBool() {
		data.NoTranscodeUtf8 = types.BoolNull()
	}
	if value := res.Get(pathRoot + `NamedInOutLocationType`); value.Exists() && !data.NamedInOutLocationType.IsNull() {
		data.NamedInOutLocationType = tfutils.ParseStringFromGJSON(value)
	} else if data.NamedInOutLocationType.ValueString() != "default" {
		data.NamedInOutLocationType = types.StringNull()
	}
	if value := res.Get(pathRoot + `NamedInputs`); value.Exists() && !data.NamedInputs.IsNull() {
		l := []DmNamedInOut{}
		for _, v := range value.Array() {
			item := DmNamedInOut{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
		if len(l) > 0 {
			data.NamedInputs, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmNamedInOutObjectType}, l)
		} else {
			data.NamedInputs = types.ListNull(types.ObjectType{AttrTypes: DmNamedInOutObjectType})
		}
	} else {
		data.NamedInputs = types.ListNull(types.ObjectType{AttrTypes: DmNamedInOutObjectType})
	}
	if value := res.Get(pathRoot + `NamedOutputs`); value.Exists() && !data.NamedOutputs.IsNull() {
		l := []DmNamedInOut{}
		for _, v := range value.Array() {
			item := DmNamedInOut{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
		if len(l) > 0 {
			data.NamedOutputs, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmNamedInOutObjectType}, l)
		} else {
			data.NamedOutputs = types.ListNull(types.ObjectType{AttrTypes: DmNamedInOutObjectType})
		}
	} else {
		data.NamedOutputs = types.ListNull(types.ObjectType{AttrTypes: DmNamedInOutObjectType})
	}
	if value := res.Get(pathRoot + `Destination`); value.Exists() && !data.Destination.IsNull() {
		data.Destination = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Destination = types.StringNull()
	}
	if value := res.Get(pathRoot + `SchemaURL`); value.Exists() && !data.SchemaUrl.IsNull() {
		data.SchemaUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SchemaUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `JSONSchemaURL`); value.Exists() && !data.JsonSchemaUrl.IsNull() {
		data.JsonSchemaUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.JsonSchemaUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `WsdlURL`); value.Exists() && !data.WsdlUrl.IsNull() {
		data.WsdlUrl = tfutils.ParseStringFromGJSON(value)
	} else {
		data.WsdlUrl = types.StringNull()
	}
	if value := res.Get(pathRoot + `Policy`); value.Exists() && !data.Policy.IsNull() {
		data.Policy = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Policy = types.StringNull()
	}
	if value := res.Get(pathRoot + `AAA`); value.Exists() && !data.Aaa.IsNull() {
		data.Aaa = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Aaa = types.StringNull()
	}
	if value := res.Get(pathRoot + `DynamicSchema`); value.Exists() && !data.DynamicSchema.IsNull() {
		data.DynamicSchema = tfutils.ParseStringFromGJSON(value)
	} else {
		data.DynamicSchema = types.StringNull()
	}
	if value := res.Get(pathRoot + `DynamicStylesheet`); value.Exists() && !data.DynamicStylesheet.IsNull() {
		data.DynamicStylesheet = tfutils.ParseStringFromGJSON(value)
	} else {
		data.DynamicStylesheet = types.StringNull()
	}
	if value := res.Get(pathRoot + `InputConversion`); value.Exists() && !data.InputConversion.IsNull() {
		data.InputConversion = tfutils.ParseStringFromGJSON(value)
	} else {
		data.InputConversion = types.StringNull()
	}
	if value := res.Get(pathRoot + `XPath`); value.Exists() && !data.Xpath.IsNull() {
		data.Xpath = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Xpath = types.StringNull()
	}
	if value := res.Get(pathRoot + `Variable`); value.Exists() && !data.Variable.IsNull() {
		data.Variable = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Variable = types.StringNull()
	}
	if value := res.Get(pathRoot + `Value`); value.Exists() && !data.Value.IsNull() {
		data.Value = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Value = types.StringNull()
	}
	if value := res.Get(pathRoot + `SSLClientConfigType`); value.Exists() && !data.SslClientConfigType.IsNull() {
		data.SslClientConfigType = tfutils.ParseStringFromGJSON(value)
	} else if data.SslClientConfigType.ValueString() != "client" {
		data.SslClientConfigType = types.StringNull()
	}
	if value := res.Get(pathRoot + `SSLClientCred`); value.Exists() && !data.SslClientCred.IsNull() {
		data.SslClientCred = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SslClientCred = types.StringNull()
	}
	if value := res.Get(pathRoot + `AttachmentURI`); value.Exists() && !data.AttachmentUri.IsNull() {
		data.AttachmentUri = tfutils.ParseStringFromGJSON(value)
	} else {
		data.AttachmentUri = types.StringNull()
	}
	if value := res.Get(pathRoot + `StylesheetParameters`); value.Exists() && !data.StylesheetParameters.IsNull() {
		l := []DmStylesheetParameter{}
		for _, v := range value.Array() {
			item := DmStylesheetParameter{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
		if len(l) > 0 {
			data.StylesheetParameters, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmStylesheetParameterObjectType}, l)
		} else {
			data.StylesheetParameters = types.ListNull(types.ObjectType{AttrTypes: DmStylesheetParameterObjectType})
		}
	} else {
		data.StylesheetParameters = types.ListNull(types.ObjectType{AttrTypes: DmStylesheetParameterObjectType})
	}
	if value := res.Get(pathRoot + `ErrorMode`); value.Exists() && !data.ErrorMode.IsNull() {
		data.ErrorMode = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ErrorMode = types.StringNull()
	}
	if value := res.Get(pathRoot + `ErrorInput`); value.Exists() && !data.ErrorInput.IsNull() {
		data.ErrorInput = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ErrorInput = types.StringNull()
	}
	if value := res.Get(pathRoot + `ErrorOutput`); value.Exists() && !data.ErrorOutput.IsNull() {
		data.ErrorOutput = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ErrorOutput = types.StringNull()
	}
	if value := res.Get(pathRoot + `Rule`); value.Exists() && !data.Rule.IsNull() {
		data.Rule = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Rule = types.StringNull()
	}
	if value := res.Get(pathRoot + `OutputType`); value.Exists() && !data.OutputType.IsNull() {
		data.OutputType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.OutputType = types.StringNull()
	}
	if value := res.Get(pathRoot + `LogLevel`); value.Exists() && !data.LogLevel.IsNull() {
		data.LogLevel = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LogLevel = types.StringNull()
	}
	if value := res.Get(pathRoot + `LogType`); value.Exists() && !data.LogType.IsNull() {
		data.LogType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LogType = types.StringNull()
	}
	if value := res.Get(pathRoot + `Transactional`); value.Exists() && !data.Transactional.IsNull() {
		data.Transactional = tfutils.BoolFromString(value.String())
	} else if data.Transactional.ValueBool() {
		data.Transactional = types.BoolNull()
	}
	if value := res.Get(pathRoot + `CheckpointEvent`); value.Exists() && !data.CheckpointEvent.IsNull() {
		data.CheckpointEvent = tfutils.ParseStringFromGJSON(value)
	} else {
		data.CheckpointEvent = types.StringNull()
	}
	if value := res.Get(pathRoot + `SLMPolicy`); value.Exists() && !data.SlmPolicy.IsNull() {
		data.SlmPolicy = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SlmPolicy = types.StringNull()
	}
	if value := res.Get(pathRoot + `SQLDataSource`); value.Exists() && !data.SqlDataSource.IsNull() {
		data.SqlDataSource = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SqlDataSource = types.StringNull()
	}
	if value := res.Get(pathRoot + `SQLText`); value.Exists() && !data.SqlText.IsNull() {
		data.SqlText = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SqlText = types.StringNull()
	}
	if value := res.Get(pathRoot + `SOAPValidation`); value.Exists() && !data.SoapValidation.IsNull() {
		data.SoapValidation = tfutils.ParseStringFromGJSON(value)
	} else if data.SoapValidation.ValueString() != "body" {
		data.SoapValidation = types.StringNull()
	}
	if value := res.Get(pathRoot + `SQLSourceType`); value.Exists() && !data.SqlSourceType.IsNull() {
		data.SqlSourceType = tfutils.ParseStringFromGJSON(value)
	} else if data.SqlSourceType.ValueString() != "static" {
		data.SqlSourceType = types.StringNull()
	}
	if value := res.Get(pathRoot + `JOSESerializationType`); value.Exists() && !data.JoseSerializationType.IsNull() {
		data.JoseSerializationType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.JoseSerializationType = types.StringNull()
	}
	if value := res.Get(pathRoot + `JWEEncAlgorithm`); value.Exists() && !data.JweEncAlgorithm.IsNull() {
		data.JweEncAlgorithm = tfutils.ParseStringFromGJSON(value)
	} else {
		data.JweEncAlgorithm = types.StringNull()
	}
	if value := res.Get(pathRoot + `JWSSignatureObject`); value.Exists() && !data.JwsSignatureObject.IsNull() {
		data.JwsSignatureObject = tfutils.ParseStringFromGJSON(value)
	} else {
		data.JwsSignatureObject = types.StringNull()
	}
	if value := res.Get(pathRoot + `JWEHeaderObject`); value.Exists() && !data.JweHeaderObject.IsNull() {
		data.JweHeaderObject = tfutils.ParseStringFromGJSON(value)
	} else {
		data.JweHeaderObject = types.StringNull()
	}
	if value := res.Get(pathRoot + `JOSEVerifyType`); value.Exists() && !data.JoseVerifyType.IsNull() {
		data.JoseVerifyType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.JoseVerifyType = types.StringNull()
	}
	if value := res.Get(pathRoot + `JOSEDecryptType`); value.Exists() && !data.JoseDecryptType.IsNull() {
		data.JoseDecryptType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.JoseDecryptType = types.StringNull()
	}
	if value := res.Get(pathRoot + `SignatureIdentifier`); value.Exists() && !data.SignatureIdentifier.IsNull() {
		data.SignatureIdentifier = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.SignatureIdentifier = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `RecipientIdentifier`); value.Exists() && !data.RecipientIdentifier.IsNull() {
		data.RecipientIdentifier = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.RecipientIdentifier = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `SingleCertificate`); value.Exists() && !data.SingleCertificate.IsNull() {
		data.SingleCertificate = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SingleCertificate = types.StringNull()
	}
	if value := res.Get(pathRoot + `SingleKey`); value.Exists() && !data.SingleKey.IsNull() {
		data.SingleKey = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SingleKey = types.StringNull()
	}
	if value := res.Get(pathRoot + `SingleSSKey`); value.Exists() && !data.SingleSskey.IsNull() {
		data.SingleSskey = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SingleSskey = types.StringNull()
	}
	if value := res.Get(pathRoot + `JWEDirectKeyObject`); value.Exists() && !data.JweDirectKeyObject.IsNull() {
		data.JweDirectKeyObject = tfutils.ParseStringFromGJSON(value)
	} else {
		data.JweDirectKeyObject = types.StringNull()
	}
	if value := res.Get(pathRoot + `JWSVerifyStripSignature`); value.Exists() && !data.JwsVerifyStripSignature.IsNull() {
		data.JwsVerifyStripSignature = tfutils.BoolFromString(value.String())
	} else if !data.JwsVerifyStripSignature.ValueBool() {
		data.JwsVerifyStripSignature = types.BoolNull()
	}
	if value := res.Get(pathRoot + `Asynchronous`); value.Exists() && !data.Asynchronous.IsNull() {
		data.Asynchronous = tfutils.BoolFromString(value.String())
	} else if data.Asynchronous.ValueBool() {
		data.Asynchronous = types.BoolNull()
	}
	if value := res.Get(pathRoot + `Condition`); value.Exists() && !data.Condition.IsNull() {
		l := []DmCondition{}
		for _, v := range value.Array() {
			item := DmCondition{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
		if len(l) > 0 {
			data.Condition, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmConditionObjectType}, l)
		} else {
			data.Condition = types.ListNull(types.ObjectType{AttrTypes: DmConditionObjectType})
		}
	} else {
		data.Condition = types.ListNull(types.ObjectType{AttrTypes: DmConditionObjectType})
	}
	if value := res.Get(pathRoot + `ResultsMode`); value.Exists() && !data.ResultsMode.IsNull() {
		data.ResultsMode = tfutils.ParseStringFromGJSON(value)
	} else if data.ResultsMode.ValueString() != "first-available" {
		data.ResultsMode = types.StringNull()
	}
	if value := res.Get(pathRoot + `RetryCount`); value.Exists() && !data.RetryCount.IsNull() {
		data.RetryCount = types.Int64Value(value.Int())
	} else {
		data.RetryCount = types.Int64Null()
	}
	if value := res.Get(pathRoot + `RetryInterval`); value.Exists() && !data.RetryInterval.IsNull() {
		data.RetryInterval = types.Int64Value(value.Int())
	} else if data.RetryInterval.ValueInt64() != 1000 {
		data.RetryInterval = types.Int64Null()
	}
	if value := res.Get(pathRoot + `MultipleOutputs`); value.Exists() && !data.MultipleOutputs.IsNull() {
		data.MultipleOutputs = tfutils.BoolFromString(value.String())
	} else if data.MultipleOutputs.ValueBool() {
		data.MultipleOutputs = types.BoolNull()
	}
	if value := res.Get(pathRoot + `IteratorType`); value.Exists() && !data.IteratorType.IsNull() {
		data.IteratorType = tfutils.ParseStringFromGJSON(value)
	} else if data.IteratorType.ValueString() != "XPATH" {
		data.IteratorType = types.StringNull()
	}
	if value := res.Get(pathRoot + `IteratorExpression`); value.Exists() && !data.IteratorExpression.IsNull() {
		data.IteratorExpression = tfutils.ParseStringFromGJSON(value)
	} else {
		data.IteratorExpression = types.StringNull()
	}
	if value := res.Get(pathRoot + `IteratorCount`); value.Exists() && !data.IteratorCount.IsNull() {
		data.IteratorCount = types.Int64Value(value.Int())
	} else {
		data.IteratorCount = types.Int64Null()
	}
	if value := res.Get(pathRoot + `LoopAction`); value.Exists() && !data.LoopAction.IsNull() {
		data.LoopAction = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LoopAction = types.StringNull()
	}
	if value := res.Get(pathRoot + `AsyncAction`); value.Exists() && !data.AsyncAction.IsNull() {
		data.AsyncAction = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.AsyncAction = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `Timeout`); value.Exists() && !data.Timeout.IsNull() {
		data.Timeout = types.Int64Value(value.Int())
	} else {
		data.Timeout = types.Int64Null()
	}
	if value := res.Get(pathRoot + `WSDLPortQName`); value.Exists() && !data.WsdlPortQName.IsNull() {
		data.WsdlPortQName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.WsdlPortQName = types.StringNull()
	}
	if value := res.Get(pathRoot + `WSDLOperationName`); value.Exists() && !data.WsdlOperationName.IsNull() {
		data.WsdlOperationName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.WsdlOperationName = types.StringNull()
	}
	if value := res.Get(pathRoot + `WSDLMessageDirectionOrName`); value.Exists() && !data.WsdlMessageDirectionOrName.IsNull() {
		data.WsdlMessageDirectionOrName = tfutils.ParseStringFromGJSON(value)
	} else {
		data.WsdlMessageDirectionOrName = types.StringNull()
	}
	if value := res.Get(pathRoot + `WSDLAttachmentPart`); value.Exists() && !data.WsdlAttachmentPart.IsNull() {
		data.WsdlAttachmentPart = tfutils.ParseStringFromGJSON(value)
	} else {
		data.WsdlAttachmentPart = types.StringNull()
	}
	if value := res.Get(pathRoot + `MethodRewriteType`); value.Exists() && !data.MethodRewriteType.IsNull() {
		data.MethodRewriteType = tfutils.ParseStringFromGJSON(value)
	} else if data.MethodRewriteType.ValueString() != "GET" {
		data.MethodRewriteType = types.StringNull()
	}
	if value := res.Get(pathRoot + `MethodType`); value.Exists() && !data.MethodType.IsNull() {
		data.MethodType = tfutils.ParseStringFromGJSON(value)
	} else if data.MethodType.ValueString() != "POST" {
		data.MethodType = types.StringNull()
	}
	if value := res.Get(pathRoot + `MethodType2`); value.Exists() && !data.MethodType2.IsNull() {
		data.MethodType2 = tfutils.ParseStringFromGJSON(value)
	} else if data.MethodType2.ValueString() != "POST" {
		data.MethodType2 = types.StringNull()
	}
	if value := res.Get(pathRoot + `PolicyKey`); value.Exists() && !data.PolicyKey.IsNull() {
		data.PolicyKey = tfutils.ParseStringFromGJSON(value)
	} else {
		data.PolicyKey = types.StringNull()
	}
}

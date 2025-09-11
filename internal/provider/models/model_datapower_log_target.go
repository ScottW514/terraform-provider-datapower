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

type LogTarget struct {
	Id                      types.String                `tfsdk:"id"`
	AppDomain               types.String                `tfsdk:"app_domain"`
	Backup                  types.String                `tfsdk:"backup"`
	LogEvents               types.List                  `tfsdk:"log_events"`
	UserSummary             types.String                `tfsdk:"user_summary"`
	Type                    types.String                `tfsdk:"type"`
	Priority                types.String                `tfsdk:"priority"`
	SoapVersion             types.String                `tfsdk:"soap_version"`
	Format                  types.String                `tfsdk:"format"`
	TimestampFormat         types.String                `tfsdk:"timestamp_format"`
	FixedFormat             types.Bool                  `tfsdk:"fixed_format"`
	LocalIdentifier         types.String                `tfsdk:"local_identifier"`
	EmailAddress            types.String                `tfsdk:"email_address"`
	SenderAddress           types.String                `tfsdk:"sender_address"`
	SmtpDomain              types.String                `tfsdk:"smtp_domain"`
	Size                    types.Int64                 `tfsdk:"size"`
	Url                     types.String                `tfsdk:"url"`
	NfSmOunt                types.String                `tfsdk:"nf_sm_ount"`
	LocalFile               types.String                `tfsdk:"local_file"`
	NfsFile                 types.String                `tfsdk:"nfs_file"`
	ArchiveMode             types.String                `tfsdk:"archive_mode"`
	UploadMethod            types.String                `tfsdk:"upload_method"`
	Rotate                  types.Int64                 `tfsdk:"rotate"`
	UseAnsiColor            types.Bool                  `tfsdk:"use_ansi_color"`
	RemoteAddress           types.String                `tfsdk:"remote_address"`
	RemotePort              types.Int64                 `tfsdk:"remote_port"`
	RemoteLogin             types.String                `tfsdk:"remote_login"`
	RemotePasswordWo        types.String                `tfsdk:"remote_password_wo"`
	RemotePasswordWoVersion types.Int64                 `tfsdk:"remote_password_wo_version"`
	RemoteDirectory         types.String                `tfsdk:"remote_directory"`
	LocalAddress            types.String                `tfsdk:"local_address"`
	SyslogFacility          types.String                `tfsdk:"syslog_facility"`
	RateLimit               types.Int64                 `tfsdk:"rate_limit"`
	MaxConnections          types.Int64                 `tfsdk:"max_connections"`
	ConnectTimeout          types.Int64                 `tfsdk:"connect_timeout"`
	IdleTimeout             types.Int64                 `tfsdk:"idle_timeout"`
	ActiveTimeout           types.Int64                 `tfsdk:"active_timeout"`
	FeedbackDetection       types.Bool                  `tfsdk:"feedback_detection"`
	LogEventCode            types.List                  `tfsdk:"log_event_code"`
	LogEventFilter          types.List                  `tfsdk:"log_event_filter"`
	LogObjects              types.List                  `tfsdk:"log_objects"`
	LogIpFilter             types.List                  `tfsdk:"log_ip_filter"`
	LogTriggers             types.List                  `tfsdk:"log_triggers"`
	SslClientProfile        types.String                `tfsdk:"ssl_client_profile"`
	SslClientConfigType     types.String                `tfsdk:"ssl_client_config_type"`
	RetryInterval           types.Int64                 `tfsdk:"retry_interval"`
	RetryAttempts           types.Int64                 `tfsdk:"retry_attempts"`
	LongRetryInterval       types.Int64                 `tfsdk:"long_retry_interval"`
	LogPrecision            types.String                `tfsdk:"log_precision"`
	EventBufferSize         types.String                `tfsdk:"event_buffer_size"`
	DependencyActions       []*actions.DependencyAction `tfsdk:"dependency_actions"`
}
type LogTargetWO struct {
	Id                  types.String                `tfsdk:"id"`
	AppDomain           types.String                `tfsdk:"app_domain"`
	Backup              types.String                `tfsdk:"backup"`
	LogEvents           types.List                  `tfsdk:"log_events"`
	UserSummary         types.String                `tfsdk:"user_summary"`
	Type                types.String                `tfsdk:"type"`
	Priority            types.String                `tfsdk:"priority"`
	SoapVersion         types.String                `tfsdk:"soap_version"`
	Format              types.String                `tfsdk:"format"`
	TimestampFormat     types.String                `tfsdk:"timestamp_format"`
	FixedFormat         types.Bool                  `tfsdk:"fixed_format"`
	LocalIdentifier     types.String                `tfsdk:"local_identifier"`
	EmailAddress        types.String                `tfsdk:"email_address"`
	SenderAddress       types.String                `tfsdk:"sender_address"`
	SmtpDomain          types.String                `tfsdk:"smtp_domain"`
	Size                types.Int64                 `tfsdk:"size"`
	Url                 types.String                `tfsdk:"url"`
	NfSmOunt            types.String                `tfsdk:"nf_sm_ount"`
	LocalFile           types.String                `tfsdk:"local_file"`
	NfsFile             types.String                `tfsdk:"nfs_file"`
	ArchiveMode         types.String                `tfsdk:"archive_mode"`
	UploadMethod        types.String                `tfsdk:"upload_method"`
	Rotate              types.Int64                 `tfsdk:"rotate"`
	UseAnsiColor        types.Bool                  `tfsdk:"use_ansi_color"`
	RemoteAddress       types.String                `tfsdk:"remote_address"`
	RemotePort          types.Int64                 `tfsdk:"remote_port"`
	RemoteLogin         types.String                `tfsdk:"remote_login"`
	RemoteDirectory     types.String                `tfsdk:"remote_directory"`
	LocalAddress        types.String                `tfsdk:"local_address"`
	SyslogFacility      types.String                `tfsdk:"syslog_facility"`
	RateLimit           types.Int64                 `tfsdk:"rate_limit"`
	MaxConnections      types.Int64                 `tfsdk:"max_connections"`
	ConnectTimeout      types.Int64                 `tfsdk:"connect_timeout"`
	IdleTimeout         types.Int64                 `tfsdk:"idle_timeout"`
	ActiveTimeout       types.Int64                 `tfsdk:"active_timeout"`
	FeedbackDetection   types.Bool                  `tfsdk:"feedback_detection"`
	LogEventCode        types.List                  `tfsdk:"log_event_code"`
	LogEventFilter      types.List                  `tfsdk:"log_event_filter"`
	LogObjects          types.List                  `tfsdk:"log_objects"`
	LogIpFilter         types.List                  `tfsdk:"log_ip_filter"`
	LogTriggers         types.List                  `tfsdk:"log_triggers"`
	SslClientProfile    types.String                `tfsdk:"ssl_client_profile"`
	SslClientConfigType types.String                `tfsdk:"ssl_client_config_type"`
	RetryInterval       types.Int64                 `tfsdk:"retry_interval"`
	RetryAttempts       types.Int64                 `tfsdk:"retry_attempts"`
	LongRetryInterval   types.Int64                 `tfsdk:"long_retry_interval"`
	LogPrecision        types.String                `tfsdk:"log_precision"`
	EventBufferSize     types.String                `tfsdk:"event_buffer_size"`
	DependencyActions   []*actions.DependencyAction `tfsdk:"dependency_actions"`
}

var LogTargetLocalIdentifierCondVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "format",
			AttrType:    "String",
			AttrDefault: "xml",
			Value:       []string{"cbe"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "type",
			AttrType:    "String",
			AttrDefault: "file",
			Value:       []string{"syslog", "smtp"},
		},
		{
			Evaluation: "logical-and",
			Conditions: []validators.Evaluation{
				{
					Evaluation:  "property-value-in-list",
					Attribute:   "type",
					AttrType:    "String",
					AttrDefault: "file",
					Value:       []string{"file"},
				},
				{
					Evaluation:  "property-value-in-list",
					Attribute:   "upload_method",
					AttrType:    "String",
					AttrDefault: "ftp",
					Value:       []string{"smtp"},
				},
			},
		},
	},
}
var LogTargetEmailAddressCondVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "type",
			AttrType:    "String",
			AttrDefault: "file",
			Value:       []string{"smtp"},
		},
		{
			Evaluation: "logical-and",
			Conditions: []validators.Evaluation{
				{
					Evaluation:  "property-value-in-list",
					Attribute:   "type",
					AttrType:    "String",
					AttrDefault: "file",
					Value:       []string{"file"},
				},
				{
					Evaluation:  "property-value-in-list",
					Attribute:   "upload_method",
					AttrType:    "String",
					AttrDefault: "ftp",
					Value:       []string{"smtp"},
				},
			},
		},
	},
}
var LogTargetSenderAddressCondVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "type",
			AttrType:    "String",
			AttrDefault: "file",
			Value:       []string{"smtp"},
		},
		{
			Evaluation: "logical-and",
			Conditions: []validators.Evaluation{
				{
					Evaluation:  "property-value-in-list",
					Attribute:   "type",
					AttrType:    "String",
					AttrDefault: "file",
					Value:       []string{"file"},
				},
				{
					Evaluation:  "property-value-in-list",
					Attribute:   "upload_method",
					AttrType:    "String",
					AttrDefault: "ftp",
					Value:       []string{"smtp"},
				},
			},
		},
	},
}
var LogTargetSMTPDomainCondVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "type",
			AttrType:    "String",
			AttrDefault: "file",
			Value:       []string{"smtp"},
		},
		{
			Evaluation: "logical-and",
			Conditions: []validators.Evaluation{
				{
					Evaluation:  "property-value-in-list",
					Attribute:   "type",
					AttrType:    "String",
					AttrDefault: "file",
					Value:       []string{"file"},
				},
				{
					Evaluation:  "property-value-in-list",
					Attribute:   "upload_method",
					AttrType:    "String",
					AttrDefault: "ftp",
					Value:       []string{"smtp"},
				},
			},
		},
	},
}
var LogTargetSizeCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "type",
	AttrType:    "String",
	AttrDefault: "file",
	Value:       []string{"file", "nfs"},
}
var LogTargetURLCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "type",
	AttrType:    "String",
	AttrDefault: "file",
	Value:       []string{"soap"},
}
var LogTargetNFSMountCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "type",
	AttrType:    "String",
	AttrDefault: "file",
	Value:       []string{"nfs"},
}
var LogTargetLocalFileCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "type",
	AttrType:    "String",
	AttrDefault: "file",
	Value:       []string{"file"},
}
var LogTargetNFSFileCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "type",
	AttrType:    "String",
	AttrDefault: "file",
	Value:       []string{"nfs"},
}
var LogTargetArchiveModeCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "type",
	AttrType:    "String",
	AttrDefault: "file",
	Value:       []string{"file"},
}
var LogTargetUploadMethodCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "archive_mode",
			AttrType:    "String",
			AttrDefault: "rotate",
			Value:       []string{"upload"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "type",
			AttrType:    "String",
			AttrDefault: "file",
			Value:       []string{"file"},
		},
	},
}
var LogTargetRotateCondVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "type",
			AttrType:    "String",
			AttrDefault: "file",
			Value:       []string{"nfs"},
		},
		{
			Evaluation: "logical-and",
			Conditions: []validators.Evaluation{
				{
					Evaluation:  "property-value-in-list",
					Attribute:   "archive_mode",
					AttrType:    "String",
					AttrDefault: "rotate",
					Value:       []string{"rotate"},
				},
				{
					Evaluation:  "property-value-in-list",
					Attribute:   "type",
					AttrType:    "String",
					AttrDefault: "file",
					Value:       []string{"file"},
				},
			},
		},
	},
}
var LogTargetRemoteAddressCondVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "type",
			AttrType:    "String",
			AttrDefault: "file",
			Value:       []string{"syslog", "syslog-tcp", "smtp"},
		},
		{
			Evaluation: "logical-and",
			Conditions: []validators.Evaluation{
				{
					Evaluation:  "property-value-in-list",
					Attribute:   "archive_mode",
					AttrType:    "String",
					AttrDefault: "rotate",
					Value:       []string{"upload"},
				},
				{
					Evaluation:  "property-value-in-list",
					Attribute:   "type",
					AttrType:    "String",
					AttrDefault: "file",
					Value:       []string{"file"},
				},
			},
		},
	},
}
var LogTargetRemoteLoginCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "upload_method",
			AttrType:    "String",
			AttrDefault: "ftp",
			Value:       []string{"ftp", "scp", "sftp"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "archive_mode",
			AttrType:    "String",
			AttrDefault: "rotate",
			Value:       []string{"upload"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "type",
			AttrType:    "String",
			AttrDefault: "file",
			Value:       []string{"file"},
		},
	},
}
var LogTargetRemotePasswordCondVal = validators.Evaluation{
	Evaluation: "logical-and",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "upload_method",
			AttrType:    "String",
			AttrDefault: "ftp",
			Value:       []string{"ftp"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "archive_mode",
			AttrType:    "String",
			AttrDefault: "rotate",
			Value:       []string{"upload"},
		},
		{
			Evaluation:  "property-value-in-list",
			Attribute:   "type",
			AttrType:    "String",
			AttrDefault: "file",
			Value:       []string{"file"},
		},
	},
}
var LogTargetLocalAddressCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "type",
	AttrType:    "String",
	AttrDefault: "file",
	Value:       []string{"smtp", "syslog-tcp"},
}
var LogTargetSyslogFacilityCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "type",
	AttrType:    "String",
	AttrDefault: "file",
	Value:       []string{"syslog-tcp", "syslog"},
}
var LogTargetLogPrecisionCondVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "type",
	AttrType:    "String",
	AttrDefault: "file",
	Value:       []string{"syslog-tcp"},
}
var LogTargetBackupIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "type",
	AttrType:    "String",
	AttrDefault: "file",
	Value:       []string{"file"},
}
var LogTargetPriorityIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "type",
	AttrType:    "String",
	AttrDefault: "file",
	Value:       []string{"syslog-tcp"},
}
var LogTargetSoapVersionIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "type",
	AttrType:    "String",
	AttrDefault: "file",
	Value:       []string{"soap"},
}
var LogTargetFormatIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "type",
	AttrType:    "String",
	AttrDefault: "file",
	Value:       []string{"snmp", "syslog", "syslog-tcp"},
}
var LogTargetTimestampFormatIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "type",
	AttrType:    "String",
	AttrDefault: "file",
	Value:       []string{"smtp", "syslog", "syslog-tcp"},
}
var LogTargetFixedFormatIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-in-list",
	Attribute:   "format",
	AttrType:    "String",
	AttrDefault: "xml",
	Value:       []string{"json-icp"},
}
var LogTargetLocalIdentifierIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "type",
	AttrType:    "String",
	AttrDefault: "file",
	Value:       []string{"syslog-tcp"},
}
var LogTargetEmailAddressIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}
var LogTargetSenderAddressIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}
var LogTargetSMTPDomainIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "type",
	AttrType:    "String",
	AttrDefault: "file",
	Value:       []string{"smtp"},
}
var LogTargetSizeIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}
var LogTargetURLIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}
var LogTargetNFSMountIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}
var LogTargetLocalFileIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}
var LogTargetNFSFileIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}
var LogTargetArchiveModeIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}
var LogTargetUploadMethodIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}
var LogTargetRotateIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}
var LogTargetUseANSIColorIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "type",
	AttrType:    "String",
	AttrDefault: "file",
	Value:       []string{"console"},
}
var LogTargetRemoteAddressIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}
var LogTargetRemotePortIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation: "logical-and",
			Conditions: []validators.Evaluation{
				{
					Evaluation:  "property-value-in-list",
					Attribute:   "type",
					AttrType:    "String",
					AttrDefault: "file",
					Value:       []string{"file"},
				},
				{
					Evaluation:  "property-value-not-in-list",
					Attribute:   "archive_mode",
					AttrType:    "String",
					AttrDefault: "rotate",
					Value:       []string{"upload"},
				},
			},
		},
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "type",
			AttrType:    "String",
			AttrDefault: "file",
			Value:       []string{"syslog", "syslog-tcp", "smtp", "file"},
		},
	},
}
var LogTargetRemoteLoginIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}
var LogTargetRemotePasswordIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "upload_method",
			AttrType:    "String",
			AttrDefault: "ftp",
			Value:       []string{"ftp", "scp", "sftp"},
		},
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "archive_mode",
			AttrType:    "String",
			AttrDefault: "rotate",
			Value:       []string{"upload"},
		},
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "type",
			AttrType:    "String",
			AttrDefault: "file",
			Value:       []string{"file"},
		},
	},
}
var LogTargetRemoteDirectoryIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "upload_method",
			AttrType:    "String",
			AttrDefault: "ftp",
			Value:       []string{"ftp", "scp", "sftp"},
		},
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "archive_mode",
			AttrType:    "String",
			AttrDefault: "rotate",
			Value:       []string{"upload"},
		},
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "type",
			AttrType:    "String",
			AttrDefault: "file",
			Value:       []string{"file"},
		},
	},
}
var LogTargetLocalAddressIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "type",
	AttrType:    "String",
	AttrDefault: "file",
	Value:       []string{"syslog", "soap"},
}
var LogTargetSyslogFacilityIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}
var LogTargetRateLimitIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "type",
	AttrType:    "String",
	AttrDefault: "file",
	Value:       []string{"nfs", "smtp", "soap", "snmp", "syslog", "syslog-tcp"},
}
var LogTargetMaxConnectionsIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "type",
	AttrType:    "String",
	AttrDefault: "file",
	Value:       []string{"syslog-tcp"},
}
var LogTargetConnectTimeoutIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "type",
	AttrType:    "String",
	AttrDefault: "file",
	Value:       []string{"syslog-tcp"},
}
var LogTargetIdleTimeoutIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "type",
	AttrType:    "String",
	AttrDefault: "file",
	Value:       []string{"syslog-tcp"},
}
var LogTargetActiveTimeoutIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "type",
	AttrType:    "String",
	AttrDefault: "file",
	Value:       []string{"syslog-tcp"},
}
var LogTargetSSLClientProfileIgnoreVal = validators.Evaluation{
	Evaluation: "logical-or",
	Conditions: []validators.Evaluation{
		{
			Evaluation:  "property-value-not-in-list",
			Attribute:   "type",
			AttrType:    "String",
			AttrDefault: "file",
			Value:       []string{"soap", "syslog-tcp"},
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
var LogTargetSSLClientConfigTypeIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "type",
	AttrType:    "String",
	AttrDefault: "file",
	Value:       []string{"soap", "syslog-tcp"},
}
var LogTargetRetryIntervalIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "type",
	AttrType:    "String",
	AttrDefault: "file",
	Value:       []string{"syslog-tcp"},
}
var LogTargetRetryAttemptsIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "type",
	AttrType:    "String",
	AttrDefault: "file",
	Value:       []string{"syslog-tcp"},
}
var LogTargetLongRetryIntervalIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "type",
	AttrType:    "String",
	AttrDefault: "file",
	Value:       []string{"syslog-tcp"},
}
var LogTargetLogPrecisionIgnoreVal = validators.Evaluation{
	Evaluation: "logical-true",
}
var LogTargetEventBufferSizeIgnoreVal = validators.Evaluation{
	Evaluation:  "property-value-not-in-list",
	Attribute:   "type",
	AttrType:    "String",
	AttrDefault: "file",
	Value:       []string{"syslog-tcp"},
}

var LogTargetObjectType = map[string]attr.Type{
	"id":                         types.StringType,
	"app_domain":                 types.StringType,
	"backup":                     types.StringType,
	"log_events":                 types.ListType{ElemType: types.ObjectType{AttrTypes: DmLogEventObjectType}},
	"user_summary":               types.StringType,
	"type":                       types.StringType,
	"priority":                   types.StringType,
	"soap_version":               types.StringType,
	"format":                     types.StringType,
	"timestamp_format":           types.StringType,
	"fixed_format":               types.BoolType,
	"local_identifier":           types.StringType,
	"email_address":              types.StringType,
	"sender_address":             types.StringType,
	"smtp_domain":                types.StringType,
	"size":                       types.Int64Type,
	"url":                        types.StringType,
	"nf_sm_ount":                 types.StringType,
	"local_file":                 types.StringType,
	"nfs_file":                   types.StringType,
	"archive_mode":               types.StringType,
	"upload_method":              types.StringType,
	"rotate":                     types.Int64Type,
	"use_ansi_color":             types.BoolType,
	"remote_address":             types.StringType,
	"remote_port":                types.Int64Type,
	"remote_login":               types.StringType,
	"remote_password_wo":         types.StringType,
	"remote_password_wo_version": types.Int64Type,
	"remote_directory":           types.StringType,
	"local_address":              types.StringType,
	"syslog_facility":            types.StringType,
	"rate_limit":                 types.Int64Type,
	"max_connections":            types.Int64Type,
	"connect_timeout":            types.Int64Type,
	"idle_timeout":               types.Int64Type,
	"active_timeout":             types.Int64Type,
	"feedback_detection":         types.BoolType,
	"log_event_code":             types.ListType{ElemType: types.StringType},
	"log_event_filter":           types.ListType{ElemType: types.StringType},
	"log_objects":                types.ListType{ElemType: types.ObjectType{AttrTypes: DmLogObjectObjectType}},
	"log_ip_filter":              types.ListType{ElemType: types.ObjectType{AttrTypes: DmLogIPFilterObjectType}},
	"log_triggers":               types.ListType{ElemType: types.ObjectType{AttrTypes: DmLogTriggerObjectType}},
	"ssl_client_profile":         types.StringType,
	"ssl_client_config_type":     types.StringType,
	"retry_interval":             types.Int64Type,
	"retry_attempts":             types.Int64Type,
	"long_retry_interval":        types.Int64Type,
	"log_precision":              types.StringType,
	"event_buffer_size":          types.StringType,
	"dependency_actions":         actions.ActionsListType,
}
var LogTargetObjectTypeWO = map[string]attr.Type{
	"id":                     types.StringType,
	"app_domain":             types.StringType,
	"backup":                 types.StringType,
	"log_events":             types.ListType{ElemType: types.ObjectType{AttrTypes: DmLogEventObjectType}},
	"user_summary":           types.StringType,
	"type":                   types.StringType,
	"priority":               types.StringType,
	"soap_version":           types.StringType,
	"format":                 types.StringType,
	"timestamp_format":       types.StringType,
	"fixed_format":           types.BoolType,
	"local_identifier":       types.StringType,
	"email_address":          types.StringType,
	"sender_address":         types.StringType,
	"smtp_domain":            types.StringType,
	"size":                   types.Int64Type,
	"url":                    types.StringType,
	"nf_sm_ount":             types.StringType,
	"local_file":             types.StringType,
	"nfs_file":               types.StringType,
	"archive_mode":           types.StringType,
	"upload_method":          types.StringType,
	"rotate":                 types.Int64Type,
	"use_ansi_color":         types.BoolType,
	"remote_address":         types.StringType,
	"remote_port":            types.Int64Type,
	"remote_login":           types.StringType,
	"remote_directory":       types.StringType,
	"local_address":          types.StringType,
	"syslog_facility":        types.StringType,
	"rate_limit":             types.Int64Type,
	"max_connections":        types.Int64Type,
	"connect_timeout":        types.Int64Type,
	"idle_timeout":           types.Int64Type,
	"active_timeout":         types.Int64Type,
	"feedback_detection":     types.BoolType,
	"log_event_code":         types.ListType{ElemType: types.StringType},
	"log_event_filter":       types.ListType{ElemType: types.StringType},
	"log_objects":            types.ListType{ElemType: types.ObjectType{AttrTypes: DmLogObjectObjectType}},
	"log_ip_filter":          types.ListType{ElemType: types.ObjectType{AttrTypes: DmLogIPFilterObjectType}},
	"log_triggers":           types.ListType{ElemType: types.ObjectType{AttrTypes: DmLogTriggerObjectType}},
	"ssl_client_profile":     types.StringType,
	"ssl_client_config_type": types.StringType,
	"retry_interval":         types.Int64Type,
	"retry_attempts":         types.Int64Type,
	"long_retry_interval":    types.Int64Type,
	"log_precision":          types.StringType,
	"event_buffer_size":      types.StringType,
	"dependency_actions":     actions.ActionsListType,
}

func (data LogTarget) GetPath() string {
	rest_path := "/mgmt/config/{domain}/LogTarget"
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data LogTargetWO) GetPath() string {
	rest_path := "/mgmt/config/{domain}/LogTarget"
	rest_path = strings.ReplaceAll(rest_path, "{name}", url.QueryEscape(data.Id.ValueString()))
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data LogTarget) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.Backup.IsNull() {
		return false
	}
	if !data.LogEvents.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.Type.IsNull() {
		return false
	}
	if !data.Priority.IsNull() {
		return false
	}
	if !data.SoapVersion.IsNull() {
		return false
	}
	if !data.Format.IsNull() {
		return false
	}
	if !data.TimestampFormat.IsNull() {
		return false
	}
	if !data.FixedFormat.IsNull() {
		return false
	}
	if !data.LocalIdentifier.IsNull() {
		return false
	}
	if !data.EmailAddress.IsNull() {
		return false
	}
	if !data.SenderAddress.IsNull() {
		return false
	}
	if !data.SmtpDomain.IsNull() {
		return false
	}
	if !data.Size.IsNull() {
		return false
	}
	if !data.Url.IsNull() {
		return false
	}
	if !data.NfSmOunt.IsNull() {
		return false
	}
	if !data.LocalFile.IsNull() {
		return false
	}
	if !data.NfsFile.IsNull() {
		return false
	}
	if !data.ArchiveMode.IsNull() {
		return false
	}
	if !data.UploadMethod.IsNull() {
		return false
	}
	if !data.Rotate.IsNull() {
		return false
	}
	if !data.UseAnsiColor.IsNull() {
		return false
	}
	if !data.RemoteAddress.IsNull() {
		return false
	}
	if !data.RemotePort.IsNull() {
		return false
	}
	if !data.RemoteLogin.IsNull() {
		return false
	}
	if !data.RemotePasswordWo.IsNull() {
		return false
	}
	if !data.RemoteDirectory.IsNull() {
		return false
	}
	if !data.LocalAddress.IsNull() {
		return false
	}
	if !data.SyslogFacility.IsNull() {
		return false
	}
	if !data.RateLimit.IsNull() {
		return false
	}
	if !data.MaxConnections.IsNull() {
		return false
	}
	if !data.ConnectTimeout.IsNull() {
		return false
	}
	if !data.IdleTimeout.IsNull() {
		return false
	}
	if !data.ActiveTimeout.IsNull() {
		return false
	}
	if !data.FeedbackDetection.IsNull() {
		return false
	}
	if !data.LogEventCode.IsNull() {
		return false
	}
	if !data.LogEventFilter.IsNull() {
		return false
	}
	if !data.LogObjects.IsNull() {
		return false
	}
	if !data.LogIpFilter.IsNull() {
		return false
	}
	if !data.LogTriggers.IsNull() {
		return false
	}
	if !data.SslClientProfile.IsNull() {
		return false
	}
	if !data.SslClientConfigType.IsNull() {
		return false
	}
	if !data.RetryInterval.IsNull() {
		return false
	}
	if !data.RetryAttempts.IsNull() {
		return false
	}
	if !data.LongRetryInterval.IsNull() {
		return false
	}
	if !data.LogPrecision.IsNull() {
		return false
	}
	if !data.EventBufferSize.IsNull() {
		return false
	}
	return true
}
func (data LogTargetWO) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.Backup.IsNull() {
		return false
	}
	if !data.LogEvents.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.Type.IsNull() {
		return false
	}
	if !data.Priority.IsNull() {
		return false
	}
	if !data.SoapVersion.IsNull() {
		return false
	}
	if !data.Format.IsNull() {
		return false
	}
	if !data.TimestampFormat.IsNull() {
		return false
	}
	if !data.FixedFormat.IsNull() {
		return false
	}
	if !data.LocalIdentifier.IsNull() {
		return false
	}
	if !data.EmailAddress.IsNull() {
		return false
	}
	if !data.SenderAddress.IsNull() {
		return false
	}
	if !data.SmtpDomain.IsNull() {
		return false
	}
	if !data.Size.IsNull() {
		return false
	}
	if !data.Url.IsNull() {
		return false
	}
	if !data.NfSmOunt.IsNull() {
		return false
	}
	if !data.LocalFile.IsNull() {
		return false
	}
	if !data.NfsFile.IsNull() {
		return false
	}
	if !data.ArchiveMode.IsNull() {
		return false
	}
	if !data.UploadMethod.IsNull() {
		return false
	}
	if !data.Rotate.IsNull() {
		return false
	}
	if !data.UseAnsiColor.IsNull() {
		return false
	}
	if !data.RemoteAddress.IsNull() {
		return false
	}
	if !data.RemotePort.IsNull() {
		return false
	}
	if !data.RemoteLogin.IsNull() {
		return false
	}
	if !data.RemoteDirectory.IsNull() {
		return false
	}
	if !data.LocalAddress.IsNull() {
		return false
	}
	if !data.SyslogFacility.IsNull() {
		return false
	}
	if !data.RateLimit.IsNull() {
		return false
	}
	if !data.MaxConnections.IsNull() {
		return false
	}
	if !data.ConnectTimeout.IsNull() {
		return false
	}
	if !data.IdleTimeout.IsNull() {
		return false
	}
	if !data.ActiveTimeout.IsNull() {
		return false
	}
	if !data.FeedbackDetection.IsNull() {
		return false
	}
	if !data.LogEventCode.IsNull() {
		return false
	}
	if !data.LogEventFilter.IsNull() {
		return false
	}
	if !data.LogObjects.IsNull() {
		return false
	}
	if !data.LogIpFilter.IsNull() {
		return false
	}
	if !data.LogTriggers.IsNull() {
		return false
	}
	if !data.SslClientProfile.IsNull() {
		return false
	}
	if !data.SslClientConfigType.IsNull() {
		return false
	}
	if !data.RetryInterval.IsNull() {
		return false
	}
	if !data.RetryAttempts.IsNull() {
		return false
	}
	if !data.LongRetryInterval.IsNull() {
		return false
	}
	if !data.LogPrecision.IsNull() {
		return false
	}
	if !data.EventBufferSize.IsNull() {
		return false
	}
	return true
}

func (data LogTarget) ToBody(ctx context.Context, pathRoot string, config *LogTarget) string {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	body := ""

	if !data.Id.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`name`, data.Id.ValueString())
	}
	if !data.Backup.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Backup`, data.Backup.ValueString())
	}
	if !data.LogEvents.IsNull() {
		var dataValues []DmLogEvent
		data.LogEvents.ElementsAs(ctx, &dataValues, false)
		for _, val := range dataValues {
			body, _ = sjson.SetRaw(body, pathRoot+`LogEvents`+".-1", val.ToBody(ctx, ""))
		}
	}
	if !data.UserSummary.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UserSummary`, data.UserSummary.ValueString())
	}
	if !data.Type.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Type`, data.Type.ValueString())
	}
	if !data.Priority.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Priority`, data.Priority.ValueString())
	}
	if !data.SoapVersion.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SoapVersion`, data.SoapVersion.ValueString())
	}
	if !data.Format.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Format`, data.Format.ValueString())
	}
	if !data.TimestampFormat.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`TimestampFormat`, data.TimestampFormat.ValueString())
	}
	if !data.FixedFormat.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`FixedFormat`, tfutils.StringFromBool(data.FixedFormat, ""))
	}
	if !data.LocalIdentifier.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LocalIdentifier`, data.LocalIdentifier.ValueString())
	}
	if !data.EmailAddress.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EmailAddress`, data.EmailAddress.ValueString())
	}
	if !data.SenderAddress.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SenderAddress`, data.SenderAddress.ValueString())
	}
	if !data.SmtpDomain.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SMTPDomain`, data.SmtpDomain.ValueString())
	}
	if !data.Size.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Size`, data.Size.ValueInt64())
	}
	if !data.Url.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`URL`, data.Url.ValueString())
	}
	if !data.NfSmOunt.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`NFSMount`, data.NfSmOunt.ValueString())
	}
	if !data.LocalFile.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LocalFile`, data.LocalFile.ValueString())
	}
	if !data.NfsFile.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`NFSFile`, data.NfsFile.ValueString())
	}
	if !data.ArchiveMode.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ArchiveMode`, data.ArchiveMode.ValueString())
	}
	if !data.UploadMethod.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UploadMethod`, data.UploadMethod.ValueString())
	}
	if !data.Rotate.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Rotate`, data.Rotate.ValueInt64())
	}
	if !data.UseAnsiColor.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`UseANSIColor`, tfutils.StringFromBool(data.UseAnsiColor, ""))
	}
	if !data.RemoteAddress.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`RemoteAddress`, data.RemoteAddress.ValueString())
	}
	if !data.RemotePort.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`RemotePort`, data.RemotePort.ValueInt64())
	}
	if !data.RemoteLogin.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`RemoteLogin`, data.RemoteLogin.ValueString())
	}
	if !data.RemotePasswordWo.IsNull() || !data.RemotePasswordWoVersion.IsNull() {
		if data.RemotePasswordWo.IsNull() && config != nil {
			data.RemotePasswordWo = config.RemotePasswordWo
		}
		body, _ = sjson.Set(body, pathRoot+`RemotePassword`, data.RemotePasswordWo.ValueString())
	}
	if !data.RemoteDirectory.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`RemoteDirectory`, data.RemoteDirectory.ValueString())
	}
	if !data.LocalAddress.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LocalAddress`, data.LocalAddress.ValueString())
	}
	if !data.SyslogFacility.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SyslogFacility`, data.SyslogFacility.ValueString())
	}
	if !data.RateLimit.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`RateLimit`, data.RateLimit.ValueInt64())
	}
	if !data.MaxConnections.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MaxConnections`, data.MaxConnections.ValueInt64())
	}
	if !data.ConnectTimeout.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ConnectTimeout`, data.ConnectTimeout.ValueInt64())
	}
	if !data.IdleTimeout.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`IdleTimeout`, data.IdleTimeout.ValueInt64())
	}
	if !data.ActiveTimeout.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`ActiveTimeout`, data.ActiveTimeout.ValueInt64())
	}
	if !data.FeedbackDetection.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`FeedbackDetection`, tfutils.StringFromBool(data.FeedbackDetection, ""))
	}
	if !data.LogEventCode.IsNull() {
		var dataValues []string
		data.LogEventCode.ElementsAs(ctx, &dataValues, false)
		for _, val := range dataValues {
			body, _ = sjson.Set(body, pathRoot+`LogEventCode`+".-1", map[string]string{"value": val})
		}
	}
	if !data.LogEventFilter.IsNull() {
		var dataValues []string
		data.LogEventFilter.ElementsAs(ctx, &dataValues, false)
		for _, val := range dataValues {
			body, _ = sjson.Set(body, pathRoot+`LogEventFilter`+".-1", map[string]string{"value": val})
		}
	}
	if !data.LogObjects.IsNull() {
		var dataValues []DmLogObject
		data.LogObjects.ElementsAs(ctx, &dataValues, false)
		for _, val := range dataValues {
			body, _ = sjson.SetRaw(body, pathRoot+`LogObjects`+".-1", val.ToBody(ctx, ""))
		}
	}
	if !data.LogIpFilter.IsNull() {
		var dataValues []DmLogIPFilter
		data.LogIpFilter.ElementsAs(ctx, &dataValues, false)
		for _, val := range dataValues {
			body, _ = sjson.SetRaw(body, pathRoot+`LogIPFilter`+".-1", val.ToBody(ctx, ""))
		}
	}
	if !data.LogTriggers.IsNull() {
		var dataValues []DmLogTrigger
		data.LogTriggers.ElementsAs(ctx, &dataValues, false)
		for _, val := range dataValues {
			body, _ = sjson.SetRaw(body, pathRoot+`LogTriggers`+".-1", val.ToBody(ctx, ""))
		}
	}
	if !data.SslClientProfile.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SSLClientProfile`, data.SslClientProfile.ValueString())
	}
	if !data.SslClientConfigType.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`SSLClientConfigType`, data.SslClientConfigType.ValueString())
	}
	if !data.RetryInterval.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`RetryInterval`, data.RetryInterval.ValueInt64())
	}
	if !data.RetryAttempts.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`RetryAttempts`, data.RetryAttempts.ValueInt64())
	}
	if !data.LongRetryInterval.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LongRetryInterval`, data.LongRetryInterval.ValueInt64())
	}
	if !data.LogPrecision.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`LogPrecision`, data.LogPrecision.ValueString())
	}
	if !data.EventBufferSize.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`EventBufferSize`, data.EventBufferSize.ValueString())
	}
	return body
}

func (data *LogTarget) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `name`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Id = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Id = types.StringNull()
	}
	if value := res.Get(pathRoot + `Backup`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Backup = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Backup = types.StringNull()
	}
	if value := res.Get(pathRoot + `LogEvents`); value.Exists() {
		l := []DmLogEvent{}
		if value := res.Get(`LogEvents`); value.Exists() {
			for _, v := range value.Array() {
				item := DmLogEvent{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.LogEvents, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmLogEventObjectType}, l)
		} else {
			data.LogEvents = types.ListNull(types.ObjectType{AttrTypes: DmLogEventObjectType})
		}
	} else {
		data.LogEvents = types.ListNull(types.ObjectType{AttrTypes: DmLogEventObjectType})
	}
	if value := res.Get(pathRoot + `UserSummary`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.UserSummary = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UserSummary = types.StringNull()
	}
	if value := res.Get(pathRoot + `Type`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Type = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Type = types.StringValue("file")
	}
	if value := res.Get(pathRoot + `Priority`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Priority = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Priority = types.StringValue("normal")
	}
	if value := res.Get(pathRoot + `SoapVersion`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SoapVersion = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SoapVersion = types.StringValue("soap11")
	}
	if value := res.Get(pathRoot + `Format`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Format = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Format = types.StringValue("xml")
	}
	if value := res.Get(pathRoot + `TimestampFormat`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.TimestampFormat = tfutils.ParseStringFromGJSON(value)
	} else {
		data.TimestampFormat = types.StringValue("zulu")
	}
	if value := res.Get(pathRoot + `FixedFormat`); value.Exists() {
		data.FixedFormat = tfutils.BoolFromString(value.String())
	} else {
		data.FixedFormat = types.BoolNull()
	}
	if value := res.Get(pathRoot + `LocalIdentifier`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.LocalIdentifier = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LocalIdentifier = types.StringNull()
	}
	if value := res.Get(pathRoot + `EmailAddress`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.EmailAddress = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EmailAddress = types.StringNull()
	}
	if value := res.Get(pathRoot + `SenderAddress`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SenderAddress = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SenderAddress = types.StringNull()
	}
	if value := res.Get(pathRoot + `SMTPDomain`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SmtpDomain = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SmtpDomain = types.StringNull()
	}
	if value := res.Get(pathRoot + `Size`); value.Exists() {
		data.Size = types.Int64Value(value.Int())
	} else {
		data.Size = types.Int64Value(500)
	}
	if value := res.Get(pathRoot + `URL`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Url = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Url = types.StringNull()
	}
	if value := res.Get(pathRoot + `NFSMount`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.NfSmOunt = tfutils.ParseStringFromGJSON(value)
	} else {
		data.NfSmOunt = types.StringNull()
	}
	if value := res.Get(pathRoot + `LocalFile`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.LocalFile = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LocalFile = types.StringNull()
	}
	if value := res.Get(pathRoot + `NFSFile`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.NfsFile = tfutils.ParseStringFromGJSON(value)
	} else {
		data.NfsFile = types.StringNull()
	}
	if value := res.Get(pathRoot + `ArchiveMode`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ArchiveMode = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ArchiveMode = types.StringValue("rotate")
	}
	if value := res.Get(pathRoot + `UploadMethod`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.UploadMethod = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UploadMethod = types.StringValue("ftp")
	}
	if value := res.Get(pathRoot + `Rotate`); value.Exists() {
		data.Rotate = types.Int64Value(value.Int())
	} else {
		data.Rotate = types.Int64Value(3)
	}
	if value := res.Get(pathRoot + `UseANSIColor`); value.Exists() {
		data.UseAnsiColor = tfutils.BoolFromString(value.String())
	} else {
		data.UseAnsiColor = types.BoolNull()
	}
	if value := res.Get(pathRoot + `RemoteAddress`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.RemoteAddress = tfutils.ParseStringFromGJSON(value)
	} else {
		data.RemoteAddress = types.StringNull()
	}
	if value := res.Get(pathRoot + `RemotePort`); value.Exists() {
		data.RemotePort = types.Int64Value(value.Int())
	} else {
		data.RemotePort = types.Int64Null()
	}
	if value := res.Get(pathRoot + `RemoteLogin`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.RemoteLogin = tfutils.ParseStringFromGJSON(value)
	} else {
		data.RemoteLogin = types.StringNull()
	}
	if value := res.Get(pathRoot + `RemotePassword`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.RemotePasswordWo = tfutils.ParseStringFromGJSON(value)
	} else {
		data.RemotePasswordWo = types.StringNull()
	}
	if value := res.Get(pathRoot + `RemoteDirectory`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.RemoteDirectory = tfutils.ParseStringFromGJSON(value)
	} else {
		data.RemoteDirectory = types.StringNull()
	}
	if value := res.Get(pathRoot + `LocalAddress`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.LocalAddress = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LocalAddress = types.StringNull()
	}
	if value := res.Get(pathRoot + `SyslogFacility`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SyslogFacility = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SyslogFacility = types.StringValue("user")
	}
	if value := res.Get(pathRoot + `RateLimit`); value.Exists() {
		data.RateLimit = types.Int64Value(value.Int())
	} else {
		data.RateLimit = types.Int64Value(100)
	}
	if value := res.Get(pathRoot + `MaxConnections`); value.Exists() {
		data.MaxConnections = types.Int64Value(value.Int())
	} else {
		data.MaxConnections = types.Int64Value(1)
	}
	if value := res.Get(pathRoot + `ConnectTimeout`); value.Exists() {
		data.ConnectTimeout = types.Int64Value(value.Int())
	} else {
		data.ConnectTimeout = types.Int64Value(60)
	}
	if value := res.Get(pathRoot + `IdleTimeout`); value.Exists() {
		data.IdleTimeout = types.Int64Value(value.Int())
	} else {
		data.IdleTimeout = types.Int64Value(15)
	}
	if value := res.Get(pathRoot + `ActiveTimeout`); value.Exists() {
		data.ActiveTimeout = types.Int64Value(value.Int())
	} else {
		data.ActiveTimeout = types.Int64Value(0)
	}
	if value := res.Get(pathRoot + `FeedbackDetection`); value.Exists() {
		data.FeedbackDetection = tfutils.BoolFromString(value.String())
	} else {
		data.FeedbackDetection = types.BoolNull()
	}
	if value := res.Get(pathRoot + `LogEventCode`); value.Exists() {
		data.LogEventCode = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.LogEventCode = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `LogEventFilter`); value.Exists() {
		data.LogEventFilter = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.LogEventFilter = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `LogObjects`); value.Exists() {
		l := []DmLogObject{}
		if value := res.Get(`LogObjects`); value.Exists() {
			for _, v := range value.Array() {
				item := DmLogObject{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.LogObjects, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmLogObjectObjectType}, l)
		} else {
			data.LogObjects = types.ListNull(types.ObjectType{AttrTypes: DmLogObjectObjectType})
		}
	} else {
		data.LogObjects = types.ListNull(types.ObjectType{AttrTypes: DmLogObjectObjectType})
	}
	if value := res.Get(pathRoot + `LogIPFilter`); value.Exists() {
		l := []DmLogIPFilter{}
		if value := res.Get(`LogIPFilter`); value.Exists() {
			for _, v := range value.Array() {
				item := DmLogIPFilter{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.LogIpFilter, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmLogIPFilterObjectType}, l)
		} else {
			data.LogIpFilter = types.ListNull(types.ObjectType{AttrTypes: DmLogIPFilterObjectType})
		}
	} else {
		data.LogIpFilter = types.ListNull(types.ObjectType{AttrTypes: DmLogIPFilterObjectType})
	}
	if value := res.Get(pathRoot + `LogTriggers`); value.Exists() {
		l := []DmLogTrigger{}
		if value := res.Get(`LogTriggers`); value.Exists() {
			for _, v := range value.Array() {
				item := DmLogTrigger{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.LogTriggers, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmLogTriggerObjectType}, l)
		} else {
			data.LogTriggers = types.ListNull(types.ObjectType{AttrTypes: DmLogTriggerObjectType})
		}
	} else {
		data.LogTriggers = types.ListNull(types.ObjectType{AttrTypes: DmLogTriggerObjectType})
	}
	if value := res.Get(pathRoot + `SSLClientProfile`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SslClientProfile = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SslClientProfile = types.StringNull()
	}
	if value := res.Get(pathRoot + `SSLClientConfigType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SslClientConfigType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SslClientConfigType = types.StringValue("client")
	}
	if value := res.Get(pathRoot + `RetryInterval`); value.Exists() {
		data.RetryInterval = types.Int64Value(value.Int())
	} else {
		data.RetryInterval = types.Int64Value(1)
	}
	if value := res.Get(pathRoot + `RetryAttempts`); value.Exists() {
		data.RetryAttempts = types.Int64Value(value.Int())
	} else {
		data.RetryAttempts = types.Int64Value(1)
	}
	if value := res.Get(pathRoot + `LongRetryInterval`); value.Exists() {
		data.LongRetryInterval = types.Int64Value(value.Int())
	} else {
		data.LongRetryInterval = types.Int64Value(20)
	}
	if value := res.Get(pathRoot + `LogPrecision`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.LogPrecision = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LogPrecision = types.StringValue("second")
	}
	if value := res.Get(pathRoot + `EventBufferSize`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.EventBufferSize = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EventBufferSize = types.StringValue("2048")
	}
}
func (data *LogTargetWO) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `name`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Id = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Id = types.StringNull()
	}
	if value := res.Get(pathRoot + `Backup`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Backup = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Backup = types.StringNull()
	}
	if value := res.Get(pathRoot + `LogEvents`); value.Exists() {
		l := []DmLogEvent{}
		if value := res.Get(`LogEvents`); value.Exists() {
			for _, v := range value.Array() {
				item := DmLogEvent{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.LogEvents, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmLogEventObjectType}, l)
		} else {
			data.LogEvents = types.ListNull(types.ObjectType{AttrTypes: DmLogEventObjectType})
		}
	} else {
		data.LogEvents = types.ListNull(types.ObjectType{AttrTypes: DmLogEventObjectType})
	}
	if value := res.Get(pathRoot + `UserSummary`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.UserSummary = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UserSummary = types.StringNull()
	}
	if value := res.Get(pathRoot + `Type`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Type = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Type = types.StringValue("file")
	}
	if value := res.Get(pathRoot + `Priority`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Priority = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Priority = types.StringValue("normal")
	}
	if value := res.Get(pathRoot + `SoapVersion`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SoapVersion = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SoapVersion = types.StringValue("soap11")
	}
	if value := res.Get(pathRoot + `Format`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Format = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Format = types.StringValue("xml")
	}
	if value := res.Get(pathRoot + `TimestampFormat`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.TimestampFormat = tfutils.ParseStringFromGJSON(value)
	} else {
		data.TimestampFormat = types.StringValue("zulu")
	}
	if value := res.Get(pathRoot + `FixedFormat`); value.Exists() {
		data.FixedFormat = tfutils.BoolFromString(value.String())
	} else {
		data.FixedFormat = types.BoolNull()
	}
	if value := res.Get(pathRoot + `LocalIdentifier`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.LocalIdentifier = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LocalIdentifier = types.StringNull()
	}
	if value := res.Get(pathRoot + `EmailAddress`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.EmailAddress = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EmailAddress = types.StringNull()
	}
	if value := res.Get(pathRoot + `SenderAddress`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SenderAddress = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SenderAddress = types.StringNull()
	}
	if value := res.Get(pathRoot + `SMTPDomain`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SmtpDomain = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SmtpDomain = types.StringNull()
	}
	if value := res.Get(pathRoot + `Size`); value.Exists() {
		data.Size = types.Int64Value(value.Int())
	} else {
		data.Size = types.Int64Value(500)
	}
	if value := res.Get(pathRoot + `URL`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Url = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Url = types.StringNull()
	}
	if value := res.Get(pathRoot + `NFSMount`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.NfSmOunt = tfutils.ParseStringFromGJSON(value)
	} else {
		data.NfSmOunt = types.StringNull()
	}
	if value := res.Get(pathRoot + `LocalFile`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.LocalFile = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LocalFile = types.StringNull()
	}
	if value := res.Get(pathRoot + `NFSFile`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.NfsFile = tfutils.ParseStringFromGJSON(value)
	} else {
		data.NfsFile = types.StringNull()
	}
	if value := res.Get(pathRoot + `ArchiveMode`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.ArchiveMode = tfutils.ParseStringFromGJSON(value)
	} else {
		data.ArchiveMode = types.StringValue("rotate")
	}
	if value := res.Get(pathRoot + `UploadMethod`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.UploadMethod = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UploadMethod = types.StringValue("ftp")
	}
	if value := res.Get(pathRoot + `Rotate`); value.Exists() {
		data.Rotate = types.Int64Value(value.Int())
	} else {
		data.Rotate = types.Int64Value(3)
	}
	if value := res.Get(pathRoot + `UseANSIColor`); value.Exists() {
		data.UseAnsiColor = tfutils.BoolFromString(value.String())
	} else {
		data.UseAnsiColor = types.BoolNull()
	}
	if value := res.Get(pathRoot + `RemoteAddress`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.RemoteAddress = tfutils.ParseStringFromGJSON(value)
	} else {
		data.RemoteAddress = types.StringNull()
	}
	if value := res.Get(pathRoot + `RemotePort`); value.Exists() {
		data.RemotePort = types.Int64Value(value.Int())
	} else {
		data.RemotePort = types.Int64Null()
	}
	if value := res.Get(pathRoot + `RemoteLogin`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.RemoteLogin = tfutils.ParseStringFromGJSON(value)
	} else {
		data.RemoteLogin = types.StringNull()
	}
	if value := res.Get(pathRoot + `RemoteDirectory`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.RemoteDirectory = tfutils.ParseStringFromGJSON(value)
	} else {
		data.RemoteDirectory = types.StringNull()
	}
	if value := res.Get(pathRoot + `LocalAddress`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.LocalAddress = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LocalAddress = types.StringNull()
	}
	if value := res.Get(pathRoot + `SyslogFacility`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SyslogFacility = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SyslogFacility = types.StringValue("user")
	}
	if value := res.Get(pathRoot + `RateLimit`); value.Exists() {
		data.RateLimit = types.Int64Value(value.Int())
	} else {
		data.RateLimit = types.Int64Value(100)
	}
	if value := res.Get(pathRoot + `MaxConnections`); value.Exists() {
		data.MaxConnections = types.Int64Value(value.Int())
	} else {
		data.MaxConnections = types.Int64Value(1)
	}
	if value := res.Get(pathRoot + `ConnectTimeout`); value.Exists() {
		data.ConnectTimeout = types.Int64Value(value.Int())
	} else {
		data.ConnectTimeout = types.Int64Value(60)
	}
	if value := res.Get(pathRoot + `IdleTimeout`); value.Exists() {
		data.IdleTimeout = types.Int64Value(value.Int())
	} else {
		data.IdleTimeout = types.Int64Value(15)
	}
	if value := res.Get(pathRoot + `ActiveTimeout`); value.Exists() {
		data.ActiveTimeout = types.Int64Value(value.Int())
	} else {
		data.ActiveTimeout = types.Int64Value(0)
	}
	if value := res.Get(pathRoot + `FeedbackDetection`); value.Exists() {
		data.FeedbackDetection = tfutils.BoolFromString(value.String())
	} else {
		data.FeedbackDetection = types.BoolNull()
	}
	if value := res.Get(pathRoot + `LogEventCode`); value.Exists() {
		data.LogEventCode = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.LogEventCode = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `LogEventFilter`); value.Exists() {
		data.LogEventFilter = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.LogEventFilter = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `LogObjects`); value.Exists() {
		l := []DmLogObject{}
		if value := res.Get(`LogObjects`); value.Exists() {
			for _, v := range value.Array() {
				item := DmLogObject{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.LogObjects, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmLogObjectObjectType}, l)
		} else {
			data.LogObjects = types.ListNull(types.ObjectType{AttrTypes: DmLogObjectObjectType})
		}
	} else {
		data.LogObjects = types.ListNull(types.ObjectType{AttrTypes: DmLogObjectObjectType})
	}
	if value := res.Get(pathRoot + `LogIPFilter`); value.Exists() {
		l := []DmLogIPFilter{}
		if value := res.Get(`LogIPFilter`); value.Exists() {
			for _, v := range value.Array() {
				item := DmLogIPFilter{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.LogIpFilter, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmLogIPFilterObjectType}, l)
		} else {
			data.LogIpFilter = types.ListNull(types.ObjectType{AttrTypes: DmLogIPFilterObjectType})
		}
	} else {
		data.LogIpFilter = types.ListNull(types.ObjectType{AttrTypes: DmLogIPFilterObjectType})
	}
	if value := res.Get(pathRoot + `LogTriggers`); value.Exists() {
		l := []DmLogTrigger{}
		if value := res.Get(`LogTriggers`); value.Exists() {
			for _, v := range value.Array() {
				item := DmLogTrigger{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.LogTriggers, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmLogTriggerObjectType}, l)
		} else {
			data.LogTriggers = types.ListNull(types.ObjectType{AttrTypes: DmLogTriggerObjectType})
		}
	} else {
		data.LogTriggers = types.ListNull(types.ObjectType{AttrTypes: DmLogTriggerObjectType})
	}
	if value := res.Get(pathRoot + `SSLClientProfile`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SslClientProfile = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SslClientProfile = types.StringNull()
	}
	if value := res.Get(pathRoot + `SSLClientConfigType`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.SslClientConfigType = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SslClientConfigType = types.StringValue("client")
	}
	if value := res.Get(pathRoot + `RetryInterval`); value.Exists() {
		data.RetryInterval = types.Int64Value(value.Int())
	} else {
		data.RetryInterval = types.Int64Value(1)
	}
	if value := res.Get(pathRoot + `RetryAttempts`); value.Exists() {
		data.RetryAttempts = types.Int64Value(value.Int())
	} else {
		data.RetryAttempts = types.Int64Value(1)
	}
	if value := res.Get(pathRoot + `LongRetryInterval`); value.Exists() {
		data.LongRetryInterval = types.Int64Value(value.Int())
	} else {
		data.LongRetryInterval = types.Int64Value(20)
	}
	if value := res.Get(pathRoot + `LogPrecision`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.LogPrecision = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LogPrecision = types.StringValue("second")
	}
	if value := res.Get(pathRoot + `EventBufferSize`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.EventBufferSize = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EventBufferSize = types.StringValue("2048")
	}
}

func (data *LogTarget) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if value := res.Get(pathRoot + `name`); value.Exists() && !data.Id.IsNull() {
		data.Id = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Id = types.StringNull()
	}
	if value := res.Get(pathRoot + `Backup`); value.Exists() && !data.Backup.IsNull() {
		data.Backup = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Backup = types.StringNull()
	}
	if value := res.Get(pathRoot + `LogEvents`); value.Exists() && !data.LogEvents.IsNull() {
		l := []DmLogEvent{}
		for _, v := range value.Array() {
			item := DmLogEvent{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
		if len(l) > 0 {
			data.LogEvents, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmLogEventObjectType}, l)
		} else {
			data.LogEvents = types.ListNull(types.ObjectType{AttrTypes: DmLogEventObjectType})
		}
	} else {
		data.LogEvents = types.ListNull(types.ObjectType{AttrTypes: DmLogEventObjectType})
	}
	if value := res.Get(pathRoot + `UserSummary`); value.Exists() && !data.UserSummary.IsNull() {
		data.UserSummary = tfutils.ParseStringFromGJSON(value)
	} else {
		data.UserSummary = types.StringNull()
	}
	if value := res.Get(pathRoot + `Type`); value.Exists() && !data.Type.IsNull() {
		data.Type = tfutils.ParseStringFromGJSON(value)
	} else if data.Type.ValueString() != "file" {
		data.Type = types.StringNull()
	}
	if value := res.Get(pathRoot + `Priority`); value.Exists() && !data.Priority.IsNull() {
		data.Priority = tfutils.ParseStringFromGJSON(value)
	} else if data.Priority.ValueString() != "normal" {
		data.Priority = types.StringNull()
	}
	if value := res.Get(pathRoot + `SoapVersion`); value.Exists() && !data.SoapVersion.IsNull() {
		data.SoapVersion = tfutils.ParseStringFromGJSON(value)
	} else if data.SoapVersion.ValueString() != "soap11" {
		data.SoapVersion = types.StringNull()
	}
	if value := res.Get(pathRoot + `Format`); value.Exists() && !data.Format.IsNull() {
		data.Format = tfutils.ParseStringFromGJSON(value)
	} else if data.Format.ValueString() != "xml" {
		data.Format = types.StringNull()
	}
	if value := res.Get(pathRoot + `TimestampFormat`); value.Exists() && !data.TimestampFormat.IsNull() {
		data.TimestampFormat = tfutils.ParseStringFromGJSON(value)
	} else if data.TimestampFormat.ValueString() != "zulu" {
		data.TimestampFormat = types.StringNull()
	}
	if value := res.Get(pathRoot + `FixedFormat`); value.Exists() && !data.FixedFormat.IsNull() {
		data.FixedFormat = tfutils.BoolFromString(value.String())
	} else if data.FixedFormat.ValueBool() {
		data.FixedFormat = types.BoolNull()
	}
	if value := res.Get(pathRoot + `LocalIdentifier`); value.Exists() && !data.LocalIdentifier.IsNull() {
		data.LocalIdentifier = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LocalIdentifier = types.StringNull()
	}
	if value := res.Get(pathRoot + `EmailAddress`); value.Exists() && !data.EmailAddress.IsNull() {
		data.EmailAddress = tfutils.ParseStringFromGJSON(value)
	} else {
		data.EmailAddress = types.StringNull()
	}
	if value := res.Get(pathRoot + `SenderAddress`); value.Exists() && !data.SenderAddress.IsNull() {
		data.SenderAddress = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SenderAddress = types.StringNull()
	}
	if value := res.Get(pathRoot + `SMTPDomain`); value.Exists() && !data.SmtpDomain.IsNull() {
		data.SmtpDomain = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SmtpDomain = types.StringNull()
	}
	if value := res.Get(pathRoot + `Size`); value.Exists() && !data.Size.IsNull() {
		data.Size = types.Int64Value(value.Int())
	} else if data.Size.ValueInt64() != 500 {
		data.Size = types.Int64Null()
	}
	if value := res.Get(pathRoot + `URL`); value.Exists() && !data.Url.IsNull() {
		data.Url = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Url = types.StringNull()
	}
	if value := res.Get(pathRoot + `NFSMount`); value.Exists() && !data.NfSmOunt.IsNull() {
		data.NfSmOunt = tfutils.ParseStringFromGJSON(value)
	} else {
		data.NfSmOunt = types.StringNull()
	}
	if value := res.Get(pathRoot + `LocalFile`); value.Exists() && !data.LocalFile.IsNull() {
		data.LocalFile = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LocalFile = types.StringNull()
	}
	if value := res.Get(pathRoot + `NFSFile`); value.Exists() && !data.NfsFile.IsNull() {
		data.NfsFile = tfutils.ParseStringFromGJSON(value)
	} else {
		data.NfsFile = types.StringNull()
	}
	if value := res.Get(pathRoot + `ArchiveMode`); value.Exists() && !data.ArchiveMode.IsNull() {
		data.ArchiveMode = tfutils.ParseStringFromGJSON(value)
	} else if data.ArchiveMode.ValueString() != "rotate" {
		data.ArchiveMode = types.StringNull()
	}
	if value := res.Get(pathRoot + `UploadMethod`); value.Exists() && !data.UploadMethod.IsNull() {
		data.UploadMethod = tfutils.ParseStringFromGJSON(value)
	} else if data.UploadMethod.ValueString() != "ftp" {
		data.UploadMethod = types.StringNull()
	}
	if value := res.Get(pathRoot + `Rotate`); value.Exists() && !data.Rotate.IsNull() {
		data.Rotate = types.Int64Value(value.Int())
	} else if data.Rotate.ValueInt64() != 3 {
		data.Rotate = types.Int64Null()
	}
	if value := res.Get(pathRoot + `UseANSIColor`); value.Exists() && !data.UseAnsiColor.IsNull() {
		data.UseAnsiColor = tfutils.BoolFromString(value.String())
	} else if data.UseAnsiColor.ValueBool() {
		data.UseAnsiColor = types.BoolNull()
	}
	if value := res.Get(pathRoot + `RemoteAddress`); value.Exists() && !data.RemoteAddress.IsNull() {
		data.RemoteAddress = tfutils.ParseStringFromGJSON(value)
	} else {
		data.RemoteAddress = types.StringNull()
	}
	if value := res.Get(pathRoot + `RemotePort`); value.Exists() && !data.RemotePort.IsNull() {
		data.RemotePort = types.Int64Value(value.Int())
	} else {
		data.RemotePort = types.Int64Null()
	}
	if value := res.Get(pathRoot + `RemoteLogin`); value.Exists() && !data.RemoteLogin.IsNull() {
		data.RemoteLogin = tfutils.ParseStringFromGJSON(value)
	} else {
		data.RemoteLogin = types.StringNull()
	}
	if value := res.Get(pathRoot + `RemotePassword`); value.Exists() && !data.RemotePasswordWo.IsNull() {
		data.RemotePasswordWo = tfutils.ParseStringFromGJSON(value)
	} else {
		data.RemotePasswordWo = types.StringNull()
	}
	if value := res.Get(pathRoot + `RemoteDirectory`); value.Exists() && !data.RemoteDirectory.IsNull() {
		data.RemoteDirectory = tfutils.ParseStringFromGJSON(value)
	} else {
		data.RemoteDirectory = types.StringNull()
	}
	if value := res.Get(pathRoot + `LocalAddress`); value.Exists() && !data.LocalAddress.IsNull() {
		data.LocalAddress = tfutils.ParseStringFromGJSON(value)
	} else {
		data.LocalAddress = types.StringNull()
	}
	if value := res.Get(pathRoot + `SyslogFacility`); value.Exists() && !data.SyslogFacility.IsNull() {
		data.SyslogFacility = tfutils.ParseStringFromGJSON(value)
	} else if data.SyslogFacility.ValueString() != "user" {
		data.SyslogFacility = types.StringNull()
	}
	if value := res.Get(pathRoot + `RateLimit`); value.Exists() && !data.RateLimit.IsNull() {
		data.RateLimit = types.Int64Value(value.Int())
	} else if data.RateLimit.ValueInt64() != 100 {
		data.RateLimit = types.Int64Null()
	}
	if value := res.Get(pathRoot + `MaxConnections`); value.Exists() && !data.MaxConnections.IsNull() {
		data.MaxConnections = types.Int64Value(value.Int())
	} else if data.MaxConnections.ValueInt64() != 1 {
		data.MaxConnections = types.Int64Null()
	}
	if value := res.Get(pathRoot + `ConnectTimeout`); value.Exists() && !data.ConnectTimeout.IsNull() {
		data.ConnectTimeout = types.Int64Value(value.Int())
	} else if data.ConnectTimeout.ValueInt64() != 60 {
		data.ConnectTimeout = types.Int64Null()
	}
	if value := res.Get(pathRoot + `IdleTimeout`); value.Exists() && !data.IdleTimeout.IsNull() {
		data.IdleTimeout = types.Int64Value(value.Int())
	} else if data.IdleTimeout.ValueInt64() != 15 {
		data.IdleTimeout = types.Int64Null()
	}
	if value := res.Get(pathRoot + `ActiveTimeout`); value.Exists() && !data.ActiveTimeout.IsNull() {
		data.ActiveTimeout = types.Int64Value(value.Int())
	} else if data.ActiveTimeout.ValueInt64() != 0 {
		data.ActiveTimeout = types.Int64Null()
	}
	if value := res.Get(pathRoot + `FeedbackDetection`); value.Exists() && !data.FeedbackDetection.IsNull() {
		data.FeedbackDetection = tfutils.BoolFromString(value.String())
	} else if data.FeedbackDetection.ValueBool() {
		data.FeedbackDetection = types.BoolNull()
	}
	if value := res.Get(pathRoot + `LogEventCode`); value.Exists() && !data.LogEventCode.IsNull() {
		data.LogEventCode = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.LogEventCode = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `LogEventFilter`); value.Exists() && !data.LogEventFilter.IsNull() {
		data.LogEventFilter = tfutils.ParseStringListFromGJSON(value)
	} else {
		data.LogEventFilter = types.ListNull(types.StringType)
	}
	if value := res.Get(pathRoot + `LogObjects`); value.Exists() && !data.LogObjects.IsNull() {
		l := []DmLogObject{}
		for _, v := range value.Array() {
			item := DmLogObject{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
		if len(l) > 0 {
			data.LogObjects, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmLogObjectObjectType}, l)
		} else {
			data.LogObjects = types.ListNull(types.ObjectType{AttrTypes: DmLogObjectObjectType})
		}
	} else {
		data.LogObjects = types.ListNull(types.ObjectType{AttrTypes: DmLogObjectObjectType})
	}
	if value := res.Get(pathRoot + `LogIPFilter`); value.Exists() && !data.LogIpFilter.IsNull() {
		l := []DmLogIPFilter{}
		for _, v := range value.Array() {
			item := DmLogIPFilter{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
		if len(l) > 0 {
			data.LogIpFilter, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmLogIPFilterObjectType}, l)
		} else {
			data.LogIpFilter = types.ListNull(types.ObjectType{AttrTypes: DmLogIPFilterObjectType})
		}
	} else {
		data.LogIpFilter = types.ListNull(types.ObjectType{AttrTypes: DmLogIPFilterObjectType})
	}
	if value := res.Get(pathRoot + `LogTriggers`); value.Exists() && !data.LogTriggers.IsNull() {
		l := []DmLogTrigger{}
		for _, v := range value.Array() {
			item := DmLogTrigger{}
			item.FromBody(ctx, "", v)
			if !item.IsNull() {
				l = append(l, item)
			}
		}
		if len(l) > 0 {
			data.LogTriggers, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmLogTriggerObjectType}, l)
		} else {
			data.LogTriggers = types.ListNull(types.ObjectType{AttrTypes: DmLogTriggerObjectType})
		}
	} else {
		data.LogTriggers = types.ListNull(types.ObjectType{AttrTypes: DmLogTriggerObjectType})
	}
	if value := res.Get(pathRoot + `SSLClientProfile`); value.Exists() && !data.SslClientProfile.IsNull() {
		data.SslClientProfile = tfutils.ParseStringFromGJSON(value)
	} else {
		data.SslClientProfile = types.StringNull()
	}
	if value := res.Get(pathRoot + `SSLClientConfigType`); value.Exists() && !data.SslClientConfigType.IsNull() {
		data.SslClientConfigType = tfutils.ParseStringFromGJSON(value)
	} else if data.SslClientConfigType.ValueString() != "client" {
		data.SslClientConfigType = types.StringNull()
	}
	if value := res.Get(pathRoot + `RetryInterval`); value.Exists() && !data.RetryInterval.IsNull() {
		data.RetryInterval = types.Int64Value(value.Int())
	} else if data.RetryInterval.ValueInt64() != 1 {
		data.RetryInterval = types.Int64Null()
	}
	if value := res.Get(pathRoot + `RetryAttempts`); value.Exists() && !data.RetryAttempts.IsNull() {
		data.RetryAttempts = types.Int64Value(value.Int())
	} else if data.RetryAttempts.ValueInt64() != 1 {
		data.RetryAttempts = types.Int64Null()
	}
	if value := res.Get(pathRoot + `LongRetryInterval`); value.Exists() && !data.LongRetryInterval.IsNull() {
		data.LongRetryInterval = types.Int64Value(value.Int())
	} else if data.LongRetryInterval.ValueInt64() != 20 {
		data.LongRetryInterval = types.Int64Null()
	}
	if value := res.Get(pathRoot + `LogPrecision`); value.Exists() && !data.LogPrecision.IsNull() {
		data.LogPrecision = tfutils.ParseStringFromGJSON(value)
	} else if data.LogPrecision.ValueString() != "second" {
		data.LogPrecision = types.StringNull()
	}
	if value := res.Get(pathRoot + `EventBufferSize`); value.Exists() && !data.EventBufferSize.IsNull() {
		data.EventBufferSize = tfutils.ParseStringFromGJSON(value)
	} else if data.EventBufferSize.ValueString() != "2048" {
		data.EventBufferSize = types.StringNull()
	}
}
func (data *LogTarget) UpdateUnknownFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
	if pathRoot != "" {
		pathRoot = pathRoot + "."
	}
	if data.Id.IsUnknown() {
		if value := res.Get(pathRoot + `name`); value.Exists() && !data.Id.IsNull() {
			data.Id = tfutils.ParseStringFromGJSON(value)
		} else {
			data.Id = types.StringNull()
		}
	}
	if data.Backup.IsUnknown() {
		if value := res.Get(pathRoot + `Backup`); value.Exists() && !data.Backup.IsNull() {
			data.Backup = tfutils.ParseStringFromGJSON(value)
		} else {
			data.Backup = types.StringNull()
		}
	}
	if data.LogEvents.IsUnknown() {
		if value := res.Get(pathRoot + `LogEvents`); value.Exists() && !data.LogEvents.IsNull() {
			l := []DmLogEvent{}
			if value := res.Get(`LogEvents`); value.Exists() {
				for _, v := range value.Array() {
					item := DmLogEvent{}
					item.FromBody(ctx, "", v)
					if !item.IsNull() {
						l = append(l, item)
					}
				}
			}
			if len(l) > 0 {
				data.LogEvents, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmLogEventObjectType}, l)
			} else {
				data.LogEvents = types.ListNull(types.ObjectType{AttrTypes: DmLogEventObjectType})
			}
		} else {
			data.LogEvents = types.ListNull(types.ObjectType{AttrTypes: DmLogEventObjectType})
		}
	}
	if data.UserSummary.IsUnknown() {
		if value := res.Get(pathRoot + `UserSummary`); value.Exists() && !data.UserSummary.IsNull() {
			data.UserSummary = tfutils.ParseStringFromGJSON(value)
		} else {
			data.UserSummary = types.StringNull()
		}
	}
	if data.Type.IsUnknown() {
		if value := res.Get(pathRoot + `Type`); value.Exists() && !data.Type.IsNull() {
			data.Type = tfutils.ParseStringFromGJSON(value)
		} else if data.Type.ValueString() != "file" {
			data.Type = types.StringNull()
		}
	}
	if data.Priority.IsUnknown() {
		if value := res.Get(pathRoot + `Priority`); value.Exists() && !data.Priority.IsNull() {
			data.Priority = tfutils.ParseStringFromGJSON(value)
		} else if data.Priority.ValueString() != "normal" {
			data.Priority = types.StringNull()
		}
	}
	if data.SoapVersion.IsUnknown() {
		if value := res.Get(pathRoot + `SoapVersion`); value.Exists() && !data.SoapVersion.IsNull() {
			data.SoapVersion = tfutils.ParseStringFromGJSON(value)
		} else if data.SoapVersion.ValueString() != "soap11" {
			data.SoapVersion = types.StringNull()
		}
	}
	if data.Format.IsUnknown() {
		if value := res.Get(pathRoot + `Format`); value.Exists() && !data.Format.IsNull() {
			data.Format = tfutils.ParseStringFromGJSON(value)
		} else if data.Format.ValueString() != "xml" {
			data.Format = types.StringNull()
		}
	}
	if data.TimestampFormat.IsUnknown() {
		if value := res.Get(pathRoot + `TimestampFormat`); value.Exists() && !data.TimestampFormat.IsNull() {
			data.TimestampFormat = tfutils.ParseStringFromGJSON(value)
		} else if data.TimestampFormat.ValueString() != "zulu" {
			data.TimestampFormat = types.StringNull()
		}
	}
	if data.FixedFormat.IsUnknown() {
		if value := res.Get(pathRoot + `FixedFormat`); value.Exists() && !data.FixedFormat.IsNull() {
			data.FixedFormat = tfutils.BoolFromString(value.String())
		} else {
			data.FixedFormat = types.BoolNull()
		}
	}
	if data.LocalIdentifier.IsUnknown() {
		if value := res.Get(pathRoot + `LocalIdentifier`); value.Exists() && !data.LocalIdentifier.IsNull() {
			data.LocalIdentifier = tfutils.ParseStringFromGJSON(value)
		} else {
			data.LocalIdentifier = types.StringNull()
		}
	}
	if data.EmailAddress.IsUnknown() {
		if value := res.Get(pathRoot + `EmailAddress`); value.Exists() && !data.EmailAddress.IsNull() {
			data.EmailAddress = tfutils.ParseStringFromGJSON(value)
		} else {
			data.EmailAddress = types.StringNull()
		}
	}
	if data.SenderAddress.IsUnknown() {
		if value := res.Get(pathRoot + `SenderAddress`); value.Exists() && !data.SenderAddress.IsNull() {
			data.SenderAddress = tfutils.ParseStringFromGJSON(value)
		} else {
			data.SenderAddress = types.StringNull()
		}
	}
	if data.SmtpDomain.IsUnknown() {
		if value := res.Get(pathRoot + `SMTPDomain`); value.Exists() && !data.SmtpDomain.IsNull() {
			data.SmtpDomain = tfutils.ParseStringFromGJSON(value)
		} else {
			data.SmtpDomain = types.StringNull()
		}
	}
	if data.Size.IsUnknown() {
		if value := res.Get(pathRoot + `Size`); value.Exists() && !data.Size.IsNull() {
			data.Size = types.Int64Value(value.Int())
		} else if data.Size.ValueInt64() != 500 {
			data.Size = types.Int64Null()
		}
	}
	if data.Url.IsUnknown() {
		if value := res.Get(pathRoot + `URL`); value.Exists() && !data.Url.IsNull() {
			data.Url = tfutils.ParseStringFromGJSON(value)
		} else {
			data.Url = types.StringNull()
		}
	}
	if data.NfSmOunt.IsUnknown() {
		if value := res.Get(pathRoot + `NFSMount`); value.Exists() && !data.NfSmOunt.IsNull() {
			data.NfSmOunt = tfutils.ParseStringFromGJSON(value)
		} else {
			data.NfSmOunt = types.StringNull()
		}
	}
	if data.LocalFile.IsUnknown() {
		if value := res.Get(pathRoot + `LocalFile`); value.Exists() && !data.LocalFile.IsNull() {
			data.LocalFile = tfutils.ParseStringFromGJSON(value)
		} else {
			data.LocalFile = types.StringNull()
		}
	}
	if data.NfsFile.IsUnknown() {
		if value := res.Get(pathRoot + `NFSFile`); value.Exists() && !data.NfsFile.IsNull() {
			data.NfsFile = tfutils.ParseStringFromGJSON(value)
		} else {
			data.NfsFile = types.StringNull()
		}
	}
	if data.ArchiveMode.IsUnknown() {
		if value := res.Get(pathRoot + `ArchiveMode`); value.Exists() && !data.ArchiveMode.IsNull() {
			data.ArchiveMode = tfutils.ParseStringFromGJSON(value)
		} else if data.ArchiveMode.ValueString() != "rotate" {
			data.ArchiveMode = types.StringNull()
		}
	}
	if data.UploadMethod.IsUnknown() {
		if value := res.Get(pathRoot + `UploadMethod`); value.Exists() && !data.UploadMethod.IsNull() {
			data.UploadMethod = tfutils.ParseStringFromGJSON(value)
		} else if data.UploadMethod.ValueString() != "ftp" {
			data.UploadMethod = types.StringNull()
		}
	}
	if data.Rotate.IsUnknown() {
		if value := res.Get(pathRoot + `Rotate`); value.Exists() && !data.Rotate.IsNull() {
			data.Rotate = types.Int64Value(value.Int())
		} else if data.Rotate.ValueInt64() != 3 {
			data.Rotate = types.Int64Null()
		}
	}
	if data.UseAnsiColor.IsUnknown() {
		if value := res.Get(pathRoot + `UseANSIColor`); value.Exists() && !data.UseAnsiColor.IsNull() {
			data.UseAnsiColor = tfutils.BoolFromString(value.String())
		} else {
			data.UseAnsiColor = types.BoolNull()
		}
	}
	if data.RemoteAddress.IsUnknown() {
		if value := res.Get(pathRoot + `RemoteAddress`); value.Exists() && !data.RemoteAddress.IsNull() {
			data.RemoteAddress = tfutils.ParseStringFromGJSON(value)
		} else {
			data.RemoteAddress = types.StringNull()
		}
	}
	if data.RemotePort.IsUnknown() {
		if value := res.Get(pathRoot + `RemotePort`); value.Exists() && !data.RemotePort.IsNull() {
			data.RemotePort = types.Int64Value(value.Int())
		} else {
			data.RemotePort = types.Int64Null()
		}
	}
	if data.RemoteLogin.IsUnknown() {
		if value := res.Get(pathRoot + `RemoteLogin`); value.Exists() && !data.RemoteLogin.IsNull() {
			data.RemoteLogin = tfutils.ParseStringFromGJSON(value)
		} else {
			data.RemoteLogin = types.StringNull()
		}
	}
	if data.RemotePasswordWo.IsUnknown() {
		if value := res.Get(pathRoot + `RemotePassword`); value.Exists() && !data.RemotePasswordWo.IsNull() {
			data.RemotePasswordWo = tfutils.ParseStringFromGJSON(value)
		} else {
			data.RemotePasswordWo = types.StringNull()
		}
	}
	if data.RemoteDirectory.IsUnknown() {
		if value := res.Get(pathRoot + `RemoteDirectory`); value.Exists() && !data.RemoteDirectory.IsNull() {
			data.RemoteDirectory = tfutils.ParseStringFromGJSON(value)
		} else {
			data.RemoteDirectory = types.StringNull()
		}
	}
	if data.LocalAddress.IsUnknown() {
		if value := res.Get(pathRoot + `LocalAddress`); value.Exists() && !data.LocalAddress.IsNull() {
			data.LocalAddress = tfutils.ParseStringFromGJSON(value)
		} else {
			data.LocalAddress = types.StringNull()
		}
	}
	if data.SyslogFacility.IsUnknown() {
		if value := res.Get(pathRoot + `SyslogFacility`); value.Exists() && !data.SyslogFacility.IsNull() {
			data.SyslogFacility = tfutils.ParseStringFromGJSON(value)
		} else if data.SyslogFacility.ValueString() != "user" {
			data.SyslogFacility = types.StringNull()
		}
	}
	if data.RateLimit.IsUnknown() {
		if value := res.Get(pathRoot + `RateLimit`); value.Exists() && !data.RateLimit.IsNull() {
			data.RateLimit = types.Int64Value(value.Int())
		} else if data.RateLimit.ValueInt64() != 100 {
			data.RateLimit = types.Int64Null()
		}
	}
	if data.MaxConnections.IsUnknown() {
		if value := res.Get(pathRoot + `MaxConnections`); value.Exists() && !data.MaxConnections.IsNull() {
			data.MaxConnections = types.Int64Value(value.Int())
		} else if data.MaxConnections.ValueInt64() != 1 {
			data.MaxConnections = types.Int64Null()
		}
	}
	if data.ConnectTimeout.IsUnknown() {
		if value := res.Get(pathRoot + `ConnectTimeout`); value.Exists() && !data.ConnectTimeout.IsNull() {
			data.ConnectTimeout = types.Int64Value(value.Int())
		} else if data.ConnectTimeout.ValueInt64() != 60 {
			data.ConnectTimeout = types.Int64Null()
		}
	}
	if data.IdleTimeout.IsUnknown() {
		if value := res.Get(pathRoot + `IdleTimeout`); value.Exists() && !data.IdleTimeout.IsNull() {
			data.IdleTimeout = types.Int64Value(value.Int())
		} else if data.IdleTimeout.ValueInt64() != 15 {
			data.IdleTimeout = types.Int64Null()
		}
	}
	if data.ActiveTimeout.IsUnknown() {
		if value := res.Get(pathRoot + `ActiveTimeout`); value.Exists() && !data.ActiveTimeout.IsNull() {
			data.ActiveTimeout = types.Int64Value(value.Int())
		} else if data.ActiveTimeout.ValueInt64() != 0 {
			data.ActiveTimeout = types.Int64Null()
		}
	}
	if data.FeedbackDetection.IsUnknown() {
		if value := res.Get(pathRoot + `FeedbackDetection`); value.Exists() && !data.FeedbackDetection.IsNull() {
			data.FeedbackDetection = tfutils.BoolFromString(value.String())
		} else {
			data.FeedbackDetection = types.BoolNull()
		}
	}
	if data.LogEventCode.IsUnknown() {
		if value := res.Get(pathRoot + `LogEventCode`); value.Exists() && !data.LogEventCode.IsNull() {
			data.LogEventCode = tfutils.ParseStringListFromGJSON(value)
		} else {
			data.LogEventCode = types.ListNull(types.StringType)
		}
	}
	if data.LogEventFilter.IsUnknown() {
		if value := res.Get(pathRoot + `LogEventFilter`); value.Exists() && !data.LogEventFilter.IsNull() {
			data.LogEventFilter = tfutils.ParseStringListFromGJSON(value)
		} else {
			data.LogEventFilter = types.ListNull(types.StringType)
		}
	}
	if data.LogObjects.IsUnknown() {
		if value := res.Get(pathRoot + `LogObjects`); value.Exists() && !data.LogObjects.IsNull() {
			l := []DmLogObject{}
			if value := res.Get(`LogObjects`); value.Exists() {
				for _, v := range value.Array() {
					item := DmLogObject{}
					item.FromBody(ctx, "", v)
					if !item.IsNull() {
						l = append(l, item)
					}
				}
			}
			if len(l) > 0 {
				data.LogObjects, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmLogObjectObjectType}, l)
			} else {
				data.LogObjects = types.ListNull(types.ObjectType{AttrTypes: DmLogObjectObjectType})
			}
		} else {
			data.LogObjects = types.ListNull(types.ObjectType{AttrTypes: DmLogObjectObjectType})
		}
	}
	if data.LogIpFilter.IsUnknown() {
		if value := res.Get(pathRoot + `LogIPFilter`); value.Exists() && !data.LogIpFilter.IsNull() {
			l := []DmLogIPFilter{}
			if value := res.Get(`LogIPFilter`); value.Exists() {
				for _, v := range value.Array() {
					item := DmLogIPFilter{}
					item.FromBody(ctx, "", v)
					if !item.IsNull() {
						l = append(l, item)
					}
				}
			}
			if len(l) > 0 {
				data.LogIpFilter, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmLogIPFilterObjectType}, l)
			} else {
				data.LogIpFilter = types.ListNull(types.ObjectType{AttrTypes: DmLogIPFilterObjectType})
			}
		} else {
			data.LogIpFilter = types.ListNull(types.ObjectType{AttrTypes: DmLogIPFilterObjectType})
		}
	}
	if data.LogTriggers.IsUnknown() {
		if value := res.Get(pathRoot + `LogTriggers`); value.Exists() && !data.LogTriggers.IsNull() {
			l := []DmLogTrigger{}
			if value := res.Get(`LogTriggers`); value.Exists() {
				for _, v := range value.Array() {
					item := DmLogTrigger{}
					item.FromBody(ctx, "", v)
					if !item.IsNull() {
						l = append(l, item)
					}
				}
			}
			if len(l) > 0 {
				data.LogTriggers, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmLogTriggerObjectType}, l)
			} else {
				data.LogTriggers = types.ListNull(types.ObjectType{AttrTypes: DmLogTriggerObjectType})
			}
		} else {
			data.LogTriggers = types.ListNull(types.ObjectType{AttrTypes: DmLogTriggerObjectType})
		}
	}
	if data.SslClientProfile.IsUnknown() {
		if value := res.Get(pathRoot + `SSLClientProfile`); value.Exists() && !data.SslClientProfile.IsNull() {
			data.SslClientProfile = tfutils.ParseStringFromGJSON(value)
		} else {
			data.SslClientProfile = types.StringNull()
		}
	}
	if data.SslClientConfigType.IsUnknown() {
		if value := res.Get(pathRoot + `SSLClientConfigType`); value.Exists() && !data.SslClientConfigType.IsNull() {
			data.SslClientConfigType = tfutils.ParseStringFromGJSON(value)
		} else if data.SslClientConfigType.ValueString() != "client" {
			data.SslClientConfigType = types.StringNull()
		}
	}
	if data.RetryInterval.IsUnknown() {
		if value := res.Get(pathRoot + `RetryInterval`); value.Exists() && !data.RetryInterval.IsNull() {
			data.RetryInterval = types.Int64Value(value.Int())
		} else if data.RetryInterval.ValueInt64() != 1 {
			data.RetryInterval = types.Int64Null()
		}
	}
	if data.RetryAttempts.IsUnknown() {
		if value := res.Get(pathRoot + `RetryAttempts`); value.Exists() && !data.RetryAttempts.IsNull() {
			data.RetryAttempts = types.Int64Value(value.Int())
		} else if data.RetryAttempts.ValueInt64() != 1 {
			data.RetryAttempts = types.Int64Null()
		}
	}
	if data.LongRetryInterval.IsUnknown() {
		if value := res.Get(pathRoot + `LongRetryInterval`); value.Exists() && !data.LongRetryInterval.IsNull() {
			data.LongRetryInterval = types.Int64Value(value.Int())
		} else if data.LongRetryInterval.ValueInt64() != 20 {
			data.LongRetryInterval = types.Int64Null()
		}
	}
	if data.LogPrecision.IsUnknown() {
		if value := res.Get(pathRoot + `LogPrecision`); value.Exists() && !data.LogPrecision.IsNull() {
			data.LogPrecision = tfutils.ParseStringFromGJSON(value)
		} else if data.LogPrecision.ValueString() != "second" {
			data.LogPrecision = types.StringNull()
		}
	}
	if data.EventBufferSize.IsUnknown() {
		if value := res.Get(pathRoot + `EventBufferSize`); value.Exists() && !data.EventBufferSize.IsNull() {
			data.EventBufferSize = tfutils.ParseStringFromGJSON(value)
		} else if data.EventBufferSize.ValueString() != "2048" {
			data.EventBufferSize = types.StringNull()
		}
	}
}

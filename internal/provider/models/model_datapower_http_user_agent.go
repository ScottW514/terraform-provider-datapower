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
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type HTTPUserAgent struct {
	Id                       types.String                `tfsdk:"id"`
	AppDomain                types.String                `tfsdk:"app_domain"`
	UserSummary              types.String                `tfsdk:"user_summary"`
	Identifier               types.String                `tfsdk:"identifier"`
	MaxRedirects             types.Int64                 `tfsdk:"max_redirects"`
	Timeout                  types.Int64                 `tfsdk:"timeout"`
	ProxyPolicies            types.List                  `tfsdk:"proxy_policies"`
	SslPolicies              types.List                  `tfsdk:"ssl_policies"`
	BasicAuthPolicies        types.List                  `tfsdk:"basic_auth_policies"`
	SoapActionPolicies       types.List                  `tfsdk:"soap_action_policies"`
	PubkeyAuthPolicies       types.List                  `tfsdk:"pubkey_auth_policies"`
	AllowCompressionPolicies types.List                  `tfsdk:"allow_compression_policies"`
	HeaderRetentionPolicies  types.List                  `tfsdk:"header_retention_policies"`
	HttpVersionPolicies      types.List                  `tfsdk:"http_version_policies"`
	AddHeaderPolicies        types.List                  `tfsdk:"add_header_policies"`
	UploadChunkedPolicies    types.List                  `tfsdk:"upload_chunked_policies"`
	FtpPolicies              types.List                  `tfsdk:"ftp_policies"`
	SmtpPolicies             types.List                  `tfsdk:"smtp_policies"`
	SftpPolicies             types.List                  `tfsdk:"sftp_policies"`
	DependencyActions        []*actions.DependencyAction `tfsdk:"dependency_actions"`
	ProviderTarget           types.String                `tfsdk:"provider_target"`
}

var HTTPUserAgentObjectType = map[string]attr.Type{
	"provider_target":            types.StringType,
	"id":                         types.StringType,
	"app_domain":                 types.StringType,
	"user_summary":               types.StringType,
	"identifier":                 types.StringType,
	"max_redirects":              types.Int64Type,
	"timeout":                    types.Int64Type,
	"proxy_policies":             types.ListType{ElemType: types.ObjectType{AttrTypes: DmProxyPolicyObjectType}},
	"ssl_policies":               types.ListType{ElemType: types.ObjectType{AttrTypes: DmSSLPolicyObjectType}},
	"basic_auth_policies":        types.ListType{ElemType: types.ObjectType{AttrTypes: DmBasicAuthPolicyObjectType}},
	"soap_action_policies":       types.ListType{ElemType: types.ObjectType{AttrTypes: DmSoapActionPolicyObjectType}},
	"pubkey_auth_policies":       types.ListType{ElemType: types.ObjectType{AttrTypes: DmPubkeyAuthPolicyObjectType}},
	"allow_compression_policies": types.ListType{ElemType: types.ObjectType{AttrTypes: DmAllowCompressionPolicyObjectType}},
	"header_retention_policies":  types.ListType{ElemType: types.ObjectType{AttrTypes: DmHeaderRetentionPolicyObjectType}},
	"http_version_policies":      types.ListType{ElemType: types.ObjectType{AttrTypes: DmHTTPVersionPolicyObjectType}},
	"add_header_policies":        types.ListType{ElemType: types.ObjectType{AttrTypes: DmAddHeaderPolicyObjectType}},
	"upload_chunked_policies":    types.ListType{ElemType: types.ObjectType{AttrTypes: DmUploadChunkedPolicyObjectType}},
	"ftp_policies":               types.ListType{ElemType: types.ObjectType{AttrTypes: DmFTPPolicyObjectType}},
	"smtp_policies":              types.ListType{ElemType: types.ObjectType{AttrTypes: DmSMTPPolicyObjectType}},
	"sftp_policies":              types.ListType{ElemType: types.ObjectType{AttrTypes: DmSFTPPolicyObjectType}},
	"dependency_actions":         actions.ActionsListType,
}

func (data HTTPUserAgent) GetPath() string {
	rest_path := "/mgmt/config/{domain}/HTTPUserAgent"
	rest_path = strings.ReplaceAll(rest_path, "{domain}", url.QueryEscape(data.AppDomain.ValueString()))
	return rest_path
}

func (data HTTPUserAgent) IsNull() bool {
	if !data.Id.IsNull() {
		return false
	}
	if !data.AppDomain.IsNull() {
		return false
	}
	if !data.UserSummary.IsNull() {
		return false
	}
	if !data.Identifier.IsNull() {
		return false
	}
	if !data.MaxRedirects.IsNull() {
		return false
	}
	if !data.Timeout.IsNull() {
		return false
	}
	if !data.ProxyPolicies.IsNull() {
		return false
	}
	if !data.SslPolicies.IsNull() {
		return false
	}
	if !data.BasicAuthPolicies.IsNull() {
		return false
	}
	if !data.SoapActionPolicies.IsNull() {
		return false
	}
	if !data.PubkeyAuthPolicies.IsNull() {
		return false
	}
	if !data.AllowCompressionPolicies.IsNull() {
		return false
	}
	if !data.HeaderRetentionPolicies.IsNull() {
		return false
	}
	if !data.HttpVersionPolicies.IsNull() {
		return false
	}
	if !data.AddHeaderPolicies.IsNull() {
		return false
	}
	if !data.UploadChunkedPolicies.IsNull() {
		return false
	}
	if !data.FtpPolicies.IsNull() {
		return false
	}
	if !data.SmtpPolicies.IsNull() {
		return false
	}
	if !data.SftpPolicies.IsNull() {
		return false
	}
	return true
}

func (data HTTPUserAgent) ToBody(ctx context.Context, pathRoot string) string {
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
	if !data.Identifier.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Identifier`, data.Identifier.ValueString())
	}
	if !data.MaxRedirects.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`MaxRedirects`, data.MaxRedirects.ValueInt64())
	}
	if !data.Timeout.IsNull() {
		body, _ = sjson.Set(body, pathRoot+`Timeout`, data.Timeout.ValueInt64())
	}
	if !data.ProxyPolicies.IsNull() {
		var dataValues []DmProxyPolicy
		data.ProxyPolicies.ElementsAs(ctx, &dataValues, false)
		if len(dataValues) > 0 {
			for _, val := range dataValues {
				body, _ = sjson.SetRaw(body, pathRoot+`ProxyPolicies`+".-1", val.ToBody(ctx, ""))
			}
		} else {
			body, _ = sjson.SetRaw(body, pathRoot+`ProxyPolicies`, "[]")
		}
	} else {
		body, _ = sjson.SetRaw(body, pathRoot+`ProxyPolicies`, "[]")
	}
	if !data.SslPolicies.IsNull() {
		var dataValues []DmSSLPolicy
		data.SslPolicies.ElementsAs(ctx, &dataValues, false)
		if len(dataValues) > 0 {
			for _, val := range dataValues {
				body, _ = sjson.SetRaw(body, pathRoot+`SSLPolicies`+".-1", val.ToBody(ctx, ""))
			}
		} else {
			body, _ = sjson.SetRaw(body, pathRoot+`SSLPolicies`, "[]")
		}
	} else {
		body, _ = sjson.SetRaw(body, pathRoot+`SSLPolicies`, "[]")
	}
	if !data.BasicAuthPolicies.IsNull() {
		var dataValues []DmBasicAuthPolicy
		data.BasicAuthPolicies.ElementsAs(ctx, &dataValues, false)
		if len(dataValues) > 0 {
			for _, val := range dataValues {
				body, _ = sjson.SetRaw(body, pathRoot+`BasicAuthPolicies`+".-1", val.ToBody(ctx, ""))
			}
		} else {
			body, _ = sjson.SetRaw(body, pathRoot+`BasicAuthPolicies`, "[]")
		}
	} else {
		body, _ = sjson.SetRaw(body, pathRoot+`BasicAuthPolicies`, "[]")
	}
	if !data.SoapActionPolicies.IsNull() {
		var dataValues []DmSoapActionPolicy
		data.SoapActionPolicies.ElementsAs(ctx, &dataValues, false)
		if len(dataValues) > 0 {
			for _, val := range dataValues {
				body, _ = sjson.SetRaw(body, pathRoot+`SoapActionPolicies`+".-1", val.ToBody(ctx, ""))
			}
		} else {
			body, _ = sjson.SetRaw(body, pathRoot+`SoapActionPolicies`, "[]")
		}
	} else {
		body, _ = sjson.SetRaw(body, pathRoot+`SoapActionPolicies`, "[]")
	}
	if !data.PubkeyAuthPolicies.IsNull() {
		var dataValues []DmPubkeyAuthPolicy
		data.PubkeyAuthPolicies.ElementsAs(ctx, &dataValues, false)
		if len(dataValues) > 0 {
			for _, val := range dataValues {
				body, _ = sjson.SetRaw(body, pathRoot+`PubkeyAuthPolicies`+".-1", val.ToBody(ctx, ""))
			}
		} else {
			body, _ = sjson.SetRaw(body, pathRoot+`PubkeyAuthPolicies`, "[]")
		}
	} else {
		body, _ = sjson.SetRaw(body, pathRoot+`PubkeyAuthPolicies`, "[]")
	}
	if !data.AllowCompressionPolicies.IsNull() {
		var dataValues []DmAllowCompressionPolicy
		data.AllowCompressionPolicies.ElementsAs(ctx, &dataValues, false)
		if len(dataValues) > 0 {
			for _, val := range dataValues {
				body, _ = sjson.SetRaw(body, pathRoot+`AllowCompressionPolicies`+".-1", val.ToBody(ctx, ""))
			}
		} else {
			body, _ = sjson.SetRaw(body, pathRoot+`AllowCompressionPolicies`, "[]")
		}
	} else {
		body, _ = sjson.SetRaw(body, pathRoot+`AllowCompressionPolicies`, "[]")
	}
	if !data.HeaderRetentionPolicies.IsNull() {
		var dataValues []DmHeaderRetentionPolicy
		data.HeaderRetentionPolicies.ElementsAs(ctx, &dataValues, false)
		if len(dataValues) > 0 {
			for _, val := range dataValues {
				body, _ = sjson.SetRaw(body, pathRoot+`HeaderRetentionPolicies`+".-1", val.ToBody(ctx, ""))
			}
		} else {
			body, _ = sjson.SetRaw(body, pathRoot+`HeaderRetentionPolicies`, "[]")
		}
	} else {
		body, _ = sjson.SetRaw(body, pathRoot+`HeaderRetentionPolicies`, "[]")
	}
	if !data.HttpVersionPolicies.IsNull() {
		var dataValues []DmHTTPVersionPolicy
		data.HttpVersionPolicies.ElementsAs(ctx, &dataValues, false)
		if len(dataValues) > 0 {
			for _, val := range dataValues {
				body, _ = sjson.SetRaw(body, pathRoot+`HTTPVersionPolicies`+".-1", val.ToBody(ctx, ""))
			}
		} else {
			body, _ = sjson.SetRaw(body, pathRoot+`HTTPVersionPolicies`, "[]")
		}
	} else {
		body, _ = sjson.SetRaw(body, pathRoot+`HTTPVersionPolicies`, "[]")
	}
	if !data.AddHeaderPolicies.IsNull() {
		var dataValues []DmAddHeaderPolicy
		data.AddHeaderPolicies.ElementsAs(ctx, &dataValues, false)
		if len(dataValues) > 0 {
			for _, val := range dataValues {
				body, _ = sjson.SetRaw(body, pathRoot+`AddHeaderPolicies`+".-1", val.ToBody(ctx, ""))
			}
		} else {
			body, _ = sjson.SetRaw(body, pathRoot+`AddHeaderPolicies`, "[]")
		}
	} else {
		body, _ = sjson.SetRaw(body, pathRoot+`AddHeaderPolicies`, "[]")
	}
	if !data.UploadChunkedPolicies.IsNull() {
		var dataValues []DmUploadChunkedPolicy
		data.UploadChunkedPolicies.ElementsAs(ctx, &dataValues, false)
		if len(dataValues) > 0 {
			for _, val := range dataValues {
				body, _ = sjson.SetRaw(body, pathRoot+`UploadChunkedPolicies`+".-1", val.ToBody(ctx, ""))
			}
		} else {
			body, _ = sjson.SetRaw(body, pathRoot+`UploadChunkedPolicies`, "[]")
		}
	} else {
		body, _ = sjson.SetRaw(body, pathRoot+`UploadChunkedPolicies`, "[]")
	}
	if !data.FtpPolicies.IsNull() {
		var dataValues []DmFTPPolicy
		data.FtpPolicies.ElementsAs(ctx, &dataValues, false)
		if len(dataValues) > 0 {
			for _, val := range dataValues {
				body, _ = sjson.SetRaw(body, pathRoot+`FTPPolicies`+".-1", val.ToBody(ctx, ""))
			}
		} else {
			body, _ = sjson.SetRaw(body, pathRoot+`FTPPolicies`, "[]")
		}
	} else {
		body, _ = sjson.SetRaw(body, pathRoot+`FTPPolicies`, "[]")
	}
	if !data.SmtpPolicies.IsNull() {
		var dataValues []DmSMTPPolicy
		data.SmtpPolicies.ElementsAs(ctx, &dataValues, false)
		if len(dataValues) > 0 {
			for _, val := range dataValues {
				body, _ = sjson.SetRaw(body, pathRoot+`SMTPPolicies`+".-1", val.ToBody(ctx, ""))
			}
		} else {
			body, _ = sjson.SetRaw(body, pathRoot+`SMTPPolicies`, "[]")
		}
	} else {
		body, _ = sjson.SetRaw(body, pathRoot+`SMTPPolicies`, "[]")
	}
	if !data.SftpPolicies.IsNull() {
		var dataValues []DmSFTPPolicy
		data.SftpPolicies.ElementsAs(ctx, &dataValues, false)
		if len(dataValues) > 0 {
			for _, val := range dataValues {
				body, _ = sjson.SetRaw(body, pathRoot+`SFTPPolicies`+".-1", val.ToBody(ctx, ""))
			}
		} else {
			body, _ = sjson.SetRaw(body, pathRoot+`SFTPPolicies`, "[]")
		}
	} else {
		body, _ = sjson.SetRaw(body, pathRoot+`SFTPPolicies`, "[]")
	}
	return body
}

func (data *HTTPUserAgent) FromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `Identifier`); value.Exists() && tfutils.ParseStringFromGJSON(value).ValueString() != "" {
		data.Identifier = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Identifier = types.StringNull()
	}
	if value := res.Get(pathRoot + `MaxRedirects`); value.Exists() {
		data.MaxRedirects = types.Int64Value(value.Int())
	} else {
		data.MaxRedirects = types.Int64Value(8)
	}
	if value := res.Get(pathRoot + `Timeout`); value.Exists() {
		data.Timeout = types.Int64Value(value.Int())
	} else {
		data.Timeout = types.Int64Value(300)
	}
	if value := res.Get(pathRoot + `ProxyPolicies`); value.Exists() {
		l := []DmProxyPolicy{}
		if value := res.Get(`ProxyPolicies`); value.Exists() {
			for _, v := range value.Array() {
				item := DmProxyPolicy{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.ProxyPolicies, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmProxyPolicyObjectType}, l)
		} else {
			data.ProxyPolicies = types.ListNull(types.ObjectType{AttrTypes: DmProxyPolicyObjectType})
		}
	} else {
		data.ProxyPolicies = types.ListNull(types.ObjectType{AttrTypes: DmProxyPolicyObjectType})
	}
	if value := res.Get(pathRoot + `SSLPolicies`); value.Exists() {
		l := []DmSSLPolicy{}
		if value := res.Get(`SSLPolicies`); value.Exists() {
			for _, v := range value.Array() {
				item := DmSSLPolicy{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.SslPolicies, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmSSLPolicyObjectType}, l)
		} else {
			data.SslPolicies = types.ListNull(types.ObjectType{AttrTypes: DmSSLPolicyObjectType})
		}
	} else {
		data.SslPolicies = types.ListNull(types.ObjectType{AttrTypes: DmSSLPolicyObjectType})
	}
	if value := res.Get(pathRoot + `BasicAuthPolicies`); value.Exists() {
		l := []DmBasicAuthPolicy{}
		if value := res.Get(`BasicAuthPolicies`); value.Exists() {
			for _, v := range value.Array() {
				item := DmBasicAuthPolicy{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.BasicAuthPolicies, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmBasicAuthPolicyObjectType}, l)
		} else {
			data.BasicAuthPolicies = types.ListNull(types.ObjectType{AttrTypes: DmBasicAuthPolicyObjectType})
		}
	} else {
		data.BasicAuthPolicies = types.ListNull(types.ObjectType{AttrTypes: DmBasicAuthPolicyObjectType})
	}
	if value := res.Get(pathRoot + `SoapActionPolicies`); value.Exists() {
		l := []DmSoapActionPolicy{}
		if value := res.Get(`SoapActionPolicies`); value.Exists() {
			for _, v := range value.Array() {
				item := DmSoapActionPolicy{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.SoapActionPolicies, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmSoapActionPolicyObjectType}, l)
		} else {
			data.SoapActionPolicies = types.ListNull(types.ObjectType{AttrTypes: DmSoapActionPolicyObjectType})
		}
	} else {
		data.SoapActionPolicies = types.ListNull(types.ObjectType{AttrTypes: DmSoapActionPolicyObjectType})
	}
	if value := res.Get(pathRoot + `PubkeyAuthPolicies`); value.Exists() {
		l := []DmPubkeyAuthPolicy{}
		if value := res.Get(`PubkeyAuthPolicies`); value.Exists() {
			for _, v := range value.Array() {
				item := DmPubkeyAuthPolicy{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.PubkeyAuthPolicies, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmPubkeyAuthPolicyObjectType}, l)
		} else {
			data.PubkeyAuthPolicies = types.ListNull(types.ObjectType{AttrTypes: DmPubkeyAuthPolicyObjectType})
		}
	} else {
		data.PubkeyAuthPolicies = types.ListNull(types.ObjectType{AttrTypes: DmPubkeyAuthPolicyObjectType})
	}
	if value := res.Get(pathRoot + `AllowCompressionPolicies`); value.Exists() {
		l := []DmAllowCompressionPolicy{}
		if value := res.Get(`AllowCompressionPolicies`); value.Exists() {
			for _, v := range value.Array() {
				item := DmAllowCompressionPolicy{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.AllowCompressionPolicies, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmAllowCompressionPolicyObjectType}, l)
		} else {
			data.AllowCompressionPolicies = types.ListNull(types.ObjectType{AttrTypes: DmAllowCompressionPolicyObjectType})
		}
	} else {
		data.AllowCompressionPolicies = types.ListNull(types.ObjectType{AttrTypes: DmAllowCompressionPolicyObjectType})
	}
	if value := res.Get(pathRoot + `HeaderRetentionPolicies`); value.Exists() {
		l := []DmHeaderRetentionPolicy{}
		if value := res.Get(`HeaderRetentionPolicies`); value.Exists() {
			for _, v := range value.Array() {
				item := DmHeaderRetentionPolicy{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.HeaderRetentionPolicies, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmHeaderRetentionPolicyObjectType}, l)
		} else {
			data.HeaderRetentionPolicies = types.ListNull(types.ObjectType{AttrTypes: DmHeaderRetentionPolicyObjectType})
		}
	} else {
		data.HeaderRetentionPolicies = types.ListNull(types.ObjectType{AttrTypes: DmHeaderRetentionPolicyObjectType})
	}
	if value := res.Get(pathRoot + `HTTPVersionPolicies`); value.Exists() {
		l := []DmHTTPVersionPolicy{}
		if value := res.Get(`HTTPVersionPolicies`); value.Exists() {
			for _, v := range value.Array() {
				item := DmHTTPVersionPolicy{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.HttpVersionPolicies, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmHTTPVersionPolicyObjectType}, l)
		} else {
			data.HttpVersionPolicies = types.ListNull(types.ObjectType{AttrTypes: DmHTTPVersionPolicyObjectType})
		}
	} else {
		data.HttpVersionPolicies = types.ListNull(types.ObjectType{AttrTypes: DmHTTPVersionPolicyObjectType})
	}
	if value := res.Get(pathRoot + `AddHeaderPolicies`); value.Exists() {
		l := []DmAddHeaderPolicy{}
		if value := res.Get(`AddHeaderPolicies`); value.Exists() {
			for _, v := range value.Array() {
				item := DmAddHeaderPolicy{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.AddHeaderPolicies, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmAddHeaderPolicyObjectType}, l)
		} else {
			data.AddHeaderPolicies = types.ListNull(types.ObjectType{AttrTypes: DmAddHeaderPolicyObjectType})
		}
	} else {
		data.AddHeaderPolicies = types.ListNull(types.ObjectType{AttrTypes: DmAddHeaderPolicyObjectType})
	}
	if value := res.Get(pathRoot + `UploadChunkedPolicies`); value.Exists() {
		l := []DmUploadChunkedPolicy{}
		if value := res.Get(`UploadChunkedPolicies`); value.Exists() {
			for _, v := range value.Array() {
				item := DmUploadChunkedPolicy{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.UploadChunkedPolicies, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmUploadChunkedPolicyObjectType}, l)
		} else {
			data.UploadChunkedPolicies = types.ListNull(types.ObjectType{AttrTypes: DmUploadChunkedPolicyObjectType})
		}
	} else {
		data.UploadChunkedPolicies = types.ListNull(types.ObjectType{AttrTypes: DmUploadChunkedPolicyObjectType})
	}
	if value := res.Get(pathRoot + `FTPPolicies`); value.Exists() {
		l := []DmFTPPolicy{}
		if value := res.Get(`FTPPolicies`); value.Exists() {
			for _, v := range value.Array() {
				item := DmFTPPolicy{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.FtpPolicies, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmFTPPolicyObjectType}, l)
		} else {
			data.FtpPolicies = types.ListNull(types.ObjectType{AttrTypes: DmFTPPolicyObjectType})
		}
	} else {
		data.FtpPolicies = types.ListNull(types.ObjectType{AttrTypes: DmFTPPolicyObjectType})
	}
	if value := res.Get(pathRoot + `SMTPPolicies`); value.Exists() {
		l := []DmSMTPPolicy{}
		if value := res.Get(`SMTPPolicies`); value.Exists() {
			for _, v := range value.Array() {
				item := DmSMTPPolicy{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.SmtpPolicies, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmSMTPPolicyObjectType}, l)
		} else {
			data.SmtpPolicies = types.ListNull(types.ObjectType{AttrTypes: DmSMTPPolicyObjectType})
		}
	} else {
		data.SmtpPolicies = types.ListNull(types.ObjectType{AttrTypes: DmSMTPPolicyObjectType})
	}
	if value := res.Get(pathRoot + `SFTPPolicies`); value.Exists() {
		l := []DmSFTPPolicy{}
		if value := res.Get(`SFTPPolicies`); value.Exists() {
			for _, v := range value.Array() {
				item := DmSFTPPolicy{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.SftpPolicies, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmSFTPPolicyObjectType}, l)
		} else {
			data.SftpPolicies = types.ListNull(types.ObjectType{AttrTypes: DmSFTPPolicyObjectType})
		}
	} else {
		data.SftpPolicies = types.ListNull(types.ObjectType{AttrTypes: DmSFTPPolicyObjectType})
	}
}

func (data *HTTPUserAgent) UpdateFromBody(ctx context.Context, pathRoot string, res gjson.Result) {
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
	if value := res.Get(pathRoot + `Identifier`); value.Exists() && !data.Identifier.IsNull() {
		data.Identifier = tfutils.ParseStringFromGJSON(value)
	} else {
		data.Identifier = types.StringNull()
	}
	if value := res.Get(pathRoot + `MaxRedirects`); value.Exists() && !data.MaxRedirects.IsNull() {
		data.MaxRedirects = types.Int64Value(value.Int())
	} else if data.MaxRedirects.ValueInt64() != 8 {
		data.MaxRedirects = types.Int64Null()
	}
	if value := res.Get(pathRoot + `Timeout`); value.Exists() && !data.Timeout.IsNull() {
		data.Timeout = types.Int64Value(value.Int())
	} else if data.Timeout.ValueInt64() != 300 {
		data.Timeout = types.Int64Null()
	}
	if value := res.Get(pathRoot + `ProxyPolicies`); value.Exists() && !data.ProxyPolicies.IsNull() {
		l := []DmProxyPolicy{}
		e := []DmProxyPolicy{}
		data.ProxyPolicies.ElementsAs(ctx, &e, false)
		if len(value.Array()) == len(e) {
			for i, v := range value.Array() {
				item := e[i]
				item.UpdateFromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		} else {
			for _, v := range value.Array() {
				item := DmProxyPolicy{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.ProxyPolicies, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmProxyPolicyObjectType}, l)
		} else {
			data.ProxyPolicies = types.ListNull(types.ObjectType{AttrTypes: DmProxyPolicyObjectType})
		}
	} else {
		data.ProxyPolicies = types.ListNull(types.ObjectType{AttrTypes: DmProxyPolicyObjectType})
	}
	if value := res.Get(pathRoot + `SSLPolicies`); value.Exists() && !data.SslPolicies.IsNull() {
		l := []DmSSLPolicy{}
		e := []DmSSLPolicy{}
		data.SslPolicies.ElementsAs(ctx, &e, false)
		if len(value.Array()) == len(e) {
			for i, v := range value.Array() {
				item := e[i]
				item.UpdateFromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		} else {
			for _, v := range value.Array() {
				item := DmSSLPolicy{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.SslPolicies, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmSSLPolicyObjectType}, l)
		} else {
			data.SslPolicies = types.ListNull(types.ObjectType{AttrTypes: DmSSLPolicyObjectType})
		}
	} else {
		data.SslPolicies = types.ListNull(types.ObjectType{AttrTypes: DmSSLPolicyObjectType})
	}
	if value := res.Get(pathRoot + `BasicAuthPolicies`); value.Exists() && !data.BasicAuthPolicies.IsNull() {
		l := []DmBasicAuthPolicy{}
		e := []DmBasicAuthPolicy{}
		data.BasicAuthPolicies.ElementsAs(ctx, &e, false)
		if len(value.Array()) == len(e) {
			for i, v := range value.Array() {
				item := e[i]
				item.UpdateFromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		} else {
			for _, v := range value.Array() {
				item := DmBasicAuthPolicy{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.BasicAuthPolicies, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmBasicAuthPolicyObjectType}, l)
		} else {
			data.BasicAuthPolicies = types.ListNull(types.ObjectType{AttrTypes: DmBasicAuthPolicyObjectType})
		}
	} else {
		data.BasicAuthPolicies = types.ListNull(types.ObjectType{AttrTypes: DmBasicAuthPolicyObjectType})
	}
	if value := res.Get(pathRoot + `SoapActionPolicies`); value.Exists() && !data.SoapActionPolicies.IsNull() {
		l := []DmSoapActionPolicy{}
		e := []DmSoapActionPolicy{}
		data.SoapActionPolicies.ElementsAs(ctx, &e, false)
		if len(value.Array()) == len(e) {
			for i, v := range value.Array() {
				item := e[i]
				item.UpdateFromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		} else {
			for _, v := range value.Array() {
				item := DmSoapActionPolicy{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.SoapActionPolicies, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmSoapActionPolicyObjectType}, l)
		} else {
			data.SoapActionPolicies = types.ListNull(types.ObjectType{AttrTypes: DmSoapActionPolicyObjectType})
		}
	} else {
		data.SoapActionPolicies = types.ListNull(types.ObjectType{AttrTypes: DmSoapActionPolicyObjectType})
	}
	if value := res.Get(pathRoot + `PubkeyAuthPolicies`); value.Exists() && !data.PubkeyAuthPolicies.IsNull() {
		l := []DmPubkeyAuthPolicy{}
		e := []DmPubkeyAuthPolicy{}
		data.PubkeyAuthPolicies.ElementsAs(ctx, &e, false)
		if len(value.Array()) == len(e) {
			for i, v := range value.Array() {
				item := e[i]
				item.UpdateFromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		} else {
			for _, v := range value.Array() {
				item := DmPubkeyAuthPolicy{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.PubkeyAuthPolicies, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmPubkeyAuthPolicyObjectType}, l)
		} else {
			data.PubkeyAuthPolicies = types.ListNull(types.ObjectType{AttrTypes: DmPubkeyAuthPolicyObjectType})
		}
	} else {
		data.PubkeyAuthPolicies = types.ListNull(types.ObjectType{AttrTypes: DmPubkeyAuthPolicyObjectType})
	}
	if value := res.Get(pathRoot + `AllowCompressionPolicies`); value.Exists() && !data.AllowCompressionPolicies.IsNull() {
		l := []DmAllowCompressionPolicy{}
		e := []DmAllowCompressionPolicy{}
		data.AllowCompressionPolicies.ElementsAs(ctx, &e, false)
		if len(value.Array()) == len(e) {
			for i, v := range value.Array() {
				item := e[i]
				item.UpdateFromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		} else {
			for _, v := range value.Array() {
				item := DmAllowCompressionPolicy{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.AllowCompressionPolicies, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmAllowCompressionPolicyObjectType}, l)
		} else {
			data.AllowCompressionPolicies = types.ListNull(types.ObjectType{AttrTypes: DmAllowCompressionPolicyObjectType})
		}
	} else {
		data.AllowCompressionPolicies = types.ListNull(types.ObjectType{AttrTypes: DmAllowCompressionPolicyObjectType})
	}
	if value := res.Get(pathRoot + `HeaderRetentionPolicies`); value.Exists() && !data.HeaderRetentionPolicies.IsNull() {
		l := []DmHeaderRetentionPolicy{}
		e := []DmHeaderRetentionPolicy{}
		data.HeaderRetentionPolicies.ElementsAs(ctx, &e, false)
		if len(value.Array()) == len(e) {
			for i, v := range value.Array() {
				item := e[i]
				item.UpdateFromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		} else {
			for _, v := range value.Array() {
				item := DmHeaderRetentionPolicy{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.HeaderRetentionPolicies, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmHeaderRetentionPolicyObjectType}, l)
		} else {
			data.HeaderRetentionPolicies = types.ListNull(types.ObjectType{AttrTypes: DmHeaderRetentionPolicyObjectType})
		}
	} else {
		data.HeaderRetentionPolicies = types.ListNull(types.ObjectType{AttrTypes: DmHeaderRetentionPolicyObjectType})
	}
	if value := res.Get(pathRoot + `HTTPVersionPolicies`); value.Exists() && !data.HttpVersionPolicies.IsNull() {
		l := []DmHTTPVersionPolicy{}
		e := []DmHTTPVersionPolicy{}
		data.HttpVersionPolicies.ElementsAs(ctx, &e, false)
		if len(value.Array()) == len(e) {
			for i, v := range value.Array() {
				item := e[i]
				item.UpdateFromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		} else {
			for _, v := range value.Array() {
				item := DmHTTPVersionPolicy{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.HttpVersionPolicies, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmHTTPVersionPolicyObjectType}, l)
		} else {
			data.HttpVersionPolicies = types.ListNull(types.ObjectType{AttrTypes: DmHTTPVersionPolicyObjectType})
		}
	} else {
		data.HttpVersionPolicies = types.ListNull(types.ObjectType{AttrTypes: DmHTTPVersionPolicyObjectType})
	}
	if value := res.Get(pathRoot + `AddHeaderPolicies`); value.Exists() && !data.AddHeaderPolicies.IsNull() {
		l := []DmAddHeaderPolicy{}
		e := []DmAddHeaderPolicy{}
		data.AddHeaderPolicies.ElementsAs(ctx, &e, false)
		if len(value.Array()) == len(e) {
			for i, v := range value.Array() {
				item := e[i]
				item.UpdateFromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		} else {
			for _, v := range value.Array() {
				item := DmAddHeaderPolicy{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.AddHeaderPolicies, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmAddHeaderPolicyObjectType}, l)
		} else {
			data.AddHeaderPolicies = types.ListNull(types.ObjectType{AttrTypes: DmAddHeaderPolicyObjectType})
		}
	} else {
		data.AddHeaderPolicies = types.ListNull(types.ObjectType{AttrTypes: DmAddHeaderPolicyObjectType})
	}
	if value := res.Get(pathRoot + `UploadChunkedPolicies`); value.Exists() && !data.UploadChunkedPolicies.IsNull() {
		l := []DmUploadChunkedPolicy{}
		e := []DmUploadChunkedPolicy{}
		data.UploadChunkedPolicies.ElementsAs(ctx, &e, false)
		if len(value.Array()) == len(e) {
			for i, v := range value.Array() {
				item := e[i]
				item.UpdateFromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		} else {
			for _, v := range value.Array() {
				item := DmUploadChunkedPolicy{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.UploadChunkedPolicies, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmUploadChunkedPolicyObjectType}, l)
		} else {
			data.UploadChunkedPolicies = types.ListNull(types.ObjectType{AttrTypes: DmUploadChunkedPolicyObjectType})
		}
	} else {
		data.UploadChunkedPolicies = types.ListNull(types.ObjectType{AttrTypes: DmUploadChunkedPolicyObjectType})
	}
	if value := res.Get(pathRoot + `FTPPolicies`); value.Exists() && !data.FtpPolicies.IsNull() {
		l := []DmFTPPolicy{}
		e := []DmFTPPolicy{}
		data.FtpPolicies.ElementsAs(ctx, &e, false)
		if len(value.Array()) == len(e) {
			for i, v := range value.Array() {
				item := e[i]
				item.UpdateFromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		} else {
			for _, v := range value.Array() {
				item := DmFTPPolicy{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.FtpPolicies, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmFTPPolicyObjectType}, l)
		} else {
			data.FtpPolicies = types.ListNull(types.ObjectType{AttrTypes: DmFTPPolicyObjectType})
		}
	} else {
		data.FtpPolicies = types.ListNull(types.ObjectType{AttrTypes: DmFTPPolicyObjectType})
	}
	if value := res.Get(pathRoot + `SMTPPolicies`); value.Exists() && !data.SmtpPolicies.IsNull() {
		l := []DmSMTPPolicy{}
		e := []DmSMTPPolicy{}
		data.SmtpPolicies.ElementsAs(ctx, &e, false)
		if len(value.Array()) == len(e) {
			for i, v := range value.Array() {
				item := e[i]
				item.UpdateFromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		} else {
			for _, v := range value.Array() {
				item := DmSMTPPolicy{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.SmtpPolicies, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmSMTPPolicyObjectType}, l)
		} else {
			data.SmtpPolicies = types.ListNull(types.ObjectType{AttrTypes: DmSMTPPolicyObjectType})
		}
	} else {
		data.SmtpPolicies = types.ListNull(types.ObjectType{AttrTypes: DmSMTPPolicyObjectType})
	}
	if value := res.Get(pathRoot + `SFTPPolicies`); value.Exists() && !data.SftpPolicies.IsNull() {
		l := []DmSFTPPolicy{}
		e := []DmSFTPPolicy{}
		data.SftpPolicies.ElementsAs(ctx, &e, false)
		if len(value.Array()) == len(e) {
			for i, v := range value.Array() {
				item := e[i]
				item.UpdateFromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		} else {
			for _, v := range value.Array() {
				item := DmSFTPPolicy{}
				item.FromBody(ctx, "", v)
				if !item.IsNull() {
					l = append(l, item)
				}
			}
		}
		if len(l) > 0 {
			data.SftpPolicies, _ = types.ListValueFrom(ctx, types.ObjectType{AttrTypes: DmSFTPPolicyObjectType}, l)
		} else {
			data.SftpPolicies = types.ListNull(types.ObjectType{AttrTypes: DmSFTPPolicyObjectType})
		}
	} else {
		data.SftpPolicies = types.ListNull(types.ObjectType{AttrTypes: DmSFTPPolicyObjectType})
	}
}

resource "datapower_file" "mpgw-proxy-url-xml" {
  app_domain  = datapower_domain.mpgw.app_domain
  remote_path = "local:///Hello.xml"
  local_path  = "./files/HelloWorld.xml"
}

resource "datapower_file" "mpgw-proxy-url-xslt" {
  app_domain  = datapower_domain.mpgw.app_domain
  remote_path = "local:///Hello.xslt"
  local_path  = "./files/HelloWorld.xslt"
}

resource "datapower_matching" "mpgw" {
  id         = "mpgw-matching"
  app_domain = datapower_domain.mpgw.app_domain
  match_rules = [
    {
      type = "url"
      url  = "*"
    },
    {
      type   = "method"
      method = "GET"
    }
  ]
}

resource "datapower_stylepolicyaction" "mpgw_fetch_0" {
  id          = "mpgw_fetch_0"
  app_domain  = datapower_domain.mpgw.app_domain
  type        = "fetch"
  destination = datapower_file.mpgw-proxy-url-xml.remote_path
  input       = "INPUT"
  output      = "PIPE"
}

resource "datapower_stylepolicyaction" "mpgw_transform_1" {
  id         = "mpgw_transform_1"
  app_domain = datapower_domain.mpgw.app_domain
  type       = "xform"
  transform  = datapower_file.mpgw-proxy-url-xslt.remote_path
  input      = "PIPE"
  output     = "dpvar_1"
}

resource "datapower_stylepolicyaction" "mpgw_setvar_2" {
  id         = "mpgw_setvar_2"
  app_domain = datapower_domain.mpgw.app_domain
  type       = "setvar"
  variable   = "var://service/mpgw/skip-backside"
  value      = "1"
  input      = "dpvar_1"
}

resource "datapower_stylepolicyaction" "mpgw_output_3" {
  id         = "mpgw_output_3"
  app_domain = datapower_domain.mpgw.app_domain
  type       = "results"
  input      = "dpvar_1"
}

resource "datapower_stylepolicyrule" "mpgw" {
  id         = "mpgw-style-policy-rule"
  app_domain = datapower_domain.mpgw.app_domain
  actions = [
    datapower_stylepolicyaction.mpgw_fetch_0.id,
    datapower_stylepolicyaction.mpgw_transform_1.id,
    datapower_stylepolicyaction.mpgw_setvar_2.id,
    datapower_stylepolicyaction.mpgw_output_3.id
  ]
}

resource "datapower_stylepolicy" "mpgw" {
  id         = "mpgw-style-policy"
  app_domain = datapower_domain.mpgw.app_domain
  policy_maps = [
    {
      match = datapower_matching.mpgw.id
      rule  = datapower_stylepolicyrule.mpgw.id
    }
  ]
}

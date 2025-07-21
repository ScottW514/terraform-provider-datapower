---
subcategory: "Guides"
page_title: "Getting Started"
description: |-
    Getting Started
---

# Getting Started

This example demonstrates how the provider can be used to configure simple multiprotocol gateway. 

The full example can be found here: [https://github.com/ScottW514/terraform-provider-datapower/tree/main/examples/basic/getting_started](https://github.com/ScottW514/terraform-provider-datapower/tree/main/examples/basic/getting_started)

## Provider Configuration

```hcl
terraform {
  required_providers {
    datapower = {
      source  = "scottw514/datapower"
    }
  }
}

provider "datapower" {
  hostname  = "datapower-dev"
  username  = "admin"
  password  = "admin"
}
```

## Style Policy

This creates the style policy that describes how the MPGW will behave.

First, we'll add an XML file and an XLST to transform it with:

```hcl
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
```

Next, we create our policy action steps, and add them to a policy rule:

```hcl

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
```

Next, we create a matching rule, and attach the rule and actions to a style policy:

```hcl
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
```

### Create the Multiprotocol Gateway

Finally, well create a domain, an HTTP protocol handler, and attach them and style policy to a MPGW:

```hcl
resource "datapower_domain" "mpgw" {
  app_domain = "basic_example"
}

resource "datapower_httpsourceprotocolhandler" "mpgw" {
  id         = "http_sph"
  app_domain = datapower_domain.mpgw.app_domain
  local_port = 8899
  allowed_features = {
    post                 = false
    get                  = true
    put                  = false
    query_string         = false
    fragment_identifiers = false
  }
}

resource "datapower_multiprotocolgateway" "mpgw" {
  id         = "mpgw"
  app_domain = datapower_domain.mpgw.app_domain
  front_protocol = [
    datapower_httpsourceprotocolhandler.mpgw.id
  ]
  style_policy  = datapower_stylepolicy.mpgw.id
  type          = "dynamic-backend"
  request_type  = "preprocessed"
  response_type = "preprocessed"
}
```

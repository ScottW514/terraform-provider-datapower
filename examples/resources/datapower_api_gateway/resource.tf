
resource "datapower_api_gateway" "test" {
  id             = "ResTestAPIGateway"
  app_domain     = "acceptance_test"
  front_protocol = ["AccTest_HTTPSourceProtocolHandler"]
  doc_cache_policy = [{
    type     = "protocol"
    priority = 128
  }]
  scheduled_rule = [{
    rule = "__dp-policy-begin__"
  }]
  assembly_burst_limit = [{
    name  = "apiburstlimit"
    burst = 1000
  }]
  assembly_rate_limit = [{
    name = "apiratelimit"
    rate = 1000
  }]
  assembly_count_limit = [{
    name  = "countlimit"
    count = 1000
  }]
  proxy_policies = [{
    reg_exp        = "*"
    skip           = false
    remote_address = "10.10.10.10"
    remote_port    = 8888
  }]
  front_timeout            = 120
  front_persistent_timeout = 180
  open_telemetry_resource_attribute = [{
    value = "value"
  }]
}
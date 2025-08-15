resource "datapower_domain" "mpgw" {
  app_domain = "basic_example"
}

resource "datapower_http_source_protocol_handler" "mpgw" {
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

resource "datapower_multi_protocol_gateway" "mpgw" {
  id         = "mpgw"
  app_domain = datapower_domain.mpgw.app_domain
  front_protocol = [
    datapower_http_source_protocol_handler.mpgw.id
  ]
  style_policy  = datapower_style_policy.mpgw.id
  type          = "dynamic-backend"
  request_type  = "preprocessed"
  response_type = "preprocessed"
}

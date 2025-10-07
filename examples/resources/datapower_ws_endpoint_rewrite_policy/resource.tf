
resource "datapower_ws_endpoint_rewrite_policy" "test" {
  id         = "ResTestWSEndpointRewritePolicy"
  app_domain = "acceptance_test"
  ws_endpoint_local_rewrite_rule = [{
    local_endpoint_hostname = "0.0.0.0"
  }]
  ws_endpoint_remote_rewrite_rule = [{
    remote_endpoint_hostname = "10.10.10.10"
  }]
  ws_endpoint_publish_rewrite_rule = [{
  }]
  ws_endpoint_subscription_local_rule = [{
    local_endpoint_hostname = "0.0.0.0"
  }]
  ws_endpoint_subscription_remote_rule = [{
    remote_endpoint_hostname = "10.10.10.10"
  }]
  ws_endpoint_subscription_publish_rule = [{
  }]
}
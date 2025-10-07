
resource "datapower_style_policy" "test" {
  id         = "ResTestStylePolicy"
  app_domain = "acceptance_test"
  policy_maps = [{
    match = "__default-accept-service-providers__"
    rule  = "__dp-policy-begin__"
  }]
}
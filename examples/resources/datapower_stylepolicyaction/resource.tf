
resource "datapower_stylepolicyaction" "test" {
  id            = "ResTestStylePolicyAction"
  app_domain    = "acceptance_test"
  type          = "xform"
  input         = "INPUT"
  transform     = "tfile"
  output        = "OUTPUT"
  named_inputs  = null
  named_outputs = null
}
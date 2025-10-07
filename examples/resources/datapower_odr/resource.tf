
resource "datapower_odr" "test" {
  odr_server_name = "dp_set"
  odr_custom_properties = [{
    prop_name  = "propname"
    prop_value = "propvalue"
  }]
}
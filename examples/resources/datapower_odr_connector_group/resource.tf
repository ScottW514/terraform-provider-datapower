
resource "datapower_odr_connector_group" "test" {
  id = "ResTestODRConnectorGroup"
  odr_group_connectors = [{
    dmgr_hostname = "localhost"
    dmgr_port     = 8888
  }]
  xml_manager = "default"
  odr_conn_group_properties = [{
    conn_group_prop_name  = "propname"
    conn_group_prop_value = "value"
  }]
}
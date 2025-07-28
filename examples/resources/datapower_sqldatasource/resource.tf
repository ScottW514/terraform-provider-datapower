
resource "datapower_sqldatasource" "test" {
  id               = "SQLDataSource_name"
  app_domain       = "acc_test_domain"
  database         = "Oracle"
  username         = "username"
  password_alias   = "TestAccPasswordAlias"
  data_source_id   = "datasource_id"
  data_source_host = "datasource.host"
  data_source_port = 1488
  max_connection   = 10
  connect_timeout  = 15
  query_timeout    = 30
  idle_timeout     = 180
}
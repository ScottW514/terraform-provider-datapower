
resource "datapower_sql_data_source" "test" {
  id               = "ResTestSQLDataSource"
  app_domain       = "acceptance_test"
  database         = "Oracle"
  username         = "username"
  password_alias   = "AccTest_PasswordAlias"
  data_source_id   = "datasource_id"
  data_source_host = "datasource.host"
  data_source_port = 1488
  max_connection   = 10
  connect_timeout  = 15
  query_timeout    = 30
  idle_timeout     = 180
}

resource "datapower_lunapartition" "test" {
  id               = "LunaPartition_name"
  app_domain       = "acc_test_domain"
  partition_name   = "partitionname"
  partition_serial = "serial"
  password_alias   = "TestAccPasswordAlias"
  login_role       = "co"
}
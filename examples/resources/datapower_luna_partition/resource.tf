
resource "datapower_luna_partition" "test" {
  id               = "ResTestLunaPartition"
  app_domain       = "acceptance_test"
  partition_name   = "partitionname"
  partition_serial = "serial"
  password_alias   = "AccTest_PasswordAlias"
  login_role       = "co"
}
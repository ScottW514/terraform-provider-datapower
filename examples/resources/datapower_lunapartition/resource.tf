
resource "datapower_lunapartition" "test" {
  id               = "LunaPartition_name"
  app_domain       = "acc_test_domain"
  partition_name   = "partitionname"
  partition_serial = "serial"
  password_alias   = datapower_passwordalias.test.id
  login_role       = "co"
}
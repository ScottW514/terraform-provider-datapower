
resource "datapower_b2bpersistence" "test" {
  raid_volume  = "raid0"
  storage_size = 1024
  ha_enabled   = false
}
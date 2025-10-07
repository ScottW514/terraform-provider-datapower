
resource "datapower_b2b_persistence" "test" {
  raid_volume  = "raid0"
  storage_size = 1024
  ha_enabled   = false
  ha_other_hosts = {
  }
}
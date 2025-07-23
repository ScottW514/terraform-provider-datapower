
resource "datapower_filesystemusagemonitor" "test" {
  polling_interval = 60
  all_system       = true
}

resource "datapower_dnsnameservice" "test" {
  load_balance_algorithm = "first-alive"
}
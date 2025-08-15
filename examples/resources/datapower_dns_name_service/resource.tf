
resource "datapower_dns_name_service" "test" {
  load_balance_algorithm = "first-alive"
}
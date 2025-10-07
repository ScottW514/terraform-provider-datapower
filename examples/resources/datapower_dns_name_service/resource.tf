
resource "datapower_dns_name_service" "test" {
  search_domains = [{
    search_domain = "datapower.com"
  }]
  name_servers = [{
    udp_port = 53
    tcp_port = 53
  }]
  static_hosts = [{
  }]
  load_balance_algorithm = "first-alive"
}
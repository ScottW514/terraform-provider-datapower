
resource "datapower_appsecuritypolicy" "test" {
  id         = "AppSecurityPolicy_test"
  app_domain = "acc_test_domain"
  request_maps = [{
    match = datapower_matching.test.id
    rule  = datapower_webapprequest.test.id
  }]
  response_maps = [{
    match = datapower_matching.test.id
    rule  = datapower_webappresponse.test.id
  }]
}
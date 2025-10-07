
resource "datapower_web_app_request" "test" {
  id         = "ResTest_WebAppRequest"
  app_domain = "acceptance_test"
  ok_methods = {
  }
  ok_versions = {
  }
  content_types = ["application/www-url-encoded", ]
  multipart_form_data = {
  }
  cookie_profile = {
  }
}
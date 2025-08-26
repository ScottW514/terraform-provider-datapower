
resource "datapower_ws_rr_saved_search_subscription" "test" {
  id                = "ResTestWSRRSavedSearchSubscription"
  app_domain        = "acceptance_test"
  server            = "AccTest_WSRRServer"
  saved_search_name = "ResTestsaved_search"
}
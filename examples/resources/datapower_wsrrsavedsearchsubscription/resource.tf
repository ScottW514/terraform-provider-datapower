
resource "datapower_wsrrsavedsearchsubscription" "test" {
  id                = "ResTestWSRRSavedSearchSubscription"
  app_domain        = "acceptance_test"
  server            = "AccTest_WSRRServer"
  saved_search_name = "ResTestsaved_search"
}
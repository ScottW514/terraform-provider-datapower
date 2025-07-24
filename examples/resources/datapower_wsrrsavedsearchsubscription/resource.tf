
resource "datapower_wsrrsavedsearchsubscription" "test" {
  id                = "WSRRSavedSearchSubscription_name"
  app_domain        = "acc_test_domain"
  server            = datapower_wsrrserver.test.id
  saved_search_name = "saved_search_name"
}
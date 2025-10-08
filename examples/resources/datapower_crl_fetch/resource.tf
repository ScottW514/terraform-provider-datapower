
resource "datapower_crl_fetch" "test" {
  crl_fetch_config = [{
    name             = "crlfetch"
    fetch_type       = "http"
    issuer_valcred   = "DefaultTest_CryptoValCred"
    refresh_interval = 240
    url              = "http://foo.com/crl"
  }]
}
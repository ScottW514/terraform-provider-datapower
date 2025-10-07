
resource "datapower_http_user_agent" "test" {
  id         = "ResTestHTTPUserAgent"
  app_domain = "acceptance_test"
  proxy_policies = [{
    reg_exp        = "*"
    skip           = false
    remote_address = "remote.host"
    remote_port    = 443
  }]
  ssl_policies = [{
    reg_exp    = "*"
    ssl_client = "AccTest_SSLClientProfile"
  }]
  basic_auth_policies = [{
    reg_exp   = "*"
    user_name = "someuser"
  }]
  soap_action_policies = [{
    reg_exp     = "*"
    soap_action = "*"
  }]
  pubkey_auth_policies = [{
    reg_exp    = "*"
    crypto_key = "AccTest_CryptoKey"
  }]
  allow_compression_policies = [{
    reg_exp           = "*"
    allow_compression = false
  }]
  header_retention_policies = [{
    reg_exp = "*"
    header_retention = {
    }
  }]
  http_version_policies = [{
    reg_exp = "*"
    version = "HTTP/1.1"
  }]
  add_header_policies = [{
    reg_exp    = "*"
    add_header = "HEADER"
    add_value  = "VALUE"
  }]
  upload_chunked_policies = [{
    reg_exp        = "*"
    upload_chunked = false
  }]
  ftp_policies = [{
    reg_exp    = "*"
    passive    = "pasv-req"
    auth_tls   = "auth-off"
    data_type  = "binary"
    slash_stou = "slash-stou-on"
  }]
  smtp_policies = [{
    reg_exp = "dpsmtp://*"
    options = {
    }
  }]
  sftp_policies = [{
    reg_exp              = "*"
    ssh_client_profile   = "AccTest_SSHClientProfile"
    use_unique_filenames = false
  }]
}
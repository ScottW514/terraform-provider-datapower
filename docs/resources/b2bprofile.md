---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "datapower_b2bprofile Resource - terraform-provider-datapower"
subcategory: ""
description: |-
  B2B partner profile
  CLI Alias: b2b-profile
---

# datapower_b2bprofile (Resource)

B2B partner profile
  - CLI Alias: `b2b-profile`

## Example Usage

```terraform
resource "datapower_b2bprofile" "test" {
  id         = "B2BProfile_name"
  app_domain = "acc_test_domain"
  destinations = [{
    dest_name = "b2bdestinationname"
    dest_url  = "https://localhost"
  }]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `app_domain` (String) The name of the application domain the object belongs to
- `destinations` (Attributes List) Destinations
  - CLI Alias: `destination` (see [below for nested schema](#nestedatt--destinations))
- `id` (String) Name of the object. Must be unique among object types in application domain.

### Optional

- `as_allow_duplicate_message` (String) Allow duplicate AS inbound message
  - CLI Alias: `as-allow-dup-msg`
  - Choices: `never`, `always`, `on-error`
  - Default value: `never`
- `business_i_ds` (List of String) Partner business IDs
  - CLI Alias: `business-id`
- `business_i_ds_duns` (List of String) Partner business IDs (DUNS)
  - CLI Alias: `business-id-duns`
- `business_i_ds_duns_plus4` (List of String) Partner business IDs (DUNS+4)
  - CLI Alias: `business-id-duns4`
- `contacts` (Attributes List) Contacts
  - CLI Alias: `contact` (see [below for nested schema](#nestedatt--contacts))
- `custom_style_policy` (String) Processing policy
  - CLI Alias: `stylepolicy`
  - Reference to: `datapower_stylepolicy:id`
- `ebms3_allow_duplicate_message` (String) Allow duplicate ebMS3 inbound message
  - CLI Alias: `ebms3-allow-dup-msg`
  - Choices: `never`, `always`, `on-error`
  - Default value: `never`
- `ebms3_default_signer_cert` (String) Default inbound signature validation certificate
  - CLI Alias: `ebms3-default-signer-cert`
  - Reference to: `datapower_cryptocertificate:id`
- `ebms3_duplicate_detection_notification` (Boolean) Duplicate detection notification
  - CLI Alias: `ebms3-duplicate-detection-notification`
  - Default value: `false`
- `ebms3_inbound_decrypt_id_cred` (String) Inbound decryption identification credentials
  - CLI Alias: `ebms3-decrypt-idcred`
  - Reference to: `datapower_cryptoidentcred:id`
- `ebms3_inbound_require_compressed` (Boolean) Require compression
  - CLI Alias: `ebms3-require-compressed`
  - Default value: `false`
- `ebms3_inbound_require_encrypted` (Boolean) Require encryption
  - CLI Alias: `ebms3-require-encrypted`
  - Default value: `false`
- `ebms3_inbound_require_signed` (Boolean) Require signature
  - CLI Alias: `ebms3-require-signed`
  - Default value: `false`
- `ebms3_inbound_verify_val_cred` (String) Inbound signature validation credentials
  - CLI Alias: `ebms3-verify-valcred`
  - Reference to: `datapower_cryptovalcred:id`
- `ebms3_outbound_sign` (Boolean) Sign outbound messages
  - CLI Alias: `ebms3-sign`
  - Default value: `false`
- `ebms3_outbound_sign_digest_alg` (String) Signing digest algorithm
  - CLI Alias: `ebms3-sign-digest-alg`
  - Choices: `sha1`, `sha256`, `sha512`, `ripemd160`, `sha224`, `sha384`, `md5`
  - Default value: `sha1`
- `ebms3_outbound_sign_id_cred` (String) Signing identification credentials
  - CLI Alias: `ebms3-sign-idcred`
  - Reference to: `datapower_cryptoidentcred:id`
- `ebms3_outbound_signature_alg` (String) Signature algorithm
  - CLI Alias: `ebms3-signature-alg`
  - Choices: `dsa-sha1`, `rsa-sha1`, `rsa-sha256`, `rsa-sha384`, `rsa-sha512`, `rsa-ripemd160`, `rsa-ripemd160-2010`, `sha256-rsa-MGF1`, `rsa-md5`, `ecdsa-sha1`, `ecdsa-sha224`, `ecdsa-sha256`, `ecdsa-sha384`, `ecdsa-sha512`
  - Default value: `rsa-sha1`
- `ebms3_outbound_signature_c14n_alg` (String) Signature canonicalization method
  - CLI Alias: `ebms3-signature-c14n-alg`
  - Choices: `c14n`, `exc-c14n`, `c14n-comments`, `exc-c14n-comments`, `c14n11`, `c14n11-comments`
- `ebms3_receipt_ssl_client` (String) Receipt/Error TLS client profile
  - CLI Alias: `ebms3-receipt-ssl-client`
  - Reference to: `datapower_sslclientprofile:id`
- `ebms3_receipt_ssl_client_config_type` (String) Receipt/Error TLS client type
  - CLI Alias: `ebms3-receipt-ssl-client-type`
  - Choices: `client`
  - Default value: `client`
- `ebms_ack_ssl_client` (String) Acknowledgment/Error TLS client profile
  - CLI Alias: `ebmsack-ssl-client`
  - Reference to: `datapower_sslclientprofile:id`
- `ebms_ack_ssl_client_config_type` (String) Acknowledgment/Error TLS client type
  - CLI Alias: `ebmsack-ssl-client-type`
  - Choices: `client`
  - Default value: `client`
- `ebms_ack_url` (String) Acknowledgment URL
  - CLI Alias: `ebms-ack-url`
- `ebms_action` (String) Default action
  - CLI Alias: `ebms-action`
- `ebms_allow_duplicate_message` (String) Allow duplicate ebMS2 inbound message
  - CLI Alias: `ebms-allow-dup-msg`
  - Choices: `never`, `always`, `on-error`
  - Default value: `never`
- `ebms_cpa_id` (String) Default CPA ID
  - CLI Alias: `ebms-cpa-id`
- `ebms_default_signer_cert` (String) Default inbound signature validation certificate
  - CLI Alias: `ebms-default-signer-cert`
  - Reference to: `datapower_cryptocertificate:id`
- `ebms_enable_cpa_binding` (Boolean) Enable CPA bindings
  - CLI Alias: `ebms-enable-cpa-bindings`
  - Default value: `false`
- `ebms_error_url` (String) Error URL
  - CLI Alias: `ebms-error-url`
- `ebms_inbound_decrypt_id_cred` (String) Inbound decryption identification credentials
  - CLI Alias: `ebms-decrypt-idcred`
  - Reference to: `datapower_cryptoidentcred:id`
- `ebms_inbound_error_url` (String) Error URL
  - CLI Alias: `ebms-inbound-error-url`
- `ebms_inbound_receipt_reply_pattern` (String) Receipt reply pattern
  - CLI Alias: `ebms-inbound-receipt-reply-pattern`
  - Choices: `Response`, `Callback`
  - Default value: `Response`
- `ebms_inbound_require_encrypted` (Boolean) Require encryption
  - CLI Alias: `ebms-require-encrypted`
  - Default value: `false`
- `ebms_inbound_require_signed` (Boolean) Require signature
  - CLI Alias: `ebms-require-signed`
  - Default value: `false`
- `ebms_inbound_send_receipt` (Boolean) Reply with receipt
  - CLI Alias: `ebms-inbound-send-receipt`
  - Default value: `false`
- `ebms_inbound_send_signed_receipt` (Boolean) Reply with signed receipt
  - CLI Alias: `ebms-inbound-send-signed-receipt`
  - Default value: `false`
- `ebms_inbound_verify_val_cred` (String) Inbound signature validation credentials
  - CLI Alias: `ebms-verify-valcred`
  - Reference to: `datapower_cryptovalcred:id`
- `ebms_message_properties` (Attributes List) ebMS3 message properties
  - CLI Alias: `ebms-messageproperties` (see [below for nested schema](#nestedatt--ebms_message_properties))
- `ebms_notification` (Boolean) Enable notification
  - CLI Alias: `ebms-notification`
  - Default value: `false`
- `ebms_notification_ssl_client` (String) Notification TLS client profile
  - CLI Alias: `ebms-notification-ssl-client`
  - Reference to: `datapower_sslclientprofile:id`
- `ebms_notification_ssl_client_config_type` (String) Notification TLS client type
  - CLI Alias: `ebms-notification-ssl-client-type`
  - Choices: `client`
  - Default value: `client`
- `ebms_notification_url` (String) Notification URL
  - CLI Alias: `ebms-notification-url`
- `ebms_outbound_sign` (Boolean) Sign outbound messages
  - CLI Alias: `ebms-sign`
  - Default value: `false`
- `ebms_outbound_sign_digest_alg` (String) Signing digest algorithm
  - CLI Alias: `ebms-sign-digest-alg`
  - Choices: `sha1`, `sha256`, `sha512`, `ripemd160`, `sha224`, `sha384`, `md5`
  - Default value: `sha1`
- `ebms_outbound_sign_id_cred` (String) Signing identification credentials
  - CLI Alias: `ebms-sign-idcred`
  - Reference to: `datapower_cryptoidentcred:id`
- `ebms_outbound_signature_alg` (String) Signature algorithm
  - CLI Alias: `ebms-signature-alg`
  - Choices: `dsa-sha1`, `rsa-sha1`, `rsa-sha256`, `rsa-sha384`, `rsa-sha512`, `rsa-ripemd160`, `rsa-ripemd160-2010`, `sha256-rsa-MGF1`, `rsa-md5`, `ecdsa-sha1`, `ecdsa-sha224`, `ecdsa-sha256`, `ecdsa-sha384`, `ecdsa-sha512`
  - Default value: `dsa-sha1`
- `ebms_outbound_signature_c14n_alg` (String) Signature canonicalization method
  - CLI Alias: `ebms-signature-c14n-alg`
  - Choices: `c14n`, `exc-c14n`, `c14n-comments`, `exc-c14n-comments`, `c14n11`, `c14n11-comments`
  - Default value: `c14n`
- `ebms_persist_duration` (Number) Persist duration
  - CLI Alias: `ebms-persist-duration`
  - Range: `0`-`6000000`
- `ebms_profile_cpa_bindings` (Attributes List) CPA bindings
  - CLI Alias: `profile-cpa-binding` (see [below for nested schema](#nestedatt--ebms_profile_cpa_bindings))
- `ebms_receipt_url` (String) Receipt URL
  - CLI Alias: `ebms-receipt-url`
- `ebms_role` (String) Role
  - CLI Alias: `ebms-role`
- `ebms_service` (String) Default service
  - CLI Alias: `ebms-service`
- `ebms_start_parameter` (Boolean) Generate start parameter
  - CLI Alias: `ebms-start-parameter`
  - Default value: `false`
- `email_addresses` (List of String) Partner email addresses
  - CLI Alias: `email-address`
- `inbound_decrypt_id_cred` (String) Inbound decryption identification credentials
  - CLI Alias: `decrypt-idcred`
  - Reference to: `datapower_cryptoidentcred:id`
- `inbound_require_encrypted` (Boolean) Require encryption
  - CLI Alias: `require-encrypted`
  - Default value: `false`
- `inbound_require_signed` (Boolean) Require signature
  - CLI Alias: `require-signed`
  - Default value: `false`
- `inbound_verify_val_cred` (String) Inbound signature validation credentials
  - CLI Alias: `verify-valcred`
  - Reference to: `datapower_cryptovalcred:id`
- `mdnssl_client` (String) MDN TLS client profile
  - CLI Alias: `mdn-ssl-client`
  - Reference to: `datapower_sslclientprofile:id`
- `mdnssl_client_config_type` (String) MDN TLS client type
  - CLI Alias: `mdn-ssl-client-type`
  - Choices: `client`
  - Default value: `client`
- `outbound_sign` (Boolean) Sign outbound messages
  - CLI Alias: `sign`
  - Default value: `false`
- `outbound_sign_digest_alg` (String) Signing digest algorithm
  - CLI Alias: `sign-digest-alg`
  - Choices: `sha1`, `md5`, `sha256`, `sha384`, `sha512`
  - Default value: `sha1`
- `outbound_sign_id_cred` (String) Signing identification credentials
  - CLI Alias: `sign-idcred`
  - Reference to: `datapower_cryptoidentcred:id`
- `outbound_sign_mic_alg_version` (String) Signing S/MIME version
  - CLI Alias: `sign-micalg-version`
  - Choices: `SMIME3.1`, `SMIME3.2`
- `override_asid` (String) Override AS identifier
  - CLI Alias: `override-as-identifier`
- `preserve_filename` (Boolean) Preserve file name
  - CLI Alias: `preserve-filename`
  - Default value: `false`
- `profile_type` (String) Profile type
  - CLI Alias: `profile-type`
  - Choices: `internal`, `external`
  - Default value: `internal`
- `response_type` (String) Response traffic type
  - CLI Alias: `response-type`
  - Choices: `soap`, `xml`, `unprocessed`, `preprocessed`
  - Default value: `preprocessed`
- `user_summary` (String) Comments
  - CLI Alias: `summary`

<a id="nestedatt--destinations"></a>
### Nested Schema for `destinations`

Required:

- `dest_name` (String) Destination name
  - CLI Alias: `name`
- `dest_url` (String) Destination URL
  - CLI Alias: `dest-url`

Optional:

- `ack_time` (Number) Time to acknowledge
  - CLI Alias: `ack-time`
  - Range: `1`-`3600`
  - Default value: `1800`
- `as1mdn_redirect_email` (String) AS1 MDN redirection E-mail
  - CLI Alias: `as1-mdn-email`
- `as2mdn_redirect_url` (String) AS2 MDN redirection URL
  - CLI Alias: `as2-mdn-url`
- `as3mdn_redirect_url` (String) AS3 MDN redirection URL
  - CLI Alias: `as3-mdn-url`
- `as_compress` (Boolean) Compress messages
  - CLI Alias: `as-compress`
  - Default value: `false`
- `as_compress_before_sign` (Boolean) Compress before sign
  - CLI Alias: `as-compress-before-sign`
  - Default value: `false`
- `as_encrypt` (Boolean) Encrypt messages
  - CLI Alias: `as-encrypt`
  - Default value: `false`
- `as_encrypt_alg` (String) Encryption algorithm
  - CLI Alias: `as-encrypt-alg`
  - Choices: `3des`, `des`, `rc2-128`, `rc2-64`, `rc2-40`, `aes-128`, `aes-192`, `aes-256`
  - Default value: `3des`
- `as_encrypt_cert` (String) Encryption certificate
  - CLI Alias: `as-encrypt-cert`
  - Reference to: `datapower_cryptocertificate:id`
- `as_send_unsigned` (Boolean) Send messages unsigned
  - CLI Alias: `as-send-unsigned`
  - Default value: `false`
- `asmdn_request` (Boolean) Request MDN
  - CLI Alias: `as-mdn-request`
  - Default value: `false`
- `asmdn_request_async` (Boolean) Request asynchronous MDN
  - CLI Alias: `as-mdn-request-async`
  - Default value: `false`
- `asmdn_request_signed` (Boolean) Request signed MDN
  - CLI Alias: `as-mdn-request-signed`
  - Default value: `false`
- `asmdn_request_signed_algs` (String) Request MDN signing algorithms
  - CLI Alias: `as-mdn-request-signed-algs`
  - Default value: `sha1,md5`
- `auth_tls` (String) Encrypt command connection
  - CLI Alias: `ftp-auth-tls`
  - Choices: `auth-off`, `auth-tls-opt`, `auth-tls-req`, `auth-tls-imp`
  - Default value: `auth-off`
- `binary_transfer_mode` (String) Binary transfer
  - CLI Alias: `binary-transfer-mode`
  - Choices: `auto-detect`, `enforce`
  - Default value: `auto-detect`
- `data_type` (String) Data type
  - CLI Alias: `ftp-data-type`
  - Choices: `ascii`, `binary`
  - Default value: `binary`
- `ebms_ack_request` (Boolean) Request acknowledgment
  - CLI Alias: `ebms-ack-request`
  - Default value: `false`
- `ebms_ack_request_signed` (Boolean) Request signed acknowledgment
  - CLI Alias: `ebms-ack-request-signed`
  - Default value: `false`
- `ebms_action` (String) Action
  - CLI Alias: `ebms-action`
- `ebms_agreement_ref` (String) PMode AgreementRef
  - CLI Alias: `ebms-agreementref`
- `ebms_compress` (Boolean) Compress messages
  - CLI Alias: `ebms-compress`
  - Default value: `false`
- `ebms_cpa_id` (String) CPA ID
  - CLI Alias: `ebms-cpa-id`
- `ebms_duplicate_elimination_request` (Boolean) Request duplicate elimination
  - CLI Alias: `ebms-duplicate-elimination-request`
  - Default value: `true`
- `ebms_encrypt` (Boolean) Encrypt messages
  - CLI Alias: `ebms-encrypt`
  - Default value: `false`
- `ebms_encrypt_alg` (String) Encryption algorithm
  - CLI Alias: `ebms-encrypt-alg`
  - Choices: `http://www.w3.org/2001/04/xmlenc#tripledes-cbc`, `http://www.w3.org/2001/04/xmlenc#aes128-cbc`, `http://www.w3.org/2001/04/xmlenc#aes192-cbc`, `http://www.w3.org/2001/04/xmlenc#aes256-cbc`, `http://www.w3.org/2009/xmlenc11#aes128-gcm`, `http://www.w3.org/2009/xmlenc11#aes192-gcm`, `http://www.w3.org/2009/xmlenc11#aes256-gcm`
  - Default value: `http://www.w3.org/2001/04/xmlenc#tripledes-cbc`
- `ebms_encrypt_cert` (String) Encryption certificate
  - CLI Alias: `ebms-encrypt-cert`
  - Reference to: `datapower_cryptocertificate:id`
- `ebms_include_time_to_live` (Boolean) Include TimeToLive element
  - CLI Alias: `ebms-include-time-to-live`
  - Default value: `true`
- `ebms_key_encrypt_alg` (String) Asymmetric key encryption algorithm
  - CLI Alias: `ebms-key-encrypt-alg`
  - Choices: `http://www.w3.org/2001/04/xmlenc#rsa-1_5`, `http://www.w3.org/2001/04/xmlenc#rsa-oaep-mgf1p`, `http://www.w3.org/2009/xmlenc11#rsa-oaep`
  - Default value: `http://www.w3.org/2001/04/xmlenc#rsa-1_5`
- `ebms_max_retries` (Number) Max retransmit attempts
  - CLI Alias: `ebms-max-retries`
  - Range: `1`-`30`
  - Default value: `3`
- `ebms_message_exchange_pattern` (String) Message exchange pattern
  - CLI Alias: `ebms-mep`
  - Choices: `one-way-push`, `one-way-pull`
  - Default value: `one-way-push`
- `ebms_message_partition_channel` (String) Message partition channel
  - CLI Alias: `ebms-mpc`
- `ebms_outbound_receipt_reply_pattern` (String) Requested receipt reply pattern
  - CLI Alias: `ebms-outbound-receipt-reply-pattern`
  - Choices: `Response`, `Callback`
  - Default value: `Response`
- `ebms_outbound_reception_awareness_notification` (Boolean) Reception awareness error notification
  - CLI Alias: `ebms-reception-awareness-notification`
  - Default value: `false`
- `ebms_outbound_reception_awareness_timeout` (Number) Reception awareness timeout
  - CLI Alias: `ebms-reception-awareness-timeout`
  - Range: `3`-`7200`
  - Default value: `300`
- `ebms_outbound_request_receipt` (Boolean) Request receipt
  - CLI Alias: `ebms-outbound-request-receipt`
  - Default value: `false`
- `ebms_outbound_request_signed_receipt` (Boolean) Request signed receipt
  - CLI Alias: `ebms-outbound-request-signed-receipt`
  - Default value: `false`
- `ebms_retry` (Boolean) Attempt message retransmit
  - CLI Alias: `ebms-retry`
  - Default value: `false`
- `ebms_retry_interval` (Number) Retry interval
  - CLI Alias: `ebms-retry-interval`
  - Range: `1`-`3600`
  - Default value: `1800`
- `ebms_send_unsigned` (Boolean) Send messages unsigned
  - CLI Alias: `ebms-send-unsigned`
  - Default value: `false`
- `ebms_service` (String) Service
  - CLI Alias: `ebms-service`
- `ebms_service_type` (String) Service type
  - CLI Alias: `ebms-service-type`
- `ebms_sync_reply_mode` (String) SyncReply mode
  - CLI Alias: `ebms-syncreply-mode`
  - Choices: `mshSignalsOnly`, `none`
  - Default value: `none`
- `ebmsmpc_auth_method` (String) MPC authentication method
  - CLI Alias: `embs-mpc-auth-method`
  - Choices: `username-token`, `cert`
  - Default value: `username-token`
- `ebmsmpc_verify_val_cred` (String) MPC validation credentials
  - CLI Alias: `ebms-mpc-verify-valcred`
  - Reference to: `datapower_cryptovalcred:id`
- `ebmsp_mode` (String) PMode ID
  - CLI Alias: `ebms-pmode`
- `ebmssoap_body` (Boolean) Messages in SOAP Body
  - CLI Alias: `ebms-soapbody`
  - Default value: `false`
- `email_address` (String) Email address
  - CLI Alias: `email-address`
- `enable_ftp_settings` (Boolean) Enable advanced AS3/FTP settings
  - CLI Alias: `enable-ftp-settings`
  - Default value: `false`
- `enabled_doc_type` (Attributes) Enabled document type
  - CLI Alias: `enabled-doc-type` (see [below for nested schema](#nestedatt--destinations--enabled_doc_type))
- `encrypt_data` (String) Encrypt file transfers
  - CLI Alias: `ftp-encrypt-data`
  - Choices: `enc-data-off`, `enc-data-opt`, `enc-data-req`
  - Default value: `enc-data-off`
- `max_resends` (Number) Retransmit attempts
  - CLI Alias: `max-resends`
  - Range: `1`-`30`
  - Default value: `3`
- `override_timeout` (Number) Connection timeout
  - CLI Alias: `timeout`
  - Range: `3`-`7200`
  - Default value: `300`
- `passive` (String) Passive mode
  - CLI Alias: `ftp-passive`
  - Choices: `pasv-off`, `pasv-opt`, `pasv-req`
  - Default value: `pasv-req`
- `password_alias` (String) Password alias
  - CLI Alias: `password-alias`
  - Reference to: `datapower_passwordalias:id`
- `quoted_commands` (String) Quoted commands
  - CLI Alias: `ftp-quoted-commands`
  - Reference to: `datapower_ftpquotecommands:id`
- `retransmit` (Boolean) Attempt message retransmit
  - CLI Alias: `retransmit`
  - Default value: `false`
- `size_check` (String) Size check
  - CLI Alias: `ftp-size-check`
  - Choices: `size-check-optional`, `size-check-disabled`
  - Default value: `size-check-optional`
- `slash_stou` (String) Write unique filename if trailing slash
  - CLI Alias: `ftp-slash-stou`
  - Choices: `slash-stou-off`, `slash-stou-on`
  - Default value: `slash-stou-on`
- `smtp_server_connection` (String) SMTP server connection
  - CLI Alias: `smtp-server-connection`
  - Reference to: `datapower_smtpserverconnection:id`
  - Default value: `default`
- `ssh_client_connection` (String) SSH client connection
  - CLI Alias: `ssh-client-connection`
  - Reference to: `datapower_sshclientprofile:id`
- `ssl_client` (String) TLS client profile
  - CLI Alias: `ssl-client`
  - Reference to: `datapower_sslclientprofile:id`
- `ssl_client_config_type` (String) TLS client type
  - CLI Alias: `ssl-client-type`
  - Choices: `client`
  - Default value: `client`
- `use_ccc` (String) Stop command encryption after authentication
  - CLI Alias: `ftp-use-ccc`
  - Choices: `ccc-off`, `ccc-opt`, `ccc-req`
  - Default value: `ccc-off`
- `use_unique_filenames` (Boolean) Use unique file names
  - CLI Alias: `use-unique-filenames`
  - Default value: `false`
- `user_name` (String) Username
  - CLI Alias: `username`
- `user_name_token` (String) Username token
  - CLI Alias: `username-token`
- `user_name_token_password_alias` (String) Username token password alias
  - CLI Alias: `username-token-password-alias`
  - Reference to: `datapower_passwordalias:id`

<a id="nestedatt--destinations--enabled_doc_type"></a>
### Nested Schema for `destinations.enabled_doc_type`

Optional:

- `enable_binary` (Boolean) Binary
  - Default value: `true`
- `enable_edifact` (Boolean) EDIFACT
  - Default value: `true`
- `enable_x12` (Boolean) X12
  - Default value: `true`
- `enable_xml` (Boolean) XML
  - Default value: `true`



<a id="nestedatt--contacts"></a>
### Nested Schema for `contacts`

Optional:

- `email` (String) Email
  - CLI Alias: `email`
- `family_name` (String) Family name
  - CLI Alias: `family-name`
- `given_name` (String) Given name
  - CLI Alias: `given-name`
- `phone` (String) Phone
  - CLI Alias: `phone`
- `title` (String) Title
  - CLI Alias: `title`


<a id="nestedatt--ebms_message_properties"></a>
### Nested Schema for `ebms_message_properties`

Required:

- `name` (String) Name
- `value` (String) Value

Optional:

- `type` (String) Type


<a id="nestedatt--ebms_profile_cpa_bindings"></a>
### Nested Schema for `ebms_profile_cpa_bindings`

Required:

- `collaboration` (String) Service
  - CLI Alias: `collaboration`
  - Reference to: `datapower_b2bcpacollaboration:id`
- `cpa` (String) CPA
  - CLI Alias: `cpa`
  - Reference to: `datapower_b2bcpa:id`
- `internal_partner` (String) CPA sender (internal partner profile)
  - CLI Alias: `internal-partner`
  - Reference to: `datapower_b2bprofile:id`

Optional:

- `action` (String) Action
  - CLI Alias: `action`

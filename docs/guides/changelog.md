---
subcategory: "Guides"
page_title: "Changelog"
description: |-
    Changelog
---

# Changelog

## Unreleased

- Fix model test config generator to properly handle list of reference objects
- API Additions (Objects)
  - AMQPBroker
  - AMQPSourceProtocolHandler
  - AnalyticsEndpoint
  - APIApplicationType
  - APIAuthURLRegistry
  - APIClientIdentification
  - APICollection
  - APICORS
  - APIDefinition
  - APIExecute
  - APIFinal
  - APIGateway
  - APILDAPRegistry
  - APIOperation
  - APIPath
  - APIPlan
  - APIRateLimit
  - APIResult
  - APIRouting
  - APIRule
  - APISchema
  - APISecurity
  - APISecurityAPIKey
  - APISecurityBasicAuth
  - APISecurityHTTPScheme
  - APISecurityOAuth
  - APISecurityOAuthReq
  - APISecurityRequirement
  - AppSecurityPolicy
  - NameValueProfile
  - WebAppErrorHandlingPolicy
  - WebAppFW
  - WebAppRequest
  - WebAppResponse
  - WebAppSessionPolicy

## 0.2.0

- Update getting started guide
- Enhanced `datasource` and `resource` testing
- Enhanced string/bool conversion functions to handle admin state in addition to toggle
- Enhanced code generation schema to lock API endpoints to `default` domain only
- Refactored `read-only` attribute handling
- API Additions (Domain Settings):
  - APIConnectGatewayService
  - APISecurityTokenManager
  - AuditLog
  - B2BPersistence
  - CertMonitor
  - CRLFetch
  - DistributedVariable
  - DNSNameService
  - DomainAvailability
  - DomainSettings
  - ErrorReportSettings
  - FileSystemUsageMonitor
  - GatewayPeeringManager
  - GitOps
  - GitOpsVariables
  - GWScriptSettings
  - GWSRemoteDebug
  - InteropService
  - MgmtInterface
  - NFSClientSettings
  - NFSDynamicMounts
  - ODR
  - Probe
  - QuotaEnforcementServer
  - RADIUSSettings
  - RaidVolume
  - RateLimitConfiguration
  - RBMSettings
  - SNMPSettings
  - SSHDomainClientProfile
  - SSHServerProfile
  - SSHService
  - Statistics
  - SystemSettings
- API Additions (Objects)
  - GatewayPeering

## 0.1.2

- Fix resource client error handling

## 0.1.1

- Fix file path handling on Windows in `datapower_file` resource
- BREAKING CHANGE: Provider configuration values for `username` and `password` now take precedence over environment variables
- Fix `DP_INSECURE` environment parameter to only disable certificate validation if set to true
- Fix provider port value validated for proper range
- Fix client GET path normalized to remove trailing `/` in client
- Fix JSON error handling in GET response
- Enhanced client by adding retry max of 3 and timeout of 30s
- Enhanced client test suite
- Merged dependabot PR1, PR2, PR3
- Fix resource client error handling

## 0.1.0

- Initial release


---
subcategory: "Guides"
page_title: "Changelog"
description: |-
    Changelog
---

# Changelog

## UNRELEASED

- Bump github.com/hashicorp/terraform-plugin-framework from 1.16.0 to 1.16.1 in the terraform-plugin group (PR16)
- Added Terraform extensions to devcontainer
- Added `import` capability to all resources and changed resource creation to update existing resource if it already exists (with warning)
- Added warning when resource deletion fails due to conflict
- Fixed error when updating an existing `datapower_user` resource with the existing password
- Added reset to default values when deleting singleton resource
- Updates to test bed configuration
- Refactored `UpdateFromBody` method to better handle lists of objects

## v0.11.4

- Updated string validation to treat empty strings as null when value not expected.

## v0.11.3

- Bump the terraform-plugin group with 3 updates (PR15)
- Reverted `variable` `default` `null` value handling in modules
- Updates to test bed configuration

## v0.11.2

- Fixed `variable` `default` `null` value handling in modules (workaround for [known](https://github.com/hashicorp/terraform/issues/21702) Terraform oddity)

## v0.11.1

- Added retrieve object by name option to datasources
- Bump actions/setup-go from 5.5.0 to 6.0.0 in the github-actions group (PR13)
- Bump golang.org/x/tools from 0.36.0 to 0.37.0 (PR14)

## v0.11.0

- Fixed case for `required_when` `values` in `ConformancePolicy`, `dmAAAPExtractIdentity`, `dmAAAPExtractResource` and `dmSMTPPolicy` definitions.
- Corrected `do_s_valve` attribute to `dos_valve` in `AAAPolicy`
- Added default `document_cache_size` in `XMLManager`
- Added `ignored_when` var definitions to models
- Added `ignored_when` logic and attribute validators for conditionally ignoring attributes
- Added devcontainer for development
- Bump github.com/stretchr/testify from 1.11.0 to 1.11.1 (PR11)
- Bump golang.org/x/crypto from 0.41.0 to 0.42.0 (PR12)

## v0.10.0

- Added validation criteria to definitions
- BREAKING CHANGE: Refactored `write-only` handling to align with common provider practice
- BREAKING CHANGE: `sensitive` and `write-only` properties in `DomainSettings:Passphrase`, `dmRadiusServer:Secret`, `dmSnmpCred:AuthSecret`, `dmSnmpCred:PrivSecret`, `LogTarget:RemotePassword` definitions
- Bump github.com/hashicorp/terraform-plugin-testing from 1.13.2 to 1.13.3 in the terraform-plugin group (PR8)
- Added `required_when` logic and attribute validators for conditionally requiring attributes
- BREAKING CHANGE: Changed attribute names to aid in readability.
- Bump goreleaser/goreleaser-action from 6.3.0 to 6.4.0 in the github-actions group (PR9)
- Bump github.com/stretchr/testify from 1.10.0 to 1.11.0 (PR10)

## v0.9.1

- Fixed `dependency_actions` handling of non-existent resources (silently proceed)

## v0.9.0

- BREAKING CHANGE: Changed `data_source` and `resource` names to aid in readability.

## v0.7.0

- Cleaned up formatting in resource template
- Added string validation for remote file in file resource
- Refactored resource and datasource directory structure
- Refactored internal/provider/models/testconfig to testutils
- Added `dependency_actions` to all resources
- Moved domain `SaveConfig` operations to provider exit
- Test suite refactors/enhancements
- Updated github.com/hashicorp/terraform-plugin-framework from 1.15.0 to 1.15.1 (PR4)
- Bump golang.org/x/tools from 0.35.0 to 0.36.0 (PR5)
- Bump golang.org/x/crypto from 0.40.0 to 0.41.0 (PR6)
- Fixed `Sensitive` attribute property
- BREAKING CHANGE: Removed `restart_domain_on_update` attribute from `datapower_domainavailability`. Should now use `dependency_actions`.
- Enhanced attribute descriptions
- Bump actions/checkout from 4.2.2 to 5.0.0 in the github-actions group (PR7)

## v0.6.2

- Added internal use only attribute flag
- Added domain restart on change model flag
- Added domain restart flag to DomainAvailability
- Added string validations to Id and AppDomain attributes for all models

## v0.6.1

- Update test fixtures for efficiency

## v0.6.0

- Fix test generator handling of test_value for non-string types
- API Additions (Domain Settings):
  - SecureBackupMode
- API Additions (Objects)
  - CompileSettings
  - ConfigSequence
  - ConformancePolicy
  - ControlList
  - CryptoKerberosKDC
  - DocumentCryptoMap
  - FTPFilePollerSourceProtocolHandler
  - FTPServerSourceProtocolHandler
  - GatewayPeeringGroup
  - GitOpsTemplate
  - GraphQLSchemaOptions
  - HostAlias
  - ImportPackage
  - IncludeConfig
  - KafkaCluster
  - KafkaSourceProtocolHandler
  - LogTarget
  - MTOMPolicy
  - NFSFilePollerSourceProtocolHandler
  - NFSStaticMount
  - OAuthProviderSettings
  - ODRConnectorGroup
  - OpenTelemetry
  - OpenTelemetryExporter
  - OpenTelemetrySampler
  - OperationRateLimit
  - PeerGroup
  - POPPollerSourceProtocolHandler
  - RateLimitDefinitionGroup
  - SchemaExceptionMap
  - SFTPFilePollerSourceProtocolHandler
  - SMTPServerConnection
  - SOAPHeaderDisposition
  - SSHServerSourceProtocolHandler
  - SSLProxyService
  - StatelessTCPSourceProtocolHandler
  - TCPProxyService
  - VisibilityList
  - WebSphereJMSServer
  - WebSphereJMSSourceProtocolHandler
  - WebTokenService
  - WSEndpointRewritePolicy
  - WSGateway
  - WSStylePolicy
  - WSStylePolicyRule
  - WXSGrid
  - XPathRoutingMap
  - XSLCoprocService
  - XSLProxyService
  - XTCProtocolHandler

## v0.5.0

- Remove deprecated API's
  - APIRateLimit
  - TFIMEndpoint
- API Additions (Domain Settings):
  - LunaHASettings
- API Additions (Objects)
  - AS1PollerSourceProtocolHandler
  - AS2ProxySourceProtocolHandler
  - AS2SourceProtocolHandler
  - AS3SourceProtocolHandler
  - B2BCPA
  - B2BCPACollaboration
  - B2BCPAReceiverSetting
  - B2BCPASenderSetting
  - B2BGateway
  - B2BProfile
  - B2BProfileGroupmake
  - B2BXPathRoutingPolicy
  - CORSPolicy
  - CORSRule
  - EBMS2SourceProtocolHandler
  - EBMS3SourceProtocolHandler
  - Luna
  - LunaHAGroup
  - LunaPartition
  - MCFCustomRule
  - MCFHttpHeader
  - MCFHttpMethod
  - MCFHttpURL
  - MCFXPath
  - MQManager
  - MQManagerGroup
  - MQv9PlusMFTSourceProtocolHandler
  - MQv9PlusSourceProtocolHandler
  - SLMAction
  - SLMCredClass
  - SLMPolicy
  - SLMRsrcClass
  - SLMSchedule

## v0.4.0

- Removed deprecated SSL proxy settings
- Removed unneccesary `sensitive` markings
- Updated example names
- API Additions (Objects)
  - Assembly
  - AssemblyActionClientSecurity
  - AssemblyActionExtract
  - AssemblyActionFunctionCall
  - AssemblyActionGatewayScript
  - AssemblyActionGraphQLCostAnalysis
  - AssemblyActionGraphQLExecute
  - AssemblyActionGraphQLIntrospect
  - AssemblyActionHtmlPage
  - AssemblyActionInvoke
  - AssemblyActionJson2Xml
  - AssemblyActionJWTGenerate
  - AssemblyActionJWTValidate
  - AssemblyActionLog
  - AssemblyActionMap
  - AssemblyActionOAuth
  - AssemblyActionParse
  - AssemblyActionRateLimit
  - AssemblyActionRateLimitInfo
  - AssemblyActionRedact
  - AssemblyActionSetVar
  - AssemblyActionThrow
  - AssemblyActionUserSecurity
  - AssemblyActionValidate
  - AssemblyActionWebSocketUpgrade
  - AssemblyActionWSDL
  - AssemblyActionXml2Json
  - AssemblyActionXSLT
  - AssemblyFunction
  - AssemblyLogicOperationSwitch
  - AssemblyLogicSwitch

## 0.3.0

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
  - XMLFirewallService

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


## Unreleased
- Cleaned up formatting in resource template
- Added string validation for remote file in file resource
- Refactored resource and datasource directory structure
- Updated github.com/hashicorp/terraform-plugin-framework from 1.15.0 to 1.15.1 (PR4)
- Added `object_actions`

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

# ClusterConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GpuDriverVersion** | Pointer to **string** | GPU driver version. | [optional] 
**ClientAuth** | Pointer to [**ClientAuth**](ClientAuth.md) |  | [optional] 
**AuthorizedPublicKeyList** | Pointer to [**[]PublicKey**](PublicKey.md) | List of valid ssh keys for the cluster. | [optional] 
**SoftwareMap** | Pointer to [**map[string]ClusterSoftware**](ClusterSoftware.md) | Map of software on the cluster with software type as the key.  | [optional] 
**EncryptionStatus** | Pointer to **string** | Cluster encryption status. | [optional] [default to "NOT_SUPPORTED"]
**SslKey** | Pointer to [**SslKey**](SslKey.md) |  | [optional] 
**ServiceList** | Pointer to **[]string** | Array of enabled cluster services. For example, a cluster can function as both AOS and cloud data gateway. - &#39;AOS&#39;: Regular Prism Element - &#39;PRISM_CENTRAL&#39;: Prism Central - &#39;CLOUD_DATA_GATEWAY&#39;: Cloud backup and DR gateway - &#39;AFS&#39;: Cluster for file server - &#39;WITNESS&#39; : Witness cluster - &#39;XI_PORTAL&#39;: Xi cluster - &#39;ONE_NODE_CLUSTER&#39;: Single node backup cluster - &#39;TWO_NODE_CLUSTER&#39;: Two node cluster  | [optional] 
**SupportedInformationVerbosity** | Pointer to **string** | Verbosity level settings for populating support information. - &#39;Nothing&#39;: Send nothing - &#39;Basic&#39;: Send basic information - skip core dump and hypervisor            stats information - &#39;BasicPlusCoreDump&#39;: Send basic and core dump information - &#39;All&#39;: Send all information  | [optional] [default to "BASIC_PLUS_CORE_DUMP"]
**CertificationSigningInfo** | Pointer to [**CertificationSigningInfo**](CertificationSigningInfo.md) |  | [optional] 
**RedundancyFactor** | Pointer to **int32** | Cluster supported redundancy factor. | [optional] 
**ExternalConfigurations** | Pointer to [**ExternalConfigurations**](ExternalConfigurations.md) |  | [optional] 
**OperationMode** | Pointer to **string** | Cluster operation mode. - &#39;NORMAL&#39;: Cluster is operating normally. - &#39;READ_ONLY&#39;: Cluster is operating in read only mode. - &#39;STAND_ALONE&#39;: Only one node is operational in the cluster. This is                  valid only for single node or two node clusters. - &#39;SWITCH_TO_TWO_NODE&#39;: Cluster is moving from single node to two node                         cluster. - &#39;OVERRIDE&#39;: Valid only for single node cluster. If the user wants to               run vms on a single node cluster in read only mode, he               can set the cluster peration mode to override. Writes               will be allowed in override mode.  | [optional] 
**CaCertificateList** | Pointer to [**[]CaCert**](CaCert.md) | List of cluster trusted CA certificates. | [optional] [readonly] 
**DomainAwarenessLevel** | Pointer to **string** | Domain awareness supported on cluster. | [optional] [default to "NODE"]
**EnabledFeatureList** | Pointer to **[]string** | Array of enabled features. | [optional] 
**IsAvailable** | Pointer to **bool** | Indicates if cluster is available to contact. | [optional] [readonly] 
**Build** | Pointer to [**BuildInfo**](BuildInfo.md) |  | [optional] 
**Timezone** | Pointer to **string** | Zone name used in value of TZ environment variable. | [optional] 
**EnableEfficientMetricSync** | Pointer to **bool** | Indicates if downsampling of metrics syncing between PE and PC is enabled or not.  | [optional] 
**ClusterArch** | Pointer to **string** | Cluster architecture. | [optional] [readonly] 
**ManagementServerList** | Pointer to [**[]ClusterManagementServer**](ClusterManagementServer.md) | List of cluster management servers. | [optional] [readonly] 

## Methods

### NewClusterConfig

`func NewClusterConfig() *ClusterConfig`

NewClusterConfig instantiates a new ClusterConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterConfigWithDefaults

`func NewClusterConfigWithDefaults() *ClusterConfig`

NewClusterConfigWithDefaults instantiates a new ClusterConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGpuDriverVersion

`func (o *ClusterConfig) GetGpuDriverVersion() string`

GetGpuDriverVersion returns the GpuDriverVersion field if non-nil, zero value otherwise.

### GetGpuDriverVersionOk

`func (o *ClusterConfig) GetGpuDriverVersionOk() (*string, bool)`

GetGpuDriverVersionOk returns a tuple with the GpuDriverVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuDriverVersion

`func (o *ClusterConfig) SetGpuDriverVersion(v string)`

SetGpuDriverVersion sets GpuDriverVersion field to given value.

### HasGpuDriverVersion

`func (o *ClusterConfig) HasGpuDriverVersion() bool`

HasGpuDriverVersion returns a boolean if a field has been set.

### GetClientAuth

`func (o *ClusterConfig) GetClientAuth() ClientAuth`

GetClientAuth returns the ClientAuth field if non-nil, zero value otherwise.

### GetClientAuthOk

`func (o *ClusterConfig) GetClientAuthOk() (*ClientAuth, bool)`

GetClientAuthOk returns a tuple with the ClientAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientAuth

`func (o *ClusterConfig) SetClientAuth(v ClientAuth)`

SetClientAuth sets ClientAuth field to given value.

### HasClientAuth

`func (o *ClusterConfig) HasClientAuth() bool`

HasClientAuth returns a boolean if a field has been set.

### GetAuthorizedPublicKeyList

`func (o *ClusterConfig) GetAuthorizedPublicKeyList() []PublicKey`

GetAuthorizedPublicKeyList returns the AuthorizedPublicKeyList field if non-nil, zero value otherwise.

### GetAuthorizedPublicKeyListOk

`func (o *ClusterConfig) GetAuthorizedPublicKeyListOk() (*[]PublicKey, bool)`

GetAuthorizedPublicKeyListOk returns a tuple with the AuthorizedPublicKeyList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizedPublicKeyList

`func (o *ClusterConfig) SetAuthorizedPublicKeyList(v []PublicKey)`

SetAuthorizedPublicKeyList sets AuthorizedPublicKeyList field to given value.

### HasAuthorizedPublicKeyList

`func (o *ClusterConfig) HasAuthorizedPublicKeyList() bool`

HasAuthorizedPublicKeyList returns a boolean if a field has been set.

### GetSoftwareMap

`func (o *ClusterConfig) GetSoftwareMap() map[string]ClusterSoftware`

GetSoftwareMap returns the SoftwareMap field if non-nil, zero value otherwise.

### GetSoftwareMapOk

`func (o *ClusterConfig) GetSoftwareMapOk() (*map[string]ClusterSoftware, bool)`

GetSoftwareMapOk returns a tuple with the SoftwareMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareMap

`func (o *ClusterConfig) SetSoftwareMap(v map[string]ClusterSoftware)`

SetSoftwareMap sets SoftwareMap field to given value.

### HasSoftwareMap

`func (o *ClusterConfig) HasSoftwareMap() bool`

HasSoftwareMap returns a boolean if a field has been set.

### GetEncryptionStatus

`func (o *ClusterConfig) GetEncryptionStatus() string`

GetEncryptionStatus returns the EncryptionStatus field if non-nil, zero value otherwise.

### GetEncryptionStatusOk

`func (o *ClusterConfig) GetEncryptionStatusOk() (*string, bool)`

GetEncryptionStatusOk returns a tuple with the EncryptionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionStatus

`func (o *ClusterConfig) SetEncryptionStatus(v string)`

SetEncryptionStatus sets EncryptionStatus field to given value.

### HasEncryptionStatus

`func (o *ClusterConfig) HasEncryptionStatus() bool`

HasEncryptionStatus returns a boolean if a field has been set.

### GetSslKey

`func (o *ClusterConfig) GetSslKey() SslKey`

GetSslKey returns the SslKey field if non-nil, zero value otherwise.

### GetSslKeyOk

`func (o *ClusterConfig) GetSslKeyOk() (*SslKey, bool)`

GetSslKeyOk returns a tuple with the SslKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslKey

`func (o *ClusterConfig) SetSslKey(v SslKey)`

SetSslKey sets SslKey field to given value.

### HasSslKey

`func (o *ClusterConfig) HasSslKey() bool`

HasSslKey returns a boolean if a field has been set.

### GetServiceList

`func (o *ClusterConfig) GetServiceList() []string`

GetServiceList returns the ServiceList field if non-nil, zero value otherwise.

### GetServiceListOk

`func (o *ClusterConfig) GetServiceListOk() (*[]string, bool)`

GetServiceListOk returns a tuple with the ServiceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceList

`func (o *ClusterConfig) SetServiceList(v []string)`

SetServiceList sets ServiceList field to given value.

### HasServiceList

`func (o *ClusterConfig) HasServiceList() bool`

HasServiceList returns a boolean if a field has been set.

### GetSupportedInformationVerbosity

`func (o *ClusterConfig) GetSupportedInformationVerbosity() string`

GetSupportedInformationVerbosity returns the SupportedInformationVerbosity field if non-nil, zero value otherwise.

### GetSupportedInformationVerbosityOk

`func (o *ClusterConfig) GetSupportedInformationVerbosityOk() (*string, bool)`

GetSupportedInformationVerbosityOk returns a tuple with the SupportedInformationVerbosity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedInformationVerbosity

`func (o *ClusterConfig) SetSupportedInformationVerbosity(v string)`

SetSupportedInformationVerbosity sets SupportedInformationVerbosity field to given value.

### HasSupportedInformationVerbosity

`func (o *ClusterConfig) HasSupportedInformationVerbosity() bool`

HasSupportedInformationVerbosity returns a boolean if a field has been set.

### GetCertificationSigningInfo

`func (o *ClusterConfig) GetCertificationSigningInfo() CertificationSigningInfo`

GetCertificationSigningInfo returns the CertificationSigningInfo field if non-nil, zero value otherwise.

### GetCertificationSigningInfoOk

`func (o *ClusterConfig) GetCertificationSigningInfoOk() (*CertificationSigningInfo, bool)`

GetCertificationSigningInfoOk returns a tuple with the CertificationSigningInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificationSigningInfo

`func (o *ClusterConfig) SetCertificationSigningInfo(v CertificationSigningInfo)`

SetCertificationSigningInfo sets CertificationSigningInfo field to given value.

### HasCertificationSigningInfo

`func (o *ClusterConfig) HasCertificationSigningInfo() bool`

HasCertificationSigningInfo returns a boolean if a field has been set.

### GetRedundancyFactor

`func (o *ClusterConfig) GetRedundancyFactor() int32`

GetRedundancyFactor returns the RedundancyFactor field if non-nil, zero value otherwise.

### GetRedundancyFactorOk

`func (o *ClusterConfig) GetRedundancyFactorOk() (*int32, bool)`

GetRedundancyFactorOk returns a tuple with the RedundancyFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedundancyFactor

`func (o *ClusterConfig) SetRedundancyFactor(v int32)`

SetRedundancyFactor sets RedundancyFactor field to given value.

### HasRedundancyFactor

`func (o *ClusterConfig) HasRedundancyFactor() bool`

HasRedundancyFactor returns a boolean if a field has been set.

### GetExternalConfigurations

`func (o *ClusterConfig) GetExternalConfigurations() ExternalConfigurations`

GetExternalConfigurations returns the ExternalConfigurations field if non-nil, zero value otherwise.

### GetExternalConfigurationsOk

`func (o *ClusterConfig) GetExternalConfigurationsOk() (*ExternalConfigurations, bool)`

GetExternalConfigurationsOk returns a tuple with the ExternalConfigurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalConfigurations

`func (o *ClusterConfig) SetExternalConfigurations(v ExternalConfigurations)`

SetExternalConfigurations sets ExternalConfigurations field to given value.

### HasExternalConfigurations

`func (o *ClusterConfig) HasExternalConfigurations() bool`

HasExternalConfigurations returns a boolean if a field has been set.

### GetOperationMode

`func (o *ClusterConfig) GetOperationMode() string`

GetOperationMode returns the OperationMode field if non-nil, zero value otherwise.

### GetOperationModeOk

`func (o *ClusterConfig) GetOperationModeOk() (*string, bool)`

GetOperationModeOk returns a tuple with the OperationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationMode

`func (o *ClusterConfig) SetOperationMode(v string)`

SetOperationMode sets OperationMode field to given value.

### HasOperationMode

`func (o *ClusterConfig) HasOperationMode() bool`

HasOperationMode returns a boolean if a field has been set.

### GetCaCertificateList

`func (o *ClusterConfig) GetCaCertificateList() []CaCert`

GetCaCertificateList returns the CaCertificateList field if non-nil, zero value otherwise.

### GetCaCertificateListOk

`func (o *ClusterConfig) GetCaCertificateListOk() (*[]CaCert, bool)`

GetCaCertificateListOk returns a tuple with the CaCertificateList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaCertificateList

`func (o *ClusterConfig) SetCaCertificateList(v []CaCert)`

SetCaCertificateList sets CaCertificateList field to given value.

### HasCaCertificateList

`func (o *ClusterConfig) HasCaCertificateList() bool`

HasCaCertificateList returns a boolean if a field has been set.

### GetDomainAwarenessLevel

`func (o *ClusterConfig) GetDomainAwarenessLevel() string`

GetDomainAwarenessLevel returns the DomainAwarenessLevel field if non-nil, zero value otherwise.

### GetDomainAwarenessLevelOk

`func (o *ClusterConfig) GetDomainAwarenessLevelOk() (*string, bool)`

GetDomainAwarenessLevelOk returns a tuple with the DomainAwarenessLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainAwarenessLevel

`func (o *ClusterConfig) SetDomainAwarenessLevel(v string)`

SetDomainAwarenessLevel sets DomainAwarenessLevel field to given value.

### HasDomainAwarenessLevel

`func (o *ClusterConfig) HasDomainAwarenessLevel() bool`

HasDomainAwarenessLevel returns a boolean if a field has been set.

### GetEnabledFeatureList

`func (o *ClusterConfig) GetEnabledFeatureList() []string`

GetEnabledFeatureList returns the EnabledFeatureList field if non-nil, zero value otherwise.

### GetEnabledFeatureListOk

`func (o *ClusterConfig) GetEnabledFeatureListOk() (*[]string, bool)`

GetEnabledFeatureListOk returns a tuple with the EnabledFeatureList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledFeatureList

`func (o *ClusterConfig) SetEnabledFeatureList(v []string)`

SetEnabledFeatureList sets EnabledFeatureList field to given value.

### HasEnabledFeatureList

`func (o *ClusterConfig) HasEnabledFeatureList() bool`

HasEnabledFeatureList returns a boolean if a field has been set.

### GetIsAvailable

`func (o *ClusterConfig) GetIsAvailable() bool`

GetIsAvailable returns the IsAvailable field if non-nil, zero value otherwise.

### GetIsAvailableOk

`func (o *ClusterConfig) GetIsAvailableOk() (*bool, bool)`

GetIsAvailableOk returns a tuple with the IsAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAvailable

`func (o *ClusterConfig) SetIsAvailable(v bool)`

SetIsAvailable sets IsAvailable field to given value.

### HasIsAvailable

`func (o *ClusterConfig) HasIsAvailable() bool`

HasIsAvailable returns a boolean if a field has been set.

### GetBuild

`func (o *ClusterConfig) GetBuild() BuildInfo`

GetBuild returns the Build field if non-nil, zero value otherwise.

### GetBuildOk

`func (o *ClusterConfig) GetBuildOk() (*BuildInfo, bool)`

GetBuildOk returns a tuple with the Build field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuild

`func (o *ClusterConfig) SetBuild(v BuildInfo)`

SetBuild sets Build field to given value.

### HasBuild

`func (o *ClusterConfig) HasBuild() bool`

HasBuild returns a boolean if a field has been set.

### GetTimezone

`func (o *ClusterConfig) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *ClusterConfig) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *ClusterConfig) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *ClusterConfig) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetEnableEfficientMetricSync

`func (o *ClusterConfig) GetEnableEfficientMetricSync() bool`

GetEnableEfficientMetricSync returns the EnableEfficientMetricSync field if non-nil, zero value otherwise.

### GetEnableEfficientMetricSyncOk

`func (o *ClusterConfig) GetEnableEfficientMetricSyncOk() (*bool, bool)`

GetEnableEfficientMetricSyncOk returns a tuple with the EnableEfficientMetricSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableEfficientMetricSync

`func (o *ClusterConfig) SetEnableEfficientMetricSync(v bool)`

SetEnableEfficientMetricSync sets EnableEfficientMetricSync field to given value.

### HasEnableEfficientMetricSync

`func (o *ClusterConfig) HasEnableEfficientMetricSync() bool`

HasEnableEfficientMetricSync returns a boolean if a field has been set.

### GetClusterArch

`func (o *ClusterConfig) GetClusterArch() string`

GetClusterArch returns the ClusterArch field if non-nil, zero value otherwise.

### GetClusterArchOk

`func (o *ClusterConfig) GetClusterArchOk() (*string, bool)`

GetClusterArchOk returns a tuple with the ClusterArch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterArch

`func (o *ClusterConfig) SetClusterArch(v string)`

SetClusterArch sets ClusterArch field to given value.

### HasClusterArch

`func (o *ClusterConfig) HasClusterArch() bool`

HasClusterArch returns a boolean if a field has been set.

### GetManagementServerList

`func (o *ClusterConfig) GetManagementServerList() []ClusterManagementServer`

GetManagementServerList returns the ManagementServerList field if non-nil, zero value otherwise.

### GetManagementServerListOk

`func (o *ClusterConfig) GetManagementServerListOk() (*[]ClusterManagementServer, bool)`

GetManagementServerListOk returns a tuple with the ManagementServerList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementServerList

`func (o *ClusterConfig) SetManagementServerList(v []ClusterManagementServer)`

SetManagementServerList sets ManagementServerList field to given value.

### HasManagementServerList

`func (o *ClusterConfig) HasManagementServerList() bool`

HasManagementServerList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



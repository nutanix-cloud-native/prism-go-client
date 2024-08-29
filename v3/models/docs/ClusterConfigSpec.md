# ClusterConfigSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GpuDriverVersion** | Pointer to **string** | GPU driver version. | [optional] 
**ClientAuth** | Pointer to [**ClientAuth**](ClientAuth.md) |  | [optional] 
**AuthorizedPublicKeyList** | Pointer to [**[]PublicKey**](PublicKey.md) | List of valid ssh keys for the cluster. | [optional] 
**SoftwareMap** | Pointer to [**map[string]ClusterSoftware**](ClusterSoftware.md) | Map of software on the cluster with software type as the key.  | [optional] 
**EncryptionStatus** | Pointer to **string** | Cluster encryption status. | [optional] [default to "NOT_SUPPORTED"]
**RedundancyFactor** | Pointer to **int32** | Cluster supported redundancy factor. Default is 2. | [optional] 
**CertificationSigningInfo** | Pointer to [**CertificationSigningInfo**](CertificationSigningInfo.md) |  | [optional] 
**SupportedInformationVerbosity** | Pointer to **string** | Verbosity level settings for populating support information. - &#39;Nothing&#39;: Send nothing - &#39;Basic&#39;: Send basic information - skip core dump and hypervisor            stats information - &#39;BasicPlusCoreDump&#39;: Send basic and core dump information - &#39;All&#39;: Send all information  | [optional] [default to "BASIC_PLUS_CORE_DUMP"]
**ExternalConfigurations** | Pointer to [**ExternalConfigurationsSpec**](ExternalConfigurationsSpec.md) |  | [optional] 
**DomainAwarenessLevel** | Pointer to **string** | Domain awareness supported on cluster. | [optional] [default to "NODE"]
**EnabledFeatureList** | Pointer to **[]string** | Array of enabled features. | [optional] 
**Timezone** | Pointer to **string** | Zone name used in value of TZ environment variable. | [optional] 
**EnableEfficientMetricSync** | Pointer to **bool** | Indicates if downsampling of metrics syncing between PE and PC is enabled or not.  | [optional] 
**OperationMode** | Pointer to **string** | Cluster operation mode. - &#39;NORMAL&#39;: Cluster is operating normally. - &#39;READ_ONLY&#39;: Cluster is operating in read only mode. - &#39;STAND_ALONE&#39;: Only one node is operational in the cluster. This is                  valid only for single node or two node clusters. - &#39;SWITCH_TO_TWO_NODE&#39;: Cluster is moving from single node to two node                         cluster. - &#39;OVERRIDE&#39;: Valid only for single node cluster. If the user wants to               run vms on a single node cluster in read only mode, he               can set the cluster peration mode to override. Writes               will be allowed in override mode.  | [optional] 

## Methods

### NewClusterConfigSpec

`func NewClusterConfigSpec() *ClusterConfigSpec`

NewClusterConfigSpec instantiates a new ClusterConfigSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterConfigSpecWithDefaults

`func NewClusterConfigSpecWithDefaults() *ClusterConfigSpec`

NewClusterConfigSpecWithDefaults instantiates a new ClusterConfigSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGpuDriverVersion

`func (o *ClusterConfigSpec) GetGpuDriverVersion() string`

GetGpuDriverVersion returns the GpuDriverVersion field if non-nil, zero value otherwise.

### GetGpuDriverVersionOk

`func (o *ClusterConfigSpec) GetGpuDriverVersionOk() (*string, bool)`

GetGpuDriverVersionOk returns a tuple with the GpuDriverVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuDriverVersion

`func (o *ClusterConfigSpec) SetGpuDriverVersion(v string)`

SetGpuDriverVersion sets GpuDriverVersion field to given value.

### HasGpuDriverVersion

`func (o *ClusterConfigSpec) HasGpuDriverVersion() bool`

HasGpuDriverVersion returns a boolean if a field has been set.

### GetClientAuth

`func (o *ClusterConfigSpec) GetClientAuth() ClientAuth`

GetClientAuth returns the ClientAuth field if non-nil, zero value otherwise.

### GetClientAuthOk

`func (o *ClusterConfigSpec) GetClientAuthOk() (*ClientAuth, bool)`

GetClientAuthOk returns a tuple with the ClientAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientAuth

`func (o *ClusterConfigSpec) SetClientAuth(v ClientAuth)`

SetClientAuth sets ClientAuth field to given value.

### HasClientAuth

`func (o *ClusterConfigSpec) HasClientAuth() bool`

HasClientAuth returns a boolean if a field has been set.

### GetAuthorizedPublicKeyList

`func (o *ClusterConfigSpec) GetAuthorizedPublicKeyList() []PublicKey`

GetAuthorizedPublicKeyList returns the AuthorizedPublicKeyList field if non-nil, zero value otherwise.

### GetAuthorizedPublicKeyListOk

`func (o *ClusterConfigSpec) GetAuthorizedPublicKeyListOk() (*[]PublicKey, bool)`

GetAuthorizedPublicKeyListOk returns a tuple with the AuthorizedPublicKeyList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizedPublicKeyList

`func (o *ClusterConfigSpec) SetAuthorizedPublicKeyList(v []PublicKey)`

SetAuthorizedPublicKeyList sets AuthorizedPublicKeyList field to given value.

### HasAuthorizedPublicKeyList

`func (o *ClusterConfigSpec) HasAuthorizedPublicKeyList() bool`

HasAuthorizedPublicKeyList returns a boolean if a field has been set.

### GetSoftwareMap

`func (o *ClusterConfigSpec) GetSoftwareMap() map[string]ClusterSoftware`

GetSoftwareMap returns the SoftwareMap field if non-nil, zero value otherwise.

### GetSoftwareMapOk

`func (o *ClusterConfigSpec) GetSoftwareMapOk() (*map[string]ClusterSoftware, bool)`

GetSoftwareMapOk returns a tuple with the SoftwareMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareMap

`func (o *ClusterConfigSpec) SetSoftwareMap(v map[string]ClusterSoftware)`

SetSoftwareMap sets SoftwareMap field to given value.

### HasSoftwareMap

`func (o *ClusterConfigSpec) HasSoftwareMap() bool`

HasSoftwareMap returns a boolean if a field has been set.

### GetEncryptionStatus

`func (o *ClusterConfigSpec) GetEncryptionStatus() string`

GetEncryptionStatus returns the EncryptionStatus field if non-nil, zero value otherwise.

### GetEncryptionStatusOk

`func (o *ClusterConfigSpec) GetEncryptionStatusOk() (*string, bool)`

GetEncryptionStatusOk returns a tuple with the EncryptionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionStatus

`func (o *ClusterConfigSpec) SetEncryptionStatus(v string)`

SetEncryptionStatus sets EncryptionStatus field to given value.

### HasEncryptionStatus

`func (o *ClusterConfigSpec) HasEncryptionStatus() bool`

HasEncryptionStatus returns a boolean if a field has been set.

### GetRedundancyFactor

`func (o *ClusterConfigSpec) GetRedundancyFactor() int32`

GetRedundancyFactor returns the RedundancyFactor field if non-nil, zero value otherwise.

### GetRedundancyFactorOk

`func (o *ClusterConfigSpec) GetRedundancyFactorOk() (*int32, bool)`

GetRedundancyFactorOk returns a tuple with the RedundancyFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedundancyFactor

`func (o *ClusterConfigSpec) SetRedundancyFactor(v int32)`

SetRedundancyFactor sets RedundancyFactor field to given value.

### HasRedundancyFactor

`func (o *ClusterConfigSpec) HasRedundancyFactor() bool`

HasRedundancyFactor returns a boolean if a field has been set.

### GetCertificationSigningInfo

`func (o *ClusterConfigSpec) GetCertificationSigningInfo() CertificationSigningInfo`

GetCertificationSigningInfo returns the CertificationSigningInfo field if non-nil, zero value otherwise.

### GetCertificationSigningInfoOk

`func (o *ClusterConfigSpec) GetCertificationSigningInfoOk() (*CertificationSigningInfo, bool)`

GetCertificationSigningInfoOk returns a tuple with the CertificationSigningInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificationSigningInfo

`func (o *ClusterConfigSpec) SetCertificationSigningInfo(v CertificationSigningInfo)`

SetCertificationSigningInfo sets CertificationSigningInfo field to given value.

### HasCertificationSigningInfo

`func (o *ClusterConfigSpec) HasCertificationSigningInfo() bool`

HasCertificationSigningInfo returns a boolean if a field has been set.

### GetSupportedInformationVerbosity

`func (o *ClusterConfigSpec) GetSupportedInformationVerbosity() string`

GetSupportedInformationVerbosity returns the SupportedInformationVerbosity field if non-nil, zero value otherwise.

### GetSupportedInformationVerbosityOk

`func (o *ClusterConfigSpec) GetSupportedInformationVerbosityOk() (*string, bool)`

GetSupportedInformationVerbosityOk returns a tuple with the SupportedInformationVerbosity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedInformationVerbosity

`func (o *ClusterConfigSpec) SetSupportedInformationVerbosity(v string)`

SetSupportedInformationVerbosity sets SupportedInformationVerbosity field to given value.

### HasSupportedInformationVerbosity

`func (o *ClusterConfigSpec) HasSupportedInformationVerbosity() bool`

HasSupportedInformationVerbosity returns a boolean if a field has been set.

### GetExternalConfigurations

`func (o *ClusterConfigSpec) GetExternalConfigurations() ExternalConfigurationsSpec`

GetExternalConfigurations returns the ExternalConfigurations field if non-nil, zero value otherwise.

### GetExternalConfigurationsOk

`func (o *ClusterConfigSpec) GetExternalConfigurationsOk() (*ExternalConfigurationsSpec, bool)`

GetExternalConfigurationsOk returns a tuple with the ExternalConfigurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalConfigurations

`func (o *ClusterConfigSpec) SetExternalConfigurations(v ExternalConfigurationsSpec)`

SetExternalConfigurations sets ExternalConfigurations field to given value.

### HasExternalConfigurations

`func (o *ClusterConfigSpec) HasExternalConfigurations() bool`

HasExternalConfigurations returns a boolean if a field has been set.

### GetDomainAwarenessLevel

`func (o *ClusterConfigSpec) GetDomainAwarenessLevel() string`

GetDomainAwarenessLevel returns the DomainAwarenessLevel field if non-nil, zero value otherwise.

### GetDomainAwarenessLevelOk

`func (o *ClusterConfigSpec) GetDomainAwarenessLevelOk() (*string, bool)`

GetDomainAwarenessLevelOk returns a tuple with the DomainAwarenessLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainAwarenessLevel

`func (o *ClusterConfigSpec) SetDomainAwarenessLevel(v string)`

SetDomainAwarenessLevel sets DomainAwarenessLevel field to given value.

### HasDomainAwarenessLevel

`func (o *ClusterConfigSpec) HasDomainAwarenessLevel() bool`

HasDomainAwarenessLevel returns a boolean if a field has been set.

### GetEnabledFeatureList

`func (o *ClusterConfigSpec) GetEnabledFeatureList() []string`

GetEnabledFeatureList returns the EnabledFeatureList field if non-nil, zero value otherwise.

### GetEnabledFeatureListOk

`func (o *ClusterConfigSpec) GetEnabledFeatureListOk() (*[]string, bool)`

GetEnabledFeatureListOk returns a tuple with the EnabledFeatureList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledFeatureList

`func (o *ClusterConfigSpec) SetEnabledFeatureList(v []string)`

SetEnabledFeatureList sets EnabledFeatureList field to given value.

### HasEnabledFeatureList

`func (o *ClusterConfigSpec) HasEnabledFeatureList() bool`

HasEnabledFeatureList returns a boolean if a field has been set.

### GetTimezone

`func (o *ClusterConfigSpec) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *ClusterConfigSpec) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *ClusterConfigSpec) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *ClusterConfigSpec) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetEnableEfficientMetricSync

`func (o *ClusterConfigSpec) GetEnableEfficientMetricSync() bool`

GetEnableEfficientMetricSync returns the EnableEfficientMetricSync field if non-nil, zero value otherwise.

### GetEnableEfficientMetricSyncOk

`func (o *ClusterConfigSpec) GetEnableEfficientMetricSyncOk() (*bool, bool)`

GetEnableEfficientMetricSyncOk returns a tuple with the EnableEfficientMetricSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableEfficientMetricSync

`func (o *ClusterConfigSpec) SetEnableEfficientMetricSync(v bool)`

SetEnableEfficientMetricSync sets EnableEfficientMetricSync field to given value.

### HasEnableEfficientMetricSync

`func (o *ClusterConfigSpec) HasEnableEfficientMetricSync() bool`

HasEnableEfficientMetricSync returns a boolean if a field has been set.

### GetOperationMode

`func (o *ClusterConfigSpec) GetOperationMode() string`

GetOperationMode returns the OperationMode field if non-nil, zero value otherwise.

### GetOperationModeOk

`func (o *ClusterConfigSpec) GetOperationModeOk() (*string, bool)`

GetOperationModeOk returns a tuple with the OperationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationMode

`func (o *ClusterConfigSpec) SetOperationMode(v string)`

SetOperationMode sets OperationMode field to given value.

### HasOperationMode

`func (o *ClusterConfigSpec) HasOperationMode() bool`

HasOperationMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



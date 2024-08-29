# NutanixGuestToolsStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvailableVersion** | Pointer to **string** | Version of Nutanix Guest Tools available on the cluster. | [optional] 
**NgtState** | Pointer to **string** | Nutanix guest tools is installed or not. | [optional] 
**IsoMountState** | Pointer to **string** | Desired mount state of Nutanix Guest Tools ISO.  | [optional] 
**GuestOsVersion** | Pointer to **string** | Version of the operating system on the VM. | [optional] 
**State** | Pointer to **string** | Nutanix Guest Tools is enabled or not. | [optional] 
**Version** | Pointer to **string** | Version of Nutanix Guest Tools installed on the VM. | [optional] 
**EnabledCapabilityList** | Pointer to **[]string** | Application names that are enabled. | [optional] 
**Credentials** | Pointer to [**Credentials**](Credentials.md) |  | [optional] 
**VssSnapshotCapable** | Pointer to **bool** | Whether the VM is configured to take VSS snapshots through NGT.  | [optional] 
**IsReachable** | Pointer to **bool** | Communication from VM to CVM is active or not. | [optional] 
**VmMobilityDriversInstalled** | Pointer to **bool** | Whether VM mobility drivers are installed in the VM. | [optional] 

## Methods

### NewNutanixGuestToolsStatus

`func NewNutanixGuestToolsStatus() *NutanixGuestToolsStatus`

NewNutanixGuestToolsStatus instantiates a new NutanixGuestToolsStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNutanixGuestToolsStatusWithDefaults

`func NewNutanixGuestToolsStatusWithDefaults() *NutanixGuestToolsStatus`

NewNutanixGuestToolsStatusWithDefaults instantiates a new NutanixGuestToolsStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailableVersion

`func (o *NutanixGuestToolsStatus) GetAvailableVersion() string`

GetAvailableVersion returns the AvailableVersion field if non-nil, zero value otherwise.

### GetAvailableVersionOk

`func (o *NutanixGuestToolsStatus) GetAvailableVersionOk() (*string, bool)`

GetAvailableVersionOk returns a tuple with the AvailableVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableVersion

`func (o *NutanixGuestToolsStatus) SetAvailableVersion(v string)`

SetAvailableVersion sets AvailableVersion field to given value.

### HasAvailableVersion

`func (o *NutanixGuestToolsStatus) HasAvailableVersion() bool`

HasAvailableVersion returns a boolean if a field has been set.

### GetNgtState

`func (o *NutanixGuestToolsStatus) GetNgtState() string`

GetNgtState returns the NgtState field if non-nil, zero value otherwise.

### GetNgtStateOk

`func (o *NutanixGuestToolsStatus) GetNgtStateOk() (*string, bool)`

GetNgtStateOk returns a tuple with the NgtState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNgtState

`func (o *NutanixGuestToolsStatus) SetNgtState(v string)`

SetNgtState sets NgtState field to given value.

### HasNgtState

`func (o *NutanixGuestToolsStatus) HasNgtState() bool`

HasNgtState returns a boolean if a field has been set.

### GetIsoMountState

`func (o *NutanixGuestToolsStatus) GetIsoMountState() string`

GetIsoMountState returns the IsoMountState field if non-nil, zero value otherwise.

### GetIsoMountStateOk

`func (o *NutanixGuestToolsStatus) GetIsoMountStateOk() (*string, bool)`

GetIsoMountStateOk returns a tuple with the IsoMountState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsoMountState

`func (o *NutanixGuestToolsStatus) SetIsoMountState(v string)`

SetIsoMountState sets IsoMountState field to given value.

### HasIsoMountState

`func (o *NutanixGuestToolsStatus) HasIsoMountState() bool`

HasIsoMountState returns a boolean if a field has been set.

### GetGuestOsVersion

`func (o *NutanixGuestToolsStatus) GetGuestOsVersion() string`

GetGuestOsVersion returns the GuestOsVersion field if non-nil, zero value otherwise.

### GetGuestOsVersionOk

`func (o *NutanixGuestToolsStatus) GetGuestOsVersionOk() (*string, bool)`

GetGuestOsVersionOk returns a tuple with the GuestOsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestOsVersion

`func (o *NutanixGuestToolsStatus) SetGuestOsVersion(v string)`

SetGuestOsVersion sets GuestOsVersion field to given value.

### HasGuestOsVersion

`func (o *NutanixGuestToolsStatus) HasGuestOsVersion() bool`

HasGuestOsVersion returns a boolean if a field has been set.

### GetState

`func (o *NutanixGuestToolsStatus) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *NutanixGuestToolsStatus) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *NutanixGuestToolsStatus) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *NutanixGuestToolsStatus) HasState() bool`

HasState returns a boolean if a field has been set.

### GetVersion

`func (o *NutanixGuestToolsStatus) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *NutanixGuestToolsStatus) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *NutanixGuestToolsStatus) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *NutanixGuestToolsStatus) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetEnabledCapabilityList

`func (o *NutanixGuestToolsStatus) GetEnabledCapabilityList() []string`

GetEnabledCapabilityList returns the EnabledCapabilityList field if non-nil, zero value otherwise.

### GetEnabledCapabilityListOk

`func (o *NutanixGuestToolsStatus) GetEnabledCapabilityListOk() (*[]string, bool)`

GetEnabledCapabilityListOk returns a tuple with the EnabledCapabilityList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledCapabilityList

`func (o *NutanixGuestToolsStatus) SetEnabledCapabilityList(v []string)`

SetEnabledCapabilityList sets EnabledCapabilityList field to given value.

### HasEnabledCapabilityList

`func (o *NutanixGuestToolsStatus) HasEnabledCapabilityList() bool`

HasEnabledCapabilityList returns a boolean if a field has been set.

### GetCredentials

`func (o *NutanixGuestToolsStatus) GetCredentials() Credentials`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *NutanixGuestToolsStatus) GetCredentialsOk() (*Credentials, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *NutanixGuestToolsStatus) SetCredentials(v Credentials)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *NutanixGuestToolsStatus) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetVssSnapshotCapable

`func (o *NutanixGuestToolsStatus) GetVssSnapshotCapable() bool`

GetVssSnapshotCapable returns the VssSnapshotCapable field if non-nil, zero value otherwise.

### GetVssSnapshotCapableOk

`func (o *NutanixGuestToolsStatus) GetVssSnapshotCapableOk() (*bool, bool)`

GetVssSnapshotCapableOk returns a tuple with the VssSnapshotCapable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVssSnapshotCapable

`func (o *NutanixGuestToolsStatus) SetVssSnapshotCapable(v bool)`

SetVssSnapshotCapable sets VssSnapshotCapable field to given value.

### HasVssSnapshotCapable

`func (o *NutanixGuestToolsStatus) HasVssSnapshotCapable() bool`

HasVssSnapshotCapable returns a boolean if a field has been set.

### GetIsReachable

`func (o *NutanixGuestToolsStatus) GetIsReachable() bool`

GetIsReachable returns the IsReachable field if non-nil, zero value otherwise.

### GetIsReachableOk

`func (o *NutanixGuestToolsStatus) GetIsReachableOk() (*bool, bool)`

GetIsReachableOk returns a tuple with the IsReachable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReachable

`func (o *NutanixGuestToolsStatus) SetIsReachable(v bool)`

SetIsReachable sets IsReachable field to given value.

### HasIsReachable

`func (o *NutanixGuestToolsStatus) HasIsReachable() bool`

HasIsReachable returns a boolean if a field has been set.

### GetVmMobilityDriversInstalled

`func (o *NutanixGuestToolsStatus) GetVmMobilityDriversInstalled() bool`

GetVmMobilityDriversInstalled returns the VmMobilityDriversInstalled field if non-nil, zero value otherwise.

### GetVmMobilityDriversInstalledOk

`func (o *NutanixGuestToolsStatus) GetVmMobilityDriversInstalledOk() (*bool, bool)`

GetVmMobilityDriversInstalledOk returns a tuple with the VmMobilityDriversInstalled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmMobilityDriversInstalled

`func (o *NutanixGuestToolsStatus) SetVmMobilityDriversInstalled(v bool)`

SetVmMobilityDriversInstalled sets VmMobilityDriversInstalled field to given value.

### HasVmMobilityDriversInstalled

`func (o *NutanixGuestToolsStatus) HasVmMobilityDriversInstalled() bool`

HasVmMobilityDriversInstalled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



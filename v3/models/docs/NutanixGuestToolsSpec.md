# NutanixGuestToolsSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NgtState** | Pointer to **string** | Nutanix guest tools is installed or not. | [optional] 
**IsoMountState** | Pointer to **string** | Desired mount state of Nutanix Guest Tools ISO.  | [optional] 
**State** | Pointer to **string** | Nutanix Guest Tools is enabled or not. | [optional] 
**Version** | Pointer to **string** | Desired Version of Nutanix Guest Tools installed on the VM. | [optional] 
**EnabledCapabilityList** | Pointer to **[]string** | Application names that are enabled. | [optional] 
**Credentials** | Pointer to [**Credentials**](Credentials.md) |  | [optional] 

## Methods

### NewNutanixGuestToolsSpec

`func NewNutanixGuestToolsSpec() *NutanixGuestToolsSpec`

NewNutanixGuestToolsSpec instantiates a new NutanixGuestToolsSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNutanixGuestToolsSpecWithDefaults

`func NewNutanixGuestToolsSpecWithDefaults() *NutanixGuestToolsSpec`

NewNutanixGuestToolsSpecWithDefaults instantiates a new NutanixGuestToolsSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNgtState

`func (o *NutanixGuestToolsSpec) GetNgtState() string`

GetNgtState returns the NgtState field if non-nil, zero value otherwise.

### GetNgtStateOk

`func (o *NutanixGuestToolsSpec) GetNgtStateOk() (*string, bool)`

GetNgtStateOk returns a tuple with the NgtState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNgtState

`func (o *NutanixGuestToolsSpec) SetNgtState(v string)`

SetNgtState sets NgtState field to given value.

### HasNgtState

`func (o *NutanixGuestToolsSpec) HasNgtState() bool`

HasNgtState returns a boolean if a field has been set.

### GetIsoMountState

`func (o *NutanixGuestToolsSpec) GetIsoMountState() string`

GetIsoMountState returns the IsoMountState field if non-nil, zero value otherwise.

### GetIsoMountStateOk

`func (o *NutanixGuestToolsSpec) GetIsoMountStateOk() (*string, bool)`

GetIsoMountStateOk returns a tuple with the IsoMountState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsoMountState

`func (o *NutanixGuestToolsSpec) SetIsoMountState(v string)`

SetIsoMountState sets IsoMountState field to given value.

### HasIsoMountState

`func (o *NutanixGuestToolsSpec) HasIsoMountState() bool`

HasIsoMountState returns a boolean if a field has been set.

### GetState

`func (o *NutanixGuestToolsSpec) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *NutanixGuestToolsSpec) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *NutanixGuestToolsSpec) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *NutanixGuestToolsSpec) HasState() bool`

HasState returns a boolean if a field has been set.

### GetVersion

`func (o *NutanixGuestToolsSpec) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *NutanixGuestToolsSpec) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *NutanixGuestToolsSpec) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *NutanixGuestToolsSpec) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetEnabledCapabilityList

`func (o *NutanixGuestToolsSpec) GetEnabledCapabilityList() []string`

GetEnabledCapabilityList returns the EnabledCapabilityList field if non-nil, zero value otherwise.

### GetEnabledCapabilityListOk

`func (o *NutanixGuestToolsSpec) GetEnabledCapabilityListOk() (*[]string, bool)`

GetEnabledCapabilityListOk returns a tuple with the EnabledCapabilityList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledCapabilityList

`func (o *NutanixGuestToolsSpec) SetEnabledCapabilityList(v []string)`

SetEnabledCapabilityList sets EnabledCapabilityList field to given value.

### HasEnabledCapabilityList

`func (o *NutanixGuestToolsSpec) HasEnabledCapabilityList() bool`

HasEnabledCapabilityList returns a boolean if a field has been set.

### GetCredentials

`func (o *NutanixGuestToolsSpec) GetCredentials() Credentials`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *NutanixGuestToolsSpec) GetCredentialsOk() (*Credentials, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *NutanixGuestToolsSpec) SetCredentials(v Credentials)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *NutanixGuestToolsSpec) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



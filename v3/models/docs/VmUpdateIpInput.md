# VmUpdateIpInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UpdateList** | Pointer to [**[]NicUpdateIpInfo**](NicUpdateIpInfo.md) | List of NICs to update new IP for. | [optional] 

## Methods

### NewVmUpdateIpInput

`func NewVmUpdateIpInput() *VmUpdateIpInput`

NewVmUpdateIpInput instantiates a new VmUpdateIpInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmUpdateIpInputWithDefaults

`func NewVmUpdateIpInputWithDefaults() *VmUpdateIpInput`

NewVmUpdateIpInputWithDefaults instantiates a new VmUpdateIpInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUpdateList

`func (o *VmUpdateIpInput) GetUpdateList() []NicUpdateIpInfo`

GetUpdateList returns the UpdateList field if non-nil, zero value otherwise.

### GetUpdateListOk

`func (o *VmUpdateIpInput) GetUpdateListOk() (*[]NicUpdateIpInfo, bool)`

GetUpdateListOk returns a tuple with the UpdateList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateList

`func (o *VmUpdateIpInput) SetUpdateList(v []NicUpdateIpInfo)`

SetUpdateList sets UpdateList field to given value.

### HasUpdateList

`func (o *VmUpdateIpInput) HasUpdateList() bool`

HasUpdateList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



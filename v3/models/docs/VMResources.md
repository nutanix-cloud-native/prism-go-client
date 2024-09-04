# VmResources

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NicList** | Pointer to [**[]VmNic**](VmNic.md) | NICs attached to the VM. | [optional] 

## Methods

### NewVmResources

`func NewVmResources() *VmResources`

NewVmResources instantiates a new VmResources object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmResourcesWithDefaults

`func NewVmResourcesWithDefaults() *VmResources`

NewVmResourcesWithDefaults instantiates a new VmResources object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNicList

`func (o *VmResources) GetNicList() []VmNic`

GetNicList returns the NicList field if non-nil, zero value otherwise.

### GetNicListOk

`func (o *VmResources) GetNicListOk() (*[]VmNic, bool)`

GetNicListOk returns a tuple with the NicList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNicList

`func (o *VmResources) SetNicList(v []VmNic)`

SetNicList sets NicList field to given value.

### HasNicList

`func (o *VmResources) HasNicList() bool`

HasNicList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



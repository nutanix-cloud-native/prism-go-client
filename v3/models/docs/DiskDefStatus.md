# DiskDefStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | Pointer to **string** | The state of the disk. | [optional] 
**MessageList** | Pointer to [**[]MessageResource**](MessageResource.md) | Any error messages for the disk, if in an error state. | [optional] 
**Resources** | [**DiskDefStatusResources**](DiskDefStatusResources.md) |  | 

## Methods

### NewDiskDefStatus

`func NewDiskDefStatus(resources DiskDefStatusResources, ) *DiskDefStatus`

NewDiskDefStatus instantiates a new DiskDefStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiskDefStatusWithDefaults

`func NewDiskDefStatusWithDefaults() *DiskDefStatus`

NewDiskDefStatusWithDefaults instantiates a new DiskDefStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *DiskDefStatus) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *DiskDefStatus) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *DiskDefStatus) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *DiskDefStatus) HasState() bool`

HasState returns a boolean if a field has been set.

### GetMessageList

`func (o *DiskDefStatus) GetMessageList() []MessageResource`

GetMessageList returns the MessageList field if non-nil, zero value otherwise.

### GetMessageListOk

`func (o *DiskDefStatus) GetMessageListOk() (*[]MessageResource, bool)`

GetMessageListOk returns a tuple with the MessageList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageList

`func (o *DiskDefStatus) SetMessageList(v []MessageResource)`

SetMessageList sets MessageList field to given value.

### HasMessageList

`func (o *DiskDefStatus) HasMessageList() bool`

HasMessageList returns a boolean if a field has been set.

### GetResources

`func (o *DiskDefStatus) GetResources() DiskDefStatusResources`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *DiskDefStatus) GetResourcesOk() (*DiskDefStatusResources, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *DiskDefStatus) SetResources(v DiskDefStatusResources)`

SetResources sets Resources field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



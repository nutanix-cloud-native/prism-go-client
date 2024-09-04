# DirectoryServiceDefStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | Pointer to **string** | The state of the directory service configuration. | [optional] 
**MessageList** | Pointer to [**[]MessageResource**](MessageResource.md) |  | [optional] 
**Name** | **string** | Name of the directory service. | 
**Resources** | [**DirectoryServiceResources**](DirectoryServiceResources.md) |  | 

## Methods

### NewDirectoryServiceDefStatus

`func NewDirectoryServiceDefStatus(name string, resources DirectoryServiceResources, ) *DirectoryServiceDefStatus`

NewDirectoryServiceDefStatus instantiates a new DirectoryServiceDefStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDirectoryServiceDefStatusWithDefaults

`func NewDirectoryServiceDefStatusWithDefaults() *DirectoryServiceDefStatus`

NewDirectoryServiceDefStatusWithDefaults instantiates a new DirectoryServiceDefStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *DirectoryServiceDefStatus) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *DirectoryServiceDefStatus) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *DirectoryServiceDefStatus) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *DirectoryServiceDefStatus) HasState() bool`

HasState returns a boolean if a field has been set.

### GetMessageList

`func (o *DirectoryServiceDefStatus) GetMessageList() []MessageResource`

GetMessageList returns the MessageList field if non-nil, zero value otherwise.

### GetMessageListOk

`func (o *DirectoryServiceDefStatus) GetMessageListOk() (*[]MessageResource, bool)`

GetMessageListOk returns a tuple with the MessageList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageList

`func (o *DirectoryServiceDefStatus) SetMessageList(v []MessageResource)`

SetMessageList sets MessageList field to given value.

### HasMessageList

`func (o *DirectoryServiceDefStatus) HasMessageList() bool`

HasMessageList returns a boolean if a field has been set.

### GetName

`func (o *DirectoryServiceDefStatus) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DirectoryServiceDefStatus) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DirectoryServiceDefStatus) SetName(v string)`

SetName sets Name field to given value.


### GetResources

`func (o *DirectoryServiceDefStatus) GetResources() DirectoryServiceResources`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *DirectoryServiceDefStatus) GetResourcesOk() (*DirectoryServiceResources, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *DirectoryServiceDefStatus) SetResources(v DirectoryServiceResources)`

SetResources sets Resources field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



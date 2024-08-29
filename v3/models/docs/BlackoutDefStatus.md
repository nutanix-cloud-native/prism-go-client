# BlackoutDefStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | Pointer to **string** | The state of the entity. | [optional] 
**MessageList** | Pointer to [**[]MessageResource**](MessageResource.md) |  | [optional] 
**Resources** | [**BlackoutResources**](BlackoutResources.md) |  | 

## Methods

### NewBlackoutDefStatus

`func NewBlackoutDefStatus(resources BlackoutResources, ) *BlackoutDefStatus`

NewBlackoutDefStatus instantiates a new BlackoutDefStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlackoutDefStatusWithDefaults

`func NewBlackoutDefStatusWithDefaults() *BlackoutDefStatus`

NewBlackoutDefStatusWithDefaults instantiates a new BlackoutDefStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *BlackoutDefStatus) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *BlackoutDefStatus) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *BlackoutDefStatus) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *BlackoutDefStatus) HasState() bool`

HasState returns a boolean if a field has been set.

### GetMessageList

`func (o *BlackoutDefStatus) GetMessageList() []MessageResource`

GetMessageList returns the MessageList field if non-nil, zero value otherwise.

### GetMessageListOk

`func (o *BlackoutDefStatus) GetMessageListOk() (*[]MessageResource, bool)`

GetMessageListOk returns a tuple with the MessageList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageList

`func (o *BlackoutDefStatus) SetMessageList(v []MessageResource)`

SetMessageList sets MessageList field to given value.

### HasMessageList

`func (o *BlackoutDefStatus) HasMessageList() bool`

HasMessageList returns a boolean if a field has been set.

### GetResources

`func (o *BlackoutDefStatus) GetResources() BlackoutResources`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *BlackoutDefStatus) GetResourcesOk() (*BlackoutResources, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *BlackoutDefStatus) SetResources(v BlackoutResources)`

SetResources sets Resources field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



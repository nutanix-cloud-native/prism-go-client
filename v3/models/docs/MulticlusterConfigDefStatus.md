# MulticlusterConfigDefStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | Pointer to **string** | The state of the multicluster configuration request. | [optional] 
**MessageList** | Pointer to [**[]MessageResource**](MessageResource.md) |  | [optional] 

## Methods

### NewMulticlusterConfigDefStatus

`func NewMulticlusterConfigDefStatus() *MulticlusterConfigDefStatus`

NewMulticlusterConfigDefStatus instantiates a new MulticlusterConfigDefStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMulticlusterConfigDefStatusWithDefaults

`func NewMulticlusterConfigDefStatusWithDefaults() *MulticlusterConfigDefStatus`

NewMulticlusterConfigDefStatusWithDefaults instantiates a new MulticlusterConfigDefStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *MulticlusterConfigDefStatus) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *MulticlusterConfigDefStatus) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *MulticlusterConfigDefStatus) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *MulticlusterConfigDefStatus) HasState() bool`

HasState returns a boolean if a field has been set.

### GetMessageList

`func (o *MulticlusterConfigDefStatus) GetMessageList() []MessageResource`

GetMessageList returns the MessageList field if non-nil, zero value otherwise.

### GetMessageListOk

`func (o *MulticlusterConfigDefStatus) GetMessageListOk() (*[]MessageResource, bool)`

GetMessageListOk returns a tuple with the MessageList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageList

`func (o *MulticlusterConfigDefStatus) SetMessageList(v []MessageResource)`

SetMessageList sets MessageList field to given value.

### HasMessageList

`func (o *MulticlusterConfigDefStatus) HasMessageList() bool`

HasMessageList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



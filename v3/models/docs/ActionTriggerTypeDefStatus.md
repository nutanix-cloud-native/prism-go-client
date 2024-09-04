# ActionTriggerTypeDefStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | Pointer to **string** | The state of the entity. | [optional] 
**MessageList** | Pointer to [**[]MessageResource**](MessageResource.md) |  | [optional] 
**Resources** | [**ActionTriggerType**](ActionTriggerType.md) |  | 

## Methods

### NewActionTriggerTypeDefStatus

`func NewActionTriggerTypeDefStatus(resources ActionTriggerType, ) *ActionTriggerTypeDefStatus`

NewActionTriggerTypeDefStatus instantiates a new ActionTriggerTypeDefStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionTriggerTypeDefStatusWithDefaults

`func NewActionTriggerTypeDefStatusWithDefaults() *ActionTriggerTypeDefStatus`

NewActionTriggerTypeDefStatusWithDefaults instantiates a new ActionTriggerTypeDefStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *ActionTriggerTypeDefStatus) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ActionTriggerTypeDefStatus) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ActionTriggerTypeDefStatus) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ActionTriggerTypeDefStatus) HasState() bool`

HasState returns a boolean if a field has been set.

### GetMessageList

`func (o *ActionTriggerTypeDefStatus) GetMessageList() []MessageResource`

GetMessageList returns the MessageList field if non-nil, zero value otherwise.

### GetMessageListOk

`func (o *ActionTriggerTypeDefStatus) GetMessageListOk() (*[]MessageResource, bool)`

GetMessageListOk returns a tuple with the MessageList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageList

`func (o *ActionTriggerTypeDefStatus) SetMessageList(v []MessageResource)`

SetMessageList sets MessageList field to given value.

### HasMessageList

`func (o *ActionTriggerTypeDefStatus) HasMessageList() bool`

HasMessageList returns a boolean if a field has been set.

### GetResources

`func (o *ActionTriggerTypeDefStatus) GetResources() ActionTriggerType`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *ActionTriggerTypeDefStatus) GetResourcesOk() (*ActionTriggerType, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *ActionTriggerTypeDefStatus) SetResources(v ActionTriggerType)`

SetResources sets Resources field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



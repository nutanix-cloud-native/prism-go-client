# ActionTemplateDefStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | Pointer to **string** | The state of the entity. | [optional] 
**MessageList** | Pointer to [**[]MessageResource**](MessageResource.md) |  | [optional] 
**Resources** | [**ActionTemplateResources**](ActionTemplateResources.md) |  | 

## Methods

### NewActionTemplateDefStatus

`func NewActionTemplateDefStatus(resources ActionTemplateResources, ) *ActionTemplateDefStatus`

NewActionTemplateDefStatus instantiates a new ActionTemplateDefStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionTemplateDefStatusWithDefaults

`func NewActionTemplateDefStatusWithDefaults() *ActionTemplateDefStatus`

NewActionTemplateDefStatusWithDefaults instantiates a new ActionTemplateDefStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *ActionTemplateDefStatus) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ActionTemplateDefStatus) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ActionTemplateDefStatus) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ActionTemplateDefStatus) HasState() bool`

HasState returns a boolean if a field has been set.

### GetMessageList

`func (o *ActionTemplateDefStatus) GetMessageList() []MessageResource`

GetMessageList returns the MessageList field if non-nil, zero value otherwise.

### GetMessageListOk

`func (o *ActionTemplateDefStatus) GetMessageListOk() (*[]MessageResource, bool)`

GetMessageListOk returns a tuple with the MessageList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageList

`func (o *ActionTemplateDefStatus) SetMessageList(v []MessageResource)`

SetMessageList sets MessageList field to given value.

### HasMessageList

`func (o *ActionTemplateDefStatus) HasMessageList() bool`

HasMessageList returns a boolean if a field has been set.

### GetResources

`func (o *ActionTemplateDefStatus) GetResources() ActionTemplateResources`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *ActionTemplateDefStatus) GetResourcesOk() (*ActionTemplateResources, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *ActionTemplateDefStatus) SetResources(v ActionTemplateResources)`

SetResources sets Resources field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



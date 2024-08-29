# StreamingPolicyDefStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | Pointer to **string** | The state of the streaming policy. | [optional] 
**MessageList** | Pointer to [**[]MessageResource**](MessageResource.md) | Any error messages for the streaming policy, if in an error state.  | [optional] 
**Name** | **string** | Policy name | 
**Resources** | [**StreamingPolicyDetails**](StreamingPolicyDetails.md) |  | 
**Description** | **string** | Policy description | 

## Methods

### NewStreamingPolicyDefStatus

`func NewStreamingPolicyDefStatus(name string, resources StreamingPolicyDetails, description string, ) *StreamingPolicyDefStatus`

NewStreamingPolicyDefStatus instantiates a new StreamingPolicyDefStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamingPolicyDefStatusWithDefaults

`func NewStreamingPolicyDefStatusWithDefaults() *StreamingPolicyDefStatus`

NewStreamingPolicyDefStatusWithDefaults instantiates a new StreamingPolicyDefStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *StreamingPolicyDefStatus) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *StreamingPolicyDefStatus) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *StreamingPolicyDefStatus) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *StreamingPolicyDefStatus) HasState() bool`

HasState returns a boolean if a field has been set.

### GetMessageList

`func (o *StreamingPolicyDefStatus) GetMessageList() []MessageResource`

GetMessageList returns the MessageList field if non-nil, zero value otherwise.

### GetMessageListOk

`func (o *StreamingPolicyDefStatus) GetMessageListOk() (*[]MessageResource, bool)`

GetMessageListOk returns a tuple with the MessageList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageList

`func (o *StreamingPolicyDefStatus) SetMessageList(v []MessageResource)`

SetMessageList sets MessageList field to given value.

### HasMessageList

`func (o *StreamingPolicyDefStatus) HasMessageList() bool`

HasMessageList returns a boolean if a field has been set.

### GetName

`func (o *StreamingPolicyDefStatus) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StreamingPolicyDefStatus) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StreamingPolicyDefStatus) SetName(v string)`

SetName sets Name field to given value.


### GetResources

`func (o *StreamingPolicyDefStatus) GetResources() StreamingPolicyDetails`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *StreamingPolicyDefStatus) GetResourcesOk() (*StreamingPolicyDetails, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *StreamingPolicyDefStatus) SetResources(v StreamingPolicyDetails)`

SetResources sets Resources field to given value.


### GetDescription

`func (o *StreamingPolicyDefStatus) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *StreamingPolicyDefStatus) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *StreamingPolicyDefStatus) SetDescription(v string)`

SetDescription sets Description field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



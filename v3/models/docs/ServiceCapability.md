# ServiceCapability

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | Pointer to **bool** | The current state of the capability. For example, if a service can be enabled or not.  | [optional] 
**ValidationMessageList** | Pointer to [**[]ServiceCapabilityValidationMessageListInner**](ServiceCapabilityValidationMessageListInner.md) | List of messages explaining the current state of the capability. For example, if a service cannot be enabled, this list will contain the reasons for the same.  | [optional] 

## Methods

### NewServiceCapability

`func NewServiceCapability() *ServiceCapability`

NewServiceCapability instantiates a new ServiceCapability object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceCapabilityWithDefaults

`func NewServiceCapabilityWithDefaults() *ServiceCapability`

NewServiceCapabilityWithDefaults instantiates a new ServiceCapability object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *ServiceCapability) GetState() bool`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ServiceCapability) GetStateOk() (*bool, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ServiceCapability) SetState(v bool)`

SetState sets State field to given value.

### HasState

`func (o *ServiceCapability) HasState() bool`

HasState returns a boolean if a field has been set.

### GetValidationMessageList

`func (o *ServiceCapability) GetValidationMessageList() []ServiceCapabilityValidationMessageListInner`

GetValidationMessageList returns the ValidationMessageList field if non-nil, zero value otherwise.

### GetValidationMessageListOk

`func (o *ServiceCapability) GetValidationMessageListOk() (*[]ServiceCapabilityValidationMessageListInner, bool)`

GetValidationMessageListOk returns a tuple with the ValidationMessageList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationMessageList

`func (o *ServiceCapability) SetValidationMessageList(v []ServiceCapabilityValidationMessageListInner)`

SetValidationMessageList sets ValidationMessageList field to given value.

### HasValidationMessageList

`func (o *ServiceCapability) HasValidationMessageList() bool`

HasValidationMessageList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



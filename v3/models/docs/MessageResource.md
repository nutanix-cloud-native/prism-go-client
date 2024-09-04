# MessageResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | If state is ERROR, a message describing the error. | 
**Reason** | **string** | If state is ERROR, a machine-readable snake-cased string. | 
**Details** | Pointer to **map[string]string** | Custom key-value details relevant to the status. | [optional] [readonly] 

## Methods

### NewMessageResource

`func NewMessageResource(message string, reason string, ) *MessageResource`

NewMessageResource instantiates a new MessageResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageResourceWithDefaults

`func NewMessageResourceWithDefaults() *MessageResource`

NewMessageResourceWithDefaults instantiates a new MessageResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *MessageResource) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *MessageResource) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *MessageResource) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetReason

`func (o *MessageResource) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *MessageResource) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *MessageResource) SetReason(v string)`

SetReason sets Reason field to given value.


### GetDetails

`func (o *MessageResource) GetDetails() map[string]string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *MessageResource) GetDetailsOk() (*map[string]string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *MessageResource) SetDetails(v map[string]string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *MessageResource) HasDetails() bool`

HasDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



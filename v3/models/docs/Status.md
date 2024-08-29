# Status

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | Only value possible is \&quot;failure\&quot;. | [optional] [readonly] 
**Kind** | Pointer to **string** | The entitys kind. i.e. \&quot;status\&quot;. | [optional] [readonly] 
**Code** | Pointer to **int32** | The HTTP error code. | [optional] [readonly] 
**Reason** | Pointer to **string** | One snake case word. | [optional] [readonly] 
**Details** | Pointer to **map[string]string** | Custom key-value details relevant to the status. | [optional] [readonly] 
**Message** | Pointer to **string** | A sentence explaining the reason for the status. | [optional] [readonly] 
**ApiVersion** | Pointer to **string** | API Version of the Nutanix v3 API framework. | [optional] [readonly] [default to "3.1.0"]

## Methods

### NewStatus

`func NewStatus() *Status`

NewStatus instantiates a new Status object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatusWithDefaults

`func NewStatusWithDefaults() *Status`

NewStatusWithDefaults instantiates a new Status object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *Status) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Status) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Status) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Status) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetKind

`func (o *Status) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *Status) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *Status) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *Status) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetCode

`func (o *Status) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *Status) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *Status) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *Status) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetReason

`func (o *Status) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *Status) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *Status) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *Status) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetDetails

`func (o *Status) GetDetails() map[string]string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *Status) GetDetailsOk() (*map[string]string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *Status) SetDetails(v map[string]string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *Status) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetMessage

`func (o *Status) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *Status) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *Status) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *Status) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetApiVersion

`func (o *Status) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *Status) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *Status) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *Status) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# MicrosegServiceConfigInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsValidationOnly** | Pointer to **bool** | Flag indicating whether to do Microsegmentation enablement validation only.  | [optional] 
**State** | Pointer to **string** | The desired state of Microsegmentation. | [optional] 

## Methods

### NewMicrosegServiceConfigInput

`func NewMicrosegServiceConfigInput() *MicrosegServiceConfigInput`

NewMicrosegServiceConfigInput instantiates a new MicrosegServiceConfigInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosegServiceConfigInputWithDefaults

`func NewMicrosegServiceConfigInputWithDefaults() *MicrosegServiceConfigInput`

NewMicrosegServiceConfigInputWithDefaults instantiates a new MicrosegServiceConfigInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsValidationOnly

`func (o *MicrosegServiceConfigInput) GetIsValidationOnly() bool`

GetIsValidationOnly returns the IsValidationOnly field if non-nil, zero value otherwise.

### GetIsValidationOnlyOk

`func (o *MicrosegServiceConfigInput) GetIsValidationOnlyOk() (*bool, bool)`

GetIsValidationOnlyOk returns a tuple with the IsValidationOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsValidationOnly

`func (o *MicrosegServiceConfigInput) SetIsValidationOnly(v bool)`

SetIsValidationOnly sets IsValidationOnly field to given value.

### HasIsValidationOnly

`func (o *MicrosegServiceConfigInput) HasIsValidationOnly() bool`

HasIsValidationOnly returns a boolean if a field has been set.

### GetState

`func (o *MicrosegServiceConfigInput) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *MicrosegServiceConfigInput) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *MicrosegServiceConfigInput) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *MicrosegServiceConfigInput) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



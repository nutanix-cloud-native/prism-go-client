# XfitServiceInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | Pointer to **string** | The desired state of xfit service. | [optional] 
**ValidationOnly** | Pointer to **bool** | Flag indicating whether to do Xfit enablement validation only.  | [optional] 

## Methods

### NewXfitServiceInput

`func NewXfitServiceInput() *XfitServiceInput`

NewXfitServiceInput instantiates a new XfitServiceInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewXfitServiceInputWithDefaults

`func NewXfitServiceInputWithDefaults() *XfitServiceInput`

NewXfitServiceInputWithDefaults instantiates a new XfitServiceInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *XfitServiceInput) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *XfitServiceInput) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *XfitServiceInput) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *XfitServiceInput) HasState() bool`

HasState returns a boolean if a field has been set.

### GetValidationOnly

`func (o *XfitServiceInput) GetValidationOnly() bool`

GetValidationOnly returns the ValidationOnly field if non-nil, zero value otherwise.

### GetValidationOnlyOk

`func (o *XfitServiceInput) GetValidationOnlyOk() (*bool, bool)`

GetValidationOnlyOk returns a tuple with the ValidationOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationOnly

`func (o *XfitServiceInput) SetValidationOnly(v bool)`

SetValidationOnly sets ValidationOnly field to given value.

### HasValidationOnly

`func (o *XfitServiceInput) HasValidationOnly() bool`

HasValidationOnly returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



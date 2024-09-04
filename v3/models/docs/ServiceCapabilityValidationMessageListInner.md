# ServiceCapabilityValidationMessageListInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorDetail** | Pointer to **string** | The error detail for the validation failure.  | [optional] 
**ValidationType** | Pointer to **string** | The type of validation performed.  | [optional] 

## Methods

### NewServiceCapabilityValidationMessageListInner

`func NewServiceCapabilityValidationMessageListInner() *ServiceCapabilityValidationMessageListInner`

NewServiceCapabilityValidationMessageListInner instantiates a new ServiceCapabilityValidationMessageListInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceCapabilityValidationMessageListInnerWithDefaults

`func NewServiceCapabilityValidationMessageListInnerWithDefaults() *ServiceCapabilityValidationMessageListInner`

NewServiceCapabilityValidationMessageListInnerWithDefaults instantiates a new ServiceCapabilityValidationMessageListInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorDetail

`func (o *ServiceCapabilityValidationMessageListInner) GetErrorDetail() string`

GetErrorDetail returns the ErrorDetail field if non-nil, zero value otherwise.

### GetErrorDetailOk

`func (o *ServiceCapabilityValidationMessageListInner) GetErrorDetailOk() (*string, bool)`

GetErrorDetailOk returns a tuple with the ErrorDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDetail

`func (o *ServiceCapabilityValidationMessageListInner) SetErrorDetail(v string)`

SetErrorDetail sets ErrorDetail field to given value.

### HasErrorDetail

`func (o *ServiceCapabilityValidationMessageListInner) HasErrorDetail() bool`

HasErrorDetail returns a boolean if a field has been set.

### GetValidationType

`func (o *ServiceCapabilityValidationMessageListInner) GetValidationType() string`

GetValidationType returns the ValidationType field if non-nil, zero value otherwise.

### GetValidationTypeOk

`func (o *ServiceCapabilityValidationMessageListInner) GetValidationTypeOk() (*string, bool)`

GetValidationTypeOk returns a tuple with the ValidationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationType

`func (o *ServiceCapabilityValidationMessageListInner) SetValidationType(v string)`

SetValidationType sets ValidationType field to given value.

### HasValidationType

`func (o *ServiceCapabilityValidationMessageListInner) HasValidationType() bool`

HasValidationType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



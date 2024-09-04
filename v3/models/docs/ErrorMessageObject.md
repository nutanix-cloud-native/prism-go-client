# ErrorMessageObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ParamList** | Pointer to **[]string** | params to enhance error message | [optional] 
**ErrorMessage** | Pointer to **string** | error message with param place holders | [optional] 

## Methods

### NewErrorMessageObject

`func NewErrorMessageObject() *ErrorMessageObject`

NewErrorMessageObject instantiates a new ErrorMessageObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorMessageObjectWithDefaults

`func NewErrorMessageObjectWithDefaults() *ErrorMessageObject`

NewErrorMessageObjectWithDefaults instantiates a new ErrorMessageObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParamList

`func (o *ErrorMessageObject) GetParamList() []string`

GetParamList returns the ParamList field if non-nil, zero value otherwise.

### GetParamListOk

`func (o *ErrorMessageObject) GetParamListOk() (*[]string, bool)`

GetParamListOk returns a tuple with the ParamList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParamList

`func (o *ErrorMessageObject) SetParamList(v []string)`

SetParamList sets ParamList field to given value.

### HasParamList

`func (o *ErrorMessageObject) HasParamList() bool`

HasParamList returns a boolean if a field has been set.

### GetErrorMessage

`func (o *ErrorMessageObject) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *ErrorMessageObject) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *ErrorMessageObject) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *ErrorMessageObject) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



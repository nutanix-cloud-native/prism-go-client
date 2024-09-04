# ParameterError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorType** | Pointer to **string** | type of error like ValidationError | [optional] 
**ErrorObject** | Pointer to [**ErrorMessageObject**](ErrorMessageObject.md) |  | [optional] 
**ParameterName** | Pointer to **string** | name of the parameter. | [optional] 

## Methods

### NewParameterError

`func NewParameterError() *ParameterError`

NewParameterError instantiates a new ParameterError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParameterErrorWithDefaults

`func NewParameterErrorWithDefaults() *ParameterError`

NewParameterErrorWithDefaults instantiates a new ParameterError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorType

`func (o *ParameterError) GetErrorType() string`

GetErrorType returns the ErrorType field if non-nil, zero value otherwise.

### GetErrorTypeOk

`func (o *ParameterError) GetErrorTypeOk() (*string, bool)`

GetErrorTypeOk returns a tuple with the ErrorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorType

`func (o *ParameterError) SetErrorType(v string)`

SetErrorType sets ErrorType field to given value.

### HasErrorType

`func (o *ParameterError) HasErrorType() bool`

HasErrorType returns a boolean if a field has been set.

### GetErrorObject

`func (o *ParameterError) GetErrorObject() ErrorMessageObject`

GetErrorObject returns the ErrorObject field if non-nil, zero value otherwise.

### GetErrorObjectOk

`func (o *ParameterError) GetErrorObjectOk() (*ErrorMessageObject, bool)`

GetErrorObjectOk returns a tuple with the ErrorObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorObject

`func (o *ParameterError) SetErrorObject(v ErrorMessageObject)`

SetErrorObject sets ErrorObject field to given value.

### HasErrorObject

`func (o *ParameterError) HasErrorObject() bool`

HasErrorObject returns a boolean if a field has been set.

### GetParameterName

`func (o *ParameterError) GetParameterName() string`

GetParameterName returns the ParameterName field if non-nil, zero value otherwise.

### GetParameterNameOk

`func (o *ParameterError) GetParameterNameOk() (*string, bool)`

GetParameterNameOk returns a tuple with the ParameterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameterName

`func (o *ParameterError) SetParameterName(v string)`

SetParameterName sets ParameterName field to given value.

### HasParameterName

`func (o *ParameterError) HasParameterName() bool`

HasParameterName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



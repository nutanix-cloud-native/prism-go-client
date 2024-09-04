# ComponentError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceUuid** | Pointer to **string** | Id to uniquely identify action/trigger in instance list. | [optional] 
**ErrorObjectList** | Pointer to [**[]ErrorMessageObject**](ErrorMessageObject.md) | list of error message objects. | [optional] 
**ParameterErrorList** | Pointer to [**[]ParameterError**](ParameterError.md) | list of parameter error messages. | [optional] 
**ComponentType** | Pointer to **string** | component type like action or trigger. | [optional] 
**ComponentName** | Pointer to **string** | name of the component. | [optional] 

## Methods

### NewComponentError

`func NewComponentError() *ComponentError`

NewComponentError instantiates a new ComponentError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComponentErrorWithDefaults

`func NewComponentErrorWithDefaults() *ComponentError`

NewComponentErrorWithDefaults instantiates a new ComponentError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceUuid

`func (o *ComponentError) GetInstanceUuid() string`

GetInstanceUuid returns the InstanceUuid field if non-nil, zero value otherwise.

### GetInstanceUuidOk

`func (o *ComponentError) GetInstanceUuidOk() (*string, bool)`

GetInstanceUuidOk returns a tuple with the InstanceUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceUuid

`func (o *ComponentError) SetInstanceUuid(v string)`

SetInstanceUuid sets InstanceUuid field to given value.

### HasInstanceUuid

`func (o *ComponentError) HasInstanceUuid() bool`

HasInstanceUuid returns a boolean if a field has been set.

### GetErrorObjectList

`func (o *ComponentError) GetErrorObjectList() []ErrorMessageObject`

GetErrorObjectList returns the ErrorObjectList field if non-nil, zero value otherwise.

### GetErrorObjectListOk

`func (o *ComponentError) GetErrorObjectListOk() (*[]ErrorMessageObject, bool)`

GetErrorObjectListOk returns a tuple with the ErrorObjectList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorObjectList

`func (o *ComponentError) SetErrorObjectList(v []ErrorMessageObject)`

SetErrorObjectList sets ErrorObjectList field to given value.

### HasErrorObjectList

`func (o *ComponentError) HasErrorObjectList() bool`

HasErrorObjectList returns a boolean if a field has been set.

### GetParameterErrorList

`func (o *ComponentError) GetParameterErrorList() []ParameterError`

GetParameterErrorList returns the ParameterErrorList field if non-nil, zero value otherwise.

### GetParameterErrorListOk

`func (o *ComponentError) GetParameterErrorListOk() (*[]ParameterError, bool)`

GetParameterErrorListOk returns a tuple with the ParameterErrorList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameterErrorList

`func (o *ComponentError) SetParameterErrorList(v []ParameterError)`

SetParameterErrorList sets ParameterErrorList field to given value.

### HasParameterErrorList

`func (o *ComponentError) HasParameterErrorList() bool`

HasParameterErrorList returns a boolean if a field has been set.

### GetComponentType

`func (o *ComponentError) GetComponentType() string`

GetComponentType returns the ComponentType field if non-nil, zero value otherwise.

### GetComponentTypeOk

`func (o *ComponentError) GetComponentTypeOk() (*string, bool)`

GetComponentTypeOk returns a tuple with the ComponentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentType

`func (o *ComponentError) SetComponentType(v string)`

SetComponentType sets ComponentType field to given value.

### HasComponentType

`func (o *ComponentError) HasComponentType() bool`

HasComponentType returns a boolean if a field has been set.

### GetComponentName

`func (o *ComponentError) GetComponentName() string`

GetComponentName returns the ComponentName field if non-nil, zero value otherwise.

### GetComponentNameOk

`func (o *ComponentError) GetComponentNameOk() (*string, bool)`

GetComponentNameOk returns a tuple with the ComponentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentName

`func (o *ComponentError) SetComponentName(v string)`

SetComponentName sets ComponentName field to given value.

### HasComponentName

`func (o *ComponentError) HasComponentName() bool`

HasComponentName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# ServiceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TaskUuid** | Pointer to **string** | UUID of the task created for handling the request.  | [optional] 

## Methods

### NewServiceResponse

`func NewServiceResponse() *ServiceResponse`

NewServiceResponse instantiates a new ServiceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceResponseWithDefaults

`func NewServiceResponseWithDefaults() *ServiceResponse`

NewServiceResponseWithDefaults instantiates a new ServiceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaskUuid

`func (o *ServiceResponse) GetTaskUuid() string`

GetTaskUuid returns the TaskUuid field if non-nil, zero value otherwise.

### GetTaskUuidOk

`func (o *ServiceResponse) GetTaskUuidOk() (*string, bool)`

GetTaskUuidOk returns a tuple with the TaskUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskUuid

`func (o *ServiceResponse) SetTaskUuid(v string)`

SetTaskUuid sets TaskUuid field to given value.

### HasTaskUuid

`func (o *ServiceResponse) HasTaskUuid() bool`

HasTaskUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



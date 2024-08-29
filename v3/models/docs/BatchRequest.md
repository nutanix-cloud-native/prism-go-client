# BatchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExecutionOrder** | Pointer to **string** | The order of execution of the APIs in the batch. Can be either Sequential (default value) or Parallel.  | [optional] [default to "SEQUENTIAL"]
**ActionOnFailure** | Pointer to **string** | If the specified parameter is CONTINUE, the remaining APIs in the batch continue to be executed.  | [optional] [default to "CONTINUE"]
**ApiRequestList** | [**[]ApiRequest**](ApiRequest.md) | A list of API requests in the batch. | 
**ApiVersion** | Pointer to **string** | API Version of the Nutanix v3 API framework. | [optional] [readonly] [default to "3.1.0"]

## Methods

### NewBatchRequest

`func NewBatchRequest(apiRequestList []ApiRequest, ) *BatchRequest`

NewBatchRequest instantiates a new BatchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchRequestWithDefaults

`func NewBatchRequestWithDefaults() *BatchRequest`

NewBatchRequestWithDefaults instantiates a new BatchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExecutionOrder

`func (o *BatchRequest) GetExecutionOrder() string`

GetExecutionOrder returns the ExecutionOrder field if non-nil, zero value otherwise.

### GetExecutionOrderOk

`func (o *BatchRequest) GetExecutionOrderOk() (*string, bool)`

GetExecutionOrderOk returns a tuple with the ExecutionOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionOrder

`func (o *BatchRequest) SetExecutionOrder(v string)`

SetExecutionOrder sets ExecutionOrder field to given value.

### HasExecutionOrder

`func (o *BatchRequest) HasExecutionOrder() bool`

HasExecutionOrder returns a boolean if a field has been set.

### GetActionOnFailure

`func (o *BatchRequest) GetActionOnFailure() string`

GetActionOnFailure returns the ActionOnFailure field if non-nil, zero value otherwise.

### GetActionOnFailureOk

`func (o *BatchRequest) GetActionOnFailureOk() (*string, bool)`

GetActionOnFailureOk returns a tuple with the ActionOnFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionOnFailure

`func (o *BatchRequest) SetActionOnFailure(v string)`

SetActionOnFailure sets ActionOnFailure field to given value.

### HasActionOnFailure

`func (o *BatchRequest) HasActionOnFailure() bool`

HasActionOnFailure returns a boolean if a field has been set.

### GetApiRequestList

`func (o *BatchRequest) GetApiRequestList() []ApiRequest`

GetApiRequestList returns the ApiRequestList field if non-nil, zero value otherwise.

### GetApiRequestListOk

`func (o *BatchRequest) GetApiRequestListOk() (*[]ApiRequest, bool)`

GetApiRequestListOk returns a tuple with the ApiRequestList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiRequestList

`func (o *BatchRequest) SetApiRequestList(v []ApiRequest)`

SetApiRequestList sets ApiRequestList field to given value.


### GetApiVersion

`func (o *BatchRequest) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *BatchRequest) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *BatchRequest) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *BatchRequest) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



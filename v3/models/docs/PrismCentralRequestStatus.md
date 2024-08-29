# PrismCentralRequestStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TaskUuid** | **string** | Task UUID. | 
**Resources** | [**PrismCentralNodes**](PrismCentralNodes.md) |  | 

## Methods

### NewPrismCentralRequestStatus

`func NewPrismCentralRequestStatus(taskUuid string, resources PrismCentralNodes, ) *PrismCentralRequestStatus`

NewPrismCentralRequestStatus instantiates a new PrismCentralRequestStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrismCentralRequestStatusWithDefaults

`func NewPrismCentralRequestStatusWithDefaults() *PrismCentralRequestStatus`

NewPrismCentralRequestStatusWithDefaults instantiates a new PrismCentralRequestStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaskUuid

`func (o *PrismCentralRequestStatus) GetTaskUuid() string`

GetTaskUuid returns the TaskUuid field if non-nil, zero value otherwise.

### GetTaskUuidOk

`func (o *PrismCentralRequestStatus) GetTaskUuidOk() (*string, bool)`

GetTaskUuidOk returns a tuple with the TaskUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskUuid

`func (o *PrismCentralRequestStatus) SetTaskUuid(v string)`

SetTaskUuid sets TaskUuid field to given value.


### GetResources

`func (o *PrismCentralRequestStatus) GetResources() PrismCentralNodes`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *PrismCentralRequestStatus) GetResourcesOk() (*PrismCentralNodes, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *PrismCentralRequestStatus) SetResources(v PrismCentralNodes)`

SetResources sets Resources field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



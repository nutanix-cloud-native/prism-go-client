# DpdConfigStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operation** | Pointer to **string** | Operation to be performed on detecting a dead peer. | [optional] 
**IntervalSecs** | Pointer to **int32** | The amount of time the peer waits for traffic before sending a DPD request. | [optional] 
**TimeoutSecs** | Pointer to **int32** | The maximum amount of time to wait for a DPD response before marking the peer as dead.  | [optional] 

## Methods

### NewDpdConfigStatus

`func NewDpdConfigStatus() *DpdConfigStatus`

NewDpdConfigStatus instantiates a new DpdConfigStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDpdConfigStatusWithDefaults

`func NewDpdConfigStatusWithDefaults() *DpdConfigStatus`

NewDpdConfigStatusWithDefaults instantiates a new DpdConfigStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperation

`func (o *DpdConfigStatus) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *DpdConfigStatus) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *DpdConfigStatus) SetOperation(v string)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *DpdConfigStatus) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### GetIntervalSecs

`func (o *DpdConfigStatus) GetIntervalSecs() int32`

GetIntervalSecs returns the IntervalSecs field if non-nil, zero value otherwise.

### GetIntervalSecsOk

`func (o *DpdConfigStatus) GetIntervalSecsOk() (*int32, bool)`

GetIntervalSecsOk returns a tuple with the IntervalSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervalSecs

`func (o *DpdConfigStatus) SetIntervalSecs(v int32)`

SetIntervalSecs sets IntervalSecs field to given value.

### HasIntervalSecs

`func (o *DpdConfigStatus) HasIntervalSecs() bool`

HasIntervalSecs returns a boolean if a field has been set.

### GetTimeoutSecs

`func (o *DpdConfigStatus) GetTimeoutSecs() int32`

GetTimeoutSecs returns the TimeoutSecs field if non-nil, zero value otherwise.

### GetTimeoutSecsOk

`func (o *DpdConfigStatus) GetTimeoutSecsOk() (*int32, bool)`

GetTimeoutSecsOk returns a tuple with the TimeoutSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutSecs

`func (o *DpdConfigStatus) SetTimeoutSecs(v int32)`

SetTimeoutSecs sets TimeoutSecs field to given value.

### HasTimeoutSecs

`func (o *DpdConfigStatus) HasTimeoutSecs() bool`

HasTimeoutSecs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



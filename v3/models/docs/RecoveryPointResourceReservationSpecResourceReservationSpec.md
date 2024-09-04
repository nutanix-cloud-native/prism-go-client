# RecoveryPointResourceReservationSpecResourceReservationSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OperationType** | **string** | Operation to be performed on the resource reservation made by Lazan for the recovery point.  | 
**ReservationLeaseTimeoutSecs** | Pointer to **int64** | Amount of time in seconds for which the reservation lease made by Lazan need to be refreshed.  | [optional] 

## Methods

### NewRecoveryPointResourceReservationSpecResourceReservationSpec

`func NewRecoveryPointResourceReservationSpecResourceReservationSpec(operationType string, ) *RecoveryPointResourceReservationSpecResourceReservationSpec`

NewRecoveryPointResourceReservationSpecResourceReservationSpec instantiates a new RecoveryPointResourceReservationSpecResourceReservationSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoveryPointResourceReservationSpecResourceReservationSpecWithDefaults

`func NewRecoveryPointResourceReservationSpecResourceReservationSpecWithDefaults() *RecoveryPointResourceReservationSpecResourceReservationSpec`

NewRecoveryPointResourceReservationSpecResourceReservationSpecWithDefaults instantiates a new RecoveryPointResourceReservationSpecResourceReservationSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperationType

`func (o *RecoveryPointResourceReservationSpecResourceReservationSpec) GetOperationType() string`

GetOperationType returns the OperationType field if non-nil, zero value otherwise.

### GetOperationTypeOk

`func (o *RecoveryPointResourceReservationSpecResourceReservationSpec) GetOperationTypeOk() (*string, bool)`

GetOperationTypeOk returns a tuple with the OperationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationType

`func (o *RecoveryPointResourceReservationSpecResourceReservationSpec) SetOperationType(v string)`

SetOperationType sets OperationType field to given value.


### GetReservationLeaseTimeoutSecs

`func (o *RecoveryPointResourceReservationSpecResourceReservationSpec) GetReservationLeaseTimeoutSecs() int64`

GetReservationLeaseTimeoutSecs returns the ReservationLeaseTimeoutSecs field if non-nil, zero value otherwise.

### GetReservationLeaseTimeoutSecsOk

`func (o *RecoveryPointResourceReservationSpecResourceReservationSpec) GetReservationLeaseTimeoutSecsOk() (*int64, bool)`

GetReservationLeaseTimeoutSecsOk returns a tuple with the ReservationLeaseTimeoutSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservationLeaseTimeoutSecs

`func (o *RecoveryPointResourceReservationSpecResourceReservationSpec) SetReservationLeaseTimeoutSecs(v int64)`

SetReservationLeaseTimeoutSecs sets ReservationLeaseTimeoutSecs field to given value.

### HasReservationLeaseTimeoutSecs

`func (o *RecoveryPointResourceReservationSpecResourceReservationSpec) HasReservationLeaseTimeoutSecs() bool`

HasReservationLeaseTimeoutSecs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



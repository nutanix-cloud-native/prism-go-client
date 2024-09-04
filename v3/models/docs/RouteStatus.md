# RouteStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Priority** | **int32** | The preference value assigned to this route. A higher value means greater preference.  | 
**Nexthop** | [**RouteNexthopStatus**](RouteNexthopStatus.md) |  | 
**Destination** | **string** | The destination subnet of this route. | 
**IsActive** | **bool** | Whether this route is currently active. | 

## Methods

### NewRouteStatus

`func NewRouteStatus(priority int32, nexthop RouteNexthopStatus, destination string, isActive bool, ) *RouteStatus`

NewRouteStatus instantiates a new RouteStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouteStatusWithDefaults

`func NewRouteStatusWithDefaults() *RouteStatus`

NewRouteStatusWithDefaults instantiates a new RouteStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPriority

`func (o *RouteStatus) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *RouteStatus) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *RouteStatus) SetPriority(v int32)`

SetPriority sets Priority field to given value.


### GetNexthop

`func (o *RouteStatus) GetNexthop() RouteNexthopStatus`

GetNexthop returns the Nexthop field if non-nil, zero value otherwise.

### GetNexthopOk

`func (o *RouteStatus) GetNexthopOk() (*RouteNexthopStatus, bool)`

GetNexthopOk returns a tuple with the Nexthop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNexthop

`func (o *RouteStatus) SetNexthop(v RouteNexthopStatus)`

SetNexthop sets Nexthop field to given value.


### GetDestination

`func (o *RouteStatus) GetDestination() string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *RouteStatus) GetDestinationOk() (*string, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *RouteStatus) SetDestination(v string)`

SetDestination sets Destination field to given value.


### GetIsActive

`func (o *RouteStatus) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *RouteStatus) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *RouteStatus) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



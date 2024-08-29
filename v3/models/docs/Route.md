# Route

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nexthop** | [**RouteNexthopReference**](RouteNexthopReference.md) |  | 
**Destination** | **string** |  | 

## Methods

### NewRoute

`func NewRoute(nexthop RouteNexthopReference, destination string, ) *Route`

NewRoute instantiates a new Route object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouteWithDefaults

`func NewRouteWithDefaults() *Route`

NewRouteWithDefaults instantiates a new Route object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNexthop

`func (o *Route) GetNexthop() RouteNexthopReference`

GetNexthop returns the Nexthop field if non-nil, zero value otherwise.

### GetNexthopOk

`func (o *Route) GetNexthopOk() (*RouteNexthopReference, bool)`

GetNexthopOk returns a tuple with the Nexthop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNexthop

`func (o *Route) SetNexthop(v RouteNexthopReference)`

SetNexthop sets Nexthop field to given value.


### GetDestination

`func (o *Route) GetDestination() string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *Route) GetDestinationOk() (*string, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *Route) SetDestination(v string)`

SetDestination sets Destination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



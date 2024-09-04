# VpcRouteTableResources

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StaticRoutesList** | Pointer to [**[]Route**](Route.md) |  | [optional] 
**DefaultRouteNexthop** | Pointer to [**RouteNexthopReference**](RouteNexthopReference.md) |  | [optional] 

## Methods

### NewVpcRouteTableResources

`func NewVpcRouteTableResources() *VpcRouteTableResources`

NewVpcRouteTableResources instantiates a new VpcRouteTableResources object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpcRouteTableResourcesWithDefaults

`func NewVpcRouteTableResourcesWithDefaults() *VpcRouteTableResources`

NewVpcRouteTableResourcesWithDefaults instantiates a new VpcRouteTableResources object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStaticRoutesList

`func (o *VpcRouteTableResources) GetStaticRoutesList() []Route`

GetStaticRoutesList returns the StaticRoutesList field if non-nil, zero value otherwise.

### GetStaticRoutesListOk

`func (o *VpcRouteTableResources) GetStaticRoutesListOk() (*[]Route, bool)`

GetStaticRoutesListOk returns a tuple with the StaticRoutesList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticRoutesList

`func (o *VpcRouteTableResources) SetStaticRoutesList(v []Route)`

SetStaticRoutesList sets StaticRoutesList field to given value.

### HasStaticRoutesList

`func (o *VpcRouteTableResources) HasStaticRoutesList() bool`

HasStaticRoutesList returns a boolean if a field has been set.

### GetDefaultRouteNexthop

`func (o *VpcRouteTableResources) GetDefaultRouteNexthop() RouteNexthopReference`

GetDefaultRouteNexthop returns the DefaultRouteNexthop field if non-nil, zero value otherwise.

### GetDefaultRouteNexthopOk

`func (o *VpcRouteTableResources) GetDefaultRouteNexthopOk() (*RouteNexthopReference, bool)`

GetDefaultRouteNexthopOk returns a tuple with the DefaultRouteNexthop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRouteNexthop

`func (o *VpcRouteTableResources) SetDefaultRouteNexthop(v RouteNexthopReference)`

SetDefaultRouteNexthop sets DefaultRouteNexthop field to given value.

### HasDefaultRouteNexthop

`func (o *VpcRouteTableResources) HasDefaultRouteNexthop() bool`

HasDefaultRouteNexthop returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



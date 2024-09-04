# VpcRouteTableResourcesDefStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StaticRoutesList** | Pointer to [**[]RouteStatus**](RouteStatus.md) | Set of statically configured routes in this table. | [optional] 
**DynamicRoutesList** | Pointer to [**[]RouteStatus**](RouteStatus.md) | Set of dynamically received routes in this table. | [optional] 
**DefaultRoute** | Pointer to [**RouteStatus**](RouteStatus.md) |  | [optional] 
**LocalRoutesList** | Pointer to [**[]RouteStatus**](RouteStatus.md) | Set of locally defined routes in this table. | [optional] 

## Methods

### NewVpcRouteTableResourcesDefStatus

`func NewVpcRouteTableResourcesDefStatus() *VpcRouteTableResourcesDefStatus`

NewVpcRouteTableResourcesDefStatus instantiates a new VpcRouteTableResourcesDefStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpcRouteTableResourcesDefStatusWithDefaults

`func NewVpcRouteTableResourcesDefStatusWithDefaults() *VpcRouteTableResourcesDefStatus`

NewVpcRouteTableResourcesDefStatusWithDefaults instantiates a new VpcRouteTableResourcesDefStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStaticRoutesList

`func (o *VpcRouteTableResourcesDefStatus) GetStaticRoutesList() []RouteStatus`

GetStaticRoutesList returns the StaticRoutesList field if non-nil, zero value otherwise.

### GetStaticRoutesListOk

`func (o *VpcRouteTableResourcesDefStatus) GetStaticRoutesListOk() (*[]RouteStatus, bool)`

GetStaticRoutesListOk returns a tuple with the StaticRoutesList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticRoutesList

`func (o *VpcRouteTableResourcesDefStatus) SetStaticRoutesList(v []RouteStatus)`

SetStaticRoutesList sets StaticRoutesList field to given value.

### HasStaticRoutesList

`func (o *VpcRouteTableResourcesDefStatus) HasStaticRoutesList() bool`

HasStaticRoutesList returns a boolean if a field has been set.

### GetDynamicRoutesList

`func (o *VpcRouteTableResourcesDefStatus) GetDynamicRoutesList() []RouteStatus`

GetDynamicRoutesList returns the DynamicRoutesList field if non-nil, zero value otherwise.

### GetDynamicRoutesListOk

`func (o *VpcRouteTableResourcesDefStatus) GetDynamicRoutesListOk() (*[]RouteStatus, bool)`

GetDynamicRoutesListOk returns a tuple with the DynamicRoutesList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicRoutesList

`func (o *VpcRouteTableResourcesDefStatus) SetDynamicRoutesList(v []RouteStatus)`

SetDynamicRoutesList sets DynamicRoutesList field to given value.

### HasDynamicRoutesList

`func (o *VpcRouteTableResourcesDefStatus) HasDynamicRoutesList() bool`

HasDynamicRoutesList returns a boolean if a field has been set.

### GetDefaultRoute

`func (o *VpcRouteTableResourcesDefStatus) GetDefaultRoute() RouteStatus`

GetDefaultRoute returns the DefaultRoute field if non-nil, zero value otherwise.

### GetDefaultRouteOk

`func (o *VpcRouteTableResourcesDefStatus) GetDefaultRouteOk() (*RouteStatus, bool)`

GetDefaultRouteOk returns a tuple with the DefaultRoute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRoute

`func (o *VpcRouteTableResourcesDefStatus) SetDefaultRoute(v RouteStatus)`

SetDefaultRoute sets DefaultRoute field to given value.

### HasDefaultRoute

`func (o *VpcRouteTableResourcesDefStatus) HasDefaultRoute() bool`

HasDefaultRoute returns a boolean if a field has been set.

### GetLocalRoutesList

`func (o *VpcRouteTableResourcesDefStatus) GetLocalRoutesList() []RouteStatus`

GetLocalRoutesList returns the LocalRoutesList field if non-nil, zero value otherwise.

### GetLocalRoutesListOk

`func (o *VpcRouteTableResourcesDefStatus) GetLocalRoutesListOk() (*[]RouteStatus, bool)`

GetLocalRoutesListOk returns a tuple with the LocalRoutesList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalRoutesList

`func (o *VpcRouteTableResourcesDefStatus) SetLocalRoutesList(v []RouteStatus)`

SetLocalRoutesList sets LocalRoutesList field to given value.

### HasLocalRoutesList

`func (o *VpcRouteTableResourcesDefStatus) HasLocalRoutesList() bool`

HasLocalRoutesList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



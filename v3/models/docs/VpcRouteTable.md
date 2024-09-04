# VpcRouteTable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvailabilityZoneReference** | Pointer to [**AvailabilityZoneReference**](AvailabilityZoneReference.md) |  | [optional] 
**Description** | Pointer to **string** | A description for vpc_route_table. | [optional] 
**Resources** | [**VpcRouteTableResources**](VpcRouteTableResources.md) |  | 
**Name** | Pointer to **string** | vpc_route_table Name. | [optional] 

## Methods

### NewVpcRouteTable

`func NewVpcRouteTable(resources VpcRouteTableResources, ) *VpcRouteTable`

NewVpcRouteTable instantiates a new VpcRouteTable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpcRouteTableWithDefaults

`func NewVpcRouteTableWithDefaults() *VpcRouteTable`

NewVpcRouteTableWithDefaults instantiates a new VpcRouteTable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailabilityZoneReference

`func (o *VpcRouteTable) GetAvailabilityZoneReference() AvailabilityZoneReference`

GetAvailabilityZoneReference returns the AvailabilityZoneReference field if non-nil, zero value otherwise.

### GetAvailabilityZoneReferenceOk

`func (o *VpcRouteTable) GetAvailabilityZoneReferenceOk() (*AvailabilityZoneReference, bool)`

GetAvailabilityZoneReferenceOk returns a tuple with the AvailabilityZoneReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZoneReference

`func (o *VpcRouteTable) SetAvailabilityZoneReference(v AvailabilityZoneReference)`

SetAvailabilityZoneReference sets AvailabilityZoneReference field to given value.

### HasAvailabilityZoneReference

`func (o *VpcRouteTable) HasAvailabilityZoneReference() bool`

HasAvailabilityZoneReference returns a boolean if a field has been set.

### GetDescription

`func (o *VpcRouteTable) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VpcRouteTable) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VpcRouteTable) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VpcRouteTable) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetResources

`func (o *VpcRouteTable) GetResources() VpcRouteTableResources`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *VpcRouteTable) GetResourcesOk() (*VpcRouteTableResources, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *VpcRouteTable) SetResources(v VpcRouteTableResources)`

SetResources sets Resources field to given value.


### GetName

`func (o *VpcRouteTable) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VpcRouteTable) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VpcRouteTable) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VpcRouteTable) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



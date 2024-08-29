# VpcRouteTableDefStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | vpc_route_table Name. | 
**State** | Pointer to **string** | The state of the vpc_route_table. | [optional] 
**AvailabilityZoneReference** | Pointer to [**AvailabilityZoneReference**](AvailabilityZoneReference.md) |  | [optional] 
**MessageList** | Pointer to [**[]MessageResource**](MessageResource.md) | Any error messages for the vpc_route_table, if in an error state. | [optional] 
**ClusterReference** | Pointer to [**ClusterReference**](ClusterReference.md) |  | [optional] 
**Resources** | [**VpcRouteTableResourcesDefStatus**](VpcRouteTableResourcesDefStatus.md) |  | 
**Description** | Pointer to **string** | A description for vpc_route_table. | [optional] 

## Methods

### NewVpcRouteTableDefStatus

`func NewVpcRouteTableDefStatus(name string, resources VpcRouteTableResourcesDefStatus, ) *VpcRouteTableDefStatus`

NewVpcRouteTableDefStatus instantiates a new VpcRouteTableDefStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpcRouteTableDefStatusWithDefaults

`func NewVpcRouteTableDefStatusWithDefaults() *VpcRouteTableDefStatus`

NewVpcRouteTableDefStatusWithDefaults instantiates a new VpcRouteTableDefStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *VpcRouteTableDefStatus) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VpcRouteTableDefStatus) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VpcRouteTableDefStatus) SetName(v string)`

SetName sets Name field to given value.


### GetState

`func (o *VpcRouteTableDefStatus) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *VpcRouteTableDefStatus) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *VpcRouteTableDefStatus) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *VpcRouteTableDefStatus) HasState() bool`

HasState returns a boolean if a field has been set.

### GetAvailabilityZoneReference

`func (o *VpcRouteTableDefStatus) GetAvailabilityZoneReference() AvailabilityZoneReference`

GetAvailabilityZoneReference returns the AvailabilityZoneReference field if non-nil, zero value otherwise.

### GetAvailabilityZoneReferenceOk

`func (o *VpcRouteTableDefStatus) GetAvailabilityZoneReferenceOk() (*AvailabilityZoneReference, bool)`

GetAvailabilityZoneReferenceOk returns a tuple with the AvailabilityZoneReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZoneReference

`func (o *VpcRouteTableDefStatus) SetAvailabilityZoneReference(v AvailabilityZoneReference)`

SetAvailabilityZoneReference sets AvailabilityZoneReference field to given value.

### HasAvailabilityZoneReference

`func (o *VpcRouteTableDefStatus) HasAvailabilityZoneReference() bool`

HasAvailabilityZoneReference returns a boolean if a field has been set.

### GetMessageList

`func (o *VpcRouteTableDefStatus) GetMessageList() []MessageResource`

GetMessageList returns the MessageList field if non-nil, zero value otherwise.

### GetMessageListOk

`func (o *VpcRouteTableDefStatus) GetMessageListOk() (*[]MessageResource, bool)`

GetMessageListOk returns a tuple with the MessageList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageList

`func (o *VpcRouteTableDefStatus) SetMessageList(v []MessageResource)`

SetMessageList sets MessageList field to given value.

### HasMessageList

`func (o *VpcRouteTableDefStatus) HasMessageList() bool`

HasMessageList returns a boolean if a field has been set.

### GetClusterReference

`func (o *VpcRouteTableDefStatus) GetClusterReference() ClusterReference`

GetClusterReference returns the ClusterReference field if non-nil, zero value otherwise.

### GetClusterReferenceOk

`func (o *VpcRouteTableDefStatus) GetClusterReferenceOk() (*ClusterReference, bool)`

GetClusterReferenceOk returns a tuple with the ClusterReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterReference

`func (o *VpcRouteTableDefStatus) SetClusterReference(v ClusterReference)`

SetClusterReference sets ClusterReference field to given value.

### HasClusterReference

`func (o *VpcRouteTableDefStatus) HasClusterReference() bool`

HasClusterReference returns a boolean if a field has been set.

### GetResources

`func (o *VpcRouteTableDefStatus) GetResources() VpcRouteTableResourcesDefStatus`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *VpcRouteTableDefStatus) GetResourcesOk() (*VpcRouteTableResourcesDefStatus, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *VpcRouteTableDefStatus) SetResources(v VpcRouteTableResourcesDefStatus)`

SetResources sets Resources field to given value.


### GetDescription

`func (o *VpcRouteTableDefStatus) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VpcRouteTableDefStatus) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VpcRouteTableDefStatus) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VpcRouteTableDefStatus) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# Blueprint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | blueprint Name. | 
**AvailabilityZoneReference** | Pointer to [**AvailabilityZoneReference**](AvailabilityZoneReference.md) |  | [optional] 
**Description** | Pointer to **string** | A description for blueprint. | [optional] 
**Resources** | [**BlueprintResources**](BlueprintResources.md) |  | 
**ClusterReference** | Pointer to [**ClusterReference**](ClusterReference.md) |  | [optional] 

## Methods

### NewBlueprint

`func NewBlueprint(name string, resources BlueprintResources, ) *Blueprint`

NewBlueprint instantiates a new Blueprint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintWithDefaults

`func NewBlueprintWithDefaults() *Blueprint`

NewBlueprintWithDefaults instantiates a new Blueprint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Blueprint) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Blueprint) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Blueprint) SetName(v string)`

SetName sets Name field to given value.


### GetAvailabilityZoneReference

`func (o *Blueprint) GetAvailabilityZoneReference() AvailabilityZoneReference`

GetAvailabilityZoneReference returns the AvailabilityZoneReference field if non-nil, zero value otherwise.

### GetAvailabilityZoneReferenceOk

`func (o *Blueprint) GetAvailabilityZoneReferenceOk() (*AvailabilityZoneReference, bool)`

GetAvailabilityZoneReferenceOk returns a tuple with the AvailabilityZoneReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZoneReference

`func (o *Blueprint) SetAvailabilityZoneReference(v AvailabilityZoneReference)`

SetAvailabilityZoneReference sets AvailabilityZoneReference field to given value.

### HasAvailabilityZoneReference

`func (o *Blueprint) HasAvailabilityZoneReference() bool`

HasAvailabilityZoneReference returns a boolean if a field has been set.

### GetDescription

`func (o *Blueprint) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Blueprint) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Blueprint) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Blueprint) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetResources

`func (o *Blueprint) GetResources() BlueprintResources`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *Blueprint) GetResourcesOk() (*BlueprintResources, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *Blueprint) SetResources(v BlueprintResources)`

SetResources sets Resources field to given value.


### GetClusterReference

`func (o *Blueprint) GetClusterReference() ClusterReference`

GetClusterReference returns the ClusterReference field if non-nil, zero value otherwise.

### GetClusterReferenceOk

`func (o *Blueprint) GetClusterReferenceOk() (*ClusterReference, bool)`

GetClusterReferenceOk returns a tuple with the ClusterReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterReference

`func (o *Blueprint) SetClusterReference(v ClusterReference)`

SetClusterReference sets ClusterReference field to given value.

### HasClusterReference

`func (o *Blueprint) HasClusterReference() bool`

HasClusterReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



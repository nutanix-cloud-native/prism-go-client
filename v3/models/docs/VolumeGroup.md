# VolumeGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Volume group name. | 
**Description** | Pointer to **string** | Volume group description. | [optional] 
**Resources** | [**VolumeGroupResourcesInput**](VolumeGroupResourcesInput.md) |  | 
**ClusterReference** | Pointer to [**ClusterReference**](ClusterReference.md) |  | [optional] 

## Methods

### NewVolumeGroup

`func NewVolumeGroup(name string, resources VolumeGroupResourcesInput, ) *VolumeGroup`

NewVolumeGroup instantiates a new VolumeGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeGroupWithDefaults

`func NewVolumeGroupWithDefaults() *VolumeGroup`

NewVolumeGroupWithDefaults instantiates a new VolumeGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *VolumeGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VolumeGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VolumeGroup) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *VolumeGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VolumeGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VolumeGroup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VolumeGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetResources

`func (o *VolumeGroup) GetResources() VolumeGroupResourcesInput`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *VolumeGroup) GetResourcesOk() (*VolumeGroupResourcesInput, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *VolumeGroup) SetResources(v VolumeGroupResourcesInput)`

SetResources sets Resources field to given value.


### GetClusterReference

`func (o *VolumeGroup) GetClusterReference() ClusterReference`

GetClusterReference returns the ClusterReference field if non-nil, zero value otherwise.

### GetClusterReferenceOk

`func (o *VolumeGroup) GetClusterReferenceOk() (*ClusterReference, bool)`

GetClusterReferenceOk returns a tuple with the ClusterReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterReference

`func (o *VolumeGroup) SetClusterReference(v ClusterReference)`

SetClusterReference sets ClusterReference field to given value.

### HasClusterReference

`func (o *VolumeGroup) HasClusterReference() bool`

HasClusterReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



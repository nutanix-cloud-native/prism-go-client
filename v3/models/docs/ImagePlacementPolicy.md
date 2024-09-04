# ImagePlacementPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Image placement policy name. | 
**Resources** | [**ImagePlacementPolicyResources**](ImagePlacementPolicyResources.md) |  | 
**Description** | Pointer to **string** | A description for Image placement policy. | [optional] 

## Methods

### NewImagePlacementPolicy

`func NewImagePlacementPolicy(name string, resources ImagePlacementPolicyResources, ) *ImagePlacementPolicy`

NewImagePlacementPolicy instantiates a new ImagePlacementPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImagePlacementPolicyWithDefaults

`func NewImagePlacementPolicyWithDefaults() *ImagePlacementPolicy`

NewImagePlacementPolicyWithDefaults instantiates a new ImagePlacementPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ImagePlacementPolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ImagePlacementPolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ImagePlacementPolicy) SetName(v string)`

SetName sets Name field to given value.


### GetResources

`func (o *ImagePlacementPolicy) GetResources() ImagePlacementPolicyResources`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *ImagePlacementPolicy) GetResourcesOk() (*ImagePlacementPolicyResources, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *ImagePlacementPolicy) SetResources(v ImagePlacementPolicyResources)`

SetResources sets Resources field to given value.


### GetDescription

`func (o *ImagePlacementPolicy) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ImagePlacementPolicy) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ImagePlacementPolicy) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ImagePlacementPolicy) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



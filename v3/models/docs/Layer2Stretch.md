# Layer2Stretch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvailabilityZoneReference** | Pointer to [**AvailabilityZoneReference**](AvailabilityZoneReference.md) |  | [optional] 
**Description** | Pointer to **string** | A description for layer2_stretch. | [optional] 
**Resources** | [**Layer2StretchResources**](Layer2StretchResources.md) |  | 
**Name** | **string** | layer2_stretch Name. | 

## Methods

### NewLayer2Stretch

`func NewLayer2Stretch(resources Layer2StretchResources, name string, ) *Layer2Stretch`

NewLayer2Stretch instantiates a new Layer2Stretch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLayer2StretchWithDefaults

`func NewLayer2StretchWithDefaults() *Layer2Stretch`

NewLayer2StretchWithDefaults instantiates a new Layer2Stretch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailabilityZoneReference

`func (o *Layer2Stretch) GetAvailabilityZoneReference() AvailabilityZoneReference`

GetAvailabilityZoneReference returns the AvailabilityZoneReference field if non-nil, zero value otherwise.

### GetAvailabilityZoneReferenceOk

`func (o *Layer2Stretch) GetAvailabilityZoneReferenceOk() (*AvailabilityZoneReference, bool)`

GetAvailabilityZoneReferenceOk returns a tuple with the AvailabilityZoneReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZoneReference

`func (o *Layer2Stretch) SetAvailabilityZoneReference(v AvailabilityZoneReference)`

SetAvailabilityZoneReference sets AvailabilityZoneReference field to given value.

### HasAvailabilityZoneReference

`func (o *Layer2Stretch) HasAvailabilityZoneReference() bool`

HasAvailabilityZoneReference returns a boolean if a field has been set.

### GetDescription

`func (o *Layer2Stretch) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Layer2Stretch) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Layer2Stretch) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Layer2Stretch) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetResources

`func (o *Layer2Stretch) GetResources() Layer2StretchResources`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *Layer2Stretch) GetResourcesOk() (*Layer2StretchResources, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *Layer2Stretch) SetResources(v Layer2StretchResources)`

SetResources sets Resources field to given value.


### GetName

`func (o *Layer2Stretch) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Layer2Stretch) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Layer2Stretch) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



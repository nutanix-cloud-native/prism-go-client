# AvailabilityZoneDefStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Availability Zone Name | 
**Resources** | [**AvailabilityZoneResources**](AvailabilityZoneResources.md) |  | 

## Methods

### NewAvailabilityZoneDefStatus

`func NewAvailabilityZoneDefStatus(name string, resources AvailabilityZoneResources, ) *AvailabilityZoneDefStatus`

NewAvailabilityZoneDefStatus instantiates a new AvailabilityZoneDefStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAvailabilityZoneDefStatusWithDefaults

`func NewAvailabilityZoneDefStatusWithDefaults() *AvailabilityZoneDefStatus`

NewAvailabilityZoneDefStatusWithDefaults instantiates a new AvailabilityZoneDefStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AvailabilityZoneDefStatus) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AvailabilityZoneDefStatus) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AvailabilityZoneDefStatus) SetName(v string)`

SetName sets Name field to given value.


### GetResources

`func (o *AvailabilityZoneDefStatus) GetResources() AvailabilityZoneResources`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *AvailabilityZoneDefStatus) GetResourcesOk() (*AvailabilityZoneResources, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *AvailabilityZoneDefStatus) SetResources(v AvailabilityZoneResources)`

SetResources sets Resources field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



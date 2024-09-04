# FloatingIp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvailabilityZoneReference** | Pointer to [**AvailabilityZoneReference**](AvailabilityZoneReference.md) |  | [optional] 
**Description** | Pointer to **string** | A description for floating_ip. | [optional] 
**Resources** | [**FloatingIpResources**](FloatingIpResources.md) |  | 
**Name** | Pointer to **string** | floating_ip Name. | [optional] 

## Methods

### NewFloatingIp

`func NewFloatingIp(resources FloatingIpResources, ) *FloatingIp`

NewFloatingIp instantiates a new FloatingIp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFloatingIpWithDefaults

`func NewFloatingIpWithDefaults() *FloatingIp`

NewFloatingIpWithDefaults instantiates a new FloatingIp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailabilityZoneReference

`func (o *FloatingIp) GetAvailabilityZoneReference() AvailabilityZoneReference`

GetAvailabilityZoneReference returns the AvailabilityZoneReference field if non-nil, zero value otherwise.

### GetAvailabilityZoneReferenceOk

`func (o *FloatingIp) GetAvailabilityZoneReferenceOk() (*AvailabilityZoneReference, bool)`

GetAvailabilityZoneReferenceOk returns a tuple with the AvailabilityZoneReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZoneReference

`func (o *FloatingIp) SetAvailabilityZoneReference(v AvailabilityZoneReference)`

SetAvailabilityZoneReference sets AvailabilityZoneReference field to given value.

### HasAvailabilityZoneReference

`func (o *FloatingIp) HasAvailabilityZoneReference() bool`

HasAvailabilityZoneReference returns a boolean if a field has been set.

### GetDescription

`func (o *FloatingIp) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FloatingIp) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FloatingIp) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FloatingIp) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetResources

`func (o *FloatingIp) GetResources() FloatingIpResources`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *FloatingIp) GetResourcesOk() (*FloatingIpResources, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *FloatingIp) SetResources(v FloatingIpResources)`

SetResources sets Resources field to given value.


### GetName

`func (o *FloatingIp) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FloatingIp) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FloatingIp) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FloatingIp) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



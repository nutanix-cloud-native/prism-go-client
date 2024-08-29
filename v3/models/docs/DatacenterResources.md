# DatacenterResources

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvailabilityZoneReference** | [**AvailabilityZoneReference**](AvailabilityZoneReference.md) |  | 
**Address** | Pointer to [**PostalAddress**](PostalAddress.md) |  | [optional] 

## Methods

### NewDatacenterResources

`func NewDatacenterResources(availabilityZoneReference AvailabilityZoneReference, ) *DatacenterResources`

NewDatacenterResources instantiates a new DatacenterResources object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatacenterResourcesWithDefaults

`func NewDatacenterResourcesWithDefaults() *DatacenterResources`

NewDatacenterResourcesWithDefaults instantiates a new DatacenterResources object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailabilityZoneReference

`func (o *DatacenterResources) GetAvailabilityZoneReference() AvailabilityZoneReference`

GetAvailabilityZoneReference returns the AvailabilityZoneReference field if non-nil, zero value otherwise.

### GetAvailabilityZoneReferenceOk

`func (o *DatacenterResources) GetAvailabilityZoneReferenceOk() (*AvailabilityZoneReference, bool)`

GetAvailabilityZoneReferenceOk returns a tuple with the AvailabilityZoneReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZoneReference

`func (o *DatacenterResources) SetAvailabilityZoneReference(v AvailabilityZoneReference)`

SetAvailabilityZoneReference sets AvailabilityZoneReference field to given value.


### GetAddress

`func (o *DatacenterResources) GetAddress() PostalAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *DatacenterResources) GetAddressOk() (*PostalAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *DatacenterResources) SetAddress(v PostalAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *DatacenterResources) HasAddress() bool`

HasAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



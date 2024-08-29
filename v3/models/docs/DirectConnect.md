# DirectConnect

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvailabilityZoneReference** | Pointer to [**AvailabilityZoneReference**](AvailabilityZoneReference.md) |  | [optional] 
**Description** | Pointer to **string** | A description for direct_connect. | [optional] 
**Resources** | [**DirectConnectResources**](DirectConnectResources.md) |  | 
**Name** | **string** | direct_connect Name. | 

## Methods

### NewDirectConnect

`func NewDirectConnect(resources DirectConnectResources, name string, ) *DirectConnect`

NewDirectConnect instantiates a new DirectConnect object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDirectConnectWithDefaults

`func NewDirectConnectWithDefaults() *DirectConnect`

NewDirectConnectWithDefaults instantiates a new DirectConnect object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailabilityZoneReference

`func (o *DirectConnect) GetAvailabilityZoneReference() AvailabilityZoneReference`

GetAvailabilityZoneReference returns the AvailabilityZoneReference field if non-nil, zero value otherwise.

### GetAvailabilityZoneReferenceOk

`func (o *DirectConnect) GetAvailabilityZoneReferenceOk() (*AvailabilityZoneReference, bool)`

GetAvailabilityZoneReferenceOk returns a tuple with the AvailabilityZoneReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZoneReference

`func (o *DirectConnect) SetAvailabilityZoneReference(v AvailabilityZoneReference)`

SetAvailabilityZoneReference sets AvailabilityZoneReference field to given value.

### HasAvailabilityZoneReference

`func (o *DirectConnect) HasAvailabilityZoneReference() bool`

HasAvailabilityZoneReference returns a boolean if a field has been set.

### GetDescription

`func (o *DirectConnect) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DirectConnect) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DirectConnect) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DirectConnect) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetResources

`func (o *DirectConnect) GetResources() DirectConnectResources`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *DirectConnect) GetResourcesOk() (*DirectConnectResources, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *DirectConnect) SetResources(v DirectConnectResources)`

SetResources sets Resources field to given value.


### GetName

`func (o *DirectConnect) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DirectConnect) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DirectConnect) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# RecoveryPlanResourcesParametersNetworkMappingListInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AreNetworksStretched** | Pointer to **bool** | Whether the networks across the Availability Zones in above mapping are stretched.  | [optional] 
**AvailabilityZoneNetworkMappingList** | Pointer to [**[]RecoveryPlanResourcesParametersNetworkMappingListInnerAvailabilityZoneNetworkMappingListInner**](RecoveryPlanResourcesParametersNetworkMappingListInnerAvailabilityZoneNetworkMappingListInner.md) | Mapping of networks across the Availability Zones.  | [optional] 

## Methods

### NewRecoveryPlanResourcesParametersNetworkMappingListInner

`func NewRecoveryPlanResourcesParametersNetworkMappingListInner() *RecoveryPlanResourcesParametersNetworkMappingListInner`

NewRecoveryPlanResourcesParametersNetworkMappingListInner instantiates a new RecoveryPlanResourcesParametersNetworkMappingListInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoveryPlanResourcesParametersNetworkMappingListInnerWithDefaults

`func NewRecoveryPlanResourcesParametersNetworkMappingListInnerWithDefaults() *RecoveryPlanResourcesParametersNetworkMappingListInner`

NewRecoveryPlanResourcesParametersNetworkMappingListInnerWithDefaults instantiates a new RecoveryPlanResourcesParametersNetworkMappingListInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAreNetworksStretched

`func (o *RecoveryPlanResourcesParametersNetworkMappingListInner) GetAreNetworksStretched() bool`

GetAreNetworksStretched returns the AreNetworksStretched field if non-nil, zero value otherwise.

### GetAreNetworksStretchedOk

`func (o *RecoveryPlanResourcesParametersNetworkMappingListInner) GetAreNetworksStretchedOk() (*bool, bool)`

GetAreNetworksStretchedOk returns a tuple with the AreNetworksStretched field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAreNetworksStretched

`func (o *RecoveryPlanResourcesParametersNetworkMappingListInner) SetAreNetworksStretched(v bool)`

SetAreNetworksStretched sets AreNetworksStretched field to given value.

### HasAreNetworksStretched

`func (o *RecoveryPlanResourcesParametersNetworkMappingListInner) HasAreNetworksStretched() bool`

HasAreNetworksStretched returns a boolean if a field has been set.

### GetAvailabilityZoneNetworkMappingList

`func (o *RecoveryPlanResourcesParametersNetworkMappingListInner) GetAvailabilityZoneNetworkMappingList() []RecoveryPlanResourcesParametersNetworkMappingListInnerAvailabilityZoneNetworkMappingListInner`

GetAvailabilityZoneNetworkMappingList returns the AvailabilityZoneNetworkMappingList field if non-nil, zero value otherwise.

### GetAvailabilityZoneNetworkMappingListOk

`func (o *RecoveryPlanResourcesParametersNetworkMappingListInner) GetAvailabilityZoneNetworkMappingListOk() (*[]RecoveryPlanResourcesParametersNetworkMappingListInnerAvailabilityZoneNetworkMappingListInner, bool)`

GetAvailabilityZoneNetworkMappingListOk returns a tuple with the AvailabilityZoneNetworkMappingList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZoneNetworkMappingList

`func (o *RecoveryPlanResourcesParametersNetworkMappingListInner) SetAvailabilityZoneNetworkMappingList(v []RecoveryPlanResourcesParametersNetworkMappingListInnerAvailabilityZoneNetworkMappingListInner)`

SetAvailabilityZoneNetworkMappingList sets AvailabilityZoneNetworkMappingList field to given value.

### HasAvailabilityZoneNetworkMappingList

`func (o *RecoveryPlanResourcesParametersNetworkMappingListInner) HasAvailabilityZoneNetworkMappingList() bool`

HasAvailabilityZoneNetworkMappingList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



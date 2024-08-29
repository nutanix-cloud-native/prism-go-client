# RecoveryPlanResourcesParametersFloatingIpAssignmentListInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvailabilityZoneUrl** | **string** | URL of the Availability Zone.  | 
**VmIpAssignmentList** | Pointer to [**[]RecoveryPlanResourcesParametersFloatingIpAssignmentListInnerVmIpAssignmentListInner**](RecoveryPlanResourcesParametersFloatingIpAssignmentListInnerVmIpAssignmentListInner.md) | IP assignment for VMs upon recovery in the specified Availability Zone.  | [optional] 

## Methods

### NewRecoveryPlanResourcesParametersFloatingIpAssignmentListInner

`func NewRecoveryPlanResourcesParametersFloatingIpAssignmentListInner(availabilityZoneUrl string, ) *RecoveryPlanResourcesParametersFloatingIpAssignmentListInner`

NewRecoveryPlanResourcesParametersFloatingIpAssignmentListInner instantiates a new RecoveryPlanResourcesParametersFloatingIpAssignmentListInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoveryPlanResourcesParametersFloatingIpAssignmentListInnerWithDefaults

`func NewRecoveryPlanResourcesParametersFloatingIpAssignmentListInnerWithDefaults() *RecoveryPlanResourcesParametersFloatingIpAssignmentListInner`

NewRecoveryPlanResourcesParametersFloatingIpAssignmentListInnerWithDefaults instantiates a new RecoveryPlanResourcesParametersFloatingIpAssignmentListInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailabilityZoneUrl

`func (o *RecoveryPlanResourcesParametersFloatingIpAssignmentListInner) GetAvailabilityZoneUrl() string`

GetAvailabilityZoneUrl returns the AvailabilityZoneUrl field if non-nil, zero value otherwise.

### GetAvailabilityZoneUrlOk

`func (o *RecoveryPlanResourcesParametersFloatingIpAssignmentListInner) GetAvailabilityZoneUrlOk() (*string, bool)`

GetAvailabilityZoneUrlOk returns a tuple with the AvailabilityZoneUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZoneUrl

`func (o *RecoveryPlanResourcesParametersFloatingIpAssignmentListInner) SetAvailabilityZoneUrl(v string)`

SetAvailabilityZoneUrl sets AvailabilityZoneUrl field to given value.


### GetVmIpAssignmentList

`func (o *RecoveryPlanResourcesParametersFloatingIpAssignmentListInner) GetVmIpAssignmentList() []RecoveryPlanResourcesParametersFloatingIpAssignmentListInnerVmIpAssignmentListInner`

GetVmIpAssignmentList returns the VmIpAssignmentList field if non-nil, zero value otherwise.

### GetVmIpAssignmentListOk

`func (o *RecoveryPlanResourcesParametersFloatingIpAssignmentListInner) GetVmIpAssignmentListOk() (*[]RecoveryPlanResourcesParametersFloatingIpAssignmentListInnerVmIpAssignmentListInner, bool)`

GetVmIpAssignmentListOk returns a tuple with the VmIpAssignmentList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmIpAssignmentList

`func (o *RecoveryPlanResourcesParametersFloatingIpAssignmentListInner) SetVmIpAssignmentList(v []RecoveryPlanResourcesParametersFloatingIpAssignmentListInnerVmIpAssignmentListInner)`

SetVmIpAssignmentList sets VmIpAssignmentList field to given value.

### HasVmIpAssignmentList

`func (o *RecoveryPlanResourcesParametersFloatingIpAssignmentListInner) HasVmIpAssignmentList() bool`

HasVmIpAssignmentList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# RecoveryPlanVmIpAssignment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IpConfigList** | [**[]RecoveryPlanVmIpAssignmentIpConfigListInner**](RecoveryPlanVmIpAssignmentIpConfigListInner.md) | List of IP configurations for a VM.  | 
**VmReference** | [**VmReference**](VmReference.md) |  | 

## Methods

### NewRecoveryPlanVmIpAssignment

`func NewRecoveryPlanVmIpAssignment(ipConfigList []RecoveryPlanVmIpAssignmentIpConfigListInner, vmReference VmReference, ) *RecoveryPlanVmIpAssignment`

NewRecoveryPlanVmIpAssignment instantiates a new RecoveryPlanVmIpAssignment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoveryPlanVmIpAssignmentWithDefaults

`func NewRecoveryPlanVmIpAssignmentWithDefaults() *RecoveryPlanVmIpAssignment`

NewRecoveryPlanVmIpAssignmentWithDefaults instantiates a new RecoveryPlanVmIpAssignment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpConfigList

`func (o *RecoveryPlanVmIpAssignment) GetIpConfigList() []RecoveryPlanVmIpAssignmentIpConfigListInner`

GetIpConfigList returns the IpConfigList field if non-nil, zero value otherwise.

### GetIpConfigListOk

`func (o *RecoveryPlanVmIpAssignment) GetIpConfigListOk() (*[]RecoveryPlanVmIpAssignmentIpConfigListInner, bool)`

GetIpConfigListOk returns a tuple with the IpConfigList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpConfigList

`func (o *RecoveryPlanVmIpAssignment) SetIpConfigList(v []RecoveryPlanVmIpAssignmentIpConfigListInner)`

SetIpConfigList sets IpConfigList field to given value.


### GetVmReference

`func (o *RecoveryPlanVmIpAssignment) GetVmReference() VmReference`

GetVmReference returns the VmReference field if non-nil, zero value otherwise.

### GetVmReferenceOk

`func (o *RecoveryPlanVmIpAssignment) GetVmReferenceOk() (*VmReference, bool)`

GetVmReferenceOk returns a tuple with the VmReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmReference

`func (o *RecoveryPlanVmIpAssignment) SetVmReference(v VmReference)`

SetVmReference sets VmReference field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# RecoveryPlanResourcesParametersFloatingIpAssignmentListInnerVmIpAssignmentListInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TestFloatingIpConfig** | Pointer to [**RecoveryPlanFloatingIpConfig**](RecoveryPlanFloatingIpConfig.md) |  | [optional] 
**RecoveryFloatingIpConfig** | Pointer to [**RecoveryPlanFloatingIpConfig**](RecoveryPlanFloatingIpConfig.md) |  | [optional] 
**VmReference** | [**VmReference**](VmReference.md) |  | 
**VmNicInformation** | [**RecoveryPlanResourcesParametersFloatingIpAssignmentListInnerVmIpAssignmentListInnerVmNicInformation**](RecoveryPlanResourcesParametersFloatingIpAssignmentListInnerVmIpAssignmentListInnerVmNicInformation.md) |  | 

## Methods

### NewRecoveryPlanResourcesParametersFloatingIpAssignmentListInnerVmIpAssignmentListInner

`func NewRecoveryPlanResourcesParametersFloatingIpAssignmentListInnerVmIpAssignmentListInner(vmReference VmReference, vmNicInformation RecoveryPlanResourcesParametersFloatingIpAssignmentListInnerVmIpAssignmentListInnerVmNicInformation, ) *RecoveryPlanResourcesParametersFloatingIpAssignmentListInnerVmIpAssignmentListInner`

NewRecoveryPlanResourcesParametersFloatingIpAssignmentListInnerVmIpAssignmentListInner instantiates a new RecoveryPlanResourcesParametersFloatingIpAssignmentListInnerVmIpAssignmentListInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoveryPlanResourcesParametersFloatingIpAssignmentListInnerVmIpAssignmentListInnerWithDefaults

`func NewRecoveryPlanResourcesParametersFloatingIpAssignmentListInnerVmIpAssignmentListInnerWithDefaults() *RecoveryPlanResourcesParametersFloatingIpAssignmentListInnerVmIpAssignmentListInner`

NewRecoveryPlanResourcesParametersFloatingIpAssignmentListInnerVmIpAssignmentListInnerWithDefaults instantiates a new RecoveryPlanResourcesParametersFloatingIpAssignmentListInnerVmIpAssignmentListInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTestFloatingIpConfig

`func (o *RecoveryPlanResourcesParametersFloatingIpAssignmentListInnerVmIpAssignmentListInner) GetTestFloatingIpConfig() RecoveryPlanFloatingIpConfig`

GetTestFloatingIpConfig returns the TestFloatingIpConfig field if non-nil, zero value otherwise.

### GetTestFloatingIpConfigOk

`func (o *RecoveryPlanResourcesParametersFloatingIpAssignmentListInnerVmIpAssignmentListInner) GetTestFloatingIpConfigOk() (*RecoveryPlanFloatingIpConfig, bool)`

GetTestFloatingIpConfigOk returns a tuple with the TestFloatingIpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestFloatingIpConfig

`func (o *RecoveryPlanResourcesParametersFloatingIpAssignmentListInnerVmIpAssignmentListInner) SetTestFloatingIpConfig(v RecoveryPlanFloatingIpConfig)`

SetTestFloatingIpConfig sets TestFloatingIpConfig field to given value.

### HasTestFloatingIpConfig

`func (o *RecoveryPlanResourcesParametersFloatingIpAssignmentListInnerVmIpAssignmentListInner) HasTestFloatingIpConfig() bool`

HasTestFloatingIpConfig returns a boolean if a field has been set.

### GetRecoveryFloatingIpConfig

`func (o *RecoveryPlanResourcesParametersFloatingIpAssignmentListInnerVmIpAssignmentListInner) GetRecoveryFloatingIpConfig() RecoveryPlanFloatingIpConfig`

GetRecoveryFloatingIpConfig returns the RecoveryFloatingIpConfig field if non-nil, zero value otherwise.

### GetRecoveryFloatingIpConfigOk

`func (o *RecoveryPlanResourcesParametersFloatingIpAssignmentListInnerVmIpAssignmentListInner) GetRecoveryFloatingIpConfigOk() (*RecoveryPlanFloatingIpConfig, bool)`

GetRecoveryFloatingIpConfigOk returns a tuple with the RecoveryFloatingIpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryFloatingIpConfig

`func (o *RecoveryPlanResourcesParametersFloatingIpAssignmentListInnerVmIpAssignmentListInner) SetRecoveryFloatingIpConfig(v RecoveryPlanFloatingIpConfig)`

SetRecoveryFloatingIpConfig sets RecoveryFloatingIpConfig field to given value.

### HasRecoveryFloatingIpConfig

`func (o *RecoveryPlanResourcesParametersFloatingIpAssignmentListInnerVmIpAssignmentListInner) HasRecoveryFloatingIpConfig() bool`

HasRecoveryFloatingIpConfig returns a boolean if a field has been set.

### GetVmReference

`func (o *RecoveryPlanResourcesParametersFloatingIpAssignmentListInnerVmIpAssignmentListInner) GetVmReference() VmReference`

GetVmReference returns the VmReference field if non-nil, zero value otherwise.

### GetVmReferenceOk

`func (o *RecoveryPlanResourcesParametersFloatingIpAssignmentListInnerVmIpAssignmentListInner) GetVmReferenceOk() (*VmReference, bool)`

GetVmReferenceOk returns a tuple with the VmReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmReference

`func (o *RecoveryPlanResourcesParametersFloatingIpAssignmentListInnerVmIpAssignmentListInner) SetVmReference(v VmReference)`

SetVmReference sets VmReference field to given value.


### GetVmNicInformation

`func (o *RecoveryPlanResourcesParametersFloatingIpAssignmentListInnerVmIpAssignmentListInner) GetVmNicInformation() RecoveryPlanResourcesParametersFloatingIpAssignmentListInnerVmIpAssignmentListInnerVmNicInformation`

GetVmNicInformation returns the VmNicInformation field if non-nil, zero value otherwise.

### GetVmNicInformationOk

`func (o *RecoveryPlanResourcesParametersFloatingIpAssignmentListInnerVmIpAssignmentListInner) GetVmNicInformationOk() (*RecoveryPlanResourcesParametersFloatingIpAssignmentListInnerVmIpAssignmentListInnerVmNicInformation, bool)`

GetVmNicInformationOk returns a tuple with the VmNicInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmNicInformation

`func (o *RecoveryPlanResourcesParametersFloatingIpAssignmentListInnerVmIpAssignmentListInner) SetVmNicInformation(v RecoveryPlanResourcesParametersFloatingIpAssignmentListInnerVmIpAssignmentListInnerVmNicInformation)`

SetVmNicInformation sets VmNicInformation field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# VmHostAffinityLegacyPolicyConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HostReferenceList** | Pointer to [**[]HostReference**](HostReference.md) | List of host references which are affined to VM. | [optional] 
**VmReference** | Pointer to [**VmReference**](VmReference.md) |  | [optional] 

## Methods

### NewVmHostAffinityLegacyPolicyConfig

`func NewVmHostAffinityLegacyPolicyConfig() *VmHostAffinityLegacyPolicyConfig`

NewVmHostAffinityLegacyPolicyConfig instantiates a new VmHostAffinityLegacyPolicyConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmHostAffinityLegacyPolicyConfigWithDefaults

`func NewVmHostAffinityLegacyPolicyConfigWithDefaults() *VmHostAffinityLegacyPolicyConfig`

NewVmHostAffinityLegacyPolicyConfigWithDefaults instantiates a new VmHostAffinityLegacyPolicyConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostReferenceList

`func (o *VmHostAffinityLegacyPolicyConfig) GetHostReferenceList() []HostReference`

GetHostReferenceList returns the HostReferenceList field if non-nil, zero value otherwise.

### GetHostReferenceListOk

`func (o *VmHostAffinityLegacyPolicyConfig) GetHostReferenceListOk() (*[]HostReference, bool)`

GetHostReferenceListOk returns a tuple with the HostReferenceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostReferenceList

`func (o *VmHostAffinityLegacyPolicyConfig) SetHostReferenceList(v []HostReference)`

SetHostReferenceList sets HostReferenceList field to given value.

### HasHostReferenceList

`func (o *VmHostAffinityLegacyPolicyConfig) HasHostReferenceList() bool`

HasHostReferenceList returns a boolean if a field has been set.

### GetVmReference

`func (o *VmHostAffinityLegacyPolicyConfig) GetVmReference() VmReference`

GetVmReference returns the VmReference field if non-nil, zero value otherwise.

### GetVmReferenceOk

`func (o *VmHostAffinityLegacyPolicyConfig) GetVmReferenceOk() (*VmReference, bool)`

GetVmReferenceOk returns a tuple with the VmReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmReference

`func (o *VmHostAffinityLegacyPolicyConfig) SetVmReference(v VmReference)`

SetVmReference sets VmReference field to given value.

### HasVmReference

`func (o *VmHostAffinityLegacyPolicyConfig) HasVmReference() bool`

HasVmReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



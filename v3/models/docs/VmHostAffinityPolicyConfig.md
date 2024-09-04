# VmHostAffinityPolicyConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VmCategories** | **map[string][]string** | VM categories across whom affinity is being defined.  | 
**Description** | Pointer to **string** | Description of the policy. | [optional] 
**HostCategories** | **map[string][]string** | Host categories across whom affinity is being defined.  | 
**Name** | **string** | Name of the policy. | 

## Methods

### NewVmHostAffinityPolicyConfig

`func NewVmHostAffinityPolicyConfig(vmCategories map[string][]string, hostCategories map[string][]string, name string, ) *VmHostAffinityPolicyConfig`

NewVmHostAffinityPolicyConfig instantiates a new VmHostAffinityPolicyConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmHostAffinityPolicyConfigWithDefaults

`func NewVmHostAffinityPolicyConfigWithDefaults() *VmHostAffinityPolicyConfig`

NewVmHostAffinityPolicyConfigWithDefaults instantiates a new VmHostAffinityPolicyConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVmCategories

`func (o *VmHostAffinityPolicyConfig) GetVmCategories() map[string][]string`

GetVmCategories returns the VmCategories field if non-nil, zero value otherwise.

### GetVmCategoriesOk

`func (o *VmHostAffinityPolicyConfig) GetVmCategoriesOk() (*map[string][]string, bool)`

GetVmCategoriesOk returns a tuple with the VmCategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmCategories

`func (o *VmHostAffinityPolicyConfig) SetVmCategories(v map[string][]string)`

SetVmCategories sets VmCategories field to given value.


### GetDescription

`func (o *VmHostAffinityPolicyConfig) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VmHostAffinityPolicyConfig) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VmHostAffinityPolicyConfig) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VmHostAffinityPolicyConfig) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetHostCategories

`func (o *VmHostAffinityPolicyConfig) GetHostCategories() map[string][]string`

GetHostCategories returns the HostCategories field if non-nil, zero value otherwise.

### GetHostCategoriesOk

`func (o *VmHostAffinityPolicyConfig) GetHostCategoriesOk() (*map[string][]string, bool)`

GetHostCategoriesOk returns a tuple with the HostCategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostCategories

`func (o *VmHostAffinityPolicyConfig) SetHostCategories(v map[string][]string)`

SetHostCategories sets HostCategories field to given value.


### GetName

`func (o *VmHostAffinityPolicyConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VmHostAffinityPolicyConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VmHostAffinityPolicyConfig) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



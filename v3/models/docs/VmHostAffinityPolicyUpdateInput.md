# VmHostAffinityPolicyUpdateInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | [**VmHostAffinityPolicyConfig**](VmHostAffinityPolicyConfig.md) |  | 
**Metadata** | [**VmHostAffinityPolicyUpdateMetadata**](VmHostAffinityPolicyUpdateMetadata.md) |  | 

## Methods

### NewVmHostAffinityPolicyUpdateInput

`func NewVmHostAffinityPolicyUpdateInput(config VmHostAffinityPolicyConfig, metadata VmHostAffinityPolicyUpdateMetadata, ) *VmHostAffinityPolicyUpdateInput`

NewVmHostAffinityPolicyUpdateInput instantiates a new VmHostAffinityPolicyUpdateInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmHostAffinityPolicyUpdateInputWithDefaults

`func NewVmHostAffinityPolicyUpdateInputWithDefaults() *VmHostAffinityPolicyUpdateInput`

NewVmHostAffinityPolicyUpdateInputWithDefaults instantiates a new VmHostAffinityPolicyUpdateInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *VmHostAffinityPolicyUpdateInput) GetConfig() VmHostAffinityPolicyConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *VmHostAffinityPolicyUpdateInput) GetConfigOk() (*VmHostAffinityPolicyConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *VmHostAffinityPolicyUpdateInput) SetConfig(v VmHostAffinityPolicyConfig)`

SetConfig sets Config field to given value.


### GetMetadata

`func (o *VmHostAffinityPolicyUpdateInput) GetMetadata() VmHostAffinityPolicyUpdateMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *VmHostAffinityPolicyUpdateInput) GetMetadataOk() (*VmHostAffinityPolicyUpdateMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *VmHostAffinityPolicyUpdateInput) SetMetadata(v VmHostAffinityPolicyUpdateMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



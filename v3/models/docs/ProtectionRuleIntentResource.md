# ProtectionRuleIntentResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**ProtectionRuleDefStatus**](ProtectionRuleDefStatus.md) |  | [optional] 
**Spec** | Pointer to [**ProtectionRule**](ProtectionRule.md) |  | [optional] 
**ApiVersion** | Pointer to **string** | API Version of the Nutanix v3 API framework. | [optional] [readonly] [default to "3.1.0"]
**Metadata** | [**ProtectionRuleMetadata**](ProtectionRuleMetadata.md) |  | 

## Methods

### NewProtectionRuleIntentResource

`func NewProtectionRuleIntentResource(metadata ProtectionRuleMetadata, ) *ProtectionRuleIntentResource`

NewProtectionRuleIntentResource instantiates a new ProtectionRuleIntentResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProtectionRuleIntentResourceWithDefaults

`func NewProtectionRuleIntentResourceWithDefaults() *ProtectionRuleIntentResource`

NewProtectionRuleIntentResourceWithDefaults instantiates a new ProtectionRuleIntentResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ProtectionRuleIntentResource) GetStatus() ProtectionRuleDefStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ProtectionRuleIntentResource) GetStatusOk() (*ProtectionRuleDefStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ProtectionRuleIntentResource) SetStatus(v ProtectionRuleDefStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ProtectionRuleIntentResource) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSpec

`func (o *ProtectionRuleIntentResource) GetSpec() ProtectionRule`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *ProtectionRuleIntentResource) GetSpecOk() (*ProtectionRule, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *ProtectionRuleIntentResource) SetSpec(v ProtectionRule)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *ProtectionRuleIntentResource) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetApiVersion

`func (o *ProtectionRuleIntentResource) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ProtectionRuleIntentResource) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ProtectionRuleIntentResource) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *ProtectionRuleIntentResource) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetMetadata

`func (o *ProtectionRuleIntentResource) GetMetadata() ProtectionRuleMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ProtectionRuleIntentResource) GetMetadataOk() (*ProtectionRuleMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ProtectionRuleIntentResource) SetMetadata(v ProtectionRuleMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



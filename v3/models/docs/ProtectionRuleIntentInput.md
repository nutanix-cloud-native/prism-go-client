# ProtectionRuleIntentInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Spec** | [**ProtectionRule**](ProtectionRule.md) |  | 
**ApiVersion** | Pointer to **string** | API Version of the Nutanix v3 API framework. | [optional] [readonly] [default to "3.1.0"]
**Metadata** | [**ProtectionRuleMetadata**](ProtectionRuleMetadata.md) |  | 

## Methods

### NewProtectionRuleIntentInput

`func NewProtectionRuleIntentInput(spec ProtectionRule, metadata ProtectionRuleMetadata, ) *ProtectionRuleIntentInput`

NewProtectionRuleIntentInput instantiates a new ProtectionRuleIntentInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProtectionRuleIntentInputWithDefaults

`func NewProtectionRuleIntentInputWithDefaults() *ProtectionRuleIntentInput`

NewProtectionRuleIntentInputWithDefaults instantiates a new ProtectionRuleIntentInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpec

`func (o *ProtectionRuleIntentInput) GetSpec() ProtectionRule`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *ProtectionRuleIntentInput) GetSpecOk() (*ProtectionRule, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *ProtectionRuleIntentInput) SetSpec(v ProtectionRule)`

SetSpec sets Spec field to given value.


### GetApiVersion

`func (o *ProtectionRuleIntentInput) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ProtectionRuleIntentInput) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ProtectionRuleIntentInput) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *ProtectionRuleIntentInput) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetMetadata

`func (o *ProtectionRuleIntentInput) GetMetadata() ProtectionRuleMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ProtectionRuleIntentInput) GetMetadataOk() (*ProtectionRuleMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ProtectionRuleIntentInput) SetMetadata(v ProtectionRuleMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



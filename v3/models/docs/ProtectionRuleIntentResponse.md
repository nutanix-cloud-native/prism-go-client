# ProtectionRuleIntentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**ProtectionRuleDefStatus**](ProtectionRuleDefStatus.md) |  | [optional] 
**Spec** | Pointer to [**ProtectionRule**](ProtectionRule.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**ProtectionRuleMetadata**](ProtectionRuleMetadata.md) |  | 

## Methods

### NewProtectionRuleIntentResponse

`func NewProtectionRuleIntentResponse(apiVersion string, metadata ProtectionRuleMetadata, ) *ProtectionRuleIntentResponse`

NewProtectionRuleIntentResponse instantiates a new ProtectionRuleIntentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProtectionRuleIntentResponseWithDefaults

`func NewProtectionRuleIntentResponseWithDefaults() *ProtectionRuleIntentResponse`

NewProtectionRuleIntentResponseWithDefaults instantiates a new ProtectionRuleIntentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ProtectionRuleIntentResponse) GetStatus() ProtectionRuleDefStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ProtectionRuleIntentResponse) GetStatusOk() (*ProtectionRuleDefStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ProtectionRuleIntentResponse) SetStatus(v ProtectionRuleDefStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ProtectionRuleIntentResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSpec

`func (o *ProtectionRuleIntentResponse) GetSpec() ProtectionRule`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *ProtectionRuleIntentResponse) GetSpecOk() (*ProtectionRule, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *ProtectionRuleIntentResponse) SetSpec(v ProtectionRule)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *ProtectionRuleIntentResponse) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetApiVersion

`func (o *ProtectionRuleIntentResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ProtectionRuleIntentResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ProtectionRuleIntentResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *ProtectionRuleIntentResponse) GetMetadata() ProtectionRuleMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ProtectionRuleIntentResponse) GetMetadataOk() (*ProtectionRuleMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ProtectionRuleIntentResponse) SetMetadata(v ProtectionRuleMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



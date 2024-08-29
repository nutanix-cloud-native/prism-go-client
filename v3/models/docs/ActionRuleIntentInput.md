# ActionRuleIntentInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Spec** | [**ActionRule**](ActionRule.md) |  | 
**ApiVersion** | Pointer to **string** | API Version of the Nutanix v3 API framework. | [optional] [readonly] [default to "3.1.0"]
**Metadata** | [**ActionRuleMetadata**](ActionRuleMetadata.md) |  | 

## Methods

### NewActionRuleIntentInput

`func NewActionRuleIntentInput(spec ActionRule, metadata ActionRuleMetadata, ) *ActionRuleIntentInput`

NewActionRuleIntentInput instantiates a new ActionRuleIntentInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionRuleIntentInputWithDefaults

`func NewActionRuleIntentInputWithDefaults() *ActionRuleIntentInput`

NewActionRuleIntentInputWithDefaults instantiates a new ActionRuleIntentInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpec

`func (o *ActionRuleIntentInput) GetSpec() ActionRule`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *ActionRuleIntentInput) GetSpecOk() (*ActionRule, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *ActionRuleIntentInput) SetSpec(v ActionRule)`

SetSpec sets Spec field to given value.


### GetApiVersion

`func (o *ActionRuleIntentInput) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ActionRuleIntentInput) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ActionRuleIntentInput) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *ActionRuleIntentInput) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetMetadata

`func (o *ActionRuleIntentInput) GetMetadata() ActionRuleMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ActionRuleIntentInput) GetMetadataOk() (*ActionRuleMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ActionRuleIntentInput) SetMetadata(v ActionRuleMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



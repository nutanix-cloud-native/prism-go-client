# ActionRuleIntentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**ActionRuleDefStatus**](ActionRuleDefStatus.md) |  | [optional] 
**Spec** | Pointer to [**ActionRule**](ActionRule.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**ActionRuleMetadata**](ActionRuleMetadata.md) |  | 

## Methods

### NewActionRuleIntentResponse

`func NewActionRuleIntentResponse(apiVersion string, metadata ActionRuleMetadata, ) *ActionRuleIntentResponse`

NewActionRuleIntentResponse instantiates a new ActionRuleIntentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionRuleIntentResponseWithDefaults

`func NewActionRuleIntentResponseWithDefaults() *ActionRuleIntentResponse`

NewActionRuleIntentResponseWithDefaults instantiates a new ActionRuleIntentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ActionRuleIntentResponse) GetStatus() ActionRuleDefStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ActionRuleIntentResponse) GetStatusOk() (*ActionRuleDefStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ActionRuleIntentResponse) SetStatus(v ActionRuleDefStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ActionRuleIntentResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSpec

`func (o *ActionRuleIntentResponse) GetSpec() ActionRule`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *ActionRuleIntentResponse) GetSpecOk() (*ActionRule, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *ActionRuleIntentResponse) SetSpec(v ActionRule)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *ActionRuleIntentResponse) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetApiVersion

`func (o *ActionRuleIntentResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ActionRuleIntentResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ActionRuleIntentResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *ActionRuleIntentResponse) GetMetadata() ActionRuleMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ActionRuleIntentResponse) GetMetadataOk() (*ActionRuleMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ActionRuleIntentResponse) SetMetadata(v ActionRuleMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



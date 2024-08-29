# ActionRuleResultIntentResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**ActionRuleResultDefStatus**](ActionRuleResultDefStatus.md) |  | [optional] 
**Spec** | Pointer to [**ActionRuleResult**](ActionRuleResult.md) |  | [optional] 
**ApiVersion** | Pointer to **string** | API Version of the Nutanix v3 API framework. | [optional] [readonly] [default to "3.1.0"]
**Metadata** | [**ActionRuleResultMetadata**](ActionRuleResultMetadata.md) |  | 

## Methods

### NewActionRuleResultIntentResource

`func NewActionRuleResultIntentResource(metadata ActionRuleResultMetadata, ) *ActionRuleResultIntentResource`

NewActionRuleResultIntentResource instantiates a new ActionRuleResultIntentResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionRuleResultIntentResourceWithDefaults

`func NewActionRuleResultIntentResourceWithDefaults() *ActionRuleResultIntentResource`

NewActionRuleResultIntentResourceWithDefaults instantiates a new ActionRuleResultIntentResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ActionRuleResultIntentResource) GetStatus() ActionRuleResultDefStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ActionRuleResultIntentResource) GetStatusOk() (*ActionRuleResultDefStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ActionRuleResultIntentResource) SetStatus(v ActionRuleResultDefStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ActionRuleResultIntentResource) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSpec

`func (o *ActionRuleResultIntentResource) GetSpec() ActionRuleResult`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *ActionRuleResultIntentResource) GetSpecOk() (*ActionRuleResult, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *ActionRuleResultIntentResource) SetSpec(v ActionRuleResult)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *ActionRuleResultIntentResource) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetApiVersion

`func (o *ActionRuleResultIntentResource) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ActionRuleResultIntentResource) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ActionRuleResultIntentResource) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *ActionRuleResultIntentResource) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetMetadata

`func (o *ActionRuleResultIntentResource) GetMetadata() ActionRuleResultMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ActionRuleResultIntentResource) GetMetadataOk() (*ActionRuleResultMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ActionRuleResultIntentResource) SetMetadata(v ActionRuleResultMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



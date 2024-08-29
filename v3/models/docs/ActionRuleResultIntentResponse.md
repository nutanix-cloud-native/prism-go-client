# ActionRuleResultIntentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**ActionRuleResultDefStatus**](ActionRuleResultDefStatus.md) |  | [optional] 
**Spec** | Pointer to [**ActionRuleResult**](ActionRuleResult.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**ActionRuleResultMetadata**](ActionRuleResultMetadata.md) |  | 

## Methods

### NewActionRuleResultIntentResponse

`func NewActionRuleResultIntentResponse(apiVersion string, metadata ActionRuleResultMetadata, ) *ActionRuleResultIntentResponse`

NewActionRuleResultIntentResponse instantiates a new ActionRuleResultIntentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionRuleResultIntentResponseWithDefaults

`func NewActionRuleResultIntentResponseWithDefaults() *ActionRuleResultIntentResponse`

NewActionRuleResultIntentResponseWithDefaults instantiates a new ActionRuleResultIntentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ActionRuleResultIntentResponse) GetStatus() ActionRuleResultDefStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ActionRuleResultIntentResponse) GetStatusOk() (*ActionRuleResultDefStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ActionRuleResultIntentResponse) SetStatus(v ActionRuleResultDefStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ActionRuleResultIntentResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSpec

`func (o *ActionRuleResultIntentResponse) GetSpec() ActionRuleResult`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *ActionRuleResultIntentResponse) GetSpecOk() (*ActionRuleResult, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *ActionRuleResultIntentResponse) SetSpec(v ActionRuleResult)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *ActionRuleResultIntentResponse) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetApiVersion

`func (o *ActionRuleResultIntentResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ActionRuleResultIntentResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ActionRuleResultIntentResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *ActionRuleResultIntentResponse) GetMetadata() ActionRuleResultMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ActionRuleResultIntentResponse) GetMetadataOk() (*ActionRuleResultMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ActionRuleResultIntentResponse) SetMetadata(v ActionRuleResultMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



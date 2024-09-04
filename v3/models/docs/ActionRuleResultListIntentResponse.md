# ActionRuleResultListIntentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entities** | Pointer to [**[]ActionRuleResultIntentResource**](ActionRuleResultIntentResource.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**ActionRuleResultListMetadataOutput**](ActionRuleResultListMetadataOutput.md) |  | 

## Methods

### NewActionRuleResultListIntentResponse

`func NewActionRuleResultListIntentResponse(apiVersion string, metadata ActionRuleResultListMetadataOutput, ) *ActionRuleResultListIntentResponse`

NewActionRuleResultListIntentResponse instantiates a new ActionRuleResultListIntentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionRuleResultListIntentResponseWithDefaults

`func NewActionRuleResultListIntentResponseWithDefaults() *ActionRuleResultListIntentResponse`

NewActionRuleResultListIntentResponseWithDefaults instantiates a new ActionRuleResultListIntentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntities

`func (o *ActionRuleResultListIntentResponse) GetEntities() []ActionRuleResultIntentResource`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *ActionRuleResultListIntentResponse) GetEntitiesOk() (*[]ActionRuleResultIntentResource, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *ActionRuleResultListIntentResponse) SetEntities(v []ActionRuleResultIntentResource)`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *ActionRuleResultListIntentResponse) HasEntities() bool`

HasEntities returns a boolean if a field has been set.

### GetApiVersion

`func (o *ActionRuleResultListIntentResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ActionRuleResultListIntentResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ActionRuleResultListIntentResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *ActionRuleResultListIntentResponse) GetMetadata() ActionRuleResultListMetadataOutput`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ActionRuleResultListIntentResponse) GetMetadataOk() (*ActionRuleResultListMetadataOutput, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ActionRuleResultListIntentResponse) SetMetadata(v ActionRuleResultListMetadataOutput)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



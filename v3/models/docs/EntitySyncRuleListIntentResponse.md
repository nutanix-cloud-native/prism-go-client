# EntitySyncRuleListIntentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entities** | Pointer to [**[]EntitySyncRuleIntentResource**](EntitySyncRuleIntentResource.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**EntitySyncRuleListMetadataOutput**](EntitySyncRuleListMetadataOutput.md) |  | 

## Methods

### NewEntitySyncRuleListIntentResponse

`func NewEntitySyncRuleListIntentResponse(apiVersion string, metadata EntitySyncRuleListMetadataOutput, ) *EntitySyncRuleListIntentResponse`

NewEntitySyncRuleListIntentResponse instantiates a new EntitySyncRuleListIntentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitySyncRuleListIntentResponseWithDefaults

`func NewEntitySyncRuleListIntentResponseWithDefaults() *EntitySyncRuleListIntentResponse`

NewEntitySyncRuleListIntentResponseWithDefaults instantiates a new EntitySyncRuleListIntentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntities

`func (o *EntitySyncRuleListIntentResponse) GetEntities() []EntitySyncRuleIntentResource`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *EntitySyncRuleListIntentResponse) GetEntitiesOk() (*[]EntitySyncRuleIntentResource, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *EntitySyncRuleListIntentResponse) SetEntities(v []EntitySyncRuleIntentResource)`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *EntitySyncRuleListIntentResponse) HasEntities() bool`

HasEntities returns a boolean if a field has been set.

### GetApiVersion

`func (o *EntitySyncRuleListIntentResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *EntitySyncRuleListIntentResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *EntitySyncRuleListIntentResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *EntitySyncRuleListIntentResponse) GetMetadata() EntitySyncRuleListMetadataOutput`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *EntitySyncRuleListIntentResponse) GetMetadataOk() (*EntitySyncRuleListMetadataOutput, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *EntitySyncRuleListIntentResponse) SetMetadata(v EntitySyncRuleListMetadataOutput)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



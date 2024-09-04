# EntitySyncRuleIntentResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**EntitySyncRuleDefStatus**](EntitySyncRuleDefStatus.md) |  | [optional] 
**Spec** | Pointer to [**EntitySyncRule**](EntitySyncRule.md) |  | [optional] 
**ApiVersion** | Pointer to **string** | API Version of the Nutanix v3 API framework. | [optional] [readonly] [default to "3.1.0"]
**Metadata** | [**EntitySyncRuleMetadata**](EntitySyncRuleMetadata.md) |  | 

## Methods

### NewEntitySyncRuleIntentResource

`func NewEntitySyncRuleIntentResource(metadata EntitySyncRuleMetadata, ) *EntitySyncRuleIntentResource`

NewEntitySyncRuleIntentResource instantiates a new EntitySyncRuleIntentResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitySyncRuleIntentResourceWithDefaults

`func NewEntitySyncRuleIntentResourceWithDefaults() *EntitySyncRuleIntentResource`

NewEntitySyncRuleIntentResourceWithDefaults instantiates a new EntitySyncRuleIntentResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *EntitySyncRuleIntentResource) GetStatus() EntitySyncRuleDefStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EntitySyncRuleIntentResource) GetStatusOk() (*EntitySyncRuleDefStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EntitySyncRuleIntentResource) SetStatus(v EntitySyncRuleDefStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EntitySyncRuleIntentResource) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSpec

`func (o *EntitySyncRuleIntentResource) GetSpec() EntitySyncRule`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *EntitySyncRuleIntentResource) GetSpecOk() (*EntitySyncRule, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *EntitySyncRuleIntentResource) SetSpec(v EntitySyncRule)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *EntitySyncRuleIntentResource) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetApiVersion

`func (o *EntitySyncRuleIntentResource) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *EntitySyncRuleIntentResource) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *EntitySyncRuleIntentResource) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *EntitySyncRuleIntentResource) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetMetadata

`func (o *EntitySyncRuleIntentResource) GetMetadata() EntitySyncRuleMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *EntitySyncRuleIntentResource) GetMetadataOk() (*EntitySyncRuleMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *EntitySyncRuleIntentResource) SetMetadata(v EntitySyncRuleMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



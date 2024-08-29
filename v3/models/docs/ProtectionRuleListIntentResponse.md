# ProtectionRuleListIntentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entities** | Pointer to [**[]ProtectionRuleIntentResource**](ProtectionRuleIntentResource.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**ProtectionRuleListMetadataOutput**](ProtectionRuleListMetadataOutput.md) |  | 

## Methods

### NewProtectionRuleListIntentResponse

`func NewProtectionRuleListIntentResponse(apiVersion string, metadata ProtectionRuleListMetadataOutput, ) *ProtectionRuleListIntentResponse`

NewProtectionRuleListIntentResponse instantiates a new ProtectionRuleListIntentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProtectionRuleListIntentResponseWithDefaults

`func NewProtectionRuleListIntentResponseWithDefaults() *ProtectionRuleListIntentResponse`

NewProtectionRuleListIntentResponseWithDefaults instantiates a new ProtectionRuleListIntentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntities

`func (o *ProtectionRuleListIntentResponse) GetEntities() []ProtectionRuleIntentResource`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *ProtectionRuleListIntentResponse) GetEntitiesOk() (*[]ProtectionRuleIntentResource, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *ProtectionRuleListIntentResponse) SetEntities(v []ProtectionRuleIntentResource)`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *ProtectionRuleListIntentResponse) HasEntities() bool`

HasEntities returns a boolean if a field has been set.

### GetApiVersion

`func (o *ProtectionRuleListIntentResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ProtectionRuleListIntentResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ProtectionRuleListIntentResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *ProtectionRuleListIntentResponse) GetMetadata() ProtectionRuleListMetadataOutput`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ProtectionRuleListIntentResponse) GetMetadataOk() (*ProtectionRuleListMetadataOutput, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ProtectionRuleListIntentResponse) SetMetadata(v ProtectionRuleListMetadataOutput)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



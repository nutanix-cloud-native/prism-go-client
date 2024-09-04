# NetworkSecurityRuleListIntentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entities** | Pointer to [**[]NetworkSecurityRuleIntentResource**](NetworkSecurityRuleIntentResource.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**NetworkSecurityRuleListMetadataOutput**](NetworkSecurityRuleListMetadataOutput.md) |  | 

## Methods

### NewNetworkSecurityRuleListIntentResponse

`func NewNetworkSecurityRuleListIntentResponse(apiVersion string, metadata NetworkSecurityRuleListMetadataOutput, ) *NetworkSecurityRuleListIntentResponse`

NewNetworkSecurityRuleListIntentResponse instantiates a new NetworkSecurityRuleListIntentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkSecurityRuleListIntentResponseWithDefaults

`func NewNetworkSecurityRuleListIntentResponseWithDefaults() *NetworkSecurityRuleListIntentResponse`

NewNetworkSecurityRuleListIntentResponseWithDefaults instantiates a new NetworkSecurityRuleListIntentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntities

`func (o *NetworkSecurityRuleListIntentResponse) GetEntities() []NetworkSecurityRuleIntentResource`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *NetworkSecurityRuleListIntentResponse) GetEntitiesOk() (*[]NetworkSecurityRuleIntentResource, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *NetworkSecurityRuleListIntentResponse) SetEntities(v []NetworkSecurityRuleIntentResource)`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *NetworkSecurityRuleListIntentResponse) HasEntities() bool`

HasEntities returns a boolean if a field has been set.

### GetApiVersion

`func (o *NetworkSecurityRuleListIntentResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *NetworkSecurityRuleListIntentResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *NetworkSecurityRuleListIntentResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *NetworkSecurityRuleListIntentResponse) GetMetadata() NetworkSecurityRuleListMetadataOutput`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *NetworkSecurityRuleListIntentResponse) GetMetadataOk() (*NetworkSecurityRuleListMetadataOutput, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *NetworkSecurityRuleListIntentResponse) SetMetadata(v NetworkSecurityRuleListMetadataOutput)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



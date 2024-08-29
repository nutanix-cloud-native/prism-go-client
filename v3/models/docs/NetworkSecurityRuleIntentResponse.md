# NetworkSecurityRuleIntentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**NetworkSecurityRuleDefStatus**](NetworkSecurityRuleDefStatus.md) |  | [optional] 
**Spec** | Pointer to [**NetworkSecurityRule**](NetworkSecurityRule.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**NetworkSecurityRuleMetadata**](NetworkSecurityRuleMetadata.md) |  | 

## Methods

### NewNetworkSecurityRuleIntentResponse

`func NewNetworkSecurityRuleIntentResponse(apiVersion string, metadata NetworkSecurityRuleMetadata, ) *NetworkSecurityRuleIntentResponse`

NewNetworkSecurityRuleIntentResponse instantiates a new NetworkSecurityRuleIntentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkSecurityRuleIntentResponseWithDefaults

`func NewNetworkSecurityRuleIntentResponseWithDefaults() *NetworkSecurityRuleIntentResponse`

NewNetworkSecurityRuleIntentResponseWithDefaults instantiates a new NetworkSecurityRuleIntentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *NetworkSecurityRuleIntentResponse) GetStatus() NetworkSecurityRuleDefStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NetworkSecurityRuleIntentResponse) GetStatusOk() (*NetworkSecurityRuleDefStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NetworkSecurityRuleIntentResponse) SetStatus(v NetworkSecurityRuleDefStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *NetworkSecurityRuleIntentResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSpec

`func (o *NetworkSecurityRuleIntentResponse) GetSpec() NetworkSecurityRule`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *NetworkSecurityRuleIntentResponse) GetSpecOk() (*NetworkSecurityRule, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *NetworkSecurityRuleIntentResponse) SetSpec(v NetworkSecurityRule)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *NetworkSecurityRuleIntentResponse) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetApiVersion

`func (o *NetworkSecurityRuleIntentResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *NetworkSecurityRuleIntentResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *NetworkSecurityRuleIntentResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *NetworkSecurityRuleIntentResponse) GetMetadata() NetworkSecurityRuleMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *NetworkSecurityRuleIntentResponse) GetMetadataOk() (*NetworkSecurityRuleMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *NetworkSecurityRuleIntentResponse) SetMetadata(v NetworkSecurityRuleMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



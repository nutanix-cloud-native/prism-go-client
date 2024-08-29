# NetworkSecurityRuleImportResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityList** | Pointer to [**[]NetworkSecurityRuleImportEntity**](NetworkSecurityRuleImportEntity.md) | entity_list is returned if the imported data is not applied (commit_mode is dryrun). | [optional] 

## Methods

### NewNetworkSecurityRuleImportResponse

`func NewNetworkSecurityRuleImportResponse() *NetworkSecurityRuleImportResponse`

NewNetworkSecurityRuleImportResponse instantiates a new NetworkSecurityRuleImportResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkSecurityRuleImportResponseWithDefaults

`func NewNetworkSecurityRuleImportResponseWithDefaults() *NetworkSecurityRuleImportResponse`

NewNetworkSecurityRuleImportResponseWithDefaults instantiates a new NetworkSecurityRuleImportResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityList

`func (o *NetworkSecurityRuleImportResponse) GetEntityList() []NetworkSecurityRuleImportEntity`

GetEntityList returns the EntityList field if non-nil, zero value otherwise.

### GetEntityListOk

`func (o *NetworkSecurityRuleImportResponse) GetEntityListOk() (*[]NetworkSecurityRuleImportEntity, bool)`

GetEntityListOk returns a tuple with the EntityList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityList

`func (o *NetworkSecurityRuleImportResponse) SetEntityList(v []NetworkSecurityRuleImportEntity)`

SetEntityList sets EntityList field to given value.

### HasEntityList

`func (o *NetworkSecurityRuleImportResponse) HasEntityList() bool`

HasEntityList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# IdentityCategorizationDirectoryConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MatchingCriteriaList** | Pointer to [**[]IdCategorizationMatchingCriteria**](IdCategorizationMatchingCriteria.md) | The matching criteria used to determine whether an entity will be affected by identity categorization. If not provided, no entity will be affected. Only a single entry in this list is supported today.  | [optional] 
**DomainControllerList** | Pointer to [**[]IdCategorizationDomainController**](IdCategorizationDomainController.md) | List of domain controllers to be used for event scraping. | [optional] 
**DirectoryServiceReference** | [**DirectoryServiceReference**](DirectoryServiceReference.md) |  | 

## Methods

### NewIdentityCategorizationDirectoryConfig

`func NewIdentityCategorizationDirectoryConfig(directoryServiceReference DirectoryServiceReference, ) *IdentityCategorizationDirectoryConfig`

NewIdentityCategorizationDirectoryConfig instantiates a new IdentityCategorizationDirectoryConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityCategorizationDirectoryConfigWithDefaults

`func NewIdentityCategorizationDirectoryConfigWithDefaults() *IdentityCategorizationDirectoryConfig`

NewIdentityCategorizationDirectoryConfigWithDefaults instantiates a new IdentityCategorizationDirectoryConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMatchingCriteriaList

`func (o *IdentityCategorizationDirectoryConfig) GetMatchingCriteriaList() []IdCategorizationMatchingCriteria`

GetMatchingCriteriaList returns the MatchingCriteriaList field if non-nil, zero value otherwise.

### GetMatchingCriteriaListOk

`func (o *IdentityCategorizationDirectoryConfig) GetMatchingCriteriaListOk() (*[]IdCategorizationMatchingCriteria, bool)`

GetMatchingCriteriaListOk returns a tuple with the MatchingCriteriaList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchingCriteriaList

`func (o *IdentityCategorizationDirectoryConfig) SetMatchingCriteriaList(v []IdCategorizationMatchingCriteria)`

SetMatchingCriteriaList sets MatchingCriteriaList field to given value.

### HasMatchingCriteriaList

`func (o *IdentityCategorizationDirectoryConfig) HasMatchingCriteriaList() bool`

HasMatchingCriteriaList returns a boolean if a field has been set.

### GetDomainControllerList

`func (o *IdentityCategorizationDirectoryConfig) GetDomainControllerList() []IdCategorizationDomainController`

GetDomainControllerList returns the DomainControllerList field if non-nil, zero value otherwise.

### GetDomainControllerListOk

`func (o *IdentityCategorizationDirectoryConfig) GetDomainControllerListOk() (*[]IdCategorizationDomainController, bool)`

GetDomainControllerListOk returns a tuple with the DomainControllerList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainControllerList

`func (o *IdentityCategorizationDirectoryConfig) SetDomainControllerList(v []IdCategorizationDomainController)`

SetDomainControllerList sets DomainControllerList field to given value.

### HasDomainControllerList

`func (o *IdentityCategorizationDirectoryConfig) HasDomainControllerList() bool`

HasDomainControllerList returns a boolean if a field has been set.

### GetDirectoryServiceReference

`func (o *IdentityCategorizationDirectoryConfig) GetDirectoryServiceReference() DirectoryServiceReference`

GetDirectoryServiceReference returns the DirectoryServiceReference field if non-nil, zero value otherwise.

### GetDirectoryServiceReferenceOk

`func (o *IdentityCategorizationDirectoryConfig) GetDirectoryServiceReferenceOk() (*DirectoryServiceReference, bool)`

GetDirectoryServiceReferenceOk returns a tuple with the DirectoryServiceReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryServiceReference

`func (o *IdentityCategorizationDirectoryConfig) SetDirectoryServiceReference(v DirectoryServiceReference)`

SetDirectoryServiceReference sets DirectoryServiceReference field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



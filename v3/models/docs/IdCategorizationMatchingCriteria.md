# IdCategorizationMatchingCriteria

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MatchField** | **string** | The field to match on. Today only NAME is supported, which matches on an entity&#39;s name.  | 
**Criteria** | Pointer to **string** | The criteria to use for matching entities to be categorized. Note that depending on the match type, the usage of this value may differ.  | [optional] 
**MatchType** | **string** | The type of match. Today only CONTAINS and ALL are supported. CONTAINS performs a substring match on the given entity and field for the criteria value whereas ALL allows all string to match on the given entity.  | 
**MatchEntity** | **string** | The entity to perform matching on. Currently, only the target VM that a logon occurred on is supported.  | 

## Methods

### NewIdCategorizationMatchingCriteria

`func NewIdCategorizationMatchingCriteria(matchField string, matchType string, matchEntity string, ) *IdCategorizationMatchingCriteria`

NewIdCategorizationMatchingCriteria instantiates a new IdCategorizationMatchingCriteria object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdCategorizationMatchingCriteriaWithDefaults

`func NewIdCategorizationMatchingCriteriaWithDefaults() *IdCategorizationMatchingCriteria`

NewIdCategorizationMatchingCriteriaWithDefaults instantiates a new IdCategorizationMatchingCriteria object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMatchField

`func (o *IdCategorizationMatchingCriteria) GetMatchField() string`

GetMatchField returns the MatchField field if non-nil, zero value otherwise.

### GetMatchFieldOk

`func (o *IdCategorizationMatchingCriteria) GetMatchFieldOk() (*string, bool)`

GetMatchFieldOk returns a tuple with the MatchField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchField

`func (o *IdCategorizationMatchingCriteria) SetMatchField(v string)`

SetMatchField sets MatchField field to given value.


### GetCriteria

`func (o *IdCategorizationMatchingCriteria) GetCriteria() string`

GetCriteria returns the Criteria field if non-nil, zero value otherwise.

### GetCriteriaOk

`func (o *IdCategorizationMatchingCriteria) GetCriteriaOk() (*string, bool)`

GetCriteriaOk returns a tuple with the Criteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriteria

`func (o *IdCategorizationMatchingCriteria) SetCriteria(v string)`

SetCriteria sets Criteria field to given value.

### HasCriteria

`func (o *IdCategorizationMatchingCriteria) HasCriteria() bool`

HasCriteria returns a boolean if a field has been set.

### GetMatchType

`func (o *IdCategorizationMatchingCriteria) GetMatchType() string`

GetMatchType returns the MatchType field if non-nil, zero value otherwise.

### GetMatchTypeOk

`func (o *IdCategorizationMatchingCriteria) GetMatchTypeOk() (*string, bool)`

GetMatchTypeOk returns a tuple with the MatchType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchType

`func (o *IdCategorizationMatchingCriteria) SetMatchType(v string)`

SetMatchType sets MatchType field to given value.


### GetMatchEntity

`func (o *IdCategorizationMatchingCriteria) GetMatchEntity() string`

GetMatchEntity returns the MatchEntity field if non-nil, zero value otherwise.

### GetMatchEntityOk

`func (o *IdCategorizationMatchingCriteria) GetMatchEntityOk() (*string, bool)`

GetMatchEntityOk returns a tuple with the MatchEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchEntity

`func (o *IdCategorizationMatchingCriteria) SetMatchEntity(v string)`

SetMatchEntity sets MatchEntity field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



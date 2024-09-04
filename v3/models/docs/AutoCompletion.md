# AutoCompletion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SearchType** | Pointer to **string** | Search type for the autocompletion.  | [optional] 
**CompleteQuery** | Pointer to **string** | The complete query corresponding the the auto-completion. | [optional] 
**QueryTermList** | Pointer to [**[]QueryTerm**](QueryTerm.md) | Structured representation that infers query intent unambiguously. Client will echo this information back to the backend. Essentially, it is like a search result link. The list has an item corresponding to every query term. One user query is nothing but a collection of multiple query terms.  | [optional] 
**EntityType** | Pointer to **string** | Enity type or focus for the autocompletion. | [optional] 

## Methods

### NewAutoCompletion

`func NewAutoCompletion() *AutoCompletion`

NewAutoCompletion instantiates a new AutoCompletion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoCompletionWithDefaults

`func NewAutoCompletionWithDefaults() *AutoCompletion`

NewAutoCompletionWithDefaults instantiates a new AutoCompletion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSearchType

`func (o *AutoCompletion) GetSearchType() string`

GetSearchType returns the SearchType field if non-nil, zero value otherwise.

### GetSearchTypeOk

`func (o *AutoCompletion) GetSearchTypeOk() (*string, bool)`

GetSearchTypeOk returns a tuple with the SearchType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchType

`func (o *AutoCompletion) SetSearchType(v string)`

SetSearchType sets SearchType field to given value.

### HasSearchType

`func (o *AutoCompletion) HasSearchType() bool`

HasSearchType returns a boolean if a field has been set.

### GetCompleteQuery

`func (o *AutoCompletion) GetCompleteQuery() string`

GetCompleteQuery returns the CompleteQuery field if non-nil, zero value otherwise.

### GetCompleteQueryOk

`func (o *AutoCompletion) GetCompleteQueryOk() (*string, bool)`

GetCompleteQueryOk returns a tuple with the CompleteQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleteQuery

`func (o *AutoCompletion) SetCompleteQuery(v string)`

SetCompleteQuery sets CompleteQuery field to given value.

### HasCompleteQuery

`func (o *AutoCompletion) HasCompleteQuery() bool`

HasCompleteQuery returns a boolean if a field has been set.

### GetQueryTermList

`func (o *AutoCompletion) GetQueryTermList() []QueryTerm`

GetQueryTermList returns the QueryTermList field if non-nil, zero value otherwise.

### GetQueryTermListOk

`func (o *AutoCompletion) GetQueryTermListOk() (*[]QueryTerm, bool)`

GetQueryTermListOk returns a tuple with the QueryTermList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryTermList

`func (o *AutoCompletion) SetQueryTermList(v []QueryTerm)`

SetQueryTermList sets QueryTermList field to given value.

### HasQueryTermList

`func (o *AutoCompletion) HasQueryTermList() bool`

HasQueryTermList returns a boolean if a field has been set.

### GetEntityType

`func (o *AutoCompletion) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *AutoCompletion) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *AutoCompletion) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.

### HasEntityType

`func (o *AutoCompletion) HasEntityType() bool`

HasEntityType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



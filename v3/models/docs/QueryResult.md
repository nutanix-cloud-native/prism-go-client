# QueryResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QueryExplaination** | Pointer to **string** | A description of the interpretation done by the search engine.  | [optional] 
**AutoCompletionList** | Pointer to [**[]AutoCompletion**](AutoCompletion.md) | List of all the autocompletions for a given user query.  | [optional] 
**QueryContext** | Pointer to **string** | Common context for query which needs to be shared across client and search service. Client can pass some context which will be echoed back along with the response.  | [optional] 
**QueryTermList** | Pointer to [**[]QueryTerm**](QueryTerm.md) | Structured representation that infers query intent unambiguously. Client will echo this information back to the backend. Essentially, it is like a search result link. The list has an item corresponding to every query term. One user query is nothing but a collection of multiple query terms.  | [optional] 
**Result** | Pointer to [**Result**](Result.md) |  | [optional] 
**UserQuery** | Pointer to **string** | User query in simple text. | [optional] 
**TimeTaken** | Pointer to **int64** | Time taken(ms) in executing the search query. | [optional] 
**PlainTextQuery** | Pointer to **bool** | A boolean flag to distinguish between a plain text search_for query and the search_for query with query term list.  | [optional] 

## Methods

### NewQueryResult

`func NewQueryResult() *QueryResult`

NewQueryResult instantiates a new QueryResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryResultWithDefaults

`func NewQueryResultWithDefaults() *QueryResult`

NewQueryResultWithDefaults instantiates a new QueryResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQueryExplaination

`func (o *QueryResult) GetQueryExplaination() string`

GetQueryExplaination returns the QueryExplaination field if non-nil, zero value otherwise.

### GetQueryExplainationOk

`func (o *QueryResult) GetQueryExplainationOk() (*string, bool)`

GetQueryExplainationOk returns a tuple with the QueryExplaination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryExplaination

`func (o *QueryResult) SetQueryExplaination(v string)`

SetQueryExplaination sets QueryExplaination field to given value.

### HasQueryExplaination

`func (o *QueryResult) HasQueryExplaination() bool`

HasQueryExplaination returns a boolean if a field has been set.

### GetAutoCompletionList

`func (o *QueryResult) GetAutoCompletionList() []AutoCompletion`

GetAutoCompletionList returns the AutoCompletionList field if non-nil, zero value otherwise.

### GetAutoCompletionListOk

`func (o *QueryResult) GetAutoCompletionListOk() (*[]AutoCompletion, bool)`

GetAutoCompletionListOk returns a tuple with the AutoCompletionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoCompletionList

`func (o *QueryResult) SetAutoCompletionList(v []AutoCompletion)`

SetAutoCompletionList sets AutoCompletionList field to given value.

### HasAutoCompletionList

`func (o *QueryResult) HasAutoCompletionList() bool`

HasAutoCompletionList returns a boolean if a field has been set.

### GetQueryContext

`func (o *QueryResult) GetQueryContext() string`

GetQueryContext returns the QueryContext field if non-nil, zero value otherwise.

### GetQueryContextOk

`func (o *QueryResult) GetQueryContextOk() (*string, bool)`

GetQueryContextOk returns a tuple with the QueryContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryContext

`func (o *QueryResult) SetQueryContext(v string)`

SetQueryContext sets QueryContext field to given value.

### HasQueryContext

`func (o *QueryResult) HasQueryContext() bool`

HasQueryContext returns a boolean if a field has been set.

### GetQueryTermList

`func (o *QueryResult) GetQueryTermList() []QueryTerm`

GetQueryTermList returns the QueryTermList field if non-nil, zero value otherwise.

### GetQueryTermListOk

`func (o *QueryResult) GetQueryTermListOk() (*[]QueryTerm, bool)`

GetQueryTermListOk returns a tuple with the QueryTermList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryTermList

`func (o *QueryResult) SetQueryTermList(v []QueryTerm)`

SetQueryTermList sets QueryTermList field to given value.

### HasQueryTermList

`func (o *QueryResult) HasQueryTermList() bool`

HasQueryTermList returns a boolean if a field has been set.

### GetResult

`func (o *QueryResult) GetResult() Result`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *QueryResult) GetResultOk() (*Result, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *QueryResult) SetResult(v Result)`

SetResult sets Result field to given value.

### HasResult

`func (o *QueryResult) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetUserQuery

`func (o *QueryResult) GetUserQuery() string`

GetUserQuery returns the UserQuery field if non-nil, zero value otherwise.

### GetUserQueryOk

`func (o *QueryResult) GetUserQueryOk() (*string, bool)`

GetUserQueryOk returns a tuple with the UserQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserQuery

`func (o *QueryResult) SetUserQuery(v string)`

SetUserQuery sets UserQuery field to given value.

### HasUserQuery

`func (o *QueryResult) HasUserQuery() bool`

HasUserQuery returns a boolean if a field has been set.

### GetTimeTaken

`func (o *QueryResult) GetTimeTaken() int64`

GetTimeTaken returns the TimeTaken field if non-nil, zero value otherwise.

### GetTimeTakenOk

`func (o *QueryResult) GetTimeTakenOk() (*int64, bool)`

GetTimeTakenOk returns a tuple with the TimeTaken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeTaken

`func (o *QueryResult) SetTimeTaken(v int64)`

SetTimeTaken sets TimeTaken field to given value.

### HasTimeTaken

`func (o *QueryResult) HasTimeTaken() bool`

HasTimeTaken returns a boolean if a field has been set.

### GetPlainTextQuery

`func (o *QueryResult) GetPlainTextQuery() bool`

GetPlainTextQuery returns the PlainTextQuery field if non-nil, zero value otherwise.

### GetPlainTextQueryOk

`func (o *QueryResult) GetPlainTextQueryOk() (*bool, bool)`

GetPlainTextQueryOk returns a tuple with the PlainTextQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlainTextQuery

`func (o *QueryResult) SetPlainTextQuery(v bool)`

SetPlainTextQuery sets PlainTextQuery field to given value.

### HasPlainTextQuery

`func (o *QueryResult) HasPlainTextQuery() bool`

HasPlainTextQuery returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



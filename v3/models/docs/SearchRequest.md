# SearchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QueryContext** | Pointer to **string** | Common context for query which needs to be shared across client and search service. Client can pass some context which will be echoed back along with the response.  | [optional] 
**ExplicitQuery** | Pointer to **bool** | Flag to indicate the user explicitly made this query (e.g by pressing enter) and is not still typing. Helpful for tracking concrete queries fired by the user.  | [optional] 
**GenerateAutocompletionsOnly** | Pointer to **bool** | Flag to specify  if user is interested only in autocompletions.  | [optional] 
**QueryTermList** | Pointer to [**[]QueryTerm**](QueryTerm.md) | Structured representation that infers query intent unambiguously. Client will echo this information back to the backend. Essentially, it is like a search result link. The list has an item corresponding to every query term. One user query is nothing but a collection of multiple query terms.  | [optional] 
**IsAutocompleteSelection** | Pointer to **bool** | Flag to indicate the user selected an autocomplete. Helpful for tracking concrete autocomplete selections.  | [optional] 
**Timezone** | Pointer to **string** | Timezone in which the query is getting excecuted.  | [optional] 
**UserQuery** | Pointer to **string** | User query in simple text. | [optional] 
**WidgetIdList** | Pointer to **[]string** | Optional list of widgets that a client can request for a specific query.The list is meant to be populated with IDs based on the previous searchresponse. For instance the first response can indicate that the result consists of widget ids \&quot;property_summary\&quot;, \&quot;metric_summary\&quot; (or some other form of unique identifier but without any actual data). The client can then make a second request for those widgets. This helps in performance reasons as well as for refreshing content on demand.  | [optional] 

## Methods

### NewSearchRequest

`func NewSearchRequest() *SearchRequest`

NewSearchRequest instantiates a new SearchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchRequestWithDefaults

`func NewSearchRequestWithDefaults() *SearchRequest`

NewSearchRequestWithDefaults instantiates a new SearchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQueryContext

`func (o *SearchRequest) GetQueryContext() string`

GetQueryContext returns the QueryContext field if non-nil, zero value otherwise.

### GetQueryContextOk

`func (o *SearchRequest) GetQueryContextOk() (*string, bool)`

GetQueryContextOk returns a tuple with the QueryContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryContext

`func (o *SearchRequest) SetQueryContext(v string)`

SetQueryContext sets QueryContext field to given value.

### HasQueryContext

`func (o *SearchRequest) HasQueryContext() bool`

HasQueryContext returns a boolean if a field has been set.

### GetExplicitQuery

`func (o *SearchRequest) GetExplicitQuery() bool`

GetExplicitQuery returns the ExplicitQuery field if non-nil, zero value otherwise.

### GetExplicitQueryOk

`func (o *SearchRequest) GetExplicitQueryOk() (*bool, bool)`

GetExplicitQueryOk returns a tuple with the ExplicitQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplicitQuery

`func (o *SearchRequest) SetExplicitQuery(v bool)`

SetExplicitQuery sets ExplicitQuery field to given value.

### HasExplicitQuery

`func (o *SearchRequest) HasExplicitQuery() bool`

HasExplicitQuery returns a boolean if a field has been set.

### GetGenerateAutocompletionsOnly

`func (o *SearchRequest) GetGenerateAutocompletionsOnly() bool`

GetGenerateAutocompletionsOnly returns the GenerateAutocompletionsOnly field if non-nil, zero value otherwise.

### GetGenerateAutocompletionsOnlyOk

`func (o *SearchRequest) GetGenerateAutocompletionsOnlyOk() (*bool, bool)`

GetGenerateAutocompletionsOnlyOk returns a tuple with the GenerateAutocompletionsOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerateAutocompletionsOnly

`func (o *SearchRequest) SetGenerateAutocompletionsOnly(v bool)`

SetGenerateAutocompletionsOnly sets GenerateAutocompletionsOnly field to given value.

### HasGenerateAutocompletionsOnly

`func (o *SearchRequest) HasGenerateAutocompletionsOnly() bool`

HasGenerateAutocompletionsOnly returns a boolean if a field has been set.

### GetQueryTermList

`func (o *SearchRequest) GetQueryTermList() []QueryTerm`

GetQueryTermList returns the QueryTermList field if non-nil, zero value otherwise.

### GetQueryTermListOk

`func (o *SearchRequest) GetQueryTermListOk() (*[]QueryTerm, bool)`

GetQueryTermListOk returns a tuple with the QueryTermList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryTermList

`func (o *SearchRequest) SetQueryTermList(v []QueryTerm)`

SetQueryTermList sets QueryTermList field to given value.

### HasQueryTermList

`func (o *SearchRequest) HasQueryTermList() bool`

HasQueryTermList returns a boolean if a field has been set.

### GetIsAutocompleteSelection

`func (o *SearchRequest) GetIsAutocompleteSelection() bool`

GetIsAutocompleteSelection returns the IsAutocompleteSelection field if non-nil, zero value otherwise.

### GetIsAutocompleteSelectionOk

`func (o *SearchRequest) GetIsAutocompleteSelectionOk() (*bool, bool)`

GetIsAutocompleteSelectionOk returns a tuple with the IsAutocompleteSelection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutocompleteSelection

`func (o *SearchRequest) SetIsAutocompleteSelection(v bool)`

SetIsAutocompleteSelection sets IsAutocompleteSelection field to given value.

### HasIsAutocompleteSelection

`func (o *SearchRequest) HasIsAutocompleteSelection() bool`

HasIsAutocompleteSelection returns a boolean if a field has been set.

### GetTimezone

`func (o *SearchRequest) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *SearchRequest) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *SearchRequest) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *SearchRequest) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetUserQuery

`func (o *SearchRequest) GetUserQuery() string`

GetUserQuery returns the UserQuery field if non-nil, zero value otherwise.

### GetUserQueryOk

`func (o *SearchRequest) GetUserQueryOk() (*string, bool)`

GetUserQueryOk returns a tuple with the UserQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserQuery

`func (o *SearchRequest) SetUserQuery(v string)`

SetUserQuery sets UserQuery field to given value.

### HasUserQuery

`func (o *SearchRequest) HasUserQuery() bool`

HasUserQuery returns a boolean if a field has been set.

### GetWidgetIdList

`func (o *SearchRequest) GetWidgetIdList() []string`

GetWidgetIdList returns the WidgetIdList field if non-nil, zero value otherwise.

### GetWidgetIdListOk

`func (o *SearchRequest) GetWidgetIdListOk() (*[]string, bool)`

GetWidgetIdListOk returns a tuple with the WidgetIdList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidgetIdList

`func (o *SearchRequest) SetWidgetIdList(v []string)`

SetWidgetIdList sets WidgetIdList field to given value.

### HasWidgetIdList

`func (o *SearchRequest) HasWidgetIdList() bool`

HasWidgetIdList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# Favorite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Route** | Pointer to **string** | Optional route if present to directly link to the favorite. | [optional] 
**QueryTermList** | Pointer to [**[]QueryTerm**](QueryTerm.md) | Structured representation that infers query intent unambiguously. Client will echo this information back to the backend. Essentially, it is like a search result link. The list has an item corresponding to every query term. One user query is nothing but a collection of multiple query terms.  | [optional] 
**CompleteQuery** | Pointer to **string** | Actual query string. | [optional] 
**UUID** | Pointer to **string** | Entity id. | [optional] 

## Methods

### NewFavorite

`func NewFavorite() *Favorite`

NewFavorite instantiates a new Favorite object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFavoriteWithDefaults

`func NewFavoriteWithDefaults() *Favorite`

NewFavoriteWithDefaults instantiates a new Favorite object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoute

`func (o *Favorite) GetRoute() string`

GetRoute returns the Route field if non-nil, zero value otherwise.

### GetRouteOk

`func (o *Favorite) GetRouteOk() (*string, bool)`

GetRouteOk returns a tuple with the Route field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoute

`func (o *Favorite) SetRoute(v string)`

SetRoute sets Route field to given value.

### HasRoute

`func (o *Favorite) HasRoute() bool`

HasRoute returns a boolean if a field has been set.

### GetQueryTermList

`func (o *Favorite) GetQueryTermList() []QueryTerm`

GetQueryTermList returns the QueryTermList field if non-nil, zero value otherwise.

### GetQueryTermListOk

`func (o *Favorite) GetQueryTermListOk() (*[]QueryTerm, bool)`

GetQueryTermListOk returns a tuple with the QueryTermList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryTermList

`func (o *Favorite) SetQueryTermList(v []QueryTerm)`

SetQueryTermList sets QueryTermList field to given value.

### HasQueryTermList

`func (o *Favorite) HasQueryTermList() bool`

HasQueryTermList returns a boolean if a field has been set.

### GetCompleteQuery

`func (o *Favorite) GetCompleteQuery() string`

GetCompleteQuery returns the CompleteQuery field if non-nil, zero value otherwise.

### GetCompleteQueryOk

`func (o *Favorite) GetCompleteQueryOk() (*string, bool)`

GetCompleteQueryOk returns a tuple with the CompleteQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleteQuery

`func (o *Favorite) SetCompleteQuery(v string)`

SetCompleteQuery sets CompleteQuery field to given value.

### HasCompleteQuery

`func (o *Favorite) HasCompleteQuery() bool`

HasCompleteQuery returns a boolean if a field has been set.

### GetUUID

`func (o *Favorite) GetUUID() string`

GetUUID returns the UUID field if non-nil, zero value otherwise.

### GetUUIDOk

`func (o *Favorite) GetUUIDOk() (*string, bool)`

GetUUIDOk returns a tuple with the UUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUUID

`func (o *Favorite) SetUUID(v string)`

SetUUID sets UUID field to given value.

### HasUUID

`func (o *Favorite) HasUUID() bool`

HasUUID returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



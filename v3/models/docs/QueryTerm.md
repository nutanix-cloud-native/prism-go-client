# QueryTerm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Term** | Pointer to **string** | User passed string for the query term. | [optional] 
**TokenList** | Pointer to [**[]Token**](Token.md) | All possible interpretations of the term. | [optional] 

## Methods

### NewQueryTerm

`func NewQueryTerm() *QueryTerm`

NewQueryTerm instantiates a new QueryTerm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryTermWithDefaults

`func NewQueryTermWithDefaults() *QueryTerm`

NewQueryTermWithDefaults instantiates a new QueryTerm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTerm

`func (o *QueryTerm) GetTerm() string`

GetTerm returns the Term field if non-nil, zero value otherwise.

### GetTermOk

`func (o *QueryTerm) GetTermOk() (*string, bool)`

GetTermOk returns a tuple with the Term field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerm

`func (o *QueryTerm) SetTerm(v string)`

SetTerm sets Term field to given value.

### HasTerm

`func (o *QueryTerm) HasTerm() bool`

HasTerm returns a boolean if a field has been set.

### GetTokenList

`func (o *QueryTerm) GetTokenList() []Token`

GetTokenList returns the TokenList field if non-nil, zero value otherwise.

### GetTokenListOk

`func (o *QueryTerm) GetTokenListOk() (*[]Token, bool)`

GetTokenListOk returns a tuple with the TokenList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenList

`func (o *QueryTerm) SetTokenList(v []Token)`

SetTokenList sets TokenList field to given value.

### HasTokenList

`func (o *QueryTerm) HasTokenList() bool`

HasTokenList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



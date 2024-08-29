# Token

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayContext** | Pointer to [**Expression**](Expression.md) |  | [optional] 
**Identifier** | Pointer to [**Expression**](Expression.md) |  | [optional] 
**MatchType** | Pointer to **string** | Match type of the query term (e.g. exact, prefix). | [optional] 
**IsChildEntityType** | Pointer to **bool** | An indication whether the token is a child entity type. | [optional] 
**AdditionalContext** | Pointer to [**[]Expression**](Expression.md) | Any additional information we have about the token. | [optional] 

## Methods

### NewToken

`func NewToken() *Token`

NewToken instantiates a new Token object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenWithDefaults

`func NewTokenWithDefaults() *Token`

NewTokenWithDefaults instantiates a new Token object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayContext

`func (o *Token) GetDisplayContext() Expression`

GetDisplayContext returns the DisplayContext field if non-nil, zero value otherwise.

### GetDisplayContextOk

`func (o *Token) GetDisplayContextOk() (*Expression, bool)`

GetDisplayContextOk returns a tuple with the DisplayContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayContext

`func (o *Token) SetDisplayContext(v Expression)`

SetDisplayContext sets DisplayContext field to given value.

### HasDisplayContext

`func (o *Token) HasDisplayContext() bool`

HasDisplayContext returns a boolean if a field has been set.

### GetIdentifier

`func (o *Token) GetIdentifier() Expression`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *Token) GetIdentifierOk() (*Expression, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *Token) SetIdentifier(v Expression)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *Token) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetMatchType

`func (o *Token) GetMatchType() string`

GetMatchType returns the MatchType field if non-nil, zero value otherwise.

### GetMatchTypeOk

`func (o *Token) GetMatchTypeOk() (*string, bool)`

GetMatchTypeOk returns a tuple with the MatchType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchType

`func (o *Token) SetMatchType(v string)`

SetMatchType sets MatchType field to given value.

### HasMatchType

`func (o *Token) HasMatchType() bool`

HasMatchType returns a boolean if a field has been set.

### GetIsChildEntityType

`func (o *Token) GetIsChildEntityType() bool`

GetIsChildEntityType returns the IsChildEntityType field if non-nil, zero value otherwise.

### GetIsChildEntityTypeOk

`func (o *Token) GetIsChildEntityTypeOk() (*bool, bool)`

GetIsChildEntityTypeOk returns a tuple with the IsChildEntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsChildEntityType

`func (o *Token) SetIsChildEntityType(v bool)`

SetIsChildEntityType sets IsChildEntityType field to given value.

### HasIsChildEntityType

`func (o *Token) HasIsChildEntityType() bool`

HasIsChildEntityType returns a boolean if a field has been set.

### GetAdditionalContext

`func (o *Token) GetAdditionalContext() []Expression`

GetAdditionalContext returns the AdditionalContext field if non-nil, zero value otherwise.

### GetAdditionalContextOk

`func (o *Token) GetAdditionalContextOk() (*[]Expression, bool)`

GetAdditionalContextOk returns a tuple with the AdditionalContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalContext

`func (o *Token) SetAdditionalContext(v []Expression)`

SetAdditionalContext sets AdditionalContext field to given value.

### HasAdditionalContext

`func (o *Token) HasAdditionalContext() bool`

HasAdditionalContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



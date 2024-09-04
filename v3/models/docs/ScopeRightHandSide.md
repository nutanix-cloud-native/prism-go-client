# ScopeRightHandSide

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ValueList** | Pointer to **[]string** | The explicit list of values for the given kind. | [optional] 
**Categories** | Pointer to **map[string][]string** | The category values represented as a dictionary of key -&gt; list of values. e.g.{\&quot;env\&quot;:[\&quot;env1\&quot;, \&quot;env2\&quot;]}  | [optional] 
**Collection** | Pointer to **string** | A representative term for supported groupings of entities. ALL &#x3D; All the entities of a given kind.  | [optional] 
**UuidList** | Pointer to **[]string** | The explicit list of UUIDs for the given kind. | [optional] 

## Methods

### NewScopeRightHandSide

`func NewScopeRightHandSide() *ScopeRightHandSide`

NewScopeRightHandSide instantiates a new ScopeRightHandSide object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScopeRightHandSideWithDefaults

`func NewScopeRightHandSideWithDefaults() *ScopeRightHandSide`

NewScopeRightHandSideWithDefaults instantiates a new ScopeRightHandSide object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValueList

`func (o *ScopeRightHandSide) GetValueList() []string`

GetValueList returns the ValueList field if non-nil, zero value otherwise.

### GetValueListOk

`func (o *ScopeRightHandSide) GetValueListOk() (*[]string, bool)`

GetValueListOk returns a tuple with the ValueList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueList

`func (o *ScopeRightHandSide) SetValueList(v []string)`

SetValueList sets ValueList field to given value.

### HasValueList

`func (o *ScopeRightHandSide) HasValueList() bool`

HasValueList returns a boolean if a field has been set.

### GetCategories

`func (o *ScopeRightHandSide) GetCategories() map[string][]string`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *ScopeRightHandSide) GetCategoriesOk() (*map[string][]string, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *ScopeRightHandSide) SetCategories(v map[string][]string)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *ScopeRightHandSide) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetCollection

`func (o *ScopeRightHandSide) GetCollection() string`

GetCollection returns the Collection field if non-nil, zero value otherwise.

### GetCollectionOk

`func (o *ScopeRightHandSide) GetCollectionOk() (*string, bool)`

GetCollectionOk returns a tuple with the Collection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollection

`func (o *ScopeRightHandSide) SetCollection(v string)`

SetCollection sets Collection field to given value.

### HasCollection

`func (o *ScopeRightHandSide) HasCollection() bool`

HasCollection returns a boolean if a field has been set.

### GetUuidList

`func (o *ScopeRightHandSide) GetUuidList() []string`

GetUuidList returns the UuidList field if non-nil, zero value otherwise.

### GetUuidListOk

`func (o *ScopeRightHandSide) GetUuidListOk() (*[]string, bool)`

GetUuidListOk returns a tuple with the UuidList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuidList

`func (o *ScopeRightHandSide) SetUuidList(v []string)`

SetUuidList sets UuidList field to given value.

### HasUuidList

`func (o *ScopeRightHandSide) HasUuidList() bool`

HasUuidList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



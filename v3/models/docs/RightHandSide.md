# RightHandSide

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Categories** | Pointer to **map[string][]string** | The category values represented as a dictionary of key -&gt; list of values. e.g.{\&quot;env\&quot;:[\&quot;env1\&quot;, \&quot;env2\&quot;]}  | [optional] 
**Collection** | Pointer to **string** | A representative term for supported groupings of entities. ALL &#x3D; All the entities of a given kind. SELF_OWNED &#x3D; The entities of a given kind, that are owned by the user.  | [optional] 
**UuidList** | Pointer to **[]string** | The explicit list of UUIDs for the given kind. | [optional] 

## Methods

### NewRightHandSide

`func NewRightHandSide() *RightHandSide`

NewRightHandSide instantiates a new RightHandSide object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRightHandSideWithDefaults

`func NewRightHandSideWithDefaults() *RightHandSide`

NewRightHandSideWithDefaults instantiates a new RightHandSide object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategories

`func (o *RightHandSide) GetCategories() map[string][]string`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *RightHandSide) GetCategoriesOk() (*map[string][]string, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *RightHandSide) SetCategories(v map[string][]string)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *RightHandSide) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetCollection

`func (o *RightHandSide) GetCollection() string`

GetCollection returns the Collection field if non-nil, zero value otherwise.

### GetCollectionOk

`func (o *RightHandSide) GetCollectionOk() (*string, bool)`

GetCollectionOk returns a tuple with the Collection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollection

`func (o *RightHandSide) SetCollection(v string)`

SetCollection sets Collection field to given value.

### HasCollection

`func (o *RightHandSide) HasCollection() bool`

HasCollection returns a boolean if a field has been set.

### GetUuidList

`func (o *RightHandSide) GetUuidList() []string`

GetUuidList returns the UuidList field if non-nil, zero value otherwise.

### GetUuidListOk

`func (o *RightHandSide) GetUuidListOk() (*[]string, bool)`

GetUuidListOk returns a tuple with the UuidList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuidList

`func (o *RightHandSide) SetUuidList(v []string)`

SetUuidList sets UuidList field to given value.

### HasUuidList

`func (o *RightHandSide) HasUuidList() bool`

HasUuidList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



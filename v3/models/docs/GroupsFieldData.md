# GroupsFieldData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Buckets** | Pointer to **map[string]int64** | Group Bucket Summary Map. | [optional] 
**Values** | Pointer to [**[]GroupsTimevaluePair**](GroupsTimevaluePair.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 

## Methods

### NewGroupsFieldData

`func NewGroupsFieldData() *GroupsFieldData`

NewGroupsFieldData instantiates a new GroupsFieldData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupsFieldDataWithDefaults

`func NewGroupsFieldDataWithDefaults() *GroupsFieldData`

NewGroupsFieldDataWithDefaults instantiates a new GroupsFieldData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuckets

`func (o *GroupsFieldData) GetBuckets() map[string]int64`

GetBuckets returns the Buckets field if non-nil, zero value otherwise.

### GetBucketsOk

`func (o *GroupsFieldData) GetBucketsOk() (*map[string]int64, bool)`

GetBucketsOk returns a tuple with the Buckets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuckets

`func (o *GroupsFieldData) SetBuckets(v map[string]int64)`

SetBuckets sets Buckets field to given value.

### HasBuckets

`func (o *GroupsFieldData) HasBuckets() bool`

HasBuckets returns a boolean if a field has been set.

### GetValues

`func (o *GroupsFieldData) GetValues() []GroupsTimevaluePair`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *GroupsFieldData) GetValuesOk() (*[]GroupsTimevaluePair, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *GroupsFieldData) SetValues(v []GroupsTimevaluePair)`

SetValues sets Values field to given value.

### HasValues

`func (o *GroupsFieldData) HasValues() bool`

HasValues returns a boolean if a field has been set.

### GetName

`func (o *GroupsFieldData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GroupsFieldData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GroupsFieldData) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GroupsFieldData) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# GroupsEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityId** | Pointer to **string** |  | [optional] 
**Data** | Pointer to [**[]GroupsFieldData**](GroupsFieldData.md) |  | [optional] 

## Methods

### NewGroupsEntity

`func NewGroupsEntity() *GroupsEntity`

NewGroupsEntity instantiates a new GroupsEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupsEntityWithDefaults

`func NewGroupsEntityWithDefaults() *GroupsEntity`

NewGroupsEntityWithDefaults instantiates a new GroupsEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityId

`func (o *GroupsEntity) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *GroupsEntity) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *GroupsEntity) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *GroupsEntity) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### GetData

`func (o *GroupsEntity) GetData() []GroupsFieldData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GroupsEntity) GetDataOk() (*[]GroupsFieldData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GroupsEntity) SetData(v []GroupsFieldData)`

SetData sets Data field to given value.

### HasData

`func (o *GroupsEntity) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



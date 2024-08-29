# QueryGroupsEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityId** | Pointer to **string** |  | [optional] 
**Data** | Pointer to [**[]QueryGroupsFieldData**](QueryGroupsFieldData.md) |  | [optional] 

## Methods

### NewQueryGroupsEntity

`func NewQueryGroupsEntity() *QueryGroupsEntity`

NewQueryGroupsEntity instantiates a new QueryGroupsEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryGroupsEntityWithDefaults

`func NewQueryGroupsEntityWithDefaults() *QueryGroupsEntity`

NewQueryGroupsEntityWithDefaults instantiates a new QueryGroupsEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityId

`func (o *QueryGroupsEntity) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *QueryGroupsEntity) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *QueryGroupsEntity) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *QueryGroupsEntity) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### GetData

`func (o *QueryGroupsEntity) GetData() []QueryGroupsFieldData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *QueryGroupsEntity) GetDataOk() (*[]QueryGroupsFieldData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *QueryGroupsEntity) SetData(v []QueryGroupsFieldData)`

SetData sets Data field to given value.

### HasData

`func (o *QueryGroupsEntity) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



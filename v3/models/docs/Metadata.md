# Metadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceList** | Pointer to [**[]EntityMetadata**](EntityMetadata.md) | List of entity instances. | [optional] 
**UserQuery** | Pointer to **string** | User query in simple text. | [optional] 
**FilterList** | Pointer to [**[]Expression**](Expression.md) | List of filters. | [optional] 
**EntityType** | Pointer to **string** | Type of the entity. | [optional] 

## Methods

### NewMetadata

`func NewMetadata() *Metadata`

NewMetadata instantiates a new Metadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetadataWithDefaults

`func NewMetadataWithDefaults() *Metadata`

NewMetadataWithDefaults instantiates a new Metadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceList

`func (o *Metadata) GetInstanceList() []EntityMetadata`

GetInstanceList returns the InstanceList field if non-nil, zero value otherwise.

### GetInstanceListOk

`func (o *Metadata) GetInstanceListOk() (*[]EntityMetadata, bool)`

GetInstanceListOk returns a tuple with the InstanceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceList

`func (o *Metadata) SetInstanceList(v []EntityMetadata)`

SetInstanceList sets InstanceList field to given value.

### HasInstanceList

`func (o *Metadata) HasInstanceList() bool`

HasInstanceList returns a boolean if a field has been set.

### GetUserQuery

`func (o *Metadata) GetUserQuery() string`

GetUserQuery returns the UserQuery field if non-nil, zero value otherwise.

### GetUserQueryOk

`func (o *Metadata) GetUserQueryOk() (*string, bool)`

GetUserQueryOk returns a tuple with the UserQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserQuery

`func (o *Metadata) SetUserQuery(v string)`

SetUserQuery sets UserQuery field to given value.

### HasUserQuery

`func (o *Metadata) HasUserQuery() bool`

HasUserQuery returns a boolean if a field has been set.

### GetFilterList

`func (o *Metadata) GetFilterList() []Expression`

GetFilterList returns the FilterList field if non-nil, zero value otherwise.

### GetFilterListOk

`func (o *Metadata) GetFilterListOk() (*[]Expression, bool)`

GetFilterListOk returns a tuple with the FilterList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterList

`func (o *Metadata) SetFilterList(v []Expression)`

SetFilterList sets FilterList field to given value.

### HasFilterList

`func (o *Metadata) HasFilterList() bool`

HasFilterList returns a boolean if a field has been set.

### GetEntityType

`func (o *Metadata) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *Metadata) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *Metadata) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.

### HasEntityType

`func (o *Metadata) HasEntityType() bool`

HasEntityType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



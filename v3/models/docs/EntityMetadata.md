# EntityMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterUuid** | Pointer to **string** | Cluster uuid on which the entity is present. | [optional] 
**EntityId** | Pointer to **string** | Unique identifier of the entity. | [optional] 
**EntityName** | Pointer to **string** | Name of the entity. | [optional] 
**EntityType** | Pointer to **string** | Type of the entity. | [optional] 

## Methods

### NewEntityMetadata

`func NewEntityMetadata() *EntityMetadata`

NewEntityMetadata instantiates a new EntityMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntityMetadataWithDefaults

`func NewEntityMetadataWithDefaults() *EntityMetadata`

NewEntityMetadataWithDefaults instantiates a new EntityMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterUuid

`func (o *EntityMetadata) GetClusterUuid() string`

GetClusterUuid returns the ClusterUuid field if non-nil, zero value otherwise.

### GetClusterUuidOk

`func (o *EntityMetadata) GetClusterUuidOk() (*string, bool)`

GetClusterUuidOk returns a tuple with the ClusterUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterUuid

`func (o *EntityMetadata) SetClusterUuid(v string)`

SetClusterUuid sets ClusterUuid field to given value.

### HasClusterUuid

`func (o *EntityMetadata) HasClusterUuid() bool`

HasClusterUuid returns a boolean if a field has been set.

### GetEntityId

`func (o *EntityMetadata) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *EntityMetadata) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *EntityMetadata) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *EntityMetadata) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### GetEntityName

`func (o *EntityMetadata) GetEntityName() string`

GetEntityName returns the EntityName field if non-nil, zero value otherwise.

### GetEntityNameOk

`func (o *EntityMetadata) GetEntityNameOk() (*string, bool)`

GetEntityNameOk returns a tuple with the EntityName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityName

`func (o *EntityMetadata) SetEntityName(v string)`

SetEntityName sets EntityName field to given value.

### HasEntityName

`func (o *EntityMetadata) HasEntityName() bool`

HasEntityName returns a boolean if a field has been set.

### GetEntityType

`func (o *EntityMetadata) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *EntityMetadata) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *EntityMetadata) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.

### HasEntityType

`func (o *EntityMetadata) HasEntityType() bool`

HasEntityType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



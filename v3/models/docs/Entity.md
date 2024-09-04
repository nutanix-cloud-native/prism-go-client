# Entity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | User or group or OU in the directory service. | [optional] 
**Name** | Pointer to **string** | The name of the entity in canonical format. | [optional] 
**AttributeList** | Pointer to [**[]Attribute**](Attribute.md) |  | [optional] 

## Methods

### NewEntity

`func NewEntity() *Entity`

NewEntity instantiates a new Entity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntityWithDefaults

`func NewEntityWithDefaults() *Entity`

NewEntityWithDefaults instantiates a new Entity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Entity) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Entity) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Entity) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Entity) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *Entity) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Entity) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Entity) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Entity) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAttributeList

`func (o *Entity) GetAttributeList() []Attribute`

GetAttributeList returns the AttributeList field if non-nil, zero value otherwise.

### GetAttributeListOk

`func (o *Entity) GetAttributeListOk() (*[]Attribute, bool)`

GetAttributeListOk returns a tuple with the AttributeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeList

`func (o *Entity) SetAttributeList(v []Attribute)`

SetAttributeList sets AttributeList field to given value.

### HasAttributeList

`func (o *Entity) HasAttributeList() bool`

HasAttributeList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



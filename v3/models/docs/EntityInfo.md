# EntityInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | Entity type | [optional] 
**Name** | Pointer to **string** | Entity name | [optional] 
**UUID** | Pointer to **string** | Entity UUID | [optional] 

## Methods

### NewEntityInfo

`func NewEntityInfo() *EntityInfo`

NewEntityInfo instantiates a new EntityInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntityInfoWithDefaults

`func NewEntityInfoWithDefaults() *EntityInfo`

NewEntityInfoWithDefaults instantiates a new EntityInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *EntityInfo) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EntityInfo) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EntityInfo) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *EntityInfo) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *EntityInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EntityInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EntityInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EntityInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUUID

`func (o *EntityInfo) GetUUID() string`

GetUUID returns the UUID field if non-nil, zero value otherwise.

### GetUUIDOk

`func (o *EntityInfo) GetUUIDOk() (*string, bool)`

GetUUIDOk returns a tuple with the UUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUUID

`func (o *EntityInfo) SetUUID(v string)`

SetUUID sets UUID field to given value.

### HasUUID

`func (o *EntityInfo) HasUUID() bool`

HasUUID returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



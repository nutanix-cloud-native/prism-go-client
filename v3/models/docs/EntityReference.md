# EntityReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | Pointer to **string** | Kind of the reference. | [optional] 
**Type** | Pointer to **string** | The type of filter being used. | [optional] 
**UUID** | Pointer to **string** | UUID of the entity. | [optional] 
**Categories** | Pointer to **map[string]string** | Categories for the entity. | [optional] 
**Name** | Pointer to **string** | Name of the entity. | [optional] 

## Methods

### NewEntityReference

`func NewEntityReference() *EntityReference`

NewEntityReference instantiates a new EntityReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntityReferenceWithDefaults

`func NewEntityReferenceWithDefaults() *EntityReference`

NewEntityReferenceWithDefaults instantiates a new EntityReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *EntityReference) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *EntityReference) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *EntityReference) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *EntityReference) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetType

`func (o *EntityReference) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EntityReference) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EntityReference) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *EntityReference) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUUID

`func (o *EntityReference) GetUUID() string`

GetUUID returns the UUID field if non-nil, zero value otherwise.

### GetUUIDOk

`func (o *EntityReference) GetUUIDOk() (*string, bool)`

GetUUIDOk returns a tuple with the UUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUUID

`func (o *EntityReference) SetUUID(v string)`

SetUUID sets UUID field to given value.

### HasUUID

`func (o *EntityReference) HasUUID() bool`

HasUUID returns a boolean if a field has been set.

### GetCategories

`func (o *EntityReference) GetCategories() map[string]string`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *EntityReference) GetCategoriesOk() (*map[string]string, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *EntityReference) SetCategories(v map[string]string)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *EntityReference) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetName

`func (o *EntityReference) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EntityReference) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EntityReference) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EntityReference) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



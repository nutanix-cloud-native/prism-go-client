# IdentifierResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NameUuidList** | Pointer to [**[]NameIdentifierMap**](NameIdentifierMap.md) | The list of name to salted UUID5 mapping(s). | [optional] 

## Methods

### NewIdentifierResponse

`func NewIdentifierResponse() *IdentifierResponse`

NewIdentifierResponse instantiates a new IdentifierResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentifierResponseWithDefaults

`func NewIdentifierResponseWithDefaults() *IdentifierResponse`

NewIdentifierResponseWithDefaults instantiates a new IdentifierResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNameUuidList

`func (o *IdentifierResponse) GetNameUuidList() []NameIdentifierMap`

GetNameUuidList returns the NameUuidList field if non-nil, zero value otherwise.

### GetNameUuidListOk

`func (o *IdentifierResponse) GetNameUuidListOk() (*[]NameIdentifierMap, bool)`

GetNameUuidListOk returns a tuple with the NameUuidList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameUuidList

`func (o *IdentifierResponse) SetNameUuidList(v []NameIdentifierMap)`

SetNameUuidList sets NameUuidList field to given value.

### HasNameUuidList

`func (o *IdentifierResponse) HasNameUuidList() bool`

HasNameUuidList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



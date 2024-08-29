# EntitySyncInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShouldOverride** | Pointer to **bool** | Indicates whether to override entities in case of conflicts.  | [optional] 

## Methods

### NewEntitySyncInput

`func NewEntitySyncInput() *EntitySyncInput`

NewEntitySyncInput instantiates a new EntitySyncInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitySyncInputWithDefaults

`func NewEntitySyncInputWithDefaults() *EntitySyncInput`

NewEntitySyncInputWithDefaults instantiates a new EntitySyncInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShouldOverride

`func (o *EntitySyncInput) GetShouldOverride() bool`

GetShouldOverride returns the ShouldOverride field if non-nil, zero value otherwise.

### GetShouldOverrideOk

`func (o *EntitySyncInput) GetShouldOverrideOk() (*bool, bool)`

GetShouldOverrideOk returns a tuple with the ShouldOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldOverride

`func (o *EntitySyncInput) SetShouldOverride(v bool)`

SetShouldOverride sets ShouldOverride field to given value.

### HasShouldOverride

`func (o *EntitySyncInput) HasShouldOverride() bool`

HasShouldOverride returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



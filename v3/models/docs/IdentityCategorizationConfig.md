# IdentityCategorizationConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsDefaultCategoryEnabled** | Pointer to **bool** | Enablement status of the default category. | [optional] [default to false]
**KeepDefaultCategoryOnLogin** | Pointer to **bool** | Retain default category on user login. | [optional] [default to false]
**DirectoryServiceConfigList** | Pointer to [**[]IdentityCategorizationDirectoryConfig**](IdentityCategorizationDirectoryConfig.md) | Directory service specific config. | [optional] 

## Methods

### NewIdentityCategorizationConfig

`func NewIdentityCategorizationConfig() *IdentityCategorizationConfig`

NewIdentityCategorizationConfig instantiates a new IdentityCategorizationConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityCategorizationConfigWithDefaults

`func NewIdentityCategorizationConfigWithDefaults() *IdentityCategorizationConfig`

NewIdentityCategorizationConfigWithDefaults instantiates a new IdentityCategorizationConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsDefaultCategoryEnabled

`func (o *IdentityCategorizationConfig) GetIsDefaultCategoryEnabled() bool`

GetIsDefaultCategoryEnabled returns the IsDefaultCategoryEnabled field if non-nil, zero value otherwise.

### GetIsDefaultCategoryEnabledOk

`func (o *IdentityCategorizationConfig) GetIsDefaultCategoryEnabledOk() (*bool, bool)`

GetIsDefaultCategoryEnabledOk returns a tuple with the IsDefaultCategoryEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefaultCategoryEnabled

`func (o *IdentityCategorizationConfig) SetIsDefaultCategoryEnabled(v bool)`

SetIsDefaultCategoryEnabled sets IsDefaultCategoryEnabled field to given value.

### HasIsDefaultCategoryEnabled

`func (o *IdentityCategorizationConfig) HasIsDefaultCategoryEnabled() bool`

HasIsDefaultCategoryEnabled returns a boolean if a field has been set.

### GetKeepDefaultCategoryOnLogin

`func (o *IdentityCategorizationConfig) GetKeepDefaultCategoryOnLogin() bool`

GetKeepDefaultCategoryOnLogin returns the KeepDefaultCategoryOnLogin field if non-nil, zero value otherwise.

### GetKeepDefaultCategoryOnLoginOk

`func (o *IdentityCategorizationConfig) GetKeepDefaultCategoryOnLoginOk() (*bool, bool)`

GetKeepDefaultCategoryOnLoginOk returns a tuple with the KeepDefaultCategoryOnLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepDefaultCategoryOnLogin

`func (o *IdentityCategorizationConfig) SetKeepDefaultCategoryOnLogin(v bool)`

SetKeepDefaultCategoryOnLogin sets KeepDefaultCategoryOnLogin field to given value.

### HasKeepDefaultCategoryOnLogin

`func (o *IdentityCategorizationConfig) HasKeepDefaultCategoryOnLogin() bool`

HasKeepDefaultCategoryOnLogin returns a boolean if a field has been set.

### GetDirectoryServiceConfigList

`func (o *IdentityCategorizationConfig) GetDirectoryServiceConfigList() []IdentityCategorizationDirectoryConfig`

GetDirectoryServiceConfigList returns the DirectoryServiceConfigList field if non-nil, zero value otherwise.

### GetDirectoryServiceConfigListOk

`func (o *IdentityCategorizationConfig) GetDirectoryServiceConfigListOk() (*[]IdentityCategorizationDirectoryConfig, bool)`

GetDirectoryServiceConfigListOk returns a tuple with the DirectoryServiceConfigList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryServiceConfigList

`func (o *IdentityCategorizationConfig) SetDirectoryServiceConfigList(v []IdentityCategorizationDirectoryConfig)`

SetDirectoryServiceConfigList sets DirectoryServiceConfigList field to given value.

### HasDirectoryServiceConfigList

`func (o *IdentityCategorizationConfig) HasDirectoryServiceConfigList() bool`

HasDirectoryServiceConfigList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



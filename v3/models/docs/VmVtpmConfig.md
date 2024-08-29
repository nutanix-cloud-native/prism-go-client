# VmVtpmConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VtpmEnabled** | Pointer to **bool** | Indicates whether virtual trusted platform module should be enabled for the Guest OS.  | [optional] 
**DataSourceReference** | Pointer to [**Reference**](Reference.md) |  | [optional] 
**VtpmSecret** | Pointer to **string** | Virtual trusted platform module secret. | [optional] 

## Methods

### NewVmVtpmConfig

`func NewVmVtpmConfig() *VmVtpmConfig`

NewVmVtpmConfig instantiates a new VmVtpmConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmVtpmConfigWithDefaults

`func NewVmVtpmConfigWithDefaults() *VmVtpmConfig`

NewVmVtpmConfigWithDefaults instantiates a new VmVtpmConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVtpmEnabled

`func (o *VmVtpmConfig) GetVtpmEnabled() bool`

GetVtpmEnabled returns the VtpmEnabled field if non-nil, zero value otherwise.

### GetVtpmEnabledOk

`func (o *VmVtpmConfig) GetVtpmEnabledOk() (*bool, bool)`

GetVtpmEnabledOk returns a tuple with the VtpmEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVtpmEnabled

`func (o *VmVtpmConfig) SetVtpmEnabled(v bool)`

SetVtpmEnabled sets VtpmEnabled field to given value.

### HasVtpmEnabled

`func (o *VmVtpmConfig) HasVtpmEnabled() bool`

HasVtpmEnabled returns a boolean if a field has been set.

### GetDataSourceReference

`func (o *VmVtpmConfig) GetDataSourceReference() Reference`

GetDataSourceReference returns the DataSourceReference field if non-nil, zero value otherwise.

### GetDataSourceReferenceOk

`func (o *VmVtpmConfig) GetDataSourceReferenceOk() (*Reference, bool)`

GetDataSourceReferenceOk returns a tuple with the DataSourceReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSourceReference

`func (o *VmVtpmConfig) SetDataSourceReference(v Reference)`

SetDataSourceReference sets DataSourceReference field to given value.

### HasDataSourceReference

`func (o *VmVtpmConfig) HasDataSourceReference() bool`

HasDataSourceReference returns a boolean if a field has been set.

### GetVtpmSecret

`func (o *VmVtpmConfig) GetVtpmSecret() string`

GetVtpmSecret returns the VtpmSecret field if non-nil, zero value otherwise.

### GetVtpmSecretOk

`func (o *VmVtpmConfig) GetVtpmSecretOk() (*string, bool)`

GetVtpmSecretOk returns a tuple with the VtpmSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVtpmSecret

`func (o *VmVtpmConfig) SetVtpmSecret(v string)`

SetVtpmSecret sets VtpmSecret field to given value.

### HasVtpmSecret

`func (o *VmVtpmConfig) HasVtpmSecret() bool`

HasVtpmSecret returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



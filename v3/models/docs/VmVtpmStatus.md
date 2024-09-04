# VmVtpmStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VtpmEnabled** | Pointer to **bool** | Indicates whether virtual trusted platform module is enabled for the the Guest OS.  | [optional] 
**DataSourceReference** | Pointer to [**Reference**](Reference.md) |  | [optional] 
**Version** | Pointer to **string** | Virtual trusted platform module version. | [optional] 

## Methods

### NewVmVtpmStatus

`func NewVmVtpmStatus() *VmVtpmStatus`

NewVmVtpmStatus instantiates a new VmVtpmStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmVtpmStatusWithDefaults

`func NewVmVtpmStatusWithDefaults() *VmVtpmStatus`

NewVmVtpmStatusWithDefaults instantiates a new VmVtpmStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVtpmEnabled

`func (o *VmVtpmStatus) GetVtpmEnabled() bool`

GetVtpmEnabled returns the VtpmEnabled field if non-nil, zero value otherwise.

### GetVtpmEnabledOk

`func (o *VmVtpmStatus) GetVtpmEnabledOk() (*bool, bool)`

GetVtpmEnabledOk returns a tuple with the VtpmEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVtpmEnabled

`func (o *VmVtpmStatus) SetVtpmEnabled(v bool)`

SetVtpmEnabled sets VtpmEnabled field to given value.

### HasVtpmEnabled

`func (o *VmVtpmStatus) HasVtpmEnabled() bool`

HasVtpmEnabled returns a boolean if a field has been set.

### GetDataSourceReference

`func (o *VmVtpmStatus) GetDataSourceReference() Reference`

GetDataSourceReference returns the DataSourceReference field if non-nil, zero value otherwise.

### GetDataSourceReferenceOk

`func (o *VmVtpmStatus) GetDataSourceReferenceOk() (*Reference, bool)`

GetDataSourceReferenceOk returns a tuple with the DataSourceReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSourceReference

`func (o *VmVtpmStatus) SetDataSourceReference(v Reference)`

SetDataSourceReference sets DataSourceReference field to given value.

### HasDataSourceReference

`func (o *VmVtpmStatus) HasDataSourceReference() bool`

HasDataSourceReference returns a boolean if a field has been set.

### GetVersion

`func (o *VmVtpmStatus) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *VmVtpmStatus) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *VmVtpmStatus) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *VmVtpmStatus) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



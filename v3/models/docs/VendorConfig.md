# VendorConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VersionList** | Pointer to **[]string** | Vendor device version list. | [optional] 
**Name** | Pointer to **string** | Vendor device name. | [optional] 

## Methods

### NewVendorConfig

`func NewVendorConfig() *VendorConfig`

NewVendorConfig instantiates a new VendorConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVendorConfigWithDefaults

`func NewVendorConfigWithDefaults() *VendorConfig`

NewVendorConfigWithDefaults instantiates a new VendorConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersionList

`func (o *VendorConfig) GetVersionList() []string`

GetVersionList returns the VersionList field if non-nil, zero value otherwise.

### GetVersionListOk

`func (o *VendorConfig) GetVersionListOk() (*[]string, bool)`

GetVersionListOk returns a tuple with the VersionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionList

`func (o *VendorConfig) SetVersionList(v []string)`

SetVersionList sets VersionList field to given value.

### HasVersionList

`func (o *VendorConfig) HasVersionList() bool`

HasVersionList returns a boolean if a field has been set.

### GetName

`func (o *VendorConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VendorConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VendorConfig) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VendorConfig) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



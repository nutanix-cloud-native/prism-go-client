# XenWorkload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OperatingSystem** | Pointer to **string** |  | [optional] 
**Vendor** | Pointer to **string** |  | [optional] 
**SystemData** | Pointer to **int32** |  | [optional] 
**RdshProvisioningType** | Pointer to **string** |  | [optional] 
**NumUsers** | Pointer to **int32** |  | [optional] 
**PvsWriteCacheSize** | Pointer to **int32** |  | [optional] 
**McsDiffSize** | Pointer to **int32** |  | [optional] 
**UserProfileData** | Pointer to **int32** |  | [optional] 

## Methods

### NewXenWorkload

`func NewXenWorkload() *XenWorkload`

NewXenWorkload instantiates a new XenWorkload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewXenWorkloadWithDefaults

`func NewXenWorkloadWithDefaults() *XenWorkload`

NewXenWorkloadWithDefaults instantiates a new XenWorkload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperatingSystem

`func (o *XenWorkload) GetOperatingSystem() string`

GetOperatingSystem returns the OperatingSystem field if non-nil, zero value otherwise.

### GetOperatingSystemOk

`func (o *XenWorkload) GetOperatingSystemOk() (*string, bool)`

GetOperatingSystemOk returns a tuple with the OperatingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystem

`func (o *XenWorkload) SetOperatingSystem(v string)`

SetOperatingSystem sets OperatingSystem field to given value.

### HasOperatingSystem

`func (o *XenWorkload) HasOperatingSystem() bool`

HasOperatingSystem returns a boolean if a field has been set.

### GetVendor

`func (o *XenWorkload) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *XenWorkload) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *XenWorkload) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *XenWorkload) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetSystemData

`func (o *XenWorkload) GetSystemData() int32`

GetSystemData returns the SystemData field if non-nil, zero value otherwise.

### GetSystemDataOk

`func (o *XenWorkload) GetSystemDataOk() (*int32, bool)`

GetSystemDataOk returns a tuple with the SystemData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemData

`func (o *XenWorkload) SetSystemData(v int32)`

SetSystemData sets SystemData field to given value.

### HasSystemData

`func (o *XenWorkload) HasSystemData() bool`

HasSystemData returns a boolean if a field has been set.

### GetRdshProvisioningType

`func (o *XenWorkload) GetRdshProvisioningType() string`

GetRdshProvisioningType returns the RdshProvisioningType field if non-nil, zero value otherwise.

### GetRdshProvisioningTypeOk

`func (o *XenWorkload) GetRdshProvisioningTypeOk() (*string, bool)`

GetRdshProvisioningTypeOk returns a tuple with the RdshProvisioningType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdshProvisioningType

`func (o *XenWorkload) SetRdshProvisioningType(v string)`

SetRdshProvisioningType sets RdshProvisioningType field to given value.

### HasRdshProvisioningType

`func (o *XenWorkload) HasRdshProvisioningType() bool`

HasRdshProvisioningType returns a boolean if a field has been set.

### GetNumUsers

`func (o *XenWorkload) GetNumUsers() int32`

GetNumUsers returns the NumUsers field if non-nil, zero value otherwise.

### GetNumUsersOk

`func (o *XenWorkload) GetNumUsersOk() (*int32, bool)`

GetNumUsersOk returns a tuple with the NumUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumUsers

`func (o *XenWorkload) SetNumUsers(v int32)`

SetNumUsers sets NumUsers field to given value.

### HasNumUsers

`func (o *XenWorkload) HasNumUsers() bool`

HasNumUsers returns a boolean if a field has been set.

### GetPvsWriteCacheSize

`func (o *XenWorkload) GetPvsWriteCacheSize() int32`

GetPvsWriteCacheSize returns the PvsWriteCacheSize field if non-nil, zero value otherwise.

### GetPvsWriteCacheSizeOk

`func (o *XenWorkload) GetPvsWriteCacheSizeOk() (*int32, bool)`

GetPvsWriteCacheSizeOk returns a tuple with the PvsWriteCacheSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPvsWriteCacheSize

`func (o *XenWorkload) SetPvsWriteCacheSize(v int32)`

SetPvsWriteCacheSize sets PvsWriteCacheSize field to given value.

### HasPvsWriteCacheSize

`func (o *XenWorkload) HasPvsWriteCacheSize() bool`

HasPvsWriteCacheSize returns a boolean if a field has been set.

### GetMcsDiffSize

`func (o *XenWorkload) GetMcsDiffSize() int32`

GetMcsDiffSize returns the McsDiffSize field if non-nil, zero value otherwise.

### GetMcsDiffSizeOk

`func (o *XenWorkload) GetMcsDiffSizeOk() (*int32, bool)`

GetMcsDiffSizeOk returns a tuple with the McsDiffSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcsDiffSize

`func (o *XenWorkload) SetMcsDiffSize(v int32)`

SetMcsDiffSize sets McsDiffSize field to given value.

### HasMcsDiffSize

`func (o *XenWorkload) HasMcsDiffSize() bool`

HasMcsDiffSize returns a boolean if a field has been set.

### GetUserProfileData

`func (o *XenWorkload) GetUserProfileData() int32`

GetUserProfileData returns the UserProfileData field if non-nil, zero value otherwise.

### GetUserProfileDataOk

`func (o *XenWorkload) GetUserProfileDataOk() (*int32, bool)`

GetUserProfileDataOk returns a tuple with the UserProfileData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserProfileData

`func (o *XenWorkload) SetUserProfileData(v int32)`

SetUserProfileData sets UserProfileData field to given value.

### HasUserProfileData

`func (o *XenWorkload) HasUserProfileData() bool`

HasUserProfileData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# PortalSoftware

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReleaseNoteUrl** | Pointer to **string** | URL to point to the support portal release note of this software. Currently only set and used for NOS releases  | [optional] 
**UpgradeNotification** | Pointer to [**UpgradeNotification**](UpgradeNotification.md) |  | [optional] 
**ReleaseDate** | Pointer to **time.Time** | Release date of this software in RFC3339 format.  | [optional] 
**Md5sum** | Pointer to **string** | MD5 checksum of the software file | [optional] 
**CompatibleVersionList** | Pointer to **[]string** | List of software versions that this version can be upgraded from  | [optional] 
**Version** | Pointer to **string** | Software version string | [optional] 
**CompatiblePeVersionList** | Pointer to **[]string** | List of Prism Element compatible versions | [optional] 
**SoftwareType** | Pointer to **string** | Software type | [optional] 
**SizeBytes** | Pointer to **int64** | Total size of the software file in bytes | [optional] 

## Methods

### NewPortalSoftware

`func NewPortalSoftware() *PortalSoftware`

NewPortalSoftware instantiates a new PortalSoftware object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortalSoftwareWithDefaults

`func NewPortalSoftwareWithDefaults() *PortalSoftware`

NewPortalSoftwareWithDefaults instantiates a new PortalSoftware object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReleaseNoteUrl

`func (o *PortalSoftware) GetReleaseNoteUrl() string`

GetReleaseNoteUrl returns the ReleaseNoteUrl field if non-nil, zero value otherwise.

### GetReleaseNoteUrlOk

`func (o *PortalSoftware) GetReleaseNoteUrlOk() (*string, bool)`

GetReleaseNoteUrlOk returns a tuple with the ReleaseNoteUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseNoteUrl

`func (o *PortalSoftware) SetReleaseNoteUrl(v string)`

SetReleaseNoteUrl sets ReleaseNoteUrl field to given value.

### HasReleaseNoteUrl

`func (o *PortalSoftware) HasReleaseNoteUrl() bool`

HasReleaseNoteUrl returns a boolean if a field has been set.

### GetUpgradeNotification

`func (o *PortalSoftware) GetUpgradeNotification() UpgradeNotification`

GetUpgradeNotification returns the UpgradeNotification field if non-nil, zero value otherwise.

### GetUpgradeNotificationOk

`func (o *PortalSoftware) GetUpgradeNotificationOk() (*UpgradeNotification, bool)`

GetUpgradeNotificationOk returns a tuple with the UpgradeNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeNotification

`func (o *PortalSoftware) SetUpgradeNotification(v UpgradeNotification)`

SetUpgradeNotification sets UpgradeNotification field to given value.

### HasUpgradeNotification

`func (o *PortalSoftware) HasUpgradeNotification() bool`

HasUpgradeNotification returns a boolean if a field has been set.

### GetReleaseDate

`func (o *PortalSoftware) GetReleaseDate() time.Time`

GetReleaseDate returns the ReleaseDate field if non-nil, zero value otherwise.

### GetReleaseDateOk

`func (o *PortalSoftware) GetReleaseDateOk() (*time.Time, bool)`

GetReleaseDateOk returns a tuple with the ReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseDate

`func (o *PortalSoftware) SetReleaseDate(v time.Time)`

SetReleaseDate sets ReleaseDate field to given value.

### HasReleaseDate

`func (o *PortalSoftware) HasReleaseDate() bool`

HasReleaseDate returns a boolean if a field has been set.

### GetMd5sum

`func (o *PortalSoftware) GetMd5sum() string`

GetMd5sum returns the Md5sum field if non-nil, zero value otherwise.

### GetMd5sumOk

`func (o *PortalSoftware) GetMd5sumOk() (*string, bool)`

GetMd5sumOk returns a tuple with the Md5sum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMd5sum

`func (o *PortalSoftware) SetMd5sum(v string)`

SetMd5sum sets Md5sum field to given value.

### HasMd5sum

`func (o *PortalSoftware) HasMd5sum() bool`

HasMd5sum returns a boolean if a field has been set.

### GetCompatibleVersionList

`func (o *PortalSoftware) GetCompatibleVersionList() []string`

GetCompatibleVersionList returns the CompatibleVersionList field if non-nil, zero value otherwise.

### GetCompatibleVersionListOk

`func (o *PortalSoftware) GetCompatibleVersionListOk() (*[]string, bool)`

GetCompatibleVersionListOk returns a tuple with the CompatibleVersionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompatibleVersionList

`func (o *PortalSoftware) SetCompatibleVersionList(v []string)`

SetCompatibleVersionList sets CompatibleVersionList field to given value.

### HasCompatibleVersionList

`func (o *PortalSoftware) HasCompatibleVersionList() bool`

HasCompatibleVersionList returns a boolean if a field has been set.

### GetVersion

`func (o *PortalSoftware) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *PortalSoftware) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *PortalSoftware) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *PortalSoftware) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetCompatiblePeVersionList

`func (o *PortalSoftware) GetCompatiblePeVersionList() []string`

GetCompatiblePeVersionList returns the CompatiblePeVersionList field if non-nil, zero value otherwise.

### GetCompatiblePeVersionListOk

`func (o *PortalSoftware) GetCompatiblePeVersionListOk() (*[]string, bool)`

GetCompatiblePeVersionListOk returns a tuple with the CompatiblePeVersionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompatiblePeVersionList

`func (o *PortalSoftware) SetCompatiblePeVersionList(v []string)`

SetCompatiblePeVersionList sets CompatiblePeVersionList field to given value.

### HasCompatiblePeVersionList

`func (o *PortalSoftware) HasCompatiblePeVersionList() bool`

HasCompatiblePeVersionList returns a boolean if a field has been set.

### GetSoftwareType

`func (o *PortalSoftware) GetSoftwareType() string`

GetSoftwareType returns the SoftwareType field if non-nil, zero value otherwise.

### GetSoftwareTypeOk

`func (o *PortalSoftware) GetSoftwareTypeOk() (*string, bool)`

GetSoftwareTypeOk returns a tuple with the SoftwareType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareType

`func (o *PortalSoftware) SetSoftwareType(v string)`

SetSoftwareType sets SoftwareType field to given value.

### HasSoftwareType

`func (o *PortalSoftware) HasSoftwareType() bool`

HasSoftwareType returns a boolean if a field has been set.

### GetSizeBytes

`func (o *PortalSoftware) GetSizeBytes() int64`

GetSizeBytes returns the SizeBytes field if non-nil, zero value otherwise.

### GetSizeBytesOk

`func (o *PortalSoftware) GetSizeBytesOk() (*int64, bool)`

GetSizeBytesOk returns a tuple with the SizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeBytes

`func (o *PortalSoftware) SetSizeBytes(v int64)`

SetSizeBytes sets SizeBytes field to given value.

### HasSizeBytes

`func (o *PortalSoftware) HasSizeBytes() bool`

HasSizeBytes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# BuildInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommitId** | **string** | Last Git commit id which the build is based on. | 
**FullVersion** | Pointer to **string** | Full version name. | [optional] 
**CommitDate** | Pointer to **time.Time** | Date/time of the last commit. | [optional] 
**IsLongTermSupport** | Pointer to **bool** | Flag to indicate if the AOS release is qualified as long term support.  | [optional] 
**Version** | **string** | Numeric version. e.g. \&quot;5.5\&quot; | 
**ShortCommitId** | **string** | First 6 characters of the last Git commit id. | 
**BuildType** | **string** | Build type, one of {dbg, opt, release}. | 

## Methods

### NewBuildInfo

`func NewBuildInfo(commitId string, version string, shortCommitId string, buildType string, ) *BuildInfo`

NewBuildInfo instantiates a new BuildInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuildInfoWithDefaults

`func NewBuildInfoWithDefaults() *BuildInfo`

NewBuildInfoWithDefaults instantiates a new BuildInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommitId

`func (o *BuildInfo) GetCommitId() string`

GetCommitId returns the CommitId field if non-nil, zero value otherwise.

### GetCommitIdOk

`func (o *BuildInfo) GetCommitIdOk() (*string, bool)`

GetCommitIdOk returns a tuple with the CommitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitId

`func (o *BuildInfo) SetCommitId(v string)`

SetCommitId sets CommitId field to given value.


### GetFullVersion

`func (o *BuildInfo) GetFullVersion() string`

GetFullVersion returns the FullVersion field if non-nil, zero value otherwise.

### GetFullVersionOk

`func (o *BuildInfo) GetFullVersionOk() (*string, bool)`

GetFullVersionOk returns a tuple with the FullVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullVersion

`func (o *BuildInfo) SetFullVersion(v string)`

SetFullVersion sets FullVersion field to given value.

### HasFullVersion

`func (o *BuildInfo) HasFullVersion() bool`

HasFullVersion returns a boolean if a field has been set.

### GetCommitDate

`func (o *BuildInfo) GetCommitDate() time.Time`

GetCommitDate returns the CommitDate field if non-nil, zero value otherwise.

### GetCommitDateOk

`func (o *BuildInfo) GetCommitDateOk() (*time.Time, bool)`

GetCommitDateOk returns a tuple with the CommitDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitDate

`func (o *BuildInfo) SetCommitDate(v time.Time)`

SetCommitDate sets CommitDate field to given value.

### HasCommitDate

`func (o *BuildInfo) HasCommitDate() bool`

HasCommitDate returns a boolean if a field has been set.

### GetIsLongTermSupport

`func (o *BuildInfo) GetIsLongTermSupport() bool`

GetIsLongTermSupport returns the IsLongTermSupport field if non-nil, zero value otherwise.

### GetIsLongTermSupportOk

`func (o *BuildInfo) GetIsLongTermSupportOk() (*bool, bool)`

GetIsLongTermSupportOk returns a tuple with the IsLongTermSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLongTermSupport

`func (o *BuildInfo) SetIsLongTermSupport(v bool)`

SetIsLongTermSupport sets IsLongTermSupport field to given value.

### HasIsLongTermSupport

`func (o *BuildInfo) HasIsLongTermSupport() bool`

HasIsLongTermSupport returns a boolean if a field has been set.

### GetVersion

`func (o *BuildInfo) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *BuildInfo) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *BuildInfo) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetShortCommitId

`func (o *BuildInfo) GetShortCommitId() string`

GetShortCommitId returns the ShortCommitId field if non-nil, zero value otherwise.

### GetShortCommitIdOk

`func (o *BuildInfo) GetShortCommitIdOk() (*string, bool)`

GetShortCommitIdOk returns a tuple with the ShortCommitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortCommitId

`func (o *BuildInfo) SetShortCommitId(v string)`

SetShortCommitId sets ShortCommitId field to given value.


### GetBuildType

`func (o *BuildInfo) GetBuildType() string`

GetBuildType returns the BuildType field if non-nil, zero value otherwise.

### GetBuildTypeOk

`func (o *BuildInfo) GetBuildTypeOk() (*string, bool)`

GetBuildTypeOk returns a tuple with the BuildType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildType

`func (o *BuildInfo) SetBuildType(v string)`

SetBuildType sets BuildType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



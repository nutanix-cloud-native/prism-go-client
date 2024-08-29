# EpochConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AocUrl** | Pointer to **string** | IP address or Fully Qualified Domain Name(FQDN) obtained by user on deploying Epoch AOC.  | [optional] 
**OrgId** | Pointer to **string** | Organization ID obtained by user on deploying Epoch AOC.  | [optional] 
**AppKey** | Pointer to **string** | Application Key obtained by user on deploying Epoch AOC.  | [optional] 

## Methods

### NewEpochConfig

`func NewEpochConfig() *EpochConfig`

NewEpochConfig instantiates a new EpochConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEpochConfigWithDefaults

`func NewEpochConfigWithDefaults() *EpochConfig`

NewEpochConfigWithDefaults instantiates a new EpochConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAocUrl

`func (o *EpochConfig) GetAocUrl() string`

GetAocUrl returns the AocUrl field if non-nil, zero value otherwise.

### GetAocUrlOk

`func (o *EpochConfig) GetAocUrlOk() (*string, bool)`

GetAocUrlOk returns a tuple with the AocUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAocUrl

`func (o *EpochConfig) SetAocUrl(v string)`

SetAocUrl sets AocUrl field to given value.

### HasAocUrl

`func (o *EpochConfig) HasAocUrl() bool`

HasAocUrl returns a boolean if a field has been set.

### GetOrgId

`func (o *EpochConfig) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *EpochConfig) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *EpochConfig) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *EpochConfig) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetAppKey

`func (o *EpochConfig) GetAppKey() string`

GetAppKey returns the AppKey field if non-nil, zero value otherwise.

### GetAppKeyOk

`func (o *EpochConfig) GetAppKeyOk() (*string, bool)`

GetAppKeyOk returns a tuple with the AppKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppKey

`func (o *EpochConfig) SetAppKey(v string)`

SetAppKey sets AppKey field to given value.

### HasAppKey

`func (o *EpochConfig) HasAppKey() bool`

HasAppKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



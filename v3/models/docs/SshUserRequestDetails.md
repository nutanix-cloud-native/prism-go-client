# SshUserRequestDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | **string** | SSH User name | 
**EntityList** | [**[]SshUserRequestDetailsEntityListInner**](SshUserRequestDetailsEntityListInner.md) | List containing entity IP and entity type for each entity | 
**UserKeyName** | **string** | Name for the user key to add | 
**UserKey** | **string** | key to add for the user | 
**UserUnixId** | Pointer to **int32** | Linux UID for user | [optional] 
**AccessRequestUuid** | **string** | UUID of the access request requesting SSH access | 
**UserUuid** | **string** | SSH User UUID | 

## Methods

### NewSshUserRequestDetails

`func NewSshUserRequestDetails(username string, entityList []SshUserRequestDetailsEntityListInner, userKeyName string, userKey string, accessRequestUuid string, userUuid string, ) *SshUserRequestDetails`

NewSshUserRequestDetails instantiates a new SshUserRequestDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSshUserRequestDetailsWithDefaults

`func NewSshUserRequestDetailsWithDefaults() *SshUserRequestDetails`

NewSshUserRequestDetailsWithDefaults instantiates a new SshUserRequestDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *SshUserRequestDetails) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *SshUserRequestDetails) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *SshUserRequestDetails) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetEntityList

`func (o *SshUserRequestDetails) GetEntityList() []SshUserRequestDetailsEntityListInner`

GetEntityList returns the EntityList field if non-nil, zero value otherwise.

### GetEntityListOk

`func (o *SshUserRequestDetails) GetEntityListOk() (*[]SshUserRequestDetailsEntityListInner, bool)`

GetEntityListOk returns a tuple with the EntityList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityList

`func (o *SshUserRequestDetails) SetEntityList(v []SshUserRequestDetailsEntityListInner)`

SetEntityList sets EntityList field to given value.


### GetUserKeyName

`func (o *SshUserRequestDetails) GetUserKeyName() string`

GetUserKeyName returns the UserKeyName field if non-nil, zero value otherwise.

### GetUserKeyNameOk

`func (o *SshUserRequestDetails) GetUserKeyNameOk() (*string, bool)`

GetUserKeyNameOk returns a tuple with the UserKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserKeyName

`func (o *SshUserRequestDetails) SetUserKeyName(v string)`

SetUserKeyName sets UserKeyName field to given value.


### GetUserKey

`func (o *SshUserRequestDetails) GetUserKey() string`

GetUserKey returns the UserKey field if non-nil, zero value otherwise.

### GetUserKeyOk

`func (o *SshUserRequestDetails) GetUserKeyOk() (*string, bool)`

GetUserKeyOk returns a tuple with the UserKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserKey

`func (o *SshUserRequestDetails) SetUserKey(v string)`

SetUserKey sets UserKey field to given value.


### GetUserUnixId

`func (o *SshUserRequestDetails) GetUserUnixId() int32`

GetUserUnixId returns the UserUnixId field if non-nil, zero value otherwise.

### GetUserUnixIdOk

`func (o *SshUserRequestDetails) GetUserUnixIdOk() (*int32, bool)`

GetUserUnixIdOk returns a tuple with the UserUnixId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserUnixId

`func (o *SshUserRequestDetails) SetUserUnixId(v int32)`

SetUserUnixId sets UserUnixId field to given value.

### HasUserUnixId

`func (o *SshUserRequestDetails) HasUserUnixId() bool`

HasUserUnixId returns a boolean if a field has been set.

### GetAccessRequestUuid

`func (o *SshUserRequestDetails) GetAccessRequestUuid() string`

GetAccessRequestUuid returns the AccessRequestUuid field if non-nil, zero value otherwise.

### GetAccessRequestUuidOk

`func (o *SshUserRequestDetails) GetAccessRequestUuidOk() (*string, bool)`

GetAccessRequestUuidOk returns a tuple with the AccessRequestUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessRequestUuid

`func (o *SshUserRequestDetails) SetAccessRequestUuid(v string)`

SetAccessRequestUuid sets AccessRequestUuid field to given value.


### GetUserUuid

`func (o *SshUserRequestDetails) GetUserUuid() string`

GetUserUuid returns the UserUuid field if non-nil, zero value otherwise.

### GetUserUuidOk

`func (o *SshUserRequestDetails) GetUserUuidOk() (*string, bool)`

GetUserUuidOk returns a tuple with the UserUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserUuid

`func (o *SshUserRequestDetails) SetUserUuid(v string)`

SetUserUuid sets UserUuid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



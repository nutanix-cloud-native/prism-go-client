# SshUserDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | **string** | SSH User name | 
**UserUuid** | **string** | SSH User UUID | 
**AccessRequestUuid** | **string** | UUID of the access request requesting SSH access | 
**EntityList** | [**[]SshUserDetailsEntityListInner**](SshUserDetailsEntityListInner.md) | List of IP addresses to the entity which the user requested access to | 

## Methods

### NewSshUserDetails

`func NewSshUserDetails(username string, userUuid string, accessRequestUuid string, entityList []SshUserDetailsEntityListInner, ) *SshUserDetails`

NewSshUserDetails instantiates a new SshUserDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSshUserDetailsWithDefaults

`func NewSshUserDetailsWithDefaults() *SshUserDetails`

NewSshUserDetailsWithDefaults instantiates a new SshUserDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *SshUserDetails) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *SshUserDetails) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *SshUserDetails) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetUserUuid

`func (o *SshUserDetails) GetUserUuid() string`

GetUserUuid returns the UserUuid field if non-nil, zero value otherwise.

### GetUserUuidOk

`func (o *SshUserDetails) GetUserUuidOk() (*string, bool)`

GetUserUuidOk returns a tuple with the UserUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserUuid

`func (o *SshUserDetails) SetUserUuid(v string)`

SetUserUuid sets UserUuid field to given value.


### GetAccessRequestUuid

`func (o *SshUserDetails) GetAccessRequestUuid() string`

GetAccessRequestUuid returns the AccessRequestUuid field if non-nil, zero value otherwise.

### GetAccessRequestUuidOk

`func (o *SshUserDetails) GetAccessRequestUuidOk() (*string, bool)`

GetAccessRequestUuidOk returns a tuple with the AccessRequestUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessRequestUuid

`func (o *SshUserDetails) SetAccessRequestUuid(v string)`

SetAccessRequestUuid sets AccessRequestUuid field to given value.


### GetEntityList

`func (o *SshUserDetails) GetEntityList() []SshUserDetailsEntityListInner`

GetEntityList returns the EntityList field if non-nil, zero value otherwise.

### GetEntityListOk

`func (o *SshUserDetails) GetEntityListOk() (*[]SshUserDetailsEntityListInner, bool)`

GetEntityListOk returns a tuple with the EntityList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityList

`func (o *SshUserDetails) SetEntityList(v []SshUserDetailsEntityListInner)`

SetEntityList sets EntityList field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



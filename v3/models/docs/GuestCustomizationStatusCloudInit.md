# GuestCustomizationStatusCloudInit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MetaData** | Pointer to **string** | The contents of the meta_data configuration for cloud-init. This can be formatted as YAML or JSON. The value must be base64 encoded.  | [optional] 
**UserData** | Pointer to **string** | The contents of the user_data configuration for cloud-init. This can be formatted as YAML, JSON, or could be a shell script. The value must be base64 encoded.  | [optional] 
**CustomKeyValues** | Pointer to **map[string]string** | Generic key value pair used for custom attributes | [optional] 

## Methods

### NewGuestCustomizationStatusCloudInit

`func NewGuestCustomizationStatusCloudInit() *GuestCustomizationStatusCloudInit`

NewGuestCustomizationStatusCloudInit instantiates a new GuestCustomizationStatusCloudInit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGuestCustomizationStatusCloudInitWithDefaults

`func NewGuestCustomizationStatusCloudInitWithDefaults() *GuestCustomizationStatusCloudInit`

NewGuestCustomizationStatusCloudInitWithDefaults instantiates a new GuestCustomizationStatusCloudInit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetaData

`func (o *GuestCustomizationStatusCloudInit) GetMetaData() string`

GetMetaData returns the MetaData field if non-nil, zero value otherwise.

### GetMetaDataOk

`func (o *GuestCustomizationStatusCloudInit) GetMetaDataOk() (*string, bool)`

GetMetaDataOk returns a tuple with the MetaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaData

`func (o *GuestCustomizationStatusCloudInit) SetMetaData(v string)`

SetMetaData sets MetaData field to given value.

### HasMetaData

`func (o *GuestCustomizationStatusCloudInit) HasMetaData() bool`

HasMetaData returns a boolean if a field has been set.

### GetUserData

`func (o *GuestCustomizationStatusCloudInit) GetUserData() string`

GetUserData returns the UserData field if non-nil, zero value otherwise.

### GetUserDataOk

`func (o *GuestCustomizationStatusCloudInit) GetUserDataOk() (*string, bool)`

GetUserDataOk returns a tuple with the UserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserData

`func (o *GuestCustomizationStatusCloudInit) SetUserData(v string)`

SetUserData sets UserData field to given value.

### HasUserData

`func (o *GuestCustomizationStatusCloudInit) HasUserData() bool`

HasUserData returns a boolean if a field has been set.

### GetCustomKeyValues

`func (o *GuestCustomizationStatusCloudInit) GetCustomKeyValues() map[string]string`

GetCustomKeyValues returns the CustomKeyValues field if non-nil, zero value otherwise.

### GetCustomKeyValuesOk

`func (o *GuestCustomizationStatusCloudInit) GetCustomKeyValuesOk() (*map[string]string, bool)`

GetCustomKeyValuesOk returns a tuple with the CustomKeyValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomKeyValues

`func (o *GuestCustomizationStatusCloudInit) SetCustomKeyValues(v map[string]string)`

SetCustomKeyValues sets CustomKeyValues field to given value.

### HasCustomKeyValues

`func (o *GuestCustomizationStatusCloudInit) HasCustomKeyValues() bool`

HasCustomKeyValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



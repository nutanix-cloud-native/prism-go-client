# InternalServiceAccountConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceAccountUsername** | Pointer to **string** | The username to connect to the directory service. | [optional] 
**ServiceAccountPassword** | Pointer to **string** | The password to authenticate the request. | [optional] 

## Methods

### NewInternalServiceAccountConfig

`func NewInternalServiceAccountConfig() *InternalServiceAccountConfig`

NewInternalServiceAccountConfig instantiates a new InternalServiceAccountConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternalServiceAccountConfigWithDefaults

`func NewInternalServiceAccountConfigWithDefaults() *InternalServiceAccountConfig`

NewInternalServiceAccountConfigWithDefaults instantiates a new InternalServiceAccountConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceAccountUsername

`func (o *InternalServiceAccountConfig) GetServiceAccountUsername() string`

GetServiceAccountUsername returns the ServiceAccountUsername field if non-nil, zero value otherwise.

### GetServiceAccountUsernameOk

`func (o *InternalServiceAccountConfig) GetServiceAccountUsernameOk() (*string, bool)`

GetServiceAccountUsernameOk returns a tuple with the ServiceAccountUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountUsername

`func (o *InternalServiceAccountConfig) SetServiceAccountUsername(v string)`

SetServiceAccountUsername sets ServiceAccountUsername field to given value.

### HasServiceAccountUsername

`func (o *InternalServiceAccountConfig) HasServiceAccountUsername() bool`

HasServiceAccountUsername returns a boolean if a field has been set.

### GetServiceAccountPassword

`func (o *InternalServiceAccountConfig) GetServiceAccountPassword() string`

GetServiceAccountPassword returns the ServiceAccountPassword field if non-nil, zero value otherwise.

### GetServiceAccountPasswordOk

`func (o *InternalServiceAccountConfig) GetServiceAccountPasswordOk() (*string, bool)`

GetServiceAccountPasswordOk returns a tuple with the ServiceAccountPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountPassword

`func (o *InternalServiceAccountConfig) SetServiceAccountPassword(v string)`

SetServiceAccountPassword sets ServiceAccountPassword field to given value.

### HasServiceAccountPassword

`func (o *InternalServiceAccountConfig) HasServiceAccountPassword() bool`

HasServiceAccountPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



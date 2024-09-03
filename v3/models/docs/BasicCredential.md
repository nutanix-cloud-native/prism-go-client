# BasicCredential

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | Pointer to **string** | username of authorized user on remote cluster, must for BASIC auth_type  | [optional] 
**Password** | Pointer to **string** | password of authorized user on remote cluster, must for BASIC auth_type  | [optional] 

## Methods

### NewBasicCredential

`func NewBasicCredential() *BasicCredential`

NewBasicCredential instantiates a new BasicCredential object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBasicCredentialWithDefaults

`func NewBasicCredentialWithDefaults() *BasicCredential`

NewBasicCredentialWithDefaults instantiates a new BasicCredential object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *BasicCredential) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *BasicCredential) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *BasicCredential) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *BasicCredential) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *BasicCredential) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *BasicCredential) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *BasicCredential) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *BasicCredential) HasPassword() bool`

HasPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



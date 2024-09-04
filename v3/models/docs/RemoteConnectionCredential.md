# RemoteConnectionCredential

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthType** | Pointer to **string** | Type of credentials to use | [optional] 
**BearerToken** | Pointer to **string** | token to use with bearer authorization, must for BEARER auth_type  | [optional] 
**BasicCredential** | Pointer to [**BasicCredential**](BasicCredential.md) |  | [optional] 

## Methods

### NewRemoteConnectionCredential

`func NewRemoteConnectionCredential() *RemoteConnectionCredential`

NewRemoteConnectionCredential instantiates a new RemoteConnectionCredential object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoteConnectionCredentialWithDefaults

`func NewRemoteConnectionCredentialWithDefaults() *RemoteConnectionCredential`

NewRemoteConnectionCredentialWithDefaults instantiates a new RemoteConnectionCredential object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthType

`func (o *RemoteConnectionCredential) GetAuthType() string`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *RemoteConnectionCredential) GetAuthTypeOk() (*string, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *RemoteConnectionCredential) SetAuthType(v string)`

SetAuthType sets AuthType field to given value.

### HasAuthType

`func (o *RemoteConnectionCredential) HasAuthType() bool`

HasAuthType returns a boolean if a field has been set.

### GetBearerToken

`func (o *RemoteConnectionCredential) GetBearerToken() string`

GetBearerToken returns the BearerToken field if non-nil, zero value otherwise.

### GetBearerTokenOk

`func (o *RemoteConnectionCredential) GetBearerTokenOk() (*string, bool)`

GetBearerTokenOk returns a tuple with the BearerToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBearerToken

`func (o *RemoteConnectionCredential) SetBearerToken(v string)`

SetBearerToken sets BearerToken field to given value.

### HasBearerToken

`func (o *RemoteConnectionCredential) HasBearerToken() bool`

HasBearerToken returns a boolean if a field has been set.

### GetBasicCredential

`func (o *RemoteConnectionCredential) GetBasicCredential() BasicCredential`

GetBasicCredential returns the BasicCredential field if non-nil, zero value otherwise.

### GetBasicCredentialOk

`func (o *RemoteConnectionCredential) GetBasicCredentialOk() (*BasicCredential, bool)`

GetBasicCredentialOk returns a tuple with the BasicCredential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicCredential

`func (o *RemoteConnectionCredential) SetBasicCredential(v BasicCredential)`

SetBasicCredential sets BasicCredential field to given value.

### HasBasicCredential

`func (o *RemoteConnectionCredential) HasBasicCredential() bool`

HasBasicCredential returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



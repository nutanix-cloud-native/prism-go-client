# OauthGenerateTokenRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Scopes** | **[]string** |  | 
**ClientName** | **string** | client name of the Oauth Client | 
**ClientDescription** | **string** | client description of the Oauth Client | 
**ExpiresIn** | Pointer to **int64** | Token expiration time in seconds | [optional] 

## Methods

### NewOauthGenerateTokenRequest

`func NewOauthGenerateTokenRequest(scopes []string, clientName string, clientDescription string, ) *OauthGenerateTokenRequest`

NewOauthGenerateTokenRequest instantiates a new OauthGenerateTokenRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOauthGenerateTokenRequestWithDefaults

`func NewOauthGenerateTokenRequestWithDefaults() *OauthGenerateTokenRequest`

NewOauthGenerateTokenRequestWithDefaults instantiates a new OauthGenerateTokenRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScopes

`func (o *OauthGenerateTokenRequest) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *OauthGenerateTokenRequest) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *OauthGenerateTokenRequest) SetScopes(v []string)`

SetScopes sets Scopes field to given value.


### GetClientName

`func (o *OauthGenerateTokenRequest) GetClientName() string`

GetClientName returns the ClientName field if non-nil, zero value otherwise.

### GetClientNameOk

`func (o *OauthGenerateTokenRequest) GetClientNameOk() (*string, bool)`

GetClientNameOk returns a tuple with the ClientName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientName

`func (o *OauthGenerateTokenRequest) SetClientName(v string)`

SetClientName sets ClientName field to given value.


### GetClientDescription

`func (o *OauthGenerateTokenRequest) GetClientDescription() string`

GetClientDescription returns the ClientDescription field if non-nil, zero value otherwise.

### GetClientDescriptionOk

`func (o *OauthGenerateTokenRequest) GetClientDescriptionOk() (*string, bool)`

GetClientDescriptionOk returns a tuple with the ClientDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientDescription

`func (o *OauthGenerateTokenRequest) SetClientDescription(v string)`

SetClientDescription sets ClientDescription field to given value.


### GetExpiresIn

`func (o *OauthGenerateTokenRequest) GetExpiresIn() int64`

GetExpiresIn returns the ExpiresIn field if non-nil, zero value otherwise.

### GetExpiresInOk

`func (o *OauthGenerateTokenRequest) GetExpiresInOk() (*int64, bool)`

GetExpiresInOk returns a tuple with the ExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresIn

`func (o *OauthGenerateTokenRequest) SetExpiresIn(v int64)`

SetExpiresIn sets ExpiresIn field to given value.

### HasExpiresIn

`func (o *OauthGenerateTokenRequest) HasExpiresIn() bool`

HasExpiresIn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



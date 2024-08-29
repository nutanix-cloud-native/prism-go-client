# ProviderOauthToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | **string** | access_token to be used for accessing gateway | 
**TokenType** | **string** | Only value possible is \&quot;bearer\&quot;. | [readonly] 
**ExpiresIn** | **int64** | Token expiration time in seconds | 
**RefreshToken** | **string** | Refresh_token which can used to get new token | 
**Scope** | **string** | Scope string with individual scopes separated by space | 

## Methods

### NewProviderOauthToken

`func NewProviderOauthToken(accessToken string, tokenType string, expiresIn int64, refreshToken string, scope string, ) *ProviderOauthToken`

NewProviderOauthToken instantiates a new ProviderOauthToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderOauthTokenWithDefaults

`func NewProviderOauthTokenWithDefaults() *ProviderOauthToken`

NewProviderOauthTokenWithDefaults instantiates a new ProviderOauthToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessToken

`func (o *ProviderOauthToken) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *ProviderOauthToken) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *ProviderOauthToken) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.


### GetTokenType

`func (o *ProviderOauthToken) GetTokenType() string`

GetTokenType returns the TokenType field if non-nil, zero value otherwise.

### GetTokenTypeOk

`func (o *ProviderOauthToken) GetTokenTypeOk() (*string, bool)`

GetTokenTypeOk returns a tuple with the TokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenType

`func (o *ProviderOauthToken) SetTokenType(v string)`

SetTokenType sets TokenType field to given value.


### GetExpiresIn

`func (o *ProviderOauthToken) GetExpiresIn() int64`

GetExpiresIn returns the ExpiresIn field if non-nil, zero value otherwise.

### GetExpiresInOk

`func (o *ProviderOauthToken) GetExpiresInOk() (*int64, bool)`

GetExpiresInOk returns a tuple with the ExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresIn

`func (o *ProviderOauthToken) SetExpiresIn(v int64)`

SetExpiresIn sets ExpiresIn field to given value.


### GetRefreshToken

`func (o *ProviderOauthToken) GetRefreshToken() string`

GetRefreshToken returns the RefreshToken field if non-nil, zero value otherwise.

### GetRefreshTokenOk

`func (o *ProviderOauthToken) GetRefreshTokenOk() (*string, bool)`

GetRefreshTokenOk returns a tuple with the RefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshToken

`func (o *ProviderOauthToken) SetRefreshToken(v string)`

SetRefreshToken sets RefreshToken field to given value.


### GetScope

`func (o *ProviderOauthToken) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *ProviderOauthToken) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *ProviderOauthToken) SetScope(v string)`

SetScope sets Scope field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



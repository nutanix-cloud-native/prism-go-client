# OauthGenerateTokenResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientSecret** | **string** | client secret of the Oauth Client | 
**OauthToken** | [**OauthToken**](OauthToken.md) |  | 
**ClientId** | **string** | client id of the Oauth Client | 

## Methods

### NewOauthGenerateTokenResponse

`func NewOauthGenerateTokenResponse(clientSecret string, oauthToken OauthToken, clientId string, ) *OauthGenerateTokenResponse`

NewOauthGenerateTokenResponse instantiates a new OauthGenerateTokenResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOauthGenerateTokenResponseWithDefaults

`func NewOauthGenerateTokenResponseWithDefaults() *OauthGenerateTokenResponse`

NewOauthGenerateTokenResponseWithDefaults instantiates a new OauthGenerateTokenResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientSecret

`func (o *OauthGenerateTokenResponse) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *OauthGenerateTokenResponse) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *OauthGenerateTokenResponse) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.


### GetOauthToken

`func (o *OauthGenerateTokenResponse) GetOauthToken() OauthToken`

GetOauthToken returns the OauthToken field if non-nil, zero value otherwise.

### GetOauthTokenOk

`func (o *OauthGenerateTokenResponse) GetOauthTokenOk() (*OauthToken, bool)`

GetOauthTokenOk returns a tuple with the OauthToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthToken

`func (o *OauthGenerateTokenResponse) SetOauthToken(v OauthToken)`

SetOauthToken sets OauthToken field to given value.


### GetClientId

`func (o *OauthGenerateTokenResponse) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *OauthGenerateTokenResponse) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *OauthGenerateTokenResponse) SetClientId(v string)`

SetClientId sets ClientId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# OauthClientResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RedirectUris** | **[]string** |  | 
**ClientName** | **string** | client name of the Oauth Client | 
**ClientId** | **string** | client id of the Oauth Client | 
**ClientSecret** | **string** | client secret of the Oauth Client | 
**ClientDescription** | **string** | client description of the Oauth Client | 
**DefaultScopes** | **[]string** |  | 

## Methods

### NewOauthClientResponse

`func NewOauthClientResponse(redirectUris []string, clientName string, clientId string, clientSecret string, clientDescription string, defaultScopes []string, ) *OauthClientResponse`

NewOauthClientResponse instantiates a new OauthClientResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOauthClientResponseWithDefaults

`func NewOauthClientResponseWithDefaults() *OauthClientResponse`

NewOauthClientResponseWithDefaults instantiates a new OauthClientResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRedirectUris

`func (o *OauthClientResponse) GetRedirectUris() []string`

GetRedirectUris returns the RedirectUris field if non-nil, zero value otherwise.

### GetRedirectUrisOk

`func (o *OauthClientResponse) GetRedirectUrisOk() (*[]string, bool)`

GetRedirectUrisOk returns a tuple with the RedirectUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUris

`func (o *OauthClientResponse) SetRedirectUris(v []string)`

SetRedirectUris sets RedirectUris field to given value.


### GetClientName

`func (o *OauthClientResponse) GetClientName() string`

GetClientName returns the ClientName field if non-nil, zero value otherwise.

### GetClientNameOk

`func (o *OauthClientResponse) GetClientNameOk() (*string, bool)`

GetClientNameOk returns a tuple with the ClientName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientName

`func (o *OauthClientResponse) SetClientName(v string)`

SetClientName sets ClientName field to given value.


### GetClientId

`func (o *OauthClientResponse) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *OauthClientResponse) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *OauthClientResponse) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetClientSecret

`func (o *OauthClientResponse) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *OauthClientResponse) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *OauthClientResponse) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.


### GetClientDescription

`func (o *OauthClientResponse) GetClientDescription() string`

GetClientDescription returns the ClientDescription field if non-nil, zero value otherwise.

### GetClientDescriptionOk

`func (o *OauthClientResponse) GetClientDescriptionOk() (*string, bool)`

GetClientDescriptionOk returns a tuple with the ClientDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientDescription

`func (o *OauthClientResponse) SetClientDescription(v string)`

SetClientDescription sets ClientDescription field to given value.


### GetDefaultScopes

`func (o *OauthClientResponse) GetDefaultScopes() []string`

GetDefaultScopes returns the DefaultScopes field if non-nil, zero value otherwise.

### GetDefaultScopesOk

`func (o *OauthClientResponse) GetDefaultScopesOk() (*[]string, bool)`

GetDefaultScopesOk returns a tuple with the DefaultScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultScopes

`func (o *OauthClientResponse) SetDefaultScopes(v []string)`

SetDefaultScopes sets DefaultScopes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



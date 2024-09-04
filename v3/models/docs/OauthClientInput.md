# OauthClientInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientName** | **string** | client name of the Oauth Client | 
**RedirectUris** | **[]string** |  | 
**ClientDescription** | **string** | client description of the Oauth Client | 
**DefaultScopes** | **[]string** |  | 

## Methods

### NewOauthClientInput

`func NewOauthClientInput(clientName string, redirectUris []string, clientDescription string, defaultScopes []string, ) *OauthClientInput`

NewOauthClientInput instantiates a new OauthClientInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOauthClientInputWithDefaults

`func NewOauthClientInputWithDefaults() *OauthClientInput`

NewOauthClientInputWithDefaults instantiates a new OauthClientInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientName

`func (o *OauthClientInput) GetClientName() string`

GetClientName returns the ClientName field if non-nil, zero value otherwise.

### GetClientNameOk

`func (o *OauthClientInput) GetClientNameOk() (*string, bool)`

GetClientNameOk returns a tuple with the ClientName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientName

`func (o *OauthClientInput) SetClientName(v string)`

SetClientName sets ClientName field to given value.


### GetRedirectUris

`func (o *OauthClientInput) GetRedirectUris() []string`

GetRedirectUris returns the RedirectUris field if non-nil, zero value otherwise.

### GetRedirectUrisOk

`func (o *OauthClientInput) GetRedirectUrisOk() (*[]string, bool)`

GetRedirectUrisOk returns a tuple with the RedirectUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUris

`func (o *OauthClientInput) SetRedirectUris(v []string)`

SetRedirectUris sets RedirectUris field to given value.


### GetClientDescription

`func (o *OauthClientInput) GetClientDescription() string`

GetClientDescription returns the ClientDescription field if non-nil, zero value otherwise.

### GetClientDescriptionOk

`func (o *OauthClientInput) GetClientDescriptionOk() (*string, bool)`

GetClientDescriptionOk returns a tuple with the ClientDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientDescription

`func (o *OauthClientInput) SetClientDescription(v string)`

SetClientDescription sets ClientDescription field to given value.


### GetDefaultScopes

`func (o *OauthClientInput) GetDefaultScopes() []string`

GetDefaultScopes returns the DefaultScopes field if non-nil, zero value otherwise.

### GetDefaultScopesOk

`func (o *OauthClientInput) GetDefaultScopesOk() (*[]string, bool)`

GetDefaultScopesOk returns a tuple with the DefaultScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultScopes

`func (o *OauthClientInput) SetDefaultScopes(v []string)`

SetDefaultScopes sets DefaultScopes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



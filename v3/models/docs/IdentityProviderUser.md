# IdentityProviderUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | Pointer to **string** | The username from the identity provider. Name Id for SAML Identity Provider.  | [optional] 
**IdentityProviderReference** | Pointer to [**IdentityProviderReference**](IdentityProviderReference.md) |  | [optional] 

## Methods

### NewIdentityProviderUser

`func NewIdentityProviderUser() *IdentityProviderUser`

NewIdentityProviderUser instantiates a new IdentityProviderUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityProviderUserWithDefaults

`func NewIdentityProviderUserWithDefaults() *IdentityProviderUser`

NewIdentityProviderUserWithDefaults instantiates a new IdentityProviderUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *IdentityProviderUser) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *IdentityProviderUser) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *IdentityProviderUser) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *IdentityProviderUser) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetIdentityProviderReference

`func (o *IdentityProviderUser) GetIdentityProviderReference() IdentityProviderReference`

GetIdentityProviderReference returns the IdentityProviderReference field if non-nil, zero value otherwise.

### GetIdentityProviderReferenceOk

`func (o *IdentityProviderUser) GetIdentityProviderReferenceOk() (*IdentityProviderReference, bool)`

GetIdentityProviderReferenceOk returns a tuple with the IdentityProviderReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityProviderReference

`func (o *IdentityProviderUser) SetIdentityProviderReference(v IdentityProviderReference)`

SetIdentityProviderReference sets IdentityProviderReference field to given value.

### HasIdentityProviderReference

`func (o *IdentityProviderUser) HasIdentityProviderReference() bool`

HasIdentityProviderReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



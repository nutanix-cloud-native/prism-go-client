# SamlUserGroupInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdpUuid** | Pointer to **string** | The UUID of the Identity Provider that the group belongs to | [optional] 
**Name** | Pointer to **string** | The name of the SAML group which the IDP provides as attribute in SAML response | [optional] 

## Methods

### NewSamlUserGroupInput

`func NewSamlUserGroupInput() *SamlUserGroupInput`

NewSamlUserGroupInput instantiates a new SamlUserGroupInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSamlUserGroupInputWithDefaults

`func NewSamlUserGroupInputWithDefaults() *SamlUserGroupInput`

NewSamlUserGroupInputWithDefaults instantiates a new SamlUserGroupInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdpUuid

`func (o *SamlUserGroupInput) GetIdpUuid() string`

GetIdpUuid returns the IdpUuid field if non-nil, zero value otherwise.

### GetIdpUuidOk

`func (o *SamlUserGroupInput) GetIdpUuidOk() (*string, bool)`

GetIdpUuidOk returns a tuple with the IdpUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpUuid

`func (o *SamlUserGroupInput) SetIdpUuid(v string)`

SetIdpUuid sets IdpUuid field to given value.

### HasIdpUuid

`func (o *SamlUserGroupInput) HasIdpUuid() bool`

HasIdpUuid returns a boolean if a field has been set.

### GetName

`func (o *SamlUserGroupInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SamlUserGroupInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SamlUserGroupInput) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SamlUserGroupInput) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



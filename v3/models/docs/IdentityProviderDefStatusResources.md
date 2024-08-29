# IdentityProviderDefStatusResources

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupsDelim** | Pointer to **string** | If groups delimiter is provided groups are assumed to be represented as a single attribute and the delimiter is used to split the attribute&#39;s value into multiple groups.  | [optional] 
**IdpProperties** | Pointer to [**IdentityProviderDefStatusResourcesIdpProperties**](IdentityProviderDefStatusResourcesIdpProperties.md) |  | [optional] 
**IdpMetadata** | Pointer to **string** | Metadata in xml format with IDP details. | [optional] 
**GroupsAttr** | Pointer to **string** | Saml assertion groups attribute element. | [optional] 

## Methods

### NewIdentityProviderDefStatusResources

`func NewIdentityProviderDefStatusResources() *IdentityProviderDefStatusResources`

NewIdentityProviderDefStatusResources instantiates a new IdentityProviderDefStatusResources object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityProviderDefStatusResourcesWithDefaults

`func NewIdentityProviderDefStatusResourcesWithDefaults() *IdentityProviderDefStatusResources`

NewIdentityProviderDefStatusResourcesWithDefaults instantiates a new IdentityProviderDefStatusResources object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupsDelim

`func (o *IdentityProviderDefStatusResources) GetGroupsDelim() string`

GetGroupsDelim returns the GroupsDelim field if non-nil, zero value otherwise.

### GetGroupsDelimOk

`func (o *IdentityProviderDefStatusResources) GetGroupsDelimOk() (*string, bool)`

GetGroupsDelimOk returns a tuple with the GroupsDelim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupsDelim

`func (o *IdentityProviderDefStatusResources) SetGroupsDelim(v string)`

SetGroupsDelim sets GroupsDelim field to given value.

### HasGroupsDelim

`func (o *IdentityProviderDefStatusResources) HasGroupsDelim() bool`

HasGroupsDelim returns a boolean if a field has been set.

### GetIdpProperties

`func (o *IdentityProviderDefStatusResources) GetIdpProperties() IdentityProviderDefStatusResourcesIdpProperties`

GetIdpProperties returns the IdpProperties field if non-nil, zero value otherwise.

### GetIdpPropertiesOk

`func (o *IdentityProviderDefStatusResources) GetIdpPropertiesOk() (*IdentityProviderDefStatusResourcesIdpProperties, bool)`

GetIdpPropertiesOk returns a tuple with the IdpProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpProperties

`func (o *IdentityProviderDefStatusResources) SetIdpProperties(v IdentityProviderDefStatusResourcesIdpProperties)`

SetIdpProperties sets IdpProperties field to given value.

### HasIdpProperties

`func (o *IdentityProviderDefStatusResources) HasIdpProperties() bool`

HasIdpProperties returns a boolean if a field has been set.

### GetIdpMetadata

`func (o *IdentityProviderDefStatusResources) GetIdpMetadata() string`

GetIdpMetadata returns the IdpMetadata field if non-nil, zero value otherwise.

### GetIdpMetadataOk

`func (o *IdentityProviderDefStatusResources) GetIdpMetadataOk() (*string, bool)`

GetIdpMetadataOk returns a tuple with the IdpMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpMetadata

`func (o *IdentityProviderDefStatusResources) SetIdpMetadata(v string)`

SetIdpMetadata sets IdpMetadata field to given value.

### HasIdpMetadata

`func (o *IdentityProviderDefStatusResources) HasIdpMetadata() bool`

HasIdpMetadata returns a boolean if a field has been set.

### GetGroupsAttr

`func (o *IdentityProviderDefStatusResources) GetGroupsAttr() string`

GetGroupsAttr returns the GroupsAttr field if non-nil, zero value otherwise.

### GetGroupsAttrOk

`func (o *IdentityProviderDefStatusResources) GetGroupsAttrOk() (*string, bool)`

GetGroupsAttrOk returns a tuple with the GroupsAttr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupsAttr

`func (o *IdentityProviderDefStatusResources) SetGroupsAttr(v string)`

SetGroupsAttr sets GroupsAttr field to given value.

### HasGroupsAttr

`func (o *IdentityProviderDefStatusResources) HasGroupsAttr() bool`

HasGroupsAttr returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



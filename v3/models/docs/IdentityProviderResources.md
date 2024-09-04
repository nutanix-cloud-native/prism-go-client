# IdentityProviderResources

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupsDelim** | Pointer to **string** | If groups delimiter is provided groups are assumed to be represented as a single attribute and the delimiter is used to split the attribute&#39;s value into multiple groups.  | [optional] 
**IdpProperties** | Pointer to [**IdentityProviderDefStatusResourcesIdpProperties**](IdentityProviderDefStatusResourcesIdpProperties.md) |  | [optional] 
**IdpMetadata** | Pointer to **string** | Metadata in xml format with IDP details. | [optional] 
**GroupsAttr** | Pointer to **string** | Saml assertion groups attribute element. | [optional] 

## Methods

### NewIdentityProviderResources

`func NewIdentityProviderResources() *IdentityProviderResources`

NewIdentityProviderResources instantiates a new IdentityProviderResources object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityProviderResourcesWithDefaults

`func NewIdentityProviderResourcesWithDefaults() *IdentityProviderResources`

NewIdentityProviderResourcesWithDefaults instantiates a new IdentityProviderResources object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupsDelim

`func (o *IdentityProviderResources) GetGroupsDelim() string`

GetGroupsDelim returns the GroupsDelim field if non-nil, zero value otherwise.

### GetGroupsDelimOk

`func (o *IdentityProviderResources) GetGroupsDelimOk() (*string, bool)`

GetGroupsDelimOk returns a tuple with the GroupsDelim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupsDelim

`func (o *IdentityProviderResources) SetGroupsDelim(v string)`

SetGroupsDelim sets GroupsDelim field to given value.

### HasGroupsDelim

`func (o *IdentityProviderResources) HasGroupsDelim() bool`

HasGroupsDelim returns a boolean if a field has been set.

### GetIdpProperties

`func (o *IdentityProviderResources) GetIdpProperties() IdentityProviderDefStatusResourcesIdpProperties`

GetIdpProperties returns the IdpProperties field if non-nil, zero value otherwise.

### GetIdpPropertiesOk

`func (o *IdentityProviderResources) GetIdpPropertiesOk() (*IdentityProviderDefStatusResourcesIdpProperties, bool)`

GetIdpPropertiesOk returns a tuple with the IdpProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpProperties

`func (o *IdentityProviderResources) SetIdpProperties(v IdentityProviderDefStatusResourcesIdpProperties)`

SetIdpProperties sets IdpProperties field to given value.

### HasIdpProperties

`func (o *IdentityProviderResources) HasIdpProperties() bool`

HasIdpProperties returns a boolean if a field has been set.

### GetIdpMetadata

`func (o *IdentityProviderResources) GetIdpMetadata() string`

GetIdpMetadata returns the IdpMetadata field if non-nil, zero value otherwise.

### GetIdpMetadataOk

`func (o *IdentityProviderResources) GetIdpMetadataOk() (*string, bool)`

GetIdpMetadataOk returns a tuple with the IdpMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpMetadata

`func (o *IdentityProviderResources) SetIdpMetadata(v string)`

SetIdpMetadata sets IdpMetadata field to given value.

### HasIdpMetadata

`func (o *IdentityProviderResources) HasIdpMetadata() bool`

HasIdpMetadata returns a boolean if a field has been set.

### GetGroupsAttr

`func (o *IdentityProviderResources) GetGroupsAttr() string`

GetGroupsAttr returns the GroupsAttr field if non-nil, zero value otherwise.

### GetGroupsAttrOk

`func (o *IdentityProviderResources) GetGroupsAttrOk() (*string, bool)`

GetGroupsAttrOk returns a tuple with the GroupsAttr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupsAttr

`func (o *IdentityProviderResources) SetGroupsAttr(v string)`

SetGroupsAttr sets GroupsAttr field to given value.

### HasGroupsAttr

`func (o *IdentityProviderResources) HasGroupsAttr() bool`

HasGroupsAttr returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



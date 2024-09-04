# DirectoryServiceUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserPrincipalName** | Pointer to **string** | The UserPrincipalName of the user from the directory service.  | [optional] 
**DirectoryServiceReference** | Pointer to [**DirectoryServiceReference**](DirectoryServiceReference.md) |  | [optional] 
**UserAttributeValue** | Pointer to **string** | The Unique identifier for each user from the directory service.  | [optional] 

## Methods

### NewDirectoryServiceUser

`func NewDirectoryServiceUser() *DirectoryServiceUser`

NewDirectoryServiceUser instantiates a new DirectoryServiceUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDirectoryServiceUserWithDefaults

`func NewDirectoryServiceUserWithDefaults() *DirectoryServiceUser`

NewDirectoryServiceUserWithDefaults instantiates a new DirectoryServiceUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserPrincipalName

`func (o *DirectoryServiceUser) GetUserPrincipalName() string`

GetUserPrincipalName returns the UserPrincipalName field if non-nil, zero value otherwise.

### GetUserPrincipalNameOk

`func (o *DirectoryServiceUser) GetUserPrincipalNameOk() (*string, bool)`

GetUserPrincipalNameOk returns a tuple with the UserPrincipalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPrincipalName

`func (o *DirectoryServiceUser) SetUserPrincipalName(v string)`

SetUserPrincipalName sets UserPrincipalName field to given value.

### HasUserPrincipalName

`func (o *DirectoryServiceUser) HasUserPrincipalName() bool`

HasUserPrincipalName returns a boolean if a field has been set.

### GetDirectoryServiceReference

`func (o *DirectoryServiceUser) GetDirectoryServiceReference() DirectoryServiceReference`

GetDirectoryServiceReference returns the DirectoryServiceReference field if non-nil, zero value otherwise.

### GetDirectoryServiceReferenceOk

`func (o *DirectoryServiceUser) GetDirectoryServiceReferenceOk() (*DirectoryServiceReference, bool)`

GetDirectoryServiceReferenceOk returns a tuple with the DirectoryServiceReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryServiceReference

`func (o *DirectoryServiceUser) SetDirectoryServiceReference(v DirectoryServiceReference)`

SetDirectoryServiceReference sets DirectoryServiceReference field to given value.

### HasDirectoryServiceReference

`func (o *DirectoryServiceUser) HasDirectoryServiceReference() bool`

HasDirectoryServiceReference returns a boolean if a field has been set.

### GetUserAttributeValue

`func (o *DirectoryServiceUser) GetUserAttributeValue() string`

GetUserAttributeValue returns the UserAttributeValue field if non-nil, zero value otherwise.

### GetUserAttributeValueOk

`func (o *DirectoryServiceUser) GetUserAttributeValueOk() (*string, bool)`

GetUserAttributeValueOk returns a tuple with the UserAttributeValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAttributeValue

`func (o *DirectoryServiceUser) SetUserAttributeValue(v string)`

SetUserAttributeValue sets UserAttributeValue field to given value.

### HasUserAttributeValue

`func (o *DirectoryServiceUser) HasUserAttributeValue() bool`

HasUserAttributeValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



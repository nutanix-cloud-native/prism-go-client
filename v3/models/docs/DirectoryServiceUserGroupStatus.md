# DirectoryServiceUserGroupStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DistinguishedName** | Pointer to **string** | The Distinguished name for the user group. | [optional] 
**DirectoryServiceReference** | Pointer to [**DirectoryServiceReference**](DirectoryServiceReference.md) |  | [optional] 

## Methods

### NewDirectoryServiceUserGroupStatus

`func NewDirectoryServiceUserGroupStatus() *DirectoryServiceUserGroupStatus`

NewDirectoryServiceUserGroupStatus instantiates a new DirectoryServiceUserGroupStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDirectoryServiceUserGroupStatusWithDefaults

`func NewDirectoryServiceUserGroupStatusWithDefaults() *DirectoryServiceUserGroupStatus`

NewDirectoryServiceUserGroupStatusWithDefaults instantiates a new DirectoryServiceUserGroupStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDistinguishedName

`func (o *DirectoryServiceUserGroupStatus) GetDistinguishedName() string`

GetDistinguishedName returns the DistinguishedName field if non-nil, zero value otherwise.

### GetDistinguishedNameOk

`func (o *DirectoryServiceUserGroupStatus) GetDistinguishedNameOk() (*string, bool)`

GetDistinguishedNameOk returns a tuple with the DistinguishedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistinguishedName

`func (o *DirectoryServiceUserGroupStatus) SetDistinguishedName(v string)`

SetDistinguishedName sets DistinguishedName field to given value.

### HasDistinguishedName

`func (o *DirectoryServiceUserGroupStatus) HasDistinguishedName() bool`

HasDistinguishedName returns a boolean if a field has been set.

### GetDirectoryServiceReference

`func (o *DirectoryServiceUserGroupStatus) GetDirectoryServiceReference() DirectoryServiceReference`

GetDirectoryServiceReference returns the DirectoryServiceReference field if non-nil, zero value otherwise.

### GetDirectoryServiceReferenceOk

`func (o *DirectoryServiceUserGroupStatus) GetDirectoryServiceReferenceOk() (*DirectoryServiceReference, bool)`

GetDirectoryServiceReferenceOk returns a tuple with the DirectoryServiceReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryServiceReference

`func (o *DirectoryServiceUserGroupStatus) SetDirectoryServiceReference(v DirectoryServiceReference)`

SetDirectoryServiceReference sets DirectoryServiceReference field to given value.

### HasDirectoryServiceReference

`func (o *DirectoryServiceUserGroupStatus) HasDirectoryServiceReference() bool`

HasDirectoryServiceReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



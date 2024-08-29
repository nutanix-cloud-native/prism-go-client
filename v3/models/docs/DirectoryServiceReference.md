# DirectoryServiceReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | The kind name | [readonly] [default to "directory_service"]
**Name** | Pointer to **string** |  | [optional] [readonly] 
**UUID** | **string** |  | 

## Methods

### NewDirectoryServiceReference

`func NewDirectoryServiceReference(kind string, uUID string, ) *DirectoryServiceReference`

NewDirectoryServiceReference instantiates a new DirectoryServiceReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDirectoryServiceReferenceWithDefaults

`func NewDirectoryServiceReferenceWithDefaults() *DirectoryServiceReference`

NewDirectoryServiceReferenceWithDefaults instantiates a new DirectoryServiceReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *DirectoryServiceReference) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *DirectoryServiceReference) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *DirectoryServiceReference) SetKind(v string)`

SetKind sets Kind field to given value.


### GetName

`func (o *DirectoryServiceReference) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DirectoryServiceReference) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DirectoryServiceReference) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DirectoryServiceReference) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUUID

`func (o *DirectoryServiceReference) GetUUID() string`

GetUUID returns the UUID field if non-nil, zero value otherwise.

### GetUUIDOk

`func (o *DirectoryServiceReference) GetUUIDOk() (*string, bool)`

GetUUIDOk returns a tuple with the UUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUUID

`func (o *DirectoryServiceReference) SetUUID(v string)`

SetUUID sets UUID field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



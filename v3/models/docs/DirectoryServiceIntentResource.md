# DirectoryServiceIntentResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**DirectoryServiceDefStatus**](DirectoryServiceDefStatus.md) |  | [optional] 
**Spec** | Pointer to [**DirectoryService**](DirectoryService.md) |  | [optional] 
**ApiVersion** | Pointer to **string** | API Version of the Nutanix v3 API framework. | [optional] [readonly] [default to "3.1.0"]
**Metadata** | [**DirectoryServiceMetadata**](DirectoryServiceMetadata.md) |  | 

## Methods

### NewDirectoryServiceIntentResource

`func NewDirectoryServiceIntentResource(metadata DirectoryServiceMetadata, ) *DirectoryServiceIntentResource`

NewDirectoryServiceIntentResource instantiates a new DirectoryServiceIntentResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDirectoryServiceIntentResourceWithDefaults

`func NewDirectoryServiceIntentResourceWithDefaults() *DirectoryServiceIntentResource`

NewDirectoryServiceIntentResourceWithDefaults instantiates a new DirectoryServiceIntentResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *DirectoryServiceIntentResource) GetStatus() DirectoryServiceDefStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DirectoryServiceIntentResource) GetStatusOk() (*DirectoryServiceDefStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DirectoryServiceIntentResource) SetStatus(v DirectoryServiceDefStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DirectoryServiceIntentResource) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSpec

`func (o *DirectoryServiceIntentResource) GetSpec() DirectoryService`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *DirectoryServiceIntentResource) GetSpecOk() (*DirectoryService, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *DirectoryServiceIntentResource) SetSpec(v DirectoryService)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *DirectoryServiceIntentResource) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetApiVersion

`func (o *DirectoryServiceIntentResource) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *DirectoryServiceIntentResource) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *DirectoryServiceIntentResource) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *DirectoryServiceIntentResource) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetMetadata

`func (o *DirectoryServiceIntentResource) GetMetadata() DirectoryServiceMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DirectoryServiceIntentResource) GetMetadataOk() (*DirectoryServiceMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DirectoryServiceIntentResource) SetMetadata(v DirectoryServiceMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



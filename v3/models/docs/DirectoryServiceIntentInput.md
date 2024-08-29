# DirectoryServiceIntentInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Spec** | [**DirectoryService**](DirectoryService.md) |  | 
**ApiVersion** | Pointer to **string** | API Version of the Nutanix v3 API framework. | [optional] [readonly] [default to "3.1.0"]
**Metadata** | [**DirectoryServiceMetadata**](DirectoryServiceMetadata.md) |  | 

## Methods

### NewDirectoryServiceIntentInput

`func NewDirectoryServiceIntentInput(spec DirectoryService, metadata DirectoryServiceMetadata, ) *DirectoryServiceIntentInput`

NewDirectoryServiceIntentInput instantiates a new DirectoryServiceIntentInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDirectoryServiceIntentInputWithDefaults

`func NewDirectoryServiceIntentInputWithDefaults() *DirectoryServiceIntentInput`

NewDirectoryServiceIntentInputWithDefaults instantiates a new DirectoryServiceIntentInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpec

`func (o *DirectoryServiceIntentInput) GetSpec() DirectoryService`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *DirectoryServiceIntentInput) GetSpecOk() (*DirectoryService, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *DirectoryServiceIntentInput) SetSpec(v DirectoryService)`

SetSpec sets Spec field to given value.


### GetApiVersion

`func (o *DirectoryServiceIntentInput) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *DirectoryServiceIntentInput) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *DirectoryServiceIntentInput) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *DirectoryServiceIntentInput) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetMetadata

`func (o *DirectoryServiceIntentInput) GetMetadata() DirectoryServiceMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DirectoryServiceIntentInput) GetMetadataOk() (*DirectoryServiceMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DirectoryServiceIntentInput) SetMetadata(v DirectoryServiceMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



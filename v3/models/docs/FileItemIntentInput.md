# FileItemIntentInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Spec** | [**FileItem**](FileItem.md) |  | 
**ApiVersion** | Pointer to **string** | API Version of the Nutanix v3 API framework. | [optional] [readonly] [default to "3.1.0"]
**Metadata** | [**FileItemMetadata**](FileItemMetadata.md) |  | 

## Methods

### NewFileItemIntentInput

`func NewFileItemIntentInput(spec FileItem, metadata FileItemMetadata, ) *FileItemIntentInput`

NewFileItemIntentInput instantiates a new FileItemIntentInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileItemIntentInputWithDefaults

`func NewFileItemIntentInputWithDefaults() *FileItemIntentInput`

NewFileItemIntentInputWithDefaults instantiates a new FileItemIntentInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpec

`func (o *FileItemIntentInput) GetSpec() FileItem`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *FileItemIntentInput) GetSpecOk() (*FileItem, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *FileItemIntentInput) SetSpec(v FileItem)`

SetSpec sets Spec field to given value.


### GetApiVersion

`func (o *FileItemIntentInput) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *FileItemIntentInput) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *FileItemIntentInput) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *FileItemIntentInput) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetMetadata

`func (o *FileItemIntentInput) GetMetadata() FileItemMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *FileItemIntentInput) GetMetadataOk() (*FileItemMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *FileItemIntentInput) SetMetadata(v FileItemMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



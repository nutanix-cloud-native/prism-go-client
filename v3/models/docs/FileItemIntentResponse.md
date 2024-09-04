# FileItemIntentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**FileItemDefStatus**](FileItemDefStatus.md) |  | [optional] 
**Spec** | Pointer to [**FileItem**](FileItem.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**FileItemMetadata**](FileItemMetadata.md) |  | 

## Methods

### NewFileItemIntentResponse

`func NewFileItemIntentResponse(apiVersion string, metadata FileItemMetadata, ) *FileItemIntentResponse`

NewFileItemIntentResponse instantiates a new FileItemIntentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileItemIntentResponseWithDefaults

`func NewFileItemIntentResponseWithDefaults() *FileItemIntentResponse`

NewFileItemIntentResponseWithDefaults instantiates a new FileItemIntentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *FileItemIntentResponse) GetStatus() FileItemDefStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FileItemIntentResponse) GetStatusOk() (*FileItemDefStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FileItemIntentResponse) SetStatus(v FileItemDefStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FileItemIntentResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSpec

`func (o *FileItemIntentResponse) GetSpec() FileItem`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *FileItemIntentResponse) GetSpecOk() (*FileItem, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *FileItemIntentResponse) SetSpec(v FileItem)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *FileItemIntentResponse) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetApiVersion

`func (o *FileItemIntentResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *FileItemIntentResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *FileItemIntentResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *FileItemIntentResponse) GetMetadata() FileItemMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *FileItemIntentResponse) GetMetadataOk() (*FileItemMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *FileItemIntentResponse) SetMetadata(v FileItemMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



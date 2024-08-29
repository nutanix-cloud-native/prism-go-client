# StoragePolicyIntentResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**StoragePolicyDefStatus**](StoragePolicyDefStatus.md) |  | [optional] 
**Spec** | Pointer to [**StoragePolicy**](StoragePolicy.md) |  | [optional] 
**ApiVersion** | Pointer to **string** | API Version of the Nutanix v3 API framework. | [optional] [readonly] [default to "3.1.0"]
**Metadata** | [**StoragePolicyMetadata**](StoragePolicyMetadata.md) |  | 

## Methods

### NewStoragePolicyIntentResource

`func NewStoragePolicyIntentResource(metadata StoragePolicyMetadata, ) *StoragePolicyIntentResource`

NewStoragePolicyIntentResource instantiates a new StoragePolicyIntentResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoragePolicyIntentResourceWithDefaults

`func NewStoragePolicyIntentResourceWithDefaults() *StoragePolicyIntentResource`

NewStoragePolicyIntentResourceWithDefaults instantiates a new StoragePolicyIntentResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *StoragePolicyIntentResource) GetStatus() StoragePolicyDefStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *StoragePolicyIntentResource) GetStatusOk() (*StoragePolicyDefStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *StoragePolicyIntentResource) SetStatus(v StoragePolicyDefStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *StoragePolicyIntentResource) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSpec

`func (o *StoragePolicyIntentResource) GetSpec() StoragePolicy`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *StoragePolicyIntentResource) GetSpecOk() (*StoragePolicy, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *StoragePolicyIntentResource) SetSpec(v StoragePolicy)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *StoragePolicyIntentResource) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetApiVersion

`func (o *StoragePolicyIntentResource) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *StoragePolicyIntentResource) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *StoragePolicyIntentResource) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *StoragePolicyIntentResource) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetMetadata

`func (o *StoragePolicyIntentResource) GetMetadata() StoragePolicyMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *StoragePolicyIntentResource) GetMetadataOk() (*StoragePolicyMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *StoragePolicyIntentResource) SetMetadata(v StoragePolicyMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



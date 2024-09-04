# StoragePolicyIntentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**StoragePolicyDefStatus**](StoragePolicyDefStatus.md) |  | [optional] 
**Spec** | Pointer to [**StoragePolicy**](StoragePolicy.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**StoragePolicyMetadata**](StoragePolicyMetadata.md) |  | 

## Methods

### NewStoragePolicyIntentResponse

`func NewStoragePolicyIntentResponse(apiVersion string, metadata StoragePolicyMetadata, ) *StoragePolicyIntentResponse`

NewStoragePolicyIntentResponse instantiates a new StoragePolicyIntentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoragePolicyIntentResponseWithDefaults

`func NewStoragePolicyIntentResponseWithDefaults() *StoragePolicyIntentResponse`

NewStoragePolicyIntentResponseWithDefaults instantiates a new StoragePolicyIntentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *StoragePolicyIntentResponse) GetStatus() StoragePolicyDefStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *StoragePolicyIntentResponse) GetStatusOk() (*StoragePolicyDefStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *StoragePolicyIntentResponse) SetStatus(v StoragePolicyDefStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *StoragePolicyIntentResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSpec

`func (o *StoragePolicyIntentResponse) GetSpec() StoragePolicy`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *StoragePolicyIntentResponse) GetSpecOk() (*StoragePolicy, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *StoragePolicyIntentResponse) SetSpec(v StoragePolicy)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *StoragePolicyIntentResponse) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetApiVersion

`func (o *StoragePolicyIntentResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *StoragePolicyIntentResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *StoragePolicyIntentResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *StoragePolicyIntentResponse) GetMetadata() StoragePolicyMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *StoragePolicyIntentResponse) GetMetadataOk() (*StoragePolicyMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *StoragePolicyIntentResponse) SetMetadata(v StoragePolicyMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



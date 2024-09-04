# StoragePolicyIntentInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Spec** | [**StoragePolicy**](StoragePolicy.md) |  | 
**ApiVersion** | Pointer to **string** | API Version of the Nutanix v3 API framework. | [optional] [readonly] [default to "3.1.0"]
**Metadata** | [**StoragePolicyMetadata**](StoragePolicyMetadata.md) |  | 

## Methods

### NewStoragePolicyIntentInput

`func NewStoragePolicyIntentInput(spec StoragePolicy, metadata StoragePolicyMetadata, ) *StoragePolicyIntentInput`

NewStoragePolicyIntentInput instantiates a new StoragePolicyIntentInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoragePolicyIntentInputWithDefaults

`func NewStoragePolicyIntentInputWithDefaults() *StoragePolicyIntentInput`

NewStoragePolicyIntentInputWithDefaults instantiates a new StoragePolicyIntentInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpec

`func (o *StoragePolicyIntentInput) GetSpec() StoragePolicy`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *StoragePolicyIntentInput) GetSpecOk() (*StoragePolicy, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *StoragePolicyIntentInput) SetSpec(v StoragePolicy)`

SetSpec sets Spec field to given value.


### GetApiVersion

`func (o *StoragePolicyIntentInput) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *StoragePolicyIntentInput) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *StoragePolicyIntentInput) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *StoragePolicyIntentInput) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetMetadata

`func (o *StoragePolicyIntentInput) GetMetadata() StoragePolicyMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *StoragePolicyIntentInput) GetMetadataOk() (*StoragePolicyMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *StoragePolicyIntentInput) SetMetadata(v StoragePolicyMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



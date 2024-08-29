# StoragePolicyListIntentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entities** | Pointer to [**[]StoragePolicyIntentResource**](StoragePolicyIntentResource.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**StoragePolicyListMetadataOutput**](StoragePolicyListMetadataOutput.md) |  | 

## Methods

### NewStoragePolicyListIntentResponse

`func NewStoragePolicyListIntentResponse(apiVersion string, metadata StoragePolicyListMetadataOutput, ) *StoragePolicyListIntentResponse`

NewStoragePolicyListIntentResponse instantiates a new StoragePolicyListIntentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoragePolicyListIntentResponseWithDefaults

`func NewStoragePolicyListIntentResponseWithDefaults() *StoragePolicyListIntentResponse`

NewStoragePolicyListIntentResponseWithDefaults instantiates a new StoragePolicyListIntentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntities

`func (o *StoragePolicyListIntentResponse) GetEntities() []StoragePolicyIntentResource`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *StoragePolicyListIntentResponse) GetEntitiesOk() (*[]StoragePolicyIntentResource, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *StoragePolicyListIntentResponse) SetEntities(v []StoragePolicyIntentResource)`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *StoragePolicyListIntentResponse) HasEntities() bool`

HasEntities returns a boolean if a field has been set.

### GetApiVersion

`func (o *StoragePolicyListIntentResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *StoragePolicyListIntentResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *StoragePolicyListIntentResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *StoragePolicyListIntentResponse) GetMetadata() StoragePolicyListMetadataOutput`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *StoragePolicyListIntentResponse) GetMetadataOk() (*StoragePolicyListMetadataOutput, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *StoragePolicyListIntentResponse) SetMetadata(v StoragePolicyListMetadataOutput)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



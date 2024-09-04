# DiskListIntentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entities** | Pointer to [**[]DiskIntentResource**](DiskIntentResource.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**DiskListMetadataOutput**](DiskListMetadataOutput.md) |  | 

## Methods

### NewDiskListIntentResponse

`func NewDiskListIntentResponse(apiVersion string, metadata DiskListMetadataOutput, ) *DiskListIntentResponse`

NewDiskListIntentResponse instantiates a new DiskListIntentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiskListIntentResponseWithDefaults

`func NewDiskListIntentResponseWithDefaults() *DiskListIntentResponse`

NewDiskListIntentResponseWithDefaults instantiates a new DiskListIntentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntities

`func (o *DiskListIntentResponse) GetEntities() []DiskIntentResource`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *DiskListIntentResponse) GetEntitiesOk() (*[]DiskIntentResource, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *DiskListIntentResponse) SetEntities(v []DiskIntentResource)`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *DiskListIntentResponse) HasEntities() bool`

HasEntities returns a boolean if a field has been set.

### GetApiVersion

`func (o *DiskListIntentResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *DiskListIntentResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *DiskListIntentResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *DiskListIntentResponse) GetMetadata() DiskListMetadataOutput`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DiskListIntentResponse) GetMetadataOk() (*DiskListMetadataOutput, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DiskListIntentResponse) SetMetadata(v DiskListMetadataOutput)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



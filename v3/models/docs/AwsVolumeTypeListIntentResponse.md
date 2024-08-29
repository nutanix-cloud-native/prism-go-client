# AwsVolumeTypeListIntentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entities** | Pointer to [**[]AwsVolumeTypeIntentResource**](AwsVolumeTypeIntentResource.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**AwsVolumeTypeListMetadataOutput**](AwsVolumeTypeListMetadataOutput.md) |  | 

## Methods

### NewAwsVolumeTypeListIntentResponse

`func NewAwsVolumeTypeListIntentResponse(apiVersion string, metadata AwsVolumeTypeListMetadataOutput, ) *AwsVolumeTypeListIntentResponse`

NewAwsVolumeTypeListIntentResponse instantiates a new AwsVolumeTypeListIntentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsVolumeTypeListIntentResponseWithDefaults

`func NewAwsVolumeTypeListIntentResponseWithDefaults() *AwsVolumeTypeListIntentResponse`

NewAwsVolumeTypeListIntentResponseWithDefaults instantiates a new AwsVolumeTypeListIntentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntities

`func (o *AwsVolumeTypeListIntentResponse) GetEntities() []AwsVolumeTypeIntentResource`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *AwsVolumeTypeListIntentResponse) GetEntitiesOk() (*[]AwsVolumeTypeIntentResource, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *AwsVolumeTypeListIntentResponse) SetEntities(v []AwsVolumeTypeIntentResource)`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *AwsVolumeTypeListIntentResponse) HasEntities() bool`

HasEntities returns a boolean if a field has been set.

### GetApiVersion

`func (o *AwsVolumeTypeListIntentResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *AwsVolumeTypeListIntentResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *AwsVolumeTypeListIntentResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *AwsVolumeTypeListIntentResponse) GetMetadata() AwsVolumeTypeListMetadataOutput`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AwsVolumeTypeListIntentResponse) GetMetadataOk() (*AwsVolumeTypeListMetadataOutput, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AwsVolumeTypeListIntentResponse) SetMetadata(v AwsVolumeTypeListMetadataOutput)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



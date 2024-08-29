# CloudTrustListIntentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entities** | Pointer to [**[]CloudTrustIntentResource**](CloudTrustIntentResource.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**CloudTrustListMetadataOutput**](CloudTrustListMetadataOutput.md) |  | 

## Methods

### NewCloudTrustListIntentResponse

`func NewCloudTrustListIntentResponse(apiVersion string, metadata CloudTrustListMetadataOutput, ) *CloudTrustListIntentResponse`

NewCloudTrustListIntentResponse instantiates a new CloudTrustListIntentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudTrustListIntentResponseWithDefaults

`func NewCloudTrustListIntentResponseWithDefaults() *CloudTrustListIntentResponse`

NewCloudTrustListIntentResponseWithDefaults instantiates a new CloudTrustListIntentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntities

`func (o *CloudTrustListIntentResponse) GetEntities() []CloudTrustIntentResource`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *CloudTrustListIntentResponse) GetEntitiesOk() (*[]CloudTrustIntentResource, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *CloudTrustListIntentResponse) SetEntities(v []CloudTrustIntentResource)`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *CloudTrustListIntentResponse) HasEntities() bool`

HasEntities returns a boolean if a field has been set.

### GetApiVersion

`func (o *CloudTrustListIntentResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *CloudTrustListIntentResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *CloudTrustListIntentResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *CloudTrustListIntentResponse) GetMetadata() CloudTrustListMetadataOutput`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CloudTrustListIntentResponse) GetMetadataOk() (*CloudTrustListMetadataOutput, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CloudTrustListIntentResponse) SetMetadata(v CloudTrustListMetadataOutput)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



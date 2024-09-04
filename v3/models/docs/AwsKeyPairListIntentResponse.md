# AwsKeyPairListIntentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entities** | Pointer to [**[]AwsKeyPairIntentResource**](AwsKeyPairIntentResource.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**AwsKeyPairListMetadataOutput**](AwsKeyPairListMetadataOutput.md) |  | 

## Methods

### NewAwsKeyPairListIntentResponse

`func NewAwsKeyPairListIntentResponse(apiVersion string, metadata AwsKeyPairListMetadataOutput, ) *AwsKeyPairListIntentResponse`

NewAwsKeyPairListIntentResponse instantiates a new AwsKeyPairListIntentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsKeyPairListIntentResponseWithDefaults

`func NewAwsKeyPairListIntentResponseWithDefaults() *AwsKeyPairListIntentResponse`

NewAwsKeyPairListIntentResponseWithDefaults instantiates a new AwsKeyPairListIntentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntities

`func (o *AwsKeyPairListIntentResponse) GetEntities() []AwsKeyPairIntentResource`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *AwsKeyPairListIntentResponse) GetEntitiesOk() (*[]AwsKeyPairIntentResource, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *AwsKeyPairListIntentResponse) SetEntities(v []AwsKeyPairIntentResource)`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *AwsKeyPairListIntentResponse) HasEntities() bool`

HasEntities returns a boolean if a field has been set.

### GetApiVersion

`func (o *AwsKeyPairListIntentResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *AwsKeyPairListIntentResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *AwsKeyPairListIntentResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *AwsKeyPairListIntentResponse) GetMetadata() AwsKeyPairListMetadataOutput`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AwsKeyPairListIntentResponse) GetMetadataOk() (*AwsKeyPairListMetadataOutput, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AwsKeyPairListIntentResponse) SetMetadata(v AwsKeyPairListMetadataOutput)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



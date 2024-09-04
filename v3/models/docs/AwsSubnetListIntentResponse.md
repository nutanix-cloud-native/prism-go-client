# AwsSubnetListIntentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entities** | Pointer to [**[]AwsSubnetIntentResource**](AwsSubnetIntentResource.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**AwsSubnetListMetadataOutput**](AwsSubnetListMetadataOutput.md) |  | 

## Methods

### NewAwsSubnetListIntentResponse

`func NewAwsSubnetListIntentResponse(apiVersion string, metadata AwsSubnetListMetadataOutput, ) *AwsSubnetListIntentResponse`

NewAwsSubnetListIntentResponse instantiates a new AwsSubnetListIntentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsSubnetListIntentResponseWithDefaults

`func NewAwsSubnetListIntentResponseWithDefaults() *AwsSubnetListIntentResponse`

NewAwsSubnetListIntentResponseWithDefaults instantiates a new AwsSubnetListIntentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntities

`func (o *AwsSubnetListIntentResponse) GetEntities() []AwsSubnetIntentResource`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *AwsSubnetListIntentResponse) GetEntitiesOk() (*[]AwsSubnetIntentResource, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *AwsSubnetListIntentResponse) SetEntities(v []AwsSubnetIntentResource)`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *AwsSubnetListIntentResponse) HasEntities() bool`

HasEntities returns a boolean if a field has been set.

### GetApiVersion

`func (o *AwsSubnetListIntentResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *AwsSubnetListIntentResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *AwsSubnetListIntentResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *AwsSubnetListIntentResponse) GetMetadata() AwsSubnetListMetadataOutput`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AwsSubnetListIntentResponse) GetMetadataOk() (*AwsSubnetListMetadataOutput, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AwsSubnetListIntentResponse) SetMetadata(v AwsSubnetListMetadataOutput)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



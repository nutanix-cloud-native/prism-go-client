# AwsVmListIntentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entities** | Pointer to [**[]AwsVmIntentResource**](AwsVmIntentResource.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**AwsVmListMetadataOutput**](AwsVmListMetadataOutput.md) |  | 

## Methods

### NewAwsVmListIntentResponse

`func NewAwsVmListIntentResponse(apiVersion string, metadata AwsVmListMetadataOutput, ) *AwsVmListIntentResponse`

NewAwsVmListIntentResponse instantiates a new AwsVmListIntentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsVmListIntentResponseWithDefaults

`func NewAwsVmListIntentResponseWithDefaults() *AwsVmListIntentResponse`

NewAwsVmListIntentResponseWithDefaults instantiates a new AwsVmListIntentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntities

`func (o *AwsVmListIntentResponse) GetEntities() []AwsVmIntentResource`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *AwsVmListIntentResponse) GetEntitiesOk() (*[]AwsVmIntentResource, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *AwsVmListIntentResponse) SetEntities(v []AwsVmIntentResource)`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *AwsVmListIntentResponse) HasEntities() bool`

HasEntities returns a boolean if a field has been set.

### GetApiVersion

`func (o *AwsVmListIntentResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *AwsVmListIntentResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *AwsVmListIntentResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *AwsVmListIntentResponse) GetMetadata() AwsVmListMetadataOutput`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AwsVmListIntentResponse) GetMetadataOk() (*AwsVmListMetadataOutput, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AwsVmListIntentResponse) SetMetadata(v AwsVmListMetadataOutput)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



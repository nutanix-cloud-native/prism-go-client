# EnvironmentListIntentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entities** | Pointer to [**[]EnvironmentIntentResource**](EnvironmentIntentResource.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**EnvironmentListMetadataOutput**](EnvironmentListMetadataOutput.md) |  | 

## Methods

### NewEnvironmentListIntentResponse

`func NewEnvironmentListIntentResponse(apiVersion string, metadata EnvironmentListMetadataOutput, ) *EnvironmentListIntentResponse`

NewEnvironmentListIntentResponse instantiates a new EnvironmentListIntentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentListIntentResponseWithDefaults

`func NewEnvironmentListIntentResponseWithDefaults() *EnvironmentListIntentResponse`

NewEnvironmentListIntentResponseWithDefaults instantiates a new EnvironmentListIntentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntities

`func (o *EnvironmentListIntentResponse) GetEntities() []EnvironmentIntentResource`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *EnvironmentListIntentResponse) GetEntitiesOk() (*[]EnvironmentIntentResource, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *EnvironmentListIntentResponse) SetEntities(v []EnvironmentIntentResource)`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *EnvironmentListIntentResponse) HasEntities() bool`

HasEntities returns a boolean if a field has been set.

### GetApiVersion

`func (o *EnvironmentListIntentResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *EnvironmentListIntentResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *EnvironmentListIntentResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *EnvironmentListIntentResponse) GetMetadata() EnvironmentListMetadataOutput`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *EnvironmentListIntentResponse) GetMetadataOk() (*EnvironmentListMetadataOutput, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *EnvironmentListIntentResponse) SetMetadata(v EnvironmentListMetadataOutput)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



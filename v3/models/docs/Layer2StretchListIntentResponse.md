# Layer2StretchListIntentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entities** | Pointer to [**[]Layer2StretchIntentResource**](Layer2StretchIntentResource.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**Layer2StretchListMetadataOutput**](Layer2StretchListMetadataOutput.md) |  | 

## Methods

### NewLayer2StretchListIntentResponse

`func NewLayer2StretchListIntentResponse(apiVersion string, metadata Layer2StretchListMetadataOutput, ) *Layer2StretchListIntentResponse`

NewLayer2StretchListIntentResponse instantiates a new Layer2StretchListIntentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLayer2StretchListIntentResponseWithDefaults

`func NewLayer2StretchListIntentResponseWithDefaults() *Layer2StretchListIntentResponse`

NewLayer2StretchListIntentResponseWithDefaults instantiates a new Layer2StretchListIntentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntities

`func (o *Layer2StretchListIntentResponse) GetEntities() []Layer2StretchIntentResource`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *Layer2StretchListIntentResponse) GetEntitiesOk() (*[]Layer2StretchIntentResource, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *Layer2StretchListIntentResponse) SetEntities(v []Layer2StretchIntentResource)`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *Layer2StretchListIntentResponse) HasEntities() bool`

HasEntities returns a boolean if a field has been set.

### GetApiVersion

`func (o *Layer2StretchListIntentResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *Layer2StretchListIntentResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *Layer2StretchListIntentResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *Layer2StretchListIntentResponse) GetMetadata() Layer2StretchListMetadataOutput`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Layer2StretchListIntentResponse) GetMetadataOk() (*Layer2StretchListMetadataOutput, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Layer2StretchListIntentResponse) SetMetadata(v Layer2StretchListMetadataOutput)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



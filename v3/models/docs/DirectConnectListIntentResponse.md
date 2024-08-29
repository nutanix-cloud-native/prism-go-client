# DirectConnectListIntentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entities** | Pointer to [**[]DirectConnectIntentResource**](DirectConnectIntentResource.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**DirectConnectListMetadataOutput**](DirectConnectListMetadataOutput.md) |  | 

## Methods

### NewDirectConnectListIntentResponse

`func NewDirectConnectListIntentResponse(apiVersion string, metadata DirectConnectListMetadataOutput, ) *DirectConnectListIntentResponse`

NewDirectConnectListIntentResponse instantiates a new DirectConnectListIntentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDirectConnectListIntentResponseWithDefaults

`func NewDirectConnectListIntentResponseWithDefaults() *DirectConnectListIntentResponse`

NewDirectConnectListIntentResponseWithDefaults instantiates a new DirectConnectListIntentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntities

`func (o *DirectConnectListIntentResponse) GetEntities() []DirectConnectIntentResource`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *DirectConnectListIntentResponse) GetEntitiesOk() (*[]DirectConnectIntentResource, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *DirectConnectListIntentResponse) SetEntities(v []DirectConnectIntentResource)`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *DirectConnectListIntentResponse) HasEntities() bool`

HasEntities returns a boolean if a field has been set.

### GetApiVersion

`func (o *DirectConnectListIntentResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *DirectConnectListIntentResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *DirectConnectListIntentResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *DirectConnectListIntentResponse) GetMetadata() DirectConnectListMetadataOutput`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DirectConnectListIntentResponse) GetMetadataOk() (*DirectConnectListMetadataOutput, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DirectConnectListIntentResponse) SetMetadata(v DirectConnectListMetadataOutput)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



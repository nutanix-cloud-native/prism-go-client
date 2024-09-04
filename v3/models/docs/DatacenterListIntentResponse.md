# DatacenterListIntentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entities** | Pointer to [**[]DatacenterIntentResource**](DatacenterIntentResource.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**DatacenterListMetadataOutput**](DatacenterListMetadataOutput.md) |  | 

## Methods

### NewDatacenterListIntentResponse

`func NewDatacenterListIntentResponse(apiVersion string, metadata DatacenterListMetadataOutput, ) *DatacenterListIntentResponse`

NewDatacenterListIntentResponse instantiates a new DatacenterListIntentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatacenterListIntentResponseWithDefaults

`func NewDatacenterListIntentResponseWithDefaults() *DatacenterListIntentResponse`

NewDatacenterListIntentResponseWithDefaults instantiates a new DatacenterListIntentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntities

`func (o *DatacenterListIntentResponse) GetEntities() []DatacenterIntentResource`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *DatacenterListIntentResponse) GetEntitiesOk() (*[]DatacenterIntentResource, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *DatacenterListIntentResponse) SetEntities(v []DatacenterIntentResource)`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *DatacenterListIntentResponse) HasEntities() bool`

HasEntities returns a boolean if a field has been set.

### GetApiVersion

`func (o *DatacenterListIntentResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *DatacenterListIntentResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *DatacenterListIntentResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *DatacenterListIntentResponse) GetMetadata() DatacenterListMetadataOutput`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DatacenterListIntentResponse) GetMetadataOk() (*DatacenterListMetadataOutput, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DatacenterListIntentResponse) SetMetadata(v DatacenterListMetadataOutput)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



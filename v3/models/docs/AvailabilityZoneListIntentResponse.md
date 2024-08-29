# AvailabilityZoneListIntentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entities** | Pointer to [**[]AvailabilityZoneIntentResource**](AvailabilityZoneIntentResource.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**AvailabilityZoneListMetadataOutput**](AvailabilityZoneListMetadataOutput.md) |  | 

## Methods

### NewAvailabilityZoneListIntentResponse

`func NewAvailabilityZoneListIntentResponse(apiVersion string, metadata AvailabilityZoneListMetadataOutput, ) *AvailabilityZoneListIntentResponse`

NewAvailabilityZoneListIntentResponse instantiates a new AvailabilityZoneListIntentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAvailabilityZoneListIntentResponseWithDefaults

`func NewAvailabilityZoneListIntentResponseWithDefaults() *AvailabilityZoneListIntentResponse`

NewAvailabilityZoneListIntentResponseWithDefaults instantiates a new AvailabilityZoneListIntentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntities

`func (o *AvailabilityZoneListIntentResponse) GetEntities() []AvailabilityZoneIntentResource`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *AvailabilityZoneListIntentResponse) GetEntitiesOk() (*[]AvailabilityZoneIntentResource, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *AvailabilityZoneListIntentResponse) SetEntities(v []AvailabilityZoneIntentResource)`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *AvailabilityZoneListIntentResponse) HasEntities() bool`

HasEntities returns a boolean if a field has been set.

### GetApiVersion

`func (o *AvailabilityZoneListIntentResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *AvailabilityZoneListIntentResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *AvailabilityZoneListIntentResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *AvailabilityZoneListIntentResponse) GetMetadata() AvailabilityZoneListMetadataOutput`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AvailabilityZoneListIntentResponse) GetMetadataOk() (*AvailabilityZoneListMetadataOutput, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AvailabilityZoneListIntentResponse) SetMetadata(v AvailabilityZoneListMetadataOutput)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



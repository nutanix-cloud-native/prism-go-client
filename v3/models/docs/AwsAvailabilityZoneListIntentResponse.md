# AwsAvailabilityZoneListIntentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entities** | Pointer to [**[]AwsAvailabilityZoneIntentResource**](AwsAvailabilityZoneIntentResource.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**AwsAvailabilityZoneListMetadataOutput**](AwsAvailabilityZoneListMetadataOutput.md) |  | 

## Methods

### NewAwsAvailabilityZoneListIntentResponse

`func NewAwsAvailabilityZoneListIntentResponse(apiVersion string, metadata AwsAvailabilityZoneListMetadataOutput, ) *AwsAvailabilityZoneListIntentResponse`

NewAwsAvailabilityZoneListIntentResponse instantiates a new AwsAvailabilityZoneListIntentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsAvailabilityZoneListIntentResponseWithDefaults

`func NewAwsAvailabilityZoneListIntentResponseWithDefaults() *AwsAvailabilityZoneListIntentResponse`

NewAwsAvailabilityZoneListIntentResponseWithDefaults instantiates a new AwsAvailabilityZoneListIntentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntities

`func (o *AwsAvailabilityZoneListIntentResponse) GetEntities() []AwsAvailabilityZoneIntentResource`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *AwsAvailabilityZoneListIntentResponse) GetEntitiesOk() (*[]AwsAvailabilityZoneIntentResource, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *AwsAvailabilityZoneListIntentResponse) SetEntities(v []AwsAvailabilityZoneIntentResource)`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *AwsAvailabilityZoneListIntentResponse) HasEntities() bool`

HasEntities returns a boolean if a field has been set.

### GetApiVersion

`func (o *AwsAvailabilityZoneListIntentResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *AwsAvailabilityZoneListIntentResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *AwsAvailabilityZoneListIntentResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *AwsAvailabilityZoneListIntentResponse) GetMetadata() AwsAvailabilityZoneListMetadataOutput`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AwsAvailabilityZoneListIntentResponse) GetMetadataOk() (*AwsAvailabilityZoneListMetadataOutput, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AwsAvailabilityZoneListIntentResponse) SetMetadata(v AwsAvailabilityZoneListMetadataOutput)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



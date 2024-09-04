# AvailabilityZoneIntentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**AvailabilityZoneDefStatus**](AvailabilityZoneDefStatus.md) |  | [optional] 
**Spec** | Pointer to [**AvailabilityZone**](AvailabilityZone.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**AvailabilityZoneMetadata**](AvailabilityZoneMetadata.md) |  | 

## Methods

### NewAvailabilityZoneIntentResponse

`func NewAvailabilityZoneIntentResponse(apiVersion string, metadata AvailabilityZoneMetadata, ) *AvailabilityZoneIntentResponse`

NewAvailabilityZoneIntentResponse instantiates a new AvailabilityZoneIntentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAvailabilityZoneIntentResponseWithDefaults

`func NewAvailabilityZoneIntentResponseWithDefaults() *AvailabilityZoneIntentResponse`

NewAvailabilityZoneIntentResponseWithDefaults instantiates a new AvailabilityZoneIntentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *AvailabilityZoneIntentResponse) GetStatus() AvailabilityZoneDefStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AvailabilityZoneIntentResponse) GetStatusOk() (*AvailabilityZoneDefStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AvailabilityZoneIntentResponse) SetStatus(v AvailabilityZoneDefStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AvailabilityZoneIntentResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSpec

`func (o *AvailabilityZoneIntentResponse) GetSpec() AvailabilityZone`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *AvailabilityZoneIntentResponse) GetSpecOk() (*AvailabilityZone, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *AvailabilityZoneIntentResponse) SetSpec(v AvailabilityZone)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *AvailabilityZoneIntentResponse) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetApiVersion

`func (o *AvailabilityZoneIntentResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *AvailabilityZoneIntentResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *AvailabilityZoneIntentResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *AvailabilityZoneIntentResponse) GetMetadata() AvailabilityZoneMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AvailabilityZoneIntentResponse) GetMetadataOk() (*AvailabilityZoneMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AvailabilityZoneIntentResponse) SetMetadata(v AvailabilityZoneMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



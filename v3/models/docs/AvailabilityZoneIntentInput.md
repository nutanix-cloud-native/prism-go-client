# AvailabilityZoneIntentInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Spec** | [**AvailabilityZone**](AvailabilityZone.md) |  | 
**ApiVersion** | Pointer to **string** | API Version of the Nutanix v3 API framework. | [optional] [readonly] [default to "3.1.0"]
**Metadata** | [**AvailabilityZoneMetadata**](AvailabilityZoneMetadata.md) |  | 

## Methods

### NewAvailabilityZoneIntentInput

`func NewAvailabilityZoneIntentInput(spec AvailabilityZone, metadata AvailabilityZoneMetadata, ) *AvailabilityZoneIntentInput`

NewAvailabilityZoneIntentInput instantiates a new AvailabilityZoneIntentInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAvailabilityZoneIntentInputWithDefaults

`func NewAvailabilityZoneIntentInputWithDefaults() *AvailabilityZoneIntentInput`

NewAvailabilityZoneIntentInputWithDefaults instantiates a new AvailabilityZoneIntentInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpec

`func (o *AvailabilityZoneIntentInput) GetSpec() AvailabilityZone`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *AvailabilityZoneIntentInput) GetSpecOk() (*AvailabilityZone, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *AvailabilityZoneIntentInput) SetSpec(v AvailabilityZone)`

SetSpec sets Spec field to given value.


### GetApiVersion

`func (o *AvailabilityZoneIntentInput) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *AvailabilityZoneIntentInput) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *AvailabilityZoneIntentInput) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *AvailabilityZoneIntentInput) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetMetadata

`func (o *AvailabilityZoneIntentInput) GetMetadata() AvailabilityZoneMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AvailabilityZoneIntentInput) GetMetadataOk() (*AvailabilityZoneMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AvailabilityZoneIntentInput) SetMetadata(v AvailabilityZoneMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



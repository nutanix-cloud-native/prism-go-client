# AwsAvailabilityZoneIntentResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**AwsAvailabilityZoneDefStatus**](AwsAvailabilityZoneDefStatus.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**AwsAvailabilityZoneMetadata**](AwsAvailabilityZoneMetadata.md) |  | 

## Methods

### NewAwsAvailabilityZoneIntentResource

`func NewAwsAvailabilityZoneIntentResource(apiVersion string, metadata AwsAvailabilityZoneMetadata, ) *AwsAvailabilityZoneIntentResource`

NewAwsAvailabilityZoneIntentResource instantiates a new AwsAvailabilityZoneIntentResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsAvailabilityZoneIntentResourceWithDefaults

`func NewAwsAvailabilityZoneIntentResourceWithDefaults() *AwsAvailabilityZoneIntentResource`

NewAwsAvailabilityZoneIntentResourceWithDefaults instantiates a new AwsAvailabilityZoneIntentResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *AwsAvailabilityZoneIntentResource) GetStatus() AwsAvailabilityZoneDefStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AwsAvailabilityZoneIntentResource) GetStatusOk() (*AwsAvailabilityZoneDefStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AwsAvailabilityZoneIntentResource) SetStatus(v AwsAvailabilityZoneDefStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AwsAvailabilityZoneIntentResource) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetApiVersion

`func (o *AwsAvailabilityZoneIntentResource) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *AwsAvailabilityZoneIntentResource) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *AwsAvailabilityZoneIntentResource) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *AwsAvailabilityZoneIntentResource) GetMetadata() AwsAvailabilityZoneMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AwsAvailabilityZoneIntentResource) GetMetadataOk() (*AwsAvailabilityZoneMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AwsAvailabilityZoneIntentResource) SetMetadata(v AwsAvailabilityZoneMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



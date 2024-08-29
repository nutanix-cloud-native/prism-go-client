# DatacenterIntentResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**DatacenterDefStatus**](DatacenterDefStatus.md) |  | [optional] 
**Spec** | Pointer to [**Datacenter**](Datacenter.md) |  | [optional] 
**ApiVersion** | Pointer to **string** | API Version of the Nutanix v3 API framework. | [optional] [readonly] [default to "3.1.0"]
**Metadata** | [**DatacenterMetadata**](DatacenterMetadata.md) |  | 

## Methods

### NewDatacenterIntentResource

`func NewDatacenterIntentResource(metadata DatacenterMetadata, ) *DatacenterIntentResource`

NewDatacenterIntentResource instantiates a new DatacenterIntentResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatacenterIntentResourceWithDefaults

`func NewDatacenterIntentResourceWithDefaults() *DatacenterIntentResource`

NewDatacenterIntentResourceWithDefaults instantiates a new DatacenterIntentResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *DatacenterIntentResource) GetStatus() DatacenterDefStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DatacenterIntentResource) GetStatusOk() (*DatacenterDefStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DatacenterIntentResource) SetStatus(v DatacenterDefStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DatacenterIntentResource) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSpec

`func (o *DatacenterIntentResource) GetSpec() Datacenter`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *DatacenterIntentResource) GetSpecOk() (*Datacenter, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *DatacenterIntentResource) SetSpec(v Datacenter)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *DatacenterIntentResource) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetApiVersion

`func (o *DatacenterIntentResource) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *DatacenterIntentResource) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *DatacenterIntentResource) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *DatacenterIntentResource) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetMetadata

`func (o *DatacenterIntentResource) GetMetadata() DatacenterMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DatacenterIntentResource) GetMetadataOk() (*DatacenterMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DatacenterIntentResource) SetMetadata(v DatacenterMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



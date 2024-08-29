# DatacenterIntentInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Spec** | [**Datacenter**](Datacenter.md) |  | 
**ApiVersion** | Pointer to **string** | API Version of the Nutanix v3 API framework. | [optional] [readonly] [default to "3.1.0"]
**Metadata** | [**DatacenterMetadata**](DatacenterMetadata.md) |  | 

## Methods

### NewDatacenterIntentInput

`func NewDatacenterIntentInput(spec Datacenter, metadata DatacenterMetadata, ) *DatacenterIntentInput`

NewDatacenterIntentInput instantiates a new DatacenterIntentInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatacenterIntentInputWithDefaults

`func NewDatacenterIntentInputWithDefaults() *DatacenterIntentInput`

NewDatacenterIntentInputWithDefaults instantiates a new DatacenterIntentInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpec

`func (o *DatacenterIntentInput) GetSpec() Datacenter`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *DatacenterIntentInput) GetSpecOk() (*Datacenter, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *DatacenterIntentInput) SetSpec(v Datacenter)`

SetSpec sets Spec field to given value.


### GetApiVersion

`func (o *DatacenterIntentInput) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *DatacenterIntentInput) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *DatacenterIntentInput) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *DatacenterIntentInput) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetMetadata

`func (o *DatacenterIntentInput) GetMetadata() DatacenterMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DatacenterIntentInput) GetMetadataOk() (*DatacenterMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DatacenterIntentInput) SetMetadata(v DatacenterMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



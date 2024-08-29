# RackIntentInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Spec** | [**Rack**](Rack.md) |  | 
**ApiVersion** | Pointer to **string** | API Version of the Nutanix v3 API framework. | [optional] [readonly] [default to "3.1.0"]
**Metadata** | [**RackMetadata**](RackMetadata.md) |  | 

## Methods

### NewRackIntentInput

`func NewRackIntentInput(spec Rack, metadata RackMetadata, ) *RackIntentInput`

NewRackIntentInput instantiates a new RackIntentInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRackIntentInputWithDefaults

`func NewRackIntentInputWithDefaults() *RackIntentInput`

NewRackIntentInputWithDefaults instantiates a new RackIntentInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpec

`func (o *RackIntentInput) GetSpec() Rack`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *RackIntentInput) GetSpecOk() (*Rack, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *RackIntentInput) SetSpec(v Rack)`

SetSpec sets Spec field to given value.


### GetApiVersion

`func (o *RackIntentInput) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *RackIntentInput) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *RackIntentInput) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *RackIntentInput) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetMetadata

`func (o *RackIntentInput) GetMetadata() RackMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *RackIntentInput) GetMetadataOk() (*RackMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *RackIntentInput) SetMetadata(v RackMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



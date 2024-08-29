# HostIntentInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Spec** | [**Host**](Host.md) |  | 
**ApiVersion** | Pointer to **string** | API Version of the Nutanix v3 API framework. | [optional] [readonly] [default to "3.1.0"]
**Metadata** | [**HostMetadata**](HostMetadata.md) |  | 

## Methods

### NewHostIntentInput

`func NewHostIntentInput(spec Host, metadata HostMetadata, ) *HostIntentInput`

NewHostIntentInput instantiates a new HostIntentInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostIntentInputWithDefaults

`func NewHostIntentInputWithDefaults() *HostIntentInput`

NewHostIntentInputWithDefaults instantiates a new HostIntentInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpec

`func (o *HostIntentInput) GetSpec() Host`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *HostIntentInput) GetSpecOk() (*Host, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *HostIntentInput) SetSpec(v Host)`

SetSpec sets Spec field to given value.


### GetApiVersion

`func (o *HostIntentInput) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *HostIntentInput) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *HostIntentInput) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *HostIntentInput) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetMetadata

`func (o *HostIntentInput) GetMetadata() HostMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *HostIntentInput) GetMetadataOk() (*HostMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *HostIntentInput) SetMetadata(v HostMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



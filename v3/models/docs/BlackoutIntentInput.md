# BlackoutIntentInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Spec** | [**Blackout**](Blackout.md) |  | 
**ApiVersion** | Pointer to **string** | API Version of the Nutanix v3 API framework. | [optional] [readonly] [default to "3.1.0"]
**Metadata** | [**BlackoutMetadata**](BlackoutMetadata.md) |  | 

## Methods

### NewBlackoutIntentInput

`func NewBlackoutIntentInput(spec Blackout, metadata BlackoutMetadata, ) *BlackoutIntentInput`

NewBlackoutIntentInput instantiates a new BlackoutIntentInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlackoutIntentInputWithDefaults

`func NewBlackoutIntentInputWithDefaults() *BlackoutIntentInput`

NewBlackoutIntentInputWithDefaults instantiates a new BlackoutIntentInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpec

`func (o *BlackoutIntentInput) GetSpec() Blackout`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *BlackoutIntentInput) GetSpecOk() (*Blackout, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *BlackoutIntentInput) SetSpec(v Blackout)`

SetSpec sets Spec field to given value.


### GetApiVersion

`func (o *BlackoutIntentInput) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *BlackoutIntentInput) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *BlackoutIntentInput) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *BlackoutIntentInput) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetMetadata

`func (o *BlackoutIntentInput) GetMetadata() BlackoutMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *BlackoutIntentInput) GetMetadataOk() (*BlackoutMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *BlackoutIntentInput) SetMetadata(v BlackoutMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



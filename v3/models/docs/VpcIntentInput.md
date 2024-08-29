# VpcIntentInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Spec** | [**Vpc**](Vpc.md) |  | 
**ApiVersion** | Pointer to **string** | API Version of the Nutanix v3 API framework. | [optional] [readonly] [default to "3.1.0"]
**Metadata** | [**VpcMetadata**](VpcMetadata.md) |  | 

## Methods

### NewVpcIntentInput

`func NewVpcIntentInput(spec Vpc, metadata VpcMetadata, ) *VpcIntentInput`

NewVpcIntentInput instantiates a new VpcIntentInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpcIntentInputWithDefaults

`func NewVpcIntentInputWithDefaults() *VpcIntentInput`

NewVpcIntentInputWithDefaults instantiates a new VpcIntentInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpec

`func (o *VpcIntentInput) GetSpec() Vpc`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *VpcIntentInput) GetSpecOk() (*Vpc, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *VpcIntentInput) SetSpec(v Vpc)`

SetSpec sets Spec field to given value.


### GetApiVersion

`func (o *VpcIntentInput) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *VpcIntentInput) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *VpcIntentInput) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *VpcIntentInput) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetMetadata

`func (o *VpcIntentInput) GetMetadata() VpcMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *VpcIntentInput) GetMetadataOk() (*VpcMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *VpcIntentInput) SetMetadata(v VpcMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



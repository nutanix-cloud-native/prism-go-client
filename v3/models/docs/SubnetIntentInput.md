# SubnetIntentInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Spec** | [**Subnet**](Subnet.md) |  | 
**ApiVersion** | Pointer to **string** | API Version of the Nutanix v3 API framework. | [optional] [readonly] [default to "3.1.0"]
**Metadata** | [**SubnetMetadata**](SubnetMetadata.md) |  | 

## Methods

### NewSubnetIntentInput

`func NewSubnetIntentInput(spec Subnet, metadata SubnetMetadata, ) *SubnetIntentInput`

NewSubnetIntentInput instantiates a new SubnetIntentInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubnetIntentInputWithDefaults

`func NewSubnetIntentInputWithDefaults() *SubnetIntentInput`

NewSubnetIntentInputWithDefaults instantiates a new SubnetIntentInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpec

`func (o *SubnetIntentInput) GetSpec() Subnet`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *SubnetIntentInput) GetSpecOk() (*Subnet, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *SubnetIntentInput) SetSpec(v Subnet)`

SetSpec sets Spec field to given value.


### GetApiVersion

`func (o *SubnetIntentInput) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *SubnetIntentInput) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *SubnetIntentInput) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *SubnetIntentInput) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetMetadata

`func (o *SubnetIntentInput) GetMetadata() SubnetMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SubnetIntentInput) GetMetadataOk() (*SubnetMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SubnetIntentInput) SetMetadata(v SubnetMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



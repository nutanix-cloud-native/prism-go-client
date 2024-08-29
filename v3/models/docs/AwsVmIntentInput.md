# AwsVmIntentInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Spec** | [**AwsVm**](AwsVm.md) |  | 
**ApiVersion** | Pointer to **string** | API Version of the Nutanix v3 API framework. | [optional] [readonly] [default to "3.1.0"]
**Metadata** | [**AwsVmMetadata**](AwsVmMetadata.md) |  | 

## Methods

### NewAwsVmIntentInput

`func NewAwsVmIntentInput(spec AwsVm, metadata AwsVmMetadata, ) *AwsVmIntentInput`

NewAwsVmIntentInput instantiates a new AwsVmIntentInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsVmIntentInputWithDefaults

`func NewAwsVmIntentInputWithDefaults() *AwsVmIntentInput`

NewAwsVmIntentInputWithDefaults instantiates a new AwsVmIntentInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpec

`func (o *AwsVmIntentInput) GetSpec() AwsVm`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *AwsVmIntentInput) GetSpecOk() (*AwsVm, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *AwsVmIntentInput) SetSpec(v AwsVm)`

SetSpec sets Spec field to given value.


### GetApiVersion

`func (o *AwsVmIntentInput) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *AwsVmIntentInput) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *AwsVmIntentInput) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *AwsVmIntentInput) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetMetadata

`func (o *AwsVmIntentInput) GetMetadata() AwsVmMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AwsVmIntentInput) GetMetadataOk() (*AwsVmMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AwsVmIntentInput) SetMetadata(v AwsVmMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



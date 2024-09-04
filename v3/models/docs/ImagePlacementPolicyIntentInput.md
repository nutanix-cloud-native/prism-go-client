# ImagePlacementPolicyIntentInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Spec** | [**ImagePlacementPolicy**](ImagePlacementPolicy.md) |  | 
**ApiVersion** | Pointer to **string** | API Version of the Nutanix v3 API framework. | [optional] [readonly] [default to "3.1.0"]
**Metadata** | [**ImagePlacementPolicyMetadata**](ImagePlacementPolicyMetadata.md) |  | 

## Methods

### NewImagePlacementPolicyIntentInput

`func NewImagePlacementPolicyIntentInput(spec ImagePlacementPolicy, metadata ImagePlacementPolicyMetadata, ) *ImagePlacementPolicyIntentInput`

NewImagePlacementPolicyIntentInput instantiates a new ImagePlacementPolicyIntentInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImagePlacementPolicyIntentInputWithDefaults

`func NewImagePlacementPolicyIntentInputWithDefaults() *ImagePlacementPolicyIntentInput`

NewImagePlacementPolicyIntentInputWithDefaults instantiates a new ImagePlacementPolicyIntentInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpec

`func (o *ImagePlacementPolicyIntentInput) GetSpec() ImagePlacementPolicy`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *ImagePlacementPolicyIntentInput) GetSpecOk() (*ImagePlacementPolicy, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *ImagePlacementPolicyIntentInput) SetSpec(v ImagePlacementPolicy)`

SetSpec sets Spec field to given value.


### GetApiVersion

`func (o *ImagePlacementPolicyIntentInput) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ImagePlacementPolicyIntentInput) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ImagePlacementPolicyIntentInput) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *ImagePlacementPolicyIntentInput) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetMetadata

`func (o *ImagePlacementPolicyIntentInput) GetMetadata() ImagePlacementPolicyMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ImagePlacementPolicyIntentInput) GetMetadataOk() (*ImagePlacementPolicyMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ImagePlacementPolicyIntentInput) SetMetadata(v ImagePlacementPolicyMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



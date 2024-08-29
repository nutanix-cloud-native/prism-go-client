# CloudTrustIntentInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Spec** | [**CloudTrust**](CloudTrust.md) |  | 
**ApiVersion** | Pointer to **string** | API Version of the Nutanix v3 API framework. | [optional] [readonly] [default to "3.1.0"]
**Metadata** | [**CloudTrustMetadata**](CloudTrustMetadata.md) |  | 

## Methods

### NewCloudTrustIntentInput

`func NewCloudTrustIntentInput(spec CloudTrust, metadata CloudTrustMetadata, ) *CloudTrustIntentInput`

NewCloudTrustIntentInput instantiates a new CloudTrustIntentInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudTrustIntentInputWithDefaults

`func NewCloudTrustIntentInputWithDefaults() *CloudTrustIntentInput`

NewCloudTrustIntentInputWithDefaults instantiates a new CloudTrustIntentInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpec

`func (o *CloudTrustIntentInput) GetSpec() CloudTrust`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *CloudTrustIntentInput) GetSpecOk() (*CloudTrust, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *CloudTrustIntentInput) SetSpec(v CloudTrust)`

SetSpec sets Spec field to given value.


### GetApiVersion

`func (o *CloudTrustIntentInput) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *CloudTrustIntentInput) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *CloudTrustIntentInput) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *CloudTrustIntentInput) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetMetadata

`func (o *CloudTrustIntentInput) GetMetadata() CloudTrustMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CloudTrustIntentInput) GetMetadataOk() (*CloudTrustMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CloudTrustIntentInput) SetMetadata(v CloudTrustMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



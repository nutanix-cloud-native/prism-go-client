# CloudCredentialsIntentInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Spec** | [**CloudCredentials**](CloudCredentials.md) |  | 
**ApiVersion** | Pointer to **string** | API Version of the Nutanix v3 API framework. | [optional] [readonly] [default to "3.1.0"]
**Metadata** | [**CloudCredentialsMetadata**](CloudCredentialsMetadata.md) |  | 

## Methods

### NewCloudCredentialsIntentInput

`func NewCloudCredentialsIntentInput(spec CloudCredentials, metadata CloudCredentialsMetadata, ) *CloudCredentialsIntentInput`

NewCloudCredentialsIntentInput instantiates a new CloudCredentialsIntentInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudCredentialsIntentInputWithDefaults

`func NewCloudCredentialsIntentInputWithDefaults() *CloudCredentialsIntentInput`

NewCloudCredentialsIntentInputWithDefaults instantiates a new CloudCredentialsIntentInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpec

`func (o *CloudCredentialsIntentInput) GetSpec() CloudCredentials`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *CloudCredentialsIntentInput) GetSpecOk() (*CloudCredentials, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *CloudCredentialsIntentInput) SetSpec(v CloudCredentials)`

SetSpec sets Spec field to given value.


### GetApiVersion

`func (o *CloudCredentialsIntentInput) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *CloudCredentialsIntentInput) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *CloudCredentialsIntentInput) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *CloudCredentialsIntentInput) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetMetadata

`func (o *CloudCredentialsIntentInput) GetMetadata() CloudCredentialsMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CloudCredentialsIntentInput) GetMetadataOk() (*CloudCredentialsMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CloudCredentialsIntentInput) SetMetadata(v CloudCredentialsMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



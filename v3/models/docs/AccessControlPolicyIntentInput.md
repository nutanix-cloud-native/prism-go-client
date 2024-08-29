# AccessControlPolicyIntentInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Spec** | [**AccessControlPolicy**](AccessControlPolicy.md) |  | 
**ApiVersion** | Pointer to **string** | API Version of the Nutanix v3 API framework. | [optional] [readonly] [default to "3.1.0"]
**Metadata** | [**AccessControlPolicyMetadata**](AccessControlPolicyMetadata.md) |  | 

## Methods

### NewAccessControlPolicyIntentInput

`func NewAccessControlPolicyIntentInput(spec AccessControlPolicy, metadata AccessControlPolicyMetadata, ) *AccessControlPolicyIntentInput`

NewAccessControlPolicyIntentInput instantiates a new AccessControlPolicyIntentInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessControlPolicyIntentInputWithDefaults

`func NewAccessControlPolicyIntentInputWithDefaults() *AccessControlPolicyIntentInput`

NewAccessControlPolicyIntentInputWithDefaults instantiates a new AccessControlPolicyIntentInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpec

`func (o *AccessControlPolicyIntentInput) GetSpec() AccessControlPolicy`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *AccessControlPolicyIntentInput) GetSpecOk() (*AccessControlPolicy, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *AccessControlPolicyIntentInput) SetSpec(v AccessControlPolicy)`

SetSpec sets Spec field to given value.


### GetApiVersion

`func (o *AccessControlPolicyIntentInput) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *AccessControlPolicyIntentInput) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *AccessControlPolicyIntentInput) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *AccessControlPolicyIntentInput) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetMetadata

`func (o *AccessControlPolicyIntentInput) GetMetadata() AccessControlPolicyMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AccessControlPolicyIntentInput) GetMetadataOk() (*AccessControlPolicyMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AccessControlPolicyIntentInput) SetMetadata(v AccessControlPolicyMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



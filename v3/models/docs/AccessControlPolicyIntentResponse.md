# AccessControlPolicyIntentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**AccessControlPolicyDefStatus**](AccessControlPolicyDefStatus.md) |  | [optional] 
**Spec** | Pointer to [**AccessControlPolicy**](AccessControlPolicy.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**AccessControlPolicyMetadata**](AccessControlPolicyMetadata.md) |  | 

## Methods

### NewAccessControlPolicyIntentResponse

`func NewAccessControlPolicyIntentResponse(apiVersion string, metadata AccessControlPolicyMetadata, ) *AccessControlPolicyIntentResponse`

NewAccessControlPolicyIntentResponse instantiates a new AccessControlPolicyIntentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessControlPolicyIntentResponseWithDefaults

`func NewAccessControlPolicyIntentResponseWithDefaults() *AccessControlPolicyIntentResponse`

NewAccessControlPolicyIntentResponseWithDefaults instantiates a new AccessControlPolicyIntentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *AccessControlPolicyIntentResponse) GetStatus() AccessControlPolicyDefStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AccessControlPolicyIntentResponse) GetStatusOk() (*AccessControlPolicyDefStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AccessControlPolicyIntentResponse) SetStatus(v AccessControlPolicyDefStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AccessControlPolicyIntentResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSpec

`func (o *AccessControlPolicyIntentResponse) GetSpec() AccessControlPolicy`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *AccessControlPolicyIntentResponse) GetSpecOk() (*AccessControlPolicy, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *AccessControlPolicyIntentResponse) SetSpec(v AccessControlPolicy)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *AccessControlPolicyIntentResponse) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetApiVersion

`func (o *AccessControlPolicyIntentResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *AccessControlPolicyIntentResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *AccessControlPolicyIntentResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *AccessControlPolicyIntentResponse) GetMetadata() AccessControlPolicyMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AccessControlPolicyIntentResponse) GetMetadataOk() (*AccessControlPolicyMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AccessControlPolicyIntentResponse) SetMetadata(v AccessControlPolicyMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



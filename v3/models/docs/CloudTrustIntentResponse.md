# CloudTrustIntentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**CloudTrustDefStatus**](CloudTrustDefStatus.md) |  | [optional] 
**Spec** | Pointer to [**CloudTrust**](CloudTrust.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**CloudTrustMetadata**](CloudTrustMetadata.md) |  | 

## Methods

### NewCloudTrustIntentResponse

`func NewCloudTrustIntentResponse(apiVersion string, metadata CloudTrustMetadata, ) *CloudTrustIntentResponse`

NewCloudTrustIntentResponse instantiates a new CloudTrustIntentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudTrustIntentResponseWithDefaults

`func NewCloudTrustIntentResponseWithDefaults() *CloudTrustIntentResponse`

NewCloudTrustIntentResponseWithDefaults instantiates a new CloudTrustIntentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *CloudTrustIntentResponse) GetStatus() CloudTrustDefStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudTrustIntentResponse) GetStatusOk() (*CloudTrustDefStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudTrustIntentResponse) SetStatus(v CloudTrustDefStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CloudTrustIntentResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSpec

`func (o *CloudTrustIntentResponse) GetSpec() CloudTrust`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *CloudTrustIntentResponse) GetSpecOk() (*CloudTrust, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *CloudTrustIntentResponse) SetSpec(v CloudTrust)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *CloudTrustIntentResponse) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetApiVersion

`func (o *CloudTrustIntentResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *CloudTrustIntentResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *CloudTrustIntentResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *CloudTrustIntentResponse) GetMetadata() CloudTrustMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CloudTrustIntentResponse) GetMetadataOk() (*CloudTrustMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CloudTrustIntentResponse) SetMetadata(v CloudTrustMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# AwsVmIntentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**AwsVmDefStatus**](AwsVmDefStatus.md) |  | [optional] 
**Spec** | Pointer to [**AwsVm**](AwsVm.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**AwsVmMetadata**](AwsVmMetadata.md) |  | 

## Methods

### NewAwsVmIntentResponse

`func NewAwsVmIntentResponse(apiVersion string, metadata AwsVmMetadata, ) *AwsVmIntentResponse`

NewAwsVmIntentResponse instantiates a new AwsVmIntentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsVmIntentResponseWithDefaults

`func NewAwsVmIntentResponseWithDefaults() *AwsVmIntentResponse`

NewAwsVmIntentResponseWithDefaults instantiates a new AwsVmIntentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *AwsVmIntentResponse) GetStatus() AwsVmDefStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AwsVmIntentResponse) GetStatusOk() (*AwsVmDefStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AwsVmIntentResponse) SetStatus(v AwsVmDefStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AwsVmIntentResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSpec

`func (o *AwsVmIntentResponse) GetSpec() AwsVm`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *AwsVmIntentResponse) GetSpecOk() (*AwsVm, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *AwsVmIntentResponse) SetSpec(v AwsVm)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *AwsVmIntentResponse) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetApiVersion

`func (o *AwsVmIntentResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *AwsVmIntentResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *AwsVmIntentResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *AwsVmIntentResponse) GetMetadata() AwsVmMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AwsVmIntentResponse) GetMetadataOk() (*AwsVmMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AwsVmIntentResponse) SetMetadata(v AwsVmMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



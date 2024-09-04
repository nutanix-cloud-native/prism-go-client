# MhVmIntentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**MhVmDefStatus**](MhVmDefStatus.md) |  | [optional] 
**Spec** | Pointer to [**MhVm**](MhVm.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**MhVmMetadata**](MhVmMetadata.md) |  | 

## Methods

### NewMhVmIntentResponse

`func NewMhVmIntentResponse(apiVersion string, metadata MhVmMetadata, ) *MhVmIntentResponse`

NewMhVmIntentResponse instantiates a new MhVmIntentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMhVmIntentResponseWithDefaults

`func NewMhVmIntentResponseWithDefaults() *MhVmIntentResponse`

NewMhVmIntentResponseWithDefaults instantiates a new MhVmIntentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *MhVmIntentResponse) GetStatus() MhVmDefStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MhVmIntentResponse) GetStatusOk() (*MhVmDefStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MhVmIntentResponse) SetStatus(v MhVmDefStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MhVmIntentResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSpec

`func (o *MhVmIntentResponse) GetSpec() MhVm`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *MhVmIntentResponse) GetSpecOk() (*MhVm, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *MhVmIntentResponse) SetSpec(v MhVm)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *MhVmIntentResponse) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetApiVersion

`func (o *MhVmIntentResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *MhVmIntentResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *MhVmIntentResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *MhVmIntentResponse) GetMetadata() MhVmMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *MhVmIntentResponse) GetMetadataOk() (*MhVmMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *MhVmIntentResponse) SetMetadata(v MhVmMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



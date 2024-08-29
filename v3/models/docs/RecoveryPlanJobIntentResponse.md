# RecoveryPlanJobIntentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**RecoveryPlanJobDefStatus**](RecoveryPlanJobDefStatus.md) |  | [optional] 
**Spec** | Pointer to [**RecoveryPlanJob**](RecoveryPlanJob.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**RecoveryPlanJobMetadata**](RecoveryPlanJobMetadata.md) |  | 

## Methods

### NewRecoveryPlanJobIntentResponse

`func NewRecoveryPlanJobIntentResponse(apiVersion string, metadata RecoveryPlanJobMetadata, ) *RecoveryPlanJobIntentResponse`

NewRecoveryPlanJobIntentResponse instantiates a new RecoveryPlanJobIntentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoveryPlanJobIntentResponseWithDefaults

`func NewRecoveryPlanJobIntentResponseWithDefaults() *RecoveryPlanJobIntentResponse`

NewRecoveryPlanJobIntentResponseWithDefaults instantiates a new RecoveryPlanJobIntentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *RecoveryPlanJobIntentResponse) GetStatus() RecoveryPlanJobDefStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RecoveryPlanJobIntentResponse) GetStatusOk() (*RecoveryPlanJobDefStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RecoveryPlanJobIntentResponse) SetStatus(v RecoveryPlanJobDefStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RecoveryPlanJobIntentResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSpec

`func (o *RecoveryPlanJobIntentResponse) GetSpec() RecoveryPlanJob`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *RecoveryPlanJobIntentResponse) GetSpecOk() (*RecoveryPlanJob, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *RecoveryPlanJobIntentResponse) SetSpec(v RecoveryPlanJob)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *RecoveryPlanJobIntentResponse) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetApiVersion

`func (o *RecoveryPlanJobIntentResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *RecoveryPlanJobIntentResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *RecoveryPlanJobIntentResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *RecoveryPlanJobIntentResponse) GetMetadata() RecoveryPlanJobMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *RecoveryPlanJobIntentResponse) GetMetadataOk() (*RecoveryPlanJobMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *RecoveryPlanJobIntentResponse) SetMetadata(v RecoveryPlanJobMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



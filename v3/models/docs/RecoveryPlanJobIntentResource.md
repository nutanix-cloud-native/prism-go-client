# RecoveryPlanJobIntentResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**RecoveryPlanJobDefStatus**](RecoveryPlanJobDefStatus.md) |  | [optional] 
**Spec** | Pointer to [**RecoveryPlanJob**](RecoveryPlanJob.md) |  | [optional] 
**ApiVersion** | Pointer to **string** | API Version of the Nutanix v3 API framework. | [optional] [readonly] [default to "3.1.0"]
**Metadata** | [**RecoveryPlanJobMetadata**](RecoveryPlanJobMetadata.md) |  | 

## Methods

### NewRecoveryPlanJobIntentResource

`func NewRecoveryPlanJobIntentResource(metadata RecoveryPlanJobMetadata, ) *RecoveryPlanJobIntentResource`

NewRecoveryPlanJobIntentResource instantiates a new RecoveryPlanJobIntentResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoveryPlanJobIntentResourceWithDefaults

`func NewRecoveryPlanJobIntentResourceWithDefaults() *RecoveryPlanJobIntentResource`

NewRecoveryPlanJobIntentResourceWithDefaults instantiates a new RecoveryPlanJobIntentResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *RecoveryPlanJobIntentResource) GetStatus() RecoveryPlanJobDefStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RecoveryPlanJobIntentResource) GetStatusOk() (*RecoveryPlanJobDefStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RecoveryPlanJobIntentResource) SetStatus(v RecoveryPlanJobDefStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RecoveryPlanJobIntentResource) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSpec

`func (o *RecoveryPlanJobIntentResource) GetSpec() RecoveryPlanJob`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *RecoveryPlanJobIntentResource) GetSpecOk() (*RecoveryPlanJob, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *RecoveryPlanJobIntentResource) SetSpec(v RecoveryPlanJob)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *RecoveryPlanJobIntentResource) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetApiVersion

`func (o *RecoveryPlanJobIntentResource) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *RecoveryPlanJobIntentResource) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *RecoveryPlanJobIntentResource) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *RecoveryPlanJobIntentResource) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetMetadata

`func (o *RecoveryPlanJobIntentResource) GetMetadata() RecoveryPlanJobMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *RecoveryPlanJobIntentResource) GetMetadataOk() (*RecoveryPlanJobMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *RecoveryPlanJobIntentResource) SetMetadata(v RecoveryPlanJobMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



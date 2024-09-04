# RecoveryPlanIntentResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**RecoveryPlanDefStatus**](RecoveryPlanDefStatus.md) |  | [optional] 
**Spec** | Pointer to [**RecoveryPlan**](RecoveryPlan.md) |  | [optional] 
**ApiVersion** | Pointer to **string** | API Version of the Nutanix v3 API framework. | [optional] [readonly] [default to "3.1.0"]
**Metadata** | [**RecoveryPlanMetadata**](RecoveryPlanMetadata.md) |  | 

## Methods

### NewRecoveryPlanIntentResource

`func NewRecoveryPlanIntentResource(metadata RecoveryPlanMetadata, ) *RecoveryPlanIntentResource`

NewRecoveryPlanIntentResource instantiates a new RecoveryPlanIntentResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoveryPlanIntentResourceWithDefaults

`func NewRecoveryPlanIntentResourceWithDefaults() *RecoveryPlanIntentResource`

NewRecoveryPlanIntentResourceWithDefaults instantiates a new RecoveryPlanIntentResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *RecoveryPlanIntentResource) GetStatus() RecoveryPlanDefStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RecoveryPlanIntentResource) GetStatusOk() (*RecoveryPlanDefStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RecoveryPlanIntentResource) SetStatus(v RecoveryPlanDefStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RecoveryPlanIntentResource) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSpec

`func (o *RecoveryPlanIntentResource) GetSpec() RecoveryPlan`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *RecoveryPlanIntentResource) GetSpecOk() (*RecoveryPlan, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *RecoveryPlanIntentResource) SetSpec(v RecoveryPlan)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *RecoveryPlanIntentResource) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetApiVersion

`func (o *RecoveryPlanIntentResource) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *RecoveryPlanIntentResource) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *RecoveryPlanIntentResource) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *RecoveryPlanIntentResource) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetMetadata

`func (o *RecoveryPlanIntentResource) GetMetadata() RecoveryPlanMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *RecoveryPlanIntentResource) GetMetadataOk() (*RecoveryPlanMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *RecoveryPlanIntentResource) SetMetadata(v RecoveryPlanMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



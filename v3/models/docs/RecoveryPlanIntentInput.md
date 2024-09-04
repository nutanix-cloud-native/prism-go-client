# RecoveryPlanIntentInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Spec** | [**RecoveryPlan**](RecoveryPlan.md) |  | 
**ApiVersion** | Pointer to **string** | API Version of the Nutanix v3 API framework. | [optional] [readonly] [default to "3.1.0"]
**Metadata** | [**RecoveryPlanMetadata**](RecoveryPlanMetadata.md) |  | 

## Methods

### NewRecoveryPlanIntentInput

`func NewRecoveryPlanIntentInput(spec RecoveryPlan, metadata RecoveryPlanMetadata, ) *RecoveryPlanIntentInput`

NewRecoveryPlanIntentInput instantiates a new RecoveryPlanIntentInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoveryPlanIntentInputWithDefaults

`func NewRecoveryPlanIntentInputWithDefaults() *RecoveryPlanIntentInput`

NewRecoveryPlanIntentInputWithDefaults instantiates a new RecoveryPlanIntentInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpec

`func (o *RecoveryPlanIntentInput) GetSpec() RecoveryPlan`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *RecoveryPlanIntentInput) GetSpecOk() (*RecoveryPlan, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *RecoveryPlanIntentInput) SetSpec(v RecoveryPlan)`

SetSpec sets Spec field to given value.


### GetApiVersion

`func (o *RecoveryPlanIntentInput) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *RecoveryPlanIntentInput) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *RecoveryPlanIntentInput) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *RecoveryPlanIntentInput) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetMetadata

`func (o *RecoveryPlanIntentInput) GetMetadata() RecoveryPlanMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *RecoveryPlanIntentInput) GetMetadataOk() (*RecoveryPlanMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *RecoveryPlanIntentInput) SetMetadata(v RecoveryPlanMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



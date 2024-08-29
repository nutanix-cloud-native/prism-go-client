# RecoveryPlanJobIntentInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Spec** | [**RecoveryPlanJob**](RecoveryPlanJob.md) |  | 
**ApiVersion** | Pointer to **string** | API Version of the Nutanix v3 API framework. | [optional] [readonly] [default to "3.1.0"]
**Metadata** | [**RecoveryPlanJobMetadata**](RecoveryPlanJobMetadata.md) |  | 

## Methods

### NewRecoveryPlanJobIntentInput

`func NewRecoveryPlanJobIntentInput(spec RecoveryPlanJob, metadata RecoveryPlanJobMetadata, ) *RecoveryPlanJobIntentInput`

NewRecoveryPlanJobIntentInput instantiates a new RecoveryPlanJobIntentInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoveryPlanJobIntentInputWithDefaults

`func NewRecoveryPlanJobIntentInputWithDefaults() *RecoveryPlanJobIntentInput`

NewRecoveryPlanJobIntentInputWithDefaults instantiates a new RecoveryPlanJobIntentInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpec

`func (o *RecoveryPlanJobIntentInput) GetSpec() RecoveryPlanJob`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *RecoveryPlanJobIntentInput) GetSpecOk() (*RecoveryPlanJob, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *RecoveryPlanJobIntentInput) SetSpec(v RecoveryPlanJob)`

SetSpec sets Spec field to given value.


### GetApiVersion

`func (o *RecoveryPlanJobIntentInput) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *RecoveryPlanJobIntentInput) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *RecoveryPlanJobIntentInput) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *RecoveryPlanJobIntentInput) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetMetadata

`func (o *RecoveryPlanJobIntentInput) GetMetadata() RecoveryPlanJobMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *RecoveryPlanJobIntentInput) GetMetadataOk() (*RecoveryPlanJobMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *RecoveryPlanJobIntentInput) SetMetadata(v RecoveryPlanJobMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



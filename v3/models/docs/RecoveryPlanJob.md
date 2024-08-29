# RecoveryPlanJob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Recovery Plan Job name. | 
**Resources** | [**RecoveryPlanJobResources**](RecoveryPlanJobResources.md) |  | 

## Methods

### NewRecoveryPlanJob

`func NewRecoveryPlanJob(name string, resources RecoveryPlanJobResources, ) *RecoveryPlanJob`

NewRecoveryPlanJob instantiates a new RecoveryPlanJob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoveryPlanJobWithDefaults

`func NewRecoveryPlanJobWithDefaults() *RecoveryPlanJob`

NewRecoveryPlanJobWithDefaults instantiates a new RecoveryPlanJob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RecoveryPlanJob) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RecoveryPlanJob) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RecoveryPlanJob) SetName(v string)`

SetName sets Name field to given value.


### GetResources

`func (o *RecoveryPlanJob) GetResources() RecoveryPlanJobResources`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *RecoveryPlanJob) GetResourcesOk() (*RecoveryPlanJobResources, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *RecoveryPlanJob) SetResources(v RecoveryPlanJobResources)`

SetResources sets Resources field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



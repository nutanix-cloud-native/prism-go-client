# RecoveryPlan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Recovery Plan name | 
**Resources** | [**RecoveryPlanResources**](RecoveryPlanResources.md) |  | 
**Description** | Pointer to **string** | A description or user annotation for the Recovery Plan. | [optional] 

## Methods

### NewRecoveryPlan

`func NewRecoveryPlan(name string, resources RecoveryPlanResources, ) *RecoveryPlan`

NewRecoveryPlan instantiates a new RecoveryPlan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoveryPlanWithDefaults

`func NewRecoveryPlanWithDefaults() *RecoveryPlan`

NewRecoveryPlanWithDefaults instantiates a new RecoveryPlan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RecoveryPlan) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RecoveryPlan) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RecoveryPlan) SetName(v string)`

SetName sets Name field to given value.


### GetResources

`func (o *RecoveryPlan) GetResources() RecoveryPlanResources`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *RecoveryPlan) GetResourcesOk() (*RecoveryPlanResources, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *RecoveryPlan) SetResources(v RecoveryPlanResources)`

SetResources sets Resources field to given value.


### GetDescription

`func (o *RecoveryPlan) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RecoveryPlan) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RecoveryPlan) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RecoveryPlan) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# RecoveryPlanFloatingIpConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ip** | Pointer to **string** | IP to be assigned to VM, in case of failover.  | [optional] 
**ShouldAllocateDynamically** | Pointer to **bool** | Whether to allocate the floating IPs for the VMs dynamically.  | [optional] 

## Methods

### NewRecoveryPlanFloatingIpConfig

`func NewRecoveryPlanFloatingIpConfig() *RecoveryPlanFloatingIpConfig`

NewRecoveryPlanFloatingIpConfig instantiates a new RecoveryPlanFloatingIpConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoveryPlanFloatingIpConfigWithDefaults

`func NewRecoveryPlanFloatingIpConfigWithDefaults() *RecoveryPlanFloatingIpConfig`

NewRecoveryPlanFloatingIpConfigWithDefaults instantiates a new RecoveryPlanFloatingIpConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIp

`func (o *RecoveryPlanFloatingIpConfig) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *RecoveryPlanFloatingIpConfig) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *RecoveryPlanFloatingIpConfig) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *RecoveryPlanFloatingIpConfig) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetShouldAllocateDynamically

`func (o *RecoveryPlanFloatingIpConfig) GetShouldAllocateDynamically() bool`

GetShouldAllocateDynamically returns the ShouldAllocateDynamically field if non-nil, zero value otherwise.

### GetShouldAllocateDynamicallyOk

`func (o *RecoveryPlanFloatingIpConfig) GetShouldAllocateDynamicallyOk() (*bool, bool)`

GetShouldAllocateDynamicallyOk returns a tuple with the ShouldAllocateDynamically field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldAllocateDynamically

`func (o *RecoveryPlanFloatingIpConfig) SetShouldAllocateDynamically(v bool)`

SetShouldAllocateDynamically sets ShouldAllocateDynamically field to given value.

### HasShouldAllocateDynamically

`func (o *RecoveryPlanFloatingIpConfig) HasShouldAllocateDynamically() bool`

HasShouldAllocateDynamically returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



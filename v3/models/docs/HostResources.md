# HostResources

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WindowsDomain** | Pointer to [**WindowsDomain**](WindowsDomain.md) |  | [optional] 
**ControllerVm** | Pointer to [**ControllerVm**](ControllerVm.md) |  | [optional] 
**FailoverCluster** | Pointer to [**FailoverCluster**](FailoverCluster.md) |  | [optional] 

## Methods

### NewHostResources

`func NewHostResources() *HostResources`

NewHostResources instantiates a new HostResources object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostResourcesWithDefaults

`func NewHostResourcesWithDefaults() *HostResources`

NewHostResourcesWithDefaults instantiates a new HostResources object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWindowsDomain

`func (o *HostResources) GetWindowsDomain() WindowsDomain`

GetWindowsDomain returns the WindowsDomain field if non-nil, zero value otherwise.

### GetWindowsDomainOk

`func (o *HostResources) GetWindowsDomainOk() (*WindowsDomain, bool)`

GetWindowsDomainOk returns a tuple with the WindowsDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowsDomain

`func (o *HostResources) SetWindowsDomain(v WindowsDomain)`

SetWindowsDomain sets WindowsDomain field to given value.

### HasWindowsDomain

`func (o *HostResources) HasWindowsDomain() bool`

HasWindowsDomain returns a boolean if a field has been set.

### GetControllerVm

`func (o *HostResources) GetControllerVm() ControllerVm`

GetControllerVm returns the ControllerVm field if non-nil, zero value otherwise.

### GetControllerVmOk

`func (o *HostResources) GetControllerVmOk() (*ControllerVm, bool)`

GetControllerVmOk returns a tuple with the ControllerVm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerVm

`func (o *HostResources) SetControllerVm(v ControllerVm)`

SetControllerVm sets ControllerVm field to given value.

### HasControllerVm

`func (o *HostResources) HasControllerVm() bool`

HasControllerVm returns a boolean if a field has been set.

### GetFailoverCluster

`func (o *HostResources) GetFailoverCluster() FailoverCluster`

GetFailoverCluster returns the FailoverCluster field if non-nil, zero value otherwise.

### GetFailoverClusterOk

`func (o *HostResources) GetFailoverClusterOk() (*FailoverCluster, bool)`

GetFailoverClusterOk returns a tuple with the FailoverCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverCluster

`func (o *HostResources) SetFailoverCluster(v FailoverCluster)`

SetFailoverCluster sets FailoverCluster field to given value.

### HasFailoverCluster

`func (o *HostResources) HasFailoverCluster() bool`

HasFailoverCluster returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



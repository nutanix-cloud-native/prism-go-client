# ControllerVm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ip** | **string** | Controller VM IP address. | [readonly] 
**NatIp** | Pointer to **string** | Controller VM NAT IP address. | [optional] 
**OplogUsage** | Pointer to [**OplogUsage**](OplogUsage.md) |  | [optional] 
**NatPort** | Pointer to **int32** | Controller VM NAT port. | [optional] 

## Methods

### NewControllerVm

`func NewControllerVm(ip string, ) *ControllerVm`

NewControllerVm instantiates a new ControllerVm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewControllerVmWithDefaults

`func NewControllerVmWithDefaults() *ControllerVm`

NewControllerVmWithDefaults instantiates a new ControllerVm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIp

`func (o *ControllerVm) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *ControllerVm) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *ControllerVm) SetIp(v string)`

SetIp sets Ip field to given value.


### GetNatIp

`func (o *ControllerVm) GetNatIp() string`

GetNatIp returns the NatIp field if non-nil, zero value otherwise.

### GetNatIpOk

`func (o *ControllerVm) GetNatIpOk() (*string, bool)`

GetNatIpOk returns a tuple with the NatIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatIp

`func (o *ControllerVm) SetNatIp(v string)`

SetNatIp sets NatIp field to given value.

### HasNatIp

`func (o *ControllerVm) HasNatIp() bool`

HasNatIp returns a boolean if a field has been set.

### GetOplogUsage

`func (o *ControllerVm) GetOplogUsage() OplogUsage`

GetOplogUsage returns the OplogUsage field if non-nil, zero value otherwise.

### GetOplogUsageOk

`func (o *ControllerVm) GetOplogUsageOk() (*OplogUsage, bool)`

GetOplogUsageOk returns a tuple with the OplogUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOplogUsage

`func (o *ControllerVm) SetOplogUsage(v OplogUsage)`

SetOplogUsage sets OplogUsage field to given value.

### HasOplogUsage

`func (o *ControllerVm) HasOplogUsage() bool`

HasOplogUsage returns a boolean if a field has been set.

### GetNatPort

`func (o *ControllerVm) GetNatPort() int32`

GetNatPort returns the NatPort field if non-nil, zero value otherwise.

### GetNatPortOk

`func (o *ControllerVm) GetNatPortOk() (*int32, bool)`

GetNatPortOk returns a tuple with the NatPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatPort

`func (o *ControllerVm) SetNatPort(v int32)`

SetNatPort sets NatPort field to given value.

### HasNatPort

`func (o *ControllerVm) HasNatPort() bool`

HasNatPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



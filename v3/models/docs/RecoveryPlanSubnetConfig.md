# RecoveryPlanSubnetConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**L2StretchConfig** | Pointer to [**[]RecoveryPlanL2StretchConfig**](RecoveryPlanL2StretchConfig.md) | Configurations for creating Layer2 stretches for the subnet.  | [optional] 
**ExternalConnectivityState** | Pointer to **string** | External connectivity state of the subnet. This is applicable only for the subnet to be created in public cloud Availability Zone.  | [optional] 
**GatewayIp** | **string** | Gateway IP address for the subnet.  | 
**PrefixLength** | **int32** | Prefix length for the subnet.  | 

## Methods

### NewRecoveryPlanSubnetConfig

`func NewRecoveryPlanSubnetConfig(gatewayIp string, prefixLength int32, ) *RecoveryPlanSubnetConfig`

NewRecoveryPlanSubnetConfig instantiates a new RecoveryPlanSubnetConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoveryPlanSubnetConfigWithDefaults

`func NewRecoveryPlanSubnetConfigWithDefaults() *RecoveryPlanSubnetConfig`

NewRecoveryPlanSubnetConfigWithDefaults instantiates a new RecoveryPlanSubnetConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetL2StretchConfig

`func (o *RecoveryPlanSubnetConfig) GetL2StretchConfig() []RecoveryPlanL2StretchConfig`

GetL2StretchConfig returns the L2StretchConfig field if non-nil, zero value otherwise.

### GetL2StretchConfigOk

`func (o *RecoveryPlanSubnetConfig) GetL2StretchConfigOk() (*[]RecoveryPlanL2StretchConfig, bool)`

GetL2StretchConfigOk returns a tuple with the L2StretchConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetL2StretchConfig

`func (o *RecoveryPlanSubnetConfig) SetL2StretchConfig(v []RecoveryPlanL2StretchConfig)`

SetL2StretchConfig sets L2StretchConfig field to given value.

### HasL2StretchConfig

`func (o *RecoveryPlanSubnetConfig) HasL2StretchConfig() bool`

HasL2StretchConfig returns a boolean if a field has been set.

### GetExternalConnectivityState

`func (o *RecoveryPlanSubnetConfig) GetExternalConnectivityState() string`

GetExternalConnectivityState returns the ExternalConnectivityState field if non-nil, zero value otherwise.

### GetExternalConnectivityStateOk

`func (o *RecoveryPlanSubnetConfig) GetExternalConnectivityStateOk() (*string, bool)`

GetExternalConnectivityStateOk returns a tuple with the ExternalConnectivityState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalConnectivityState

`func (o *RecoveryPlanSubnetConfig) SetExternalConnectivityState(v string)`

SetExternalConnectivityState sets ExternalConnectivityState field to given value.

### HasExternalConnectivityState

`func (o *RecoveryPlanSubnetConfig) HasExternalConnectivityState() bool`

HasExternalConnectivityState returns a boolean if a field has been set.

### GetGatewayIp

`func (o *RecoveryPlanSubnetConfig) GetGatewayIp() string`

GetGatewayIp returns the GatewayIp field if non-nil, zero value otherwise.

### GetGatewayIpOk

`func (o *RecoveryPlanSubnetConfig) GetGatewayIpOk() (*string, bool)`

GetGatewayIpOk returns a tuple with the GatewayIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayIp

`func (o *RecoveryPlanSubnetConfig) SetGatewayIp(v string)`

SetGatewayIp sets GatewayIp field to given value.


### GetPrefixLength

`func (o *RecoveryPlanSubnetConfig) GetPrefixLength() int32`

GetPrefixLength returns the PrefixLength field if non-nil, zero value otherwise.

### GetPrefixLengthOk

`func (o *RecoveryPlanSubnetConfig) GetPrefixLengthOk() (*int32, bool)`

GetPrefixLengthOk returns a tuple with the PrefixLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixLength

`func (o *RecoveryPlanSubnetConfig) SetPrefixLength(v int32)`

SetPrefixLength sets PrefixLength field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



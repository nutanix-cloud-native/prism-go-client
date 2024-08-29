# RecoveryPlanL2StretchConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocalVtepInterfaceIpAddress** | **string** | VTEP Interface IP address on site for which subnet configuration is being specified.  | 
**RemoteVtepGatewayReference** | [**VtepGatewayReference**](VtepGatewayReference.md) |  | 
**LocalVtepGatewayConfig** | [**VTEPGatewayConfiguration**](VTEPGatewayConfiguration.md) |  | 
**RemoteVtepInterfaceIpAddress** | **string** | VTEP Interface IP address on site to which subnet is extended.  | 

## Methods

### NewRecoveryPlanL2StretchConfig

`func NewRecoveryPlanL2StretchConfig(localVtepInterfaceIpAddress string, remoteVtepGatewayReference VtepGatewayReference, localVtepGatewayConfig VTEPGatewayConfiguration, remoteVtepInterfaceIpAddress string, ) *RecoveryPlanL2StretchConfig`

NewRecoveryPlanL2StretchConfig instantiates a new RecoveryPlanL2StretchConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoveryPlanL2StretchConfigWithDefaults

`func NewRecoveryPlanL2StretchConfigWithDefaults() *RecoveryPlanL2StretchConfig`

NewRecoveryPlanL2StretchConfigWithDefaults instantiates a new RecoveryPlanL2StretchConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocalVtepInterfaceIpAddress

`func (o *RecoveryPlanL2StretchConfig) GetLocalVtepInterfaceIpAddress() string`

GetLocalVtepInterfaceIpAddress returns the LocalVtepInterfaceIpAddress field if non-nil, zero value otherwise.

### GetLocalVtepInterfaceIpAddressOk

`func (o *RecoveryPlanL2StretchConfig) GetLocalVtepInterfaceIpAddressOk() (*string, bool)`

GetLocalVtepInterfaceIpAddressOk returns a tuple with the LocalVtepInterfaceIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalVtepInterfaceIpAddress

`func (o *RecoveryPlanL2StretchConfig) SetLocalVtepInterfaceIpAddress(v string)`

SetLocalVtepInterfaceIpAddress sets LocalVtepInterfaceIpAddress field to given value.


### GetRemoteVtepGatewayReference

`func (o *RecoveryPlanL2StretchConfig) GetRemoteVtepGatewayReference() VtepGatewayReference`

GetRemoteVtepGatewayReference returns the RemoteVtepGatewayReference field if non-nil, zero value otherwise.

### GetRemoteVtepGatewayReferenceOk

`func (o *RecoveryPlanL2StretchConfig) GetRemoteVtepGatewayReferenceOk() (*VtepGatewayReference, bool)`

GetRemoteVtepGatewayReferenceOk returns a tuple with the RemoteVtepGatewayReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteVtepGatewayReference

`func (o *RecoveryPlanL2StretchConfig) SetRemoteVtepGatewayReference(v VtepGatewayReference)`

SetRemoteVtepGatewayReference sets RemoteVtepGatewayReference field to given value.


### GetLocalVtepGatewayConfig

`func (o *RecoveryPlanL2StretchConfig) GetLocalVtepGatewayConfig() VTEPGatewayConfiguration`

GetLocalVtepGatewayConfig returns the LocalVtepGatewayConfig field if non-nil, zero value otherwise.

### GetLocalVtepGatewayConfigOk

`func (o *RecoveryPlanL2StretchConfig) GetLocalVtepGatewayConfigOk() (*VTEPGatewayConfiguration, bool)`

GetLocalVtepGatewayConfigOk returns a tuple with the LocalVtepGatewayConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalVtepGatewayConfig

`func (o *RecoveryPlanL2StretchConfig) SetLocalVtepGatewayConfig(v VTEPGatewayConfiguration)`

SetLocalVtepGatewayConfig sets LocalVtepGatewayConfig field to given value.


### GetRemoteVtepInterfaceIpAddress

`func (o *RecoveryPlanL2StretchConfig) GetRemoteVtepInterfaceIpAddress() string`

GetRemoteVtepInterfaceIpAddress returns the RemoteVtepInterfaceIpAddress field if non-nil, zero value otherwise.

### GetRemoteVtepInterfaceIpAddressOk

`func (o *RecoveryPlanL2StretchConfig) GetRemoteVtepInterfaceIpAddressOk() (*string, bool)`

GetRemoteVtepInterfaceIpAddressOk returns a tuple with the RemoteVtepInterfaceIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteVtepInterfaceIpAddress

`func (o *RecoveryPlanL2StretchConfig) SetRemoteVtepInterfaceIpAddress(v string)`

SetRemoteVtepInterfaceIpAddress sets RemoteVtepInterfaceIpAddress field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# RecoveryPlanNetwork

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the network.  | [optional] 
**VpcReference** | Pointer to [**VpcReference**](VpcReference.md) |  | [optional] 
**UseVpcReference** | Pointer to **bool** | Client need to specify this field as true while using vpc_reference for specifying the VPC for the network. Without this values in vpc_reference will be ignored.  | [optional] 
**VirtualNetworkReference** | Pointer to [**VirtualNetworkReference**](VirtualNetworkReference.md) |  | [optional] 
**SubnetList** | Pointer to [**[]RecoveryPlanSubnetConfig**](RecoveryPlanSubnetConfig.md) | List of subnets for the network.  | [optional] 

## Methods

### NewRecoveryPlanNetwork

`func NewRecoveryPlanNetwork() *RecoveryPlanNetwork`

NewRecoveryPlanNetwork instantiates a new RecoveryPlanNetwork object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoveryPlanNetworkWithDefaults

`func NewRecoveryPlanNetworkWithDefaults() *RecoveryPlanNetwork`

NewRecoveryPlanNetworkWithDefaults instantiates a new RecoveryPlanNetwork object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RecoveryPlanNetwork) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RecoveryPlanNetwork) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RecoveryPlanNetwork) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RecoveryPlanNetwork) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVpcReference

`func (o *RecoveryPlanNetwork) GetVpcReference() VpcReference`

GetVpcReference returns the VpcReference field if non-nil, zero value otherwise.

### GetVpcReferenceOk

`func (o *RecoveryPlanNetwork) GetVpcReferenceOk() (*VpcReference, bool)`

GetVpcReferenceOk returns a tuple with the VpcReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcReference

`func (o *RecoveryPlanNetwork) SetVpcReference(v VpcReference)`

SetVpcReference sets VpcReference field to given value.

### HasVpcReference

`func (o *RecoveryPlanNetwork) HasVpcReference() bool`

HasVpcReference returns a boolean if a field has been set.

### GetUseVpcReference

`func (o *RecoveryPlanNetwork) GetUseVpcReference() bool`

GetUseVpcReference returns the UseVpcReference field if non-nil, zero value otherwise.

### GetUseVpcReferenceOk

`func (o *RecoveryPlanNetwork) GetUseVpcReferenceOk() (*bool, bool)`

GetUseVpcReferenceOk returns a tuple with the UseVpcReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseVpcReference

`func (o *RecoveryPlanNetwork) SetUseVpcReference(v bool)`

SetUseVpcReference sets UseVpcReference field to given value.

### HasUseVpcReference

`func (o *RecoveryPlanNetwork) HasUseVpcReference() bool`

HasUseVpcReference returns a boolean if a field has been set.

### GetVirtualNetworkReference

`func (o *RecoveryPlanNetwork) GetVirtualNetworkReference() VirtualNetworkReference`

GetVirtualNetworkReference returns the VirtualNetworkReference field if non-nil, zero value otherwise.

### GetVirtualNetworkReferenceOk

`func (o *RecoveryPlanNetwork) GetVirtualNetworkReferenceOk() (*VirtualNetworkReference, bool)`

GetVirtualNetworkReferenceOk returns a tuple with the VirtualNetworkReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualNetworkReference

`func (o *RecoveryPlanNetwork) SetVirtualNetworkReference(v VirtualNetworkReference)`

SetVirtualNetworkReference sets VirtualNetworkReference field to given value.

### HasVirtualNetworkReference

`func (o *RecoveryPlanNetwork) HasVirtualNetworkReference() bool`

HasVirtualNetworkReference returns a boolean if a field has been set.

### GetSubnetList

`func (o *RecoveryPlanNetwork) GetSubnetList() []RecoveryPlanSubnetConfig`

GetSubnetList returns the SubnetList field if non-nil, zero value otherwise.

### GetSubnetListOk

`func (o *RecoveryPlanNetwork) GetSubnetListOk() (*[]RecoveryPlanSubnetConfig, bool)`

GetSubnetListOk returns a tuple with the SubnetList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetList

`func (o *RecoveryPlanNetwork) SetSubnetList(v []RecoveryPlanSubnetConfig)`

SetSubnetList sets SubnetList field to given value.

### HasSubnetList

`func (o *RecoveryPlanNetwork) HasSubnetList() bool`

HasSubnetList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



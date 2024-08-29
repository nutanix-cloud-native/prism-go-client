# AwsVmResources

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceType** | Pointer to **string** | The type of instance e.g.&#39;t1.micro&#39;, &#39;m1.small&#39; | [optional] 
**InstanceProfileName** | Pointer to **string** | The name of the IAM Instance Profile (IIP) associated with the instance  | [optional] 
**AvailabilityZone** | Pointer to **string** | The zone on which the instance is to be created. | [optional] 
**SubnetId** | Pointer to **string** | The subnet within the VPC the instance belongs to. | [optional] 
**KeyName** | Pointer to **string** | The name of the key pair used to launch the instance | [optional] 
**Region** | Pointer to **string** | The region to which the instance belongs. | [optional] 
**UserData** | Pointer to **string** | User data passed to launch the instance | [optional] 
**InstanceInitiatedShutdownBehavior** | Pointer to **string** | Specifies whether the instance stops or terminates on instance-initiated shutdown.  | [optional] 
**ImageId** | Pointer to **string** | The AWS ID of the AMI on the instance. | [optional] 
**State** | Pointer to **string** | Instance&#39;s desired state. | [optional] 
**SecurityGroupList** | Pointer to [**[]AwsSecurityGroupListInner**](AwsSecurityGroupListInner.md) | List of AWS security group IDs. | [optional] 
**BlockDeviceMap** | Pointer to [**AwsBlockDeviceMap**](AwsBlockDeviceMap.md) |  | [optional] 
**PrivateIpAddress** | Pointer to **string** | The specific available IP from the subnet assigned to the instance.  | [optional] 
**VpcId** | Pointer to **string** | The VPC AWS ID, if running in VPC. | [optional] 
**TagList** | Pointer to [**[]AwsTagListInner**](AwsTagListInner.md) | The AWS Tags associated with any AWS resource | [optional] 
**AccountUuid** | Pointer to **string** | The AWS account to which the instance belongs. | [optional] 
**AssociatePublicIpAddress** | Pointer to **bool** | Indicates whether the network interface receives a public IP address.Can associate a public IP address with a network interface only if it has a device index of eth0 and if it is a new network interface (not an existing one).  | [optional] 

## Methods

### NewAwsVmResources

`func NewAwsVmResources() *AwsVmResources`

NewAwsVmResources instantiates a new AwsVmResources object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsVmResourcesWithDefaults

`func NewAwsVmResourcesWithDefaults() *AwsVmResources`

NewAwsVmResourcesWithDefaults instantiates a new AwsVmResources object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceType

`func (o *AwsVmResources) GetInstanceType() string`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *AwsVmResources) GetInstanceTypeOk() (*string, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *AwsVmResources) SetInstanceType(v string)`

SetInstanceType sets InstanceType field to given value.

### HasInstanceType

`func (o *AwsVmResources) HasInstanceType() bool`

HasInstanceType returns a boolean if a field has been set.

### GetInstanceProfileName

`func (o *AwsVmResources) GetInstanceProfileName() string`

GetInstanceProfileName returns the InstanceProfileName field if non-nil, zero value otherwise.

### GetInstanceProfileNameOk

`func (o *AwsVmResources) GetInstanceProfileNameOk() (*string, bool)`

GetInstanceProfileNameOk returns a tuple with the InstanceProfileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceProfileName

`func (o *AwsVmResources) SetInstanceProfileName(v string)`

SetInstanceProfileName sets InstanceProfileName field to given value.

### HasInstanceProfileName

`func (o *AwsVmResources) HasInstanceProfileName() bool`

HasInstanceProfileName returns a boolean if a field has been set.

### GetAvailabilityZone

`func (o *AwsVmResources) GetAvailabilityZone() string`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *AwsVmResources) GetAvailabilityZoneOk() (*string, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *AwsVmResources) SetAvailabilityZone(v string)`

SetAvailabilityZone sets AvailabilityZone field to given value.

### HasAvailabilityZone

`func (o *AwsVmResources) HasAvailabilityZone() bool`

HasAvailabilityZone returns a boolean if a field has been set.

### GetSubnetId

`func (o *AwsVmResources) GetSubnetId() string`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *AwsVmResources) GetSubnetIdOk() (*string, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *AwsVmResources) SetSubnetId(v string)`

SetSubnetId sets SubnetId field to given value.

### HasSubnetId

`func (o *AwsVmResources) HasSubnetId() bool`

HasSubnetId returns a boolean if a field has been set.

### GetKeyName

`func (o *AwsVmResources) GetKeyName() string`

GetKeyName returns the KeyName field if non-nil, zero value otherwise.

### GetKeyNameOk

`func (o *AwsVmResources) GetKeyNameOk() (*string, bool)`

GetKeyNameOk returns a tuple with the KeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyName

`func (o *AwsVmResources) SetKeyName(v string)`

SetKeyName sets KeyName field to given value.

### HasKeyName

`func (o *AwsVmResources) HasKeyName() bool`

HasKeyName returns a boolean if a field has been set.

### GetRegion

`func (o *AwsVmResources) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *AwsVmResources) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *AwsVmResources) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *AwsVmResources) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetUserData

`func (o *AwsVmResources) GetUserData() string`

GetUserData returns the UserData field if non-nil, zero value otherwise.

### GetUserDataOk

`func (o *AwsVmResources) GetUserDataOk() (*string, bool)`

GetUserDataOk returns a tuple with the UserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserData

`func (o *AwsVmResources) SetUserData(v string)`

SetUserData sets UserData field to given value.

### HasUserData

`func (o *AwsVmResources) HasUserData() bool`

HasUserData returns a boolean if a field has been set.

### GetInstanceInitiatedShutdownBehavior

`func (o *AwsVmResources) GetInstanceInitiatedShutdownBehavior() string`

GetInstanceInitiatedShutdownBehavior returns the InstanceInitiatedShutdownBehavior field if non-nil, zero value otherwise.

### GetInstanceInitiatedShutdownBehaviorOk

`func (o *AwsVmResources) GetInstanceInitiatedShutdownBehaviorOk() (*string, bool)`

GetInstanceInitiatedShutdownBehaviorOk returns a tuple with the InstanceInitiatedShutdownBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceInitiatedShutdownBehavior

`func (o *AwsVmResources) SetInstanceInitiatedShutdownBehavior(v string)`

SetInstanceInitiatedShutdownBehavior sets InstanceInitiatedShutdownBehavior field to given value.

### HasInstanceInitiatedShutdownBehavior

`func (o *AwsVmResources) HasInstanceInitiatedShutdownBehavior() bool`

HasInstanceInitiatedShutdownBehavior returns a boolean if a field has been set.

### GetImageId

`func (o *AwsVmResources) GetImageId() string`

GetImageId returns the ImageId field if non-nil, zero value otherwise.

### GetImageIdOk

`func (o *AwsVmResources) GetImageIdOk() (*string, bool)`

GetImageIdOk returns a tuple with the ImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageId

`func (o *AwsVmResources) SetImageId(v string)`

SetImageId sets ImageId field to given value.

### HasImageId

`func (o *AwsVmResources) HasImageId() bool`

HasImageId returns a boolean if a field has been set.

### GetState

`func (o *AwsVmResources) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AwsVmResources) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AwsVmResources) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *AwsVmResources) HasState() bool`

HasState returns a boolean if a field has been set.

### GetSecurityGroupList

`func (o *AwsVmResources) GetSecurityGroupList() []AwsSecurityGroupListInner`

GetSecurityGroupList returns the SecurityGroupList field if non-nil, zero value otherwise.

### GetSecurityGroupListOk

`func (o *AwsVmResources) GetSecurityGroupListOk() (*[]AwsSecurityGroupListInner, bool)`

GetSecurityGroupListOk returns a tuple with the SecurityGroupList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroupList

`func (o *AwsVmResources) SetSecurityGroupList(v []AwsSecurityGroupListInner)`

SetSecurityGroupList sets SecurityGroupList field to given value.

### HasSecurityGroupList

`func (o *AwsVmResources) HasSecurityGroupList() bool`

HasSecurityGroupList returns a boolean if a field has been set.

### GetBlockDeviceMap

`func (o *AwsVmResources) GetBlockDeviceMap() AwsBlockDeviceMap`

GetBlockDeviceMap returns the BlockDeviceMap field if non-nil, zero value otherwise.

### GetBlockDeviceMapOk

`func (o *AwsVmResources) GetBlockDeviceMapOk() (*AwsBlockDeviceMap, bool)`

GetBlockDeviceMapOk returns a tuple with the BlockDeviceMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockDeviceMap

`func (o *AwsVmResources) SetBlockDeviceMap(v AwsBlockDeviceMap)`

SetBlockDeviceMap sets BlockDeviceMap field to given value.

### HasBlockDeviceMap

`func (o *AwsVmResources) HasBlockDeviceMap() bool`

HasBlockDeviceMap returns a boolean if a field has been set.

### GetPrivateIpAddress

`func (o *AwsVmResources) GetPrivateIpAddress() string`

GetPrivateIpAddress returns the PrivateIpAddress field if non-nil, zero value otherwise.

### GetPrivateIpAddressOk

`func (o *AwsVmResources) GetPrivateIpAddressOk() (*string, bool)`

GetPrivateIpAddressOk returns a tuple with the PrivateIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateIpAddress

`func (o *AwsVmResources) SetPrivateIpAddress(v string)`

SetPrivateIpAddress sets PrivateIpAddress field to given value.

### HasPrivateIpAddress

`func (o *AwsVmResources) HasPrivateIpAddress() bool`

HasPrivateIpAddress returns a boolean if a field has been set.

### GetVpcId

`func (o *AwsVmResources) GetVpcId() string`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *AwsVmResources) GetVpcIdOk() (*string, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *AwsVmResources) SetVpcId(v string)`

SetVpcId sets VpcId field to given value.

### HasVpcId

`func (o *AwsVmResources) HasVpcId() bool`

HasVpcId returns a boolean if a field has been set.

### GetTagList

`func (o *AwsVmResources) GetTagList() []AwsTagListInner`

GetTagList returns the TagList field if non-nil, zero value otherwise.

### GetTagListOk

`func (o *AwsVmResources) GetTagListOk() (*[]AwsTagListInner, bool)`

GetTagListOk returns a tuple with the TagList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagList

`func (o *AwsVmResources) SetTagList(v []AwsTagListInner)`

SetTagList sets TagList field to given value.

### HasTagList

`func (o *AwsVmResources) HasTagList() bool`

HasTagList returns a boolean if a field has been set.

### GetAccountUuid

`func (o *AwsVmResources) GetAccountUuid() string`

GetAccountUuid returns the AccountUuid field if non-nil, zero value otherwise.

### GetAccountUuidOk

`func (o *AwsVmResources) GetAccountUuidOk() (*string, bool)`

GetAccountUuidOk returns a tuple with the AccountUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountUuid

`func (o *AwsVmResources) SetAccountUuid(v string)`

SetAccountUuid sets AccountUuid field to given value.

### HasAccountUuid

`func (o *AwsVmResources) HasAccountUuid() bool`

HasAccountUuid returns a boolean if a field has been set.

### GetAssociatePublicIpAddress

`func (o *AwsVmResources) GetAssociatePublicIpAddress() bool`

GetAssociatePublicIpAddress returns the AssociatePublicIpAddress field if non-nil, zero value otherwise.

### GetAssociatePublicIpAddressOk

`func (o *AwsVmResources) GetAssociatePublicIpAddressOk() (*bool, bool)`

GetAssociatePublicIpAddressOk returns a tuple with the AssociatePublicIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatePublicIpAddress

`func (o *AwsVmResources) SetAssociatePublicIpAddress(v bool)`

SetAssociatePublicIpAddress sets AssociatePublicIpAddress field to given value.

### HasAssociatePublicIpAddress

`func (o *AwsVmResources) HasAssociatePublicIpAddress() bool`

HasAssociatePublicIpAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



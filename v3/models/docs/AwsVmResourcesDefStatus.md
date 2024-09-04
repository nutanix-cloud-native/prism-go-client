# AwsVmResourcesDefStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RootDeviceType** | Pointer to **string** | The root device type | [optional] 
**PrivateDnsName** | Pointer to **string** | The private DNS name of the instance. | [optional] 
**AvailabilityZone** | Pointer to **string** | The zone on which the instance is created | [optional] 
**BlockDeviceMap** | Pointer to [**AwsBlockDeviceMapOutputStatus**](AwsBlockDeviceMapOutputStatus.md) |  | [optional] 
**Id** | Pointer to **string** | The AWS ID of the instance. | [optional] 
**SubnetId** | Pointer to **string** | The subnet within the VPC the instance belongs to. | [optional] 
**Platform** | Pointer to **string** | Platform of the instance e.g. Windows | [optional] 
**State** | Pointer to **string** | Instance&#39;s current state. | [optional] 
**ConsoleOutput** | Pointer to **string** | Console output of the instance | [optional] 
**PublicIpAddress** | Pointer to **string** | The public IP of the instance | [optional] 
**KeyName** | Pointer to **string** | The name of the key pair used to launch the instance | [optional] 
**ImageId** | Pointer to **string** | The AWS ID of the AMI on the instance. | [optional] 
**PublicDnsName** | Pointer to **string** | The public DNS name of the instance. | [optional] 
**AccountUuid** | Pointer to **string** | The AWS account to which the instance belongs. | [optional] 
**InstanceProfileName** | Pointer to **string** | The name of the IAM Instance Profile (IIP) associated with the instance.  | [optional] 
**Region** | Pointer to **string** | The region to which the instance belongs. | [optional] 
**LaunchTime** | Pointer to **string** | The time the instance was launched | [optional] 
**InstanceInitiatedShutdownBehavior** | Pointer to **string** | Specifies whether the instance stops or terminates on instance-initiated shutdown.  | [optional] 
**InstanceType** | Pointer to **string** | The type of instance e.g.&#39;t1.micro&#39;, &#39;m1.small&#39; | [optional] 
**SecurityGroupList** | Pointer to [**[]AwsSecurityGroupListInner**](AwsSecurityGroupListInner.md) | List of AWS security group IDs. | [optional] 
**PrivateIpAddress** | Pointer to **string** | The specific available IP from the subnet assigned to the instance.  | [optional] 
**VpcId** | Pointer to **string** | The VPC AWS ID, if running in VPC. | [optional] 
**TagList** | Pointer to [**[]AwsTagListInner**](AwsTagListInner.md) | The AWS Tags associated with any AWS resource | [optional] 

## Methods

### NewAwsVmResourcesDefStatus

`func NewAwsVmResourcesDefStatus() *AwsVmResourcesDefStatus`

NewAwsVmResourcesDefStatus instantiates a new AwsVmResourcesDefStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsVmResourcesDefStatusWithDefaults

`func NewAwsVmResourcesDefStatusWithDefaults() *AwsVmResourcesDefStatus`

NewAwsVmResourcesDefStatusWithDefaults instantiates a new AwsVmResourcesDefStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRootDeviceType

`func (o *AwsVmResourcesDefStatus) GetRootDeviceType() string`

GetRootDeviceType returns the RootDeviceType field if non-nil, zero value otherwise.

### GetRootDeviceTypeOk

`func (o *AwsVmResourcesDefStatus) GetRootDeviceTypeOk() (*string, bool)`

GetRootDeviceTypeOk returns a tuple with the RootDeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootDeviceType

`func (o *AwsVmResourcesDefStatus) SetRootDeviceType(v string)`

SetRootDeviceType sets RootDeviceType field to given value.

### HasRootDeviceType

`func (o *AwsVmResourcesDefStatus) HasRootDeviceType() bool`

HasRootDeviceType returns a boolean if a field has been set.

### GetPrivateDnsName

`func (o *AwsVmResourcesDefStatus) GetPrivateDnsName() string`

GetPrivateDnsName returns the PrivateDnsName field if non-nil, zero value otherwise.

### GetPrivateDnsNameOk

`func (o *AwsVmResourcesDefStatus) GetPrivateDnsNameOk() (*string, bool)`

GetPrivateDnsNameOk returns a tuple with the PrivateDnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateDnsName

`func (o *AwsVmResourcesDefStatus) SetPrivateDnsName(v string)`

SetPrivateDnsName sets PrivateDnsName field to given value.

### HasPrivateDnsName

`func (o *AwsVmResourcesDefStatus) HasPrivateDnsName() bool`

HasPrivateDnsName returns a boolean if a field has been set.

### GetAvailabilityZone

`func (o *AwsVmResourcesDefStatus) GetAvailabilityZone() string`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *AwsVmResourcesDefStatus) GetAvailabilityZoneOk() (*string, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *AwsVmResourcesDefStatus) SetAvailabilityZone(v string)`

SetAvailabilityZone sets AvailabilityZone field to given value.

### HasAvailabilityZone

`func (o *AwsVmResourcesDefStatus) HasAvailabilityZone() bool`

HasAvailabilityZone returns a boolean if a field has been set.

### GetBlockDeviceMap

`func (o *AwsVmResourcesDefStatus) GetBlockDeviceMap() AwsBlockDeviceMapOutputStatus`

GetBlockDeviceMap returns the BlockDeviceMap field if non-nil, zero value otherwise.

### GetBlockDeviceMapOk

`func (o *AwsVmResourcesDefStatus) GetBlockDeviceMapOk() (*AwsBlockDeviceMapOutputStatus, bool)`

GetBlockDeviceMapOk returns a tuple with the BlockDeviceMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockDeviceMap

`func (o *AwsVmResourcesDefStatus) SetBlockDeviceMap(v AwsBlockDeviceMapOutputStatus)`

SetBlockDeviceMap sets BlockDeviceMap field to given value.

### HasBlockDeviceMap

`func (o *AwsVmResourcesDefStatus) HasBlockDeviceMap() bool`

HasBlockDeviceMap returns a boolean if a field has been set.

### GetId

`func (o *AwsVmResourcesDefStatus) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AwsVmResourcesDefStatus) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AwsVmResourcesDefStatus) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AwsVmResourcesDefStatus) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSubnetId

`func (o *AwsVmResourcesDefStatus) GetSubnetId() string`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *AwsVmResourcesDefStatus) GetSubnetIdOk() (*string, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *AwsVmResourcesDefStatus) SetSubnetId(v string)`

SetSubnetId sets SubnetId field to given value.

### HasSubnetId

`func (o *AwsVmResourcesDefStatus) HasSubnetId() bool`

HasSubnetId returns a boolean if a field has been set.

### GetPlatform

`func (o *AwsVmResourcesDefStatus) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *AwsVmResourcesDefStatus) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *AwsVmResourcesDefStatus) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *AwsVmResourcesDefStatus) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetState

`func (o *AwsVmResourcesDefStatus) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AwsVmResourcesDefStatus) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AwsVmResourcesDefStatus) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *AwsVmResourcesDefStatus) HasState() bool`

HasState returns a boolean if a field has been set.

### GetConsoleOutput

`func (o *AwsVmResourcesDefStatus) GetConsoleOutput() string`

GetConsoleOutput returns the ConsoleOutput field if non-nil, zero value otherwise.

### GetConsoleOutputOk

`func (o *AwsVmResourcesDefStatus) GetConsoleOutputOk() (*string, bool)`

GetConsoleOutputOk returns a tuple with the ConsoleOutput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsoleOutput

`func (o *AwsVmResourcesDefStatus) SetConsoleOutput(v string)`

SetConsoleOutput sets ConsoleOutput field to given value.

### HasConsoleOutput

`func (o *AwsVmResourcesDefStatus) HasConsoleOutput() bool`

HasConsoleOutput returns a boolean if a field has been set.

### GetPublicIpAddress

`func (o *AwsVmResourcesDefStatus) GetPublicIpAddress() string`

GetPublicIpAddress returns the PublicIpAddress field if non-nil, zero value otherwise.

### GetPublicIpAddressOk

`func (o *AwsVmResourcesDefStatus) GetPublicIpAddressOk() (*string, bool)`

GetPublicIpAddressOk returns a tuple with the PublicIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIpAddress

`func (o *AwsVmResourcesDefStatus) SetPublicIpAddress(v string)`

SetPublicIpAddress sets PublicIpAddress field to given value.

### HasPublicIpAddress

`func (o *AwsVmResourcesDefStatus) HasPublicIpAddress() bool`

HasPublicIpAddress returns a boolean if a field has been set.

### GetKeyName

`func (o *AwsVmResourcesDefStatus) GetKeyName() string`

GetKeyName returns the KeyName field if non-nil, zero value otherwise.

### GetKeyNameOk

`func (o *AwsVmResourcesDefStatus) GetKeyNameOk() (*string, bool)`

GetKeyNameOk returns a tuple with the KeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyName

`func (o *AwsVmResourcesDefStatus) SetKeyName(v string)`

SetKeyName sets KeyName field to given value.

### HasKeyName

`func (o *AwsVmResourcesDefStatus) HasKeyName() bool`

HasKeyName returns a boolean if a field has been set.

### GetImageId

`func (o *AwsVmResourcesDefStatus) GetImageId() string`

GetImageId returns the ImageId field if non-nil, zero value otherwise.

### GetImageIdOk

`func (o *AwsVmResourcesDefStatus) GetImageIdOk() (*string, bool)`

GetImageIdOk returns a tuple with the ImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageId

`func (o *AwsVmResourcesDefStatus) SetImageId(v string)`

SetImageId sets ImageId field to given value.

### HasImageId

`func (o *AwsVmResourcesDefStatus) HasImageId() bool`

HasImageId returns a boolean if a field has been set.

### GetPublicDnsName

`func (o *AwsVmResourcesDefStatus) GetPublicDnsName() string`

GetPublicDnsName returns the PublicDnsName field if non-nil, zero value otherwise.

### GetPublicDnsNameOk

`func (o *AwsVmResourcesDefStatus) GetPublicDnsNameOk() (*string, bool)`

GetPublicDnsNameOk returns a tuple with the PublicDnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicDnsName

`func (o *AwsVmResourcesDefStatus) SetPublicDnsName(v string)`

SetPublicDnsName sets PublicDnsName field to given value.

### HasPublicDnsName

`func (o *AwsVmResourcesDefStatus) HasPublicDnsName() bool`

HasPublicDnsName returns a boolean if a field has been set.

### GetAccountUuid

`func (o *AwsVmResourcesDefStatus) GetAccountUuid() string`

GetAccountUuid returns the AccountUuid field if non-nil, zero value otherwise.

### GetAccountUuidOk

`func (o *AwsVmResourcesDefStatus) GetAccountUuidOk() (*string, bool)`

GetAccountUuidOk returns a tuple with the AccountUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountUuid

`func (o *AwsVmResourcesDefStatus) SetAccountUuid(v string)`

SetAccountUuid sets AccountUuid field to given value.

### HasAccountUuid

`func (o *AwsVmResourcesDefStatus) HasAccountUuid() bool`

HasAccountUuid returns a boolean if a field has been set.

### GetInstanceProfileName

`func (o *AwsVmResourcesDefStatus) GetInstanceProfileName() string`

GetInstanceProfileName returns the InstanceProfileName field if non-nil, zero value otherwise.

### GetInstanceProfileNameOk

`func (o *AwsVmResourcesDefStatus) GetInstanceProfileNameOk() (*string, bool)`

GetInstanceProfileNameOk returns a tuple with the InstanceProfileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceProfileName

`func (o *AwsVmResourcesDefStatus) SetInstanceProfileName(v string)`

SetInstanceProfileName sets InstanceProfileName field to given value.

### HasInstanceProfileName

`func (o *AwsVmResourcesDefStatus) HasInstanceProfileName() bool`

HasInstanceProfileName returns a boolean if a field has been set.

### GetRegion

`func (o *AwsVmResourcesDefStatus) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *AwsVmResourcesDefStatus) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *AwsVmResourcesDefStatus) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *AwsVmResourcesDefStatus) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetLaunchTime

`func (o *AwsVmResourcesDefStatus) GetLaunchTime() string`

GetLaunchTime returns the LaunchTime field if non-nil, zero value otherwise.

### GetLaunchTimeOk

`func (o *AwsVmResourcesDefStatus) GetLaunchTimeOk() (*string, bool)`

GetLaunchTimeOk returns a tuple with the LaunchTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLaunchTime

`func (o *AwsVmResourcesDefStatus) SetLaunchTime(v string)`

SetLaunchTime sets LaunchTime field to given value.

### HasLaunchTime

`func (o *AwsVmResourcesDefStatus) HasLaunchTime() bool`

HasLaunchTime returns a boolean if a field has been set.

### GetInstanceInitiatedShutdownBehavior

`func (o *AwsVmResourcesDefStatus) GetInstanceInitiatedShutdownBehavior() string`

GetInstanceInitiatedShutdownBehavior returns the InstanceInitiatedShutdownBehavior field if non-nil, zero value otherwise.

### GetInstanceInitiatedShutdownBehaviorOk

`func (o *AwsVmResourcesDefStatus) GetInstanceInitiatedShutdownBehaviorOk() (*string, bool)`

GetInstanceInitiatedShutdownBehaviorOk returns a tuple with the InstanceInitiatedShutdownBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceInitiatedShutdownBehavior

`func (o *AwsVmResourcesDefStatus) SetInstanceInitiatedShutdownBehavior(v string)`

SetInstanceInitiatedShutdownBehavior sets InstanceInitiatedShutdownBehavior field to given value.

### HasInstanceInitiatedShutdownBehavior

`func (o *AwsVmResourcesDefStatus) HasInstanceInitiatedShutdownBehavior() bool`

HasInstanceInitiatedShutdownBehavior returns a boolean if a field has been set.

### GetInstanceType

`func (o *AwsVmResourcesDefStatus) GetInstanceType() string`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *AwsVmResourcesDefStatus) GetInstanceTypeOk() (*string, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *AwsVmResourcesDefStatus) SetInstanceType(v string)`

SetInstanceType sets InstanceType field to given value.

### HasInstanceType

`func (o *AwsVmResourcesDefStatus) HasInstanceType() bool`

HasInstanceType returns a boolean if a field has been set.

### GetSecurityGroupList

`func (o *AwsVmResourcesDefStatus) GetSecurityGroupList() []AwsSecurityGroupListInner`

GetSecurityGroupList returns the SecurityGroupList field if non-nil, zero value otherwise.

### GetSecurityGroupListOk

`func (o *AwsVmResourcesDefStatus) GetSecurityGroupListOk() (*[]AwsSecurityGroupListInner, bool)`

GetSecurityGroupListOk returns a tuple with the SecurityGroupList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroupList

`func (o *AwsVmResourcesDefStatus) SetSecurityGroupList(v []AwsSecurityGroupListInner)`

SetSecurityGroupList sets SecurityGroupList field to given value.

### HasSecurityGroupList

`func (o *AwsVmResourcesDefStatus) HasSecurityGroupList() bool`

HasSecurityGroupList returns a boolean if a field has been set.

### GetPrivateIpAddress

`func (o *AwsVmResourcesDefStatus) GetPrivateIpAddress() string`

GetPrivateIpAddress returns the PrivateIpAddress field if non-nil, zero value otherwise.

### GetPrivateIpAddressOk

`func (o *AwsVmResourcesDefStatus) GetPrivateIpAddressOk() (*string, bool)`

GetPrivateIpAddressOk returns a tuple with the PrivateIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateIpAddress

`func (o *AwsVmResourcesDefStatus) SetPrivateIpAddress(v string)`

SetPrivateIpAddress sets PrivateIpAddress field to given value.

### HasPrivateIpAddress

`func (o *AwsVmResourcesDefStatus) HasPrivateIpAddress() bool`

HasPrivateIpAddress returns a boolean if a field has been set.

### GetVpcId

`func (o *AwsVmResourcesDefStatus) GetVpcId() string`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *AwsVmResourcesDefStatus) GetVpcIdOk() (*string, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *AwsVmResourcesDefStatus) SetVpcId(v string)`

SetVpcId sets VpcId field to given value.

### HasVpcId

`func (o *AwsVmResourcesDefStatus) HasVpcId() bool`

HasVpcId returns a boolean if a field has been set.

### GetTagList

`func (o *AwsVmResourcesDefStatus) GetTagList() []AwsTagListInner`

GetTagList returns the TagList field if non-nil, zero value otherwise.

### GetTagListOk

`func (o *AwsVmResourcesDefStatus) GetTagListOk() (*[]AwsTagListInner, bool)`

GetTagListOk returns a tuple with the TagList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagList

`func (o *AwsVmResourcesDefStatus) SetTagList(v []AwsTagListInner)`

SetTagList sets TagList field to given value.

### HasTagList

`func (o *AwsVmResourcesDefStatus) HasTagList() bool`

HasTagList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



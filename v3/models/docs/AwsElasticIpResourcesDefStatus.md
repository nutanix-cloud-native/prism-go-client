# AwsElasticIpResourcesDefStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | Pointer to **string** | Indicates whether the address is an EC2 or VPC address. | [optional] 
**PublicIpAddress** | Pointer to **string** | The Elastic IP address | [optional] 
**AllocationId** | Pointer to **string** | The allocation AWS ID for the address (only for VPC) | [optional] 
**InstanceId** | Pointer to **string** | The AWS ID of the instance the address is associated with.  | [optional] 
**AssociationId** | Pointer to **string** | The association AWS ID for the address (only for VPC) | [optional] 
**NetworkInterfaceId** | Pointer to **string** | The network interface (if any) that the address is associated with (only for VPC).  | [optional] 
**PrivateIpAddress** | Pointer to **string** | The private IP address associated with the Elastic IP (only for VPC)  | [optional] 
**NetworkInterfaceOwnerId** | Pointer to **string** | The owner AWS ID (only for VPC). | [optional] 

## Methods

### NewAwsElasticIpResourcesDefStatus

`func NewAwsElasticIpResourcesDefStatus() *AwsElasticIpResourcesDefStatus`

NewAwsElasticIpResourcesDefStatus instantiates a new AwsElasticIpResourcesDefStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsElasticIpResourcesDefStatusWithDefaults

`func NewAwsElasticIpResourcesDefStatusWithDefaults() *AwsElasticIpResourcesDefStatus`

NewAwsElasticIpResourcesDefStatusWithDefaults instantiates a new AwsElasticIpResourcesDefStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *AwsElasticIpResourcesDefStatus) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *AwsElasticIpResourcesDefStatus) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *AwsElasticIpResourcesDefStatus) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *AwsElasticIpResourcesDefStatus) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetPublicIpAddress

`func (o *AwsElasticIpResourcesDefStatus) GetPublicIpAddress() string`

GetPublicIpAddress returns the PublicIpAddress field if non-nil, zero value otherwise.

### GetPublicIpAddressOk

`func (o *AwsElasticIpResourcesDefStatus) GetPublicIpAddressOk() (*string, bool)`

GetPublicIpAddressOk returns a tuple with the PublicIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIpAddress

`func (o *AwsElasticIpResourcesDefStatus) SetPublicIpAddress(v string)`

SetPublicIpAddress sets PublicIpAddress field to given value.

### HasPublicIpAddress

`func (o *AwsElasticIpResourcesDefStatus) HasPublicIpAddress() bool`

HasPublicIpAddress returns a boolean if a field has been set.

### GetAllocationId

`func (o *AwsElasticIpResourcesDefStatus) GetAllocationId() string`

GetAllocationId returns the AllocationId field if non-nil, zero value otherwise.

### GetAllocationIdOk

`func (o *AwsElasticIpResourcesDefStatus) GetAllocationIdOk() (*string, bool)`

GetAllocationIdOk returns a tuple with the AllocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocationId

`func (o *AwsElasticIpResourcesDefStatus) SetAllocationId(v string)`

SetAllocationId sets AllocationId field to given value.

### HasAllocationId

`func (o *AwsElasticIpResourcesDefStatus) HasAllocationId() bool`

HasAllocationId returns a boolean if a field has been set.

### GetInstanceId

`func (o *AwsElasticIpResourcesDefStatus) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *AwsElasticIpResourcesDefStatus) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *AwsElasticIpResourcesDefStatus) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.

### HasInstanceId

`func (o *AwsElasticIpResourcesDefStatus) HasInstanceId() bool`

HasInstanceId returns a boolean if a field has been set.

### GetAssociationId

`func (o *AwsElasticIpResourcesDefStatus) GetAssociationId() string`

GetAssociationId returns the AssociationId field if non-nil, zero value otherwise.

### GetAssociationIdOk

`func (o *AwsElasticIpResourcesDefStatus) GetAssociationIdOk() (*string, bool)`

GetAssociationIdOk returns a tuple with the AssociationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociationId

`func (o *AwsElasticIpResourcesDefStatus) SetAssociationId(v string)`

SetAssociationId sets AssociationId field to given value.

### HasAssociationId

`func (o *AwsElasticIpResourcesDefStatus) HasAssociationId() bool`

HasAssociationId returns a boolean if a field has been set.

### GetNetworkInterfaceId

`func (o *AwsElasticIpResourcesDefStatus) GetNetworkInterfaceId() string`

GetNetworkInterfaceId returns the NetworkInterfaceId field if non-nil, zero value otherwise.

### GetNetworkInterfaceIdOk

`func (o *AwsElasticIpResourcesDefStatus) GetNetworkInterfaceIdOk() (*string, bool)`

GetNetworkInterfaceIdOk returns a tuple with the NetworkInterfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInterfaceId

`func (o *AwsElasticIpResourcesDefStatus) SetNetworkInterfaceId(v string)`

SetNetworkInterfaceId sets NetworkInterfaceId field to given value.

### HasNetworkInterfaceId

`func (o *AwsElasticIpResourcesDefStatus) HasNetworkInterfaceId() bool`

HasNetworkInterfaceId returns a boolean if a field has been set.

### GetPrivateIpAddress

`func (o *AwsElasticIpResourcesDefStatus) GetPrivateIpAddress() string`

GetPrivateIpAddress returns the PrivateIpAddress field if non-nil, zero value otherwise.

### GetPrivateIpAddressOk

`func (o *AwsElasticIpResourcesDefStatus) GetPrivateIpAddressOk() (*string, bool)`

GetPrivateIpAddressOk returns a tuple with the PrivateIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateIpAddress

`func (o *AwsElasticIpResourcesDefStatus) SetPrivateIpAddress(v string)`

SetPrivateIpAddress sets PrivateIpAddress field to given value.

### HasPrivateIpAddress

`func (o *AwsElasticIpResourcesDefStatus) HasPrivateIpAddress() bool`

HasPrivateIpAddress returns a boolean if a field has been set.

### GetNetworkInterfaceOwnerId

`func (o *AwsElasticIpResourcesDefStatus) GetNetworkInterfaceOwnerId() string`

GetNetworkInterfaceOwnerId returns the NetworkInterfaceOwnerId field if non-nil, zero value otherwise.

### GetNetworkInterfaceOwnerIdOk

`func (o *AwsElasticIpResourcesDefStatus) GetNetworkInterfaceOwnerIdOk() (*string, bool)`

GetNetworkInterfaceOwnerIdOk returns a tuple with the NetworkInterfaceOwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInterfaceOwnerId

`func (o *AwsElasticIpResourcesDefStatus) SetNetworkInterfaceOwnerId(v string)`

SetNetworkInterfaceOwnerId sets NetworkInterfaceOwnerId field to given value.

### HasNetworkInterfaceOwnerId

`func (o *AwsElasticIpResourcesDefStatus) HasNetworkInterfaceOwnerId() bool`

HasNetworkInterfaceOwnerId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# VcenterDeploymentDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VcenterNetworkName** | Pointer to **string** | vcenter network name to which the vpn vm nic is attached.  | [optional] 
**VcenterDatacenterName** | Pointer to **string** | vcenter datacenter that the cluster belongs to. required when interacting with vcenter apis to deploy the vpn vm.  | [optional] 
**VcenterDatastoreName** | Pointer to **string** | vcenter datastore to which the vpn disks and images will be uploaded during deployment.  | [optional] 

## Methods

### NewVcenterDeploymentDetails

`func NewVcenterDeploymentDetails() *VcenterDeploymentDetails`

NewVcenterDeploymentDetails instantiates a new VcenterDeploymentDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterDeploymentDetailsWithDefaults

`func NewVcenterDeploymentDetailsWithDefaults() *VcenterDeploymentDetails`

NewVcenterDeploymentDetailsWithDefaults instantiates a new VcenterDeploymentDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVcenterNetworkName

`func (o *VcenterDeploymentDetails) GetVcenterNetworkName() string`

GetVcenterNetworkName returns the VcenterNetworkName field if non-nil, zero value otherwise.

### GetVcenterNetworkNameOk

`func (o *VcenterDeploymentDetails) GetVcenterNetworkNameOk() (*string, bool)`

GetVcenterNetworkNameOk returns a tuple with the VcenterNetworkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcenterNetworkName

`func (o *VcenterDeploymentDetails) SetVcenterNetworkName(v string)`

SetVcenterNetworkName sets VcenterNetworkName field to given value.

### HasVcenterNetworkName

`func (o *VcenterDeploymentDetails) HasVcenterNetworkName() bool`

HasVcenterNetworkName returns a boolean if a field has been set.

### GetVcenterDatacenterName

`func (o *VcenterDeploymentDetails) GetVcenterDatacenterName() string`

GetVcenterDatacenterName returns the VcenterDatacenterName field if non-nil, zero value otherwise.

### GetVcenterDatacenterNameOk

`func (o *VcenterDeploymentDetails) GetVcenterDatacenterNameOk() (*string, bool)`

GetVcenterDatacenterNameOk returns a tuple with the VcenterDatacenterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcenterDatacenterName

`func (o *VcenterDeploymentDetails) SetVcenterDatacenterName(v string)`

SetVcenterDatacenterName sets VcenterDatacenterName field to given value.

### HasVcenterDatacenterName

`func (o *VcenterDeploymentDetails) HasVcenterDatacenterName() bool`

HasVcenterDatacenterName returns a boolean if a field has been set.

### GetVcenterDatastoreName

`func (o *VcenterDeploymentDetails) GetVcenterDatastoreName() string`

GetVcenterDatastoreName returns the VcenterDatastoreName field if non-nil, zero value otherwise.

### GetVcenterDatastoreNameOk

`func (o *VcenterDeploymentDetails) GetVcenterDatastoreNameOk() (*string, bool)`

GetVcenterDatastoreNameOk returns a tuple with the VcenterDatastoreName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcenterDatastoreName

`func (o *VcenterDeploymentDetails) SetVcenterDatastoreName(v string)`

SetVcenterDatastoreName sets VcenterDatastoreName field to given value.

### HasVcenterDatastoreName

`func (o *VcenterDeploymentDetails) HasVcenterDatastoreName() bool`

HasVcenterDatastoreName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



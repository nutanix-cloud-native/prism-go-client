# RecoveryPlanResourcesParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FloatingIpAssignmentList** | Pointer to [**[]RecoveryPlanResourcesParametersFloatingIpAssignmentListInner**](RecoveryPlanResourcesParametersFloatingIpAssignmentListInner.md) | Floating IP assignment for VMs upon recovery in an Availability Zone. This is applicable only for the public cloud Availability Zones.  | [optional] 
**WitnessConfigurationList** | Pointer to [**[]WitnessConfiguration**](WitnessConfiguration.md) | A list containing witness configuration.  | [optional] 
**CutoverMode** | Pointer to **string** | An option to specify cutover mode for VM/Volume Group when restoring data from recovery point present at any storage over the network. EARLY - VM/Volume Group will be made available for active   consumption before completely hydrating the data for them from   the recovery point available over the remote network storage.   Hydration of data will keep happening in the background for the   recovered VMs/Volume Groups. LATE - VM/Volume Group will be made available for active   consumption after completely hydrating the data for them from   the recovery point available over the remote network storage.  | [optional] 
**NetworkMappingList** | Pointer to [**[]RecoveryPlanResourcesParametersNetworkMappingListInner**](RecoveryPlanResourcesParametersNetworkMappingListInner.md) | Network mappings to be used for the Recovery Plan. This will be represented by array of network mappings across the Availability Zones. Each entry of network mapping will have Availability Zone URL, recovery and test network information, static IP assignment for the VMs for the recovery and test networks. For example, Let RNx, TNx denote the recovery and test network information and RIPMx, TIPMx denote the static IP assignment for the VMs. As per below matrix, while performing failover action from Availability Zone AZ1 to AZ2, RN1 will be mapped to RN4 and static IP \&quot;a.b.c.d\&quot; of the VM with reference VMx will be mapped to \&quot;i.j.k.l\&quot;. On performing test failover from AZ1 to AZ2, RN1 will be mapped to TN4 and static IP \&quot;a.b.c.d\&quot; of the VM with reference VMx will be mapped to \&quot;I.J.K.L\&quot;. [[(AZ1 URL, RN1, TN1, [{VMx, \&quot;a.b.c.d\&quot;}, {VMy, \&quot;e.f.g.h\&quot;}],    [{VMx, \&quot;A.B.C.D\&quot;}, {VMy, \&quot;E.F.G.H\&quot;}]),   (AZ2 URL, RN4, TN4, [{VMx, \&quot;i.j.k.l\&quot;}, {VMy, \&quot;m.n.p.q\&quot;}],    [{VMx, \&quot;I.J.K.L\&quot;}, {VMy, \&quot;M.N.P.Q\&quot;}]),   (AZ3 URL, RN7, TN7)],  [(AZ1 URL, RN2, TN2), (AZ2 URL, RN5, TN5),   (AZ3 URL, RN8, TN8)],  [(AZ1 URL, RN3, TN3), (AZ2 URL, RN6, TN6)]] The order of the static IP assignment for the VMs should remain same across all the networks provided in a network mapping. It forms a matrix containing IP mapping for the VMs across the Availability Zones. If a VM has multiple static IP addresses, static IP mapping will happen on the array indices. For example, Let IPrxy denote an IP address in subnet RNz and IPtxy denote an IP address in the subnet TNz. Lets VMx has two static IP address and VMy, VMz has one static IP associated with a vNIC created in network RN1. The IP mapping for the VMs will be as below.     AZ1, RN1        AZ1, TN1        AZ2, RN3        AZ2, TN2 [(VMx, [IPr11]), (VMx, [IPt11]), (VMx, [IPr12]), (VMx,  [IPt12])] [(VMy, [IPr21]), (VMy, [IPt21]), (VMy, [IPr22]), (VMy,  [IPt22])] [(VMz, [IPr31]), (VMz, [IPt31]), (VMz, [IPr32]), (VMz,  [IPt32])] In case of recovery of VMs from one Prism Element to the other within the same Availability Zone, a list of cluster references where the network exists can be specified. The network mapping to be used for a vNIC is decided as follows - 1. If a VM that has a vNIC in a network N1 on cluster C1, then    the network mapping of N1 that has C1 in the cluster    references list will be used.  2. In case there is no network mapping for N1 with cluster C1,    the default mapping of N1 for the Availability Zone (in which    cluster is not specified), will be used if present.  | [optional] 
**DataServiceIpMappingList** | Pointer to [**[]RecoveryPlanResourcesParametersDataServiceIpMappingListInner**](RecoveryPlanResourcesParametersDataServiceIpMappingListInner.md) | IP address mappings for attaching Volume Groups to VMs upon failover.  | [optional] 
**AvailabilityZoneList** | Pointer to [**[]AvailabilityZoneInformation**](AvailabilityZoneInformation.md) | A list containing information about primary and secondary Availability zones.  | [optional] 
**PrimaryLocationIndex** | Pointer to **int32** | This field should be set to the index of the Availability Zone in the availability_zone_list which should be considered as a protected Availability Zone.  | [optional] 

## Methods

### NewRecoveryPlanResourcesParameters

`func NewRecoveryPlanResourcesParameters() *RecoveryPlanResourcesParameters`

NewRecoveryPlanResourcesParameters instantiates a new RecoveryPlanResourcesParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoveryPlanResourcesParametersWithDefaults

`func NewRecoveryPlanResourcesParametersWithDefaults() *RecoveryPlanResourcesParameters`

NewRecoveryPlanResourcesParametersWithDefaults instantiates a new RecoveryPlanResourcesParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFloatingIpAssignmentList

`func (o *RecoveryPlanResourcesParameters) GetFloatingIpAssignmentList() []RecoveryPlanResourcesParametersFloatingIpAssignmentListInner`

GetFloatingIpAssignmentList returns the FloatingIpAssignmentList field if non-nil, zero value otherwise.

### GetFloatingIpAssignmentListOk

`func (o *RecoveryPlanResourcesParameters) GetFloatingIpAssignmentListOk() (*[]RecoveryPlanResourcesParametersFloatingIpAssignmentListInner, bool)`

GetFloatingIpAssignmentListOk returns a tuple with the FloatingIpAssignmentList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFloatingIpAssignmentList

`func (o *RecoveryPlanResourcesParameters) SetFloatingIpAssignmentList(v []RecoveryPlanResourcesParametersFloatingIpAssignmentListInner)`

SetFloatingIpAssignmentList sets FloatingIpAssignmentList field to given value.

### HasFloatingIpAssignmentList

`func (o *RecoveryPlanResourcesParameters) HasFloatingIpAssignmentList() bool`

HasFloatingIpAssignmentList returns a boolean if a field has been set.

### GetWitnessConfigurationList

`func (o *RecoveryPlanResourcesParameters) GetWitnessConfigurationList() []WitnessConfiguration`

GetWitnessConfigurationList returns the WitnessConfigurationList field if non-nil, zero value otherwise.

### GetWitnessConfigurationListOk

`func (o *RecoveryPlanResourcesParameters) GetWitnessConfigurationListOk() (*[]WitnessConfiguration, bool)`

GetWitnessConfigurationListOk returns a tuple with the WitnessConfigurationList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWitnessConfigurationList

`func (o *RecoveryPlanResourcesParameters) SetWitnessConfigurationList(v []WitnessConfiguration)`

SetWitnessConfigurationList sets WitnessConfigurationList field to given value.

### HasWitnessConfigurationList

`func (o *RecoveryPlanResourcesParameters) HasWitnessConfigurationList() bool`

HasWitnessConfigurationList returns a boolean if a field has been set.

### GetCutoverMode

`func (o *RecoveryPlanResourcesParameters) GetCutoverMode() string`

GetCutoverMode returns the CutoverMode field if non-nil, zero value otherwise.

### GetCutoverModeOk

`func (o *RecoveryPlanResourcesParameters) GetCutoverModeOk() (*string, bool)`

GetCutoverModeOk returns a tuple with the CutoverMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCutoverMode

`func (o *RecoveryPlanResourcesParameters) SetCutoverMode(v string)`

SetCutoverMode sets CutoverMode field to given value.

### HasCutoverMode

`func (o *RecoveryPlanResourcesParameters) HasCutoverMode() bool`

HasCutoverMode returns a boolean if a field has been set.

### GetNetworkMappingList

`func (o *RecoveryPlanResourcesParameters) GetNetworkMappingList() []RecoveryPlanResourcesParametersNetworkMappingListInner`

GetNetworkMappingList returns the NetworkMappingList field if non-nil, zero value otherwise.

### GetNetworkMappingListOk

`func (o *RecoveryPlanResourcesParameters) GetNetworkMappingListOk() (*[]RecoveryPlanResourcesParametersNetworkMappingListInner, bool)`

GetNetworkMappingListOk returns a tuple with the NetworkMappingList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkMappingList

`func (o *RecoveryPlanResourcesParameters) SetNetworkMappingList(v []RecoveryPlanResourcesParametersNetworkMappingListInner)`

SetNetworkMappingList sets NetworkMappingList field to given value.

### HasNetworkMappingList

`func (o *RecoveryPlanResourcesParameters) HasNetworkMappingList() bool`

HasNetworkMappingList returns a boolean if a field has been set.

### GetDataServiceIpMappingList

`func (o *RecoveryPlanResourcesParameters) GetDataServiceIpMappingList() []RecoveryPlanResourcesParametersDataServiceIpMappingListInner`

GetDataServiceIpMappingList returns the DataServiceIpMappingList field if non-nil, zero value otherwise.

### GetDataServiceIpMappingListOk

`func (o *RecoveryPlanResourcesParameters) GetDataServiceIpMappingListOk() (*[]RecoveryPlanResourcesParametersDataServiceIpMappingListInner, bool)`

GetDataServiceIpMappingListOk returns a tuple with the DataServiceIpMappingList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataServiceIpMappingList

`func (o *RecoveryPlanResourcesParameters) SetDataServiceIpMappingList(v []RecoveryPlanResourcesParametersDataServiceIpMappingListInner)`

SetDataServiceIpMappingList sets DataServiceIpMappingList field to given value.

### HasDataServiceIpMappingList

`func (o *RecoveryPlanResourcesParameters) HasDataServiceIpMappingList() bool`

HasDataServiceIpMappingList returns a boolean if a field has been set.

### GetAvailabilityZoneList

`func (o *RecoveryPlanResourcesParameters) GetAvailabilityZoneList() []AvailabilityZoneInformation`

GetAvailabilityZoneList returns the AvailabilityZoneList field if non-nil, zero value otherwise.

### GetAvailabilityZoneListOk

`func (o *RecoveryPlanResourcesParameters) GetAvailabilityZoneListOk() (*[]AvailabilityZoneInformation, bool)`

GetAvailabilityZoneListOk returns a tuple with the AvailabilityZoneList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZoneList

`func (o *RecoveryPlanResourcesParameters) SetAvailabilityZoneList(v []AvailabilityZoneInformation)`

SetAvailabilityZoneList sets AvailabilityZoneList field to given value.

### HasAvailabilityZoneList

`func (o *RecoveryPlanResourcesParameters) HasAvailabilityZoneList() bool`

HasAvailabilityZoneList returns a boolean if a field has been set.

### GetPrimaryLocationIndex

`func (o *RecoveryPlanResourcesParameters) GetPrimaryLocationIndex() int32`

GetPrimaryLocationIndex returns the PrimaryLocationIndex field if non-nil, zero value otherwise.

### GetPrimaryLocationIndexOk

`func (o *RecoveryPlanResourcesParameters) GetPrimaryLocationIndexOk() (*int32, bool)`

GetPrimaryLocationIndexOk returns a tuple with the PrimaryLocationIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryLocationIndex

`func (o *RecoveryPlanResourcesParameters) SetPrimaryLocationIndex(v int32)`

SetPrimaryLocationIndex sets PrimaryLocationIndex field to given value.

### HasPrimaryLocationIndex

`func (o *RecoveryPlanResourcesParameters) HasPrimaryLocationIndex() bool`

HasPrimaryLocationIndex returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



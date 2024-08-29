# RecoveryPlanResourcesParametersNetworkMappingListInnerAvailabilityZoneNetworkMappingListInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecoveryIpAssignmentList** | Pointer to [**[]RecoveryPlanVmIpAssignment**](RecoveryPlanVmIpAssignment.md) | Static IP configuration for the VMs to be applied post recovery in the recovery network for migrate/ failover action on the Recovery Plan.  | [optional] 
**TestIpAssignmentList** | Pointer to [**[]RecoveryPlanVmIpAssignment**](RecoveryPlanVmIpAssignment.md) | Static IP configuration for the VMs to be applied post recovery in the test network for test failover action on the Recovery Plan.  | [optional] 
**AvailabilityZoneUrl** | **string** | URL of the Availability Zone.  | 
**RecoveryNetwork** | Pointer to [**RecoveryPlanNetwork**](RecoveryPlanNetwork.md) |  | [optional] 
**TestNetwork** | Pointer to [**RecoveryPlanNetwork**](RecoveryPlanNetwork.md) |  | [optional] 
**ClusterReferenceList** | Pointer to [**[]ClusterReference**](ClusterReference.md) | The clusters where the recovery and test networks reside. This is required to specify network mapping across clusters for a Recovery Plan created to handle failover within the same Availability Zone.  | [optional] 

## Methods

### NewRecoveryPlanResourcesParametersNetworkMappingListInnerAvailabilityZoneNetworkMappingListInner

`func NewRecoveryPlanResourcesParametersNetworkMappingListInnerAvailabilityZoneNetworkMappingListInner(availabilityZoneUrl string, ) *RecoveryPlanResourcesParametersNetworkMappingListInnerAvailabilityZoneNetworkMappingListInner`

NewRecoveryPlanResourcesParametersNetworkMappingListInnerAvailabilityZoneNetworkMappingListInner instantiates a new RecoveryPlanResourcesParametersNetworkMappingListInnerAvailabilityZoneNetworkMappingListInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoveryPlanResourcesParametersNetworkMappingListInnerAvailabilityZoneNetworkMappingListInnerWithDefaults

`func NewRecoveryPlanResourcesParametersNetworkMappingListInnerAvailabilityZoneNetworkMappingListInnerWithDefaults() *RecoveryPlanResourcesParametersNetworkMappingListInnerAvailabilityZoneNetworkMappingListInner`

NewRecoveryPlanResourcesParametersNetworkMappingListInnerAvailabilityZoneNetworkMappingListInnerWithDefaults instantiates a new RecoveryPlanResourcesParametersNetworkMappingListInnerAvailabilityZoneNetworkMappingListInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecoveryIpAssignmentList

`func (o *RecoveryPlanResourcesParametersNetworkMappingListInnerAvailabilityZoneNetworkMappingListInner) GetRecoveryIpAssignmentList() []RecoveryPlanVmIpAssignment`

GetRecoveryIpAssignmentList returns the RecoveryIpAssignmentList field if non-nil, zero value otherwise.

### GetRecoveryIpAssignmentListOk

`func (o *RecoveryPlanResourcesParametersNetworkMappingListInnerAvailabilityZoneNetworkMappingListInner) GetRecoveryIpAssignmentListOk() (*[]RecoveryPlanVmIpAssignment, bool)`

GetRecoveryIpAssignmentListOk returns a tuple with the RecoveryIpAssignmentList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryIpAssignmentList

`func (o *RecoveryPlanResourcesParametersNetworkMappingListInnerAvailabilityZoneNetworkMappingListInner) SetRecoveryIpAssignmentList(v []RecoveryPlanVmIpAssignment)`

SetRecoveryIpAssignmentList sets RecoveryIpAssignmentList field to given value.

### HasRecoveryIpAssignmentList

`func (o *RecoveryPlanResourcesParametersNetworkMappingListInnerAvailabilityZoneNetworkMappingListInner) HasRecoveryIpAssignmentList() bool`

HasRecoveryIpAssignmentList returns a boolean if a field has been set.

### GetTestIpAssignmentList

`func (o *RecoveryPlanResourcesParametersNetworkMappingListInnerAvailabilityZoneNetworkMappingListInner) GetTestIpAssignmentList() []RecoveryPlanVmIpAssignment`

GetTestIpAssignmentList returns the TestIpAssignmentList field if non-nil, zero value otherwise.

### GetTestIpAssignmentListOk

`func (o *RecoveryPlanResourcesParametersNetworkMappingListInnerAvailabilityZoneNetworkMappingListInner) GetTestIpAssignmentListOk() (*[]RecoveryPlanVmIpAssignment, bool)`

GetTestIpAssignmentListOk returns a tuple with the TestIpAssignmentList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestIpAssignmentList

`func (o *RecoveryPlanResourcesParametersNetworkMappingListInnerAvailabilityZoneNetworkMappingListInner) SetTestIpAssignmentList(v []RecoveryPlanVmIpAssignment)`

SetTestIpAssignmentList sets TestIpAssignmentList field to given value.

### HasTestIpAssignmentList

`func (o *RecoveryPlanResourcesParametersNetworkMappingListInnerAvailabilityZoneNetworkMappingListInner) HasTestIpAssignmentList() bool`

HasTestIpAssignmentList returns a boolean if a field has been set.

### GetAvailabilityZoneUrl

`func (o *RecoveryPlanResourcesParametersNetworkMappingListInnerAvailabilityZoneNetworkMappingListInner) GetAvailabilityZoneUrl() string`

GetAvailabilityZoneUrl returns the AvailabilityZoneUrl field if non-nil, zero value otherwise.

### GetAvailabilityZoneUrlOk

`func (o *RecoveryPlanResourcesParametersNetworkMappingListInnerAvailabilityZoneNetworkMappingListInner) GetAvailabilityZoneUrlOk() (*string, bool)`

GetAvailabilityZoneUrlOk returns a tuple with the AvailabilityZoneUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZoneUrl

`func (o *RecoveryPlanResourcesParametersNetworkMappingListInnerAvailabilityZoneNetworkMappingListInner) SetAvailabilityZoneUrl(v string)`

SetAvailabilityZoneUrl sets AvailabilityZoneUrl field to given value.


### GetRecoveryNetwork

`func (o *RecoveryPlanResourcesParametersNetworkMappingListInnerAvailabilityZoneNetworkMappingListInner) GetRecoveryNetwork() RecoveryPlanNetwork`

GetRecoveryNetwork returns the RecoveryNetwork field if non-nil, zero value otherwise.

### GetRecoveryNetworkOk

`func (o *RecoveryPlanResourcesParametersNetworkMappingListInnerAvailabilityZoneNetworkMappingListInner) GetRecoveryNetworkOk() (*RecoveryPlanNetwork, bool)`

GetRecoveryNetworkOk returns a tuple with the RecoveryNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryNetwork

`func (o *RecoveryPlanResourcesParametersNetworkMappingListInnerAvailabilityZoneNetworkMappingListInner) SetRecoveryNetwork(v RecoveryPlanNetwork)`

SetRecoveryNetwork sets RecoveryNetwork field to given value.

### HasRecoveryNetwork

`func (o *RecoveryPlanResourcesParametersNetworkMappingListInnerAvailabilityZoneNetworkMappingListInner) HasRecoveryNetwork() bool`

HasRecoveryNetwork returns a boolean if a field has been set.

### GetTestNetwork

`func (o *RecoveryPlanResourcesParametersNetworkMappingListInnerAvailabilityZoneNetworkMappingListInner) GetTestNetwork() RecoveryPlanNetwork`

GetTestNetwork returns the TestNetwork field if non-nil, zero value otherwise.

### GetTestNetworkOk

`func (o *RecoveryPlanResourcesParametersNetworkMappingListInnerAvailabilityZoneNetworkMappingListInner) GetTestNetworkOk() (*RecoveryPlanNetwork, bool)`

GetTestNetworkOk returns a tuple with the TestNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestNetwork

`func (o *RecoveryPlanResourcesParametersNetworkMappingListInnerAvailabilityZoneNetworkMappingListInner) SetTestNetwork(v RecoveryPlanNetwork)`

SetTestNetwork sets TestNetwork field to given value.

### HasTestNetwork

`func (o *RecoveryPlanResourcesParametersNetworkMappingListInnerAvailabilityZoneNetworkMappingListInner) HasTestNetwork() bool`

HasTestNetwork returns a boolean if a field has been set.

### GetClusterReferenceList

`func (o *RecoveryPlanResourcesParametersNetworkMappingListInnerAvailabilityZoneNetworkMappingListInner) GetClusterReferenceList() []ClusterReference`

GetClusterReferenceList returns the ClusterReferenceList field if non-nil, zero value otherwise.

### GetClusterReferenceListOk

`func (o *RecoveryPlanResourcesParametersNetworkMappingListInnerAvailabilityZoneNetworkMappingListInner) GetClusterReferenceListOk() (*[]ClusterReference, bool)`

GetClusterReferenceListOk returns a tuple with the ClusterReferenceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterReferenceList

`func (o *RecoveryPlanResourcesParametersNetworkMappingListInnerAvailabilityZoneNetworkMappingListInner) SetClusterReferenceList(v []ClusterReference)`

SetClusterReferenceList sets ClusterReferenceList field to given value.

### HasClusterReferenceList

`func (o *RecoveryPlanResourcesParametersNetworkMappingListInnerAvailabilityZoneNetworkMappingListInner) HasClusterReferenceList() bool`

HasClusterReferenceList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# RecoveryPlanDataServiceIpConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecoveryDataServiceIp** | **string** | Data Services IP address to be reconfigured in VM during Planned and Unplanned Failover.  | 
**AvailabilityZoneUrl** | **string** | URL of the Availability Zone. | 
**ClusterReference** | [**ClusterReference**](ClusterReference.md) |  | 
**TestDataServiceIp** | Pointer to **string** | Data Services IP address to be reconfigured in VM during Test Failover.  | [optional] 

## Methods

### NewRecoveryPlanDataServiceIpConfig

`func NewRecoveryPlanDataServiceIpConfig(recoveryDataServiceIp string, availabilityZoneUrl string, clusterReference ClusterReference, ) *RecoveryPlanDataServiceIpConfig`

NewRecoveryPlanDataServiceIpConfig instantiates a new RecoveryPlanDataServiceIpConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoveryPlanDataServiceIpConfigWithDefaults

`func NewRecoveryPlanDataServiceIpConfigWithDefaults() *RecoveryPlanDataServiceIpConfig`

NewRecoveryPlanDataServiceIpConfigWithDefaults instantiates a new RecoveryPlanDataServiceIpConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecoveryDataServiceIp

`func (o *RecoveryPlanDataServiceIpConfig) GetRecoveryDataServiceIp() string`

GetRecoveryDataServiceIp returns the RecoveryDataServiceIp field if non-nil, zero value otherwise.

### GetRecoveryDataServiceIpOk

`func (o *RecoveryPlanDataServiceIpConfig) GetRecoveryDataServiceIpOk() (*string, bool)`

GetRecoveryDataServiceIpOk returns a tuple with the RecoveryDataServiceIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryDataServiceIp

`func (o *RecoveryPlanDataServiceIpConfig) SetRecoveryDataServiceIp(v string)`

SetRecoveryDataServiceIp sets RecoveryDataServiceIp field to given value.


### GetAvailabilityZoneUrl

`func (o *RecoveryPlanDataServiceIpConfig) GetAvailabilityZoneUrl() string`

GetAvailabilityZoneUrl returns the AvailabilityZoneUrl field if non-nil, zero value otherwise.

### GetAvailabilityZoneUrlOk

`func (o *RecoveryPlanDataServiceIpConfig) GetAvailabilityZoneUrlOk() (*string, bool)`

GetAvailabilityZoneUrlOk returns a tuple with the AvailabilityZoneUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZoneUrl

`func (o *RecoveryPlanDataServiceIpConfig) SetAvailabilityZoneUrl(v string)`

SetAvailabilityZoneUrl sets AvailabilityZoneUrl field to given value.


### GetClusterReference

`func (o *RecoveryPlanDataServiceIpConfig) GetClusterReference() ClusterReference`

GetClusterReference returns the ClusterReference field if non-nil, zero value otherwise.

### GetClusterReferenceOk

`func (o *RecoveryPlanDataServiceIpConfig) GetClusterReferenceOk() (*ClusterReference, bool)`

GetClusterReferenceOk returns a tuple with the ClusterReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterReference

`func (o *RecoveryPlanDataServiceIpConfig) SetClusterReference(v ClusterReference)`

SetClusterReference sets ClusterReference field to given value.


### GetTestDataServiceIp

`func (o *RecoveryPlanDataServiceIpConfig) GetTestDataServiceIp() string`

GetTestDataServiceIp returns the TestDataServiceIp field if non-nil, zero value otherwise.

### GetTestDataServiceIpOk

`func (o *RecoveryPlanDataServiceIpConfig) GetTestDataServiceIpOk() (*string, bool)`

GetTestDataServiceIpOk returns a tuple with the TestDataServiceIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestDataServiceIp

`func (o *RecoveryPlanDataServiceIpConfig) SetTestDataServiceIp(v string)`

SetTestDataServiceIp sets TestDataServiceIp field to given value.

### HasTestDataServiceIp

`func (o *RecoveryPlanDataServiceIpConfig) HasTestDataServiceIp() bool`

HasTestDataServiceIp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



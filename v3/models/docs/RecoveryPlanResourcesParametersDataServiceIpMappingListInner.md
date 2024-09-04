# RecoveryPlanResourcesParametersDataServiceIpMappingListInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataServiceIpMapping** | Pointer to [**[]RecoveryPlanDataServiceIpConfig**](RecoveryPlanDataServiceIpConfig.md) | Data Services IP address mapping, each entry in this mapping will include Availability Zone URL, Cluster reference, recovery and test data services IP. During the Volume Groups attachment step of Recovery Plan failover operation , the data service IP specified for the target Cluster will be reconfigured in the VMs.  | [optional] 

## Methods

### NewRecoveryPlanResourcesParametersDataServiceIpMappingListInner

`func NewRecoveryPlanResourcesParametersDataServiceIpMappingListInner() *RecoveryPlanResourcesParametersDataServiceIpMappingListInner`

NewRecoveryPlanResourcesParametersDataServiceIpMappingListInner instantiates a new RecoveryPlanResourcesParametersDataServiceIpMappingListInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoveryPlanResourcesParametersDataServiceIpMappingListInnerWithDefaults

`func NewRecoveryPlanResourcesParametersDataServiceIpMappingListInnerWithDefaults() *RecoveryPlanResourcesParametersDataServiceIpMappingListInner`

NewRecoveryPlanResourcesParametersDataServiceIpMappingListInnerWithDefaults instantiates a new RecoveryPlanResourcesParametersDataServiceIpMappingListInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataServiceIpMapping

`func (o *RecoveryPlanResourcesParametersDataServiceIpMappingListInner) GetDataServiceIpMapping() []RecoveryPlanDataServiceIpConfig`

GetDataServiceIpMapping returns the DataServiceIpMapping field if non-nil, zero value otherwise.

### GetDataServiceIpMappingOk

`func (o *RecoveryPlanResourcesParametersDataServiceIpMappingListInner) GetDataServiceIpMappingOk() (*[]RecoveryPlanDataServiceIpConfig, bool)`

GetDataServiceIpMappingOk returns a tuple with the DataServiceIpMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataServiceIpMapping

`func (o *RecoveryPlanResourcesParametersDataServiceIpMappingListInner) SetDataServiceIpMapping(v []RecoveryPlanDataServiceIpConfig)`

SetDataServiceIpMapping sets DataServiceIpMapping field to given value.

### HasDataServiceIpMapping

`func (o *RecoveryPlanResourcesParametersDataServiceIpMappingListInner) HasDataServiceIpMapping() bool`

HasDataServiceIpMapping returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# WhatifScenario

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NewCluster** | Pointer to **bool** | The flag to indicate whether it is a new cluster or not. | [optional] 
**ClusterEntityType** | Pointer to **string** | The entity type for the cluster e.g. cluster or nutanix_vcenter__cluster. | [optional] [default to "cluster"]
**UUID** | Pointer to **string** | The uuid would be automatically generated when created. | [optional] 
**VendorList** | Pointer to **[]string** |  | [optional] 
**WorkloadList** | Pointer to [**[]Workload**](Workload.md) | workload added by user. | [optional] 
**RecommendedRunway** | Pointer to [**Runway**](Runway.md) |  | [optional] 
**UpdatedTimeSec** | Pointer to **int32** | Last updated timestamp. | [optional] 
**ClusterUuid** | Pointer to **string** | The cluster uuid. | [optional] 
**TargetRunwayDays** | Pointer to **int32** | The target runway. | [optional] [default to 180]
**ClusterSpec** | Pointer to [**ClusterSpec**](ClusterSpec.md) |  | [optional] 
**Runway** | Pointer to [**Runway**](Runway.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 

## Methods

### NewWhatifScenario

`func NewWhatifScenario() *WhatifScenario`

NewWhatifScenario instantiates a new WhatifScenario object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatifScenarioWithDefaults

`func NewWhatifScenarioWithDefaults() *WhatifScenario`

NewWhatifScenarioWithDefaults instantiates a new WhatifScenario object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNewCluster

`func (o *WhatifScenario) GetNewCluster() bool`

GetNewCluster returns the NewCluster field if non-nil, zero value otherwise.

### GetNewClusterOk

`func (o *WhatifScenario) GetNewClusterOk() (*bool, bool)`

GetNewClusterOk returns a tuple with the NewCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewCluster

`func (o *WhatifScenario) SetNewCluster(v bool)`

SetNewCluster sets NewCluster field to given value.

### HasNewCluster

`func (o *WhatifScenario) HasNewCluster() bool`

HasNewCluster returns a boolean if a field has been set.

### GetClusterEntityType

`func (o *WhatifScenario) GetClusterEntityType() string`

GetClusterEntityType returns the ClusterEntityType field if non-nil, zero value otherwise.

### GetClusterEntityTypeOk

`func (o *WhatifScenario) GetClusterEntityTypeOk() (*string, bool)`

GetClusterEntityTypeOk returns a tuple with the ClusterEntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterEntityType

`func (o *WhatifScenario) SetClusterEntityType(v string)`

SetClusterEntityType sets ClusterEntityType field to given value.

### HasClusterEntityType

`func (o *WhatifScenario) HasClusterEntityType() bool`

HasClusterEntityType returns a boolean if a field has been set.

### GetUUID

`func (o *WhatifScenario) GetUUID() string`

GetUUID returns the UUID field if non-nil, zero value otherwise.

### GetUUIDOk

`func (o *WhatifScenario) GetUUIDOk() (*string, bool)`

GetUUIDOk returns a tuple with the UUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUUID

`func (o *WhatifScenario) SetUUID(v string)`

SetUUID sets UUID field to given value.

### HasUUID

`func (o *WhatifScenario) HasUUID() bool`

HasUUID returns a boolean if a field has been set.

### GetVendorList

`func (o *WhatifScenario) GetVendorList() []string`

GetVendorList returns the VendorList field if non-nil, zero value otherwise.

### GetVendorListOk

`func (o *WhatifScenario) GetVendorListOk() (*[]string, bool)`

GetVendorListOk returns a tuple with the VendorList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorList

`func (o *WhatifScenario) SetVendorList(v []string)`

SetVendorList sets VendorList field to given value.

### HasVendorList

`func (o *WhatifScenario) HasVendorList() bool`

HasVendorList returns a boolean if a field has been set.

### GetWorkloadList

`func (o *WhatifScenario) GetWorkloadList() []Workload`

GetWorkloadList returns the WorkloadList field if non-nil, zero value otherwise.

### GetWorkloadListOk

`func (o *WhatifScenario) GetWorkloadListOk() (*[]Workload, bool)`

GetWorkloadListOk returns a tuple with the WorkloadList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkloadList

`func (o *WhatifScenario) SetWorkloadList(v []Workload)`

SetWorkloadList sets WorkloadList field to given value.

### HasWorkloadList

`func (o *WhatifScenario) HasWorkloadList() bool`

HasWorkloadList returns a boolean if a field has been set.

### GetRecommendedRunway

`func (o *WhatifScenario) GetRecommendedRunway() Runway`

GetRecommendedRunway returns the RecommendedRunway field if non-nil, zero value otherwise.

### GetRecommendedRunwayOk

`func (o *WhatifScenario) GetRecommendedRunwayOk() (*Runway, bool)`

GetRecommendedRunwayOk returns a tuple with the RecommendedRunway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendedRunway

`func (o *WhatifScenario) SetRecommendedRunway(v Runway)`

SetRecommendedRunway sets RecommendedRunway field to given value.

### HasRecommendedRunway

`func (o *WhatifScenario) HasRecommendedRunway() bool`

HasRecommendedRunway returns a boolean if a field has been set.

### GetUpdatedTimeSec

`func (o *WhatifScenario) GetUpdatedTimeSec() int32`

GetUpdatedTimeSec returns the UpdatedTimeSec field if non-nil, zero value otherwise.

### GetUpdatedTimeSecOk

`func (o *WhatifScenario) GetUpdatedTimeSecOk() (*int32, bool)`

GetUpdatedTimeSecOk returns a tuple with the UpdatedTimeSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimeSec

`func (o *WhatifScenario) SetUpdatedTimeSec(v int32)`

SetUpdatedTimeSec sets UpdatedTimeSec field to given value.

### HasUpdatedTimeSec

`func (o *WhatifScenario) HasUpdatedTimeSec() bool`

HasUpdatedTimeSec returns a boolean if a field has been set.

### GetClusterUuid

`func (o *WhatifScenario) GetClusterUuid() string`

GetClusterUuid returns the ClusterUuid field if non-nil, zero value otherwise.

### GetClusterUuidOk

`func (o *WhatifScenario) GetClusterUuidOk() (*string, bool)`

GetClusterUuidOk returns a tuple with the ClusterUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterUuid

`func (o *WhatifScenario) SetClusterUuid(v string)`

SetClusterUuid sets ClusterUuid field to given value.

### HasClusterUuid

`func (o *WhatifScenario) HasClusterUuid() bool`

HasClusterUuid returns a boolean if a field has been set.

### GetTargetRunwayDays

`func (o *WhatifScenario) GetTargetRunwayDays() int32`

GetTargetRunwayDays returns the TargetRunwayDays field if non-nil, zero value otherwise.

### GetTargetRunwayDaysOk

`func (o *WhatifScenario) GetTargetRunwayDaysOk() (*int32, bool)`

GetTargetRunwayDaysOk returns a tuple with the TargetRunwayDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetRunwayDays

`func (o *WhatifScenario) SetTargetRunwayDays(v int32)`

SetTargetRunwayDays sets TargetRunwayDays field to given value.

### HasTargetRunwayDays

`func (o *WhatifScenario) HasTargetRunwayDays() bool`

HasTargetRunwayDays returns a boolean if a field has been set.

### GetClusterSpec

`func (o *WhatifScenario) GetClusterSpec() ClusterSpec`

GetClusterSpec returns the ClusterSpec field if non-nil, zero value otherwise.

### GetClusterSpecOk

`func (o *WhatifScenario) GetClusterSpecOk() (*ClusterSpec, bool)`

GetClusterSpecOk returns a tuple with the ClusterSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterSpec

`func (o *WhatifScenario) SetClusterSpec(v ClusterSpec)`

SetClusterSpec sets ClusterSpec field to given value.

### HasClusterSpec

`func (o *WhatifScenario) HasClusterSpec() bool`

HasClusterSpec returns a boolean if a field has been set.

### GetRunway

`func (o *WhatifScenario) GetRunway() Runway`

GetRunway returns the Runway field if non-nil, zero value otherwise.

### GetRunwayOk

`func (o *WhatifScenario) GetRunwayOk() (*Runway, bool)`

GetRunwayOk returns a tuple with the Runway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunway

`func (o *WhatifScenario) SetRunway(v Runway)`

SetRunway sets Runway field to given value.

### HasRunway

`func (o *WhatifScenario) HasRunway() bool`

HasRunway returns a boolean if a field has been set.

### GetName

`func (o *WhatifScenario) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WhatifScenario) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WhatifScenario) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WhatifScenario) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



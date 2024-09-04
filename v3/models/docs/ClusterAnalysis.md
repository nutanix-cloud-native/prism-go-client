# ClusterAnalysis

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VmEfficiencyMap** | Pointer to **map[string]string** | Map of cluster efficiency which includes numbers of inefficient vms. The value is populated by analytics on PC.  | [optional] [readonly] 

## Methods

### NewClusterAnalysis

`func NewClusterAnalysis() *ClusterAnalysis`

NewClusterAnalysis instantiates a new ClusterAnalysis object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterAnalysisWithDefaults

`func NewClusterAnalysisWithDefaults() *ClusterAnalysis`

NewClusterAnalysisWithDefaults instantiates a new ClusterAnalysis object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVmEfficiencyMap

`func (o *ClusterAnalysis) GetVmEfficiencyMap() map[string]string`

GetVmEfficiencyMap returns the VmEfficiencyMap field if non-nil, zero value otherwise.

### GetVmEfficiencyMapOk

`func (o *ClusterAnalysis) GetVmEfficiencyMapOk() (*map[string]string, bool)`

GetVmEfficiencyMapOk returns a tuple with the VmEfficiencyMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmEfficiencyMap

`func (o *ClusterAnalysis) SetVmEfficiencyMap(v map[string]string)`

SetVmEfficiencyMap sets VmEfficiencyMap field to given value.

### HasVmEfficiencyMap

`func (o *ClusterAnalysis) HasVmEfficiencyMap() bool`

HasVmEfficiencyMap returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



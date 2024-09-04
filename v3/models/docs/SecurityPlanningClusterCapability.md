# SecurityPlanningClusterCapability

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterUuid** | Pointer to **string** | Cluster UUID of a cluster managed by Prism Central.  | [optional] 
**ClusterName** | Pointer to **string** | Name of the cluster. | [optional] 
**DataCollectorPrechecks** | Pointer to [**SecurityPlanningPrechecks**](SecurityPlanningPrechecks.md) |  | [optional] 
**TrafficVisibilityPrechecks** | Pointer to [**TrafficVisibilityPrecheckResults**](TrafficVisibilityPrecheckResults.md) |  | [optional] 

## Methods

### NewSecurityPlanningClusterCapability

`func NewSecurityPlanningClusterCapability() *SecurityPlanningClusterCapability`

NewSecurityPlanningClusterCapability instantiates a new SecurityPlanningClusterCapability object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityPlanningClusterCapabilityWithDefaults

`func NewSecurityPlanningClusterCapabilityWithDefaults() *SecurityPlanningClusterCapability`

NewSecurityPlanningClusterCapabilityWithDefaults instantiates a new SecurityPlanningClusterCapability object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterUuid

`func (o *SecurityPlanningClusterCapability) GetClusterUuid() string`

GetClusterUuid returns the ClusterUuid field if non-nil, zero value otherwise.

### GetClusterUuidOk

`func (o *SecurityPlanningClusterCapability) GetClusterUuidOk() (*string, bool)`

GetClusterUuidOk returns a tuple with the ClusterUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterUuid

`func (o *SecurityPlanningClusterCapability) SetClusterUuid(v string)`

SetClusterUuid sets ClusterUuid field to given value.

### HasClusterUuid

`func (o *SecurityPlanningClusterCapability) HasClusterUuid() bool`

HasClusterUuid returns a boolean if a field has been set.

### GetClusterName

`func (o *SecurityPlanningClusterCapability) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *SecurityPlanningClusterCapability) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *SecurityPlanningClusterCapability) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.

### HasClusterName

`func (o *SecurityPlanningClusterCapability) HasClusterName() bool`

HasClusterName returns a boolean if a field has been set.

### GetDataCollectorPrechecks

`func (o *SecurityPlanningClusterCapability) GetDataCollectorPrechecks() SecurityPlanningPrechecks`

GetDataCollectorPrechecks returns the DataCollectorPrechecks field if non-nil, zero value otherwise.

### GetDataCollectorPrechecksOk

`func (o *SecurityPlanningClusterCapability) GetDataCollectorPrechecksOk() (*SecurityPlanningPrechecks, bool)`

GetDataCollectorPrechecksOk returns a tuple with the DataCollectorPrechecks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataCollectorPrechecks

`func (o *SecurityPlanningClusterCapability) SetDataCollectorPrechecks(v SecurityPlanningPrechecks)`

SetDataCollectorPrechecks sets DataCollectorPrechecks field to given value.

### HasDataCollectorPrechecks

`func (o *SecurityPlanningClusterCapability) HasDataCollectorPrechecks() bool`

HasDataCollectorPrechecks returns a boolean if a field has been set.

### GetTrafficVisibilityPrechecks

`func (o *SecurityPlanningClusterCapability) GetTrafficVisibilityPrechecks() TrafficVisibilityPrecheckResults`

GetTrafficVisibilityPrechecks returns the TrafficVisibilityPrechecks field if non-nil, zero value otherwise.

### GetTrafficVisibilityPrechecksOk

`func (o *SecurityPlanningClusterCapability) GetTrafficVisibilityPrechecksOk() (*TrafficVisibilityPrecheckResults, bool)`

GetTrafficVisibilityPrechecksOk returns a tuple with the TrafficVisibilityPrechecks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficVisibilityPrechecks

`func (o *SecurityPlanningClusterCapability) SetTrafficVisibilityPrechecks(v TrafficVisibilityPrecheckResults)`

SetTrafficVisibilityPrechecks sets TrafficVisibilityPrechecks field to given value.

### HasTrafficVisibilityPrechecks

`func (o *SecurityPlanningClusterCapability) HasTrafficVisibilityPrechecks() bool`

HasTrafficVisibilityPrechecks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



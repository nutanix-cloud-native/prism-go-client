# SecurityPlanningClusterCapabilityStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocalClusterPrechecks** | Pointer to [**SecurityPlanningPrechecks**](SecurityPlanningPrechecks.md) |  | [optional] 
**ClusterCapabilityList** | Pointer to [**[]SecurityPlanningClusterCapability**](SecurityPlanningClusterCapability.md) | Capability of the feature per cluster managed by Prism Central.  | [optional] 

## Methods

### NewSecurityPlanningClusterCapabilityStatus

`func NewSecurityPlanningClusterCapabilityStatus() *SecurityPlanningClusterCapabilityStatus`

NewSecurityPlanningClusterCapabilityStatus instantiates a new SecurityPlanningClusterCapabilityStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityPlanningClusterCapabilityStatusWithDefaults

`func NewSecurityPlanningClusterCapabilityStatusWithDefaults() *SecurityPlanningClusterCapabilityStatus`

NewSecurityPlanningClusterCapabilityStatusWithDefaults instantiates a new SecurityPlanningClusterCapabilityStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocalClusterPrechecks

`func (o *SecurityPlanningClusterCapabilityStatus) GetLocalClusterPrechecks() SecurityPlanningPrechecks`

GetLocalClusterPrechecks returns the LocalClusterPrechecks field if non-nil, zero value otherwise.

### GetLocalClusterPrechecksOk

`func (o *SecurityPlanningClusterCapabilityStatus) GetLocalClusterPrechecksOk() (*SecurityPlanningPrechecks, bool)`

GetLocalClusterPrechecksOk returns a tuple with the LocalClusterPrechecks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalClusterPrechecks

`func (o *SecurityPlanningClusterCapabilityStatus) SetLocalClusterPrechecks(v SecurityPlanningPrechecks)`

SetLocalClusterPrechecks sets LocalClusterPrechecks field to given value.

### HasLocalClusterPrechecks

`func (o *SecurityPlanningClusterCapabilityStatus) HasLocalClusterPrechecks() bool`

HasLocalClusterPrechecks returns a boolean if a field has been set.

### GetClusterCapabilityList

`func (o *SecurityPlanningClusterCapabilityStatus) GetClusterCapabilityList() []SecurityPlanningClusterCapability`

GetClusterCapabilityList returns the ClusterCapabilityList field if non-nil, zero value otherwise.

### GetClusterCapabilityListOk

`func (o *SecurityPlanningClusterCapabilityStatus) GetClusterCapabilityListOk() (*[]SecurityPlanningClusterCapability, bool)`

GetClusterCapabilityListOk returns a tuple with the ClusterCapabilityList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterCapabilityList

`func (o *SecurityPlanningClusterCapabilityStatus) SetClusterCapabilityList(v []SecurityPlanningClusterCapability)`

SetClusterCapabilityList sets ClusterCapabilityList field to given value.

### HasClusterCapabilityList

`func (o *SecurityPlanningClusterCapabilityStatus) HasClusterCapabilityList() bool`

HasClusterCapabilityList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



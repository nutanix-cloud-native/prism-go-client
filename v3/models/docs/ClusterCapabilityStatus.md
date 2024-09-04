# ClusterCapabilityStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterCapabilityList** | Pointer to [**[]ClusterCapability**](ClusterCapability.md) | Capability of the feature per cluster managed by Prism Central.  | [optional] 

## Methods

### NewClusterCapabilityStatus

`func NewClusterCapabilityStatus() *ClusterCapabilityStatus`

NewClusterCapabilityStatus instantiates a new ClusterCapabilityStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterCapabilityStatusWithDefaults

`func NewClusterCapabilityStatusWithDefaults() *ClusterCapabilityStatus`

NewClusterCapabilityStatusWithDefaults instantiates a new ClusterCapabilityStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterCapabilityList

`func (o *ClusterCapabilityStatus) GetClusterCapabilityList() []ClusterCapability`

GetClusterCapabilityList returns the ClusterCapabilityList field if non-nil, zero value otherwise.

### GetClusterCapabilityListOk

`func (o *ClusterCapabilityStatus) GetClusterCapabilityListOk() (*[]ClusterCapability, bool)`

GetClusterCapabilityListOk returns a tuple with the ClusterCapabilityList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterCapabilityList

`func (o *ClusterCapabilityStatus) SetClusterCapabilityList(v []ClusterCapability)`

SetClusterCapabilityList sets ClusterCapabilityList field to given value.

### HasClusterCapabilityList

`func (o *ClusterCapabilityStatus) HasClusterCapabilityList() bool`

HasClusterCapabilityList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# QosConfigStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IngressLimit** | Pointer to **int32** | Ingress traffic limit. (In Mbits/s) | [optional] 
**EgressLimit** | Pointer to **int32** | Egress traffic limit. (In Mbits/s) | [optional] 

## Methods

### NewQosConfigStatus

`func NewQosConfigStatus() *QosConfigStatus`

NewQosConfigStatus instantiates a new QosConfigStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQosConfigStatusWithDefaults

`func NewQosConfigStatusWithDefaults() *QosConfigStatus`

NewQosConfigStatusWithDefaults instantiates a new QosConfigStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIngressLimit

`func (o *QosConfigStatus) GetIngressLimit() int32`

GetIngressLimit returns the IngressLimit field if non-nil, zero value otherwise.

### GetIngressLimitOk

`func (o *QosConfigStatus) GetIngressLimitOk() (*int32, bool)`

GetIngressLimitOk returns a tuple with the IngressLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngressLimit

`func (o *QosConfigStatus) SetIngressLimit(v int32)`

SetIngressLimit sets IngressLimit field to given value.

### HasIngressLimit

`func (o *QosConfigStatus) HasIngressLimit() bool`

HasIngressLimit returns a boolean if a field has been set.

### GetEgressLimit

`func (o *QosConfigStatus) GetEgressLimit() int32`

GetEgressLimit returns the EgressLimit field if non-nil, zero value otherwise.

### GetEgressLimitOk

`func (o *QosConfigStatus) GetEgressLimitOk() (*int32, bool)`

GetEgressLimitOk returns a tuple with the EgressLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEgressLimit

`func (o *QosConfigStatus) SetEgressLimit(v int32)`

SetEgressLimit sets EgressLimit field to given value.

### HasEgressLimit

`func (o *QosConfigStatus) HasEgressLimit() bool`

HasEgressLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



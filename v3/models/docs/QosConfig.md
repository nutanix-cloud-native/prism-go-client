# QosConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IngressLimit** | Pointer to **int32** | Ingress traffic limit. (In Mbits/s) | [optional] 
**EgressLimit** | Pointer to **int32** | Egress traffic limit. (In Mbits/s) | [optional] 

## Methods

### NewQosConfig

`func NewQosConfig() *QosConfig`

NewQosConfig instantiates a new QosConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQosConfigWithDefaults

`func NewQosConfigWithDefaults() *QosConfig`

NewQosConfigWithDefaults instantiates a new QosConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIngressLimit

`func (o *QosConfig) GetIngressLimit() int32`

GetIngressLimit returns the IngressLimit field if non-nil, zero value otherwise.

### GetIngressLimitOk

`func (o *QosConfig) GetIngressLimitOk() (*int32, bool)`

GetIngressLimitOk returns a tuple with the IngressLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngressLimit

`func (o *QosConfig) SetIngressLimit(v int32)`

SetIngressLimit sets IngressLimit field to given value.

### HasIngressLimit

`func (o *QosConfig) HasIngressLimit() bool`

HasIngressLimit returns a boolean if a field has been set.

### GetEgressLimit

`func (o *QosConfig) GetEgressLimit() int32`

GetEgressLimit returns the EgressLimit field if non-nil, zero value otherwise.

### GetEgressLimitOk

`func (o *QosConfig) GetEgressLimitOk() (*int32, bool)`

GetEgressLimitOk returns a tuple with the EgressLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEgressLimit

`func (o *QosConfig) SetEgressLimit(v int32)`

SetEgressLimit sets EgressLimit field to given value.

### HasEgressLimit

`func (o *QosConfig) HasEgressLimit() bool`

HasEgressLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



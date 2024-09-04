# ServiceIpCounters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Received** | Pointer to [**RoutingPolicyCounters**](RoutingPolicyCounters.md) |  | [optional] 
**ServiceIp** | Pointer to **string** |  | [optional] 
**Sent** | Pointer to [**RoutingPolicyCounters**](RoutingPolicyCounters.md) |  | [optional] 

## Methods

### NewServiceIpCounters

`func NewServiceIpCounters() *ServiceIpCounters`

NewServiceIpCounters instantiates a new ServiceIpCounters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceIpCountersWithDefaults

`func NewServiceIpCountersWithDefaults() *ServiceIpCounters`

NewServiceIpCountersWithDefaults instantiates a new ServiceIpCounters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReceived

`func (o *ServiceIpCounters) GetReceived() RoutingPolicyCounters`

GetReceived returns the Received field if non-nil, zero value otherwise.

### GetReceivedOk

`func (o *ServiceIpCounters) GetReceivedOk() (*RoutingPolicyCounters, bool)`

GetReceivedOk returns a tuple with the Received field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceived

`func (o *ServiceIpCounters) SetReceived(v RoutingPolicyCounters)`

SetReceived sets Received field to given value.

### HasReceived

`func (o *ServiceIpCounters) HasReceived() bool`

HasReceived returns a boolean if a field has been set.

### GetServiceIp

`func (o *ServiceIpCounters) GetServiceIp() string`

GetServiceIp returns the ServiceIp field if non-nil, zero value otherwise.

### GetServiceIpOk

`func (o *ServiceIpCounters) GetServiceIpOk() (*string, bool)`

GetServiceIpOk returns a tuple with the ServiceIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceIp

`func (o *ServiceIpCounters) SetServiceIp(v string)`

SetServiceIp sets ServiceIp field to given value.

### HasServiceIp

`func (o *ServiceIpCounters) HasServiceIp() bool`

HasServiceIp returns a boolean if a field has been set.

### GetSent

`func (o *ServiceIpCounters) GetSent() RoutingPolicyCounters`

GetSent returns the Sent field if non-nil, zero value otherwise.

### GetSentOk

`func (o *ServiceIpCounters) GetSentOk() (*RoutingPolicyCounters, bool)`

GetSentOk returns a tuple with the Sent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSent

`func (o *ServiceIpCounters) SetSent(v RoutingPolicyCounters)`

SetSent sets Sent field to given value.

### HasSent

`func (o *ServiceIpCounters) HasSent() bool`

HasSent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



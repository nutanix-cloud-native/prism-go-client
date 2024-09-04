# RoutingPolicyAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** |  | 
**ServiceIpList** | Pointer to **[]string** | IP addresses of network services in the chain.&gt; This field is valid only when action is REROUTE. | [optional] 
**RerouteFallback** | Pointer to **string** | Flag denoting reroute fallback action when service VM is down. This field is valid only when action is REROUTE. Default reroute fallback action is PASSTHROUGH.  | [optional] 

## Methods

### NewRoutingPolicyAction

`func NewRoutingPolicyAction(action string, ) *RoutingPolicyAction`

NewRoutingPolicyAction instantiates a new RoutingPolicyAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoutingPolicyActionWithDefaults

`func NewRoutingPolicyActionWithDefaults() *RoutingPolicyAction`

NewRoutingPolicyActionWithDefaults instantiates a new RoutingPolicyAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *RoutingPolicyAction) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *RoutingPolicyAction) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *RoutingPolicyAction) SetAction(v string)`

SetAction sets Action field to given value.


### GetServiceIpList

`func (o *RoutingPolicyAction) GetServiceIpList() []string`

GetServiceIpList returns the ServiceIpList field if non-nil, zero value otherwise.

### GetServiceIpListOk

`func (o *RoutingPolicyAction) GetServiceIpListOk() (*[]string, bool)`

GetServiceIpListOk returns a tuple with the ServiceIpList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceIpList

`func (o *RoutingPolicyAction) SetServiceIpList(v []string)`

SetServiceIpList sets ServiceIpList field to given value.

### HasServiceIpList

`func (o *RoutingPolicyAction) HasServiceIpList() bool`

HasServiceIpList returns a boolean if a field has been set.

### GetRerouteFallback

`func (o *RoutingPolicyAction) GetRerouteFallback() string`

GetRerouteFallback returns the RerouteFallback field if non-nil, zero value otherwise.

### GetRerouteFallbackOk

`func (o *RoutingPolicyAction) GetRerouteFallbackOk() (*string, bool)`

GetRerouteFallbackOk returns a tuple with the RerouteFallback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRerouteFallback

`func (o *RoutingPolicyAction) SetRerouteFallback(v string)`

SetRerouteFallback sets RerouteFallback field to given value.

### HasRerouteFallback

`func (o *RoutingPolicyAction) HasRerouteFallback() bool`

HasRerouteFallback returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



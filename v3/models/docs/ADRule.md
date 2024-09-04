# ADRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | Type of deployment of the rule. | [optional] 
**OutboundAllowList** | Pointer to [**[]NetworkRuleStatus**](NetworkRuleStatus.md) |  | [optional] 
**TargetGroup** | Pointer to [**TargetGroup**](TargetGroup.md) |  | [optional] 
**InboundAllowList** | Pointer to [**[]NetworkRuleStatus**](NetworkRuleStatus.md) |  | [optional] 

## Methods

### NewADRule

`func NewADRule() *ADRule`

NewADRule instantiates a new ADRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewADRuleWithDefaults

`func NewADRuleWithDefaults() *ADRule`

NewADRuleWithDefaults instantiates a new ADRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *ADRule) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ADRule) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ADRule) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *ADRule) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetOutboundAllowList

`func (o *ADRule) GetOutboundAllowList() []NetworkRuleStatus`

GetOutboundAllowList returns the OutboundAllowList field if non-nil, zero value otherwise.

### GetOutboundAllowListOk

`func (o *ADRule) GetOutboundAllowListOk() (*[]NetworkRuleStatus, bool)`

GetOutboundAllowListOk returns a tuple with the OutboundAllowList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutboundAllowList

`func (o *ADRule) SetOutboundAllowList(v []NetworkRuleStatus)`

SetOutboundAllowList sets OutboundAllowList field to given value.

### HasOutboundAllowList

`func (o *ADRule) HasOutboundAllowList() bool`

HasOutboundAllowList returns a boolean if a field has been set.

### GetTargetGroup

`func (o *ADRule) GetTargetGroup() TargetGroup`

GetTargetGroup returns the TargetGroup field if non-nil, zero value otherwise.

### GetTargetGroupOk

`func (o *ADRule) GetTargetGroupOk() (*TargetGroup, bool)`

GetTargetGroupOk returns a tuple with the TargetGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetGroup

`func (o *ADRule) SetTargetGroup(v TargetGroup)`

SetTargetGroup sets TargetGroup field to given value.

### HasTargetGroup

`func (o *ADRule) HasTargetGroup() bool`

HasTargetGroup returns a boolean if a field has been set.

### GetInboundAllowList

`func (o *ADRule) GetInboundAllowList() []NetworkRuleStatus`

GetInboundAllowList returns the InboundAllowList field if non-nil, zero value otherwise.

### GetInboundAllowListOk

`func (o *ADRule) GetInboundAllowListOk() (*[]NetworkRuleStatus, bool)`

GetInboundAllowListOk returns a tuple with the InboundAllowList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInboundAllowList

`func (o *ADRule) SetInboundAllowList(v []NetworkRuleStatus)`

SetInboundAllowList sets InboundAllowList field to given value.

### HasInboundAllowList

`func (o *ADRule) HasInboundAllowList() bool`

HasInboundAllowList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



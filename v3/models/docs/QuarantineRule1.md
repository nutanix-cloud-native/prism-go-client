# QuarantineRule1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | Type of action. | [optional] 
**OutboundAllowList** | Pointer to [**[]NetworkRule**](NetworkRule.md) |  | [optional] 
**TargetGroup** | Pointer to [**TargetGroup**](TargetGroup.md) |  | [optional] 
**InboundAllowList** | Pointer to [**[]NetworkRule**](NetworkRule.md) |  | [optional] 

## Methods

### NewQuarantineRule1

`func NewQuarantineRule1() *QuarantineRule1`

NewQuarantineRule1 instantiates a new QuarantineRule1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuarantineRule1WithDefaults

`func NewQuarantineRule1WithDefaults() *QuarantineRule1`

NewQuarantineRule1WithDefaults instantiates a new QuarantineRule1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *QuarantineRule1) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *QuarantineRule1) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *QuarantineRule1) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *QuarantineRule1) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetOutboundAllowList

`func (o *QuarantineRule1) GetOutboundAllowList() []NetworkRule`

GetOutboundAllowList returns the OutboundAllowList field if non-nil, zero value otherwise.

### GetOutboundAllowListOk

`func (o *QuarantineRule1) GetOutboundAllowListOk() (*[]NetworkRule, bool)`

GetOutboundAllowListOk returns a tuple with the OutboundAllowList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutboundAllowList

`func (o *QuarantineRule1) SetOutboundAllowList(v []NetworkRule)`

SetOutboundAllowList sets OutboundAllowList field to given value.

### HasOutboundAllowList

`func (o *QuarantineRule1) HasOutboundAllowList() bool`

HasOutboundAllowList returns a boolean if a field has been set.

### GetTargetGroup

`func (o *QuarantineRule1) GetTargetGroup() TargetGroup`

GetTargetGroup returns the TargetGroup field if non-nil, zero value otherwise.

### GetTargetGroupOk

`func (o *QuarantineRule1) GetTargetGroupOk() (*TargetGroup, bool)`

GetTargetGroupOk returns a tuple with the TargetGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetGroup

`func (o *QuarantineRule1) SetTargetGroup(v TargetGroup)`

SetTargetGroup sets TargetGroup field to given value.

### HasTargetGroup

`func (o *QuarantineRule1) HasTargetGroup() bool`

HasTargetGroup returns a boolean if a field has been set.

### GetInboundAllowList

`func (o *QuarantineRule1) GetInboundAllowList() []NetworkRule`

GetInboundAllowList returns the InboundAllowList field if non-nil, zero value otherwise.

### GetInboundAllowListOk

`func (o *QuarantineRule1) GetInboundAllowListOk() (*[]NetworkRule, bool)`

GetInboundAllowListOk returns a tuple with the InboundAllowList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInboundAllowList

`func (o *QuarantineRule1) SetInboundAllowList(v []NetworkRule)`

SetInboundAllowList sets InboundAllowList field to given value.

### HasInboundAllowList

`func (o *QuarantineRule1) HasInboundAllowList() bool`

HasInboundAllowList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# NetworkSecurityRuleResourcesStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowIpv6Traffic** | Pointer to **bool** |  | [optional] 
**IsPolicyHitlogEnabled** | Pointer to **bool** |  | [optional] 
**QuarantineRule** | Pointer to [**QuarantineRule**](QuarantineRule.md) |  | [optional] 
**IsolationRule** | Pointer to [**IsolationRule**](IsolationRule.md) |  | [optional] 
**AppRule** | Pointer to [**AppRule**](AppRule.md) |  | [optional] 
**AdRule** | Pointer to [**ADRule**](ADRule.md) |  | [optional] 

## Methods

### NewNetworkSecurityRuleResourcesStatus

`func NewNetworkSecurityRuleResourcesStatus() *NetworkSecurityRuleResourcesStatus`

NewNetworkSecurityRuleResourcesStatus instantiates a new NetworkSecurityRuleResourcesStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkSecurityRuleResourcesStatusWithDefaults

`func NewNetworkSecurityRuleResourcesStatusWithDefaults() *NetworkSecurityRuleResourcesStatus`

NewNetworkSecurityRuleResourcesStatusWithDefaults instantiates a new NetworkSecurityRuleResourcesStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowIpv6Traffic

`func (o *NetworkSecurityRuleResourcesStatus) GetAllowIpv6Traffic() bool`

GetAllowIpv6Traffic returns the AllowIpv6Traffic field if non-nil, zero value otherwise.

### GetAllowIpv6TrafficOk

`func (o *NetworkSecurityRuleResourcesStatus) GetAllowIpv6TrafficOk() (*bool, bool)`

GetAllowIpv6TrafficOk returns a tuple with the AllowIpv6Traffic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowIpv6Traffic

`func (o *NetworkSecurityRuleResourcesStatus) SetAllowIpv6Traffic(v bool)`

SetAllowIpv6Traffic sets AllowIpv6Traffic field to given value.

### HasAllowIpv6Traffic

`func (o *NetworkSecurityRuleResourcesStatus) HasAllowIpv6Traffic() bool`

HasAllowIpv6Traffic returns a boolean if a field has been set.

### GetIsPolicyHitlogEnabled

`func (o *NetworkSecurityRuleResourcesStatus) GetIsPolicyHitlogEnabled() bool`

GetIsPolicyHitlogEnabled returns the IsPolicyHitlogEnabled field if non-nil, zero value otherwise.

### GetIsPolicyHitlogEnabledOk

`func (o *NetworkSecurityRuleResourcesStatus) GetIsPolicyHitlogEnabledOk() (*bool, bool)`

GetIsPolicyHitlogEnabledOk returns a tuple with the IsPolicyHitlogEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPolicyHitlogEnabled

`func (o *NetworkSecurityRuleResourcesStatus) SetIsPolicyHitlogEnabled(v bool)`

SetIsPolicyHitlogEnabled sets IsPolicyHitlogEnabled field to given value.

### HasIsPolicyHitlogEnabled

`func (o *NetworkSecurityRuleResourcesStatus) HasIsPolicyHitlogEnabled() bool`

HasIsPolicyHitlogEnabled returns a boolean if a field has been set.

### GetQuarantineRule

`func (o *NetworkSecurityRuleResourcesStatus) GetQuarantineRule() QuarantineRule`

GetQuarantineRule returns the QuarantineRule field if non-nil, zero value otherwise.

### GetQuarantineRuleOk

`func (o *NetworkSecurityRuleResourcesStatus) GetQuarantineRuleOk() (*QuarantineRule, bool)`

GetQuarantineRuleOk returns a tuple with the QuarantineRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuarantineRule

`func (o *NetworkSecurityRuleResourcesStatus) SetQuarantineRule(v QuarantineRule)`

SetQuarantineRule sets QuarantineRule field to given value.

### HasQuarantineRule

`func (o *NetworkSecurityRuleResourcesStatus) HasQuarantineRule() bool`

HasQuarantineRule returns a boolean if a field has been set.

### GetIsolationRule

`func (o *NetworkSecurityRuleResourcesStatus) GetIsolationRule() IsolationRule`

GetIsolationRule returns the IsolationRule field if non-nil, zero value otherwise.

### GetIsolationRuleOk

`func (o *NetworkSecurityRuleResourcesStatus) GetIsolationRuleOk() (*IsolationRule, bool)`

GetIsolationRuleOk returns a tuple with the IsolationRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsolationRule

`func (o *NetworkSecurityRuleResourcesStatus) SetIsolationRule(v IsolationRule)`

SetIsolationRule sets IsolationRule field to given value.

### HasIsolationRule

`func (o *NetworkSecurityRuleResourcesStatus) HasIsolationRule() bool`

HasIsolationRule returns a boolean if a field has been set.

### GetAppRule

`func (o *NetworkSecurityRuleResourcesStatus) GetAppRule() AppRule`

GetAppRule returns the AppRule field if non-nil, zero value otherwise.

### GetAppRuleOk

`func (o *NetworkSecurityRuleResourcesStatus) GetAppRuleOk() (*AppRule, bool)`

GetAppRuleOk returns a tuple with the AppRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppRule

`func (o *NetworkSecurityRuleResourcesStatus) SetAppRule(v AppRule)`

SetAppRule sets AppRule field to given value.

### HasAppRule

`func (o *NetworkSecurityRuleResourcesStatus) HasAppRule() bool`

HasAppRule returns a boolean if a field has been set.

### GetAdRule

`func (o *NetworkSecurityRuleResourcesStatus) GetAdRule() ADRule`

GetAdRule returns the AdRule field if non-nil, zero value otherwise.

### GetAdRuleOk

`func (o *NetworkSecurityRuleResourcesStatus) GetAdRuleOk() (*ADRule, bool)`

GetAdRuleOk returns a tuple with the AdRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdRule

`func (o *NetworkSecurityRuleResourcesStatus) SetAdRule(v ADRule)`

SetAdRule sets AdRule field to given value.

### HasAdRule

`func (o *NetworkSecurityRuleResourcesStatus) HasAdRule() bool`

HasAdRule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



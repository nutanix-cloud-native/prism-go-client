# NetworkSecurityRuleResources

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowIpv6Traffic** | Pointer to **bool** |  | [optional] 
**IsPolicyHitlogEnabled** | Pointer to **bool** |  | [optional] 
**QuarantineRule** | Pointer to [**QuarantineRule1**](QuarantineRule1.md) |  | [optional] 
**IsolationRule** | Pointer to [**IsolationRule**](IsolationRule.md) |  | [optional] 
**AppRule** | Pointer to [**AppRule1**](AppRule1.md) |  | [optional] 
**AdRule** | Pointer to [**ADRule1**](ADRule1.md) |  | [optional] 

## Methods

### NewNetworkSecurityRuleResources

`func NewNetworkSecurityRuleResources() *NetworkSecurityRuleResources`

NewNetworkSecurityRuleResources instantiates a new NetworkSecurityRuleResources object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkSecurityRuleResourcesWithDefaults

`func NewNetworkSecurityRuleResourcesWithDefaults() *NetworkSecurityRuleResources`

NewNetworkSecurityRuleResourcesWithDefaults instantiates a new NetworkSecurityRuleResources object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowIpv6Traffic

`func (o *NetworkSecurityRuleResources) GetAllowIpv6Traffic() bool`

GetAllowIpv6Traffic returns the AllowIpv6Traffic field if non-nil, zero value otherwise.

### GetAllowIpv6TrafficOk

`func (o *NetworkSecurityRuleResources) GetAllowIpv6TrafficOk() (*bool, bool)`

GetAllowIpv6TrafficOk returns a tuple with the AllowIpv6Traffic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowIpv6Traffic

`func (o *NetworkSecurityRuleResources) SetAllowIpv6Traffic(v bool)`

SetAllowIpv6Traffic sets AllowIpv6Traffic field to given value.

### HasAllowIpv6Traffic

`func (o *NetworkSecurityRuleResources) HasAllowIpv6Traffic() bool`

HasAllowIpv6Traffic returns a boolean if a field has been set.

### GetIsPolicyHitlogEnabled

`func (o *NetworkSecurityRuleResources) GetIsPolicyHitlogEnabled() bool`

GetIsPolicyHitlogEnabled returns the IsPolicyHitlogEnabled field if non-nil, zero value otherwise.

### GetIsPolicyHitlogEnabledOk

`func (o *NetworkSecurityRuleResources) GetIsPolicyHitlogEnabledOk() (*bool, bool)`

GetIsPolicyHitlogEnabledOk returns a tuple with the IsPolicyHitlogEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPolicyHitlogEnabled

`func (o *NetworkSecurityRuleResources) SetIsPolicyHitlogEnabled(v bool)`

SetIsPolicyHitlogEnabled sets IsPolicyHitlogEnabled field to given value.

### HasIsPolicyHitlogEnabled

`func (o *NetworkSecurityRuleResources) HasIsPolicyHitlogEnabled() bool`

HasIsPolicyHitlogEnabled returns a boolean if a field has been set.

### GetQuarantineRule

`func (o *NetworkSecurityRuleResources) GetQuarantineRule() QuarantineRule1`

GetQuarantineRule returns the QuarantineRule field if non-nil, zero value otherwise.

### GetQuarantineRuleOk

`func (o *NetworkSecurityRuleResources) GetQuarantineRuleOk() (*QuarantineRule1, bool)`

GetQuarantineRuleOk returns a tuple with the QuarantineRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuarantineRule

`func (o *NetworkSecurityRuleResources) SetQuarantineRule(v QuarantineRule1)`

SetQuarantineRule sets QuarantineRule field to given value.

### HasQuarantineRule

`func (o *NetworkSecurityRuleResources) HasQuarantineRule() bool`

HasQuarantineRule returns a boolean if a field has been set.

### GetIsolationRule

`func (o *NetworkSecurityRuleResources) GetIsolationRule() IsolationRule`

GetIsolationRule returns the IsolationRule field if non-nil, zero value otherwise.

### GetIsolationRuleOk

`func (o *NetworkSecurityRuleResources) GetIsolationRuleOk() (*IsolationRule, bool)`

GetIsolationRuleOk returns a tuple with the IsolationRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsolationRule

`func (o *NetworkSecurityRuleResources) SetIsolationRule(v IsolationRule)`

SetIsolationRule sets IsolationRule field to given value.

### HasIsolationRule

`func (o *NetworkSecurityRuleResources) HasIsolationRule() bool`

HasIsolationRule returns a boolean if a field has been set.

### GetAppRule

`func (o *NetworkSecurityRuleResources) GetAppRule() AppRule1`

GetAppRule returns the AppRule field if non-nil, zero value otherwise.

### GetAppRuleOk

`func (o *NetworkSecurityRuleResources) GetAppRuleOk() (*AppRule1, bool)`

GetAppRuleOk returns a tuple with the AppRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppRule

`func (o *NetworkSecurityRuleResources) SetAppRule(v AppRule1)`

SetAppRule sets AppRule field to given value.

### HasAppRule

`func (o *NetworkSecurityRuleResources) HasAppRule() bool`

HasAppRule returns a boolean if a field has been set.

### GetAdRule

`func (o *NetworkSecurityRuleResources) GetAdRule() ADRule1`

GetAdRule returns the AdRule field if non-nil, zero value otherwise.

### GetAdRuleOk

`func (o *NetworkSecurityRuleResources) GetAdRuleOk() (*ADRule1, bool)`

GetAdRuleOk returns a tuple with the AdRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdRule

`func (o *NetworkSecurityRuleResources) SetAdRule(v ADRule1)`

SetAdRule sets AdRule field to given value.

### HasAdRule

`func (o *NetworkSecurityRuleResources) HasAdRule() bool`

HasAdRule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



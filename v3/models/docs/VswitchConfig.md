# VswitchConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NicTeamingPolicy** | **string** | NIC teaming policy. | 
**UplinkGrouping** | **string** | Determines how the ethernet uplinks are selected for this vswitch. | 

## Methods

### NewVswitchConfig

`func NewVswitchConfig(nicTeamingPolicy string, uplinkGrouping string, ) *VswitchConfig`

NewVswitchConfig instantiates a new VswitchConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVswitchConfigWithDefaults

`func NewVswitchConfigWithDefaults() *VswitchConfig`

NewVswitchConfigWithDefaults instantiates a new VswitchConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNicTeamingPolicy

`func (o *VswitchConfig) GetNicTeamingPolicy() string`

GetNicTeamingPolicy returns the NicTeamingPolicy field if non-nil, zero value otherwise.

### GetNicTeamingPolicyOk

`func (o *VswitchConfig) GetNicTeamingPolicyOk() (*string, bool)`

GetNicTeamingPolicyOk returns a tuple with the NicTeamingPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNicTeamingPolicy

`func (o *VswitchConfig) SetNicTeamingPolicy(v string)`

SetNicTeamingPolicy sets NicTeamingPolicy field to given value.


### GetUplinkGrouping

`func (o *VswitchConfig) GetUplinkGrouping() string`

GetUplinkGrouping returns the UplinkGrouping field if non-nil, zero value otherwise.

### GetUplinkGroupingOk

`func (o *VswitchConfig) GetUplinkGroupingOk() (*string, bool)`

GetUplinkGroupingOk returns a tuple with the UplinkGrouping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplinkGrouping

`func (o *VswitchConfig) SetUplinkGrouping(v string)`

SetUplinkGrouping sets UplinkGrouping field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



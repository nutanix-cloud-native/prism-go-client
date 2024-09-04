# VpnComponentStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | Pointer to **string** |  | [optional] 
**Detail** | Pointer to **string** | Detailed informational/error string in human-readable form. | [optional] 

## Methods

### NewVpnComponentStatus

`func NewVpnComponentStatus() *VpnComponentStatus`

NewVpnComponentStatus instantiates a new VpnComponentStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpnComponentStatusWithDefaults

`func NewVpnComponentStatusWithDefaults() *VpnComponentStatus`

NewVpnComponentStatusWithDefaults instantiates a new VpnComponentStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *VpnComponentStatus) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *VpnComponentStatus) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *VpnComponentStatus) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *VpnComponentStatus) HasState() bool`

HasState returns a boolean if a field has been set.

### GetDetail

`func (o *VpnComponentStatus) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *VpnComponentStatus) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *VpnComponentStatus) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *VpnComponentStatus) HasDetail() bool`

HasDetail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



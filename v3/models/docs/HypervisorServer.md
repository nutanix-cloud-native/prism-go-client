# HypervisorServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ip** | **string** |  | 
**Version** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewHypervisorServer

`func NewHypervisorServer(ip string, ) *HypervisorServer`

NewHypervisorServer instantiates a new HypervisorServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHypervisorServerWithDefaults

`func NewHypervisorServerWithDefaults() *HypervisorServer`

NewHypervisorServerWithDefaults instantiates a new HypervisorServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIp

`func (o *HypervisorServer) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *HypervisorServer) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *HypervisorServer) SetIp(v string)`

SetIp sets Ip field to given value.


### GetVersion

`func (o *HypervisorServer) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *HypervisorServer) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *HypervisorServer) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *HypervisorServer) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetType

`func (o *HypervisorServer) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HypervisorServer) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HypervisorServer) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *HypervisorServer) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



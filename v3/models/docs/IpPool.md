# IpPool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Range** | Pointer to **string** | Range of IPs (example: 10.0.0.9 10.0.0.19).  | [optional] 

## Methods

### NewIpPool

`func NewIpPool() *IpPool`

NewIpPool instantiates a new IpPool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpPoolWithDefaults

`func NewIpPoolWithDefaults() *IpPool`

NewIpPoolWithDefaults instantiates a new IpPool object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRange

`func (o *IpPool) GetRange() string`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *IpPool) GetRangeOk() (*string, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *IpPool) SetRange(v string)`

SetRange sets Range field to given value.

### HasRange

`func (o *IpPool) HasRange() bool`

HasRange returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



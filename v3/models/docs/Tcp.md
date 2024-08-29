# Tcp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestinationPortRangeList** | Pointer to [**[]PortRange**](PortRange.md) | List of ranges of TCP destination ports. | [optional] 
**SourcePortRange** | Pointer to [**PortRange**](PortRange.md) |  | [optional] 
**SourcePortRangeList** | Pointer to [**[]PortRange**](PortRange.md) | List of ranges of TCP source ports. | [optional] 
**DestinationPortRange** | Pointer to [**PortRange**](PortRange.md) |  | [optional] 

## Methods

### NewTcp

`func NewTcp() *Tcp`

NewTcp instantiates a new Tcp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTcpWithDefaults

`func NewTcpWithDefaults() *Tcp`

NewTcpWithDefaults instantiates a new Tcp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestinationPortRangeList

`func (o *Tcp) GetDestinationPortRangeList() []PortRange`

GetDestinationPortRangeList returns the DestinationPortRangeList field if non-nil, zero value otherwise.

### GetDestinationPortRangeListOk

`func (o *Tcp) GetDestinationPortRangeListOk() (*[]PortRange, bool)`

GetDestinationPortRangeListOk returns a tuple with the DestinationPortRangeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationPortRangeList

`func (o *Tcp) SetDestinationPortRangeList(v []PortRange)`

SetDestinationPortRangeList sets DestinationPortRangeList field to given value.

### HasDestinationPortRangeList

`func (o *Tcp) HasDestinationPortRangeList() bool`

HasDestinationPortRangeList returns a boolean if a field has been set.

### GetSourcePortRange

`func (o *Tcp) GetSourcePortRange() PortRange`

GetSourcePortRange returns the SourcePortRange field if non-nil, zero value otherwise.

### GetSourcePortRangeOk

`func (o *Tcp) GetSourcePortRangeOk() (*PortRange, bool)`

GetSourcePortRangeOk returns a tuple with the SourcePortRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcePortRange

`func (o *Tcp) SetSourcePortRange(v PortRange)`

SetSourcePortRange sets SourcePortRange field to given value.

### HasSourcePortRange

`func (o *Tcp) HasSourcePortRange() bool`

HasSourcePortRange returns a boolean if a field has been set.

### GetSourcePortRangeList

`func (o *Tcp) GetSourcePortRangeList() []PortRange`

GetSourcePortRangeList returns the SourcePortRangeList field if non-nil, zero value otherwise.

### GetSourcePortRangeListOk

`func (o *Tcp) GetSourcePortRangeListOk() (*[]PortRange, bool)`

GetSourcePortRangeListOk returns a tuple with the SourcePortRangeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcePortRangeList

`func (o *Tcp) SetSourcePortRangeList(v []PortRange)`

SetSourcePortRangeList sets SourcePortRangeList field to given value.

### HasSourcePortRangeList

`func (o *Tcp) HasSourcePortRangeList() bool`

HasSourcePortRangeList returns a boolean if a field has been set.

### GetDestinationPortRange

`func (o *Tcp) GetDestinationPortRange() PortRange`

GetDestinationPortRange returns the DestinationPortRange field if non-nil, zero value otherwise.

### GetDestinationPortRangeOk

`func (o *Tcp) GetDestinationPortRangeOk() (*PortRange, bool)`

GetDestinationPortRangeOk returns a tuple with the DestinationPortRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationPortRange

`func (o *Tcp) SetDestinationPortRange(v PortRange)`

SetDestinationPortRange sets DestinationPortRange field to given value.

### HasDestinationPortRange

`func (o *Tcp) HasDestinationPortRange() bool`

HasDestinationPortRange returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



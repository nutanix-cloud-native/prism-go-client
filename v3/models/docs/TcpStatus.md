# TcpStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestinationPortRangeList** | Pointer to [**[]PortRangeStatus**](PortRangeStatus.md) | List of ranges of TCP destination ports. | [optional] 
**SourcePortRange** | Pointer to [**PortRangeStatus**](PortRangeStatus.md) |  | [optional] 
**SourcePortRangeList** | Pointer to [**[]PortRangeStatus**](PortRangeStatus.md) | List of ranges of TCP source ports. | [optional] 
**DestinationPortRange** | Pointer to [**PortRangeStatus**](PortRangeStatus.md) |  | [optional] 

## Methods

### NewTcpStatus

`func NewTcpStatus() *TcpStatus`

NewTcpStatus instantiates a new TcpStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTcpStatusWithDefaults

`func NewTcpStatusWithDefaults() *TcpStatus`

NewTcpStatusWithDefaults instantiates a new TcpStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestinationPortRangeList

`func (o *TcpStatus) GetDestinationPortRangeList() []PortRangeStatus`

GetDestinationPortRangeList returns the DestinationPortRangeList field if non-nil, zero value otherwise.

### GetDestinationPortRangeListOk

`func (o *TcpStatus) GetDestinationPortRangeListOk() (*[]PortRangeStatus, bool)`

GetDestinationPortRangeListOk returns a tuple with the DestinationPortRangeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationPortRangeList

`func (o *TcpStatus) SetDestinationPortRangeList(v []PortRangeStatus)`

SetDestinationPortRangeList sets DestinationPortRangeList field to given value.

### HasDestinationPortRangeList

`func (o *TcpStatus) HasDestinationPortRangeList() bool`

HasDestinationPortRangeList returns a boolean if a field has been set.

### GetSourcePortRange

`func (o *TcpStatus) GetSourcePortRange() PortRangeStatus`

GetSourcePortRange returns the SourcePortRange field if non-nil, zero value otherwise.

### GetSourcePortRangeOk

`func (o *TcpStatus) GetSourcePortRangeOk() (*PortRangeStatus, bool)`

GetSourcePortRangeOk returns a tuple with the SourcePortRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcePortRange

`func (o *TcpStatus) SetSourcePortRange(v PortRangeStatus)`

SetSourcePortRange sets SourcePortRange field to given value.

### HasSourcePortRange

`func (o *TcpStatus) HasSourcePortRange() bool`

HasSourcePortRange returns a boolean if a field has been set.

### GetSourcePortRangeList

`func (o *TcpStatus) GetSourcePortRangeList() []PortRangeStatus`

GetSourcePortRangeList returns the SourcePortRangeList field if non-nil, zero value otherwise.

### GetSourcePortRangeListOk

`func (o *TcpStatus) GetSourcePortRangeListOk() (*[]PortRangeStatus, bool)`

GetSourcePortRangeListOk returns a tuple with the SourcePortRangeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcePortRangeList

`func (o *TcpStatus) SetSourcePortRangeList(v []PortRangeStatus)`

SetSourcePortRangeList sets SourcePortRangeList field to given value.

### HasSourcePortRangeList

`func (o *TcpStatus) HasSourcePortRangeList() bool`

HasSourcePortRangeList returns a boolean if a field has been set.

### GetDestinationPortRange

`func (o *TcpStatus) GetDestinationPortRange() PortRangeStatus`

GetDestinationPortRange returns the DestinationPortRange field if non-nil, zero value otherwise.

### GetDestinationPortRangeOk

`func (o *TcpStatus) GetDestinationPortRangeOk() (*PortRangeStatus, bool)`

GetDestinationPortRangeOk returns a tuple with the DestinationPortRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationPortRange

`func (o *TcpStatus) SetDestinationPortRange(v PortRangeStatus)`

SetDestinationPortRange sets DestinationPortRange field to given value.

### HasDestinationPortRange

`func (o *TcpStatus) HasDestinationPortRange() bool`

HasDestinationPortRange returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



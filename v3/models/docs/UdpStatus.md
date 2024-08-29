# UdpStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestinationPortRangeList** | Pointer to [**[]PortRangeStatus**](PortRangeStatus.md) | List of ranges of UDP destination ports. | [optional] 
**SourcePortRange** | Pointer to [**PortRangeStatus**](PortRangeStatus.md) |  | [optional] 
**SourcePortRangeList** | Pointer to [**[]PortRangeStatus**](PortRangeStatus.md) | List of ranges of UDP source ports. | [optional] 
**DestinationPortRange** | Pointer to [**PortRangeStatus**](PortRangeStatus.md) |  | [optional] 

## Methods

### NewUdpStatus

`func NewUdpStatus() *UdpStatus`

NewUdpStatus instantiates a new UdpStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUdpStatusWithDefaults

`func NewUdpStatusWithDefaults() *UdpStatus`

NewUdpStatusWithDefaults instantiates a new UdpStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestinationPortRangeList

`func (o *UdpStatus) GetDestinationPortRangeList() []PortRangeStatus`

GetDestinationPortRangeList returns the DestinationPortRangeList field if non-nil, zero value otherwise.

### GetDestinationPortRangeListOk

`func (o *UdpStatus) GetDestinationPortRangeListOk() (*[]PortRangeStatus, bool)`

GetDestinationPortRangeListOk returns a tuple with the DestinationPortRangeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationPortRangeList

`func (o *UdpStatus) SetDestinationPortRangeList(v []PortRangeStatus)`

SetDestinationPortRangeList sets DestinationPortRangeList field to given value.

### HasDestinationPortRangeList

`func (o *UdpStatus) HasDestinationPortRangeList() bool`

HasDestinationPortRangeList returns a boolean if a field has been set.

### GetSourcePortRange

`func (o *UdpStatus) GetSourcePortRange() PortRangeStatus`

GetSourcePortRange returns the SourcePortRange field if non-nil, zero value otherwise.

### GetSourcePortRangeOk

`func (o *UdpStatus) GetSourcePortRangeOk() (*PortRangeStatus, bool)`

GetSourcePortRangeOk returns a tuple with the SourcePortRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcePortRange

`func (o *UdpStatus) SetSourcePortRange(v PortRangeStatus)`

SetSourcePortRange sets SourcePortRange field to given value.

### HasSourcePortRange

`func (o *UdpStatus) HasSourcePortRange() bool`

HasSourcePortRange returns a boolean if a field has been set.

### GetSourcePortRangeList

`func (o *UdpStatus) GetSourcePortRangeList() []PortRangeStatus`

GetSourcePortRangeList returns the SourcePortRangeList field if non-nil, zero value otherwise.

### GetSourcePortRangeListOk

`func (o *UdpStatus) GetSourcePortRangeListOk() (*[]PortRangeStatus, bool)`

GetSourcePortRangeListOk returns a tuple with the SourcePortRangeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcePortRangeList

`func (o *UdpStatus) SetSourcePortRangeList(v []PortRangeStatus)`

SetSourcePortRangeList sets SourcePortRangeList field to given value.

### HasSourcePortRangeList

`func (o *UdpStatus) HasSourcePortRangeList() bool`

HasSourcePortRangeList returns a boolean if a field has been set.

### GetDestinationPortRange

`func (o *UdpStatus) GetDestinationPortRange() PortRangeStatus`

GetDestinationPortRange returns the DestinationPortRange field if non-nil, zero value otherwise.

### GetDestinationPortRangeOk

`func (o *UdpStatus) GetDestinationPortRangeOk() (*PortRangeStatus, bool)`

GetDestinationPortRangeOk returns a tuple with the DestinationPortRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationPortRange

`func (o *UdpStatus) SetDestinationPortRange(v PortRangeStatus)`

SetDestinationPortRange sets DestinationPortRange field to given value.

### HasDestinationPortRange

`func (o *UdpStatus) HasDestinationPortRange() bool`

HasDestinationPortRange returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# Node

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HostReference** | [**Reference**](Reference.md) |  | 
**IpAddress** | Pointer to **string** | Node IP Address | [optional] 

## Methods

### NewNode

`func NewNode(hostReference Reference, ) *Node`

NewNode instantiates a new Node object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeWithDefaults

`func NewNodeWithDefaults() *Node`

NewNodeWithDefaults instantiates a new Node object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostReference

`func (o *Node) GetHostReference() Reference`

GetHostReference returns the HostReference field if non-nil, zero value otherwise.

### GetHostReferenceOk

`func (o *Node) GetHostReferenceOk() (*Reference, bool)`

GetHostReferenceOk returns a tuple with the HostReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostReference

`func (o *Node) SetHostReference(v Reference)`

SetHostReference sets HostReference field to given value.


### GetIpAddress

`func (o *Node) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *Node) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *Node) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *Node) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



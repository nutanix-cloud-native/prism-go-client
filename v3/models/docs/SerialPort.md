# SerialPort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | Pointer to **int32** | Index of the serial port. | [optional] 
**IsConnected** | Pointer to **bool** | Indicates whether the serial port connection is connected or not.  | [optional] 

## Methods

### NewSerialPort

`func NewSerialPort() *SerialPort`

NewSerialPort instantiates a new SerialPort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSerialPortWithDefaults

`func NewSerialPortWithDefaults() *SerialPort`

NewSerialPortWithDefaults instantiates a new SerialPort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndex

`func (o *SerialPort) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *SerialPort) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *SerialPort) SetIndex(v int32)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *SerialPort) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetIsConnected

`func (o *SerialPort) GetIsConnected() bool`

GetIsConnected returns the IsConnected field if non-nil, zero value otherwise.

### GetIsConnectedOk

`func (o *SerialPort) GetIsConnectedOk() (*bool, bool)`

GetIsConnectedOk returns a tuple with the IsConnected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsConnected

`func (o *SerialPort) SetIsConnected(v bool)`

SetIsConnected sets IsConnected field to given value.

### HasIsConnected

`func (o *SerialPort) HasIsConnected() bool`

HasIsConnected returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



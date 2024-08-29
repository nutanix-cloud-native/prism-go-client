# RackableUnitDefStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Serial** | **string** | The rackable unit serial | 
**MessageList** | Pointer to [**[]MessageResource**](MessageResource.md) |  | [optional] 
**Model** | **string** | The rackable unit model | 
**Resources** | [**RackableUnitResources**](RackableUnitResources.md) |  | 
**State** | Pointer to **string** | The state of the rackable unit entity | [optional] [readonly] 

## Methods

### NewRackableUnitDefStatus

`func NewRackableUnitDefStatus(serial string, model string, resources RackableUnitResources, ) *RackableUnitDefStatus`

NewRackableUnitDefStatus instantiates a new RackableUnitDefStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRackableUnitDefStatusWithDefaults

`func NewRackableUnitDefStatusWithDefaults() *RackableUnitDefStatus`

NewRackableUnitDefStatusWithDefaults instantiates a new RackableUnitDefStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerial

`func (o *RackableUnitDefStatus) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *RackableUnitDefStatus) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *RackableUnitDefStatus) SetSerial(v string)`

SetSerial sets Serial field to given value.


### GetMessageList

`func (o *RackableUnitDefStatus) GetMessageList() []MessageResource`

GetMessageList returns the MessageList field if non-nil, zero value otherwise.

### GetMessageListOk

`func (o *RackableUnitDefStatus) GetMessageListOk() (*[]MessageResource, bool)`

GetMessageListOk returns a tuple with the MessageList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageList

`func (o *RackableUnitDefStatus) SetMessageList(v []MessageResource)`

SetMessageList sets MessageList field to given value.

### HasMessageList

`func (o *RackableUnitDefStatus) HasMessageList() bool`

HasMessageList returns a boolean if a field has been set.

### GetModel

`func (o *RackableUnitDefStatus) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *RackableUnitDefStatus) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *RackableUnitDefStatus) SetModel(v string)`

SetModel sets Model field to given value.


### GetResources

`func (o *RackableUnitDefStatus) GetResources() RackableUnitResources`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *RackableUnitDefStatus) GetResourcesOk() (*RackableUnitResources, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *RackableUnitDefStatus) SetResources(v RackableUnitResources)`

SetResources sets Resources field to given value.


### GetState

`func (o *RackableUnitDefStatus) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *RackableUnitDefStatus) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *RackableUnitDefStatus) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *RackableUnitDefStatus) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



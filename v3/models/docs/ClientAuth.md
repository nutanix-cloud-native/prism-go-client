# ClientAuth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** | Status of client authentication. | [default to "DISABLED"]
**CaChain** | Pointer to **string** | Content of CA chain certificate. | [optional] [readonly] 
**Name** | Pointer to **string** | Name of CA chain file. | [optional] [readonly] 

## Methods

### NewClientAuth

`func NewClientAuth(status string, ) *ClientAuth`

NewClientAuth instantiates a new ClientAuth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientAuthWithDefaults

`func NewClientAuthWithDefaults() *ClientAuth`

NewClientAuthWithDefaults instantiates a new ClientAuth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ClientAuth) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ClientAuth) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ClientAuth) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetCaChain

`func (o *ClientAuth) GetCaChain() string`

GetCaChain returns the CaChain field if non-nil, zero value otherwise.

### GetCaChainOk

`func (o *ClientAuth) GetCaChainOk() (*string, bool)`

GetCaChainOk returns a tuple with the CaChain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaChain

`func (o *ClientAuth) SetCaChain(v string)`

SetCaChain sets CaChain field to given value.

### HasCaChain

`func (o *ClientAuth) HasCaChain() bool`

HasCaChain returns a boolean if a field has been set.

### GetName

`func (o *ClientAuth) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClientAuth) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClientAuth) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ClientAuth) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



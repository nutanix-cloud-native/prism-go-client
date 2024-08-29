# DirectConnectVifStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | Pointer to **string** |  | [optional] 
**ErrorDetail** | Pointer to **string** | Detailed informational/error string in human-readable form. | [optional] 

## Methods

### NewDirectConnectVifStatus

`func NewDirectConnectVifStatus() *DirectConnectVifStatus`

NewDirectConnectVifStatus instantiates a new DirectConnectVifStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDirectConnectVifStatusWithDefaults

`func NewDirectConnectVifStatusWithDefaults() *DirectConnectVifStatus`

NewDirectConnectVifStatusWithDefaults instantiates a new DirectConnectVifStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *DirectConnectVifStatus) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *DirectConnectVifStatus) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *DirectConnectVifStatus) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *DirectConnectVifStatus) HasState() bool`

HasState returns a boolean if a field has been set.

### GetErrorDetail

`func (o *DirectConnectVifStatus) GetErrorDetail() string`

GetErrorDetail returns the ErrorDetail field if non-nil, zero value otherwise.

### GetErrorDetailOk

`func (o *DirectConnectVifStatus) GetErrorDetailOk() (*string, bool)`

GetErrorDetailOk returns a tuple with the ErrorDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDetail

`func (o *DirectConnectVifStatus) SetErrorDetail(v string)`

SetErrorDetail sets ErrorDetail field to given value.

### HasErrorDetail

`func (o *DirectConnectVifStatus) HasErrorDetail() bool`

HasErrorDetail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



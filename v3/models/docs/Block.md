# Block

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlockSerialNumber** | Pointer to **string** | Rackable unit serial number. | [optional] 
**BlockModel** | Pointer to **string** | Rackable unit model name. | [optional] 

## Methods

### NewBlock

`func NewBlock() *Block`

NewBlock instantiates a new Block object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlockWithDefaults

`func NewBlockWithDefaults() *Block`

NewBlockWithDefaults instantiates a new Block object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockSerialNumber

`func (o *Block) GetBlockSerialNumber() string`

GetBlockSerialNumber returns the BlockSerialNumber field if non-nil, zero value otherwise.

### GetBlockSerialNumberOk

`func (o *Block) GetBlockSerialNumberOk() (*string, bool)`

GetBlockSerialNumberOk returns a tuple with the BlockSerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockSerialNumber

`func (o *Block) SetBlockSerialNumber(v string)`

SetBlockSerialNumber sets BlockSerialNumber field to given value.

### HasBlockSerialNumber

`func (o *Block) HasBlockSerialNumber() bool`

HasBlockSerialNumber returns a boolean if a field has been set.

### GetBlockModel

`func (o *Block) GetBlockModel() string`

GetBlockModel returns the BlockModel field if non-nil, zero value otherwise.

### GetBlockModelOk

`func (o *Block) GetBlockModelOk() (*string, bool)`

GetBlockModelOk returns a tuple with the BlockModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockModel

`func (o *Block) SetBlockModel(v string)`

SetBlockModel sets BlockModel field to given value.

### HasBlockModel

`func (o *Block) HasBlockModel() bool`

HasBlockModel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



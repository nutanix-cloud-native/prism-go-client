# CompressionParametersForTheEntities

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | Pointer to **string** | Enable or disable compression for entities governed by the policy. If user does not have an explicit preference for compression, the system decides an appropriate value.  | [optional] [default to "SYSTEM_DERIVED"]
**InlineCompression** | Pointer to **bool** | Indicates inline or post-process compression. If compression state is SYSTEM_DERIVED, then this field has no significance.  | [optional] 

## Methods

### NewCompressionParametersForTheEntities

`func NewCompressionParametersForTheEntities() *CompressionParametersForTheEntities`

NewCompressionParametersForTheEntities instantiates a new CompressionParametersForTheEntities object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompressionParametersForTheEntitiesWithDefaults

`func NewCompressionParametersForTheEntitiesWithDefaults() *CompressionParametersForTheEntities`

NewCompressionParametersForTheEntitiesWithDefaults instantiates a new CompressionParametersForTheEntities object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *CompressionParametersForTheEntities) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CompressionParametersForTheEntities) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CompressionParametersForTheEntities) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *CompressionParametersForTheEntities) HasState() bool`

HasState returns a boolean if a field has been set.

### GetInlineCompression

`func (o *CompressionParametersForTheEntities) GetInlineCompression() bool`

GetInlineCompression returns the InlineCompression field if non-nil, zero value otherwise.

### GetInlineCompressionOk

`func (o *CompressionParametersForTheEntities) GetInlineCompressionOk() (*bool, bool)`

GetInlineCompressionOk returns a tuple with the InlineCompression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInlineCompression

`func (o *CompressionParametersForTheEntities) SetInlineCompression(v bool)`

SetInlineCompression sets InlineCompression field to given value.

### HasInlineCompression

`func (o *CompressionParametersForTheEntities) HasInlineCompression() bool`

HasInlineCompression returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



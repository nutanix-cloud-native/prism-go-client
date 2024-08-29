# CapabilityDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceUuid** | Pointer to **string** | The UUID for the cluster that contained capability | [optional] 
**Notes** | Pointer to **[]string** | Notes related to this capability | [optional] 
**Value** | Pointer to [**CapabilityValue**](CapabilityValue.md) |  | [optional] 

## Methods

### NewCapabilityDetails

`func NewCapabilityDetails() *CapabilityDetails`

NewCapabilityDetails instantiates a new CapabilityDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapabilityDetailsWithDefaults

`func NewCapabilityDetailsWithDefaults() *CapabilityDetails`

NewCapabilityDetailsWithDefaults instantiates a new CapabilityDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceUuid

`func (o *CapabilityDetails) GetSourceUuid() string`

GetSourceUuid returns the SourceUuid field if non-nil, zero value otherwise.

### GetSourceUuidOk

`func (o *CapabilityDetails) GetSourceUuidOk() (*string, bool)`

GetSourceUuidOk returns a tuple with the SourceUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUuid

`func (o *CapabilityDetails) SetSourceUuid(v string)`

SetSourceUuid sets SourceUuid field to given value.

### HasSourceUuid

`func (o *CapabilityDetails) HasSourceUuid() bool`

HasSourceUuid returns a boolean if a field has been set.

### GetNotes

`func (o *CapabilityDetails) GetNotes() []string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *CapabilityDetails) GetNotesOk() (*[]string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *CapabilityDetails) SetNotes(v []string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *CapabilityDetails) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetValue

`func (o *CapabilityDetails) GetValue() CapabilityValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CapabilityDetails) GetValueOk() (*CapabilityValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CapabilityDetails) SetValue(v CapabilityValue)`

SetValue sets Value field to given value.

### HasValue

`func (o *CapabilityDetails) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



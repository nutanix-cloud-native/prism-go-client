# WitnessConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WitnessAddress** | Pointer to **string** | Address of the witness, which will be witnessing this Recovery Plan. This will be same as Availability Zone URL, on which witness is deployed.  | [optional] 
**WitnessFailoverTimeoutSecs** | Pointer to **int32** | Time in seconds after which witness will trigger failover on this Recovery Plan in case there is a failure of source cluster or pause replication for entities in case there is a failure of target cluster.  | [optional] 

## Methods

### NewWitnessConfiguration

`func NewWitnessConfiguration() *WitnessConfiguration`

NewWitnessConfiguration instantiates a new WitnessConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWitnessConfigurationWithDefaults

`func NewWitnessConfigurationWithDefaults() *WitnessConfiguration`

NewWitnessConfigurationWithDefaults instantiates a new WitnessConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWitnessAddress

`func (o *WitnessConfiguration) GetWitnessAddress() string`

GetWitnessAddress returns the WitnessAddress field if non-nil, zero value otherwise.

### GetWitnessAddressOk

`func (o *WitnessConfiguration) GetWitnessAddressOk() (*string, bool)`

GetWitnessAddressOk returns a tuple with the WitnessAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWitnessAddress

`func (o *WitnessConfiguration) SetWitnessAddress(v string)`

SetWitnessAddress sets WitnessAddress field to given value.

### HasWitnessAddress

`func (o *WitnessConfiguration) HasWitnessAddress() bool`

HasWitnessAddress returns a boolean if a field has been set.

### GetWitnessFailoverTimeoutSecs

`func (o *WitnessConfiguration) GetWitnessFailoverTimeoutSecs() int32`

GetWitnessFailoverTimeoutSecs returns the WitnessFailoverTimeoutSecs field if non-nil, zero value otherwise.

### GetWitnessFailoverTimeoutSecsOk

`func (o *WitnessConfiguration) GetWitnessFailoverTimeoutSecsOk() (*int32, bool)`

GetWitnessFailoverTimeoutSecsOk returns a tuple with the WitnessFailoverTimeoutSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWitnessFailoverTimeoutSecs

`func (o *WitnessConfiguration) SetWitnessFailoverTimeoutSecs(v int32)`

SetWitnessFailoverTimeoutSecs sets WitnessFailoverTimeoutSecs field to given value.

### HasWitnessFailoverTimeoutSecs

`func (o *WitnessConfiguration) HasWitnessFailoverTimeoutSecs() bool`

HasWitnessFailoverTimeoutSecs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



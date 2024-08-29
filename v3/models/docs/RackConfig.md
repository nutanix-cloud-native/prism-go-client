# RackConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HostList** | Pointer to [**[]HostReference**](HostReference.md) |  | [optional] 
**RackName** | Pointer to **string** |  | [optional] 

## Methods

### NewRackConfig

`func NewRackConfig() *RackConfig`

NewRackConfig instantiates a new RackConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRackConfigWithDefaults

`func NewRackConfigWithDefaults() *RackConfig`

NewRackConfigWithDefaults instantiates a new RackConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostList

`func (o *RackConfig) GetHostList() []HostReference`

GetHostList returns the HostList field if non-nil, zero value otherwise.

### GetHostListOk

`func (o *RackConfig) GetHostListOk() (*[]HostReference, bool)`

GetHostListOk returns a tuple with the HostList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostList

`func (o *RackConfig) SetHostList(v []HostReference)`

SetHostList sets HostList field to given value.

### HasHostList

`func (o *RackConfig) HasHostList() bool`

HasHostList returns a boolean if a field has been set.

### GetRackName

`func (o *RackConfig) GetRackName() string`

GetRackName returns the RackName field if non-nil, zero value otherwise.

### GetRackNameOk

`func (o *RackConfig) GetRackNameOk() (*string, bool)`

GetRackNameOk returns a tuple with the RackName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRackName

`func (o *RackConfig) SetRackName(v string)`

SetRackName sets RackName field to given value.

### HasRackName

`func (o *RackConfig) HasRackName() bool`

HasRackName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



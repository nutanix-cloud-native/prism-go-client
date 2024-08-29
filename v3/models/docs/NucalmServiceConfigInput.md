# NucalmServiceConfigInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnableNutanixApps** | Pointer to **bool** | Flag indicating whether to enable Nutanix apps. | [optional] 
**State** | Pointer to **string** | The desired state of NuCalm service. | [optional] 
**EnableLite** | Pointer to **bool** | Enable lite versions of service. | [optional] 
**PerformValidationOnly** | Pointer to **bool** | Flag indicating whether to do NuCalm enablement validation only.  | [optional] 

## Methods

### NewNucalmServiceConfigInput

`func NewNucalmServiceConfigInput() *NucalmServiceConfigInput`

NewNucalmServiceConfigInput instantiates a new NucalmServiceConfigInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNucalmServiceConfigInputWithDefaults

`func NewNucalmServiceConfigInputWithDefaults() *NucalmServiceConfigInput`

NewNucalmServiceConfigInputWithDefaults instantiates a new NucalmServiceConfigInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnableNutanixApps

`func (o *NucalmServiceConfigInput) GetEnableNutanixApps() bool`

GetEnableNutanixApps returns the EnableNutanixApps field if non-nil, zero value otherwise.

### GetEnableNutanixAppsOk

`func (o *NucalmServiceConfigInput) GetEnableNutanixAppsOk() (*bool, bool)`

GetEnableNutanixAppsOk returns a tuple with the EnableNutanixApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableNutanixApps

`func (o *NucalmServiceConfigInput) SetEnableNutanixApps(v bool)`

SetEnableNutanixApps sets EnableNutanixApps field to given value.

### HasEnableNutanixApps

`func (o *NucalmServiceConfigInput) HasEnableNutanixApps() bool`

HasEnableNutanixApps returns a boolean if a field has been set.

### GetState

`func (o *NucalmServiceConfigInput) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *NucalmServiceConfigInput) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *NucalmServiceConfigInput) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *NucalmServiceConfigInput) HasState() bool`

HasState returns a boolean if a field has been set.

### GetEnableLite

`func (o *NucalmServiceConfigInput) GetEnableLite() bool`

GetEnableLite returns the EnableLite field if non-nil, zero value otherwise.

### GetEnableLiteOk

`func (o *NucalmServiceConfigInput) GetEnableLiteOk() (*bool, bool)`

GetEnableLiteOk returns a tuple with the EnableLite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableLite

`func (o *NucalmServiceConfigInput) SetEnableLite(v bool)`

SetEnableLite sets EnableLite field to given value.

### HasEnableLite

`func (o *NucalmServiceConfigInput) HasEnableLite() bool`

HasEnableLite returns a boolean if a field has been set.

### GetPerformValidationOnly

`func (o *NucalmServiceConfigInput) GetPerformValidationOnly() bool`

GetPerformValidationOnly returns the PerformValidationOnly field if non-nil, zero value otherwise.

### GetPerformValidationOnlyOk

`func (o *NucalmServiceConfigInput) GetPerformValidationOnlyOk() (*bool, bool)`

GetPerformValidationOnlyOk returns a tuple with the PerformValidationOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformValidationOnly

`func (o *NucalmServiceConfigInput) SetPerformValidationOnly(v bool)`

SetPerformValidationOnly sets PerformValidationOnly field to given value.

### HasPerformValidationOnly

`func (o *NucalmServiceConfigInput) HasPerformValidationOnly() bool`

HasPerformValidationOnly returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# CmspConfigEnablement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RegistryImage** | Pointer to [**ImageSpec**](ImageSpec.md) |  | [optional] 
**Operation** | **string** | The type of operation is kValidate or kEnable.  | 
**Config** | [**CmspConfig**](CmspConfig.md) |  | 
**Source** | Pointer to **string** | Indicates if CMSP enablement is being made as part of PC deployment (value kPE) | [optional] [default to "kPC"]

## Methods

### NewCmspConfigEnablement

`func NewCmspConfigEnablement(operation string, config CmspConfig, ) *CmspConfigEnablement`

NewCmspConfigEnablement instantiates a new CmspConfigEnablement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCmspConfigEnablementWithDefaults

`func NewCmspConfigEnablementWithDefaults() *CmspConfigEnablement`

NewCmspConfigEnablementWithDefaults instantiates a new CmspConfigEnablement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegistryImage

`func (o *CmspConfigEnablement) GetRegistryImage() ImageSpec`

GetRegistryImage returns the RegistryImage field if non-nil, zero value otherwise.

### GetRegistryImageOk

`func (o *CmspConfigEnablement) GetRegistryImageOk() (*ImageSpec, bool)`

GetRegistryImageOk returns a tuple with the RegistryImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistryImage

`func (o *CmspConfigEnablement) SetRegistryImage(v ImageSpec)`

SetRegistryImage sets RegistryImage field to given value.

### HasRegistryImage

`func (o *CmspConfigEnablement) HasRegistryImage() bool`

HasRegistryImage returns a boolean if a field has been set.

### GetOperation

`func (o *CmspConfigEnablement) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *CmspConfigEnablement) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *CmspConfigEnablement) SetOperation(v string)`

SetOperation sets Operation field to given value.


### GetConfig

`func (o *CmspConfigEnablement) GetConfig() CmspConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *CmspConfigEnablement) GetConfigOk() (*CmspConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *CmspConfigEnablement) SetConfig(v CmspConfig)`

SetConfig sets Config field to given value.


### GetSource

`func (o *CmspConfigEnablement) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *CmspConfigEnablement) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *CmspConfigEnablement) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *CmspConfigEnablement) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



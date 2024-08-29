# BrownfieldInstanceInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceId** | **string** | Id of the vm | 
**InstanceName** | Pointer to **string** | Name of the vm | [optional] 
**PlatformData** | Pointer to **map[string]string** |  | [optional] 
**Address** | Pointer to **[]string** | Address of the vm | [optional] 

## Methods

### NewBrownfieldInstanceInput

`func NewBrownfieldInstanceInput(instanceId string, ) *BrownfieldInstanceInput`

NewBrownfieldInstanceInput instantiates a new BrownfieldInstanceInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBrownfieldInstanceInputWithDefaults

`func NewBrownfieldInstanceInputWithDefaults() *BrownfieldInstanceInput`

NewBrownfieldInstanceInputWithDefaults instantiates a new BrownfieldInstanceInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceId

`func (o *BrownfieldInstanceInput) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *BrownfieldInstanceInput) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *BrownfieldInstanceInput) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.


### GetInstanceName

`func (o *BrownfieldInstanceInput) GetInstanceName() string`

GetInstanceName returns the InstanceName field if non-nil, zero value otherwise.

### GetInstanceNameOk

`func (o *BrownfieldInstanceInput) GetInstanceNameOk() (*string, bool)`

GetInstanceNameOk returns a tuple with the InstanceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceName

`func (o *BrownfieldInstanceInput) SetInstanceName(v string)`

SetInstanceName sets InstanceName field to given value.

### HasInstanceName

`func (o *BrownfieldInstanceInput) HasInstanceName() bool`

HasInstanceName returns a boolean if a field has been set.

### GetPlatformData

`func (o *BrownfieldInstanceInput) GetPlatformData() map[string]string`

GetPlatformData returns the PlatformData field if non-nil, zero value otherwise.

### GetPlatformDataOk

`func (o *BrownfieldInstanceInput) GetPlatformDataOk() (*map[string]string, bool)`

GetPlatformDataOk returns a tuple with the PlatformData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformData

`func (o *BrownfieldInstanceInput) SetPlatformData(v map[string]string)`

SetPlatformData sets PlatformData field to given value.

### HasPlatformData

`func (o *BrownfieldInstanceInput) HasPlatformData() bool`

HasPlatformData returns a boolean if a field has been set.

### GetAddress

`func (o *BrownfieldInstanceInput) GetAddress() []string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *BrownfieldInstanceInput) GetAddressOk() (*[]string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *BrownfieldInstanceInput) SetAddress(v []string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *BrownfieldInstanceInput) HasAddress() bool`

HasAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



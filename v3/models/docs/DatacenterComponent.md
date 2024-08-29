# DatacenterComponent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RackUnitHeight** | Pointer to **int32** | rack unit height | [optional] 
**ManagementIp** | Pointer to **string** | management IP address | [optional] 
**ManagementMacAddress** | Pointer to **string** | management mac address | [optional] 
**RackUnitPosition** | Pointer to **int32** | rack unit position | [optional] 
**InstallationTime** | Pointer to **time.Time** | installation date time | [optional] 

## Methods

### NewDatacenterComponent

`func NewDatacenterComponent() *DatacenterComponent`

NewDatacenterComponent instantiates a new DatacenterComponent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatacenterComponentWithDefaults

`func NewDatacenterComponentWithDefaults() *DatacenterComponent`

NewDatacenterComponentWithDefaults instantiates a new DatacenterComponent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRackUnitHeight

`func (o *DatacenterComponent) GetRackUnitHeight() int32`

GetRackUnitHeight returns the RackUnitHeight field if non-nil, zero value otherwise.

### GetRackUnitHeightOk

`func (o *DatacenterComponent) GetRackUnitHeightOk() (*int32, bool)`

GetRackUnitHeightOk returns a tuple with the RackUnitHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRackUnitHeight

`func (o *DatacenterComponent) SetRackUnitHeight(v int32)`

SetRackUnitHeight sets RackUnitHeight field to given value.

### HasRackUnitHeight

`func (o *DatacenterComponent) HasRackUnitHeight() bool`

HasRackUnitHeight returns a boolean if a field has been set.

### GetManagementIp

`func (o *DatacenterComponent) GetManagementIp() string`

GetManagementIp returns the ManagementIp field if non-nil, zero value otherwise.

### GetManagementIpOk

`func (o *DatacenterComponent) GetManagementIpOk() (*string, bool)`

GetManagementIpOk returns a tuple with the ManagementIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementIp

`func (o *DatacenterComponent) SetManagementIp(v string)`

SetManagementIp sets ManagementIp field to given value.

### HasManagementIp

`func (o *DatacenterComponent) HasManagementIp() bool`

HasManagementIp returns a boolean if a field has been set.

### GetManagementMacAddress

`func (o *DatacenterComponent) GetManagementMacAddress() string`

GetManagementMacAddress returns the ManagementMacAddress field if non-nil, zero value otherwise.

### GetManagementMacAddressOk

`func (o *DatacenterComponent) GetManagementMacAddressOk() (*string, bool)`

GetManagementMacAddressOk returns a tuple with the ManagementMacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementMacAddress

`func (o *DatacenterComponent) SetManagementMacAddress(v string)`

SetManagementMacAddress sets ManagementMacAddress field to given value.

### HasManagementMacAddress

`func (o *DatacenterComponent) HasManagementMacAddress() bool`

HasManagementMacAddress returns a boolean if a field has been set.

### GetRackUnitPosition

`func (o *DatacenterComponent) GetRackUnitPosition() int32`

GetRackUnitPosition returns the RackUnitPosition field if non-nil, zero value otherwise.

### GetRackUnitPositionOk

`func (o *DatacenterComponent) GetRackUnitPositionOk() (*int32, bool)`

GetRackUnitPositionOk returns a tuple with the RackUnitPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRackUnitPosition

`func (o *DatacenterComponent) SetRackUnitPosition(v int32)`

SetRackUnitPosition sets RackUnitPosition field to given value.

### HasRackUnitPosition

`func (o *DatacenterComponent) HasRackUnitPosition() bool`

HasRackUnitPosition returns a boolean if a field has been set.

### GetInstallationTime

`func (o *DatacenterComponent) GetInstallationTime() time.Time`

GetInstallationTime returns the InstallationTime field if non-nil, zero value otherwise.

### GetInstallationTimeOk

`func (o *DatacenterComponent) GetInstallationTimeOk() (*time.Time, bool)`

GetInstallationTimeOk returns a tuple with the InstallationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallationTime

`func (o *DatacenterComponent) SetInstallationTime(v time.Time)`

SetInstallationTime sets InstallationTime field to given value.

### HasInstallationTime

`func (o *DatacenterComponent) HasInstallationTime() bool`

HasInstallationTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



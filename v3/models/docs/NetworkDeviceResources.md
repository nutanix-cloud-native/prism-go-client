# NetworkDeviceResources

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**CurrentFirmwareVersion** | Pointer to **string** | The current firmware version | [optional] 
**DeviceClass** | Pointer to **string** | A well known string for network device class information | [optional] 
**Component** | Pointer to [**DatacenterComponent**](DatacenterComponent.md) |  | [optional] 
**IpAddress** | Pointer to **string** | device IP address | [optional] 
**UpgradeStatus** | Pointer to **string** | upgrade status | [optional] 
**RackReference** | Pointer to [**RackReference**](RackReference.md) |  | [optional] 
**Model** | Pointer to **string** | device model | [optional] 
**DeviceSerial** | Pointer to **string** | Device serial number | [optional] 
**TargetFirmwareVersion** | Pointer to **string** | The target firmware version | [optional] 

## Methods

### NewNetworkDeviceResources

`func NewNetworkDeviceResources() *NetworkDeviceResources`

NewNetworkDeviceResources instantiates a new NetworkDeviceResources object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkDeviceResourcesWithDefaults

`func NewNetworkDeviceResourcesWithDefaults() *NetworkDeviceResources`

NewNetworkDeviceResourcesWithDefaults instantiates a new NetworkDeviceResources object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *NetworkDeviceResources) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NetworkDeviceResources) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NetworkDeviceResources) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NetworkDeviceResources) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCurrentFirmwareVersion

`func (o *NetworkDeviceResources) GetCurrentFirmwareVersion() string`

GetCurrentFirmwareVersion returns the CurrentFirmwareVersion field if non-nil, zero value otherwise.

### GetCurrentFirmwareVersionOk

`func (o *NetworkDeviceResources) GetCurrentFirmwareVersionOk() (*string, bool)`

GetCurrentFirmwareVersionOk returns a tuple with the CurrentFirmwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentFirmwareVersion

`func (o *NetworkDeviceResources) SetCurrentFirmwareVersion(v string)`

SetCurrentFirmwareVersion sets CurrentFirmwareVersion field to given value.

### HasCurrentFirmwareVersion

`func (o *NetworkDeviceResources) HasCurrentFirmwareVersion() bool`

HasCurrentFirmwareVersion returns a boolean if a field has been set.

### GetDeviceClass

`func (o *NetworkDeviceResources) GetDeviceClass() string`

GetDeviceClass returns the DeviceClass field if non-nil, zero value otherwise.

### GetDeviceClassOk

`func (o *NetworkDeviceResources) GetDeviceClassOk() (*string, bool)`

GetDeviceClassOk returns a tuple with the DeviceClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceClass

`func (o *NetworkDeviceResources) SetDeviceClass(v string)`

SetDeviceClass sets DeviceClass field to given value.

### HasDeviceClass

`func (o *NetworkDeviceResources) HasDeviceClass() bool`

HasDeviceClass returns a boolean if a field has been set.

### GetComponent

`func (o *NetworkDeviceResources) GetComponent() DatacenterComponent`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *NetworkDeviceResources) GetComponentOk() (*DatacenterComponent, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *NetworkDeviceResources) SetComponent(v DatacenterComponent)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *NetworkDeviceResources) HasComponent() bool`

HasComponent returns a boolean if a field has been set.

### GetIpAddress

`func (o *NetworkDeviceResources) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *NetworkDeviceResources) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *NetworkDeviceResources) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *NetworkDeviceResources) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetUpgradeStatus

`func (o *NetworkDeviceResources) GetUpgradeStatus() string`

GetUpgradeStatus returns the UpgradeStatus field if non-nil, zero value otherwise.

### GetUpgradeStatusOk

`func (o *NetworkDeviceResources) GetUpgradeStatusOk() (*string, bool)`

GetUpgradeStatusOk returns a tuple with the UpgradeStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeStatus

`func (o *NetworkDeviceResources) SetUpgradeStatus(v string)`

SetUpgradeStatus sets UpgradeStatus field to given value.

### HasUpgradeStatus

`func (o *NetworkDeviceResources) HasUpgradeStatus() bool`

HasUpgradeStatus returns a boolean if a field has been set.

### GetRackReference

`func (o *NetworkDeviceResources) GetRackReference() RackReference`

GetRackReference returns the RackReference field if non-nil, zero value otherwise.

### GetRackReferenceOk

`func (o *NetworkDeviceResources) GetRackReferenceOk() (*RackReference, bool)`

GetRackReferenceOk returns a tuple with the RackReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRackReference

`func (o *NetworkDeviceResources) SetRackReference(v RackReference)`

SetRackReference sets RackReference field to given value.

### HasRackReference

`func (o *NetworkDeviceResources) HasRackReference() bool`

HasRackReference returns a boolean if a field has been set.

### GetModel

`func (o *NetworkDeviceResources) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *NetworkDeviceResources) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *NetworkDeviceResources) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *NetworkDeviceResources) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetDeviceSerial

`func (o *NetworkDeviceResources) GetDeviceSerial() string`

GetDeviceSerial returns the DeviceSerial field if non-nil, zero value otherwise.

### GetDeviceSerialOk

`func (o *NetworkDeviceResources) GetDeviceSerialOk() (*string, bool)`

GetDeviceSerialOk returns a tuple with the DeviceSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceSerial

`func (o *NetworkDeviceResources) SetDeviceSerial(v string)`

SetDeviceSerial sets DeviceSerial field to given value.

### HasDeviceSerial

`func (o *NetworkDeviceResources) HasDeviceSerial() bool`

HasDeviceSerial returns a boolean if a field has been set.

### GetTargetFirmwareVersion

`func (o *NetworkDeviceResources) GetTargetFirmwareVersion() string`

GetTargetFirmwareVersion returns the TargetFirmwareVersion field if non-nil, zero value otherwise.

### GetTargetFirmwareVersionOk

`func (o *NetworkDeviceResources) GetTargetFirmwareVersionOk() (*string, bool)`

GetTargetFirmwareVersionOk returns a tuple with the TargetFirmwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetFirmwareVersion

`func (o *NetworkDeviceResources) SetTargetFirmwareVersion(v string)`

SetTargetFirmwareVersion sets TargetFirmwareVersion field to given value.

### HasTargetFirmwareVersion

`func (o *NetworkDeviceResources) HasTargetFirmwareVersion() bool`

HasTargetFirmwareVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



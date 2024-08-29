# AvailabilityZoneResources

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ManagementUrl** | Pointer to **string** | Identifier of the management plane. This could be the URL of the PC or the FQDN of Xi portal.  | [optional] 
**Region** | Pointer to **string** | Cloud region where the data will be replicated to. Based on the cloud provider type the available list of regions will differ.  | [optional] 
**ManagementPlaneType** | **string** | This defines the type of management entity. Its value can be Xi, PC, or Local. Local AZs are auto-created and cannot be deleted. How to talk to management entity will be decided based on the type of management plane.  | 
**DisplayName** | Pointer to **string** | Display name. It is mainly used by user interface to show the user-friendly name of the availability zone. If unset, default value will be used.  | [optional] 
**Credentials** | Pointer to [**AvailabilityZoneResourcesSpecCredentials**](AvailabilityZoneResourcesSpecCredentials.md) |  | [optional] 

## Methods

### NewAvailabilityZoneResources

`func NewAvailabilityZoneResources(managementPlaneType string, ) *AvailabilityZoneResources`

NewAvailabilityZoneResources instantiates a new AvailabilityZoneResources object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAvailabilityZoneResourcesWithDefaults

`func NewAvailabilityZoneResourcesWithDefaults() *AvailabilityZoneResources`

NewAvailabilityZoneResourcesWithDefaults instantiates a new AvailabilityZoneResources object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetManagementUrl

`func (o *AvailabilityZoneResources) GetManagementUrl() string`

GetManagementUrl returns the ManagementUrl field if non-nil, zero value otherwise.

### GetManagementUrlOk

`func (o *AvailabilityZoneResources) GetManagementUrlOk() (*string, bool)`

GetManagementUrlOk returns a tuple with the ManagementUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementUrl

`func (o *AvailabilityZoneResources) SetManagementUrl(v string)`

SetManagementUrl sets ManagementUrl field to given value.

### HasManagementUrl

`func (o *AvailabilityZoneResources) HasManagementUrl() bool`

HasManagementUrl returns a boolean if a field has been set.

### GetRegion

`func (o *AvailabilityZoneResources) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *AvailabilityZoneResources) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *AvailabilityZoneResources) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *AvailabilityZoneResources) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetManagementPlaneType

`func (o *AvailabilityZoneResources) GetManagementPlaneType() string`

GetManagementPlaneType returns the ManagementPlaneType field if non-nil, zero value otherwise.

### GetManagementPlaneTypeOk

`func (o *AvailabilityZoneResources) GetManagementPlaneTypeOk() (*string, bool)`

GetManagementPlaneTypeOk returns a tuple with the ManagementPlaneType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementPlaneType

`func (o *AvailabilityZoneResources) SetManagementPlaneType(v string)`

SetManagementPlaneType sets ManagementPlaneType field to given value.


### GetDisplayName

`func (o *AvailabilityZoneResources) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *AvailabilityZoneResources) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *AvailabilityZoneResources) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *AvailabilityZoneResources) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetCredentials

`func (o *AvailabilityZoneResources) GetCredentials() AvailabilityZoneResourcesSpecCredentials`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *AvailabilityZoneResources) GetCredentialsOk() (*AvailabilityZoneResourcesSpecCredentials, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *AvailabilityZoneResources) SetCredentials(v AvailabilityZoneResourcesSpecCredentials)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *AvailabilityZoneResources) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



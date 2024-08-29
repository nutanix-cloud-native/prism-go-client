# SecurityPlanningDataCollectorUpgradeStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentVersion** | Pointer to **string** | Current running version of the data collector  | [optional] 
**TaskUuid** | Pointer to **string** | UUID of the data collector upgrade task that&#39;s currently running. Absence of this field implies that an upgrade is not in progress.  | [optional] 
**IsUpgradeAvailable** | Pointer to **bool** | Is a data collector upgrade available  | [optional] 
**LatestVersion** | Pointer to **string** | Latest available version of the data collector  | [optional] 

## Methods

### NewSecurityPlanningDataCollectorUpgradeStatus

`func NewSecurityPlanningDataCollectorUpgradeStatus() *SecurityPlanningDataCollectorUpgradeStatus`

NewSecurityPlanningDataCollectorUpgradeStatus instantiates a new SecurityPlanningDataCollectorUpgradeStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityPlanningDataCollectorUpgradeStatusWithDefaults

`func NewSecurityPlanningDataCollectorUpgradeStatusWithDefaults() *SecurityPlanningDataCollectorUpgradeStatus`

NewSecurityPlanningDataCollectorUpgradeStatusWithDefaults instantiates a new SecurityPlanningDataCollectorUpgradeStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentVersion

`func (o *SecurityPlanningDataCollectorUpgradeStatus) GetCurrentVersion() string`

GetCurrentVersion returns the CurrentVersion field if non-nil, zero value otherwise.

### GetCurrentVersionOk

`func (o *SecurityPlanningDataCollectorUpgradeStatus) GetCurrentVersionOk() (*string, bool)`

GetCurrentVersionOk returns a tuple with the CurrentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentVersion

`func (o *SecurityPlanningDataCollectorUpgradeStatus) SetCurrentVersion(v string)`

SetCurrentVersion sets CurrentVersion field to given value.

### HasCurrentVersion

`func (o *SecurityPlanningDataCollectorUpgradeStatus) HasCurrentVersion() bool`

HasCurrentVersion returns a boolean if a field has been set.

### GetTaskUuid

`func (o *SecurityPlanningDataCollectorUpgradeStatus) GetTaskUuid() string`

GetTaskUuid returns the TaskUuid field if non-nil, zero value otherwise.

### GetTaskUuidOk

`func (o *SecurityPlanningDataCollectorUpgradeStatus) GetTaskUuidOk() (*string, bool)`

GetTaskUuidOk returns a tuple with the TaskUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskUuid

`func (o *SecurityPlanningDataCollectorUpgradeStatus) SetTaskUuid(v string)`

SetTaskUuid sets TaskUuid field to given value.

### HasTaskUuid

`func (o *SecurityPlanningDataCollectorUpgradeStatus) HasTaskUuid() bool`

HasTaskUuid returns a boolean if a field has been set.

### GetIsUpgradeAvailable

`func (o *SecurityPlanningDataCollectorUpgradeStatus) GetIsUpgradeAvailable() bool`

GetIsUpgradeAvailable returns the IsUpgradeAvailable field if non-nil, zero value otherwise.

### GetIsUpgradeAvailableOk

`func (o *SecurityPlanningDataCollectorUpgradeStatus) GetIsUpgradeAvailableOk() (*bool, bool)`

GetIsUpgradeAvailableOk returns a tuple with the IsUpgradeAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUpgradeAvailable

`func (o *SecurityPlanningDataCollectorUpgradeStatus) SetIsUpgradeAvailable(v bool)`

SetIsUpgradeAvailable sets IsUpgradeAvailable field to given value.

### HasIsUpgradeAvailable

`func (o *SecurityPlanningDataCollectorUpgradeStatus) HasIsUpgradeAvailable() bool`

HasIsUpgradeAvailable returns a boolean if a field has been set.

### GetLatestVersion

`func (o *SecurityPlanningDataCollectorUpgradeStatus) GetLatestVersion() string`

GetLatestVersion returns the LatestVersion field if non-nil, zero value otherwise.

### GetLatestVersionOk

`func (o *SecurityPlanningDataCollectorUpgradeStatus) GetLatestVersionOk() (*string, bool)`

GetLatestVersionOk returns a tuple with the LatestVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestVersion

`func (o *SecurityPlanningDataCollectorUpgradeStatus) SetLatestVersion(v string)`

SetLatestVersion sets LatestVersion field to given value.

### HasLatestVersion

`func (o *SecurityPlanningDataCollectorUpgradeStatus) HasLatestVersion() bool`

HasLatestVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



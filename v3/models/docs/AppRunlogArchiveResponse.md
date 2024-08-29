# AppRunlogArchiveResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Archive file name. | [optional] 
**StartTime** | Pointer to **time.Time** | From time for archive. | [optional] 
**CreationTime** | Pointer to **time.Time** | Creation time of archive. | [optional] 
**IsAvailable** | Pointer to **bool** | Archive available flag. | [optional] 
**EndTime** | Pointer to **time.Time** | Till time for archive. | [optional] 
**UUID** | Pointer to **string** | Archive file uuid. | [optional] 

## Methods

### NewAppRunlogArchiveResponse

`func NewAppRunlogArchiveResponse() *AppRunlogArchiveResponse`

NewAppRunlogArchiveResponse instantiates a new AppRunlogArchiveResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppRunlogArchiveResponseWithDefaults

`func NewAppRunlogArchiveResponseWithDefaults() *AppRunlogArchiveResponse`

NewAppRunlogArchiveResponseWithDefaults instantiates a new AppRunlogArchiveResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AppRunlogArchiveResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppRunlogArchiveResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppRunlogArchiveResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AppRunlogArchiveResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStartTime

`func (o *AppRunlogArchiveResponse) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *AppRunlogArchiveResponse) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *AppRunlogArchiveResponse) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *AppRunlogArchiveResponse) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetCreationTime

`func (o *AppRunlogArchiveResponse) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *AppRunlogArchiveResponse) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *AppRunlogArchiveResponse) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *AppRunlogArchiveResponse) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetIsAvailable

`func (o *AppRunlogArchiveResponse) GetIsAvailable() bool`

GetIsAvailable returns the IsAvailable field if non-nil, zero value otherwise.

### GetIsAvailableOk

`func (o *AppRunlogArchiveResponse) GetIsAvailableOk() (*bool, bool)`

GetIsAvailableOk returns a tuple with the IsAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAvailable

`func (o *AppRunlogArchiveResponse) SetIsAvailable(v bool)`

SetIsAvailable sets IsAvailable field to given value.

### HasIsAvailable

`func (o *AppRunlogArchiveResponse) HasIsAvailable() bool`

HasIsAvailable returns a boolean if a field has been set.

### GetEndTime

`func (o *AppRunlogArchiveResponse) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *AppRunlogArchiveResponse) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *AppRunlogArchiveResponse) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *AppRunlogArchiveResponse) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetUUID

`func (o *AppRunlogArchiveResponse) GetUUID() string`

GetUUID returns the UUID field if non-nil, zero value otherwise.

### GetUUIDOk

`func (o *AppRunlogArchiveResponse) GetUUIDOk() (*string, bool)`

GetUUIDOk returns a tuple with the UUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUUID

`func (o *AppRunlogArchiveResponse) SetUUID(v string)`

SetUUID sets UUID field to given value.

### HasUUID

`func (o *AppRunlogArchiveResponse) HasUUID() bool`

HasUUID returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



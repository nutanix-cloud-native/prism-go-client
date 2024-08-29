# AlertState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsTrue** | Pointer to **bool** | A boolean status for acknowledgement or resolution | [optional] 
**IsAutoResolved** | Pointer to **bool** | It is automatically resolved or user manually resolved.  Currently, it is limtied to the resolution only.  | [optional] 
**User** | Pointer to **string** | Name of the user who change this alert status. | [optional] 
**Time** | Pointer to **time.Time** | The time the alert status was changed. | [optional] 

## Methods

### NewAlertState

`func NewAlertState() *AlertState`

NewAlertState instantiates a new AlertState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertStateWithDefaults

`func NewAlertStateWithDefaults() *AlertState`

NewAlertStateWithDefaults instantiates a new AlertState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsTrue

`func (o *AlertState) GetIsTrue() bool`

GetIsTrue returns the IsTrue field if non-nil, zero value otherwise.

### GetIsTrueOk

`func (o *AlertState) GetIsTrueOk() (*bool, bool)`

GetIsTrueOk returns a tuple with the IsTrue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTrue

`func (o *AlertState) SetIsTrue(v bool)`

SetIsTrue sets IsTrue field to given value.

### HasIsTrue

`func (o *AlertState) HasIsTrue() bool`

HasIsTrue returns a boolean if a field has been set.

### GetIsAutoResolved

`func (o *AlertState) GetIsAutoResolved() bool`

GetIsAutoResolved returns the IsAutoResolved field if non-nil, zero value otherwise.

### GetIsAutoResolvedOk

`func (o *AlertState) GetIsAutoResolvedOk() (*bool, bool)`

GetIsAutoResolvedOk returns a tuple with the IsAutoResolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutoResolved

`func (o *AlertState) SetIsAutoResolved(v bool)`

SetIsAutoResolved sets IsAutoResolved field to given value.

### HasIsAutoResolved

`func (o *AlertState) HasIsAutoResolved() bool`

HasIsAutoResolved returns a boolean if a field has been set.

### GetUser

`func (o *AlertState) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *AlertState) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *AlertState) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *AlertState) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetTime

`func (o *AlertState) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *AlertState) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *AlertState) SetTime(v time.Time)`

SetTime sets Time field to given value.

### HasTime

`func (o *AlertState) HasTime() bool`

HasTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



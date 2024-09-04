# AlertNotificationInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NotificationType** | **string** | The notification definition for this type of alerts. | 
**SourceEntity** | Pointer to [**Reference**](Reference.md) |  | [optional] 
**Severity** | **string** | Alert severity | 
**Parameters** | Pointer to [**map[string]ParamValue**](ParamValue.md) | Alert notification type specific parameters. | [optional] 

## Methods

### NewAlertNotificationInput

`func NewAlertNotificationInput(notificationType string, severity string, ) *AlertNotificationInput`

NewAlertNotificationInput instantiates a new AlertNotificationInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertNotificationInputWithDefaults

`func NewAlertNotificationInputWithDefaults() *AlertNotificationInput`

NewAlertNotificationInputWithDefaults instantiates a new AlertNotificationInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotificationType

`func (o *AlertNotificationInput) GetNotificationType() string`

GetNotificationType returns the NotificationType field if non-nil, zero value otherwise.

### GetNotificationTypeOk

`func (o *AlertNotificationInput) GetNotificationTypeOk() (*string, bool)`

GetNotificationTypeOk returns a tuple with the NotificationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationType

`func (o *AlertNotificationInput) SetNotificationType(v string)`

SetNotificationType sets NotificationType field to given value.


### GetSourceEntity

`func (o *AlertNotificationInput) GetSourceEntity() Reference`

GetSourceEntity returns the SourceEntity field if non-nil, zero value otherwise.

### GetSourceEntityOk

`func (o *AlertNotificationInput) GetSourceEntityOk() (*Reference, bool)`

GetSourceEntityOk returns a tuple with the SourceEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceEntity

`func (o *AlertNotificationInput) SetSourceEntity(v Reference)`

SetSourceEntity sets SourceEntity field to given value.

### HasSourceEntity

`func (o *AlertNotificationInput) HasSourceEntity() bool`

HasSourceEntity returns a boolean if a field has been set.

### GetSeverity

`func (o *AlertNotificationInput) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *AlertNotificationInput) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *AlertNotificationInput) SetSeverity(v string)`

SetSeverity sets Severity field to given value.


### GetParameters

`func (o *AlertNotificationInput) GetParameters() map[string]ParamValue`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *AlertNotificationInput) GetParametersOk() (*map[string]ParamValue, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *AlertNotificationInput) SetParameters(v map[string]ParamValue)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *AlertNotificationInput) HasParameters() bool`

HasParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



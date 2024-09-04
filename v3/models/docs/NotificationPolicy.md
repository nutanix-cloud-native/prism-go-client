# NotificationPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmailConfig** | Pointer to [**EmailConfig**](EmailConfig.md) |  | [optional] 

## Methods

### NewNotificationPolicy

`func NewNotificationPolicy() *NotificationPolicy`

NewNotificationPolicy instantiates a new NotificationPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationPolicyWithDefaults

`func NewNotificationPolicyWithDefaults() *NotificationPolicy`

NewNotificationPolicyWithDefaults instantiates a new NotificationPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmailConfig

`func (o *NotificationPolicy) GetEmailConfig() EmailConfig`

GetEmailConfig returns the EmailConfig field if non-nil, zero value otherwise.

### GetEmailConfigOk

`func (o *NotificationPolicy) GetEmailConfigOk() (*EmailConfig, bool)`

GetEmailConfigOk returns a tuple with the EmailConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailConfig

`func (o *NotificationPolicy) SetEmailConfig(v EmailConfig)`

SetEmailConfig sets EmailConfig field to given value.

### HasEmailConfig

`func (o *NotificationPolicy) HasEmailConfig() bool`

HasEmailConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



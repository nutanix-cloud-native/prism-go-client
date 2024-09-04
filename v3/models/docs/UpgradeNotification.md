# UpgradeNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProceedMessage** | Pointer to **string** |  | [optional] 
**CancelMessage** | Pointer to **string** |  | [optional] 
**DescriptionItemList** | Pointer to [**[]DescriptionItem**](DescriptionItem.md) | Description items of the notification | [optional] 
**Title** | Pointer to **string** | Title of the notification | [optional] 

## Methods

### NewUpgradeNotification

`func NewUpgradeNotification() *UpgradeNotification`

NewUpgradeNotification instantiates a new UpgradeNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpgradeNotificationWithDefaults

`func NewUpgradeNotificationWithDefaults() *UpgradeNotification`

NewUpgradeNotificationWithDefaults instantiates a new UpgradeNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProceedMessage

`func (o *UpgradeNotification) GetProceedMessage() string`

GetProceedMessage returns the ProceedMessage field if non-nil, zero value otherwise.

### GetProceedMessageOk

`func (o *UpgradeNotification) GetProceedMessageOk() (*string, bool)`

GetProceedMessageOk returns a tuple with the ProceedMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProceedMessage

`func (o *UpgradeNotification) SetProceedMessage(v string)`

SetProceedMessage sets ProceedMessage field to given value.

### HasProceedMessage

`func (o *UpgradeNotification) HasProceedMessage() bool`

HasProceedMessage returns a boolean if a field has been set.

### GetCancelMessage

`func (o *UpgradeNotification) GetCancelMessage() string`

GetCancelMessage returns the CancelMessage field if non-nil, zero value otherwise.

### GetCancelMessageOk

`func (o *UpgradeNotification) GetCancelMessageOk() (*string, bool)`

GetCancelMessageOk returns a tuple with the CancelMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelMessage

`func (o *UpgradeNotification) SetCancelMessage(v string)`

SetCancelMessage sets CancelMessage field to given value.

### HasCancelMessage

`func (o *UpgradeNotification) HasCancelMessage() bool`

HasCancelMessage returns a boolean if a field has been set.

### GetDescriptionItemList

`func (o *UpgradeNotification) GetDescriptionItemList() []DescriptionItem`

GetDescriptionItemList returns the DescriptionItemList field if non-nil, zero value otherwise.

### GetDescriptionItemListOk

`func (o *UpgradeNotification) GetDescriptionItemListOk() (*[]DescriptionItem, bool)`

GetDescriptionItemListOk returns a tuple with the DescriptionItemList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionItemList

`func (o *UpgradeNotification) SetDescriptionItemList(v []DescriptionItem)`

SetDescriptionItemList sets DescriptionItemList field to given value.

### HasDescriptionItemList

`func (o *UpgradeNotification) HasDescriptionItemList() bool`

HasDescriptionItemList returns a boolean if a field has been set.

### GetTitle

`func (o *UpgradeNotification) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *UpgradeNotification) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *UpgradeNotification) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *UpgradeNotification) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



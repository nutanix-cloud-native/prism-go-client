# AppCostResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**MessageList** | [**[]MessageResource**](MessageResource.md) | Message list | 
**Interval** | **string** |  | 
**Value** | **string** |  | 
**Name** | **string** |  | 
**State** | **string** |  | 
**Type** | **string** |  | 
**UUID** | **string** |  | 

## Methods

### NewAppCostResponse

`func NewAppCostResponse(messageList []MessageResource, interval string, value string, name string, state string, type_ string, uUID string, ) *AppCostResponse`

NewAppCostResponse instantiates a new AppCostResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppCostResponseWithDefaults

`func NewAppCostResponseWithDefaults() *AppCostResponse`

NewAppCostResponseWithDefaults instantiates a new AppCostResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *AppCostResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AppCostResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AppCostResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AppCostResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMessageList

`func (o *AppCostResponse) GetMessageList() []MessageResource`

GetMessageList returns the MessageList field if non-nil, zero value otherwise.

### GetMessageListOk

`func (o *AppCostResponse) GetMessageListOk() (*[]MessageResource, bool)`

GetMessageListOk returns a tuple with the MessageList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageList

`func (o *AppCostResponse) SetMessageList(v []MessageResource)`

SetMessageList sets MessageList field to given value.


### GetInterval

`func (o *AppCostResponse) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *AppCostResponse) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *AppCostResponse) SetInterval(v string)`

SetInterval sets Interval field to given value.


### GetValue

`func (o *AppCostResponse) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *AppCostResponse) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *AppCostResponse) SetValue(v string)`

SetValue sets Value field to given value.


### GetName

`func (o *AppCostResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppCostResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppCostResponse) SetName(v string)`

SetName sets Name field to given value.


### GetState

`func (o *AppCostResponse) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AppCostResponse) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AppCostResponse) SetState(v string)`

SetState sets State field to given value.


### GetType

`func (o *AppCostResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AppCostResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AppCostResponse) SetType(v string)`

SetType sets Type field to given value.


### GetUUID

`func (o *AppCostResponse) GetUUID() string`

GetUUID returns the UUID field if non-nil, zero value otherwise.

### GetUUIDOk

`func (o *AppCostResponse) GetUUIDOk() (*string, bool)`

GetUUIDOk returns a tuple with the UUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUUID

`func (o *AppCostResponse) SetUUID(v string)`

SetUUID sets UUID field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



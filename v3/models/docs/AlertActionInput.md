# AlertActionInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlertUuidList** | Pointer to **[]string** | A list of alert UUIDs to be acknowledged or resolved.  If the list is empty, that means resolve all alerts in the system.  | [optional] 

## Methods

### NewAlertActionInput

`func NewAlertActionInput() *AlertActionInput`

NewAlertActionInput instantiates a new AlertActionInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertActionInputWithDefaults

`func NewAlertActionInputWithDefaults() *AlertActionInput`

NewAlertActionInputWithDefaults instantiates a new AlertActionInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlertUuidList

`func (o *AlertActionInput) GetAlertUuidList() []string`

GetAlertUuidList returns the AlertUuidList field if non-nil, zero value otherwise.

### GetAlertUuidListOk

`func (o *AlertActionInput) GetAlertUuidListOk() (*[]string, bool)`

GetAlertUuidListOk returns a tuple with the AlertUuidList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertUuidList

`func (o *AlertActionInput) SetAlertUuidList(v []string)`

SetAlertUuidList sets AlertUuidList field to given value.

### HasAlertUuidList

`func (o *AlertActionInput) HasAlertUuidList() bool`

HasAlertUuidList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



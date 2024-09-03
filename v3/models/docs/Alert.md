# Alert

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Resources** | [**AlertResources**](AlertResources.md) |  | 

## Methods

### NewAlert

`func NewAlert(resources AlertResources, ) *Alert`

NewAlert instantiates a new Alert object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertWithDefaults

`func NewAlertWithDefaults() *Alert`

NewAlertWithDefaults instantiates a new Alert object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResources

`func (o *Alert) GetResources() AlertResources`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *Alert) GetResourcesOk() (*AlertResources, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *Alert) SetResources(v AlertResources)`

SetResources sets Resources field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



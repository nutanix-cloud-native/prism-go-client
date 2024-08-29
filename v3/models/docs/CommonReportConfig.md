# CommonReportConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the common report config. | 
**Resources** | [**CommonReportConfigResources1**](CommonReportConfigResources1.md) |  | 

## Methods

### NewCommonReportConfig

`func NewCommonReportConfig(name string, resources CommonReportConfigResources1, ) *CommonReportConfig`

NewCommonReportConfig instantiates a new CommonReportConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonReportConfigWithDefaults

`func NewCommonReportConfigWithDefaults() *CommonReportConfig`

NewCommonReportConfigWithDefaults instantiates a new CommonReportConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CommonReportConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CommonReportConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CommonReportConfig) SetName(v string)`

SetName sets Name field to given value.


### GetResources

`func (o *CommonReportConfig) GetResources() CommonReportConfigResources1`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *CommonReportConfig) GetResourcesOk() (*CommonReportConfigResources1, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *CommonReportConfig) SetResources(v CommonReportConfigResources1)`

SetResources sets Resources field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



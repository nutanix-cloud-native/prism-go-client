# ReportConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the report config. | 
**Resources** | [**ReportConfigResources**](ReportConfigResources.md) |  | 

## Methods

### NewReportConfig

`func NewReportConfig(name string, resources ReportConfigResources, ) *ReportConfig`

NewReportConfig instantiates a new ReportConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportConfigWithDefaults

`func NewReportConfigWithDefaults() *ReportConfig`

NewReportConfigWithDefaults instantiates a new ReportConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ReportConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ReportConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ReportConfig) SetName(v string)`

SetName sets Name field to given value.


### GetResources

`func (o *ReportConfig) GetResources() ReportConfigResources`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *ReportConfig) GetResourcesOk() (*ReportConfigResources, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *ReportConfig) SetResources(v ReportConfigResources)`

SetResources sets Resources field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



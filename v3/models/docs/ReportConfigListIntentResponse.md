# ReportConfigListIntentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entities** | Pointer to [**[]ReportConfigIntentResource**](ReportConfigIntentResource.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**ReportConfigListMetadataOutput**](ReportConfigListMetadataOutput.md) |  | 

## Methods

### NewReportConfigListIntentResponse

`func NewReportConfigListIntentResponse(apiVersion string, metadata ReportConfigListMetadataOutput, ) *ReportConfigListIntentResponse`

NewReportConfigListIntentResponse instantiates a new ReportConfigListIntentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportConfigListIntentResponseWithDefaults

`func NewReportConfigListIntentResponseWithDefaults() *ReportConfigListIntentResponse`

NewReportConfigListIntentResponseWithDefaults instantiates a new ReportConfigListIntentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntities

`func (o *ReportConfigListIntentResponse) GetEntities() []ReportConfigIntentResource`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *ReportConfigListIntentResponse) GetEntitiesOk() (*[]ReportConfigIntentResource, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *ReportConfigListIntentResponse) SetEntities(v []ReportConfigIntentResource)`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *ReportConfigListIntentResponse) HasEntities() bool`

HasEntities returns a boolean if a field has been set.

### GetApiVersion

`func (o *ReportConfigListIntentResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ReportConfigListIntentResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ReportConfigListIntentResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *ReportConfigListIntentResponse) GetMetadata() ReportConfigListMetadataOutput`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ReportConfigListIntentResponse) GetMetadataOk() (*ReportConfigListMetadataOutput, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ReportConfigListIntentResponse) SetMetadata(v ReportConfigListMetadataOutput)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



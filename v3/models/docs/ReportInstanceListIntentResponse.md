# ReportInstanceListIntentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entities** | Pointer to [**[]ReportInstanceIntentResource**](ReportInstanceIntentResource.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**ReportInstanceListMetadataOutput**](ReportInstanceListMetadataOutput.md) |  | 

## Methods

### NewReportInstanceListIntentResponse

`func NewReportInstanceListIntentResponse(apiVersion string, metadata ReportInstanceListMetadataOutput, ) *ReportInstanceListIntentResponse`

NewReportInstanceListIntentResponse instantiates a new ReportInstanceListIntentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportInstanceListIntentResponseWithDefaults

`func NewReportInstanceListIntentResponseWithDefaults() *ReportInstanceListIntentResponse`

NewReportInstanceListIntentResponseWithDefaults instantiates a new ReportInstanceListIntentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntities

`func (o *ReportInstanceListIntentResponse) GetEntities() []ReportInstanceIntentResource`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *ReportInstanceListIntentResponse) GetEntitiesOk() (*[]ReportInstanceIntentResource, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *ReportInstanceListIntentResponse) SetEntities(v []ReportInstanceIntentResource)`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *ReportInstanceListIntentResponse) HasEntities() bool`

HasEntities returns a boolean if a field has been set.

### GetApiVersion

`func (o *ReportInstanceListIntentResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ReportInstanceListIntentResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ReportInstanceListIntentResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *ReportInstanceListIntentResponse) GetMetadata() ReportInstanceListMetadataOutput`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ReportInstanceListIntentResponse) GetMetadataOk() (*ReportInstanceListMetadataOutput, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ReportInstanceListIntentResponse) SetMetadata(v ReportInstanceListMetadataOutput)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



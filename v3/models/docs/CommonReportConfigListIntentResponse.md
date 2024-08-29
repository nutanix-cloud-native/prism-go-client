# CommonReportConfigListIntentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entities** | Pointer to [**[]CommonReportConfigIntentResource**](CommonReportConfigIntentResource.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**CommonReportConfigListMetadataOutput**](CommonReportConfigListMetadataOutput.md) |  | 

## Methods

### NewCommonReportConfigListIntentResponse

`func NewCommonReportConfigListIntentResponse(apiVersion string, metadata CommonReportConfigListMetadataOutput, ) *CommonReportConfigListIntentResponse`

NewCommonReportConfigListIntentResponse instantiates a new CommonReportConfigListIntentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonReportConfigListIntentResponseWithDefaults

`func NewCommonReportConfigListIntentResponseWithDefaults() *CommonReportConfigListIntentResponse`

NewCommonReportConfigListIntentResponseWithDefaults instantiates a new CommonReportConfigListIntentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntities

`func (o *CommonReportConfigListIntentResponse) GetEntities() []CommonReportConfigIntentResource`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *CommonReportConfigListIntentResponse) GetEntitiesOk() (*[]CommonReportConfigIntentResource, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *CommonReportConfigListIntentResponse) SetEntities(v []CommonReportConfigIntentResource)`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *CommonReportConfigListIntentResponse) HasEntities() bool`

HasEntities returns a boolean if a field has been set.

### GetApiVersion

`func (o *CommonReportConfigListIntentResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *CommonReportConfigListIntentResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *CommonReportConfigListIntentResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *CommonReportConfigListIntentResponse) GetMetadata() CommonReportConfigListMetadataOutput`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CommonReportConfigListIntentResponse) GetMetadataOk() (*CommonReportConfigListMetadataOutput, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CommonReportConfigListIntentResponse) SetMetadata(v CommonReportConfigListMetadataOutput)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



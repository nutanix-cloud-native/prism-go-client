# CommonReportConfigIntentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**CommonReportConfigDefStatus**](CommonReportConfigDefStatus.md) |  | [optional] 
**Spec** | Pointer to [**CommonReportConfig**](CommonReportConfig.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**CommonReportConfigMetadata**](CommonReportConfigMetadata.md) |  | 

## Methods

### NewCommonReportConfigIntentResponse

`func NewCommonReportConfigIntentResponse(apiVersion string, metadata CommonReportConfigMetadata, ) *CommonReportConfigIntentResponse`

NewCommonReportConfigIntentResponse instantiates a new CommonReportConfigIntentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonReportConfigIntentResponseWithDefaults

`func NewCommonReportConfigIntentResponseWithDefaults() *CommonReportConfigIntentResponse`

NewCommonReportConfigIntentResponseWithDefaults instantiates a new CommonReportConfigIntentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *CommonReportConfigIntentResponse) GetStatus() CommonReportConfigDefStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CommonReportConfigIntentResponse) GetStatusOk() (*CommonReportConfigDefStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CommonReportConfigIntentResponse) SetStatus(v CommonReportConfigDefStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CommonReportConfigIntentResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSpec

`func (o *CommonReportConfigIntentResponse) GetSpec() CommonReportConfig`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *CommonReportConfigIntentResponse) GetSpecOk() (*CommonReportConfig, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *CommonReportConfigIntentResponse) SetSpec(v CommonReportConfig)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *CommonReportConfigIntentResponse) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetApiVersion

`func (o *CommonReportConfigIntentResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *CommonReportConfigIntentResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *CommonReportConfigIntentResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *CommonReportConfigIntentResponse) GetMetadata() CommonReportConfigMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CommonReportConfigIntentResponse) GetMetadataOk() (*CommonReportConfigMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CommonReportConfigIntentResponse) SetMetadata(v CommonReportConfigMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



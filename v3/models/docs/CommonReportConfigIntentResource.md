# CommonReportConfigIntentResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**CommonReportConfigDefStatus**](CommonReportConfigDefStatus.md) |  | [optional] 
**Spec** | Pointer to [**CommonReportConfig**](CommonReportConfig.md) |  | [optional] 
**ApiVersion** | Pointer to **string** | API Version of the Nutanix v3 API framework. | [optional] [readonly] [default to "3.1.0"]
**Metadata** | [**CommonReportConfigMetadata**](CommonReportConfigMetadata.md) |  | 

## Methods

### NewCommonReportConfigIntentResource

`func NewCommonReportConfigIntentResource(metadata CommonReportConfigMetadata, ) *CommonReportConfigIntentResource`

NewCommonReportConfigIntentResource instantiates a new CommonReportConfigIntentResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonReportConfigIntentResourceWithDefaults

`func NewCommonReportConfigIntentResourceWithDefaults() *CommonReportConfigIntentResource`

NewCommonReportConfigIntentResourceWithDefaults instantiates a new CommonReportConfigIntentResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *CommonReportConfigIntentResource) GetStatus() CommonReportConfigDefStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CommonReportConfigIntentResource) GetStatusOk() (*CommonReportConfigDefStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CommonReportConfigIntentResource) SetStatus(v CommonReportConfigDefStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CommonReportConfigIntentResource) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSpec

`func (o *CommonReportConfigIntentResource) GetSpec() CommonReportConfig`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *CommonReportConfigIntentResource) GetSpecOk() (*CommonReportConfig, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *CommonReportConfigIntentResource) SetSpec(v CommonReportConfig)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *CommonReportConfigIntentResource) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetApiVersion

`func (o *CommonReportConfigIntentResource) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *CommonReportConfigIntentResource) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *CommonReportConfigIntentResource) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *CommonReportConfigIntentResource) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetMetadata

`func (o *CommonReportConfigIntentResource) GetMetadata() CommonReportConfigMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CommonReportConfigIntentResource) GetMetadataOk() (*CommonReportConfigMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CommonReportConfigIntentResource) SetMetadata(v CommonReportConfigMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



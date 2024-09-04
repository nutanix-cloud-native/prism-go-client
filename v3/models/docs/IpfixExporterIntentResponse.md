# IpfixExporterIntentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**IpfixExporterDefStatus**](IpfixExporterDefStatus.md) |  | [optional] 
**Spec** | Pointer to [**IpfixExporter**](IpfixExporter.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**IpfixExporterMetadata**](IpfixExporterMetadata.md) |  | 

## Methods

### NewIpfixExporterIntentResponse

`func NewIpfixExporterIntentResponse(apiVersion string, metadata IpfixExporterMetadata, ) *IpfixExporterIntentResponse`

NewIpfixExporterIntentResponse instantiates a new IpfixExporterIntentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpfixExporterIntentResponseWithDefaults

`func NewIpfixExporterIntentResponseWithDefaults() *IpfixExporterIntentResponse`

NewIpfixExporterIntentResponseWithDefaults instantiates a new IpfixExporterIntentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *IpfixExporterIntentResponse) GetStatus() IpfixExporterDefStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IpfixExporterIntentResponse) GetStatusOk() (*IpfixExporterDefStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IpfixExporterIntentResponse) SetStatus(v IpfixExporterDefStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *IpfixExporterIntentResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSpec

`func (o *IpfixExporterIntentResponse) GetSpec() IpfixExporter`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *IpfixExporterIntentResponse) GetSpecOk() (*IpfixExporter, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *IpfixExporterIntentResponse) SetSpec(v IpfixExporter)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *IpfixExporterIntentResponse) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetApiVersion

`func (o *IpfixExporterIntentResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *IpfixExporterIntentResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *IpfixExporterIntentResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *IpfixExporterIntentResponse) GetMetadata() IpfixExporterMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *IpfixExporterIntentResponse) GetMetadataOk() (*IpfixExporterMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *IpfixExporterIntentResponse) SetMetadata(v IpfixExporterMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



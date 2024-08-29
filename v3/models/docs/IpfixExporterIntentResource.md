# IpfixExporterIntentResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**IpfixExporterDefStatus**](IpfixExporterDefStatus.md) |  | [optional] 
**Spec** | Pointer to [**IpfixExporter**](IpfixExporter.md) |  | [optional] 
**ApiVersion** | Pointer to **string** | API Version of the Nutanix v3 API framework. | [optional] [readonly] [default to "3.1.0"]
**Metadata** | [**IpfixExporterMetadata**](IpfixExporterMetadata.md) |  | 

## Methods

### NewIpfixExporterIntentResource

`func NewIpfixExporterIntentResource(metadata IpfixExporterMetadata, ) *IpfixExporterIntentResource`

NewIpfixExporterIntentResource instantiates a new IpfixExporterIntentResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpfixExporterIntentResourceWithDefaults

`func NewIpfixExporterIntentResourceWithDefaults() *IpfixExporterIntentResource`

NewIpfixExporterIntentResourceWithDefaults instantiates a new IpfixExporterIntentResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *IpfixExporterIntentResource) GetStatus() IpfixExporterDefStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IpfixExporterIntentResource) GetStatusOk() (*IpfixExporterDefStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IpfixExporterIntentResource) SetStatus(v IpfixExporterDefStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *IpfixExporterIntentResource) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSpec

`func (o *IpfixExporterIntentResource) GetSpec() IpfixExporter`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *IpfixExporterIntentResource) GetSpecOk() (*IpfixExporter, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *IpfixExporterIntentResource) SetSpec(v IpfixExporter)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *IpfixExporterIntentResource) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetApiVersion

`func (o *IpfixExporterIntentResource) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *IpfixExporterIntentResource) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *IpfixExporterIntentResource) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *IpfixExporterIntentResource) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetMetadata

`func (o *IpfixExporterIntentResource) GetMetadata() IpfixExporterMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *IpfixExporterIntentResource) GetMetadataOk() (*IpfixExporterMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *IpfixExporterIntentResource) SetMetadata(v IpfixExporterMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



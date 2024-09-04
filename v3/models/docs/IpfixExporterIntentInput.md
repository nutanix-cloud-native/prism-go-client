# IpfixExporterIntentInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Spec** | [**IpfixExporter**](IpfixExporter.md) |  | 
**ApiVersion** | Pointer to **string** | API Version of the Nutanix v3 API framework. | [optional] [readonly] [default to "3.1.0"]
**Metadata** | [**IpfixExporterMetadata**](IpfixExporterMetadata.md) |  | 

## Methods

### NewIpfixExporterIntentInput

`func NewIpfixExporterIntentInput(spec IpfixExporter, metadata IpfixExporterMetadata, ) *IpfixExporterIntentInput`

NewIpfixExporterIntentInput instantiates a new IpfixExporterIntentInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpfixExporterIntentInputWithDefaults

`func NewIpfixExporterIntentInputWithDefaults() *IpfixExporterIntentInput`

NewIpfixExporterIntentInputWithDefaults instantiates a new IpfixExporterIntentInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpec

`func (o *IpfixExporterIntentInput) GetSpec() IpfixExporter`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *IpfixExporterIntentInput) GetSpecOk() (*IpfixExporter, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *IpfixExporterIntentInput) SetSpec(v IpfixExporter)`

SetSpec sets Spec field to given value.


### GetApiVersion

`func (o *IpfixExporterIntentInput) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *IpfixExporterIntentInput) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *IpfixExporterIntentInput) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *IpfixExporterIntentInput) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetMetadata

`func (o *IpfixExporterIntentInput) GetMetadata() IpfixExporterMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *IpfixExporterIntentInput) GetMetadataOk() (*IpfixExporterMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *IpfixExporterIntentInput) SetMetadata(v IpfixExporterMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



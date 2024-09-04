# IpfixExporterListIntentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entities** | Pointer to [**[]IpfixExporterIntentResource**](IpfixExporterIntentResource.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**IpfixExporterListMetadataOutput**](IpfixExporterListMetadataOutput.md) |  | 

## Methods

### NewIpfixExporterListIntentResponse

`func NewIpfixExporterListIntentResponse(apiVersion string, metadata IpfixExporterListMetadataOutput, ) *IpfixExporterListIntentResponse`

NewIpfixExporterListIntentResponse instantiates a new IpfixExporterListIntentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpfixExporterListIntentResponseWithDefaults

`func NewIpfixExporterListIntentResponseWithDefaults() *IpfixExporterListIntentResponse`

NewIpfixExporterListIntentResponseWithDefaults instantiates a new IpfixExporterListIntentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntities

`func (o *IpfixExporterListIntentResponse) GetEntities() []IpfixExporterIntentResource`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *IpfixExporterListIntentResponse) GetEntitiesOk() (*[]IpfixExporterIntentResource, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *IpfixExporterListIntentResponse) SetEntities(v []IpfixExporterIntentResource)`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *IpfixExporterListIntentResponse) HasEntities() bool`

HasEntities returns a boolean if a field has been set.

### GetApiVersion

`func (o *IpfixExporterListIntentResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *IpfixExporterListIntentResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *IpfixExporterListIntentResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *IpfixExporterListIntentResponse) GetMetadata() IpfixExporterListMetadataOutput`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *IpfixExporterListIntentResponse) GetMetadataOk() (*IpfixExporterListMetadataOutput, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *IpfixExporterListIntentResponse) SetMetadata(v IpfixExporterListMetadataOutput)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



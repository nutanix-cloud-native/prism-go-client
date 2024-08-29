# SoftwareListIntentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entities** | [**[]PortalSoftware**](PortalSoftware.md) |  | 
**ApiVersion** | Pointer to **string** | API Version of the Nutanix v3 API framework. | [optional] [readonly] [default to "3.1.0"]
**Metadata** | [**SoftwareListMetadataOutput**](SoftwareListMetadataOutput.md) |  | 

## Methods

### NewSoftwareListIntentResponse

`func NewSoftwareListIntentResponse(entities []PortalSoftware, metadata SoftwareListMetadataOutput, ) *SoftwareListIntentResponse`

NewSoftwareListIntentResponse instantiates a new SoftwareListIntentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSoftwareListIntentResponseWithDefaults

`func NewSoftwareListIntentResponseWithDefaults() *SoftwareListIntentResponse`

NewSoftwareListIntentResponseWithDefaults instantiates a new SoftwareListIntentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntities

`func (o *SoftwareListIntentResponse) GetEntities() []PortalSoftware`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *SoftwareListIntentResponse) GetEntitiesOk() (*[]PortalSoftware, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *SoftwareListIntentResponse) SetEntities(v []PortalSoftware)`

SetEntities sets Entities field to given value.


### GetApiVersion

`func (o *SoftwareListIntentResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *SoftwareListIntentResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *SoftwareListIntentResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *SoftwareListIntentResponse) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetMetadata

`func (o *SoftwareListIntentResponse) GetMetadata() SoftwareListMetadataOutput`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SoftwareListIntentResponse) GetMetadataOk() (*SoftwareListMetadataOutput, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SoftwareListIntentResponse) SetMetadata(v SoftwareListMetadataOutput)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



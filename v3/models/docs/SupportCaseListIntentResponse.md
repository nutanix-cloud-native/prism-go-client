# SupportCaseListIntentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entities** | Pointer to [**[]SupportCaseIntentResource**](SupportCaseIntentResource.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**SupportCaseListMetadataOutput**](SupportCaseListMetadataOutput.md) |  | 

## Methods

### NewSupportCaseListIntentResponse

`func NewSupportCaseListIntentResponse(apiVersion string, metadata SupportCaseListMetadataOutput, ) *SupportCaseListIntentResponse`

NewSupportCaseListIntentResponse instantiates a new SupportCaseListIntentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupportCaseListIntentResponseWithDefaults

`func NewSupportCaseListIntentResponseWithDefaults() *SupportCaseListIntentResponse`

NewSupportCaseListIntentResponseWithDefaults instantiates a new SupportCaseListIntentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntities

`func (o *SupportCaseListIntentResponse) GetEntities() []SupportCaseIntentResource`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *SupportCaseListIntentResponse) GetEntitiesOk() (*[]SupportCaseIntentResource, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *SupportCaseListIntentResponse) SetEntities(v []SupportCaseIntentResource)`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *SupportCaseListIntentResponse) HasEntities() bool`

HasEntities returns a boolean if a field has been set.

### GetApiVersion

`func (o *SupportCaseListIntentResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *SupportCaseListIntentResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *SupportCaseListIntentResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *SupportCaseListIntentResponse) GetMetadata() SupportCaseListMetadataOutput`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SupportCaseListIntentResponse) GetMetadataOk() (*SupportCaseListMetadataOutput, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SupportCaseListIntentResponse) SetMetadata(v SupportCaseListMetadataOutput)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



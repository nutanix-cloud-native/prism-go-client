# HostListIntentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entities** | Pointer to [**[]HostIntentResource**](HostIntentResource.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**HostListMetadataOutput**](HostListMetadataOutput.md) |  | 

## Methods

### NewHostListIntentResponse

`func NewHostListIntentResponse(apiVersion string, metadata HostListMetadataOutput, ) *HostListIntentResponse`

NewHostListIntentResponse instantiates a new HostListIntentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostListIntentResponseWithDefaults

`func NewHostListIntentResponseWithDefaults() *HostListIntentResponse`

NewHostListIntentResponseWithDefaults instantiates a new HostListIntentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntities

`func (o *HostListIntentResponse) GetEntities() []HostIntentResource`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *HostListIntentResponse) GetEntitiesOk() (*[]HostIntentResource, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *HostListIntentResponse) SetEntities(v []HostIntentResource)`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *HostListIntentResponse) HasEntities() bool`

HasEntities returns a boolean if a field has been set.

### GetApiVersion

`func (o *HostListIntentResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *HostListIntentResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *HostListIntentResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *HostListIntentResponse) GetMetadata() HostListMetadataOutput`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *HostListIntentResponse) GetMetadataOk() (*HostListMetadataOutput, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *HostListIntentResponse) SetMetadata(v HostListMetadataOutput)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



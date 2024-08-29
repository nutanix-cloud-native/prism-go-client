# DirectoryServiceListIntentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entities** | Pointer to [**[]DirectoryServiceIntentResource**](DirectoryServiceIntentResource.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**DirectoryServiceListMetadataOutput**](DirectoryServiceListMetadataOutput.md) |  | 

## Methods

### NewDirectoryServiceListIntentResponse

`func NewDirectoryServiceListIntentResponse(apiVersion string, metadata DirectoryServiceListMetadataOutput, ) *DirectoryServiceListIntentResponse`

NewDirectoryServiceListIntentResponse instantiates a new DirectoryServiceListIntentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDirectoryServiceListIntentResponseWithDefaults

`func NewDirectoryServiceListIntentResponseWithDefaults() *DirectoryServiceListIntentResponse`

NewDirectoryServiceListIntentResponseWithDefaults instantiates a new DirectoryServiceListIntentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntities

`func (o *DirectoryServiceListIntentResponse) GetEntities() []DirectoryServiceIntentResource`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *DirectoryServiceListIntentResponse) GetEntitiesOk() (*[]DirectoryServiceIntentResource, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *DirectoryServiceListIntentResponse) SetEntities(v []DirectoryServiceIntentResource)`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *DirectoryServiceListIntentResponse) HasEntities() bool`

HasEntities returns a boolean if a field has been set.

### GetApiVersion

`func (o *DirectoryServiceListIntentResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *DirectoryServiceListIntentResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *DirectoryServiceListIntentResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *DirectoryServiceListIntentResponse) GetMetadata() DirectoryServiceListMetadataOutput`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DirectoryServiceListIntentResponse) GetMetadataOk() (*DirectoryServiceListMetadataOutput, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DirectoryServiceListIntentResponse) SetMetadata(v DirectoryServiceListMetadataOutput)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



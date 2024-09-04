# PermissionListIntentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entities** | Pointer to [**[]PermissionIntentResource**](PermissionIntentResource.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**PermissionListMetadataOutput**](PermissionListMetadataOutput.md) |  | 

## Methods

### NewPermissionListIntentResponse

`func NewPermissionListIntentResponse(apiVersion string, metadata PermissionListMetadataOutput, ) *PermissionListIntentResponse`

NewPermissionListIntentResponse instantiates a new PermissionListIntentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPermissionListIntentResponseWithDefaults

`func NewPermissionListIntentResponseWithDefaults() *PermissionListIntentResponse`

NewPermissionListIntentResponseWithDefaults instantiates a new PermissionListIntentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntities

`func (o *PermissionListIntentResponse) GetEntities() []PermissionIntentResource`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *PermissionListIntentResponse) GetEntitiesOk() (*[]PermissionIntentResource, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *PermissionListIntentResponse) SetEntities(v []PermissionIntentResource)`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *PermissionListIntentResponse) HasEntities() bool`

HasEntities returns a boolean if a field has been set.

### GetApiVersion

`func (o *PermissionListIntentResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *PermissionListIntentResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *PermissionListIntentResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *PermissionListIntentResponse) GetMetadata() PermissionListMetadataOutput`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PermissionListIntentResponse) GetMetadataOk() (*PermissionListMetadataOutput, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PermissionListIntentResponse) SetMetadata(v PermissionListMetadataOutput)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



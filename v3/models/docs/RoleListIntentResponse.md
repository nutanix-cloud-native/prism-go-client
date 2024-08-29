# RoleListIntentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entities** | Pointer to [**[]RoleIntentResource**](RoleIntentResource.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**RoleListMetadataOutput**](RoleListMetadataOutput.md) |  | 

## Methods

### NewRoleListIntentResponse

`func NewRoleListIntentResponse(apiVersion string, metadata RoleListMetadataOutput, ) *RoleListIntentResponse`

NewRoleListIntentResponse instantiates a new RoleListIntentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleListIntentResponseWithDefaults

`func NewRoleListIntentResponseWithDefaults() *RoleListIntentResponse`

NewRoleListIntentResponseWithDefaults instantiates a new RoleListIntentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntities

`func (o *RoleListIntentResponse) GetEntities() []RoleIntentResource`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *RoleListIntentResponse) GetEntitiesOk() (*[]RoleIntentResource, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *RoleListIntentResponse) SetEntities(v []RoleIntentResource)`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *RoleListIntentResponse) HasEntities() bool`

HasEntities returns a boolean if a field has been set.

### GetApiVersion

`func (o *RoleListIntentResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *RoleListIntentResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *RoleListIntentResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *RoleListIntentResponse) GetMetadata() RoleListMetadataOutput`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *RoleListIntentResponse) GetMetadataOk() (*RoleListMetadataOutput, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *RoleListIntentResponse) SetMetadata(v RoleListMetadataOutput)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



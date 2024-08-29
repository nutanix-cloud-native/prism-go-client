# AwsRoleListIntentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entities** | Pointer to [**[]AwsRoleIntentResource**](AwsRoleIntentResource.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**AwsRoleListMetadataOutput**](AwsRoleListMetadataOutput.md) |  | 

## Methods

### NewAwsRoleListIntentResponse

`func NewAwsRoleListIntentResponse(apiVersion string, metadata AwsRoleListMetadataOutput, ) *AwsRoleListIntentResponse`

NewAwsRoleListIntentResponse instantiates a new AwsRoleListIntentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsRoleListIntentResponseWithDefaults

`func NewAwsRoleListIntentResponseWithDefaults() *AwsRoleListIntentResponse`

NewAwsRoleListIntentResponseWithDefaults instantiates a new AwsRoleListIntentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntities

`func (o *AwsRoleListIntentResponse) GetEntities() []AwsRoleIntentResource`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *AwsRoleListIntentResponse) GetEntitiesOk() (*[]AwsRoleIntentResource, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *AwsRoleListIntentResponse) SetEntities(v []AwsRoleIntentResource)`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *AwsRoleListIntentResponse) HasEntities() bool`

HasEntities returns a boolean if a field has been set.

### GetApiVersion

`func (o *AwsRoleListIntentResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *AwsRoleListIntentResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *AwsRoleListIntentResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *AwsRoleListIntentResponse) GetMetadata() AwsRoleListMetadataOutput`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AwsRoleListIntentResponse) GetMetadataOk() (*AwsRoleListMetadataOutput, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AwsRoleListIntentResponse) SetMetadata(v AwsRoleListMetadataOutput)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



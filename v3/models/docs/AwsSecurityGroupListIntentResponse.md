# AwsSecurityGroupListIntentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entities** | Pointer to [**[]AwsSecurityGroupIntentResource**](AwsSecurityGroupIntentResource.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**AwsSecurityGroupListMetadataOutput**](AwsSecurityGroupListMetadataOutput.md) |  | 

## Methods

### NewAwsSecurityGroupListIntentResponse

`func NewAwsSecurityGroupListIntentResponse(apiVersion string, metadata AwsSecurityGroupListMetadataOutput, ) *AwsSecurityGroupListIntentResponse`

NewAwsSecurityGroupListIntentResponse instantiates a new AwsSecurityGroupListIntentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsSecurityGroupListIntentResponseWithDefaults

`func NewAwsSecurityGroupListIntentResponseWithDefaults() *AwsSecurityGroupListIntentResponse`

NewAwsSecurityGroupListIntentResponseWithDefaults instantiates a new AwsSecurityGroupListIntentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntities

`func (o *AwsSecurityGroupListIntentResponse) GetEntities() []AwsSecurityGroupIntentResource`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *AwsSecurityGroupListIntentResponse) GetEntitiesOk() (*[]AwsSecurityGroupIntentResource, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *AwsSecurityGroupListIntentResponse) SetEntities(v []AwsSecurityGroupIntentResource)`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *AwsSecurityGroupListIntentResponse) HasEntities() bool`

HasEntities returns a boolean if a field has been set.

### GetApiVersion

`func (o *AwsSecurityGroupListIntentResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *AwsSecurityGroupListIntentResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *AwsSecurityGroupListIntentResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *AwsSecurityGroupListIntentResponse) GetMetadata() AwsSecurityGroupListMetadataOutput`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AwsSecurityGroupListIntentResponse) GetMetadataOk() (*AwsSecurityGroupListMetadataOutput, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AwsSecurityGroupListIntentResponse) SetMetadata(v AwsSecurityGroupListMetadataOutput)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



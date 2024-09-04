# SubnetListIntentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entities** | Pointer to [**[]SubnetIntentResource**](SubnetIntentResource.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**SubnetListMetadataOutput**](SubnetListMetadataOutput.md) |  | 

## Methods

### NewSubnetListIntentResponse

`func NewSubnetListIntentResponse(apiVersion string, metadata SubnetListMetadataOutput, ) *SubnetListIntentResponse`

NewSubnetListIntentResponse instantiates a new SubnetListIntentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubnetListIntentResponseWithDefaults

`func NewSubnetListIntentResponseWithDefaults() *SubnetListIntentResponse`

NewSubnetListIntentResponseWithDefaults instantiates a new SubnetListIntentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntities

`func (o *SubnetListIntentResponse) GetEntities() []SubnetIntentResource`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *SubnetListIntentResponse) GetEntitiesOk() (*[]SubnetIntentResource, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *SubnetListIntentResponse) SetEntities(v []SubnetIntentResource)`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *SubnetListIntentResponse) HasEntities() bool`

HasEntities returns a boolean if a field has been set.

### GetApiVersion

`func (o *SubnetListIntentResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *SubnetListIntentResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *SubnetListIntentResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *SubnetListIntentResponse) GetMetadata() SubnetListMetadataOutput`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SubnetListIntentResponse) GetMetadataOk() (*SubnetListMetadataOutput, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SubnetListIntentResponse) SetMetadata(v SubnetListMetadataOutput)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



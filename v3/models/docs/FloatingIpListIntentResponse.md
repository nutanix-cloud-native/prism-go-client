# FloatingIpListIntentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entities** | Pointer to [**[]FloatingIpIntentResource**](FloatingIpIntentResource.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**FloatingIpListMetadataOutput**](FloatingIpListMetadataOutput.md) |  | 

## Methods

### NewFloatingIpListIntentResponse

`func NewFloatingIpListIntentResponse(apiVersion string, metadata FloatingIpListMetadataOutput, ) *FloatingIpListIntentResponse`

NewFloatingIpListIntentResponse instantiates a new FloatingIpListIntentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFloatingIpListIntentResponseWithDefaults

`func NewFloatingIpListIntentResponseWithDefaults() *FloatingIpListIntentResponse`

NewFloatingIpListIntentResponseWithDefaults instantiates a new FloatingIpListIntentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntities

`func (o *FloatingIpListIntentResponse) GetEntities() []FloatingIpIntentResource`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *FloatingIpListIntentResponse) GetEntitiesOk() (*[]FloatingIpIntentResource, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *FloatingIpListIntentResponse) SetEntities(v []FloatingIpIntentResource)`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *FloatingIpListIntentResponse) HasEntities() bool`

HasEntities returns a boolean if a field has been set.

### GetApiVersion

`func (o *FloatingIpListIntentResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *FloatingIpListIntentResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *FloatingIpListIntentResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *FloatingIpListIntentResponse) GetMetadata() FloatingIpListMetadataOutput`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *FloatingIpListIntentResponse) GetMetadataOk() (*FloatingIpListMetadataOutput, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *FloatingIpListIntentResponse) SetMetadata(v FloatingIpListMetadataOutput)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



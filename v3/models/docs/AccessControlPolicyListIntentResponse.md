# AccessControlPolicyListIntentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entities** | Pointer to [**[]AccessControlPolicyIntentResource**](AccessControlPolicyIntentResource.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**AccessControlPolicyListMetadataOutput**](AccessControlPolicyListMetadataOutput.md) |  | 

## Methods

### NewAccessControlPolicyListIntentResponse

`func NewAccessControlPolicyListIntentResponse(apiVersion string, metadata AccessControlPolicyListMetadataOutput, ) *AccessControlPolicyListIntentResponse`

NewAccessControlPolicyListIntentResponse instantiates a new AccessControlPolicyListIntentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessControlPolicyListIntentResponseWithDefaults

`func NewAccessControlPolicyListIntentResponseWithDefaults() *AccessControlPolicyListIntentResponse`

NewAccessControlPolicyListIntentResponseWithDefaults instantiates a new AccessControlPolicyListIntentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntities

`func (o *AccessControlPolicyListIntentResponse) GetEntities() []AccessControlPolicyIntentResource`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *AccessControlPolicyListIntentResponse) GetEntitiesOk() (*[]AccessControlPolicyIntentResource, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *AccessControlPolicyListIntentResponse) SetEntities(v []AccessControlPolicyIntentResource)`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *AccessControlPolicyListIntentResponse) HasEntities() bool`

HasEntities returns a boolean if a field has been set.

### GetApiVersion

`func (o *AccessControlPolicyListIntentResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *AccessControlPolicyListIntentResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *AccessControlPolicyListIntentResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *AccessControlPolicyListIntentResponse) GetMetadata() AccessControlPolicyListMetadataOutput`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AccessControlPolicyListIntentResponse) GetMetadataOk() (*AccessControlPolicyListMetadataOutput, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AccessControlPolicyListIntentResponse) SetMetadata(v AccessControlPolicyListMetadataOutput)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



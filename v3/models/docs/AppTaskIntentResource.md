# AppTaskIntentResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**AppTaskDefStatus**](AppTaskDefStatus.md) |  | [optional] 
**Spec** | Pointer to [**AppTask**](AppTask.md) |  | [optional] 
**ApiVersion** | Pointer to **string** | API Version of the Nutanix v3 API framework. | [optional] [readonly] [default to "3.1.0"]
**Metadata** | [**AppTaskMetadata**](AppTaskMetadata.md) |  | 

## Methods

### NewAppTaskIntentResource

`func NewAppTaskIntentResource(metadata AppTaskMetadata, ) *AppTaskIntentResource`

NewAppTaskIntentResource instantiates a new AppTaskIntentResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppTaskIntentResourceWithDefaults

`func NewAppTaskIntentResourceWithDefaults() *AppTaskIntentResource`

NewAppTaskIntentResourceWithDefaults instantiates a new AppTaskIntentResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *AppTaskIntentResource) GetStatus() AppTaskDefStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AppTaskIntentResource) GetStatusOk() (*AppTaskDefStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AppTaskIntentResource) SetStatus(v AppTaskDefStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AppTaskIntentResource) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSpec

`func (o *AppTaskIntentResource) GetSpec() AppTask`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *AppTaskIntentResource) GetSpecOk() (*AppTask, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *AppTaskIntentResource) SetSpec(v AppTask)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *AppTaskIntentResource) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetApiVersion

`func (o *AppTaskIntentResource) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *AppTaskIntentResource) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *AppTaskIntentResource) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *AppTaskIntentResource) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetMetadata

`func (o *AppTaskIntentResource) GetMetadata() AppTaskMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AppTaskIntentResource) GetMetadataOk() (*AppTaskMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AppTaskIntentResource) SetMetadata(v AppTaskMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



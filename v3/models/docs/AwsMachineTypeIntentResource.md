# AwsMachineTypeIntentResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**AwsMachineTypeDefStatus**](AwsMachineTypeDefStatus.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**AwsMachineTypeMetadata**](AwsMachineTypeMetadata.md) |  | 

## Methods

### NewAwsMachineTypeIntentResource

`func NewAwsMachineTypeIntentResource(apiVersion string, metadata AwsMachineTypeMetadata, ) *AwsMachineTypeIntentResource`

NewAwsMachineTypeIntentResource instantiates a new AwsMachineTypeIntentResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsMachineTypeIntentResourceWithDefaults

`func NewAwsMachineTypeIntentResourceWithDefaults() *AwsMachineTypeIntentResource`

NewAwsMachineTypeIntentResourceWithDefaults instantiates a new AwsMachineTypeIntentResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *AwsMachineTypeIntentResource) GetStatus() AwsMachineTypeDefStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AwsMachineTypeIntentResource) GetStatusOk() (*AwsMachineTypeDefStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AwsMachineTypeIntentResource) SetStatus(v AwsMachineTypeDefStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AwsMachineTypeIntentResource) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetApiVersion

`func (o *AwsMachineTypeIntentResource) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *AwsMachineTypeIntentResource) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *AwsMachineTypeIntentResource) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *AwsMachineTypeIntentResource) GetMetadata() AwsMachineTypeMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AwsMachineTypeIntentResource) GetMetadataOk() (*AwsMachineTypeMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AwsMachineTypeIntentResource) SetMetadata(v AwsMachineTypeMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



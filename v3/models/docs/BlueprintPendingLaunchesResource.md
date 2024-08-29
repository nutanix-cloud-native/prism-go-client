# BlueprintPendingLaunchesResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**BlueprintPendingLaunchesStatus**](BlueprintPendingLaunchesStatus.md) |  | [optional] 
**ApiVersion** | Pointer to **string** | API Version of the Nutanix v3 API framework. | [optional] [readonly] [default to "3.1.0"]
**Metadata** | [**BlueprintMetadata**](BlueprintMetadata.md) |  | 

## Methods

### NewBlueprintPendingLaunchesResource

`func NewBlueprintPendingLaunchesResource(metadata BlueprintMetadata, ) *BlueprintPendingLaunchesResource`

NewBlueprintPendingLaunchesResource instantiates a new BlueprintPendingLaunchesResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintPendingLaunchesResourceWithDefaults

`func NewBlueprintPendingLaunchesResourceWithDefaults() *BlueprintPendingLaunchesResource`

NewBlueprintPendingLaunchesResourceWithDefaults instantiates a new BlueprintPendingLaunchesResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *BlueprintPendingLaunchesResource) GetStatus() BlueprintPendingLaunchesStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BlueprintPendingLaunchesResource) GetStatusOk() (*BlueprintPendingLaunchesStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BlueprintPendingLaunchesResource) SetStatus(v BlueprintPendingLaunchesStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BlueprintPendingLaunchesResource) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetApiVersion

`func (o *BlueprintPendingLaunchesResource) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *BlueprintPendingLaunchesResource) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *BlueprintPendingLaunchesResource) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *BlueprintPendingLaunchesResource) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetMetadata

`func (o *BlueprintPendingLaunchesResource) GetMetadata() BlueprintMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *BlueprintPendingLaunchesResource) GetMetadataOk() (*BlueprintMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *BlueprintPendingLaunchesResource) SetMetadata(v BlueprintMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



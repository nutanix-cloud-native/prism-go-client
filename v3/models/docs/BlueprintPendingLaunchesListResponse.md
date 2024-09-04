# BlueprintPendingLaunchesListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entites** | Pointer to [**[]BlueprintPendingLaunchesResource**](BlueprintPendingLaunchesResource.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**BlueprintListMetadataOutput**](BlueprintListMetadataOutput.md) |  | 

## Methods

### NewBlueprintPendingLaunchesListResponse

`func NewBlueprintPendingLaunchesListResponse(apiVersion string, metadata BlueprintListMetadataOutput, ) *BlueprintPendingLaunchesListResponse`

NewBlueprintPendingLaunchesListResponse instantiates a new BlueprintPendingLaunchesListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintPendingLaunchesListResponseWithDefaults

`func NewBlueprintPendingLaunchesListResponseWithDefaults() *BlueprintPendingLaunchesListResponse`

NewBlueprintPendingLaunchesListResponseWithDefaults instantiates a new BlueprintPendingLaunchesListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntites

`func (o *BlueprintPendingLaunchesListResponse) GetEntites() []BlueprintPendingLaunchesResource`

GetEntites returns the Entites field if non-nil, zero value otherwise.

### GetEntitesOk

`func (o *BlueprintPendingLaunchesListResponse) GetEntitesOk() (*[]BlueprintPendingLaunchesResource, bool)`

GetEntitesOk returns a tuple with the Entites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntites

`func (o *BlueprintPendingLaunchesListResponse) SetEntites(v []BlueprintPendingLaunchesResource)`

SetEntites sets Entites field to given value.

### HasEntites

`func (o *BlueprintPendingLaunchesListResponse) HasEntites() bool`

HasEntites returns a boolean if a field has been set.

### GetApiVersion

`func (o *BlueprintPendingLaunchesListResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *BlueprintPendingLaunchesListResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *BlueprintPendingLaunchesListResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *BlueprintPendingLaunchesListResponse) GetMetadata() BlueprintListMetadataOutput`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *BlueprintPendingLaunchesListResponse) GetMetadataOk() (*BlueprintListMetadataOutput, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *BlueprintPendingLaunchesListResponse) SetMetadata(v BlueprintListMetadataOutput)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# CellIntentResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**CellDefStatus**](CellDefStatus.md) |  | [optional] 
**Spec** | Pointer to [**Cell**](Cell.md) |  | [optional] 
**ApiVersion** | Pointer to **string** | API Version of the Nutanix v3 API framework. | [optional] [readonly] [default to "3.1.0"]
**Metadata** | [**CellMetadata**](CellMetadata.md) |  | 

## Methods

### NewCellIntentResource

`func NewCellIntentResource(metadata CellMetadata, ) *CellIntentResource`

NewCellIntentResource instantiates a new CellIntentResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCellIntentResourceWithDefaults

`func NewCellIntentResourceWithDefaults() *CellIntentResource`

NewCellIntentResourceWithDefaults instantiates a new CellIntentResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *CellIntentResource) GetStatus() CellDefStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CellIntentResource) GetStatusOk() (*CellDefStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CellIntentResource) SetStatus(v CellDefStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CellIntentResource) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSpec

`func (o *CellIntentResource) GetSpec() Cell`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *CellIntentResource) GetSpecOk() (*Cell, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *CellIntentResource) SetSpec(v Cell)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *CellIntentResource) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetApiVersion

`func (o *CellIntentResource) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *CellIntentResource) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *CellIntentResource) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *CellIntentResource) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetMetadata

`func (o *CellIntentResource) GetMetadata() CellMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CellIntentResource) GetMetadataOk() (*CellMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CellIntentResource) SetMetadata(v CellMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



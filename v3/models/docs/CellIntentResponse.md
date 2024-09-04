# CellIntentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**CellDefStatus**](CellDefStatus.md) |  | [optional] 
**Spec** | Pointer to [**Cell**](Cell.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**CellMetadata**](CellMetadata.md) |  | 

## Methods

### NewCellIntentResponse

`func NewCellIntentResponse(apiVersion string, metadata CellMetadata, ) *CellIntentResponse`

NewCellIntentResponse instantiates a new CellIntentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCellIntentResponseWithDefaults

`func NewCellIntentResponseWithDefaults() *CellIntentResponse`

NewCellIntentResponseWithDefaults instantiates a new CellIntentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *CellIntentResponse) GetStatus() CellDefStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CellIntentResponse) GetStatusOk() (*CellDefStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CellIntentResponse) SetStatus(v CellDefStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CellIntentResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSpec

`func (o *CellIntentResponse) GetSpec() Cell`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *CellIntentResponse) GetSpecOk() (*Cell, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *CellIntentResponse) SetSpec(v Cell)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *CellIntentResponse) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetApiVersion

`func (o *CellIntentResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *CellIntentResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *CellIntentResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *CellIntentResponse) GetMetadata() CellMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CellIntentResponse) GetMetadataOk() (*CellMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CellIntentResponse) SetMetadata(v CellMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



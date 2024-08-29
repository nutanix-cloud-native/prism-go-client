# CellIntentInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Spec** | [**Cell**](Cell.md) |  | 
**ApiVersion** | Pointer to **string** | API Version of the Nutanix v3 API framework. | [optional] [readonly] [default to "3.1.0"]
**Metadata** | [**CellMetadata**](CellMetadata.md) |  | 

## Methods

### NewCellIntentInput

`func NewCellIntentInput(spec Cell, metadata CellMetadata, ) *CellIntentInput`

NewCellIntentInput instantiates a new CellIntentInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCellIntentInputWithDefaults

`func NewCellIntentInputWithDefaults() *CellIntentInput`

NewCellIntentInputWithDefaults instantiates a new CellIntentInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpec

`func (o *CellIntentInput) GetSpec() Cell`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *CellIntentInput) GetSpecOk() (*Cell, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *CellIntentInput) SetSpec(v Cell)`

SetSpec sets Spec field to given value.


### GetApiVersion

`func (o *CellIntentInput) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *CellIntentInput) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *CellIntentInput) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *CellIntentInput) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetMetadata

`func (o *CellIntentInput) GetMetadata() CellMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CellIntentInput) GetMetadataOk() (*CellMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CellIntentInput) SetMetadata(v CellMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# UnderlaySubnetIntentInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Spec** | [**UnderlaySubnet**](UnderlaySubnet.md) |  | 
**ApiVersion** | Pointer to **string** | API Version of the Nutanix v3 API framework. | [optional] [readonly] [default to "3.1.0"]
**Metadata** | [**UnderlaySubnetMetadata**](UnderlaySubnetMetadata.md) |  | 

## Methods

### NewUnderlaySubnetIntentInput

`func NewUnderlaySubnetIntentInput(spec UnderlaySubnet, metadata UnderlaySubnetMetadata, ) *UnderlaySubnetIntentInput`

NewUnderlaySubnetIntentInput instantiates a new UnderlaySubnetIntentInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnderlaySubnetIntentInputWithDefaults

`func NewUnderlaySubnetIntentInputWithDefaults() *UnderlaySubnetIntentInput`

NewUnderlaySubnetIntentInputWithDefaults instantiates a new UnderlaySubnetIntentInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpec

`func (o *UnderlaySubnetIntentInput) GetSpec() UnderlaySubnet`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *UnderlaySubnetIntentInput) GetSpecOk() (*UnderlaySubnet, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *UnderlaySubnetIntentInput) SetSpec(v UnderlaySubnet)`

SetSpec sets Spec field to given value.


### GetApiVersion

`func (o *UnderlaySubnetIntentInput) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *UnderlaySubnetIntentInput) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *UnderlaySubnetIntentInput) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *UnderlaySubnetIntentInput) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetMetadata

`func (o *UnderlaySubnetIntentInput) GetMetadata() UnderlaySubnetMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UnderlaySubnetIntentInput) GetMetadataOk() (*UnderlaySubnetMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UnderlaySubnetIntentInput) SetMetadata(v UnderlaySubnetMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



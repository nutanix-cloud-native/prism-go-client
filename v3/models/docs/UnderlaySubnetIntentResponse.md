# UnderlaySubnetIntentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**UnderlaySubnetDefStatus**](UnderlaySubnetDefStatus.md) |  | [optional] 
**Spec** | Pointer to [**UnderlaySubnet**](UnderlaySubnet.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**UnderlaySubnetMetadata**](UnderlaySubnetMetadata.md) |  | 

## Methods

### NewUnderlaySubnetIntentResponse

`func NewUnderlaySubnetIntentResponse(apiVersion string, metadata UnderlaySubnetMetadata, ) *UnderlaySubnetIntentResponse`

NewUnderlaySubnetIntentResponse instantiates a new UnderlaySubnetIntentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnderlaySubnetIntentResponseWithDefaults

`func NewUnderlaySubnetIntentResponseWithDefaults() *UnderlaySubnetIntentResponse`

NewUnderlaySubnetIntentResponseWithDefaults instantiates a new UnderlaySubnetIntentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *UnderlaySubnetIntentResponse) GetStatus() UnderlaySubnetDefStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UnderlaySubnetIntentResponse) GetStatusOk() (*UnderlaySubnetDefStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UnderlaySubnetIntentResponse) SetStatus(v UnderlaySubnetDefStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UnderlaySubnetIntentResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSpec

`func (o *UnderlaySubnetIntentResponse) GetSpec() UnderlaySubnet`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *UnderlaySubnetIntentResponse) GetSpecOk() (*UnderlaySubnet, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *UnderlaySubnetIntentResponse) SetSpec(v UnderlaySubnet)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *UnderlaySubnetIntentResponse) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetApiVersion

`func (o *UnderlaySubnetIntentResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *UnderlaySubnetIntentResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *UnderlaySubnetIntentResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *UnderlaySubnetIntentResponse) GetMetadata() UnderlaySubnetMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UnderlaySubnetIntentResponse) GetMetadataOk() (*UnderlaySubnetMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UnderlaySubnetIntentResponse) SetMetadata(v UnderlaySubnetMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# ActionTypeIntentResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**ActionTypeDefStatus**](ActionTypeDefStatus.md) |  | [optional] 
**Spec** | Pointer to [**ActionType**](ActionType.md) |  | [optional] 
**ApiVersion** | Pointer to **string** | API Version of the Nutanix v3 API framework. | [optional] [readonly] [default to "3.1.0"]
**Metadata** | [**ActionTypeMetadata**](ActionTypeMetadata.md) |  | 

## Methods

### NewActionTypeIntentResource

`func NewActionTypeIntentResource(metadata ActionTypeMetadata, ) *ActionTypeIntentResource`

NewActionTypeIntentResource instantiates a new ActionTypeIntentResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionTypeIntentResourceWithDefaults

`func NewActionTypeIntentResourceWithDefaults() *ActionTypeIntentResource`

NewActionTypeIntentResourceWithDefaults instantiates a new ActionTypeIntentResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ActionTypeIntentResource) GetStatus() ActionTypeDefStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ActionTypeIntentResource) GetStatusOk() (*ActionTypeDefStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ActionTypeIntentResource) SetStatus(v ActionTypeDefStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ActionTypeIntentResource) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSpec

`func (o *ActionTypeIntentResource) GetSpec() ActionType`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *ActionTypeIntentResource) GetSpecOk() (*ActionType, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *ActionTypeIntentResource) SetSpec(v ActionType)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *ActionTypeIntentResource) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetApiVersion

`func (o *ActionTypeIntentResource) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ActionTypeIntentResource) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ActionTypeIntentResource) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *ActionTypeIntentResource) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetMetadata

`func (o *ActionTypeIntentResource) GetMetadata() ActionTypeMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ActionTypeIntentResource) GetMetadataOk() (*ActionTypeMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ActionTypeIntentResource) SetMetadata(v ActionTypeMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



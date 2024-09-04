# ActionTypeIntentInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Spec** | [**ActionType**](ActionType.md) |  | 
**ApiVersion** | Pointer to **string** | API Version of the Nutanix v3 API framework. | [optional] [readonly] [default to "3.1.0"]
**Metadata** | [**ActionTypeMetadata**](ActionTypeMetadata.md) |  | 

## Methods

### NewActionTypeIntentInput

`func NewActionTypeIntentInput(spec ActionType, metadata ActionTypeMetadata, ) *ActionTypeIntentInput`

NewActionTypeIntentInput instantiates a new ActionTypeIntentInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionTypeIntentInputWithDefaults

`func NewActionTypeIntentInputWithDefaults() *ActionTypeIntentInput`

NewActionTypeIntentInputWithDefaults instantiates a new ActionTypeIntentInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpec

`func (o *ActionTypeIntentInput) GetSpec() ActionType`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *ActionTypeIntentInput) GetSpecOk() (*ActionType, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *ActionTypeIntentInput) SetSpec(v ActionType)`

SetSpec sets Spec field to given value.


### GetApiVersion

`func (o *ActionTypeIntentInput) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ActionTypeIntentInput) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ActionTypeIntentInput) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *ActionTypeIntentInput) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetMetadata

`func (o *ActionTypeIntentInput) GetMetadata() ActionTypeMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ActionTypeIntentInput) GetMetadataOk() (*ActionTypeMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ActionTypeIntentInput) SetMetadata(v ActionTypeMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



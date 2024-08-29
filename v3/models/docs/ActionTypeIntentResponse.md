# ActionTypeIntentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**ActionTypeDefStatus**](ActionTypeDefStatus.md) |  | [optional] 
**Spec** | Pointer to [**ActionType**](ActionType.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**ActionTypeMetadata**](ActionTypeMetadata.md) |  | 

## Methods

### NewActionTypeIntentResponse

`func NewActionTypeIntentResponse(apiVersion string, metadata ActionTypeMetadata, ) *ActionTypeIntentResponse`

NewActionTypeIntentResponse instantiates a new ActionTypeIntentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionTypeIntentResponseWithDefaults

`func NewActionTypeIntentResponseWithDefaults() *ActionTypeIntentResponse`

NewActionTypeIntentResponseWithDefaults instantiates a new ActionTypeIntentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ActionTypeIntentResponse) GetStatus() ActionTypeDefStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ActionTypeIntentResponse) GetStatusOk() (*ActionTypeDefStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ActionTypeIntentResponse) SetStatus(v ActionTypeDefStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ActionTypeIntentResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSpec

`func (o *ActionTypeIntentResponse) GetSpec() ActionType`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *ActionTypeIntentResponse) GetSpecOk() (*ActionType, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *ActionTypeIntentResponse) SetSpec(v ActionType)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *ActionTypeIntentResponse) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetApiVersion

`func (o *ActionTypeIntentResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ActionTypeIntentResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ActionTypeIntentResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *ActionTypeIntentResponse) GetMetadata() ActionTypeMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ActionTypeIntentResponse) GetMetadataOk() (*ActionTypeMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ActionTypeIntentResponse) SetMetadata(v ActionTypeMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



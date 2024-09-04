# ActionTriggerTypeIntentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**ActionTriggerTypeDefStatus**](ActionTriggerTypeDefStatus.md) |  | [optional] 
**Spec** | Pointer to [**ActionTriggerType**](ActionTriggerType.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**ActionTriggerTypeMetadata**](ActionTriggerTypeMetadata.md) |  | 

## Methods

### NewActionTriggerTypeIntentResponse

`func NewActionTriggerTypeIntentResponse(apiVersion string, metadata ActionTriggerTypeMetadata, ) *ActionTriggerTypeIntentResponse`

NewActionTriggerTypeIntentResponse instantiates a new ActionTriggerTypeIntentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionTriggerTypeIntentResponseWithDefaults

`func NewActionTriggerTypeIntentResponseWithDefaults() *ActionTriggerTypeIntentResponse`

NewActionTriggerTypeIntentResponseWithDefaults instantiates a new ActionTriggerTypeIntentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ActionTriggerTypeIntentResponse) GetStatus() ActionTriggerTypeDefStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ActionTriggerTypeIntentResponse) GetStatusOk() (*ActionTriggerTypeDefStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ActionTriggerTypeIntentResponse) SetStatus(v ActionTriggerTypeDefStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ActionTriggerTypeIntentResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSpec

`func (o *ActionTriggerTypeIntentResponse) GetSpec() ActionTriggerType`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *ActionTriggerTypeIntentResponse) GetSpecOk() (*ActionTriggerType, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *ActionTriggerTypeIntentResponse) SetSpec(v ActionTriggerType)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *ActionTriggerTypeIntentResponse) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetApiVersion

`func (o *ActionTriggerTypeIntentResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ActionTriggerTypeIntentResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ActionTriggerTypeIntentResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *ActionTriggerTypeIntentResponse) GetMetadata() ActionTriggerTypeMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ActionTriggerTypeIntentResponse) GetMetadataOk() (*ActionTriggerTypeMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ActionTriggerTypeIntentResponse) SetMetadata(v ActionTriggerTypeMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



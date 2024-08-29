# ActionTriggerTypeListIntentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entities** | Pointer to [**[]ActionTriggerTypeIntentResource**](ActionTriggerTypeIntentResource.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**ActionTriggerTypeListMetadataOutput**](ActionTriggerTypeListMetadataOutput.md) |  | 

## Methods

### NewActionTriggerTypeListIntentResponse

`func NewActionTriggerTypeListIntentResponse(apiVersion string, metadata ActionTriggerTypeListMetadataOutput, ) *ActionTriggerTypeListIntentResponse`

NewActionTriggerTypeListIntentResponse instantiates a new ActionTriggerTypeListIntentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionTriggerTypeListIntentResponseWithDefaults

`func NewActionTriggerTypeListIntentResponseWithDefaults() *ActionTriggerTypeListIntentResponse`

NewActionTriggerTypeListIntentResponseWithDefaults instantiates a new ActionTriggerTypeListIntentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntities

`func (o *ActionTriggerTypeListIntentResponse) GetEntities() []ActionTriggerTypeIntentResource`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *ActionTriggerTypeListIntentResponse) GetEntitiesOk() (*[]ActionTriggerTypeIntentResource, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *ActionTriggerTypeListIntentResponse) SetEntities(v []ActionTriggerTypeIntentResource)`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *ActionTriggerTypeListIntentResponse) HasEntities() bool`

HasEntities returns a boolean if a field has been set.

### GetApiVersion

`func (o *ActionTriggerTypeListIntentResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ActionTriggerTypeListIntentResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ActionTriggerTypeListIntentResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *ActionTriggerTypeListIntentResponse) GetMetadata() ActionTriggerTypeListMetadataOutput`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ActionTriggerTypeListIntentResponse) GetMetadataOk() (*ActionTriggerTypeListMetadataOutput, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ActionTriggerTypeListIntentResponse) SetMetadata(v ActionTriggerTypeListMetadataOutput)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



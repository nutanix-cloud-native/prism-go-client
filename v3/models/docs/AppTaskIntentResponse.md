# AppTaskIntentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**AppTaskDefStatus**](AppTaskDefStatus.md) |  | [optional] 
**Spec** | Pointer to [**AppTask**](AppTask.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**AppTaskMetadata**](AppTaskMetadata.md) |  | 

## Methods

### NewAppTaskIntentResponse

`func NewAppTaskIntentResponse(apiVersion string, metadata AppTaskMetadata, ) *AppTaskIntentResponse`

NewAppTaskIntentResponse instantiates a new AppTaskIntentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppTaskIntentResponseWithDefaults

`func NewAppTaskIntentResponseWithDefaults() *AppTaskIntentResponse`

NewAppTaskIntentResponseWithDefaults instantiates a new AppTaskIntentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *AppTaskIntentResponse) GetStatus() AppTaskDefStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AppTaskIntentResponse) GetStatusOk() (*AppTaskDefStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AppTaskIntentResponse) SetStatus(v AppTaskDefStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AppTaskIntentResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSpec

`func (o *AppTaskIntentResponse) GetSpec() AppTask`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *AppTaskIntentResponse) GetSpecOk() (*AppTask, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *AppTaskIntentResponse) SetSpec(v AppTask)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *AppTaskIntentResponse) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetApiVersion

`func (o *AppTaskIntentResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *AppTaskIntentResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *AppTaskIntentResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *AppTaskIntentResponse) GetMetadata() AppTaskMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AppTaskIntentResponse) GetMetadataOk() (*AppTaskMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AppTaskIntentResponse) SetMetadata(v AppTaskMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



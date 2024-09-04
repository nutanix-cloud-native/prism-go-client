# AlertIntentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**AlertDefStatus**](AlertDefStatus.md) |  | [optional] 
**Spec** | Pointer to [**Alert**](Alert.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**AlertMetadata**](AlertMetadata.md) |  | 

## Methods

### NewAlertIntentResponse

`func NewAlertIntentResponse(apiVersion string, metadata AlertMetadata, ) *AlertIntentResponse`

NewAlertIntentResponse instantiates a new AlertIntentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertIntentResponseWithDefaults

`func NewAlertIntentResponseWithDefaults() *AlertIntentResponse`

NewAlertIntentResponseWithDefaults instantiates a new AlertIntentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *AlertIntentResponse) GetStatus() AlertDefStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AlertIntentResponse) GetStatusOk() (*AlertDefStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AlertIntentResponse) SetStatus(v AlertDefStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AlertIntentResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSpec

`func (o *AlertIntentResponse) GetSpec() Alert`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *AlertIntentResponse) GetSpecOk() (*Alert, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *AlertIntentResponse) SetSpec(v Alert)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *AlertIntentResponse) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetApiVersion

`func (o *AlertIntentResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *AlertIntentResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *AlertIntentResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *AlertIntentResponse) GetMetadata() AlertMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AlertIntentResponse) GetMetadataOk() (*AlertMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AlertIntentResponse) SetMetadata(v AlertMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



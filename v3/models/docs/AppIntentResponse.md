# AppIntentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**AppDefStatus**](AppDefStatus.md) |  | [optional] 
**Spec** | Pointer to [**App**](App.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**AppMetadata**](AppMetadata.md) |  | 

## Methods

### NewAppIntentResponse

`func NewAppIntentResponse(apiVersion string, metadata AppMetadata, ) *AppIntentResponse`

NewAppIntentResponse instantiates a new AppIntentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppIntentResponseWithDefaults

`func NewAppIntentResponseWithDefaults() *AppIntentResponse`

NewAppIntentResponseWithDefaults instantiates a new AppIntentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *AppIntentResponse) GetStatus() AppDefStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AppIntentResponse) GetStatusOk() (*AppDefStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AppIntentResponse) SetStatus(v AppDefStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AppIntentResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSpec

`func (o *AppIntentResponse) GetSpec() App`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *AppIntentResponse) GetSpecOk() (*App, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *AppIntentResponse) SetSpec(v App)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *AppIntentResponse) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetApiVersion

`func (o *AppIntentResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *AppIntentResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *AppIntentResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *AppIntentResponse) GetMetadata() AppMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AppIntentResponse) GetMetadataOk() (*AppMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AppIntentResponse) SetMetadata(v AppMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# AccountIntentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**AccountDefStatus**](AccountDefStatus.md) |  | [optional] 
**Spec** | Pointer to [**Account**](Account.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**AccountMetadata**](AccountMetadata.md) |  | 

## Methods

### NewAccountIntentResponse

`func NewAccountIntentResponse(apiVersion string, metadata AccountMetadata, ) *AccountIntentResponse`

NewAccountIntentResponse instantiates a new AccountIntentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountIntentResponseWithDefaults

`func NewAccountIntentResponseWithDefaults() *AccountIntentResponse`

NewAccountIntentResponseWithDefaults instantiates a new AccountIntentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *AccountIntentResponse) GetStatus() AccountDefStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AccountIntentResponse) GetStatusOk() (*AccountDefStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AccountIntentResponse) SetStatus(v AccountDefStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AccountIntentResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSpec

`func (o *AccountIntentResponse) GetSpec() Account`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *AccountIntentResponse) GetSpecOk() (*Account, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *AccountIntentResponse) SetSpec(v Account)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *AccountIntentResponse) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetApiVersion

`func (o *AccountIntentResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *AccountIntentResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *AccountIntentResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *AccountIntentResponse) GetMetadata() AccountMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AccountIntentResponse) GetMetadataOk() (*AccountMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AccountIntentResponse) SetMetadata(v AccountMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



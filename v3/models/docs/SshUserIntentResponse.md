# SshUserIntentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**SshUserDefStatus**](SshUserDefStatus.md) |  | [optional] 
**Spec** | Pointer to [**SshUser**](SshUser.md) |  | [optional] 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**SshUserMetadata**](SshUserMetadata.md) |  | 

## Methods

### NewSshUserIntentResponse

`func NewSshUserIntentResponse(apiVersion string, metadata SshUserMetadata, ) *SshUserIntentResponse`

NewSshUserIntentResponse instantiates a new SshUserIntentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSshUserIntentResponseWithDefaults

`func NewSshUserIntentResponseWithDefaults() *SshUserIntentResponse`

NewSshUserIntentResponseWithDefaults instantiates a new SshUserIntentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *SshUserIntentResponse) GetStatus() SshUserDefStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SshUserIntentResponse) GetStatusOk() (*SshUserDefStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SshUserIntentResponse) SetStatus(v SshUserDefStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SshUserIntentResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSpec

`func (o *SshUserIntentResponse) GetSpec() SshUser`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *SshUserIntentResponse) GetSpecOk() (*SshUser, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *SshUserIntentResponse) SetSpec(v SshUser)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *SshUserIntentResponse) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetApiVersion

`func (o *SshUserIntentResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *SshUserIntentResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *SshUserIntentResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *SshUserIntentResponse) GetMetadata() SshUserMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SshUserIntentResponse) GetMetadataOk() (*SshUserMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SshUserIntentResponse) SetMetadata(v SshUserMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



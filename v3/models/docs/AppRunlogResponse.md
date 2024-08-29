# AppRunlogResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**AppRunlogResources**](AppRunlogResources.md) |  | 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**AppRunlogMetadata**](AppRunlogMetadata.md) |  | 

## Methods

### NewAppRunlogResponse

`func NewAppRunlogResponse(status AppRunlogResources, apiVersion string, metadata AppRunlogMetadata, ) *AppRunlogResponse`

NewAppRunlogResponse instantiates a new AppRunlogResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppRunlogResponseWithDefaults

`func NewAppRunlogResponseWithDefaults() *AppRunlogResponse`

NewAppRunlogResponseWithDefaults instantiates a new AppRunlogResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *AppRunlogResponse) GetStatus() AppRunlogResources`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AppRunlogResponse) GetStatusOk() (*AppRunlogResources, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AppRunlogResponse) SetStatus(v AppRunlogResources)`

SetStatus sets Status field to given value.


### GetApiVersion

`func (o *AppRunlogResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *AppRunlogResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *AppRunlogResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *AppRunlogResponse) GetMetadata() AppRunlogMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AppRunlogResponse) GetMetadataOk() (*AppRunlogMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AppRunlogResponse) SetMetadata(v AppRunlogMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



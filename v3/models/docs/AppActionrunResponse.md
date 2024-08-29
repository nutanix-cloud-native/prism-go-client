# AppActionrunResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**AppActionrunStatus**](AppActionrunStatus.md) |  | 
**Spec** | [**AppActionrunSpec**](AppActionrunSpec.md) |  | 
**ApiVersion** | **string** | API Version of the Nutanix v3 API framework. | [readonly] [default to "3.1.0"]
**Metadata** | [**AppMetadata**](AppMetadata.md) |  | 

## Methods

### NewAppActionrunResponse

`func NewAppActionrunResponse(status AppActionrunStatus, spec AppActionrunSpec, apiVersion string, metadata AppMetadata, ) *AppActionrunResponse`

NewAppActionrunResponse instantiates a new AppActionrunResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppActionrunResponseWithDefaults

`func NewAppActionrunResponseWithDefaults() *AppActionrunResponse`

NewAppActionrunResponseWithDefaults instantiates a new AppActionrunResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *AppActionrunResponse) GetStatus() AppActionrunStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AppActionrunResponse) GetStatusOk() (*AppActionrunStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AppActionrunResponse) SetStatus(v AppActionrunStatus)`

SetStatus sets Status field to given value.


### GetSpec

`func (o *AppActionrunResponse) GetSpec() AppActionrunSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *AppActionrunResponse) GetSpecOk() (*AppActionrunSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *AppActionrunResponse) SetSpec(v AppActionrunSpec)`

SetSpec sets Spec field to given value.


### GetApiVersion

`func (o *AppActionrunResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *AppActionrunResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *AppActionrunResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetMetadata

`func (o *AppActionrunResponse) GetMetadata() AppMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AppActionrunResponse) GetMetadataOk() (*AppMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AppActionrunResponse) SetMetadata(v AppMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



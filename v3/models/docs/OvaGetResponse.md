# OvaGetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Info** | Pointer to [**OvaInfo**](OvaInfo.md) |  | [optional] 
**ApiVersion** | Pointer to **string** | API Version of the Nutanix v3 API framework. | [optional] [readonly] [default to "3.1.0"]
**Metadata** | Pointer to [**OvaMetadata**](OvaMetadata.md) |  | [optional] 

## Methods

### NewOvaGetResponse

`func NewOvaGetResponse() *OvaGetResponse`

NewOvaGetResponse instantiates a new OvaGetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOvaGetResponseWithDefaults

`func NewOvaGetResponseWithDefaults() *OvaGetResponse`

NewOvaGetResponseWithDefaults instantiates a new OvaGetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInfo

`func (o *OvaGetResponse) GetInfo() OvaInfo`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *OvaGetResponse) GetInfoOk() (*OvaInfo, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *OvaGetResponse) SetInfo(v OvaInfo)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *OvaGetResponse) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetApiVersion

`func (o *OvaGetResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *OvaGetResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *OvaGetResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *OvaGetResponse) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetMetadata

`func (o *OvaGetResponse) GetMetadata() OvaMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *OvaGetResponse) GetMetadataOk() (*OvaMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *OvaGetResponse) SetMetadata(v OvaMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *OvaGetResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



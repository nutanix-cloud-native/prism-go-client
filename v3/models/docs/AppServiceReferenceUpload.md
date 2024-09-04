# AppServiceReferenceUpload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | The kind name | [readonly] [default to "app_service"]
**Name** | **string** |  | [readonly] 
**UUID** | Pointer to **string** |  | [optional] 

## Methods

### NewAppServiceReferenceUpload

`func NewAppServiceReferenceUpload(kind string, name string, ) *AppServiceReferenceUpload`

NewAppServiceReferenceUpload instantiates a new AppServiceReferenceUpload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppServiceReferenceUploadWithDefaults

`func NewAppServiceReferenceUploadWithDefaults() *AppServiceReferenceUpload`

NewAppServiceReferenceUploadWithDefaults instantiates a new AppServiceReferenceUpload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *AppServiceReferenceUpload) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *AppServiceReferenceUpload) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *AppServiceReferenceUpload) SetKind(v string)`

SetKind sets Kind field to given value.


### GetName

`func (o *AppServiceReferenceUpload) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppServiceReferenceUpload) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppServiceReferenceUpload) SetName(v string)`

SetName sets Name field to given value.


### GetUUID

`func (o *AppServiceReferenceUpload) GetUUID() string`

GetUUID returns the UUID field if non-nil, zero value otherwise.

### GetUUIDOk

`func (o *AppServiceReferenceUpload) GetUUIDOk() (*string, bool)`

GetUUIDOk returns a tuple with the UUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUUID

`func (o *AppServiceReferenceUpload) SetUUID(v string)`

SetUUID sets UUID field to given value.

### HasUUID

`func (o *AppServiceReferenceUpload) HasUUID() bool`

HasUUID returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# AccountReferenceUpload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | The kind name | [readonly] [default to "account"]
**Name** | **string** |  | [readonly] 
**UUID** | Pointer to **string** |  | [optional] 

## Methods

### NewAccountReferenceUpload

`func NewAccountReferenceUpload(kind string, name string, ) *AccountReferenceUpload`

NewAccountReferenceUpload instantiates a new AccountReferenceUpload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountReferenceUploadWithDefaults

`func NewAccountReferenceUploadWithDefaults() *AccountReferenceUpload`

NewAccountReferenceUploadWithDefaults instantiates a new AccountReferenceUpload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *AccountReferenceUpload) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *AccountReferenceUpload) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *AccountReferenceUpload) SetKind(v string)`

SetKind sets Kind field to given value.


### GetName

`func (o *AccountReferenceUpload) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccountReferenceUpload) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccountReferenceUpload) SetName(v string)`

SetName sets Name field to given value.


### GetUUID

`func (o *AccountReferenceUpload) GetUUID() string`

GetUUID returns the UUID field if non-nil, zero value otherwise.

### GetUUIDOk

`func (o *AccountReferenceUpload) GetUUIDOk() (*string, bool)`

GetUUIDOk returns a tuple with the UUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUUID

`func (o *AccountReferenceUpload) SetUUID(v string)`

SetUUID sets UUID field to given value.

### HasUUID

`func (o *AccountReferenceUpload) HasUUID() bool`

HasUUID returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



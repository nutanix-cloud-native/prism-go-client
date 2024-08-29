# Reference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** | GET query on the URL will provide information on the source.  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**UUID** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewReference

`func NewReference() *Reference`

NewReference instantiates a new Reference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReferenceWithDefaults

`func NewReferenceWithDefaults() *Reference`

NewReferenceWithDefaults instantiates a new Reference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *Reference) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Reference) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Reference) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *Reference) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetKind

`func (o *Reference) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *Reference) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *Reference) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *Reference) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetUUID

`func (o *Reference) GetUUID() string`

GetUUID returns the UUID field if non-nil, zero value otherwise.

### GetUUIDOk

`func (o *Reference) GetUUIDOk() (*string, bool)`

GetUUIDOk returns a tuple with the UUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUUID

`func (o *Reference) SetUUID(v string)`

SetUUID sets UUID field to given value.

### HasUUID

`func (o *Reference) HasUUID() bool`

HasUUID returns a boolean if a field has been set.

### GetName

`func (o *Reference) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Reference) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Reference) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Reference) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



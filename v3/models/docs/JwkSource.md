# JwkSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Api** | Pointer to **string** | XIG will poll on this api for JSON web keys. | [optional] 

## Methods

### NewJwkSource

`func NewJwkSource() *JwkSource`

NewJwkSource instantiates a new JwkSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJwkSourceWithDefaults

`func NewJwkSourceWithDefaults() *JwkSource`

NewJwkSourceWithDefaults instantiates a new JwkSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApi

`func (o *JwkSource) GetApi() string`

GetApi returns the Api field if non-nil, zero value otherwise.

### GetApiOk

`func (o *JwkSource) GetApiOk() (*string, bool)`

GetApiOk returns a tuple with the Api field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApi

`func (o *JwkSource) SetApi(v string)`

SetApi sets Api field to given value.

### HasApi

`func (o *JwkSource) HasApi() bool`

HasApi returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



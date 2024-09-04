# SourceOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowInsecureConnection** | Pointer to **bool** | allow_insecure_connection is an option for the user to ignore the server certificate verification while accessing source_uri (image location). If it has value &#x3D; true, it ignores server certificate verification. If it has value &#x3D; false, it does regular server certificate verification.  | [optional] 
**BasicAuth** | Pointer to [**Credentials**](Credentials.md) |  | [optional] 

## Methods

### NewSourceOptions

`func NewSourceOptions() *SourceOptions`

NewSourceOptions instantiates a new SourceOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceOptionsWithDefaults

`func NewSourceOptionsWithDefaults() *SourceOptions`

NewSourceOptionsWithDefaults instantiates a new SourceOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowInsecureConnection

`func (o *SourceOptions) GetAllowInsecureConnection() bool`

GetAllowInsecureConnection returns the AllowInsecureConnection field if non-nil, zero value otherwise.

### GetAllowInsecureConnectionOk

`func (o *SourceOptions) GetAllowInsecureConnectionOk() (*bool, bool)`

GetAllowInsecureConnectionOk returns a tuple with the AllowInsecureConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowInsecureConnection

`func (o *SourceOptions) SetAllowInsecureConnection(v bool)`

SetAllowInsecureConnection sets AllowInsecureConnection field to given value.

### HasAllowInsecureConnection

`func (o *SourceOptions) HasAllowInsecureConnection() bool`

HasAllowInsecureConnection returns a boolean if a field has been set.

### GetBasicAuth

`func (o *SourceOptions) GetBasicAuth() Credentials`

GetBasicAuth returns the BasicAuth field if non-nil, zero value otherwise.

### GetBasicAuthOk

`func (o *SourceOptions) GetBasicAuthOk() (*Credentials, bool)`

GetBasicAuthOk returns a tuple with the BasicAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicAuth

`func (o *SourceOptions) SetBasicAuth(v Credentials)`

SetBasicAuth sets BasicAuth field to given value.

### HasBasicAuth

`func (o *SourceOptions) HasBasicAuth() bool`

HasBasicAuth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



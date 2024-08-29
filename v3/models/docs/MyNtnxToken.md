# MyNtnxToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KeyId** | **string** | Unique identifier for the API key. | 
**ApiKey** | **string** | API Key issue by my.nutanix API Keys app. | 
**ScopeId** | **string** | my.nutanix scope Id assigned to the deployed component. | 
**SubscopeId** | **string** | my.nutanix subscope Id assigned to the deployed component. | 

## Methods

### NewMyNtnxToken

`func NewMyNtnxToken(keyId string, apiKey string, scopeId string, subscopeId string, ) *MyNtnxToken`

NewMyNtnxToken instantiates a new MyNtnxToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMyNtnxTokenWithDefaults

`func NewMyNtnxTokenWithDefaults() *MyNtnxToken`

NewMyNtnxTokenWithDefaults instantiates a new MyNtnxToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeyId

`func (o *MyNtnxToken) GetKeyId() string`

GetKeyId returns the KeyId field if non-nil, zero value otherwise.

### GetKeyIdOk

`func (o *MyNtnxToken) GetKeyIdOk() (*string, bool)`

GetKeyIdOk returns a tuple with the KeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyId

`func (o *MyNtnxToken) SetKeyId(v string)`

SetKeyId sets KeyId field to given value.


### GetApiKey

`func (o *MyNtnxToken) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *MyNtnxToken) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *MyNtnxToken) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.


### GetScopeId

`func (o *MyNtnxToken) GetScopeId() string`

GetScopeId returns the ScopeId field if non-nil, zero value otherwise.

### GetScopeIdOk

`func (o *MyNtnxToken) GetScopeIdOk() (*string, bool)`

GetScopeIdOk returns a tuple with the ScopeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeId

`func (o *MyNtnxToken) SetScopeId(v string)`

SetScopeId sets ScopeId field to given value.


### GetSubscopeId

`func (o *MyNtnxToken) GetSubscopeId() string`

GetSubscopeId returns the SubscopeId field if non-nil, zero value otherwise.

### GetSubscopeIdOk

`func (o *MyNtnxToken) GetSubscopeIdOk() (*string, bool)`

GetSubscopeIdOk returns a tuple with the SubscopeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscopeId

`func (o *MyNtnxToken) SetSubscopeId(v string)`

SetSubscopeId sets SubscopeId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



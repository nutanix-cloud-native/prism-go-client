# AppCredentialResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | **string** |  | 
**UUID** | **string** |  | 
**MessageList** | Pointer to [**[]MessageResource**](MessageResource.md) | Message list | [optional] 
**State** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**Secret** | **map[string]interface{}** | Credential secret object | 
**Editables** | Pointer to **map[string]interface{}** | Runtime editable attributes for this entity. The structure for this is a dictionary. The keys in this dictionary should be the name of the attribute on the entity. If the attribute is editable, the value should be true, else false. If the attribute is a nested dictionary, the value can contain a nested dictionary with the same key value structure described above.  | [optional] 
**Type** | **string** |  | 
**Passphrase** | Pointer to **map[string]interface{}** | Credential passphrase object associated with the provided key | [optional] 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewAppCredentialResponse

`func NewAppCredentialResponse(username string, uUID string, name string, secret map[string]interface{}, type_ string, ) *AppCredentialResponse`

NewAppCredentialResponse instantiates a new AppCredentialResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppCredentialResponseWithDefaults

`func NewAppCredentialResponseWithDefaults() *AppCredentialResponse`

NewAppCredentialResponseWithDefaults instantiates a new AppCredentialResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *AppCredentialResponse) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *AppCredentialResponse) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *AppCredentialResponse) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetUUID

`func (o *AppCredentialResponse) GetUUID() string`

GetUUID returns the UUID field if non-nil, zero value otherwise.

### GetUUIDOk

`func (o *AppCredentialResponse) GetUUIDOk() (*string, bool)`

GetUUIDOk returns a tuple with the UUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUUID

`func (o *AppCredentialResponse) SetUUID(v string)`

SetUUID sets UUID field to given value.


### GetMessageList

`func (o *AppCredentialResponse) GetMessageList() []MessageResource`

GetMessageList returns the MessageList field if non-nil, zero value otherwise.

### GetMessageListOk

`func (o *AppCredentialResponse) GetMessageListOk() (*[]MessageResource, bool)`

GetMessageListOk returns a tuple with the MessageList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageList

`func (o *AppCredentialResponse) SetMessageList(v []MessageResource)`

SetMessageList sets MessageList field to given value.

### HasMessageList

`func (o *AppCredentialResponse) HasMessageList() bool`

HasMessageList returns a boolean if a field has been set.

### GetState

`func (o *AppCredentialResponse) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AppCredentialResponse) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AppCredentialResponse) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *AppCredentialResponse) HasState() bool`

HasState returns a boolean if a field has been set.

### GetName

`func (o *AppCredentialResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppCredentialResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppCredentialResponse) SetName(v string)`

SetName sets Name field to given value.


### GetSecret

`func (o *AppCredentialResponse) GetSecret() map[string]interface{}`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *AppCredentialResponse) GetSecretOk() (*map[string]interface{}, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *AppCredentialResponse) SetSecret(v map[string]interface{})`

SetSecret sets Secret field to given value.


### GetEditables

`func (o *AppCredentialResponse) GetEditables() map[string]interface{}`

GetEditables returns the Editables field if non-nil, zero value otherwise.

### GetEditablesOk

`func (o *AppCredentialResponse) GetEditablesOk() (*map[string]interface{}, bool)`

GetEditablesOk returns a tuple with the Editables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditables

`func (o *AppCredentialResponse) SetEditables(v map[string]interface{})`

SetEditables sets Editables field to given value.

### HasEditables

`func (o *AppCredentialResponse) HasEditables() bool`

HasEditables returns a boolean if a field has been set.

### GetType

`func (o *AppCredentialResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AppCredentialResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AppCredentialResponse) SetType(v string)`

SetType sets Type field to given value.


### GetPassphrase

`func (o *AppCredentialResponse) GetPassphrase() map[string]interface{}`

GetPassphrase returns the Passphrase field if non-nil, zero value otherwise.

### GetPassphraseOk

`func (o *AppCredentialResponse) GetPassphraseOk() (*map[string]interface{}, bool)`

GetPassphraseOk returns a tuple with the Passphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassphrase

`func (o *AppCredentialResponse) SetPassphrase(v map[string]interface{})`

SetPassphrase sets Passphrase field to given value.

### HasPassphrase

`func (o *AppCredentialResponse) HasPassphrase() bool`

HasPassphrase returns a boolean if a field has been set.

### GetDescription

`func (o *AppCredentialResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AppCredentialResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AppCredentialResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AppCredentialResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



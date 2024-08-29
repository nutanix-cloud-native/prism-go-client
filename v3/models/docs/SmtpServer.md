# SmtpServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] [default to "PLAIN"]
**EmailAddress** | **string** |  | 
**Server** | [**ClusterNetworkEntity**](ClusterNetworkEntity.md) |  | 

## Methods

### NewSmtpServer

`func NewSmtpServer(emailAddress string, server ClusterNetworkEntity, ) *SmtpServer`

NewSmtpServer instantiates a new SmtpServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmtpServerWithDefaults

`func NewSmtpServerWithDefaults() *SmtpServer`

NewSmtpServerWithDefaults instantiates a new SmtpServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SmtpServer) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SmtpServer) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SmtpServer) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SmtpServer) HasType() bool`

HasType returns a boolean if a field has been set.

### GetEmailAddress

`func (o *SmtpServer) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *SmtpServer) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *SmtpServer) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.


### GetServer

`func (o *SmtpServer) GetServer() ClusterNetworkEntity`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *SmtpServer) GetServerOk() (*ClusterNetworkEntity, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *SmtpServer) SetServer(v ClusterNetworkEntity)`

SetServer sets Server field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



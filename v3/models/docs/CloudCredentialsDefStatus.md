# CloudCredentialsDefStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | Pointer to **string** | The state of the entity. | [optional] [readonly] 
**MessageList** | Pointer to [**[]MessageResource**](MessageResource.md) |  | [optional] [readonly] 
**Id** | **int64** | ID to uniquely identify cloud credentials | 
**Resources** | [**CloudCredentialsResources**](CloudCredentialsResources.md) |  | 
**Name** | **string** | Credentials name. | 

## Methods

### NewCloudCredentialsDefStatus

`func NewCloudCredentialsDefStatus(id int64, resources CloudCredentialsResources, name string, ) *CloudCredentialsDefStatus`

NewCloudCredentialsDefStatus instantiates a new CloudCredentialsDefStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudCredentialsDefStatusWithDefaults

`func NewCloudCredentialsDefStatusWithDefaults() *CloudCredentialsDefStatus`

NewCloudCredentialsDefStatusWithDefaults instantiates a new CloudCredentialsDefStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *CloudCredentialsDefStatus) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CloudCredentialsDefStatus) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CloudCredentialsDefStatus) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *CloudCredentialsDefStatus) HasState() bool`

HasState returns a boolean if a field has been set.

### GetMessageList

`func (o *CloudCredentialsDefStatus) GetMessageList() []MessageResource`

GetMessageList returns the MessageList field if non-nil, zero value otherwise.

### GetMessageListOk

`func (o *CloudCredentialsDefStatus) GetMessageListOk() (*[]MessageResource, bool)`

GetMessageListOk returns a tuple with the MessageList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageList

`func (o *CloudCredentialsDefStatus) SetMessageList(v []MessageResource)`

SetMessageList sets MessageList field to given value.

### HasMessageList

`func (o *CloudCredentialsDefStatus) HasMessageList() bool`

HasMessageList returns a boolean if a field has been set.

### GetId

`func (o *CloudCredentialsDefStatus) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudCredentialsDefStatus) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudCredentialsDefStatus) SetId(v int64)`

SetId sets Id field to given value.


### GetResources

`func (o *CloudCredentialsDefStatus) GetResources() CloudCredentialsResources`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *CloudCredentialsDefStatus) GetResourcesOk() (*CloudCredentialsResources, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *CloudCredentialsDefStatus) SetResources(v CloudCredentialsResources)`

SetResources sets Resources field to given value.


### GetName

`func (o *CloudCredentialsDefStatus) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudCredentialsDefStatus) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudCredentialsDefStatus) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



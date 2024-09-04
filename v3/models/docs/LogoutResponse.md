# LogoutResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalLogoutUrlList** | Pointer to **[]string** | List of addition logout urls which should be called by UI | [optional] 

## Methods

### NewLogoutResponse

`func NewLogoutResponse() *LogoutResponse`

NewLogoutResponse instantiates a new LogoutResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogoutResponseWithDefaults

`func NewLogoutResponseWithDefaults() *LogoutResponse`

NewLogoutResponseWithDefaults instantiates a new LogoutResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalLogoutUrlList

`func (o *LogoutResponse) GetExternalLogoutUrlList() []string`

GetExternalLogoutUrlList returns the ExternalLogoutUrlList field if non-nil, zero value otherwise.

### GetExternalLogoutUrlListOk

`func (o *LogoutResponse) GetExternalLogoutUrlListOk() (*[]string, bool)`

GetExternalLogoutUrlListOk returns a tuple with the ExternalLogoutUrlList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalLogoutUrlList

`func (o *LogoutResponse) SetExternalLogoutUrlList(v []string)`

SetExternalLogoutUrlList sets ExternalLogoutUrlList field to given value.

### HasExternalLogoutUrlList

`func (o *LogoutResponse) HasExternalLogoutUrlList() bool`

HasExternalLogoutUrlList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



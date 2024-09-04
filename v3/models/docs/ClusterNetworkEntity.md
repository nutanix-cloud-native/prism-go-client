# ClusterNetworkEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credentials** | Pointer to [**Credentials**](Credentials.md) |  | [optional] 
**Name** | Pointer to **string** | Name for the network entity (optional) | [optional] 
**ProxyTypeList** | Pointer to **[]string** |  | [optional] 
**Address** | [**Address**](Address.md) |  | 

## Methods

### NewClusterNetworkEntity

`func NewClusterNetworkEntity(address Address, ) *ClusterNetworkEntity`

NewClusterNetworkEntity instantiates a new ClusterNetworkEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterNetworkEntityWithDefaults

`func NewClusterNetworkEntityWithDefaults() *ClusterNetworkEntity`

NewClusterNetworkEntityWithDefaults instantiates a new ClusterNetworkEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentials

`func (o *ClusterNetworkEntity) GetCredentials() Credentials`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *ClusterNetworkEntity) GetCredentialsOk() (*Credentials, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *ClusterNetworkEntity) SetCredentials(v Credentials)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *ClusterNetworkEntity) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetName

`func (o *ClusterNetworkEntity) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClusterNetworkEntity) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClusterNetworkEntity) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ClusterNetworkEntity) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProxyTypeList

`func (o *ClusterNetworkEntity) GetProxyTypeList() []string`

GetProxyTypeList returns the ProxyTypeList field if non-nil, zero value otherwise.

### GetProxyTypeListOk

`func (o *ClusterNetworkEntity) GetProxyTypeListOk() (*[]string, bool)`

GetProxyTypeListOk returns a tuple with the ProxyTypeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyTypeList

`func (o *ClusterNetworkEntity) SetProxyTypeList(v []string)`

SetProxyTypeList sets ProxyTypeList field to given value.

### HasProxyTypeList

`func (o *ClusterNetworkEntity) HasProxyTypeList() bool`

HasProxyTypeList returns a boolean if a field has been set.

### GetAddress

`func (o *ClusterNetworkEntity) GetAddress() Address`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ClusterNetworkEntity) GetAddressOk() (*Address, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ClusterNetworkEntity) SetAddress(v Address)`

SetAddress sets Address field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



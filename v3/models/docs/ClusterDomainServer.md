# ClusterDomainServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nameserver** | Pointer to **string** | The IP of the nameserver that can resolve the domain name. Must set when joining the domain.  | [optional] 
**Name** | **string** | Joined domain name. In &#39;put&#39; request, empty name will unjoin the cluster from current domain.  | 
**DomainCredentials** | Pointer to [**Credentials**](Credentials.md) |  | [optional] 

## Methods

### NewClusterDomainServer

`func NewClusterDomainServer(name string, ) *ClusterDomainServer`

NewClusterDomainServer instantiates a new ClusterDomainServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterDomainServerWithDefaults

`func NewClusterDomainServerWithDefaults() *ClusterDomainServer`

NewClusterDomainServerWithDefaults instantiates a new ClusterDomainServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNameserver

`func (o *ClusterDomainServer) GetNameserver() string`

GetNameserver returns the Nameserver field if non-nil, zero value otherwise.

### GetNameserverOk

`func (o *ClusterDomainServer) GetNameserverOk() (*string, bool)`

GetNameserverOk returns a tuple with the Nameserver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameserver

`func (o *ClusterDomainServer) SetNameserver(v string)`

SetNameserver sets Nameserver field to given value.

### HasNameserver

`func (o *ClusterDomainServer) HasNameserver() bool`

HasNameserver returns a boolean if a field has been set.

### GetName

`func (o *ClusterDomainServer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClusterDomainServer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClusterDomainServer) SetName(v string)`

SetName sets Name field to given value.


### GetDomainCredentials

`func (o *ClusterDomainServer) GetDomainCredentials() Credentials`

GetDomainCredentials returns the DomainCredentials field if non-nil, zero value otherwise.

### GetDomainCredentialsOk

`func (o *ClusterDomainServer) GetDomainCredentialsOk() (*Credentials, bool)`

GetDomainCredentialsOk returns a tuple with the DomainCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainCredentials

`func (o *ClusterDomainServer) SetDomainCredentials(v Credentials)`

SetDomainCredentials sets DomainCredentials field to given value.

### HasDomainCredentials

`func (o *ClusterDomainServer) HasDomainCredentials() bool`

HasDomainCredentials returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



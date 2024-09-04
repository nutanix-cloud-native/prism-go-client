# FailoverCluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ip** | Pointer to **string** | IP address of the failover cluster. | [optional] 
**Name** | Pointer to **string** | Name of the failover cluster. | [optional] 
**DomainCredential** | [**Credentials**](Credentials.md) |  | 

## Methods

### NewFailoverCluster

`func NewFailoverCluster(domainCredential Credentials, ) *FailoverCluster`

NewFailoverCluster instantiates a new FailoverCluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFailoverClusterWithDefaults

`func NewFailoverClusterWithDefaults() *FailoverCluster`

NewFailoverClusterWithDefaults instantiates a new FailoverCluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIp

`func (o *FailoverCluster) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *FailoverCluster) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *FailoverCluster) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *FailoverCluster) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetName

`func (o *FailoverCluster) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FailoverCluster) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FailoverCluster) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FailoverCluster) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDomainCredential

`func (o *FailoverCluster) GetDomainCredential() Credentials`

GetDomainCredential returns the DomainCredential field if non-nil, zero value otherwise.

### GetDomainCredentialOk

`func (o *FailoverCluster) GetDomainCredentialOk() (*Credentials, bool)`

GetDomainCredentialOk returns a tuple with the DomainCredential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainCredential

`func (o *FailoverCluster) SetDomainCredential(v Credentials)`

SetDomainCredential sets DomainCredential field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



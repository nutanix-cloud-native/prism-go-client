# DataCollectorConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterUuid** | Pointer to **string** | UUID of the cluster on which the Data Collector will be installed.  | [optional] 
**DataCollectorIpAddress** | Pointer to **string** | The IP address of collector.  | [optional] 
**DataCollectorPortNumber** | Pointer to **int32** | The port number of collector.  | [optional] 
**NetworkUuid** | Pointer to **string** | UUID of the managed network which will be used to deploy MSP cluster.  | [optional] 

## Methods

### NewDataCollectorConfig

`func NewDataCollectorConfig() *DataCollectorConfig`

NewDataCollectorConfig instantiates a new DataCollectorConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataCollectorConfigWithDefaults

`func NewDataCollectorConfigWithDefaults() *DataCollectorConfig`

NewDataCollectorConfigWithDefaults instantiates a new DataCollectorConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterUuid

`func (o *DataCollectorConfig) GetClusterUuid() string`

GetClusterUuid returns the ClusterUuid field if non-nil, zero value otherwise.

### GetClusterUuidOk

`func (o *DataCollectorConfig) GetClusterUuidOk() (*string, bool)`

GetClusterUuidOk returns a tuple with the ClusterUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterUuid

`func (o *DataCollectorConfig) SetClusterUuid(v string)`

SetClusterUuid sets ClusterUuid field to given value.

### HasClusterUuid

`func (o *DataCollectorConfig) HasClusterUuid() bool`

HasClusterUuid returns a boolean if a field has been set.

### GetDataCollectorIpAddress

`func (o *DataCollectorConfig) GetDataCollectorIpAddress() string`

GetDataCollectorIpAddress returns the DataCollectorIpAddress field if non-nil, zero value otherwise.

### GetDataCollectorIpAddressOk

`func (o *DataCollectorConfig) GetDataCollectorIpAddressOk() (*string, bool)`

GetDataCollectorIpAddressOk returns a tuple with the DataCollectorIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataCollectorIpAddress

`func (o *DataCollectorConfig) SetDataCollectorIpAddress(v string)`

SetDataCollectorIpAddress sets DataCollectorIpAddress field to given value.

### HasDataCollectorIpAddress

`func (o *DataCollectorConfig) HasDataCollectorIpAddress() bool`

HasDataCollectorIpAddress returns a boolean if a field has been set.

### GetDataCollectorPortNumber

`func (o *DataCollectorConfig) GetDataCollectorPortNumber() int32`

GetDataCollectorPortNumber returns the DataCollectorPortNumber field if non-nil, zero value otherwise.

### GetDataCollectorPortNumberOk

`func (o *DataCollectorConfig) GetDataCollectorPortNumberOk() (*int32, bool)`

GetDataCollectorPortNumberOk returns a tuple with the DataCollectorPortNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataCollectorPortNumber

`func (o *DataCollectorConfig) SetDataCollectorPortNumber(v int32)`

SetDataCollectorPortNumber sets DataCollectorPortNumber field to given value.

### HasDataCollectorPortNumber

`func (o *DataCollectorConfig) HasDataCollectorPortNumber() bool`

HasDataCollectorPortNumber returns a boolean if a field has been set.

### GetNetworkUuid

`func (o *DataCollectorConfig) GetNetworkUuid() string`

GetNetworkUuid returns the NetworkUuid field if non-nil, zero value otherwise.

### GetNetworkUuidOk

`func (o *DataCollectorConfig) GetNetworkUuidOk() (*string, bool)`

GetNetworkUuidOk returns a tuple with the NetworkUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkUuid

`func (o *DataCollectorConfig) SetNetworkUuid(v string)`

SetNetworkUuid sets NetworkUuid field to given value.

### HasNetworkUuid

`func (o *DataCollectorConfig) HasNetworkUuid() bool`

HasNetworkUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



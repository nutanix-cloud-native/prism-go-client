# AlertResourcesSourceEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterUuid** | Pointer to **string** | The UUID for the cluster that contained the source entity at the alert creation time.  | [optional] 
**Entity** | Pointer to [**EntityInfo**](EntityInfo.md) |  | [optional] 

## Methods

### NewAlertResourcesSourceEntity

`func NewAlertResourcesSourceEntity() *AlertResourcesSourceEntity`

NewAlertResourcesSourceEntity instantiates a new AlertResourcesSourceEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertResourcesSourceEntityWithDefaults

`func NewAlertResourcesSourceEntityWithDefaults() *AlertResourcesSourceEntity`

NewAlertResourcesSourceEntityWithDefaults instantiates a new AlertResourcesSourceEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterUuid

`func (o *AlertResourcesSourceEntity) GetClusterUuid() string`

GetClusterUuid returns the ClusterUuid field if non-nil, zero value otherwise.

### GetClusterUuidOk

`func (o *AlertResourcesSourceEntity) GetClusterUuidOk() (*string, bool)`

GetClusterUuidOk returns a tuple with the ClusterUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterUuid

`func (o *AlertResourcesSourceEntity) SetClusterUuid(v string)`

SetClusterUuid sets ClusterUuid field to given value.

### HasClusterUuid

`func (o *AlertResourcesSourceEntity) HasClusterUuid() bool`

HasClusterUuid returns a boolean if a field has been set.

### GetEntity

`func (o *AlertResourcesSourceEntity) GetEntity() EntityInfo`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *AlertResourcesSourceEntity) GetEntityOk() (*EntityInfo, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *AlertResourcesSourceEntity) SetEntity(v EntityInfo)`

SetEntity sets Entity field to given value.

### HasEntity

`func (o *AlertResourcesSourceEntity) HasEntity() bool`

HasEntity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



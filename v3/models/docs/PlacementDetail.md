# PlacementDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterIp** | Pointer to **string** | IP of the cluster. | [optional] 
**ClusterDrNetworkSegmentedVip** | Pointer to **string** | DR network segmentation virtual IP of the cluster. | [optional] 
**ClusterPort** | Pointer to **int32** | Port of the cluster. | [optional] 
**ClusterReference** | Pointer to [**ClusterReference**](ClusterReference.md) |  | [optional] 
**LtssServiceDetails** | Pointer to [**LtssServiceInfo**](LtssServiceInfo.md) |  | [optional] 

## Methods

### NewPlacementDetail

`func NewPlacementDetail() *PlacementDetail`

NewPlacementDetail instantiates a new PlacementDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlacementDetailWithDefaults

`func NewPlacementDetailWithDefaults() *PlacementDetail`

NewPlacementDetailWithDefaults instantiates a new PlacementDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterIp

`func (o *PlacementDetail) GetClusterIp() string`

GetClusterIp returns the ClusterIp field if non-nil, zero value otherwise.

### GetClusterIpOk

`func (o *PlacementDetail) GetClusterIpOk() (*string, bool)`

GetClusterIpOk returns a tuple with the ClusterIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterIp

`func (o *PlacementDetail) SetClusterIp(v string)`

SetClusterIp sets ClusterIp field to given value.

### HasClusterIp

`func (o *PlacementDetail) HasClusterIp() bool`

HasClusterIp returns a boolean if a field has been set.

### GetClusterDrNetworkSegmentedVip

`func (o *PlacementDetail) GetClusterDrNetworkSegmentedVip() string`

GetClusterDrNetworkSegmentedVip returns the ClusterDrNetworkSegmentedVip field if non-nil, zero value otherwise.

### GetClusterDrNetworkSegmentedVipOk

`func (o *PlacementDetail) GetClusterDrNetworkSegmentedVipOk() (*string, bool)`

GetClusterDrNetworkSegmentedVipOk returns a tuple with the ClusterDrNetworkSegmentedVip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterDrNetworkSegmentedVip

`func (o *PlacementDetail) SetClusterDrNetworkSegmentedVip(v string)`

SetClusterDrNetworkSegmentedVip sets ClusterDrNetworkSegmentedVip field to given value.

### HasClusterDrNetworkSegmentedVip

`func (o *PlacementDetail) HasClusterDrNetworkSegmentedVip() bool`

HasClusterDrNetworkSegmentedVip returns a boolean if a field has been set.

### GetClusterPort

`func (o *PlacementDetail) GetClusterPort() int32`

GetClusterPort returns the ClusterPort field if non-nil, zero value otherwise.

### GetClusterPortOk

`func (o *PlacementDetail) GetClusterPortOk() (*int32, bool)`

GetClusterPortOk returns a tuple with the ClusterPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterPort

`func (o *PlacementDetail) SetClusterPort(v int32)`

SetClusterPort sets ClusterPort field to given value.

### HasClusterPort

`func (o *PlacementDetail) HasClusterPort() bool`

HasClusterPort returns a boolean if a field has been set.

### GetClusterReference

`func (o *PlacementDetail) GetClusterReference() ClusterReference`

GetClusterReference returns the ClusterReference field if non-nil, zero value otherwise.

### GetClusterReferenceOk

`func (o *PlacementDetail) GetClusterReferenceOk() (*ClusterReference, bool)`

GetClusterReferenceOk returns a tuple with the ClusterReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterReference

`func (o *PlacementDetail) SetClusterReference(v ClusterReference)`

SetClusterReference sets ClusterReference field to given value.

### HasClusterReference

`func (o *PlacementDetail) HasClusterReference() bool`

HasClusterReference returns a boolean if a field has been set.

### GetLtssServiceDetails

`func (o *PlacementDetail) GetLtssServiceDetails() LtssServiceInfo`

GetLtssServiceDetails returns the LtssServiceDetails field if non-nil, zero value otherwise.

### GetLtssServiceDetailsOk

`func (o *PlacementDetail) GetLtssServiceDetailsOk() (*LtssServiceInfo, bool)`

GetLtssServiceDetailsOk returns a tuple with the LtssServiceDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLtssServiceDetails

`func (o *PlacementDetail) SetLtssServiceDetails(v LtssServiceInfo)`

SetLtssServiceDetails sets LtssServiceDetails field to given value.

### HasLtssServiceDetails

`func (o *PlacementDetail) HasLtssServiceDetails() bool`

HasLtssServiceDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



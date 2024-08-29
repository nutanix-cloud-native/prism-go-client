# NodeSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecommendedOnlineTimestampSecs** | Pointer to **int32** |  | [optional] 
**ToRemoved** | Pointer to **bool** | Indicate if the node is set for removal. | [optional] 
**Model** | Pointer to **string** |  | [optional] 
**ResourceSpec** | Pointer to [**GenericResourceSpec**](GenericResourceSpec.md) |  | [optional] 
**NumOfNodes** | Pointer to **int32** |  | [optional] 

## Methods

### NewNodeSpec

`func NewNodeSpec() *NodeSpec`

NewNodeSpec instantiates a new NodeSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeSpecWithDefaults

`func NewNodeSpecWithDefaults() *NodeSpec`

NewNodeSpecWithDefaults instantiates a new NodeSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecommendedOnlineTimestampSecs

`func (o *NodeSpec) GetRecommendedOnlineTimestampSecs() int32`

GetRecommendedOnlineTimestampSecs returns the RecommendedOnlineTimestampSecs field if non-nil, zero value otherwise.

### GetRecommendedOnlineTimestampSecsOk

`func (o *NodeSpec) GetRecommendedOnlineTimestampSecsOk() (*int32, bool)`

GetRecommendedOnlineTimestampSecsOk returns a tuple with the RecommendedOnlineTimestampSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendedOnlineTimestampSecs

`func (o *NodeSpec) SetRecommendedOnlineTimestampSecs(v int32)`

SetRecommendedOnlineTimestampSecs sets RecommendedOnlineTimestampSecs field to given value.

### HasRecommendedOnlineTimestampSecs

`func (o *NodeSpec) HasRecommendedOnlineTimestampSecs() bool`

HasRecommendedOnlineTimestampSecs returns a boolean if a field has been set.

### GetToRemoved

`func (o *NodeSpec) GetToRemoved() bool`

GetToRemoved returns the ToRemoved field if non-nil, zero value otherwise.

### GetToRemovedOk

`func (o *NodeSpec) GetToRemovedOk() (*bool, bool)`

GetToRemovedOk returns a tuple with the ToRemoved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToRemoved

`func (o *NodeSpec) SetToRemoved(v bool)`

SetToRemoved sets ToRemoved field to given value.

### HasToRemoved

`func (o *NodeSpec) HasToRemoved() bool`

HasToRemoved returns a boolean if a field has been set.

### GetModel

`func (o *NodeSpec) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *NodeSpec) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *NodeSpec) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *NodeSpec) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetResourceSpec

`func (o *NodeSpec) GetResourceSpec() GenericResourceSpec`

GetResourceSpec returns the ResourceSpec field if non-nil, zero value otherwise.

### GetResourceSpecOk

`func (o *NodeSpec) GetResourceSpecOk() (*GenericResourceSpec, bool)`

GetResourceSpecOk returns a tuple with the ResourceSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceSpec

`func (o *NodeSpec) SetResourceSpec(v GenericResourceSpec)`

SetResourceSpec sets ResourceSpec field to given value.

### HasResourceSpec

`func (o *NodeSpec) HasResourceSpec() bool`

HasResourceSpec returns a boolean if a field has been set.

### GetNumOfNodes

`func (o *NodeSpec) GetNumOfNodes() int32`

GetNumOfNodes returns the NumOfNodes field if non-nil, zero value otherwise.

### GetNumOfNodesOk

`func (o *NodeSpec) GetNumOfNodesOk() (*int32, bool)`

GetNumOfNodesOk returns a tuple with the NumOfNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumOfNodes

`func (o *NodeSpec) SetNumOfNodes(v int32)`

SetNumOfNodes sets NumOfNodes field to given value.

### HasNumOfNodes

`func (o *NodeSpec) HasNumOfNodes() bool`

HasNumOfNodes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# CaChainSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CaChain** | **string** | Content of CA chain | 
**Name** | **string** | The name of the CA Chain file | 

## Methods

### NewCaChainSpec

`func NewCaChainSpec(caChain string, name string, ) *CaChainSpec`

NewCaChainSpec instantiates a new CaChainSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCaChainSpecWithDefaults

`func NewCaChainSpecWithDefaults() *CaChainSpec`

NewCaChainSpecWithDefaults instantiates a new CaChainSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCaChain

`func (o *CaChainSpec) GetCaChain() string`

GetCaChain returns the CaChain field if non-nil, zero value otherwise.

### GetCaChainOk

`func (o *CaChainSpec) GetCaChainOk() (*string, bool)`

GetCaChainOk returns a tuple with the CaChain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaChain

`func (o *CaChainSpec) SetCaChain(v string)`

SetCaChain sets CaChain field to given value.


### GetName

`func (o *CaChainSpec) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CaChainSpec) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CaChainSpec) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



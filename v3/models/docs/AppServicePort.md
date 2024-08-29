# AppServicePort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Protocol** | **string** |  | 
**ExposedPort** | Pointer to **string** |  | [optional] 
**ExposedAddress** | Pointer to **string** |  | [optional] 
**TargetPort** | **string** |  | 
**EndpointName** | Pointer to **string** |  | [optional] 
**ContainerSpec** | Pointer to **map[string]interface{}** | Additional properties for k8s continaer spec | [optional] 

## Methods

### NewAppServicePort

`func NewAppServicePort(protocol string, targetPort string, ) *AppServicePort`

NewAppServicePort instantiates a new AppServicePort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppServicePortWithDefaults

`func NewAppServicePortWithDefaults() *AppServicePort`

NewAppServicePortWithDefaults instantiates a new AppServicePort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProtocol

`func (o *AppServicePort) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *AppServicePort) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *AppServicePort) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.


### GetExposedPort

`func (o *AppServicePort) GetExposedPort() string`

GetExposedPort returns the ExposedPort field if non-nil, zero value otherwise.

### GetExposedPortOk

`func (o *AppServicePort) GetExposedPortOk() (*string, bool)`

GetExposedPortOk returns a tuple with the ExposedPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExposedPort

`func (o *AppServicePort) SetExposedPort(v string)`

SetExposedPort sets ExposedPort field to given value.

### HasExposedPort

`func (o *AppServicePort) HasExposedPort() bool`

HasExposedPort returns a boolean if a field has been set.

### GetExposedAddress

`func (o *AppServicePort) GetExposedAddress() string`

GetExposedAddress returns the ExposedAddress field if non-nil, zero value otherwise.

### GetExposedAddressOk

`func (o *AppServicePort) GetExposedAddressOk() (*string, bool)`

GetExposedAddressOk returns a tuple with the ExposedAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExposedAddress

`func (o *AppServicePort) SetExposedAddress(v string)`

SetExposedAddress sets ExposedAddress field to given value.

### HasExposedAddress

`func (o *AppServicePort) HasExposedAddress() bool`

HasExposedAddress returns a boolean if a field has been set.

### GetTargetPort

`func (o *AppServicePort) GetTargetPort() string`

GetTargetPort returns the TargetPort field if non-nil, zero value otherwise.

### GetTargetPortOk

`func (o *AppServicePort) GetTargetPortOk() (*string, bool)`

GetTargetPortOk returns a tuple with the TargetPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetPort

`func (o *AppServicePort) SetTargetPort(v string)`

SetTargetPort sets TargetPort field to given value.


### GetEndpointName

`func (o *AppServicePort) GetEndpointName() string`

GetEndpointName returns the EndpointName field if non-nil, zero value otherwise.

### GetEndpointNameOk

`func (o *AppServicePort) GetEndpointNameOk() (*string, bool)`

GetEndpointNameOk returns a tuple with the EndpointName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointName

`func (o *AppServicePort) SetEndpointName(v string)`

SetEndpointName sets EndpointName field to given value.

### HasEndpointName

`func (o *AppServicePort) HasEndpointName() bool`

HasEndpointName returns a boolean if a field has been set.

### GetContainerSpec

`func (o *AppServicePort) GetContainerSpec() map[string]interface{}`

GetContainerSpec returns the ContainerSpec field if non-nil, zero value otherwise.

### GetContainerSpecOk

`func (o *AppServicePort) GetContainerSpecOk() (*map[string]interface{}, bool)`

GetContainerSpecOk returns a tuple with the ContainerSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerSpec

`func (o *AppServicePort) SetContainerSpec(v map[string]interface{})`

SetContainerSpec sets ContainerSpec field to given value.

### HasContainerSpec

`func (o *AppServicePort) HasContainerSpec() bool`

HasContainerSpec returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# CitrixConnectorConfigDetailsSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CitrixVmReferenceList** | Pointer to [**[]VmReference**](VmReference.md) | Reference to the list of vm ids registered with citrix cloud. | [optional] 
**ClientSecret** | Pointer to **string** | The client secret for the Citrix Cloud. | [optional] 
**CustomerId** | Pointer to **string** | The customer id registered with Citrix Cloud. | [optional] 
**ClientId** | Pointer to **string** | The client id for the Citrix Cloud. | [optional] 
**ResourceLocation** | Pointer to [**CitrixResourceLocationSpec**](CitrixResourceLocationSpec.md) |  | [optional] 

## Methods

### NewCitrixConnectorConfigDetailsSpec

`func NewCitrixConnectorConfigDetailsSpec() *CitrixConnectorConfigDetailsSpec`

NewCitrixConnectorConfigDetailsSpec instantiates a new CitrixConnectorConfigDetailsSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCitrixConnectorConfigDetailsSpecWithDefaults

`func NewCitrixConnectorConfigDetailsSpecWithDefaults() *CitrixConnectorConfigDetailsSpec`

NewCitrixConnectorConfigDetailsSpecWithDefaults instantiates a new CitrixConnectorConfigDetailsSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCitrixVmReferenceList

`func (o *CitrixConnectorConfigDetailsSpec) GetCitrixVmReferenceList() []VmReference`

GetCitrixVmReferenceList returns the CitrixVmReferenceList field if non-nil, zero value otherwise.

### GetCitrixVmReferenceListOk

`func (o *CitrixConnectorConfigDetailsSpec) GetCitrixVmReferenceListOk() (*[]VmReference, bool)`

GetCitrixVmReferenceListOk returns a tuple with the CitrixVmReferenceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCitrixVmReferenceList

`func (o *CitrixConnectorConfigDetailsSpec) SetCitrixVmReferenceList(v []VmReference)`

SetCitrixVmReferenceList sets CitrixVmReferenceList field to given value.

### HasCitrixVmReferenceList

`func (o *CitrixConnectorConfigDetailsSpec) HasCitrixVmReferenceList() bool`

HasCitrixVmReferenceList returns a boolean if a field has been set.

### GetClientSecret

`func (o *CitrixConnectorConfigDetailsSpec) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *CitrixConnectorConfigDetailsSpec) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *CitrixConnectorConfigDetailsSpec) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *CitrixConnectorConfigDetailsSpec) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetCustomerId

`func (o *CitrixConnectorConfigDetailsSpec) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *CitrixConnectorConfigDetailsSpec) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *CitrixConnectorConfigDetailsSpec) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *CitrixConnectorConfigDetailsSpec) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetClientId

`func (o *CitrixConnectorConfigDetailsSpec) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *CitrixConnectorConfigDetailsSpec) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *CitrixConnectorConfigDetailsSpec) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *CitrixConnectorConfigDetailsSpec) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetResourceLocation

`func (o *CitrixConnectorConfigDetailsSpec) GetResourceLocation() CitrixResourceLocationSpec`

GetResourceLocation returns the ResourceLocation field if non-nil, zero value otherwise.

### GetResourceLocationOk

`func (o *CitrixConnectorConfigDetailsSpec) GetResourceLocationOk() (*CitrixResourceLocationSpec, bool)`

GetResourceLocationOk returns a tuple with the ResourceLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceLocation

`func (o *CitrixConnectorConfigDetailsSpec) SetResourceLocation(v CitrixResourceLocationSpec)`

SetResourceLocation sets ResourceLocation field to given value.

### HasResourceLocation

`func (o *CitrixConnectorConfigDetailsSpec) HasResourceLocation() bool`

HasResourceLocation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# CategoryKeyStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Capabilities** | Pointer to [**Capabilities**](Capabilities.md) |  | [optional] 
**SystemDefined** | Pointer to **bool** | Specifying whether its a system defined category. | [optional] [readonly] 
**Description** | Pointer to **string** | Description of the category. | [optional] 
**ApiVersion** | Pointer to **string** | API Version of the Nutanix v3 API framework. | [optional] [readonly] [default to "3.1.0"]
**Name** | **string** | Name of the category. | 

## Methods

### NewCategoryKeyStatus

`func NewCategoryKeyStatus(name string, ) *CategoryKeyStatus`

NewCategoryKeyStatus instantiates a new CategoryKeyStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCategoryKeyStatusWithDefaults

`func NewCategoryKeyStatusWithDefaults() *CategoryKeyStatus`

NewCategoryKeyStatusWithDefaults instantiates a new CategoryKeyStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapabilities

`func (o *CategoryKeyStatus) GetCapabilities() Capabilities`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *CategoryKeyStatus) GetCapabilitiesOk() (*Capabilities, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *CategoryKeyStatus) SetCapabilities(v Capabilities)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *CategoryKeyStatus) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### GetSystemDefined

`func (o *CategoryKeyStatus) GetSystemDefined() bool`

GetSystemDefined returns the SystemDefined field if non-nil, zero value otherwise.

### GetSystemDefinedOk

`func (o *CategoryKeyStatus) GetSystemDefinedOk() (*bool, bool)`

GetSystemDefinedOk returns a tuple with the SystemDefined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemDefined

`func (o *CategoryKeyStatus) SetSystemDefined(v bool)`

SetSystemDefined sets SystemDefined field to given value.

### HasSystemDefined

`func (o *CategoryKeyStatus) HasSystemDefined() bool`

HasSystemDefined returns a boolean if a field has been set.

### GetDescription

`func (o *CategoryKeyStatus) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CategoryKeyStatus) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CategoryKeyStatus) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CategoryKeyStatus) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetApiVersion

`func (o *CategoryKeyStatus) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *CategoryKeyStatus) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *CategoryKeyStatus) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *CategoryKeyStatus) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetName

`func (o *CategoryKeyStatus) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CategoryKeyStatus) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CategoryKeyStatus) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



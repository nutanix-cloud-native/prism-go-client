# SupportCase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Description of the support case. | [optional] 
**Resources** | [**SupportCaseResources**](SupportCaseResources.md) |  | 
**Subject** | **string** | Subject of the support case. | 

## Methods

### NewSupportCase

`func NewSupportCase(resources SupportCaseResources, subject string, ) *SupportCase`

NewSupportCase instantiates a new SupportCase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupportCaseWithDefaults

`func NewSupportCaseWithDefaults() *SupportCase`

NewSupportCaseWithDefaults instantiates a new SupportCase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *SupportCase) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SupportCase) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SupportCase) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SupportCase) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetResources

`func (o *SupportCase) GetResources() SupportCaseResources`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *SupportCase) GetResourcesOk() (*SupportCaseResources, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *SupportCase) SetResources(v SupportCaseResources)`

SetResources sets Resources field to given value.


### GetSubject

`func (o *SupportCase) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *SupportCase) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *SupportCase) SetSubject(v string)`

SetSubject sets Subject field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



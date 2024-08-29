# AttachmentReferenceInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IscsiInitiatorNetworkId** | Pointer to **string** | Ip address of the external client. | [optional] 
**ClientSecret** | Pointer to **string** | Client secret for CHAP authentication. | [optional] 
**VmReference** | Pointer to [**VmReference**](VmReference.md) |  | [optional] 
**IscsiInitiatorName** | Pointer to **string** | Name of the iSCSI initiator of the workload outside Nutanix cluster.  | [optional] 

## Methods

### NewAttachmentReferenceInput

`func NewAttachmentReferenceInput() *AttachmentReferenceInput`

NewAttachmentReferenceInput instantiates a new AttachmentReferenceInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttachmentReferenceInputWithDefaults

`func NewAttachmentReferenceInputWithDefaults() *AttachmentReferenceInput`

NewAttachmentReferenceInputWithDefaults instantiates a new AttachmentReferenceInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIscsiInitiatorNetworkId

`func (o *AttachmentReferenceInput) GetIscsiInitiatorNetworkId() string`

GetIscsiInitiatorNetworkId returns the IscsiInitiatorNetworkId field if non-nil, zero value otherwise.

### GetIscsiInitiatorNetworkIdOk

`func (o *AttachmentReferenceInput) GetIscsiInitiatorNetworkIdOk() (*string, bool)`

GetIscsiInitiatorNetworkIdOk returns a tuple with the IscsiInitiatorNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIscsiInitiatorNetworkId

`func (o *AttachmentReferenceInput) SetIscsiInitiatorNetworkId(v string)`

SetIscsiInitiatorNetworkId sets IscsiInitiatorNetworkId field to given value.

### HasIscsiInitiatorNetworkId

`func (o *AttachmentReferenceInput) HasIscsiInitiatorNetworkId() bool`

HasIscsiInitiatorNetworkId returns a boolean if a field has been set.

### GetClientSecret

`func (o *AttachmentReferenceInput) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *AttachmentReferenceInput) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *AttachmentReferenceInput) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *AttachmentReferenceInput) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetVmReference

`func (o *AttachmentReferenceInput) GetVmReference() VmReference`

GetVmReference returns the VmReference field if non-nil, zero value otherwise.

### GetVmReferenceOk

`func (o *AttachmentReferenceInput) GetVmReferenceOk() (*VmReference, bool)`

GetVmReferenceOk returns a tuple with the VmReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmReference

`func (o *AttachmentReferenceInput) SetVmReference(v VmReference)`

SetVmReference sets VmReference field to given value.

### HasVmReference

`func (o *AttachmentReferenceInput) HasVmReference() bool`

HasVmReference returns a boolean if a field has been set.

### GetIscsiInitiatorName

`func (o *AttachmentReferenceInput) GetIscsiInitiatorName() string`

GetIscsiInitiatorName returns the IscsiInitiatorName field if non-nil, zero value otherwise.

### GetIscsiInitiatorNameOk

`func (o *AttachmentReferenceInput) GetIscsiInitiatorNameOk() (*string, bool)`

GetIscsiInitiatorNameOk returns a tuple with the IscsiInitiatorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIscsiInitiatorName

`func (o *AttachmentReferenceInput) SetIscsiInitiatorName(v string)`

SetIscsiInitiatorName sets IscsiInitiatorName field to given value.

### HasIscsiInitiatorName

`func (o *AttachmentReferenceInput) HasIscsiInitiatorName() bool`

HasIscsiInitiatorName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



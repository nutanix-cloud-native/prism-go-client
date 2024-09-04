# AttachmentReferenceOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IscsiInitiatorNetworkId** | Pointer to **string** | IPv4 address of the external client. | [optional] 
**EnabledAuthentications** | Pointer to **string** | Which authentication is enabled for client. | [optional] 
**VmReference** | Pointer to [**VmReference**](VmReference.md) |  | [optional] 
**IscsiInitiatorName** | Pointer to **string** | Name of the iSCSI initiator of the workload outside Nutanix cluster.  | [optional] 

## Methods

### NewAttachmentReferenceOutput

`func NewAttachmentReferenceOutput() *AttachmentReferenceOutput`

NewAttachmentReferenceOutput instantiates a new AttachmentReferenceOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttachmentReferenceOutputWithDefaults

`func NewAttachmentReferenceOutputWithDefaults() *AttachmentReferenceOutput`

NewAttachmentReferenceOutputWithDefaults instantiates a new AttachmentReferenceOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIscsiInitiatorNetworkId

`func (o *AttachmentReferenceOutput) GetIscsiInitiatorNetworkId() string`

GetIscsiInitiatorNetworkId returns the IscsiInitiatorNetworkId field if non-nil, zero value otherwise.

### GetIscsiInitiatorNetworkIdOk

`func (o *AttachmentReferenceOutput) GetIscsiInitiatorNetworkIdOk() (*string, bool)`

GetIscsiInitiatorNetworkIdOk returns a tuple with the IscsiInitiatorNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIscsiInitiatorNetworkId

`func (o *AttachmentReferenceOutput) SetIscsiInitiatorNetworkId(v string)`

SetIscsiInitiatorNetworkId sets IscsiInitiatorNetworkId field to given value.

### HasIscsiInitiatorNetworkId

`func (o *AttachmentReferenceOutput) HasIscsiInitiatorNetworkId() bool`

HasIscsiInitiatorNetworkId returns a boolean if a field has been set.

### GetEnabledAuthentications

`func (o *AttachmentReferenceOutput) GetEnabledAuthentications() string`

GetEnabledAuthentications returns the EnabledAuthentications field if non-nil, zero value otherwise.

### GetEnabledAuthenticationsOk

`func (o *AttachmentReferenceOutput) GetEnabledAuthenticationsOk() (*string, bool)`

GetEnabledAuthenticationsOk returns a tuple with the EnabledAuthentications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledAuthentications

`func (o *AttachmentReferenceOutput) SetEnabledAuthentications(v string)`

SetEnabledAuthentications sets EnabledAuthentications field to given value.

### HasEnabledAuthentications

`func (o *AttachmentReferenceOutput) HasEnabledAuthentications() bool`

HasEnabledAuthentications returns a boolean if a field has been set.

### GetVmReference

`func (o *AttachmentReferenceOutput) GetVmReference() VmReference`

GetVmReference returns the VmReference field if non-nil, zero value otherwise.

### GetVmReferenceOk

`func (o *AttachmentReferenceOutput) GetVmReferenceOk() (*VmReference, bool)`

GetVmReferenceOk returns a tuple with the VmReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmReference

`func (o *AttachmentReferenceOutput) SetVmReference(v VmReference)`

SetVmReference sets VmReference field to given value.

### HasVmReference

`func (o *AttachmentReferenceOutput) HasVmReference() bool`

HasVmReference returns a boolean if a field has been set.

### GetIscsiInitiatorName

`func (o *AttachmentReferenceOutput) GetIscsiInitiatorName() string`

GetIscsiInitiatorName returns the IscsiInitiatorName field if non-nil, zero value otherwise.

### GetIscsiInitiatorNameOk

`func (o *AttachmentReferenceOutput) GetIscsiInitiatorNameOk() (*string, bool)`

GetIscsiInitiatorNameOk returns a tuple with the IscsiInitiatorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIscsiInitiatorName

`func (o *AttachmentReferenceOutput) SetIscsiInitiatorName(v string)`

SetIscsiInitiatorName sets IscsiInitiatorName field to given value.

### HasIscsiInitiatorName

`func (o *AttachmentReferenceOutput) HasIscsiInitiatorName() bool`

HasIscsiInitiatorName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



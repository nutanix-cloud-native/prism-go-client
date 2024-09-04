# RecoverEntitiesEntityInfoListInnerVolumeGroupAttachmentListInnerVolumeGroupAttachmentInfoListInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthenticationType** | Pointer to **string** | Type of authentication protocol to be used.  | [optional] 
**ClientSecret** | Pointer to **string** | Client secret in case of CHAP authentication is required.  | [optional] 
**AttachmentType** | **string** | Mechanism to be used for Volume Group attachment. The allowed attachment types are IQN, HYPERVISOR. Specify IQN in case of \&quot;iSCSI Qualified Name\&quot; based attachments. In case of IQN based attachment, authentication_type, client_secret can be specified for client authentication. Specify HYPERVISOR for directly attaching Volumes to VM through the hypervisor, this is only supported when VM is running on AHV hypervisor.\&quot;  | [default to "IQN"]
**VolumeGroupReference** | [**VolumeGroupReference**](VolumeGroupReference.md) |  | 

## Methods

### NewRecoverEntitiesEntityInfoListInnerVolumeGroupAttachmentListInnerVolumeGroupAttachmentInfoListInner

`func NewRecoverEntitiesEntityInfoListInnerVolumeGroupAttachmentListInnerVolumeGroupAttachmentInfoListInner(attachmentType string, volumeGroupReference VolumeGroupReference, ) *RecoverEntitiesEntityInfoListInnerVolumeGroupAttachmentListInnerVolumeGroupAttachmentInfoListInner`

NewRecoverEntitiesEntityInfoListInnerVolumeGroupAttachmentListInnerVolumeGroupAttachmentInfoListInner instantiates a new RecoverEntitiesEntityInfoListInnerVolumeGroupAttachmentListInnerVolumeGroupAttachmentInfoListInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverEntitiesEntityInfoListInnerVolumeGroupAttachmentListInnerVolumeGroupAttachmentInfoListInnerWithDefaults

`func NewRecoverEntitiesEntityInfoListInnerVolumeGroupAttachmentListInnerVolumeGroupAttachmentInfoListInnerWithDefaults() *RecoverEntitiesEntityInfoListInnerVolumeGroupAttachmentListInnerVolumeGroupAttachmentInfoListInner`

NewRecoverEntitiesEntityInfoListInnerVolumeGroupAttachmentListInnerVolumeGroupAttachmentInfoListInnerWithDefaults instantiates a new RecoverEntitiesEntityInfoListInnerVolumeGroupAttachmentListInnerVolumeGroupAttachmentInfoListInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthenticationType

`func (o *RecoverEntitiesEntityInfoListInnerVolumeGroupAttachmentListInnerVolumeGroupAttachmentInfoListInner) GetAuthenticationType() string`

GetAuthenticationType returns the AuthenticationType field if non-nil, zero value otherwise.

### GetAuthenticationTypeOk

`func (o *RecoverEntitiesEntityInfoListInnerVolumeGroupAttachmentListInnerVolumeGroupAttachmentInfoListInner) GetAuthenticationTypeOk() (*string, bool)`

GetAuthenticationTypeOk returns a tuple with the AuthenticationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationType

`func (o *RecoverEntitiesEntityInfoListInnerVolumeGroupAttachmentListInnerVolumeGroupAttachmentInfoListInner) SetAuthenticationType(v string)`

SetAuthenticationType sets AuthenticationType field to given value.

### HasAuthenticationType

`func (o *RecoverEntitiesEntityInfoListInnerVolumeGroupAttachmentListInnerVolumeGroupAttachmentInfoListInner) HasAuthenticationType() bool`

HasAuthenticationType returns a boolean if a field has been set.

### GetClientSecret

`func (o *RecoverEntitiesEntityInfoListInnerVolumeGroupAttachmentListInnerVolumeGroupAttachmentInfoListInner) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *RecoverEntitiesEntityInfoListInnerVolumeGroupAttachmentListInnerVolumeGroupAttachmentInfoListInner) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *RecoverEntitiesEntityInfoListInnerVolumeGroupAttachmentListInnerVolumeGroupAttachmentInfoListInner) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *RecoverEntitiesEntityInfoListInnerVolumeGroupAttachmentListInnerVolumeGroupAttachmentInfoListInner) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetAttachmentType

`func (o *RecoverEntitiesEntityInfoListInnerVolumeGroupAttachmentListInnerVolumeGroupAttachmentInfoListInner) GetAttachmentType() string`

GetAttachmentType returns the AttachmentType field if non-nil, zero value otherwise.

### GetAttachmentTypeOk

`func (o *RecoverEntitiesEntityInfoListInnerVolumeGroupAttachmentListInnerVolumeGroupAttachmentInfoListInner) GetAttachmentTypeOk() (*string, bool)`

GetAttachmentTypeOk returns a tuple with the AttachmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachmentType

`func (o *RecoverEntitiesEntityInfoListInnerVolumeGroupAttachmentListInnerVolumeGroupAttachmentInfoListInner) SetAttachmentType(v string)`

SetAttachmentType sets AttachmentType field to given value.


### GetVolumeGroupReference

`func (o *RecoverEntitiesEntityInfoListInnerVolumeGroupAttachmentListInnerVolumeGroupAttachmentInfoListInner) GetVolumeGroupReference() VolumeGroupReference`

GetVolumeGroupReference returns the VolumeGroupReference field if non-nil, zero value otherwise.

### GetVolumeGroupReferenceOk

`func (o *RecoverEntitiesEntityInfoListInnerVolumeGroupAttachmentListInnerVolumeGroupAttachmentInfoListInner) GetVolumeGroupReferenceOk() (*VolumeGroupReference, bool)`

GetVolumeGroupReferenceOk returns a tuple with the VolumeGroupReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeGroupReference

`func (o *RecoverEntitiesEntityInfoListInnerVolumeGroupAttachmentListInnerVolumeGroupAttachmentInfoListInner) SetVolumeGroupReference(v VolumeGroupReference)`

SetVolumeGroupReference sets VolumeGroupReference field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



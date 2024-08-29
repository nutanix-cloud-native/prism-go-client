# ProjectInternalAccessControlPolicyListInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operation** | **string** | Indicates the action(add, delete, update) | 
**Acp** | [**AccessControlPolicyInput**](AccessControlPolicyInput.md) |  | 
**Metadata** | [**ProjectInternalAccessControlPolicyListInnerMetadata**](ProjectInternalAccessControlPolicyListInnerMetadata.md) |  | 

## Methods

### NewProjectInternalAccessControlPolicyListInner

`func NewProjectInternalAccessControlPolicyListInner(operation string, acp AccessControlPolicyInput, metadata ProjectInternalAccessControlPolicyListInnerMetadata, ) *ProjectInternalAccessControlPolicyListInner`

NewProjectInternalAccessControlPolicyListInner instantiates a new ProjectInternalAccessControlPolicyListInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectInternalAccessControlPolicyListInnerWithDefaults

`func NewProjectInternalAccessControlPolicyListInnerWithDefaults() *ProjectInternalAccessControlPolicyListInner`

NewProjectInternalAccessControlPolicyListInnerWithDefaults instantiates a new ProjectInternalAccessControlPolicyListInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperation

`func (o *ProjectInternalAccessControlPolicyListInner) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *ProjectInternalAccessControlPolicyListInner) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *ProjectInternalAccessControlPolicyListInner) SetOperation(v string)`

SetOperation sets Operation field to given value.


### GetAcp

`func (o *ProjectInternalAccessControlPolicyListInner) GetAcp() AccessControlPolicyInput`

GetAcp returns the Acp field if non-nil, zero value otherwise.

### GetAcpOk

`func (o *ProjectInternalAccessControlPolicyListInner) GetAcpOk() (*AccessControlPolicyInput, bool)`

GetAcpOk returns a tuple with the Acp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcp

`func (o *ProjectInternalAccessControlPolicyListInner) SetAcp(v AccessControlPolicyInput)`

SetAcp sets Acp field to given value.


### GetMetadata

`func (o *ProjectInternalAccessControlPolicyListInner) GetMetadata() ProjectInternalAccessControlPolicyListInnerMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ProjectInternalAccessControlPolicyListInner) GetMetadataOk() (*ProjectInternalAccessControlPolicyListInnerMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ProjectInternalAccessControlPolicyListInner) SetMetadata(v ProjectInternalAccessControlPolicyListInnerMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



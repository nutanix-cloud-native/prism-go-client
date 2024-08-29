# BlueprintDeploymentInputUpload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PercentFaultTolerance** | Pointer to **int64** |  | [optional] 
**FaultDomainScope** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ActionList** | Pointer to [**[]AppActionInputUpload**](AppActionInputUpload.md) | List of references to action | [optional] 
**PackageLocalReferenceList** | Pointer to [**[]AppPackageReferenceUpload**](AppPackageReferenceUpload.md) | List of references for the packages | [optional] 
**PublishedServiceLocalReferenceList** | Pointer to [**[]AppPublishedServiceReferenceUpload**](AppPublishedServiceReferenceUpload.md) | List of references for published services | [optional] 
**UUID** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**Editables** | Pointer to **map[string]interface{}** | Runtime editable attributes for this entity. The structure for this is a dictionary. The keys in this dictionary should be the name of the attribute on the entity. If the attribute is editable, the value should be true, else false. If the attribute is a nested dictionary, the value can contain a nested dictionary with the same key value structure described above.  | [optional] 
**DependsOnList** | Pointer to [**[]EntityReference**](EntityReference.md) |  | [optional] 
**MaxReplicas** | **string** | Maximum replicas for the deployment. | [default to "1"]
**Type** | Pointer to **string** |  | [optional] [default to "GREENFIELD"]
**SubstrateLocalReference** | [**AppSubstrateReferenceUpload**](AppSubstrateReferenceUpload.md) |  | 
**VariableList** | Pointer to [**[]AppVariableInputUpload**](AppVariableInputUpload.md) |  | [optional] 
**MinReplicas** | **string** | Minimum replicas for the deployment. | [default to "1"]
**Options** | Pointer to **map[string]interface{}** | Additional deployment options | [optional] 
**NumFaultTolerance** | Pointer to **int64** |  | [optional] 
**BrownfieldInstanceList** | Pointer to [**[]BrownfieldInstanceInput**](BrownfieldInstanceInput.md) | brownfield map | [optional] 

## Methods

### NewBlueprintDeploymentInputUpload

`func NewBlueprintDeploymentInputUpload(name string, maxReplicas string, substrateLocalReference AppSubstrateReferenceUpload, minReplicas string, ) *BlueprintDeploymentInputUpload`

NewBlueprintDeploymentInputUpload instantiates a new BlueprintDeploymentInputUpload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintDeploymentInputUploadWithDefaults

`func NewBlueprintDeploymentInputUploadWithDefaults() *BlueprintDeploymentInputUpload`

NewBlueprintDeploymentInputUploadWithDefaults instantiates a new BlueprintDeploymentInputUpload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPercentFaultTolerance

`func (o *BlueprintDeploymentInputUpload) GetPercentFaultTolerance() int64`

GetPercentFaultTolerance returns the PercentFaultTolerance field if non-nil, zero value otherwise.

### GetPercentFaultToleranceOk

`func (o *BlueprintDeploymentInputUpload) GetPercentFaultToleranceOk() (*int64, bool)`

GetPercentFaultToleranceOk returns a tuple with the PercentFaultTolerance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentFaultTolerance

`func (o *BlueprintDeploymentInputUpload) SetPercentFaultTolerance(v int64)`

SetPercentFaultTolerance sets PercentFaultTolerance field to given value.

### HasPercentFaultTolerance

`func (o *BlueprintDeploymentInputUpload) HasPercentFaultTolerance() bool`

HasPercentFaultTolerance returns a boolean if a field has been set.

### GetFaultDomainScope

`func (o *BlueprintDeploymentInputUpload) GetFaultDomainScope() string`

GetFaultDomainScope returns the FaultDomainScope field if non-nil, zero value otherwise.

### GetFaultDomainScopeOk

`func (o *BlueprintDeploymentInputUpload) GetFaultDomainScopeOk() (*string, bool)`

GetFaultDomainScopeOk returns a tuple with the FaultDomainScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaultDomainScope

`func (o *BlueprintDeploymentInputUpload) SetFaultDomainScope(v string)`

SetFaultDomainScope sets FaultDomainScope field to given value.

### HasFaultDomainScope

`func (o *BlueprintDeploymentInputUpload) HasFaultDomainScope() bool`

HasFaultDomainScope returns a boolean if a field has been set.

### GetDescription

`func (o *BlueprintDeploymentInputUpload) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BlueprintDeploymentInputUpload) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BlueprintDeploymentInputUpload) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BlueprintDeploymentInputUpload) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetActionList

`func (o *BlueprintDeploymentInputUpload) GetActionList() []AppActionInputUpload`

GetActionList returns the ActionList field if non-nil, zero value otherwise.

### GetActionListOk

`func (o *BlueprintDeploymentInputUpload) GetActionListOk() (*[]AppActionInputUpload, bool)`

GetActionListOk returns a tuple with the ActionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionList

`func (o *BlueprintDeploymentInputUpload) SetActionList(v []AppActionInputUpload)`

SetActionList sets ActionList field to given value.

### HasActionList

`func (o *BlueprintDeploymentInputUpload) HasActionList() bool`

HasActionList returns a boolean if a field has been set.

### GetPackageLocalReferenceList

`func (o *BlueprintDeploymentInputUpload) GetPackageLocalReferenceList() []AppPackageReferenceUpload`

GetPackageLocalReferenceList returns the PackageLocalReferenceList field if non-nil, zero value otherwise.

### GetPackageLocalReferenceListOk

`func (o *BlueprintDeploymentInputUpload) GetPackageLocalReferenceListOk() (*[]AppPackageReferenceUpload, bool)`

GetPackageLocalReferenceListOk returns a tuple with the PackageLocalReferenceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageLocalReferenceList

`func (o *BlueprintDeploymentInputUpload) SetPackageLocalReferenceList(v []AppPackageReferenceUpload)`

SetPackageLocalReferenceList sets PackageLocalReferenceList field to given value.

### HasPackageLocalReferenceList

`func (o *BlueprintDeploymentInputUpload) HasPackageLocalReferenceList() bool`

HasPackageLocalReferenceList returns a boolean if a field has been set.

### GetPublishedServiceLocalReferenceList

`func (o *BlueprintDeploymentInputUpload) GetPublishedServiceLocalReferenceList() []AppPublishedServiceReferenceUpload`

GetPublishedServiceLocalReferenceList returns the PublishedServiceLocalReferenceList field if non-nil, zero value otherwise.

### GetPublishedServiceLocalReferenceListOk

`func (o *BlueprintDeploymentInputUpload) GetPublishedServiceLocalReferenceListOk() (*[]AppPublishedServiceReferenceUpload, bool)`

GetPublishedServiceLocalReferenceListOk returns a tuple with the PublishedServiceLocalReferenceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedServiceLocalReferenceList

`func (o *BlueprintDeploymentInputUpload) SetPublishedServiceLocalReferenceList(v []AppPublishedServiceReferenceUpload)`

SetPublishedServiceLocalReferenceList sets PublishedServiceLocalReferenceList field to given value.

### HasPublishedServiceLocalReferenceList

`func (o *BlueprintDeploymentInputUpload) HasPublishedServiceLocalReferenceList() bool`

HasPublishedServiceLocalReferenceList returns a boolean if a field has been set.

### GetUUID

`func (o *BlueprintDeploymentInputUpload) GetUUID() string`

GetUUID returns the UUID field if non-nil, zero value otherwise.

### GetUUIDOk

`func (o *BlueprintDeploymentInputUpload) GetUUIDOk() (*string, bool)`

GetUUIDOk returns a tuple with the UUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUUID

`func (o *BlueprintDeploymentInputUpload) SetUUID(v string)`

SetUUID sets UUID field to given value.

### HasUUID

`func (o *BlueprintDeploymentInputUpload) HasUUID() bool`

HasUUID returns a boolean if a field has been set.

### GetName

`func (o *BlueprintDeploymentInputUpload) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BlueprintDeploymentInputUpload) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BlueprintDeploymentInputUpload) SetName(v string)`

SetName sets Name field to given value.


### GetEditables

`func (o *BlueprintDeploymentInputUpload) GetEditables() map[string]interface{}`

GetEditables returns the Editables field if non-nil, zero value otherwise.

### GetEditablesOk

`func (o *BlueprintDeploymentInputUpload) GetEditablesOk() (*map[string]interface{}, bool)`

GetEditablesOk returns a tuple with the Editables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditables

`func (o *BlueprintDeploymentInputUpload) SetEditables(v map[string]interface{})`

SetEditables sets Editables field to given value.

### HasEditables

`func (o *BlueprintDeploymentInputUpload) HasEditables() bool`

HasEditables returns a boolean if a field has been set.

### GetDependsOnList

`func (o *BlueprintDeploymentInputUpload) GetDependsOnList() []EntityReference`

GetDependsOnList returns the DependsOnList field if non-nil, zero value otherwise.

### GetDependsOnListOk

`func (o *BlueprintDeploymentInputUpload) GetDependsOnListOk() (*[]EntityReference, bool)`

GetDependsOnListOk returns a tuple with the DependsOnList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependsOnList

`func (o *BlueprintDeploymentInputUpload) SetDependsOnList(v []EntityReference)`

SetDependsOnList sets DependsOnList field to given value.

### HasDependsOnList

`func (o *BlueprintDeploymentInputUpload) HasDependsOnList() bool`

HasDependsOnList returns a boolean if a field has been set.

### GetMaxReplicas

`func (o *BlueprintDeploymentInputUpload) GetMaxReplicas() string`

GetMaxReplicas returns the MaxReplicas field if non-nil, zero value otherwise.

### GetMaxReplicasOk

`func (o *BlueprintDeploymentInputUpload) GetMaxReplicasOk() (*string, bool)`

GetMaxReplicasOk returns a tuple with the MaxReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxReplicas

`func (o *BlueprintDeploymentInputUpload) SetMaxReplicas(v string)`

SetMaxReplicas sets MaxReplicas field to given value.


### GetType

`func (o *BlueprintDeploymentInputUpload) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BlueprintDeploymentInputUpload) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BlueprintDeploymentInputUpload) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *BlueprintDeploymentInputUpload) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSubstrateLocalReference

`func (o *BlueprintDeploymentInputUpload) GetSubstrateLocalReference() AppSubstrateReferenceUpload`

GetSubstrateLocalReference returns the SubstrateLocalReference field if non-nil, zero value otherwise.

### GetSubstrateLocalReferenceOk

`func (o *BlueprintDeploymentInputUpload) GetSubstrateLocalReferenceOk() (*AppSubstrateReferenceUpload, bool)`

GetSubstrateLocalReferenceOk returns a tuple with the SubstrateLocalReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubstrateLocalReference

`func (o *BlueprintDeploymentInputUpload) SetSubstrateLocalReference(v AppSubstrateReferenceUpload)`

SetSubstrateLocalReference sets SubstrateLocalReference field to given value.


### GetVariableList

`func (o *BlueprintDeploymentInputUpload) GetVariableList() []AppVariableInputUpload`

GetVariableList returns the VariableList field if non-nil, zero value otherwise.

### GetVariableListOk

`func (o *BlueprintDeploymentInputUpload) GetVariableListOk() (*[]AppVariableInputUpload, bool)`

GetVariableListOk returns a tuple with the VariableList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableList

`func (o *BlueprintDeploymentInputUpload) SetVariableList(v []AppVariableInputUpload)`

SetVariableList sets VariableList field to given value.

### HasVariableList

`func (o *BlueprintDeploymentInputUpload) HasVariableList() bool`

HasVariableList returns a boolean if a field has been set.

### GetMinReplicas

`func (o *BlueprintDeploymentInputUpload) GetMinReplicas() string`

GetMinReplicas returns the MinReplicas field if non-nil, zero value otherwise.

### GetMinReplicasOk

`func (o *BlueprintDeploymentInputUpload) GetMinReplicasOk() (*string, bool)`

GetMinReplicasOk returns a tuple with the MinReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinReplicas

`func (o *BlueprintDeploymentInputUpload) SetMinReplicas(v string)`

SetMinReplicas sets MinReplicas field to given value.


### GetOptions

`func (o *BlueprintDeploymentInputUpload) GetOptions() map[string]interface{}`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *BlueprintDeploymentInputUpload) GetOptionsOk() (*map[string]interface{}, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *BlueprintDeploymentInputUpload) SetOptions(v map[string]interface{})`

SetOptions sets Options field to given value.

### HasOptions

`func (o *BlueprintDeploymentInputUpload) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetNumFaultTolerance

`func (o *BlueprintDeploymentInputUpload) GetNumFaultTolerance() int64`

GetNumFaultTolerance returns the NumFaultTolerance field if non-nil, zero value otherwise.

### GetNumFaultToleranceOk

`func (o *BlueprintDeploymentInputUpload) GetNumFaultToleranceOk() (*int64, bool)`

GetNumFaultToleranceOk returns a tuple with the NumFaultTolerance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumFaultTolerance

`func (o *BlueprintDeploymentInputUpload) SetNumFaultTolerance(v int64)`

SetNumFaultTolerance sets NumFaultTolerance field to given value.

### HasNumFaultTolerance

`func (o *BlueprintDeploymentInputUpload) HasNumFaultTolerance() bool`

HasNumFaultTolerance returns a boolean if a field has been set.

### GetBrownfieldInstanceList

`func (o *BlueprintDeploymentInputUpload) GetBrownfieldInstanceList() []BrownfieldInstanceInput`

GetBrownfieldInstanceList returns the BrownfieldInstanceList field if non-nil, zero value otherwise.

### GetBrownfieldInstanceListOk

`func (o *BlueprintDeploymentInputUpload) GetBrownfieldInstanceListOk() (*[]BrownfieldInstanceInput, bool)`

GetBrownfieldInstanceListOk returns a tuple with the BrownfieldInstanceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrownfieldInstanceList

`func (o *BlueprintDeploymentInputUpload) SetBrownfieldInstanceList(v []BrownfieldInstanceInput)`

SetBrownfieldInstanceList sets BrownfieldInstanceList field to given value.

### HasBrownfieldInstanceList

`func (o *BlueprintDeploymentInputUpload) HasBrownfieldInstanceList() bool`

HasBrownfieldInstanceList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



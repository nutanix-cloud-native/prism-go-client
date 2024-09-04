# BlueprintDeploymentInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PercentFaultTolerance** | Pointer to **int64** |  | [optional] 
**FaultDomainScope** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ActionList** | Pointer to [**[]AppActionInput**](AppActionInput.md) | List of references to action | [optional] 
**PackageLocalReferenceList** | Pointer to [**[]AppPackageReference**](AppPackageReference.md) | List of references for the packages | [optional] 
**PublishedServiceLocalReferenceList** | Pointer to [**[]AppPublishedServiceReference**](AppPublishedServiceReference.md) | List of references for published services | [optional] 
**UUID** | **string** |  | 
**Name** | **string** |  | 
**Editables** | Pointer to **map[string]interface{}** | Runtime editable attributes for this entity. The structure for this is a dictionary. The keys in this dictionary should be the name of the attribute on the entity. If the attribute is editable, the value should be true, else false. If the attribute is a nested dictionary, the value can contain a nested dictionary with the same key value structure described above.  | [optional] 
**DependsOnList** | Pointer to [**[]EntityReference**](EntityReference.md) |  | [optional] 
**MaxReplicas** | **string** | Maximum replicas for the deployment. | [default to "1"]
**Type** | Pointer to **string** |  | [optional] [default to "GREENFIELD"]
**SubstrateLocalReference** | [**AppSubstrateReference**](AppSubstrateReference.md) |  | 
**VariableList** | Pointer to [**[]AppVariableInput**](AppVariableInput.md) |  | [optional] 
**MinReplicas** | **string** | Minimum replicas for the deployment. | [default to "1"]
**Options** | Pointer to **map[string]interface{}** | Additional deployment options | [optional] 
**NumFaultTolerance** | Pointer to **int64** |  | [optional] 
**BrownfieldInstanceList** | Pointer to [**[]BrownfieldInstanceInput**](BrownfieldInstanceInput.md) | brownfield map | [optional] 

## Methods

### NewBlueprintDeploymentInput

`func NewBlueprintDeploymentInput(uUID string, name string, maxReplicas string, substrateLocalReference AppSubstrateReference, minReplicas string, ) *BlueprintDeploymentInput`

NewBlueprintDeploymentInput instantiates a new BlueprintDeploymentInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintDeploymentInputWithDefaults

`func NewBlueprintDeploymentInputWithDefaults() *BlueprintDeploymentInput`

NewBlueprintDeploymentInputWithDefaults instantiates a new BlueprintDeploymentInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPercentFaultTolerance

`func (o *BlueprintDeploymentInput) GetPercentFaultTolerance() int64`

GetPercentFaultTolerance returns the PercentFaultTolerance field if non-nil, zero value otherwise.

### GetPercentFaultToleranceOk

`func (o *BlueprintDeploymentInput) GetPercentFaultToleranceOk() (*int64, bool)`

GetPercentFaultToleranceOk returns a tuple with the PercentFaultTolerance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentFaultTolerance

`func (o *BlueprintDeploymentInput) SetPercentFaultTolerance(v int64)`

SetPercentFaultTolerance sets PercentFaultTolerance field to given value.

### HasPercentFaultTolerance

`func (o *BlueprintDeploymentInput) HasPercentFaultTolerance() bool`

HasPercentFaultTolerance returns a boolean if a field has been set.

### GetFaultDomainScope

`func (o *BlueprintDeploymentInput) GetFaultDomainScope() string`

GetFaultDomainScope returns the FaultDomainScope field if non-nil, zero value otherwise.

### GetFaultDomainScopeOk

`func (o *BlueprintDeploymentInput) GetFaultDomainScopeOk() (*string, bool)`

GetFaultDomainScopeOk returns a tuple with the FaultDomainScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaultDomainScope

`func (o *BlueprintDeploymentInput) SetFaultDomainScope(v string)`

SetFaultDomainScope sets FaultDomainScope field to given value.

### HasFaultDomainScope

`func (o *BlueprintDeploymentInput) HasFaultDomainScope() bool`

HasFaultDomainScope returns a boolean if a field has been set.

### GetDescription

`func (o *BlueprintDeploymentInput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BlueprintDeploymentInput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BlueprintDeploymentInput) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BlueprintDeploymentInput) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetActionList

`func (o *BlueprintDeploymentInput) GetActionList() []AppActionInput`

GetActionList returns the ActionList field if non-nil, zero value otherwise.

### GetActionListOk

`func (o *BlueprintDeploymentInput) GetActionListOk() (*[]AppActionInput, bool)`

GetActionListOk returns a tuple with the ActionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionList

`func (o *BlueprintDeploymentInput) SetActionList(v []AppActionInput)`

SetActionList sets ActionList field to given value.

### HasActionList

`func (o *BlueprintDeploymentInput) HasActionList() bool`

HasActionList returns a boolean if a field has been set.

### GetPackageLocalReferenceList

`func (o *BlueprintDeploymentInput) GetPackageLocalReferenceList() []AppPackageReference`

GetPackageLocalReferenceList returns the PackageLocalReferenceList field if non-nil, zero value otherwise.

### GetPackageLocalReferenceListOk

`func (o *BlueprintDeploymentInput) GetPackageLocalReferenceListOk() (*[]AppPackageReference, bool)`

GetPackageLocalReferenceListOk returns a tuple with the PackageLocalReferenceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageLocalReferenceList

`func (o *BlueprintDeploymentInput) SetPackageLocalReferenceList(v []AppPackageReference)`

SetPackageLocalReferenceList sets PackageLocalReferenceList field to given value.

### HasPackageLocalReferenceList

`func (o *BlueprintDeploymentInput) HasPackageLocalReferenceList() bool`

HasPackageLocalReferenceList returns a boolean if a field has been set.

### GetPublishedServiceLocalReferenceList

`func (o *BlueprintDeploymentInput) GetPublishedServiceLocalReferenceList() []AppPublishedServiceReference`

GetPublishedServiceLocalReferenceList returns the PublishedServiceLocalReferenceList field if non-nil, zero value otherwise.

### GetPublishedServiceLocalReferenceListOk

`func (o *BlueprintDeploymentInput) GetPublishedServiceLocalReferenceListOk() (*[]AppPublishedServiceReference, bool)`

GetPublishedServiceLocalReferenceListOk returns a tuple with the PublishedServiceLocalReferenceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedServiceLocalReferenceList

`func (o *BlueprintDeploymentInput) SetPublishedServiceLocalReferenceList(v []AppPublishedServiceReference)`

SetPublishedServiceLocalReferenceList sets PublishedServiceLocalReferenceList field to given value.

### HasPublishedServiceLocalReferenceList

`func (o *BlueprintDeploymentInput) HasPublishedServiceLocalReferenceList() bool`

HasPublishedServiceLocalReferenceList returns a boolean if a field has been set.

### GetUUID

`func (o *BlueprintDeploymentInput) GetUUID() string`

GetUUID returns the UUID field if non-nil, zero value otherwise.

### GetUUIDOk

`func (o *BlueprintDeploymentInput) GetUUIDOk() (*string, bool)`

GetUUIDOk returns a tuple with the UUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUUID

`func (o *BlueprintDeploymentInput) SetUUID(v string)`

SetUUID sets UUID field to given value.


### GetName

`func (o *BlueprintDeploymentInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BlueprintDeploymentInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BlueprintDeploymentInput) SetName(v string)`

SetName sets Name field to given value.


### GetEditables

`func (o *BlueprintDeploymentInput) GetEditables() map[string]interface{}`

GetEditables returns the Editables field if non-nil, zero value otherwise.

### GetEditablesOk

`func (o *BlueprintDeploymentInput) GetEditablesOk() (*map[string]interface{}, bool)`

GetEditablesOk returns a tuple with the Editables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditables

`func (o *BlueprintDeploymentInput) SetEditables(v map[string]interface{})`

SetEditables sets Editables field to given value.

### HasEditables

`func (o *BlueprintDeploymentInput) HasEditables() bool`

HasEditables returns a boolean if a field has been set.

### GetDependsOnList

`func (o *BlueprintDeploymentInput) GetDependsOnList() []EntityReference`

GetDependsOnList returns the DependsOnList field if non-nil, zero value otherwise.

### GetDependsOnListOk

`func (o *BlueprintDeploymentInput) GetDependsOnListOk() (*[]EntityReference, bool)`

GetDependsOnListOk returns a tuple with the DependsOnList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependsOnList

`func (o *BlueprintDeploymentInput) SetDependsOnList(v []EntityReference)`

SetDependsOnList sets DependsOnList field to given value.

### HasDependsOnList

`func (o *BlueprintDeploymentInput) HasDependsOnList() bool`

HasDependsOnList returns a boolean if a field has been set.

### GetMaxReplicas

`func (o *BlueprintDeploymentInput) GetMaxReplicas() string`

GetMaxReplicas returns the MaxReplicas field if non-nil, zero value otherwise.

### GetMaxReplicasOk

`func (o *BlueprintDeploymentInput) GetMaxReplicasOk() (*string, bool)`

GetMaxReplicasOk returns a tuple with the MaxReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxReplicas

`func (o *BlueprintDeploymentInput) SetMaxReplicas(v string)`

SetMaxReplicas sets MaxReplicas field to given value.


### GetType

`func (o *BlueprintDeploymentInput) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BlueprintDeploymentInput) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BlueprintDeploymentInput) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *BlueprintDeploymentInput) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSubstrateLocalReference

`func (o *BlueprintDeploymentInput) GetSubstrateLocalReference() AppSubstrateReference`

GetSubstrateLocalReference returns the SubstrateLocalReference field if non-nil, zero value otherwise.

### GetSubstrateLocalReferenceOk

`func (o *BlueprintDeploymentInput) GetSubstrateLocalReferenceOk() (*AppSubstrateReference, bool)`

GetSubstrateLocalReferenceOk returns a tuple with the SubstrateLocalReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubstrateLocalReference

`func (o *BlueprintDeploymentInput) SetSubstrateLocalReference(v AppSubstrateReference)`

SetSubstrateLocalReference sets SubstrateLocalReference field to given value.


### GetVariableList

`func (o *BlueprintDeploymentInput) GetVariableList() []AppVariableInput`

GetVariableList returns the VariableList field if non-nil, zero value otherwise.

### GetVariableListOk

`func (o *BlueprintDeploymentInput) GetVariableListOk() (*[]AppVariableInput, bool)`

GetVariableListOk returns a tuple with the VariableList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableList

`func (o *BlueprintDeploymentInput) SetVariableList(v []AppVariableInput)`

SetVariableList sets VariableList field to given value.

### HasVariableList

`func (o *BlueprintDeploymentInput) HasVariableList() bool`

HasVariableList returns a boolean if a field has been set.

### GetMinReplicas

`func (o *BlueprintDeploymentInput) GetMinReplicas() string`

GetMinReplicas returns the MinReplicas field if non-nil, zero value otherwise.

### GetMinReplicasOk

`func (o *BlueprintDeploymentInput) GetMinReplicasOk() (*string, bool)`

GetMinReplicasOk returns a tuple with the MinReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinReplicas

`func (o *BlueprintDeploymentInput) SetMinReplicas(v string)`

SetMinReplicas sets MinReplicas field to given value.


### GetOptions

`func (o *BlueprintDeploymentInput) GetOptions() map[string]interface{}`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *BlueprintDeploymentInput) GetOptionsOk() (*map[string]interface{}, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *BlueprintDeploymentInput) SetOptions(v map[string]interface{})`

SetOptions sets Options field to given value.

### HasOptions

`func (o *BlueprintDeploymentInput) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetNumFaultTolerance

`func (o *BlueprintDeploymentInput) GetNumFaultTolerance() int64`

GetNumFaultTolerance returns the NumFaultTolerance field if non-nil, zero value otherwise.

### GetNumFaultToleranceOk

`func (o *BlueprintDeploymentInput) GetNumFaultToleranceOk() (*int64, bool)`

GetNumFaultToleranceOk returns a tuple with the NumFaultTolerance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumFaultTolerance

`func (o *BlueprintDeploymentInput) SetNumFaultTolerance(v int64)`

SetNumFaultTolerance sets NumFaultTolerance field to given value.

### HasNumFaultTolerance

`func (o *BlueprintDeploymentInput) HasNumFaultTolerance() bool`

HasNumFaultTolerance returns a boolean if a field has been set.

### GetBrownfieldInstanceList

`func (o *BlueprintDeploymentInput) GetBrownfieldInstanceList() []BrownfieldInstanceInput`

GetBrownfieldInstanceList returns the BrownfieldInstanceList field if non-nil, zero value otherwise.

### GetBrownfieldInstanceListOk

`func (o *BlueprintDeploymentInput) GetBrownfieldInstanceListOk() (*[]BrownfieldInstanceInput, bool)`

GetBrownfieldInstanceListOk returns a tuple with the BrownfieldInstanceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrownfieldInstanceList

`func (o *BlueprintDeploymentInput) SetBrownfieldInstanceList(v []BrownfieldInstanceInput)`

SetBrownfieldInstanceList sets BrownfieldInstanceList field to given value.

### HasBrownfieldInstanceList

`func (o *BlueprintDeploymentInput) HasBrownfieldInstanceList() bool`

HasBrownfieldInstanceList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



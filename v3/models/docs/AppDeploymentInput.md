# AppDeploymentInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PercentFaultTolerance** | Pointer to **int64** |  | [optional] 
**PublishedServiceList** | Pointer to [**[]AppPublishedServiceInput**](AppPublishedServiceInput.md) | List of references for published services | [optional] 
**Substrate** | [**AppSubstrateInput**](AppSubstrateInput.md) |  | 
**Description** | Pointer to **string** |  | [optional] 
**ActionList** | Pointer to [**[]AppActionInput**](AppActionInput.md) | List of references to action  | [optional] 
**ServiceList** | Pointer to [**[]AppServiceInput**](AppServiceInput.md) | List of references for services | [optional] 
**Name** | **string** |  | 
**UUID** | **string** |  | 
**PackageList** | Pointer to [**[]AppPackageInput**](AppPackageInput.md) | List of references for the packages | [optional] 
**State** | Pointer to **string** |  | [optional] 
**DependsOnList** | Pointer to [**[]EntityReference**](EntityReference.md) |  | [optional] 
**MaxReplicas** | **string** | Maximum replicas for the deployment. | [default to "1"]
**Type** | Pointer to **string** |  | [optional] [default to "GREENFIELD"]
**ConfigReference** | Pointer to [**AppBlueprintDeploymentReference**](AppBlueprintDeploymentReference.md) |  | [optional] 
**FaultDomainScope** | Pointer to **string** |  | [optional] 
**VariableList** | Pointer to [**[]AppVariableInput**](AppVariableInput.md) |  | [optional] 
**MinReplicas** | **string** | Minimum replicas for the deployment. | [default to "1"]
**Options** | Pointer to **map[string]interface{}** | Additional deployment options | [optional] 
**NumFaultTolerance** | Pointer to **int64** |  | [optional] 
**BrownfieldInstanceList** | Pointer to [**[]BrownfieldInstanceInput**](BrownfieldInstanceInput.md) | list of brownfield elements | [optional] 

## Methods

### NewAppDeploymentInput

`func NewAppDeploymentInput(substrate AppSubstrateInput, name string, uUID string, maxReplicas string, minReplicas string, ) *AppDeploymentInput`

NewAppDeploymentInput instantiates a new AppDeploymentInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppDeploymentInputWithDefaults

`func NewAppDeploymentInputWithDefaults() *AppDeploymentInput`

NewAppDeploymentInputWithDefaults instantiates a new AppDeploymentInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPercentFaultTolerance

`func (o *AppDeploymentInput) GetPercentFaultTolerance() int64`

GetPercentFaultTolerance returns the PercentFaultTolerance field if non-nil, zero value otherwise.

### GetPercentFaultToleranceOk

`func (o *AppDeploymentInput) GetPercentFaultToleranceOk() (*int64, bool)`

GetPercentFaultToleranceOk returns a tuple with the PercentFaultTolerance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentFaultTolerance

`func (o *AppDeploymentInput) SetPercentFaultTolerance(v int64)`

SetPercentFaultTolerance sets PercentFaultTolerance field to given value.

### HasPercentFaultTolerance

`func (o *AppDeploymentInput) HasPercentFaultTolerance() bool`

HasPercentFaultTolerance returns a boolean if a field has been set.

### GetPublishedServiceList

`func (o *AppDeploymentInput) GetPublishedServiceList() []AppPublishedServiceInput`

GetPublishedServiceList returns the PublishedServiceList field if non-nil, zero value otherwise.

### GetPublishedServiceListOk

`func (o *AppDeploymentInput) GetPublishedServiceListOk() (*[]AppPublishedServiceInput, bool)`

GetPublishedServiceListOk returns a tuple with the PublishedServiceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedServiceList

`func (o *AppDeploymentInput) SetPublishedServiceList(v []AppPublishedServiceInput)`

SetPublishedServiceList sets PublishedServiceList field to given value.

### HasPublishedServiceList

`func (o *AppDeploymentInput) HasPublishedServiceList() bool`

HasPublishedServiceList returns a boolean if a field has been set.

### GetSubstrate

`func (o *AppDeploymentInput) GetSubstrate() AppSubstrateInput`

GetSubstrate returns the Substrate field if non-nil, zero value otherwise.

### GetSubstrateOk

`func (o *AppDeploymentInput) GetSubstrateOk() (*AppSubstrateInput, bool)`

GetSubstrateOk returns a tuple with the Substrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubstrate

`func (o *AppDeploymentInput) SetSubstrate(v AppSubstrateInput)`

SetSubstrate sets Substrate field to given value.


### GetDescription

`func (o *AppDeploymentInput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AppDeploymentInput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AppDeploymentInput) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AppDeploymentInput) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetActionList

`func (o *AppDeploymentInput) GetActionList() []AppActionInput`

GetActionList returns the ActionList field if non-nil, zero value otherwise.

### GetActionListOk

`func (o *AppDeploymentInput) GetActionListOk() (*[]AppActionInput, bool)`

GetActionListOk returns a tuple with the ActionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionList

`func (o *AppDeploymentInput) SetActionList(v []AppActionInput)`

SetActionList sets ActionList field to given value.

### HasActionList

`func (o *AppDeploymentInput) HasActionList() bool`

HasActionList returns a boolean if a field has been set.

### GetServiceList

`func (o *AppDeploymentInput) GetServiceList() []AppServiceInput`

GetServiceList returns the ServiceList field if non-nil, zero value otherwise.

### GetServiceListOk

`func (o *AppDeploymentInput) GetServiceListOk() (*[]AppServiceInput, bool)`

GetServiceListOk returns a tuple with the ServiceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceList

`func (o *AppDeploymentInput) SetServiceList(v []AppServiceInput)`

SetServiceList sets ServiceList field to given value.

### HasServiceList

`func (o *AppDeploymentInput) HasServiceList() bool`

HasServiceList returns a boolean if a field has been set.

### GetName

`func (o *AppDeploymentInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppDeploymentInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppDeploymentInput) SetName(v string)`

SetName sets Name field to given value.


### GetUUID

`func (o *AppDeploymentInput) GetUUID() string`

GetUUID returns the UUID field if non-nil, zero value otherwise.

### GetUUIDOk

`func (o *AppDeploymentInput) GetUUIDOk() (*string, bool)`

GetUUIDOk returns a tuple with the UUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUUID

`func (o *AppDeploymentInput) SetUUID(v string)`

SetUUID sets UUID field to given value.


### GetPackageList

`func (o *AppDeploymentInput) GetPackageList() []AppPackageInput`

GetPackageList returns the PackageList field if non-nil, zero value otherwise.

### GetPackageListOk

`func (o *AppDeploymentInput) GetPackageListOk() (*[]AppPackageInput, bool)`

GetPackageListOk returns a tuple with the PackageList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageList

`func (o *AppDeploymentInput) SetPackageList(v []AppPackageInput)`

SetPackageList sets PackageList field to given value.

### HasPackageList

`func (o *AppDeploymentInput) HasPackageList() bool`

HasPackageList returns a boolean if a field has been set.

### GetState

`func (o *AppDeploymentInput) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AppDeploymentInput) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AppDeploymentInput) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *AppDeploymentInput) HasState() bool`

HasState returns a boolean if a field has been set.

### GetDependsOnList

`func (o *AppDeploymentInput) GetDependsOnList() []EntityReference`

GetDependsOnList returns the DependsOnList field if non-nil, zero value otherwise.

### GetDependsOnListOk

`func (o *AppDeploymentInput) GetDependsOnListOk() (*[]EntityReference, bool)`

GetDependsOnListOk returns a tuple with the DependsOnList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependsOnList

`func (o *AppDeploymentInput) SetDependsOnList(v []EntityReference)`

SetDependsOnList sets DependsOnList field to given value.

### HasDependsOnList

`func (o *AppDeploymentInput) HasDependsOnList() bool`

HasDependsOnList returns a boolean if a field has been set.

### GetMaxReplicas

`func (o *AppDeploymentInput) GetMaxReplicas() string`

GetMaxReplicas returns the MaxReplicas field if non-nil, zero value otherwise.

### GetMaxReplicasOk

`func (o *AppDeploymentInput) GetMaxReplicasOk() (*string, bool)`

GetMaxReplicasOk returns a tuple with the MaxReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxReplicas

`func (o *AppDeploymentInput) SetMaxReplicas(v string)`

SetMaxReplicas sets MaxReplicas field to given value.


### GetType

`func (o *AppDeploymentInput) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AppDeploymentInput) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AppDeploymentInput) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AppDeploymentInput) HasType() bool`

HasType returns a boolean if a field has been set.

### GetConfigReference

`func (o *AppDeploymentInput) GetConfigReference() AppBlueprintDeploymentReference`

GetConfigReference returns the ConfigReference field if non-nil, zero value otherwise.

### GetConfigReferenceOk

`func (o *AppDeploymentInput) GetConfigReferenceOk() (*AppBlueprintDeploymentReference, bool)`

GetConfigReferenceOk returns a tuple with the ConfigReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigReference

`func (o *AppDeploymentInput) SetConfigReference(v AppBlueprintDeploymentReference)`

SetConfigReference sets ConfigReference field to given value.

### HasConfigReference

`func (o *AppDeploymentInput) HasConfigReference() bool`

HasConfigReference returns a boolean if a field has been set.

### GetFaultDomainScope

`func (o *AppDeploymentInput) GetFaultDomainScope() string`

GetFaultDomainScope returns the FaultDomainScope field if non-nil, zero value otherwise.

### GetFaultDomainScopeOk

`func (o *AppDeploymentInput) GetFaultDomainScopeOk() (*string, bool)`

GetFaultDomainScopeOk returns a tuple with the FaultDomainScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaultDomainScope

`func (o *AppDeploymentInput) SetFaultDomainScope(v string)`

SetFaultDomainScope sets FaultDomainScope field to given value.

### HasFaultDomainScope

`func (o *AppDeploymentInput) HasFaultDomainScope() bool`

HasFaultDomainScope returns a boolean if a field has been set.

### GetVariableList

`func (o *AppDeploymentInput) GetVariableList() []AppVariableInput`

GetVariableList returns the VariableList field if non-nil, zero value otherwise.

### GetVariableListOk

`func (o *AppDeploymentInput) GetVariableListOk() (*[]AppVariableInput, bool)`

GetVariableListOk returns a tuple with the VariableList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableList

`func (o *AppDeploymentInput) SetVariableList(v []AppVariableInput)`

SetVariableList sets VariableList field to given value.

### HasVariableList

`func (o *AppDeploymentInput) HasVariableList() bool`

HasVariableList returns a boolean if a field has been set.

### GetMinReplicas

`func (o *AppDeploymentInput) GetMinReplicas() string`

GetMinReplicas returns the MinReplicas field if non-nil, zero value otherwise.

### GetMinReplicasOk

`func (o *AppDeploymentInput) GetMinReplicasOk() (*string, bool)`

GetMinReplicasOk returns a tuple with the MinReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinReplicas

`func (o *AppDeploymentInput) SetMinReplicas(v string)`

SetMinReplicas sets MinReplicas field to given value.


### GetOptions

`func (o *AppDeploymentInput) GetOptions() map[string]interface{}`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *AppDeploymentInput) GetOptionsOk() (*map[string]interface{}, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *AppDeploymentInput) SetOptions(v map[string]interface{})`

SetOptions sets Options field to given value.

### HasOptions

`func (o *AppDeploymentInput) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetNumFaultTolerance

`func (o *AppDeploymentInput) GetNumFaultTolerance() int64`

GetNumFaultTolerance returns the NumFaultTolerance field if non-nil, zero value otherwise.

### GetNumFaultToleranceOk

`func (o *AppDeploymentInput) GetNumFaultToleranceOk() (*int64, bool)`

GetNumFaultToleranceOk returns a tuple with the NumFaultTolerance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumFaultTolerance

`func (o *AppDeploymentInput) SetNumFaultTolerance(v int64)`

SetNumFaultTolerance sets NumFaultTolerance field to given value.

### HasNumFaultTolerance

`func (o *AppDeploymentInput) HasNumFaultTolerance() bool`

HasNumFaultTolerance returns a boolean if a field has been set.

### GetBrownfieldInstanceList

`func (o *AppDeploymentInput) GetBrownfieldInstanceList() []BrownfieldInstanceInput`

GetBrownfieldInstanceList returns the BrownfieldInstanceList field if non-nil, zero value otherwise.

### GetBrownfieldInstanceListOk

`func (o *AppDeploymentInput) GetBrownfieldInstanceListOk() (*[]BrownfieldInstanceInput, bool)`

GetBrownfieldInstanceListOk returns a tuple with the BrownfieldInstanceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrownfieldInstanceList

`func (o *AppDeploymentInput) SetBrownfieldInstanceList(v []BrownfieldInstanceInput)`

SetBrownfieldInstanceList sets BrownfieldInstanceList field to given value.

### HasBrownfieldInstanceList

`func (o *AppDeploymentInput) HasBrownfieldInstanceList() bool`

HasBrownfieldInstanceList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



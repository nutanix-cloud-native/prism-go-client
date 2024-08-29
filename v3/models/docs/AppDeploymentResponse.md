# AppDeploymentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PercentFaultTolerance** | Pointer to **int64** |  | [optional] 
**PublishedServiceList** | Pointer to [**[]AppPublishedServiceInput**](AppPublishedServiceInput.md) | List of references for published services | [optional] 
**FaultDomainScope** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ActionList** | Pointer to [**[]AppActionResponse**](AppActionResponse.md) | List of references to action  | [optional] 
**ElementList** | [**[]AppDeploymentElement**](AppDeploymentElement.md) |  | 
**ServiceList** | Pointer to [**[]AppServiceResponse**](AppServiceResponse.md) | List of references for services | [optional] 
**SubstrateConfiguration** | [**AppSubstrateResponse**](AppSubstrateResponse.md) |  | 
**DependsOnList** | Pointer to [**[]EntityReference**](EntityReference.md) |  | [optional] 
**PackageList** | Pointer to [**[]AppPackageResponse**](AppPackageResponse.md) | List of references for the packages | [optional] 
**MinReplicas** | **string** | Minimum replicas for the deployment. | [default to "1"]
**State** | **string** |  | 
**CurrentReplicas** | **int32** |  | 
**MaxReplicas** | **string** | Maximum replicas for the deployment. | [default to "1"]
**MessageList** | [**[]MessageResource**](MessageResource.md) | Message list for deployment | 
**ConfigReference** | Pointer to [**AppBlueprintDeploymentReference**](AppBlueprintDeploymentReference.md) |  | [optional] 
**VariableList** | Pointer to [**[]AppVariableResponse**](AppVariableResponse.md) |  | [optional] 
**Name** | **string** |  | 
**Options** | Pointer to **map[string]interface{}** | Additional deployment options | [optional] 
**NumFaultTolerance** | Pointer to **int64** |  | [optional] 
**UUID** | **string** |  | 

## Methods

### NewAppDeploymentResponse

`func NewAppDeploymentResponse(elementList []AppDeploymentElement, substrateConfiguration AppSubstrateResponse, minReplicas string, state string, currentReplicas int32, maxReplicas string, messageList []MessageResource, name string, uUID string, ) *AppDeploymentResponse`

NewAppDeploymentResponse instantiates a new AppDeploymentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppDeploymentResponseWithDefaults

`func NewAppDeploymentResponseWithDefaults() *AppDeploymentResponse`

NewAppDeploymentResponseWithDefaults instantiates a new AppDeploymentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPercentFaultTolerance

`func (o *AppDeploymentResponse) GetPercentFaultTolerance() int64`

GetPercentFaultTolerance returns the PercentFaultTolerance field if non-nil, zero value otherwise.

### GetPercentFaultToleranceOk

`func (o *AppDeploymentResponse) GetPercentFaultToleranceOk() (*int64, bool)`

GetPercentFaultToleranceOk returns a tuple with the PercentFaultTolerance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentFaultTolerance

`func (o *AppDeploymentResponse) SetPercentFaultTolerance(v int64)`

SetPercentFaultTolerance sets PercentFaultTolerance field to given value.

### HasPercentFaultTolerance

`func (o *AppDeploymentResponse) HasPercentFaultTolerance() bool`

HasPercentFaultTolerance returns a boolean if a field has been set.

### GetPublishedServiceList

`func (o *AppDeploymentResponse) GetPublishedServiceList() []AppPublishedServiceInput`

GetPublishedServiceList returns the PublishedServiceList field if non-nil, zero value otherwise.

### GetPublishedServiceListOk

`func (o *AppDeploymentResponse) GetPublishedServiceListOk() (*[]AppPublishedServiceInput, bool)`

GetPublishedServiceListOk returns a tuple with the PublishedServiceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedServiceList

`func (o *AppDeploymentResponse) SetPublishedServiceList(v []AppPublishedServiceInput)`

SetPublishedServiceList sets PublishedServiceList field to given value.

### HasPublishedServiceList

`func (o *AppDeploymentResponse) HasPublishedServiceList() bool`

HasPublishedServiceList returns a boolean if a field has been set.

### GetFaultDomainScope

`func (o *AppDeploymentResponse) GetFaultDomainScope() string`

GetFaultDomainScope returns the FaultDomainScope field if non-nil, zero value otherwise.

### GetFaultDomainScopeOk

`func (o *AppDeploymentResponse) GetFaultDomainScopeOk() (*string, bool)`

GetFaultDomainScopeOk returns a tuple with the FaultDomainScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaultDomainScope

`func (o *AppDeploymentResponse) SetFaultDomainScope(v string)`

SetFaultDomainScope sets FaultDomainScope field to given value.

### HasFaultDomainScope

`func (o *AppDeploymentResponse) HasFaultDomainScope() bool`

HasFaultDomainScope returns a boolean if a field has been set.

### GetDescription

`func (o *AppDeploymentResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AppDeploymentResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AppDeploymentResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AppDeploymentResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetActionList

`func (o *AppDeploymentResponse) GetActionList() []AppActionResponse`

GetActionList returns the ActionList field if non-nil, zero value otherwise.

### GetActionListOk

`func (o *AppDeploymentResponse) GetActionListOk() (*[]AppActionResponse, bool)`

GetActionListOk returns a tuple with the ActionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionList

`func (o *AppDeploymentResponse) SetActionList(v []AppActionResponse)`

SetActionList sets ActionList field to given value.

### HasActionList

`func (o *AppDeploymentResponse) HasActionList() bool`

HasActionList returns a boolean if a field has been set.

### GetElementList

`func (o *AppDeploymentResponse) GetElementList() []AppDeploymentElement`

GetElementList returns the ElementList field if non-nil, zero value otherwise.

### GetElementListOk

`func (o *AppDeploymentResponse) GetElementListOk() (*[]AppDeploymentElement, bool)`

GetElementListOk returns a tuple with the ElementList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementList

`func (o *AppDeploymentResponse) SetElementList(v []AppDeploymentElement)`

SetElementList sets ElementList field to given value.


### GetServiceList

`func (o *AppDeploymentResponse) GetServiceList() []AppServiceResponse`

GetServiceList returns the ServiceList field if non-nil, zero value otherwise.

### GetServiceListOk

`func (o *AppDeploymentResponse) GetServiceListOk() (*[]AppServiceResponse, bool)`

GetServiceListOk returns a tuple with the ServiceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceList

`func (o *AppDeploymentResponse) SetServiceList(v []AppServiceResponse)`

SetServiceList sets ServiceList field to given value.

### HasServiceList

`func (o *AppDeploymentResponse) HasServiceList() bool`

HasServiceList returns a boolean if a field has been set.

### GetSubstrateConfiguration

`func (o *AppDeploymentResponse) GetSubstrateConfiguration() AppSubstrateResponse`

GetSubstrateConfiguration returns the SubstrateConfiguration field if non-nil, zero value otherwise.

### GetSubstrateConfigurationOk

`func (o *AppDeploymentResponse) GetSubstrateConfigurationOk() (*AppSubstrateResponse, bool)`

GetSubstrateConfigurationOk returns a tuple with the SubstrateConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubstrateConfiguration

`func (o *AppDeploymentResponse) SetSubstrateConfiguration(v AppSubstrateResponse)`

SetSubstrateConfiguration sets SubstrateConfiguration field to given value.


### GetDependsOnList

`func (o *AppDeploymentResponse) GetDependsOnList() []EntityReference`

GetDependsOnList returns the DependsOnList field if non-nil, zero value otherwise.

### GetDependsOnListOk

`func (o *AppDeploymentResponse) GetDependsOnListOk() (*[]EntityReference, bool)`

GetDependsOnListOk returns a tuple with the DependsOnList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependsOnList

`func (o *AppDeploymentResponse) SetDependsOnList(v []EntityReference)`

SetDependsOnList sets DependsOnList field to given value.

### HasDependsOnList

`func (o *AppDeploymentResponse) HasDependsOnList() bool`

HasDependsOnList returns a boolean if a field has been set.

### GetPackageList

`func (o *AppDeploymentResponse) GetPackageList() []AppPackageResponse`

GetPackageList returns the PackageList field if non-nil, zero value otherwise.

### GetPackageListOk

`func (o *AppDeploymentResponse) GetPackageListOk() (*[]AppPackageResponse, bool)`

GetPackageListOk returns a tuple with the PackageList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageList

`func (o *AppDeploymentResponse) SetPackageList(v []AppPackageResponse)`

SetPackageList sets PackageList field to given value.

### HasPackageList

`func (o *AppDeploymentResponse) HasPackageList() bool`

HasPackageList returns a boolean if a field has been set.

### GetMinReplicas

`func (o *AppDeploymentResponse) GetMinReplicas() string`

GetMinReplicas returns the MinReplicas field if non-nil, zero value otherwise.

### GetMinReplicasOk

`func (o *AppDeploymentResponse) GetMinReplicasOk() (*string, bool)`

GetMinReplicasOk returns a tuple with the MinReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinReplicas

`func (o *AppDeploymentResponse) SetMinReplicas(v string)`

SetMinReplicas sets MinReplicas field to given value.


### GetState

`func (o *AppDeploymentResponse) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AppDeploymentResponse) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AppDeploymentResponse) SetState(v string)`

SetState sets State field to given value.


### GetCurrentReplicas

`func (o *AppDeploymentResponse) GetCurrentReplicas() int32`

GetCurrentReplicas returns the CurrentReplicas field if non-nil, zero value otherwise.

### GetCurrentReplicasOk

`func (o *AppDeploymentResponse) GetCurrentReplicasOk() (*int32, bool)`

GetCurrentReplicasOk returns a tuple with the CurrentReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentReplicas

`func (o *AppDeploymentResponse) SetCurrentReplicas(v int32)`

SetCurrentReplicas sets CurrentReplicas field to given value.


### GetMaxReplicas

`func (o *AppDeploymentResponse) GetMaxReplicas() string`

GetMaxReplicas returns the MaxReplicas field if non-nil, zero value otherwise.

### GetMaxReplicasOk

`func (o *AppDeploymentResponse) GetMaxReplicasOk() (*string, bool)`

GetMaxReplicasOk returns a tuple with the MaxReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxReplicas

`func (o *AppDeploymentResponse) SetMaxReplicas(v string)`

SetMaxReplicas sets MaxReplicas field to given value.


### GetMessageList

`func (o *AppDeploymentResponse) GetMessageList() []MessageResource`

GetMessageList returns the MessageList field if non-nil, zero value otherwise.

### GetMessageListOk

`func (o *AppDeploymentResponse) GetMessageListOk() (*[]MessageResource, bool)`

GetMessageListOk returns a tuple with the MessageList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageList

`func (o *AppDeploymentResponse) SetMessageList(v []MessageResource)`

SetMessageList sets MessageList field to given value.


### GetConfigReference

`func (o *AppDeploymentResponse) GetConfigReference() AppBlueprintDeploymentReference`

GetConfigReference returns the ConfigReference field if non-nil, zero value otherwise.

### GetConfigReferenceOk

`func (o *AppDeploymentResponse) GetConfigReferenceOk() (*AppBlueprintDeploymentReference, bool)`

GetConfigReferenceOk returns a tuple with the ConfigReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigReference

`func (o *AppDeploymentResponse) SetConfigReference(v AppBlueprintDeploymentReference)`

SetConfigReference sets ConfigReference field to given value.

### HasConfigReference

`func (o *AppDeploymentResponse) HasConfigReference() bool`

HasConfigReference returns a boolean if a field has been set.

### GetVariableList

`func (o *AppDeploymentResponse) GetVariableList() []AppVariableResponse`

GetVariableList returns the VariableList field if non-nil, zero value otherwise.

### GetVariableListOk

`func (o *AppDeploymentResponse) GetVariableListOk() (*[]AppVariableResponse, bool)`

GetVariableListOk returns a tuple with the VariableList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableList

`func (o *AppDeploymentResponse) SetVariableList(v []AppVariableResponse)`

SetVariableList sets VariableList field to given value.

### HasVariableList

`func (o *AppDeploymentResponse) HasVariableList() bool`

HasVariableList returns a boolean if a field has been set.

### GetName

`func (o *AppDeploymentResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppDeploymentResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppDeploymentResponse) SetName(v string)`

SetName sets Name field to given value.


### GetOptions

`func (o *AppDeploymentResponse) GetOptions() map[string]interface{}`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *AppDeploymentResponse) GetOptionsOk() (*map[string]interface{}, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *AppDeploymentResponse) SetOptions(v map[string]interface{})`

SetOptions sets Options field to given value.

### HasOptions

`func (o *AppDeploymentResponse) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetNumFaultTolerance

`func (o *AppDeploymentResponse) GetNumFaultTolerance() int64`

GetNumFaultTolerance returns the NumFaultTolerance field if non-nil, zero value otherwise.

### GetNumFaultToleranceOk

`func (o *AppDeploymentResponse) GetNumFaultToleranceOk() (*int64, bool)`

GetNumFaultToleranceOk returns a tuple with the NumFaultTolerance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumFaultTolerance

`func (o *AppDeploymentResponse) SetNumFaultTolerance(v int64)`

SetNumFaultTolerance sets NumFaultTolerance field to given value.

### HasNumFaultTolerance

`func (o *AppDeploymentResponse) HasNumFaultTolerance() bool`

HasNumFaultTolerance returns a boolean if a field has been set.

### GetUUID

`func (o *AppDeploymentResponse) GetUUID() string`

GetUUID returns the UUID field if non-nil, zero value otherwise.

### GetUUIDOk

`func (o *AppDeploymentResponse) GetUUIDOk() (*string, bool)`

GetUUIDOk returns a tuple with the UUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUUID

`func (o *AppDeploymentResponse) SetUUID(v string)`

SetUUID sets UUID field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



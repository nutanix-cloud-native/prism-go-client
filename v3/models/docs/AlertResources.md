# AlertResources

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastUpdateTime** | Pointer to **time.Time** | The last time this alert was updated.  System sets this. | [optional] 
**SourceEntity** | Pointer to [**AlertResourcesSourceEntity**](AlertResourcesSourceEntity.md) |  | [optional] 
**RcaMetadataList** | Pointer to [**[]RcaMetadata**](RcaMetadata.md) | List of cause and resolution object from alert RCA metadata. | [optional] 
**AffectedEntityList** | Pointer to [**[]EntityInfo**](EntityInfo.md) | A list of entities causing and/or related to this alert.  | [optional] 
**Severity** | **string** | Alert severity | 
**Title** | **string** | Alert title | 
**DefaultMessage** | Pointer to **string** | Alert message. | [optional] 
**CreationTime** | Pointer to **time.Time** | The time that this alert was created. | [optional] 
**ResolutionStatus** | Pointer to [**AlertState**](AlertState.md) |  | [optional] 
**IndicatorList** | Pointer to [**[]Indicator**](Indicator.md) | The symptoms that caused this alert | [optional] 
**AcknowledgedStatus** | Pointer to [**AlertState**](AlertState.md) |  | [optional] 
**IsUserDefined** | Pointer to **bool** | The alert is raised by user defined policy or not. | [optional] 
**SeverityTrailList** | Pointer to [**[]AlertResourcesSeverityTrailListInner**](AlertResourcesSeverityTrailListInner.md) | The field has a list of information alert severity change history. If the alert is duplicated without severity change, then, that instance will not be saved here.  | [optional] 
**ClassificationList** | Pointer to **[]string** | Component classification | [optional] 
**PossibleCauseList** | Pointer to [**[]CauseAnalysis**](CauseAnalysis.md) | An ordered list of the possible causes and resolutions for the alert.  | [optional] 
**ImpactTypeList** | Pointer to **[]string** | The area this alert could impact.  | [optional] 
**Parameters** | Pointer to [**map[string]ParamValue**](ParamValue.md) | Alert notification type specific parameters. | [optional] 
**Type** | **string** | A preconfigured, or dynamically created alert type or alert type UUID. For example, A1128 for the storage pool space usage exceeded alerts, or, a real UUID for user defined alert policy.  | 
**ImpactList** | Pointer to **[]string** | The impact of the alert. | [optional] [readonly] 
**LatestOccurrenceTime** | Pointer to **time.Time** | Alert instances could be dedupped by the system.  However, the last time this similar alert was received is still be updated in this field.  | [optional] 

## Methods

### NewAlertResources

`func NewAlertResources(severity string, title string, type_ string, ) *AlertResources`

NewAlertResources instantiates a new AlertResources object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertResourcesWithDefaults

`func NewAlertResourcesWithDefaults() *AlertResources`

NewAlertResourcesWithDefaults instantiates a new AlertResources object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastUpdateTime

`func (o *AlertResources) GetLastUpdateTime() time.Time`

GetLastUpdateTime returns the LastUpdateTime field if non-nil, zero value otherwise.

### GetLastUpdateTimeOk

`func (o *AlertResources) GetLastUpdateTimeOk() (*time.Time, bool)`

GetLastUpdateTimeOk returns a tuple with the LastUpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateTime

`func (o *AlertResources) SetLastUpdateTime(v time.Time)`

SetLastUpdateTime sets LastUpdateTime field to given value.

### HasLastUpdateTime

`func (o *AlertResources) HasLastUpdateTime() bool`

HasLastUpdateTime returns a boolean if a field has been set.

### GetSourceEntity

`func (o *AlertResources) GetSourceEntity() AlertResourcesSourceEntity`

GetSourceEntity returns the SourceEntity field if non-nil, zero value otherwise.

### GetSourceEntityOk

`func (o *AlertResources) GetSourceEntityOk() (*AlertResourcesSourceEntity, bool)`

GetSourceEntityOk returns a tuple with the SourceEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceEntity

`func (o *AlertResources) SetSourceEntity(v AlertResourcesSourceEntity)`

SetSourceEntity sets SourceEntity field to given value.

### HasSourceEntity

`func (o *AlertResources) HasSourceEntity() bool`

HasSourceEntity returns a boolean if a field has been set.

### GetRcaMetadataList

`func (o *AlertResources) GetRcaMetadataList() []RcaMetadata`

GetRcaMetadataList returns the RcaMetadataList field if non-nil, zero value otherwise.

### GetRcaMetadataListOk

`func (o *AlertResources) GetRcaMetadataListOk() (*[]RcaMetadata, bool)`

GetRcaMetadataListOk returns a tuple with the RcaMetadataList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRcaMetadataList

`func (o *AlertResources) SetRcaMetadataList(v []RcaMetadata)`

SetRcaMetadataList sets RcaMetadataList field to given value.

### HasRcaMetadataList

`func (o *AlertResources) HasRcaMetadataList() bool`

HasRcaMetadataList returns a boolean if a field has been set.

### GetAffectedEntityList

`func (o *AlertResources) GetAffectedEntityList() []EntityInfo`

GetAffectedEntityList returns the AffectedEntityList field if non-nil, zero value otherwise.

### GetAffectedEntityListOk

`func (o *AlertResources) GetAffectedEntityListOk() (*[]EntityInfo, bool)`

GetAffectedEntityListOk returns a tuple with the AffectedEntityList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedEntityList

`func (o *AlertResources) SetAffectedEntityList(v []EntityInfo)`

SetAffectedEntityList sets AffectedEntityList field to given value.

### HasAffectedEntityList

`func (o *AlertResources) HasAffectedEntityList() bool`

HasAffectedEntityList returns a boolean if a field has been set.

### GetSeverity

`func (o *AlertResources) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *AlertResources) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *AlertResources) SetSeverity(v string)`

SetSeverity sets Severity field to given value.


### GetTitle

`func (o *AlertResources) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AlertResources) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AlertResources) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDefaultMessage

`func (o *AlertResources) GetDefaultMessage() string`

GetDefaultMessage returns the DefaultMessage field if non-nil, zero value otherwise.

### GetDefaultMessageOk

`func (o *AlertResources) GetDefaultMessageOk() (*string, bool)`

GetDefaultMessageOk returns a tuple with the DefaultMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultMessage

`func (o *AlertResources) SetDefaultMessage(v string)`

SetDefaultMessage sets DefaultMessage field to given value.

### HasDefaultMessage

`func (o *AlertResources) HasDefaultMessage() bool`

HasDefaultMessage returns a boolean if a field has been set.

### GetCreationTime

`func (o *AlertResources) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *AlertResources) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *AlertResources) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *AlertResources) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetResolutionStatus

`func (o *AlertResources) GetResolutionStatus() AlertState`

GetResolutionStatus returns the ResolutionStatus field if non-nil, zero value otherwise.

### GetResolutionStatusOk

`func (o *AlertResources) GetResolutionStatusOk() (*AlertState, bool)`

GetResolutionStatusOk returns a tuple with the ResolutionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolutionStatus

`func (o *AlertResources) SetResolutionStatus(v AlertState)`

SetResolutionStatus sets ResolutionStatus field to given value.

### HasResolutionStatus

`func (o *AlertResources) HasResolutionStatus() bool`

HasResolutionStatus returns a boolean if a field has been set.

### GetIndicatorList

`func (o *AlertResources) GetIndicatorList() []Indicator`

GetIndicatorList returns the IndicatorList field if non-nil, zero value otherwise.

### GetIndicatorListOk

`func (o *AlertResources) GetIndicatorListOk() (*[]Indicator, bool)`

GetIndicatorListOk returns a tuple with the IndicatorList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndicatorList

`func (o *AlertResources) SetIndicatorList(v []Indicator)`

SetIndicatorList sets IndicatorList field to given value.

### HasIndicatorList

`func (o *AlertResources) HasIndicatorList() bool`

HasIndicatorList returns a boolean if a field has been set.

### GetAcknowledgedStatus

`func (o *AlertResources) GetAcknowledgedStatus() AlertState`

GetAcknowledgedStatus returns the AcknowledgedStatus field if non-nil, zero value otherwise.

### GetAcknowledgedStatusOk

`func (o *AlertResources) GetAcknowledgedStatusOk() (*AlertState, bool)`

GetAcknowledgedStatusOk returns a tuple with the AcknowledgedStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcknowledgedStatus

`func (o *AlertResources) SetAcknowledgedStatus(v AlertState)`

SetAcknowledgedStatus sets AcknowledgedStatus field to given value.

### HasAcknowledgedStatus

`func (o *AlertResources) HasAcknowledgedStatus() bool`

HasAcknowledgedStatus returns a boolean if a field has been set.

### GetIsUserDefined

`func (o *AlertResources) GetIsUserDefined() bool`

GetIsUserDefined returns the IsUserDefined field if non-nil, zero value otherwise.

### GetIsUserDefinedOk

`func (o *AlertResources) GetIsUserDefinedOk() (*bool, bool)`

GetIsUserDefinedOk returns a tuple with the IsUserDefined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUserDefined

`func (o *AlertResources) SetIsUserDefined(v bool)`

SetIsUserDefined sets IsUserDefined field to given value.

### HasIsUserDefined

`func (o *AlertResources) HasIsUserDefined() bool`

HasIsUserDefined returns a boolean if a field has been set.

### GetSeverityTrailList

`func (o *AlertResources) GetSeverityTrailList() []AlertResourcesSeverityTrailListInner`

GetSeverityTrailList returns the SeverityTrailList field if non-nil, zero value otherwise.

### GetSeverityTrailListOk

`func (o *AlertResources) GetSeverityTrailListOk() (*[]AlertResourcesSeverityTrailListInner, bool)`

GetSeverityTrailListOk returns a tuple with the SeverityTrailList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverityTrailList

`func (o *AlertResources) SetSeverityTrailList(v []AlertResourcesSeverityTrailListInner)`

SetSeverityTrailList sets SeverityTrailList field to given value.

### HasSeverityTrailList

`func (o *AlertResources) HasSeverityTrailList() bool`

HasSeverityTrailList returns a boolean if a field has been set.

### GetClassificationList

`func (o *AlertResources) GetClassificationList() []string`

GetClassificationList returns the ClassificationList field if non-nil, zero value otherwise.

### GetClassificationListOk

`func (o *AlertResources) GetClassificationListOk() (*[]string, bool)`

GetClassificationListOk returns a tuple with the ClassificationList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassificationList

`func (o *AlertResources) SetClassificationList(v []string)`

SetClassificationList sets ClassificationList field to given value.

### HasClassificationList

`func (o *AlertResources) HasClassificationList() bool`

HasClassificationList returns a boolean if a field has been set.

### GetPossibleCauseList

`func (o *AlertResources) GetPossibleCauseList() []CauseAnalysis`

GetPossibleCauseList returns the PossibleCauseList field if non-nil, zero value otherwise.

### GetPossibleCauseListOk

`func (o *AlertResources) GetPossibleCauseListOk() (*[]CauseAnalysis, bool)`

GetPossibleCauseListOk returns a tuple with the PossibleCauseList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPossibleCauseList

`func (o *AlertResources) SetPossibleCauseList(v []CauseAnalysis)`

SetPossibleCauseList sets PossibleCauseList field to given value.

### HasPossibleCauseList

`func (o *AlertResources) HasPossibleCauseList() bool`

HasPossibleCauseList returns a boolean if a field has been set.

### GetImpactTypeList

`func (o *AlertResources) GetImpactTypeList() []string`

GetImpactTypeList returns the ImpactTypeList field if non-nil, zero value otherwise.

### GetImpactTypeListOk

`func (o *AlertResources) GetImpactTypeListOk() (*[]string, bool)`

GetImpactTypeListOk returns a tuple with the ImpactTypeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpactTypeList

`func (o *AlertResources) SetImpactTypeList(v []string)`

SetImpactTypeList sets ImpactTypeList field to given value.

### HasImpactTypeList

`func (o *AlertResources) HasImpactTypeList() bool`

HasImpactTypeList returns a boolean if a field has been set.

### GetParameters

`func (o *AlertResources) GetParameters() map[string]ParamValue`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *AlertResources) GetParametersOk() (*map[string]ParamValue, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *AlertResources) SetParameters(v map[string]ParamValue)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *AlertResources) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetType

`func (o *AlertResources) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AlertResources) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AlertResources) SetType(v string)`

SetType sets Type field to given value.


### GetImpactList

`func (o *AlertResources) GetImpactList() []string`

GetImpactList returns the ImpactList field if non-nil, zero value otherwise.

### GetImpactListOk

`func (o *AlertResources) GetImpactListOk() (*[]string, bool)`

GetImpactListOk returns a tuple with the ImpactList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpactList

`func (o *AlertResources) SetImpactList(v []string)`

SetImpactList sets ImpactList field to given value.

### HasImpactList

`func (o *AlertResources) HasImpactList() bool`

HasImpactList returns a boolean if a field has been set.

### GetLatestOccurrenceTime

`func (o *AlertResources) GetLatestOccurrenceTime() time.Time`

GetLatestOccurrenceTime returns the LatestOccurrenceTime field if non-nil, zero value otherwise.

### GetLatestOccurrenceTimeOk

`func (o *AlertResources) GetLatestOccurrenceTimeOk() (*time.Time, bool)`

GetLatestOccurrenceTimeOk returns a tuple with the LatestOccurrenceTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestOccurrenceTime

`func (o *AlertResources) SetLatestOccurrenceTime(v time.Time)`

SetLatestOccurrenceTime sets LatestOccurrenceTime field to given value.

### HasLatestOccurrenceTime

`func (o *AlertResources) HasLatestOccurrenceTime() bool`

HasLatestOccurrenceTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



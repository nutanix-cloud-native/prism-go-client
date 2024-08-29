# RecoverEntities

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityInfoList** | [**[]RecoverEntitiesEntityInfoListInner**](RecoverEntitiesEntityInfoListInner.md) | Information about entities to be recovered as part of this stage. For VM, entity information will include set of scripts to be executed after recovery of VM. Only one of categories or any_entity_reference has to be provided.  | 

## Methods

### NewRecoverEntities

`func NewRecoverEntities(entityInfoList []RecoverEntitiesEntityInfoListInner, ) *RecoverEntities`

NewRecoverEntities instantiates a new RecoverEntities object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverEntitiesWithDefaults

`func NewRecoverEntitiesWithDefaults() *RecoverEntities`

NewRecoverEntitiesWithDefaults instantiates a new RecoverEntities object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityInfoList

`func (o *RecoverEntities) GetEntityInfoList() []RecoverEntitiesEntityInfoListInner`

GetEntityInfoList returns the EntityInfoList field if non-nil, zero value otherwise.

### GetEntityInfoListOk

`func (o *RecoverEntities) GetEntityInfoListOk() (*[]RecoverEntitiesEntityInfoListInner, bool)`

GetEntityInfoListOk returns a tuple with the EntityInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityInfoList

`func (o *RecoverEntities) SetEntityInfoList(v []RecoverEntitiesEntityInfoListInner)`

SetEntityInfoList sets EntityInfoList field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



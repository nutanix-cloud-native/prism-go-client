# ClusterCapabilityPrecheck

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CheckName** | Pointer to **string** | Name of the check | [optional] 
**HasPassed** | Pointer to **bool** | Whether this precheck has passed | [optional] 
**Reason** | Pointer to **string** | Reason of failed validation. Will only be populated when validation fails.  | [optional] 

## Methods

### NewClusterCapabilityPrecheck

`func NewClusterCapabilityPrecheck() *ClusterCapabilityPrecheck`

NewClusterCapabilityPrecheck instantiates a new ClusterCapabilityPrecheck object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterCapabilityPrecheckWithDefaults

`func NewClusterCapabilityPrecheckWithDefaults() *ClusterCapabilityPrecheck`

NewClusterCapabilityPrecheckWithDefaults instantiates a new ClusterCapabilityPrecheck object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCheckName

`func (o *ClusterCapabilityPrecheck) GetCheckName() string`

GetCheckName returns the CheckName field if non-nil, zero value otherwise.

### GetCheckNameOk

`func (o *ClusterCapabilityPrecheck) GetCheckNameOk() (*string, bool)`

GetCheckNameOk returns a tuple with the CheckName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckName

`func (o *ClusterCapabilityPrecheck) SetCheckName(v string)`

SetCheckName sets CheckName field to given value.

### HasCheckName

`func (o *ClusterCapabilityPrecheck) HasCheckName() bool`

HasCheckName returns a boolean if a field has been set.

### GetHasPassed

`func (o *ClusterCapabilityPrecheck) GetHasPassed() bool`

GetHasPassed returns the HasPassed field if non-nil, zero value otherwise.

### GetHasPassedOk

`func (o *ClusterCapabilityPrecheck) GetHasPassedOk() (*bool, bool)`

GetHasPassedOk returns a tuple with the HasPassed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPassed

`func (o *ClusterCapabilityPrecheck) SetHasPassed(v bool)`

SetHasPassed sets HasPassed field to given value.

### HasHasPassed

`func (o *ClusterCapabilityPrecheck) HasHasPassed() bool`

HasHasPassed returns a boolean if a field has been set.

### GetReason

`func (o *ClusterCapabilityPrecheck) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ClusterCapabilityPrecheck) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ClusterCapabilityPrecheck) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *ClusterCapabilityPrecheck) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



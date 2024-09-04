# StreamingPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Policy name | 
**Resources** | [**StreamingPolicyRequestDetails**](StreamingPolicyRequestDetails.md) |  | 
**Description** | **string** | Policy description | 

## Methods

### NewStreamingPolicy

`func NewStreamingPolicy(name string, resources StreamingPolicyRequestDetails, description string, ) *StreamingPolicy`

NewStreamingPolicy instantiates a new StreamingPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamingPolicyWithDefaults

`func NewStreamingPolicyWithDefaults() *StreamingPolicy`

NewStreamingPolicyWithDefaults instantiates a new StreamingPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *StreamingPolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StreamingPolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StreamingPolicy) SetName(v string)`

SetName sets Name field to given value.


### GetResources

`func (o *StreamingPolicy) GetResources() StreamingPolicyRequestDetails`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *StreamingPolicy) GetResourcesOk() (*StreamingPolicyRequestDetails, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *StreamingPolicy) SetResources(v StreamingPolicyRequestDetails)`

SetResources sets Resources field to given value.


### GetDescription

`func (o *StreamingPolicy) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *StreamingPolicy) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *StreamingPolicy) SetDescription(v string)`

SetDescription sets Description field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



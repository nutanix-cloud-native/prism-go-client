# Datacenter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The datacenter name | 
**Resources** | [**DatacenterResources**](DatacenterResources.md) |  | 

## Methods

### NewDatacenter

`func NewDatacenter(name string, resources DatacenterResources, ) *Datacenter`

NewDatacenter instantiates a new Datacenter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatacenterWithDefaults

`func NewDatacenterWithDefaults() *Datacenter`

NewDatacenterWithDefaults instantiates a new Datacenter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Datacenter) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Datacenter) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Datacenter) SetName(v string)`

SetName sets Name field to given value.


### GetResources

`func (o *Datacenter) GetResources() DatacenterResources`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *Datacenter) GetResourcesOk() (*DatacenterResources, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *Datacenter) SetResources(v DatacenterResources)`

SetResources sets Resources field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



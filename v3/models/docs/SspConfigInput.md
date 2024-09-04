# SspConfigInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DirectoryServiceServiceAccount** | Pointer to [**Credentials**](Credentials.md) |  | [optional] 
**ShouldSkipMigration** | Pointer to **bool** | Migration can be skipped or not. If this flag is set to True, migration will be skipped and the Prism Element will be directly marked as migration completed. Default is False.  | [optional] 

## Methods

### NewSspConfigInput

`func NewSspConfigInput() *SspConfigInput`

NewSspConfigInput instantiates a new SspConfigInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSspConfigInputWithDefaults

`func NewSspConfigInputWithDefaults() *SspConfigInput`

NewSspConfigInputWithDefaults instantiates a new SspConfigInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDirectoryServiceServiceAccount

`func (o *SspConfigInput) GetDirectoryServiceServiceAccount() Credentials`

GetDirectoryServiceServiceAccount returns the DirectoryServiceServiceAccount field if non-nil, zero value otherwise.

### GetDirectoryServiceServiceAccountOk

`func (o *SspConfigInput) GetDirectoryServiceServiceAccountOk() (*Credentials, bool)`

GetDirectoryServiceServiceAccountOk returns a tuple with the DirectoryServiceServiceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryServiceServiceAccount

`func (o *SspConfigInput) SetDirectoryServiceServiceAccount(v Credentials)`

SetDirectoryServiceServiceAccount sets DirectoryServiceServiceAccount field to given value.

### HasDirectoryServiceServiceAccount

`func (o *SspConfigInput) HasDirectoryServiceServiceAccount() bool`

HasDirectoryServiceServiceAccount returns a boolean if a field has been set.

### GetShouldSkipMigration

`func (o *SspConfigInput) GetShouldSkipMigration() bool`

GetShouldSkipMigration returns the ShouldSkipMigration field if non-nil, zero value otherwise.

### GetShouldSkipMigrationOk

`func (o *SspConfigInput) GetShouldSkipMigrationOk() (*bool, bool)`

GetShouldSkipMigrationOk returns a tuple with the ShouldSkipMigration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldSkipMigration

`func (o *SspConfigInput) SetShouldSkipMigration(v bool)`

SetShouldSkipMigration sets ShouldSkipMigration field to given value.

### HasShouldSkipMigration

`func (o *SspConfigInput) HasShouldSkipMigration() bool`

HasShouldSkipMigration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



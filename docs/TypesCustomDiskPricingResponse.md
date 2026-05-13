# TypesCustomDiskPricingResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entries** | Pointer to [**[]SchemaDiskPricingsRow**](SchemaDiskPricingsRow.md) |  | [optional] 
**Status** | Pointer to [**GithubComBaselinehqPricingapiPkgTypesStatus**](GithubComBaselinehqPricingapiPkgTypesStatus.md) |  | [optional] 

## Methods

### NewTypesCustomDiskPricingResponse

`func NewTypesCustomDiskPricingResponse() *TypesCustomDiskPricingResponse`

NewTypesCustomDiskPricingResponse instantiates a new TypesCustomDiskPricingResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTypesCustomDiskPricingResponseWithDefaults

`func NewTypesCustomDiskPricingResponseWithDefaults() *TypesCustomDiskPricingResponse`

NewTypesCustomDiskPricingResponseWithDefaults instantiates a new TypesCustomDiskPricingResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntries

`func (o *TypesCustomDiskPricingResponse) GetEntries() []SchemaDiskPricingsRow`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *TypesCustomDiskPricingResponse) GetEntriesOk() (*[]SchemaDiskPricingsRow, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *TypesCustomDiskPricingResponse) SetEntries(v []SchemaDiskPricingsRow)`

SetEntries sets Entries field to given value.

### HasEntries

`func (o *TypesCustomDiskPricingResponse) HasEntries() bool`

HasEntries returns a boolean if a field has been set.

### GetStatus

`func (o *TypesCustomDiskPricingResponse) GetStatus() GithubComBaselinehqPricingapiPkgTypesStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TypesCustomDiskPricingResponse) GetStatusOk() (*GithubComBaselinehqPricingapiPkgTypesStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TypesCustomDiskPricingResponse) SetStatus(v GithubComBaselinehqPricingapiPkgTypesStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TypesCustomDiskPricingResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



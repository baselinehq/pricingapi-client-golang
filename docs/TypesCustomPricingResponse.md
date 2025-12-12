# TypesCustomPricingResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entries** | Pointer to [**[]SchemaComputePricingsRow**](SchemaComputePricingsRow.md) |  | [optional] 
**Status** | Pointer to [**GithubComBaselinehqPricingapiPkgTypesStatus**](GithubComBaselinehqPricingapiPkgTypesStatus.md) |  | [optional] 

## Methods

### NewTypesCustomPricingResponse

`func NewTypesCustomPricingResponse() *TypesCustomPricingResponse`

NewTypesCustomPricingResponse instantiates a new TypesCustomPricingResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTypesCustomPricingResponseWithDefaults

`func NewTypesCustomPricingResponseWithDefaults() *TypesCustomPricingResponse`

NewTypesCustomPricingResponseWithDefaults instantiates a new TypesCustomPricingResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntries

`func (o *TypesCustomPricingResponse) GetEntries() []SchemaComputePricingsRow`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *TypesCustomPricingResponse) GetEntriesOk() (*[]SchemaComputePricingsRow, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *TypesCustomPricingResponse) SetEntries(v []SchemaComputePricingsRow)`

SetEntries sets Entries field to given value.

### HasEntries

`func (o *TypesCustomPricingResponse) HasEntries() bool`

HasEntries returns a boolean if a field has been set.

### GetStatus

`func (o *TypesCustomPricingResponse) GetStatus() GithubComBaselinehqPricingapiPkgTypesStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TypesCustomPricingResponse) GetStatusOk() (*GithubComBaselinehqPricingapiPkgTypesStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TypesCustomPricingResponse) SetStatus(v GithubComBaselinehqPricingapiPkgTypesStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TypesCustomPricingResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



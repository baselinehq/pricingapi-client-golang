# FilterSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filters** | Pointer to **[]string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 

## Methods

### NewFilterSet

`func NewFilterSet() *FilterSet`

NewFilterSet instantiates a new FilterSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilterSetWithDefaults

`func NewFilterSetWithDefaults() *FilterSet`

NewFilterSetWithDefaults instantiates a new FilterSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilters

`func (o *FilterSet) GetFilters() []string`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *FilterSet) GetFiltersOk() (*[]string, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *FilterSet) SetFilters(v []string)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *FilterSet) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetKind

`func (o *FilterSet) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *FilterSet) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *FilterSet) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *FilterSet) HasKind() bool`

HasKind returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



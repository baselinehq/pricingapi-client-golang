# ProviderConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Predicates** | Pointer to [**TypesPredicates**](TypesPredicates.md) |  | [optional] 
**Provider** | Pointer to [**GithubComBaselinehqGolangSharedTypesProvider**](GithubComBaselinehqGolangSharedTypesProvider.md) |  | [optional] 

## Methods

### NewProviderConfig

`func NewProviderConfig() *ProviderConfig`

NewProviderConfig instantiates a new ProviderConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderConfigWithDefaults

`func NewProviderConfigWithDefaults() *ProviderConfig`

NewProviderConfigWithDefaults instantiates a new ProviderConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPredicates

`func (o *ProviderConfig) GetPredicates() TypesPredicates`

GetPredicates returns the Predicates field if non-nil, zero value otherwise.

### GetPredicatesOk

`func (o *ProviderConfig) GetPredicatesOk() (*TypesPredicates, bool)`

GetPredicatesOk returns a tuple with the Predicates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredicates

`func (o *ProviderConfig) SetPredicates(v TypesPredicates)`

SetPredicates sets Predicates field to given value.

### HasPredicates

`func (o *ProviderConfig) HasPredicates() bool`

HasPredicates returns a boolean if a field has been set.

### GetProvider

`func (o *ProviderConfig) GetProvider() GithubComBaselinehqGolangSharedTypesProvider`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *ProviderConfig) GetProviderOk() (*GithubComBaselinehqGolangSharedTypesProvider, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *ProviderConfig) SetProvider(v GithubComBaselinehqGolangSharedTypesProvider)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *ProviderConfig) HasProvider() bool`

HasProvider returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



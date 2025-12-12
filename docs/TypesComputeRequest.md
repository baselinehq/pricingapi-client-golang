# TypesComputeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IncludeMetadata** | Pointer to **bool** |  | [optional] 
**Instance** | Pointer to [**GithubComBaselinehqGolangSharedTypesInstance**](GithubComBaselinehqGolangSharedTypesInstance.md) |  | [optional] 
**Predicates** | Pointer to [**TypesPredicates**](TypesPredicates.md) |  | [optional] 
**Usage** | Pointer to [**GithubComBaselinehqGolangSharedTypesVM**](GithubComBaselinehqGolangSharedTypesVM.md) |  | [optional] 

## Methods

### NewTypesComputeRequest

`func NewTypesComputeRequest() *TypesComputeRequest`

NewTypesComputeRequest instantiates a new TypesComputeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTypesComputeRequestWithDefaults

`func NewTypesComputeRequestWithDefaults() *TypesComputeRequest`

NewTypesComputeRequestWithDefaults instantiates a new TypesComputeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncludeMetadata

`func (o *TypesComputeRequest) GetIncludeMetadata() bool`

GetIncludeMetadata returns the IncludeMetadata field if non-nil, zero value otherwise.

### GetIncludeMetadataOk

`func (o *TypesComputeRequest) GetIncludeMetadataOk() (*bool, bool)`

GetIncludeMetadataOk returns a tuple with the IncludeMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeMetadata

`func (o *TypesComputeRequest) SetIncludeMetadata(v bool)`

SetIncludeMetadata sets IncludeMetadata field to given value.

### HasIncludeMetadata

`func (o *TypesComputeRequest) HasIncludeMetadata() bool`

HasIncludeMetadata returns a boolean if a field has been set.

### GetInstance

`func (o *TypesComputeRequest) GetInstance() GithubComBaselinehqGolangSharedTypesInstance`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *TypesComputeRequest) GetInstanceOk() (*GithubComBaselinehqGolangSharedTypesInstance, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *TypesComputeRequest) SetInstance(v GithubComBaselinehqGolangSharedTypesInstance)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *TypesComputeRequest) HasInstance() bool`

HasInstance returns a boolean if a field has been set.

### GetPredicates

`func (o *TypesComputeRequest) GetPredicates() TypesPredicates`

GetPredicates returns the Predicates field if non-nil, zero value otherwise.

### GetPredicatesOk

`func (o *TypesComputeRequest) GetPredicatesOk() (*TypesPredicates, bool)`

GetPredicatesOk returns a tuple with the Predicates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredicates

`func (o *TypesComputeRequest) SetPredicates(v TypesPredicates)`

SetPredicates sets Predicates field to given value.

### HasPredicates

`func (o *TypesComputeRequest) HasPredicates() bool`

HasPredicates returns a boolean if a field has been set.

### GetUsage

`func (o *TypesComputeRequest) GetUsage() GithubComBaselinehqGolangSharedTypesVM`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *TypesComputeRequest) GetUsageOk() (*GithubComBaselinehqGolangSharedTypesVM, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *TypesComputeRequest) SetUsage(v GithubComBaselinehqGolangSharedTypesVM)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *TypesComputeRequest) HasUsage() bool`

HasUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



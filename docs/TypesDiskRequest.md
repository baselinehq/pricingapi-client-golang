# TypesDiskRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IncludeMetadata** | Pointer to **bool** |  | [optional] 
**Instance** | Pointer to [**TypesDisk**](TypesDisk.md) |  | [optional] 
**Predicates** | Pointer to [**TypesPredicates**](TypesPredicates.md) |  | [optional] 
**Usage** | Pointer to [**TypesDisk**](TypesDisk.md) |  | [optional] 

## Methods

### NewTypesDiskRequest

`func NewTypesDiskRequest() *TypesDiskRequest`

NewTypesDiskRequest instantiates a new TypesDiskRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTypesDiskRequestWithDefaults

`func NewTypesDiskRequestWithDefaults() *TypesDiskRequest`

NewTypesDiskRequestWithDefaults instantiates a new TypesDiskRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncludeMetadata

`func (o *TypesDiskRequest) GetIncludeMetadata() bool`

GetIncludeMetadata returns the IncludeMetadata field if non-nil, zero value otherwise.

### GetIncludeMetadataOk

`func (o *TypesDiskRequest) GetIncludeMetadataOk() (*bool, bool)`

GetIncludeMetadataOk returns a tuple with the IncludeMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeMetadata

`func (o *TypesDiskRequest) SetIncludeMetadata(v bool)`

SetIncludeMetadata sets IncludeMetadata field to given value.

### HasIncludeMetadata

`func (o *TypesDiskRequest) HasIncludeMetadata() bool`

HasIncludeMetadata returns a boolean if a field has been set.

### GetInstance

`func (o *TypesDiskRequest) GetInstance() TypesDisk`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *TypesDiskRequest) GetInstanceOk() (*TypesDisk, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *TypesDiskRequest) SetInstance(v TypesDisk)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *TypesDiskRequest) HasInstance() bool`

HasInstance returns a boolean if a field has been set.

### GetPredicates

`func (o *TypesDiskRequest) GetPredicates() TypesPredicates`

GetPredicates returns the Predicates field if non-nil, zero value otherwise.

### GetPredicatesOk

`func (o *TypesDiskRequest) GetPredicatesOk() (*TypesPredicates, bool)`

GetPredicatesOk returns a tuple with the Predicates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredicates

`func (o *TypesDiskRequest) SetPredicates(v TypesPredicates)`

SetPredicates sets Predicates field to given value.

### HasPredicates

`func (o *TypesDiskRequest) HasPredicates() bool`

HasPredicates returns a boolean if a field has been set.

### GetUsage

`func (o *TypesDiskRequest) GetUsage() TypesDisk`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *TypesDiskRequest) GetUsageOk() (*TypesDisk, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *TypesDiskRequest) SetUsage(v TypesDisk)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *TypesDiskRequest) HasUsage() bool`

HasUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# ComputePricesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entries** | [**[]ComputePrice**](ComputePrice.md) |  | 

## Methods

### NewComputePricesRequest

`func NewComputePricesRequest(entries []ComputePrice, ) *ComputePricesRequest`

NewComputePricesRequest instantiates a new ComputePricesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputePricesRequestWithDefaults

`func NewComputePricesRequestWithDefaults() *ComputePricesRequest`

NewComputePricesRequestWithDefaults instantiates a new ComputePricesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntries

`func (o *ComputePricesRequest) GetEntries() []ComputePrice`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *ComputePricesRequest) GetEntriesOk() (*[]ComputePrice, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *ComputePricesRequest) SetEntries(v []ComputePrice)`

SetEntries sets Entries field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



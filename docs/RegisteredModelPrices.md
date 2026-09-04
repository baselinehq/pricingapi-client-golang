# RegisteredModelPrices

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entries** | Pointer to [**[]ModelPrice**](ModelPrice.md) |  | [optional] 
**Status** | Pointer to [**RegistrationStatus**](RegistrationStatus.md) |  | [optional] 

## Methods

### NewRegisteredModelPrices

`func NewRegisteredModelPrices() *RegisteredModelPrices`

NewRegisteredModelPrices instantiates a new RegisteredModelPrices object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegisteredModelPricesWithDefaults

`func NewRegisteredModelPricesWithDefaults() *RegisteredModelPrices`

NewRegisteredModelPricesWithDefaults instantiates a new RegisteredModelPrices object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntries

`func (o *RegisteredModelPrices) GetEntries() []ModelPrice`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *RegisteredModelPrices) GetEntriesOk() (*[]ModelPrice, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *RegisteredModelPrices) SetEntries(v []ModelPrice)`

SetEntries sets Entries field to given value.

### HasEntries

`func (o *RegisteredModelPrices) HasEntries() bool`

HasEntries returns a boolean if a field has been set.

### GetStatus

`func (o *RegisteredModelPrices) GetStatus() RegistrationStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RegisteredModelPrices) GetStatusOk() (*RegistrationStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RegisteredModelPrices) SetStatus(v RegistrationStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RegisteredModelPrices) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



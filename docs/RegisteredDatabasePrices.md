# RegisteredDatabasePrices

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entries** | Pointer to [**[]DatabasePrice**](DatabasePrice.md) |  | [optional] 
**Status** | Pointer to [**RegistrationStatus**](RegistrationStatus.md) |  | [optional] 

## Methods

### NewRegisteredDatabasePrices

`func NewRegisteredDatabasePrices() *RegisteredDatabasePrices`

NewRegisteredDatabasePrices instantiates a new RegisteredDatabasePrices object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegisteredDatabasePricesWithDefaults

`func NewRegisteredDatabasePricesWithDefaults() *RegisteredDatabasePrices`

NewRegisteredDatabasePricesWithDefaults instantiates a new RegisteredDatabasePrices object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntries

`func (o *RegisteredDatabasePrices) GetEntries() []DatabasePrice`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *RegisteredDatabasePrices) GetEntriesOk() (*[]DatabasePrice, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *RegisteredDatabasePrices) SetEntries(v []DatabasePrice)`

SetEntries sets Entries field to given value.

### HasEntries

`func (o *RegisteredDatabasePrices) HasEntries() bool`

HasEntries returns a boolean if a field has been set.

### GetStatus

`func (o *RegisteredDatabasePrices) GetStatus() RegistrationStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RegisteredDatabasePrices) GetStatusOk() (*RegistrationStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RegisteredDatabasePrices) SetStatus(v RegistrationStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RegisteredDatabasePrices) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



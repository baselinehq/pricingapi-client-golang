# RegisteredDiskPrices

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entries** | Pointer to [**[]DiskPrice**](DiskPrice.md) |  | [optional] 
**Status** | Pointer to [**RegistrationStatus**](RegistrationStatus.md) |  | [optional] 

## Methods

### NewRegisteredDiskPrices

`func NewRegisteredDiskPrices() *RegisteredDiskPrices`

NewRegisteredDiskPrices instantiates a new RegisteredDiskPrices object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegisteredDiskPricesWithDefaults

`func NewRegisteredDiskPricesWithDefaults() *RegisteredDiskPrices`

NewRegisteredDiskPricesWithDefaults instantiates a new RegisteredDiskPrices object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntries

`func (o *RegisteredDiskPrices) GetEntries() []DiskPrice`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *RegisteredDiskPrices) GetEntriesOk() (*[]DiskPrice, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *RegisteredDiskPrices) SetEntries(v []DiskPrice)`

SetEntries sets Entries field to given value.

### HasEntries

`func (o *RegisteredDiskPrices) HasEntries() bool`

HasEntries returns a boolean if a field has been set.

### GetStatus

`func (o *RegisteredDiskPrices) GetStatus() RegistrationStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RegisteredDiskPrices) GetStatusOk() (*RegistrationStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RegisteredDiskPrices) SetStatus(v RegistrationStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RegisteredDiskPrices) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



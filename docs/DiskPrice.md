# DiskPrice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvailabilityZone** | Pointer to **string** |  | [optional] 
**CostPerGbHour** | Pointer to **float32** |  | [optional] 
**CostPerIopsHour** | Pointer to **float32** |  | [optional] 
**CostPerThroughputMbpsHour** | Pointer to **float32** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**MaxCapacityGb** | Pointer to **float32** |  | [optional] 
**MaxIops** | Pointer to **float32** |  | [optional] 
**MaxThroughputMbps** | Pointer to **float32** |  | [optional] 
**MinCapacityGb** | Pointer to **float32** |  | [optional] 
**MinIops** | Pointer to **float32** |  | [optional] 
**MinThroughputMbps** | Pointer to **float32** |  | [optional] 
**NativeCostPerGbHour** | Pointer to **float32** |  | [optional] 
**NativeCostPerIopsHour** | Pointer to **float32** |  | [optional] 
**NativeCostPerThroughputMbpsHour** | Pointer to **float32** |  | [optional] 
**PeriodBillingHours** | Pointer to **float32** |  | [optional] 
**Provider** | Pointer to **string** |  | [optional] 
**RawPricingData** | Pointer to **map[string]interface{}** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**Service** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **map[string]interface{}** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**UsageType** | Pointer to **string** |  | [optional] 

## Methods

### NewDiskPrice

`func NewDiskPrice() *DiskPrice`

NewDiskPrice instantiates a new DiskPrice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiskPriceWithDefaults

`func NewDiskPriceWithDefaults() *DiskPrice`

NewDiskPriceWithDefaults instantiates a new DiskPrice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailabilityZone

`func (o *DiskPrice) GetAvailabilityZone() string`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *DiskPrice) GetAvailabilityZoneOk() (*string, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *DiskPrice) SetAvailabilityZone(v string)`

SetAvailabilityZone sets AvailabilityZone field to given value.

### HasAvailabilityZone

`func (o *DiskPrice) HasAvailabilityZone() bool`

HasAvailabilityZone returns a boolean if a field has been set.

### GetCostPerGbHour

`func (o *DiskPrice) GetCostPerGbHour() float32`

GetCostPerGbHour returns the CostPerGbHour field if non-nil, zero value otherwise.

### GetCostPerGbHourOk

`func (o *DiskPrice) GetCostPerGbHourOk() (*float32, bool)`

GetCostPerGbHourOk returns a tuple with the CostPerGbHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostPerGbHour

`func (o *DiskPrice) SetCostPerGbHour(v float32)`

SetCostPerGbHour sets CostPerGbHour field to given value.

### HasCostPerGbHour

`func (o *DiskPrice) HasCostPerGbHour() bool`

HasCostPerGbHour returns a boolean if a field has been set.

### GetCostPerIopsHour

`func (o *DiskPrice) GetCostPerIopsHour() float32`

GetCostPerIopsHour returns the CostPerIopsHour field if non-nil, zero value otherwise.

### GetCostPerIopsHourOk

`func (o *DiskPrice) GetCostPerIopsHourOk() (*float32, bool)`

GetCostPerIopsHourOk returns a tuple with the CostPerIopsHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostPerIopsHour

`func (o *DiskPrice) SetCostPerIopsHour(v float32)`

SetCostPerIopsHour sets CostPerIopsHour field to given value.

### HasCostPerIopsHour

`func (o *DiskPrice) HasCostPerIopsHour() bool`

HasCostPerIopsHour returns a boolean if a field has been set.

### GetCostPerThroughputMbpsHour

`func (o *DiskPrice) GetCostPerThroughputMbpsHour() float32`

GetCostPerThroughputMbpsHour returns the CostPerThroughputMbpsHour field if non-nil, zero value otherwise.

### GetCostPerThroughputMbpsHourOk

`func (o *DiskPrice) GetCostPerThroughputMbpsHourOk() (*float32, bool)`

GetCostPerThroughputMbpsHourOk returns a tuple with the CostPerThroughputMbpsHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostPerThroughputMbpsHour

`func (o *DiskPrice) SetCostPerThroughputMbpsHour(v float32)`

SetCostPerThroughputMbpsHour sets CostPerThroughputMbpsHour field to given value.

### HasCostPerThroughputMbpsHour

`func (o *DiskPrice) HasCostPerThroughputMbpsHour() bool`

HasCostPerThroughputMbpsHour returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DiskPrice) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DiskPrice) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DiskPrice) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DiskPrice) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCurrency

`func (o *DiskPrice) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *DiskPrice) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *DiskPrice) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *DiskPrice) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetId

`func (o *DiskPrice) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DiskPrice) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DiskPrice) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DiskPrice) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMaxCapacityGb

`func (o *DiskPrice) GetMaxCapacityGb() float32`

GetMaxCapacityGb returns the MaxCapacityGb field if non-nil, zero value otherwise.

### GetMaxCapacityGbOk

`func (o *DiskPrice) GetMaxCapacityGbOk() (*float32, bool)`

GetMaxCapacityGbOk returns a tuple with the MaxCapacityGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCapacityGb

`func (o *DiskPrice) SetMaxCapacityGb(v float32)`

SetMaxCapacityGb sets MaxCapacityGb field to given value.

### HasMaxCapacityGb

`func (o *DiskPrice) HasMaxCapacityGb() bool`

HasMaxCapacityGb returns a boolean if a field has been set.

### GetMaxIops

`func (o *DiskPrice) GetMaxIops() float32`

GetMaxIops returns the MaxIops field if non-nil, zero value otherwise.

### GetMaxIopsOk

`func (o *DiskPrice) GetMaxIopsOk() (*float32, bool)`

GetMaxIopsOk returns a tuple with the MaxIops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxIops

`func (o *DiskPrice) SetMaxIops(v float32)`

SetMaxIops sets MaxIops field to given value.

### HasMaxIops

`func (o *DiskPrice) HasMaxIops() bool`

HasMaxIops returns a boolean if a field has been set.

### GetMaxThroughputMbps

`func (o *DiskPrice) GetMaxThroughputMbps() float32`

GetMaxThroughputMbps returns the MaxThroughputMbps field if non-nil, zero value otherwise.

### GetMaxThroughputMbpsOk

`func (o *DiskPrice) GetMaxThroughputMbpsOk() (*float32, bool)`

GetMaxThroughputMbpsOk returns a tuple with the MaxThroughputMbps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxThroughputMbps

`func (o *DiskPrice) SetMaxThroughputMbps(v float32)`

SetMaxThroughputMbps sets MaxThroughputMbps field to given value.

### HasMaxThroughputMbps

`func (o *DiskPrice) HasMaxThroughputMbps() bool`

HasMaxThroughputMbps returns a boolean if a field has been set.

### GetMinCapacityGb

`func (o *DiskPrice) GetMinCapacityGb() float32`

GetMinCapacityGb returns the MinCapacityGb field if non-nil, zero value otherwise.

### GetMinCapacityGbOk

`func (o *DiskPrice) GetMinCapacityGbOk() (*float32, bool)`

GetMinCapacityGbOk returns a tuple with the MinCapacityGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinCapacityGb

`func (o *DiskPrice) SetMinCapacityGb(v float32)`

SetMinCapacityGb sets MinCapacityGb field to given value.

### HasMinCapacityGb

`func (o *DiskPrice) HasMinCapacityGb() bool`

HasMinCapacityGb returns a boolean if a field has been set.

### GetMinIops

`func (o *DiskPrice) GetMinIops() float32`

GetMinIops returns the MinIops field if non-nil, zero value otherwise.

### GetMinIopsOk

`func (o *DiskPrice) GetMinIopsOk() (*float32, bool)`

GetMinIopsOk returns a tuple with the MinIops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinIops

`func (o *DiskPrice) SetMinIops(v float32)`

SetMinIops sets MinIops field to given value.

### HasMinIops

`func (o *DiskPrice) HasMinIops() bool`

HasMinIops returns a boolean if a field has been set.

### GetMinThroughputMbps

`func (o *DiskPrice) GetMinThroughputMbps() float32`

GetMinThroughputMbps returns the MinThroughputMbps field if non-nil, zero value otherwise.

### GetMinThroughputMbpsOk

`func (o *DiskPrice) GetMinThroughputMbpsOk() (*float32, bool)`

GetMinThroughputMbpsOk returns a tuple with the MinThroughputMbps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinThroughputMbps

`func (o *DiskPrice) SetMinThroughputMbps(v float32)`

SetMinThroughputMbps sets MinThroughputMbps field to given value.

### HasMinThroughputMbps

`func (o *DiskPrice) HasMinThroughputMbps() bool`

HasMinThroughputMbps returns a boolean if a field has been set.

### GetNativeCostPerGbHour

`func (o *DiskPrice) GetNativeCostPerGbHour() float32`

GetNativeCostPerGbHour returns the NativeCostPerGbHour field if non-nil, zero value otherwise.

### GetNativeCostPerGbHourOk

`func (o *DiskPrice) GetNativeCostPerGbHourOk() (*float32, bool)`

GetNativeCostPerGbHourOk returns a tuple with the NativeCostPerGbHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNativeCostPerGbHour

`func (o *DiskPrice) SetNativeCostPerGbHour(v float32)`

SetNativeCostPerGbHour sets NativeCostPerGbHour field to given value.

### HasNativeCostPerGbHour

`func (o *DiskPrice) HasNativeCostPerGbHour() bool`

HasNativeCostPerGbHour returns a boolean if a field has been set.

### GetNativeCostPerIopsHour

`func (o *DiskPrice) GetNativeCostPerIopsHour() float32`

GetNativeCostPerIopsHour returns the NativeCostPerIopsHour field if non-nil, zero value otherwise.

### GetNativeCostPerIopsHourOk

`func (o *DiskPrice) GetNativeCostPerIopsHourOk() (*float32, bool)`

GetNativeCostPerIopsHourOk returns a tuple with the NativeCostPerIopsHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNativeCostPerIopsHour

`func (o *DiskPrice) SetNativeCostPerIopsHour(v float32)`

SetNativeCostPerIopsHour sets NativeCostPerIopsHour field to given value.

### HasNativeCostPerIopsHour

`func (o *DiskPrice) HasNativeCostPerIopsHour() bool`

HasNativeCostPerIopsHour returns a boolean if a field has been set.

### GetNativeCostPerThroughputMbpsHour

`func (o *DiskPrice) GetNativeCostPerThroughputMbpsHour() float32`

GetNativeCostPerThroughputMbpsHour returns the NativeCostPerThroughputMbpsHour field if non-nil, zero value otherwise.

### GetNativeCostPerThroughputMbpsHourOk

`func (o *DiskPrice) GetNativeCostPerThroughputMbpsHourOk() (*float32, bool)`

GetNativeCostPerThroughputMbpsHourOk returns a tuple with the NativeCostPerThroughputMbpsHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNativeCostPerThroughputMbpsHour

`func (o *DiskPrice) SetNativeCostPerThroughputMbpsHour(v float32)`

SetNativeCostPerThroughputMbpsHour sets NativeCostPerThroughputMbpsHour field to given value.

### HasNativeCostPerThroughputMbpsHour

`func (o *DiskPrice) HasNativeCostPerThroughputMbpsHour() bool`

HasNativeCostPerThroughputMbpsHour returns a boolean if a field has been set.

### GetPeriodBillingHours

`func (o *DiskPrice) GetPeriodBillingHours() float32`

GetPeriodBillingHours returns the PeriodBillingHours field if non-nil, zero value otherwise.

### GetPeriodBillingHoursOk

`func (o *DiskPrice) GetPeriodBillingHoursOk() (*float32, bool)`

GetPeriodBillingHoursOk returns a tuple with the PeriodBillingHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodBillingHours

`func (o *DiskPrice) SetPeriodBillingHours(v float32)`

SetPeriodBillingHours sets PeriodBillingHours field to given value.

### HasPeriodBillingHours

`func (o *DiskPrice) HasPeriodBillingHours() bool`

HasPeriodBillingHours returns a boolean if a field has been set.

### GetProvider

`func (o *DiskPrice) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *DiskPrice) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *DiskPrice) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *DiskPrice) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetRawPricingData

`func (o *DiskPrice) GetRawPricingData() map[string]interface{}`

GetRawPricingData returns the RawPricingData field if non-nil, zero value otherwise.

### GetRawPricingDataOk

`func (o *DiskPrice) GetRawPricingDataOk() (*map[string]interface{}, bool)`

GetRawPricingDataOk returns a tuple with the RawPricingData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawPricingData

`func (o *DiskPrice) SetRawPricingData(v map[string]interface{})`

SetRawPricingData sets RawPricingData field to given value.

### HasRawPricingData

`func (o *DiskPrice) HasRawPricingData() bool`

HasRawPricingData returns a boolean if a field has been set.

### GetRegion

`func (o *DiskPrice) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *DiskPrice) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *DiskPrice) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *DiskPrice) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetService

`func (o *DiskPrice) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *DiskPrice) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *DiskPrice) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *DiskPrice) HasService() bool`

HasService returns a boolean if a field has been set.

### GetTags

`func (o *DiskPrice) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DiskPrice) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DiskPrice) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *DiskPrice) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetType

`func (o *DiskPrice) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DiskPrice) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DiskPrice) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DiskPrice) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DiskPrice) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DiskPrice) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DiskPrice) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DiskPrice) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUsageType

`func (o *DiskPrice) GetUsageType() string`

GetUsageType returns the UsageType field if non-nil, zero value otherwise.

### GetUsageTypeOk

`func (o *DiskPrice) GetUsageTypeOk() (*string, bool)`

GetUsageTypeOk returns a tuple with the UsageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageType

`func (o *DiskPrice) SetUsageType(v string)`

SetUsageType sets UsageType field to given value.

### HasUsageType

`func (o *DiskPrice) HasUsageType() bool`

HasUsageType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



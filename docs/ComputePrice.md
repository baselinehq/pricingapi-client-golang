# ComputePrice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Architecture** | Pointer to **string** |  | [optional] 
**AvailabilityZone** | Pointer to **string** |  | [optional] 
**CostPerHour** | Pointer to **float32** |  | [optional] 
**CpuCores** | Pointer to **float32** |  | [optional] 
**CpuCoresCostPerHour** | Pointer to **float32** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**GpuCostPerHour** | Pointer to **float32** |  | [optional] 
**GpuCount** | Pointer to **float32** |  | [optional] 
**GpuType** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**InstanceType** | Pointer to **string** |  | [optional] 
**NativeCostPerHour** | Pointer to **float32** |  | [optional] 
**OperatingSystem** | Pointer to **string** |  | [optional] 
**PeriodBillingHours** | Pointer to **float32** |  | [optional] 
**Provider** | Pointer to **string** |  | [optional] 
**RamGb** | Pointer to **float32** |  | [optional] 
**RamGbCostPerHour** | Pointer to **float32** |  | [optional] 
**RawPricingData** | Pointer to **map[string]interface{}** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**Service** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **map[string]interface{}** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**UsageType** | Pointer to **string** |  | [optional] 

## Methods

### NewComputePrice

`func NewComputePrice() *ComputePrice`

NewComputePrice instantiates a new ComputePrice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputePriceWithDefaults

`func NewComputePriceWithDefaults() *ComputePrice`

NewComputePriceWithDefaults instantiates a new ComputePrice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchitecture

`func (o *ComputePrice) GetArchitecture() string`

GetArchitecture returns the Architecture field if non-nil, zero value otherwise.

### GetArchitectureOk

`func (o *ComputePrice) GetArchitectureOk() (*string, bool)`

GetArchitectureOk returns a tuple with the Architecture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchitecture

`func (o *ComputePrice) SetArchitecture(v string)`

SetArchitecture sets Architecture field to given value.

### HasArchitecture

`func (o *ComputePrice) HasArchitecture() bool`

HasArchitecture returns a boolean if a field has been set.

### GetAvailabilityZone

`func (o *ComputePrice) GetAvailabilityZone() string`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *ComputePrice) GetAvailabilityZoneOk() (*string, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *ComputePrice) SetAvailabilityZone(v string)`

SetAvailabilityZone sets AvailabilityZone field to given value.

### HasAvailabilityZone

`func (o *ComputePrice) HasAvailabilityZone() bool`

HasAvailabilityZone returns a boolean if a field has been set.

### GetCostPerHour

`func (o *ComputePrice) GetCostPerHour() float32`

GetCostPerHour returns the CostPerHour field if non-nil, zero value otherwise.

### GetCostPerHourOk

`func (o *ComputePrice) GetCostPerHourOk() (*float32, bool)`

GetCostPerHourOk returns a tuple with the CostPerHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostPerHour

`func (o *ComputePrice) SetCostPerHour(v float32)`

SetCostPerHour sets CostPerHour field to given value.

### HasCostPerHour

`func (o *ComputePrice) HasCostPerHour() bool`

HasCostPerHour returns a boolean if a field has been set.

### GetCpuCores

`func (o *ComputePrice) GetCpuCores() float32`

GetCpuCores returns the CpuCores field if non-nil, zero value otherwise.

### GetCpuCoresOk

`func (o *ComputePrice) GetCpuCoresOk() (*float32, bool)`

GetCpuCoresOk returns a tuple with the CpuCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuCores

`func (o *ComputePrice) SetCpuCores(v float32)`

SetCpuCores sets CpuCores field to given value.

### HasCpuCores

`func (o *ComputePrice) HasCpuCores() bool`

HasCpuCores returns a boolean if a field has been set.

### GetCpuCoresCostPerHour

`func (o *ComputePrice) GetCpuCoresCostPerHour() float32`

GetCpuCoresCostPerHour returns the CpuCoresCostPerHour field if non-nil, zero value otherwise.

### GetCpuCoresCostPerHourOk

`func (o *ComputePrice) GetCpuCoresCostPerHourOk() (*float32, bool)`

GetCpuCoresCostPerHourOk returns a tuple with the CpuCoresCostPerHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuCoresCostPerHour

`func (o *ComputePrice) SetCpuCoresCostPerHour(v float32)`

SetCpuCoresCostPerHour sets CpuCoresCostPerHour field to given value.

### HasCpuCoresCostPerHour

`func (o *ComputePrice) HasCpuCoresCostPerHour() bool`

HasCpuCoresCostPerHour returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ComputePrice) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ComputePrice) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ComputePrice) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ComputePrice) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCurrency

`func (o *ComputePrice) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ComputePrice) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *ComputePrice) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *ComputePrice) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetGpuCostPerHour

`func (o *ComputePrice) GetGpuCostPerHour() float32`

GetGpuCostPerHour returns the GpuCostPerHour field if non-nil, zero value otherwise.

### GetGpuCostPerHourOk

`func (o *ComputePrice) GetGpuCostPerHourOk() (*float32, bool)`

GetGpuCostPerHourOk returns a tuple with the GpuCostPerHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuCostPerHour

`func (o *ComputePrice) SetGpuCostPerHour(v float32)`

SetGpuCostPerHour sets GpuCostPerHour field to given value.

### HasGpuCostPerHour

`func (o *ComputePrice) HasGpuCostPerHour() bool`

HasGpuCostPerHour returns a boolean if a field has been set.

### GetGpuCount

`func (o *ComputePrice) GetGpuCount() float32`

GetGpuCount returns the GpuCount field if non-nil, zero value otherwise.

### GetGpuCountOk

`func (o *ComputePrice) GetGpuCountOk() (*float32, bool)`

GetGpuCountOk returns a tuple with the GpuCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuCount

`func (o *ComputePrice) SetGpuCount(v float32)`

SetGpuCount sets GpuCount field to given value.

### HasGpuCount

`func (o *ComputePrice) HasGpuCount() bool`

HasGpuCount returns a boolean if a field has been set.

### GetGpuType

`func (o *ComputePrice) GetGpuType() string`

GetGpuType returns the GpuType field if non-nil, zero value otherwise.

### GetGpuTypeOk

`func (o *ComputePrice) GetGpuTypeOk() (*string, bool)`

GetGpuTypeOk returns a tuple with the GpuType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuType

`func (o *ComputePrice) SetGpuType(v string)`

SetGpuType sets GpuType field to given value.

### HasGpuType

`func (o *ComputePrice) HasGpuType() bool`

HasGpuType returns a boolean if a field has been set.

### GetId

`func (o *ComputePrice) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ComputePrice) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ComputePrice) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ComputePrice) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInstanceType

`func (o *ComputePrice) GetInstanceType() string`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *ComputePrice) GetInstanceTypeOk() (*string, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *ComputePrice) SetInstanceType(v string)`

SetInstanceType sets InstanceType field to given value.

### HasInstanceType

`func (o *ComputePrice) HasInstanceType() bool`

HasInstanceType returns a boolean if a field has been set.

### GetNativeCostPerHour

`func (o *ComputePrice) GetNativeCostPerHour() float32`

GetNativeCostPerHour returns the NativeCostPerHour field if non-nil, zero value otherwise.

### GetNativeCostPerHourOk

`func (o *ComputePrice) GetNativeCostPerHourOk() (*float32, bool)`

GetNativeCostPerHourOk returns a tuple with the NativeCostPerHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNativeCostPerHour

`func (o *ComputePrice) SetNativeCostPerHour(v float32)`

SetNativeCostPerHour sets NativeCostPerHour field to given value.

### HasNativeCostPerHour

`func (o *ComputePrice) HasNativeCostPerHour() bool`

HasNativeCostPerHour returns a boolean if a field has been set.

### GetOperatingSystem

`func (o *ComputePrice) GetOperatingSystem() string`

GetOperatingSystem returns the OperatingSystem field if non-nil, zero value otherwise.

### GetOperatingSystemOk

`func (o *ComputePrice) GetOperatingSystemOk() (*string, bool)`

GetOperatingSystemOk returns a tuple with the OperatingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystem

`func (o *ComputePrice) SetOperatingSystem(v string)`

SetOperatingSystem sets OperatingSystem field to given value.

### HasOperatingSystem

`func (o *ComputePrice) HasOperatingSystem() bool`

HasOperatingSystem returns a boolean if a field has been set.

### GetPeriodBillingHours

`func (o *ComputePrice) GetPeriodBillingHours() float32`

GetPeriodBillingHours returns the PeriodBillingHours field if non-nil, zero value otherwise.

### GetPeriodBillingHoursOk

`func (o *ComputePrice) GetPeriodBillingHoursOk() (*float32, bool)`

GetPeriodBillingHoursOk returns a tuple with the PeriodBillingHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodBillingHours

`func (o *ComputePrice) SetPeriodBillingHours(v float32)`

SetPeriodBillingHours sets PeriodBillingHours field to given value.

### HasPeriodBillingHours

`func (o *ComputePrice) HasPeriodBillingHours() bool`

HasPeriodBillingHours returns a boolean if a field has been set.

### GetProvider

`func (o *ComputePrice) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *ComputePrice) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *ComputePrice) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *ComputePrice) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetRamGb

`func (o *ComputePrice) GetRamGb() float32`

GetRamGb returns the RamGb field if non-nil, zero value otherwise.

### GetRamGbOk

`func (o *ComputePrice) GetRamGbOk() (*float32, bool)`

GetRamGbOk returns a tuple with the RamGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRamGb

`func (o *ComputePrice) SetRamGb(v float32)`

SetRamGb sets RamGb field to given value.

### HasRamGb

`func (o *ComputePrice) HasRamGb() bool`

HasRamGb returns a boolean if a field has been set.

### GetRamGbCostPerHour

`func (o *ComputePrice) GetRamGbCostPerHour() float32`

GetRamGbCostPerHour returns the RamGbCostPerHour field if non-nil, zero value otherwise.

### GetRamGbCostPerHourOk

`func (o *ComputePrice) GetRamGbCostPerHourOk() (*float32, bool)`

GetRamGbCostPerHourOk returns a tuple with the RamGbCostPerHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRamGbCostPerHour

`func (o *ComputePrice) SetRamGbCostPerHour(v float32)`

SetRamGbCostPerHour sets RamGbCostPerHour field to given value.

### HasRamGbCostPerHour

`func (o *ComputePrice) HasRamGbCostPerHour() bool`

HasRamGbCostPerHour returns a boolean if a field has been set.

### GetRawPricingData

`func (o *ComputePrice) GetRawPricingData() map[string]interface{}`

GetRawPricingData returns the RawPricingData field if non-nil, zero value otherwise.

### GetRawPricingDataOk

`func (o *ComputePrice) GetRawPricingDataOk() (*map[string]interface{}, bool)`

GetRawPricingDataOk returns a tuple with the RawPricingData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawPricingData

`func (o *ComputePrice) SetRawPricingData(v map[string]interface{})`

SetRawPricingData sets RawPricingData field to given value.

### HasRawPricingData

`func (o *ComputePrice) HasRawPricingData() bool`

HasRawPricingData returns a boolean if a field has been set.

### GetRegion

`func (o *ComputePrice) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *ComputePrice) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *ComputePrice) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *ComputePrice) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetService

`func (o *ComputePrice) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *ComputePrice) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *ComputePrice) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *ComputePrice) HasService() bool`

HasService returns a boolean if a field has been set.

### GetTags

`func (o *ComputePrice) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ComputePrice) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ComputePrice) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *ComputePrice) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ComputePrice) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ComputePrice) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ComputePrice) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ComputePrice) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUsageType

`func (o *ComputePrice) GetUsageType() string`

GetUsageType returns the UsageType field if non-nil, zero value otherwise.

### GetUsageTypeOk

`func (o *ComputePrice) GetUsageTypeOk() (*string, bool)`

GetUsageTypeOk returns a tuple with the UsageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageType

`func (o *ComputePrice) SetUsageType(v string)`

SetUsageType sets UsageType field to given value.

### HasUsageType

`func (o *ComputePrice) HasUsageType() bool`

HasUsageType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



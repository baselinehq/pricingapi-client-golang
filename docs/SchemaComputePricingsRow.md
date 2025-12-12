# SchemaComputePricingsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvailabilityZone** | Pointer to **string** |  | [optional] 
**CostPerHour** | Pointer to **float32** |  | [optional] 
**CpuCores** | Pointer to **float32** |  | [optional] 
**CpuCoresCostPerHour** | Pointer to **float32** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**InstanceType** | Pointer to **string** |  | [optional] 
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

### NewSchemaComputePricingsRow

`func NewSchemaComputePricingsRow() *SchemaComputePricingsRow`

NewSchemaComputePricingsRow instantiates a new SchemaComputePricingsRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchemaComputePricingsRowWithDefaults

`func NewSchemaComputePricingsRowWithDefaults() *SchemaComputePricingsRow`

NewSchemaComputePricingsRowWithDefaults instantiates a new SchemaComputePricingsRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailabilityZone

`func (o *SchemaComputePricingsRow) GetAvailabilityZone() string`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *SchemaComputePricingsRow) GetAvailabilityZoneOk() (*string, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *SchemaComputePricingsRow) SetAvailabilityZone(v string)`

SetAvailabilityZone sets AvailabilityZone field to given value.

### HasAvailabilityZone

`func (o *SchemaComputePricingsRow) HasAvailabilityZone() bool`

HasAvailabilityZone returns a boolean if a field has been set.

### GetCostPerHour

`func (o *SchemaComputePricingsRow) GetCostPerHour() float32`

GetCostPerHour returns the CostPerHour field if non-nil, zero value otherwise.

### GetCostPerHourOk

`func (o *SchemaComputePricingsRow) GetCostPerHourOk() (*float32, bool)`

GetCostPerHourOk returns a tuple with the CostPerHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostPerHour

`func (o *SchemaComputePricingsRow) SetCostPerHour(v float32)`

SetCostPerHour sets CostPerHour field to given value.

### HasCostPerHour

`func (o *SchemaComputePricingsRow) HasCostPerHour() bool`

HasCostPerHour returns a boolean if a field has been set.

### GetCpuCores

`func (o *SchemaComputePricingsRow) GetCpuCores() float32`

GetCpuCores returns the CpuCores field if non-nil, zero value otherwise.

### GetCpuCoresOk

`func (o *SchemaComputePricingsRow) GetCpuCoresOk() (*float32, bool)`

GetCpuCoresOk returns a tuple with the CpuCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuCores

`func (o *SchemaComputePricingsRow) SetCpuCores(v float32)`

SetCpuCores sets CpuCores field to given value.

### HasCpuCores

`func (o *SchemaComputePricingsRow) HasCpuCores() bool`

HasCpuCores returns a boolean if a field has been set.

### GetCpuCoresCostPerHour

`func (o *SchemaComputePricingsRow) GetCpuCoresCostPerHour() float32`

GetCpuCoresCostPerHour returns the CpuCoresCostPerHour field if non-nil, zero value otherwise.

### GetCpuCoresCostPerHourOk

`func (o *SchemaComputePricingsRow) GetCpuCoresCostPerHourOk() (*float32, bool)`

GetCpuCoresCostPerHourOk returns a tuple with the CpuCoresCostPerHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuCoresCostPerHour

`func (o *SchemaComputePricingsRow) SetCpuCoresCostPerHour(v float32)`

SetCpuCoresCostPerHour sets CpuCoresCostPerHour field to given value.

### HasCpuCoresCostPerHour

`func (o *SchemaComputePricingsRow) HasCpuCoresCostPerHour() bool`

HasCpuCoresCostPerHour returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SchemaComputePricingsRow) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SchemaComputePricingsRow) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SchemaComputePricingsRow) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SchemaComputePricingsRow) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetId

`func (o *SchemaComputePricingsRow) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SchemaComputePricingsRow) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SchemaComputePricingsRow) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SchemaComputePricingsRow) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInstanceType

`func (o *SchemaComputePricingsRow) GetInstanceType() string`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *SchemaComputePricingsRow) GetInstanceTypeOk() (*string, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *SchemaComputePricingsRow) SetInstanceType(v string)`

SetInstanceType sets InstanceType field to given value.

### HasInstanceType

`func (o *SchemaComputePricingsRow) HasInstanceType() bool`

HasInstanceType returns a boolean if a field has been set.

### GetOperatingSystem

`func (o *SchemaComputePricingsRow) GetOperatingSystem() string`

GetOperatingSystem returns the OperatingSystem field if non-nil, zero value otherwise.

### GetOperatingSystemOk

`func (o *SchemaComputePricingsRow) GetOperatingSystemOk() (*string, bool)`

GetOperatingSystemOk returns a tuple with the OperatingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystem

`func (o *SchemaComputePricingsRow) SetOperatingSystem(v string)`

SetOperatingSystem sets OperatingSystem field to given value.

### HasOperatingSystem

`func (o *SchemaComputePricingsRow) HasOperatingSystem() bool`

HasOperatingSystem returns a boolean if a field has been set.

### GetPeriodBillingHours

`func (o *SchemaComputePricingsRow) GetPeriodBillingHours() float32`

GetPeriodBillingHours returns the PeriodBillingHours field if non-nil, zero value otherwise.

### GetPeriodBillingHoursOk

`func (o *SchemaComputePricingsRow) GetPeriodBillingHoursOk() (*float32, bool)`

GetPeriodBillingHoursOk returns a tuple with the PeriodBillingHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodBillingHours

`func (o *SchemaComputePricingsRow) SetPeriodBillingHours(v float32)`

SetPeriodBillingHours sets PeriodBillingHours field to given value.

### HasPeriodBillingHours

`func (o *SchemaComputePricingsRow) HasPeriodBillingHours() bool`

HasPeriodBillingHours returns a boolean if a field has been set.

### GetProvider

`func (o *SchemaComputePricingsRow) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *SchemaComputePricingsRow) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *SchemaComputePricingsRow) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *SchemaComputePricingsRow) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetRamGb

`func (o *SchemaComputePricingsRow) GetRamGb() float32`

GetRamGb returns the RamGb field if non-nil, zero value otherwise.

### GetRamGbOk

`func (o *SchemaComputePricingsRow) GetRamGbOk() (*float32, bool)`

GetRamGbOk returns a tuple with the RamGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRamGb

`func (o *SchemaComputePricingsRow) SetRamGb(v float32)`

SetRamGb sets RamGb field to given value.

### HasRamGb

`func (o *SchemaComputePricingsRow) HasRamGb() bool`

HasRamGb returns a boolean if a field has been set.

### GetRamGbCostPerHour

`func (o *SchemaComputePricingsRow) GetRamGbCostPerHour() float32`

GetRamGbCostPerHour returns the RamGbCostPerHour field if non-nil, zero value otherwise.

### GetRamGbCostPerHourOk

`func (o *SchemaComputePricingsRow) GetRamGbCostPerHourOk() (*float32, bool)`

GetRamGbCostPerHourOk returns a tuple with the RamGbCostPerHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRamGbCostPerHour

`func (o *SchemaComputePricingsRow) SetRamGbCostPerHour(v float32)`

SetRamGbCostPerHour sets RamGbCostPerHour field to given value.

### HasRamGbCostPerHour

`func (o *SchemaComputePricingsRow) HasRamGbCostPerHour() bool`

HasRamGbCostPerHour returns a boolean if a field has been set.

### GetRawPricingData

`func (o *SchemaComputePricingsRow) GetRawPricingData() map[string]interface{}`

GetRawPricingData returns the RawPricingData field if non-nil, zero value otherwise.

### GetRawPricingDataOk

`func (o *SchemaComputePricingsRow) GetRawPricingDataOk() (*map[string]interface{}, bool)`

GetRawPricingDataOk returns a tuple with the RawPricingData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawPricingData

`func (o *SchemaComputePricingsRow) SetRawPricingData(v map[string]interface{})`

SetRawPricingData sets RawPricingData field to given value.

### HasRawPricingData

`func (o *SchemaComputePricingsRow) HasRawPricingData() bool`

HasRawPricingData returns a boolean if a field has been set.

### GetRegion

`func (o *SchemaComputePricingsRow) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *SchemaComputePricingsRow) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *SchemaComputePricingsRow) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *SchemaComputePricingsRow) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetService

`func (o *SchemaComputePricingsRow) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *SchemaComputePricingsRow) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *SchemaComputePricingsRow) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *SchemaComputePricingsRow) HasService() bool`

HasService returns a boolean if a field has been set.

### GetTags

`func (o *SchemaComputePricingsRow) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *SchemaComputePricingsRow) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *SchemaComputePricingsRow) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *SchemaComputePricingsRow) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *SchemaComputePricingsRow) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SchemaComputePricingsRow) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SchemaComputePricingsRow) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *SchemaComputePricingsRow) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUsageType

`func (o *SchemaComputePricingsRow) GetUsageType() string`

GetUsageType returns the UsageType field if non-nil, zero value otherwise.

### GetUsageTypeOk

`func (o *SchemaComputePricingsRow) GetUsageTypeOk() (*string, bool)`

GetUsageTypeOk returns a tuple with the UsageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageType

`func (o *SchemaComputePricingsRow) SetUsageType(v string)`

SetUsageType sets UsageType field to given value.

### HasUsageType

`func (o *SchemaComputePricingsRow) HasUsageType() bool`

HasUsageType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



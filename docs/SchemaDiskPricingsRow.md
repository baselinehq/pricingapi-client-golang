# SchemaDiskPricingsRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvailabilityZone** | Pointer to **string** |  | [optional] 
**CostPerGbHour** | Pointer to **float32** |  | [optional] 
**CostPerIopsHour** | Pointer to **float32** |  | [optional] 
**CostPerThroughputMbpsHour** | Pointer to **float32** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**MaxCapacityGb** | Pointer to **float32** |  | [optional] 
**MaxIops** | Pointer to **float32** |  | [optional] 
**MaxThroughputMbps** | Pointer to **float32** |  | [optional] 
**MinCapacityGb** | Pointer to **float32** |  | [optional] 
**MinIops** | Pointer to **float32** |  | [optional] 
**MinThroughputMbps** | Pointer to **float32** |  | [optional] 
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

### NewSchemaDiskPricingsRow

`func NewSchemaDiskPricingsRow() *SchemaDiskPricingsRow`

NewSchemaDiskPricingsRow instantiates a new SchemaDiskPricingsRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchemaDiskPricingsRowWithDefaults

`func NewSchemaDiskPricingsRowWithDefaults() *SchemaDiskPricingsRow`

NewSchemaDiskPricingsRowWithDefaults instantiates a new SchemaDiskPricingsRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailabilityZone

`func (o *SchemaDiskPricingsRow) GetAvailabilityZone() string`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *SchemaDiskPricingsRow) GetAvailabilityZoneOk() (*string, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *SchemaDiskPricingsRow) SetAvailabilityZone(v string)`

SetAvailabilityZone sets AvailabilityZone field to given value.

### HasAvailabilityZone

`func (o *SchemaDiskPricingsRow) HasAvailabilityZone() bool`

HasAvailabilityZone returns a boolean if a field has been set.

### GetCostPerGbHour

`func (o *SchemaDiskPricingsRow) GetCostPerGbHour() float32`

GetCostPerGbHour returns the CostPerGbHour field if non-nil, zero value otherwise.

### GetCostPerGbHourOk

`func (o *SchemaDiskPricingsRow) GetCostPerGbHourOk() (*float32, bool)`

GetCostPerGbHourOk returns a tuple with the CostPerGbHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostPerGbHour

`func (o *SchemaDiskPricingsRow) SetCostPerGbHour(v float32)`

SetCostPerGbHour sets CostPerGbHour field to given value.

### HasCostPerGbHour

`func (o *SchemaDiskPricingsRow) HasCostPerGbHour() bool`

HasCostPerGbHour returns a boolean if a field has been set.

### GetCostPerIopsHour

`func (o *SchemaDiskPricingsRow) GetCostPerIopsHour() float32`

GetCostPerIopsHour returns the CostPerIopsHour field if non-nil, zero value otherwise.

### GetCostPerIopsHourOk

`func (o *SchemaDiskPricingsRow) GetCostPerIopsHourOk() (*float32, bool)`

GetCostPerIopsHourOk returns a tuple with the CostPerIopsHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostPerIopsHour

`func (o *SchemaDiskPricingsRow) SetCostPerIopsHour(v float32)`

SetCostPerIopsHour sets CostPerIopsHour field to given value.

### HasCostPerIopsHour

`func (o *SchemaDiskPricingsRow) HasCostPerIopsHour() bool`

HasCostPerIopsHour returns a boolean if a field has been set.

### GetCostPerThroughputMbpsHour

`func (o *SchemaDiskPricingsRow) GetCostPerThroughputMbpsHour() float32`

GetCostPerThroughputMbpsHour returns the CostPerThroughputMbpsHour field if non-nil, zero value otherwise.

### GetCostPerThroughputMbpsHourOk

`func (o *SchemaDiskPricingsRow) GetCostPerThroughputMbpsHourOk() (*float32, bool)`

GetCostPerThroughputMbpsHourOk returns a tuple with the CostPerThroughputMbpsHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostPerThroughputMbpsHour

`func (o *SchemaDiskPricingsRow) SetCostPerThroughputMbpsHour(v float32)`

SetCostPerThroughputMbpsHour sets CostPerThroughputMbpsHour field to given value.

### HasCostPerThroughputMbpsHour

`func (o *SchemaDiskPricingsRow) HasCostPerThroughputMbpsHour() bool`

HasCostPerThroughputMbpsHour returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SchemaDiskPricingsRow) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SchemaDiskPricingsRow) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SchemaDiskPricingsRow) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SchemaDiskPricingsRow) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetId

`func (o *SchemaDiskPricingsRow) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SchemaDiskPricingsRow) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SchemaDiskPricingsRow) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SchemaDiskPricingsRow) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMaxCapacityGb

`func (o *SchemaDiskPricingsRow) GetMaxCapacityGb() float32`

GetMaxCapacityGb returns the MaxCapacityGb field if non-nil, zero value otherwise.

### GetMaxCapacityGbOk

`func (o *SchemaDiskPricingsRow) GetMaxCapacityGbOk() (*float32, bool)`

GetMaxCapacityGbOk returns a tuple with the MaxCapacityGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCapacityGb

`func (o *SchemaDiskPricingsRow) SetMaxCapacityGb(v float32)`

SetMaxCapacityGb sets MaxCapacityGb field to given value.

### HasMaxCapacityGb

`func (o *SchemaDiskPricingsRow) HasMaxCapacityGb() bool`

HasMaxCapacityGb returns a boolean if a field has been set.

### GetMaxIops

`func (o *SchemaDiskPricingsRow) GetMaxIops() float32`

GetMaxIops returns the MaxIops field if non-nil, zero value otherwise.

### GetMaxIopsOk

`func (o *SchemaDiskPricingsRow) GetMaxIopsOk() (*float32, bool)`

GetMaxIopsOk returns a tuple with the MaxIops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxIops

`func (o *SchemaDiskPricingsRow) SetMaxIops(v float32)`

SetMaxIops sets MaxIops field to given value.

### HasMaxIops

`func (o *SchemaDiskPricingsRow) HasMaxIops() bool`

HasMaxIops returns a boolean if a field has been set.

### GetMaxThroughputMbps

`func (o *SchemaDiskPricingsRow) GetMaxThroughputMbps() float32`

GetMaxThroughputMbps returns the MaxThroughputMbps field if non-nil, zero value otherwise.

### GetMaxThroughputMbpsOk

`func (o *SchemaDiskPricingsRow) GetMaxThroughputMbpsOk() (*float32, bool)`

GetMaxThroughputMbpsOk returns a tuple with the MaxThroughputMbps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxThroughputMbps

`func (o *SchemaDiskPricingsRow) SetMaxThroughputMbps(v float32)`

SetMaxThroughputMbps sets MaxThroughputMbps field to given value.

### HasMaxThroughputMbps

`func (o *SchemaDiskPricingsRow) HasMaxThroughputMbps() bool`

HasMaxThroughputMbps returns a boolean if a field has been set.

### GetMinCapacityGb

`func (o *SchemaDiskPricingsRow) GetMinCapacityGb() float32`

GetMinCapacityGb returns the MinCapacityGb field if non-nil, zero value otherwise.

### GetMinCapacityGbOk

`func (o *SchemaDiskPricingsRow) GetMinCapacityGbOk() (*float32, bool)`

GetMinCapacityGbOk returns a tuple with the MinCapacityGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinCapacityGb

`func (o *SchemaDiskPricingsRow) SetMinCapacityGb(v float32)`

SetMinCapacityGb sets MinCapacityGb field to given value.

### HasMinCapacityGb

`func (o *SchemaDiskPricingsRow) HasMinCapacityGb() bool`

HasMinCapacityGb returns a boolean if a field has been set.

### GetMinIops

`func (o *SchemaDiskPricingsRow) GetMinIops() float32`

GetMinIops returns the MinIops field if non-nil, zero value otherwise.

### GetMinIopsOk

`func (o *SchemaDiskPricingsRow) GetMinIopsOk() (*float32, bool)`

GetMinIopsOk returns a tuple with the MinIops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinIops

`func (o *SchemaDiskPricingsRow) SetMinIops(v float32)`

SetMinIops sets MinIops field to given value.

### HasMinIops

`func (o *SchemaDiskPricingsRow) HasMinIops() bool`

HasMinIops returns a boolean if a field has been set.

### GetMinThroughputMbps

`func (o *SchemaDiskPricingsRow) GetMinThroughputMbps() float32`

GetMinThroughputMbps returns the MinThroughputMbps field if non-nil, zero value otherwise.

### GetMinThroughputMbpsOk

`func (o *SchemaDiskPricingsRow) GetMinThroughputMbpsOk() (*float32, bool)`

GetMinThroughputMbpsOk returns a tuple with the MinThroughputMbps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinThroughputMbps

`func (o *SchemaDiskPricingsRow) SetMinThroughputMbps(v float32)`

SetMinThroughputMbps sets MinThroughputMbps field to given value.

### HasMinThroughputMbps

`func (o *SchemaDiskPricingsRow) HasMinThroughputMbps() bool`

HasMinThroughputMbps returns a boolean if a field has been set.

### GetPeriodBillingHours

`func (o *SchemaDiskPricingsRow) GetPeriodBillingHours() float32`

GetPeriodBillingHours returns the PeriodBillingHours field if non-nil, zero value otherwise.

### GetPeriodBillingHoursOk

`func (o *SchemaDiskPricingsRow) GetPeriodBillingHoursOk() (*float32, bool)`

GetPeriodBillingHoursOk returns a tuple with the PeriodBillingHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodBillingHours

`func (o *SchemaDiskPricingsRow) SetPeriodBillingHours(v float32)`

SetPeriodBillingHours sets PeriodBillingHours field to given value.

### HasPeriodBillingHours

`func (o *SchemaDiskPricingsRow) HasPeriodBillingHours() bool`

HasPeriodBillingHours returns a boolean if a field has been set.

### GetProvider

`func (o *SchemaDiskPricingsRow) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *SchemaDiskPricingsRow) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *SchemaDiskPricingsRow) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *SchemaDiskPricingsRow) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetRawPricingData

`func (o *SchemaDiskPricingsRow) GetRawPricingData() map[string]interface{}`

GetRawPricingData returns the RawPricingData field if non-nil, zero value otherwise.

### GetRawPricingDataOk

`func (o *SchemaDiskPricingsRow) GetRawPricingDataOk() (*map[string]interface{}, bool)`

GetRawPricingDataOk returns a tuple with the RawPricingData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawPricingData

`func (o *SchemaDiskPricingsRow) SetRawPricingData(v map[string]interface{})`

SetRawPricingData sets RawPricingData field to given value.

### HasRawPricingData

`func (o *SchemaDiskPricingsRow) HasRawPricingData() bool`

HasRawPricingData returns a boolean if a field has been set.

### GetRegion

`func (o *SchemaDiskPricingsRow) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *SchemaDiskPricingsRow) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *SchemaDiskPricingsRow) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *SchemaDiskPricingsRow) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetService

`func (o *SchemaDiskPricingsRow) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *SchemaDiskPricingsRow) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *SchemaDiskPricingsRow) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *SchemaDiskPricingsRow) HasService() bool`

HasService returns a boolean if a field has been set.

### GetTags

`func (o *SchemaDiskPricingsRow) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *SchemaDiskPricingsRow) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *SchemaDiskPricingsRow) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *SchemaDiskPricingsRow) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetType

`func (o *SchemaDiskPricingsRow) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SchemaDiskPricingsRow) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SchemaDiskPricingsRow) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SchemaDiskPricingsRow) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *SchemaDiskPricingsRow) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SchemaDiskPricingsRow) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SchemaDiskPricingsRow) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *SchemaDiskPricingsRow) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUsageType

`func (o *SchemaDiskPricingsRow) GetUsageType() string`

GetUsageType returns the UsageType field if non-nil, zero value otherwise.

### GetUsageTypeOk

`func (o *SchemaDiskPricingsRow) GetUsageTypeOk() (*string, bool)`

GetUsageTypeOk returns a tuple with the UsageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageType

`func (o *SchemaDiskPricingsRow) SetUsageType(v string)`

SetUsageType sets UsageType field to given value.

### HasUsageType

`func (o *SchemaDiskPricingsRow) HasUsageType() bool`

HasUsageType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



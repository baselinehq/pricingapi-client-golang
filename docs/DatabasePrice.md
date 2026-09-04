# DatabasePrice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Architecture** | Pointer to **string** |  | [optional] 
**AvailabilityZone** | Pointer to **string** |  | [optional] 
**BillingConfig** | Pointer to **string** |  | [optional] 
**CapacityUnit** | Pointer to **string** |  | [optional] 
**CostPerCapacityUnit** | Pointer to **float32** |  | [optional] 
**CostPerHour** | Pointer to **float32** |  | [optional] 
**CostPerIoRequest** | Pointer to **float32** |  | [optional] 
**CostPerIopsHour** | Pointer to **float32** |  | [optional] 
**CostPerStorageGbHour** | Pointer to **float32** |  | [optional] 
**CostPerThroughputMbpsHour** | Pointer to **float32** |  | [optional] 
**CostPerVcpuLicenseHour** | Pointer to **float32** |  | [optional] 
**CpuCores** | Pointer to **float32** |  | [optional] 
**CpuCoresCostPerHour** | Pointer to **float32** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**DeploymentOption** | Pointer to **string** |  | [optional] 
**Edition** | Pointer to **string** |  | [optional] 
**Engine** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**InstanceType** | Pointer to **string** |  | [optional] 
**MaxIops** | Pointer to **float32** |  | [optional] 
**MaxStorageGb** | Pointer to **float32** |  | [optional] 
**MaxThroughputMbps** | Pointer to **float32** |  | [optional] 
**MinIops** | Pointer to **float32** |  | [optional] 
**MinStorageGb** | Pointer to **float32** |  | [optional] 
**MinThroughputMbps** | Pointer to **float32** |  | [optional] 
**NativeCostPerHour** | Pointer to **float32** |  | [optional] 
**NativeCostPerStorageGbHour** | Pointer to **float32** |  | [optional] 
**PeriodBillingHours** | Pointer to **float32** |  | [optional] 
**Provider** | Pointer to **string** |  | [optional] 
**RamGb** | Pointer to **float32** |  | [optional] 
**RamGbCostPerHour** | Pointer to **float32** |  | [optional] 
**RawPricingData** | Pointer to **map[string]interface{}** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**Service** | Pointer to **string** |  | [optional] 
**StorageType** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **map[string]interface{}** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**UsageType** | Pointer to **string** |  | [optional] 

## Methods

### NewDatabasePrice

`func NewDatabasePrice() *DatabasePrice`

NewDatabasePrice instantiates a new DatabasePrice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabasePriceWithDefaults

`func NewDatabasePriceWithDefaults() *DatabasePrice`

NewDatabasePriceWithDefaults instantiates a new DatabasePrice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchitecture

`func (o *DatabasePrice) GetArchitecture() string`

GetArchitecture returns the Architecture field if non-nil, zero value otherwise.

### GetArchitectureOk

`func (o *DatabasePrice) GetArchitectureOk() (*string, bool)`

GetArchitectureOk returns a tuple with the Architecture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchitecture

`func (o *DatabasePrice) SetArchitecture(v string)`

SetArchitecture sets Architecture field to given value.

### HasArchitecture

`func (o *DatabasePrice) HasArchitecture() bool`

HasArchitecture returns a boolean if a field has been set.

### GetAvailabilityZone

`func (o *DatabasePrice) GetAvailabilityZone() string`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *DatabasePrice) GetAvailabilityZoneOk() (*string, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *DatabasePrice) SetAvailabilityZone(v string)`

SetAvailabilityZone sets AvailabilityZone field to given value.

### HasAvailabilityZone

`func (o *DatabasePrice) HasAvailabilityZone() bool`

HasAvailabilityZone returns a boolean if a field has been set.

### GetBillingConfig

`func (o *DatabasePrice) GetBillingConfig() string`

GetBillingConfig returns the BillingConfig field if non-nil, zero value otherwise.

### GetBillingConfigOk

`func (o *DatabasePrice) GetBillingConfigOk() (*string, bool)`

GetBillingConfigOk returns a tuple with the BillingConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingConfig

`func (o *DatabasePrice) SetBillingConfig(v string)`

SetBillingConfig sets BillingConfig field to given value.

### HasBillingConfig

`func (o *DatabasePrice) HasBillingConfig() bool`

HasBillingConfig returns a boolean if a field has been set.

### GetCapacityUnit

`func (o *DatabasePrice) GetCapacityUnit() string`

GetCapacityUnit returns the CapacityUnit field if non-nil, zero value otherwise.

### GetCapacityUnitOk

`func (o *DatabasePrice) GetCapacityUnitOk() (*string, bool)`

GetCapacityUnitOk returns a tuple with the CapacityUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacityUnit

`func (o *DatabasePrice) SetCapacityUnit(v string)`

SetCapacityUnit sets CapacityUnit field to given value.

### HasCapacityUnit

`func (o *DatabasePrice) HasCapacityUnit() bool`

HasCapacityUnit returns a boolean if a field has been set.

### GetCostPerCapacityUnit

`func (o *DatabasePrice) GetCostPerCapacityUnit() float32`

GetCostPerCapacityUnit returns the CostPerCapacityUnit field if non-nil, zero value otherwise.

### GetCostPerCapacityUnitOk

`func (o *DatabasePrice) GetCostPerCapacityUnitOk() (*float32, bool)`

GetCostPerCapacityUnitOk returns a tuple with the CostPerCapacityUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostPerCapacityUnit

`func (o *DatabasePrice) SetCostPerCapacityUnit(v float32)`

SetCostPerCapacityUnit sets CostPerCapacityUnit field to given value.

### HasCostPerCapacityUnit

`func (o *DatabasePrice) HasCostPerCapacityUnit() bool`

HasCostPerCapacityUnit returns a boolean if a field has been set.

### GetCostPerHour

`func (o *DatabasePrice) GetCostPerHour() float32`

GetCostPerHour returns the CostPerHour field if non-nil, zero value otherwise.

### GetCostPerHourOk

`func (o *DatabasePrice) GetCostPerHourOk() (*float32, bool)`

GetCostPerHourOk returns a tuple with the CostPerHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostPerHour

`func (o *DatabasePrice) SetCostPerHour(v float32)`

SetCostPerHour sets CostPerHour field to given value.

### HasCostPerHour

`func (o *DatabasePrice) HasCostPerHour() bool`

HasCostPerHour returns a boolean if a field has been set.

### GetCostPerIoRequest

`func (o *DatabasePrice) GetCostPerIoRequest() float32`

GetCostPerIoRequest returns the CostPerIoRequest field if non-nil, zero value otherwise.

### GetCostPerIoRequestOk

`func (o *DatabasePrice) GetCostPerIoRequestOk() (*float32, bool)`

GetCostPerIoRequestOk returns a tuple with the CostPerIoRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostPerIoRequest

`func (o *DatabasePrice) SetCostPerIoRequest(v float32)`

SetCostPerIoRequest sets CostPerIoRequest field to given value.

### HasCostPerIoRequest

`func (o *DatabasePrice) HasCostPerIoRequest() bool`

HasCostPerIoRequest returns a boolean if a field has been set.

### GetCostPerIopsHour

`func (o *DatabasePrice) GetCostPerIopsHour() float32`

GetCostPerIopsHour returns the CostPerIopsHour field if non-nil, zero value otherwise.

### GetCostPerIopsHourOk

`func (o *DatabasePrice) GetCostPerIopsHourOk() (*float32, bool)`

GetCostPerIopsHourOk returns a tuple with the CostPerIopsHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostPerIopsHour

`func (o *DatabasePrice) SetCostPerIopsHour(v float32)`

SetCostPerIopsHour sets CostPerIopsHour field to given value.

### HasCostPerIopsHour

`func (o *DatabasePrice) HasCostPerIopsHour() bool`

HasCostPerIopsHour returns a boolean if a field has been set.

### GetCostPerStorageGbHour

`func (o *DatabasePrice) GetCostPerStorageGbHour() float32`

GetCostPerStorageGbHour returns the CostPerStorageGbHour field if non-nil, zero value otherwise.

### GetCostPerStorageGbHourOk

`func (o *DatabasePrice) GetCostPerStorageGbHourOk() (*float32, bool)`

GetCostPerStorageGbHourOk returns a tuple with the CostPerStorageGbHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostPerStorageGbHour

`func (o *DatabasePrice) SetCostPerStorageGbHour(v float32)`

SetCostPerStorageGbHour sets CostPerStorageGbHour field to given value.

### HasCostPerStorageGbHour

`func (o *DatabasePrice) HasCostPerStorageGbHour() bool`

HasCostPerStorageGbHour returns a boolean if a field has been set.

### GetCostPerThroughputMbpsHour

`func (o *DatabasePrice) GetCostPerThroughputMbpsHour() float32`

GetCostPerThroughputMbpsHour returns the CostPerThroughputMbpsHour field if non-nil, zero value otherwise.

### GetCostPerThroughputMbpsHourOk

`func (o *DatabasePrice) GetCostPerThroughputMbpsHourOk() (*float32, bool)`

GetCostPerThroughputMbpsHourOk returns a tuple with the CostPerThroughputMbpsHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostPerThroughputMbpsHour

`func (o *DatabasePrice) SetCostPerThroughputMbpsHour(v float32)`

SetCostPerThroughputMbpsHour sets CostPerThroughputMbpsHour field to given value.

### HasCostPerThroughputMbpsHour

`func (o *DatabasePrice) HasCostPerThroughputMbpsHour() bool`

HasCostPerThroughputMbpsHour returns a boolean if a field has been set.

### GetCostPerVcpuLicenseHour

`func (o *DatabasePrice) GetCostPerVcpuLicenseHour() float32`

GetCostPerVcpuLicenseHour returns the CostPerVcpuLicenseHour field if non-nil, zero value otherwise.

### GetCostPerVcpuLicenseHourOk

`func (o *DatabasePrice) GetCostPerVcpuLicenseHourOk() (*float32, bool)`

GetCostPerVcpuLicenseHourOk returns a tuple with the CostPerVcpuLicenseHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostPerVcpuLicenseHour

`func (o *DatabasePrice) SetCostPerVcpuLicenseHour(v float32)`

SetCostPerVcpuLicenseHour sets CostPerVcpuLicenseHour field to given value.

### HasCostPerVcpuLicenseHour

`func (o *DatabasePrice) HasCostPerVcpuLicenseHour() bool`

HasCostPerVcpuLicenseHour returns a boolean if a field has been set.

### GetCpuCores

`func (o *DatabasePrice) GetCpuCores() float32`

GetCpuCores returns the CpuCores field if non-nil, zero value otherwise.

### GetCpuCoresOk

`func (o *DatabasePrice) GetCpuCoresOk() (*float32, bool)`

GetCpuCoresOk returns a tuple with the CpuCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuCores

`func (o *DatabasePrice) SetCpuCores(v float32)`

SetCpuCores sets CpuCores field to given value.

### HasCpuCores

`func (o *DatabasePrice) HasCpuCores() bool`

HasCpuCores returns a boolean if a field has been set.

### GetCpuCoresCostPerHour

`func (o *DatabasePrice) GetCpuCoresCostPerHour() float32`

GetCpuCoresCostPerHour returns the CpuCoresCostPerHour field if non-nil, zero value otherwise.

### GetCpuCoresCostPerHourOk

`func (o *DatabasePrice) GetCpuCoresCostPerHourOk() (*float32, bool)`

GetCpuCoresCostPerHourOk returns a tuple with the CpuCoresCostPerHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuCoresCostPerHour

`func (o *DatabasePrice) SetCpuCoresCostPerHour(v float32)`

SetCpuCoresCostPerHour sets CpuCoresCostPerHour field to given value.

### HasCpuCoresCostPerHour

`func (o *DatabasePrice) HasCpuCoresCostPerHour() bool`

HasCpuCoresCostPerHour returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DatabasePrice) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DatabasePrice) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DatabasePrice) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DatabasePrice) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCurrency

`func (o *DatabasePrice) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *DatabasePrice) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *DatabasePrice) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *DatabasePrice) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDeploymentOption

`func (o *DatabasePrice) GetDeploymentOption() string`

GetDeploymentOption returns the DeploymentOption field if non-nil, zero value otherwise.

### GetDeploymentOptionOk

`func (o *DatabasePrice) GetDeploymentOptionOk() (*string, bool)`

GetDeploymentOptionOk returns a tuple with the DeploymentOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentOption

`func (o *DatabasePrice) SetDeploymentOption(v string)`

SetDeploymentOption sets DeploymentOption field to given value.

### HasDeploymentOption

`func (o *DatabasePrice) HasDeploymentOption() bool`

HasDeploymentOption returns a boolean if a field has been set.

### GetEdition

`func (o *DatabasePrice) GetEdition() string`

GetEdition returns the Edition field if non-nil, zero value otherwise.

### GetEditionOk

`func (o *DatabasePrice) GetEditionOk() (*string, bool)`

GetEditionOk returns a tuple with the Edition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdition

`func (o *DatabasePrice) SetEdition(v string)`

SetEdition sets Edition field to given value.

### HasEdition

`func (o *DatabasePrice) HasEdition() bool`

HasEdition returns a boolean if a field has been set.

### GetEngine

`func (o *DatabasePrice) GetEngine() string`

GetEngine returns the Engine field if non-nil, zero value otherwise.

### GetEngineOk

`func (o *DatabasePrice) GetEngineOk() (*string, bool)`

GetEngineOk returns a tuple with the Engine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngine

`func (o *DatabasePrice) SetEngine(v string)`

SetEngine sets Engine field to given value.

### HasEngine

`func (o *DatabasePrice) HasEngine() bool`

HasEngine returns a boolean if a field has been set.

### GetId

`func (o *DatabasePrice) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DatabasePrice) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DatabasePrice) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DatabasePrice) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInstanceType

`func (o *DatabasePrice) GetInstanceType() string`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *DatabasePrice) GetInstanceTypeOk() (*string, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *DatabasePrice) SetInstanceType(v string)`

SetInstanceType sets InstanceType field to given value.

### HasInstanceType

`func (o *DatabasePrice) HasInstanceType() bool`

HasInstanceType returns a boolean if a field has been set.

### GetMaxIops

`func (o *DatabasePrice) GetMaxIops() float32`

GetMaxIops returns the MaxIops field if non-nil, zero value otherwise.

### GetMaxIopsOk

`func (o *DatabasePrice) GetMaxIopsOk() (*float32, bool)`

GetMaxIopsOk returns a tuple with the MaxIops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxIops

`func (o *DatabasePrice) SetMaxIops(v float32)`

SetMaxIops sets MaxIops field to given value.

### HasMaxIops

`func (o *DatabasePrice) HasMaxIops() bool`

HasMaxIops returns a boolean if a field has been set.

### GetMaxStorageGb

`func (o *DatabasePrice) GetMaxStorageGb() float32`

GetMaxStorageGb returns the MaxStorageGb field if non-nil, zero value otherwise.

### GetMaxStorageGbOk

`func (o *DatabasePrice) GetMaxStorageGbOk() (*float32, bool)`

GetMaxStorageGbOk returns a tuple with the MaxStorageGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStorageGb

`func (o *DatabasePrice) SetMaxStorageGb(v float32)`

SetMaxStorageGb sets MaxStorageGb field to given value.

### HasMaxStorageGb

`func (o *DatabasePrice) HasMaxStorageGb() bool`

HasMaxStorageGb returns a boolean if a field has been set.

### GetMaxThroughputMbps

`func (o *DatabasePrice) GetMaxThroughputMbps() float32`

GetMaxThroughputMbps returns the MaxThroughputMbps field if non-nil, zero value otherwise.

### GetMaxThroughputMbpsOk

`func (o *DatabasePrice) GetMaxThroughputMbpsOk() (*float32, bool)`

GetMaxThroughputMbpsOk returns a tuple with the MaxThroughputMbps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxThroughputMbps

`func (o *DatabasePrice) SetMaxThroughputMbps(v float32)`

SetMaxThroughputMbps sets MaxThroughputMbps field to given value.

### HasMaxThroughputMbps

`func (o *DatabasePrice) HasMaxThroughputMbps() bool`

HasMaxThroughputMbps returns a boolean if a field has been set.

### GetMinIops

`func (o *DatabasePrice) GetMinIops() float32`

GetMinIops returns the MinIops field if non-nil, zero value otherwise.

### GetMinIopsOk

`func (o *DatabasePrice) GetMinIopsOk() (*float32, bool)`

GetMinIopsOk returns a tuple with the MinIops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinIops

`func (o *DatabasePrice) SetMinIops(v float32)`

SetMinIops sets MinIops field to given value.

### HasMinIops

`func (o *DatabasePrice) HasMinIops() bool`

HasMinIops returns a boolean if a field has been set.

### GetMinStorageGb

`func (o *DatabasePrice) GetMinStorageGb() float32`

GetMinStorageGb returns the MinStorageGb field if non-nil, zero value otherwise.

### GetMinStorageGbOk

`func (o *DatabasePrice) GetMinStorageGbOk() (*float32, bool)`

GetMinStorageGbOk returns a tuple with the MinStorageGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinStorageGb

`func (o *DatabasePrice) SetMinStorageGb(v float32)`

SetMinStorageGb sets MinStorageGb field to given value.

### HasMinStorageGb

`func (o *DatabasePrice) HasMinStorageGb() bool`

HasMinStorageGb returns a boolean if a field has been set.

### GetMinThroughputMbps

`func (o *DatabasePrice) GetMinThroughputMbps() float32`

GetMinThroughputMbps returns the MinThroughputMbps field if non-nil, zero value otherwise.

### GetMinThroughputMbpsOk

`func (o *DatabasePrice) GetMinThroughputMbpsOk() (*float32, bool)`

GetMinThroughputMbpsOk returns a tuple with the MinThroughputMbps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinThroughputMbps

`func (o *DatabasePrice) SetMinThroughputMbps(v float32)`

SetMinThroughputMbps sets MinThroughputMbps field to given value.

### HasMinThroughputMbps

`func (o *DatabasePrice) HasMinThroughputMbps() bool`

HasMinThroughputMbps returns a boolean if a field has been set.

### GetNativeCostPerHour

`func (o *DatabasePrice) GetNativeCostPerHour() float32`

GetNativeCostPerHour returns the NativeCostPerHour field if non-nil, zero value otherwise.

### GetNativeCostPerHourOk

`func (o *DatabasePrice) GetNativeCostPerHourOk() (*float32, bool)`

GetNativeCostPerHourOk returns a tuple with the NativeCostPerHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNativeCostPerHour

`func (o *DatabasePrice) SetNativeCostPerHour(v float32)`

SetNativeCostPerHour sets NativeCostPerHour field to given value.

### HasNativeCostPerHour

`func (o *DatabasePrice) HasNativeCostPerHour() bool`

HasNativeCostPerHour returns a boolean if a field has been set.

### GetNativeCostPerStorageGbHour

`func (o *DatabasePrice) GetNativeCostPerStorageGbHour() float32`

GetNativeCostPerStorageGbHour returns the NativeCostPerStorageGbHour field if non-nil, zero value otherwise.

### GetNativeCostPerStorageGbHourOk

`func (o *DatabasePrice) GetNativeCostPerStorageGbHourOk() (*float32, bool)`

GetNativeCostPerStorageGbHourOk returns a tuple with the NativeCostPerStorageGbHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNativeCostPerStorageGbHour

`func (o *DatabasePrice) SetNativeCostPerStorageGbHour(v float32)`

SetNativeCostPerStorageGbHour sets NativeCostPerStorageGbHour field to given value.

### HasNativeCostPerStorageGbHour

`func (o *DatabasePrice) HasNativeCostPerStorageGbHour() bool`

HasNativeCostPerStorageGbHour returns a boolean if a field has been set.

### GetPeriodBillingHours

`func (o *DatabasePrice) GetPeriodBillingHours() float32`

GetPeriodBillingHours returns the PeriodBillingHours field if non-nil, zero value otherwise.

### GetPeriodBillingHoursOk

`func (o *DatabasePrice) GetPeriodBillingHoursOk() (*float32, bool)`

GetPeriodBillingHoursOk returns a tuple with the PeriodBillingHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodBillingHours

`func (o *DatabasePrice) SetPeriodBillingHours(v float32)`

SetPeriodBillingHours sets PeriodBillingHours field to given value.

### HasPeriodBillingHours

`func (o *DatabasePrice) HasPeriodBillingHours() bool`

HasPeriodBillingHours returns a boolean if a field has been set.

### GetProvider

`func (o *DatabasePrice) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *DatabasePrice) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *DatabasePrice) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *DatabasePrice) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetRamGb

`func (o *DatabasePrice) GetRamGb() float32`

GetRamGb returns the RamGb field if non-nil, zero value otherwise.

### GetRamGbOk

`func (o *DatabasePrice) GetRamGbOk() (*float32, bool)`

GetRamGbOk returns a tuple with the RamGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRamGb

`func (o *DatabasePrice) SetRamGb(v float32)`

SetRamGb sets RamGb field to given value.

### HasRamGb

`func (o *DatabasePrice) HasRamGb() bool`

HasRamGb returns a boolean if a field has been set.

### GetRamGbCostPerHour

`func (o *DatabasePrice) GetRamGbCostPerHour() float32`

GetRamGbCostPerHour returns the RamGbCostPerHour field if non-nil, zero value otherwise.

### GetRamGbCostPerHourOk

`func (o *DatabasePrice) GetRamGbCostPerHourOk() (*float32, bool)`

GetRamGbCostPerHourOk returns a tuple with the RamGbCostPerHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRamGbCostPerHour

`func (o *DatabasePrice) SetRamGbCostPerHour(v float32)`

SetRamGbCostPerHour sets RamGbCostPerHour field to given value.

### HasRamGbCostPerHour

`func (o *DatabasePrice) HasRamGbCostPerHour() bool`

HasRamGbCostPerHour returns a boolean if a field has been set.

### GetRawPricingData

`func (o *DatabasePrice) GetRawPricingData() map[string]interface{}`

GetRawPricingData returns the RawPricingData field if non-nil, zero value otherwise.

### GetRawPricingDataOk

`func (o *DatabasePrice) GetRawPricingDataOk() (*map[string]interface{}, bool)`

GetRawPricingDataOk returns a tuple with the RawPricingData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawPricingData

`func (o *DatabasePrice) SetRawPricingData(v map[string]interface{})`

SetRawPricingData sets RawPricingData field to given value.

### HasRawPricingData

`func (o *DatabasePrice) HasRawPricingData() bool`

HasRawPricingData returns a boolean if a field has been set.

### GetRegion

`func (o *DatabasePrice) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *DatabasePrice) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *DatabasePrice) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *DatabasePrice) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetService

`func (o *DatabasePrice) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *DatabasePrice) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *DatabasePrice) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *DatabasePrice) HasService() bool`

HasService returns a boolean if a field has been set.

### GetStorageType

`func (o *DatabasePrice) GetStorageType() string`

GetStorageType returns the StorageType field if non-nil, zero value otherwise.

### GetStorageTypeOk

`func (o *DatabasePrice) GetStorageTypeOk() (*string, bool)`

GetStorageTypeOk returns a tuple with the StorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageType

`func (o *DatabasePrice) SetStorageType(v string)`

SetStorageType sets StorageType field to given value.

### HasStorageType

`func (o *DatabasePrice) HasStorageType() bool`

HasStorageType returns a boolean if a field has been set.

### GetTags

`func (o *DatabasePrice) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DatabasePrice) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DatabasePrice) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *DatabasePrice) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DatabasePrice) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DatabasePrice) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DatabasePrice) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DatabasePrice) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUsageType

`func (o *DatabasePrice) GetUsageType() string`

GetUsageType returns the UsageType field if non-nil, zero value otherwise.

### GetUsageTypeOk

`func (o *DatabasePrice) GetUsageTypeOk() (*string, bool)`

GetUsageTypeOk returns a tuple with the UsageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageType

`func (o *DatabasePrice) SetUsageType(v string)`

SetUsageType sets UsageType field to given value.

### HasUsageType

`func (o *DatabasePrice) HasUsageType() bool`

HasUsageType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



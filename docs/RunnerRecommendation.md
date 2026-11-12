# RunnerRecommendation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Architecture** | Pointer to **string** |  | [optional] 
**BilledSeconds** | Pointer to **float32** |  | [optional] 
**CandidateClass** | Pointer to **string** |  | [optional] 
**ComputeSource** | Pointer to **string** |  | [optional] 
**Connected** | Pointer to **bool** |  | [optional] 
**Cost** | Pointer to **float32** |  | [optional] 
**CostPerMonth** | Pointer to **float32** |  | [optional] 
**CpuCores** | Pointer to **int32** |  | [optional] 
**CurrentP75JobCost** | Pointer to **float32** |  | [optional] 
**CurrentP75WallSeconds** | Pointer to **float32** |  | [optional] 
**CurrentWallSeconds** | Pointer to **float32** |  | [optional] 
**Docker** | Pointer to **bool** |  | [optional] 
**FetchedOn** | Pointer to **string** |  | [optional] 
**Flags** | Pointer to **[]string** |  | [optional] 
**GpuCount** | Pointer to **int32** |  | [optional] 
**IncludedApplied** | Pointer to **float32** |  | [optional] 
**NestedVirtualization** | Pointer to **bool** |  | [optional] 
**OperatingSystem** | Pointer to **string** |  | [optional] 
**P75JobCost** | Pointer to **float32** |  | [optional] 
**P75WallSeconds** | Pointer to **float32** |  | [optional] 
**PlanFees** | Pointer to [**[]TypesRunnerPlanFee**](TypesRunnerPlanFee.md) |  | [optional] 
**PricingModel** | Pointer to **string** |  | [optional] 
**Provider** | Pointer to **string** |  | [optional] 
**RamGb** | Pointer to **float32** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**RuntimeSeconds** | Pointer to **float32** |  | [optional] 
**SavingsVsCurrent** | Pointer to **float32** |  | [optional] 
**Sku** | Pointer to **string** |  | [optional] 
**SourceUrl** | Pointer to **string** |  | [optional] 
**SpeedMeasured** | Pointer to **bool** |  | [optional] 
**VmInstanceType** | Pointer to **string** |  | [optional] 
**WallSeconds** | Pointer to **float32** |  | [optional] 

## Methods

### NewRunnerRecommendation

`func NewRunnerRecommendation() *RunnerRecommendation`

NewRunnerRecommendation instantiates a new RunnerRecommendation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunnerRecommendationWithDefaults

`func NewRunnerRecommendationWithDefaults() *RunnerRecommendation`

NewRunnerRecommendationWithDefaults instantiates a new RunnerRecommendation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchitecture

`func (o *RunnerRecommendation) GetArchitecture() string`

GetArchitecture returns the Architecture field if non-nil, zero value otherwise.

### GetArchitectureOk

`func (o *RunnerRecommendation) GetArchitectureOk() (*string, bool)`

GetArchitectureOk returns a tuple with the Architecture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchitecture

`func (o *RunnerRecommendation) SetArchitecture(v string)`

SetArchitecture sets Architecture field to given value.

### HasArchitecture

`func (o *RunnerRecommendation) HasArchitecture() bool`

HasArchitecture returns a boolean if a field has been set.

### GetBilledSeconds

`func (o *RunnerRecommendation) GetBilledSeconds() float32`

GetBilledSeconds returns the BilledSeconds field if non-nil, zero value otherwise.

### GetBilledSecondsOk

`func (o *RunnerRecommendation) GetBilledSecondsOk() (*float32, bool)`

GetBilledSecondsOk returns a tuple with the BilledSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBilledSeconds

`func (o *RunnerRecommendation) SetBilledSeconds(v float32)`

SetBilledSeconds sets BilledSeconds field to given value.

### HasBilledSeconds

`func (o *RunnerRecommendation) HasBilledSeconds() bool`

HasBilledSeconds returns a boolean if a field has been set.

### GetCandidateClass

`func (o *RunnerRecommendation) GetCandidateClass() string`

GetCandidateClass returns the CandidateClass field if non-nil, zero value otherwise.

### GetCandidateClassOk

`func (o *RunnerRecommendation) GetCandidateClassOk() (*string, bool)`

GetCandidateClassOk returns a tuple with the CandidateClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCandidateClass

`func (o *RunnerRecommendation) SetCandidateClass(v string)`

SetCandidateClass sets CandidateClass field to given value.

### HasCandidateClass

`func (o *RunnerRecommendation) HasCandidateClass() bool`

HasCandidateClass returns a boolean if a field has been set.

### GetComputeSource

`func (o *RunnerRecommendation) GetComputeSource() string`

GetComputeSource returns the ComputeSource field if non-nil, zero value otherwise.

### GetComputeSourceOk

`func (o *RunnerRecommendation) GetComputeSourceOk() (*string, bool)`

GetComputeSourceOk returns a tuple with the ComputeSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeSource

`func (o *RunnerRecommendation) SetComputeSource(v string)`

SetComputeSource sets ComputeSource field to given value.

### HasComputeSource

`func (o *RunnerRecommendation) HasComputeSource() bool`

HasComputeSource returns a boolean if a field has been set.

### GetConnected

`func (o *RunnerRecommendation) GetConnected() bool`

GetConnected returns the Connected field if non-nil, zero value otherwise.

### GetConnectedOk

`func (o *RunnerRecommendation) GetConnectedOk() (*bool, bool)`

GetConnectedOk returns a tuple with the Connected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnected

`func (o *RunnerRecommendation) SetConnected(v bool)`

SetConnected sets Connected field to given value.

### HasConnected

`func (o *RunnerRecommendation) HasConnected() bool`

HasConnected returns a boolean if a field has been set.

### GetCost

`func (o *RunnerRecommendation) GetCost() float32`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *RunnerRecommendation) GetCostOk() (*float32, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *RunnerRecommendation) SetCost(v float32)`

SetCost sets Cost field to given value.

### HasCost

`func (o *RunnerRecommendation) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetCostPerMonth

`func (o *RunnerRecommendation) GetCostPerMonth() float32`

GetCostPerMonth returns the CostPerMonth field if non-nil, zero value otherwise.

### GetCostPerMonthOk

`func (o *RunnerRecommendation) GetCostPerMonthOk() (*float32, bool)`

GetCostPerMonthOk returns a tuple with the CostPerMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostPerMonth

`func (o *RunnerRecommendation) SetCostPerMonth(v float32)`

SetCostPerMonth sets CostPerMonth field to given value.

### HasCostPerMonth

`func (o *RunnerRecommendation) HasCostPerMonth() bool`

HasCostPerMonth returns a boolean if a field has been set.

### GetCpuCores

`func (o *RunnerRecommendation) GetCpuCores() int32`

GetCpuCores returns the CpuCores field if non-nil, zero value otherwise.

### GetCpuCoresOk

`func (o *RunnerRecommendation) GetCpuCoresOk() (*int32, bool)`

GetCpuCoresOk returns a tuple with the CpuCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuCores

`func (o *RunnerRecommendation) SetCpuCores(v int32)`

SetCpuCores sets CpuCores field to given value.

### HasCpuCores

`func (o *RunnerRecommendation) HasCpuCores() bool`

HasCpuCores returns a boolean if a field has been set.

### GetCurrentP75JobCost

`func (o *RunnerRecommendation) GetCurrentP75JobCost() float32`

GetCurrentP75JobCost returns the CurrentP75JobCost field if non-nil, zero value otherwise.

### GetCurrentP75JobCostOk

`func (o *RunnerRecommendation) GetCurrentP75JobCostOk() (*float32, bool)`

GetCurrentP75JobCostOk returns a tuple with the CurrentP75JobCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentP75JobCost

`func (o *RunnerRecommendation) SetCurrentP75JobCost(v float32)`

SetCurrentP75JobCost sets CurrentP75JobCost field to given value.

### HasCurrentP75JobCost

`func (o *RunnerRecommendation) HasCurrentP75JobCost() bool`

HasCurrentP75JobCost returns a boolean if a field has been set.

### GetCurrentP75WallSeconds

`func (o *RunnerRecommendation) GetCurrentP75WallSeconds() float32`

GetCurrentP75WallSeconds returns the CurrentP75WallSeconds field if non-nil, zero value otherwise.

### GetCurrentP75WallSecondsOk

`func (o *RunnerRecommendation) GetCurrentP75WallSecondsOk() (*float32, bool)`

GetCurrentP75WallSecondsOk returns a tuple with the CurrentP75WallSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentP75WallSeconds

`func (o *RunnerRecommendation) SetCurrentP75WallSeconds(v float32)`

SetCurrentP75WallSeconds sets CurrentP75WallSeconds field to given value.

### HasCurrentP75WallSeconds

`func (o *RunnerRecommendation) HasCurrentP75WallSeconds() bool`

HasCurrentP75WallSeconds returns a boolean if a field has been set.

### GetCurrentWallSeconds

`func (o *RunnerRecommendation) GetCurrentWallSeconds() float32`

GetCurrentWallSeconds returns the CurrentWallSeconds field if non-nil, zero value otherwise.

### GetCurrentWallSecondsOk

`func (o *RunnerRecommendation) GetCurrentWallSecondsOk() (*float32, bool)`

GetCurrentWallSecondsOk returns a tuple with the CurrentWallSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentWallSeconds

`func (o *RunnerRecommendation) SetCurrentWallSeconds(v float32)`

SetCurrentWallSeconds sets CurrentWallSeconds field to given value.

### HasCurrentWallSeconds

`func (o *RunnerRecommendation) HasCurrentWallSeconds() bool`

HasCurrentWallSeconds returns a boolean if a field has been set.

### GetDocker

`func (o *RunnerRecommendation) GetDocker() bool`

GetDocker returns the Docker field if non-nil, zero value otherwise.

### GetDockerOk

`func (o *RunnerRecommendation) GetDockerOk() (*bool, bool)`

GetDockerOk returns a tuple with the Docker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocker

`func (o *RunnerRecommendation) SetDocker(v bool)`

SetDocker sets Docker field to given value.

### HasDocker

`func (o *RunnerRecommendation) HasDocker() bool`

HasDocker returns a boolean if a field has been set.

### GetFetchedOn

`func (o *RunnerRecommendation) GetFetchedOn() string`

GetFetchedOn returns the FetchedOn field if non-nil, zero value otherwise.

### GetFetchedOnOk

`func (o *RunnerRecommendation) GetFetchedOnOk() (*string, bool)`

GetFetchedOnOk returns a tuple with the FetchedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFetchedOn

`func (o *RunnerRecommendation) SetFetchedOn(v string)`

SetFetchedOn sets FetchedOn field to given value.

### HasFetchedOn

`func (o *RunnerRecommendation) HasFetchedOn() bool`

HasFetchedOn returns a boolean if a field has been set.

### GetFlags

`func (o *RunnerRecommendation) GetFlags() []string`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *RunnerRecommendation) GetFlagsOk() (*[]string, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *RunnerRecommendation) SetFlags(v []string)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *RunnerRecommendation) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### GetGpuCount

`func (o *RunnerRecommendation) GetGpuCount() int32`

GetGpuCount returns the GpuCount field if non-nil, zero value otherwise.

### GetGpuCountOk

`func (o *RunnerRecommendation) GetGpuCountOk() (*int32, bool)`

GetGpuCountOk returns a tuple with the GpuCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuCount

`func (o *RunnerRecommendation) SetGpuCount(v int32)`

SetGpuCount sets GpuCount field to given value.

### HasGpuCount

`func (o *RunnerRecommendation) HasGpuCount() bool`

HasGpuCount returns a boolean if a field has been set.

### GetIncludedApplied

`func (o *RunnerRecommendation) GetIncludedApplied() float32`

GetIncludedApplied returns the IncludedApplied field if non-nil, zero value otherwise.

### GetIncludedAppliedOk

`func (o *RunnerRecommendation) GetIncludedAppliedOk() (*float32, bool)`

GetIncludedAppliedOk returns a tuple with the IncludedApplied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedApplied

`func (o *RunnerRecommendation) SetIncludedApplied(v float32)`

SetIncludedApplied sets IncludedApplied field to given value.

### HasIncludedApplied

`func (o *RunnerRecommendation) HasIncludedApplied() bool`

HasIncludedApplied returns a boolean if a field has been set.

### GetNestedVirtualization

`func (o *RunnerRecommendation) GetNestedVirtualization() bool`

GetNestedVirtualization returns the NestedVirtualization field if non-nil, zero value otherwise.

### GetNestedVirtualizationOk

`func (o *RunnerRecommendation) GetNestedVirtualizationOk() (*bool, bool)`

GetNestedVirtualizationOk returns a tuple with the NestedVirtualization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNestedVirtualization

`func (o *RunnerRecommendation) SetNestedVirtualization(v bool)`

SetNestedVirtualization sets NestedVirtualization field to given value.

### HasNestedVirtualization

`func (o *RunnerRecommendation) HasNestedVirtualization() bool`

HasNestedVirtualization returns a boolean if a field has been set.

### GetOperatingSystem

`func (o *RunnerRecommendation) GetOperatingSystem() string`

GetOperatingSystem returns the OperatingSystem field if non-nil, zero value otherwise.

### GetOperatingSystemOk

`func (o *RunnerRecommendation) GetOperatingSystemOk() (*string, bool)`

GetOperatingSystemOk returns a tuple with the OperatingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystem

`func (o *RunnerRecommendation) SetOperatingSystem(v string)`

SetOperatingSystem sets OperatingSystem field to given value.

### HasOperatingSystem

`func (o *RunnerRecommendation) HasOperatingSystem() bool`

HasOperatingSystem returns a boolean if a field has been set.

### GetP75JobCost

`func (o *RunnerRecommendation) GetP75JobCost() float32`

GetP75JobCost returns the P75JobCost field if non-nil, zero value otherwise.

### GetP75JobCostOk

`func (o *RunnerRecommendation) GetP75JobCostOk() (*float32, bool)`

GetP75JobCostOk returns a tuple with the P75JobCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP75JobCost

`func (o *RunnerRecommendation) SetP75JobCost(v float32)`

SetP75JobCost sets P75JobCost field to given value.

### HasP75JobCost

`func (o *RunnerRecommendation) HasP75JobCost() bool`

HasP75JobCost returns a boolean if a field has been set.

### GetP75WallSeconds

`func (o *RunnerRecommendation) GetP75WallSeconds() float32`

GetP75WallSeconds returns the P75WallSeconds field if non-nil, zero value otherwise.

### GetP75WallSecondsOk

`func (o *RunnerRecommendation) GetP75WallSecondsOk() (*float32, bool)`

GetP75WallSecondsOk returns a tuple with the P75WallSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP75WallSeconds

`func (o *RunnerRecommendation) SetP75WallSeconds(v float32)`

SetP75WallSeconds sets P75WallSeconds field to given value.

### HasP75WallSeconds

`func (o *RunnerRecommendation) HasP75WallSeconds() bool`

HasP75WallSeconds returns a boolean if a field has been set.

### GetPlanFees

`func (o *RunnerRecommendation) GetPlanFees() []TypesRunnerPlanFee`

GetPlanFees returns the PlanFees field if non-nil, zero value otherwise.

### GetPlanFeesOk

`func (o *RunnerRecommendation) GetPlanFeesOk() (*[]TypesRunnerPlanFee, bool)`

GetPlanFeesOk returns a tuple with the PlanFees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanFees

`func (o *RunnerRecommendation) SetPlanFees(v []TypesRunnerPlanFee)`

SetPlanFees sets PlanFees field to given value.

### HasPlanFees

`func (o *RunnerRecommendation) HasPlanFees() bool`

HasPlanFees returns a boolean if a field has been set.

### GetPricingModel

`func (o *RunnerRecommendation) GetPricingModel() string`

GetPricingModel returns the PricingModel field if non-nil, zero value otherwise.

### GetPricingModelOk

`func (o *RunnerRecommendation) GetPricingModelOk() (*string, bool)`

GetPricingModelOk returns a tuple with the PricingModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricingModel

`func (o *RunnerRecommendation) SetPricingModel(v string)`

SetPricingModel sets PricingModel field to given value.

### HasPricingModel

`func (o *RunnerRecommendation) HasPricingModel() bool`

HasPricingModel returns a boolean if a field has been set.

### GetProvider

`func (o *RunnerRecommendation) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *RunnerRecommendation) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *RunnerRecommendation) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *RunnerRecommendation) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetRamGb

`func (o *RunnerRecommendation) GetRamGb() float32`

GetRamGb returns the RamGb field if non-nil, zero value otherwise.

### GetRamGbOk

`func (o *RunnerRecommendation) GetRamGbOk() (*float32, bool)`

GetRamGbOk returns a tuple with the RamGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRamGb

`func (o *RunnerRecommendation) SetRamGb(v float32)`

SetRamGb sets RamGb field to given value.

### HasRamGb

`func (o *RunnerRecommendation) HasRamGb() bool`

HasRamGb returns a boolean if a field has been set.

### GetRegion

`func (o *RunnerRecommendation) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *RunnerRecommendation) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *RunnerRecommendation) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *RunnerRecommendation) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetRuntimeSeconds

`func (o *RunnerRecommendation) GetRuntimeSeconds() float32`

GetRuntimeSeconds returns the RuntimeSeconds field if non-nil, zero value otherwise.

### GetRuntimeSecondsOk

`func (o *RunnerRecommendation) GetRuntimeSecondsOk() (*float32, bool)`

GetRuntimeSecondsOk returns a tuple with the RuntimeSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntimeSeconds

`func (o *RunnerRecommendation) SetRuntimeSeconds(v float32)`

SetRuntimeSeconds sets RuntimeSeconds field to given value.

### HasRuntimeSeconds

`func (o *RunnerRecommendation) HasRuntimeSeconds() bool`

HasRuntimeSeconds returns a boolean if a field has been set.

### GetSavingsVsCurrent

`func (o *RunnerRecommendation) GetSavingsVsCurrent() float32`

GetSavingsVsCurrent returns the SavingsVsCurrent field if non-nil, zero value otherwise.

### GetSavingsVsCurrentOk

`func (o *RunnerRecommendation) GetSavingsVsCurrentOk() (*float32, bool)`

GetSavingsVsCurrentOk returns a tuple with the SavingsVsCurrent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSavingsVsCurrent

`func (o *RunnerRecommendation) SetSavingsVsCurrent(v float32)`

SetSavingsVsCurrent sets SavingsVsCurrent field to given value.

### HasSavingsVsCurrent

`func (o *RunnerRecommendation) HasSavingsVsCurrent() bool`

HasSavingsVsCurrent returns a boolean if a field has been set.

### GetSku

`func (o *RunnerRecommendation) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *RunnerRecommendation) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *RunnerRecommendation) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *RunnerRecommendation) HasSku() bool`

HasSku returns a boolean if a field has been set.

### GetSourceUrl

`func (o *RunnerRecommendation) GetSourceUrl() string`

GetSourceUrl returns the SourceUrl field if non-nil, zero value otherwise.

### GetSourceUrlOk

`func (o *RunnerRecommendation) GetSourceUrlOk() (*string, bool)`

GetSourceUrlOk returns a tuple with the SourceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUrl

`func (o *RunnerRecommendation) SetSourceUrl(v string)`

SetSourceUrl sets SourceUrl field to given value.

### HasSourceUrl

`func (o *RunnerRecommendation) HasSourceUrl() bool`

HasSourceUrl returns a boolean if a field has been set.

### GetSpeedMeasured

`func (o *RunnerRecommendation) GetSpeedMeasured() bool`

GetSpeedMeasured returns the SpeedMeasured field if non-nil, zero value otherwise.

### GetSpeedMeasuredOk

`func (o *RunnerRecommendation) GetSpeedMeasuredOk() (*bool, bool)`

GetSpeedMeasuredOk returns a tuple with the SpeedMeasured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeedMeasured

`func (o *RunnerRecommendation) SetSpeedMeasured(v bool)`

SetSpeedMeasured sets SpeedMeasured field to given value.

### HasSpeedMeasured

`func (o *RunnerRecommendation) HasSpeedMeasured() bool`

HasSpeedMeasured returns a boolean if a field has been set.

### GetVmInstanceType

`func (o *RunnerRecommendation) GetVmInstanceType() string`

GetVmInstanceType returns the VmInstanceType field if non-nil, zero value otherwise.

### GetVmInstanceTypeOk

`func (o *RunnerRecommendation) GetVmInstanceTypeOk() (*string, bool)`

GetVmInstanceTypeOk returns a tuple with the VmInstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmInstanceType

`func (o *RunnerRecommendation) SetVmInstanceType(v string)`

SetVmInstanceType sets VmInstanceType field to given value.

### HasVmInstanceType

`func (o *RunnerRecommendation) HasVmInstanceType() bool`

HasVmInstanceType returns a boolean if a field has been set.

### GetWallSeconds

`func (o *RunnerRecommendation) GetWallSeconds() float32`

GetWallSeconds returns the WallSeconds field if non-nil, zero value otherwise.

### GetWallSecondsOk

`func (o *RunnerRecommendation) GetWallSecondsOk() (*float32, bool)`

GetWallSecondsOk returns a tuple with the WallSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWallSeconds

`func (o *RunnerRecommendation) SetWallSeconds(v float32)`

SetWallSeconds sets WallSeconds field to given value.

### HasWallSeconds

`func (o *RunnerRecommendation) HasWallSeconds() bool`

HasWallSeconds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# RunnerRecommendation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BilledSeconds** | Pointer to **float32** |  | [optional] 
**ComputeSource** | Pointer to **string** |  | [optional] 
**Cost** | Pointer to **float32** |  | [optional] 
**CostPerMonth** | Pointer to **float32** |  | [optional] 
**FetchedOn** | Pointer to **string** |  | [optional] 
**Flags** | Pointer to **[]string** |  | [optional] 
**PlanFees** | Pointer to [**[]TypesRunnerPlanFee**](TypesRunnerPlanFee.md) |  | [optional] 
**PricingModel** | Pointer to **string** |  | [optional] 
**Provider** | Pointer to **string** |  | [optional] 
**RuntimeSeconds** | Pointer to **float32** |  | [optional] 
**SavingsVsCurrent** | Pointer to **float32** |  | [optional] 
**Sku** | Pointer to **string** |  | [optional] 
**SourceUrl** | Pointer to **string** |  | [optional] 
**VmInstanceType** | Pointer to **string** |  | [optional] 

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



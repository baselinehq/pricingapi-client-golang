# ComputeRecommendation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CostPerWorkUnit** | Pointer to **float32** |  | [optional] 
**CpuClass** | Pointer to **string** |  | [optional] 
**Pricing** | Pointer to [**ComputePrice**](ComputePrice.md) |  | [optional] 
**Savings** | Pointer to [**TypesSavings**](TypesSavings.md) |  | [optional] 
**WorkUnit** | Pointer to **string** |  | [optional] 

## Methods

### NewComputeRecommendation

`func NewComputeRecommendation() *ComputeRecommendation`

NewComputeRecommendation instantiates a new ComputeRecommendation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputeRecommendationWithDefaults

`func NewComputeRecommendationWithDefaults() *ComputeRecommendation`

NewComputeRecommendationWithDefaults instantiates a new ComputeRecommendation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCostPerWorkUnit

`func (o *ComputeRecommendation) GetCostPerWorkUnit() float32`

GetCostPerWorkUnit returns the CostPerWorkUnit field if non-nil, zero value otherwise.

### GetCostPerWorkUnitOk

`func (o *ComputeRecommendation) GetCostPerWorkUnitOk() (*float32, bool)`

GetCostPerWorkUnitOk returns a tuple with the CostPerWorkUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostPerWorkUnit

`func (o *ComputeRecommendation) SetCostPerWorkUnit(v float32)`

SetCostPerWorkUnit sets CostPerWorkUnit field to given value.

### HasCostPerWorkUnit

`func (o *ComputeRecommendation) HasCostPerWorkUnit() bool`

HasCostPerWorkUnit returns a boolean if a field has been set.

### GetCpuClass

`func (o *ComputeRecommendation) GetCpuClass() string`

GetCpuClass returns the CpuClass field if non-nil, zero value otherwise.

### GetCpuClassOk

`func (o *ComputeRecommendation) GetCpuClassOk() (*string, bool)`

GetCpuClassOk returns a tuple with the CpuClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuClass

`func (o *ComputeRecommendation) SetCpuClass(v string)`

SetCpuClass sets CpuClass field to given value.

### HasCpuClass

`func (o *ComputeRecommendation) HasCpuClass() bool`

HasCpuClass returns a boolean if a field has been set.

### GetPricing

`func (o *ComputeRecommendation) GetPricing() ComputePrice`

GetPricing returns the Pricing field if non-nil, zero value otherwise.

### GetPricingOk

`func (o *ComputeRecommendation) GetPricingOk() (*ComputePrice, bool)`

GetPricingOk returns a tuple with the Pricing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricing

`func (o *ComputeRecommendation) SetPricing(v ComputePrice)`

SetPricing sets Pricing field to given value.

### HasPricing

`func (o *ComputeRecommendation) HasPricing() bool`

HasPricing returns a boolean if a field has been set.

### GetSavings

`func (o *ComputeRecommendation) GetSavings() TypesSavings`

GetSavings returns the Savings field if non-nil, zero value otherwise.

### GetSavingsOk

`func (o *ComputeRecommendation) GetSavingsOk() (*TypesSavings, bool)`

GetSavingsOk returns a tuple with the Savings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSavings

`func (o *ComputeRecommendation) SetSavings(v TypesSavings)`

SetSavings sets Savings field to given value.

### HasSavings

`func (o *ComputeRecommendation) HasSavings() bool`

HasSavings returns a boolean if a field has been set.

### GetWorkUnit

`func (o *ComputeRecommendation) GetWorkUnit() string`

GetWorkUnit returns the WorkUnit field if non-nil, zero value otherwise.

### GetWorkUnitOk

`func (o *ComputeRecommendation) GetWorkUnitOk() (*string, bool)`

GetWorkUnitOk returns a tuple with the WorkUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkUnit

`func (o *ComputeRecommendation) SetWorkUnit(v string)`

SetWorkUnit sets WorkUnit field to given value.

### HasWorkUnit

`func (o *ComputeRecommendation) HasWorkUnit() bool`

HasWorkUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



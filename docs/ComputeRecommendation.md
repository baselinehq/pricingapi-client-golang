# ComputeRecommendation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pricing** | Pointer to [**ComputePrice**](ComputePrice.md) |  | [optional] 
**Savings** | Pointer to [**TypesSavings**](TypesSavings.md) |  | [optional] 

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



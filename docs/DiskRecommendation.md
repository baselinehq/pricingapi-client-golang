# DiskRecommendation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pricing** | Pointer to [**DiskPrice**](DiskPrice.md) |  | [optional] 
**Savings** | Pointer to [**TypesSavings**](TypesSavings.md) |  | [optional] 

## Methods

### NewDiskRecommendation

`func NewDiskRecommendation() *DiskRecommendation`

NewDiskRecommendation instantiates a new DiskRecommendation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiskRecommendationWithDefaults

`func NewDiskRecommendationWithDefaults() *DiskRecommendation`

NewDiskRecommendationWithDefaults instantiates a new DiskRecommendation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPricing

`func (o *DiskRecommendation) GetPricing() DiskPrice`

GetPricing returns the Pricing field if non-nil, zero value otherwise.

### GetPricingOk

`func (o *DiskRecommendation) GetPricingOk() (*DiskPrice, bool)`

GetPricingOk returns a tuple with the Pricing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricing

`func (o *DiskRecommendation) SetPricing(v DiskPrice)`

SetPricing sets Pricing field to given value.

### HasPricing

`func (o *DiskRecommendation) HasPricing() bool`

HasPricing returns a boolean if a field has been set.

### GetSavings

`func (o *DiskRecommendation) GetSavings() TypesSavings`

GetSavings returns the Savings field if non-nil, zero value otherwise.

### GetSavingsOk

`func (o *DiskRecommendation) GetSavingsOk() (*TypesSavings, bool)`

GetSavingsOk returns a tuple with the Savings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSavings

`func (o *DiskRecommendation) SetSavings(v TypesSavings)`

SetSavings sets Savings field to given value.

### HasSavings

`func (o *DiskRecommendation) HasSavings() bool`

HasSavings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



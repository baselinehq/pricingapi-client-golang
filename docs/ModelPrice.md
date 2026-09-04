# ModelPrice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContextLimit** | Pointer to **int32** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**EffectiveFrom** | Pointer to **string** |  | [optional] 
**EffectiveTo** | Pointer to **string** |  | [optional] 
**FetchedAt** | Pointer to **string** |  | [optional] 
**Host** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**ModalitiesInput** | Pointer to **[]string** |  | [optional] 
**ModalitiesOutput** | Pointer to **[]string** |  | [optional] 
**Model** | Pointer to **string** |  | [optional] 
**ModelId** | Pointer to **string** |  | [optional] 
**OpenWeights** | Pointer to **bool** |  | [optional] 
**OutputLimit** | Pointer to **int32** |  | [optional] 
**Provider** | Pointer to **string** |  | [optional] 
**Publisher** | Pointer to **string** |  | [optional] 
**RawPricingData** | Pointer to **map[string]interface{}** |  | [optional] 
**Reasoning** | Pointer to **bool** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**SkuPriceId** | Pointer to **string** |  | [optional] 
**Source** | Pointer to **string** |  | [optional] 
**SourceUrl** | Pointer to **string** |  | [optional] 
**StructuredOutput** | Pointer to **bool** |  | [optional] 
**Tags** | Pointer to **map[string]interface{}** |  | [optional] 
**TokenBucket** | Pointer to **string** |  | [optional] 
**ToolCall** | Pointer to **bool** |  | [optional] 
**UnitPricePerMtok** | Pointer to **float32** |  | [optional] 

## Methods

### NewModelPrice

`func NewModelPrice() *ModelPrice`

NewModelPrice instantiates a new ModelPrice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelPriceWithDefaults

`func NewModelPriceWithDefaults() *ModelPrice`

NewModelPriceWithDefaults instantiates a new ModelPrice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContextLimit

`func (o *ModelPrice) GetContextLimit() int32`

GetContextLimit returns the ContextLimit field if non-nil, zero value otherwise.

### GetContextLimitOk

`func (o *ModelPrice) GetContextLimitOk() (*int32, bool)`

GetContextLimitOk returns a tuple with the ContextLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextLimit

`func (o *ModelPrice) SetContextLimit(v int32)`

SetContextLimit sets ContextLimit field to given value.

### HasContextLimit

`func (o *ModelPrice) HasContextLimit() bool`

HasContextLimit returns a boolean if a field has been set.

### GetCurrency

`func (o *ModelPrice) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ModelPrice) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *ModelPrice) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *ModelPrice) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetEffectiveFrom

`func (o *ModelPrice) GetEffectiveFrom() string`

GetEffectiveFrom returns the EffectiveFrom field if non-nil, zero value otherwise.

### GetEffectiveFromOk

`func (o *ModelPrice) GetEffectiveFromOk() (*string, bool)`

GetEffectiveFromOk returns a tuple with the EffectiveFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveFrom

`func (o *ModelPrice) SetEffectiveFrom(v string)`

SetEffectiveFrom sets EffectiveFrom field to given value.

### HasEffectiveFrom

`func (o *ModelPrice) HasEffectiveFrom() bool`

HasEffectiveFrom returns a boolean if a field has been set.

### GetEffectiveTo

`func (o *ModelPrice) GetEffectiveTo() string`

GetEffectiveTo returns the EffectiveTo field if non-nil, zero value otherwise.

### GetEffectiveToOk

`func (o *ModelPrice) GetEffectiveToOk() (*string, bool)`

GetEffectiveToOk returns a tuple with the EffectiveTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveTo

`func (o *ModelPrice) SetEffectiveTo(v string)`

SetEffectiveTo sets EffectiveTo field to given value.

### HasEffectiveTo

`func (o *ModelPrice) HasEffectiveTo() bool`

HasEffectiveTo returns a boolean if a field has been set.

### GetFetchedAt

`func (o *ModelPrice) GetFetchedAt() string`

GetFetchedAt returns the FetchedAt field if non-nil, zero value otherwise.

### GetFetchedAtOk

`func (o *ModelPrice) GetFetchedAtOk() (*string, bool)`

GetFetchedAtOk returns a tuple with the FetchedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFetchedAt

`func (o *ModelPrice) SetFetchedAt(v string)`

SetFetchedAt sets FetchedAt field to given value.

### HasFetchedAt

`func (o *ModelPrice) HasFetchedAt() bool`

HasFetchedAt returns a boolean if a field has been set.

### GetHost

`func (o *ModelPrice) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *ModelPrice) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *ModelPrice) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *ModelPrice) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetId

`func (o *ModelPrice) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelPrice) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelPrice) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ModelPrice) HasId() bool`

HasId returns a boolean if a field has been set.

### GetModalitiesInput

`func (o *ModelPrice) GetModalitiesInput() []string`

GetModalitiesInput returns the ModalitiesInput field if non-nil, zero value otherwise.

### GetModalitiesInputOk

`func (o *ModelPrice) GetModalitiesInputOk() (*[]string, bool)`

GetModalitiesInputOk returns a tuple with the ModalitiesInput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModalitiesInput

`func (o *ModelPrice) SetModalitiesInput(v []string)`

SetModalitiesInput sets ModalitiesInput field to given value.

### HasModalitiesInput

`func (o *ModelPrice) HasModalitiesInput() bool`

HasModalitiesInput returns a boolean if a field has been set.

### GetModalitiesOutput

`func (o *ModelPrice) GetModalitiesOutput() []string`

GetModalitiesOutput returns the ModalitiesOutput field if non-nil, zero value otherwise.

### GetModalitiesOutputOk

`func (o *ModelPrice) GetModalitiesOutputOk() (*[]string, bool)`

GetModalitiesOutputOk returns a tuple with the ModalitiesOutput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModalitiesOutput

`func (o *ModelPrice) SetModalitiesOutput(v []string)`

SetModalitiesOutput sets ModalitiesOutput field to given value.

### HasModalitiesOutput

`func (o *ModelPrice) HasModalitiesOutput() bool`

HasModalitiesOutput returns a boolean if a field has been set.

### GetModel

`func (o *ModelPrice) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ModelPrice) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ModelPrice) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *ModelPrice) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetModelId

`func (o *ModelPrice) GetModelId() string`

GetModelId returns the ModelId field if non-nil, zero value otherwise.

### GetModelIdOk

`func (o *ModelPrice) GetModelIdOk() (*string, bool)`

GetModelIdOk returns a tuple with the ModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelId

`func (o *ModelPrice) SetModelId(v string)`

SetModelId sets ModelId field to given value.

### HasModelId

`func (o *ModelPrice) HasModelId() bool`

HasModelId returns a boolean if a field has been set.

### GetOpenWeights

`func (o *ModelPrice) GetOpenWeights() bool`

GetOpenWeights returns the OpenWeights field if non-nil, zero value otherwise.

### GetOpenWeightsOk

`func (o *ModelPrice) GetOpenWeightsOk() (*bool, bool)`

GetOpenWeightsOk returns a tuple with the OpenWeights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenWeights

`func (o *ModelPrice) SetOpenWeights(v bool)`

SetOpenWeights sets OpenWeights field to given value.

### HasOpenWeights

`func (o *ModelPrice) HasOpenWeights() bool`

HasOpenWeights returns a boolean if a field has been set.

### GetOutputLimit

`func (o *ModelPrice) GetOutputLimit() int32`

GetOutputLimit returns the OutputLimit field if non-nil, zero value otherwise.

### GetOutputLimitOk

`func (o *ModelPrice) GetOutputLimitOk() (*int32, bool)`

GetOutputLimitOk returns a tuple with the OutputLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputLimit

`func (o *ModelPrice) SetOutputLimit(v int32)`

SetOutputLimit sets OutputLimit field to given value.

### HasOutputLimit

`func (o *ModelPrice) HasOutputLimit() bool`

HasOutputLimit returns a boolean if a field has been set.

### GetProvider

`func (o *ModelPrice) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *ModelPrice) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *ModelPrice) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *ModelPrice) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetPublisher

`func (o *ModelPrice) GetPublisher() string`

GetPublisher returns the Publisher field if non-nil, zero value otherwise.

### GetPublisherOk

`func (o *ModelPrice) GetPublisherOk() (*string, bool)`

GetPublisherOk returns a tuple with the Publisher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublisher

`func (o *ModelPrice) SetPublisher(v string)`

SetPublisher sets Publisher field to given value.

### HasPublisher

`func (o *ModelPrice) HasPublisher() bool`

HasPublisher returns a boolean if a field has been set.

### GetRawPricingData

`func (o *ModelPrice) GetRawPricingData() map[string]interface{}`

GetRawPricingData returns the RawPricingData field if non-nil, zero value otherwise.

### GetRawPricingDataOk

`func (o *ModelPrice) GetRawPricingDataOk() (*map[string]interface{}, bool)`

GetRawPricingDataOk returns a tuple with the RawPricingData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawPricingData

`func (o *ModelPrice) SetRawPricingData(v map[string]interface{})`

SetRawPricingData sets RawPricingData field to given value.

### HasRawPricingData

`func (o *ModelPrice) HasRawPricingData() bool`

HasRawPricingData returns a boolean if a field has been set.

### GetReasoning

`func (o *ModelPrice) GetReasoning() bool`

GetReasoning returns the Reasoning field if non-nil, zero value otherwise.

### GetReasoningOk

`func (o *ModelPrice) GetReasoningOk() (*bool, bool)`

GetReasoningOk returns a tuple with the Reasoning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasoning

`func (o *ModelPrice) SetReasoning(v bool)`

SetReasoning sets Reasoning field to given value.

### HasReasoning

`func (o *ModelPrice) HasReasoning() bool`

HasReasoning returns a boolean if a field has been set.

### GetRegion

`func (o *ModelPrice) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *ModelPrice) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *ModelPrice) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *ModelPrice) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetSkuPriceId

`func (o *ModelPrice) GetSkuPriceId() string`

GetSkuPriceId returns the SkuPriceId field if non-nil, zero value otherwise.

### GetSkuPriceIdOk

`func (o *ModelPrice) GetSkuPriceIdOk() (*string, bool)`

GetSkuPriceIdOk returns a tuple with the SkuPriceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkuPriceId

`func (o *ModelPrice) SetSkuPriceId(v string)`

SetSkuPriceId sets SkuPriceId field to given value.

### HasSkuPriceId

`func (o *ModelPrice) HasSkuPriceId() bool`

HasSkuPriceId returns a boolean if a field has been set.

### GetSource

`func (o *ModelPrice) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ModelPrice) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ModelPrice) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *ModelPrice) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetSourceUrl

`func (o *ModelPrice) GetSourceUrl() string`

GetSourceUrl returns the SourceUrl field if non-nil, zero value otherwise.

### GetSourceUrlOk

`func (o *ModelPrice) GetSourceUrlOk() (*string, bool)`

GetSourceUrlOk returns a tuple with the SourceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUrl

`func (o *ModelPrice) SetSourceUrl(v string)`

SetSourceUrl sets SourceUrl field to given value.

### HasSourceUrl

`func (o *ModelPrice) HasSourceUrl() bool`

HasSourceUrl returns a boolean if a field has been set.

### GetStructuredOutput

`func (o *ModelPrice) GetStructuredOutput() bool`

GetStructuredOutput returns the StructuredOutput field if non-nil, zero value otherwise.

### GetStructuredOutputOk

`func (o *ModelPrice) GetStructuredOutputOk() (*bool, bool)`

GetStructuredOutputOk returns a tuple with the StructuredOutput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStructuredOutput

`func (o *ModelPrice) SetStructuredOutput(v bool)`

SetStructuredOutput sets StructuredOutput field to given value.

### HasStructuredOutput

`func (o *ModelPrice) HasStructuredOutput() bool`

HasStructuredOutput returns a boolean if a field has been set.

### GetTags

`func (o *ModelPrice) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ModelPrice) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ModelPrice) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *ModelPrice) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTokenBucket

`func (o *ModelPrice) GetTokenBucket() string`

GetTokenBucket returns the TokenBucket field if non-nil, zero value otherwise.

### GetTokenBucketOk

`func (o *ModelPrice) GetTokenBucketOk() (*string, bool)`

GetTokenBucketOk returns a tuple with the TokenBucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenBucket

`func (o *ModelPrice) SetTokenBucket(v string)`

SetTokenBucket sets TokenBucket field to given value.

### HasTokenBucket

`func (o *ModelPrice) HasTokenBucket() bool`

HasTokenBucket returns a boolean if a field has been set.

### GetToolCall

`func (o *ModelPrice) GetToolCall() bool`

GetToolCall returns the ToolCall field if non-nil, zero value otherwise.

### GetToolCallOk

`func (o *ModelPrice) GetToolCallOk() (*bool, bool)`

GetToolCallOk returns a tuple with the ToolCall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolCall

`func (o *ModelPrice) SetToolCall(v bool)`

SetToolCall sets ToolCall field to given value.

### HasToolCall

`func (o *ModelPrice) HasToolCall() bool`

HasToolCall returns a boolean if a field has been set.

### GetUnitPricePerMtok

`func (o *ModelPrice) GetUnitPricePerMtok() float32`

GetUnitPricePerMtok returns the UnitPricePerMtok field if non-nil, zero value otherwise.

### GetUnitPricePerMtokOk

`func (o *ModelPrice) GetUnitPricePerMtokOk() (*float32, bool)`

GetUnitPricePerMtokOk returns a tuple with the UnitPricePerMtok field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitPricePerMtok

`func (o *ModelPrice) SetUnitPricePerMtok(v float32)`

SetUnitPricePerMtok sets UnitPricePerMtok field to given value.

### HasUnitPricePerMtok

`func (o *ModelPrice) HasUnitPricePerMtok() bool`

HasUnitPricePerMtok returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



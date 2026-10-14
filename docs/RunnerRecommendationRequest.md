# RunnerRecommendationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BilledMinutes** | Pointer to **float32** |  | [optional] 
**Current** | Pointer to [**TypesRunnerCurrent**](TypesRunnerCurrent.md) |  | [optional] 
**JobCount** | Pointer to **int32** |  | [optional] 
**JobsSeconds** | Pointer to **[]float32** |  | [optional] 
**MonthlyMinutes** | Pointer to **float32** |  | [optional] 
**PublicRepo** | Pointer to **bool** |  | [optional] 
**SelfHosted** | Pointer to [**TypesRunnerSelfHosted**](TypesRunnerSelfHosted.md) |  | [optional] 
**Shape** | Pointer to [**TypesRunnerShape**](TypesRunnerShape.md) |  | [optional] 

## Methods

### NewRunnerRecommendationRequest

`func NewRunnerRecommendationRequest() *RunnerRecommendationRequest`

NewRunnerRecommendationRequest instantiates a new RunnerRecommendationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunnerRecommendationRequestWithDefaults

`func NewRunnerRecommendationRequestWithDefaults() *RunnerRecommendationRequest`

NewRunnerRecommendationRequestWithDefaults instantiates a new RunnerRecommendationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBilledMinutes

`func (o *RunnerRecommendationRequest) GetBilledMinutes() float32`

GetBilledMinutes returns the BilledMinutes field if non-nil, zero value otherwise.

### GetBilledMinutesOk

`func (o *RunnerRecommendationRequest) GetBilledMinutesOk() (*float32, bool)`

GetBilledMinutesOk returns a tuple with the BilledMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBilledMinutes

`func (o *RunnerRecommendationRequest) SetBilledMinutes(v float32)`

SetBilledMinutes sets BilledMinutes field to given value.

### HasBilledMinutes

`func (o *RunnerRecommendationRequest) HasBilledMinutes() bool`

HasBilledMinutes returns a boolean if a field has been set.

### GetCurrent

`func (o *RunnerRecommendationRequest) GetCurrent() TypesRunnerCurrent`

GetCurrent returns the Current field if non-nil, zero value otherwise.

### GetCurrentOk

`func (o *RunnerRecommendationRequest) GetCurrentOk() (*TypesRunnerCurrent, bool)`

GetCurrentOk returns a tuple with the Current field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrent

`func (o *RunnerRecommendationRequest) SetCurrent(v TypesRunnerCurrent)`

SetCurrent sets Current field to given value.

### HasCurrent

`func (o *RunnerRecommendationRequest) HasCurrent() bool`

HasCurrent returns a boolean if a field has been set.

### GetJobCount

`func (o *RunnerRecommendationRequest) GetJobCount() int32`

GetJobCount returns the JobCount field if non-nil, zero value otherwise.

### GetJobCountOk

`func (o *RunnerRecommendationRequest) GetJobCountOk() (*int32, bool)`

GetJobCountOk returns a tuple with the JobCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobCount

`func (o *RunnerRecommendationRequest) SetJobCount(v int32)`

SetJobCount sets JobCount field to given value.

### HasJobCount

`func (o *RunnerRecommendationRequest) HasJobCount() bool`

HasJobCount returns a boolean if a field has been set.

### GetJobsSeconds

`func (o *RunnerRecommendationRequest) GetJobsSeconds() []float32`

GetJobsSeconds returns the JobsSeconds field if non-nil, zero value otherwise.

### GetJobsSecondsOk

`func (o *RunnerRecommendationRequest) GetJobsSecondsOk() (*[]float32, bool)`

GetJobsSecondsOk returns a tuple with the JobsSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobsSeconds

`func (o *RunnerRecommendationRequest) SetJobsSeconds(v []float32)`

SetJobsSeconds sets JobsSeconds field to given value.

### HasJobsSeconds

`func (o *RunnerRecommendationRequest) HasJobsSeconds() bool`

HasJobsSeconds returns a boolean if a field has been set.

### GetMonthlyMinutes

`func (o *RunnerRecommendationRequest) GetMonthlyMinutes() float32`

GetMonthlyMinutes returns the MonthlyMinutes field if non-nil, zero value otherwise.

### GetMonthlyMinutesOk

`func (o *RunnerRecommendationRequest) GetMonthlyMinutesOk() (*float32, bool)`

GetMonthlyMinutesOk returns a tuple with the MonthlyMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthlyMinutes

`func (o *RunnerRecommendationRequest) SetMonthlyMinutes(v float32)`

SetMonthlyMinutes sets MonthlyMinutes field to given value.

### HasMonthlyMinutes

`func (o *RunnerRecommendationRequest) HasMonthlyMinutes() bool`

HasMonthlyMinutes returns a boolean if a field has been set.

### GetPublicRepo

`func (o *RunnerRecommendationRequest) GetPublicRepo() bool`

GetPublicRepo returns the PublicRepo field if non-nil, zero value otherwise.

### GetPublicRepoOk

`func (o *RunnerRecommendationRequest) GetPublicRepoOk() (*bool, bool)`

GetPublicRepoOk returns a tuple with the PublicRepo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicRepo

`func (o *RunnerRecommendationRequest) SetPublicRepo(v bool)`

SetPublicRepo sets PublicRepo field to given value.

### HasPublicRepo

`func (o *RunnerRecommendationRequest) HasPublicRepo() bool`

HasPublicRepo returns a boolean if a field has been set.

### GetSelfHosted

`func (o *RunnerRecommendationRequest) GetSelfHosted() TypesRunnerSelfHosted`

GetSelfHosted returns the SelfHosted field if non-nil, zero value otherwise.

### GetSelfHostedOk

`func (o *RunnerRecommendationRequest) GetSelfHostedOk() (*TypesRunnerSelfHosted, bool)`

GetSelfHostedOk returns a tuple with the SelfHosted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfHosted

`func (o *RunnerRecommendationRequest) SetSelfHosted(v TypesRunnerSelfHosted)`

SetSelfHosted sets SelfHosted field to given value.

### HasSelfHosted

`func (o *RunnerRecommendationRequest) HasSelfHosted() bool`

HasSelfHosted returns a boolean if a field has been set.

### GetShape

`func (o *RunnerRecommendationRequest) GetShape() TypesRunnerShape`

GetShape returns the Shape field if non-nil, zero value otherwise.

### GetShapeOk

`func (o *RunnerRecommendationRequest) GetShapeOk() (*TypesRunnerShape, bool)`

GetShapeOk returns a tuple with the Shape field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShape

`func (o *RunnerRecommendationRequest) SetShape(v TypesRunnerShape)`

SetShape sets Shape field to given value.

### HasShape

`func (o *RunnerRecommendationRequest) HasShape() bool`

HasShape returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



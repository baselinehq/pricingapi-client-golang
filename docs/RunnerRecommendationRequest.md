# RunnerRecommendationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Clouds** | Pointer to [**[]TypesRunnerSelfHosted**](TypesRunnerSelfHosted.md) |  | [optional] 
**Connected** | Pointer to **[]string** |  | [optional] 
**Current** | Pointer to [**TypesCIQuoteCurrent**](TypesCIQuoteCurrent.md) |  | [optional] 
**GithubSelfHostedFee** | Pointer to **bool** |  | [optional] 
**Included** | Pointer to [**map[string]TypesCIIncluded**](TypesCIIncluded.md) |  | [optional] 
**Jobs** | Pointer to [**[]TypesCIJobBucket**](TypesCIJobBucket.md) |  | [optional] 
**MonthlyMinutes** | Pointer to **float32** |  | [optional] 
**PublicRepo** | Pointer to **bool** |  | [optional] 
**Shape** | Pointer to [**TypesCIRunnerShape**](TypesCIRunnerShape.md) |  | [optional] 
**Speed** | Pointer to [**map[string]TypesCISpeed**](TypesCISpeed.md) |  | [optional] 

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

### GetClouds

`func (o *RunnerRecommendationRequest) GetClouds() []TypesRunnerSelfHosted`

GetClouds returns the Clouds field if non-nil, zero value otherwise.

### GetCloudsOk

`func (o *RunnerRecommendationRequest) GetCloudsOk() (*[]TypesRunnerSelfHosted, bool)`

GetCloudsOk returns a tuple with the Clouds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClouds

`func (o *RunnerRecommendationRequest) SetClouds(v []TypesRunnerSelfHosted)`

SetClouds sets Clouds field to given value.

### HasClouds

`func (o *RunnerRecommendationRequest) HasClouds() bool`

HasClouds returns a boolean if a field has been set.

### GetConnected

`func (o *RunnerRecommendationRequest) GetConnected() []string`

GetConnected returns the Connected field if non-nil, zero value otherwise.

### GetConnectedOk

`func (o *RunnerRecommendationRequest) GetConnectedOk() (*[]string, bool)`

GetConnectedOk returns a tuple with the Connected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnected

`func (o *RunnerRecommendationRequest) SetConnected(v []string)`

SetConnected sets Connected field to given value.

### HasConnected

`func (o *RunnerRecommendationRequest) HasConnected() bool`

HasConnected returns a boolean if a field has been set.

### GetCurrent

`func (o *RunnerRecommendationRequest) GetCurrent() TypesCIQuoteCurrent`

GetCurrent returns the Current field if non-nil, zero value otherwise.

### GetCurrentOk

`func (o *RunnerRecommendationRequest) GetCurrentOk() (*TypesCIQuoteCurrent, bool)`

GetCurrentOk returns a tuple with the Current field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrent

`func (o *RunnerRecommendationRequest) SetCurrent(v TypesCIQuoteCurrent)`

SetCurrent sets Current field to given value.

### HasCurrent

`func (o *RunnerRecommendationRequest) HasCurrent() bool`

HasCurrent returns a boolean if a field has been set.

### GetGithubSelfHostedFee

`func (o *RunnerRecommendationRequest) GetGithubSelfHostedFee() bool`

GetGithubSelfHostedFee returns the GithubSelfHostedFee field if non-nil, zero value otherwise.

### GetGithubSelfHostedFeeOk

`func (o *RunnerRecommendationRequest) GetGithubSelfHostedFeeOk() (*bool, bool)`

GetGithubSelfHostedFeeOk returns a tuple with the GithubSelfHostedFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGithubSelfHostedFee

`func (o *RunnerRecommendationRequest) SetGithubSelfHostedFee(v bool)`

SetGithubSelfHostedFee sets GithubSelfHostedFee field to given value.

### HasGithubSelfHostedFee

`func (o *RunnerRecommendationRequest) HasGithubSelfHostedFee() bool`

HasGithubSelfHostedFee returns a boolean if a field has been set.

### GetIncluded

`func (o *RunnerRecommendationRequest) GetIncluded() map[string]TypesCIIncluded`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *RunnerRecommendationRequest) GetIncludedOk() (*map[string]TypesCIIncluded, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *RunnerRecommendationRequest) SetIncluded(v map[string]TypesCIIncluded)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *RunnerRecommendationRequest) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetJobs

`func (o *RunnerRecommendationRequest) GetJobs() []TypesCIJobBucket`

GetJobs returns the Jobs field if non-nil, zero value otherwise.

### GetJobsOk

`func (o *RunnerRecommendationRequest) GetJobsOk() (*[]TypesCIJobBucket, bool)`

GetJobsOk returns a tuple with the Jobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobs

`func (o *RunnerRecommendationRequest) SetJobs(v []TypesCIJobBucket)`

SetJobs sets Jobs field to given value.

### HasJobs

`func (o *RunnerRecommendationRequest) HasJobs() bool`

HasJobs returns a boolean if a field has been set.

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

### GetShape

`func (o *RunnerRecommendationRequest) GetShape() TypesCIRunnerShape`

GetShape returns the Shape field if non-nil, zero value otherwise.

### GetShapeOk

`func (o *RunnerRecommendationRequest) GetShapeOk() (*TypesCIRunnerShape, bool)`

GetShapeOk returns a tuple with the Shape field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShape

`func (o *RunnerRecommendationRequest) SetShape(v TypesCIRunnerShape)`

SetShape sets Shape field to given value.

### HasShape

`func (o *RunnerRecommendationRequest) HasShape() bool`

HasShape returns a boolean if a field has been set.

### GetSpeed

`func (o *RunnerRecommendationRequest) GetSpeed() map[string]TypesCISpeed`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *RunnerRecommendationRequest) GetSpeedOk() (*map[string]TypesCISpeed, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *RunnerRecommendationRequest) SetSpeed(v map[string]TypesCISpeed)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *RunnerRecommendationRequest) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



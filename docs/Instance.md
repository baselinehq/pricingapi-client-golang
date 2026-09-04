# Instance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvailabilityZone** | Pointer to **string** |  | [optional] 
**FallbackToBasePricing** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**InstanceType** | Pointer to **string** |  | [optional] 
**OperatingSystem** | Pointer to **string** |  | [optional] 
**Provider** | Pointer to [**GithubComBaselinehqGolangSharedTypesProvider**](GithubComBaselinehqGolangSharedTypesProvider.md) |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**Service** | Pointer to [**TypesService**](TypesService.md) |  | [optional] 
**UsageType** | Pointer to [**GithubComBaselinehqGolangSharedTypesUsageType**](GithubComBaselinehqGolangSharedTypesUsageType.md) |  | [optional] 
**Vm** | Pointer to [**GithubComBaselinehqGolangSharedTypesVM**](GithubComBaselinehqGolangSharedTypesVM.md) |  | [optional] 

## Methods

### NewInstance

`func NewInstance() *Instance`

NewInstance instantiates a new Instance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceWithDefaults

`func NewInstanceWithDefaults() *Instance`

NewInstanceWithDefaults instantiates a new Instance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailabilityZone

`func (o *Instance) GetAvailabilityZone() string`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *Instance) GetAvailabilityZoneOk() (*string, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *Instance) SetAvailabilityZone(v string)`

SetAvailabilityZone sets AvailabilityZone field to given value.

### HasAvailabilityZone

`func (o *Instance) HasAvailabilityZone() bool`

HasAvailabilityZone returns a boolean if a field has been set.

### GetFallbackToBasePricing

`func (o *Instance) GetFallbackToBasePricing() bool`

GetFallbackToBasePricing returns the FallbackToBasePricing field if non-nil, zero value otherwise.

### GetFallbackToBasePricingOk

`func (o *Instance) GetFallbackToBasePricingOk() (*bool, bool)`

GetFallbackToBasePricingOk returns a tuple with the FallbackToBasePricing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackToBasePricing

`func (o *Instance) SetFallbackToBasePricing(v bool)`

SetFallbackToBasePricing sets FallbackToBasePricing field to given value.

### HasFallbackToBasePricing

`func (o *Instance) HasFallbackToBasePricing() bool`

HasFallbackToBasePricing returns a boolean if a field has been set.

### GetId

`func (o *Instance) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Instance) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Instance) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Instance) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInstanceType

`func (o *Instance) GetInstanceType() string`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *Instance) GetInstanceTypeOk() (*string, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *Instance) SetInstanceType(v string)`

SetInstanceType sets InstanceType field to given value.

### HasInstanceType

`func (o *Instance) HasInstanceType() bool`

HasInstanceType returns a boolean if a field has been set.

### GetOperatingSystem

`func (o *Instance) GetOperatingSystem() string`

GetOperatingSystem returns the OperatingSystem field if non-nil, zero value otherwise.

### GetOperatingSystemOk

`func (o *Instance) GetOperatingSystemOk() (*string, bool)`

GetOperatingSystemOk returns a tuple with the OperatingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystem

`func (o *Instance) SetOperatingSystem(v string)`

SetOperatingSystem sets OperatingSystem field to given value.

### HasOperatingSystem

`func (o *Instance) HasOperatingSystem() bool`

HasOperatingSystem returns a boolean if a field has been set.

### GetProvider

`func (o *Instance) GetProvider() GithubComBaselinehqGolangSharedTypesProvider`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *Instance) GetProviderOk() (*GithubComBaselinehqGolangSharedTypesProvider, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *Instance) SetProvider(v GithubComBaselinehqGolangSharedTypesProvider)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *Instance) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetRegion

`func (o *Instance) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *Instance) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *Instance) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *Instance) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetService

`func (o *Instance) GetService() TypesService`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *Instance) GetServiceOk() (*TypesService, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *Instance) SetService(v TypesService)`

SetService sets Service field to given value.

### HasService

`func (o *Instance) HasService() bool`

HasService returns a boolean if a field has been set.

### GetUsageType

`func (o *Instance) GetUsageType() GithubComBaselinehqGolangSharedTypesUsageType`

GetUsageType returns the UsageType field if non-nil, zero value otherwise.

### GetUsageTypeOk

`func (o *Instance) GetUsageTypeOk() (*GithubComBaselinehqGolangSharedTypesUsageType, bool)`

GetUsageTypeOk returns a tuple with the UsageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageType

`func (o *Instance) SetUsageType(v GithubComBaselinehqGolangSharedTypesUsageType)`

SetUsageType sets UsageType field to given value.

### HasUsageType

`func (o *Instance) HasUsageType() bool`

HasUsageType returns a boolean if a field has been set.

### GetVm

`func (o *Instance) GetVm() GithubComBaselinehqGolangSharedTypesVM`

GetVm returns the Vm field if non-nil, zero value otherwise.

### GetVmOk

`func (o *Instance) GetVmOk() (*GithubComBaselinehqGolangSharedTypesVM, bool)`

GetVmOk returns a tuple with the Vm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVm

`func (o *Instance) SetVm(v GithubComBaselinehqGolangSharedTypesVM)`

SetVm sets Vm field to given value.

### HasVm

`func (o *Instance) HasVm() bool`

HasVm returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



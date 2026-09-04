# Disk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvailabilityZone** | Pointer to **string** |  | [optional] 
**CapacityGb** | Pointer to **float32** |  | [optional] 
**FallbackToBasePricing** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Iops** | Pointer to **float32** |  | [optional] 
**Provider** | Pointer to [**GithubComBaselinehqGolangSharedTypesProvider**](GithubComBaselinehqGolangSharedTypesProvider.md) |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**Service** | Pointer to [**TypesService**](TypesService.md) |  | [optional] 
**ThroughputMbps** | Pointer to **float32** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**UsageType** | Pointer to [**GithubComBaselinehqGolangSharedTypesUsageType**](GithubComBaselinehqGolangSharedTypesUsageType.md) |  | [optional] 

## Methods

### NewDisk

`func NewDisk() *Disk`

NewDisk instantiates a new Disk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiskWithDefaults

`func NewDiskWithDefaults() *Disk`

NewDiskWithDefaults instantiates a new Disk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailabilityZone

`func (o *Disk) GetAvailabilityZone() string`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *Disk) GetAvailabilityZoneOk() (*string, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *Disk) SetAvailabilityZone(v string)`

SetAvailabilityZone sets AvailabilityZone field to given value.

### HasAvailabilityZone

`func (o *Disk) HasAvailabilityZone() bool`

HasAvailabilityZone returns a boolean if a field has been set.

### GetCapacityGb

`func (o *Disk) GetCapacityGb() float32`

GetCapacityGb returns the CapacityGb field if non-nil, zero value otherwise.

### GetCapacityGbOk

`func (o *Disk) GetCapacityGbOk() (*float32, bool)`

GetCapacityGbOk returns a tuple with the CapacityGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacityGb

`func (o *Disk) SetCapacityGb(v float32)`

SetCapacityGb sets CapacityGb field to given value.

### HasCapacityGb

`func (o *Disk) HasCapacityGb() bool`

HasCapacityGb returns a boolean if a field has been set.

### GetFallbackToBasePricing

`func (o *Disk) GetFallbackToBasePricing() bool`

GetFallbackToBasePricing returns the FallbackToBasePricing field if non-nil, zero value otherwise.

### GetFallbackToBasePricingOk

`func (o *Disk) GetFallbackToBasePricingOk() (*bool, bool)`

GetFallbackToBasePricingOk returns a tuple with the FallbackToBasePricing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackToBasePricing

`func (o *Disk) SetFallbackToBasePricing(v bool)`

SetFallbackToBasePricing sets FallbackToBasePricing field to given value.

### HasFallbackToBasePricing

`func (o *Disk) HasFallbackToBasePricing() bool`

HasFallbackToBasePricing returns a boolean if a field has been set.

### GetId

`func (o *Disk) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Disk) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Disk) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Disk) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIops

`func (o *Disk) GetIops() float32`

GetIops returns the Iops field if non-nil, zero value otherwise.

### GetIopsOk

`func (o *Disk) GetIopsOk() (*float32, bool)`

GetIopsOk returns a tuple with the Iops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIops

`func (o *Disk) SetIops(v float32)`

SetIops sets Iops field to given value.

### HasIops

`func (o *Disk) HasIops() bool`

HasIops returns a boolean if a field has been set.

### GetProvider

`func (o *Disk) GetProvider() GithubComBaselinehqGolangSharedTypesProvider`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *Disk) GetProviderOk() (*GithubComBaselinehqGolangSharedTypesProvider, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *Disk) SetProvider(v GithubComBaselinehqGolangSharedTypesProvider)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *Disk) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetRegion

`func (o *Disk) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *Disk) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *Disk) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *Disk) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetService

`func (o *Disk) GetService() TypesService`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *Disk) GetServiceOk() (*TypesService, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *Disk) SetService(v TypesService)`

SetService sets Service field to given value.

### HasService

`func (o *Disk) HasService() bool`

HasService returns a boolean if a field has been set.

### GetThroughputMbps

`func (o *Disk) GetThroughputMbps() float32`

GetThroughputMbps returns the ThroughputMbps field if non-nil, zero value otherwise.

### GetThroughputMbpsOk

`func (o *Disk) GetThroughputMbpsOk() (*float32, bool)`

GetThroughputMbpsOk returns a tuple with the ThroughputMbps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThroughputMbps

`func (o *Disk) SetThroughputMbps(v float32)`

SetThroughputMbps sets ThroughputMbps field to given value.

### HasThroughputMbps

`func (o *Disk) HasThroughputMbps() bool`

HasThroughputMbps returns a boolean if a field has been set.

### GetType

`func (o *Disk) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Disk) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Disk) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Disk) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUsageType

`func (o *Disk) GetUsageType() GithubComBaselinehqGolangSharedTypesUsageType`

GetUsageType returns the UsageType field if non-nil, zero value otherwise.

### GetUsageTypeOk

`func (o *Disk) GetUsageTypeOk() (*GithubComBaselinehqGolangSharedTypesUsageType, bool)`

GetUsageTypeOk returns a tuple with the UsageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageType

`func (o *Disk) SetUsageType(v GithubComBaselinehqGolangSharedTypesUsageType)`

SetUsageType sets UsageType field to given value.

### HasUsageType

`func (o *Disk) HasUsageType() bool`

HasUsageType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



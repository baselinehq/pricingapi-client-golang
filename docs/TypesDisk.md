# TypesDisk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvailabilityZone** | Pointer to **string** |  | [optional] 
**CapacityGb** | Pointer to **float32** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Iops** | Pointer to **float32** |  | [optional] 
**Provider** | Pointer to [**GithubComBaselinehqGolangSharedTypesProvider**](GithubComBaselinehqGolangSharedTypesProvider.md) |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**Service** | Pointer to [**GithubComBaselinehqGolangSharedTypesService**](GithubComBaselinehqGolangSharedTypesService.md) |  | [optional] 
**ThroughputMbps** | Pointer to **float32** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**UsageType** | Pointer to [**GithubComBaselinehqGolangSharedTypesUsageType**](GithubComBaselinehqGolangSharedTypesUsageType.md) |  | [optional] 

## Methods

### NewTypesDisk

`func NewTypesDisk() *TypesDisk`

NewTypesDisk instantiates a new TypesDisk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTypesDiskWithDefaults

`func NewTypesDiskWithDefaults() *TypesDisk`

NewTypesDiskWithDefaults instantiates a new TypesDisk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailabilityZone

`func (o *TypesDisk) GetAvailabilityZone() string`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *TypesDisk) GetAvailabilityZoneOk() (*string, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *TypesDisk) SetAvailabilityZone(v string)`

SetAvailabilityZone sets AvailabilityZone field to given value.

### HasAvailabilityZone

`func (o *TypesDisk) HasAvailabilityZone() bool`

HasAvailabilityZone returns a boolean if a field has been set.

### GetCapacityGb

`func (o *TypesDisk) GetCapacityGb() float32`

GetCapacityGb returns the CapacityGb field if non-nil, zero value otherwise.

### GetCapacityGbOk

`func (o *TypesDisk) GetCapacityGbOk() (*float32, bool)`

GetCapacityGbOk returns a tuple with the CapacityGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacityGb

`func (o *TypesDisk) SetCapacityGb(v float32)`

SetCapacityGb sets CapacityGb field to given value.

### HasCapacityGb

`func (o *TypesDisk) HasCapacityGb() bool`

HasCapacityGb returns a boolean if a field has been set.

### GetId

`func (o *TypesDisk) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TypesDisk) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TypesDisk) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TypesDisk) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIops

`func (o *TypesDisk) GetIops() float32`

GetIops returns the Iops field if non-nil, zero value otherwise.

### GetIopsOk

`func (o *TypesDisk) GetIopsOk() (*float32, bool)`

GetIopsOk returns a tuple with the Iops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIops

`func (o *TypesDisk) SetIops(v float32)`

SetIops sets Iops field to given value.

### HasIops

`func (o *TypesDisk) HasIops() bool`

HasIops returns a boolean if a field has been set.

### GetProvider

`func (o *TypesDisk) GetProvider() GithubComBaselinehqGolangSharedTypesProvider`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *TypesDisk) GetProviderOk() (*GithubComBaselinehqGolangSharedTypesProvider, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *TypesDisk) SetProvider(v GithubComBaselinehqGolangSharedTypesProvider)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *TypesDisk) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetRegion

`func (o *TypesDisk) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *TypesDisk) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *TypesDisk) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *TypesDisk) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetService

`func (o *TypesDisk) GetService() GithubComBaselinehqGolangSharedTypesService`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *TypesDisk) GetServiceOk() (*GithubComBaselinehqGolangSharedTypesService, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *TypesDisk) SetService(v GithubComBaselinehqGolangSharedTypesService)`

SetService sets Service field to given value.

### HasService

`func (o *TypesDisk) HasService() bool`

HasService returns a boolean if a field has been set.

### GetThroughputMbps

`func (o *TypesDisk) GetThroughputMbps() float32`

GetThroughputMbps returns the ThroughputMbps field if non-nil, zero value otherwise.

### GetThroughputMbpsOk

`func (o *TypesDisk) GetThroughputMbpsOk() (*float32, bool)`

GetThroughputMbpsOk returns a tuple with the ThroughputMbps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThroughputMbps

`func (o *TypesDisk) SetThroughputMbps(v float32)`

SetThroughputMbps sets ThroughputMbps field to given value.

### HasThroughputMbps

`func (o *TypesDisk) HasThroughputMbps() bool`

HasThroughputMbps returns a boolean if a field has been set.

### GetType

`func (o *TypesDisk) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TypesDisk) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TypesDisk) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TypesDisk) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUsageType

`func (o *TypesDisk) GetUsageType() GithubComBaselinehqGolangSharedTypesUsageType`

GetUsageType returns the UsageType field if non-nil, zero value otherwise.

### GetUsageTypeOk

`func (o *TypesDisk) GetUsageTypeOk() (*GithubComBaselinehqGolangSharedTypesUsageType, bool)`

GetUsageTypeOk returns a tuple with the UsageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageType

`func (o *TypesDisk) SetUsageType(v GithubComBaselinehqGolangSharedTypesUsageType)`

SetUsageType sets UsageType field to given value.

### HasUsageType

`func (o *TypesDisk) HasUsageType() bool`

HasUsageType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



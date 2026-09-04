# \DefaultAPI

All URIs are relative to *https://pricing.baselinehq.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**HealthzGet**](DefaultAPI.md#HealthzGet) | **Get** /healthz | Health check endpoint
[**MarketplaceProvidersComputeDelete**](DefaultAPI.md#MarketplaceProvidersComputeDelete) | **Delete** /marketplace/providers/compute | Delete a custom provider instance
[**MarketplaceProvidersComputeGet**](DefaultAPI.md#MarketplaceProvidersComputeGet) | **Get** /marketplace/providers/compute | Get your custom pricing entries
[**MarketplaceProvidersComputePost**](DefaultAPI.md#MarketplaceProvidersComputePost) | **Post** /marketplace/providers/compute | Register a custom provider
[**MarketplaceProvidersDatabasesDelete**](DefaultAPI.md#MarketplaceProvidersDatabasesDelete) | **Delete** /marketplace/providers/databases | Delete a custom database pricing entry
[**MarketplaceProvidersDatabasesGet**](DefaultAPI.md#MarketplaceProvidersDatabasesGet) | **Get** /marketplace/providers/databases | Get your custom database pricing entries
[**MarketplaceProvidersDatabasesPost**](DefaultAPI.md#MarketplaceProvidersDatabasesPost) | **Post** /marketplace/providers/databases | Register custom database pricing
[**MarketplaceProvidersDisksDelete**](DefaultAPI.md#MarketplaceProvidersDisksDelete) | **Delete** /marketplace/providers/disks | Delete a custom disk provider entry
[**MarketplaceProvidersDisksGet**](DefaultAPI.md#MarketplaceProvidersDisksGet) | **Get** /marketplace/providers/disks | Get your custom disk pricing entries
[**MarketplaceProvidersDisksPost**](DefaultAPI.md#MarketplaceProvidersDisksPost) | **Post** /marketplace/providers/disks | Register a custom disk provider
[**MarketplaceProvidersModelsDelete**](DefaultAPI.md#MarketplaceProvidersModelsDelete) | **Delete** /marketplace/providers/models | Delete a custom model pricing entry
[**MarketplaceProvidersModelsGet**](DefaultAPI.md#MarketplaceProvidersModelsGet) | **Get** /marketplace/providers/models | Get your custom model pricing entries
[**MarketplaceProvidersModelsPost**](DefaultAPI.md#MarketplaceProvidersModelsPost) | **Post** /marketplace/providers/models | Register custom model pricing
[**PricingComputePost**](DefaultAPI.md#PricingComputePost) | **Post** /pricing/compute | Get pricing for an instance
[**PricingDisksPost**](DefaultAPI.md#PricingDisksPost) | **Post** /pricing/disks | Get pricing for a disk
[**PricingPost**](DefaultAPI.md#PricingPost) | **Post** /pricing | Get pricing for an instance
[**ProvidersGet**](DefaultAPI.md#ProvidersGet) | **Get** /providers | Get details for the providers
[**RecommendationsComputePost**](DefaultAPI.md#RecommendationsComputePost) | **Post** /recommendations/compute | Get recommendations for compute instances
[**RecommendationsDisksPost**](DefaultAPI.md#RecommendationsDisksPost) | **Post** /recommendations/disks | Get recommendations for disks
[**RecommendationsPost**](DefaultAPI.md#RecommendationsPost) | **Post** /recommendations | Get recommendations for compute instances



## HealthzGet

> map[string]string HealthzGet(ctx).Execute()

Health check endpoint



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/baselinehq/pricingapi-client-golang"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.HealthzGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.HealthzGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HealthzGet`: map[string]string
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.HealthzGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiHealthzGetRequest struct via the builder pattern


### Return type

**map[string]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarketplaceProvidersComputeDelete

> RegisteredComputePrices MarketplaceProvidersComputeDelete(ctx).Id(id).Execute()

Delete a custom provider instance



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/baselinehq/pricingapi-client-golang"
)

func main() {
	id := "id_example" // string | Instance ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.MarketplaceProvidersComputeDelete(context.Background()).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.MarketplaceProvidersComputeDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarketplaceProvidersComputeDelete`: RegisteredComputePrices
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.MarketplaceProvidersComputeDelete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarketplaceProvidersComputeDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | Instance ID | 

### Return type

[**RegisteredComputePrices**](RegisteredComputePrices.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarketplaceProvidersComputeGet

> ComputePriceList MarketplaceProvidersComputeGet(ctx).Limit(limit).Offset(offset).Service(service).Region(region).InstanceType(instanceType).UsageType(usageType).Execute()

Get your custom pricing entries



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/baselinehq/pricingapi-client-golang"
)

func main() {
	limit := int32(56) // int32 | Maximum entries to return (default 500) (optional)
	offset := int32(56) // int32 | Entries to skip (optional)
	service := "service_example" // string | Filter by service (optional)
	region := "region_example" // string | Filter by region (optional)
	instanceType := "instanceType_example" // string | Filter by instance type (optional)
	usageType := "usageType_example" // string | Filter by usage type (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.MarketplaceProvidersComputeGet(context.Background()).Limit(limit).Offset(offset).Service(service).Region(region).InstanceType(instanceType).UsageType(usageType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.MarketplaceProvidersComputeGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarketplaceProvidersComputeGet`: ComputePriceList
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.MarketplaceProvidersComputeGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarketplaceProvidersComputeGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Maximum entries to return (default 500) | 
 **offset** | **int32** | Entries to skip | 
 **service** | **string** | Filter by service | 
 **region** | **string** | Filter by region | 
 **instanceType** | **string** | Filter by instance type | 
 **usageType** | **string** | Filter by usage type | 

### Return type

[**ComputePriceList**](ComputePriceList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarketplaceProvidersComputePost

> RegisteredComputePrices MarketplaceProvidersComputePost(ctx).Instance(instance).Execute()

Register a custom provider



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/baselinehq/pricingapi-client-golang"
)

func main() {
	instance := *openapiclient.NewComputePricesRequest([]openapiclient.ComputePrice{*openapiclient.NewComputePrice()}) // ComputePricesRequest | Custom pricing request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.MarketplaceProvidersComputePost(context.Background()).Instance(instance).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.MarketplaceProvidersComputePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarketplaceProvidersComputePost`: RegisteredComputePrices
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.MarketplaceProvidersComputePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarketplaceProvidersComputePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instance** | [**ComputePricesRequest**](ComputePricesRequest.md) | Custom pricing request | 

### Return type

[**RegisteredComputePrices**](RegisteredComputePrices.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarketplaceProvidersDatabasesDelete

> RegisteredDatabasePrices MarketplaceProvidersDatabasesDelete(ctx).Id(id).Execute()

Delete a custom database pricing entry



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/baselinehq/pricingapi-client-golang"
)

func main() {
	id := "id_example" // string | Database pricing ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.MarketplaceProvidersDatabasesDelete(context.Background()).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.MarketplaceProvidersDatabasesDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarketplaceProvidersDatabasesDelete`: RegisteredDatabasePrices
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.MarketplaceProvidersDatabasesDelete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarketplaceProvidersDatabasesDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | Database pricing ID | 

### Return type

[**RegisteredDatabasePrices**](RegisteredDatabasePrices.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarketplaceProvidersDatabasesGet

> DatabasePriceList MarketplaceProvidersDatabasesGet(ctx).Limit(limit).Offset(offset).Service(service).Region(region).Engine(engine).Edition(edition).InstanceType(instanceType).UsageType(usageType).Execute()

Get your custom database pricing entries



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/baselinehq/pricingapi-client-golang"
)

func main() {
	limit := int32(56) // int32 | Maximum entries to return (default 500) (optional)
	offset := int32(56) // int32 | Entries to skip (optional)
	service := "service_example" // string | Filter by service (optional)
	region := "region_example" // string | Filter by region (optional)
	engine := "engine_example" // string | Filter by engine (optional)
	edition := "edition_example" // string | Filter by edition (optional)
	instanceType := "instanceType_example" // string | Filter by instance type (optional)
	usageType := "usageType_example" // string | Filter by usage type (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.MarketplaceProvidersDatabasesGet(context.Background()).Limit(limit).Offset(offset).Service(service).Region(region).Engine(engine).Edition(edition).InstanceType(instanceType).UsageType(usageType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.MarketplaceProvidersDatabasesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarketplaceProvidersDatabasesGet`: DatabasePriceList
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.MarketplaceProvidersDatabasesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarketplaceProvidersDatabasesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Maximum entries to return (default 500) | 
 **offset** | **int32** | Entries to skip | 
 **service** | **string** | Filter by service | 
 **region** | **string** | Filter by region | 
 **engine** | **string** | Filter by engine | 
 **edition** | **string** | Filter by edition | 
 **instanceType** | **string** | Filter by instance type | 
 **usageType** | **string** | Filter by usage type | 

### Return type

[**DatabasePriceList**](DatabasePriceList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarketplaceProvidersDatabasesPost

> RegisteredDatabasePrices MarketplaceProvidersDatabasesPost(ctx).Instance(instance).Execute()

Register custom database pricing



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/baselinehq/pricingapi-client-golang"
)

func main() {
	instance := *openapiclient.NewDatabasePricesRequest([]openapiclient.DatabasePrice{*openapiclient.NewDatabasePrice()}) // DatabasePricesRequest | Custom database pricing request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.MarketplaceProvidersDatabasesPost(context.Background()).Instance(instance).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.MarketplaceProvidersDatabasesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarketplaceProvidersDatabasesPost`: RegisteredDatabasePrices
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.MarketplaceProvidersDatabasesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarketplaceProvidersDatabasesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instance** | [**DatabasePricesRequest**](DatabasePricesRequest.md) | Custom database pricing request | 

### Return type

[**RegisteredDatabasePrices**](RegisteredDatabasePrices.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarketplaceProvidersDisksDelete

> RegisteredDiskPrices MarketplaceProvidersDisksDelete(ctx).Id(id).Execute()

Delete a custom disk provider entry



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/baselinehq/pricingapi-client-golang"
)

func main() {
	id := "id_example" // string | Disk pricing ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.MarketplaceProvidersDisksDelete(context.Background()).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.MarketplaceProvidersDisksDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarketplaceProvidersDisksDelete`: RegisteredDiskPrices
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.MarketplaceProvidersDisksDelete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarketplaceProvidersDisksDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | Disk pricing ID | 

### Return type

[**RegisteredDiskPrices**](RegisteredDiskPrices.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarketplaceProvidersDisksGet

> DiskPriceList MarketplaceProvidersDisksGet(ctx).Limit(limit).Offset(offset).Service(service).Region(region).Type_(type_).UsageType(usageType).Execute()

Get your custom disk pricing entries



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/baselinehq/pricingapi-client-golang"
)

func main() {
	limit := int32(56) // int32 | Maximum entries to return (default 500) (optional)
	offset := int32(56) // int32 | Entries to skip (optional)
	service := "service_example" // string | Filter by service (optional)
	region := "region_example" // string | Filter by region (optional)
	type_ := "type__example" // string | Filter by disk type (optional)
	usageType := "usageType_example" // string | Filter by usage type (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.MarketplaceProvidersDisksGet(context.Background()).Limit(limit).Offset(offset).Service(service).Region(region).Type_(type_).UsageType(usageType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.MarketplaceProvidersDisksGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarketplaceProvidersDisksGet`: DiskPriceList
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.MarketplaceProvidersDisksGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarketplaceProvidersDisksGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Maximum entries to return (default 500) | 
 **offset** | **int32** | Entries to skip | 
 **service** | **string** | Filter by service | 
 **region** | **string** | Filter by region | 
 **type_** | **string** | Filter by disk type | 
 **usageType** | **string** | Filter by usage type | 

### Return type

[**DiskPriceList**](DiskPriceList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarketplaceProvidersDisksPost

> RegisteredDiskPrices MarketplaceProvidersDisksPost(ctx).Instance(instance).Execute()

Register a custom disk provider



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/baselinehq/pricingapi-client-golang"
)

func main() {
	instance := *openapiclient.NewDiskPricesRequest([]openapiclient.DiskPrice{*openapiclient.NewDiskPrice()}) // DiskPricesRequest | Custom disk pricing request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.MarketplaceProvidersDisksPost(context.Background()).Instance(instance).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.MarketplaceProvidersDisksPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarketplaceProvidersDisksPost`: RegisteredDiskPrices
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.MarketplaceProvidersDisksPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarketplaceProvidersDisksPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instance** | [**DiskPricesRequest**](DiskPricesRequest.md) | Custom disk pricing request | 

### Return type

[**RegisteredDiskPrices**](RegisteredDiskPrices.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarketplaceProvidersModelsDelete

> RegisteredModelPrices MarketplaceProvidersModelsDelete(ctx).Id(id).Execute()

Delete a custom model pricing entry



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/baselinehq/pricingapi-client-golang"
)

func main() {
	id := "id_example" // string | Model pricing ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.MarketplaceProvidersModelsDelete(context.Background()).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.MarketplaceProvidersModelsDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarketplaceProvidersModelsDelete`: RegisteredModelPrices
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.MarketplaceProvidersModelsDelete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarketplaceProvidersModelsDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | Model pricing ID | 

### Return type

[**RegisteredModelPrices**](RegisteredModelPrices.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarketplaceProvidersModelsGet

> ModelPriceList MarketplaceProvidersModelsGet(ctx).Limit(limit).Offset(offset).Model(model).Provider(provider).Region(region).TokenBucket(tokenBucket).Execute()

Get your custom model pricing entries



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/baselinehq/pricingapi-client-golang"
)

func main() {
	limit := int32(56) // int32 | Maximum entries to return (default 500) (optional)
	offset := int32(56) // int32 | Entries to skip (optional)
	model := "model_example" // string | Filter by model (optional)
	provider := "provider_example" // string | Filter by model provider (optional)
	region := "region_example" // string | Filter by region (optional)
	tokenBucket := "tokenBucket_example" // string | Filter by token bucket (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.MarketplaceProvidersModelsGet(context.Background()).Limit(limit).Offset(offset).Model(model).Provider(provider).Region(region).TokenBucket(tokenBucket).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.MarketplaceProvidersModelsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarketplaceProvidersModelsGet`: ModelPriceList
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.MarketplaceProvidersModelsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarketplaceProvidersModelsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Maximum entries to return (default 500) | 
 **offset** | **int32** | Entries to skip | 
 **model** | **string** | Filter by model | 
 **provider** | **string** | Filter by model provider | 
 **region** | **string** | Filter by region | 
 **tokenBucket** | **string** | Filter by token bucket | 

### Return type

[**ModelPriceList**](ModelPriceList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarketplaceProvidersModelsPost

> RegisteredModelPrices MarketplaceProvidersModelsPost(ctx).Instance(instance).Execute()

Register custom model pricing



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/baselinehq/pricingapi-client-golang"
)

func main() {
	instance := *openapiclient.NewModelPricesRequest([]openapiclient.ModelPrice{*openapiclient.NewModelPrice()}) // ModelPricesRequest | Custom model pricing request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.MarketplaceProvidersModelsPost(context.Background()).Instance(instance).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.MarketplaceProvidersModelsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarketplaceProvidersModelsPost`: RegisteredModelPrices
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.MarketplaceProvidersModelsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarketplaceProvidersModelsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instance** | [**ModelPricesRequest**](ModelPricesRequest.md) | Custom model pricing request | 

### Return type

[**RegisteredModelPrices**](RegisteredModelPrices.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PricingComputePost

> ComputePrice PricingComputePost(ctx).Instance(instance).Execute()

Get pricing for an instance



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/baselinehq/pricingapi-client-golang"
)

func main() {
	instance := *openapiclient.NewInstance() // Instance | Instance

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.PricingComputePost(context.Background()).Instance(instance).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.PricingComputePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PricingComputePost`: ComputePrice
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.PricingComputePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPricingComputePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instance** | [**Instance**](Instance.md) | Instance | 

### Return type

[**ComputePrice**](ComputePrice.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PricingDisksPost

> DiskPrice PricingDisksPost(ctx).Instance(instance).Execute()

Get pricing for a disk



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/baselinehq/pricingapi-client-golang"
)

func main() {
	instance := *openapiclient.NewDisk() // Disk | Disk

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.PricingDisksPost(context.Background()).Instance(instance).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.PricingDisksPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PricingDisksPost`: DiskPrice
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.PricingDisksPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPricingDisksPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instance** | [**Disk**](Disk.md) | Disk | 

### Return type

[**DiskPrice**](DiskPrice.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PricingPost

> ComputePrice PricingPost(ctx).Instance(instance).Execute()

Get pricing for an instance



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/baselinehq/pricingapi-client-golang"
)

func main() {
	instance := *openapiclient.NewInstance() // Instance | Instance

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.PricingPost(context.Background()).Instance(instance).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.PricingPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PricingPost`: ComputePrice
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.PricingPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPricingPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instance** | [**Instance**](Instance.md) | Instance | 

### Return type

[**ComputePrice**](ComputePrice.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProvidersGet

> ProviderConfigs ProvidersGet(ctx).Execute()

Get details for the providers



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/baselinehq/pricingapi-client-golang"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ProvidersGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ProvidersGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProvidersGet`: ProviderConfigs
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ProvidersGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiProvidersGetRequest struct via the builder pattern


### Return type

[**ProviderConfigs**](ProviderConfigs.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecommendationsComputePost

> []ComputeRecommendation RecommendationsComputePost(ctx).Instance(instance).Execute()

Get recommendations for compute instances



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/baselinehq/pricingapi-client-golang"
)

func main() {
	instance := *openapiclient.NewComputeRecommendationRequest() // ComputeRecommendationRequest | Instance

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.RecommendationsComputePost(context.Background()).Instance(instance).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.RecommendationsComputePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecommendationsComputePost`: []ComputeRecommendation
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.RecommendationsComputePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRecommendationsComputePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instance** | [**ComputeRecommendationRequest**](ComputeRecommendationRequest.md) | Instance | 

### Return type

[**[]ComputeRecommendation**](ComputeRecommendation.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecommendationsDisksPost

> []DiskRecommendation RecommendationsDisksPost(ctx).Instance(instance).Execute()

Get recommendations for disks



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/baselinehq/pricingapi-client-golang"
)

func main() {
	instance := *openapiclient.NewDiskRecommendationRequest() // DiskRecommendationRequest | Instance

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.RecommendationsDisksPost(context.Background()).Instance(instance).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.RecommendationsDisksPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecommendationsDisksPost`: []DiskRecommendation
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.RecommendationsDisksPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRecommendationsDisksPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instance** | [**DiskRecommendationRequest**](DiskRecommendationRequest.md) | Instance | 

### Return type

[**[]DiskRecommendation**](DiskRecommendation.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecommendationsPost

> []ComputeRecommendation RecommendationsPost(ctx).Instance(instance).Execute()

Get recommendations for compute instances



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/baselinehq/pricingapi-client-golang"
)

func main() {
	instance := *openapiclient.NewComputeRecommendationRequest() // ComputeRecommendationRequest | Instance

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.RecommendationsPost(context.Background()).Instance(instance).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.RecommendationsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecommendationsPost`: []ComputeRecommendation
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.RecommendationsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRecommendationsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instance** | [**ComputeRecommendationRequest**](ComputeRecommendationRequest.md) | Instance | 

### Return type

[**[]ComputeRecommendation**](ComputeRecommendation.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


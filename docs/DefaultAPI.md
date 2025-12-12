# \DefaultAPI

All URIs are relative to *https://pricing.baselinehq.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**HealthzGet**](DefaultAPI.md#HealthzGet) | **Get** /healthz | Health check endpoint
[**MarketplaceProvidersComputeDelete**](DefaultAPI.md#MarketplaceProvidersComputeDelete) | **Delete** /marketplace/providers/compute | Delete a custom provider instance
[**MarketplaceProvidersComputeGet**](DefaultAPI.md#MarketplaceProvidersComputeGet) | **Get** /marketplace/providers/compute | Get your custom pricing entries
[**MarketplaceProvidersComputePost**](DefaultAPI.md#MarketplaceProvidersComputePost) | **Post** /marketplace/providers/compute | Register a custom provider
[**PricingComputePost**](DefaultAPI.md#PricingComputePost) | **Post** /pricing/compute | Get  pricing for an instance
[**PricingDisksPost**](DefaultAPI.md#PricingDisksPost) | **Post** /pricing/disks | Get  pricing for a disk
[**ProvidersGet**](DefaultAPI.md#ProvidersGet) | **Get** /providers | Get details for the providers
[**RecommendationsComputePost**](DefaultAPI.md#RecommendationsComputePost) | **Post** /recommendations/compute | Get recommendations for compute instances
[**RecommendationsDisksPost**](DefaultAPI.md#RecommendationsDisksPost) | **Post** /recommendations/disks | Get recommendations for disks



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

> TypesCustomPricingResponse MarketplaceProvidersComputeDelete(ctx).Id(id).Execute()

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
	// response from `MarketplaceProvidersComputeDelete`: TypesCustomPricingResponse
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

[**TypesCustomPricingResponse**](TypesCustomPricingResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarketplaceProvidersComputeGet

> TypesMarketplaceProvidersResponse MarketplaceProvidersComputeGet(ctx).Execute()

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.MarketplaceProvidersComputeGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.MarketplaceProvidersComputeGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarketplaceProvidersComputeGet`: TypesMarketplaceProvidersResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.MarketplaceProvidersComputeGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiMarketplaceProvidersComputeGetRequest struct via the builder pattern


### Return type

[**TypesMarketplaceProvidersResponse**](TypesMarketplaceProvidersResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarketplaceProvidersComputePost

> TypesCustomPricingResponse MarketplaceProvidersComputePost(ctx).Instance(instance).Execute()

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
	instance := *openapiclient.NewTypesCustomPriceRequest([]openapiclient.SchemaComputePricingsRow{*openapiclient.NewSchemaComputePricingsRow()}) // TypesCustomPriceRequest | Pricing RecommendationRequest

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.MarketplaceProvidersComputePost(context.Background()).Instance(instance).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.MarketplaceProvidersComputePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarketplaceProvidersComputePost`: TypesCustomPricingResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.MarketplaceProvidersComputePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarketplaceProvidersComputePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instance** | [**TypesCustomPriceRequest**](TypesCustomPriceRequest.md) | Pricing RecommendationRequest | 

### Return type

[**TypesCustomPricingResponse**](TypesCustomPricingResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PricingComputePost

> SchemaComputePricingsRow PricingComputePost(ctx).Instance(instance).Execute()

Get  pricing for an instance



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
	instance := *openapiclient.NewGithubComBaselinehqGolangSharedTypesInstance() // GithubComBaselinehqGolangSharedTypesInstance | Instance

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.PricingComputePost(context.Background()).Instance(instance).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.PricingComputePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PricingComputePost`: SchemaComputePricingsRow
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.PricingComputePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPricingComputePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instance** | [**GithubComBaselinehqGolangSharedTypesInstance**](GithubComBaselinehqGolangSharedTypesInstance.md) | Instance | 

### Return type

[**SchemaComputePricingsRow**](SchemaComputePricingsRow.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PricingDisksPost

> SchemaDiskPricingsRow PricingDisksPost(ctx).Instance(instance).Execute()

Get  pricing for a disk



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
	instance := *openapiclient.NewTypesDisk() // TypesDisk | Disks

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.PricingDisksPost(context.Background()).Instance(instance).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.PricingDisksPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PricingDisksPost`: SchemaDiskPricingsRow
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.PricingDisksPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPricingDisksPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instance** | [**TypesDisk**](TypesDisk.md) | Disks | 

### Return type

[**SchemaDiskPricingsRow**](SchemaDiskPricingsRow.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

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

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecommendationsComputePost

> TypesComputeResults RecommendationsComputePost(ctx).Instance(instance).Execute()

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
	instance := *openapiclient.NewTypesComputeRequest() // TypesComputeRequest | Instance

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.RecommendationsComputePost(context.Background()).Instance(instance).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.RecommendationsComputePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecommendationsComputePost`: TypesComputeResults
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.RecommendationsComputePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRecommendationsComputePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instance** | [**TypesComputeRequest**](TypesComputeRequest.md) | Instance | 

### Return type

[**TypesComputeResults**](TypesComputeResults.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecommendationsDisksPost

> TypesDiskResults RecommendationsDisksPost(ctx).Instance(instance).Execute()

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
	instance := *openapiclient.NewTypesDiskRequest() // TypesDiskRequest | Instance

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.RecommendationsDisksPost(context.Background()).Instance(instance).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.RecommendationsDisksPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecommendationsDisksPost`: TypesDiskResults
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.RecommendationsDisksPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRecommendationsDisksPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instance** | [**TypesDiskRequest**](TypesDiskRequest.md) | Instance | 

### Return type

[**TypesDiskResults**](TypesDiskResults.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


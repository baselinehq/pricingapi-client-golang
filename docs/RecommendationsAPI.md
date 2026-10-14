# \RecommendationsAPI

All URIs are relative to *https://pricing.baselinehq.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1RecommendationsComputePost**](RecommendationsAPI.md#V1RecommendationsComputePost) | **Post** /v1/recommendations/compute | Find cheaper machines
[**V1RecommendationsDisksPost**](RecommendationsAPI.md#V1RecommendationsDisksPost) | **Post** /v1/recommendations/disks | Find cheaper volumes
[**V1RecommendationsRunnerPost**](RecommendationsAPI.md#V1RecommendationsRunnerPost) | **Post** /v1/recommendations/runner | Find cheaper CI runners



## V1RecommendationsComputePost

> []ComputeRecommendation V1RecommendationsComputePost(ctx).Instance(instance).Persist(persist).Execute()

Find cheaper machines



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
	instance := *openapiclient.NewComputeRecommendationRequest() // ComputeRecommendationRequest | Machine to improve on
	persist := true // bool | Save a price computed for a shape the catalogue does not list. Defaults to true; false returns the same recommendations without saving. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RecommendationsAPI.V1RecommendationsComputePost(context.Background()).Instance(instance).Persist(persist).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecommendationsAPI.V1RecommendationsComputePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1RecommendationsComputePost`: []ComputeRecommendation
	fmt.Fprintf(os.Stdout, "Response from `RecommendationsAPI.V1RecommendationsComputePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1RecommendationsComputePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instance** | [**ComputeRecommendationRequest**](ComputeRecommendationRequest.md) | Machine to improve on | 
 **persist** | **bool** | Save a price computed for a shape the catalogue does not list. Defaults to true; false returns the same recommendations without saving. | 

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


## V1RecommendationsDisksPost

> []DiskRecommendation V1RecommendationsDisksPost(ctx).Instance(instance).Execute()

Find cheaper volumes



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
	instance := *openapiclient.NewDiskRecommendationRequest() // DiskRecommendationRequest | Volume to improve on

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RecommendationsAPI.V1RecommendationsDisksPost(context.Background()).Instance(instance).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecommendationsAPI.V1RecommendationsDisksPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1RecommendationsDisksPost`: []DiskRecommendation
	fmt.Fprintf(os.Stdout, "Response from `RecommendationsAPI.V1RecommendationsDisksPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1RecommendationsDisksPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instance** | [**DiskRecommendationRequest**](DiskRecommendationRequest.md) | Volume to improve on | 

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


## V1RecommendationsRunnerPost

> []RunnerRecommendation V1RecommendationsRunnerPost(ctx).Request(request).Execute()

Find cheaper CI runners



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
	request := *openapiclient.NewRunnerRecommendationRequest() // RunnerRecommendationRequest | Jobs to price

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RecommendationsAPI.V1RecommendationsRunnerPost(context.Background()).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecommendationsAPI.V1RecommendationsRunnerPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1RecommendationsRunnerPost`: []RunnerRecommendation
	fmt.Fprintf(os.Stdout, "Response from `RecommendationsAPI.V1RecommendationsRunnerPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1RecommendationsRunnerPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**RunnerRecommendationRequest**](RunnerRecommendationRequest.md) | Jobs to price | 

### Return type

[**[]RunnerRecommendation**](RunnerRecommendation.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


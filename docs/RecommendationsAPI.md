# \RecommendationsAPI

All URIs are relative to *https://pricing.baselinehq.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1RecommendationsComputePost**](RecommendationsAPI.md#V1RecommendationsComputePost) | **Post** /v1/recommendations/compute | Find cheaper machines
[**V1RecommendationsDisksPost**](RecommendationsAPI.md#V1RecommendationsDisksPost) | **Post** /v1/recommendations/disks | Find cheaper volumes



## V1RecommendationsComputePost

> []ComputeRecommendation V1RecommendationsComputePost(ctx).Instance(instance).Execute()

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RecommendationsAPI.V1RecommendationsComputePost(context.Background()).Instance(instance).Execute()
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


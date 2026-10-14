# \PricingAPI

All URIs are relative to *https://pricing.baselinehq.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1PricingComputePost**](PricingAPI.md#V1PricingComputePost) | **Post** /v1/pricing/compute | Price a machine
[**V1PricingDisksPost**](PricingAPI.md#V1PricingDisksPost) | **Post** /v1/pricing/disks | Price a volume



## V1PricingComputePost

> ComputePrice V1PricingComputePost(ctx).Instance(instance).Persist(persist).Execute()

Price a machine



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
	instance := *openapiclient.NewInstance() // Instance | Machine to price
	persist := true // bool | Save a price computed for a shape the catalogue does not list, so later lookups return the same row. Defaults to true; false returns the same price without saving it. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PricingAPI.V1PricingComputePost(context.Background()).Instance(instance).Persist(persist).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricingAPI.V1PricingComputePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1PricingComputePost`: ComputePrice
	fmt.Fprintf(os.Stdout, "Response from `PricingAPI.V1PricingComputePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1PricingComputePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instance** | [**Instance**](Instance.md) | Machine to price | 
 **persist** | **bool** | Save a price computed for a shape the catalogue does not list, so later lookups return the same row. Defaults to true; false returns the same price without saving it. | 

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


## V1PricingDisksPost

> DiskPrice V1PricingDisksPost(ctx).Instance(instance).Execute()

Price a volume



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
	instance := *openapiclient.NewDisk() // Disk | Volume to price

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PricingAPI.V1PricingDisksPost(context.Background()).Instance(instance).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricingAPI.V1PricingDisksPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1PricingDisksPost`: DiskPrice
	fmt.Fprintf(os.Stdout, "Response from `PricingAPI.V1PricingDisksPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1PricingDisksPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instance** | [**Disk**](Disk.md) | Volume to price | 

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


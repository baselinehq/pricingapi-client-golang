# \ProvidersAPI

All URIs are relative to *https://pricing.baselinehq.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1ProvidersGet**](ProvidersAPI.md#V1ProvidersGet) | **Get** /v1/providers | List the providers CostGraph prices



## V1ProvidersGet

> ProviderConfigs V1ProvidersGet(ctx).Execute()

List the providers CostGraph prices



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
	resp, r, err := apiClient.ProvidersAPI.V1ProvidersGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProvidersAPI.V1ProvidersGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ProvidersGet`: ProviderConfigs
	fmt.Fprintf(os.Stdout, "Response from `ProvidersAPI.V1ProvidersGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1ProvidersGetRequest struct via the builder pattern


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


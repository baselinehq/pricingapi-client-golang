# \MarketplaceAPI

All URIs are relative to *https://pricing.baselinehq.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1MarketplaceComputeGet**](MarketplaceAPI.md#V1MarketplaceComputeGet) | **Get** /v1/marketplace/compute | List your compute prices
[**V1MarketplaceComputeIdDelete**](MarketplaceAPI.md#V1MarketplaceComputeIdDelete) | **Delete** /v1/marketplace/compute/{id} | Delete one compute price
[**V1MarketplaceComputeIdGet**](MarketplaceAPI.md#V1MarketplaceComputeIdGet) | **Get** /v1/marketplace/compute/{id} | Read one compute price
[**V1MarketplaceComputeIdPatch**](MarketplaceAPI.md#V1MarketplaceComputeIdPatch) | **Patch** /v1/marketplace/compute/{id} | Change one compute price
[**V1MarketplaceComputePost**](MarketplaceAPI.md#V1MarketplaceComputePost) | **Post** /v1/marketplace/compute | Publish compute prices
[**V1MarketplaceDatabasesGet**](MarketplaceAPI.md#V1MarketplaceDatabasesGet) | **Get** /v1/marketplace/databases | List your database prices
[**V1MarketplaceDatabasesIdDelete**](MarketplaceAPI.md#V1MarketplaceDatabasesIdDelete) | **Delete** /v1/marketplace/databases/{id} | Delete one database price
[**V1MarketplaceDatabasesIdGet**](MarketplaceAPI.md#V1MarketplaceDatabasesIdGet) | **Get** /v1/marketplace/databases/{id} | Read one database price
[**V1MarketplaceDatabasesIdPatch**](MarketplaceAPI.md#V1MarketplaceDatabasesIdPatch) | **Patch** /v1/marketplace/databases/{id} | Change one database price
[**V1MarketplaceDatabasesPost**](MarketplaceAPI.md#V1MarketplaceDatabasesPost) | **Post** /v1/marketplace/databases | Publish database prices
[**V1MarketplaceDisksGet**](MarketplaceAPI.md#V1MarketplaceDisksGet) | **Get** /v1/marketplace/disks | List your disk prices
[**V1MarketplaceDisksIdDelete**](MarketplaceAPI.md#V1MarketplaceDisksIdDelete) | **Delete** /v1/marketplace/disks/{id} | Delete one disk price
[**V1MarketplaceDisksIdGet**](MarketplaceAPI.md#V1MarketplaceDisksIdGet) | **Get** /v1/marketplace/disks/{id} | Read one disk price
[**V1MarketplaceDisksIdPatch**](MarketplaceAPI.md#V1MarketplaceDisksIdPatch) | **Patch** /v1/marketplace/disks/{id} | Change one disk price
[**V1MarketplaceDisksPost**](MarketplaceAPI.md#V1MarketplaceDisksPost) | **Post** /v1/marketplace/disks | Publish disk prices
[**V1MarketplaceFiltersKindGet**](MarketplaceAPI.md#V1MarketplaceFiltersKindGet) | **Get** /v1/marketplace/filters/{kind} | List the filters a price kind accepts
[**V1MarketplaceModelsGet**](MarketplaceAPI.md#V1MarketplaceModelsGet) | **Get** /v1/marketplace/models | List your model prices
[**V1MarketplaceModelsIdDelete**](MarketplaceAPI.md#V1MarketplaceModelsIdDelete) | **Delete** /v1/marketplace/models/{id} | Delete one model price
[**V1MarketplaceModelsIdGet**](MarketplaceAPI.md#V1MarketplaceModelsIdGet) | **Get** /v1/marketplace/models/{id} | Read one model price
[**V1MarketplaceModelsIdPatch**](MarketplaceAPI.md#V1MarketplaceModelsIdPatch) | **Patch** /v1/marketplace/models/{id} | Change one model price
[**V1MarketplaceModelsPost**](MarketplaceAPI.md#V1MarketplaceModelsPost) | **Post** /v1/marketplace/models | Publish model prices



## V1MarketplaceComputeGet

> ComputePricePage V1MarketplaceComputeGet(ctx).Limit(limit).Cursor(cursor).Service(service).Region(region).InstanceType(instanceType).UsageType(usageType).Execute()

List your compute prices



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
	limit := int32(56) // int32 | Maximum prices to return (default 500) (optional)
	cursor := "cursor_example" // string | Cursor from the previous page (optional)
	service := "service_example" // string | Filter by service (optional)
	region := "region_example" // string | Filter by region (optional)
	instanceType := "instanceType_example" // string | Filter by instance type (optional)
	usageType := "usageType_example" // string | Filter by usage type (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketplaceAPI.V1MarketplaceComputeGet(context.Background()).Limit(limit).Cursor(cursor).Service(service).Region(region).InstanceType(instanceType).UsageType(usageType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketplaceAPI.V1MarketplaceComputeGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1MarketplaceComputeGet`: ComputePricePage
	fmt.Fprintf(os.Stdout, "Response from `MarketplaceAPI.V1MarketplaceComputeGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1MarketplaceComputeGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Maximum prices to return (default 500) | 
 **cursor** | **string** | Cursor from the previous page | 
 **service** | **string** | Filter by service | 
 **region** | **string** | Filter by region | 
 **instanceType** | **string** | Filter by instance type | 
 **usageType** | **string** | Filter by usage type | 

### Return type

[**ComputePricePage**](ComputePricePage.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1MarketplaceComputeIdDelete

> V1MarketplaceComputeIdDelete(ctx, id).Execute()

Delete one compute price



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
	id := "id_example" // string | Price id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MarketplaceAPI.V1MarketplaceComputeIdDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketplaceAPI.V1MarketplaceComputeIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Price id | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1MarketplaceComputeIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1MarketplaceComputeIdGet

> ComputePrice V1MarketplaceComputeIdGet(ctx, id).Execute()

Read one compute price



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
	id := "id_example" // string | Price id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketplaceAPI.V1MarketplaceComputeIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketplaceAPI.V1MarketplaceComputeIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1MarketplaceComputeIdGet`: ComputePrice
	fmt.Fprintf(os.Stdout, "Response from `MarketplaceAPI.V1MarketplaceComputeIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Price id | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1MarketplaceComputeIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ComputePrice**](ComputePrice.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1MarketplaceComputeIdPatch

> ComputePrice V1MarketplaceComputeIdPatch(ctx, id).Body(body).Execute()

Change one compute price



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
	id := "id_example" // string | Price id
	body := map[string]interface{}{"key": interface{}(123)} // map[string]interface{} | Fields to change

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketplaceAPI.V1MarketplaceComputeIdPatch(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketplaceAPI.V1MarketplaceComputeIdPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1MarketplaceComputeIdPatch`: ComputePrice
	fmt.Fprintf(os.Stdout, "Response from `MarketplaceAPI.V1MarketplaceComputeIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Price id | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1MarketplaceComputeIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** | Fields to change | 

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


## V1MarketplaceComputePost

> PublishedComputePrices V1MarketplaceComputePost(ctx).Body(body).Execute()

Publish compute prices



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
	body := *openapiclient.NewComputePricesRequest([]openapiclient.ComputePrice{*openapiclient.NewComputePrice()}) // ComputePricesRequest | Prices to publish

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketplaceAPI.V1MarketplaceComputePost(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketplaceAPI.V1MarketplaceComputePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1MarketplaceComputePost`: PublishedComputePrices
	fmt.Fprintf(os.Stdout, "Response from `MarketplaceAPI.V1MarketplaceComputePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1MarketplaceComputePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ComputePricesRequest**](ComputePricesRequest.md) | Prices to publish | 

### Return type

[**PublishedComputePrices**](PublishedComputePrices.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1MarketplaceDatabasesGet

> DatabasePricePage V1MarketplaceDatabasesGet(ctx).Limit(limit).Cursor(cursor).Service(service).Region(region).Engine(engine).Edition(edition).InstanceType(instanceType).UsageType(usageType).Execute()

List your database prices



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
	limit := int32(56) // int32 | Maximum prices to return (default 500) (optional)
	cursor := "cursor_example" // string | Cursor from the previous page (optional)
	service := "service_example" // string | Filter by service (optional)
	region := "region_example" // string | Filter by region (optional)
	engine := "engine_example" // string | Filter by engine (optional)
	edition := "edition_example" // string | Filter by edition (optional)
	instanceType := "instanceType_example" // string | Filter by instance type (optional)
	usageType := "usageType_example" // string | Filter by usage type (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketplaceAPI.V1MarketplaceDatabasesGet(context.Background()).Limit(limit).Cursor(cursor).Service(service).Region(region).Engine(engine).Edition(edition).InstanceType(instanceType).UsageType(usageType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketplaceAPI.V1MarketplaceDatabasesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1MarketplaceDatabasesGet`: DatabasePricePage
	fmt.Fprintf(os.Stdout, "Response from `MarketplaceAPI.V1MarketplaceDatabasesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1MarketplaceDatabasesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Maximum prices to return (default 500) | 
 **cursor** | **string** | Cursor from the previous page | 
 **service** | **string** | Filter by service | 
 **region** | **string** | Filter by region | 
 **engine** | **string** | Filter by engine | 
 **edition** | **string** | Filter by edition | 
 **instanceType** | **string** | Filter by instance type | 
 **usageType** | **string** | Filter by usage type | 

### Return type

[**DatabasePricePage**](DatabasePricePage.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1MarketplaceDatabasesIdDelete

> V1MarketplaceDatabasesIdDelete(ctx, id).Execute()

Delete one database price



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
	id := "id_example" // string | Price id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MarketplaceAPI.V1MarketplaceDatabasesIdDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketplaceAPI.V1MarketplaceDatabasesIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Price id | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1MarketplaceDatabasesIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1MarketplaceDatabasesIdGet

> DatabasePrice V1MarketplaceDatabasesIdGet(ctx, id).Execute()

Read one database price



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
	id := "id_example" // string | Price id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketplaceAPI.V1MarketplaceDatabasesIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketplaceAPI.V1MarketplaceDatabasesIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1MarketplaceDatabasesIdGet`: DatabasePrice
	fmt.Fprintf(os.Stdout, "Response from `MarketplaceAPI.V1MarketplaceDatabasesIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Price id | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1MarketplaceDatabasesIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DatabasePrice**](DatabasePrice.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1MarketplaceDatabasesIdPatch

> DatabasePrice V1MarketplaceDatabasesIdPatch(ctx, id).Body(body).Execute()

Change one database price



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
	id := "id_example" // string | Price id
	body := map[string]interface{}{"key": interface{}(123)} // map[string]interface{} | Fields to change

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketplaceAPI.V1MarketplaceDatabasesIdPatch(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketplaceAPI.V1MarketplaceDatabasesIdPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1MarketplaceDatabasesIdPatch`: DatabasePrice
	fmt.Fprintf(os.Stdout, "Response from `MarketplaceAPI.V1MarketplaceDatabasesIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Price id | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1MarketplaceDatabasesIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** | Fields to change | 

### Return type

[**DatabasePrice**](DatabasePrice.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1MarketplaceDatabasesPost

> PublishedDatabasePrices V1MarketplaceDatabasesPost(ctx).Body(body).Execute()

Publish database prices



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
	body := *openapiclient.NewDatabasePricesRequest([]openapiclient.DatabasePrice{*openapiclient.NewDatabasePrice()}) // DatabasePricesRequest | Prices to publish

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketplaceAPI.V1MarketplaceDatabasesPost(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketplaceAPI.V1MarketplaceDatabasesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1MarketplaceDatabasesPost`: PublishedDatabasePrices
	fmt.Fprintf(os.Stdout, "Response from `MarketplaceAPI.V1MarketplaceDatabasesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1MarketplaceDatabasesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DatabasePricesRequest**](DatabasePricesRequest.md) | Prices to publish | 

### Return type

[**PublishedDatabasePrices**](PublishedDatabasePrices.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1MarketplaceDisksGet

> DiskPricePage V1MarketplaceDisksGet(ctx).Limit(limit).Cursor(cursor).Service(service).Region(region).Type_(type_).UsageType(usageType).Execute()

List your disk prices



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
	limit := int32(56) // int32 | Maximum prices to return (default 500) (optional)
	cursor := "cursor_example" // string | Cursor from the previous page (optional)
	service := "service_example" // string | Filter by service (optional)
	region := "region_example" // string | Filter by region (optional)
	type_ := "type__example" // string | Filter by disk type (optional)
	usageType := "usageType_example" // string | Filter by usage type (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketplaceAPI.V1MarketplaceDisksGet(context.Background()).Limit(limit).Cursor(cursor).Service(service).Region(region).Type_(type_).UsageType(usageType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketplaceAPI.V1MarketplaceDisksGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1MarketplaceDisksGet`: DiskPricePage
	fmt.Fprintf(os.Stdout, "Response from `MarketplaceAPI.V1MarketplaceDisksGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1MarketplaceDisksGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Maximum prices to return (default 500) | 
 **cursor** | **string** | Cursor from the previous page | 
 **service** | **string** | Filter by service | 
 **region** | **string** | Filter by region | 
 **type_** | **string** | Filter by disk type | 
 **usageType** | **string** | Filter by usage type | 

### Return type

[**DiskPricePage**](DiskPricePage.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1MarketplaceDisksIdDelete

> V1MarketplaceDisksIdDelete(ctx, id).Execute()

Delete one disk price



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
	id := "id_example" // string | Price id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MarketplaceAPI.V1MarketplaceDisksIdDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketplaceAPI.V1MarketplaceDisksIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Price id | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1MarketplaceDisksIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1MarketplaceDisksIdGet

> DiskPrice V1MarketplaceDisksIdGet(ctx, id).Execute()

Read one disk price



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
	id := "id_example" // string | Price id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketplaceAPI.V1MarketplaceDisksIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketplaceAPI.V1MarketplaceDisksIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1MarketplaceDisksIdGet`: DiskPrice
	fmt.Fprintf(os.Stdout, "Response from `MarketplaceAPI.V1MarketplaceDisksIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Price id | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1MarketplaceDisksIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DiskPrice**](DiskPrice.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1MarketplaceDisksIdPatch

> DiskPrice V1MarketplaceDisksIdPatch(ctx, id).Body(body).Execute()

Change one disk price



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
	id := "id_example" // string | Price id
	body := map[string]interface{}{"key": interface{}(123)} // map[string]interface{} | Fields to change

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketplaceAPI.V1MarketplaceDisksIdPatch(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketplaceAPI.V1MarketplaceDisksIdPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1MarketplaceDisksIdPatch`: DiskPrice
	fmt.Fprintf(os.Stdout, "Response from `MarketplaceAPI.V1MarketplaceDisksIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Price id | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1MarketplaceDisksIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** | Fields to change | 

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


## V1MarketplaceDisksPost

> PublishedDiskPrices V1MarketplaceDisksPost(ctx).Body(body).Execute()

Publish disk prices



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
	body := *openapiclient.NewDiskPricesRequest([]openapiclient.DiskPrice{*openapiclient.NewDiskPrice()}) // DiskPricesRequest | Prices to publish

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketplaceAPI.V1MarketplaceDisksPost(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketplaceAPI.V1MarketplaceDisksPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1MarketplaceDisksPost`: PublishedDiskPrices
	fmt.Fprintf(os.Stdout, "Response from `MarketplaceAPI.V1MarketplaceDisksPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1MarketplaceDisksPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DiskPricesRequest**](DiskPricesRequest.md) | Prices to publish | 

### Return type

[**PublishedDiskPrices**](PublishedDiskPrices.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1MarketplaceFiltersKindGet

> FilterSet V1MarketplaceFiltersKindGet(ctx, kind).Execute()

List the filters a price kind accepts



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
	kind := "kind_example" // string | Price kind

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketplaceAPI.V1MarketplaceFiltersKindGet(context.Background(), kind).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketplaceAPI.V1MarketplaceFiltersKindGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1MarketplaceFiltersKindGet`: FilterSet
	fmt.Fprintf(os.Stdout, "Response from `MarketplaceAPI.V1MarketplaceFiltersKindGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**kind** | **string** | Price kind | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1MarketplaceFiltersKindGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FilterSet**](FilterSet.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1MarketplaceModelsGet

> ModelPricePage V1MarketplaceModelsGet(ctx).Limit(limit).Cursor(cursor).Model(model).Provider(provider).Region(region).TokenBucket(tokenBucket).Execute()

List your model prices



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
	limit := int32(56) // int32 | Maximum prices to return (default 500) (optional)
	cursor := "cursor_example" // string | Cursor from the previous page (optional)
	model := "model_example" // string | Filter by model (optional)
	provider := "provider_example" // string | Filter by model provider (optional)
	region := "region_example" // string | Filter by region (optional)
	tokenBucket := "tokenBucket_example" // string | Filter by token bucket (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketplaceAPI.V1MarketplaceModelsGet(context.Background()).Limit(limit).Cursor(cursor).Model(model).Provider(provider).Region(region).TokenBucket(tokenBucket).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketplaceAPI.V1MarketplaceModelsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1MarketplaceModelsGet`: ModelPricePage
	fmt.Fprintf(os.Stdout, "Response from `MarketplaceAPI.V1MarketplaceModelsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1MarketplaceModelsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Maximum prices to return (default 500) | 
 **cursor** | **string** | Cursor from the previous page | 
 **model** | **string** | Filter by model | 
 **provider** | **string** | Filter by model provider | 
 **region** | **string** | Filter by region | 
 **tokenBucket** | **string** | Filter by token bucket | 

### Return type

[**ModelPricePage**](ModelPricePage.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1MarketplaceModelsIdDelete

> V1MarketplaceModelsIdDelete(ctx, id).Execute()

Delete one model price



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
	id := "id_example" // string | Price id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MarketplaceAPI.V1MarketplaceModelsIdDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketplaceAPI.V1MarketplaceModelsIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Price id | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1MarketplaceModelsIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1MarketplaceModelsIdGet

> ModelPrice V1MarketplaceModelsIdGet(ctx, id).Execute()

Read one model price



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
	id := "id_example" // string | Price id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketplaceAPI.V1MarketplaceModelsIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketplaceAPI.V1MarketplaceModelsIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1MarketplaceModelsIdGet`: ModelPrice
	fmt.Fprintf(os.Stdout, "Response from `MarketplaceAPI.V1MarketplaceModelsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Price id | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1MarketplaceModelsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ModelPrice**](ModelPrice.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1MarketplaceModelsIdPatch

> ModelPrice V1MarketplaceModelsIdPatch(ctx, id).Body(body).Execute()

Change one model price



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
	id := "id_example" // string | Price id
	body := map[string]interface{}{"key": interface{}(123)} // map[string]interface{} | Fields to change

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketplaceAPI.V1MarketplaceModelsIdPatch(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketplaceAPI.V1MarketplaceModelsIdPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1MarketplaceModelsIdPatch`: ModelPrice
	fmt.Fprintf(os.Stdout, "Response from `MarketplaceAPI.V1MarketplaceModelsIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Price id | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1MarketplaceModelsIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** | Fields to change | 

### Return type

[**ModelPrice**](ModelPrice.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1MarketplaceModelsPost

> PublishedModelPrices V1MarketplaceModelsPost(ctx).Body(body).Execute()

Publish model prices



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
	body := *openapiclient.NewModelPricesRequest([]openapiclient.ModelPrice{*openapiclient.NewModelPrice()}) // ModelPricesRequest | Prices to publish

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketplaceAPI.V1MarketplaceModelsPost(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketplaceAPI.V1MarketplaceModelsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1MarketplaceModelsPost`: PublishedModelPrices
	fmt.Fprintf(os.Stdout, "Response from `MarketplaceAPI.V1MarketplaceModelsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1MarketplaceModelsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ModelPricesRequest**](ModelPricesRequest.md) | Prices to publish | 

### Return type

[**PublishedModelPrices**](PublishedModelPrices.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


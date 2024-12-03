# \AdminAPI

All URIs are relative to *http://localhost:4000*

 Method                                                     | HTTP request                         | Description
------------------------------------------------------------|--------------------------------------|--------------------------
 [**CreateAdminPayment**](AdminAPI.md#CreateAdminPayment)   | **Post** /apps/payments              | Create admin payment
 [**GetAdminPaymentById**](AdminAPI.md#GetAdminPaymentById) | **Post** /apps/payments/{payment_id} | Get payment
 [**GetAppEntities**](AdminAPI.md#GetAppEntities)           | **Get** /apps                        | Get application entities
 [**GetAppEntity**](AdminAPI.md#GetAppEntity)               | **Get** /apps/{app_entity_id}        | Get application entity



## CreateAdminPayment

> Payment CreateAdminPayment(ctx).Execute()

Create admin payment



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdminAPI.CreateAdminPayment(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.CreateAdminPayment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAdminPayment`: Payment
    fmt.Fprintf(os.Stdout, "Response from `AdminAPI.CreateAdminPayment`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAdminPaymentRequest struct via the builder pattern


### Return type

[**Payment**](Payment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAdminPaymentById

> Payment GetAdminPaymentById(ctx, paymentId).Execute()

Get payment



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    paymentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdminAPI.GetAdminPaymentById(context.Background(), paymentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.GetAdminPaymentById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAdminPaymentById`: Payment
    fmt.Fprintf(os.Stdout, "Response from `AdminAPI.GetAdminPaymentById`: %v\n", resp)
}
```

### Path Parameters


 Name          | Type                | Description                                                                 | Notes
---------------|---------------------|-----------------------------------------------------------------------------|-------
 **ctx**       | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **paymentId** | **string**          |                                                                             |

### Other Parameters

Other parameters are passed through a pointer to a apiGetAdminPaymentByIdRequest struct via the builder pattern


 Name | Type | Description | Notes
------|------|-------------|-------


### Return type

[**Payment**](Payment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAppEntities

> []AppEntity GetAppEntities(ctx).Execute()

Get application entities



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdminAPI.GetAppEntities(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.GetAppEntities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAppEntities`: []AppEntity
    fmt.Fprintf(os.Stdout, "Response from `AdminAPI.GetAppEntities`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAppEntitiesRequest struct via the builder pattern


### Return type

[**[]AppEntity**](AppEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAppEntity

> AppEntity GetAppEntity(ctx, appEntityId).Execute()

Get application entity



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    appEntityId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdminAPI.GetAppEntity(context.Background(), appEntityId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.GetAppEntity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAppEntity`: AppEntity
    fmt.Fprintf(os.Stdout, "Response from `AdminAPI.GetAppEntity`: %v\n", resp)
}
```

### Path Parameters


 Name            | Type                | Description                                                                 | Notes
-----------------|---------------------|-----------------------------------------------------------------------------|-------
 **ctx**         | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **appEntityId** | **string**          |                                                                             |

### Other Parameters

Other parameters are passed through a pointer to a apiGetAppEntityRequest struct via the builder pattern


 Name | Type | Description | Notes
------|------|-------------|-------


### Return type

[**AppEntity**](AppEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


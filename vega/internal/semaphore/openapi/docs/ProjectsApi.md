# \ProjectsApi

All URIs are relative to *https://localhost:3000/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProjectsGet**](ProjectsApi.md#ProjectsGet) | **Get** /projects | Get projects
[**ProjectsPost**](ProjectsApi.md#ProjectsPost) | **Post** /projects | Create a new project



## ProjectsGet

> []Project ProjectsGet(ctx).Execute()

Get projects

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProjectsApi.ProjectsGet(context.Background()).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectsGet`: []Project
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ProjectsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsGetRequest struct via the builder pattern


### Return type

[**[]Project**](Project.md)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsPost

> ProjectsPost(ctx).Project(project).Execute()

Create a new project

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    project := *openapiclient.NewProjectRequest() // ProjectRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProjectsApi.ProjectsPost(context.Background()).Project(project).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProjectsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **project** | [**ProjectRequest**](ProjectRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer), [cookie](../README.md#cookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


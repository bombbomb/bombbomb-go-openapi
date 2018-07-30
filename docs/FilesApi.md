# \FilesApi

All URIs are relative to *https://api.bombbomb.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DocHostDelete**](FilesApi.md#DocHostDelete) | **Delete** /files/{docId} | Delete file
[**DocHostGet**](FilesApi.md#DocHostGet) | **Get** /files/{docId} | Get file
[**DocHostList**](FilesApi.md#DocHostList) | **Get** /files | List all files
[**DocHostUploadV2**](FilesApi.md#DocHostUploadV2) | **Post** /files | Upload a file
[**GetHostedImagesPaged**](FilesApi.md#GetHostedImagesPaged) | **Get** /files/images/paged | Get paged hosted images


# **DocHostDelete**
> DocHostDelete(ctx, docId)
Delete file

Deletes a users file

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **docId** | **string**| Id of document | 

### Return type

 (empty response body)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DocHostGet**
> HostedDoc DocHostGet(ctx, docId)
Get file

Get a single file by id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **docId** | **string**| Id of document | 

### Return type

[**HostedDoc**](HostedDoc.md)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DocHostList**
> []HostedDoc DocHostList(ctx, )
List all files

List all uploaded user files

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]HostedDoc**](HostedDoc.md)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DocHostUploadV2**
> []HostedDoc DocHostUploadV2(ctx, file)
Upload a file

Upload a new file

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **file** | **string**| The file being uploaded | 

### Return type

[**[]HostedDoc**](HostedDoc.md)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetHostedImagesPaged**
> GetHostedImagesPaged(ctx, pageSize, page, optional)
Get paged hosted images

Get a specific page of uploaded images available to the user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **pageSize** | **string**| The number of items to retrieve in a single db query. | 
  **page** | **string**| Zero-based index of the page of data to retrieve from the db. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **string**| The number of items to retrieve in a single db query. | 
 **page** | **string**| Zero-based index of the page of data to retrieve from the db. | 
 **search** | **string**| Filter results with names that match the search term. | 

### Return type

 (empty response body)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


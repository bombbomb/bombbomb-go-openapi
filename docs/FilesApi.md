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
> DocHostDelete($docId)

Delete file

Deletes a users file


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **docId** | **string**| Id of document | 

### Return type

void (empty response body)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DocHostGet**
> HostedDoc DocHostGet($docId)

Get file

Get a single file by id


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
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
> []HostedDoc DocHostList()

List all files

List all uploaded user files


### Parameters
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
> []HostedDoc DocHostUploadV2($file)

Upload a file

Upload a new file


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
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
> GetHostedImagesPaged($pageSize, $page, $search)

Get paged hosted images

Get a specific page of uploaded images available to the user.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **string**| The number of items to retrieve in a single db query. | 
 **page** | **string**| Zero-based index of the page of data to retrieve from the db. | 
 **search** | **string**| Filter results with names that match the search term. | [optional] 

### Return type

void (empty response body)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


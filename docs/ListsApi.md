# \ListsApi

All URIs are relative to *https://api.bombbomb.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddNewList**](ListsApi.md#AddNewList) | **Post** /lists/ | Add list.
[**ClearList**](ListsApi.md#ClearList) | **Put** /lists/{listId}/clear | Clear Contacts from List
[**CopyListContacts**](ListsApi.md#CopyListContacts) | **Post** /lists/{listId}/copy | Copy All Contacts from a List
[**GetAllLists**](ListsApi.md#GetAllLists) | **Get** /lists/ | Get all Lists
[**SuppressAllInList**](ListsApi.md#SuppressAllInList) | **Put** /lists/{listId}/suppress | Suppress All Contacts from List


# **AddNewList**
> AddNewList(ctx, listName)
Add list.

Add a list to the users account.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **listName** | **string**| Name of the new list being added | 

### Return type

 (empty response body)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClearList**
> ClearList(ctx, listId)
Clear Contacts from List

Clears all contacts from a list.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **listId** | **string**| The list to be cleared. | 

### Return type

 (empty response body)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CopyListContacts**
> CopyListContacts(ctx, fromListId, listId)
Copy All Contacts from a List

Copy all contacts from a list.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **fromListId** | **string**| The list to be cleared. | 
  **listId** | **string**| The list to be cleared. | 

### Return type

 (empty response body)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllLists**
> GetAllLists(ctx, )
Get all Lists

Get all the lists for a specific user.

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SuppressAllInList**
> SuppressAllInList(ctx, listId)
Suppress All Contacts from List

Suppresses all contacts in a list.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **listId** | **string**| The list to be cleared. | 

### Return type

 (empty response body)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


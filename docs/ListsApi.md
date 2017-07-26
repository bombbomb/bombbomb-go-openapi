# \ListsApi

All URIs are relative to *https://api.bombbomb.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClearList**](ListsApi.md#ClearList) | **Put** /lists/{listId}/clear | Clear Contacts from List
[**CopyListContacts**](ListsApi.md#CopyListContacts) | **Post** /lists/{listId}/copy | Copy All Contacts from a List
[**SuppressAllInList**](ListsApi.md#SuppressAllInList) | **Put** /lists/{listId}/suppress | Suppress All Contacts from List


# **ClearList**
> ClearList($listId)

Clear Contacts from List

Clears all contacts from a list.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **listId** | **string**| The list to be cleared. | 

### Return type

void (empty response body)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CopyListContacts**
> CopyListContacts($fromListId, $listId)

Copy All Contacts from a List

Copy all contacts from a list.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fromListId** | **string**| The list to be cleared. | 
 **listId** | **string**| The list to be cleared. | 

### Return type

void (empty response body)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SuppressAllInList**
> SuppressAllInList($listId)

Suppress All Contacts from List

Suppresses all contacts in a list.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **listId** | **string**| The list to be cleared. | 

### Return type

void (empty response body)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


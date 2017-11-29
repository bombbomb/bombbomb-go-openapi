# \IntegrationsApi

All URIs are relative to *https://api.bombbomb.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SyncUsersIntegratedLists**](IntegrationsApi.md#SyncUsersIntegratedLists) | **Get** /integrations/sync | Synchronize your integration list or lists.


# **SyncUsersIntegratedLists**
> string SyncUsersIntegratedLists($integrationId)

Synchronize your integration list or lists.

Synchronize your integration contact list with the service you are integrated with. If no integration code is provided, all integrations will be synchronized.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **integrationId** | **string**| The integration to sync lists for. All integrations will sync if nothing is provided. | [optional] 

### Return type

**string**

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


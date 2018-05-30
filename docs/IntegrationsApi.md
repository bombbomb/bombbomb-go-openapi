# \IntegrationsApi

All URIs are relative to *https://api.bombbomb.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConnectIntegration**](IntegrationsApi.md#ConnectIntegration) | **Post** /integrations | Activate an integration for a user.
[**DeleteIntegration**](IntegrationsApi.md#DeleteIntegration) | **Delete** /integrations | Remove an integration for a user.
[**GetIntegrationHealth**](IntegrationsApi.md#GetIntegrationHealth) | **Get** /integrations/health/{code} | Get health for a given integration
[**GetIntegrationPageComponents**](IntegrationsApi.md#GetIntegrationPageComponents) | **Get** /integrations/pageComponents | Get page components for a given integration
[**SyncUsersIntegratedLists**](IntegrationsApi.md#SyncUsersIntegratedLists) | **Get** /integrations/sync | Synchronize your integration list or lists.


# **ConnectIntegration**
> ConnectIntegration($code, $key, $secret, $token, $data, $overwrite)

Activate an integration for a user.

Provide the correct parameters to enable an integration. Required Parameters vary based on the desired          integration. Integrations requiring OAuth will provide the OAuth link that the user must be presented.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **code** | **string**| The identifier of the integration. | 
 **key** | **string**| The key value. | [optional] 
 **secret** | **string**| The secret value. | [optional] 
 **token** | **string**| The token value. | [optional] 
 **data** | **string**| The data value as JSON. | [optional] 
 **overwrite** | **string**| Boolean value to know whether or not to delete the integration if it already exists | [optional] 

### Return type

void (empty response body)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteIntegration**
> DeleteIntegration($id, $code)

Remove an integration for a user.

Remove an integration by providing the integration id or integration code. Only provide one of the             parameters.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| Integration ID | [optional] 
 **code** | **string**| Integration Code | [optional] 

### Return type

void (empty response body)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIntegrationHealth**
> GetIntegrationHealth($code)

Get health for a given integration

Get health for an integration.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **code** | **string**| The integration code for which to retrieve the information from | 

### Return type

void (empty response body)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIntegrationPageComponents**
> GetIntegrationPageComponents($integrationName)

Get page components for a given integration

Get all page components for an integration.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **integrationName** | **string**| The integration for which to retrieve HTML page components. | 

### Return type

void (empty response body)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

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


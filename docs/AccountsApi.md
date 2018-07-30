# \AccountsApi

All URIs are relative to *https://api.bombbomb.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AccountDetails**](AccountsApi.md#AccountDetails) | **Get** /accounts | Get account details.
[**CreateAccount**](AccountsApi.md#CreateAccount) | **Post** /accounts | Create Account
[**GetClientStatistics**](AccountsApi.md#GetClientStatistics) | **Get** /accounts/stats | Get Client Statistics
[**GetUserCountry**](AccountsApi.md#GetUserCountry) | **Get** /accounts/{clientId}/country | Gets user country
[**ResetApiKey**](AccountsApi.md#ResetApiKey) | **Put** /accounts/apikey | Reset API key
[**SubscriptionPurchaseAllowed**](AccountsApi.md#SubscriptionPurchaseAllowed) | **Get** /accounts/purchaseable | Check if subscription purchase allowed.
[**UpdateProfileData**](AccountsApi.md#UpdateProfileData) | **Post** /account/profile/ | Add profile information.


# **AccountDetails**
> AccountDetails(ctx, )
Get account details.

Get the details of the user's account.

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

# **CreateAccount**
> string CreateAccount(ctx, teamId, firstName, lastName, emailAddress, companyName, phone, optional)
Create Account

Creates a new BombBomb account. This method is currently only available to paid seat admins.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **teamId** | **string**| The team id | 
  **firstName** | **string**| First name of the user. | 
  **lastName** | **string**| Last name of the user. | 
  **emailAddress** | **string**| Email address of the user. | 
  **companyName** | **string**| Company of the user. | 
  **phone** | **string**| Phone number of the user. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **teamId** | **string**| The team id | 
 **firstName** | **string**| First name of the user. | 
 **lastName** | **string**| Last name of the user. | 
 **emailAddress** | **string**| Email address of the user. | 
 **companyName** | **string**| Company of the user. | 
 **phone** | **string**| Phone number of the user. | 
 **country** | **string**| Country of the user. | 
 **industry** | **string**| Industry of the user. | 
 **address** | **string**| Street Address of the user. | 
 **city** | **string**| City of the user. | 
 **postalCode** | **string**| Postal/Zip code of the user. | 
 **preventWelcomeEmail** | **bool**| prevent an email with login credentials from being sent to the new account. must be set to &#39;true&#39; | 

### Return type

**string**

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetClientStatistics**
> GetClientStatistics(ctx, optional)
Get Client Statistics

Gets general statics for a Client

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientId** | **string**| Client ID of the account to retrieve. Defaults to yourself. | 
 **refresh** | **bool**| Boolean for whether data returned should be from cache or not. | 
 **statisticValues** | **string**| Array of data that should be returned, used exclusively for cacheless data | 

### Return type

 (empty response body)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUserCountry**
> GetUserCountry(ctx, )
Gets user country

Gets the users country

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

# **ResetApiKey**
> ResetApiKey(ctx, )
Reset API key

Resets the current user's API key and returns the new key

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

# **SubscriptionPurchaseAllowed**
> SubscriptionPurchaseAllowed(ctx, )
Check if subscription purchase allowed.

Check whether the user can purchase a subscription.

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

# **UpdateProfileData**
> UpdateProfileData(ctx, optional)
Add profile information.

Add profile information to this users account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **profileData** | **string**| Profile field information for the account | 

### Return type

 (empty response body)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


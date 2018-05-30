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
> AccountDetails()

Get account details.

Get the details of the user's account.


### Parameters
This endpoint does not need any parameter.

### Return type

void (empty response body)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateAccount**
> string CreateAccount($teamId, $firstName, $lastName, $emailAddress, $companyName, $phone, $country, $industry, $address, $city, $postalCode, $preventWelcomeEmail)

Create Account

Creates a new BombBomb account. This method is currently only available to paid seat admins.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **teamId** | **string**| The team id | 
 **firstName** | **string**| First name of the user. | 
 **lastName** | **string**| Last name of the user. | 
 **emailAddress** | **string**| Email address of the user. | 
 **companyName** | **string**| Company of the user. | 
 **phone** | **string**| Phone number of the user. | 
 **country** | **string**| Country of the user. | [optional] 
 **industry** | **string**| Industry of the user. | [optional] 
 **address** | **string**| Street Address of the user. | [optional] 
 **city** | **string**| City of the user. | [optional] 
 **postalCode** | **string**| Postal/Zip code of the user. | [optional] 
 **preventWelcomeEmail** | **bool**| prevent an email with login credentials from being sent to the new account. must be set to &#39;true&#39; | [optional] 

### Return type

**string**

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetClientStatistics**
> GetClientStatistics($clientId, $refresh, $statisticValues)

Get Client Statistics

Gets general statics for a Client


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientId** | **string**| Client ID of the account to retrieve. Defaults to yourself. | [optional] 
 **refresh** | **bool**| Boolean for whether data returned should be from cache or not. | [optional] 
 **statisticValues** | **string**| Array of data that should be returned, used exclusively for cacheless data | [optional] 

### Return type

void (empty response body)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUserCountry**
> GetUserCountry()

Gets user country

Gets the users country


### Parameters
This endpoint does not need any parameter.

### Return type

void (empty response body)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ResetApiKey**
> ResetApiKey()

Reset API key

Resets the current user's API key and returns the new key


### Parameters
This endpoint does not need any parameter.

### Return type

void (empty response body)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SubscriptionPurchaseAllowed**
> SubscriptionPurchaseAllowed()

Check if subscription purchase allowed.

Check whether the user can purchase a subscription.


### Parameters
This endpoint does not need any parameter.

### Return type

void (empty response body)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateProfileData**
> UpdateProfileData($profileData)

Add profile information.

Add profile information to this users account


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **profileData** | **string**| Profile field information for the account | [optional] 

### Return type

void (empty response body)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


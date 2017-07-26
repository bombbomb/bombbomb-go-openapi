# \AccountsApi

All URIs are relative to *https://api.bombbomb.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AccountDetails**](AccountsApi.md#AccountDetails) | **Get** /accounts | Get account details.
[**CreateAccount**](AccountsApi.md#CreateAccount) | **Post** /accounts | Create Account
[**SubscriptionPurchaseAllowed**](AccountsApi.md#SubscriptionPurchaseAllowed) | **Get** /accounts/purchaseable | Check if subscription purchase allowed.


# **AccountDetails**
> AccountDetails($email, $pw, $apiKey)

Get account details.

Get the details of the user's account.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string**| Your login email address | [optional] 
 **pw** | **string**| Your password | [optional] 
 **apiKey** | **string**| Your Api Key | [optional] 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateAccount**
> string CreateAccount($teamId, $firstName, $lastName, $emailAddress, $companyName, $phone, $country, $industry, $address, $city, $postalCode)

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

### Return type

**string**

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SubscriptionPurchaseAllowed**
> SubscriptionPurchaseAllowed($email, $pw, $apiKey)

Check if subscription purchase allowed.

Check whether the user can purchase a subscription.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string**| Your login email address | [optional] 
 **pw** | **string**| Your password | [optional] 
 **apiKey** | **string**| Your Api Key | [optional] 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


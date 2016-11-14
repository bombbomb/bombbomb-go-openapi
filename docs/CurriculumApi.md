# \CurriculumApi

All URIs are relative to *https://api.bombbomb.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCurricula**](CurriculumApi.md#GetCurricula) | **Get** /curricula/ | Get Curricula
[**GetUserCurriculumWithProgress**](CurriculumApi.md#GetUserCurriculumWithProgress) | **Get** /curriculum/getForUserWithProgress | Get Detailed For User


# **GetCurricula**
> []Curriculum GetCurricula($includeProgress)

Get Curricula

Get Curricula, optionally with progress included.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **includeProgress** | **bool**| Whether to return progress with the curriculum. | [optional] 

### Return type

[**[]Curriculum**](Curriculum.md)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUserCurriculumWithProgress**
> []CurriculumWithProgress GetUserCurriculumWithProgress()

Get Detailed For User

Get all curricula for user including progress for each curriculum.


### Parameters
This endpoint does not need any parameter.

### Return type

[**[]CurriculumWithProgress**](CurriculumWithProgress.md)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


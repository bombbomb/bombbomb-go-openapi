# \VideosApi

All URIs are relative to *https://api.bombbomb.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetVideoEncodingStatus**](VideosApi.md#GetVideoEncodingStatus) | **Get** /videos/{videoId}/status | Video Encoding Status
[**GetVideoRecorder**](VideosApi.md#GetVideoRecorder) | **Get** /videos/live/getRecorder | Get Live Video Recorder HTML
[**MarkLiveRecordingComplete**](VideosApi.md#MarkLiveRecordingComplete) | **Post** /videos/live/markComplete | Completes a live recording
[**SignUpload**](VideosApi.md#SignUpload) | **Post** /video/signedUpload | Generate Signed Url
[**UpdateVideoThumbnailV2**](VideosApi.md#UpdateVideoThumbnailV2) | **Put** /videos/thumbnail | Upload thumbnail


# **GetVideoEncodingStatus**
> VideoEncodingStatusResponse GetVideoEncodingStatus(ctx, videoId)
Video Encoding Status

Get information about the current state of encoding for a given video id.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **videoId** | **string**| The video&#39;s id. | 

### Return type

[**VideoEncodingStatusResponse**](VideoEncodingStatusResponse.md)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVideoRecorder**
> VideoRecorderMethodResponse GetVideoRecorder(ctx, optional)
Get Live Video Recorder HTML

Returns an object with a number of properties to help you put a video recorder on your site.         This is to be used in conjunction with the VideoRecordedLive call one the user has confirmed in your UI that         the video is how they want it.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **width** | **int32**| The width of the recorder to present. | 
 **videoId** | **string**| The id of the video to record | 

### Return type

[**VideoRecorderMethodResponse**](VideoRecorderMethodResponse.md)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MarkLiveRecordingComplete**
> VideoPublicRepresentation MarkLiveRecordingComplete(ctx, videoId, filename, title)
Completes a live recording

Used in conjunction with the live recorder method to mark a video recording as complete.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **videoId** | **string**| The id of the video to mark as done. | 
  **filename** | **string**| The filename that was chosen as the final video. | 
  **title** | **string**| The title to give the video | 

### Return type

[**VideoPublicRepresentation**](VideoPublicRepresentation.md)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SignUpload**
> string SignUpload(ctx, policy, optional)
Generate Signed Url

Generates a signed url to be used for video uploads.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **policy** | [**SignUploadRequest**](SignUploadRequest.md)| The policy to sign | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **policy** | [**SignUploadRequest**](SignUploadRequest.md)| The policy to sign | 
 **v4** | **bool**| Whether to do v4 signing | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateVideoThumbnailV2**
> UpdateVideoThumbnailV2(ctx, videoId, thumbnail, optional)
Upload thumbnail

Upload a new video thumbnail

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **videoId** | **string**| The id of the video | 
  **thumbnail** | **string**| The thumbnail being uploaded | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **videoId** | **string**| The id of the video | 
 **thumbnail** | **string**| The thumbnail being uploaded | 
 **custom** | **bool**| The default email to use. | 

### Return type

 (empty response body)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


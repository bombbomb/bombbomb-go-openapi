# \ContactsApi

All URIs are relative to *https://api.bombbomb.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddContactsCSV**](ContactsApi.md#AddContactsCSV) | **Post** /contacts/import_csv | Add contacts from a CSV file.
[**AddNewContact**](ContactsApi.md#AddNewContact) | **Post** /contacts/ | Add a contact.
[**AddNewCustomField**](ContactsApi.md#AddNewCustomField) | **Post** /contacts/custom_fields/ | Add custom fields.
[**AddPastedContacts**](ContactsApi.md#AddPastedContacts) | **Post** /contacts/paste | Add pasted contacts.
[**CSVToObject**](ContactsApi.md#CSVToObject) | **Post** /csv-to-object | Format CSV.
[**DeleteContacts**](ContactsApi.md#DeleteContacts) | **Put** /contacts/delete | Delete Contacts
[**GetContactById**](ContactsApi.md#GetContactById) | **Get** /contact/{id} | Get Contact Details
[**GetCustomFields**](ContactsApi.md#GetCustomFields) | **Get** /contacts/custom_fields/ | Get custom fields.


# **AddContactsCSV**
> AddContactsCSV(ctx, mappingData, listData)
Add contacts from a CSV file.

Add multiple contacts through the upload importer from a CSV file.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **mappingData** | **string**| The info sent for the contacts | 
  **listData** | **string**| The info sent with the import for the list | 

### Return type

 (empty response body)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddNewContact**
> AddNewContact(ctx, contactEmail, optional)
Add a contact.

Add a contact to the users list.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **contactEmail** | **string**| Email of the new contact we are adding | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contactEmail** | **string**| Email of the new contact we are adding | 
 **contactInfo** | **string**| The info sent for this contact | 

### Return type

 (empty response body)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddNewCustomField**
> AddNewCustomField(ctx, fieldName, optional)
Add custom fields.

Add a new custom field.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **fieldName** | **string**| Custom field name to be added | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fieldName** | **string**| Custom field name to be added | 
 **fieldType** | **string**| Custom field type for the field to be added | 

### Return type

 (empty response body)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddPastedContacts**
> AddPastedContacts(ctx, contactEmails, optional)
Add pasted contacts.

Add the pasted contacts to the users list.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **contactEmails** | **string**| Emails array of the new contacts we are adding | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contactEmails** | **string**| Emails array of the new contacts we are adding | 
 **listInfo** | **string**| Information about the lists id, recalculations on totals, consent etc | 

### Return type

 (empty response body)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CSVToObject**
> CSVToObject(ctx, file)
Format CSV.

Format a CSV file to an object.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **file** | **string**| The CSV file being uploaded | 

### Return type

 (empty response body)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteContacts**
> DeleteContacts(ctx, optional)
Delete Contacts

Delete all contacts within a list, or provide a comma separated list of contactIds to delete.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **listId** | **string**| The list of contacts to be deleted. | 
 **contactIds** | **string**| comma separated list of contact ids to delete | 

### Return type

 (empty response body)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetContactById**
> GetContactById(ctx, id)
Get Contact Details

Get the contact details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **string**| Guid for the contact. | 

### Return type

 (empty response body)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCustomFields**
> GetCustomFields(ctx, )
Get custom fields.

Get the current users custom fields.

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


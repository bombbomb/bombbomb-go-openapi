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
> AddContactsCSV($mappingData, $listData)

Add contacts from a CSV file.

Add multiple contacts through the upload importer from a CSV file.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mappingData** | **string**| The info sent for the contacts | 
 **listData** | **string**| The info sent with the import for the list | 

### Return type

void (empty response body)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddNewContact**
> AddNewContact($contactEmail, $contactInfo)

Add a contact.

Add a contact to the users list.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contactEmail** | **string**| Email of the new contact we are adding | 
 **contactInfo** | **string**| The info sent for this contact | [optional] 

### Return type

void (empty response body)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddNewCustomField**
> AddNewCustomField($fieldName, $fieldType)

Add custom fields.

Add a new custom field.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fieldName** | **string**| Custom field name to be added | 
 **fieldType** | **string**| Custom field type for the field to be added | [optional] 

### Return type

void (empty response body)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddPastedContacts**
> AddPastedContacts($contactEmails, $listInfo)

Add pasted contacts.

Add the pasted contacts to the users list.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contactEmails** | **string**| Emails array of the new contacts we are adding | 
 **listInfo** | **string**| Information about the lists id, recalculations on totals, consent etc | [optional] 

### Return type

void (empty response body)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CSVToObject**
> CSVToObject($file)

Format CSV.

Format a CSV file to an object.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | **string**| The CSV file being uploaded | 

### Return type

void (empty response body)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteContacts**
> DeleteContacts($listId, $contactIds)

Delete Contacts

Delete all contacts within a list, or provide a comma separated list of contactIds to delete.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **listId** | **string**| The list of contacts to be deleted. | [optional] 
 **contactIds** | **string**| comma separated list of contact ids to delete | [optional] 

### Return type

void (empty response body)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetContactById**
> GetContactById($id)

Get Contact Details

Get the contact details


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| Guid for the contact. | 

### Return type

void (empty response body)

### Authorization

[BBOAuth2](../README.md#BBOAuth2)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCustomFields**
> GetCustomFields()

Get custom fields.

Get the current users custom fields.


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


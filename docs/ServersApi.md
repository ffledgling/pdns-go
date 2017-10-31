# \ServersApi

All URIs are relative to *http://localhost:8081/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListServer**](ServersApi.md#ListServer) | **Get** /servers/{server_id} | List a server
[**ListServers**](ServersApi.md#ListServers) | **Get** /servers | List all servers


# **ListServer**
> Server ListServer($serverId)

List a server


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serverId** | **string**| The id of the server to retrieve | 

### Return type

[**Server**](Server.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListServers**
> Servers ListServers()

List all servers


### Parameters
This endpoint does not need any parameter.

### Return type

[**Servers**](Servers.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


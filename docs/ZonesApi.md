# \ZonesApi

All URIs are relative to *http://localhost:8081/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListZone**](ZonesApi.md#ListZone) | **Get** /servers/{server_id}/zones/{zone_id} | zone managed by a server
[**ListZones**](ZonesApi.md#ListZones) | **Get** /servers/{server_id}/zones | List all Zones in a server


# **ListZone**
> Zone ListZone($serverId, $zoneId)

zone managed by a server


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serverId** | **string**| The id of the server to retrieve | 
 **zoneId** | **string**| The id of the zone to retrieve | 

### Return type

[**Zone**](Zone.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListZones**
> Zones ListZones($serverId)

List all Zones in a server


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serverId** | **string**| The id of the server to retrieve | 

### Return type

[**Zones**](Zones.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


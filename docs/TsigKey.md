# TsigKey

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the key | [optional] [default to null]
**Id** | **string** | The ID for this key, used in the TSIGkey URL endpoint. | [optional] [default to null]
**Algorithm** | **string** | The algorithm of the TSIG key | [optional] [default to null]
**Key** | **string** | The Base64 encoded secret key, empty when listing keys. MAY be empty when POSTing to have the server generate the key material | [optional] [default to null]
**Type_** | **string** | Set to \&quot;TSIGKey\&quot; | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# PaymentRefund

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique ID for this refund, generated by Square. | [default to null]
**Status** | **string** | The refund&#x27;s status: - &#x60;PENDING&#x60; - Awaiting approval. - &#x60;COMPLETED&#x60; - Successfully completed. - &#x60;REJECTED&#x60; - The refund was rejected. - &#x60;FAILED&#x60; - An error occurred. | [optional] [default to null]
**LocationId** | **string** | The location ID associated with the payment this refund is attached to. | [optional] [default to null]
**AmountMoney** | [***Money**](Money.md) |  | [default to null]
**AppFeeMoney** | [***Money**](Money.md) |  | [optional] [default to null]
**ProcessingFee** | [**[]ProcessingFee**](ProcessingFee.md) | Processing fees and fee adjustments assessed by Square for this refund. | [optional] [default to null]
**PaymentId** | **string** | The ID of the payment associated with this refund. | [optional] [default to null]
**OrderId** | **string** | The ID of the order associated with the refund. | [optional] [default to null]
**Reason** | **string** | The reason for the refund. | [optional] [default to null]
**CreatedAt** | **string** | The timestamp of when the refund was created, in RFC 3339 format. | [optional] [default to null]
**UpdatedAt** | **string** | The timestamp of when the refund was last updated, in RFC 3339 format. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

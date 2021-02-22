# Payment

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A unique ID for the payment. | [optional] [default to null]
**CreatedAt** | **string** | The timestamp of when the payment was created, in RFC 3339 format. | [optional] [default to null]
**UpdatedAt** | **string** | The timestamp of when the payment was last updated, in RFC 3339 format. | [optional] [default to null]
**AmountMoney** | [***Money**](Money.md) |  | [optional] [default to null]
**TipMoney** | [***Money**](Money.md) |  | [optional] [default to null]
**TotalMoney** | [***Money**](Money.md) |  | [optional] [default to null]
**AppFeeMoney** | [***Money**](Money.md) |  | [optional] [default to null]
**ProcessingFee** | [**[]ProcessingFee**](ProcessingFee.md) | The processing fees and fee adjustments assessed by Square for this payment. | [optional] [default to null]
**RefundedMoney** | [***Money**](Money.md) |  | [optional] [default to null]
**Status** | **string** | Indicates whether the payment is APPROVED, COMPLETED, CANCELED, or FAILED. | [optional] [default to null]
**DelayDuration** | **string** | The duration of time after the payment&#x27;s creation when Square automatically applies the &#x60;delay_action&#x60; to the payment. This automatic &#x60;delay_action&#x60; applies only to payments that do not reach a terminal state (COMPLETED, CANCELED, or FAILED) before the &#x60;delay_duration&#x60; time period.  This field is specified as a time duration, in RFC 3339 format.  Notes: This feature is only supported for card payments.  Default:  - Card-present payments: \&quot;PT36H\&quot; (36 hours) from the creation time. - Card-not-present payments: \&quot;P7D\&quot; (7 days) from the creation time. | [optional] [default to null]
**DelayAction** | **string** | The action to be applied to the payment when the &#x60;delay_duration&#x60; has elapsed. This field is read-only.  Current values include &#x60;CANCEL&#x60;. | [optional] [default to null]
**DelayedUntil** | **string** | The read-only timestamp of when the &#x60;delay_action&#x60; is automatically applied, in RFC 3339 format.  Note that this field is calculated by summing the payment&#x27;s &#x60;delay_duration&#x60; and &#x60;created_at&#x60; fields. The &#x60;created_at&#x60; field is generated by Square and might not exactly match the time on your local machine. | [optional] [default to null]
**SourceType** | **string** | The source type for this payment.  Current values include &#x60;CARD&#x60;. | [optional] [default to null]
**CardDetails** | [***CardPaymentDetails**](CardPaymentDetails.md) |  | [optional] [default to null]
**LocationId** | **string** | The ID of the location associated with the payment. | [optional] [default to null]
**OrderId** | **string** | The ID of the order associated with the payment. | [optional] [default to null]
**ReferenceId** | **string** | An optional ID that associates the payment with an entity in another system. | [optional] [default to null]
**CustomerId** | **string** | The &#x60;Customer&#x60; ID of the customer associated with the payment. | [optional] [default to null]
**EmployeeId** | **string** | An optional ID of the employee associated with taking the payment. | [optional] [default to null]
**RefundIds** | **[]string** | A list of &#x60;refund_id&#x60;s identifying refunds for the payment. | [optional] [default to null]
**RiskEvaluation** | [***RiskEvaluation**](RiskEvaluation.md) |  | [optional] [default to null]
**BuyerEmailAddress** | **string** | The buyer&#x27;s email address. | [optional] [default to null]
**BillingAddress** | [***Address**](Address.md) |  | [optional] [default to null]
**ShippingAddress** | [***Address**](Address.md) |  | [optional] [default to null]
**Note** | **string** | An optional note to include when creating a payment. | [optional] [default to null]
**StatementDescriptionIdentifier** | **string** | Additional payment information that gets added to the customer&#x27;s card statement as part of the statement description.  Note that the &#x60;statement_description_identifier&#x60; might get truncated on the statement description to fit the required information including the Square identifier (SQ *) and the name of the seller taking the payment. | [optional] [default to null]
**ReceiptNumber** | **string** | The payment&#x27;s receipt number. The field is missing if a payment is canceled. | [optional] [default to null]
**ReceiptUrl** | **string** | The URL for the payment&#x27;s receipt. The field is only populated for COMPLETED payments. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

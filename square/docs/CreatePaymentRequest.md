# CreatePaymentRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceId** | **string** | The ID for the source of funds for this payment. This can be a nonce generated by the Square payment form or a card on file made with the Customers API. | [default to null]
**IdempotencyKey** | **string** | A unique string that identifies this &#x60;CreatePayment&#x60; request. Keys can be any valid string but must be unique for every &#x60;CreatePayment&#x60; request.  Max: 45 characters  Note: The number of allowed characters might be less than the stated maximum, if multi-byte characters are used.  For more information, see [Idempotency](https://developer.squareup.com/docs/working-with-apis/idempotency). | [default to null]
**AmountMoney** | [***Money**](Money.md) |  | [default to null]
**TipMoney** | [***Money**](Money.md) |  | [optional] [default to null]
**AppFeeMoney** | [***Money**](Money.md) |  | [optional] [default to null]
**DelayDuration** | **string** | The duration of time after the payment&#x27;s creation when Square automatically cancels the payment. This automatic cancellation applies only to payments that do not reach a terminal state (COMPLETED, CANCELED, or FAILED) before the &#x60;delay_duration&#x60; time period.  This parameter should be specified as a time duration, in RFC 3339 format, with a minimum value of 1 minute.  Note: This feature is only supported for card payments. This parameter can only be set for a delayed capture payment (&#x60;autocomplete&#x3D;false&#x60;).  Default:  - Card-present payments: \&quot;PT36H\&quot; (36 hours) from the creation time. - Card-not-present payments: \&quot;P7D\&quot; (7 days) from the creation time. | [optional] [default to null]
**Autocomplete** | **bool** | If set to &#x60;true&#x60;, this payment will be completed when possible. If set to &#x60;false&#x60;, this payment is held in an approved state until either explicitly completed (captured) or canceled (voided). For more information, see [Delayed capture](https://developer.squareup.com/docs/payments-api/take-payments#delayed-payments).  Default: true | [optional] [default to null]
**OrderId** | **string** | Associates a previously created order with this payment. | [optional] [default to null]
**CustomerId** | **string** | The &#x60;Customer&#x60; ID of the customer associated with the payment.  This is required if the &#x60;source_id&#x60; refers to a card on file created using the Customers API. | [optional] [default to null]
**LocationId** | **string** | The location ID to associate with the payment. If not specified, the default location is used. | [optional] [default to null]
**ReferenceId** | **string** | A user-defined ID to associate with the payment.  You can use this field to associate the payment to an entity in an external system  (for example, you might specify an order ID that is generated by a third-party shopping cart).  Limit 40 characters. | [optional] [default to null]
**VerificationToken** | **string** | An identifying token generated by &#x60;SqPaymentForm.verifyBuyer()&#x60;. Verification tokens encapsulate customer device information and 3-D Secure challenge results to indicate that Square has verified the buyer identity.  For more information, see [SCA Overview](https://developer.squareup.com/docs/sca-overview). | [optional] [default to null]
**AcceptPartialAuthorization** | **bool** | If set to &#x60;true&#x60; and charging a Square Gift Card, a payment might be returned with &#x60;amount_money&#x60; equal to less than what was requested. For example, a request for $20 when charging a Square Gift Card with a balance of $5 results in an APPROVED payment of $5. You might choose to prompt the buyer for an additional payment to cover the remainder or cancel the Gift Card payment. This field cannot be &#x60;true&#x60; when &#x60;autocomplete &#x3D; true&#x60;.  For more information, see [Partial amount with Square Gift Cards](https://developer.squareup.com/docs/payments-api/take-payments#partial-payment-gift-card).  Default: false | [optional] [default to null]
**BuyerEmailAddress** | **string** | The buyer&#x27;s email address. | [optional] [default to null]
**BillingAddress** | [***Address**](Address.md) |  | [optional] [default to null]
**ShippingAddress** | [***Address**](Address.md) |  | [optional] [default to null]
**Note** | **string** | An optional note to be entered by the developer when creating a payment.  Limit 500 characters. | [optional] [default to null]
**StatementDescriptionIdentifier** | **string** | Optional additional payment information to include on the customer&#x27;s card statement as part of the statement description. This can be, for example, an invoice number, ticket number, or short description that uniquely identifies the purchase.  Note that the &#x60;statement_description_identifier&#x60; might get truncated on the statement description to fit the required information including the Square identifier (SQ *) and name of the seller taking the payment. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

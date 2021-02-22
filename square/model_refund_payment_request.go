/*
 * Square Connect API
 *
 * Client library for accessing the Square Connect APIs
 *
 * API version: 2.0
 * Contact: developers@squareup.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package square

// Refunds a payment.
type RefundPaymentRequest struct {
	//  A unique string that identifies this `RefundPayment` request. The key can be any valid string but must be unique for every `RefundPayment` request.  For more information, see [Idempotency](https://developer.squareup.com/docs/working-with-apis/idempotency).
	IdempotencyKey string `json:"idempotency_key"`
	AmountMoney *Money `json:"amount_money"`
	AppFeeMoney *Money `json:"app_fee_money,omitempty"`
	// The unique ID of the payment being refunded.
	PaymentId string `json:"payment_id"`
	// A description of the reason for the refund.
	Reason string `json:"reason,omitempty"`
}
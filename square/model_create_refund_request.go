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

// Defines the body parameters that can be included in a request to the [CreateRefund](#endpoint-createrefund) endpoint.  Deprecated - recommend using [RefundPayment](#endpoint-refunds-refundpayment)
type CreateRefundRequest struct {
	// A value you specify that uniquely identifies this refund among refunds you've created for the tender.  If you're unsure whether a particular refund succeeded, you can reattempt it with the same idempotency key without worrying about duplicating the refund.  See [Idempotency keys](#idempotencykeys) for more information.
	IdempotencyKey string `json:"idempotency_key"`
	// The ID of the tender to refund.  A ``Transaction`` has one or more `tenders` (i.e., methods of payment) associated with it, and you refund each tender separately with the Connect API.
	TenderId string `json:"tender_id"`
	// A description of the reason for the refund.  Default value: `Refund via API`
	Reason string `json:"reason,omitempty"`
	AmountMoney *Money `json:"amount_money"`
}
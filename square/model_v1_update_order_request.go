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

// V1UpdateOrderRequest
type V1UpdateOrderRequest struct {
	// The action to perform on the order (COMPLETE, CANCEL, or REFUND). See [V1UpdateOrderRequestAction](#type-v1updateorderrequestaction) for possible values
	Action string `json:"action"`
	// The tracking number of the shipment associated with the order. Only valid if action is COMPLETE.
	ShippedTrackingNumber string `json:"shipped_tracking_number,omitempty"`
	// A merchant-specified note about the completion of the order. Only valid if action is COMPLETE.
	CompletedNote string `json:"completed_note,omitempty"`
	// A merchant-specified note about the refunding of the order. Only valid if action is REFUND.
	RefundedNote string `json:"refunded_note,omitempty"`
	// A merchant-specified note about the canceling of the order. Only valid if action is CANCEL.
	CanceledNote string `json:"canceled_note,omitempty"`
}
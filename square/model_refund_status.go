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
// RefundStatus : Indicates a refund's current status.
type RefundStatus string

// List of RefundStatus
const (
	PENDING_RefundStatus RefundStatus = "PENDING"
	APPROVED_RefundStatus RefundStatus = "APPROVED"
	REJECTED_RefundStatus RefundStatus = "REJECTED"
	FAILED_RefundStatus RefundStatus = "FAILED"
)

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
// DisputeState : The list of possible dispute states.
type DisputeState string

// List of DisputeState
const (
	UNKNOWN_STATE_DisputeState DisputeState = "UNKNOWN_STATE"
	INQUIRY_EVIDENCE_REQUIRED_DisputeState DisputeState = "INQUIRY_EVIDENCE_REQUIRED"
	INQUIRY_PROCESSING_DisputeState DisputeState = "INQUIRY_PROCESSING"
	INQUIRY_CLOSED_DisputeState DisputeState = "INQUIRY_CLOSED"
	EVIDENCE_REQUIRED_DisputeState DisputeState = "EVIDENCE_REQUIRED"
	PROCESSING_DisputeState DisputeState = "PROCESSING"
	WON_DisputeState DisputeState = "WON"
	LOST_DisputeState DisputeState = "LOST"
	ACCEPTED_DisputeState DisputeState = "ACCEPTED"
	WAITING_THIRD_PARTY_DisputeState DisputeState = "WAITING_THIRD_PARTY"
)
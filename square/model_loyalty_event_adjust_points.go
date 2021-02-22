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

// Provides metadata when the event `type` is `ADJUST_POINTS`.
type LoyaltyEventAdjustPoints struct {
	// The Square-assigned ID of the `loyalty program`.
	LoyaltyProgramId string `json:"loyalty_program_id,omitempty"`
	// The number of points added or removed.
	Points int32 `json:"points"`
	// The reason for the adjustment of points.
	Reason string `json:"reason,omitempty"`
}

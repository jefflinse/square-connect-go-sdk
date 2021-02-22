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

// A request to accumulate points for a purchase.
type AccumulateLoyaltyPointsRequest struct {
	AccumulatePoints *LoyaltyEventAccumulatePoints `json:"accumulate_points"`
	// A unique string that identifies the `AccumulateLoyaltyPoints` request.  Keys can be any valid string but must be unique for every request.
	IdempotencyKey string `json:"idempotency_key"`
	// The `location` where the purchase was made.
	LocationId string `json:"location_id"`
}
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

// A response that includes the loyalty rewards satisfying the search criteria.
type SearchLoyaltyRewardsResponse struct {
	// Any errors that occurred during the request.
	Errors []ModelError `json:"errors,omitempty"`
	// The loyalty rewards that satisfy the search criteria. These are returned in descending order by `updated_at`.
	Rewards []LoyaltyReward `json:"rewards,omitempty"`
	// The pagination cursor to be used in a subsequent  request. If empty, this is the final response.
	Cursor string `json:"cursor,omitempty"`
}

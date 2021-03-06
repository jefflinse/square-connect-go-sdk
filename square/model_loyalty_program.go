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

type LoyaltyProgram struct {
	// The Square-assigned ID of the loyalty program. Updates to  the loyalty program do not modify the identifier.
	Id string `json:"id"`
	// Whether the program is currently active. See [LoyaltyProgramStatus](#type-loyaltyprogramstatus) for possible values
	Status string `json:"status"`
	// The list of rewards for buyers, sorted by ascending points.
	RewardTiers []LoyaltyProgramRewardTier `json:"reward_tiers"`
	ExpirationPolicy *LoyaltyProgramExpirationPolicy `json:"expiration_policy,omitempty"`
	Terminology *LoyaltyProgramTerminology `json:"terminology"`
	// The `locations` at which the program is active.
	LocationIds []string `json:"location_ids"`
	// The timestamp when the program was created, in RFC 3339 format.
	CreatedAt string `json:"created_at"`
	// The timestamp when the reward was last updated, in RFC 3339 format.
	UpdatedAt string `json:"updated_at"`
	// Defines how buyers can earn loyalty points.
	AccrualRules []LoyaltyProgramAccrualRule `json:"accrual_rules"`
}

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
// LoyaltyProgramRewardDefinitionScope : Indicates the scope of the reward tier. DEPRECATED at version 2020-12-16. Discount details are now defined using a catalog pricing rule and other catalog objects. For more information, see [Get discount details for the reward](https://developer.squareup.com/docs/docs/loyalty-api/overview#get-discount-details).
type LoyaltyProgramRewardDefinitionScope string

// List of LoyaltyProgramRewardDefinitionScope
const (
	ORDER_LoyaltyProgramRewardDefinitionScope LoyaltyProgramRewardDefinitionScope = "ORDER"
	ITEM_VARIATION_LoyaltyProgramRewardDefinitionScope LoyaltyProgramRewardDefinitionScope = "ITEM_VARIATION"
	CATEGORY_LoyaltyProgramRewardDefinitionScope LoyaltyProgramRewardDefinitionScope = "CATEGORY"
)

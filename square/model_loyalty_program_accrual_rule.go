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

// Defines an accrual rule, which is how buyers can earn points.
type LoyaltyProgramAccrualRule struct {
	// The type of the accrual rule that defines how buyers can earn points. See [LoyaltyProgramAccrualRuleType](#type-loyaltyprogramaccrualruletype) for possible values
	AccrualType string `json:"accrual_type"`
	// The number of points that  buyers earn based on the `accrual_type`.
	Points int32 `json:"points,omitempty"`
	VisitMinimumAmountMoney *Money `json:"visit_minimum_amount_money,omitempty"`
	SpendAmountMoney *Money `json:"spend_amount_money,omitempty"`
	// The ID of the `catalog object` to purchase to earn the number of points defined by the rule. This is either an item variation or a category, depending on the type. This is defined on `ITEM_VARIATION` rules and `CATEGORY` rules.
	CatalogObjectId string `json:"catalog_object_id,omitempty"`
}

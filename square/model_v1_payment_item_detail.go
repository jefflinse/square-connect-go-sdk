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

// V1PaymentItemDetail
type V1PaymentItemDetail struct {
	// The name of the item's merchant-defined category, if any.
	CategoryName string `json:"category_name,omitempty"`
	//  The item's merchant-defined SKU, if any.
	Sku string `json:"sku,omitempty"`
	// The unique ID of the item purchased, if any.
	ItemId string `json:"item_id,omitempty"`
	// The unique ID of the item variation purchased, if any.
	ItemVariationId string `json:"item_variation_id,omitempty"`
}
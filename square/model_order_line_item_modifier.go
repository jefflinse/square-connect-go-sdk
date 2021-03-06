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

// A [CatalogModifier](#type-catalogmodifier).
type OrderLineItemModifier struct {
	// Unique ID that identifies the modifier only within this order.
	Uid string `json:"uid,omitempty"`
	// The catalog object id referencing `CatalogModifier`.
	CatalogObjectId string `json:"catalog_object_id,omitempty"`
	// The name of the item modifier.
	Name string `json:"name,omitempty"`
	BasePriceMoney *Money `json:"base_price_money,omitempty"`
	TotalPriceMoney *Money `json:"total_price_money,omitempty"`
}

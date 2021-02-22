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
// OrderLineItemTaxType : Indicates how the tax is applied to the associated line item or order.
type OrderLineItemTaxType string

// List of OrderLineItemTaxType
const (
	UNKNOWN_TAX_OrderLineItemTaxType OrderLineItemTaxType = "UNKNOWN_TAX"
	ADDITIVE_OrderLineItemTaxType OrderLineItemTaxType = "ADDITIVE"
	INCLUSIVE_OrderLineItemTaxType OrderLineItemTaxType = "INCLUSIVE"
)
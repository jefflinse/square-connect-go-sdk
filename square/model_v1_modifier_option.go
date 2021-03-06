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

// V1ModifierOption
type V1ModifierOption struct {
	// The modifier option's unique ID.
	Id string `json:"id,omitempty"`
	// The modifier option's name.
	Name string `json:"name,omitempty"`
	PriceMoney *V1Money `json:"price_money,omitempty"`
	// If true, the modifier option is the default option in a modifier list for which selection_type is SINGLE.
	OnByDefault bool `json:"on_by_default,omitempty"`
	// Indicates the modifier option's list position when displayed in Square Point of Sale and the merchant dashboard. If more than one modifier option in the same modifier list has the same ordinal value, those options are displayed in alphabetical order.
	Ordinal int32 `json:"ordinal,omitempty"`
	// The ID of the modifier list the option belongs to.
	ModifierListId string `json:"modifier_list_id,omitempty"`
	// The ID of the CatalogObject in the Connect v2 API. Objects that are shared across multiple locations share the same v2 ID.
	V2Id string `json:"v2_id,omitempty"`
}

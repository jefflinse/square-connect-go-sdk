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

// V1PaymentModifier
type V1PaymentModifier struct {
	// The modifier option's name.
	Name string `json:"name,omitempty"`
	AppliedMoney *V1Money `json:"applied_money,omitempty"`
	// TThe ID of the applied modifier option, if available. Modifier options applied in older versions of Square Register might not have an ID.
	ModifierOptionId string `json:"modifier_option_id,omitempty"`
}
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

type PaymentOptions struct {
	// Indicates whether the Payment objects created from this `TerminalCheckout` will automatically be `COMPLETED` or left in an `APPROVED` state for later modification.
	Autocomplete bool `json:"autocomplete,omitempty"`
}

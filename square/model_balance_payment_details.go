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

// Reflects the current status of a balance payment.
type BalancePaymentDetails struct {
	// The ID of the account used to fund the payment.
	AccountId string `json:"account_id,omitempty"`
	// The balance payment’s current state. The state can be COMPLETED or FAILED.
	Status string `json:"status,omitempty"`
}

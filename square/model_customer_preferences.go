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

// Represents communication preferences for the customer profile.
type CustomerPreferences struct {
	// The customer has unsubscribed from receiving marketing campaign emails.
	EmailUnsubscribed bool `json:"email_unsubscribed,omitempty"`
}

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

// Defines the fields that are included in the response from the [CreateSubscription](#endpoint-subscriptions-createsubscription) endpoint.
type CreateSubscriptionResponse struct {
	// Information about errors encountered during the request.
	Errors []ModelError `json:"errors,omitempty"`
	Subscription *Subscription `json:"subscription,omitempty"`
}
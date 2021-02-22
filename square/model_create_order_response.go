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

// Defines the fields that are included in the response body of a request to the CreateOrder endpoint.  One of `errors` or `order` is present in a given response (never both).
type CreateOrderResponse struct {
	Order *Order `json:"order,omitempty"`
	// Any errors that occurred during the request.
	Errors []ModelError `json:"errors,omitempty"`
}

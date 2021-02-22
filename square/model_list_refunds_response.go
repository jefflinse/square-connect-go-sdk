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

// Defines the fields that are included in the response body of a request to the [ListRefunds](#endpoint-listrefunds) endpoint.  One of `errors` or `refunds` is present in a given response (never both).
type ListRefundsResponse struct {
	// Any errors that occurred during the request.
	Errors []ModelError `json:"errors,omitempty"`
	// An array of refunds that match your query.
	Refunds []Refund `json:"refunds,omitempty"`
	// A pagination cursor for retrieving the next set of results, if any remain. Provide this value as the `cursor` parameter in a subsequent request to this endpoint.  See [Paginating results](#paginatingresults) for more information.
	Cursor string `json:"cursor,omitempty"`
}

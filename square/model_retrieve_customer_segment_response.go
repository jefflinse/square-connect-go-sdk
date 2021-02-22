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

// Defines the fields included in the response body for requests to __RetrieveCustomerSegment__.  One of `errors` or `segment` is present in a given response (never both).
type RetrieveCustomerSegmentResponse struct {
	// Any errors that occurred during the request.
	Errors []ModelError `json:"errors,omitempty"`
	Segment *CustomerSegment `json:"segment,omitempty"`
}

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

// The response to a request for a set of `BreakTypes`. Contains the requested `BreakType` objects. May contain a set of `Error` objects if the request resulted in errors.
type ListBreakTypesResponse struct {
	//  A page of `BreakType` results.
	BreakTypes []BreakType `json:"break_types,omitempty"`
	// Value supplied in the subsequent request to fetch the next next page of Break Type results.
	Cursor string `json:"cursor,omitempty"`
	// Any errors that occurred during the request.
	Errors []ModelError `json:"errors,omitempty"`
}

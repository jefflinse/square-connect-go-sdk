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

type ListCashDrawerShiftsResponse struct {
	// A collection of CashDrawerShiftSummary objects for shifts that match the query.
	Items []CashDrawerShiftSummary `json:"items,omitempty"`
	// Opaque cursor for fetching the next page of results. Cursor is not present in the last page of results.
	Cursor string `json:"cursor,omitempty"`
	// Any errors that occurred during the request.
	Errors []ModelError `json:"errors,omitempty"`
}

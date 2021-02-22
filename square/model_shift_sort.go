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

// Sets the sort order of search results.
type ShiftSort struct {
	// The field to sort on. See [ShiftSortField](#type-shiftsortfield) for possible values
	Field string `json:"field,omitempty"`
	// The order in which results are returned. Defaults to DESC. See [SortOrder](#type-sortorder) for possible values
	Order string `json:"order,omitempty"`
}

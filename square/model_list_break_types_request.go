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

// A request for a filtered set of `BreakType` objects
type ListBreakTypesRequest struct {
	// Filter Break Types returned to only those that are associated with the specified location.
	LocationId string `json:"location_id,omitempty"`
	// Maximum number of Break Types to return per page. Can range between 1 and 200. The default is the maximum at 200.
	Limit int32 `json:"limit,omitempty"`
	// Pointer to the next page of Break Type results to fetch.
	Cursor string `json:"cursor,omitempty"`
}
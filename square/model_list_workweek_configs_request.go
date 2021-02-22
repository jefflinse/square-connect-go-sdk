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

// A request for a set of `WorkweekConfig` objects
type ListWorkweekConfigsRequest struct {
	// Maximum number of Workweek Configs to return per page.
	Limit int32 `json:"limit,omitempty"`
	// Pointer to the next page of Workweek Config results to fetch.
	Cursor string `json:"cursor,omitempty"`
}

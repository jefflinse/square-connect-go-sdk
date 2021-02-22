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

type RetrieveInventoryChangesRequest struct {
	// The `Location` IDs to look up as a comma-separated list. An empty list queries all locations.
	LocationIds string `json:"location_ids,omitempty"`
	// A pagination cursor returned by a previous call to this endpoint. Provide this to retrieve the next set of results for the original query.  See the [Pagination](https://developer.squareup.com/docs/working-with-apis/pagination) guide for more information.
	Cursor string `json:"cursor,omitempty"`
}

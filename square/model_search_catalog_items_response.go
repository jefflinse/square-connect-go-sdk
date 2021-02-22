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

// Defines the response body returned from the [SearchCatalogItems](#endpoint-Catalog-SearchCatalogItems) endpoint.
type SearchCatalogItemsResponse struct {
	// Any errors that occurred during the request.
	Errors []ModelError `json:"errors,omitempty"`
	// Returned items matching the specified query expressions.
	Items []CatalogObject `json:"items,omitempty"`
	// Pagination token used in the next request to return more of the search result.
	Cursor string `json:"cursor,omitempty"`
	// Ids of returned item variations matching the specified query expression.
	MatchedVariationIds []string `json:"matched_variation_ids,omitempty"`
}

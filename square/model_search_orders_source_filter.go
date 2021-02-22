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

// Filter based on order `source` information.
type SearchOrdersSourceFilter struct {
	// Filters by `Source` `name`. Will return any orders with with a `source.name` that matches any of the listed source names.  Max: 10 source names.
	SourceNames []string `json:"source_names,omitempty"`
}
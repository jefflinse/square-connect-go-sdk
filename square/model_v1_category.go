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

// V1Category
type V1Category struct {
	// The category's unique ID.
	Id string `json:"id,omitempty"`
	// The category's name.
	Name string `json:"name,omitempty"`
	// The ID of the CatalogObject in the Connect v2 API. Objects that are shared across multiple locations share the same v2 ID.
	V2Id string `json:"v2_id,omitempty"`
}

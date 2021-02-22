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

// A named selection for this `SELECTION`-type custom attribute definition.
type CatalogCustomAttributeDefinitionSelectionConfigCustomAttributeSelection struct {
	// Unique ID set by Square.
	Uid string `json:"uid,omitempty"`
	// Selection name, unique within `allowed_selections`.
	Name string `json:"name"`
}
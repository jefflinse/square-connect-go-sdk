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

// V1UpdateModifierListRequest
type V1UpdateModifierListRequest struct {
	// The modifier list's name.
	Name string `json:"name,omitempty"`
	// Indicates whether multiple options from the modifier list can be applied to a single item. See [V1UpdateModifierListRequestSelectionType](#type-v1updatemodifierlistrequestselectiontype) for possible values
	SelectionType string `json:"selection_type,omitempty"`
}

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

// Represents a query (including filtering criteria, sorting criteria, or both) used to search for customer profiles.
type CustomerQuery struct {
	Filter *CustomerFilter `json:"filter,omitempty"`
	Sort *CustomerSort `json:"sort,omitempty"`
}

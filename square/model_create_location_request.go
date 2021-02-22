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

// Request object for the [CreateLocation](#endpoint-createlocation) endpoint.
type CreateLocationRequest struct {
	Location *Location `json:"location,omitempty"`
}
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

// Defines the fields that are included in the response body of a request to the [RegisterDomain](#endpoint-registerdomain) endpoint.  Either `errors` or `status` will be present in a given response (never both).
type RegisterDomainResponse struct {
	// Any errors that occurred during the request.
	Errors []ModelError `json:"errors,omitempty"`
	// Status of the domain registration.  See `RegisterDomainResponseStatus` for possible values. See [RegisterDomainResponseStatus](#type-registerdomainresponsestatus) for possible values
	Status string `json:"status,omitempty"`
}

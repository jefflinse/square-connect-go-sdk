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

// Defines fields in a `RetrieveDispute` response.
type RetrieveDisputeResponse struct {
	// Information about errors encountered during the request.
	Errors []ModelError `json:"errors,omitempty"`
	Dispute *Dispute `json:"dispute,omitempty"`
}

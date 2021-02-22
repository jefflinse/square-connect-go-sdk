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
// RegisterDomainResponseStatus : The status of domain registration.
type RegisterDomainResponseStatus string

// List of RegisterDomainResponseStatus
const (
	PENDING_RegisterDomainResponseStatus RegisterDomainResponseStatus = "PENDING"
	VERIFIED_RegisterDomainResponseStatus RegisterDomainResponseStatus = "VERIFIED"
)
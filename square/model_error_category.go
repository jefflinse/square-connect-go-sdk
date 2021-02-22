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
// ErrorCategory : Indicates which high-level category of error has occurred during a request to the Connect API.
type ErrorCategory string

// List of ErrorCategory
const (
	API_ERROR_ErrorCategory ErrorCategory = "API_ERROR"
	AUTHENTICATION_ERROR_ErrorCategory ErrorCategory = "AUTHENTICATION_ERROR"
	INVALID_REQUEST_ERROR_ErrorCategory ErrorCategory = "INVALID_REQUEST_ERROR"
	RATE_LIMIT_ERROR_ErrorCategory ErrorCategory = "RATE_LIMIT_ERROR"
	PAYMENT_METHOD_ERROR_ErrorCategory ErrorCategory = "PAYMENT_METHOD_ERROR"
	REFUND_ERROR_ErrorCategory ErrorCategory = "REFUND_ERROR"
)

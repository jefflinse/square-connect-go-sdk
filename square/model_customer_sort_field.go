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
// CustomerSortField : Specifies customer attributes as the sort key to customer profiles returned from a search.
type CustomerSortField string

// List of CustomerSortField
const (
	DEFAULT__CustomerSortField CustomerSortField = "DEFAULT"
	CREATED_AT_CustomerSortField CustomerSortField = "CREATED_AT"
)

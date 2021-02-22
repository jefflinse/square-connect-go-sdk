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

// Identifies the  sort field and sort order.
type InvoiceSort struct {
	// The field to sort on. See [InvoiceSortField](#type-invoicesortfield) for possible values
	Field string `json:"field"`
	// The order to use for sorting the results. See [SortOrder](#type-sortorder) for possible values
	Order string `json:"order,omitempty"`
}
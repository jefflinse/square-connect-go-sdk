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

// An additional seller-defined and customer-facing field to include on the invoice. For more information,  see [Custom fields](/docs/invoices-api/overview#custom-fields).
type InvoiceCustomField struct {
	// The label or title of the custom field. This field is required for a custom field.
	Label string `json:"label,omitempty"`
	// The text of the custom field. If omitted, only the label is rendered.
	Value string `json:"value,omitempty"`
	// The location of the custom field on the invoice. This field is required for a custom field. See [InvoiceCustomFieldPlacement](#type-invoicecustomfieldplacement) for possible values
	Placement string `json:"placement,omitempty"`
}

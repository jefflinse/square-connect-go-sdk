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

// Stores information about an invoice. You use the Invoices API to create and manage invoices. For more information, see [Manage Invoices Using the Invoices API](/docs/invoices-api/overview).
type Invoice struct {
	// The Square-assigned ID of the invoice.
	Id string `json:"id,omitempty"`
	// The Square-assigned version number, which is incremented each time an update is committed to the invoice.
	Version int32 `json:"version,omitempty"`
	// The ID of the location that this invoice is associated with. This field is required when creating an invoice.
	LocationId string `json:"location_id,omitempty"`
	// The ID of the `order` for which the invoice is created.  This order must be in the `OPEN` state and must belong to the `location_id` specified for this invoice. This field is required when creating an invoice.
	OrderId string `json:"order_id,omitempty"`
	PrimaryRecipient *InvoiceRecipient `json:"primary_recipient,omitempty"`
	// The payment schedule for the invoice, represented by one or more payment requests that define payment settings, such as amount due and due date. You can specify a maximum of 13 payment requests, with up to 12 `INSTALLMENT` request types. For more information, see [Payment requests](https://developer.squareup.com/docs/docs/invoices-api/overview#payment-requests).  This field is required when creating an invoice. It must contain at least one payment request.
	PaymentRequests []InvoicePaymentRequest `json:"payment_requests,omitempty"`
	DeliveryMethod *InvoiceDeliveryMethodInvoiceDeliveryMethod `json:"delivery_method,omitempty"`
	// A user-friendly invoice number. The value is unique within a location. If not provided when creating an invoice, Square assigns a value. It increments from 1 and padded with zeros making it 7 characters long for example, 0000001, 0000002.
	InvoiceNumber string `json:"invoice_number,omitempty"`
	// The title of the invoice.
	Title string `json:"title,omitempty"`
	// The description of the invoice. This is visible to the customer receiving the invoice.
	Description string `json:"description,omitempty"`
	// The timestamp when the invoice is scheduled for processing, in RFC 3339 format. After the invoice is published, Square processes the invoice on the specified date, according to the delivery method and payment request settings.  If the field is not set, Square processes the invoice immediately after it is published.
	ScheduledAt string `json:"scheduled_at,omitempty"`
	// The URL of the Square-hosted invoice page. After you publish the invoice using the `PublishInvoice` endpoint, Square hosts the invoice page and returns the page URL in the response.
	PublicUrl string `json:"public_url,omitempty"`
	NextPaymentAmountMoney *Money `json:"next_payment_amount_money,omitempty"`
	// The status of the invoice. See [InvoiceStatus](#type-invoicestatus) for possible values
	Status string `json:"status,omitempty"`
	// The time zone of the date values (for example, `due_date`) specified in the invoice.
	Timezone string `json:"timezone,omitempty"`
	// The timestamp when the invoice was created, in RFC 3339 format.
	CreatedAt string `json:"created_at,omitempty"`
	// The timestamp when the invoice was last updated, in RFC 3339 format.
	UpdatedAt string `json:"updated_at,omitempty"`
	// Additional seller-defined fields to render on the invoice. These fields are visible to sellers and buyers on the Square-hosted invoice page and in emailed or PDF copies of invoices. For more information, see [Custom fields](https://developer.squareup.com/docs/docs/invoices-api/overview#custom-fields).  Max: 2 custom fields
	CustomFields []InvoiceCustomField `json:"custom_fields,omitempty"`
}
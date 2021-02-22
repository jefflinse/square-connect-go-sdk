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

type CashDrawerShiftEvent struct {
	// The unique ID of the event.
	Id string `json:"id,omitempty"`
	// The ID of the employee that created the event.
	EmployeeId string `json:"employee_id,omitempty"`
	// The type of cash drawer shift event. See [CashDrawerEventType](#type-cashdrawereventtype) for possible values
	EventType string `json:"event_type,omitempty"`
	EventMoney *Money `json:"event_money,omitempty"`
	// The event time in ISO 8601 format.
	CreatedAt string `json:"created_at,omitempty"`
	// An optional description of the event, entered by the employee that created the event.
	Description string `json:"description,omitempty"`
}
